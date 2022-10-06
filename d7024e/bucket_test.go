package d7024e

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBucket(t *testing.T) {
	bucket := newBucket()
	assert.NotNil(t, bucket)
}

func TestAddContact(t *testing.T) {
	bucket := newBucket()
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	bucket.AddContact(defaultCon)
	assert.Equal(t, defaultCon, bucket.list.Front().Value)
}

func TestAddContactTwice(t *testing.T) {
	bucket := newBucket()
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	bucket.AddContact(defaultCon)
	bucket.AddContact(defaultCon)
	assert.Equal(t, defaultCon, bucket.list.Front().Value)
}

func TestLen(t *testing.T) {
	bucket := newBucket()
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	bucket.AddContact(defaultCon)
	assert.Equal(t, 1, bucket.Len())
}
