package simulator

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"simblock-go/printfile"
	"simblock-go/settings"
	"strconv"
)

type BitcoinCoreTable struct {
	*AbstractRoutingTable
	outbound *arraylist.List
	inbound  *arraylist.List
}

var _ IAbstractRoutingTable = new(BitcoinCoreTable)

func NewBitcoinCoreTable(selfNode *Node) *BitcoinCoreTable {
	return &BitcoinCoreTable{
		NewAbstractRoutingTable(selfNode),
		arraylist.New(),
		arraylist.New(),
	}
}

/**
 * Gets all known outbound and inbound nodes.
 *
 * @return a list of known neighbors
 */
func (bct *BitcoinCoreTable) GetNeighbors() *arraylist.List {
	neighbors := arraylist.New()
	bct.outbound.Each(func(index int, v interface{}) {
		neighbors.Add(v)
	})
	bct.inbound.Each(func(index int, v interface{}) {
		neighbors.Add(v)
	})
	return neighbors
}

/**
 * Initializes a new BitcoinCore routing table. From a pool of
 * all available nodes, choose candidates at random and
 * fill the table using the allowed outbound connections
 * amount.
 */
//TODO this should be done using the bootstrap node
func (bct *BitcoinCoreTable) InitTable() {
	var candidates []int

	// 拷贝一份节点的id
	for i := 0; i < Simulator.GetSimulatedNodes().Size(); i++ {
		candidates = append(candidates, i)
	}
	// 打乱所有节点的顺序
	// 随机打乱原来的顺序
	settings.Rand.Shuffle(len(candidates), func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})

	for _, candidate := range candidates {
		// 如果该节点(路由表，1个节点一个)的出度小于该节点(路由表)的最大连接数
		// 就将该节点加入到路由表中
		if bct.outbound.Size() < bct.GetNumConnection() {
			if v, ok := Simulator.GetSimulatedNodes().Get(candidate); ok {
				bct.AddNeighbor(v.(*Node))
			}
		} else {
			// 出度表中满，就立刻停止循环
			break
		}
	}
}

/**
* Adds the provided node to the list of outbound connections of self node.The provided node
* will not be added if it is the self node, it exists as an outbound connection of the self node,
* it exists as an inbound connection of the self node or the self node does not allow for
* additional outbound connections. Otherwise, the self node will add the provided node to the
* list of outbound connections and the provided node will add the self node to the list of
* inbound connections.
*
* @param node the node to be connected to the self node.
* @return the success state
 */
func (bct *BitcoinCoreTable) AddNeighbor(node *Node) bool {
	if node == bct.GetSelfNode() || bct.outbound.Contains(node) || bct.inbound.Contains(
		// 如果该节点是 自身、出口表已存在(你已经加别人了)、入口表已存在(别人已经加你了)
		// 、你的出口表已经满了(之前已经判定过了) 则返回添加失败
		node) || bct.outbound.Size() >= bct.GetNumConnection() {
		return false
	} else if bct.outbound.Add(node); node.getRoutingTable().AddInbound(bct.GetSelfNode()) {
		// 否则该节点的出口表添加节点，添加节点的入口表添加你，打印连接关系，返回真
		bct.printAddLink(node)
		return true
	} else {
		// 其余条件一律判断为假
		return false
	}
}

/**
 * Remove the provided node from the list of outbound connections of the self node and the
 * self node from the list inbound connections from the provided node.
 *
 * @param node the node to be disconnected from the self node.
 * @return the success state of the operation
 */
func (bct *BitcoinCoreTable) RemoveNeighbor(node *Node) bool {
	index := bct.outbound.IndexOf(node)
	if index == -1 {
		return false
	}
	bct.outbound.Remove(index)
	if node.getRoutingTable().RemoveInbound(bct.GetSelfNode()) {
		bct.printRemoveLink(node)
		return true
	}
	return false
}

/**
 * Adds the provided node as an inbound connection.
 *
 * @param from the node to be added as an inbound connection
 * @return the success state of the operation
 */
func (bct *BitcoinCoreTable) AddInbound(from *Node) bool {
	bct.inbound.Add(from)
	bct.printAddLink(from)
	return true

}

/**
 * Removes the provided node as an inbound connection.
 *
 * @param from the node to be removed as an inbound connection
 * @return the success state of the operation
 */
func (bct *BitcoinCoreTable) RemoveInbound(from *Node) bool {
	index := bct.inbound.IndexOf(from)
	if index == -1 {
		return false
	} else {
		bct.inbound.Remove(index)
		return true
	}
}

//TODO add example
func (bct *BitcoinCoreTable) printAddLink(endNode *Node) {
	printfile.OUT_JSON_FILE.Print("{")
	printfile.OUT_JSON_FILE.Print("\"kind\":\"add-link\",")
	printfile.OUT_JSON_FILE.Print("\"content\":{")
	printfile.OUT_JSON_FILE.Print("\"timestamp\":" + strconv.FormatInt(GetCurrentTime(), 10) + ",")
	printfile.OUT_JSON_FILE.Print("\"begin-node-id\":" + strconv.Itoa(bct.GetSelfNode().GetNodeID()) + ",")
	printfile.OUT_JSON_FILE.Print("\"end-node-id\":" + strconv.Itoa(endNode.GetNodeID()))
	printfile.OUT_JSON_FILE.Print("}")
	printfile.OUT_JSON_FILE.Print("},")
	printfile.OUT_JSON_FILE.Flush()
}

//TODO add example
func (bct *BitcoinCoreTable) printRemoveLink(endNode *Node) {
	printfile.OUT_JSON_FILE.Print("{")
	printfile.OUT_JSON_FILE.Print("\"kind\":\"remove-link\",")
	printfile.OUT_JSON_FILE.Print("\"content\":{")
	printfile.OUT_JSON_FILE.Print("\"timestamp\":" + strconv.FormatInt(GetCurrentTime(), 10) + ",")
	printfile.OUT_JSON_FILE.Print("\"begin-node-id\":" + strconv.Itoa(bct.GetSelfNode().GetNodeID()) + ",")
	printfile.OUT_JSON_FILE.Print("\"end-node-id\":" + strconv.Itoa(endNode.GetNodeID()))
	printfile.OUT_JSON_FILE.Print("}")
	printfile.OUT_JSON_FILE.Print("},")
	printfile.OUT_JSON_FILE.Flush()
	//printfile.OUT_JSON_FILE.F();
}
