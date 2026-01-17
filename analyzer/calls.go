package analyzer

import (
	"os"
	"strings"
)

func FindCalls(files []string, target string) []string {
	var calls []string

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		lines := strings.Split(string(content), "\n")
		inTarget := false
		braceCount := 0

		for _, line := range lines {
			line = strings.TrimSpace(line)

			// Detect function start
			if strings.HasPrefix(line, "func ") && strings.Contains(line, target+"(") {
				inTarget = true
			}

			if inTarget {
				braceCount += strings.Count(line, "{")
				braceCount -= strings.Count(line, "}")

				// Detect calls
				if strings.Contains(line, "(") {
					name := extractCallName(line)
					if name != "" && name != target {
						calls = appendIfMissing(calls, name)
					}
				}

				// Function ended
				if braceCount == 0 && strings.Contains(line, "}") {
					inTarget = false
				}
			}
		}
	}

	return calls
}

func extractCallName(line string) string {
	parts := strings.Split(line, "(")
	if len(parts) == 0 {
		return ""
	}

	name := parts[0]
	name = strings.TrimSpace(name)
	name = strings.Trim(name, "{")

	// Skip control structures
	if name == "if" || name == "for" || name == "switch" || name == "return" {
		return ""
	}

	return name
}

func appendIfMissing(list []string, item string) []string {
	for _, v := range list {
		if v == item {
			return list
		}
	}
	return append(list, item)
}
