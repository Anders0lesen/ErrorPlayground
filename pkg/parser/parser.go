package parser

import "fmt"

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

