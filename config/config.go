package config

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/sessions"
	"gopkg.in/yaml.v2"
)

var config *Config
var store *sessions.FilesystemStore

type DBConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Database string `yaml:"database"`
	Password string `yaml:"-"`
	SSLMode  string `yaml:"ssl_mode"`
}

type WebConfig struct {
	SessionKey   string `yaml:"-"`
	CipherSecret string `yaml:"-"`
}

type MQQTConfig struct {
	Protocol string `yaml:"protocol"`
	IP       string `yaml:"ip"`
	Port     uint16 `yaml:"port"`
}

type Config struct {
	DB   DBConfig   `yaml:"db"`
	Web  WebConfig  `yaml:"-"`
	MQQT MQQTConfig `yaml:"mqqt"`
}

func makeConfig(appPath string, configName string) (*Config, error) {
	data, err := ioutil.ReadFile(path.Join(appPath, configName))
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	if config.DB.Password, err = ReadSecret(appPath, "db_password"); err != nil {
		return config, err
	}
	if config.Web.SessionKey, err = ReadSecret(appPath, "session_key"); err != nil {
		return config, err
	}
	if config.Web.CipherSecret, err = ReadSecret(appPath, "cipher_key"); err != nil {
		return config, err
	}
	return config, err
}

func GetConfig() *Config {
	return config
}

func InitConfig(configName string) error {
	var err error
	if config == nil {
		var applicationPath string
		applicationPath, err = os.Getwd()
		config, err = makeConfig(applicationPath, configName)
	}
	return err
}

func ReadSecret(secretPath string, secretName string) (string, error) {
	if data, err := ioutil.ReadFile(path.Join(secretPath, secretName)); err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

func GetStore() *sessions.FilesystemStore {
	return store
}

func InitStore() error {

	if config == nil {
		return fmt.Errorf("congfig is not init")
	}
	/* TODO: Change second key */
	store = sessions.NewFilesystemStore("./session", []byte(config.Web.SessionKey), []byte(config.Web.SessionKey))

	store.Options = &sessions.Options{
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   3600,
	}
	return nil
}
