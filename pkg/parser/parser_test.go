package parser

import "testing"

const samplePlan = `{
  "resource_changes": [
    {
      "type": "aws_instance",
      "name": "example",
      "provider_name": "registry.terraform.io/hashicorp/aws",
      "change": {"actions": ["create"]}
    }
  ]
}`

func TestParsePlan(t *testing.T) {
	p, err := ParsePlan([]byte(samplePlan))
	if err != nil {
		t.Fatalf("ParsePlan failed: %v", err)
	}
	if len(p.ResourceChanges) != 1 {
		t.Fatalf("expected 1 change, got %d", len(p.ResourceChanges))
	}
}

func TestSummarize(t *testing.T) {
	p, _ := ParsePlan([]byte(samplePlan))
	summary := Summarize(p)
	if summary == "" {
		t.Fatal("summary should not be empty")
	}
}
