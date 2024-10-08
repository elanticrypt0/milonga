# Milonga

Basic bootstrap to create simple and powerfull golangs apps

# You can use

default port = 9000 but you can change it 
[http://localhost:9000/public](http://localhost:9000/public)

htmlx
[http://localhost:9000/public/examplex.html](http://localhost:9000/public/examplex.html)

# Build

Carpetas necesarias del build:

- config
  - app_config.toml
  - db_config.toml
- public


## Binario de la api

Para construir el binario de la api ejecutar:

TODO:
```sh

go build [opciones]

```

## Web User Interface

Para construir el build de wui.

==Para que la construcción se produzca la API debe estar corriendo==

```sh
bun run build
```

El resultado estará es puesto en /public