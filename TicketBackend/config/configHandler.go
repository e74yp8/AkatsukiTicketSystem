package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func GenerateYAMLConfig(config Config, filename string) error {
	data, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	filename := "config.yaml"
	if _, err := ioutil.ReadFile(filename); err != nil {
		config := Config{
			MySQL: MySQLConfig{
				Host:     "127.0.0.1",
				Port:     3306,
				Database: "akatsuki",
				Username: "ticketadmin",
				Password: "password",
			},
			ServicePort: 8080,
		}
		if err := GenerateYAMLConfig(config, filename); err != nil {
			log.Fatalf("generating config file error: %v", err)
		}
		log.Println("config file has been generated")
		os.Exit(1)
	}
}

func LoadConfig() *Config {
	cfg := new(Config)
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("read config file error: %v", err)
		return nil
	}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		log.Fatalf("yaml unmarshal error: %v", err)
		return nil
	}

	return cfg
}
