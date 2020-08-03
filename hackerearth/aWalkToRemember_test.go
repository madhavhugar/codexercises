package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalkToRemember(t *testing.T) {
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
		{&vertex2, &vertex3},
		{&vertex3, &vertex4},
		{&vertex4, &vertex5},
		{&vertex4, &vertex2},
	}
	walkAbleVertices := []int{0, 1, 1, 1, 0}
	graph := Graph{vertices, edges}
	got := walkToRemember(graph)
	assert.Equal(t, walkAbleVertices, got)

	vertex1 = Vertex{Value: 0, Explored: false}
	vertex2 = Vertex{Value: 0, Explored: false}
	vertex3 = Vertex{Value: 0, Explored: false}
	vertex4 = Vertex{Value: 0, Explored: false}
	vertex5 = Vertex{Value: 0, Explored: false}
	vertex6 := Vertex{Value: 0, Explored: false}
	vertex7 := Vertex{Value: 0, Explored: false}
	vertex8 := Vertex{Value: 0, Explored: false}
	vertex9 := Vertex{Value: 0, Explored: false}
	vertex10 := Vertex{Value: 0, Explored: false}
	vertex11 := Vertex{Value: 0, Explored: false}
	vertices = []*Vertex{
		&vertex1,
		&vertex2,
		&vertex3,
		&vertex4,
		&vertex5,
		&vertex6,
		&vertex7,
		&vertex8,
		&vertex9,
		&vertex10,
		&vertex11,
	}
	edges = []Edge{
		{&vertex1, &vertex5},
		{&vertex2, &vertex9},
		{&vertex3, &vertex1},
		{&vertex4, &vertex2},
		{&vertex4, &vertex9},
		{&vertex5, &vertex3},
		{&vertex6, &vertex8},
		{&vertex6, &vertex11},
		{&vertex7, &vertex4},
		{&vertex7, &vertex5},
		{&vertex8, &vertex9},
		{&vertex8, &vertex10},
		{&vertex8, &vertex11},
		{&vertex9, &vertex5},
		{&vertex9, &vertex7},
		{&vertex10, &vertex2},
		{&vertex10, &vertex6},
		{&vertex11, &vertex3},
	}
	graph = Graph{vertices, edges}
	walkAbleVertices = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	got = walkToRemember(graph)
	assert.Equal(t, walkAbleVertices, got)
}

func TestReverseTopologicalOrder(t *testing.T) {
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
		{&vertex2, &vertex3},
		{&vertex3, &vertex4},
		{&vertex4, &vertex5},
		{&vertex4, &vertex2},
	}
	graph := Graph{vertices, edges}
	order := reverseTopologicalOrder(graph)
	assert.Equal(t, &vertex5, order[0])
	assert.Equal(t, &vertex1, order[4])
}

func TestResetGraph(t *testing.T) {
	vertex1 := Vertex{1, true}
	vertex2 := Vertex{2, false}
	vertex3 := Vertex{3, false}
	vertex4 := Vertex{4, true}
	vertex5 := Vertex{5, true}
	vertices := Vertices{
		&vertex1,
		&vertex2,
		&vertex3,
		&vertex4,
		&vertex5,
	}
	edges := Edges{
		{&vertex1, &vertex2},
		{&vertex2, &vertex3},
		{&vertex3, &vertex4},
		{&vertex4, &vertex5},
		{&vertex4, &vertex2},
	}
	graph := Graph{vertices, edges}
	resetGraph(graph)
	assert.Equal(t, false, graph.Vertices[0].Explored)
	assert.Equal(t, false, graph.Vertices[1].Explored)
	assert.Equal(t, false, graph.Vertices[2].Explored)
	assert.Equal(t, false, graph.Vertices[3].Explored)
	assert.Equal(t, false, graph.Vertices[4].Explored)
}
