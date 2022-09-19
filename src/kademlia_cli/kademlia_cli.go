package kademlia_cli

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli"
)

func KademliaCLI() {
	app := cli.NewApp()
	app.Name = "kademlia_cli"
	app.Usage = "A CLI for running Kademlia commands"
	app.Author = "Casper Lundberg, Simon Malmström Berghem, & Emil Viklander"
	app.Version = "0.0.0"
	app.Commands = []cli.Command{
		{
			Name:    "put",
			Aliases: []string{"p"},
			Usage:   "Takes a single argument, the contents of the file you are uploading, and outputs the hash of the object, if it could be uploaded successfully",
			Action: func(c *cli.Context) {
				fileContents := c.Args().Get(0)
				fmt.Println(fileContents)
			},
		}, {
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Takes a hash as its only argument, and outputs the contents of the object and the node it was retrieved from, if it could be downloaded successfully",
			Action: func(c *cli.Context) {
				hash := c.Args().Get(0)
				fmt.Println(hash)
			},
		}, {
			Name:    "exit",
			Aliases: []string{"e"},
			Usage:   "Terminates the node",
			Action: func(c *cli.Context) {
				// TODO
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
