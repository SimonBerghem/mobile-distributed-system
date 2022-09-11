package main

import (
	"log"
    "net"
	"fmt"
    // "strings"
	// "kademlia"
)

func main (){
	ip := GetOutboundIP()
	fmt.Println(ip)
	// Create node
	// Add node to network
	// Node listens for requests
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}