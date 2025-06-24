# 🚀 Deploy GymBro API en Railway

## 📋 Índice

1. [Prerrequisitos](#prerrequisitos)
2. [Paso a Paso Completo](#paso-a-paso-completo)
3. [Configuración de Variables](#configuración-de-variables)
4. [Verificación Post-Deploy](#verificación-post-deploy)
5. [Comandos Útiles](#comandos-útiles)
6. [Troubleshooting](#troubleshooting)

## ✅ Prerrequisitos

- [ ] Cuenta en [Railway](https://railway.app)
- [ ] Cuenta en [GitHub](https://github.com)
- [ ] Railway CLI instalado: `npm install -g @railway/cli`
- [ ] Repositorio subido a GitHub

## 🚀 Paso a Paso Completo

### **Paso 1: Preparar el Repositorio**

```bash
# 1. Asegúrate de estar en el directorio del proyecto
cd ProjectGym.Backend

# 2. Verificar que todos los archivos estén commitados
git status

# 3. Si hay cambios, hacer commit
git add .
git commit -m "feat: prepare for Railway deployment"

# 4. Push a GitHub
git push origin main
```

### **Paso 2: Conectar Repositorio en Railway**

1. **Ir a Railway Dashboard**

   - Ve a [railway.app](https://railway.app)
   - Inicia sesión con tu cuenta

2. **Crear Nuevo Proyecto**

   - Click en "New Project"
   - Selecciona "Deploy from GitHub repo"

3. **Conectar GitHub**

   - Autoriza Railway para acceder a tu GitHub
   - Selecciona tu repositorio `ProjectGym.Backend`

4. **Configurar Deploy**
   - Railway detectará automáticamente el Dockerfile
   - Click en "Deploy Now"

### **Paso 3: Agregar Base de Datos PostgreSQL**

1. **En Railway Dashboard**

   - Ve a tu proyecto
   - Click en "New Service"
   - Selecciona "Database" → "PostgreSQL"

2. **Conectar Servicios**
   - Railway configurará automáticamente las variables de PostgreSQL
   - No necesitas hacer nada más

### **Paso 4: Configurar Variables de Entorno**

1. **Ir a Variables**

   - En tu servicio principal, ve a la pestaña "Variables"

2. **Agregar Variables Obligatorias**

   ```bash
   # JWT (¡OBLIGATORIO CAMBIAR!)
   JWT_SECRET=tu-secret-super-seguro-aqui-cambialo-en-produccion
   JWT_EXPIRATION_MINUTES=60
   REFRESH_EXPIRATION_HOURS=168
   REFRESH_MAX_AGE=604800

   # Servidor
   SERVER_PORT=8080
   ```

3. **Agregar Variables de Usuarios (Opcional)**

   ```bash
   # Usuario Administrador
   ADMIN_EMAIL=admin@tuempresa.com
   ADMIN_PASSWORD=MiContraseñaSegura123!
   ADMIN_NAME=Administrador

   # Usuario Desarrollador
   DEV_EMAIL=dev@tuempresa.com
   DEV_PASSWORD=DevContraseña456!
   DEV_NAME=Desarrollador
   ```

### **Paso 5: Deploy Automático**

```bash
# 1. Instalar Railway CLI (si no lo tienes)
npm install -g @railway/cli

# 2. Login en Railway
railway login

# 3. Navegar al proyecto
cd ProjectGym.Backend

# 4. Ejecutar deploy
./scripts/deploy.sh
```

### **Paso 6: Verificar Deploy**

```bash
# Ver estado del proyecto
railway status

# Ver logs
railway logs

# Abrir en navegador
railway open
```

## 🔧 Configuración de Variables

### **Variables Automáticas (Railway las configura)**

```bash
DB_HOST=${{Postgres.DATABASE_HOST}}
DB_PORT=${{Postgres.DATABASE_PORT}}
DB_USER=${{Postgres.DATABASE_USERNAME}}
DB_PASSWORD=${{Postgres.DATABASE_PASSWORD}}
DB_NAME=${{Postgres.DATABASE_NAME}}
```

### **Variables que DEBES configurar**

```bash
# JWT (¡CAMBIAR EN PRODUCCIÓN!)
JWT_SECRET=tu-secret-super-seguro-aqui-cambialo-en-produccion
JWT_EXPIRATION_MINUTES=60
REFRESH_EXPIRATION_HOURS=168
REFRESH_MAX_AGE=604800

# Servidor
SERVER_PORT=8080
```

### **Variables Opcionales (Para Seeding)**

```bash
# Solo si quieres usuarios automáticos
ADMIN_EMAIL=admin@tuempresa.com
ADMIN_PASSWORD=MiContraseñaSegura123!
ADMIN_NAME=Administrador

DEV_EMAIL=dev@tuempresa.com
DEV_PASSWORD=DevContraseña456!
DEV_NAME=Desarrollador
```

## ✅ Verificación Post-Deploy

### **1. Health Check**

```bash
# Obtener URL del servicio
railway status

# Probar health check
curl https://tu-app.railway.app/api/v1/health
```

### **2. Swagger Documentation**

```bash
# Abrir en navegador
https://tu-app.railway.app/swagger/index.html
```

### **3. Verificar Seeding**

```bash
# Ver logs de seeding
railway logs | grep -i seed

# Ejecutar seeding manual si es necesario
railway run go run cmd/seed/main.go
```

### **4. Probar Login**

```bash
# Si configuraste usuarios automáticos
curl -X POST https://tu-app.railway.app/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@tuempresa.com",
    "password": "MiContraseñaSegura123!"
  }'
```

## 🛠️ Comandos Útiles

### **Desarrollo Local**

```bash
# Configurar proyecto
make setup

# Ejecutar aplicación
make run

# Ejecutar seeding
make seed

# Ver ayuda
make help
```

### **Railway**

```bash
# Estado del proyecto
railway status

# Ver logs
railway logs

# Logs en tiempo real
railway logs --follow

# Abrir en navegador
railway open

# Ejecutar comando
railway run <comando>

# Deploy
railway up
```

### **Seeding**

```bash
# Seeding automático (al iniciar la app)
go run cmd/main.go

# Seeding manual
go run cmd/seed/main.go

# Seeding en Railway
railway run go run cmd/seed/main.go
```

## 🔍 Troubleshooting

### **Error: "connection refused"**

- Verificar que PostgreSQL esté activo
- Verificar variables de entorno de base de datos
- Verificar configuración de SSL

### **Error: "JWT secret not configured"**

- Configurar JWT_SECRET en Railway Variables
- Verificar que no esté vacío

### **Error: "No se crean usuarios"**

- Verificar variables ADMIN_EMAIL, ADMIN_PASSWORD
- Ejecutar seeding manual: `railway run go run cmd/seed/main.go`
- Verificar logs: `railway logs`

### **Error: "CORS policy"**

- Configurar CORS_ALLOWED_ORIGINS
- Usar wildcard temporal: `*`

### **Error: "Service not found"**

- Verificar que el proyecto esté vinculado: `railway link`
- Verificar que estés en el directorio correcto

## 📊 Monitoreo

### **Logs**

```bash
# Ver logs en tiempo real
railway logs --follow

# Ver logs de errores
railway logs | grep ERROR

# Ver logs de seeding
railway logs | grep -i seed
```

### **Métricas**

- Railway Dashboard → Metrics
- Uso de CPU y memoria
- Requests por minuto
- Tiempo de respuesta

## 🔒 Seguridad

### **Buenas Prácticas**

1. **JWT Secret**: Usar secret único y complejo (mínimo 32 caracteres)
2. **Contraseñas**: Cambiar después del primer login
3. **Variables**: Nunca commitear en código
4. **CORS**: Configurar orígenes específicos

### **Auditoría**

```bash
# Verificar variables sensibles
railway variables | grep -i secret

# Verificar logs de autenticación
railway logs | grep -i auth
```

## 📞 Soporte

### **Recursos**

- [Railway Documentation](https://docs.railway.app/)
- [Go Documentation](https://golang.org/doc/)
- [Gin Framework](https://gin-gonic.com/docs/)

### **Comandos de Emergencia**

```bash
# Reiniciar servicio
railway service restart

# Ver variables
railway variables

# Ver logs detallados
railway logs --json
```

---

**🎉 ¡Tu GymBro API está lista para producción!**
