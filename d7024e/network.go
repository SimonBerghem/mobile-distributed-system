package d7024e

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
)

type Network struct {
}

type Protocol struct {
	Rpc string 			// PING, STORE, FIND_NODE, FIND_VALUE 
	Contacts []Contact 	
	Data []byte 		
	Sender Contact
	Target Contact
}

func NewNetwork() *Network {
	return &Network{}
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

	defer conn.Close()

	fmt.Println("Listening on " + addrStr)

	// i := 0
	for {
		// if node.routing.me.Address == "172.20.0.2:4000"{
		// 	time.Sleep(20 * time.Second)
		// 	fmt.Println("CONN: ", conn)
		// }
		// 		fmt.Println("Looping: ", i)
		// 		i = i + 1
		// 	}
		// defer conn.Close()
		network.HandleConn(conn, node)
	}
}

// Check which message has been recevied and handle it accordingly
func (network *Network) HandleConn(conn *net.UDPConn, node *Kademlia) {
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
		response = network.handlePingMessage(proto, node)
	case "STORE":
		response = network.handleStoreMessage(proto, node)
	case "FIND_NODE":
		response = network.handleFindContactMessage(proto, node)
	case "FIND_VALUE":
		fmt.Println("FIND_VALUE")
	default:
		fmt.Println("Unknown RPC: ", rpc)
	}

	time.Sleep(5 * time.Second)
	_, err = conn.WriteToUDP(response, addr)
	if err != nil {
		fmt.Println(err)
	}
}

// PING
func (network *Network) SendPingMessage(contact *Contact, node *Kademlia) bool {

	conn, err := net.Dial("udp4", contact.Address)
	if err != nil {
		fmt.Println(err)
		// panic(err)
		return false
	}

	msg := network.createPingMessage(*contact, node)
	conn.Write(msg)
	time.Sleep(5 * time.Second)

	buf := make([]byte, 1024)
	rlen, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		// panic(err)
		return false
	}
	message := buf[:rlen]

	var proto Protocol

	err = json.Unmarshal(message, &proto)
	if err != nil {
		fmt.Println(err)
		// panic(err)
		return false
	}

	node.routing.AddContact(proto.Sender)
	return true
}

// FIND_NODE
func (network *Network) SendFindContactMessage(contact *Contact, target *KademliaID, node *Kademlia) []Contact {

	conn, err := net.Dial("udp4", contact.Address)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}

	msg := network.createFindContactMessage(*contact, target, nil, node)
	conn.Write(msg)

	buf := make([]byte, 8192)

	// conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	rlen, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	message := buf[:rlen]
	var proto Protocol

	// fmt.Println("MSG: ", string(message))

	err = json.Unmarshal(message, &proto)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}

	// network.addContacts(proto.Contacts, node)
	if len(proto.Contacts) > 0 {
		node.routing.AddContact(proto.Sender)
	}
	return proto.Contacts
}

// FIND_VALUE
func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

// STORE
func (network *Network) SendStoreMessage(contact *Contact, data []byte, node *Kademlia) {

	conn, err := net.Dial("udp4", contact.Address)
	if err != nil {
		fmt.Println(err)
	}

	originNode := make([]Contact, 0)
	originNode = append(originNode, node.routing.me)


	msg := node.network.createStoreMessage(originNode, data, *contact, node)
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

}

// Handles a recieved PING protocol
func (network *Network) handlePingMessage(proto Protocol, node *Kademlia) []byte {
	// Add to routing
	// network.addContacts(proto.Contacts, node)
	node.routing.AddContact(proto.Sender)

	// Send back message with my ip 
	return network.createPingMessage(proto.Target, node)
}

func (network *Network) handleFindContactMessage(proto Protocol, node *Kademlia) []byte {

	// network.addContacts(proto.Contacts, node)
	node.routing.AddContact(proto.Sender)
	// target := NewKademliaID(proto.target)
	contacts := node.routing.FindClosestContacts(proto.Target.ID, bucketSize)
	return network.createFindContactMessage(proto.Target, NewKademliaID(string(proto.Data)), contacts, node)
}

func (network *Network) handleFindDataMessage() []byte {
	return make([]byte, 1024)
}

func (network *Network) handleStoreMessage(proto Protocol, node *Kademlia) []byte {
	/*
		network.addContacts(proto.Contacts, node)
		target := NewKademliaID(string(proto.Data))
		contacts := node.routing.FindClosestContacts(target, bucketSize)
	*/
	node.Store(proto.Data)

	return make([]byte, 1024)
}

// Creates a byte array containing a PING protocol
func (network *Network) createPingMessage(target Contact, node *Kademlia) []byte {
	return CreateProtocol("PING", nil, nil, node.routing.me, target)
}

// Creates a byte array containing a FIND_NODE protocol
func (network *Network) createFindContactMessage(contact Contact, target *KademliaID, contacts []Contact, node *Kademlia) []byte {
	return CreateProtocol("FIND_NODE", contacts, []byte(target.String()), node.routing.me, contact)
	
}

func (network *Network) createFindDataMessage() []byte {
	return make([]byte, 1024)
}

func(network *Network) createStoreMessage(contacts []Contact, data []byte, target Contact, node *Kademlia) []byte {
	return CreateProtocol("STORE", contacts, data, node.routing.me, target)
}

func (network *Network) addContacts(contacts []Contact, node *Kademlia) {
	for _, con := range contacts {
		node.routing.AddContact(con)
	}
}

func CreateProtocol(rpc string, contacts []Contact, data []byte, sender Contact, target Contact) []byte{
	protocol, err := 
		json.Marshal( 
			&Protocol{
				Rpc:      rpc,
				Contacts: contacts,
				Data: data,
				Sender: sender,
				Target: target})

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return protocol
}
