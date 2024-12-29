# RunPy Go Package

runpy es un paquete en Go que facilita la ejecución asíncrona de scripts Python y la captura de sus resultados. Permite ejecutar scripts de Python desde Go, manejar su salida y procesar los resultados de forma estructurada.

## Características

- Ejecución asíncrona de scripts Python
- Captura de salida estándar (stdout) y errores (stderr)
- Soporte para resultados en formato JSON
- Manejo de contexto para cancelación de operaciones
- Ejecución paralela de múltiples scripts
- Gestión de errores robusta
- Soporte para timeout y cancelación de operaciones

## Instalación

```bash
go get github.com/yourusername/runpy
```

## Uso Básico

### Inicialización

```go
import "github.com/yourusername/runpy"

ps := runpy.New("/path/to/scripts", "/path/to/storage")
```

### Ejecutar un Solo Script

```go
ctx := context.Background()
result, err := ps.ExecuteOne(ctx, "script.py", "input.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Resultado: %+v\n", result.Data)
```

### Ejecutar Múltiples Scripts en Paralelo

```go
inputs := []string{"file1.txt", "file2.txt", "file3.txt"}
results, err := ps.ApplyScript(ctx, inputs, "process.py")
if err != nil {
    log.Fatal(err)
}

for _, result := range results {
    if result.Error != nil {
        fmt.Printf("Error processing %s: %v\n", result.Input, result.Error)
        continue
    }
    fmt.Printf("Success processing %s: %v\n", result.Input, result.Data)
}
```

## Estructura de Datos

### runpy

```go
type runpy struct {
    runtime     string
    scriptsPath string
    storagePath string
}
```

### ScriptResult

```go
type ScriptResult struct {
    Input   string          `json:"input"`
    Output  string          `json:"output"`
    Error   error          `json:"error,omitempty"`
    Data    map[string]any `json:"data,omitempty"`
}
```

## Escribiendo Scripts Python Compatibles

Para obtener el mejor resultado, los scripts Python deben escribir su salida en formato JSON. Ejemplo:

```python
# script.py
import sys
import json

def main():
    input_file = sys.argv[1]
    # Procesar el input
    result = {
        "processed_file": input_file,
        "status": "success",
        "data": {
            "key1": "value1",
            "key2": 123
        }
    }
    # Imprimir el resultado como JSON
    print(json.dumps(result))

if __name__ == "__main__":
    main()
```

## Manejo de Contexto y Timeouts

```go
// Crear contexto con timeout
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

// Ejecutar script con timeout
result, err := ps.ExecuteOne(ctx, "long_script.py", "input.txt")
if err != nil {
    if err == context.DeadlineExceeded {
        log.Println("Script execution timed out")
    }
    log.Fatal(err)
}
```

## Manejo de Errores

El paquete proporciona un manejo de errores detallado:

```go
result, err := ps.ExecuteOne(ctx, "script.py", "input.txt")
if err != nil {
    switch {
    case errors.Is(err, context.DeadlineExceeded):
        // Manejar timeout
    case errors.Is(err, context.Canceled):
        // Manejar cancelación
    default:
        // Manejar otros errores
    }
}
```

## Mejores Prácticas

1. **Siempre usar contexto**: Permite manejar timeouts y cancelaciones.
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), timeout)
   defer cancel()
   ```

2. **Verificar errores**: Siempre verificar y manejar los errores retornados.
   ```go
   if result.Error != nil {
       log.Printf("Error en el script: %v", result.Error)
   }
   ```

3. **Scripts Python**: Estructurar la salida en JSON para mejor integración.
   ```python
   print(json.dumps({"status": "success", "data": result_data}))
   ```

4. **Recursos**: Liberar recursos usando `defer` cuando sea necesario.

## Ejemplo de uso

```go
func main() {
    ctx := context.Background()
    rp := runpy.New("/path/to/scripts", "/path/to/storage")

    // Ejecutar un solo script
    result, err := rp.ExecuteOne(ctx, "ejemplo_script.py", "input.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Resultado: %+v\n", result.Data)

    // Ejecutar múltiples scripts
    inputs := []string{"file1.txt", "file2.txt", "file3.txt"}
    results, err := ps.ApplyScript(ctx, inputs, "ejemplo_script.py")
    if err != nil {
        log.Fatal(err)
    }

    for _, result := range results {
        if result.Error != nil {
            fmt.Printf("Error processing %s: %v\n", result.Input, result.Error)
            continue
        }
        fmt.Printf("Success processing %s: %v\n", result.Input, result.Data)
    }
}
```

## Limitaciones

- Los scripts Python deben estar instalados en el sistema
- Requiere Python en el sistema host
- La salida del script debe ser en formato texto o JSON
- No soporta entrada interactiva con los scripts

## Contribuir

Las contribuciones son bienvenidas. Por favor, abrir un issue para discutir cambios mayores antes de enviar un pull request.
