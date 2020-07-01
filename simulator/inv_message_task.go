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

var _ IAbstractMessageTask = new(InvMessageTask)

type InvMessageTask struct {
	*AbstractMessageTask
	/**
	 * Block to be advertised.
	 */
	block IBlock
}

func NewInvMessageTask(from *Node, to *Node, block IBlock) *InvMessageTask {
	return &InvMessageTask{
		NewAbstractMessageTask(from, to, settings.INV_MESSAGE),
		block,
	}
}

func (imt *InvMessageTask) GetProtocol() string {
	return imt.protocol
}

func (imt *InvMessageTask) GetBlock() IBlock {
	return imt.block
}

/**
 * Receive message at the <em>to</em> side.
 */
func (imt *InvMessageTask) Run() {
	imt.to.ReceiveMessage(imt)
}
