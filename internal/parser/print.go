package parser

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"

	"github.com/sprkweb/finaplan-cli/pkg/finaplan"
)

func PrintPlan(p *finaplan.FinancialPlan) (string, error) {
	var builder strings.Builder
	err := printConfig(&builder, p.Config)
	if err != nil {
		return "", err
	}

	for _, v := range p.Projection {
		builder.WriteString(fmt.Sprintf("%v\n", v))
	}
	return builder.String(), nil
}

func PrintPlanToStdout(plan *finaplan.FinancialPlan) error {
	planStr, err := PrintPlan(plan)
	if err != nil {
		return err
	}
	fmt.Print(planStr)
	return nil
}

const ConfigDelimiter = "---\n"

func printConfig(builder *strings.Builder, config *finaplan.PlanConfig) error {
	yamlConfig, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("unable to format config as YAML: %w", err)
	}

	builder.WriteString(ConfigDelimiter)
	builder.Write(yamlConfig)
	builder.WriteString(ConfigDelimiter)
	return nil
}
