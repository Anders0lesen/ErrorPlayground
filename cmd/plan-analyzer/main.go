package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "log"

    "github.com/example/plan-analyzer/pkg/parser"
)

func main() {
    planPath := flag.String("plan", "", "Path to terraform plan JSON output")
    flag.Parse()

    if *planPath == "" {
        log.Fatal("plan path required")
    }

    data, err := ioutil.ReadFile(*planPath)
    if err != nil {
        log.Fatalf("failed to read plan: %v", err)
    }

    var plan parser.Plan
    if err := json.Unmarshal(data, &plan); err != nil {
        log.Fatalf("failed to parse plan: %v", err)
    }

    summary := parser.Summarize(plan)
    fmt.Println(summary)
}

