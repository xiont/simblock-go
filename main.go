package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/sets/hashset"
	"math"
	"math/rand"
	"reflect"
	"simblock-go/printfile"
	"simblock-go/settings"
	"simblock-go/simulator"
	"strconv"
	"time"
)

func main() {

	simulator.SimulatorInit()
	/**
	 * The constant to be used as the simulation seed.
	 */
	rand.Seed(rand.Int63())

	simulationTime := time.Now() // get current time

	simulator.Simulator.SetTargetInterval(settings.INTERVAL)

	//start json format
	// 写入可视化信息的文件
	printfile.OUT_JSON_FILE.Print("[")
	//OUT_JSON_FILE.flush()

	// Log regions
	// 写入能产生节点位置的地域信息
	simulator.PrintRegion()

	// Setup network
	// 根据节点个数初始化节点网络
	constructNetworkWithAllNodes(settings.NUM_OF_NODES)

	// Initial block height, we stop at END_BLOCK_HEIGHT
	currentBlockHeight := 1

	// Iterate over tasks and handle
	// Timer.getTask  此处开始跑TASK
	for simulator.GetTask() != nil {

		// 如果是挖矿任务
		// TODO AbstractMintingTask
		if task, ok := simulator.GetTask().(simulator.IAbstractMintingTask); ok {
			// 获取权值最小的挖矿任务
			if task.GetParent() != nil {

				if task.GetParent().GetHeight() == currentBlockHeight {
					// 如果获取到的挖矿任务所属区块的高度等于当前高度，则更新当前高度
					currentBlockHeight++
				}
			}

			if currentBlockHeight > settings.END_BLOCK_HEIGHT {
				// 如果当前高度大于模拟器终止的高度，那么停止循环
				break
			}
			// Log every 100 blocks and at the second block
			// TODO use constants here
			// 日志每100个区块以及在第二个区块时画图
			if currentBlockHeight%100 == 0 || currentBlockHeight == 2 {
				writeGraph(currentBlockHeight)
			}
		}
		// Execute task
		// Timer.runTask
		simulator.RunTask()
	}

	// Print propagation information about all blocks
	// 此处打印在标准输出窗口，应该打印在日志中，这里其实是打印剩余的区块
	simulator.Simulator.PrintAllPropagation()

	println()

	//<IBLOCK>
	blocks := hashset.New()

	// Get the latest block from the first simulated node
	nodeTemp, _ := simulator.Simulator.GetSimulatedNodes().Get(0)
	block := nodeTemp.(*simulator.Node).GetBlock()

	//Update the list of known blocks by adding the parents of the aforementioned block
	//通过添加上述区块的父代，更新已知区块列表。
	for !reflect.ValueOf(block.GetParent()).IsNil() {
		blocks.Add(block)
		block = block.GetParent()
	}

	//<IBlock>
	orphans := hashset.New()
	averageOrphansSize := 0
	// Gather all known orphans
	simulator.Simulator.GetSimulatedNodes().Each(func(index int, value interface{}) {
		for _, item := range value.(*simulator.Node).GetOrphans().Values() {
			orphans.Add(item)
		}
		averageOrphansSize += value.(*simulator.Node).GetOrphans().Size()
	})
	println("all discovered orphans : " + strconv.Itoa(averageOrphansSize))
	println("SimulatedNode nums:" + strconv.Itoa(simulator.Simulator.GetSimulatedNodes().Size()))
	//当前高度下，平均每个节点发现的孤块数
	//averageOrphansSize = averageOrphansSize / simulator.Simulator.GetSimulatedNodes().Size()

	// Record orphans to the list of all known blocks
	// 将所有的orphan添加到HashSet中
	for _, item := range orphans.Values() {
		blocks.Add(item)
	}

	//<IBlock>
	blockList := arraylist.New()
	for _, block := range blocks.Values() {
		blockList.Add(block)
	}
	//Sort the blocks first by time, then by hash code
	// 排序，首先按照时间排，时间相同按照hash排

	blockList.Sort(func(a interface{}, b interface{}) int {
		if a.(simulator.IBlock).GetTime() > b.(simulator.IBlock).GetTime() {
			return 1
		} else {
			return -1
		}
	})

	//Log all orphans
	// TODO move to method and use logger
	println("Orphan Information:")
	for _, orphan := range orphans.Values() {
		// 打印出所有孤块的高度信息
		println(orphan.(simulator.IBlock).GetUUID() + ":" + strconv.Itoa(orphan.(simulator.IBlock).GetHeight()))
	}
	// 平均每个节点上的孤块数量
	//当前高度下，平均每个节点发现的孤块数
	println(fmt.Sprintf("Current Height %d : All discovered Orphan nums(%d)/SimulatedNode nums(%d) = %f",
		settings.END_BLOCK_HEIGHT, averageOrphansSize, simulator.Simulator.GetSimulatedNodes().Size(),
		float64(averageOrphansSize)/float64(simulator.Simulator.GetSimulatedNodes().Size())))

	/*
	   Log in format:
	    ＜fork_information, block height, block ID＞
	   fork_information: One of "OnChain" and "Orphan". "OnChain" denote block is on Main chain.
	   "Orphan" denote block is an orphan block.
	*/
	blockList.Each(func(index int, value interface{}) {
		if orphans.Contains(value) {
			printfile.BLOCK_LIST.Println(fmt.Sprintf("Orphan : %d : %s : %d",
				value.(simulator.IBlock).GetHeight(), value.(simulator.IBlock).GetUUID(), value.(simulator.IBlock).GetTime()))
		} else {
			printfile.BLOCK_LIST.Println(fmt.Sprintf("OnChain : %d : %s : %d",
				value.(simulator.IBlock).GetHeight(), value.(simulator.IBlock).GetUUID(), value.(simulator.IBlock).GetTime()))
		}
	})
	printfile.BLOCK_LIST.Flush()
	printfile.BLOCK_LIST.Close()

	printfile.OUT_JSON_FILE.Print("{")
	printfile.OUT_JSON_FILE.Print("\"kind\":\"simulation-end\",")
	printfile.OUT_JSON_FILE.Print("\"content\":{")
	printfile.OUT_JSON_FILE.Print("\"timestamp\":" + strconv.FormatInt(simulator.GetCurrentTime(), 10))
	printfile.OUT_JSON_FILE.Print("}")
	printfile.OUT_JSON_FILE.Print("}")
	//end json format
	printfile.OUT_JSON_FILE.Print("]")
	printfile.OUT_JSON_FILE.Flush()
	printfile.OUT_JSON_FILE.Close()

	elapsedTime := time.Since(simulationTime)
	// Log simulation time in milliseconds
	println(fmt.Sprintf("Simulation Time（ms）: %d", elapsedTime/time.Millisecond))
	/**
	 * The initial simulation time.
	 */
	//var simulationTime int64 = 0

}

