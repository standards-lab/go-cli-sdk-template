# Changelog

All notable changes to the `github.com/standards-lab/go-cli-sdk-template/template` module are
documented here. The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and the module adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html). A
generated CLI starts its own changelog; this one records the template's.

## [Unreleased]

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

[Unreleased]: https://github.com/standards-lab/go-cli-sdk-template/compare/template/v0.1.0...HEAD
[v0.1.0]: https://github.com/standards-lab/go-cli-sdk-template/releases/tag/template/v0.1.0
