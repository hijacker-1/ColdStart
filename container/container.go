// package container implements a simple definition and functions for container
// defines relates container types
package container

import "time"

type ContainerType uint64
type ContainerState uint8

var (
	containerTypeMemoryUsed = map[ContainerType]uint32{
		0: 1024,
		1: 2048,
		2: 4096,
		3: 8192,
	}
)

func (ct ContainerType) Memory() uint32 {
	return containerTypeMemoryUsed[ct]
}

const (
	RUNNING ContainerState = 0
	PAUSED  ContainerState = 1
)

type Container struct {
	id         uint64
	cType      ContainerType
	startTime  time.Time
	memoryUsed uint32
	state      ContainerState
}

// NewContainer creates a new Container object.
//
// Parameters:
// - id: the ID of the container (uint64)
// - cType: the type of the container (ContainerType)
// - memoryUsed: the amount of memory used by the container (uint32)
//
// Returns:
// - a pointer to the newly created Container object (*Container)
func NewContainer(id uint64, cType ContainerType, memoryUsed uint32) *Container {
	return &Container{
		id:         id,
		cType:      cType,
		startTime:  time.Now(),
		memoryUsed: memoryUsed,
		state:      PAUSED,
	}
}

// Run sets the state of the container to RUNNING.
func (c *Container) Run() {
	c.state = RUNNING
}

// Pause pauses the container.
//
// No parameters.
// No return types.
func (c *Container) Pause() {
	c.state = PAUSED
}

// Reuse sets the state of the Container to RUNNING.
func (c *Container) Reuse() {
	c.state = RUNNING
}

// Delete deletes the container and returns the amount of memory used.
//
// No parameters.
// uint32 - the amount of memory used.
func (c *Container) Delete() uint32 {
	return c.MemoryUsed()
}

// ID returns the ID of the Container.
//
// It returns a uint64.
func (c Container) ID() uint64 {
	return c.id
}

// Type returns the container type.
//
// No parameters.
// Returns the container type.
func (c Container) Type() ContainerType {
	return c.cType
}

// MemoryUsed returns the amount of memory used by the Container.
//
// It does not take any parameters.
// It returns an uint32 value representing the amount of memory used.
func (c Container) MemoryUsed() uint32 {
	return c.memoryUsed
}

// State returns the current state of the container.
//
// No parameters.
// Returns the ContainerState.
func (c Container) State() ContainerState {
	return c.state
}
