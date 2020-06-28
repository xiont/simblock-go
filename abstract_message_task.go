package main

// 检查接口方法
var _ ITask = new(AbstractMessageTask)

type IAbstractMessageTask interface {
	ITask
	GetFrom() *Node
	GetTo() *Node
}

type AbstractMessageTask struct {
	from *Node
	to   *Node
}

func NewAbstractMessageTask(from *Node, to *Node) *AbstractMessageTask {
	return &AbstractMessageTask{
		from: from,
		to:   to,
	}
}

/**
 * Get the sending node.
 *
 * @return the <em>from</em> node
 */
func (amt *AbstractMessageTask) GetFrom() *Node {
	return amt.from
}

/**
 * Get the receiving node.
 *
 * @return the <em>to</em> node
 */
func (amt *AbstractMessageTask) GetTo() *Node {
	return amt.to
}

/**
 * Get the message delay with regards to respective regions.
 *
 * @return the message sending interval
 */
func (amt *AbstractMessageTask) GetInterval() int64 {
	var latency int64 = GetLatency(amt.from.GetRegion(), amt.to.GetRegion())
	// Add 10 milliseconds here, why?  maybe some handle time
	//TODO
	return latency + 10
}

/**
 * Receive message at the <em>to</em> side.
 */
func (amt *AbstractMessageTask) Run() {
	amt.to.ReceiveMessage(amt)
}
