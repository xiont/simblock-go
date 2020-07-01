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
	"math"
	"simblock-go/printfile"
	"simblock-go/settings"
	"strconv"
)

/**
 * Gets latency according with 20% variance pallet distribution.
 *
 * @param from the from latency
 * @param to   the to latency
 * @return the calculated latency
 */

func GetLatency(from int, to int) int64 {
	// 延迟的单位是毫秒
	var mean int64 = settings.LATENCY[from][to]
	var shape float64 = 0.2 * float64(mean)
	var scale float64 = float64(mean) - 5
	// 进行了分布的调整 20% variance pallet distribution，不太清楚是什么分布
	return int64(math.Round(scale / math.Pow(settings.Rand.NextFloat64(), 1.0/shape)))
}

/**
 * Gets the minimum between the <em>from</em> upload bandwidth and <em>to</em> download
 * bandwidth.
 *
 * @param from the from index in the {@link NetworkConfiguration#UPLOAD_BANDWIDTH} array.
 * @param to   the to index in the {@link NetworkConfiguration#UPLOAD_BANDWIDTH} array.
 * @return the bandwidth
 */

func GetBandwidth(from int, to int) int64 {
	// 获取带宽，取两者最小
	// 单位  bit/second
	if settings.UPLOAD_BANDWIDTH[from] < settings.DOWNLOAD_BANDWIDTH[to] {
		return settings.UPLOAD_BANDWIDTH[from]
	} else {
		return settings.DOWNLOAD_BANDWIDTH[to]
	}
}

/**
 * Gets region list.
 *
 * @return the {@link NetworkConfiguration#REGION_LIST} list.
 */
func getRegionList() []string {
	return settings.REGION_LIST
}

/**
 * Return the number of nodes in the corresponding region as a portion the number of all nodes.
 *
 * @return an array the distribution
 */
func GetRegionDistribution() []float64 {
	// 节点区域的分布函数
	return settings.REGION_DISTRIBUTION
}

/**
 * Get degree distribution double [ ].
 *
 * @return the double [ ]
 */
//TODO
func GetDegreeDistribution() []float64 {
	// 节点度分布
	return settings.DEGREE_DISTRIBUTION
}

//TODO
func PrintRegion() {
	// 打印区域信息
	printfile.STATIC_JSON_FILE.Print("{\"region\":[")

	id := 0
	for ; id < len(settings.REGION_LIST)-1; id++ {
		printfile.STATIC_JSON_FILE.Print("{")
		printfile.STATIC_JSON_FILE.Print("\"id\":" + strconv.Itoa(id) + ",")
		printfile.STATIC_JSON_FILE.Print("\"name\":\"" + settings.REGION_LIST[id] + "\"")
		printfile.STATIC_JSON_FILE.Print("},")
		printfile.STATIC_JSON_FILE.Flush()
	}

	printfile.STATIC_JSON_FILE.Print("{")
	printfile.STATIC_JSON_FILE.Print("\"id\":" + strconv.Itoa(id) + ",")
	printfile.STATIC_JSON_FILE.Print("\"name\":\"" + settings.REGION_LIST[id] + "\"")
	printfile.STATIC_JSON_FILE.Print("}")
	printfile.STATIC_JSON_FILE.Print("]}")
	printfile.STATIC_JSON_FILE.Flush()
	printfile.STATIC_JSON_FILE.Close()
	//STATIC_JSON_FILE.flush();
	//STATIC_JSON_FILE.close();
}
