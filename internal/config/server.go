package config

import (
	http "github.com/Diegonr1791/GymBro/interfaces/http"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server configura y maneja el servidor HTTP
type Server struct {
	router    *gin.Engine
	container *Container
}

// NewServer crea una nueva instancia del servidor
func NewServer(container *Container) *Server {
	router := gin.Default()

	server := &Server{
		router:    router,
		container: container,
	}

	// Configurar rutas
	server.setupRoutes()

	return server
}

// setupRoutes configura todas las rutas de la aplicación
func (s *Server) setupRoutes() {
	// Configurar Swagger
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configurar handlers
	http.NewUsuarioHandler(s.router, s.container.UsuarioService)
	http.NewRutinaHandler(s.router, s.container.RutinaService)
	http.NewGrupoMuscularHandler(s.router, s.container.GrupoMuscularService)
	http.NewRutinaGMHandler(s.router, s.container.RutinaGMService)
	http.NewFavoritaHandler(s.router, s.container.FavoritaService)
	http.NewMedicionHandler(s.router, s.container.MedicionService)
	http.NewTypeExerciseHandler(s.router, s.container.TipoEjercicioService)
	http.NewExerciseHandler(s.router, s.container.EjercicioService)
	http.NewSessionHandler(s.router, s.container.SesionService)
	http.NewSessionExerciseHandler(s.router, s.container.SesionEjercicioService)
}

// Run inicia el servidor en el puerto especificado
func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
