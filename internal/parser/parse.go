package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/sprkweb/finaplan-cli/finaplan/pkg/finaplan"
	"gopkg.in/yaml.v3"
)

func ParsePlan(input string) (*finaplan.FinancialPlan, error) {
	parts := strings.Split(input, ConfigDelimiter)
	if len(parts) != 3 {
		return nil, fmt.Errorf("expected 2 delimiters, got %d", len(parts)-1)
	}

	config, err := parseConfig([]byte(parts[1]))
	if err != nil {
		return nil, err
	}

	projection, err := parseProjection(parts[2])
	if err != nil {
		return nil, err
	}

	return &finaplan.FinancialPlan{
		Config:     config,
		Projection: *projection,
	}, nil
}

func ParsePlanFromStdin() (*finaplan.FinancialPlan, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return ParsePlan(string(input))
}

func parseConfig(input []byte) (*finaplan.PlanConfig, error) {
	var config finaplan.PlanConfig
	if err := yaml.Unmarshal(input, &config); err != nil {
		return nil, err
	}
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return &config, nil
}

func parseProjection(input string) (*finaplan.Projection, error) {
	reader := strings.NewReader(input)
	var line string
	var num decimal.Decimal
	var projection finaplan.Projection
	i := 0
	for {
		i++
		_, err := fmt.Fscanln(reader, &line)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("got error while parsing plan: %w", err)
		}
		num, err = decimal.NewFromString(line)
		if err != nil {
			return nil, fmt.Errorf("error parsing number: %w", err)
		}
		projection = append(projection, num)
	}
	return &projection, nil
}
