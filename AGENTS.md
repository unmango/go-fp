# AGENTS.md

This file provides guidance to coding agents when working with code in this repository.

## Project

`go.unmango.dev/fp` is a small functional programming library: `maybe`, `result`, `either`, `lazy`, `pipe`, `semigroup` and `monoid`.
It targets go1.27 and uses generic methods (`m.Map[U any](...)`), so any tool that parses the code must be go1.27 or newer.

- The library imports only the standard library. Ginkgo and Gomega are test-only.
- Every package is tested from an external `_test` package through its exported API, with Ginkgo suites plus runnable `Example` functions.
- Value types (`Maybe`, `Result`, `Either`) and `Lazy` must keep meaningful zero values.
  `Maybe` and `Either` must stay comparable with `==` when their type parameters are; `Result` compares like the error it holds.

## Commands

- `make test`: run the Ginkgo suites (`go tool ginkgo run -r` works too).
- `make check`: `nix flake check`, which evaluates the flake and checks formatting. `make build` runs the tests inside the Nix build.
- `make tidy`: after any `go.mod` change, regenerates `nix/gomod2nix.toml`. CI fails if it is stale.
- PR titles must be Conventional Commits; release-please reads them.

## Agent skills

### Issue tracker

Issues are tracked in this repo's GitHub Issues, via the `gh` CLI. See `docs/agents/issue-tracker.md`.

### Triage labels

Uses the five default triage labels (`needs-triage`, `needs-info`, `ready-for-agent`, `ready-for-human`, `wontfix`). See `docs/agents/triage-labels.md`.

### Domain docs

Single-context: one `CONTEXT.md` and `docs/adr/` at the repo root. See `docs/agents/domain.md`.
