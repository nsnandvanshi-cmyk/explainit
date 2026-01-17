package main

import (
	"fmt"
	"os"

	"explainit/cmd"
	"explainit/docs"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(docs.HelpText)
		return
	}

	arg := os.Args[1]

	if arg == "help" || arg == "-h" || arg == "--help" {
		fmt.Println(docs.HelpText)
		return
	}

	switch arg {

	case "scan":
		cmd.Scan()

	case "files":
		cmd.Files()

	case "functions":
		cmd.Functions()

	case "stats":
		cmd.Stats()

	case "calls":
		if len(os.Args) < 3 {
			fmt.Println("Usage: explainit calls <function>")
			return
		}
		cmd.Calls(os.Args[2])

	case "called-by":
		if len(os.Args) < 3 {
			fmt.Println("Usage: explainit called-by <function>")
			return
		}
		cmd.CalledBy(os.Args[2])

	case "impact":
		if len(os.Args) < 3 {
			fmt.Println("Usage: explainit impact <function>")
			return
		}
		cmd.Impact(os.Args[2])

	case "output":
		if len(os.Args) < 3 || os.Args[2] != "help" {
			fmt.Println("Usage: explainit output help")
			return
		}
		cmd.OutputHelp()

	default:
		fmt.Println("Unknown command:", arg)
		fmt.Println(docs.HelpText)
	}
}
