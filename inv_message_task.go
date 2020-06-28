package main

var _ IAbstractMessageTask = new(InvMessageTask)

type InvMessageTask struct {
	amt AbstractMessageTask
	/**
	 * Block to be advertised.
	 */
	block IBlock
}

func (imt *InvMessageTask) GetFrom() *Node {
	return imt.GetFrom()
}

func (imt *InvMessageTask) GetTo() *Node {
	return imt.GetTo()
}

func (imt *InvMessageTask) Run() {
	imt.amt.Run()
}

func (imt *InvMessageTask) GetInterval() int64 {
	return imt.amt.GetInterval()
}

func NewInvMessageTask(from *Node, to *Node, block IBlock) *InvMessageTask {
	amt := NewAbstractMessageTask(from, to)
	return &InvMessageTask{
		amt:   *amt,
		block: block,
	}
}

func (imt *InvMessageTask) GetBlock() IBlock {
	return imt.block
}
