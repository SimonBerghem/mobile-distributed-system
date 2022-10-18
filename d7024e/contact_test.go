package d7024e

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendNoDups(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	defaultCon2 := NewContact(NewKademliaID("123deabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))

	conArr := make([]Contact, 1)
	conArr[0] = defaultCon

	conArr2 := make([]Contact, 1)
	conArr2[0] = defaultCon2

	candidates := ContactCandidates{
		contacts: conArr,
	}

	candidates.AppendNoDups(conArr2)
}

func TestGetUpToContacts(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := NewContact(NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	defaultCon2 := NewContact(NewKademliaID("123deabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))

	conArr := make([]Contact, 2)
	conArr[0] = defaultCon
	conArr[1] = defaultCon2

	candidates := ContactCandidates{
		contacts: conArr,
	}

	result := candidates.GetUpToContacts(2)
	assert.NotNil(t, result)

	resultOver := candidates.GetUpToContacts(10)
	assert.NotNil(t, resultOver)
}
