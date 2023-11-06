package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/shopspring/decimal"
	"github.com/sprkweb/finaplan-cli/finaplan/pkg/finaplan"
	"gopkg.in/yaml.v3"
)

func ParsePlan(input io.Reader) (*finaplan.FinancialPlan, error) {
	scanner := bufio.NewScanner(input)
	config, err := parseConfig(scanner)
	if err != nil {
		return nil, err
	}

	projection, err := parseProjection(scanner)
	if err != nil {
		return nil, err
	}

	return &finaplan.FinancialPlan{
		Config:     config,
		Projection: *projection,
	}, nil
}

func ParsePlanFromStdin() (*finaplan.FinancialPlan, error) {
	return ParsePlan(os.Stdin)
}

func parseConfig(input *bufio.Scanner) (*finaplan.PlanConfig, error) {
	// scan first line: it must be equal to configDelimiter
	if !input.Scan() {
		if err := input.Err(); err != nil {
			return nil, fmt.Errorf("error scanning input: %w", input.Err())
		}
		return nil, fmt.Errorf("no input")
	}
	if input.Text() != configDelimiter {
		return nil, fmt.Errorf("expected config delimiter at line 0")
	}

	// copy everything to buffer until we meet an ending configDelimiter
	var configBuf bytes.Buffer
	for input.Scan() && input.Text() != configDelimiter {
		configBuf.Write(input.Bytes())
		configBuf.WriteRune('\n')
	}
	if err := input.Err(); err != nil {
		return nil, fmt.Errorf("error scanning input: %w", input.Err())
	}

	// unmarshal config from YAML
	var config finaplan.PlanConfig
	if err := yaml.Unmarshal(configBuf.Bytes(), &config); err != nil {
		return nil, err
	}
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return &config, nil
}

func parseProjection(input *bufio.Scanner) (*finaplan.Projection, error) {
	var num decimal.Decimal
	var err error
	var projection finaplan.Projection

	for input.Scan() {
		num, err = decimal.NewFromString(input.Text())
		if err != nil {
			return nil, fmt.Errorf("error parsing number: %w", err)
		}
		projection = append(projection, num)
	}
	if err := input.Err(); err != nil {
		return nil, fmt.Errorf("error scanning input: %w", err)
	}

	return &projection, nil
}
