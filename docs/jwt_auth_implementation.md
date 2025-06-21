# 🔐 Implementación de Autenticación JWT en GymBro API

## Índice

1. [Introducción](#introducción)
2. [Variables de entorno y configuración](#variables-de-entorno-y-configuración)
3. [Estructura de tokens JWT y Refresh](#estructura-de-tokens-jwt-y-refresh)
4. [Utilidades JWT](#utilidades-jwt)
5. [Middleware de autenticación](#middleware-de-autenticación)
6. [Endpoints de autenticación](#endpoints-de-autenticación)
7. [Protección de rutas](#protección-de-rutas)
8. [Flujo de autenticación completo](#flujo-de-autenticación-completo)
9. [Buenas prácticas y recomendaciones](#buenas-prácticas-y-recomendaciones)

---

## 1. Introducción

La autenticación JWT (JSON Web Token) permite proteger los endpoints de la API de manera segura y escalable. Se implementó un sistema profesional con:

- **Access Token**: para autenticación rápida y corta duración.
- **Refresh Token**: para renovar el access token sin re-login.
- **Middleware**: que protege todos los endpoints salvo login/refresh/logout.

---

## 2. Variables de entorno y configuración

En el archivo `.env` o variables de entorno del sistema:

```env
JWT_SECRET=supersecreto123
JWT_EXPIRATION_MINUTES=60
```

En `internal/config/env.go`:

```go
type Config struct {
    // ...
    JWTSecret            string
    JWTExpirationMinutes string
}

func (c *Config) GetJWTSecret() string { return c.JWTSecret }
func (c *Config) GetJWTExpirationMinutes() string { return c.JWTExpirationMinutes }
```

Esto permite que la lógica JWT sea desacoplada y configurable.

---

## 3. Estructura de tokens JWT y Refresh

En `internal/auth/jwt.go`:

```go
type CustomClaims struct {
    UserID int    `json:"user_id"`
    Email  string `json:"email"`
    jwt.RegisteredClaims
}

type RefreshClaims struct {
    UserID int    `json:"user_id"`
    Type   string `json:"type"` // siempre "refresh"
    jwt.RegisteredClaims
}
```

- **Access Token**: contiene el ID y email del usuario, y expiración corta.
- **Refresh Token**: contiene solo el ID y un claim de tipo, y expiración larga (7 días).

---

## 4. Utilidades JWT

En `internal/auth/jwt.go`:

- **GenerateJWT**: genera un access token.
- **ValidateJWT**: valida un access token.
- **GenerateRefreshToken**: genera un refresh token.
- **ValidateRefreshToken**: valida un refresh token.

Ambas funciones usan la interfaz `JWTConfig` para obtener el secreto y expiración.

---

## 5. Middleware de autenticación

En `internal/auth/middleware.go`:

```go
func JWTAuthMiddleware(cfg JWTConfig) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Extrae y valida el token del header Authorization
        // Si es válido, agrega los claims al contexto
    }
}
```

- Protege rutas automáticamente.
- Si el token es inválido o falta, responde 401.
- Los claims quedan disponibles en el contexto para los handlers.

---

## 6. Endpoints de autenticación

En `interfaces/http/auth_handler.go`:

- **POST /auth/login**: recibe email y password, retorna access y refresh token.
- **POST /auth/refresh**: recibe refresh token, retorna nuevos tokens.
- **POST /auth/logout**: (demo) invalida el refresh token (en producción, usar blacklist o base de datos).

Ejemplo de request de login:

```json
{
  "email": "usuario@ejemplo.com",
  "password": "123456"
}
```

Respuesta:

```json
{
  "access_token": "...",
  "refresh_token": "...",
  "user": { "id": 1, "email": "usuario@ejemplo.com" }
}
```

---

## 7. Protección de rutas

En `internal/config/server.go`:

```go
protected := s.router.Group("/")
protected.Use(auth.JWTAuthMiddleware(s.config))
// Todos los handlers se registran sobre 'protected'
```

- Todas las rutas (excepto `/auth/*` y `/swagger/*`) requieren JWT válido.
- El token se envía en el header:
  ```
  Authorization: Bearer <access_token>
  ```

---

## 8. Flujo de autenticación completo

1. **Login**: el usuario envía email y password a `/auth/login`.
2. **Tokens**: recibe access y refresh token.
3. **Acceso**: usa el access token en el header para acceder a endpoints protegidos.
4. **Expiración**: si el access token expira, usa `/auth/refresh` con el refresh token para obtener uno nuevo.
5. **Logout**: opcionalmente, envía el refresh token a `/auth/logout` para invalidarlo.

---

## 9. Buenas prácticas y recomendaciones

- **Nunca expongas el JWT_SECRET** en el código fuente.
- **El refresh token debe tener expiración más larga** y almacenarse de forma segura (idealmente en httpOnly cookie o almacenamiento seguro del cliente).
- **En producción**, implementa una blacklist de refresh tokens (en memoria, base de datos o Redis) para poder invalidar sesiones.
- **No expongas datos sensibles** en los claims del token.
- **Documenta los endpoints protegidos** en Swagger usando `@security BearerAuth`.

---

## Ejemplo de uso con curl

```bash
# Login
curl -X POST http://localhost:8080/auth/login -H "Content-Type: application/json" -d '{"email":"usuario@ejemplo.com","password":"123456"}'

# Usar access token
curl -H "Authorization: Bearer <access_token>" http://localhost:8080/usuarios

# Refrescar token
curl -X POST http://localhost:8080/auth/refresh -H "Content-Type: application/json" -d '{"refresh_token":"<refresh_token>"}'
```

---

_Implementación profesional de JWT y refresh token en GymBro API._
