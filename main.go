package main

import (
	"io"
	"os"

	"github.com/SimonBerghem/mobile-distributed-system/cmd"
	"github.com/SimonBerghem/mobile-distributed-system/d7024e"
)

var output io.Writer = os.Stdout

func main() {
	node := d7024e.Kademlia{}
	node.InitNode()
	cmd.InitCLI(output, node)
}
