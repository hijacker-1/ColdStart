package tracker

import (
	"errors"
	"fcas/container"
	"fcas/node"
)

// 记录所有节点的信息
// key: all nodes
// value: all containers info in every node
var (
	NodeRecord       map[*node.Node]*nodeInfo
	AllContainerInfo *Tracker = NewTracker()
)

type Tracker struct {
	record map[container.ContainerType]*containerInfo
}

func (t *Tracker) GetContainersRecord() *map[container.ContainerType]*containerInfo {
	return &(t.record)
}

// First called in starting system
func InitNodeRecord() {
	NodeRecord = make(map[*node.Node]*nodeInfo)
}

// Add all nodes into NodeRecord to maintain nodes information
func AddNode(n *node.Node) {
	NodeRecord[n] = &nodeInfo{
		node: n,
	}
}

func DeleteNode(n *node.Node) {
	delete(NodeRecord, n)
}

// nodeInfo stores information about a node and its containers
type nodeInfo struct {
	node *node.Node // Pointer to the node object

	// Map of container type to container info
	ContainerInfo map[container.ContainerType]*containerInfo
}

// containerInfo stores information about a container
type containerInfo struct {
	ContainerType container.ContainerType
	Running       uint32 // The state of running containers' number about container_type
	Paused        uint32 // The state of paused containers' number about container_type
}

func NewTracker() *Tracker {
	return &Tracker{
		record: make(map[container.ContainerType]*containerInfo),
	}
}

func (t *Tracker) AddContainer(cc *container.Container, n *node.Node) {
	node := NodeRecord[n]
	if container_info, has := node.ContainerInfo[cc.Type()]; has {
		container_info.Paused++
	} else {
		node.ContainerInfo[cc.Type()] = &containerInfo{
			ContainerType: cc.Type(),
			Paused:        1,
			Running:       0,
		}
	}
	t.addContainer(cc)
}

func (t *Tracker) DeleteContainer(cc *container.Container, n *node.Node) error {
	node := NodeRecord[n]
	if container_info, has := node.ContainerInfo[cc.Type()]; has {
		if container_info.Paused >= 1 {
			container_info.Paused--
		} else {
			return errors.New("no paused container found")
		}
	} else {
		return errors.New("no container info found")
	}
	return t.deleteContainer(cc)
}

func (t *Tracker) ReuseContainer(cc *container.Container, n *node.Node) error {
	node := NodeRecord[n]
	if container_info, has := node.ContainerInfo[cc.Type()]; has {
		container_info.Running++
		container_info.Paused--
	} else {
		return errors.New("no container info found")
	}
	return t.reuseContainer(cc)
}

func (t *Tracker) StopContainer(cc *container.Container, n *node.Node) error {
	node := NodeRecord[n]
	if container_info, has := node.ContainerInfo[cc.Type()]; has {
		container_info.Running--
		container_info.Paused++
	} else {
		return errors.New("no container info found")
	}
	return t.stopContainer(cc)
}

// addContainer increments the paused count of a container and keeps track of container types.
// If the container type is already recorded, the paused count is incremented.
// If the container type is not yet recorded, a new entry is created with the container type and a paused count of 1.
func (t *Tracker) addContainer(cc *container.Container) {
	containerType := cc.Type()
	if container, exists := t.record[containerType]; exists {
		container.Paused++ // How to change the state of the container to running?
	} else {
		t.record[containerType] = &containerInfo{
			ContainerType: containerType,
			Paused:        1,
		}
	}
}

// deleteContainer decrements the paused count of a container and keeps track of container types.
func (t *Tracker) deleteContainer(cc *container.Container) error {
	containerType := cc.Type()
	if container, exists := t.record[containerType]; exists {
		container.Paused-- // Decrement the paused count
		if container.Paused == 0 {
			delete(t.record, containerType)
		}
		return nil
	}
	return errors.New("container not found")
}

func (t *Tracker) reuseContainer(cc *container.Container) error {
	containerType := cc.Type()
	if container, exists := t.record[containerType]; exists {
		container.Running++
		return nil
	}
	return errors.New("no paused container found")
}

// stopContainer stops a container and updates the tracker record.
func (t *Tracker) stopContainer(cc *container.Container) error {
	// Get the container type.
	containerType := cc.Type()

	// Check if the container exists in the tracker record.
	if container, exists := t.record[containerType]; exists {
		container.Running--
		container.Paused++
		return nil
	}

	// Return an error if no running container is found.
	return errors.New("no running container found")
}
