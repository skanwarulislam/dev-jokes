package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	const google_clientId = "123456789101-084OlCaiLJLKJfladsffjkaldsfjala.apps.googleusercontent.com"
	os.Setenv("GOOGLE_CLIENT_ID", google_clientId)
	os.Setenv("GOOGLE_CLIENT_SECRET", "iqHZGv-K789dliUDVEbRT-AldnbEjanDAB6")
	os.Setenv("GITHUB_CLIENT_ID", "ytzscq8q4lva71gchlnt")
	os.Setenv("GITHUB_CLIENT_SECRET", "zda0jgh2h3gw9nuogpzd35zuq117zmwu3gbx3d2v")
	var cfg Config
	ReadFile(&cfg, "../config.yml")
	assert.Equal(t, cfg.Providers.Google.ClientId, google_clientId, "Google ClientId is different")
	assert.Equal(t, cfg.Providers.Google.ClientSecret, "iqHZGv-K789dliUDVEbRT-AldnbEjanDAB6", "Google ClientSecret is different")

	assert.Equal(t, cfg.Providers.Github.ClientId, "ytzscq8q4lva71gchlnt", "Github Client Id is different")
	assert.Equal(t, cfg.Providers.Github.ClientSecret, "zda0jgh2h3gw9nuogpzd35zuq117zmwu3gbx3d2v", "Github Secret is different")
}
