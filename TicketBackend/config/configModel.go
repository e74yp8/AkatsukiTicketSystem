package config

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	MySQL       MySQLConfig `yaml:"mysql"`
	ServicePort int         `yaml:"service_port"`
}
