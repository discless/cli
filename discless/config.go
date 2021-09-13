package discless

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	Host	= "localhost"
	Port	= "8443"
)

type Config struct {
	Host	string	`yaml:"host"`
	Port	string	`yaml:"port"`
	SkipTLS	bool	`yaml:"skiptls"`
	Auth struct{
		Keys	map[string]string `yaml:"keys"`
	}	`yaml:"keys"`
}

func creatConfig() (error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dir := homedir + "/.config/discless"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir,0777)
	}

	config := &Config{
		Host:    "localhost",
		Port:    "8443",
		SkipTLS: true,
	}

	out, err := yaml.Marshal(config)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dir + "/config.yaml", out, 0664)

	if err != nil {
		return err
	}

	return nil
}

func configExists() bool {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return false
	}

	dir := homedir + "/.config/discless/config.yaml"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetConfig() (*Config, error) {
	if !configExists() {
		err := creatConfig()

		if err != nil {
			return nil, err
		}
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	out, err := ioutil.ReadFile(homedir + "/.config/discless/config.yaml")

	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = yaml.Unmarshal(out, config)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func SetConfig(config *Config) error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dir := homedir + "/.config/discless/config.yaml"

	byt, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dir, byt, 0644)

	if err != nil {
		return err
	}

	return nil
}

func AddKey(bot, key string) error {
	config, err := GetConfig()

	if err != nil {
		return err
	}

	config.Auth.Keys[bot] = key

	err = SetConfig(config)

	if err != nil {
		return err
	}

	return nil
}

func GetKey(bot string) (string, error) {
	config, err := GetConfig()
	if val, ok := config.Auth.Keys[bot]; !ok {
		return "", errors.New("couldn't find key for bot \"" + bot + "\"")
	} else {
		return val, err
	}
}