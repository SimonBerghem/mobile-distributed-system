package d7024e

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

// Stores routing table
type Kademlia struct {
	routing *RoutingTable
}

func InitNode() {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewRandomKademliaID(), defaultIP+":"+strconv.Itoa(port))

	ip := GetOutboundIP()

	// Create node
	nodeID := NewRandomKademliaID()
	me := NewContact(nodeID, ip)
	routing := NewRoutingTable(me)
	node := NewKademlia(routing)
	network := NewNetwork(node)

	// fmt.Println(defaultIP != ip, " ", ip)

	if defaultIP != ip {
		rand.Seed(time.Now().UnixNano())
		port = rand.Intn(65535-1000) + 1000
	}

	go network.Listen(ip, port)
	time.Sleep(5 * time.Second)
	// fmt.Println(network, defaultCon)
	fmt.Println("yo")

	// Add node to network
	network.SendPingMessage(&defaultCon)

	for {

	}
}

func NewKademlia(table *RoutingTable) Kademlia {
	return Kademlia{table}
}

func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO
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
