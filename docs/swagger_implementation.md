# 🚀 Guía Completa de Implementación de Swagger en GymBro API

## 📋 Tabla de Contenidos

1. [Introducción](#introducción)
2. [Configuración Global](#configuración-global)
3. [Configuración del Servidor](#configuración-del-servidor)
4. [Documentación de Endpoints](#documentación-de-endpoints)
5. [Modelos de Datos](#modelos-de-datos)
6. [Flujo Completo](#flujo-completo)
7. [Comandos Útiles](#comandos-útiles)
8. [Ejemplos Prácticos](#ejemplos-prácticos)

---

## 🎯 Introducción

Swagger es una herramienta que permite documentar APIs de forma interactiva. En este proyecto, hemos implementado Swagger para documentar automáticamente todos los endpoints de la API GymBro.

### 📦 Dependencias Utilizadas

```go
github.com/swaggo/swag          // Generador de documentación
github.com/swaggo/gin-swagger   // Integración con Gin
github.com/swaggo/files         // Servidor de archivos estáticos
```

---

## ⚙️ Configuración Global

### Archivo: `cmd/main.go`

```go
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
```

### 📝 Explicación de Cada Anotación

| Anotación              | Propósito              | Ejemplo                                   |
| ---------------------- | ---------------------- | ----------------------------------------- |
| `@title`               | Título de la API       | "GymBro API"                              |
| `@version`             | Versión de la API      | "1.0"                                     |
| `@description`         | Descripción general    | "API para gestión de rutinas de gimnasio" |
| `@host`                | URL base del servidor  | "localhost:8080"                          |
| `@BasePath`            | Ruta base de endpoints | "/"                                       |
| `@securityDefinitions` | Tipo de autenticación  | Bearer token                              |

### 🔐 Configuración de Autenticación

```go
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
```

**¿Qué hace?**

- Define que la API usa autenticación por Bearer token
- El token se envía en el header `Authorization`
- Formato: `Bearer <token>`

---

## 🖥️ Configuración del Servidor

### Archivo: `internal/config/server.go`

```go
package config

import (
    http "github.com/Diegonr1791/GymBro/interfaces/http"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// Server configura y maneja el servidor HTTP
type Server struct {
    router    *gin.Engine
    container *Container
}

// setupRoutes configura todas las rutas de la aplicación
func (s *Server) setupRoutes() {
    // Configurar Swagger
    s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Configurar handlers
    http.NewUsuarioHandler(s.router, s.container.UsuarioService)
    http.NewRutinaHandler(s.router, s.container.RutinaService)
    // ... otros handlers
}
```

### 🔧 Línea Clave de Swagger

```go
s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

**¿Qué hace esta línea?**

1. **Registra la ruta**: `/swagger/*any` captura todas las rutas que empiecen con `/swagger/`
2. **WrapHandler**: Envuelve el handler de Swagger para que funcione con Gin
3. **swaggerFiles.Handler**: Sirve los archivos estáticos de Swagger UI

### 📁 Estructura de Rutas Swagger

```
/swagger/
├── index.html          # Interfaz principal de Swagger UI
├── doc.json           # Especificación JSON de la API
├── doc.yaml           # Especificación YAML de la API
└── swagger-ui/        # Archivos estáticos de la interfaz
```

---

## 📝 Documentación de Endpoints

### Archivo: `interfaces/http/ejercicio_handler.go`

```go
// @Summary Obtener todos los ejercicios
// @Description Obtiene la lista completa de ejercicios
// @Tags ejercicios
// @Accept json
// @Produce json
// @Success 200 {array} model.Ejercicio
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /ejercicio [get]
func (h *ExerciseHandler) GetAll(c *gin.Context) {
    ejercicios, err := h.uc.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, ejercicios)
}
```

### 📋 Anotaciones de Documentación

| Anotación      | Propósito                  | Ejemplo                                   |
| -------------- | -------------------------- | ----------------------------------------- |
| `@Summary`     | Título corto del endpoint  | "Obtener todos los ejercicios"            |
| `@Description` | Descripción detallada      | "Obtiene la lista completa de ejercicios" |
| `@Tags`        | Agrupa endpoints           | "ejercicios"                              |
| `@Accept`      | Tipo de contenido aceptado | "json"                                    |
| `@Produce`     | Tipo de contenido devuelto | "json"                                    |
| `@Success`     | Respuesta exitosa          | `200 {array} model.Ejercicio`             |
| `@Failure`     | Posibles errores           | `500 {object} map[string]interface{}`     |
| `@Router`      | Ruta y método HTTP         | `/ejercicio [get]`                        |

### 🔍 Ejemplo con Parámetros

```go
// @Summary Obtener ejercicio por ID
// @Description Obtiene un ejercicio específico por su ID
// @Tags ejercicios
// @Accept json
// @Produce json
// @Param id path int true "ID del ejercicio"
// @Success 200 {object} model.Ejercicio
// @Failure 400 {object} map[string]interface{} "ID inválido"
// @Failure 404 {object} map[string]interface{} "Ejercicio no encontrado"
// @Router /ejercicio/{id} [get]
func (h *ExerciseHandler) GetById(c *gin.Context) {
    // Implementación...
}
```

### 📥 Ejemplo con Body

```go
// @Summary Crear nuevo ejercicio
// @Description Crea un nuevo ejercicio en el sistema
// @Tags ejercicios
// @Accept json
// @Produce json
// @Param ejercicio body model.Ejercicio true "Datos del ejercicio"
// @Success 201 {object} model.Ejercicio
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /ejercicio [post]
func (h *ExerciseHandler) Create(c *gin.Context) {
    // Implementación...
}
```

### 📋 Tipos de Parámetros

| Tipo     | Descripción           | Ejemplo                       |
| -------- | --------------------- | ----------------------------- |
| `path`   | Parámetro en la URL   | `/ejercicio/{id}`             |
| `query`  | Parámetro de consulta | `?page=1&limit=10`            |
| `body`   | Cuerpo del request    | JSON en el body               |
| `header` | Header HTTP           | `Authorization: Bearer token` |

---

## 🏗️ Modelos de Datos

### Archivo: `internal/domain/models/ejercicio.go`

```go
package model

type Ejercicio struct {
    ID              uint   `gorm:"primaryKey" json:"id"`
    Nombre          string `json:"nombre"`
    TipoEjercicioID uint   `json:"tipo_ejercicio_id"`
    GrupoMuscularID uint   `json:"grupo_muscular_id"`
}

func (e *Ejercicio) TableName() string {
    return "ejercicios"
}
```

### 🏷️ Tags de Struct

| Tag       | Propósito                           | Ejemplo                              |
| --------- | ----------------------------------- | ------------------------------------ |
| `json`    | Serialización JSON                  | `json:"id"`                          |
| `gorm`    | Configuración de base de datos      | `gorm:"primaryKey"`                  |
| `swagger` | Documentación específica de Swagger | `swagger:"description(Descripción)"` |

### 📊 Ejemplo de Modelo Completo

```go
// @Description Modelo de ejercicio
type Ejercicio struct {
    // @Description ID único del ejercicio
    ID uint `gorm:"primaryKey" json:"id" example:"1"`

    // @Description Nombre del ejercicio
    Nombre string `json:"nombre" example:"Press de banca"`

    // @Description ID del tipo de ejercicio
    TipoEjercicioID uint `json:"tipo_ejercicio_id" example:"1"`

    // @Description ID del grupo muscular
    GrupoMuscularID uint `json:"grupo_muscular_id" example:"2"`
}
```

---

## 🔄 Flujo Completo

### 1. Generación de Documentación

```bash
# Generar documentación desde el archivo main.go
swag init -g cmd/main.go
```

**¿Qué hace este comando?**

- Escanea todos los archivos Go del proyecto
- Busca anotaciones de Swagger
- Genera archivos en la carpeta `docs/`

### 2. Archivos Generados

```
docs/
├── docs.go          # Código Go con la documentación
├── swagger.json     # Especificación JSON
└── swagger.yaml     # Especificación YAML
```

### 3. Importación en main.go

```go
import (
    _ "github.com/Diegonr1791/GymBro/docs" // Importar docs generados
)
```

**¿Por qué el `_`?**

- Solo se importa por sus efectos secundarios
- No se usa directamente en el código
- Inicializa la documentación automáticamente

### 4. Configuración del Servidor

```go
func (s *Server) setupRoutes() {
    // Configurar Swagger
    s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // ... otras rutas
}
```

### 5. Acceso a Swagger UI

```
http://localhost:8080/swagger/index.html
```

---

## 🛠️ Comandos Útiles

### Generación de Documentación

```bash
# Generar documentación básica
swag init

# Generar desde archivo específico
swag init -g cmd/main.go

# Generar con configuración personalizada
swag init -g cmd/main.go --parseDependency --parseInternal

# Regenerar documentación
swag init -g cmd/main.go --parseDependency
```

### Verificación

```bash
# Verificar que swag está instalado
swag --version

# Ver ayuda
swag --help
```

### Instalación de swag

```bash
# Instalar swag
go install github.com/swaggo/swag/cmd/swag@latest

# Verificar instalación
swag --version
```

---

## 📚 Ejemplos Prácticos

### Endpoint GET con Query Parameters

```go
// @Summary Buscar ejercicios
// @Description Busca ejercicios por nombre
// @Tags ejercicios
// @Accept json
// @Produce json
// @Param nombre query string false "Nombre del ejercicio"
// @Param limit query int false "Límite de resultados" default(10)
// @Param page query int false "Número de página" default(1)
// @Success 200 {array} model.Ejercicio
// @Router /ejercicio/buscar [get]
func (h *ExerciseHandler) Search(c *gin.Context) {
    // Implementación...
}
```

### Endpoint POST con Autenticación

```go
// @Summary Crear rutina
// @Description Crea una nueva rutina (requiere autenticación)
// @Tags rutinas
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param rutina body model.Rutina true "Datos de la rutina"
// @Success 201 {object} model.Rutina
// @Failure 401 {object} map[string]interface{} "No autorizado"
// @Router /rutina [post]
func (h *RutinaHandler) Create(c *gin.Context) {
    // Implementación...
}
```

### Endpoint con Respuestas Múltiples

```go
// @Summary Obtener usuario
// @Description Obtiene información del usuario actual
// @Tags usuarios
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.Usuario
// @Failure 401 {object} map[string]interface{} "No autorizado"
// @Failure 404 {object} map[string]interface{} "Usuario no encontrado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /usuario/perfil [get]
func (h *UsuarioHandler) GetProfile(c *gin.Context) {
    // Implementación...
}
```

---

## 🎨 Personalización de Swagger UI

### Configuración Avanzada

```go
// En server.go
import (
    "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

func (s *Server) setupRoutes() {
    // Configuración personalizada de Swagger
    s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
        ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
        ginSwagger.DefaultModelsExpandDepth(-1),
        ginSwagger.DocExpansion("none"),
    ))
}
```

### Opciones de Configuración

| Opción                     | Descripción                         | Valores                                  |
| -------------------------- | ----------------------------------- | ---------------------------------------- |
| `URL`                      | URL de la especificación            | `http://localhost:8080/swagger/doc.json` |
| `DefaultModelsExpandDepth` | Profundidad de expansión de modelos | `-1` (cerrado), `0` (abierto)            |
| `DocExpansion`             | Expansión de endpoints              | `"none"`, `"list"`, `"full"`             |

---

## 🔍 Troubleshooting

### Problemas Comunes

1. **Swagger no se muestra**

   - Verificar que `swag init` se ejecutó correctamente
   - Verificar que se importa el paquete `docs`
   - Verificar que la ruta `/swagger/*any` está registrada

2. **Documentación desactualizada**

   - Ejecutar `swag init -g cmd/main.go` después de cambios
   - Reiniciar el servidor

3. **Errores de compilación**
   - Verificar que todas las dependencias están instaladas
   - Verificar que las anotaciones están bien escritas

### Verificación de Instalación

```bash
# Verificar que swag está instalado
which swag

# Verificar versión
swag --version

# Verificar que las dependencias están en go.mod
cat go.mod | grep swaggo
```

---

## 📖 Recursos Adicionales

- [Documentación oficial de Swaggo](https://github.com/swaggo/swag)
- [Swagger UI](https://swagger.io/tools/swagger-ui/)
- [OpenAPI Specification](https://swagger.io/specification/)

---

## 🎯 Conclusión

La implementación de Swagger en GymBro API proporciona:

✅ **Documentación automática** de todos los endpoints  
✅ **Interfaz interactiva** para probar la API  
✅ **Especificación estándar** (OpenAPI)  
✅ **Fácil mantenimiento** con anotaciones en el código  
✅ **Integración perfecta** con Gin framework

Esta implementación sigue las mejores prácticas y proporciona una excelente experiencia de desarrollo para todos los usuarios de la API.

---

_Documento generado para GymBro API - Implementación de Swagger_
