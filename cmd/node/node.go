package main

import (
	"fmt"
	"log"
	"net"

	"github.com/EclesioMeloJunior/distributed/network"
	"github.com/EclesioMeloJunior/distributed/node"
)

func setupNode(c *Config) error {
	log.Printf("Node started at %s:%v", c.Network.IP, c.Network.Port)

	var ip net.IP

	if ip = net.ParseIP(c.Network.IP); ip == nil {
		return fmt.Errorf("could not parse IP addr %s", c.Network.IP)
	}

	n, err := node.NewNode(net.IP(c.Network.IP), c.Network.Port, c.Global.Basepath)
	if err != nil {
		return err
	}

	noderpc := node.NewNodeRPC(n)
	rpcserver, err := network.NewServer(noderpc)
	if err != nil {
		return err
	}

	err = rpcserver.Start("", n.Port)
	if err != nil {
		return err
	}

	return nil
}