// TODO
// TRANSLATED FROM ABOVE STATEMENT
// The following initial generation will load the scenario
// Create a task to join the node (separate the task of joining the node and the task of
// starting to paste the link)
// Add the above participating tasks with a timer in the scenario file.

/**
 * Populate the list using the distribution.
 *
 * @param distribution the distribution
 * @param facum        - what is this?
 * @return array list
 */
//TODO explanation on facum etc.
func makeRandomList(distribution []float64, facum bool) []int {

	var arrayList []int
	index := 0

	if facum {
		for ; index < len(distribution); index++ {
			for len(arrayList) <= int(float64(settings.NUM_OF_NODES)*distribution[index]) {
				arrayList = append(arrayList, index)
			}
		}
		for len(arrayList) < settings.NUM_OF_NODES {
			arrayList = append(arrayList, index)
		}
	} else {
		// 根据分布和节点总数,产生节点列表
		// input： [0.4，0.6]  10
		// output: [0,0,0,0,1,1,1,1,1,1]
		acumulative := 0.0
		for ; index < len(distribution); index++ {
			acumulative += distribution[index]
			for len(arrayList) <= int(float64(settings.NUM_OF_NODES)*acumulative) {
				arrayList = append(arrayList, index)
			}
		}
		for len(arrayList) < settings.NUM_OF_NODES {
			arrayList = append(arrayList, index)
		}
	}

	// 随机打乱原来的顺序
	settings.Rand.Shuffle(len(arrayList), func(i, j int) {
		arrayList[i], arrayList[j] = arrayList[j], arrayList[i]
	})
	//sort.Slice(arrayList , func( i int,j int) bool{
	//	//return arrayList[i] > arrayList[j]
	//	return rand.Shuffle()
	//})

	return arrayList
}

/**
 * Construct network with the provided number of nodes.
 *
 * @param numNodes the num nodes
 */
