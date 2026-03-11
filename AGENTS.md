# Project Conventions and Rules

## Operational Memory (Where Things Are)
- `cli/notify/` -> Main entry point and core logic (commands, error handling, logging).
- `cli/notify/main.go` -> CLI setup and command registration.
- `cli/notify/gntp.go` -> GNTP notification command logic.
- `cli/notify/logger.go` -> Shared logger setup.
- `cli/notify/error.go` -> Centralized error reporting.
- `.github/workflows/autorelease.yml` -> CI/CD workflow (strictly validated by 'ciborg').

## Rules

### 1. Centralized Error Reporting
- Never ignore errors. You must NEVER leave an empty catch block.
- All code paths that handle unexpected errors MUST funnel through a single, centralized error-reporting function (`reportError(err error)` in `cli/notify/error.go`).
- Never use `log.Error(err)` or `fmt.Println(err)` directly at the call site for unexpected/unrecoverable errors. Use `reportError(err)`.

### 2. CI/CD and Tooling
- Linting and formatting must be handled by `workspaced` via `mise`.
- Always pin tools in `mise.toml` to specific versions; never use 'latest' or 'lts'.
- Mise tasks (like `test`, `build`, `ci`) must ONLY depend on wildcards (e.g., `depends = ["ci:*"]`), never manually listing specific subtasks.
- Go module commands like `go test` and `go build` must be executed conditionally using full bash `if [ -f go.mod ]; then ... else ... fi` blocks in `mise.toml` to avoid swallowing errors or failing CI when the module is uninitialized.
- Commands in `mise.toml` tasks must not use `|| true` to swallow errors; they must be allowed to fail natively if the underlying code is broken so CI accurately reports the status.
- Do not downgrade dependencies (e.g., GitHub Actions versions) unless explicitly asked.
- The `.github/workflows/autorelease.yml` workflow requires exact steps: 'Install', 'Codegen', 'PR if Difference', 'CI', 'Release', 'Artifacts'.
- When using `peter-evans/create-pull-request` in GitHub Actions workflows triggered by both `pull_request` and `push` events, supply `base: ${{ github.head_ref || github.ref_name }}` to prevent detached HEAD errors.

### 3. Architecture and Structure
- Enforce directory structure: Group by domain and responsibility, not by file type.
- Things that change together should live together (colocation).
- Sparse directories (one or two files with no reason to be isolated) should be merged into their parent.
- Respect the Rule of Three: Don't abstract until duplication occurs at least three times (avoid speculative generality like premature `Plugin` interfaces for a single implementation).