# go-cli-sdk-template

The CLI template of the `go-elemental` standard. A generated CLI starts as a minimal runnable
command-line program on [go-core](https://github.com/standards-lab/go-core) and
[go-cli-sdk](https://github.com/standards-lab/go-cli-sdk), carrying:

- the command tree on go-cli-sdk's dispatcher, with help, usage errors, and exit codes, and a
  `version` command;
- configuration as nodes of go-core's dependency graph, read from `APP_` environment variables
  and validated before a command runs;
- the cold/hot start lifecycle: the graph is described with no I/O, and each command builds
  only the nodes it declares, starts them in layer order, and shuts them down in reverse when
  it returns;
- the composition root as one file per architecture layer (infrastructure, domain), each
  defining its layer's nodes and owning its mount;
- the composition root's test fixtures in `internal/apptest`, and an integration tier that runs
  the built binary one shot at a time on go-core's `processtest`.

New CLIs are generated from it:

```sh
# install the gonew command if needed
go install golang.org/x/tools/cmd/gonew@latest

# generate a new CLI
gonew github.com/standards-lab/go-cli-sdk-template/template@latest example.com/newcli
```

[`gonew`](https://pkg.go.dev/golang.org/x/tools/cmd/gonew) copies the module and rewrites its
path; the copy builds and passes its check from the first build. The module lives under
`template/` — that subtree is exactly what a generated CLI receives, and everything above it
(the template's own planning context, changelog, and tooling) stays behind by construction.
Versions are the `template/v*` tags; `@latest` follows the newest.

[`template/README.md`](template/README.md) is the starter README every generated CLI receives.
It documents running the CLI, localizing its identity, and opting in to scenarios.

## Standard

`go-cli-sdk-template` is the CLI template of
[Go Elemental](https://github.com/standards-lab/architecture/blob/main/standards/go-elemental/README.md), the
minimal-dependency Go standard. This README and each package's `doc.go` document the
repository. Its repository-level principles enhance the standard's:

- The template module depends on `go-core` and `go-cli-sdk` at pinned releases, and nothing
  else. The template is engine-free: no data engine is declared and no provider is imported; a
  generated CLI chooses its providers in its own composition root.
- Everything in the module survives the `gonew` path rewrite. The module path is the only
  identity, named only by `go.mod` and import lines, and the generated copy runs from the
  first build.
- The template stays minimal and stable: it scaffolds the baseline architecture, and the SDKs
  carry the bulk of development change. A generated CLI keeps pace by updating its `go-core`
  and `go-cli-sdk` versions. Scenarios are opt-in: go-cli-sdk's `scenario` package runs them,
  and the starter README shows how a CLI mounts its own.

## Repository layout

`template/` is the module a generated CLI copies, and development happens inside it:

- `cmd/app`: the entrypoint, which derives the signal context and calls the composition root;
- `internal/app`: the composition root, one file per layer (`infrastructure.go`, `domain.go`)
  beside `app.go` and the `version` command;
- `internal/apptest`: the fixtures internal/app's tests share;
- `integration`: the harness that runs the built binary, and the suite under the
  `integration` build tag;
- `mise.toml`, `scripts/`, `.github/workflows/`: the tasks, the currency report, and CI.

The root above it maintains the template: `context/` records its design and session state, this
README documents it, `CHANGELOG.md` sources the `template/v*` release notes, the CI workflow
runs the checks inside `template/`, and the release workflow creates a GitHub release for each
`template/v*` tag. The template's `.gitignore` and CI workflow are separate copies of the root
ones; a change to either copy is applied to the other in the same commit.

## Development

```sh
cd template
mise trust && mise install
mise run check
mise run integration
```

`template/mise.toml` carries the tasks: `check` (the one read-only check CI runs),
`integration` (the integration tier), `currency` (what trails its latest), `upgrade`, and the
plain-command wrappers; `template/README.md` lists them.

## License

[Apache License 2.0](LICENSE).
