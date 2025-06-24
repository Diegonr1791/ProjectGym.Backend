# =============================================================================
# 🚀 GYMBRO API - MAKEFILE
# =============================================================================

.PHONY: help build run test clean seed deploy

# Variables
BINARY_NAME=gymbro-api
MAIN_PATH=cmd/main.go
SEED_PATH=cmd/seed/main.go

# Colores para output
GREEN=\033[0;32m
YELLOW=\033[1;33m
RED=\033[0;31m
NC=\033[0m # No Color

# =============================================================================
# COMANDOS PRINCIPALES
# =============================================================================

help: ## Mostrar esta ayuda
	@echo "$(GREEN)🚀 GymBro API - Comandos disponibles:$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-15s$(NC) %s\n", $$1, $$2}'
	@echo ""

build: ## Construir la aplicación
	@echo "$(GREEN)🔨 Construyendo aplicación...$(NC)"
	go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "$(GREEN)✅ Aplicación construida: $(BINARY_NAME)$(NC)"

run: ## Ejecutar la aplicación
	@echo "$(GREEN)🚀 Ejecutando aplicación...$(NC)"
	go run $(MAIN_PATH)

dev: ## Ejecutar en modo desarrollo
	@echo "$(GREEN)🔧 Ejecutando en modo desarrollo...$(NC)"
	@echo "$(YELLOW)Variables de entorno:$(NC)"
	@echo "  - Cargando desde .env"
	@echo "  - Modo debug activado"
	go run $(MAIN_PATH)

# =============================================================================
# SEEDING
# =============================================================================

seed: ## Ejecutar seeding de datos iniciales
	@echo "$(GREEN)🌱 Ejecutando seeding...$(NC)"
	go run $(SEED_PATH)

seed-only: ## Solo seeding sin servidor
	@echo "$(GREEN)🌱 Ejecutando seeding manual...$(NC)"
	go run $(SEED_PATH)

# =============================================================================
# TESTING
# =============================================================================

test: ## Ejecutar tests
	@echo "$(GREEN)🧪 Ejecutando tests...$(NC)"
	go test ./...

test-verbose: ## Ejecutar tests con output detallado
	@echo "$(GREEN)🧪 Ejecutando tests con output detallado...$(NC)"
	go test -v ./...

test-coverage: ## Ejecutar tests con coverage
	@echo "$(GREEN)🧪 Ejecutando tests con coverage...$(NC)"
	go test -cover ./...

# =============================================================================
# DEPLOYMENT
# =============================================================================

deploy: ## Deploy en Railway
	@echo "$(GREEN)🚀 Deployando en Railway...$(NC)"
	./scripts/deploy.sh

deploy-railway: ## Deploy directo en Railway
	@echo "$(GREEN)🚂 Deployando en Railway...$(NC)"
	railway up

# =============================================================================
# UTILIDADES
# =============================================================================

clean: ## Limpiar archivos generados
	@echo "$(GREEN)🧹 Limpiando archivos...$(NC)"
	rm -f $(BINARY_NAME)
	rm -f main
	@echo "$(GREEN)✅ Limpieza completada$(NC)"

deps: ## Instalar dependencias
	@echo "$(GREEN)📦 Instalando dependencias...$(NC)"
	go mod download
	go mod tidy
	@echo "$(GREEN)✅ Dependencias instaladas$(NC)"

fmt: ## Formatear código
	@echo "$(GREEN)🎨 Formateando código...$(NC)"
	go fmt ./...
	@echo "$(GREEN)✅ Código formateado$(NC)"

lint: ## Ejecutar linter
	@echo "$(GREEN)🔍 Ejecutando linter...$(NC)"
	golangci-lint run
	@echo "$(GREEN)✅ Linter completado$(NC)"

# =============================================================================
# DATABASE
# =============================================================================

db-migrate: ## Ejecutar migraciones de base de datos
	@echo "$(GREEN)🗄️  Ejecutando migraciones...$(NC)"
	@echo "$(YELLOW)⚠️  Las migraciones se ejecutan automáticamente al iniciar$(NC)"

db-reset: ## Resetear base de datos (¡CUIDADO!)
	@echo "$(RED)⚠️  ¡CUIDADO! Esto eliminará todos los datos$(NC)"
	@read -p "¿Estás seguro? (y/N): " confirm && [ "$$confirm" = "y" ] || exit 1
	@echo "$(GREEN)🗄️  Reseteando base de datos...$(NC)"
	@echo "$(YELLOW)Implementar lógica de reset aquí$(NC)"

# =============================================================================
# DOCUMENTACIÓN
# =============================================================================

docs: ## Generar documentación Swagger
	@echo "$(GREEN)📚 Generando documentación Swagger...$(NC)"
	swag init -g cmd/main.go
	@echo "$(GREEN)✅ Documentación generada$(NC)"

docs-serve: ## Servir documentación localmente
	@echo "$(GREEN)📖 Sirviendo documentación en http://localhost:8080/swagger/index.html$(NC)"
	go run $(MAIN_PATH)

# =============================================================================
# MONITORING
# =============================================================================

logs: ## Ver logs en Railway
	@echo "$(GREEN)📋 Mostrando logs de Railway...$(NC)"
	railway logs

logs-follow: ## Ver logs en tiempo real
	@echo "$(GREEN)📋 Mostrando logs en tiempo real...$(NC)"
	railway logs --follow

status: ## Estado del proyecto en Railway
	@echo "$(GREEN)📊 Estado del proyecto...$(NC)"
	railway status

# =============================================================================
# DESARROLLO
# =============================================================================

setup: ## Configurar proyecto para desarrollo
	@echo "$(GREEN)⚙️  Configurando proyecto...$(NC)"
	@if [ ! -f .env ]; then \
		echo "$(YELLOW)📝 Creando archivo .env desde env.example...$(NC)"; \
		cp env.example .env; \
		echo "$(GREEN)✅ Archivo .env creado. Edítalo con tus configuraciones$(NC)"; \
	else \
		echo "$(GREEN)✅ Archivo .env ya existe$(NC)"; \
	fi
	@echo "$(GREEN)📦 Instalando dependencias...$(NC)"
	go mod download
	go mod tidy
	@echo "$(GREEN)✅ Configuración completada$(NC)"
	@echo "$(YELLOW)📋 Próximos pasos:$(NC)"
	@echo "  1. Editar archivo .env con tus configuraciones"
	@echo "  2. Configurar base de datos PostgreSQL"
	@echo "  3. Ejecutar: make run"

# =============================================================================
# DOCKER
# =============================================================================

docker-build: ## Construir imagen Docker
	@echo "$(GREEN)🐳 Construyendo imagen Docker...$(NC)"
	docker build -t gymbro-api .
	@echo "$(GREEN)✅ Imagen construida$(NC)"

docker-run: ## Ejecutar contenedor Docker
	@echo "$(GREEN)🐳 Ejecutando contenedor Docker...$(NC)"
	docker run -p 8080:8080 --env-file .env gymbro-api

docker-clean: ## Limpiar imágenes Docker
	@echo "$(GREEN)🧹 Limpiando imágenes Docker...$(NC)"
	docker rmi gymbro-api 2>/dev/null || true
	@echo "$(GREEN)✅ Limpieza completada$(NC)"

# =============================================================================
# AYUDA RÁPIDA
# =============================================================================

quick-start: ## Inicio rápido del proyecto
	@echo "$(GREEN)🚀 Inicio rápido de GymBro API$(NC)"
	@echo ""
	@echo "$(YELLOW)1. Configurar proyecto:$(NC)"
	@echo "   make setup"
	@echo ""
	@echo "$(YELLOW)2. Ejecutar aplicación:$(NC)"
	@echo "   make run"
	@echo ""
	@echo "$(YELLOW)3. Ejecutar seeding:$(NC)"
	@echo "   make seed"
	@echo ""
	@echo "$(YELLOW)4. Ver documentación:$(NC)"
	@echo "   http://localhost:8080/swagger/index.html"
	@echo ""
	@echo "$(GREEN)🎉 ¡Listo!$(NC)" 