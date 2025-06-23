# Códigos de Error de Validación para Usuarios

## Resumen

Este documento describe todos los códigos de error específicos para validaciones de usuarios en la API de GymBro.

## Códigos de Error por Categoría

### 1. Validación de Email

| Código                 | HTTP Status | Descripción                               | Ejemplo                                |
| ---------------------- | ----------- | ----------------------------------------- | -------------------------------------- |
| `EMAIL_REQUIRED`       | 400         | El email es requerido                     | "Email is required"                    |
| `INVALID_EMAIL_FORMAT` | 400         | Formato de email inválido                 | "Invalid email format"                 |
| `EMAIL_ALREADY_EXISTS` | 409         | El email ya está en uso                   | "The provided email is already in use" |
| `EMAIL_SOFT_DELETED`   | 409         | El email pertenece a un usuario eliminado | "Email belongs to a deleted user"      |

### 2. Validación de Contraseña

| Código               | HTTP Status | Descripción                | Ejemplo                                                                                     |
| -------------------- | ----------- | -------------------------- | ------------------------------------------------------------------------------------------- |
| `PASSWORD_REQUIRED`  | 400         | La contraseña es requerida | "Password is required"                                                                      |
| `PASSWORD_TOO_SHORT` | 400         | La contraseña es muy corta | "Password must be at least 8 characters long"                                               |
| `PASSWORD_TOO_LONG`  | 400         | La contraseña es muy larga | "Password must not exceed 128 characters"                                                   |
| `PASSWORD_WEAK`      | 400         | La contraseña es débil     | "Password must contain at least one uppercase letter, one lowercase letter, and one number" |

### 3. Validación de Nombre

| Código                    | HTTP Status | Descripción                             | Ejemplo                                   |
| ------------------------- | ----------- | --------------------------------------- | ----------------------------------------- |
| `NAME_REQUIRED`           | 400         | El nombre es requerido                  | "Name is required"                        |
| `NAME_TOO_SHORT`          | 400         | El nombre es muy corto                  | "Name must be at least 2 characters long" |
| `NAME_TOO_LONG`           | 400         | El nombre es muy largo                  | "Name must not exceed 100 characters"     |
| `INVALID_NAME_CHARACTERS` | 400         | El nombre contiene caracteres inválidos | "Name contains invalid characters"        |

### 4. Validación de Rol

| Código          | HTTP Status | Descripción         | Ejemplo            |
| --------------- | ----------- | ------------------- | ------------------ |
| `ROLE_REQUIRED` | 400         | El rol es requerido | "Role is required" |

### 5. Estados de Usuario

| Código                 | HTTP Status | Descripción                        | Ejemplo                    |
| ---------------------- | ----------- | ---------------------------------- | -------------------------- |
| `USER_INACTIVE`        | 401         | La cuenta de usuario está inactiva | "User account is inactive" |
| `USER_ALREADY_DELETED` | 400         | El usuario ya está eliminado       | "User is already deleted"  |
| `USER_NOT_DELETED`     | 400         | El usuario no está eliminado       | "User is not deleted"      |

### 6. Errores de Autenticación

| Código                | HTTP Status | Descripción            | Ejemplo                        |
| --------------------- | ----------- | ---------------------- | ------------------------------ |
| `INVALID_CREDENTIALS` | 401         | Credenciales inválidas | "Invalid credentials provided" |

## Ejemplos de Respuestas de Error

### Error de Validación de Email

```json
{
  "code": "INVALID_EMAIL_FORMAT",
  "message": "Invalid email format"
}
```

### Error de Contraseña Débil

```json
{
  "code": "PASSWORD_WEAK",
  "message": "Password must contain at least one uppercase letter, one lowercase letter, and one number"
}
```

### Error de Usuario Inactivo

```json
{
  "code": "USER_INACTIVE",
  "message": "User account is inactive"
}
```

### Error de Email ya Existente

```json
{
  "code": "EMAIL_ALREADY_EXISTS",
  "message": "The provided email is already in use"
}
```

## Reglas de Validación

### Email

- **Formato**: Debe ser un email válido (ejemplo@dominio.com)
- **Unicidad**: Debe ser único en el sistema
- **Longitud**: Máximo 255 caracteres

### Contraseña

- **Longitud**: Entre 8 y 128 caracteres
- **Complejidad**: Al menos una mayúscula, una minúscula y un número
- **Hashing**: Se almacena hasheada con bcrypt

### Nombre

- **Longitud**: Entre 2 y 100 caracteres
- **Caracteres**: Solo letras, espacios, guiones (-) y apóstrofes (')
- **Soporte**: Caracteres acentuados (á, é, í, ó, ú, ñ, etc.)

### Rol

- **Obligatorio**: Para nuevos usuarios
- **Preservación**: Se mantiene en actualizaciones si no se especifica

## Estados de Usuario

### Campos de Control

- **`is_active`**: Indica si la cuenta está activa
- **`is_deleted`**: Indica si el usuario está eliminado lógicamente

### Comportamiento

- Los usuarios inactivos no pueden hacer login
- Los usuarios eliminados lógicamente no aparecen en consultas normales
- Los usuarios eliminados pueden ser restaurados
- El email de usuarios eliminados no puede ser reutilizado

## Mejores Prácticas

### Para Desarrolladores Frontend

1. **Validación en Cliente**: Implementar validaciones básicas antes de enviar al servidor
2. **Manejo de Errores**: Mostrar mensajes de error específicos al usuario
3. **UX**: Proporcionar feedback inmediato sobre la fortaleza de la contraseña

### Para Consumidores de API

1. **Códigos de Error**: Usar los códigos para lógica de negocio específica
2. **Mensajes**: Mostrar los mensajes al usuario final
3. **Retry Logic**: Implementar lógica de reintento para errores temporales

## Testing

### Casos de Prueba Recomendados

- [ ] Email con formato inválido
- [ ] Email duplicado
- [ ] Contraseña muy corta
- [ ] Contraseña sin mayúsculas
- [ ] Contraseña sin números
- [ ] Nombre muy corto
- [ ] Nombre con caracteres especiales
- [ ] Usuario inactivo intentando login
- [ ] Restaurar usuario no eliminado
- [ ] Eliminar usuario ya eliminado

### Ejemplos de Testing

```bash
# Test email inválido
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"invalid-email","password":"Password123","role_id":1}'

# Test contraseña débil
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"test@example.com","password":"weak","role_id":1}'
```
