package llm

import (
	"testing"

	"github.com/example/plan-analyzer/pkg/parser"
)

func TestAnalyze(t *testing.T) {
	p := parser.Plan{ResourceChanges: []parser.ResourceChange{{Type: "x", Name: "y"}}}
	out, err := Analyze(p)
	if err != nil {
		t.Fatalf("Analyze failed: %v", err)
	}
	if out == "" {
		t.Fatal("expected analysis output")
	}
}
