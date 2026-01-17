package cmd

import (
	"fmt"
	"os"

	"explainit/analyzer"
	"explainit/storage"
)

func Stats() {
	index, err := storage.LoadIndex()
	if err != nil {
		fmt.Println("Please run `scan` first.")
		return
	}

	files := index.Files
	fileCount := len(files)

	functions := analyzer.DetectFunctions(files)
	functionCount := len(functions)

	var largestFile string
	var largestSize int64 = 0

	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			continue
		}

		if info.Size() > largestSize {
			largestSize = info.Size()
			largestFile = file
		}
	}

	fmt.Println("Repository statistics:")
	fmt.Println()
	fmt.Println("Files indexed:", fileCount)
	fmt.Println("Functions detected:", functionCount)

	if largestFile != "" {
		fmt.Printf("Largest file: %s (%.1f KB)\n", largestFile, float64(largestSize)/1024)
	}
}
