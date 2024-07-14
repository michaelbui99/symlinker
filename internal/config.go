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
