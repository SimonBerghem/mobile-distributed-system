package d7024e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKademliaID(t *testing.T) {
	kademliaID := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
	assert.NotNil(t, kademliaID)
	assert.Equal(t, kademliaID.String(), "7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
}

func TestNewRandomKademliaID(t *testing.T) {
	kademliaID := NewRandomKademliaID()
	assert.NotNil(t, kademliaID)
}

func TestLess(t *testing.T) {
	kademliaID1 := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
	kademliaID2 := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdea1111")
	assert.Equal(t, true, kademliaID2.Less(kademliaID1))
}

func TestNotLess(t *testing.T) {
	kademliaID1 := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
	kademliaID2 := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdea1111")
	assert.Equal(t, false, kademliaID1.Less(kademliaID2))
}

func TestEquals(t *testing.T) {
	kademliaID1 := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
	kademliaID2 := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
	assert.Equal(t, true, kademliaID1.Equals(kademliaID2))
}

func TestNotEquals(t *testing.T) {
	kademliaID1 := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde")
	kademliaID2 := NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdea1111")
	assert.Equal(t, false, kademliaID1.Equals(kademliaID2))
}

func TestCalcDistance(t *testing.T) {
	kademliaID1 := NewKademliaID("1111111100000000000000000000000000000000")
	kademliaID2 := NewKademliaID("1111111200000000000000000000000000000000")
	assert.Equal(t, kademliaID1.CalcDistance(kademliaID2).String(), "0000000300000000000000000000000000000000")
}

func TestCalcReturnDistance(t *testing.T) {
	kademliaID1 := NewKademliaID("1111111100000000000000000000000000000000")
	kademliaID2 := NewKademliaID("1111111200000000000000000000000000000000")
	assert.Equal(t, kademliaID1.CalcDistance(kademliaID2), kademliaID2.CalcDistance(kademliaID1))
}
