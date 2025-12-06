# BMAD-METHOD Examples

This directory contains examples demonstrating different aspects of the BMAD-METHOD framework for Go.

## Available Examples

### 1. Basic Usage (`basic/`)

Demonstrates the fundamental features of BMAD-METHOD:
- Creating a BMAD instance
- Listing built-in agents
- Listing built-in workflows
- Getting specific agents
- Creating and executing custom workflows

**Run:**
```bash
cd basic
go run main.go
```

### 2. Custom Agents & Workflows (`custom-agents/`)

Shows how to extend BMAD with custom agents and workflows:
- Registering custom agents (DevOps, Security, Data Engineer)
- Creating complex multi-step workflows
- Executing workflows with real-world scenarios
- Using context for timeouts and cancellation

**Run:**
```bash
cd custom-agents
go run main.go
```

## Common Patterns

### Creating a BMAD Instance

```go
b := bmad.New()
```

### Registering Custom Agents

```go
agent := &bmad.Agent{
    Name:        "custom-agent",
    Role:        "Custom Role",
    Description: "What this agent does",
    Expertise:   []string{"skill1", "skill2"},
}
b.RegisterAgent(agent)
```

### Creating Workflows

```go
workflow := &bmad.Workflow{
    Name:        "my-workflow",
    Description: "Description of what this does",
    Phase:       bmad.PhaseImplementation,
    Steps: []bmad.WorkflowStep{
        {
            Name:        "step-1",
            Description: "First step",
            Action: func(ctx context.Context) error {
                // Your logic here
                return nil
            },
        },
        // Add more steps...
    },
}
b.RegisterWorkflow(workflow)
```

### Executing Workflows

```go
ctx := context.Background()
err := b.ExecuteWorkflow(ctx, "my-workflow")
if err != nil {
    log.Fatal(err)
}
```

### Using Context with Timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

err := b.ExecuteWorkflow(ctx, "my-workflow")
```

## Development Phases

BMAD organizes work into four phases:

- **PhaseAnalysis**: Analysis and brainstorming
- **PhasePlanning**: Planning and documentation  
- **PhaseSolutioning**: Architecture and design
- **PhaseImplementation**: Development and testing

## Built-in Agents

The framework includes 6 specialized agents:

1. **Developer**: Code implementation and development
2. **Architect**: System architecture and design
3. **Product Manager**: Requirements and planning
4. **Test Architect**: Testing strategy and QA
5. **UX Designer**: User experience and interface design
6. **Scrum Master**: Agile process and facilitation

## Next Steps

- Read the [main README](../README.md) for API reference
- Check the [package documentation](https://pkg.go.dev/github.com/taipm/bmad-go)
- Review the [contributing guidelines](../CONTRIBUTING.md)
