# Pagination Package

El paquete `pagination` proporciona una implementación simple y eficiente para manejar la paginación en aplicaciones Go. Este paquete es útil para dividir grandes conjuntos de datos en páginas más pequeñas y manejables.

## Estructura

```go
type Pagination struct {
    PageFirst    uint // Primera página (siempre es 1)
    PageLast     uint // Última página
    PageCurrent  uint // Página actual
    PageNext     uint // Página siguiente
    PagePrev     uint // Página anterior
    ItemsPerPage uint // Elementos por página
    TotalItems   uint // Total de elementos
}
```

## Funciones y Métodos

### NewPagination

```go
func NewPagination(currentPage, itemsPerPage, totalItems uint) *Pagination
```

Crea una nueva instancia de Pagination con los parámetros proporcionados.

#### Parámetros
- `currentPage`: Número de página actual
- `itemsPerPage`: Cantidad de elementos por página
- `totalItems`: Total de elementos en el conjunto de datos

#### Ejemplo de uso

```go
// Crear una nueva paginación
// Página actual: 1
// Elementos por página: 10
// Total de elementos: 100
pagination := NewPagination(1, 10, 100)
```

### calculatePages

```go
func (p *Pagination) calculatePages(totalItems uint)
```

Método interno que calcula el número total de páginas basado en la cantidad total de elementos y elementos por página.

### ToString

```go
func (p *Pagination) ToString(elem string) string
```

Convierte los valores de paginación a string según el elemento especificado.

#### Parámetros
- `elem`: String que especifica qué valor retornar ("end", "start", "next", "prev", "current")

#### Valores de retorno
- Retorna el valor solicitado como string

#### Ejemplo de uso

```go
pagination := NewPagination(1, 10, 100)

lastPage := pagination.ToString("end")    // Obtiene la última página
firstPage := pagination.ToString("start") // Obtiene la primera página
nextPage := pagination.ToString("next")   // Obtiene la página siguiente
prevPage := pagination.ToString("prev")   // Obtiene la página anterior
currPage := pagination.ToString("current") // Obtiene la página actual
```

## Ejemplo Completo

```go
package main

import (
    "fmt"
    "yourpackage/pagination"
)

func main() {
    // Crear una nueva paginación con:
    // - Página actual: 2
    // - 10 elementos por página
    // - 100 elementos en total
    pag := pagination.NewPagination(2, 10, 100)

    fmt.Printf("Página actual: %d\n", pag.PageCurrent)
    fmt.Printf("Elementos por página: %d\n", pag.ItemsPerPage)
    fmt.Printf("Total de páginas: %d\n", pag.PageLast)
    fmt.Printf("Siguiente página: %d\n", pag.PageNext)
    fmt.Printf("Página anterior: %d\n", pag.PagePrev)
}
```

## Fórmulas de Cálculo

La paginación se calcula usando las siguientes reglas:

1. **Última página**: Se calcula dividiendo el total de elementos entre los elementos por página
   ```
   PageLast = ⌈TotalItems / ItemsPerPage⌉
   ```

2. **Página siguiente**: Se calcula si la página actual + 1 es menor o igual a la última página
   ```
   PageNext = PageCurrent + 1 (si PageCurrent + 1 ≤ PageLast)
   ```

3. **Página anterior**: Se calcula si la página actual - 1 es mayor que la primera página
   ```
   PagePrev = PageCurrent - 1 (si PageCurrent - 1 > PageFirst)
   ```

## Notas Importantes

1. La primera página siempre es 1
2. Si el total de elementos es menor que los elementos por página, `PageLast` será 1
3. Los valores de página siguiente y anterior se calculan automáticamente basados en la página actual
4. Todos los valores numéricos son de tipo `uint` para garantizar números positivos

## Limitaciones

1. No maneja paginación con offset
2. No incluye validación de entradas negativas (usa uint)
3. No proporciona métodos para calcular el offset de la base de datos

## Mejores Prácticas

1. **Inicialización**:
   ```go
   // Usar valores razonables para itemsPerPage
   pagination := NewPagination(1, 25, totalItems)
   ```

2. **Validación**:
   ```go
   // Verificar que la página actual existe
   if currentPage > pagination.PageLast {
       currentPage = pagination.PageLast
   }
   ```

3. **Uso con bases de datos**:
   ```go
   offset := (pagination.PageCurrent - 1) * pagination.ItemsPerPage
   limit := pagination.ItemsPerPage
   // Usar offset y limit en la consulta SQL
   ```