// Package app is the composition root. [New] describes the program's
// dependencies as one go-core graph and builds the command tree on
// go-cli-sdk over it; [App.Run] dispatches the process arguments over that
// tree. Each architecture layer has one file, which defines the layer's
// nodes in its define function and owns its mount, so the layer files are
// the architecture's layer list:
//
//   - app.go: [App], [New], and [Nodes], which call the layer files'
//     define and mount functions in layer order, lowest first;
//   - infrastructure.go: defineInfrastructure, the configuration, read from
//     APP_ variables, and the services the program is composed on (the
//     lifecycle configuration alone, in the template baseline);
//   - domain.go: defineDomain and mountDomain, the domain services and
//     their commands (none, in the template);
//   - version.go: the version command, which declares no nodes.
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
// nodes from it with Use. Defining the nodes constructs nothing. A command
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
