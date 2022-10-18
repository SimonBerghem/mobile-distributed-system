package d7024e

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutingTable(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	contacts := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}
}

func TestAddContacts(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	con1 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001")
	con2 := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002")

	conArr := make([]Contact, 2)
	conArr[0] = con1
	conArr[1] = con2

	rt.AddContacts(conArr)
}

func TestFindClosestContacts(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))
	bucketCon := NewContact(NewKademliaID("1111111111111111111111111111111111111111"), "localhost:8002")

	rt.buckets[100].AddContact(bucketCon)

	contacts := rt.FindClosestContacts(NewKademliaID("1111111111111111111111111111111111111111"), 0)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}
}

func TestRoutingTable_FindClosestContacts(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8010"))

	type fields struct {
		me      Contact
		buckets [IDLength * 8]*bucket
	}
	type args struct {
		target *KademliaID
		count  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Contact
	}{
		{
			name: "bucketindex > 1",
			fields: fields{
				me:      rt.me,
				buckets: rt.buckets,
			},
			args: args{
				target: NewKademliaID("1111111111111111111111111111111111111111"),
				count:  20,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routingTable := &RoutingTable{
				me:      tt.fields.me,
				buckets: tt.fields.buckets,
			}
			if got := routingTable.FindClosestContacts(tt.args.target, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoutingTable.FindClosestContacts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBucketIndex(t *testing.T) {
	kID := NewKademliaID("FFFFFFFF00000000000000000000000000000000")
	rt := NewRoutingTable(NewContact(kID, "localhost:8000"))
	target := NewKademliaID("0000000000000000000000000000000000000000")
	assert.NotNil(t, rt.getBucketIndex(target))
}
