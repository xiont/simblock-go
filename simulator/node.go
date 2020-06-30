package simulator

import (
	"container/list"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/sets/hashset"
	"simblock-go/printfile"
	"simblock-go/settings"
	"simblock-go/utils"
	"strconv"
)

type INode interface {
	getNodeId() int
	getRegion() int
	getMiningPower() int64
}

/**
 * A class representing a node in the network.
 */
type Node struct {
	/**
	 * Unique node ID.
	 */
	nodeID int

	/**
	 * Region assigned to the node.
	 */
	region int

	/**
	 * Mining power assigned to the node.
	 */
	miningPower int64

	/**
	 * A nodes routing table.
	 */
	routingTable IAbstractRoutingTable

	/**
	 * The consensus algorithm used by the node.
	 */
	consensusAlgo IAbstractConsensusAlgo

	/**
	 * The current block.
	 */
	block IBlock

	/**
	 * Orphaned blocks known to node.
	 */
	// TODO 实例化
	orphans *hashset.Set

	/**
	 * The current minting task
	 */
	mintingTask ITask

	/**
	 * In the process of sending blocks.
	 */
	// TODO verify
	sendingBlock bool

	messageQue *list.List

	downloadingBlocks *hashset.Set

	processingTime int
}

/**
 * Instantiates a new Node.
 *
 * @param nodeID            the node id
 * @param numConnection     the number of connections a node can have
 * @param region            the region
 * @param miningPower       the mining power
 * @param routingTableName  the routing table name
 * @param consensusAlgoName the consensus algorithm name
 */
func NewNode(nodeID int, numConnection int,
	region int, miningPower int64,
) *Node {

	node := &Node{
		nodeID:      nodeID,
		region:      region,
		miningPower: miningPower,
		//routingTable:      routingTable,
		//consensusAlgo:     consensusAlgo,
		block:             nil,
		orphans:           hashset.New(),
		mintingTask:       nil,
		sendingBlock:      false,
		messageQue:        new(list.List),
		downloadingBlocks: hashset.New(),
		processingTime:    2,
	}

	// Using the reflect function to find the Initial TABLE and ALGO
	r1, _ := utils.Call(*settings.FUNCS, settings.TABLE, node)
	node.SetRoutingTable(r1[0].Interface().(IAbstractRoutingTable))
	r2, _ := utils.Call(*settings.FUNCS, settings.ALGO, node)
	node.SetConsensusAlgo(r2[0].Interface().(IAbstractConsensusAlgo))
	node.setNumConnection(numConnection)

	return node
}

/**
 * Gets the node id.
 *
 * @return the node id
 */
func (n *Node) GetNodeID() int {
	return n.nodeID
}

/**
 * Gets the region ID assigned to a node.
 *
 * @return the region
 */
func (n *Node) GetRegion() int {
	return n.region
}

/**
 * Gets mining power.
 *
 * @return the mining power
 */
func (n *Node) GetMiningPower() int64 {
	return n.miningPower
} /**


 * Gets the consensus algorithm.
 *
 * @return the consensus algorithm. See {@link AbstractConsensusAlgo}
 */
func (n *Node) SetConsensusAlgo(algo IAbstractConsensusAlgo) {
	n.consensusAlgo = algo
}

/**
 * Gets routing table.
 *
 * @return the routing table
 */
func (n *Node) SetRoutingTable(table IAbstractRoutingTable) {
	n.routingTable = table
}

func (n *Node) getConsensusAlgo() IAbstractConsensusAlgo {
	return n.consensusAlgo
}

/**
 * Gets routing table.
 *
 * @return the routing table
 */
func (n *Node) getRoutingTable() IAbstractRoutingTable {
	return n.routingTable
}

/**
 * Gets the current block.
 *
 * @return the block
 */
func (n *Node) GetBlock() IBlock {
	return n.block
}

/**
 * Gets all orphans known to node.
 *
 * @return the orphans
 */
func (n *Node) GetOrphans() *hashset.Set {
	return n.orphans
}

/**
 * Gets the number of connections a node can have.
 *
 * @return the number of connection
 */
func (n *Node) GetNumConnection() int {
	return n.routingTable.GetNumConnection()
}

/**
 * Sets the number of connections a node can have.
 *
 * @param numConnection the n connection
 */
func (n *Node) setNumConnection(numConnection int) {
	n.routingTable.SetNumConnection(numConnection)
}

/**
 * Gets the nodes neighbors.
 *
 * @return the neighbors
 */
