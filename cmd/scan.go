package cmd

import (
	"fmt"

	"explainit/analyzer"
	"explainit/storage"
)

func Scan() {
	files, err := analyzer.WalkRepo(".")
	if err != nil {
		fmt.Println("Scan failed:", err)
		return
	}

	index := storage.Index{
		Files: files,
	}

	err = storage.SaveIndex(index)
	if err != nil {
		fmt.Println("Failed to save index:", err)
		return
	}

	fmt.Println("Scan complete. Files indexed:", len(files))
}
