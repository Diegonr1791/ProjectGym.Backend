package main

import (
	"fmt"
	"log"

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
