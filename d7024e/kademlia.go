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

const alpha = 3

func (kademlia *Kademlia) InitNode() {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewRandomKademliaID(), defaultIP+":"+strconv.Itoa(port))

	ip := GetOutboundIP()

	if defaultIP != ip {
		rand.Seed(time.Now().UnixNano())
		port = rand.Intn(65535-1000) + 1000
	}

	// Create node
	nodeID := NewRandomKademliaID()
	me := NewContact(nodeID, ip+":"+strconv.Itoa(port))
	routing := NewRoutingTable(me)
	network := NewNetwork()
	node := NewKademlia(routing, network)

	go network.Listen(ip, port, node)
	time.Sleep(5 * time.Second)
	// fmt.Println(network, defaultCon)

	// Add node to network
	node.routing.AddContact(defaultCon)
	// network.SendFindContactMessage(&defaultCon, nodeID, node)
	// fmt.Println("Contacts: ", node.routing.buckets[159].Len())
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

func (kademlia *Kademlia) LookupContact(target *Contact) []Contact {

	var seen ContactCandidates
	var candidates ContactCandidates

	// Pick alpha closest nodes from it knows
	queryList := kademlia.routing.FindClosestContacts(kademlia.routing.me.ID, alpha)
	currentClosest := queryList[0].ID

	fmt.Println("LEN: ", len(queryList))

	for len(queryList) < 1 {}

	for seen.Len() < bucketSize {
		if len(queryList) == 0 {
			fmt.Println("BREAK")
			break
		}

		for i := 0; i < 3 && i < len(queryList); i++ {
			fmt.Println("QUERY: ", queryList)
			candidates.Append(kademlia.lookupContactHelper(seen, currentClosest, target.ID, &queryList[i]))
		}

		// Remove queried and add to seen
		if len(queryList) - 1 < 3 {
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
				if !contains(seen.contacts, candidate) && !contains(queryList, candidate){
					queryList = append(queryList, candidate)
				}
			}
		}
	}

	seen.Sort()
	return seen.contacts
}

func (kademlia *Kademlia) lookupContactHelper(seen ContactCandidates, currentClosest *KademliaID, target *KademliaID, contact *Contact) []Contact {
	fmt.Println("ARGS: ", contact, " ", target)
	candidates := kademlia.network.SendFindContactMessage(contact, target, kademlia)
	fmt.Println("CANDIDATES: ", candidates)
	if len(candidates) > 0 && candidates[0].ID.Less(currentClosest){
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
