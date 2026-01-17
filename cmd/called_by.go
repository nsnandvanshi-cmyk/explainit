package cmd

import (
	"fmt"

	"explainit/analyzer"
	"explainit/storage"
)

func CalledBy(function string) {
	index, err := storage.LoadIndex()
	if err != nil {
		fmt.Println("Please run `scan` first.")
		return
	}

	callers := analyzer.FindCalledBy(index.Files, function)

	if len(callers) == 0 {
		fmt.Println("No callers found for:", function)
		return
	}

	fmt.Println("Function", function, "is called by:")
	for _, fn := range callers {
		fmt.Printf(" - %s (%s)\n", fn.Name, fn.File)
	}
}
