package simulator

import "simblock-go/settings"

var _ IAbstractMessageTask = new(InvMessageTask)

type InvMessageTask struct {
	*AbstractMessageTask
	/**
	 * Block to be advertised.
	 */
	block IBlock
}

func NewInvMessageTask(from *Node, to *Node, block IBlock) *InvMessageTask {
	return &InvMessageTask{
		NewAbstractMessageTask(from, to, settings.INV_MESSAGE),
		block,
	}
}

func (imt *InvMessageTask) GetProtocol() string {
	return imt.protocol
}

func (imt *InvMessageTask) GetBlock() IBlock {
	return imt.block
}

/**
 * Receive message at the <em>to</em> side.
 */
func (imt *InvMessageTask) Run() {
	imt.to.ReceiveMessage(imt)
}
