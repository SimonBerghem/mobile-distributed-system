package d7024e

import (
	"fmt"
	"encoding/json"
	"encoding/hex"
	"crypto/sha1"
	"net"
	"strconv"
	"time"
)

type Network struct {
	//Data []byte		// det finn en map funktion inbuggd i go som probs skulle vara bäst
	Data map[string]string
}


type Protocol struct {
	Rpc string 			// PING, STORE, FIND_NODE, FIND_VALUE 
	Contacts []Contact 	
	Hash string 		
	Data []byte 		
	Message string 		
   }

func NewNetwork() *Network {
	n:= &Network{}
	n.Data = make(map[string]string)
	return n
}

func (network *Network) Listen(ip string, port int, node *Kademlia) {
	
	addrStr := ip + ":" + strconv.Itoa(port)

	// udp4 only allows IPV4 addresses
	addr, err := net.ResolveUDPAddr("udp4", addrStr)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// defer conn.Close()

	fmt.Println("Listening on " + addrStr)

	// i := 0
	for {
	// if network.node.routing.me.Address == "172.20.0.2"{
	// 	time.Sleep(20 * time.Second)
	// }
	// 		fmt.Println("Looping: ", i)
	// 		i = i + 1 
	// 	}
	// defer conn.Close()
		network.HandleConn(conn, node)
	}
}

// Check which message has been recevied and handle it accordingly
func (network *Network) HandleConn(conn *net.UDPConn, node *Kademlia){
	buf := make([]byte, 1024)
	rlen, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	values := buf[:rlen]

	var proto Protocol
	var response []byte
	
	if err := json.Unmarshal(values, &proto); err != nil {
		fmt.Println(err)
		panic(err)
	}
	// fmt.Println("HMMM: ", proto.Contacts)

	switch rpc := proto.Rpc; rpc {

	case "PING":
		response = network.handlePingMessage(proto, node)
	case "STORE":
		response = network.handleStoreMessage(proto, node)
	case "FIND_NODE":
		response = network.handleFindContactMessage(proto, node)
	case "FIND_VALUE":
		fmt.Println("FIND_VALUE")
	default:
		fmt.Println("Unknown RPC")
	}

	time.Sleep(5 * time.Second)
	_, err = conn.WriteToUDP(response, addr)
	if err != nil {
		fmt.Println(err)
	}
}

// PING
func (network *Network) SendPingMessage(contact *Contact, node *Kademlia) {

	conn, err := net.Dial("udp4", contact.Address)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}

	msg := network.createPingMessage(node)
	conn.Write(msg)
	time.Sleep(5 * time.Second)

	buf := make([]byte, 1024)
	rlen, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	message := buf[:rlen]

	var proto Protocol

	err = json.Unmarshal(message, &proto)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}

	node.routing.AddContact(proto.Contacts[0])
}

// FIND_NODE
func (network *Network) SendFindContactMessage(contact *Contact, target *KademliaID, node *Kademlia) []Contact {
	
	conn, err := net.Dial("udp4", contact.Address)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	originNode := make([]Contact, 0)
	originNode = append(originNode, node.routing.me)

	msg := network.createFindContactMessage(target, originNode)
	conn.Write(msg)

	buf := make([]byte, 1024)
	rlen, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	message := buf[:rlen]
	var proto Protocol

	err = json.Unmarshal(message, &proto)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}

	network.addContacts(proto.Contacts, node)
	if len(proto.Contacts) > 0{
		node.routing.AddContact(proto.Contacts[0])
	}
	return proto.Contacts
}

// FIND_VALUE
func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}


// STORE
// verkar typ funka men är inte 100%
func (network *Network) SendStoreMessage(contact *Contact, data []byte, node *Kademlia) {

	conn, err := net.Dial("udp4", contact.Address)
	if err != nil {
		fmt.Println(err)
	}

	originNode := make([]Contact, 0)
	originNode = append(originNode, node.routing.me)


	msg := node.network.createStoreMessage(originNode, data)
	conn.Write(msg)
	//time.Sleep(5 * time.Second)

	//buf := make([]byte, 1024)
	
	//rlen, err := conn.Read(buf)

	//if err != nil {
	//	fmt.Println(err)
	//}
	//message := buf[:rlen]


	//var proto Protocol

	//err = json.Unmarshal(message, &proto)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//node.routing.AddContact(proto.Contacts[0])
}




// Handles a recieved PING protocol
func (network *Network) handlePingMessage(proto Protocol, node *Kademlia) []byte {
	// Add to routing
	network.addContacts(proto.Contacts, node)

	// Send back message with my ip 
	return network.createPingMessage(node)
}

func (network *Network) handleFindContactMessage(proto Protocol, node *Kademlia) []byte {

	network.addContacts(proto.Contacts, node)
	target := NewKademliaID(string(proto.Data))
	contacts := node.routing.FindClosestContacts(target, bucketSize)
	return network.createFindContactMessage(target, contacts)
}

func(network *Network) handleFindDataMessage() []byte {
	return make([]byte, 1024)
}

func(network *Network) handleStoreMessage(proto Protocol, node *Kademlia) []byte {

	// Hash för att skapa nyckel
	hashbytes := sha1.Sum(proto.Data)
    hash := hex.EncodeToString(hashbytes[0:IDLength])

	
	//network. = proto.Data

	fmt.Println(network.Data)
	fmt.Println("")
	network.Data[string(hash)] = string(proto.Data)
	fmt.Println(network.Data[string(hash)])

	return proto.Data
}

// Creates a byte array containing a PING protocol
func (network *Network) createPingMessage(node *Kademlia) []byte {
	originNode := make([]Contact, 0)
	originNode = append(originNode, node.routing.me)
	return CreateProtocol("PING", originNode, "", nil, "")
}

// Creates a byte array containing a FIND_NODE protocol
func (network *Network) createFindContactMessage(target *KademliaID, contacts []Contact) []byte {
	return CreateProtocol("FIND_NODE", contacts, "", []byte(target.String()), "")
	
}

func(network *Network) createFindDataMessage() []byte {
	return make([]byte, 1024)
}

func(network *Network) createStoreMessage(contacts []Contact, data []byte) []byte {
	return CreateProtocol("STORE", contacts, "", data, "")
}


func (network *Network) addContacts (contacts []Contact, node *Kademlia) {
	for _, con := range contacts{
		node.routing.AddContact(con)
	}
}

func CreateProtocol(rpc string, contacts []Contact, hash string, data []byte, msg string) []byte{
	protocol, err := 
		json.Marshal( 
			&Protocol{
				Rpc: rpc,
				Contacts: contacts,
				Hash: hash,
				Data: data,
				Message: msg})

	if err != nil{
		fmt.Println(err)
		return nil
	}

	return protocol
}


