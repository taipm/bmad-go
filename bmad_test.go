package bmad

import (
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	b := New()
	if b == nil {
		t.Fatal("New() returned nil")
	}

	if len(b.agents) == 0 {
		t.Error("Expected default agents to be registered")
	}

	if len(b.workflows) == 0 {
		t.Error("Expected default workflows to be registered")
	}
}

func TestRegisterAgent(t *testing.T) {
	b := New()
	agent := &Agent{
		Name:        "test-agent",
		Role:        "Test Agent",
		Description: "A test agent",
		Expertise:   []string{"testing"},
	}

	b.RegisterAgent(agent)

	retrieved, err := b.GetAgent("test-agent")
	if err != nil {
		t.Fatalf("Failed to get registered agent: %v", err)
	}

	if retrieved.Name != agent.Name {
		t.Errorf("Expected agent name %s, got %s", agent.Name, retrieved.Name)
	}
}

func TestGetAgent(t *testing.T) {
	b := New()

	// Test getting existing agent
	agent, err := b.GetAgent("developer")
	if err != nil {
		t.Fatalf("Failed to get developer agent: %v", err)
	}

	if agent.Role != "Developer" {
		t.Errorf("Expected role 'Developer', got '%s'", agent.Role)
	}

	// Test getting non-existent agent
	_, err = b.GetAgent("non-existent")
	if err == nil {
		t.Error("Expected error for non-existent agent")
	}
}

func TestListAgents(t *testing.T) {
	b := New()
	agents := b.ListAgents()

	if len(agents) == 0 {
		t.Error("Expected at least one agent")
	}

	// Check if developer agent is in the list
	found := false
	for _, agent := range agents {
		if agent.Name == "developer" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected to find developer agent in list")
	}
}

func TestRegisterWorkflow(t *testing.T) {
	b := New()
	workflow := &Workflow{
		Name:        "test-workflow",
		Description: "A test workflow",
		Phase:       PhaseAnalysis,
		Steps:       []WorkflowStep{},
	}

	b.RegisterWorkflow(workflow)

	retrieved, err := b.GetWorkflow("test-workflow")
	if err != nil {
		t.Fatalf("Failed to get registered workflow: %v", err)
	}

	if retrieved.Name != workflow.Name {
		t.Errorf("Expected workflow name %s, got %s", workflow.Name, retrieved.Name)
	}
}

func TestGetWorkflow(t *testing.T) {
	b := New()

	// Test getting existing workflow
	workflow, err := b.GetWorkflow("workflow-init")
	if err != nil {
		t.Fatalf("Failed to get workflow-init: %v", err)
	}

	if workflow.Phase != PhaseAnalysis {
		t.Errorf("Expected phase 'analysis', got '%s'", workflow.Phase)
	}

	// Test getting non-existent workflow
	_, err = b.GetWorkflow("non-existent")
	if err == nil {
		t.Error("Expected error for non-existent workflow")
	}
}

func TestListWorkflows(t *testing.T) {
	b := New()
	workflows := b.ListWorkflows()

	if len(workflows) == 0 {
		t.Error("Expected at least one workflow")
	}

	// Check if workflow-init is in the list
	found := false
	for _, workflow := range workflows {
		if workflow.Name == "workflow-init" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected to find workflow-init in list")
	}
}

func TestExecuteWorkflow(t *testing.T) {
	b := New()
	executed := false

	workflow := &Workflow{
		Name:        "test-exec-workflow",
		Description: "Test execution workflow",
		Phase:       PhaseImplementation,
		Steps: []WorkflowStep{
			{
				Name:        "test-step",
				Description: "A test step",
				Action: func(ctx context.Context) error {
					executed = true
					return nil
				},
			},
		},
	}

	b.RegisterWorkflow(workflow)

	ctx := context.Background()
	err := b.ExecuteWorkflow(ctx, "test-exec-workflow")
	if err != nil {
		t.Fatalf("Failed to execute workflow: %v", err)
	}

	if !executed {
		t.Error("Workflow step was not executed")
	}
}

func TestPhaseConstants(t *testing.T) {
	phases := []Phase{
		PhaseAnalysis,
		PhasePlanning,
		PhaseSolutioning,
		PhaseImplementation,
	}

	expected := []string{
		"analysis",
		"planning",
		"solutioning",
		"implementation",
	}

	for i, phase := range phases {
		if string(phase) != expected[i] {
			t.Errorf("Expected phase %s, got %s", expected[i], phase)
		}
	}
}

func TestVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version should not be empty")
	}
}
