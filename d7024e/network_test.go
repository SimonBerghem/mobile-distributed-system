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

	targetCon := NewContact(target, "localhost:8002")

	conArr := make([]Contact, 1)
	conArr[0] = targetCon

	node := NewKademlia(routing, network)
	proto := Protocol{"FIND_NODE", nil, []byte(target.String()), defaultCon, targetCon}
	created := node.network.handleFindContactMessage(proto, node)
	assert.NotNil(t, created)
}

func TestCreateFindContactMessage(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	kID := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
	defaultCon := NewContact(kID, defaultIP+":"+strconv.Itoa(port))
	routing := NewRoutingTable(defaultCon)
	network := NewNetwork()
	node := NewKademlia(routing, network)
	node.routing.AddContact(defaultCon)
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

// =====================
// NEEDS FIX, DATA IS NIL SOMEHOW
// =====================
// func TestHandleFindDataMessageCurrentNode(t *testing.T) {
// 	defaultIP := "172.20.0.2"
// 	port := 4000
// 	kID := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
// 	defaultCon := NewContact(kID, defaultIP+":"+strconv.Itoa(port))
// 	routing := NewRoutingTable(defaultCon)
// 	network := NewNetwork()
// 	node := NewKademlia(routing, network)

// 	data := []byte("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
// 	// con2 := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), "localhost:8002")
// 	node.StoreValue(data)

//		proto := Protocol{"FIND_VALUE", nil, data, defaultCon, defaultCon}
//		created := node.network.handleFindDataMessage(proto, node)
//		assert.NotNil(t, created)
//		// assert.Equal(t, proto, created)
//	}
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
