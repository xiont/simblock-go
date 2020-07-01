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
	"simblock-go/printfile"
	"simblock-go/settings"
	"strconv"
)

var _ IAbstractMessageTask = new(BlockMessageTask)

type BlockMessageTask struct {
	*AbstractMessageTask
	/**
	 * Block to be advertised.
	 */
	block IBlock
	/**
	 * The block message sending delay in milliseconds.
	 */
	interval int64
}

//override
func (bmt *BlockMessageTask) Run() {
	// 发送区块的节点继续发送区块
	bmt.GetFrom().SendNextBlockMessage()

	printfile.OUT_JSON_FILE.Print("{")
	printfile.OUT_JSON_FILE.Print("\"kind\":\"flow-block\",")
	printfile.OUT_JSON_FILE.Print("\"content\":{")
	printfile.OUT_JSON_FILE.Print("\"transmission-timestamp\":" + strconv.FormatInt(GetCurrentTime()-bmt.interval, 10) + ",")
	printfile.OUT_JSON_FILE.Print("\"reception-timestamp\":" + strconv.FormatInt(GetCurrentTime(), 10) + ",")
	printfile.OUT_JSON_FILE.Print("\"begin-node-id\":" + strconv.Itoa(bmt.GetFrom().GetNodeID()) + ",")
	printfile.OUT_JSON_FILE.Print("\"end-node-id\":" + strconv.Itoa(bmt.GetTo().GetNodeID()) + ",")
	printfile.OUT_JSON_FILE.Print("\"block-id\":" + strconv.Itoa(bmt.block.GetId()))
	printfile.OUT_JSON_FILE.Print("}")
	printfile.OUT_JSON_FILE.Print("},")
	printfile.OUT_JSON_FILE.Flush()
	//OUT_JSON_FILE.flush();

	//bmt.Run()
	bmt.to.ReceiveMessage(bmt)
}

//override
func (bmt *BlockMessageTask) GetInterval() int64 {
	return bmt.interval
}

func NewBlockMessageTask(from *Node, to *Node, block IBlock, delay int64) *BlockMessageTask {
	return &BlockMessageTask{
		NewAbstractMessageTask(from, to, settings.BLOCK_MESSAGE),
		block,
		GetLatency(from.GetRegion(), to.GetRegion()) + delay,
	}
}

func (bmt *BlockMessageTask) GetBlock() IBlock {
	return bmt.block
}

func (bmt *BlockMessageTask) GetProtocol() string {
	return bmt.protocol
}
