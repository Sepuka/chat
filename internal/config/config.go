package config

import (
	"github.com/stevenroose/gonfig"
)

type (
	httpClient struct {
		Proxy string
	}

	telegram struct {
		Token string
		Debug bool `default:false`
	}

	database struct {
		User     string
		Password string
		Name     string
		Host     string
		Port     int
	}

	log struct {
		Prod   bool
		Output string
	}

	Config struct {
		Path       string
		HttpClient httpClient `id:"http"`
		Telegram   telegram
		Database   database
		Log        log
	}
)

func GetConfig(path string) (*Config, error) {
	cfg := &Config{}
	err := gonfig.Load(cfg, gonfig.Conf{
		FileDefaultFilename: path,
		FlagIgnoreUnknown:   true,
		FlagDisable:         true,
		EnvDisable:          true,
	})
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
