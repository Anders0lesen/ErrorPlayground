package parser

import (
	"encoding/json"
	"fmt"
)

type Plan struct {
	ResourceChanges []ResourceChange `json:"resource_changes"`
}

type ResourceChange struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Provider string `json:"provider_name"`
	Change   Change `json:"change"`
}

type Change struct {
	Actions []string `json:"actions"`
}

func ParsePlan(data []byte) (Plan, error) {
	var p Plan
	if err := json.Unmarshal(data, &p); err != nil {
		return Plan{}, fmt.Errorf("unmarshal plan: %w", err)
	}
	return p, nil
}

func Summarize(p Plan) string {
	summary := make(map[string]int)
	for _, rc := range p.ResourceChanges {
		key := fmt.Sprintf("%s.%s", rc.Type, rc.Name)
		summary[key]++
	}
	result := "Summary of changes:\n"
	for k, v := range summary {
		result += fmt.Sprintf("%s: %d\n", k, v)
	}
	return result
}
