package cmd

import (
	"bytes"
	"io"
	"os"
	"testing"
)

var out io.Reader = os.Stdout

func TestPut(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "put")
	args = append(args, "abc")

	out = bytes.NewBuffer(nil)

	InitCLI(args)

	if actual := out.(*bytes.Buffer).String(); actual != "abc" {
		t.Errorf("expected abc, but got %s", actual)
	}
}
