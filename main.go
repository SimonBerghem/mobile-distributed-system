package main

import (
	"io"
	"os"

	"github.com/SimonBerghem/mobile-distributed-system/cmd"
	"github.com/SimonBerghem/mobile-distributed-system/d7024e"
)

var out io.Writer = os.Stdout

func main() {
	cmd.InitCLI(out, os.Args)
	node := d7024e.Kademlia{}
	node.InitNode()
}
