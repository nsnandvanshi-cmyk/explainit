package analyzer

func FindCalledBy(files []string, target string) []Function {
	var callers []Function

	functions := DetectFunctions(files)

	for _, fn := range functions {
		calls := FindCalls(files, fn.Name)
		for _, c := range calls {
			if c == target {
				callers = append(callers, fn)
				break
			}
		}
	}

	return callers
}
