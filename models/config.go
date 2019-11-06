package models

type Config struct {
	Port        string
	CorsEnabled bool `yaml:"corsEnabled"`
	Routes      map[string]Route
}
