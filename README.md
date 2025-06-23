# GymBro API - Documentación Completa

## Resumen Ejecutivo

GymBro es una API RESTful para la gestión de rutinas de gimnasio, construida con Go, siguiendo los principios de Clean Architecture. La API proporciona funcionalidades completas para la gestión de usuarios, rutinas, ejercicios, sesiones y mediciones.

## Características Principales

### 🔐 Autenticación y Autorización

- **JWT Authentication**: Tokens de acceso seguros
- **Refresh Tokens**: Renovación automática de sesiones
- **Role-based Access Control**: Control de acceso basado en roles
- **Secure Password Hashing**: Bcrypt para almacenamiento seguro

### 👥 Gestión de Usuarios

- **Soft Delete**: Borrado lógico con capacidad de restauración
- **Validaciones Robustas**: Email, contraseña, nombre y rol
- **Estados de Usuario**: Activo/inactivo, eliminado/restaurado
- **Gestión de Roles**: Roles del sistema y personalizados

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
│   ├── auth/              # Autenticación y autorización
│   ├── config/            # Configuración
│   ├── domain/            # Entidades y reglas de negocio
│   └── usecase/           # Casos de uso
└── pkg/                   # Paquetes compartidos
```

### Patrones de Diseño

- **Repository Pattern**: Abstracción de acceso a datos
- **Use Case Pattern**: Lógica de negocio centralizada
- **Dependency Injection**: Inyección de dependencias
- **Middleware Pattern**: Interceptores HTTP

## Endpoints de la API

### Autenticación

```
POST   /api/v1/auth/login      - Iniciar sesión
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
DELETE /api/v1/users/:id               - Borrado lógico
POST   /api/v1/users/:id/restore       - Restaurar usuario
DELETE /api/v1/users/:id/permanent     - Borrado físico
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
- Tokens JWT seguros

#### Autorización

- Control de acceso basado en roles
- Roles del sistema protegidos
- Jerarquía de permisos
- Middleware de autorización

## Códigos de Error

### Categorías de Error

1. **Validación de Email**: `EMAIL_REQUIRED`, `INVALID_EMAIL_FORMAT`, `EMAIL_ALREADY_EXISTS`
2. **Validación de Contraseña**: `PASSWORD_REQUIRED`, `PASSWORD_TOO_SHORT`, `PASSWORD_WEAK`
3. **Validación de Nombre**: `NAME_REQUIRED`, `NAME_TOO_SHORT`, `INVALID_NAME_CHARACTERS`
4. **Estados de Usuario**: `USER_INACTIVE`, `USER_ALREADY_DELETED`, `USER_NOT_DELETED`
5. **Autenticación**: `INVALID_CREDENTIALS`, `USER_INACTIVE`

### Formato de Respuesta

```json
{
  "code": "EMAIL_REQUIRED",
  "message": "Email is required"
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

### Casos de Prueba Críticos

- [ ] Validaciones de entrada
- [ ] Borrado lógico y restauración
- [ ] Autenticación y autorización
- [ ] Manejo de errores
- [ ] Conflictos de datos

## Documentación Adicional

### Archivos de Documentación

- `swagger.yaml` - Especificación OpenAPI completa
- `user_validation_errors.md` - Códigos de error específicos
- `migration_soft_delete.sql` - Script de migración
- `clean_architecture.md` - Documentación de arquitectura

### Swagger UI

- **URL**: `http://localhost:8080/swagger/index.html`
- **Especificación**: `http://localhost:8080/swagger/doc.json`
- **Documentación Interactiva**: Pruebas de endpoints en tiempo real

## Mejores Prácticas

### Desarrollo

1. **Clean Architecture**: Separación clara de responsabilidades
2. **Error Handling**: Manejo consistente de errores
3. **Validation**: Validaciones en múltiples capas
4. **Documentation**: Documentación completa y actualizada

### Seguridad

1. **Input Validation**: Validación estricta de entrada
2. **Password Security**: Hashing seguro de contraseñas
3. **JWT Security**: Tokens seguros y renovables
4. **SQL Injection**: Prevención con GORM

### Performance

1. **Database Indexing**: Índices optimizados
2. **Query Optimization**: Consultas eficientes
3. **Caching Strategy**: Estrategia de caché
4. **Connection Pooling**: Pool de conexiones

## Roadmap

### Próximas Funcionalidades

- [ ] Notificaciones push
- [ ] Reportes y analytics
- [ ] Integración con wearables
- [ ] API de terceros
- [ ] Microservicios

### Mejoras Técnicas

- [ ] GraphQL API
- [ ] Event sourcing
- [ ] CQRS pattern
- [ ] Kubernetes deployment
- [ ] Monitoring y logging

## Soporte y Contribución

### Contacto

- **Email**: support@gymbro.com
- **Documentación**: `/docs`
- **Issues**: GitHub Issues
- **Discussions**: GitHub Discussions

### Contribución

1. Fork del repositorio
2. Crear rama feature
3. Implementar cambios
4. Tests y documentación
5. Pull request

---

**GymBro API v1.0** - Sistema completo de gestión de rutinas de gimnasio con arquitectura limpia y funcionalidades avanzadas de gestión de usuarios.
