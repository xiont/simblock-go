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
