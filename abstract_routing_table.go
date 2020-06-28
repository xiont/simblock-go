package main

import (
	"container/list"
)

type IAbstractRoutingTable interface {
	GetNumConnection() int
	SetNumConnection(int)
	GetNeighbors() *list.List
	AddNeighbor(node *Node) bool
	RemoveNeighbor(node *Node) bool
	InitTable()
}

type AbstractRoutingTable struct {
	IART IAbstractRoutingTable
}
