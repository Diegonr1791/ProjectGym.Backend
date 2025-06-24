# GymBro API - Documentación Completa

## Resumen Ejecutivo

GymBro es una API RESTful para la gestión de rutinas de gimnasio, construida con Go, siguiendo los principios de Clean Architecture. La API proporciona funcionalidades completas para la gestión de usuarios, rutinas, ejercicios, sesiones y mediciones, con un sistema robusto de autenticación y autorización.

## 🚀 Características

- **🔐 Autenticación JWT** con refresh tokens
- **👥 Sistema de Roles** (Admin, Dev, User) con autorización granular
- **🛡️ Middleware de Autorización** para operaciones sensibles
- **📊 Gestión de Rutinas** y ejercicios
- **📈 Seguimiento de Mediciones** corporales
- **💪 Grupos Musculares** y tipos de ejercicios
- **⭐ Sistema de Favoritos** para rutinas
- **📝 Sesiones de Entrenamiento** con ejercicios
- **🌱 Sistema de Seeding** automático para datos iniciales
- **📚 Documentación Swagger** completa
- **🐳 Docker** listo para producción
- **☁️ Deploy automático** en Railway

## 🏗️ Arquitectura

El proyecto sigue **Clean Architecture** con las siguientes capas:

```
📁 ProjectGym.Backend/
├── 📁 cmd/                    # Punto de entrada de la aplicación
├── 📁 internal/               # Lógica interna de la aplicación
│   ├── 📁 adapters/          # Adaptadores entre capas
│   ├── 📁 auth/              # Autenticación y autorización
│   ├── 📁 config/            # Configuración y DI container
│   ├── 📁 domain/            # Entidades y reglas de negocio
│   └── 📁 usecase/           # Casos de uso de la aplicación
├── 📁 interfaces/             # Interfaces externas (HTTP, etc.)
├── 📁 infraestructure/       # Implementaciones de infraestructura
└── 📁 docs/                  # Documentación
```

## 🌱 Sistema de Seeding

La aplicación incluye un sistema de seeding automático que crea:

### **Roles del Sistema**

- **Admin**: Acceso completo al sistema
- **Dev**: Permisos de administración para desarrollo
- **User**: Usuario regular

### **Usuarios Iniciales**

Los usuarios se crean automáticamente según las variables de entorno configuradas:

```bash
# Variables para usuario administrador
ADMIN_EMAIL=tu-email@ejemplo.com
ADMIN_PASSWORD=tu-contraseña-segura
ADMIN_NAME=Tu Nombre (opcional)

# Variables para usuario desarrollador
DEV_EMAIL=dev@ejemplo.com
DEV_PASSWORD=contraseña-dev
DEV_NAME=Nombre Dev (opcional)
```

### **Ejecutar Seeding**

```bash
# Usando Makefile (recomendado)
make seed

# Comando directo
go run cmd/seed/main.go

# Seeding automático al iniciar la aplicación
go run cmd/main.go

# Seeding en Railway
railway run go run cmd/seed/main.go
```

### **Comandos Útiles**

```bash
# Ver todos los comandos disponibles
make help

# Configurar proyecto para desarrollo
make setup

# Ejecutar aplicación
make run

# Ejecutar tests
make test

# Deploy en Railway
make deploy

# Ver logs de Railway
make logs
```

⚠️ **IMPORTANTE**:

- Las credenciales se configuran mediante variables de entorno
- Cambia las contraseñas después del primer login
- Nunca commits credenciales en el código

## Características Principales

### 🔐 Autenticación y Autorización

- **JWT Authentication**: Tokens de acceso seguros con información de rol
- **Refresh Tokens**: Renovación automática de sesiones
- **Role-based Access Control**: Control de acceso basado en roles
- **Secure Password Hashing**: Bcrypt para almacenamiento seguro
- **Authorization Middleware**: Middleware de autorización para operaciones sensibles
- **User Deletion Protection**: Solo admin y dev pueden eliminar usuarios

### 👥 Gestión de Usuarios

- **Soft Delete**: Borrado lógico con capacidad de restauración
- **Validaciones Robustas**: Email, contraseña, nombre y rol
- **Estados de Usuario**: Activo/inactivo, eliminado/restaurado
- **Gestión de Roles**: Roles del sistema y personalizados
- **Protected Operations**: Eliminación de usuarios protegida por autorización

### 🏋️ Gestión de Rutinas

- **Rutinas Personalizadas**: Creación y gestión de rutinas
- **Grupos Musculares**: Asociación con ejercicios
- **Rutinas Favoritas**: Sistema de favoritos
- **Rutinas Públicas/Privadas**: Control de visibilidad

### 📊 Seguimiento y Mediciones

