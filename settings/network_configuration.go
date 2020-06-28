package settings

/**
 * Regions where nodes can exist.
 */
var REGION_LIST = []string{"NORTH_AMERICA", "EUROPE", "SOUTH_AMERICA", "ASIA_PACIFIC", "JAPAN",
	"AUSTRALIA"}

/**
 * LATENCY[i][j] is average latency from REGION_LIST[i] to REGION_LIST[j]
 * Unit: millisecond, for year 2015
 */
var LATENCY_2015 = [][]int64{
	{36, 119, 255, 310, 154, 208},
	{119, 12, 221, 242, 266, 350},
	{255, 221, 137, 347, 256, 269},
	{310, 242, 347, 99, 172, 278},
	{154, 266, 256, 172, 9, 163},
	{208, 350, 269, 278, 163, 22},
}

/**
 * LATENCY[i][j] is average latency from REGION_LIST[i] to REGION_LIST[j]
 * Unit: millisecond, for year 2019
 */
var LATENCY_2019 = [][]int64{
	{32, 124, 184, 198, 151, 189},
	{124, 11, 227, 237, 252, 294},
	{184, 227, 88, 325, 301, 322},
	{198, 237, 325, 85, 58, 198},
	{151, 252, 301, 58, 12, 126},
	{189, 294, 322, 198, 126, 16},
}

/**
 * List of latency assigned to each region. (unit: millisecond)
 */
var LATENCY = LATENCY_2019

/**
 * List of download bandwidth assigned to each region, and last element is Inter-regional
 * bandwidth. (unit: bit per second) for year 2015
 */
var DOWNLOAD_BANDWIDTH_2015 = []int64{
	25000000, 24000000, 6500000, 10000000,
	17500000, 14000000, 6 * 1000000,
}

/**
 * List of download bandwidth assigned to each region, and last element is Inter-regional
 * bandwidth. (unit: bit per second) for year 2019
 */
var DOWNLOAD_BANDWIDTH_2019 = []int64{
	52000000, 40000000, 18000000, 22800000,
	22800000, 29900000, 6 * 1000000,
}

/**
 * List of download bandwidth assigned to each region, and last element is Inter-regional
 * bandwidth. (unit: bit per second)
 */
var DOWNLOAD_BANDWIDTH = DOWNLOAD_BANDWIDTH_2019

/**
 * List of upload bandwidth assigned to each region. (unit: bit per second), and last element
 * is Inter-regional bandwidth for year 2015
 */
var UPLOAD_BANDWIDTH_2015 = []int64{
	4700000, 8100000, 1800000, 5300000,
	3400000, 5200000, 6 * 1000000,
}

/**
 * List of upload bandwidth assigned to each region. (unit: bit per second), and last element
 * is Inter-regional bandwidth for year 2019
 */
var UPLOAD_BANDWIDTH_2019 = []int64{
	19200000, 20700000, 5800000, 15700000,
	10200000, 11300000, 6 * 1000000,
}

/**
 * List of upload bandwidth assigned to each region. (unit: bit per second), and last element
 * is Inter-regional bandwidth.
 */
var UPLOAD_BANDWIDTH = UPLOAD_BANDWIDTH_2019

/**
 * Region distribution Bitcoin 2015.
 */
var REGION_DISTRIBUTION_BITCOIN_2015 = []float64{
	0.3869, 0.5159, 0.0113,
	0.0574, 0.0119, 0.0166,
}

/**
 * Region distribution Bitcoin 2019.
 */
var REGION_DISTRIBUTION_BITCOIN_2019 = []float64{
	0.3316, 0.4998, 0.0090,
	0.1177, 0.0224, 0.0195,
}

/**
 * Region distribution Litecoin.
 */
//TODO year
var REGION_DISTRIBUTION_LITECOIN = []float64{
	0.3661, 0.4791, 0.0149, 0.1022, 0.0238, 0.0139,
}

/**
 * Region distribution Dogecoin.
 */
//TODO year
var REGION_DISTRIBUTION_DOGECOIN = []float64{
	0.3924, 0.4879, 0.0212, 0.0697, 0.0106, 0.0182,
}

/**
 * The distribution of node's region. Each value means the rate of the number of nodes in the
 * corresponding region to the number of all nodes.
 */
var REGION_DISTRIBUTION = REGION_DISTRIBUTION_BITCOIN_2019

/**
 * The cumulative distribution of number of outbound links for Bitcoin 2015.
 */
/*
  节点的度分布，）0.025的度为0，0.025的度为2，0.025的度为1 ...... 0.005的度为19 。 有些孤立的节点？
*/
var DEGREE_DISTRIBUTION_BITCOIN_2015 = []float64{
	0.025, 0.050, 0.075, 0.10, 0.20, 0.30, 0.40, 0.50, 0.60, 0.70, 0.80, 0.85, 0.90, 0.95, 0.97,
	0.97, 0.98, 0.99, 0.995, 1.0,
}

/**
 * The cumulative distribution of number of outbound links for Litecoin.
 */
//TODO year
var DEGREE_DISTRIBUTION_LITECOIN = []float64{
	0.01, 0.02, 0.04, 0.07, 0.09, 0.14, 0.20, 0.28, 0.39, 0.5, 0.6, 0.69, 0.76, 0.81, 0.85, 0.87,
	0.89, 0.92, 0.93, 1.0,
}

/**
 * The cumulative distribution of number of outbound links for Dogecoin.
 */
var DEGREE_DISTRIBUTION_DOGECOIN = []float64{
	0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 1.0, 1.00, 1.00, 1.00, 1.00, 1.00, 1.00, 1.00, 1.00,
	1.00, 1.00, 1.00, 1.0,
}

/**
 * The cumulative distribution of number of outbound links. Cf. Andrew Miller et al.,
 * "Discovering bitcoin's public topology and influential nodes", 2015.
 */
var DEGREE_DISTRIBUTION = DEGREE_DISTRIBUTION_BITCOIN_2015
