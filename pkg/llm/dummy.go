package llm

import (
	"fmt"

	"github.com/example/plan-analyzer/pkg/parser"
)

func Analyze(plan parser.Plan) (string, error) {
	// Placeholder implementation for local LLM integration
	count := len(plan.ResourceChanges)
	return fmt.Sprintf("Plan contains %d resource changes", count), nil
}
