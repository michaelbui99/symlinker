package internal

import (
	"fmt"
	"os"
)

func ListFiles(dir string) ([]string, error) {
	file, err := os.Open(dir)
	if err != nil {
		return []string{}, fmt.Errorf("failed to list files in directory %s: %w", dir, err)
	}
	defer file.Close()

	files, err := file.ReadDir(-1)
	if err != nil {
		return []string{}, fmt.Errorf("failed to list files in directory %s: %w", dir, err)
	}

	res := make([]string, len(files))
	for _, f := range files {
		res = append(res, f.Name())
	}

	return res, nil
}
