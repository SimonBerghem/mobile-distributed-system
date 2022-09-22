package main

import (
	"github.com/SimonBerghem/mobile-distributed-system/d7024e"
	"github.com/SimonBerghem/mobile-distributed-system/kademlia_cli"
)

func main() {

	kademlia_cli.InitCLI()
	d7024e.InitNode()
}
