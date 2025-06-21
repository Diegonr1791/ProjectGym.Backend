package config

import (
	http "github.com/Diegonr1791/GymBro/interfaces/http"
	"github.com/Diegonr1791/GymBro/internal/auth"
	"github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server configura y maneja el servidor HTTP
type Server struct {
	router    *gin.Engine
	container *Container
	config    *Config
}

// NewServer crea una nueva instancia del servidor
func NewServer(container *Container, cfg *Config) *Server {
	router := gin.Default()

	server := &Server{
		router:    router,
		container: container,
		config:    cfg,
	}

	// Configurar rutas
	server.setupRoutes()

	return server
}

// setupRoutes configura todas las rutas de la aplicación
func (s *Server) setupRoutes() {
	// Configurar Swagger
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Grupo principal para la API versionada
	apiV1 := s.router.Group("/api/v1")

	// Inicializar el use case de refresh token aquí, donde tenemos la config
	refreshTokenUsecase := usecase.NewRefreshTokenUsecase(s.container.RefreshTokenRepo, s.container.UsuarioRepo, s.config)

	// Rutas públicas (sin autenticación) - ahora cuelgan de /api/v1
	http.NewAuthHandler(apiV1, s.container.UsuarioService, refreshTokenUsecase, s.config)

	// Grupo de rutas protegidas con JWT - ahora cuelgan de /api/v1
	protected := apiV1.Group("/")
	protected.Use(auth.JWTAuthMiddleware(s.config))

	// Configurar handlers protegidos
	http.NewUsuarioHandler(protected, s.container.UsuarioService)
	http.NewRutinaHandler(protected, s.container.RutinaService)
	http.NewGrupoMuscularHandler(protected, s.container.GrupoMuscularService)
	http.NewRutinaGMHandler(protected, s.container.RutinaGMService)
	http.NewFavoritaHandler(protected, s.container.FavoritaService)
	http.NewMedicionHandler(protected, s.container.MedicionService)
	http.NewTypeExerciseHandler(protected, s.container.TipoEjercicioService)
	http.NewExerciseHandler(protected, s.container.EjercicioService)
	http.NewSessionHandler(protected, s.container.SesionService)
	http.NewSessionExerciseHandler(protected, s.container.SesionEjercicioService)
}

// Run inicia el servidor en el puerto especificado
func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
