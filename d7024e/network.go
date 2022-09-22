package d7024e

import (
	"fmt"
	"encoding/json"
	"net"
	"strconv"
	"time"
)

type Network struct {
	node Kademlia
}

type Protocol struct {
	Rpc string 			// PING, STORE, FIND_NODE, FIND_VALUE 
	Contacts []Contact 	
	Hash string 		
	Data []byte 		
	Message string 		
   }

func NewNetwork(node Kademlia) Network {
	return Network{node}
}

func (network *Network) Listen(ip string, port int) {
	
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

	i := 0
	for {
		if network.node.routing.me.Address == "172.20.0.2"{
			fmt.Println("Looping: ", i)
			i = i + 1 
		}
		network.HandleConn(conn)
	}
}

// Check which message has been recevied and handle it accordingly
func (network *Network) HandleConn(conn *net.UDPConn){
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

	switch rpc := proto.Rpc; rpc {

	case "PING":
		response = network.handlePingMessage(proto)
	case "STORE":
		fmt.Println("STORE")
	case "FIND_NODE":
		fmt.Println("FIND_NODE")
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
func (network *Network) SendPingMessage(contact *Contact) {

	conn, err := net.Dial("udp4", contact.Address)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}

	msg := network.createPingMessage()
	conn.Write(msg)
	time.Sleep(5 * time.Second)

	buf := make([]byte, 1024)
	rlen, err := conn.Read(buf)
	fmt.Println(rlen)
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

	network.addContacts(proto.Contacts)
}

// FIND_NODE
func (network *Network) SendFindContactMessage(contact *Contact, target *KademliaID) {
	
	conn, err := net.Dial("udp4", contact.Address)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	originNode := make([]Contact, 0)
	originNode = append(originNode, network.node.routing.me)


	msg := network.createFindContactMessage(target, originNode)
	conn.Write(msg)

	buf := make([]byte, 1024)
	rlen, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	message := buf[:rlen]
	fmt.Println(message)
}

// FIND_VALUE
func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}


// STORE
func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}


// Handles a recieved PING protocol
func (network *Network) handlePingMessage(proto Protocol) []byte {
	// Add to routing
	network.addContacts(proto.Contacts)

	// Send back message with my ip 
	return network.createPingMessage()
}

func (network *Network) handleFindContactMessage(proto Protocol) []byte {

	network.addContacts(proto.Contacts)
	target := NewKademliaID(string(proto.Data))
	contacts := network.node.routing.FindClosestContacts(target, bucketSize)
	return network.createFindContactMessage(target, contacts)
}

func(network *Network) handleFindDataMessage() []byte {
	return make([]byte, 1024)
}

func(network *Network) handleStoreMessage() []byte {
	return make([]byte, 1024)
}

// Creates a byte array containing a PING protocol
func (network *Network) createPingMessage() []byte {
	originNode := make([]Contact, 0)
	originNode = append(originNode, network.node.routing.me)
	return CreateProtocol("PING", originNode, "", nil, "")
}

// Creates a byte array containing a FIND_NODE protocol
func (network *Network) createFindContactMessage(target *KademliaID, contacts []Contact) []byte {
	return CreateProtocol("FIND_NODE", contacts, "", []byte(target.String()), "")
	
}

func(network *Network) createFindDataMessage() []byte {
	return make([]byte, 1024)
}

func(network *Network) createStoreMessage() []byte {
	return make([]byte, 1024)
}


func (network *Network) addContacts (contacts []Contact) {
	for _, con := range contacts{
		network.node.routing.AddContact(con)
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

