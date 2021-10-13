package config

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	os.Setenv("GOOGLE_CLIENT_ID", "1")
	os.Setenv("GOOGLE_CLIENT_SECRET", "1")
	os.Setenv("GITHUB_CLIENT_ID", "1")
	os.Setenv("GITHUB_CLIENT_SECRET", "1")
	var cfg Config
	ReadFile(&cfg, "../config.yml")
}
