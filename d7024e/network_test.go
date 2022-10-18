package d7024e

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNetwork(t *testing.T) {
	assert.NotNil(t, NewNetwork())
}

// func TestSendPingMessage(t *testing.T) {
// 	defaultIP := "172.20.0.2"
// 	port := 4000
// 	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
// 	routing := NewRoutingTable(defaultCon)
// 	network := NewNetwork()
// 	node := NewKademlia(routing, network)

// 	// target := NewKademliaID("2111111400000000000000000000000000000000")
// 	// targetCon := NewContact(target, "localhost:8002")
// 	assert.Equal(t, true, network.SendPingMessage(&defaultCon, node))
// }

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

func TestHandleFindContactMessage(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	routing := NewRoutingTable(defaultCon)
	network := NewNetwork()
	target := NewKademliaID("2111111400000000000000000000000000000000")

	// con1 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001")
	targetCon := NewContact(target, "localhost:8002")

	conArr := make([]Contact, 1)
	conArr[0] = targetCon

	node := NewKademlia(routing, network)
	proto := Protocol{"FIND_NODE", nil, []byte(target.String()), defaultCon, targetCon}
	created := node.network.handleFindContactMessage(proto, node)
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
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	routing := NewRoutingTable(defaultCon)
	network := NewNetwork()
	node := NewKademlia(routing, network)

	con1 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001")
	con2 := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002")

	conArr := make([]Contact, 2)
	conArr[0] = con1
	conArr[1] = con2

	data := []byte("1111111100000000000000000000000000000000")

	proto := Protocol{"FIND_VALUE", conArr, data, defaultCon, con1}
	created := node.network.handleFindDataMessage(proto, node)
	assert.NotNil(t, created)
}

func TestHandleFindDataMessageCurrentNode(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	kID := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
	defaultCon := NewContact(kID, defaultIP+":"+strconv.Itoa(port))
	routing := NewRoutingTable(defaultCon)
	network := NewNetwork()
	node := NewKademlia(routing, network)

	data := []byte("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
	// con2 := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), "localhost:8002")
	node.StoreValue(data)

	proto := Protocol{"FIND_VALUE", nil, data, defaultCon, defaultCon}
	created := node.network.handleFindDataMessage(proto, node)
	assert.NotNil(t, created)
	// assert.Equal(t, proto, created)
}
func TestCreateFindDataMessage(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	routing := NewRoutingTable(defaultCon)
	network := NewNetwork()
	node := NewKademlia(routing, network)
	con1 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001")
	con2 := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002")

	conArr := make([]Contact, 2)
	conArr[0] = con1
	conArr[1] = con2

	data := []byte("1111111100000000000000000000000000000000")

	proto := Protocol{"FIND_NODE", conArr, data, defaultCon, con1}
	created := node.network.createFindDataMessage(proto.Rpc, proto.Contacts, proto.Data, proto.Sender, node)
	assert.NotNil(t, created)
}

func TestNetworkAddContacts(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	routing := NewRoutingTable(defaultCon)
	network := NewNetwork()
	node := NewKademlia(routing, network)
	con1 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001")
	con2 := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002")

	conArr := make([]Contact, 2)
	conArr[0] = con1
	conArr[1] = con2

	network.addContacts(conArr, node)
}

// func TestNetwork_SendPingMessage(t *testing.T) {
// 	defaultIP := "127.0.0.1"
// 	port := 8001
// 	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
// 	routing := NewRoutingTable(defaultCon)
// 	network := NewNetwork()
// 	node := NewKademlia(routing, network)

// 	targetCon := NewContact(NewKademliaID("FFFdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port+1))
// 	targetCon2 := NewContact(NewKademliaID("1111111111111111111111111111111111111111"), "XXXDDD:22")

// 	type args struct {
// 		contact *Contact
// 		node    *Kademlia
// 	}
// 	tests := []struct {
// 		name    string
// 		network *Network
// 		args    args
// 		want    bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name:    "test sending pings to itself",
// 			network: network,
// 			args: args{
// 				contact: &targetCon,
// 				node:    node,
// 			},
// 			want: false,
// 		},
// 		{
// 			name:    "test sending pings to unavailable ip",
// 			network: network,
// 			args: args{
// 				contact: &targetCon2,
// 				node:    node,
// 			},
// 			want: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			network := &Network{}
// 			if got := network.SendPingMessage(tt.args.contact, tt.args.node); got != tt.want {
// 				t.Errorf("Network.SendPingMessage() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestNetwork_SendStoreMessage(t *testing.T) {
// 	defaultIP := "127.0.0.1"
// 	port := 8001
// 	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
// 	routing := NewRoutingTable(defaultCon)
// 	network := NewNetwork()
// 	node := NewKademlia(routing, network)

// 	targetCon := NewContact(NewKademliaID("FFFdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port+1))
// 	// targetCon2 := NewContact(NewKademliaID("1111111111111111111111111111111111111111"), "localhost:22")

// 	data := make([]byte, 2)
// 	data[0] = 97
// 	data[0] = 98
// 	data[0] = 99

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
// 		// TODO: Add test cases.
// 		{
// 			name:    "send a store message",
// 			network: network,
// 			args: args{
// 				contact: &targetCon,
// 				data:    data,
// 				node:    node,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// network := &Network{}
// 			network.SendStoreMessage(tt.args.contact, tt.args.data, tt.args.node)
// 		})
// 	}
// }

// func TestNetwork_HandleConn(t *testing.T) {
// 	defaultIP := "127.0.0.1"
// 	port := 8001
// 	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
// 	routing := NewRoutingTable(defaultCon)
// 	network := NewNetwork()
// 	node := NewKademlia(routing, network)

// 	conn, err := net.DialUDP("udp4")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	type args struct {
// 		conn *net.UDPConn
// 		node *Kademlia
// 	}
// 	tests := []struct {
// 		name    string
// 		network *Network
// 		args    args
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name:    "test handleconn",
// 			network: network,
// 			args: args{
// 				conn: conn,
// 				node: node,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// network := &Network{}
// 			network.HandleConn(tt.args.conn, tt.args.node)
// 		})
// 	}
// }

func TestUnserialize(t *testing.T) {
	defaultIP := "127.0.0.1"
	port := 8001
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))

	msg, err := json.Marshal(&Protocol{
		Rpc:      "rpc",
		Contacts: nil,
		Data:     nil,
		Sender:   defaultCon,
		Target:   defaultCon})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	assert.NotNil(t, Unserialize(msg))
}

func TestCreateProtocol(t *testing.T) {
	defaultIP := "127.0.0.1"
	port := 8001
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	proto := CreateProtocol("string", nil, nil, defaultCon, defaultCon)
	assert.NotNil(t, proto)
}
