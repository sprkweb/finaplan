package parser

import (
	"fmt"
	"github.com/sprkweb/finaplan-cli/pkg/finaplan"
	"gopkg.in/yaml.v3"
	"io"
	"strings"
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
	var num finaplan.ProjectionUnit
	var projection finaplan.Projection
	i := 0
	for {
		i++
		n, err := fmt.Fscanln(reader, &num)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("got error while parsing plan: %w", err)
		}
		if n != 1 {
			return nil, fmt.Errorf("got 0 items on line %d", i)
		}
		projection = append(projection, num)
	}
	return &projection, nil
}
