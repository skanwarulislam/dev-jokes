package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Providers struct {
		Google struct {
			ClientId     string `yaml:"clientId"`
			ClientSecret string `yaml:"clientSecret"`
			CallbackUrl  string `yaml:"callbackUrl"`
		}
		Github struct {
			ClientId     string `yaml:"clientId"`
			ClientSecret string `yaml:"clientSecret"`
			CallbackUrl  string `yaml:"callbackUrl"`
		}
	} `yaml:"providers"`
	Server struct {
		Port uint32 `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

func ReadFile(cfg *Config, filePath string) {
	log.Println("Loading config file")
	confContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	confContent = []byte(os.ExpandEnv(string(confContent)))
	if err := yaml.Unmarshal(confContent, cfg); err != nil {
		log.Fatal(err)
	}
}
