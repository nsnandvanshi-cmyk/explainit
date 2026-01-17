package cmd

import (
	"fmt"

	"explainit/storage"
)

func Files() {
	index, err := storage.LoadIndex()
	if err != nil {
		fmt.Println("Please run `scan` first.")
		return
	}

	if len(index.Files) == 0 {
		fmt.Println("No files indexed.")
		return
	}

	fmt.Println("Indexed files:")
	for _, file := range index.Files {
		fmt.Println(" -", file)
	}
}
