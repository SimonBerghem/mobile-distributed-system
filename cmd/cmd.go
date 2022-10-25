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
		command, args := inputSplit(line)
		runCommand(output, &node, command, args)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func inputSplit(line string) (string, []string) {
	args := strings.Split(line, " ")
	command := args[0]
	args = args[1:]
	return command, args
}

func runCommand(output io.Writer, node *d7024e.Kademlia, command string, args []string) {
	n := len(args)
	switch command {
	case "put":
		if n == 1 {
			fmt.Fprintln(output, node.Store([]byte(args[0])))
		} else {
			fmt.Fprintln(output, invalidArgs)
		}
	case "p":
		if n == 1 {
			fmt.Fprintln(output, node.Store([]byte(args[0])))
		} else {
			fmt.Fprintln(output, invalidArgs)
		}
	case "get":
		if n == 1 {
			fmt.Fprintln(output, node.LookupData(args[0]))
		} else {
			fmt.Fprintln(output, invalidArgs)
		}
	case "g":
		if n == 1 {
			fmt.Fprintln(output, node.LookupData(args[0]))
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
	   0.0.0
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
