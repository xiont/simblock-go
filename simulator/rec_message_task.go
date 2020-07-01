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

import "simblock-go/settings"

var _ IAbstractMessageTask = new(RecMessageTask)

type RecMessageTask struct {
	*AbstractMessageTask
	/**
	 * Block to be advertised.
	 */
	block IBlock
}

func NewRecMessageTask(from *Node, to *Node, block IBlock) *RecMessageTask {
	return &RecMessageTask{
		NewAbstractMessageTask(from, to, settings.REC_MESSAGE),
		block,
	}
}

func (rmt *RecMessageTask) GetBlock() IBlock {
	return rmt.block
}

func (rmt *RecMessageTask) GetProtocol() string {
	return rmt.protocol
}

func (rmt *RecMessageTask) Run() {
	rmt.to.ReceiveMessage(rmt)
}
