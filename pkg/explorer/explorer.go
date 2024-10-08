package explorer

import (
	"fmt"
	"os"
	"path/filepath"
)

// bluprint to make a directory explorer

type Explorer struct {
	Files []string
}

func New(scriptsPath string) *Explorer {
	exp := &Explorer{}
	return exp
}

func (me *Explorer) Finally(path2scan string) {
	// show files
	for _, file := range me.Files {
		fmt.Printf("Dir: %q > %q\n", path2scan, file)
	}

	// fmt.Println(">> Capturas de pantalla completadas.") // Mensaje al finalizar
}

func (me *Explorer) traverseDirectory(dir string) error {

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileInfo, err := me.getFileInfo(path)
			if err != nil {
				return err
			}
			me.Files = append(me.Files, fileInfo)
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (me *Explorer) getFileInfo(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return "", err
	}
	name := info.Name()

	return name, nil
}

func (me *Explorer) Scan(path2scan string) {
	// Define un WaitGroup
	err := me.traverseDirectory(path2scan)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	me.Finally(path2scan)
}
