package main

import (
	"github.com/SimonBerghem/mobile-distributed-system/d7024e"
	"github.com/SimonBerghem/mobile-distributed-system/kademlia_cli"
)

func main() {
	node := d7024e.Kademlia{}
	node.InitNode()
	kademlia_cli.InitCLI()
}
