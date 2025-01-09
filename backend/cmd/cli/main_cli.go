package cli

import (
	"fmt"
	"os"
)

func Run() {
	if err := Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
