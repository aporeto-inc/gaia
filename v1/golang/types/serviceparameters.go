package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ServiceParameterType is the type representing the type of a parameter
type ServiceParameterType string

// Various values for ServiceParameterType.
const (
	ServiceParameterTypeBool        ServiceParameterType = "bool"
	ServiceParameterTypeDuration    ServiceParameterType = "duration"
	ServiceParameterTypeEnum        ServiceParameterType = "enum"
	ServiceParameterTypeIntSlice    ServiceParameterType = "intSlice"
	ServiceParameterTypeInt         ServiceParameterType = "int"
	ServiceParameterTypeFloat       ServiceParameterType = "float"
	ServiceParameterTypeFloatSlice  ServiceParameterType = "floatSlice"
	ServiceParameterTypePassword    ServiceParameterType = "password"
	ServiceParameterTypeString      ServiceParameterType = "string"
	ServiceParameterTypeStringSlice ServiceParameterType = "stringSlice"
)

// ServiceParameterBackend defines the link of the service parameter.
type ServiceParameterBackend int

const (
	// ServiceParameterBackendGlobalSecret defines a link to installation secret.
	ServiceParameterBackendGlobalSecret ServiceParameterBackend = iota

	// ServiceParameterBackendGlobalConfigMap defines a link to installation config map.
	ServiceParameterBackendGlobalConfigMap

	// ServiceParameterBackendLocalSecret defines a link to service secret.
	ServiceParameterBackendLocalSecret

	// ServiceParameterBackendLocalConfigMap defines a link to service config map.
	ServiceParameterBackendLocalConfigMap
)

// ServiceParameter defines a parameter for the service.
type ServiceParameter struct {
	Name              string                  `json:"name" yaml:"name"`
	Description       string                  `json:"description" yaml:"description"`
	LongDescription   string                  `json:"longDescription" yaml:"longDescription"`
	Key               string                  `json:"key" yaml:"key"`
	Value             interface{}             `json:"value" yaml:"value"`
	Env               string                  `json:"-" yaml:"-"`
	Type              ServiceParameterType    `json:"type" yaml:"type"`
	AllowedValues     []interface{}           `json:"allowedValues" yaml:"allowedValues"`
	DefaultValue      interface{}             `json:"defaultValue" yaml:"defaultValue"`
	MountPath         string                  `json:"-" yaml:"-"`
	Backend           ServiceParameterBackend `json:"-" yaml:"-"`
	Optional          bool                    `json:"optional" yaml:"optional"`
	Advanced          bool                    `json:"advanced" yaml:"advanced"`
	VersionConstraint string                  `json:"-" yaml:"-"`
}

// NewServiceParameter creates a new parameter.
func NewServiceParameter() *ServiceParameter {

	return &ServiceParameter{}
}

// Copy returns a copy of the current parameter.
func (p *ServiceParameter) Copy() *ServiceParameter {

	copy := NewServiceParameter()
	copy.Name = p.Name
	copy.Description = p.Description
	copy.LongDescription = p.LongDescription
	copy.Key = p.Key
	copy.Value = p.Value
	copy.Env = p.Env
	copy.Type = p.Type
	copy.AllowedValues = append(copy.AllowedValues, p.AllowedValues...)
	copy.DefaultValue = p.DefaultValue
	copy.MountPath = p.MountPath
	copy.Backend = p.Backend
	copy.Optional = p.Optional
	copy.Advanced = p.Advanced
	copy.VersionConstraint = p.VersionConstraint

	return copy
}

// Validate validates the service parameter.
func (p *ServiceParameter) Validate() error {

	switch p.Type {
	case ServiceParameterTypeString, ServiceParameterTypePassword:
		return p.validateStringValue()

	case ServiceParameterTypeInt:
		return p.validateIntValue()

	case ServiceParameterTypeFloat:
		return p.validateFloatValue()

	case ServiceParameterTypeBool:
		return p.validateBoolValue()

	case ServiceParameterTypeDuration:
		return p.validateDurationValue()

	case ServiceParameterTypeStringSlice,
		ServiceParameterTypeIntSlice,
		ServiceParameterTypeFloatSlice:
		return p.validateSliceValue()
	}

	return nil
}

