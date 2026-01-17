package analyzer

import (
	"os"
	"strings"
)

type Function struct {
	Name string
	File string
}

func DetectFunctions(files []string) []Function {
	var functions []Function

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)

			if strings.HasPrefix(line, "func ") {
				name := extractFuncName(line)
				if name != "" {
					functions = append(functions, Function{
						Name: name,
						File: file,
					})
				}
			}
		}
	}

	return functions
}

func extractFuncName(line string) string {
	line = strings.TrimPrefix(line, "func ")
	parts := strings.Split(line, "(")
	if len(parts) > 0 {
		return strings.TrimSpace(parts[0])
	}
	return ""
}
