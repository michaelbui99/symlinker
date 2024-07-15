package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	single = "SINGLE"
	module = "MODULE"
)

const defaultLinkType = single

type Link struct {
	Name   string `yaml:"name"`
	Source string `yaml:"source"`
	Target string `yaml:"target"`
	Type   string `yaml:"type"`
	Force  bool   `yaml:"force"`
}

func (l *Link) WillOverrideTarget() (bool, error) {
	targetPath, err := filepath.Abs(l.Target)
	if err != nil {
		return false, err
	}

	if _, err := os.Lstat(targetPath); err == nil {
		err = os.Remove(targetPath)
		if err != nil {
			return false, fmt.Errorf("failed to override target: %w", err)
		}
	}

	return false, nil
}
