package config

import (
	http "github.com/Diegonr1791/GymBro/interfaces/http"
	"github.com/Diegonr1791/GymBro/interfaces/http/middleware"
	"github.com/Diegonr1791/GymBro/internal/auth"
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

	// Registrar el middleware de errores globalmente.
	router.Use(middleware.ErrorHandler())

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

	// Rutas públicas (sin autenticación) - ahora cuelgan de /api/v1
	http.NewAuthHandler(apiV1, s.container.UsuarioService, s.container.RefreshTokenService, s.config)

	// Grupo de rutas protegidas con JWT - ahora cuelgan de /api/v1
	protected := apiV1.Group("/")
	protected.Use(auth.JWTAuthMiddleware(s.config))

	// Configurar handlers protegidos
	http.NewRoleHandler(protected, s.container.RoleService)
	http.NewUsuarioHandler(protected, s.container.UsuarioService)
	http.NewRoutineHandler(protected, s.container.RutinaService)
	http.NewGrupoMuscularHandler(protected, s.container.GrupoMuscularService)
	http.NewRoutineMuscleGroupHandler(protected, s.container.RutinaGMService)
	http.NewFavoriteHandler(protected, s.container.FavoritaService)
	http.NewMeasurementHandler(protected, s.container.MedicionService)
	http.NewTypeExerciseHandler(protected, s.container.TipoEjercicioService)
	http.NewExerciseHandler(protected, s.container.EjercicioService)
	http.NewSessionHandler(protected, s.container.SesionService)
	http.NewSessionExerciseHandler(protected, s.container.SesionEjercicioService)
}

// Run inicia el servidor en el puerto especificado
func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
