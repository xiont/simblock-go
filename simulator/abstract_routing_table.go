package simulator

import "github.com/emirpasic/gods/lists/arraylist"

type IAbstractRoutingTable interface {
	GetSelfNode() *Node
	GetNumConnection() int
	SetNumConnection(int)
	GetNeighbors() *arraylist.List
	AddNeighbor(node *Node) bool
	RemoveNeighbor(node *Node) bool
	InitTable()
	AddInbound(from *Node) bool
	RemoveInbound(from *Node) bool
	AcceptBlock()
}

var _ IAbstractRoutingTable = new(AbstractRoutingTable)

type AbstractRoutingTable struct {
	selfNode      *Node
	numConnection int
}

func NewAbstractRoutingTable(selfnode *Node) *AbstractRoutingTable {
	return &AbstractRoutingTable{
		selfNode:      selfnode,
		numConnection: 8,
	}
}

/**
 * Gets self node.
 *
 * @return the self node
 */
func (a *AbstractRoutingTable) GetSelfNode() *Node {
	return a.selfNode
}

/**
 * Gets the number of possible active connections.
 *
 * @return the connection
 */
func (a *AbstractRoutingTable) GetNumConnection() int {
	return a.numConnection
}

/**
 * Sets the number of possible active connections.
 *
 * @param numConnection the n connection
 */
func (a *AbstractRoutingTable) SetNumConnection(numConnection int) {
	a.numConnection = numConnection
}

/**
 * Gets neighbors.
 *
 * @return the neighbors
 */
func (a *AbstractRoutingTable) GetNeighbors() *arraylist.List {
	panic("abstract method")
}

/**
 * Add a neighbor to the list of neighbors.
 *
 * @param node the node
 * @return the success state of the operation
 */
func (a *AbstractRoutingTable) AddNeighbor(node *Node) bool {
	panic("abstract method")
}

/**
 * Remove the neighbor from the list of neighbors.
 *
 * @param node the node
 * @return the success state of the operation
 */
func (a *AbstractRoutingTable) RemoveNeighbor(node *Node) bool {
	panic("abstract method")
}

/**
 * Table initialization.
 */
func (a *AbstractRoutingTable) InitTable() {
	panic("abstract method")
}

/**
 * Add inbound boolean.
 *
 * @param from the from
 * @return the boolean
 */
func (a *AbstractRoutingTable) AddInbound(from *Node) bool {
	return false
}

/**
 * Remove inbound boolean.
 *
 * @param from the from
 * @return the boolean
 */
//TODO possibly incoming requests - just the hook I need
func (a *AbstractRoutingTable) RemoveInbound(from *Node) bool {
	return false
}

/**
 * Accept block.
 */
//TODO unclear what this does
func (a *AbstractRoutingTable) AcceptBlock() {
	panic("abstract method")
}
