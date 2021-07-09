package sbt

type Node struct {
	Children []*Node
	Size     int64
}

type SBTree struct {
	Root *Node
}
