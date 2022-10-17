package d7024e

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNetwork(t *testing.T) {
	assert.NotNil(t, NewNetwork())
}

func TestHandlePingMessage(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	routing := NewRoutingTable(defaultCon)
	network := NewNetwork()
	node := NewKademlia(routing, network)
	proto := Protocol{"PING", nil, nil, defaultCon, defaultCon}
	created := node.network.handlePingMessage(proto, node)
	assert.NotNil(t, created)
}

// func TestHandleFindContactMessage(t *testing.T) {
// 	defaultIP := "172.20.0.2"
// 	port := 4000
// 	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
// 	routing := NewRoutingTable(defaultCon)
// 	network := NewNetwork()
// 	node := NewKademlia(routing, network)
// 	node.routing.AddContact(defaultCon)
// 	proto := Protocol{"FIND_NODE", nil, nil, defaultCon, defaultCon}
// 	created := node.network.handleFindContactMessage(proto, node)
// 	assert.NotNil(t, created)
// }

func TestCreateFindContactMessage(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	kID := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
	defaultCon := NewContact(kID, defaultIP+":"+strconv.Itoa(port))
	routing := NewRoutingTable(defaultCon)
	network := NewNetwork()
	node := NewKademlia(routing, network)
	node.routing.AddContact(defaultCon)
	// conArray := [...]Contact{defaultCon}
	// conArray[0] = defaultCon
	// proto := Protocol{"FIND_NODE", nil, []byte(kID.String()), defaultCon, defaultCon}
	created := node.network.createFindContactMessage(defaultCon, kID, nil, node)
	assert.NotNil(t, created)
}

func TestHandleStoreMessage(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	routing := NewRoutingTable(defaultCon)
	network := NewNetwork()
	node := NewKademlia(routing, network)
	proto := Protocol{"STORE", nil, nil, defaultCon, defaultCon}
	created := node.network.handleStoreMessage(proto, node)
	assert.NotNil(t, created)
}

func TestHandleFindDataMessage(t *testing.T) {
	network := NewNetwork()
	data := network.handleFindDataMessage()
	assert.NotNil(t, data)
}
func TestCreateFindDataMessage(t *testing.T) {
	network := NewNetwork()
	data := network.createFindDataMessage()
	assert.NotNil(t, data)
}

// func TestNetwork_SendStoreMessage(t *testing.T) {
// 	defaultIP := "172.20.0.2"
// 	port := 4000
// 	kID := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
// 	defaultCon := NewContact(kID, defaultIP+":"+strconv.Itoa(port))
// 	routing := NewRoutingTable(defaultCon)
// 	network := NewNetwork()
// 	node := NewKademlia(routing, network)

// 	type args struct {
// 		contact *Contact
// 		data    []byte
// 		node    *Kademlia
// 	}
// 	tests := []struct {
// 		name    string
// 		network *Network
// 		args    args
// 	}{
// 		{
// 			name:    "test sendstore msg",
// 			network: NewNetwork(),
// 			args: args{
// 				contact: &defaultCon,
// 				data:    []byte(kID.String()),
// 				node:    node,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			network := &Network{}
// 			network.SendStoreMessage(tt.args.contact, tt.args.data, tt.args.node)
// 		})
// 	}
// }

// func TestNetwork_Listen(t *testing.T) {
// 	defaultIP := "127.0.0.1"
// 	port := 4000
// 	kID := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
// 	defaultCon := NewContact(kID, defaultIP+":"+strconv.Itoa(port))
// 	routing := NewRoutingTable(defaultCon)
// 	network := NewNetwork()
// 	node := NewKademlia(routing, network)

// 	type args struct {
// 		ip   string
// 		port int
// 		node *Kademlia
// 	}
// 	tests := []struct {
// 		name    string
// 		network *Network
// 		args    args
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name:    "test listen",
// 			network: NewNetwork(),
// 			args: args{
// 				ip:   defaultIP,
// 				port: port,
// 				node: node,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			network := &Network{}
// 			network.Listen(tt.args.ip, tt.args.port, tt.args.node)
// 		})
// 	}
// }
