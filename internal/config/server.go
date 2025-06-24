package config

import (
	handler "github.com/Diegonr1791/GymBro/interfaces/http/handler"
	middleware "github.com/Diegonr1791/GymBro/interfaces/http/middleware"
	"github.com/gin-contrib/cors"
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

	// Habilitar CORS para todos los orígenes (útil para pruebas)
	router.Use(cors.Default())

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

	// Crear factory de middlewares
	middlewareFactory := NewMiddlewareFactory(s.container)

	// Grupo principal para la API versionada
	apiV1 := s.router.Group("/api/v1")

	// Rutas públicas (sin autenticación)
	handler.NewAuthHandler(apiV1, s.container.UsuarioService, s.container.RefreshTokenService, s.config)

	// Grupo de rutas protegidas con JWT
	protected := apiV1.Group("/")
	protected.Use(middlewareFactory.CreateJWTAuthMiddleware())

	// Configurar handlers protegidos
	handler.NewRoleHandler(protected, s.container.RoleService)

	// Configurar handler de usuarios con autorización especial para eliminación
	handler.NewUsuarioHandlerWithAuth(protected, s.container.UsuarioService, middlewareFactory.CreateUserDeletionAuthMiddleware())

	// Configurar otros handlers
	s.setupOtherHandlers(protected)
}

// setupOtherHandlers configura los handlers que no requieren autorización especial
func (s *Server) setupOtherHandlers(protected *gin.RouterGroup) {
	handler.NewRoutineHandler(protected, s.container.RutinaService)
	handler.NewGrupoMuscularHandler(protected, s.container.GrupoMuscularService)
	handler.NewRoutineMuscleGroupHandler(protected, s.container.RutinaGMService)
	handler.NewFavoriteHandler(protected, s.container.FavoritaService)
	handler.NewMeasurementHandler(protected, s.container.MedicionService)
	handler.NewTypeExerciseHandler(protected, s.container.TipoEjercicioService)
	handler.NewExerciseHandler(protected, s.container.EjercicioService)
	handler.NewSessionHandler(protected, s.container.SesionService)
	handler.NewSessionExerciseHandler(protected, s.container.SesionEjercicioService)
}

// Run inicia el servidor en el puerto especificado
func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
