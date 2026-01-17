package cmd

import (
	"fmt"

	"explainit/analyzer"
	"explainit/storage"
)

func Impact(function string) {
	index, err := storage.LoadIndex()
	if err != nil {
		fmt.Println("Please run `scan` first.")
		return
	}

	affected := analyzer.AnalyzeImpact(index.Files, function)

	if len(affected) == 0 {
		fmt.Println("No impact detected for:", function)
		return
	}

	fmt.Println("Impact analysis for", function+":")
	fmt.Println()
	fmt.Println("Directly affected functions:")

	for _, fn := range affected {
		fmt.Printf(" - %s (%s)\n", fn.Name, fn.File)
	}
}
