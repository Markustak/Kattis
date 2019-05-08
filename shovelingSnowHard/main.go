package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

type node struct {
	x        int
	y        int
	distance int
	parentx  int
	parenty  int
	terrain  rune
}

func main() {

	scanner := &bufio.Scanner{}

	local := len(os.Args) > 1
	if local {
		file, err := os.Open("shovelling-1.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	width, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	height, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Width", width, " Height", height)
	aMap := make([][]rune, height)

	for i := range aMap {
		aMap[i] = make([]rune, width)
	}

	for i := 0; i < height; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j, v := range line {
			aMap[i][j] = rune(v)
		}
	}
	goal := node{}
	start := node{}
	for i := range aMap {
		for j := range aMap[i] {
			fmt.Print(string(aMap[i][j]))
			if string(aMap[i][j]) == "C" {
				goal.x = i
				goal.y = j
			} else if string(aMap[i][j]) == "B" {
				start.x = i
				start.y = j
			}
		}
		fmt.Println()
	}
	searchRoute(aMap, goal, start)
}
func searchRoute(aMap [][]rune, goal node, start node) {
	start.distance = 0
	fmt.Println("GOAL", goal, " start", start)
	dist := make([]node, 0)
	queue := NewPriorityQueue()
	visitedSet := make(map[string]node)
	for i, v := range aMap {
		for j := range v {
			if i+j == 0 {
				dist = append(dist, start)
			} else {
				dist = append(dist, node{
					x:        i,
					y:        j,
					distance: math.MaxInt32,
					terrain:  aMap[i][j],
				})
			}
			queue.Set(dist[i+j])
		}
	}
	for !queue.isEmpty() {
		current := queue.Next()

		neighbour := string(current.y-1) + "-" + string(current.x)
		alternative := 
	}
}
//https://brilliant.org/wiki/dijkstras-short-path-finder/#examples
func getKey(no node) string {
	return string(no.y) + "-" + string(no.x)
}
func isGoal(curr node, goal node) bool {
	if curr.x == goal.x && curr.y == goal.y {
		return true
	}
	return false
}

type priorityQueue struct {
	queue   []node
	nodeMap map[string]*node
}

func (p *priorityQueue) Len() int {
	return len(p.queue)
}

func (p *priorityQueue) Swap(i, j int) {
	p.queue[i], p.queue[j] = p.queue[j], p.queue[i]
}

func (p priorityQueue) Less(i, j int) bool {

	return p.queue[i].distance < p.queue[j].distance
}

func (p *priorityQueue) Set(no node) {
	key := string(no.y) + "-" + string(no.x)
	if _, ok := p.nodeMap[key]; ok {
		p.queue = append(p.queue, no)
	} else {
		p.queue = append(p.queue, no)
		p.nodeMap[key] = &no
	}

	sort.Sort(p)
}

func (p *priorityQueue) SetDIstance(key string, distance int) {
	p.nodeMap[key].distance = distance
}
func (p *priorityQueue) Next() node {

	node, left := p.queue[0], p.queue[1:]

	p.queue = left

	key := string(node.y) + "-" + string(node.x)
	delete(p.nodeMap, key)

	return node
}

func (p *priorityQueue) Get(key string) *node {
	if node, ok := p.nodeMap[key]; ok {
		return node
	}
	return nil
}
func (p *priorityQueue) isEmpty() bool {
	return len(p.queue) == 0
}

func NewPriorityQueue() *priorityQueue {
	var pq priorityQueue
	pq.queue = make([]node, 0)
	pq.nodeMap = make(map[string]*node)
	return &pq
}

/*

1	Create a node containing the goal state node_goal
2	Create a node containing the start state node_start
3	Put node_start on the open list
4	while the OPEN list is not empty
5	{
6	Get the node off the open list with the lowest f and call it node_current
7	if node_current is the same state as node_goal we have found the solution; break from the while loop
8	    Generate each state node_successor that can come after node_current
9	    for each node_successor of node_current
10	    {
11	        Set the cost of node_successor to be the cost of node_current plus the cost to get to node_successor from node_current
12	        find node_successor on the OPEN list
13	        if node_successor is on the OPEN list but the existing one is as good or better then discard this successor and continue
14	        if node_successor is on the CLOSED list but the existing one is as good or better then discard this successor and continue
15	        Remove occurences of node_successor from OPEN and CLOSED
16	        Set the parent of node_successor to node_current
17	        Set h to be the estimated distance to node_goal (Using the heuristic function)
18	         Add node_successor to the OPEN list
19	    }
20	    Add node_current to the CLOSED list
21	}

*/
