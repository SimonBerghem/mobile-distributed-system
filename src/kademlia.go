package d7024e

// Stores routing table
type Kademlia struct {
	routing *RoutingTable
}

func newKademlia(routing *RoutingTable) *Kademlia {
	kademlia := &Kademlia{}
	kademlia.routing = NewRoutingTable()
	return kademlia
}

func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO
}
