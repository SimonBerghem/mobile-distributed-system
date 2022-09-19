package d7024e

import (
	"fmt"
	"encoding/json"
	"net"
	"strconv"
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

func Listen(ip string, port int) {
	
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

	fmt.Println("Listening on " + addrStr)
	for {
		HandleConn(conn)
	}
}

// Check which message has been recevied and handle it accordingly
func HandleConn(conn *net.UDPConn){
	buf := make([]byte, 1024)
	rlen, _ , err := conn.ReadFromUDP(buf)
	fmt.Println("Got message: ", rlen)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}

	values := buf[:rlen]
	fmt.Println("Message:", values)

}

func (network *Network) SendPingMessage(contact *Contact) {

	conn, err := net.Dial("udp4", contact.Address)

	if err != nil {
		fmt.Println(err)
		// panic(err)
	}

	// TODO: create proper ping message
	conn.Write([]byte("Hello World!"))

	buf := make([]byte, 1024)
	rlen, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	message := buf[:rlen]
	fmt.Println(message)

	// TODO: add contact to routing table
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
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


