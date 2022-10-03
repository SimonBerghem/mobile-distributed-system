package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

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

func TestPutNoArg(t *testing.T) {
	assert.Equal(t, invalidArgs, runCommandTester("put"))
}

func TestPutNoArgAlias(t *testing.T) {
	assert.Equal(t, invalidArgs, runCommandTester("p"))
}

// func TestPutArg(t *testing.T) {
// 	assert.Equal(t, "'hashed'", runCommandTester("put file"))
// }

// func TestPutArgAlias(t *testing.T) {
// 	assert.Equal(t, "'hashed'", runCommandTester("p file"))
// }

func TestGetNoArg(t *testing.T) {
	assert.Equal(t, invalidArgs, runCommandTester("get"))
}

func TestGetNoArgAlias(t *testing.T) {
	assert.Equal(t, invalidArgs, runCommandTester("g"))
}

// func TestGetArg(t *testing.T) {
// 	assert.Equal(t, "'value'", runCommandTester("get hash"))
// }

// func TestGetArgAlias(t *testing.T) {
// 	assert.Equal(t, "'value'", runCommandTester("g hash"))
// }

func TestExit(t *testing.T) {
	// assert.NotNil(t, runCommandTester("exit"))
}

func TestExitAlias(t *testing.T) {
	//TODO
}

// func TestExitHelper(exitCommand string, t *testing.T) {
// 	out := ""
// 	if os.Getenv("BE_CRASHER") == "1" {
// 		out = runCommandTester(exitCommand)
// 		return
// 	}
// 	cmd := exec.Command(os.Args[0], "-test.run=TestCrasher")
// 	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
// 	err := cmd.Run()
// 	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
// 		return
// 	}
// 	t.Fatalf("process ran with err %v, want exit status 1", err)
// }

func TestHelp(t *testing.T) {
	assert.Equal(t, getHelpText(), runCommandTester("help"))
}

func TestHelpAlias(t *testing.T) {
	assert.Equal(t, getHelpText(), runCommandTester("h"))
}

func TestDefault(t *testing.T) {
	assert.Equal(t, invalidCommand, runCommandTester(""))
}

// func TestInitCLI(t *testing.T) {
// 	// input := "put"
// 	// var reader io.Reader = strings.NewReader(input)

// 	var buf bytes.Buffer
// 	log.SetOutput(&buf)
// 	defer func() {
// 		log.SetOutput(os.Stderr)
// 	}()
// 	node := d7024e.Kademlia{}
// 	node.InitNode()
// 	InitCLI(out, node)
// 	t.Log(buf.String())
// }
