// package task defines the format of tasks in FCAS

package task

import "fcas/container"

type Task struct {
	id uint64

	// every task needs to be executed with one type of container
	t_type container.ContainerType
}

func NewTask(id uint64, t_type container.ContainerType) *Task {
	return &Task{
		id:     id,
		t_type: t_type,
	}
}

func (t *Task) ID() uint64 {
	return t.id
}

func (t *Task) Type() container.ContainerType {
	return t.t_type
}
