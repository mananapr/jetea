package config

import (
	"os"
	"reflect"
	"strings"

	"github.com/mananapr/jetea/internal/util"

	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	yamlmarshaller "gopkg.in/yaml.v3"
)

var validate *validator.Validate

type NATSConfig struct {
	Name              *string `yaml:"name,omitempty" validate:"omitempty,max=100"`
	Hostname          string  `yaml:"hostname" validate:"required,hostname_port"`
	User              *string `yaml:"user,omitempty" validate:"omitempty,max=100"`
	Password          *string `yaml:"password,omitempty" validate:"omitempty,max=100"`
	TLS               bool    `yaml:"tls" validate:"omitempty,boolean"`
	TimeoutSeconds    int     `yaml:"timeout_seconds" validate:"omitempty,gte=0"`
	ReconnectAttempts int     `yaml:"reconnect_attempts" validate:"omitempty,gte=0"`
}

type Config struct {
	NATSServers []NATSConfig `yaml:"servers" validate:"omitempty,dive"`
}

type ConfigParser struct {
	k *koanf.Koanf
}

func initParser() ConfigParser {
	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.Split(fld.Tag.Get("yaml"), ",")[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return ConfigParser{
		k: koanf.New("."),
	}
}

func LoadConfig(configPath string) (Config, error) {
	parser := initParser()

	if err := parser.k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
		return Config{}, err
	}

	var config Config
	if err := parser.k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: "yaml"}); err != nil {
		return Config{}, err
	}

	if err := validate.Struct(config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (parser ConfigParser) defaultConfig() Config {
	return Config{
		NATSServers: []NATSConfig{
			{
				Name:              util.StringPtr("Default Server"),
				Hostname:          "localhost:4222",
				TLS:               false,
				TimeoutSeconds:    10,
				ReconnectAttempts: 3,
			},
		},
	}
}

func (parser ConfigParser) createDefaultConfigFile(configFilePath string) error {
	if err := os.MkdirAll(configFilePath, os.ModePerm); err != nil {
		return err
	}

	defaultConfig := parser.defaultConfig()
	content, err := yamlmarshaller.Marshal(defaultConfig)
	if err != nil {
		return err
	}

	return os.WriteFile(configFilePath, content, 0644)
}
