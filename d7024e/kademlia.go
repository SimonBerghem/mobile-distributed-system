package d7024e

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
<<<<<<< HEAD:d7024e/kademlia.go
=======
	"math/rand"

	"github.com/SimonBerghem/mobile-distributed-system/src/d7024e"
	"github.com/SimonBerghem/mobile-distributed-system/src/kademlia_cli"
>>>>>>> 21fc82c3e921e2aa1dc2fa820ac6be11552b2f56:src/main.go
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

	for{

	}
}

func NewContact(nodeID invalid type, ip string) {
	panic("unimplemented")
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
