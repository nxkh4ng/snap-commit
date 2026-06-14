package utils

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

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
