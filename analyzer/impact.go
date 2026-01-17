package analyzer

func AnalyzeImpact(files []string, target string) []Function {
	return FindCalledBy(files, target)
}
