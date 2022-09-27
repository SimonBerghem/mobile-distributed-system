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
	network *Network
}

func (kademlia *Kademlia) InitNode() {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewRandomKademliaID(), defaultIP+":"+strconv.Itoa(port))

	ip := GetOutboundIP()

	// Create node
	nodeID := NewRandomKademliaID()
	me := NewContact(nodeID, ip)
	routing := NewRoutingTable(me)
	network := NewNetwork()
	node := NewKademlia(routing, network)

	// fmt.Println(defaultIP != ip, " ", ip)

	if defaultIP != ip {
		rand.Seed(time.Now().UnixNano())
		port = rand.Intn(65535-1000) + 1000
	}

	go network.Listen(ip, port, node)
	time.Sleep(5 * time.Second)
	// fmt.Println(network, defaultCon)

	// Add node to network
	network.SendFindContactMessage(&defaultCon, nodeID, node)
	fmt.Println("Contacts: ", node.routing.buckets[159].Len())
	node.LookupContact(&node.routing.me)

	// go update()
	for {

	}
}

func update() {
	for{}
}

func NewKademlia(table *RoutingTable, network *Network) *Kademlia {
	return &Kademlia{table, network}
}

func (kademlia *Kademlia) LookupContact(target *Contact) {

	queryList := kademlia.routing.FindClosestContacts(kademlia.routing.me.ID, bucketSize)
	fmt.Println("Query len: ", len(queryList))
	// Pick alpha closest nodes from it knows

	// Send FIND_NODE to each node
	
	// Resend FIND_NODE to the new nodes
	
	// Nodes not responding quickly are not considered until they answer

	// If FIND_NODE does not return a node closer to already closest, query from non-queried k nodes

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
