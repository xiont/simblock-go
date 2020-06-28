package main

var _ IAbstractMessageTask = new(RecMessageTask)

type RecMessageTask struct {
	amt AbstractMessageTask
	/**
	 * Block to be advertised.
	 */
	block IBlock
}

func (rmt *RecMessageTask) GetFrom() *Node {
	return rmt.GetFrom()
}

func (rmt *RecMessageTask) GetTo() *Node {
	return rmt.GetTo()
}

func (rmt *RecMessageTask) Run() {
	rmt.amt.Run()
}

func (rmt *RecMessageTask) GetInterval() int64 {
	return rmt.amt.GetInterval()
}

func NewRecMessageTask(from *Node, to *Node, block IBlock) *RecMessageTask {
	amt := NewAbstractMessageTask(from, to)
	return &RecMessageTask{
		amt:   *amt,
		block: block,
	}
}

func (rmt *RecMessageTask) GetBlock() IBlock {
	return rmt.block
}
