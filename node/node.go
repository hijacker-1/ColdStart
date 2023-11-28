// package node implements a simple definition and functions for
// simulating real physical nodes
package node

import (
	"fcas/container"
)

type Reporter struct {
}

func (r Reporter) report() {
	// TODO 向追踪器报告容器被驱逐的消息
}

type Distributor struct {
}

type Node struct {
	id          uint
	memory      uint32
	containers  map[uint64]*container.Container
	reporter    Reporter
	distributor Distributor
}

func NewNode(id uint, memory uint32) *Node {
	return &Node{
		id:          id,
		memory:      memory,
		reporter:    Reporter{},
		containers:  make(map[uint64]*container.Container),
		distributor: Distributor{},
	}
}

func (n *Node) ID() uint {
	return n.id
}

func (n *Node) Memory() uint32 {
	return n.memory
}

// CreateContainer adds a new container to the node.
// It updates the node's containers map and subtracts the container's memory usage from the node's total memory.
// cc: The container to be added.
func (n *Node) CreateContainer(cc *container.Container) {
	// TODO config that if it can create a new container
	n.containers[cc.ID()] = cc
	n.memory -= cc.MemoryUsed()
}

// DeleteContainer deletes a container from the node.
// It updates the memory usage of the node, deletes the container from the list of containers,
// performs post-deletion processing, and reports the deletion to the tracker.
func (n *Node) DeleteContainer(cc *container.Container) {
	// Update the memory usage of the node
	n.memory += cc.MemoryUsed()

	// Delete the container from the list of containers
	delete(n.containers, cc.ID())

	// Perform post-deletion processing
	cc.Delete()

	// Report the deletion to the tracker
	n.reporter.report()
}

func (n *Node) ReuseContainer(cc *container.Container) {
	// TODO 重用容器
}

func (n *Node) GetListOfReleased() []*container.Container {
	// TODO 获得可以被释放的容器
	return nil
}
