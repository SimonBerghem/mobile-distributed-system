package main

import (
	"github.com/SimonBerghem/mobile-distributed-system/cmd"
	"github.com/SimonBerghem/mobile-distributed-system/d7024e"
)

func main() {
	cmd.InitCLI()
	node := d7024e.Kademlia{}
	node.InitNode()
}