func constructNetworkWithAllNodes(numNodes int) {

	// Random distribution of nodes per region
	// 获取每个地域的节点分布大小，总和为1
	// 根据分布和节点总数,产生节点列表
	// input： [0.4，0.6]  10
	// output: [0,0,0,0,1,1,1,1,1,1] => [0,1,1,1,0,0,1,1,0,1]
	regionDistribution := simulator.GetRegionDistribution()
	regionList := makeRandomList(regionDistribution, false)

	// Random distribution of node degrees
	// 根据度分布生成随机数
	// input： [0.4，1]  10
	// output: [0,0,0,0,1,1,1,1,1,1] => [0,1,1,1,0,0,1,1,0,1]
	degreeDistribution := simulator.GetDegreeDistribution()
	degreeList := makeRandomList(degreeDistribution, true)

	for id := 1; id <= numNodes; id++ {
		// Each node gets assigned a region, its degree, mining power, routing table and
		// consensus algorithm
		// 为每一个节点分配他们的地域，度，算力（每毫秒执行HASH的次数 ， 是由均值和标准差产生的正态分布），路由表 和 共识算法
		// 节点从1开始，度+1，地域，
		node := simulator.NewNode(
			id, degreeList[id-1]+1, regionList[id-1], genMiningPower())

		//node.SetConsensusAlgo()

		// Add the node to the list of simulated nodes
		// 将节点添加到，模拟节点列表中
		simulator.Simulator.AddNode(node)

		// 一添加节点，就打印该节点的信息。 所有初始化的timestamp都是0
		printfile.OUT_JSON_FILE.Print("{")
		printfile.OUT_JSON_FILE.Print("\"kind\":\"add-node\",")
		printfile.OUT_JSON_FILE.Print("\"content\":{")
		printfile.OUT_JSON_FILE.Print("\"timestamp\":0,")
		printfile.OUT_JSON_FILE.Print("\"node-id\":" + strconv.Itoa(id) + ",")
		printfile.OUT_JSON_FILE.Print("\"region-id\":" + strconv.Itoa(regionList[id-1]))
		printfile.OUT_JSON_FILE.Print("}")
		printfile.OUT_JSON_FILE.Print("},")
		printfile.OUT_JSON_FILE.Flush()
	}

	// Link newly generated nodes
	// 将节点列表中的节点连接起来
	// 使用的是抽象类的初始化方法，之后会调用其中具体的实现方法
	simulator.Simulator.GetSimulatedNodes().Each(func(index int, v interface{}) {
		v.(*simulator.Node).JoinNetWork()
	})

	// Designates a random node (nodes in list are randomized) to mint the genesis block
	// 随机抽取一个检点生成创世区块
	if v, ok := simulator.Simulator.GetSimulatedNodes().Get(0); ok {
		v.(*simulator.Node).GenesisBlock()
	}

	//getSimulatedNodes().get(0).genesisBlock();
}

/**
 * Generates a random mining power expressed as Hash Rate, and is the number of mining (hash
 * calculation) executed per millisecond.
 *
 * @return the number of hash  calculations executed per millisecond.
 */
func genMiningPower() int64 {
	// 生成正态分布
	r := rand.NormFloat64()
	//double r = random.nextGaussian()
	// 节点的平均算力和标准差，至少返回1， 每毫秒1次？
	return int64(math.Max(r*float64(settings.STDEV_OF_MINING_POWER)+float64(settings.AVERAGE_MINING_POWER), 1))
}

/**
 * Network information when block height is <em>blockHeight</em>, in format:
 *
 * <p><em>nodeID_1</em>, <em>nodeID_2</em>
 *
 * <p>meaning there is a connection from nodeID_1 to right nodeID_1.
 *
 * @param blockHeight the index of the graph and the current block height
 */
func writeGraph(blockHeight int) {
	// 画出每个高度处，节点的连接关系
	fw := printfile.NewFilePrinter("output/graph/" + strconv.Itoa(blockHeight) + ".txt")
	simulator.Simulator.GetSimulatedNodes().Each(func(index int, value interface{}) {
		node := value.(*simulator.Node)
		node.GetNeighbors().Each(func(index int, value interface{}) {
			neighbor := value.(*simulator.Node)
			fw.Println(strconv.Itoa(node.GetNodeID()) + " " + strconv.Itoa(neighbor.GetNodeID()))
		})

	})
}
