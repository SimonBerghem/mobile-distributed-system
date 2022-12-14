package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/SimonBerghem/mobile-distributed-system/d7024e"
)

const invalidCommand = "Invalid command, try again"
const invalidArgs = "Invalid argument or to many arguments"
const version = "0.0.0"

var input = os.Stdin

func InitCLI(output io.Writer, node d7024e.Kademlia) {
	fmt.Println("Starting CMD")

	scanner := bufio.NewScanner(input)

	for {
		fmt.Print(">")
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		command, arg := inputSplit(line)
		runCommand(output, &node, command, arg)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func inputSplit(line string) (string, string) {
	input := strings.Split(line, " ")
	command := input[0]
	input = input[1:]
	arg := strings.Join(input, " ")
	return command, arg
}

func runCommand(output io.Writer, node *d7024e.Kademlia, command string, arg string) {
	switch command {
	case "put":
		fmt.Fprintln(output, node.Store([]byte(arg)))

	case "p":
		fmt.Fprintln(output, node.Store([]byte(arg)))

	case "get":
		if len(arg) == 40 {
			fmt.Fprintln(output, node.LookupData(arg))
		} else {
			fmt.Fprintln(output, invalidArgs)
		}
	case "g":
		if len(arg) == 40 {
			fmt.Fprintln(output, node.LookupData(arg))
		} else {
			fmt.Fprintln(output, invalidArgs)
		}
	case "exit":
		os.Exit(1)
	case "e":
		os.Exit(1)
	case "help":
		fmt.Fprintln(output, getHelpText())
	case "h":
		fmt.Fprintln(output, getHelpText())
	case "version":
		fmt.Fprintln(output, version)
	case "v":
		fmt.Fprintln(output, version)
	default:
		fmt.Fprintln(output, invalidCommand)
	}
}

func getHelpText() string {
	return `
	NAME:
	   Kademlia CLI
	USAGE:
		A CLI for running Kademlia commands
	VERSION:
	   ` + version + `
	AUTHOR:
		Casper Lundberg, Simon Malmström Berghem & Emil Wiklander
	COMMANDS:
		put, p       Put content on network
		get, g       Get content from network by hash
		exit, e      Terminates the node one is attached to
		help, h      Show help (What ur reading right now)
		version, v   Print the version
	`
}
