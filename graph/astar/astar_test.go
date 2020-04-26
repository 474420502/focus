package astar

import (
	"testing"
)

func TestCaseSimplePath(t *testing.T) {
	graph := New2D(16, 16)
	graph.isDebug = true

	graph.SetTarget(15, 15, 0, 0)
	graph.Search()
	t.Error(graph.GetStep())
}

func TestCaseBlockPath(t *testing.T) {
	graph := New2D(16, 16)
	graph.isDebug = true

	graph.SetTarget(15, 15, 0, 0)
	graph.SetBlock(0, 1, 1)
	graph.SetBlock(1, 1, 1)
	graph.SetBlock(1, 2, 1)
	graph.SetBlock(1, 3, 1)
	graph.Search()
	t.Error(graph.GetStep())
}
