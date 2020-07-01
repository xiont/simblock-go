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
