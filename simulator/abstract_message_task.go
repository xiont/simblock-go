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

// 检查接口方法
//var _ ITask = new(AbstractMessageTask)

type IAbstractMessageTask interface {
	ITask
	GetFrom() *Node
	GetTo() *Node
	GetProtocol() string
}

// 检测抽象消息类是否实现了接口的方法
// var _ IAbstractMessageTask = new(AbstractMessageTask)

type AbstractMessageTask struct {
	from     *Node
	to       *Node
	protocol string
}

func NewAbstractMessageTask(from *Node, to *Node, protocol string) *AbstractMessageTask {
	return &AbstractMessageTask{
		from:     from,
		to:       to,
		protocol: protocol,
	}
}

/**
 * Get the sending node.
 *
 * @return the <em>from</em> node
 */
func (amt *AbstractMessageTask) GetFrom() *Node {
	return amt.from
}

/**
 * Get the receiving node.
 *
 * @return the <em>to</em> node
 */
func (amt *AbstractMessageTask) GetTo() *Node {
	return amt.to
}

/**
 * Get the message delay with regards to respective regions.
 *
 * @return the message sending interval
 */
func (amt *AbstractMessageTask) GetInterval() int64 {
	var latency int64 = GetLatency(amt.from.GetRegion(), amt.to.GetRegion())
	// Add 10 milliseconds here, why?  maybe some handle time
	//TODO
	return latency + 10
}
