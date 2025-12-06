# BMAD-METHOD for Go (Golang)

[![Go Reference](https://pkg.go.dev/badge/github.com/taipm/bmad-go.svg)](https://pkg.go.dev/github.com/taipm/bmad-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/taipm/bmad-go)](https://goreportcard.com/report/github.com/taipm/bmad-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

**Build More, Architect Dreams** (BMAD) for Go - An AI-driven agile development framework with specialized agents and workflows.

BMAD-METHOD is a revolutionary framework for human-AI collaboration in software development. This Go implementation brings the power of BMAD to the Go ecosystem, providing structured workflows and specialized agents for agile development.

## 🎯 Features

- **Specialized AI Agents**: 6 pre-configured agents (Developer, Architect, PM, Test Architect, UX Designer, Scrum Master)
- **Workflow System**: Built-in workflows for different development scenarios
- **Extensible**: Easy to add custom agents and workflows
- **Type-Safe**: Leverages Go's type system for reliability
- **Context-Aware**: Built with Go's context package for cancellation and timeouts
- **Well-Tested**: Comprehensive test coverage

## 📦 Installation

```bash
go get github.com/taipm/bmad-go
```

## 🚀 Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/taipm/bmad-go"
)

func main() {
    // Create a new BMAD instance
    b := bmad.New()

    // List available agents
    agents := b.ListAgents()
    for _, agent := range agents {
        fmt.Printf("Agent: %s - %s\n", agent.Role, agent.Description)
    }

    // Get a specific agent
    developer, err := b.GetAgent("developer")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Developer expertise: %v\n", developer.Expertise)

    // List available workflows
    workflows := b.ListWorkflows()
    for _, workflow := range workflows {
        fmt.Printf("Workflow: %s [%s]\n", workflow.Name, workflow.Phase)
    }
}
```

## 📚 Core Concepts

### Agents

Agents are specialized AI entities with specific roles and expertise:

```go
type Agent struct {
    Name        string   // Unique identifier
    Role        string   // Agent's role
    Description string   // What the agent does
    Expertise   []string // Areas of expertise
}
```

**Built-in Agents:**

- **Developer**: Code implementation and development
- **Architect**: System architecture and design
- **Product Manager**: Requirements and planning
- **Test Architect**: Testing strategy and QA
- **UX Designer**: User experience and interface design
- **Scrum Master**: Agile process and facilitation

### Workflows

Workflows represent development processes with multiple steps:

```go
type Workflow struct {
    Name        string
    Description string
    Phase       Phase
    Steps       []WorkflowStep
}
```

**Development Phases:**

- `PhaseAnalysis`: Analysis and brainstorming
- `PhasePlanning`: Planning and documentation
- `PhaseSolutioning`: Architecture and design
- `PhaseImplementation`: Development and testing

### Custom Workflows

Create and register custom workflows:

```go
customWorkflow := &bmad.Workflow{
    Name:        "feature-development",
    Description: "Full feature development cycle",
    Phase:       bmad.PhaseImplementation,
    Steps: []bmad.WorkflowStep{
        {
            Name:        "design",
            Description: "Design the feature",
            Action: func(ctx context.Context) error {
                // Implementation
                return nil
            },
        },
        {
            Name:        "implement",
            Description: "Implement the feature",
            Action: func(ctx context.Context) error {
                // Implementation
                return nil
            },
        },
        {
            Name:        "test",
            Description: "Test the feature",
            Action: func(ctx context.Context) error {
                // Implementation
                return nil
            },
        },
    },
}

b.RegisterWorkflow(customWorkflow)

// Execute the workflow
ctx := context.Background()
err := b.ExecuteWorkflow(ctx, "feature-development")
```

## 📖 Examples

See the [examples](./examples) directory for complete working examples:

- [Basic Usage](./examples/basic/main.go) - Introduction to BMAD-METHOD

## 🛠️ API Reference

### BMAD Instance

```go
// Create a new BMAD instance
b := bmad.New()
```

### Agent Management

```go
// Register a custom agent
agent := &bmad.Agent{
    Name:        "custom-agent",
    Role:        "Custom Role",
    Description: "Does custom things",
    Expertise:   []string{"custom", "specialized"},
}
b.RegisterAgent(agent)

// Get an agent
agent, err := b.GetAgent("developer")

// List all agents
agents := b.ListAgents()
```

### Workflow Management

```go
// Register a workflow
workflow := &bmad.Workflow{
    Name:        "my-workflow",
    Description: "My custom workflow",
    Phase:       bmad.PhaseImplementation,
    Steps:       []bmad.WorkflowStep{},
}
b.RegisterWorkflow(workflow)

// Get a workflow
workflow, err := b.GetWorkflow("workflow-init")

// List all workflows
workflows := b.ListWorkflows()

// Execute a workflow
ctx := context.Background()
err := b.ExecuteWorkflow(ctx, "my-workflow")
```

## 🧪 Testing

Run the test suite:

```bash
go test -v
```

Run tests with coverage:

```bash
go test -v -cover
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Related Projects

- [BMAD-METHOD](https://github.com/bmad-code-org/BMAD-METHOD) - The original BMAD-METHOD framework (JavaScript/TypeScript)

## 🌟 About BMAD

BMAD (Build More, Architect Dreams) is a breakthrough method for AI-driven agile development. It provides:

- **19 specialized AI agents** in the full framework
- **50+ guided workflows** across 4 development phases
- **Scale-adaptive intelligence** from bug fixes to enterprise systems
- **Complete development lifecycle** support

This Go implementation brings core BMAD concepts to the Go ecosystem, enabling Go developers to leverage AI-driven development workflows.

---

<p align="center">
  <sub>Built with ❤️ for the Go and AI collaboration community</sub>
</p>
