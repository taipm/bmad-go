# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-12-06

### Added

- Initial release of BMAD-METHOD for Go
- Core BMAD framework with agent and workflow systems
- 6 pre-configured specialized agents:
  - Developer
  - Architect
  - Product Manager
  - Test Architect
  - UX Designer
  - Scrum Master
- 3 built-in workflows:
  - workflow-init (Analysis phase)
  - workflow-quick-flow (Implementation phase)
  - workflow-method (Planning phase)
- Support for 4 development phases:
  - PhaseAnalysis
  - PhasePlanning
  - PhaseSolutioning
  - PhaseImplementation
- Context-aware workflow execution
- Custom agent registration
- Custom workflow creation and registration
- Comprehensive test suite (94.9% coverage)
- Complete documentation and examples
- GitHub Actions CI/CD workflow
- MIT License

### Features

- **Type-Safe**: Leverages Go's type system for reliability
- **Extensible**: Easy to add custom agents and workflows
- **Well-Tested**: Comprehensive test coverage
- **Zero Dependencies**: Uses only Go standard library
- **Context Support**: Built-in support for timeouts and cancellation

[1.0.0]: https://github.com/taipm/bmad-go/releases/tag/v1.0.0
