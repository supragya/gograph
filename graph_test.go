package gograph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gograph "github.com/supragya/gograph"
)

func TestCreation(t *testing.T) {
	graph := gograph.NewGraphStringVertexStringEdge(false)
	assert.Equal(t, graph.GetVertexCount(), 0, "vertex count")
	assert.Equal(t, graph.GetEdgeCount(), 0, "edge count")
}

func TestAddOneEdge(t *testing.T) {
	graph := gograph.NewGraphStringVertexStringEdge(false)
	graph.AddEdge("vertex1", "vertex2", "middata")
	assert.Equal(t, graph.GetVertexCount(), 2, "vertex count")
	assert.Equal(t, graph.GetEdgeCount(), 1, "edge count")
	connections := graph.GetConnectedVertices("vertex1")
	assert.Equal(t, len(connections), 1, "number of connections")
	for vertex, edge := range connections {
		assert.Equal(t, vertex, "vertex2", "vertex")
		assert.Equal(t, edge, "middata", "edge")
	}
	connections = graph.GetConnectedVertices("vertex2")
	assert.Equal(t, len(connections), 0, "number of connections")
}

func TestAddOneEdgeBidirectional(t *testing.T) {
	graph := gograph.NewGraphStringVertexStringEdge(true)
	graph.AddEdge("vertex1", "vertex2", "middata")
	assert.Equal(t, graph.GetVertexCount(), 2, "vertex count")
	assert.Equal(t, graph.GetEdgeCount(), 1, "edge count")
	connections := graph.GetConnectedVertices("vertex1")
	assert.Equal(t, len(connections), 1, "number of connections")
	for vertex, edge := range connections {
		assert.Equal(t, vertex, "vertex2", "vertex")
		assert.Equal(t, edge, "middata", "edge")
	}
	connections = graph.GetConnectedVertices("vertex2")
	assert.Equal(t, len(connections), 1, "number of connections")
	for vertex, edge := range connections {
		assert.Equal(t, vertex, "vertex1", "vertex")
		assert.Equal(t, edge, "middata", "edge")
	}
}