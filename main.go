package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	IpSets struct {
		IpSets map[string]string `mapstructure:"ipsets"`
	} `mapstructure:"ip_set_config"`
}

func NewConfig(configFilePath string) (*Config, error) {
	config := &Config{}
	err := config.parseConfig(configFilePath)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Config) parseConfig(path string) error {
	t := make(map[string]interface{})

	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fileData, &t)
	if err != nil {
		return err
	}

	err = mapstructure.Decode(t, c)
	if err != nil {
		return err
	}

	return nil
}

func initConfig() (*Config, error) {
	return NewConfig("./config.yaml")
}

func main() {
	cnf, err := initConfig()
	if err != nil {
		panic(err)
	}
	ipsets := cnf.IpSets.IpSets
	for k,v := range ipsets {
		fmt.Printf("%s -> %s\n", k, v)
	}
}