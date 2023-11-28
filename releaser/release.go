// package releaser implements a simple definition and functions
// for releasing containers

package releaser

import (
	"fcas/container"
	"fcas/node"
	"fcas/predictor"
	"fcas/tracker"
)

func GetListOfReleased(node *node.Node) []*container.Container {
	rec := make(map[container.ContainerType]uint32)
	containersInfo := tracker.NodeRecord[node].ContainerInfo
	willBeUsed := predictor.Predict()
	for containerType := range containersInfo {
		if _, has := willBeUsed[containerType]; !has {
			rec[containerType] = 0
		}
	}

	result := make([]*container.Container, 0)
	for _, container := range node.Containers {
		if _, has := rec[container.Type()]; has {
			result = append(result, container)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}
