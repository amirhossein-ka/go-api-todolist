package config

import "github.com/kelseyhightower/envconfig"

func ParseEnv(cfg *Config) (err error) {
	if err = envconfig.Process("database", &cfg.Database); err != nil {
		return err
	}
	return nil
}
