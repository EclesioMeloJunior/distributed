package main

import (
	"flag"
	"log"
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

	err = setupNode(c)
	if err != nil {
		log.Printf("could not setup node: %v", err)
	}
}
