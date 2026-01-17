package cmd

import (
	"fmt"

	"explainit/analyzer"
	"explainit/storage"
)

func Functions() {
	index, err := storage.LoadIndex()
	if err != nil {
		fmt.Println("Please run `scan` first.")
		return
	}

	functions := analyzer.DetectFunctions(index.Files)

	if len(functions) == 0 {
		fmt.Println("No functions detected.")
		return
	}

	fmt.Println("Detected functions:")
	for _, fn := range functions {
		fmt.Printf(" - %s (%s)\n", fn.Name, fn.File)
	}
}
