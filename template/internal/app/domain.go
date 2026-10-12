package app

import (
	"github.com/standards-lab/go-cli-sdk"
	"github.com/standards-lab/go-core/graph"
)

// defineDomain defines the domain services on g into n, one node per
// domain layer. Domain services are defined in the module's base-layer
// domain packages, one package per layer under domain/, and each node's
// constructor builds its service from the infrastructure nodes it Uses.
// A package that owns a resource and is built once for a command belongs
// in infrastructure.go instead. The template defines none; which services
// the program composes is the application author's decision.
func defineDomain(g *graph.Graph, n *Nodes) {}

// mountDomain mounts each domain layer's commands at root, one
// root.Add(pkg.Commands(...)...) per layer over the nodes it passes, each
// command declaring with Use the nodes it reads, so a run builds only what
// its command needs. The template mounts none.
func mountDomain(root *cli.Command, n *Nodes) {}
