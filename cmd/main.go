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
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
package main

import (
	"fmt"
	"log"

	_ "github.com/Diegonr1791/GymBro/docs" // Importar docs generados por swag
	config "github.com/Diegonr1791/GymBro/internal/config"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Inicializar contenedor de dependencias
	container := config.NewContainer()

	// Crear y configurar servidor
	server := config.NewServer(container)

	// Iniciar servidor
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("🚀 Servidor iniciando en puerto %s", cfg.ServerPort)
	if err := server.Run(addr); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
