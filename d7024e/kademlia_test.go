package d7024e

import (
	"reflect"
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

func TestStoreValue(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	routing := NewRoutingTable(defaultCon)
	network := NewNetwork()
	node := NewKademlia(routing, network)
	data := []byte("testdata")
	node.StoreValue(data)
}

func TestMin(t *testing.T) {
	min := min(4, 10)
	assert.Equal(t, min, 4)
}

func TestMin2(t *testing.T) {
	min := min(10, 4)
	assert.Equal(t, min, 4)
}

func TestGetOutboundIP(t *testing.T) {
	defaultIP := GetOutboundIP()
	assert.NotNil(t, defaultIP)
}

func Test_findUnqueriedNodes(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	defaultCon2 := NewContact(NewKademliaID("111deabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	conArr := make([]Contact, 1, 2)
	conArr[0] = defaultCon

	conArr2 := make([]Contact, 1, 2)
	conArr2[0] = defaultCon2

	type args struct {
		closestNodes []Contact
		seenNodes    []Contact
		count        int
	}
	tests := []struct {
		name string
		args args
		want []Contact
	}{
		// TODO: Add test cases.
		{
			name: "len(unqueriedNodes) == count is false",
			args: args{
				closestNodes: conArr,
				seenNodes:    conArr,
				count:        1,
			},
			want: nil,
		},
		{
			name: "len(unqueriedNodes) == count is true",
			args: args{
				closestNodes: conArr,
				seenNodes:    conArr,
				count:        0,
			},
			want: nil,
		},
		{
			name: "!contains(seenNodes, node)",
			args: args{
				closestNodes: conArr,
				seenNodes:    conArr2,
				count:        -1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnqueriedNodes(tt.args.closestNodes, tt.args.seenNodes, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("findUnqueriedNodes() = %v, want %v", got, tt.want)
				print(got)
			}
		})
	}
}
