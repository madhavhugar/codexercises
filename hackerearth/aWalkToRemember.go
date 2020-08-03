package main

//

type Vertex struct {
	Value    interface{}
	Explored bool
}

type Vertices []*Vertex

type Edge struct {
	Head *Vertex
	Tail *Vertex
}

type Edges []Edge

type Graph struct {
	Vertices []*Vertex
	Edges    []Edge
}

var order Vertices

func dfsReverse(graph Graph, startingVertex *Vertex) {
	if startingVertex.Explored {
		return
	}
	startingVertex.Explored = true
	for _, edge := range graph.Edges {
		if edge.Tail == startingVertex && !edge.Head.Explored {
			dfsReverse(graph, edge.Head)
		}
	}
	order = append(Vertices{startingVertex}, order...)
}

func reverseTopologicalOrder(graph Graph) Vertices {
	for _, vertex := range graph.Vertices {
		if !vertex.Explored {
			dfsReverse(graph, vertex)
		}
	}
	return order
}

func resetGraph(graph Graph) {
	for _, vertex := range graph.Vertices {
		vertex.Explored = false
		vertex.Value = 0
	}
}

func dfs(graph Graph, startingVertex *Vertex) {
	if startingVertex.Explored {
		return
	}
	startingVertex.Explored = true
	for _, edge := range graph.Edges {
		if edge.Head == startingVertex && !edge.Tail.Explored {
			edge.Head.Value = 1
			edge.Tail.Value = 1
			dfs(graph, edge.Tail)
		}
	}
}

const validPoint = 1
const invalidPoint = 0

func walkToRemember(graph Graph) (result []int) {
	order := reverseTopologicalOrder(graph)
	resetGraph(graph)
	for _, vertex := range order {
		if !vertex.Explored {
			dfs(graph, vertex)
		}
	}
	for _, vertex := range graph.Vertices {
		result = append(result, vertex.Value.(int))
	}
	return result
}
