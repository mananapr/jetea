package config

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"

	log "github.com/charmbracelet/log"
	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	yamlmarshaller "gopkg.in/yaml.v3"
)

type ViewType string

func (vt ViewType) String() string {
	return string(vt)
}

const (
	ServerSelectionView ViewType = "servers"
	PubSubView          ViewType = "pubsub"
	JetstreamView       ViewType = "jetstream"
	RequestReplyView    ViewType = "requestreply"
)

const (
	DEFAULT_XDG_CONFIG_DIRNAME = ".config"
	JETEA_DIR                  = "jetea"
	CONFIG_FILE_NAME           = "config.yml"
)

var (
	validate *validator.Validate
)

type NATSConfig struct {
	Name              string  `yaml:"name" validate:"required,max=100"`
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
	var config Config
	parser := initParser()

	if configPath == "" {
		configDir := os.Getenv("XDG_CONFIG_HOME")
		if configDir == "" {
			log.Debug("XDG_CONFIG_HOME not set")
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return Config{}, err
			}
			configDir = filepath.Join(homeDir, DEFAULT_XDG_CONFIG_DIRNAME)
		}

		configFilePath := filepath.Join(configDir, JETEA_DIR, CONFIG_FILE_NAME)

		if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
			log.Debug("writing default config", "path", configFilePath)
			if defaultConfErr := parser.createDefaultConfigFile(configFilePath); defaultConfErr != nil {
				return Config{}, defaultConfErr
			}
		} else {
			log.Debug("using default config", "path", configFilePath)
		}

		configPath = configFilePath
	}

	if err := parser.k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
		return Config{}, err
	}

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
				Name:              "Default Server",
				Hostname:          "localhost:4222",
				TLS:               false,
				TimeoutSeconds:    10,
				ReconnectAttempts: 3,
			},
		},
	}
}

func (parser ConfigParser) createDefaultConfigFile(configFilePath string) error {
	if err := os.MkdirAll(filepath.Dir(configFilePath), os.ModePerm); err != nil {
		return err
	}

	defaultConfig := parser.defaultConfig()
	content, err := yamlmarshaller.Marshal(defaultConfig)
	if err != nil {
		return err
	}

	return os.WriteFile(configFilePath, content, 0644)
}
