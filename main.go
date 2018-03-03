package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jmg-duarte/simmulated-annealing/graph"
)

func main() {

	tempStart := flag.Float64("t", 9000.0, "starting temperature")
	tempLimit := flag.Float64("tL", 5.0, "temperature limit")
	nIterations := flag.Int("nI", 5, "number of iterations per temperature")

	flag.Parse()

	if len(os.Args) < 2 {
		log.Fatal("usage: sim <node_file> [flags]")
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

	prob := &Problem{
		Graph:               g,
		StartingTemperature: *tempStart,
		LimitTemperature:    *tempLimit,
		NumberOfIterations:  *nIterations,
	}

	setupProblem(prob)
}

func setupProblem(p *Problem) error {
	graph := p.Graph
	path := graph.GenerateRandomPath()

	path.GenerateNeighbour()

	fmt.Println(
		graph.SimmulatedAnnealing(
			p.StartingTemperature, p.LimitTemperature, p.NumberOfIterations))

	return nil
}

//type Decay func(float64, float64)

// For each line of the file, splits with "\t",
// validates the line as a Node (a valid line would be "ID\tYCoord\tYCoord")
// then adds the node to the Graph
func parseFile(scanner *bufio.Scanner) (*graph.Graph, error) {
	// TODO Make 32 be modifiable
	// Mental note:
	// Don't be retarded and remember to use make(T, l, c)
	// for arrays that will be filled and not copied to
	nodes := graph.Graph{
		Nodes:    make(map[string]*graph.Node),
		NodeList: make(graph.NodeArray, 0, 32),
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

func validateNode(line []string) (*graph.Node, error) {
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

	return &graph.Node{
		ID:          line[0],
		XCoordinate: x,
		YCoordinate: y,
	}, nil
}
