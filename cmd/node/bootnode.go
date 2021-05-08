package main

import "log"

func setupBootnode(c *Config) error {
	log.Printf("Bootnode started at %s:%v", c.Network.IP, c.Network.Port)
	return nil
}
