package main

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

type (
	Global struct {
		Basepath string `toml:"basepath"`
	}

	Network struct {
		IP   string `toml:"ip"`
		Port int    `toml:"port"`
	}

	Config struct {
		Global  Global  `toml:"global"`
		Network Network `toml:"network"`
	}
)

// readConfig return the config defined at configpath file
func readConfig(configpath string) (*Config, error) {
	configbytes, err := ioutil.ReadFile(configpath)
	if err != nil {
		return nil, err
	}

	c := new(Config)
	err = toml.Unmarshal(configbytes, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
