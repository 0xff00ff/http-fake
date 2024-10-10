package config

import (
	"fmt"
	"fserver/routes"
)

type Config struct {
	Port        string
	CorsEnabled bool `yaml:"corsEnabled"`
	Routes      map[string]routes.Route
}

func (c Config) Validate() error {
	for k, v := range c.Routes {
		err := v.Validate()
		if err != nil {
			return fmt.Errorf("route %s returned validation error: %w", k, err)
		}
	}
	return nil
}
