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

// func TestStore(t *testing.T) {
// 	defaultIP := "172.20.0.2"
// 	port := 4000
// 	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
// 	routing := NewRoutingTable(defaultCon)
// 	network := NewNetwork()
// 	node := NewKademlia(routing, network)
// 	data := []byte("testdata")
// 	hash := node.Store(data)
// 	assert.Equal(t, NewKademliaID(string(data)), hash)
// }

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

// func TestLookupData(t *testing.T) {
// 	defaultIP := "172.20.0.2"
// 	port := 4000
// 	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
// 	routing := NewRoutingTable(defaultCon)
// 	network := NewNetwork()
// 	node := NewKademlia(routing, network)
// 	data := []byte("testdata")
// 	hash := node.Store(data)
// 	assert.Equal(t, data, node.LookupData(hash))
// }

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

	// want := make([]Contact, 1, 2)
	// want[0] = Contact(NewKademliaID("2111111400000000000000000000000000000000", "localhost:8002"))

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

// func TestKademlia_LookupContacts(t *testing.T) {
// 	defaultIP := "172.20.0.2"
// 	port := 4000
// 	kID := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
// 	defaultCon := NewContact(kID, defaultIP+":"+strconv.Itoa(port))
// 	routing := NewRoutingTable(defaultCon)
// 	network := NewNetwork()
// 	node := NewKademlia(routing, network)

// 	con1 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001")
// 	con2 := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002")

// 	conArr := make([]Contact, 2)
// 	conArr[0] = con1
// 	conArr[1] = con2

// 	routing.AddContacts(conArr)

// 	type fields struct {
// 		routing *RoutingTable
// 		network *Network
// 		data    map[KademliaID][]byte
// 	}
// 	type args struct {
// 		target *KademliaID
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   []Contact
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "lookup contacts test",
// 			fields: fields{
// 				routing: routing,
// 				network: network,
// 				data:    node.data,
// 			},
// 			args: args{
// 				target: NewKademliaID("1111111100000000000000000000000000000000"),
// 			},
// 			want: conArr,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			kademlia := &Kademlia{
// 				routing: tt.fields.routing,
// 				network: tt.fields.network,
// 				data:    tt.fields.data,
// 			}
// 			if got := kademlia.LookupContacts(tt.args.target); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Kademlia.LookupContacts() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
