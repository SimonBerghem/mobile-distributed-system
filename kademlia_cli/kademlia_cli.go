package kademlia_cli

import (
	"fmt"
	"log"
	"os"
	// "strconv

	// "github.com/SimonBerghem/mobile-distributed-system/d7024e"
	"github.com/urfave/cli"
)

func InitCLI() {
	app := cli.NewApp()
	app.Name = "kademlia_cli"
	app.Usage = "A CLI for running Kademlia commands"
	app.Author = "Caspesr Lundberg, Simon Malmström Berghem & Emil Wiklander"
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

func PingTest() {

	// ip1 := "172.21.21.21"
	// ip2 := "172.12.12.12"

	// port1 := 2121
	// port2 := 1212

	// c1 := d7024e.NewContact(d7024e.NewRandomKademliaID(), ip1+":"+strconv.Itoa(port1))
	// c2 := d7024e.NewContact(d7024e.NewRandomKademliaID(), ip2+":"+strconv.Itoa(port2))

	// n1 := d7024e.NewNetwork(d7024e.NewKademlia(d7024e.NewRoutingTable(c1)))
	// n2 := d7024e.NewNetwork(d7024e.NewKademlia(d7024e.NewRoutingTable(c2)))

	// ans1 := make(chan string)
	// ans2 := make(chan string)

	// fmt.Println("%s is pinging: %s", ip1, ip2)
	// fmt.Println("%s is pinging: %s", ip2, ip1)

	// go n2.Listen(ip2, port2)
	// go n1.Listen(ip1, port1)

	// go func() { ans1 <- n1.SendPingMessage(&c1) }()
	// go func() { ans2 <- n2.SendPingMessage(&c2) }()

	// fmt.Println("%s has responded with: ", ip2, <-ans1)
	// fmt.Println("%s has responded with: ", ip1, <-ans2)

}
