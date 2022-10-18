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

	for {
		network.HandleConn(conn, node)
	}
}

// Check which message has been recevied and handle it accordingly
func (network *Network) HandleConn(conn *net.UDPConn, node *Kademlia) {
	buf := make([]byte, 1024)
	rlen, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	values := buf[:rlen]
	proto := Unserialize(values)

	var response []byte

	switch rpc := proto.Rpc; rpc {

	case "PING":
		response = network.handlePingMessage(proto, node)
	case "STORE":
		response = network.handleStoreMessage(proto, node)
	case "FIND_NODE":
		response = network.handleFindContactMessage(proto, node)
	case "FIND_VALUE":
		response = network.handleFindDataMessage(proto, node)
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
		return false
	}

	msg := network.createPingMessage(*contact, node)
	conn.Write(msg)

	buf := make([]byte, 1024)
	message := ReadFromConnection(conn, buf)
	proto := Unserialize(message)

	node.routing.AddContact(proto.Sender)
	return true
}

// FIND_NODE
func (network *Network) SendFindContactMessage(contact *Contact, target *KademliaID, node *Kademlia) Protocol {

	conn := CreateConnection(contact.Address)
	defer conn.Close()

	msg := network.createFindContactMessage(*contact, target, nil, node)
	conn.Write(msg)

	buf := make([]byte, 8192)

	message := ReadFromConnection(conn, buf)
	proto := Unserialize(message)

	node.routing.AddContact(proto.Sender)
	return proto
}

// FIND_VALUE
func (network *Network) SendFindDataMessage(contact *Contact, target *KademliaID, node *Kademlia) Protocol {
	
	conn := CreateConnection(contact.Address)
	defer conn.Close()

	msg := network.createFindDataMessage("FIND_VALUE", nil, []byte(target.String()), *contact, node)
	conn.Write(msg)

	buf := make([]byte, 8192)

	message := ReadFromConnection(conn, buf)
	proto := Unserialize(message)

	node.routing.AddContact(proto.Sender)
	return proto
}

// STORE
func (network *Network) SendStoreMessage(contact *Contact, data []byte, node *Kademlia) {

	conn := CreateConnection(contact.Address)

	msg := node.network.createStoreMessage(data, *contact, node)
	conn.Write(msg)

	buf := make([]byte, 1024)
	message := ReadFromConnection(conn, buf)
	proto := Unserialize(message)

	node.routing.AddContact(proto.Sender)
}

// Handles a recieved PING protocol
func (network *Network) handlePingMessage(proto Protocol, node *Kademlia) []byte {
	node.routing.AddContact(proto.Sender)
	return network.createPingMessage(proto.Target, node)
}

func (network *Network) handleFindContactMessage(proto Protocol, node *Kademlia) []byte {
	node.routing.AddContact(proto.Sender)
	contacts := node.routing.FindClosestContacts(proto.Target.ID, bucketSize)
	return network.createFindContactMessage(proto.Target, NewKademliaID(string(proto.Data)), contacts, node)
}

func (network *Network) handleFindDataMessage(proto Protocol, node *Kademlia) []byte {
	node.routing.AddContact(proto.Sender)
	target := NewKademliaID(string(proto.Data))
	data := node.CheckValue(*target)

	if data != nil {
		return network.createFindDataMessage("DATA", nil, data, proto.Target, node)
	}
	contacts := node.routing.FindClosestContacts(target, bucketSize)
	return network.createFindDataMessage("FIND_VALUE", contacts, []byte(target.String()), proto.Target, node)
}

func (network *Network) handleStoreMessage(proto Protocol, node *Kademlia) []byte {
	node.routing.AddContact(proto.Sender)
	node.StoreValue(proto.Data)
	return network.createStoreMessage(proto.Data, proto.Sender, node)
}

// Creates a byte array containing a PING protocol
func (network *Network) createPingMessage(target Contact, node *Kademlia) []byte {
	return CreateProtocol("PING", nil, nil, node.routing.me, target)
}

// Creates a byte array containing a FIND_NODE protocol
func (network *Network) createFindContactMessage(contact Contact, target *KademliaID, contacts []Contact, node *Kademlia) []byte {
	return CreateProtocol("FIND_NODE", contacts, []byte(target.String()), node.routing.me, contact)
}

func (network *Network) createFindDataMessage(rpc string, contacts []Contact, data []byte, contact Contact, node *Kademlia) []byte {
	return CreateProtocol(rpc, contacts, data, node.routing.me, contact)
}

func(network *Network) createStoreMessage(data []byte, target Contact, node *Kademlia) []byte {
	return CreateProtocol("STORE", nil, data, node.routing.me, target)
}

func (network *Network) addContacts(contacts []Contact, node *Kademlia) {
	for _, con := range contacts {
		node.routing.AddContact(con)
	}
}

func CreateConnection(address string) net.Conn {
	conn, err := net.Dial("udp4", address)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}	
	return conn
}

func ReadFromConnection(conn net.Conn, buf []byte) []byte{
	rlen, err := conn.Read(buf)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	message := buf[:rlen]
	return message
}

func Unserialize(msg []byte) Protocol {
	var proto Protocol

	err := json.Unmarshal(msg, &proto)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return proto
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
		panic(err)
	}
	return protocol
}
