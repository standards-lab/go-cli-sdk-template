package app

import (
	"github.com/standards-lab/go-core/graph"
	"github.com/standards-lab/go-core/lifecycle"
)

// envPrefix is the prefix of every environment variable the program reads
// its configuration from: go-core's config convention composes the rest,
// such as APP_SHUTDOWN_TIMEOUT.
const envPrefix = "APP"

// defineInfrastructure defines the infrastructure nodes on g into n: the
// configurations, and the services the program is composed on, built from
// them; the lifecycle configuration alone, in the template baseline. It
// constructs nothing. Every value is read or constructed in its node's
// constructor, which runs only when a Build reaches the node, so
// configuration is read only for a command that declares a node, never in
// [New] and never for a run that builds nothing. A service whose value
// implements lifecycle.Subsystem takes part in startup and shutdown through
// its own methods, and its constructor opens nothing: connectivity belongs
// to its Start, so a failed Build leaks no connections.
func defineInfrastructure(g *graph.Graph, n *Nodes) {
	n.LifecycleConfig = g.Define("lifecycle config", newLifecycleConfig)
}

// newLifecycleConfig reads the lifecycle configuration, the shutdown
// budget, from its default and APP_SHUTDOWN_TIMEOUT, and validates it. The
// dispatcher adds this node to every Build, so a command that declares any
// node reads it.
func newLifecycleConfig(*graph.Scope) (lifecycle.Config, error) {
	var cfg lifecycle.Config
	err := cfg.Finalize(envPrefix)
	return cfg, err
}
