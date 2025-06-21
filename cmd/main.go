// @title           GymBro API
// @version         1.0
// @description     API para gestión de rutinas de gimnasio
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
package main

import (
	"fmt"
	"log"

	_ "github.com/Diegonr1791/GymBro/docs" // Importar docs generados por swag
	"github.com/Diegonr1791/GymBro/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno desde el archivo .env
	// Esto debe hacerse antes de cargar la configuración
	err := godotenv.Load()
	if err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env. Usando valores por defecto y variables de entorno del sistema.")
	}

	// Cargar configuración
	cfg := config.LoadConfig()

	// Inicializar contenedor de dependencias
	container := config.NewContainer()

	// Crear y configurar servidor
	server := config.NewServer(container, cfg)

	// Iniciar servidor
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("🚀 Servidor iniciando en puerto %s", cfg.ServerPort)
	if err := server.Run(addr); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
