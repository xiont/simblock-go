package settings

/**
 * The number of nodes participating in the blockchain network.
 */
//TODO revert
var NUM_OF_NODES int = 6000 //600;//800;//6000;
// public static final int NUM_OF_NODES = 600;//600;//800;//6000;

/**
 * The kind of routing table.
 */
//public static final String TABLE = "simblock.node.routing.BitcoinCoreTable";

/**
 * The consensus algorithm to be used.
 */
//TODO not documented in markdown
// TODO return to PoW
//public static final String ALGO = "simblock.node.consensus.ProofOfWork";

/**
 * The expected value of block generation interval. The difficulty of mining is automatically
 * adjusted by this value and the sum of mining power. (unit: millisecond)
 */
var INTERVAL int64 = 1000 * 60 * 10 //1000*60;//1000*30*5;//1000*60*10;

/**
 * The average mining power of each node. Mining power corresponds to Hash Rate in Bitcoin, and
 * is the number of mining (hash calculation) executed per millisecond.
 */
var AVERAGE_MINING_POWER int = 400000

/**
 * The mining power of each node is determined randomly according to the normal distribution
 * whose average is AVERAGE_MINING_POWER and standard deviation is STDEV_OF_MINING_POWER.
 */
var STDEV_OF_MINING_POWER int = 100000

/**
 * The constant AVERAGE_COINS.
 */
//TODO
var AVERAGE_COINS int = 4000

/**
 * The constant STDEV_OF_COINS.
 */
//TODO
var STDEV_OF_COINS int = 2000

/**
 * The reward a PoS minter gets for staking.
 */
var STAKING_REWARD float64 = 0.01

/**
 * The block height when a simulation ends.
 */
//TODO revert
//public static final int END_BLOCK_HEIGHT = 100;
var END_BLOCK_HEIGHT int = 300

/**
 * Block size. (unit: byte).
 */
var BLOCK_SIZE int64 = 535000 //6110;//8000;//535000;//0.5MB
