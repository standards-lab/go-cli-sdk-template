# CLI

Generated from [go-cli-sdk-template](https://github.com/standards-lab/go-cli-sdk-template).

## After generation

Four steps localize the CLI's identity:

1. Rename `cmd/app` to the command's name, and the `./cmd/app` the integration harness builds
   (`integration/harness.go`) with it.
2. Rename the root command `"app"` in `New` (`internal/app/app.go`), with its summary.
3. Rename `envPrefix` in `internal/app/infrastructure.go`, the single constant every `APP_*`
   environment-variable name derives from, and the harness's `envPrefix` with it.
4. Rewrite this README for the CLI.

Licensing and release automation are yours to define; CI arrives working
(`.github/workflows/ci.yml`).

## Getting started

[mise](https://mise.jdx.dev/) provisions the toolchain and runs the tasks:

```sh
mise trust && mise install
go run ./cmd/app            # the root help, exit code 2
go run ./cmd/app version    # (devel)
```

`go install ./cmd/app` puts the binary on your `PATH`; an installed release prints its module
version. The help (`-h`, `--help`, or no command) and a usage error read no configuration; a
command that declares graph nodes reads `APP_*` variables, such as `APP_SHUTDOWN_TIMEOUT`, and
validates them before it runs.

## Tasks

Each task wraps a plain command, so the repository works without mise:

| Task | Command | What it does |
|------|---------|--------------|
| `mise run check` | build, vet, `gofmt -l`, `go fix -diff`, `go mod tidy -diff`, `go test -race`, `golangci-lint run` | The one read-only check, as CI runs it |
| `mise run currency` | `bash scripts/currency.sh` | Report requirements, Go version, tools, and action pins behind their latest |
| `mise run upgrade` | `go mod edit -go=<current minor> -toolchain=none`, `go get <direct>@latest`, `go mod tidy`, `mise upgrade --bump --local` | Upgrade the go directive, requirements, and tools, rewriting `go.mod` and `mise.toml` |
| `mise run build` | `go build ./...` | Build the module |
| `mise run vet` | `go vet -tags integration ./...` | Compile-check and vet, the integration suite included |
| `mise run test` | `go test -race ./...` | Run the unit tier |
| `mise run integration` | `go test -race -count=1 -tags integration ./integration/` | Run the integration tier against the built CLI |
| `mise run fmt` | `gofmt -w .` | Format the source |
| `mise run tidy` | `go mod tidy` | Reconcile module requirements |
| `mise run lint` | `golangci-lint run --build-tags integration ./...` | Lint, the integration suite included |

## Testing

Two tiers. The unit tier is hermetic and runs on every pull request: `go test -race ./...`.
`internal/app`'s tests run the App in process over buffers, with the fixtures in
`internal/apptest`: `Builds` records the nodes a run builds, and `Halt` stops a run before
anything starts. The integration tier is the `integration` package: an untagged harness that
builds `cmd/app` once and runs it as a subprocess, one shot at a time, over go-core's
`process/processtest`. Under the `integration` build tag, the suite asserts the CLI through its
arguments, its two output streams, and its exit code. It runs on merge to main and on demand.

## Building out the CLI

The composition root, `internal/app`, is laid out as one file per layer of the architecture,
and those files are the build points:

- `infrastructure.go` for the configuration and the services the CLI composes on
- `domain.go` for the domain services and their commands

Each layer file defines its layer's nodes on go-core's dependency graph, in its define function
(`defineInfrastructure`, `defineDomain`), and owns its mount (`mountDomain`). The `Nodes` value
holds one handle per node. A command declares the nodes it reads with `Use`; the dispatcher
builds only those, with the lifecycle configuration, when the command runs, and shuts them down
when it returns. `cmd/app` is the entrypoint alone.

## Scenarios

A scenario is a narrated guided tour of the CLI's own commands, run as
`<cli> scenario <name>`. The template has none; go-cli-sdk's
[`scenario`](https://pkg.go.dev/github.com/standards-lab/go-cli-sdk/scenario) package runs them.
To opt in:

1. Add a root-level `scenarios` package holding the CLI's `scenario.Scenario` values, built
   over the node handles it is given.
2. Add a scenario layer file, `internal/app/scenario.go`, that mounts
   `scenario.Commands(...)` at the root with `root.Add`, and ends the root's `Footer` with
   `scenario.WriteListing` over the same scenarios, so the root help lists them too.

The `scenario` package docs state how a step reads its nodes, the narration, and when a run
stops.
