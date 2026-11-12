# go-cli-sdk-template standards

The judgement calls the standards-reviewer applies to go-cli-sdk-template, beyond what `mise -C template run check` enforces.

- A change that alters documented behavior updates the README, `template/README.md` and the affected `doc.go` in the same change.
- The root and `template/` copies of the CI workflow differ only in the lines that locate the subtree (`-C template`, `working_directory`, `working-directory`), and the two `.gitignore` files only by the root's management-layer entries.
- `architecture/standards/go-elemental/principles/dependencies.md`: the bottom-up line and no provider in a base, across `template/go.mod` and every import under `template/`.
- `architecture/standards/go-elemental/principles/tests-and-docs.md`: the doc.go inventory of `cmd`, `internal/app` and `apptest`, and the integration suite over `processtest`'s one-shot runner.
- `architecture/standards/go-elemental/principles/topology-and-naming.md`: the module rooted at `template/`, with `cmd` and `internal/app`'s layer files.
- `architecture/standards/go-elemental/principles/release-and-ci.md`: the check and currency inside `template/`, the integration job, and `template/v<semver>` tags cut from the root `CHANGELOG.md`.
- `architecture/standards/go-elemental/principles/construction-and-execution.md`: the integration harness's `Run(t, args...)`, whose argument list is the program's and whose one `processtest.Cmd` names the run's inputs as fields; the library calls the template demonstrates, `cli.Program{Root, Graph, LifecycleConfig}.Run` in `App.Run`; and every new exported function and constructor under `template/`.
- `architecture/standards/go-elemental/principles/lifecycle-and-context.md`: `internal/app`'s graph, with configuration as nodes and each command building only the nodes it declares with `Use`.
- `architecture/principles/validation-first.md`: configuration nodes validate before a command runs.
- `architecture/principles/context-architecture.md`: the README, `template/README.md` and each `doc.go` are the homes.
