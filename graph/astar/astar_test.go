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
