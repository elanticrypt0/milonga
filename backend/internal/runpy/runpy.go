package pyscript

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"sync"
)

type RunPy struct {
	runtime     string
	scriptsPath string
	storagePath string
}

// ScriptResult almacena el resultado de la ejecución del script
type ScriptResult struct {
	Input  string         `json:"input"`
	Output string         `json:"output"`
	Error  error          `json:"error,omitempty"`
	Data   map[string]any `json:"data,omitempty"`
}

func New(scriptsPath string, storagePath string) *RunPy {
	return &RunPy{
		runtime:     "python",
		scriptsPath: scriptsPath,
		storagePath: storagePath,
	}
}

func (me *RunPy) runScript(ctx context.Context, script2execute, filePathArg string) (*ScriptResult, error) {
	result := &ScriptResult{
		Input: filePathArg,
		Data:  make(map[string]any),
	}

	absolutePathPythonScript := me.scriptsPath + "/" + script2execute
	cmd := exec.CommandContext(ctx, me.runtime, absolutePathPythonScript, filePathArg)

	// Buffers para capturar stdout y stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return result, fmt.Errorf("error starting script: %w", err)
	}

	// Crear un canal para recibir el error de la ejecución
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	// Esperar a que termine el comando o se cancele el contexto
	select {
	case <-ctx.Done():
		if err := cmd.Process.Kill(); err != nil {
			return result, fmt.Errorf("error killing process: %w", err)
		}
		return result, ctx.Err()
	case err := <-done:
		if err != nil {
			result.Error = fmt.Errorf("error executing script: %w\nStderr: %s", err, stderr.String())
			return result, result.Error
		}
	}

	// Procesar la salida
	output := stdout.String()
	result.Output = output

	// Intentar parsear la salida como JSON si está en ese formato
	if err := json.Unmarshal(stdout.Bytes(), &result.Data); err != nil {
		// Si no es JSON, guardamos la salida como texto plano
		result.Data["text"] = output
	}

	return result, nil
}

func (me *RunPy) ApplyScript(ctx context.Context, list []string, script2execute string) ([]*ScriptResult, error) {
	log.Printf("Running script: %q\n", me.scriptsPath+"/"+script2execute)

	var wg sync.WaitGroup
	results := make([]*ScriptResult, len(list))
	resultsChan := make(chan *ScriptResult, len(list))
	errorsChan := make(chan error, len(list))

	// Crear un contexto cancelable
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Ejecutar scripts en goroutines
	for i, item := range list {
		wg.Add(1)
		go func(index int, input string) {
			defer wg.Done()

			result, err := me.runScript(ctx, script2execute, input)
			if err != nil {
				errorsChan <- fmt.Errorf("error processing %s: %w", input, err)
				return
			}

			resultsChan <- result
		}(i, item)
	}

	// Goroutine para esperar que terminen todas las ejecuciones
	go func() {
		wg.Wait()
		close(resultsChan)
		close(errorsChan)
	}()

	// Recolectar resultados y errores
	var errors []error
	resultsMap := make(map[string]*ScriptResult)

	for {
		select {
		case result, ok := <-resultsChan:
			if !ok {
				// Canal cerrado, terminar recolección
				if len(errors) > 0 {
					return results, fmt.Errorf("multiple errors occurred: %v", errors)
				}
				return results, nil
			}
			resultsMap[result.Input] = result

		case err, ok := <-errorsChan:
			if !ok {
				continue
			}
			errors = append(errors, err)
		}
	}
}

// Método auxiliar para ejecutar un solo script y obtener su resultado
func (me *RunPy) ExecuteOne(ctx context.Context, script2execute, input string) (*ScriptResult, error) {
	return me.runScript(ctx, script2execute, input)
}
