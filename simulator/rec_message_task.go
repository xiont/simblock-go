package simulator

import "simblock-go/settings"

var _ IAbstractMessageTask = new(RecMessageTask)

type RecMessageTask struct {
	*AbstractMessageTask
	/**
	 * Block to be advertised.
	 */
	block IBlock
}

func NewRecMessageTask(from *Node, to *Node, block IBlock) *RecMessageTask {
	return &RecMessageTask{
		NewAbstractMessageTask(from, to, settings.REC_MESSAGE),
		block,
	}
}

func (rmt *RecMessageTask) GetBlock() IBlock {
	return rmt.block
}

func (rmt *RecMessageTask) GetProtocol() string {
	return rmt.protocol
}

func (rmt *RecMessageTask) Run() {
	rmt.to.ReceiveMessage(rmt)
}
