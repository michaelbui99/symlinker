package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type SymlinkerFile struct {
	Version string `yaml:"version"`
	Links   []Link `yaml:"links"`
}

type SymlinkerFileNotFoundError struct {
	files []string
}

func (e *SymlinkerFileNotFoundError) Error() string {
	return fmt.Sprintf("no SymlinkerFile found amongst %v", e.files)
}

func ParseSymlinkerFile(symlinkerFilePath string) (SymlinkerFile, error) {
	resolvedPath, err := filepath.Abs(symlinkerFilePath)
	if err != nil {
		return SymlinkerFile{}, fmt.Errorf("failed to resolve path to SymlinkerFile, %w", err)
	}

	bytes, err := os.ReadFile(resolvedPath)
	if err != nil {
		return SymlinkerFile{}, fmt.Errorf("failed to read SymlinkerFile, %w", err)
	}

	var symlinkerFile SymlinkerFile
	err = yaml.Unmarshal(bytes, &symlinkerFile)
	if err != nil {
		return SymlinkerFile{}, fmt.Errorf("failed to deserialize SymlinkerFile, %w", err)
	}

	for _, link := range symlinkerFile.Links {
		if len(strings.TrimSpace(link.Type)) == 0 {
			link.Type = defaultLinkType
		}
	}

	return symlinkerFile, nil
}

// Returns any supported variant of "SymlinkerFile.yaml" name
func FindSymlinkerFile(files []string) (string, error) {
	supportedNames := map[string]bool{
		"SymlinkerFile.yaml":  true,
		"symlinkerFile.yaml":  true,
		"symlinkerfile.yaml":  true,
		"SYMLINKERFILE.yaml":  true,
		"SYMLINKER_FILE.yaml": true,
	}

	for _, f := range files {
		valid, ok := supportedNames[f]
		if !ok {
			continue
		}

		// Should not be possible since we use the map as a set where all entries are valid.
		if !valid {
			continue
		}

		return f, nil
	}

	return "", &SymlinkerFileNotFoundError{
		files: files,
	}
}
