package analyzer

import (
	"os"
	"path/filepath"

	"explainit/utils"
)

func WalkRepo(root string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if utils.ShouldIgnore(path) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, path)
		return nil
	})

	return files, err
}
