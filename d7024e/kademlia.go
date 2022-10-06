package d7024e

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
	"sync"
)

// Stores routing table
type Kademlia struct {
	routing *RoutingTable
	network *Network
	data    map[KademliaID][]byte
}

const alpha = 3

func (kademlia *Kademlia) InitNode() {
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

	go network.Listen(ip, port, node)

	// fmt.Println(network, defaultCon)
	contactedGateway := false

	if defaultIP != ip {
		node.routing.AddContact(defaultCon)
		contactedGateway = node.network.SendPingMessage(&defaultCon, node)
	}

	 //Add node to network
	if contactedGateway {
		nodes := node.LookupContacts(node.routing.me.ID)
		fmt.Println("Found ", len(nodes), " nodes!")
		// fmt.Printf("\n\nEmpty map\n%s\n", node.data)
		d1 := []byte("AAAAA")
		// fmt.Printf("Adding %s with hash: %s\n", d1, Hash(d1))
		network.SendStoreMessage(&con, d1, node)

		fmt.Printf("%s\n", node.data)
		fmt.Printf("Adding %s\n", []byte("123456789"))
		network.SendStoreMessage(&con, []byte("123456789"), node)

		network.SendStoreMessage(&defaultCon, []byte("AAAAA"), node)
		//network.handleStoreMessage(network.createStoreMessage(defaultCon, c), node)
		network.SendStoreMessage(&defaultCon, []byte("123456789"), node)

	}

	go update()
	// for {

	// }
}

func update() {
	for {
	}
}

func NewKademlia(table *RoutingTable, network *Network) *Kademlia {
	return &Kademlia{table, network, make(map[KademliaID][]byte)}
}

func (kademlia *Kademlia) LookupContacts(target *KademliaID) []Contact {

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


		if len(unqueried) == 0 {
			return closest
		}

		// Send find_node to the alpha closest unqueried nodes
		kademlia.contactNodes(target, unqueried, foundNodes)
		seen.AppendNoDups(unqueried)
	}
}

func (kademlia *Kademlia) contactNodes(target *KademliaID, queryList []Contact, table *RoutingTable) {
	
	var wg sync.WaitGroup
	wg.Add(len(queryList))
	for _, contact := range queryList{
		go func(contact Contact, target *KademliaID, table *RoutingTable){
			defer wg.Done()
			nodes := kademlia.network.SendFindContactMessage(&contact, target, kademlia)
			table.AddContacts(nodes)

		}(contact, target, table)
	}
	wg.Wait()
}

func contains(list []Contact, target Contact) bool {
	for _, contact := range list {
		if (contact).ID.Equals(target.ID) {
			return true
		}
	}
	return false
}

func (kademlia *Kademlia) LookupData(hash string) []byte {
	return kademlia.data[*NewKademliaID(hash)]
}

func (kademlia *Kademlia) Store(data []byte) string {
	hash := NewKademliaID(string(data))
	contacts := kademlia.LookupContacts(hash)

	var wg sync.WaitGroup
	wg.Add(len(contacts))

	for _, contact := range contacts {
		if contact.ID.Equals(kademlia.routing.me.ID){
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

func (kademlia *Kademlia) StoreValue(data []byte) {
	hash := Hash(data)
	id := NewKademliaID(hash)
	kademlia.data[*id] = data
}

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
		if len(unqueriedNodes) == count{
			break
		} else if !contains(seenNodes, node){
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
