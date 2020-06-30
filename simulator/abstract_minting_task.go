package simulator

//type IAbstractMintingTask interface {
//	ITask
//}

type IAbstractMintingTask interface {
	ITask
	GetMinter() *Node
	GetParent() IBlock
}

var _ IAbstractMintingTask = new(AbstractMintingTask)

type AbstractMintingTask struct {
	/**
	 * The node to mint the block.
	 */
	minter *Node
	/**
	 * The parent block.
	 */
	parent IBlock
	/**
	 * Block interval in milliseconds.
	 */
	interval int64
}

func NewAbstractMintingTask(minter *Node, interval int64) *AbstractMintingTask {
	return &AbstractMintingTask{
		minter:   minter,
		parent:   minter.GetBlock(),
		interval: interval,
	}
}

/**
 * Gets minter.
 *
 * @return the minter
 */
func (amt *AbstractMintingTask) GetMinter() *Node {
	return amt.minter
}

/**
 * Gets the minted blocks parent.
 *
 * @return the parent
 */
func (amt *AbstractMintingTask) GetParent() IBlock {
	return amt.parent
}

//Override
func (amt *AbstractMintingTask) GetInterval() int64 {
	return amt.interval
}

func (amt *AbstractMintingTask) Run() {
	panic("abstract method")
}
