package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

var DefaultCommitTypes = map[string]string{
	"feat":     "A new feature",
	"fix":      "A bug fix",
	"docs":     "Documentation only changes",
	"style":    "Formatting, white-space, missing semi-colons,...",
	"refactor": "Code changes that neither fix bugs nor add features",
	"perf":     "Code changes that improve performance",
	"test":     "Adding missing tests or correcting existing tests",
	"build":    "Changes that affect the build system or external dependencies",
	"ci":       "Changes to our CI configuration files and scripts",
	"chore":    "Other changes that don't modify src or test files",
	"revert":   "Reverts a previous commit",
}

type ValidationConfig struct {
	SummaryMaxLen       int
	ScopeMaxLen         int
	ScopeRequired       bool
	DescriptionRequired bool
}

type Config struct {
	CommitTypes map[string]string `toml:"commit_types"`
	Validations ValidationConfig  `toml:"validations"`
}

func LoadConfig(path string) Config {
	cfg := Config{
		CommitTypes: DefaultCommitTypes,
		Validations: ValidationConfig{
			SummaryMaxLen: 60,
			ScopeMaxLen:   30,
		},
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg
	}

	if err := toml.Unmarshal(data, &cfg); err != nil {
		return cfg
	}

	if cfg.CommitTypes == nil {
		cfg.CommitTypes = DefaultCommitTypes
	}

	return cfg
}
