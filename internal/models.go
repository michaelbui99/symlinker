package internal

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
