#!/bin/bash

# =============================================================================
# 🚀 SCRIPT DE DEPLOY PARA RAILWAY - GYMBRO API
# =============================================================================

set -e  # Salir si cualquier comando falla

# Colores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Función para imprimir mensajes con colores
print_message() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_header() {
    echo -e "${BLUE}================================${NC}"
    echo -e "${BLUE}$1${NC}"
    echo -e "${BLUE}================================${NC}"
}

# =============================================================================
# VALIDACIONES INICIALES
# =============================================================================

print_header "VALIDACIONES INICIALES"

# Verificar que estamos en el directorio correcto
if [ ! -f "go.mod" ]; then
    print_error "No se encontró go.mod. Asegúrate de estar en el directorio raíz del proyecto."
    exit 1
fi

# Verificar que Railway CLI está instalado
if ! command -v railway &> /dev/null; then
    print_error "Railway CLI no está instalado."
    print_message "Instala Railway CLI con: npm install -g @railway/cli"
    print_message "O descárgalo desde: https://railway.app/cli"
    exit 1
fi

# Verificar que estás logueado en Railway
if ! railway whoami &> /dev/null; then
    print_error "No estás logueado en Railway."
    print_message "Ejecuta: railway login"
    exit 1
fi

print_message "✅ Todas las validaciones pasaron"

# =============================================================================
# CONSTRUCCIÓN Y DEPLOY
# =============================================================================

print_header "CONSTRUCCIÓN Y DEPLOY"

# Verificar que el proyecto está vinculado a Railway
if ! railway status &> /dev/null; then
    print_warning "El proyecto no está vinculado a Railway."
    print_message "Vinculando proyecto..."
    railway link
fi

# Construir y deployar
print_message "🚀 Iniciando deploy en Railway..."
railway up

# =============================================================================
# VERIFICACIÓN POST-DEPLOY
# =============================================================================

print_header "VERIFICACIÓN POST-DEPLOY"

# Obtener la URL del servicio
SERVICE_URL=$(railway status --json | grep -o '"url":"[^"]*"' | cut -d'"' -f4)

if [ -z "$SERVICE_URL" ]; then
    print_warning "No se pudo obtener la URL del servicio automáticamente."
    print_message "Verifica manualmente en: https://railway.app/dashboard"
else
    print_message "🌐 URL del servicio: $SERVICE_URL"
fi

# =============================================================================
# SEEDING DE DATOS INICIALES
# =============================================================================

print_header "SEEDING DE DATOS INICIALES"

print_message "🌱 Configurando seeding de datos iniciales..."
print_message ""

print_message "📋 Opciones disponibles:"
print_message "1. Seeding automático (recomendado)"
print_message "2. Seeding manual"
print_message "3. Solo verificar roles"
print_message ""

# Preguntar al usuario qué opción prefiere
read -p "¿Qué opción prefieres? (1/2/3): " seed_option

case $seed_option in
    1)
        print_message "✅ Configurando seeding automático..."
        print_message "   - Se ejecutará automáticamente en cada deploy"
        print_message "   - Configura las variables de entorno en Railway Dashboard"
        print_message "   - Variables necesarias: ADMIN_EMAIL, ADMIN_PASSWORD, etc."
        ;;
    2)
        print_message "✅ Ejecutando seeding manual..."
        railway run go run cmd/seed/main.go
        ;;
    3)
        print_message "✅ Solo verificando roles del sistema..."
        railway run go run cmd/main.go
        ;;
    *)
        print_warning "Opción no válida. Usando seeding automático por defecto."
        ;;
esac

print_message "✅ Configuración de seeding completada"

# =============================================================================
# INFORMACIÓN FINAL
# =============================================================================

print_header "DEPLOY COMPLETADO"

print_message "🎉 ¡Tu aplicación está desplegada exitosamente!"
print_message ""
print_message "📋 INFORMACIÓN:"
print_message "   - Los roles del sistema han sido creados automáticamente"
print_message "   - Los usuarios se crean según las variables de entorno configuradas"
print_message ""
print_message "🔧 CONFIGURACIÓN DE USUARIOS:"
print_message "   Para crear usuarios automáticamente, configura en Railway Dashboard:"
print_message "   - ADMIN_EMAIL, ADMIN_PASSWORD, ADMIN_NAME"
print_message "   - DEV_EMAIL, DEV_PASSWORD, DEV_NAME"
print_message ""
print_message "⚠️  IMPORTANTE:"
print_message "   - Cambia las contraseñas después del primer login"
print_message "   - Configura las variables de entorno en Railway Dashboard"
print_message "   - Verifica que la base de datos esté conectada"
print_message ""
print_message "🔗 Enlaces útiles:"
print_message "   - Railway Dashboard: https://railway.app/dashboard"
print_message "   - API Docs: $SERVICE_URL/swagger/index.html"
print_message "   - Health Check: $SERVICE_URL/api/v1/health"
print_message ""
print_message "🛠️  Comandos útiles:"
print_message "   - Ver logs: railway logs"
print_message "   - Abrir en navegador: railway open"
print_message "   - Ejecutar comando: railway run <comando>"
print_message "   - Ejecutar seeding: railway run go run cmd/main.go --seed" 