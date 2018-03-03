package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("usage: sim <node_file> <temperature>")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Closes the file when "main" returns
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(bufio.NewReader(f))
	g, err := parseFile(scanner)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(g)

	p := g.GenerateRandomPath()
	//fmt.Println(p)

	p.GenerateNeighbour()

	t, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g.SimmulatedAnnealing(t, 5))
}

//type Decay func(float64, float64)

// For each line of the file, splits with "\t",
// validates the line as a Node (a valid line would be "ID\tYCoord\tYCoord")
// then adds the node to the Graph
func parseFile(scanner *bufio.Scanner) (*Graph, error) {
	// TODO Make 32 be modifiable
	// Mental note:
	// Don't be retarded and remember to use make(T, l, c)
	// for arrays that will be filled and not copied to
	nodes := Graph{
		Nodes:    make(map[string]*Node),
		NodeList: make([]*Node, 0, 32),
	}

	// Skips the header line
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		s := strings.Split(line, "\t")
		vNode, err := validateNode(s)
		if err != nil {
			return nil, err
		}
		nodes.AddNode(vNode)
	}
	return &nodes, nil
}

func validateNode(line []string) (*Node, error) {
	if len(line) != 3 {
		return nil, fmt.Errorf("invalid line: cannot contain more than 3 elements\n\t%s", line)
	}

	x, err := strconv.ParseFloat(line[1], 64)
	if err != nil {
		return nil, err
	}

	y, err := strconv.ParseFloat(line[2], 64)
	if err != nil {
		return nil, err
	}

	return &Node{line[0], x, y}, nil
}
