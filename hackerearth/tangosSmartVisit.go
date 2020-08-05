package main

func tangosSmartVisit(graph Graph, numOfFlights int) int {
	num := 0
	for _, vertex := range graph.Vertices {
		if !vertex.Explored {
			tangoDfs(graph, vertex, num)
			num++
		}
	}
	// We need (n-1) flights to reach n disconnected nodes
	if (num - 1) <= numOfFlights {
		return 1
	}
	return -1
}

func tangoDfs(graph Graph, startingVertex *Vertex, num int) {
	if startingVertex.Explored {
		return
	}
	startingVertex.Explored = true
	for _, edge := range graph.Edges {
		if edge.Head == startingVertex && !edge.Tail.Explored {
			tangoDfs(graph, edge.Tail, num)
		}
		if edge.Tail == startingVertex && !edge.Head.Explored {
			tangoDfs(graph, edge.Head, num)
		}
	}
}
