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
	graph.SetBlock(0, 2, 1)
	graph.SetBlock(1, 1, 1)
	graph.SetBlock(1, 2, 1)
	graph.SetBlock(1, 3, 1)
	graph.SetBlock(1, 4, 1)
	graph.SetBlock(7, 7, 1)
	graph.SetBlock(7, 8, 1)
	graph.Search()
	t.Error(graph.GetStep())
}

func TestCaseBlockPathFile(t *testing.T) {
	graph := New2D(16, 16)
	graph.isDebug = true
	graph.SetTarget(15, 15, 0, 0)
	graph.SetBlockFromFile("./test1.bf")
	graph.Search()
	t.Error(graph.GetStep())

	graph = New2D(16, 16)
	graph.isDebug = true
	graph.SetTarget(15, 15, 0, 0)
	graph.SetBlockFromFile("./test2.bf")
	graph.Search()
	t.Error(graph.GetStep())
}
