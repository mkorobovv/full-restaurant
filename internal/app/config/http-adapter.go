package config

import "time"

type HttpAdapter struct {
	Port              string        `yaml:"port"`
	ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout"`
	WriteTimeout      time.Duration `yaml:"writeTimeout"`
	ReadTimeout       time.Duration `yaml:"readTimeout"`
	ShutdownTimeout   time.Duration `yaml:"shutdownTimeout"`
}
