package simulator

import "math/big"

var _ IAbstractMintingTask = new(MintingTask)

type MintingTask struct {
	*AbstractMintingTask
	difficulty *big.Int
}

func NewMintingTask(minter *Node, interval int64, difficulty *big.Int) *MintingTask {
	return &MintingTask{
		NewAbstractMintingTask(minter, interval),
		difficulty,
	}
}

//Override
func (mt *MintingTask) Run() {
	var parent *ProofOfWorkBlock = nil
	if mt.GetParent() != nil {
		parent = mt.GetParent().(*ProofOfWorkBlock)
	}

	createdBlock := NewProofOfWorkBlock(parent, mt.GetMinter(), GetCurrentTime(), mt.difficulty)
	mt.GetMinter().receiveBlock(createdBlock)
}
