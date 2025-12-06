// Package bmad provides the BMAD-METHOD framework for Go (Golang).
// BMAD (Build More, Architect Dreams) is an AI-driven agile development framework
// with specialized agents and workflows.
package bmad

import (
	"context"
	"fmt"
)

// Version is the current version of the BMAD-METHOD library
const Version = "1.0.0"

// Agent represents a specialized AI agent in the BMAD framework
type Agent struct {
	Name        string
	Role        string
	Description string
	Expertise   []string
}

// Workflow represents a development workflow in the BMAD framework
type Workflow struct {
	Name        string
	Description string
	Phase       Phase
	Steps       []WorkflowStep
}

// WorkflowStep represents a single step in a workflow
type WorkflowStep struct {
	Name        string
	Description string
	Agent       *Agent
	Action      func(ctx context.Context) error
}

// Phase represents the development phase
type Phase string

const (
	// PhaseAnalysis represents the analysis phase
	PhaseAnalysis Phase = "analysis"
	// PhasePlanning represents the planning phase
	PhasePlanning Phase = "planning"
	// PhaseSolutioning represents the solutioning phase
	PhaseSolutioning Phase = "solutioning"
	// PhaseImplementation represents the implementation phase
	PhaseImplementation Phase = "implementation"
)

// BMAD represents the main BMAD framework instance
type BMAD struct {
	agents    map[string]*Agent
	workflows map[string]*Workflow
}

// New creates a new BMAD framework instance
func New() *BMAD {
	b := &BMAD{
		agents:    make(map[string]*Agent),
		workflows: make(map[string]*Workflow),
	}
	b.initializeDefaultAgents()
	b.initializeDefaultWorkflows()
	return b
}

// initializeDefaultAgents sets up the default specialized agents
func (b *BMAD) initializeDefaultAgents() {
	b.RegisterAgent(&Agent{
		Name:        "developer",
		Role:        "Developer",
		Description: "Specialized in code implementation and development",
		Expertise:   []string{"coding", "refactoring", "debugging"},
	})

	b.RegisterAgent(&Agent{
		Name:        "architect",
		Role:        "Architect",
		Description: "Specialized in system architecture and design",
		Expertise:   []string{"architecture", "design-patterns", "scalability"},
	})

	b.RegisterAgent(&Agent{
		Name:        "pm",
		Role:        "Product Manager",
		Description: "Specialized in product requirements and planning",
		Expertise:   []string{"requirements", "planning", "prioritization"},
	})

	b.RegisterAgent(&Agent{
		Name:        "test-architect",
		Role:        "Test Architect",
		Description: "Specialized in testing strategy and quality assurance",
		Expertise:   []string{"testing", "quality-assurance", "test-automation"},
	})

	b.RegisterAgent(&Agent{
		Name:        "ux-designer",
		Role:        "UX Designer",
		Description: "Specialized in user experience and interface design",
		Expertise:   []string{"ux", "ui", "user-research"},
	})

	b.RegisterAgent(&Agent{
		Name:        "scrum-master",
		Role:        "Scrum Master",
		Description: "Specialized in agile process and team facilitation",
		Expertise:   []string{"agile", "scrum", "facilitation"},
	})
}

// initializeDefaultWorkflows sets up the default workflows
func (b *BMAD) initializeDefaultWorkflows() {
	b.RegisterWorkflow(&Workflow{
		Name:        "workflow-init",
		Description: "Initialize and analyze project for the right workflow track",
		Phase:       PhaseAnalysis,
		Steps:       []WorkflowStep{},
	})

	b.RegisterWorkflow(&Workflow{
		Name:        "workflow-quick-flow",
		Description: "Quick flow for bug fixes and small features",
		Phase:       PhaseImplementation,
		Steps:       []WorkflowStep{},
	})

	b.RegisterWorkflow(&Workflow{
		Name:        "workflow-method",
		Description: "Full BMAD method for products and platforms",
		Phase:       PhasePlanning,
		Steps:       []WorkflowStep{},
	})
}

// RegisterAgent registers a new agent with the framework
func (b *BMAD) RegisterAgent(agent *Agent) {
	b.agents[agent.Name] = agent
}

// GetAgent retrieves an agent by name
func (b *BMAD) GetAgent(name string) (*Agent, error) {
	agent, ok := b.agents[name]
	if !ok {
		return nil, fmt.Errorf("agent not found: %s", name)
	}
	return agent, nil
}

// ListAgents returns all registered agents
func (b *BMAD) ListAgents() []*Agent {
	agents := make([]*Agent, 0, len(b.agents))
	for _, agent := range b.agents {
		agents = append(agents, agent)
	}
	return agents
}

// RegisterWorkflow registers a new workflow with the framework
func (b *BMAD) RegisterWorkflow(workflow *Workflow) {
	b.workflows[workflow.Name] = workflow
}

// GetWorkflow retrieves a workflow by name
func (b *BMAD) GetWorkflow(name string) (*Workflow, error) {
	workflow, ok := b.workflows[name]
	if !ok {
		return nil, fmt.Errorf("workflow not found: %s", name)
	}
	return workflow, nil
}

// ListWorkflows returns all registered workflows
func (b *BMAD) ListWorkflows() []*Workflow {
	workflows := make([]*Workflow, 0, len(b.workflows))
	for _, workflow := range b.workflows {
		workflows = append(workflows, workflow)
	}
	return workflows
}

// ExecuteWorkflow executes a workflow by name
func (b *BMAD) ExecuteWorkflow(ctx context.Context, name string) error {
	workflow, err := b.GetWorkflow(name)
	if err != nil {
		return err
	}

	for _, step := range workflow.Steps {
		if step.Action != nil {
			if err := step.Action(ctx); err != nil {
				return fmt.Errorf("workflow step '%s' failed: %w", step.Name, err)
			}
		}
	}

	return nil
}
