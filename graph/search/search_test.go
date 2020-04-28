package gsearch

import (
	"log"
	"os"
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
	// graph.CountBlocksFlag()
	graph.Search()
	t.Error(graph.GetStep())
}

func TestCaseBlockCompare(t *testing.T) {
	graph := New2DFromBlockFile("./test4.bf")
	graph.isDebug = true
	graph.SetTarget(graph.dimX-1, graph.dimY-1, 0, 0)
	pather, _ := graph.Search()
	t.Error(graph.GetStep(), pather.GetPathLen())
}

func TestCaseBlockPathFile(t *testing.T) {
	var graph *Graph
	// f, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	// log.SetOutput(f)

	graph = New2D(16, 16)
	graph.isDebug = true
	graph.SetTarget(15, 15, 0, 0)
	graph.SetBlockFromFile("./test1.bf")
	graph.Search()
	t.Error(graph.GetStep())

	graph = New2D(16, 16)
	graph.isDebug = true
	graph.SetTarget(15, 15, 0, 0)
	graph.SetBlockFromFile("./test2.bf")
	pather, _ := graph.Search()
	t.Error(graph.GetStep(), pather.GetPathLen())
}

func TestCaseBlock512x512(t *testing.T) {
	var graph *Graph
	f, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	log.SetOutput(f)

	graph = New2D(2048, 2048)
	// graph.isDebug = true
	graph.SetTarget(0, 0, 2047, 2047)
	// graph.SetBlockFromFile("./test3.bf")
	graph.Search()
	t.Error(graph.GetStep())
}

func TestCaseBlock32x32(t *testing.T) {
	var graph *Graph
	// f, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	// log.SetOutput(f)

	graph = New2DFromBlockFile("./test3.bf")
	graph.isDebug = true
	graph.SetTarget(31, 31, 0, 0)
	// graph.SetBlockFromFile("./test3.bf")
	pather, _ := graph.Search()
	t.Error(graph.GetStep(), pather.GetPathLen())
}
