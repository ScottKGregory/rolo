// Package main is the primary entrypoint for rolo
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/scottkgregory/parsley"
	"github.com/scottkgregory/rolo/internal/nodes"
)

func main() {
	args := os.Args[1:]

	ret, err := run(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(ret)
}

func run(args []string) (string, error) {
	p, err := parsley.NewParser(false)
	if err != nil {
		return "", fmt.Errorf("error creating parser: %w", err)
	}

	p.RegisterBinaryNode("d", func(left, right parsley.Node) parsley.Node {
		return nodes.NewDiceNode(left, right)
	})

	ret, err := p.ParseAsString(strings.Join(args, " "), nil)
	if err != nil {
		return "", fmt.Errorf("error parsing args: %w", err)
	}

	return ret, nil
}
