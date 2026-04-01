# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.


## v1.5.9

- Update indirect dependencies (docker, containerd, prometheus, opentelemetry)
- Add replace directives for charmbracelet/x/cellbuf, denis-tingaikin/go-header, opencontainers/runtime-spec
- Bump moby/buildkit v0.23.2 → v0.29.0, docker/docker v28.3.3 → v28.5.2
- Update go-git, klauspost/compress, prometheus stack

## v1.5.8

- fix incompatible charmbracelet/x/cellbuf dependency version

## v1.5.7

- chore: verify project health — all tests pass, linting clean, precommit succeeds

## v1.5.6

- standardize Makefile: add mocks mkdir, reorder lint, multiline trivy, add .PHONY declarations
- setup dark-factory config

## v1.5.5

- upgrade golangci-lint from v1 to v2
- update vulnerable deps (go-sdk, grpc)

## v1.5.4

- go mod update

## v1.5.3

- Update Go version from 1.25.7 to 1.26.0
- Update google/osv-scanner from v2.3.2 to v2.3.3
- Update securego/gosec from v2.22.11 to v2.23.0
- Update various indirect dependencies including anthropic-sdk-go, openai-go, and golang.org/x/* packages

## v1.5.2

- Update Go from 1.25.5 to 1.25.6
- Update ginkgo/v2 from 2.27.5 to 2.28.1
- Update gomega from 1.39.0 to 1.39.1
- Update osv-scanner/v2 and related security dependencies

## v1.5.1

- Update Go to 1.25.5
- Update dependencies
## v1.5.0

**Breaking Changes:**
- Change `HasData` interface from `map[string]string` to `map[string]any`
- Change `AddDataToError` parameter type from `map[string]string` to `map[string]any`
- Change `DataFromError` return type from `map[string]string` to `map[string]any`
- Change `AddToContext` value parameter from `string` to `any`
- Change `DataFromContext` return type from `map[string]string` to `map[string]any`

**Features:**
- Support rich error details including arrays, numbers, booleans, and nested objects
- Enable JSON error handlers to return structured data instead of comma-separated strings

## v1.4.0

- update go and deps

## v1.3.1

- Add GitHub Actions workflows for CI, code review, and Claude integration
- Add comprehensive test files for error handling patterns
- Add golangci-lint configuration
- Update dependencies
- Remove vendor directory
- Improve Makefile

## v1.3.0

- refactor
- errors.As unwrap error if not matching
- go mod update

## v1.2.0

-add errors is

## v1.1.1

-add errors is

## v1.1.0

- add errors wrap

## v1.0.0

- Initial Version
