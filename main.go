package main

import (
	"io"
	"os"

	"github.com/SimonBerghem/mobile-distributed-system/cmd"
	"github.com/SimonBerghem/mobile-distributed-system/d7024e"
)

var output io.Writer = os.Stdout

func main() {
	ch := make(chan d7024e.Kademlia)

	node := d7024e.Kademlia{}
	go node.InitNode(ch)

	createdNode := <- ch
	cmd.InitCLI(output, createdNode)
}