func (n *Node) GetNeighbors() *arraylist.List {
	return n.routingTable.GetNeighbors()
}

/**
 * Adds the node as a neighbor.
 *
 * @param node the node to be added as a neighbor
 * @return the success state of the operation
 */
func (n *Node) addNeighbor(node *Node) bool {
	return n.routingTable.AddNeighbor(node)
}

/**
 * Removes the neighbor form the node.
 *
 * @param node the node to be removed as a neighbor
 * @return the success state of the operation
 */
func (n *Node) removeNeighbor(node *Node) bool {
	return n.routingTable.RemoveNeighbor(node)
}

/**
 * Initializes the routing table.
 */
func (n *Node) JoinNetWork() {
	n.routingTable.InitTable()
}

/**
 * Mint the genesis block.
 */
func (n *Node) GenesisBlock() {
	genesis := n.consensusAlgo.GenesisBlock()
	n.receiveBlock(genesis)
}

/**
 * Adds a new block to the to chain. If node was minting that task instance is abandoned, and
 * the new block arrival is handled.
 *
 * @param newBlock the new block
 */
func (n *Node) addToChain(newBlock IBlock) {
	// If the node has been minting
	if n.mintingTask != nil {
		//Timer.removeTask
		RemoveTask(n.mintingTask)
		n.mintingTask = nil
	}
	// Update the current block
	n.block = newBlock
	n.printAddBlock(newBlock)
	// Observe and handle new block arrival
	// Simulator.arriveBlock 通知模拟器到达了区块
	Simulator.ArriveBlock(newBlock, n)
}

/**
 * Logs the provided block to the logfile.
 *
 * @param newBlock the block to be logged
 */
// TODO
func (n *Node) printAddBlock(newBlock IBlock) {
	printfile.OUT_JSON_FILE.Print("{")
	printfile.OUT_JSON_FILE.Print("\"kind\":\"add-block\",")
	printfile.OUT_JSON_FILE.Print("\"content\":{")
	printfile.OUT_JSON_FILE.Print("\"timestamp\":" + strconv.FormatInt(GetCurrentTime(), 10) + ",")
	printfile.OUT_JSON_FILE.Print("\"node-id\":" + strconv.Itoa(n.GetNodeID()) + ",")
	printfile.OUT_JSON_FILE.Print("\"block-id\":" + strconv.Itoa(newBlock.GetId()))
	printfile.OUT_JSON_FILE.Print("}")
	printfile.OUT_JSON_FILE.Print("},")
	printfile.OUT_JSON_FILE.Flush()
	//OUT_JSON_FILE.flush()
}

/**
 * Add orphans.
 *
 * @param orphanBlock the orphan block
 * @param validBlock  the valid block
 */
//TODO check this out later
func (n *Node) addOrphans(orphanBlock IBlock, validBlock IBlock) {
	if orphanBlock != validBlock {
		n.orphans.Add(orphanBlock)
		n.orphans.Remove(validBlock)
		if validBlock == nil || orphanBlock.GetHeight() > validBlock.GetHeight() {
			n.addOrphans(orphanBlock.GetParent(), validBlock)
		} else if orphanBlock.GetHeight() == validBlock.GetHeight() {
			n.addOrphans(orphanBlock.GetParent(), validBlock.GetParent())
		} else {
			n.addOrphans(orphanBlock, validBlock.GetParent())
		}
	}
}

/**
 * Generates a new minting task and registers it
 */
func (n *Node) minting() {
	mintingTask := n.consensusAlgo.Minting()
	n.mintingTask = mintingTask
	if mintingTask != nil {
		PutTask(mintingTask)
	}
}

/**
 * Send inv.
 *
 * @param block the block
 */
func (n *Node) sendInv(block IBlock) {
	//发送给所有邻居节点
	//for _,to := range n.routingTable.GetNeighbors() {
	//	invMessageTask := NewInvMessageTask(n, to, block)
	//	// Timer.putTask
	//	PutTask(invMessageTask)
	//

	it := n.routingTable.GetNeighbors().Iterator()
	for it.Next() {
		_, value := it.Index(), it.Value()
		invMessageTask := NewInvMessageTask(n, value.(*Node), block)
		// Timer.putTask
		PutTask(invMessageTask)
	}
}

/**
 * Receive block.
 *
 * @param block the block
 */
