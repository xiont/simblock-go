package main

import (
	"simblock-go/printfile"
)

var _ IAbstractMessageTask = new(BlockMessageTask)

type BlockMessageTask struct {
	amt AbstractMessageTask
	/**
	 * Block to be advertised.
	 */
	block IBlock
	/**
	 * The block message sending delay in milliseconds.
	 */
	interval int64
}

func (bmt *BlockMessageTask) GetFrom() *Node {
	return bmt.GetFrom()
}

func (bmt *BlockMessageTask) GetTo() *Node {
	return bmt.GetTo()
}

//override
func (bmt *BlockMessageTask) Run() {
	// 发送区块的节点继续发送区块
	bmt.GetFrom().SendNextBlockMessage()

	printfile.OUT_JSON_FILE.Write("{")
	printfile.OUT_JSON_FILE.Write("\"kind\":\"flow-block\",")
	printfile.OUT_JSON_FILE.Write("\"content\":{")
	printfile.OUT_JSON_FILE.Write("\"transmission-timestamp\":" + string(GetCurrentTime()-bmt.interval) + ",")
	printfile.OUT_JSON_FILE.Write("\"reception-timestamp\":" + string(GetCurrentTime()) + ",")
	printfile.OUT_JSON_FILE.Write("\"begin-node-id\":" + string(bmt.GetFrom().GetNodeID()) + ",")
	printfile.OUT_JSON_FILE.Write("\"end-node-id\":" + string(bmt.GetTo().GetNodeID()) + ",")
	printfile.OUT_JSON_FILE.Write("\"block-id\":" + string(bmt.block.GetId()))
	printfile.OUT_JSON_FILE.Write("}")
	printfile.OUT_JSON_FILE.Write("},")
	//OUT_JSON_FILE.flush();

	bmt.amt.Run()
}

//override
func (bmt *BlockMessageTask) GetInterval() int64 {
	return bmt.interval
}

func NewBlockMessageTask(from *Node, to *Node, block IBlock, delay int64) *BlockMessageTask {
	amt := NewAbstractMessageTask(from, to)
	return &BlockMessageTask{
		amt:      *amt,
		block:    block,
		interval: GetLatency(from.GetRegion(), to.GetRegion()) + delay,
	}
}

func (bmt *BlockMessageTask) GetBlock() IBlock {
	return bmt.block
}
