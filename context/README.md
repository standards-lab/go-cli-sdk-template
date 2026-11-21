# go-cli-sdk-template

go-cli-sdk-template is the CLI template of Go Elemental, the Standards Lab organization's Go
implementation of the Elemental Architecture. It scaffolds a CLI with `gonew`, built on go-core
and go-cli-sdk at pinned releases, with its module rooted at `template/`.

The README, the starter README, and each package's `doc.go` document this repository. The
[Go Elemental](https://github.com/standards-lab/architecture/blob/main/standards/go-elemental/README.md) standard states the principles it follows.
This context records only working knowledge the code and the READMEs do not express.

Scenarios are opt-in: go-cli-sdk's `scenario` package runs them, and the starter README and
`internal/app`'s `doc.go` describe how a generated CLI mounts its own. The template ships none.

## Capability map

- **Runnable baseline** (built): [template/internal/app/doc.go](../template/internal/app/doc.go).
- **Test tiers** (built): [template/internal/apptest/doc.go](../template/internal/apptest/doc.go), [template/integration/doc.go](../template/integration/doc.go).
- **Generation** (built): [README.md](../README.md), [template/README.md](../template/README.md).
- **CI and release** (built): [README.md](../README.md).
