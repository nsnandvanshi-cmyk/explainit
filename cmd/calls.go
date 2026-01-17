package cmd

import (
	"fmt"

	"explainit/analyzer"
	"explainit/storage"
)

func Calls(function string) {
	index, err := storage.LoadIndex()
	if err != nil {
		fmt.Println("Please run `scan` first.")
		return
	}

	calls := analyzer.FindCalls(index.Files, function)

	if len(calls) == 0 {
		fmt.Println("No calls found for:", function)
		return
	}

	fmt.Println("Function", function, "calls:")
	for _, c := range calls {
		fmt.Println(" -", c)
	}
}