// TODO
func (n *Node) receiveBlock(block IBlock) {
	if n.consensusAlgo.IsReceivedBlockValid(block, n.block) {
		//如果共识协议判断该区块是正确的
		if n.block != nil && !n.block.IsOnSameChainAs(block) {
			// 如果主链的最新区块不是空的 ，并且给的区块不在本条主链上， 那么加入到孤块中
			// If orphan mark orphan
			n.addOrphans(n.block, block)
		}
		// FIXME ???? 注释else，但是实际没有,只是为了打印出来，前面做了一个标记
		// Else add to canonical chain
		// 添加到达区块的同时，废弃本节点的挖矿任务，同时通知模拟器区块的到达
		n.addToChain(block)
		// Generates a new minting task
		// 新建挖矿任务
		n.minting()
		// Advertise received block
		// 广播给其他节点区块的到达
		// 此处可以优化一点点（不要再发给我了）
		n.sendInv(block)
	} else if !n.orphans.Contains(block) && !block.IsOnSameChainAs(n.block) {
		// 如果共识协议不正确，如果本节点的孤块列表中没有包含该区块 ， 并且该区块也不在主链上
		// 加入到孤块中，直接通知模拟器区块的到达
		// TODO better understand - what if orphan is not valid?
		// If the block was not valid but was an unknown orphan and is not on the same chain as the
		// current block
		// 会根据孤块的父区块 和 本区块一直追溯，找到交叉点
		n.addOrphans(block, n.block)
		Simulator.ArriveBlock(block, n)
	}

}

/**
 * Receive message.
 *
 *      BlockMessage -->
 *       \/------------|
 * InvMessage ----> RecMessage
 *              \-> RecMessage
 *       /\------------|
 *     BlockMessage -->
 * @param message the message
 */
func (n *Node) ReceiveMessage(message IAbstractMessageTask) {
	from := message.GetFrom()

	if invMessage, ok := message.(*InvMessageTask); ok {
		// 如果是广播消息
		block := invMessage.GetBlock()
		if !n.orphans.Contains(block) && !n.downloadingBlocks.Contains(block) {
			// 如果孤块列表不包含 并且 下载的区块也不包含
			if n.consensusAlgo.IsReceivedBlockValid(block, n.block) {
				// 如果接收的区块合法
				recMessageTask := NewRecMessageTask(n, from, block)
				// 发送区块已接收消息
				PutTask(recMessageTask)
				// 添加到正在下载的区块列表中
				n.downloadingBlocks.Add(block)
			} else if !block.IsOnSameChainAs(n.block) {
				// 如果区块不是在主链上，添加到正在下载的列表，表示是一个孤块
				// get new orphan block
				recMessageTask := NewRecMessageTask(n, from, block)
				PutTask(recMessageTask)
				n.downloadingBlocks.Add(block)
			}
		}
	}

	if recMessage, ok := message.(*RecMessageTask); ok {
		// 如果是 接收到区块消息
		// sending block 是一个同步位？？
		n.messageQue.PushBack(recMessage)
		if !n.sendingBlock {
			// 这里如果RecMessage消息累积过多，可能会出现问题
			// 根据区块接收消息，新建一个区块发送消息 BlockMessage (包含延迟的)
			n.SendNextBlockMessage()
		}
	}

	if blockMessage, ok := message.(*BlockMessageTask); ok {
		// 接收到区块消息，从正在下载列表中移除，同时该节点接收到区块
		block := blockMessage.GetBlock()
		n.downloadingBlocks.Remove(block)
		// 节点接收到区块后会再次开始挖矿和广播
		n.receiveBlock(block)
	}
}

/**
 * Send next block message.
 */
// send a block to the sender of the next queued recMessage
// send 之后，将队列中的消息移除
func (n *Node) SendNextBlockMessage() {
	if n.messageQue.Len() > 0 {
		n.sendingBlock = true

		to := n.messageQue.Front().Value.(*RecMessageTask).GetFrom()
		block := n.messageQue.Front().Value.(*RecMessageTask).GetBlock()
		n.messageQue.Remove(n.messageQue.Front())
		bandwidth := GetBandwidth(n.GetRegion(), to.GetRegion())

		// Convert bytes to bits and divide by the bandwidth expressed as bit per millisecond, add
		// processing time.
		delay := settings.BLOCK_SIZE*8/(bandwidth/1000) + int64(n.processingTime)

		//
		blockMessageTask := NewBlockMessageTask(n, to, block, delay)

		PutTask(blockMessageTask)
	} else {
		n.sendingBlock = false
	}
}
