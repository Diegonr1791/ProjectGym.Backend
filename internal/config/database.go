package config

import (
	"fmt"
	"log"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	config := LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos: ", err)
	}

	fmt.Println("✅ Conexión a la base de datos exitosa")
	DB = db

	// Auto-migrar las tablas
	err = db.AutoMigrate(
		&model.GrupoMuscular{},
		&model.TipoEjercicio{},
		&model.Ejercicio{},
		&model.SesionEjercicio{},
		&model.Rutina{},
		&model.Favorita{},
		&model.Medicion{},
		//&model.Usuario{},
		&model.GrupoMuscular{},
		&model.TipoEjercicio{},
		&model.Ejercicio{},
		// Agregar otros modelos aquí según sea necesario
	)
	if err != nil {
		log.Fatal("Error al migrar las tablas: ", err)
	}

	fmt.Println("✅ Migraciones completadas")

	var tables []string
	DB.Raw("SELECT tablename FROM pg_tables WHERE schemaname = 'public'").Scan(&tables)
	fmt.Println("Tablas:", tables)
}
