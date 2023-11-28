// package scheduler implements a simple definition and functions for
// FCAS's scheduler algorithm for scheduling containers to nodes
//
// Decide whether a container can be scheduled to a node based on tracker's record
// and predictor's result

package scheduler

import (
	"container/list"
	"fcas/node"
	"fcas/task"
	"fcas/tracker"
)

var (
	WaitingQueue list.List
)

func chooseNode(task *task.Task) (*node.Node, bool) {
	taskType := task.Type()
	for node, nodeInfo := range tracker.NodeRecord {
		// 如果工作节点上有空闲的该类型容器，直接返回该节点
		if nodeInfo.ContainerInfo[taskType].Paused > 0 {
			return node, true
		}

		// 获得工作节点集合中
	}

	// 如果工作节点的空闲资源足够用来创建该类型的容器，返回该节点
	for node := range tracker.NodeRecord {
		if node.Memory() >= task.Type().Memory() {
			return node, false
		}
	}

	// 如果没有拥有足够资源的节点，通过释放空闲容器可以满足计算要求，返回该节点
	for node := range tracker.NodeRecord {
		memorySum := node.Memory()
		for _, ctn := range node.GetListOfReleased() {
			memorySum += ctn.MemoryUsed()
			if memorySum >= task.Type().Memory() {
				return node, false
			}
		}
	}
	// 如果以上条件均不满足，返回nil代替随机工作节点，由Schedule函数决定
	return nil, false
}

func schedule(task *task.Task, node *node.Node) {
	// TODO
}

func Schedule(task *task.Task) {
	node, flag := chooseNode(task)
	if flag {
		// 直接将任务分配给node节点
		schedule(task, node)
	} else {
		// 将任务加入等待队列, 直到有空闲的可用容器出现
		WaitingQueue.PushBack(task)
	}
}

// 不断扫描等待列表，出现可用容器直接分配
func ScanWaitingQueueList() {
	// TODO
	for {
		for begin := WaitingQueue.Front(); begin != nil; begin = begin.Next() {
			task := begin.Value.(*task.Task)
			containersInfo := *(tracker.AllContainerInfo.GetContainersRecord())
			cInfo := containersInfo[task.Type()]
			if cInfo.Paused >= 1 {
				go Schedule(task)
				WaitingQueue.Remove(begin)
			}
		}
	}
}
