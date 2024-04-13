package enver

import (
	"fmt"
	"os"
	"strings"
)

const (
	Type_Env = "env"
)

type Source_Env struct {
}

func (s *Source_Env) Type() string {
	return Type_Env
}

func (s *Source_Env) Get() (map[string]any, error) {
	return s.readEnv()
}

func (s *Source_Env) readEnv() (map[string]any, error) {
	pairs := make(map[string]any)

	for _, pair := range os.Environ() {
		parts := strings.Split(pair, "=")
		if len(parts) != 2 {
			return pairs, fmt.Errorf("invalid environment variable format. Expected key=value format, got: %s", pair)
		}

		pairs[parts[0]] = parts[1]
	}

	return pairs, nil
}
