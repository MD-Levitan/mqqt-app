package config

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"

	"github.com/MD-Levitan/bboltstore/store"
	"go.etcd.io/bbolt"
	"gopkg.in/yaml.v2"
)

var config *Config
var session_store *store.Store
var db *bbolt.DB
var tmpl *template.Template

type DBConfig struct {
	Database string `yaml:"file"`
}

type WebConfig struct {
	SessionKey   string `yaml:"-"`
	CipherSecret string `yaml:"-"`
}

type MQQTConfig struct {
	Protocol      string `yaml:"protocol"`
	IP            string `yaml:"ip"`
	Port          uint16 `yaml:"port"`
	MultipleUsers bool   `yaml:"multiple"`
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

func GetStore() *store.Store {
	return session_store
}

func InitStore() error {

	if config == nil {
		return fmt.Errorf("congfig is not init")
	}

	/* TODO: Change second key */
	var err error
	var store_config = store.NewDefaultConfig()
	store_config.SessionOptions.MaxAge = 60 * 30 // 30 minutes
	store_config.DBOptions.FreeDB = true
	store_config.ReaperOptions.StartRoutine = true

	session_store, err = store.NewStoreWithDB(config.DB.Database, *store_config, []byte(config.Web.SessionKey), []byte(config.Web.SessionKey))
	if err != nil {
		fmt.Printf("1")
		return err
	}

	//sessions.NewFilesystemStore(config.DB.Database, []byte(config.Web.SessionKey), []byte(config.Web.SessionKey))
	return nil
}

func InitTmpl() error {
	if t, err := template.ParseGlob("templates/*.html"); err != nil {
		return err
	} else {
		tmpl = t
		return nil
	}
}

func GetTmpl() *template.Template {
	return tmpl
}
