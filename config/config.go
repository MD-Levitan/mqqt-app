package config

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

var config *Config

type DBConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Database string `yaml:"database"`
	Password string `yaml:"-"`
	SSLMode  string `yaml:"ssl_mode"`
}

type WebConfig struct {
	SessionKey string `yaml:"-"`
	JWTSecret  string `yaml:"-"`
}

type Config struct {
	DB  DBConfig  `yaml:"db"`
	Web WebConfig `yaml:"-"`
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
	if config.Web.JWTSecret, err = ReadSecret(appPath, "jwt_secret"); err != nil {
		return config, err
	}
	return config, err
}

func GetConfig() *Config {
	return config
}

func InitConfig(configName string) (*Config, error) {
	var err error
	if config == nil {
		var applicationPath string
		applicationPath, err = os.Getwd()
		config, err = makeConfig(applicationPath, configName)
	}
	return config, err
}

func ReadSecret(secretPath string, secretName string) (string, error) {
	if data, err := ioutil.ReadFile(path.Join(secretPath, secretName)); err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

//func Test()  {
//	config := Config{DB:DBConfig{User:"user", Host:"host", SSLMode:"false", Database:"db"}}
//	res, err := yaml.Marshal(&config)
//	fmt.Printf("%s", res)
//	fmt.Printf("%s", err)
//}