// ValueToString returns the value as a string.
func (p *ServiceParameter) ValueToString() string {

	switch p.Type {
	case ServiceParameterTypePassword, ServiceParameterTypeString:
		if value, ok := p.Value.(string); ok {
			return value
		}

	case ServiceParameterTypeBool:
		if value, ok := p.Value.(bool); ok {
			return strconv.FormatBool(value)
		}

	case ServiceParameterTypeInt:
		if value, ok := p.Value.(int); ok {
			return strconv.Itoa(value)
		}

		if value, ok := p.Value.(float64); ok {
			return strconv.FormatFloat(value, 'f', -1, 32)
		}

	case ServiceParameterTypeFloat:
		if value, ok := p.Value.(float64); ok {
			return strconv.FormatFloat(value, 'f', -1, 32)
		}

	case ServiceParameterTypeDuration:
		if value, ok := p.Value.(string); ok {
			return value
		}

	case ServiceParameterTypeStringSlice:
		values := []string{}
		if vs, ok := p.Value.([]interface{}); ok {
			for _, v := range vs {
				values = append(values, v.(string))
			}
		}
		return strings.Join(values, " ")

	case ServiceParameterTypeEnum:
		switch p.Value.(type) {
		case string:
			return p.Value.(string)
		case bool:
			return strconv.FormatBool(p.Value.(bool))
		case int:
			return strconv.Itoa(p.Value.(int))
		case float64:
			return strconv.FormatFloat(p.Value.(float64), 'f', -1, 64)
		}
		return ""

	case ServiceParameterTypeIntSlice:
		values := []string{}
		if vs, ok := p.Value.([]interface{}); ok {
			for _, v := range vs {
				values = append(values, strconv.Itoa(v.(int)))
			}
		}
		return strings.Join(values, " ")

	case ServiceParameterTypeFloatSlice:
		values := []string{}
		if vs, ok := p.Value.([]interface{}); ok {
			for _, v := range vs {
				values = append(values, strconv.FormatFloat(v.(float64), 'f', -1, 64))
			}
		}
		return strings.Join(values, " ")
	}

	return ""
}

// TODO: Remove Service Related Object. Keep it for now for backward compatibility. (CS 02/15/2018)
// ServiceRelatedObject defines a related object.
type ServiceRelatedObject struct {
	Namespace string `json:"-"`
	Identity  string `json:"-"`
	ID        string `json:"-"`
}

// NewServiceRelatedObject creates a new related object.
func NewServiceRelatedObject() *ServiceRelatedObject {

	return &ServiceRelatedObject{}
}

// ServiceRelatedObjectOption is to prepare the future :)
type ServiceRelatedObjectOption struct {
}

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

	if !p.Optional && (p.Value == nil || p.Value == 0) {
		return fmt.Errorf("%s is required", p.Name)
	}

	if p.Value == nil {
		return nil
	}

	// Convert from float64 as it seems that's what we receive from a json request.
	if v, ok := p.Value.(float64); ok {
		p.Value = int(v)
	}

	if _, ok := p.Value.(int); !ok {
		return fmt.Errorf("%s is not a valid integer", p.Name)
	}
	return nil
}

// validateFloatValue validates a float parameter.
func (p *ServiceParameter) validateFloatValue() error {

	if !p.Optional && p.Value == nil {
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

	if !p.Optional && p.Value == nil {
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

	if !p.Optional && p.Value == nil {
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

// validateStringSliceValue validates a string slice parameter.
func (p *ServiceParameter) validateSliceValue() error {

	if !p.Optional && p.Value == nil {
		return fmt.Errorf("%s is required", p.Name)
	}

	if p.Value == nil {
		return nil
	}

	values, ok := p.Value.([]interface{})
	if !ok {
		return fmt.Errorf("%s is not of type array", p.Name)
	}

	if !p.Optional && len(values) == 0 {
		return fmt.Errorf("%s is required", p.Name)
	}

	if len(values) == 0 {
		return nil
	}

	for _, value := range values {
		if err := isAllowedValue(p.AllowedValues, value, p.Type); err != nil {
			return fmt.Errorf("%s has not allowed values: %s", p.Name, err.Error())
		}
	}

	return nil
}

// isStringAllowedValue returns true if the value is allowed
func isAllowedValue(allowedValues []interface{}, value interface{}, parameterType ServiceParameterType) error {

	switch parameterType {
	case ServiceParameterTypeStringSlice:
		_, ok := value.(string)
		if !ok {
			return fmt.Errorf("%d is not a string", value)
		}

	case ServiceParameterTypeIntSlice:
		_, ok := value.(int)
		if !ok {
			return fmt.Errorf("%d is not an int", value)
		}

	case ServiceParameterTypeFloatSlice:
		_, ok := value.(float64)
		if !ok {
			return fmt.Errorf("%d is not a float", value)
		}
	}

	for _, allowed := range allowedValues {
		if value == allowed {
			return nil
		}
	}
	return fmt.Errorf("%s is not allowed", value)
}
