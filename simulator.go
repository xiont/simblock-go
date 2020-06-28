package main

import (
	"container/list"
	"simblock-go/settings"
)

type simulator struct {
	/**
	 * A list of nodes that will be used in a simulation.
	 */
	simulatorNodes *list.List

	/**
	 * The target block interval in milliseconds.
	 */
	targetInterval int64
}

/**
 * Get simulated nodes list.
 */
func (sim *simulator) getSimulatedNodes() *list.List {
	return sim.simulatorNodes
}

/**
 * Get target block interval.
 */
func (sim *simulator) getInterval() int64 {
	return sim.targetInterval
}

/**
 * Sets the target block interval.
 */
func (sim *simulator) SetTargetInterval(interval int64) {
	sim.targetInterval = interval
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

}

var Simulator = simulator{
	simulatorNodes: list.New(),
	targetInterval: settings.INTERVAL,
}
