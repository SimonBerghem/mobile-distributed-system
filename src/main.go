package main

import (
	"log"
    "net"
	"fmt"
    // "strings"
	// . "d7024e"
)

func main (){

    fmt.Println("HELLO")
	ip := GetOutboundIP()
	fmt.Println(ip)

	// Create node
    nodeID := NewRandomKademliaID()
    me := NewContact(nodeID , ip)
    routing := NewRoutingTable(me)
    kademlia := newKademlia(routing)

    fmt.Println(kademlia)

    // Add node to network
	// Node listens for requests
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