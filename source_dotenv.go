package enver

import (
	"fmt"
	"os"
	"strings"
)

const (
	Type_DotEnv = "dotenv"
)

type Source_DotEnv struct {
	Path string
}

func (s *Source_DotEnv) Type() string {
	return Type_DotEnv
}

func (s *Source_DotEnv) Get() (map[string]any, error) {
	return s.readDotEnvFile(s.Path)
}

func (s *Source_DotEnv) readDotEnvFile(path string) (map[string]any, error) {
	if path == "" {
		path = ".env"
	}

	pairs := make(map[string]any)

	data, err := os.ReadFile(path)
	if err != nil {
		return pairs, fmt.Errorf("failed to read .env file: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			return pairs, fmt.Errorf("invalid .env file format. Expected key=value format, got: %s", line)
		}

		pairs[parts[0]] = parts[1]
	}

	return pairs, nil
}
