package d7024e

// Stores routing table
type Kademlia struct {
	routing *RoutingTable
}

func NewKademlia(table *RoutingTable) Kademlia {
	return Kademlia{table}
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
