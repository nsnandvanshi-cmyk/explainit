package utils

func ShouldIgnore(path string) bool {
	ignoreList := []string{
		".git",
		"node_modules",
		".venv",
		"vendor",
	}

	for _, ignore := range ignoreList {
		if contains(path, ignore) {
			return true
		}
	}

	return false
}

func contains(path string, keyword string) bool {
	return len(path) >= len(keyword) &&
		(len(path) == len(keyword) ||
			(path != "" && keyword != "" &&
				stringContains(path, keyword)))
}

func stringContains(s, substr string) bool {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
