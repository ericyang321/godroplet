package main

import "github.com/ericyang321/godroplet/pkg/queue"

/*
	graph search for jug problem

	target = 4
	q = [(0, 0)]

	visited = {}

	while q.empty == false
		node = q.pop()
		if visited[node] {
			[a, b] = node
			if a == target || b == target
				for c, d in get_next_state(a, b) // all permutation
					q.add(c, d)
		}
		visited[node] = true
*/

type JugSet struct {
	One int
	Two int
}

func main() {
	q := new(queue.Queue)
}
