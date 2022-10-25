package cmd

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/SimonBerghem/mobile-distributed-system/d7024e"
	"github.com/stretchr/testify/assert"
)

var out io.Writer = os.Stdout

func runCommandTester(line string) string {
	out = bytes.NewBuffer(nil)

	command, args := inputSplit(line)
	runCommand(out, nil, command, args)

	str := out.(*bytes.Buffer).String()
	return strings.TrimSuffix(str, "\n")
}

func TestInitCLI(t *testing.T) {
	defaultIP := "172.20.0.2"
	port := 4000
	defaultCon := d7024e.NewContact(d7024e.NewKademliaID("7bcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"), defaultIP+":"+strconv.Itoa(port))
	routing := d7024e.NewRoutingTable(defaultCon)
	network := d7024e.NewNetwork()
	node := d7024e.NewKademlia(routing, network)
	InitCLI(out, *node)
}

func TestPutNoArg(t *testing.T) {
	assert.Equal(t, invalidArgs, runCommandTester("put"))
}

func TestPutNoArgAlias(t *testing.T) {
	assert.Equal(t, invalidArgs, runCommandTester("p"))
}

func TestGetNoArg(t *testing.T) {
	assert.Equal(t, invalidArgs, runCommandTester("get"))
}

func TestGetNoArgAlias(t *testing.T) {
	assert.Equal(t, invalidArgs, runCommandTester("g"))
}

func TestHelp(t *testing.T) {
	assert.Equal(t, getHelpText(), runCommandTester("help"))
}

func TestHelpAlias(t *testing.T) {
	assert.Equal(t, getHelpText(), runCommandTester("h"))
}

func TestDefault(t *testing.T) {
	assert.Equal(t, invalidCommand, runCommandTester(""))
}
