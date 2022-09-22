package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/SimonBerghem/mobile-distributed-system/src/d7024e"
	"github.com/SimonBerghem/mobile-distributed-system/src/kademlia_cli"
)

func main() {

	kademlia_cli.InitCLI()

	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := d7024e.NewContact(d7024e.NewRandomKademliaID(), defaultIP+":"+strconv.Itoa(port))

	ip := GetOutboundIP()

	// Create node
	nodeID := d7024e.NewRandomKademliaID()
	me := d7024e.NewContact(nodeID, ip)
	routing := d7024e.NewRoutingTable(me)
	node := d7024e.NewKademlia(routing)
	network := d7024e.NewNetwork(node)

	// fmt.Println(defaultIP != ip, " ", ip)

	if defaultIP != ip {
		port = 4001
	}

	go network.Listen(ip, port)
	time.Sleep(5 * time.Second)
	// fmt.Println(network, defaultCon)
	fmt.Println("yo")

	// Add node to network
	network.SendPingMessage(&defaultCon)

	// Node listens for requests
}

// Get preferred outbound ip of this machine
// Taken from https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
