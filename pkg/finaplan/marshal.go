package finaplan

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

const formatVersion = "0.1.0"

type serializablePlan struct {
	Version    string             `json:"version"`
	Config     serializableConfig `json:"config"`
	Projection []string           `json:"projection"`
}

type serializableConfig struct {
	IntervalType   string `json:"interval_type"`
	IntervalLength uint32 `json:"interval_length"`
}

var _ json.Marshaler = &FinancialPlan{}
var _ json.Unmarshaler = &FinancialPlan{}

func (p *FinancialPlan) MarshalJSON() ([]byte, error) {
	if p == nil {
		return nil, errors.New("MarshalJSON on nil pointer")
	}

	dto := p.toDto()
	encoded, err := json.Marshal(dto)
	if err != nil {
		return nil, fmt.Errorf("marshaling plan: %w", err)
	}
	return encoded, nil
}

func (p *FinancialPlan) toDto() serializablePlan {
	return serializablePlan{
		Version: formatVersion,
		Config: serializableConfig{
			IntervalType:   p.Config.IntervalType.String(),
			IntervalLength: p.Config.IntervalLength,
		},
		Projection: p.Print(),
	}
}

func (p *FinancialPlan) UnmarshalJSON(data []byte) error {
	if p == nil {
		return errors.New("UnmarshalJSON on nil pointer")
	}

	dto := &serializablePlan{}
	if err := json.Unmarshal(data, dto); err != nil {
		return fmt.Errorf("unmarshaling plan: %w", err)
	}

	if dto.Version != formatVersion {
		return fmt.Errorf("unsupported format version: %q", dto.Version)
	}

	config, err := dto.toConfig()
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}
	projection, err := dto.toProjection()
	if err != nil {
		return fmt.Errorf("parsing projection: %w", err)
	}

	*p = FinancialPlan{
		Config:     config,
		Projection: projection,
	}
	return nil
}

func (dto *serializablePlan) toConfig() (*PlanConfig, error) {
	config := &PlanConfig{
		IntervalType:   IntervalType(dto.Config.IntervalType),
		IntervalLength: dto.Config.IntervalLength,
	}
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return config, nil
}

func (dto *serializablePlan) toProjection() (Projection, error) {
	var num decimal.Decimal
	var err error
	projection := make(Projection, 0, len(dto.Projection))

	for _, v := range dto.Projection {
		num, err = decimal.NewFromString(v)
		if err != nil {
			return nil, fmt.Errorf("parsing number: %w", err)
		}
		projection = append(projection, num)
	}

	return projection, nil
}
