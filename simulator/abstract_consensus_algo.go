package simulator

type IAbstractConsensusAlgo interface {
	GenesisBlock() IBlock
	Minting() IAbstractMintingTask
	IsReceivedBlockValid(receivedBlock IBlock, currentBlock IBlock) bool
	GetSelfNode() *Node
}

var _ IAbstractConsensusAlgo = new(AbstractConsensusAlgo)

type AbstractConsensusAlgo struct {
	selfNode *Node
}

func NewAbstractConsensusAlgo(selfNode *Node) *AbstractConsensusAlgo {
	return &AbstractConsensusAlgo{
		selfNode: selfNode,
	}
}

/**
 * Gets the node using this consensus algorithm.
 *
 * @return the self node
 */
func (aca *AbstractConsensusAlgo) GetSelfNode() *Node {
	return aca.selfNode
}

/**
 * Gets the genesis block.
 *
 * @return the genesis block
 */
func (aca *AbstractConsensusAlgo) GenesisBlock() IBlock {
	panic("implement me")
}

/**
 * Minting abstract minting task.
 *
 * @return the abstract minting task
 */
func (aca *AbstractConsensusAlgo) Minting() IAbstractMintingTask {
	panic("implement me")
}

/**
 * Tests if the receivedBlock is valid with regards to the current block.
 *
 * @param receivedBlock the received block
 * @param currentBlock  the current block
 * @return true if block is valid false otherwise
 */
func (aca *AbstractConsensusAlgo) IsReceivedBlockValid(receivedBlock IBlock, currentBlock IBlock) bool {
	panic("implement me")
}
