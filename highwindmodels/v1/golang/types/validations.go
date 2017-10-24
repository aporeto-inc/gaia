package types

import (
	"fmt"
	"time"
)

// validateStringValue validates a string parameter.
func (p *ServiceParameter) validateStringValue() error {

	if !p.Optional && (p.Value == nil || p.Value.(string) == "") {
		return fmt.Errorf("%s is required", p.Name)
	}

	if p.Value == nil {
		return nil
	}

	if _, ok := p.Value.(string); !ok {
		return fmt.Errorf("%s is not a valid string", p.Name)
	}
	return nil
}

// validateIntValue validates a int parameter.
func (p *ServiceParameter) validateIntValue() error {

	if !p.Optional && (p.Value == nil) {
		return fmt.Errorf("%s is required", p.Name)
	}

	if p.Value == nil {
		return nil
	}

	if _, ok := p.Value.(int); !ok {
		return fmt.Errorf("%s is not a valid integer", p.Name)
	}
	return nil
}

// validateFloatValue validates a float parameter.
func (p *ServiceParameter) validateFloatValue() error {

	if !p.Optional && (p.Value == nil) {
		return fmt.Errorf("%s is required", p.Name)
	}

	if p.Value == nil {
		return nil
	}

	if _, ok := p.Value.(float64); !ok {
		return fmt.Errorf("%s is not a valid float", p.Name)
	}
	return nil
}

// validateBoolValue validates a bool parameter.
func (p *ServiceParameter) validateBoolValue() error {

	if !p.Optional && (p.Value == nil) {
		return fmt.Errorf("%s is required", p.Name)
	}

	if p.Value == nil {
		return nil
	}

	if _, ok := p.Value.(bool); !ok {
		return fmt.Errorf("%s is not a valid boolean", p.Name)
	}
	return nil
}

// validateDurationValue validates a duration parameter.
func (p *ServiceParameter) validateDurationValue() error {

	if !p.Optional && (p.Value == nil) {
		return fmt.Errorf("%s is required", p.Name)
	}

	if p.Value == nil {
		return nil
	}

	duration, ok := p.Value.(string)
	if !ok {
		return fmt.Errorf("%s is not a valid string", p.Name)
	}

	if _, err := time.ParseDuration(duration); err != nil {
		return fmt.Errorf("%s is not a valid duration", p.Name)
	}
	return nil
}

// isStringAllowedValue returns true if the value is allowed
func isAllowedValue(allowedValues []interface{}, value interface{}) error {

	if len(allowedValues) == 0 {
		return nil
	}

	switch allowedValues[0].(type) {
	case string:
		_, ok := value.(string)
		if !ok {
			return fmt.Errorf("%s is not a string", value)
		}
		return nil

	case int:
		_, ok := value.(int)
		if !ok {
			return fmt.Errorf("%s is not an int", value)
		}
		return nil

	case float32:
		_, ok := value.(float32)
		if !ok {
			return fmt.Errorf("%s is not a float", value)
		}
		return nil
	}

	for _, allowed := range allowedValues {
		if value == allowed {
			return fmt.Errorf("%s is not allowed", value)
		}
	}
	return nil
}
