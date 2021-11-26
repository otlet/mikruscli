package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ConfigFilePath string
	CurrentCtx     int       `yaml:"currentctx"`
	Ctx            []Context `yaml:"ctx"`
}

type Context struct {
	Name       string `yaml:"name"`
	ServerName string `yaml:"server_name"`
	Token      string `yaml:"token"`
}

func (c *Config) init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	c.ConfigFilePath = fmt.Sprintf("%v/.mikruscli.yml", dirname)
	c.read()
}

func (c *Config) read() error {
	f, err := os.Open(c.ConfigFilePath)
	if err != nil {
		log.Println(err)
		log.Println("Trying to create empy config file")
		data, err := yaml.Marshal(&c)
		if err != nil {
			return err
		}

		err2 := ioutil.WriteFile(c.ConfigFilePath, data, 0600)
		if err2 != nil {
			return err2
		}
		log.Println("Success")
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) Save() error {
	f, err := os.Create(c.ConfigFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := yaml.NewEncoder(f)
	err = encoder.Encode(&config)
	if err != nil {
		return err
	}

	return nil
}