- **Sesiones de Entrenamiento**: Registro de sesiones
- **Mediciones Corporales**: Peso, grasa, músculo
- **Historial Completo**: Seguimiento temporal
- **Métricas de Rendimiento**: Análisis de progreso

## Arquitectura del Sistema

### Clean Architecture

```
ProjectGym.Backend/
├── cmd/                    # Punto de entrada de la aplicación
├── docs/                   # Documentación (Swagger, etc.)
├── infraestructure/        # Capa de infraestructura
│   └── persistence/        # Repositorios GORM
├── interfaces/             # Capa de interfaces
│   └── http/              # Handlers HTTP
├── internal/              # Lógica de negocio interna
│   ├── adapters/          # Adaptadores para Clean Architecture
│   ├── auth/              # Autenticación y autorización
│   ├── config/            # Configuración y factory patterns
│   ├── domain/            # Entidades y reglas de negocio
│   └── usecase/           # Casos de uso
└── pkg/                   # Paquetes compartidos
```

### Patrones de Diseño

- **Repository Pattern**: Abstracción de acceso a datos
- **Use Case Pattern**: Lógica de negocio centralizada
- **Dependency Injection**: Inyección de dependencias
- **Middleware Pattern**: Interceptores HTTP
- **Factory Pattern**: Creación de middlewares
- **Adapter Pattern**: Adaptadores para Clean Architecture

## Endpoints de la API

### Autenticación

```
POST   /api/v1/auth/login      - Iniciar sesión (incluye información de rol)
POST   /api/v1/auth/logout     - Cerrar sesión
POST   /api/v1/auth/refresh    - Renovar token
```

### Usuarios

```
GET    /api/v1/users                    - Obtener usuarios activos
GET    /api/v1/users/all               - Obtener todos los usuarios
GET    /api/v1/users/deleted           - Obtener usuarios eliminados
POST   /api/v1/users                   - Crear usuario
GET    /api/v1/users/:id               - Obtener usuario por ID
PUT    /api/v1/users/:id               - Actualizar usuario
DELETE /api/v1/users/:id               - Borrado lógico (solo admin/dev)
POST   /api/v1/users/:id/restore       - Restaurar usuario
DELETE /api/v1/users/:id/permanent     - Borrado físico (solo admin/dev)
GET    /api/v1/users/email/:email      - Obtener usuario por email
```

### Roles

```
GET    /api/v1/roles                   - Obtener roles
POST   /api/v1/roles                   - Crear rol
GET    /api/v1/roles/:id               - Obtener rol por ID
PUT    /api/v1/roles/:id               - Actualizar rol
DELETE /api/v1/roles/:id               - Eliminar rol
```

### Rutinas

```
GET    /api/v1/routines                - Obtener rutinas
POST   /api/v1/routines                - Crear rutina
GET    /api/v1/routines/:id            - Obtener rutina por ID
PUT    /api/v1/routines/:id            - Actualizar rutina
DELETE /api/v1/routines/:id            - Eliminar rutina
```

### Ejercicios

```
GET    /api/v1/exercises               - Obtener ejercicios
POST   /api/v1/exercises               - Crear ejercicio
GET    /api/v1/exercises/:id           - Obtener ejercicio por ID
PUT    /api/v1/exercises/:id           - Actualizar ejercicio
DELETE /api/v1/exercises/:id           - Eliminar ejercicio
```

### Sesiones

```
GET    /api/v1/sessions                - Obtener sesiones
POST   /api/v1/sessions                - Crear sesión
GET    /api/v1/sessions/:id            - Obtener sesión por ID
PUT    /api/v1/sessions/:id            - Actualizar sesión
DELETE /api/v1/sessions/:id            - Eliminar sesión
```

### Mediciones

```
GET    /api/v1/measurements            - Obtener mediciones
POST   /api/v1/measurements            - Crear medición
GET    /api/v1/measurements/:id        - Obtener medición por ID
PUT    /api/v1/measurements/:id        - Actualizar medición
DELETE /api/v1/measurements/:id        - Eliminar medición
```

## Nuevas Funcionalidades Implementadas

### 🛡️ Sistema de Autorización Profesional

#### Características

- **Clean Architecture**: Separación clara de responsabilidades
- **Middleware de Autorización**: Verificación de roles en tiempo real
- **Factory Pattern**: Creación centralizada de middlewares
- **Adapter Pattern**: Adaptadores para evitar dependencias circulares
- **Usecase de Autorización**: Lógica de autorización centralizada

#### Roles del Sistema

```go
const (
    RoleAdmin = "admin"
    RoleUser  = "user"
    RoleDev   = "dev"
)
```

#### Operaciones Protegidas

- **Eliminación de Usuarios**: Solo admin y dev pueden eliminar usuarios
- **Borrado Físico**: Solo admin y dev pueden hacer hard delete
- **Extensible**: Fácil agregar nuevas operaciones protegidas

