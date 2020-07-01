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
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/maps/linkedhashmap"
	"simblock-go/printfile"
	"simblock-go/settings"
	"strconv"
)

type simulator struct {
	/**
	 * A list of nodes that will be used in a simulation.
	 */
	simulatorNodes *arraylist.List

	/**
	 * The target block interval in milliseconds.
	 */
	targetInterval int64

	/**
	 * A list of observed {@link Block} instances.
	 * 模拟器观察到的区块实例
	 */
	observedBlocks *arraylist.List

	/**
	 * A list of observed block propagation times. The map key represents the id of the node that
	 * has seen the
	 * block, the value represents the difference between the current time and the block minting
	 * time, effectively
	 * recording the absolute time it took for a node to witness the block.
	 * 观察到的块传播时间列表。映射键代表见过该区块的节点的id，值代表当前时间与区块铸币时间的差值，
	 * 有效地记录了一个节点见证该区块的绝对时间。
	 */
	//ArrayList<LinkedHashMap<Integer, Long>>
	observedPropagations *arraylist.List
}

/**
 * Get simulated nodes list.
 */
func (sim *simulator) GetSimulatedNodes() *arraylist.List {
	return sim.simulatorNodes
}

/**
 * Get target block interval.
 */
func (sim *simulator) GetTargetInterval() int64 {
	return sim.targetInterval
}

/**
 * Sets the target block interval.
 */
func (sim *simulator) SetTargetInterval(interval int64) {
	sim.targetInterval = interval
}

/**
 * Add node to the list of simulated nodes.
 *
 * @param node the node
 */
func (sim *simulator) AddNode(node *Node) {
	sim.simulatorNodes.Add(node)
}

/**
 * Handle the arrival of a new block. For every observed block, propagation information is
 * updated, and for a new
 * block propagation information is created.
 *
 * @param block the block
 * @param node  the node
 */
func (sim *simulator) ArriveBlock(block IBlock, node *Node) {
	// If block is already seen by any node
	if sim.observedBlocks.Contains(block) {
		// Get the propagation information for the current block
		//       observedBlocks = [block1 block2 block3]
		// observedPropagations = [LinkedHashMap<Integer, Long>,LinkedHashMap<Integer, Long>,LinkedHashMap<Integer, Long>]
		// propagation = LinkedHashMap<Integer, Long>

		//LinkedHashMap<Integer, Long> propagation = observedPropagations.get(
		//	observedBlocks.indexOf(block)
		//);
		propagation, _ := sim.observedPropagations.Get(sim.observedBlocks.IndexOf(block))

		// Update information for the new block
		// LinkedHashMap<Integer, Long> .put(node.getNodeID(), time) <谁发过来的区块，其延迟>
		propagation.(*linkedhashmap.Map).Put(node.GetNodeID(), GetCurrentTime()-block.GetTime())
	} else {
		// If the block has not been seen by any node and there is no memory allocated
		// 这里说明只缓存最近10个块的，同时会将要删除的块快速打印出来**
		//TODO move magic number to constant
		if sim.observedBlocks.Size() > 10 {
			// After the observed blocks limit is reached, log and remove old blocks by FIFO principle
			block, _ := sim.observedBlocks.Get(0)
			obpg, _ := sim.observedPropagations.Get(0)
			sim.printPropagation(block.(IBlock), obpg.(*linkedhashmap.Map))

			sim.observedBlocks.Remove(0)
			sim.observedPropagations.Remove(0)
		}
		// If the block has not been seen by any node and there is additional memory
		propagation := linkedhashmap.New()
		propagation.Put(node.GetNodeID(), GetCurrentTime()-block.GetTime())
		// Record the block as seen
		sim.observedBlocks.Add(block)
		// Record the propagation time
		sim.observedPropagations.Add(propagation)
	}

}

/**
 * Print propagation information about the propagation of the provided block  in the format:
 *
 * <p><em>node_ID, propagation_time</em>
 *
 * <p><em>propagation_time</em>: The time from when the block of the block ID is generated to
 * when the
 * node of the <em>node_ID</em> is reached.
 *
 * @param block       the block
 * @param propagation the propagation of the provided block as a list of {@link Node} IDs and
 *                    propagation times
 */
func (sim *simulator) printPropagation(block IBlock, propagation *linkedhashmap.Map) {
	// Print block and its height
	//TODO block does not have a toString method, what is printed here
	//    System.out.println(block + ":" + block.getHeight());
	//    for (Map.Entry<Integer, Long> timeEntry : propagation.entrySet()) {
	//      System.out.println(timeEntry.getKey() + "," + timeEntry.getValue());
	//    }
	//    System.out.println();

	printfile.PROPAGATION_FILE.Println(block.GetUUID() + ":" + strconv.Itoa(block.GetHeight()))

	propagation.Each(func(k interface{}, v interface{}) {
		printfile.PROPAGATION_FILE.Println(strconv.Itoa(k.(int)) + "," + strconv.FormatInt(v.(int64), 10))
	})

	printfile.PROPAGATION_FILE.Println("")
	printfile.PROPAGATION_FILE.Flush()
}

/**
 * Print propagation information about all blocks, internally relying on
 * {@link Simulator#printPropagation(Block, LinkedHashMap)}.
 */
func (sim *simulator) PrintAllPropagation() {

	sim.observedBlocks.Each(func(index int, v interface{}) {
		obpg, _ := sim.observedPropagations.Get(index)
		sim.printPropagation(v.(IBlock), obpg.(*linkedhashmap.Map))
	})
}

var Simulator = simulator{
	simulatorNodes:       arraylist.New(),
	targetInterval:       settings.INTERVAL,
	observedBlocks:       arraylist.New(),
	observedPropagations: arraylist.New(),
}
