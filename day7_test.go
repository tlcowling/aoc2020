package aoc2020

import (
	"fmt"
	"github.com/tlcowling/adventutils"
	"log"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var (
	compile = regexp.MustCompile(`(\d|no) (.*) bag.*`)
)

type Node struct {
	ID    string
	Edges map[string]int
}

func NewNode(id string) *Node {
	return &Node{
		ID:    id,
		Edges: make(map[string]int),
	}
}

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes[n.ID] = n
}

func (g *Graph) AddEdge(id1, id2 string, weight int) {
	n1 := g.Nodes[id1]
	n2 := g.Nodes[id2]
	if n1 == nil {
		g.Nodes[id1] = NewNode(id1)
		n1 = g.Nodes[id1]
	}
	if n2 == nil {
		g.Nodes[id2] = NewNode(id2)
		n2 = g.Nodes[id2]
	}
	n1.Edges[id2] = weight
}

func (g *Graph) String() {
	for k, v := range g.Nodes {
		fmt.Printf("%v\n", k)

		for d, e := range v.Edges {
			fmt.Printf("  ->%d %v\n", e, d)
		}
		fmt.Println()
	}
}

func (g *Graph) pathExistsBetween(start, end string) bool {
	//fmt.Println("considering", start, "=>", end)
	if start == end {
		//fmt.Println("start is end", start, end)
		return true
	}
	startNode := g.Nodes[start]
	if len(startNode.Edges) == 0 {
		//fmt.Println("no more edges from", start)
		return false
	}

	if len(startNode.Edges) > 0 {
		for next := range startNode.Edges {
			if g.pathExistsBetween(next, end) {
				return true
			}
		}
	}
	//fmt.Println("no edges found anywhere scenario is this", start, end)
	return false
}

func (g *Graph) TraverseAllPathsFromID(id string) int {
	return g.traverseAllPathsFromID(id, 1)
}

func (g *Graph) traverseAllPathsFromID(id string, count int) int {
	node := g.Nodes[id]
	if len(node.Edges) == 0 {
		return count
	}

	edgeCount := 0
	for destination, edgeWeight := range node.Edges {
		countBelow := g.traverseAllPathsFromID(destination, edgeWeight)
		edgeCount += count * countBelow
	}
	return edgeCount + count
}

func parseLine(g *Graph, line string) {
	vertexes := strings.Split(line, "contain")
	parent := vertexes[0]
	parentDescription := strings.Join(strings.Split(parent, " ")[0:2], " ")
	//fmt.Println("parent: ", parentDescription)
	n1 := NewNode(parentDescription)
	g.AddNode(n1)

	children := strings.Split(vertexes[1], ", ")
	//fmt.Println("children: ", children)
	for _, r := range children {
		submatch := compile.FindStringSubmatch(r)
		//fmt.Println("child", i, submatch)
		q := submatch[1]
		if q == "no" {
			continue
		}
		bagQuantity, err := strconv.Atoi(q)
		if err != nil {
			log.Fatalln(err)
		}
		bag := submatch[2]
		//childNode := NewNode(bag)
		g.AddEdge(parentDescription, bag, bagQuantity)
	}
}

func TestDay7(t *testing.T) {
	lines := adventutils.ReadInputAsLines("./inputs/day7.txt")

	g := NewGraph()
	for _, line := range lines {
		parseLine(g, line)
	}

	mapCount := make(map[string]bool)
	pathCount := 0

	requiredDestination := "shiny gold"
	for n := range g.Nodes {
		if n == requiredDestination {
			continue
		}
		between := g.pathExistsBetween(n, requiredDestination)
		mapCount[n] = between

		if between {
			pathCount++
		}
	}

	t.Log(pathCount)

	bags := g.TraverseAllPathsFromID(requiredDestination)
	t.Log(bags - 1) // not counting the shiny bag itself
}
