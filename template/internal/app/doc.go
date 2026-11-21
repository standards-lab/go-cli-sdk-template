// Package app is the composition root. [New] describes the program's
// dependencies as one go-core graph and builds the command tree on
// go-cli-sdk over it; [App.Run] runs that tree, as a go-cli-sdk
// cli.Program, on the process arguments. Each architecture layer has one
// file, which defines the layer's nodes in its define function and owns
// its mount, so the layer files are the architecture's layer list:
//
//   - app.go: [App], [New], and [Nodes], which call the layer files'
//     define and mount functions in layer order, lowest first;
//   - infrastructure.go: defineInfrastructure, the configuration, read from
//     APP_ variables, and the services the program is composed on (the
//     lifecycle configuration alone, in the template baseline);
//   - domain.go: defineDomain and mountDomain, the domain services and
//     their commands (none, in the template);
//   - version.go: the version command, which declares no nodes and prints
//     the main module's version, or (devel) when the binary carries none, as
//     under go test.
//
// envPrefix, in infrastructure.go, is the single constant every APP_
// variable name derives from; go-core's config convention composes the
// rest, such as APP_SHUTDOWN_TIMEOUT. Every infrastructure value is read or
// constructed in its node's constructor, which runs only when a Build
// reaches the node, so configuration is read only for a command that
// declares a node, never in [New]. A service whose value implements
// lifecycle.Subsystem takes part in startup and shutdown through its own
// methods, and its constructor opens nothing: connectivity belongs to its
// Start, so a failed Build leaks no connections.
//
// Domain services are defined in the module's base-layer domain packages,
// one package per layer under domain/, with one node per layer whose
// constructor builds the service from the infrastructure nodes it Uses. A
// package that owns a resource and is built once for a command belongs in
// infrastructure.go instead. mountDomain mounts each layer's commands with
// one root.Add(pkg.Commands(...)...) per layer over the nodes it passes.
// Which services the program composes is the application author's decision.
//
// The template has no admin layer and no scenario layer by default. An
// admin layer, the services that administer an infrastructure service, goes
// in its own admin.go, defined between infrastructure and domain.
//
// Scenarios are opt-in. A CLI that adds them keeps its scenario.Scenario
// values in a root-level scenarios package and adds a scenario layer file,
// scenario.go, that mounts scenario.Commands(...) at the root and ends the
// root's Footer with scenario.WriteListing over the same scenarios. The
// scenario package of go-cli-sdk documents the runner and the narration.
//
// [Nodes] holds one handle per node: each define function fills its own
// part of one Nodes value, and its constructors read the lower layers'
// nodes from it with Use. Each field's node is named for the field, so no
// comment names it; graph_test.go checks that rule with go-core's
// graphtest.CheckNames, and checks that every leaf of the tree, as
// Command.Leaves yields it, builds its nodes without dropping an order
// between layers. Defining the nodes constructs nothing. A command
// declares the nodes it needs with Use; the dispatcher builds them, with
// the lifecycle configuration node, only when the command runs, starts
// what was built layer by layer, and shuts it down in reverse when the
// command returns. A run that builds nothing, such as help, a usage error,
// or version, reads no configuration.
//
// The App publishes its composition: [App.Graph], [App.Nodes], and
// [App.Root] return the graph, a handle on each of its nodes, and the
// command tree. The program itself never calls them; they are for a caller
// that observes what a run builds, Replaces a node's constructor with a
// substitute, or adds a command over the nodes, before the App runs.
package app
