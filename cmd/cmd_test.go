package cmd

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelp(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "put")
	args = append(args, "abc")
	InitCLI(args)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		assert.Equal(t, "abc", s.Text())
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
}

// func TestPut(t *testing.T) {
// 	err := app.Run([]string{"appname", "put"})
// 	assert.Nil(t, err, "HW should not return an error")
// }

// func TestCLI(t *testing.T) {
// 	for _, test := range []struct {
// 		Args   []string
// 		Output string
// 	}{
// 		{
// 			Args:   []string{"./put", "abc"},
// 			Output: "abc\n",
// 		},
// 		{
// 			Args:   []string{"./calc", "-mode", "multiply", "3", "2", "5"},
// 			Output: "30\n",
// 		},
// 	} {
// 		t.Run("", func(t *testing.T) {
// 			os.Args = test.Args
// 			out = bytes.NewBuffer(nil)
// 			InitCLI()

// 			if actual := out.(*bytes.Buffer).String(); actual != test.Output {
// 				fmt.Println(actual, test.Output)
// 				t.Errorf("expected %s, but got %s", test.Output, actual)
// 			}
// 		})
// 	}
// }

// func TestPut(t *testing.T) {
// 	err := app.Run([]string{"appname", "put"})
// 	assert.Nil(t, err, "HW should not return an error")
// }
