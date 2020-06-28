package main

import (
	"math/rand"
	"simblock-go/printfile"
	"simblock-go/settings"
)

func main() {
	/**
	 * The constant to be used as the simulation seed.
	 */
	rand.Seed(10)

	//var simulationTime int64 = 0

	Simulator.SetTargetInterval(settings.INTERVAL)

	//start json format
	// 写入可视化信息的文件
	printfile.OUT_JSON_FILE.Write("[")
	//OUT_JSON_FILE.flush()

	// Log regions
	// 写入能产生节点位置的地域信息
	PrintRegion()

	/**
	 * The initial simulation time.
	 */
	//var simulationTime int64 = 0

	printfile.OUT_JSON_FILE.Write("aa")

}
