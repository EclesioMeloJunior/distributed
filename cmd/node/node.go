package main

import "log"

func setupNode(c *Config) error {
	log.Printf("Node started at %s:%v", c.Network.IP, c.Network.Port)
	return nil
}
