package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type SymlinkerFile struct {
	Version string              `yaml:"version"`
	Links   []SymlinkerFileLink `yaml:"links"`
}

type SymlinkerFileLink struct {
	Name   string `yaml:"name"`
	Source string `yaml:"source"`
	Target string `yaml:"target"`
	Type   string `yaml:"type"`
	Force  bool   `yaml:"force"`
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

	return symlinkerFile, nil
}
