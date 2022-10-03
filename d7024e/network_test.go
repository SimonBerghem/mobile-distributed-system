package d7024e

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNetwork(t *testing.T) {
	assert.NotNil(t, NewNetwork())
}

func TestSendPingMessage(t *testing.T) {
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
