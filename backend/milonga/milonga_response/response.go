package milonga_response

import (
	"bytes"
	"errors"
	"fmt"
	"milonga/milonga/utils"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SendHTML(c *fiber.Ctx, s string) error {
	c.Set("Content-Type", "text/html;charset=UTF-8")
	_, err := c.WriteString(s)
	return err
}

func SendHTMLFromFile(c *fiber.Ctx, filepath string) error {

	// read HTML file
	content := utils.OpenFile(filepath)

	c.Set("Content-Type", "text/html;charset=UTF-8")
	_, err := c.Write(content)
	return err
}

func SendHTMLFromFileWithData(c *fiber.Ctx, filepath string, data2parse map[string]interface{}) error {

	// read HTML file
	content := parseHTMLFile(filepath, data2parse)

	c.Set("Content-Type", "text/html;charset=UTF-8")
	_, err := c.Write(content)
	return err
}

func PrepareHTMLFromSlice(filepath string, data2parse interface{}) ([]byte, error) {
	val := reflect.ValueOf(data2parse)

	// Verificamos que el argumento sea un slice
	if val.Kind() != reflect.Slice {
		return nil, errors.New("data2parse no es un slice")
	}

	// Si el slice tiene un solo elemento
	if val.Len() == 1 {
		data := Struct2Map(val.Index(0).Interface())
		return parseHTMLFile(filepath, data), nil
	} else if val.Len() > 1 {
		var finalOutput bytes.Buffer

		for i := 0; i < val.Len(); i++ {
			dataValue := Struct2Map(val.Index(i).Interface())
			output := parseHTMLFile(filepath, dataValue)
			finalOutput.Write(output)
		}
		return finalOutput.Bytes(), nil
	} else {
		return []byte(""), errors.New("No data to parse")
	}
}

func parseHTMLFile(filepath string, data2parse map[string]interface{}) []byte {

	fileContent := utils.OpenFile(filepath)
	html := string(fileContent)

	for dataKey, dataValue := range data2parse {
		// Reemplaza el placeholder {$key} por el valor correspondiente
		placeholder := "{$" + dataKey + "}"
		html = strings.ReplaceAll(html, placeholder, fmt.Sprintf("%v", dataValue))
	}

	return []byte(html)
}

// EstructuraAMap convierte una estructura en un mapa de tipo map[string]interface{}.
func Struct2Map(estructura interface{}) map[string]interface{} {
	resultado := make(map[string]interface{})
	valor := reflect.ValueOf(estructura)

	// Asegurarse de que sea un valor válido y una estructura
	if valor.Kind() == reflect.Ptr {
		valor = valor.Elem()
	}
	if valor.Kind() != reflect.Struct {
		return resultado
	}

	// Iterar sobre los campos de la estructura
	for i := 0; i < valor.NumField(); i++ {
		campo := valor.Type().Field(i)
		valorCampo := valor.Field(i)

		// Obtener el nombre del campo desde la etiqueta JSON, si existe
		jsonTag := campo.Tag.Get("json")
		if jsonTag != "" {
			// Si la etiqueta JSON tiene una coma, tomamos la primera parte
			if commaIndex := strings.Index(jsonTag, ","); commaIndex != -1 {
				jsonTag = jsonTag[:commaIndex]
			}
			resultado[jsonTag] = valorCampo.Interface()
		} else {
			// Usar el nombre del campo como clave
			resultado[campo.Name] = valorCampo.Interface()
		}
	}

	return resultado
}

func StructSlice2Map(data interface{}) ([]interface{}, error) {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Slice {
		return nil, fmt.Errorf("se esperaba un slice, pero se recibió: %T", data)
	}

	result := make([]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		// result[i] = val.Index(i).Interface()
		result[i] = Struct2Map(val.Index(i).Interface())
	}
	return result, nil
}
