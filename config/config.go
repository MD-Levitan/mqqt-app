package config

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/boltdb/bolt"
	"github.com/gorilla/sessions"
	"github.com/skynet0590/boltstore/store"
	"github.com/yosssi/boltstore/reaper"
	"gopkg.in/yaml.v2"
)

var config *Config
var db *bolt.DB
var session_store *store.Store

type DBConfig struct {
	Database string      `yaml:"file"`
	Password os.FileMode `yaml:"filemod"`
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

func GetDB() *bolt.DB {
	return db
}

func InitDB(path string, mode os.FileMode) error {
	if db_, err := bolt.Open(path, mode, nil); err != nil {
		return err
	} else {
		db = db_

	}
	defer db.Close()
	// Invoke a reaper which checks and removes expired sessions periodically.
	defer reaper.Quit(reaper.Run(db, reaper.Options{}))
}

func GetStore() *store.Store {
	return session_store
}

func InitStore() error {

	if config == nil {
		return fmt.Errorf("congfig is not init")
	}
	/* TODO: Change second key */
	session_store = store.Store(db,
		&sessions.Options{
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			MaxAge:   60},
		[]byte(config.Web.SessionKey))
	return nil
}
