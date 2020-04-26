package astar

import (
	"testing"
)

func TestCaseSimplePath(t *testing.T) {
	graph := New2D(8, 8)
	graph.isDebug = true

	graph.SetTarget(7, 7, 0, 0)
	graph.Search()
	t.Error(graph.GetStep())
}
