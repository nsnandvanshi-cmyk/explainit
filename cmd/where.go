package cmd

import (
	"fmt"
	"os"
	"strings"

	"explainit/storage"
)

func Where(symbol string) {
	index, err := storage.LoadIndex()
	if err != nil {
		fmt.Println("Please run `scan` first.")
		return
	}

	found := false

	for _, file := range index.Files {
		content, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		if strings.Contains(string(content), symbol) {
			if !found {
				fmt.Println("Found \"" + symbol + "\" in:")
				found = true
			}
			fmt.Println(" -", file)
		}
	}

	if !found {
		fmt.Println("Symbol not found:", symbol)
	}
}
