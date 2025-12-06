/*
Package bmad provides the BMAD-METHOD framework for Go (Golang).

BMAD (Build More, Architect Dreams) is an AI-driven agile development framework
with specialized agents and workflows. This package brings the power of BMAD to
the Go ecosystem.

# Overview

BMAD-METHOD helps developers build better software through:

  - Specialized AI agents with specific roles and expertise
  - Structured workflows for different development scenarios
  - A complete development lifecycle from analysis to implementation

# Quick Start

Create a new BMAD instance and explore available agents:

	b := bmad.New()

	// List all agents
	agents := b.ListAgents()
	for _, agent := range agents {
		fmt.Printf("%s: %s\n", agent.Role, agent.Description)
	}

# Agents

The framework includes 6 specialized agents:

  - Developer: Code implementation and development
  - Architect: System architecture and design
  - Product Manager: Requirements and planning
  - Test Architect: Testing strategy and QA
  - UX Designer: User experience and interface design
  - Scrum Master: Agile process and facilitation

You can also register custom agents:

	customAgent := &bmad.Agent{
		Name:        "devops",
		Role:        "DevOps Engineer",
		Description: "Specialized in deployment and infrastructure",
		Expertise:   []string{"ci-cd", "docker", "kubernetes"},
	}
	b.RegisterAgent(customAgent)

# Workflows

Workflows represent development processes with multiple steps. The framework
includes several built-in workflows:

  - workflow-init: Initialize and analyze project
  - workflow-quick-flow: Quick flow for bug fixes
  - workflow-method: Full BMAD method for products

Create custom workflows with executable steps:

	workflow := &bmad.Workflow{
		Name:        "feature-development",
		Description: "Complete feature development cycle",
		Phase:       bmad.PhaseImplementation,
		Steps: []bmad.WorkflowStep{
			{
				Name:        "design",
				Description: "Design the feature",
				Action: func(ctx context.Context) error {
					// Design logic here
					return nil
				},
			},
			{
				Name:        "implement",
				Description: "Implement the feature",
				Action: func(ctx context.Context) error {
					// Implementation logic here
					return nil
				},
			},
		},
	}

	b.RegisterWorkflow(workflow)

	// Execute the workflow
	ctx := context.Background()
	err := b.ExecuteWorkflow(ctx, "feature-development")

# Development Phases

The framework organizes work into four phases:

  - PhaseAnalysis: Analysis and brainstorming
  - PhasePlanning: Planning and documentation
  - PhaseSolutioning: Architecture and design
  - PhaseImplementation: Development and testing

Each workflow is associated with a phase to help organize your development process.

# Context Support

All workflow actions accept a context.Context parameter, enabling:

  - Cancellation of long-running workflows
  - Timeouts for workflow steps
  - Passing request-scoped values

Example with timeout:

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := b.ExecuteWorkflow(ctx, "feature-development")
	if err != nil {
		// Handle timeout or other errors
	}

# Error Handling

Functions return errors following Go conventions. Check errors appropriately:

	agent, err := b.GetAgent("developer")
	if err != nil {
		log.Fatalf("Failed to get agent: %v", err)
	}

	workflow, err := b.GetWorkflow("workflow-init")
	if err != nil {
		log.Fatalf("Failed to get workflow: %v", err)
	}

For more information, visit: https://github.com/taipm/bmad-go
*/
package bmad
