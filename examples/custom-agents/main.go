package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/taipm/bmad-go"
)

func main() {
	// Create a new BMAD instance
	b := bmad.New()

	fmt.Println("BMAD-METHOD: Custom Agents & Workflows Example")
	fmt.Println("===============================================")
	fmt.Println()

	// 1. Register custom agents
	fmt.Println("1. Registering custom agents...")

	devOpsAgent := &bmad.Agent{
		Name:        "devops",
		Role:        "DevOps Engineer",
		Description: "Specialized in CI/CD, deployment, and infrastructure",
		Expertise:   []string{"ci-cd", "docker", "kubernetes", "terraform"},
	}
	b.RegisterAgent(devOpsAgent)

	securityAgent := &bmad.Agent{
		Name:        "security",
		Role:        "Security Engineer",
		Description: "Specialized in security audits and vulnerability assessment",
		Expertise:   []string{"security-audit", "penetration-testing", "compliance"},
	}
	b.RegisterAgent(securityAgent)

	dataEngineer := &bmad.Agent{
		Name:        "data-engineer",
		Role:        "Data Engineer",
		Description: "Specialized in data pipelines and analytics",
		Expertise:   []string{"etl", "data-modeling", "big-data"},
	}
	b.RegisterAgent(dataEngineer)

	fmt.Println("  ✓ Registered DevOps Engineer")
	fmt.Println("  ✓ Registered Security Engineer")
	fmt.Println("  ✓ Registered Data Engineer")
	fmt.Println()

	// 2. Display all agents (built-in + custom)
	fmt.Println("2. All available agents:")
	agents := b.ListAgents()
	for _, agent := range agents {
		fmt.Printf("  • %s: %s\n", agent.Role, agent.Description)
	}
	fmt.Println()

	// 3. Create a custom deployment workflow
	fmt.Println("3. Creating custom deployment workflow...")

	deploymentWorkflow := &bmad.Workflow{
		Name:        "production-deployment",
		Description: "Complete production deployment with security checks",
		Phase:       bmad.PhaseImplementation,
		Steps: []bmad.WorkflowStep{
			{
				Name:        "security-scan",
				Description: "Run security vulnerability scan",
				Action: func(ctx context.Context) error {
					fmt.Println("  → Running security scan...")
					time.Sleep(100 * time.Millisecond)
					fmt.Println("     ✓ No vulnerabilities found")
					return nil
				},
			},
			{
				Name:        "build",
				Description: "Build application artifacts",
				Action: func(ctx context.Context) error {
					fmt.Println("  → Building application...")
					time.Sleep(150 * time.Millisecond)
					fmt.Println("     ✓ Build successful")
					return nil
				},
			},
			{
				Name:        "test",
				Description: "Run integration tests",
				Action: func(ctx context.Context) error {
					fmt.Println("  → Running integration tests...")
					time.Sleep(200 * time.Millisecond)
					fmt.Println("     ✓ All tests passed")
					return nil
				},
			},
			{
				Name:        "deploy",
				Description: "Deploy to production",
				Action: func(ctx context.Context) error {
					fmt.Println("  → Deploying to production...")
					time.Sleep(250 * time.Millisecond)
					fmt.Println("     ✓ Deployment successful")
					return nil
				},
			},
			{
				Name:        "health-check",
				Description: "Verify deployment health",
				Action: func(ctx context.Context) error {
					fmt.Println("  → Running health checks...")
					time.Sleep(100 * time.Millisecond)
					fmt.Println("     ✓ All services healthy")
					return nil
				},
			},
		},
	}

	b.RegisterWorkflow(deploymentWorkflow)
	fmt.Println("  ✓ Deployment workflow registered")
	fmt.Println()

	// 4. Create a data pipeline workflow
	fmt.Println("4. Creating data pipeline workflow...")

	dataPipelineWorkflow := &bmad.Workflow{
		Name:        "data-pipeline-setup",
		Description: "Set up data processing pipeline",
		Phase:       bmad.PhaseSolutioning,
		Steps: []bmad.WorkflowStep{
			{
				Name:        "design-schema",
				Description: "Design data schema",
				Action: func(ctx context.Context) error {
					fmt.Println("  → Designing data schema...")
					time.Sleep(100 * time.Millisecond)
					fmt.Println("     ✓ Schema designed")
					return nil
				},
			},
			{
				Name:        "setup-etl",
				Description: "Configure ETL pipeline",
				Action: func(ctx context.Context) error {
					fmt.Println("  → Setting up ETL pipeline...")
					time.Sleep(150 * time.Millisecond)
					fmt.Println("     ✓ ETL pipeline configured")
					return nil
				},
			},
			{
				Name:        "test-pipeline",
				Description: "Test data pipeline",
				Action: func(ctx context.Context) error {
					fmt.Println("  → Testing data pipeline...")
					time.Sleep(120 * time.Millisecond)
					fmt.Println("     ✓ Pipeline working correctly")
					return nil
				},
			},
		},
	}

	b.RegisterWorkflow(dataPipelineWorkflow)
	fmt.Println("  ✓ Data pipeline workflow registered")
	fmt.Println()

	// 5. Execute workflows
	fmt.Println("5. Executing deployment workflow:")
	ctx := context.Background()
	if err := b.ExecuteWorkflow(ctx, "production-deployment"); err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	fmt.Println("6. Executing data pipeline workflow:")
	if err := b.ExecuteWorkflow(ctx, "data-pipeline-setup"); err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	// 6. Demonstrate context with timeout
	fmt.Println("7. Demonstrating workflow with timeout:")
	timeoutWorkflow := &bmad.Workflow{
		Name:        "long-running-task",
		Description: "A task that might take too long",
		Phase:       bmad.PhaseImplementation,
		Steps: []bmad.WorkflowStep{
			{
				Name:        "quick-step",
				Description: "This will complete",
				Action: func(ctx context.Context) error {
					fmt.Println("  → Executing quick step...")
					time.Sleep(100 * time.Millisecond)
					fmt.Println("     ✓ Quick step completed")
					return nil
				},
			},
		},
	}

	b.RegisterWorkflow(timeoutWorkflow)

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := b.ExecuteWorkflow(ctxWithTimeout, "long-running-task"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("  ✓ Workflow completed within timeout")
	fmt.Println()

	// Summary
	fmt.Println("===============================================")
	fmt.Printf("Summary: %d agents, %d workflows\n", len(b.ListAgents()), len(b.ListWorkflows()))
	fmt.Println("All workflows executed successfully! ✨")
}
