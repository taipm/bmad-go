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

	fmt.Printf("BMAD-METHOD for Go v%s\n\n", bmad.Version)

	// List all available agents
	fmt.Println("Available Agents:")
	agents := b.ListAgents()
	for _, agent := range agents {
		fmt.Printf("  - %s (%s): %s\n", agent.Role, agent.Name, agent.Description)
		fmt.Printf("    Expertise: %v\n", agent.Expertise)
	}

	fmt.Println()

	// List all available workflows
	fmt.Println("Available Workflows:")
	workflows := b.ListWorkflows()
	for _, workflow := range workflows {
		fmt.Printf("  - %s [%s]: %s\n", workflow.Name, workflow.Phase, workflow.Description)
	}

	fmt.Println()

	// Get a specific agent
	developer, err := b.GetAgent("developer")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Retrieved Agent: %s - %s\n", developer.Role, developer.Description)

	fmt.Println()

	// Register a custom workflow
	customWorkflow := &bmad.Workflow{
		Name:        "custom-feature",
		Description: "Custom feature development workflow",
		Phase:       bmad.PhaseImplementation,
		Steps: []bmad.WorkflowStep{
			{
				Name:        "analyze",
				Description: "Analyze requirements",
				Action: func(ctx context.Context) error {
					fmt.Println("  → Analyzing requirements...")
					return nil
				},
			},
			{
				Name:        "implement",
				Description: "Implement feature",
				Action: func(ctx context.Context) error {
					fmt.Println("  → Implementing feature...")
					return nil
				},
			},
			{
				Name:        "test",
				Description: "Test implementation",
				Action: func(ctx context.Context) error {
					fmt.Println("  → Testing implementation...")
					return nil
				},
			},
		},
	}

	b.RegisterWorkflow(customWorkflow)

	// Execute the custom workflow
	fmt.Println("Executing custom workflow:")
	ctx := context.Background()
	if err := b.ExecuteWorkflow(ctx, "custom-feature"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n✓ Custom workflow completed successfully!")
}
