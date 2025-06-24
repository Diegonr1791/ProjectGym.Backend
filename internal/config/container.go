package config

import (
	persistence "github.com/Diegonr1791/GymBro/infraestructure/persistence"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	usecase "github.com/Diegonr1791/GymBro/internal/usecase"
	"gorm.io/gorm"
)

// Container maneja todas las dependencias de la aplicación
type Container struct {
	DB *gorm.DB

	// Repositories
	RoleRepo            repository.RoleRepository
	UsuarioRepo         repository.UsuarioRepository
	RutinaRepo          repository.RutinaRepository
	GrupoMuscularRepo   repository.GrupoMuscularRepository
	RutinaGMRepo        repository.RutinaGrupoMuscularRepository
	FavoritaRepo        repository.FavoritaRepository
	MedicionRepo        repository.MedicionRepository
	TipoEjercicioRepo   repository.TypeExerciseRepository
	EjercicioRepo       repository.ExerciseRepository
	SesionRepo          repository.SessionRepository
	SesionEjercicioRepo repository.SessionExerciseRepository
	RefreshTokenRepo    repository.RefreshTokenRepository

	// Use Cases
	AuthorizationService   *usecase.AuthorizationUsecase
	RoleService            *usecase.RoleUseCase
	UsuarioService         *usecase.UsuarioUsecase
	RutinaService          *usecase.RutinaUsecase
	GrupoMuscularService   *usecase.GrupoMuscularUseCase
	RutinaGMService        *usecase.RoutineMuscleGroupUsecase
	FavoritaService        *usecase.FavoriteUsecase
	MedicionService        *usecase.MeasurementUsecase
	TipoEjercicioService   *usecase.TypeExerciseUsecase
	EjercicioService       *usecase.ExerciseUsecase
	SesionService          *usecase.SessionUsecase
	SesionEjercicioService *usecase.SessionExerciseUsecase
	RefreshTokenService    *usecase.RefreshTokenUsecase

	// Seeder
	Seeder *Seeder

	// Auth Configuration
	JWTConfig *Config
}

// NewContainer crea y configura todas las dependencias
func NewContainer() *Container {
	// Conectar a la base de datos
	ConnectDB()

	// Configurar JWT
	jwtConfig := LoadConfig()

	container := &Container{
		DB:        DB,
		JWTConfig: jwtConfig,
	}

	// Inicializar repositories
	container.initializeRepositories()

	// Inicializar use cases
	container.initializeUseCases()

	return container
}

// initializeRepositories configura todos los repositories
func (c *Container) initializeRepositories() {
	c.RoleRepo = persistence.NewRoleGormRepository(c.DB)
	c.UsuarioRepo = persistence.NewUsuarioGormRepository(c.DB)
	c.RutinaRepo = persistence.NewRutinaGormRepository(c.DB)
	c.GrupoMuscularRepo = persistence.NewGrupoMuscularGormRepository(c.DB)
	c.RutinaGMRepo = persistence.NewRutinaGrupoMuscularGormRepository(c.DB)
	c.FavoritaRepo = persistence.NewFavoritaGormRepository(c.DB)
	c.MedicionRepo = persistence.NewMedicionGormRepository(c.DB)
	c.TipoEjercicioRepo = persistence.NewTypeExerciseGormRepository(c.DB)
	c.EjercicioRepo = persistence.NewExerciseGormRepository(c.DB)
	c.SesionRepo = persistence.NewSessionGormRepository(c.DB)
	c.SesionEjercicioRepo = persistence.NewSessionExerciseGormRepository(c.DB)
	c.RefreshTokenRepo = persistence.NewRefreshTokenGormRepository(c.DB)
}

// initializeUseCases configura todos los use cases
func (c *Container) initializeUseCases() {
	c.AuthorizationService = usecase.NewAuthorizationUsecase(c.RoleRepo)
	c.RoleService = usecase.NewRoleUseCase(c.RoleRepo)
	c.UsuarioService = usecase.NewUsuarioUsecase(c.UsuarioRepo, c.RoleRepo)
	c.RutinaService = usecase.NewRutinaUsecase(c.RutinaRepo)
	c.GrupoMuscularService = usecase.NewGrupoMuscularUseCase(c.GrupoMuscularRepo)
	c.RutinaGMService = usecase.NewRoutineMuscleGroupUsecase(c.RutinaGMRepo)
	c.FavoritaService = usecase.NewFavoriteUsecase(c.FavoritaRepo)
	c.MedicionService = usecase.NewMeasurementUsecase(c.MedicionRepo)
	c.TipoEjercicioService = usecase.NewTypeExerciseUsecase(c.TipoEjercicioRepo)
	c.EjercicioService = usecase.NewExerciseUsecase(c.EjercicioRepo)
	c.SesionService = usecase.NewSessionUsecase(c.SesionRepo)
	c.SesionEjercicioService = usecase.NewSessionExerciseUsecase(c.SesionEjercicioRepo)
	c.RefreshTokenService = usecase.NewRefreshTokenUsecase(c.RefreshTokenRepo, c.UsuarioRepo, c.JWTConfig)

	// Inicializar seeder
	c.Seeder = NewSeeder(c.RoleRepo, c.UsuarioRepo)
}
