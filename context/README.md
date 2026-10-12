# go-cli-sdk-template

go-cli-sdk-template is the CLI template of Go Elemental, the Standards Lab organization's Go
implementation of the Elemental Architecture. It scaffolds a CLI with `gonew`, built on go-core
and go-cli-sdk at pinned releases, with its module rooted at `template/`.

The README, the starter README, and each package's `doc.go` document this repository. The
[Go Elemental](https://github.com/standards-lab/architecture/blob/main/standards/go-elemental/README.md) standard states the principles it follows.
This context records only working knowledge the code and the READMEs do not express.

## Capability map

Built as of `template/v0.1.0`; the code and each package's `doc.go` are authoritative.

- **Runnable baseline** (built): the `cmd` entrypoint, the composition root `internal/app` with
  configuration as graph nodes, and a `version` command.
- **Integration tier** (built): fixtures in `apptest`, and an integration suite on
  `processtest`'s one-shot runner.
- **Generation** (built): the constraints that keep the module cleanly copyable with `gonew`:
  the subtree boundary, every file surviving the path rewrite, and the starter README's identity
  steps.
- **CI and release** (built): CI runs `mise run check` inside `template/`; releases are
  `template/v*` tags cut from the root `CHANGELOG.md`.

Scenarios are opt-in: go-cli-sdk's `scenario` package carries them, and the starter README and
`internal/app`'s `doc.go` describe how a generated CLI mounts its own. The template ships none.
