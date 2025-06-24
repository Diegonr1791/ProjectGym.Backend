package main

import (
	"log"

	"github.com/Diegonr1791/GymBro/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env")
	}

	// Inicializar contenedor de dependencias
	container := config.NewContainer()

	// Ejecutar seeding
	log.Println("🌱 Iniciando proceso de seeding...")

	if err := container.Seeder.Seed(); err != nil {
		log.Fatalf("❌ Error durante el seeding: %v", err)
	}

	// Crear usuario dev adicional
	if err := container.Seeder.SeedDevUser(); err != nil {
		log.Printf("⚠️  Error creando usuario dev: %v", err)
	}

	log.Println("")
	log.Println("🎉 ¡Seeding completado!")
	log.Println("")
	log.Println("📋 Información:")
	log.Println("   - Los roles del sistema han sido creados")
	log.Println("   - Los usuarios se crean según las variables de entorno configuradas")
	log.Println("")
	log.Println("🔧 Variables de entorno para usuarios:")
	log.Println("   - ADMIN_EMAIL, ADMIN_PASSWORD, ADMIN_NAME")
	log.Println("   - DEV_EMAIL, DEV_PASSWORD, DEV_NAME")
	log.Println("")
	log.Println("⚠️  IMPORTANTE: Cambia las contraseñas después del primer login")
}
