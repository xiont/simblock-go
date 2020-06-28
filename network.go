package main

import (
	"math"
	"math/rand"
	"simblock-go/printfile"
	"simblock-go/settings"
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
	return int64(math.Round(scale / math.Pow(rand.Float64(), 1.0/shape)))
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

//TODO
func PrintRegion() {
	// 打印区域信息
	printfile.STATIC_JSON_FILE.Write("{\"region\":[")

	id := 0
	for ; id < len(settings.REGION_LIST)-1; id++ {
		printfile.STATIC_JSON_FILE.Write("{")
		printfile.STATIC_JSON_FILE.Write("\"id\":" + string(id) + ",")
		printfile.STATIC_JSON_FILE.Write("\"name\":\"" + settings.REGION_LIST[id] + "\"")
		printfile.STATIC_JSON_FILE.Write("},")
	}

	printfile.STATIC_JSON_FILE.Write("{")
	printfile.STATIC_JSON_FILE.Write("\"id\":" + string(id) + ",")
	printfile.STATIC_JSON_FILE.Write("\"name\":\"" + settings.REGION_LIST[id] + "\"")
	printfile.STATIC_JSON_FILE.Write("}")
	printfile.STATIC_JSON_FILE.Write("]}")
	//STATIC_JSON_FILE.flush();
	//STATIC_JSON_FILE.close();
}
