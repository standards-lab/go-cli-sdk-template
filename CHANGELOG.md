# Changelog

All notable changes to the `github.com/standards-lab/go-cli-sdk-template/template` module are
documented here. The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and the module adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html). A
generated CLI starts its own changelog; this one records the template's.

## [Unreleased]

## [v0.2.0] - 2026-11-11

The template moves onto go-core v0.11.0 and go-cli-sdk v0.3.0: the command tree runs as a
`cli.Program`, each node is named for its `Nodes` field, and graph tests check those names and
every leaf command's layer order. The module depends on
`github.com/standards-lab/go-core v0.11.0` and `github.com/standards-lab/go-cli-sdk v0.3.0`.

A generated CLI that ports this release from v0.1.0 also takes the libraries' own changes since
its previous pins; the [go-core](https://github.com/standards-lab/go-core/blob/main/CHANGELOG.md)
and [go-cli-sdk](https://github.com/standards-lab/go-cli-sdk/blob/main/CHANGELOG.md) CHANGELOGs
list them, and go-core's
[Migrating from v0.10.0](https://github.com/standards-lab/go-core/blob/main/CHANGELOG.md#migrating-from-v0100)
covers the graph a CLI defines. Porting meets these:

- `cli.Run(ctx, root, args, streams, cli.WithGraph(g, cfg))` becomes
  `cli.Program{Root: root, Graph: g, LifecycleConfig: cfg}.Run(ctx, args, streams)`;
  `cli.Run` and `cli.WithGraph` are removed. `App.Run` makes that call.
- Each node is named for its `Nodes` field, its words lowercased and spaced, so
  `LifecycleConfig` holds `"lifecycle config"`, and no comment repeats the name.
  `TestGraph_NodesNamedByTheirFields` checks the rule with go-core's `graphtest.CheckNames`; a
  node whose name differs from its field fails it.
- `TestGraph_LeafLayers` builds every leaf command, as the root's `Command.Leaves` yields it,
  from the nodes its path declares, and fails on a dropped `graph.After` with the layer table.
  A command mounted later is checked without a change to the test.
- A `Monitored` value's failure while a command runs is reported on one line,
  `<path>: run: <node>: <error>`, with no `context canceled` line before it (go-cli-sdk
  v0.3.0). A test that matches the old stderr changes.

### Added

- `internal/app/graph_test.go`: `TestGraph_NodesNamedByTheirFields`, and
  `TestGraph_LeafLayers` with `TestGraph_LeafLayersGuardFails`, which proves the layer check
  fails on a dropped order.

### Changed

- **Breaking:** the go-core pin moves to v0.11.0 (from v0.8.0) and the go-cli-sdk pin to v0.3.0
  (from v0.2.0), with the porting list above.
- **Breaking** for a CLI that ports it: `App.Run` dispatches through `cli.Program` in place of
  `cli.Run` with `cli.WithGraph`.
- The `Nodes` fields carry no name comments; `internal/app/doc.go`, both READMEs, and `App.Run`'s
  comment describe the field-name rule, the graph tests, and `cli.Program`. An App still runs
  one command at a time, now because each dispatch restores the flags the previous one set.
- `apptest.Builds` reads each built node's name from the `graph.Ref` that `Graph.Observe`
  reports.

## [v0.1.0] - 2026-10-12

The first release. A generated CLI starts as a minimal runnable command-line program on
`github.com/standards-lab/go-core v0.8.0` and `github.com/standards-lab/go-cli-sdk v0.2.0`, with
`mise.toml` pinning Go 1.27.2.

### Added

- `cmd/app`, the entrypoint: the signal context and one call into the composition root.
- `internal/app`, the composition root: `New` describes go-core's dependency graph with no I/O
  and builds the command tree on go-cli-sdk, one file per layer (`infrastructure.go`,
  `domain.go`) and the `version` command. The lifecycle configuration is a graph node read from
  `APP_SHUTDOWN_TIMEOUT` and validated before a command runs; each command builds only the nodes
  it declares with `Use`. `App.Graph`, `App.Nodes`, and `App.Root` publish the composition.
- `internal/apptest`, the fixtures internal/app's tests share: `Recorder`, `Builds`, `Halt`, and
  `ErrHalted`.
- `integration`, the integration tier: a harness that builds `cmd/app` once and runs it one shot
  at a time on go-core's `processtest`, and the suite under the `integration` build tag.
- The mise tasks (`check`, `integration`, `currency`, `upgrade`, and the plain-command
  wrappers), the CI workflow, and the starter README with the identity steps and the scenario
  opt-in.

[Unreleased]: https://github.com/standards-lab/go-cli-sdk-template/compare/template/v0.2.0...HEAD
[v0.2.0]: https://github.com/standards-lab/go-cli-sdk-template/compare/template/v0.1.0...template/v0.2.0
[v0.1.0]: https://github.com/standards-lab/go-cli-sdk-template/releases/tag/template/v0.1.0
