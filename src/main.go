package main

import (
    "log"
    "net"
    "fmt"
    // "strings"

    "github.com/SimonBerghem/mobile-distributed-system/src/d7024e"
)

func main (){

	ip := GetOutboundIP()
	fmt.Println(ip)

	// Create node
    nodeID := d7024e.NewRandomKademliaID()
    me := d7024e.NewContact(nodeID , ip)
    routing := d7024e.NewRoutingTable(me)
    node := d7024e.NewKademlia(routing)
    network := d7024e.NewNetwork()

    fmt.Println(node)
    fmt.Println(network)
    // Add node to network
    // 

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
