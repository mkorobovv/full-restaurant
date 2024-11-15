package config

type Config struct {
	Adapters Adapters `yaml:"adapters"`
}

type Adapters struct {
	Secondary Secondary `yaml:"secondary"`
	Primary   Primary   `yaml:"primary"`
}

type Primary struct {
	HttpAdapter HttpAdapter `yaml:"httpAdapter"`
}

type Secondary struct {
	Databases Databases `yaml:"databases"`
}

type Databases struct {
	Restaurant Database `yaml:"restaurant"`
}

type Database struct {
	Type     string `yaml:"type"`
	Host     string `env:"RESTAURANT_HOST"     env-required:"true" yaml:"host"`
	Port     string `env:"RESTAURANT_PORT"     env-required:"true" yaml:"port"`
	User     string `env:"RESTAURANT_USER"     env-required:"true" yaml:"user"`
	Password string `env:"RESTAURANT_PASSWORD" env-required:"true" yaml:"password"`
	Name     string `yaml:"name"`
}
