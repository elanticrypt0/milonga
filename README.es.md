# Milonga

Un framework Go potente y flexible para construir aplicaciones web modernas.

## Características

- 🚀 Servidor web de alto rendimiento usando Fiber
- 🛠 Herramienta CLI integrada para generación de código
- 📦 Soporte Docker incluido
- 🔄 Recarga en caliente para desarrollo
- 🗄️ Integración con GORM para operaciones de base de datos
- 🔒 Características de seguridad incorporadas

## Inicio Rápido

### Prerequisitos

- Go 1.19 o superior
- Docker y Docker Compose (opcional)
- Bun (para desarrollo de interfaz web)

### Instalación

```bash
# Clonar el repositorio
git clone https://github.com/yourusername/milonga.git
cd milonga

# Instalar dependencias
go mod download
```

### Desarrollo

```bash
# Ejecutar la aplicación
go run main.go

# O con recarga en caliente
air
```

### Usando Docker

```bash
# Construir la imagen
docker-compose build

# Iniciar el servicio
docker-compose up

# O en modo desacoplado
docker-compose up -d

# Ver logs
docker-compose logs -f
```

## Herramienta CLI

Milonga viene con una potente herramienta CLI para generación de código.

### Generar Modelos CRUD

```bash
# Generar un nuevo modelo con operaciones CRUD
go run main.go generate model Usuario

# Esto creará:
# - api/models/usuario.go
# - api/handlers/usuario_handler.go
# - api/routes/usuario_routes.go
```

## Estructura del Proyecto

```
.
├── api/
│   ├── handlers/    # Manejadores de peticiones
│   ├── models/      # Modelos de datos
│   └── routes/      # Definiciones de rutas
├── cmd/
│   └── cli/         # Comandos CLI
├── config/          # Archivos de configuración
├── public/          # Archivos estáticos
└── docker-compose.yaml
```

## Configuración

Carpetas necesarias para la compilación:

- config
  - app_config.toml
  - db_config.toml
- public

## Interfaz de Usuario Web

Para construir la interfaz de usuario web:

> Nota: La API debe estar ejecutándose para que la construcción funcione

```bash
bun run build
```

El resultado se colocará en el directorio `/public`.

## Acceso a la API

El puerto predeterminado es 8921 (configurable)

- API: [http://localhost:8921](http://localhost:8921)
- Archivos públicos: [http://localhost:8921/public](http://localhost:8921/public)
- Ejemplo HTMX: [http://localhost:8921/public/examplex.html](http://localhost:8921/public/examplex.html)

## Contribuir

¡Las contribuciones son bienvenidas! No dudes en enviar un Pull Request.

## Licencia

Este proyecto está licenciado bajo la Licencia MIT - ver el archivo LICENSE para más detalles.