#### Flujo de Autorización

1. **Autenticación**: JWT con RoleID incluido
2. **Validación de Token**: Middleware JWT extrae claims
3. **Verificación de Rol**: Consulta a base de datos para validar rol
4. **Validación de Permisos**: Verificación contra roles permitidos
5. **Ejecución**: Operación permitida o error 403

### 🗑️ Borrado Lógico (Soft Delete)

#### Características

- **Preservación de Datos**: Los usuarios no se eliminan físicamente
- **Restauración**: Capacidad de restaurar usuarios eliminados
- **Filtrado Automático**: Las consultas normales excluyen usuarios eliminados
- **Auditoría**: Mantenimiento de historial completo

#### Campos de Control

```go
IsDeleted bool `gorm:"default:false" json:"is_deleted"`
IsActive  bool `gorm:"default:true" json:"is_active"`
```

#### Operaciones Disponibles

- **Soft Delete**: Marcar como eliminado lógicamente
- **Restore**: Restaurar usuario eliminado
- **Hard Delete**: Eliminación física permanente
- **Get Deleted**: Obtener lista de usuarios eliminados

### ✅ Validaciones Robustas

#### Validación de Email

- Formato de email válido (regex)
- Verificación de unicidad
- Manejo de conflictos con usuarios eliminados
- Longitud máxima: 255 caracteres

#### Validación de Contraseña

- Longitud: 8-128 caracteres
- Complejidad: Mayúscula + minúscula + número
- Hashing seguro con bcrypt
- Validación de fortaleza

#### Validación de Nombre

- Longitud: 2-100 caracteres
- Caracteres válidos: letras, espacios, guiones, apóstrofes
- Soporte para caracteres acentuados
- Validación de caracteres especiales

#### Validación de Rol

- Campo obligatorio para nuevos usuarios
- Preservación en actualizaciones
- Verificación de existencia

### 🔒 Seguridad Mejorada

#### Autenticación

- Verificación de estado activo en login
- Manejo de usuarios inactivos
- Validación de credenciales mejorada
- Tokens JWT seguros con información de rol

#### Autorización

- Control de acceso basado en roles
- Roles del sistema protegidos
- Jerarquía de permisos
- Middleware de autorización profesional
- Verificación de permisos en tiempo real

## Códigos de Error

### Categorías de Error

1. **Validación de Email**: `EMAIL_REQUIRED`, `INVALID_EMAIL_FORMAT`, `EMAIL_ALREADY_EXISTS`
2. **Validación de Contraseña**: `PASSWORD_REQUIRED`, `PASSWORD_TOO_SHORT`, `PASSWORD_WEAK`
3. **Validación de Nombre**: `NAME_REQUIRED`, `NAME_TOO_SHORT`, `INVALID_NAME_CHARACTERS`
4. **Estados de Usuario**: `USER_INACTIVE`, `USER_ALREADY_DELETED`, `USER_NOT_DELETED`
5. **Autenticación**: `INVALID_CREDENTIALS`, `USER_INACTIVE`
6. **Autorización**: `FORBIDDEN`, `INSUFFICIENT_PERMISSIONS`, `ROLE_INFO_UNAVAILABLE`

### Formato de Respuesta

```json
{
  "code": "INSUFFICIENT_PERMISSIONS",
  "message": "Insufficient permissions. Only admin and dev roles can perform this action."
}
```

## Configuración y Despliegue

### Variables de Entorno

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=gymbro
JWT_SECRET=your-secret-key
SERVER_PORT=8080
```

### Base de Datos

- **PostgreSQL**: Base de datos principal
- **GORM**: ORM para Go
- **Migraciones**: Scripts de migración incluidos
- **Índices**: Optimización de consultas

### Dependencias Principales

```go
github.com/gin-gonic/gin          // Framework web
github.com/golang-jwt/jwt         // JWT tokens
golang.org/x/crypto/bcrypt        // Hashing de contraseñas
gorm.io/gorm                      // ORM
github.com/swaggo/swag            // Documentación Swagger
```

## Testing

### Tipos de Testing

- **Unit Tests**: Casos de uso y validaciones
- **Integration Tests**: Endpoints y flujos completos
- **Performance Tests**: Consultas y operaciones críticas
- **Authorization Tests**: Pruebas de permisos y roles

### Casos de Prueba Críticos

- [ ] Validaciones de entrada
- [ ] Borrado lógico y restauración
- [ ] Autenticación y autorización
- [ ] Manejo de errores
- [ ] Conflictos de datos
- [ ] Verificación de permisos por rol

## Documentación Adicional

### Archivos de Documentación

- `swagger.yaml` - Especificación OpenAPI completa
- `
