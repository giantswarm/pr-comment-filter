# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

### Fixed

- Correctly check for draft

## [0.7.0] - 2024-07-22

### Added

- Handle draft PRs and if triggered from opening add a comment stating that the trigger was ignored

## [0.6.0] - 2024-07-18

## [0.5.0] - 2024-04-05

### Added

- Added gsociprivate-pull-secret imagePullSecret to task run template

## [0.4.1] - 2024-04-04

### Changed

- Update RegEx to support optional whitespace infront of `/run` commands

## [0.4.0] - 2024-03-11

### Changed

- Updated Tekton Pipelines to latest release
- Updated the base Dockerfile to Go v1.22

## [0.3.0] - 2024-02-05

### Changed

- Added `securityContext` to task pod

## [0.2.0] - 2023-12-08

## [0.1.0] - 2023-10-23

## [0.0.1] - 2023-10-18


[Unreleased]: https://github.com/giantswarm/pr-comment-filter/compare/v0.7.0...HEAD
[0.7.0]: https://github.com/giantswarm/pr-comment-filter/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/giantswarm/pr-comment-filter/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/giantswarm/pr-comment-filter/compare/v0.4.1...v0.5.0
[0.4.1]: https://github.com/giantswarm/pr-comment-filter/compare/v0.4.0...v0.4.1
[0.4.0]: https://github.com/giantswarm/pr-comment-filter/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/pr-comment-filter/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/pr-comment-filter/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/pr-comment-filter/compare/v0.0.1...v0.1.0
[0.0.1]: https://github.com/giantswarm/pr-comment-filter/releases/tag/v0.0.1
