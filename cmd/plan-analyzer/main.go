package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/example/plan-analyzer/pkg/llm"
	"github.com/example/plan-analyzer/pkg/logging"
	"github.com/example/plan-analyzer/pkg/parser"
)

func main() {
	logging.Setup()
	planPath := flag.String("plan", "", "Path to terraform plan JSON output")
	flag.Parse()

	if *planPath == "" {
		log.Fatal("plan path required")
	}

	data, err := ioutil.ReadFile(*planPath)
	if err != nil {
		log.Fatalf("failed to read plan: %v", err)
	}

	plan, err := parser.ParsePlan(data)
	if err != nil {
		log.Fatalf("failed to parse plan: %v", err)
	}

	summary := parser.Summarize(plan)
	fmt.Println(summary)

	analysis, err := llm.Analyze(plan)
	if err != nil {
		log.Fatalf("LLM analysis failed: %v", err)
	}
	fmt.Println(analysis)
}
