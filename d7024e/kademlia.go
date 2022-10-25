package d7024e

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
	// "fmt"
)

// Stores routing table
type Kademlia struct {
	routing *RoutingTable
	network *Network
	data    map[KademliaID][]byte
}

const alpha = 3

func (kademlia *Kademlia) InitNode(ch chan Kademlia) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))

	ip := GetOutboundIP()

	var con Contact

	if defaultIP != ip {
		rand.Seed(time.Now().UnixNano())
		port = rand.Intn(65535-1000) + 1000
		nodeID := NewRandomKademliaID()
		con = NewContact(nodeID, ip+":"+strconv.Itoa(port))
	} else {
		con = defaultCon
	}

	// Create node
	routing := NewRoutingTable(con)
	network := NewNetwork()
	node := NewKademlia(routing, network)

	ch <- *node
	go network.Listen(ip, port, node)

	contactedGateway := false

	// Ping default node
	if defaultIP != ip {
		node.routing.AddContact(defaultCon)
		contactedGateway = node.network.SendPingMessage(&defaultCon, node)
	}

	// var nodes []Contact

	// Connect to the rest of the network
	if contactedGateway {
		// nodes = node.LookupContacts(node.routing.me.ID)
		node.LookupContacts(node.routing.me.ID)
	}

	// if len(nodes) > 0 {
	// 	node.Store([]byte("TEST"))
	// 	fmt.Println("string:", string(node.LookupData(Hash([]byte("TEST")))))
	// }

}

func NewKademlia(table *RoutingTable, network *Network) *Kademlia {
	return &Kademlia{table, network, make(map[KademliaID][]byte)}
}

// Lookup k closest nodes to target in network, k = bucketSize
func (kademlia *Kademlia) LookupContacts(target *KademliaID) []Contact {

	data := make(chan []byte)
	var seen ContactCandidates
	foundNodes := NewRoutingTable(kademlia.routing.me)

	// Find k initial closest nodes
	initClosest := kademlia.routing.FindClosestContacts(target, bucketSize)
	foundNodes.AddContacts(initClosest)

	for {

		// Find k currently closest nodes
		closest := foundNodes.FindClosestContacts(target, bucketSize)

		// Pick alpha nodes from k closest that have not been queried
		alphaNodes := min(alpha, len(closest))
		unqueried := findUnqueriedNodes(closest, seen.contacts, alphaNodes)

		// Check if all k closest have been queried
		if len(unqueried) == 0 {
			return closest
		}

		// Send find_node to the alpha closest unqueried nodes
		kademlia.contactNodes(target, unqueried, foundNodes, data, kademlia.network.SendFindContactMessage)
		seen.AppendNoDups(unqueried)
	}
}

// Contact nodes in querylist with contact rpc, FIND_NODE or FIND_VALUE
// FIND_NODE: adds all found nodes in candidate routing table
// FIND_VALUE: if target data is found return the data else return found nodes
func (kademlia *Kademlia) contactNodes(target *KademliaID, queryList []Contact, table *RoutingTable, ch chan []byte, ContactMessage func(*Contact, *KademliaID, *Kademlia) Protocol) {

	var data []byte
	var wg sync.WaitGroup

	wg.Add(len(queryList))
	for _, contact := range queryList {
		go func(contact Contact, target *KademliaID, table *RoutingTable) {
			defer wg.Done()
			proto := ContactMessage(&contact, target, kademlia)
			if proto.Rpc == "DATA" {
				data = proto.Data
				return
			}
			table.AddContacts(proto.Contacts)

		}(contact, target, table)
	}
	wg.Wait()
	if len(data) > 0 {
		ch <- data
	}
}

// Contains function for Contact slice
func contains(list []Contact, target Contact) bool {
	for _, contact := range list {
		if (contact).ID.Equals(target.ID) {
			return true
		}
	}
	return false
}

// Returns value stored in hash if it exists in network
// Else returns empty byte slice
func (kademlia *Kademlia) LookupData(hash string) string {
	target := NewKademliaID(hash)

	nodeData := kademlia.CheckValue(*target)

	if len(nodeData) > 0 {
		return string(nodeData)
	}

	data := make(chan []byte)
	var seen ContactCandidates
	foundNodes := NewRoutingTable(kademlia.routing.me)

	// Find k initial closest nodes
	initClosest := kademlia.routing.FindClosestContacts(target, bucketSize)
	foundNodes.AddContacts(initClosest)

	for {
		// Find k currently closest nodes
		closest := foundNodes.FindClosestContacts(target, bucketSize)

		// Check if all k closest have been queried
		alphaNodes := min(alpha, len(closest))
		unqueried := findUnqueriedNodes(closest, seen.contacts, alphaNodes)

		// Send find_value to the alpha closest unqueried nodes
		go kademlia.contactNodes(target, unqueried, foundNodes, data, kademlia.network.SendFindDataMessage)
		seen.AppendNoDups(unqueried)

		msg := <-data

		// All k closest have been queried or data is found
		if len(unqueried) == 0 || len(msg) > 0 {
			return string(msg)
		}
	}
}

// Store data at k closest nodes in network
func (kademlia *Kademlia) Store(data []byte) string {
	hash := NewKademliaID(Hash(data))
	contacts := kademlia.LookupContacts(hash)

	var wg sync.WaitGroup
	wg.Add(len(contacts))

	for _, contact := range contacts {
		if contact.ID.Equals(kademlia.routing.me.ID) {
			kademlia.StoreValue(data)
			continue
		}
		go func(contact Contact) {
			defer wg.Done()
			kademlia.network.SendStoreMessage(&contact, data, kademlia)
		}(contact)
	}
	wg.Wait()
	return hash.String()
}

// Store data in current node
func (kademlia *Kademlia) StoreValue(data []byte) {
	hash := Hash(data)
	id := NewKademliaID(hash)
	kademlia.data[*id] = data
}

// Return data if in current node
func (kademlia *Kademlia) CheckValue(id KademliaID) []byte {
	return kademlia.data[id]
}

// Hash data and set size to IDLength
func Hash(data []byte) string {
	hashbytes := sha1.Sum(data)
	return hex.EncodeToString(hashbytes[0:IDLength])
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// Returns all unqueried nodes, up to count nodes
func findUnqueriedNodes(closestNodes []Contact, seenNodes []Contact, count int) []Contact {
	var unqueriedNodes []Contact

	for _, node := range closestNodes {
		if len(unqueriedNodes) == count {
			break
		} else if !contains(seenNodes, node) {
			unqueriedNodes = append(unqueriedNodes, node)
		}
	}
	return unqueriedNodes
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
