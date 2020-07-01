/*
   Copyright 2020 LittleBear(1018589158@qq.com)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package simulator

import (
	"math"
	"simblock-go/settings"
)

type ProofOfWork struct {
	*AbstractConsensusAlgo
}

var _ IAbstractConsensusAlgo = new(ProofOfWork)

func NewProofOfWork(selfNode *Node) *ProofOfWork {
	return &ProofOfWork{
		NewAbstractConsensusAlgo(selfNode),
	}
}

/**
 * Mints a new block by simulating Proof of Work.
 */
//override
func (pow *ProofOfWork) Minting() IAbstractMintingTask {
	selfNode := pow.GetSelfNode()
	parent := selfNode.GetBlock().(*ProofOfWorkBlock)
	difficulty := parent.GetNextDifficulty()
	var p float64 = 1.0 / float64(difficulty.Int64())
	u := settings.Rand.NextFloat64()
	//double u = random.nextDouble();
	// Task(minter,interval, difficulty)。
	// Power越大 interval 越小。

	if p <= math.Pow(2, -53) {
		return nil
	} else {
		interval := math.Log(u) / math.Log(1.0-p) / float64(selfNode.GetMiningPower())
		return NewMintingTask(selfNode, int64(interval), difficulty)
	}
}

/**
 * Tests if the receivedBlock is valid with regards to the current block. The receivedBlock
 * is valid if it is an instance of a Proof of Work block and the received block needs to have
 * a bigger difficulty than its parent next difficulty and a bigger total difficulty compared to
 * the current block.
 *
 * @param receivedBlock the received block
 * @param currentBlock  the current block
 * @return true if block is valid false otherwise
 */
//Override
func (pow *ProofOfWork) IsReceivedBlockValid(receivedBlock IBlock, currentBlock IBlock) bool {

	if _, ok := receivedBlock.(*ProofOfWorkBlock); ok {

	} else {
		// 如果接收到的区块不是工作量证明的区块类型，直接返回false
		return false
	}

	recPoWBlock := receivedBlock.(*ProofOfWorkBlock)

	var currPoWBlock *ProofOfWorkBlock = nil
	if currentBlock != nil {
		currPoWBlock = currentBlock.(*ProofOfWorkBlock)
	}

	receivedBlockHeight := receivedBlock.GetHeight()

	var receivedBlockParent *ProofOfWorkBlock
	//获取接收区块的父区块
	if receivedBlockHeight == 0 {
		receivedBlockParent = nil
	} else {
		var iReceivedBlockParent IBlock = receivedBlock.GetBlockWithHeight(receivedBlockHeight - 1)
		receivedBlockParent, _ = iReceivedBlockParent.(*ProofOfWorkBlock)
	}

	//TODO - dangerous to split due to short circuit operators being used, refactor?
	// 接收到区块的高度 = 0 || 接收区块的难度大于其父区块的下一个难度
	// &&
	// 当前区块为空 || 接收到区块的总难度大于当前区块的总难度

	cond1 := receivedBlockHeight == 0

	cond2 := true
	if receivedBlockParent != nil {
		cond2 = recPoWBlock.GetDifficulty().Cmp(receivedBlockParent.GetNextDifficulty()) != -1 //>= 0
	}

	cond3 := currentBlock == nil

	cond4 := true
	if currPoWBlock != nil {
		cond4 = recPoWBlock.GetTotalDifficulty().Cmp(currPoWBlock.GetTotalDifficulty()) == 1 //>0
	}

	return (cond1 || cond2) && (cond3 || cond4)
}

//Override
func (pow *ProofOfWork) GenesisBlock() IBlock {
	return (*ProofOfWorkBlock).GenesisBlock(nil, pow.GetSelfNode())
}
