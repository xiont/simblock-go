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

import "reflect"

type IBlock interface {
	GetHeight() int
	GetParent() IBlock
	GetMinter() *Node
	GetTime() int64
	GetId() int
	GetBlockWithHeight(height int) IBlock
	IsOnSameChainAs(block IBlock) bool
	GetUUID() string
}

/**
 * Latest known block id.
 */
var LatestId int = 0

var _ IBlock = new(Block)

type Block struct {
	/**
	 * The current height of the block.
	 */
	height int

	/**
	 * The parent {@link Block}.
	 */
	parent IBlock

	/**
	 * The {@link Node} that minted the block.
	 */
	minter *Node

	/**
	 * Minting timestamp, absolute time since the beginning of the simulation.
	 */
	time int64

	/**
	 * Block unique id.
	 */
	id int
}

func (b *Block) GetUUID() string {
	panic("implement me")
}

func NewBlock(parent IBlock, minter *Node, time int64) *Block {
	//如果父区块为空，那么本区块就是创世区块，高度设置为0，否则就是父区块的高度加1

	height := 0
	//判断接口传递是否为空
	if reflect.ValueOf(parent).IsNil() {
		//		print("here")
		height = 0
	} else {
		height = parent.GetHeight() + 1
	}

	latestId := LatestId
	LatestId++
	return &Block{
		height: height,
		parent: parent,
		minter: minter,
		time:   time,
		id:     latestId,
	}
}

/**
 * Get height int.
 *
 * @return the int
 */
func (b *Block) GetHeight() int {
	return b.height
}

/**
 * Get parent block.
 *
 * @return the block
 */
func (b *Block) GetParent() IBlock {
	return b.parent
}

/**
 * Get minter node.
 *
 * @return the node
 */
func (b *Block) GetMinter() *Node {
	return b.minter
}

/**
 * Get time.
 *
 * @return the time
 */
//TODO what format
func (b *Block) GetTime() int64 {
	return b.time
}

/**
 * Gets the block id.
 *
 * @return the id
 */
//TODO what format
func (b *Block) GetId() int {
	return b.id
}

/**
 * Generates the genesis block. The parent is set to null and the time is set to 0
 *
 * @param minter the minter
 * @return the block
 */
func (b *Block) GenesisBlock(minter *Node) IBlock {
	//该方法作为一个静态方法，不需要创建实例。
	//传入一个节点就可已创建创世区块
	return NewBlock(nil, minter, 0)
}

/**
 * Recursively searches for the block at the provided height.
 *
 * @param height the height
 * @return the block with the provided height
 */
func (b *Block) GetBlockWithHeight(height int) IBlock {
	//获取指定高度的Block
	//这里采用的是递归的方法查询的
	if b.height == height {
		return b
	} else {
		return b.parent.GetBlockWithHeight(height)
	}
}

/**
 * Checks if the provided block is on the same chain as self.
 *
 * @param block the block to be checked
 * @return true if block are on the same chain false otherwise
 */
func (b *Block) IsOnSameChainAs(block IBlock) bool {
	// 检查所提供的块是否与自己的块在同一链上。
	if block == nil {
		// 如果区块为空，那么不在一条链上
		return false
	} else if b.height <= block.GetHeight() {
		// 如果该区块的高度 <= 给的区块的高度
		// genesis -> 1 -> 2 -> 3
		// 给的是 3 或者 4（4的前一节点是3） 。。。
		// 说明在同一条链上
		return b == block.GetBlockWithHeight(b.GetHeight())
	} else {
		// 如果给的区块高度比当前小
		// 那么要回溯本区块的父区块
		// genesis -> 1 -> 2 -> 3
		// 给的是 2
		// 如果2的下一个区块等于3，说明在同一个链上
		return b.GetBlockWithHeight(block.GetHeight()) == block
	}
}
