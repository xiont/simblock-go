package main

type IBlock interface {
	GetHeight() int
	GetParent() IBlock
	GetMinter() *Node
	GetTime() int64
	GetId() int
	GetBlockWithHeight(height int) IBlock
	IsOnSameChainAs(block IBlock) bool
}

type Block struct {
	IB IBlock
}
