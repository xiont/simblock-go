package simulator

import (
	"math/big"
	"simblock-go/utils"
)

var GenesisNextDifficulty *big.Int = big.NewInt(0)

var _ IBlock = new(ProofOfWorkBlock)

type ProofOfWorkBlock struct {
	*Block
	difficulty      *big.Int
	totalDifficulty *big.Int
	nextDifficulty  *big.Int
}

//Override
func (powb *ProofOfWorkBlock) GetUUID() string {
	return "PoW_Block@" + utils.CreateRandomString(10)
}

func NewProofOfWorkBlock(parent *ProofOfWorkBlock, minter *Node, time int64, difficulty *big.Int) *ProofOfWorkBlock {

	var totalDifficulty *big.Int = big.NewInt(0)
	var nextDifficulty *big.Int = big.NewInt(0)

	if parent == nil {
		totalDifficulty = totalDifficulty.Add(totalDifficulty, difficulty)
		nextDifficulty = GenesisNextDifficulty
	} else {
		totalDifficulty = totalDifficulty.Add(parent.GetTotalDifficulty(), difficulty)
		// TODO: difficulty adjustment
		nextDifficulty = parent.GetNextDifficulty()
	}

	return &ProofOfWorkBlock{
		NewBlock(parent, minter, time),
		difficulty,
		totalDifficulty,
		nextDifficulty,
	}
}

/**
 * Gets difficulty.
 *
 * @return the difficulty
 */
func (powb *ProofOfWorkBlock) GetDifficulty() *big.Int {
	return powb.difficulty
}

/**
 * Gets total difficulty.
 *
 * @return the total difficulty
 */
func (powb *ProofOfWorkBlock) GetTotalDifficulty() *big.Int {
	return powb.totalDifficulty
}

/**
 * Gets next difficulty.
 *
 * @return the next difficulty
 */
func (powb *ProofOfWorkBlock) GetNextDifficulty() *big.Int {
	return powb.nextDifficulty
}

/**
 * Generates the genesis block, gets the total mining power and adjusts the difficulty of the
 * next block accordingly.
 *
 * @param minter the minter
 * @return the genesis block
 */
func (powb *ProofOfWorkBlock) GenesisBlock(minter *Node) *ProofOfWorkBlock {
	var totalMiningPower int64 = 0
	//计算所有区块的算力总和
	//for (Node node : getSimulatedNodes()) {
	//	totalMiningPower += node.getMiningPower();
	//}
	Simulator.GetSimulatedNodes().Each(func(index int, v interface{}) {
		totalMiningPower += v.(*Node).GetMiningPower()
	})

	//难度 = 总算力*区块间隔时间
	// 相当于难度没有调整：
	// 调整策略，取前几个区块的出块时间，与预计的总出块时间
	// 难度（非难度值） = (预计的总出块时间/前几个区块的出块时间)*当前难度
	// 比特币总每次调整不超过4倍
	GenesisNextDifficulty = GenesisNextDifficulty.Mul(big.NewInt(totalMiningPower), big.NewInt(Simulator.GetTargetInterval()))

	return NewProofOfWorkBlock(nil, minter, 0, big.NewInt(0))
}
