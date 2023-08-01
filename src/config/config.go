package config

import (
	"os"
	"strings"
)

type cfg struct {
	HttpSources []string
}

func NewCfg() (*cfg, error) {
	var httpSources []string
	content, err := os.ReadFile("./config/websource/http.txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	httpSources = append(httpSources, lines...)

	return &cfg{
		HttpSources: httpSources,
	}, nil
}
