package parser

import (
	"fmt"
	"io"
	"os"

	"github.com/sprkweb/finaplan/pkg/finaplan"
	"gopkg.in/yaml.v3"
)

const (
	configDelimiter = "---"
	newline         = byte('\n')
)

func PrintPlan(w io.Writer, p *finaplan.FinancialPlan) error {
	err := printConfig(w, p.Config)
	if err != nil {
		return err
	}

	for _, v := range p.Projection {
		if _, err = io.WriteString(w, v.String()); err != nil {
			return err
		}

		if err := writeNewline(w); err != nil {
			return err
		}
	}
	return nil
}

func PrintPlanToStdout(plan *finaplan.FinancialPlan) error {
	return PrintPlan(os.Stdout, plan)
}

func printConfig(w io.Writer, config *finaplan.PlanConfig) error {
	yamlConfig, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("unable to format config as YAML: %w", err)
	}

	if _, err = io.WriteString(w, configDelimiter); err != nil {
		return err
	}
	if err := writeNewline(w); err != nil {
		return err
	}
	if _, err = w.Write(yamlConfig); err != nil {
		return err
	}
	if _, err = io.WriteString(w, configDelimiter); err != nil {
		return err
	}
	if err := writeNewline(w); err != nil {
		return err
	}
	return nil
}

func writeNewline(w io.Writer) error {
	_, err := w.Write([]byte{newline})
	if err != nil {
		return err
	}
	return nil
}
