# go-cli-sdk-template

go-cli-sdk-template is the CLI template of Go Elemental, the Standards Lab organization's Go
implementation of the Elemental Architecture. It scaffolds a CLI with `gonew`, built on go-core
and go-cli-sdk at pinned releases, with its module rooted at `template/`.

The README, the starter README, and each package's `doc.go` document this repository. The
[Go Elemental](https://github.com/standards-lab/architecture/blob/main/standards/go-elemental/README.md) standard states the principles it follows.
This context records only working knowledge the code and the READMEs do not express.

## Capability map

- **Runnable baseline** (built): the `cmd` entrypoint, the composition root with configuration
  as graph nodes, and a `version` command
  ([template/internal/app/doc.go](../template/internal/app/doc.go)).
- **Test tiers** (built): the composition root's fixtures and an integration suite on
  `processtest`'s one-shot runner
  ([template/internal/apptest/doc.go](../template/internal/apptest/doc.go),
  [template/integration/doc.go](../template/integration/doc.go)).
- **Generation** (built): the module copies cleanly with `gonew` ([README.md](../README.md),
  [template/README.md](../template/README.md)).
- **CI and release** (built): the check and integration suite in CI, and `template/v*` tags cut
  from the root `CHANGELOG.md` ([README.md](../README.md)).
