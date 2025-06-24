package config

import (
	"context"
	"log"
	"os"

	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

// Seeder maneja la creación de datos iniciales
type Seeder struct {
	roleRepo repositories.RoleRepository
	userRepo repositories.UsuarioRepository
}

// NewSeeder crea una nueva instancia del seeder
func NewSeeder(roleRepo repositories.RoleRepository, userRepo repositories.UsuarioRepository) *Seeder {
	return &Seeder{
		roleRepo: roleRepo,
		userRepo: userRepo,
	}
}

// Seed ejecuta todo el proceso de seeding
func (s *Seeder) Seed() error {
	log.Println("🌱 Iniciando proceso de seeding...")

	// Crear roles del sistema
	if err := s.seedRoles(); err != nil {
		return err
	}

	// Crear usuario admin solo si se especifica en variables de entorno
	if err := s.seedAdminUser(); err != nil {
		return err
	}

	log.Println("✅ Seeding completado exitosamente")
	return nil
}

// seedRoles crea los roles del sistema
func (s *Seeder) seedRoles() error {
	log.Println("📋 Creando roles del sistema...")

	roles := []models.Role{
		{
			Name:        models.RoleAdmin,
			Description: "Administrador del sistema con acceso completo",
			IsActive:    true,
			IsSystem:    true,
			Priority:    1,
		},
		{
			Name:        models.RoleDev,
			Description: "Desarrollador con permisos de administración",
			IsActive:    true,
			IsSystem:    true,
			Priority:    2,
		},
		{
			Name:        models.RoleUser,
			Description: "Usuario regular del sistema",
			IsActive:    true,
			IsSystem:    true,
			Priority:    3,
		},
	}

	for _, role := range roles {
		// Verificar si el rol ya existe
		existingRole, err := s.roleRepo.GetByName(context.Background(), role.Name)
		if err == nil && existingRole != nil {
			log.Printf("ℹ️  Rol '%s' ya existe, saltando...", role.Name)
			continue
		}

		// Crear el rol
		if err := s.roleRepo.Create(context.Background(), &role); err != nil {
			log.Printf("❌ Error creando rol '%s': %v", role.Name, err)
			return err
		}

		log.Printf("✅ Rol '%s' creado exitosamente", role.Name)
	}

	return nil
}

// seedAdminUser crea el usuario administrador desde variables de entorno
func (s *Seeder) seedAdminUser() error {
	// Obtener credenciales desde variables de entorno
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	adminName := os.Getenv("ADMIN_NAME")

	// Si no se especifican credenciales, no crear usuario admin
	if adminEmail == "" || adminPassword == "" {
		log.Println("ℹ️  No se especificaron credenciales de admin en variables de entorno")
		log.Println("   Para crear un usuario admin, configura:")
		log.Println("   - ADMIN_EMAIL")
		log.Println("   - ADMIN_PASSWORD")
		log.Println("   - ADMIN_NAME (opcional)")
		return nil
	}

	log.Println("👤 Creando usuario administrador...")

	// Obtener el rol admin
	adminRole, err := s.roleRepo.GetByName(context.Background(), models.RoleAdmin)
	if err != nil {
		log.Printf("❌ Error obteniendo rol admin: %v", err)
		return err
	}

	// Verificar si el usuario admin ya existe
	existingUser, err := s.userRepo.GetByEmailIncludingDeleted(adminEmail)
	if err == nil && existingUser != nil {
		log.Printf("ℹ️  Usuario admin '%s' ya existe, saltando...", adminEmail)
		return nil
	}

	// Hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("❌ Error hasheando contraseña: %v", err)
		return err
	}

	// Usar nombre por defecto si no se especifica
	if adminName == "" {
		adminName = "Administrador"
	}

	// Crear usuario admin
	adminUser := models.User{
		Name:      adminName,
		Email:     adminEmail,
		Password:  string(hashedPassword),
		RoleID:    adminRole.ID,
		IsActive:  true,
		IsDeleted: false,
	}

	if err := s.userRepo.Create(&adminUser); err != nil {
		log.Printf("❌ Error creando usuario admin: %v", err)
		return err
	}

	log.Printf("✅ Usuario administrador '%s' creado exitosamente", adminEmail)
	log.Println("⚠️  IMPORTANTE: Cambia la contraseña después del primer login")

	return nil
}

// SeedDevUser crea un usuario desarrollador desde variables de entorno
func (s *Seeder) SeedDevUser() error {
	// Obtener credenciales desde variables de entorno
	devEmail := os.Getenv("DEV_EMAIL")
	devPassword := os.Getenv("DEV_PASSWORD")
	devName := os.Getenv("DEV_NAME")

	// Si no se especifican credenciales, no crear usuario dev
	if devEmail == "" || devPassword == "" {
		log.Println("ℹ️  No se especificaron credenciales de dev en variables de entorno")
		log.Println("   Para crear un usuario dev, configura:")
		log.Println("   - DEV_EMAIL")
		log.Println("   - DEV_PASSWORD")
		log.Println("   - DEV_NAME (opcional)")
		return nil
	}

	log.Println("👨‍💻 Creando usuario desarrollador...")

	// Obtener el rol dev
	devRole, err := s.roleRepo.GetByName(context.Background(), models.RoleDev)
	if err != nil {
		log.Printf("❌ Error obteniendo rol dev: %v", err)
		return err
	}

	// Verificar si el usuario dev ya existe
	existingUser, err := s.userRepo.GetByEmailIncludingDeleted(devEmail)
	if err == nil && existingUser != nil {
		log.Printf("ℹ️  Usuario dev '%s' ya existe, saltando...", devEmail)
		return nil
	}

	// Hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(devPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("❌ Error hasheando contraseña: %v", err)
		return err
	}

	// Usar nombre por defecto si no se especifica
	if devName == "" {
		devName = "Desarrollador"
	}

	// Crear usuario dev
	devUser := models.User{
		Name:      devName,
		Email:     devEmail,
		Password:  string(hashedPassword),
		RoleID:    devRole.ID,
		IsActive:  true,
		IsDeleted: false,
	}

	if err := s.userRepo.Create(&devUser); err != nil {
		log.Printf("❌ Error creando usuario dev: %v", err)
		return err
	}

	log.Printf("✅ Usuario desarrollador '%s' creado exitosamente", devEmail)

	return nil
}
