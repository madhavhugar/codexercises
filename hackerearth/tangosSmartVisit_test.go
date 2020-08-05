package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTangosSmartVisit(t *testing.T) {
	vertex1 := Vertex{1, false}
	vertex2 := Vertex{2, false}
	vertex3 := Vertex{3, false}
	vertex4 := Vertex{4, false}
	vertex5 := Vertex{5, false}
	vertices := Vertices{
		&vertex1,
		&vertex2,
		&vertex3,
		&vertex4,
		&vertex5,
	}
	edges := Edges{
		{&vertex1, &vertex2},
		{&vertex3, &vertex4},
		{&vertex4, &vertex5},
	}
	graph := Graph{vertices, edges}
	numberOfFlights := 2
	wanted := 1
	got := tangosSmartVisit(graph, numberOfFlights)
	assert.Equal(t, wanted, got)

	vertex1 = Vertex{1, false}
	vertex2 = Vertex{2, false}
	vertex3 = Vertex{3, false}
	vertex4 = Vertex{4, false}
	vertex5 = Vertex{5, false}
	numberOfFlights = 0
	wanted = -1
	got = tangosSmartVisit(graph, numberOfFlights)
	assert.Equal(t, wanted, got)

	vertex1 = Vertex{1, false}
	vertex2 = Vertex{2, false}
	vertex3 = Vertex{3, false}
	vertex4 = Vertex{4, false}
	vertex5 = Vertex{5, false}
	numberOfFlights = 1
	wanted = 1
	got = tangosSmartVisit(graph, numberOfFlights)
	assert.Equal(t, wanted, got)
}

func TestTangoDfs(t *testing.T) {
	vertex1 := Vertex{1, false}
	vertex2 := Vertex{2, false}
	vertex3 := Vertex{3, false}
	vertex4 := Vertex{4, false}
	vertex5 := Vertex{5, false}
	vertices := Vertices{
		&vertex1,
		&vertex2,
		&vertex3,
		&vertex4,
		&vertex5,
	}
	edges := Edges{
		{&vertex1, &vertex2},
		{&vertex3, &vertex4},
		{&vertex4, &vertex5},
	}
	graph := Graph{vertices, edges}
	tangoDfs(graph, &vertex1, 0)
	assert.Equal(t, graph.Vertices[0].Explored, true)
	assert.Equal(t, graph.Vertices[1].Explored, true)
	assert.Equal(t, graph.Vertices[2].Explored, false)
	assert.Equal(t, graph.Vertices[3].Explored, false)
	assert.Equal(t, graph.Vertices[4].Explored, false)
}
