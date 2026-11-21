package app

import (
	"github.com/standards-lab/go-core/graph"
	"github.com/standards-lab/go-core/lifecycle"
)

const envPrefix = "APP"

func defineInfrastructure(g *graph.Graph, n *Nodes) {
	n.LifecycleConfig = g.Define("lifecycle config", newLifecycleConfig)
}

func newLifecycleConfig(*graph.Scope) (lifecycle.Config, error) {
	var cfg lifecycle.Config
	err := cfg.Finalize(envPrefix)
	return cfg, err
}
