#!/bin/bash

# Verificamos si el nombre del proyecto fue proporcionado
if [ -z "$1" ]; then
    echo "Por favor, proporciona el nombre del proyecto."
    echo "Uso: $0 <nombre_del_proyecto>"
    exit 1
fi

# Nombre del proyecto
PROJECT_NAME=$1

# Función para crear la estructura de directorios y archivos
create_structure() {
    # Raíz del proyecto
    mkdir -p "$PROJECT_NAME"

    # Directorios principales
    mkdir -p "$PROJECT_NAME/cmd/cli"
    mkdir -p "$PROJECT_NAME/internal/"
    mkdir -p "$PROJECT_NAME/pkg"
    mkdir -p "$PROJECT_NAME/api/handlers"
    mkdir -p "$PROJECT_NAME/api/models"
    mkdir -p "$PROJECT_NAME/web/static/css"
    mkdir -p "$PROJECT_NAME/web/static/js"
    mkdir -p "$PROJECT_NAME/web/templates"
    mkdir -p "$PROJECT_NAME/scripts"
    mkdir -p "$PROJECT_NAME/configs"
    mkdir -p "$PROJECT_NAME/database"
    mkdir -p "$PROJECT_NAME/tests/unit"
    mkdir -p "$PROJECT_NAME/tests/integration"
    mkdir -p "$PROJECT_NAME/docs"

    # Archivos dentro de 'cmd'
    touch "$PROJECT_NAME/cmd/cli/main.go"

    # Archivos dentro de 'pkg'
    touch "$PROJECT_NAME/pkg/mypackage-example.go"

    # Archivos dentro de 'api'
    touch "$PROJECT_NAME/api/handler/handler-example.go"
    touch "$PROJECT_NAME/api/model/model-example.go"

    # Archivos dentro de 'web/templates'
    echo "Hello world!" > "$PROJECT_NAME/web/templates/index.html"

    # Archivos dentro de 'scripts'
    touch "$PROJECT_NAME/scripts/build.sh"
    touch "$PROJECT_NAME/scripts/deploy.sh"

    # Archivos dentro de 'configs'
    touch "$PROJECT_NAME/configs/development.toml"
    touch "$PROJECT_NAME/configs/production.toml"

    # Archivos de configuración y documentación
    touch "$PROJECT_NAME/.gitignore"
    touch "$PROJECT_NAME/README.md"

    go mod init $PROJECT_NAME
}

# Llamamos a la función para crear la estructura
create_structure

echo "Estructura de directorios y archivos creada con éxito para el proyecto '$PROJECT_NAME'."
