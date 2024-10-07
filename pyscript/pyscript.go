package pyscript

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

type PyScript struct {
	scriptsPath string
	storagePath string
}

func New(scriptsPath string, storage2ath string) *PyScript {
	return &PyScript{
		scriptsPath: scriptsPath,
		storagePath: storage2ath,
	}
}

func (me *PyScript) runScript(script2execute, filePathArg string) error {
	absolutePathPythonScript := me.scriptsPath + "/" + script2execute
	cmd := exec.Command("python3", absolutePathPythonScript, filePathArg) // Llamar al script de Python
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error al ejecutar %q: %s\n", script2execute, err)
		return err
	}

	fmt.Printf("PYTHON3 >>> RUNNING %q %q\n", absolutePathPythonScript, filePathArg)

	return nil
}

func (me *PyScript) ApplyScript(list []string, script2execute string) error {

	log.Printf("Running script: %q \n", me.scriptsPath+"/"+script2execute)
	var wg sync.WaitGroup

	hasError := false
	var lastErr error
	for _, item := range list {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := me.runScript(script2execute, item)
			if err != nil {
				hasError = true
				lastErr = err
				fmt.Printf("Error on execute %q | arg: %q : %s\n", script2execute, item, err)
			}
		}()

	}

	wg.Wait()

	log.Printf("PyScript>> All done!\n")

	if hasError {
		return lastErr
	}
	return nil
}
