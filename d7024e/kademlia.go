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
)

// Stores routing table
type Kademlia struct {
	routing *RoutingTable
	network *Network
	data    map[string][]byte
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
	time.Sleep(5 * time.Second)
	// fmt.Println(network, defaultCon)
	contactedGateway := false
	if defaultIP != ip {
		contactedGateway = node.network.SendPingMessage(&defaultCon, node)
	}
	//Add node to network
	// node.routing.AddContact(defaultCon)
	// network.SendFindContactMessage(&defaultCon, nodeID, node)
	// fmt.Println("Contacts: ", node.routing.buckets[159].Len())
	if contactedGateway {
		node.LookupContact(&node.routing.me)
	}

	go network.Listen(ip, port, node)
	time.Sleep(5 * time.Second)
	fmt.Println(network, defaultCon)

	//Add node to network
	node.routing.AddContact(defaultCon)
	//network.SendFindContactMessage(&defaultCon, nodeID, node)
	//fmt.Println("Contacts: ", node.routing.buckets[159].Len())
	//node.LookupContact(&node.routing.me)

	fmt.Printf("\n\nEmpty map\n%s\n", node.data)
	d1 := []byte("AAAAA")
	fmt.Printf("Adding %s with hash: %s\n", d1, hash(d1))
	network.SendStoreMessage(&con, d1, node)

	fmt.Printf("%s\n", node.data)
	fmt.Printf("Adding %s\n", []byte("123456789"))
	network.SendStoreMessage(&con, []byte("123456789"), node)

	network.SendStoreMessage(&defaultCon, []byte("AAAAA"), node)
	//network.handleStoreMessage(network.createStoreMessage(defaultCon, c), node)
	network.SendStoreMessage(&defaultCon, []byte("123456789"), node)

	// go update()
	for {

	}
}

func update() {
	for {
	}
}

func NewKademlia(table *RoutingTable, network *Network) *Kademlia {
	return &Kademlia{table, network, make(map[string][]byte)}
}

func (kademlia *Kademlia) LookupContact(target *Contact) []Contact {

	var seen ContactCandidates
	var candidates ContactCandidates

	// Pick alpha closest nodes from it knows
	queryList := kademlia.routing.FindClosestContacts(target.ID, alpha)
	// fmt.Println("INIT QUERY: ", queryList)
	currentClosest := queryList[0].ID

	for len(queryList) < 1 {
	}

	for seen.Len() < bucketSize {
		fmt.Println(seen.Len())
		if len(queryList) == 0 {
			fmt.Println("BREAK")
			break
		}

		for i := 0; i < 3 && i < len(queryList); i++ {
			test := kademlia.lookupContactHelper(seen, currentClosest, target.ID, queryList[i])
			candidates.Append(test)
		}

		// Remove queried and add to seen
		if len(queryList)-1 < 3 {
			seen.Append(queryList)
			queryList = queryList[:0]
		} else {
			seen.Append(queryList[:3])
			queryList = queryList[2:]
		}

		// Set new closest and
		if candidates.Len() > 0 {
			candidates.Sort()
			currentClosest = candidates.contacts[0].ID
			for _, candidate := range candidates.contacts {
				if !contains(seen.contacts, candidate) && !contains(queryList, candidate) {
					queryList = append(queryList, candidate)
				}
			}
		}
	}

	seen.Sort()
	return seen.contacts
}

func (kademlia *Kademlia) lookupContactHelper(seen ContactCandidates, currentClosest *KademliaID, target *KademliaID, contact Contact) []Contact {
	candidates := kademlia.network.SendFindContactMessage(&contact, target, kademlia)
	if len(candidates) > 0 && candidates[0].ID.Less(currentClosest) {
		fmt.Println("FOUND CLOSER CANDIDATES")
		return candidates
	} else {
		fmt.Println("NO CLOSER CANDIDATES")
		return candidates[:0]
	}
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
	return kademlia.data[hash]
}

func (kademlia *Kademlia) Store(data []byte) string {
	hash := hash(data)
	kademlia.data[hash] = data
	return hash
}

func hash(data []byte) string {
	hashbytes := sha1.Sum(data)
	return hex.EncodeToString(hashbytes[0:IDLength])
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
