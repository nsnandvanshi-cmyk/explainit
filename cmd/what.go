package cmd

import (
	"fmt"
	"os"
	"strings"

	"explainit/storage"
)

func What(symbol string) {
	index, err := storage.LoadIndex()
	if err != nil {
		fmt.Println("Please run `scan` first.")
		return
	}

	for _, file := range index.Files {
		contentBytes, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		content := string(contentBytes)
		lines := strings.Split(content, "\n")

		for i, line := range lines {
			if strings.Contains(line, symbol) {
				fmt.Println("Symbol:", symbol)
				fmt.Println()
				fmt.Println("Found in:")
				fmt.Println(" -", file)
				fmt.Println()

				// Try to explain using nearby lines
				fmt.Println("Context:")
				start := i - 3
				if start < 0 {
					start = 0
				}
				end := i + 3
				if end >= len(lines) {
					end = len(lines) - 1
				}

				for j := start; j <= end; j++ {
					fmt.Println(lines[j])
				}
				return
			}
		}
	}

	fmt.Println("Symbol not found:", symbol)
}
