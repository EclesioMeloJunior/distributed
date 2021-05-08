package main

import (
	"flag"
)

const defaultConfigPath = "./config.toml"

var (
	c          *Config
	configpath string
)

func init() {
	flag.StringVar(&configpath, "config", defaultConfigPath, "set the path to the toml configuration file")
}

func main() {
	flag.Parse()

	var err error
	if c, err = readConfig(configpath); err != nil {
		panic(err)
	}

	if c.Global.Bootnode {
		setupBootnode(c)
	} else {
		setupNode(c)
	}
}
