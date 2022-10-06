package d7024e

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKademlia(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	routing := NewRoutingTable(defaultCon)
	network := NewNetwork()
	node := NewKademlia(routing, network)
	assert.NotNil(t, node)
}

func TestContains(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))

	var contacts []Contact
	contacts = append(contacts, defaultCon)
	assert.Equal(t, true, contains(contacts, defaultCon))
}

func TestNotContains(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))

	var contacts []Contact
	contacts = append(contacts, defaultCon)
	assert.NotEqual(t, true, contains(contacts, NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdf"), defaultIP+":"+strconv.Itoa(port))))
}

func TestStore(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	routing := NewRoutingTable(defaultCon)
	network := NewNetwork()
	node := NewKademlia(routing, network)
	data := []byte("testdata")
	hash := node.Store(data)
	assert.Equal(t, Hash(data), hash)
}

func TestLookupData(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	routing := NewRoutingTable(defaultCon)
	network := NewNetwork()
	node := NewKademlia(routing, network)
	data := []byte("testdata")
	hash := node.Store(data)
	assert.Equal(t, data, node.LookupData(hash))
}

func TestGetOutboundIP(t *testing.T) {
	defaultIP := GetOutboundIP()
	assert.NotNil(t, defaultIP)
}
