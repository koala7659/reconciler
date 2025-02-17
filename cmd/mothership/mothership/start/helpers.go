package cmd

import (
	"github.com/kyma-incubator/reconciler/pkg/keb"
	"github.com/kyma-incubator/reconciler/pkg/model"
)

func components(cfg model.ClusterConfigurationEntity) []keb.Component {
	componentsLen := len(cfg.Components)
	components := make([]keb.Component, componentsLen)

	for i := 0; i < componentsLen; i++ {
		component := cfg.Components[i]
		if component != nil {
			components[i] = *component
			continue
		}
	}
	return components
}
