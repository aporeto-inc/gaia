package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// AppParameterType is the type representing the type of a parameter
type AppParameterType string

// Various values for AppParameterType.
const (
	AppParameterTypeBool        AppParameterType = "bool"
	AppParameterTypeDuration    AppParameterType = "duration"
	AppParameterTypeEnum        AppParameterType = "enum"
	AppParameterTypeIntSlice    AppParameterType = "intSlice"
	AppParameterTypeInt         AppParameterType = "int"
	AppParameterTypeFloat       AppParameterType = "float"
	AppParameterTypeFloatSlice  AppParameterType = "floatSlice"
	AppParameterTypePassword    AppParameterType = "password"
	AppParameterTypeString      AppParameterType = "string"
	AppParameterTypeStringSlice AppParameterType = "stringSlice"
)

// AppParameterBackend defines the link of the service parameter.
type AppParameterBackend int

const (
	// AppParameterBackendGlobalSecret defines a link to installation secret.
	AppParameterBackendGlobalSecret AppParameterBackend = iota

	// AppParameterBackendGlobalConfigMap defines a link to installation config map.
	AppParameterBackendGlobalConfigMap

	// AppParameterBackendLocalSecret defines a link to service secret.
	AppParameterBackendLocalSecret

	// AppParameterBackendLocalConfigMap defines a link to service config map.
	AppParameterBackendLocalConfigMap
)

// AppParameter defines a parameter for the service.
type AppParameter struct {
	Name              string              `json:"name"`
	Description       string              `json:"description"`
	LongDescription   string              `json:"longDescription"`
	Key               string              `json:"key"`
	Value             interface{}         `json:"value"`
	Env               string              `json:"-"`
	Type              AppParameterType    `json:"type"`
	AllowedValues     []interface{}       `json:"allowedValues"`
	DefaultValue      interface{}         `json:"defaultValue"`
	MountPath         string              `json:"-"`
	Backend           AppParameterBackend `json:"-"`
	Optional          bool                `json:"optional"`
	Advanced          bool                `json:"advanced"`
	VersionConstraint string              `json:"-"`
}

// NewAppParameter creates a new parameter.
func NewAppParameter() *AppParameter {

	return &AppParameter{}
}

// Copy returns a copy of the current parameter.
func (p *AppParameter) Copy() *AppParameter {

	copy := NewAppParameter()
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
func (p *AppParameter) Validate() error {

	switch p.Type {
	case AppParameterTypeString, AppParameterTypePassword:
		return p.validateStringValue()

	case AppParameterTypeInt:
		return p.validateIntValue()

	case AppParameterTypeFloat:
		return p.validateFloatValue()

	case AppParameterTypeBool:
		return p.validateBoolValue()

	case AppParameterTypeDuration:
		return p.validateDurationValue()

	case AppParameterTypeStringSlice,
		AppParameterTypeIntSlice,
		AppParameterTypeFloatSlice:
		return p.validateSliceValue()
	}

	return nil
}

// ValueToString returns the value as a string.
func (p *AppParameter) ValueToString() string {

	switch p.Type {
	case AppParameterTypePassword, AppParameterTypeString:
		if value, ok := p.Value.(string); ok {
			return value
		}

	case AppParameterTypeBool:
		if value, ok := p.Value.(bool); ok {
			return strconv.FormatBool(value)
		}

	case AppParameterTypeInt:
		if value, ok := p.Value.(int); ok {
			return strconv.Itoa(value)
		}

		if value, ok := p.Value.(float64); ok {
			return strconv.FormatFloat(value, 'f', -1, 32)
		}

	case AppParameterTypeFloat:
		if value, ok := p.Value.(float64); ok {
			return strconv.FormatFloat(value, 'f', -1, 32)
		}

	case AppParameterTypeDuration:
		if value, ok := p.Value.(string); ok {
			return value
		}

	case AppParameterTypeStringSlice:
		values := []string{}
		if vs, ok := p.Value.([]interface{}); ok {
			for _, v := range vs {
				values = append(values, v.(string))
			}
		}
		return strings.Join(values, " ")

	case AppParameterTypeEnum:
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

	case AppParameterTypeIntSlice:
		values := []string{}
		if vs, ok := p.Value.([]interface{}); ok {
			for _, v := range vs {
				values = append(values, strconv.Itoa(v.(int)))
			}
		}
		return strings.Join(values, " ")

	case AppParameterTypeFloatSlice:
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

// AppRelatedObject defines a related object.
type AppRelatedObject struct {
	Namespace string `json:"-"`
	Identity  string `json:"-"`
	ID        string `json:"-"`
}

// NewAppRelatedObject creates a new related object.
func NewAppRelatedObject() *AppRelatedObject {

	return &AppRelatedObject{}
}

// AppRelatedObjectOption is to prepare the future :)
type AppRelatedObjectOption struct {
}

// validateStringValue validates a string parameter.
func (p *AppParameter) validateStringValue() error {

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
func (p *AppParameter) validateIntValue() error {

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
func (p *AppParameter) validateFloatValue() error {

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
func (p *AppParameter) validateBoolValue() error {

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
func (p *AppParameter) validateDurationValue() error {

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
func (p *AppParameter) validateSliceValue() error {

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
func isAllowedValue(allowedValues []interface{}, value interface{}, parameterType AppParameterType) error {

	switch parameterType {
	case AppParameterTypeStringSlice:
		_, ok := value.(string)
		if !ok {
			return fmt.Errorf("%d is not a string", value)
		}

	case AppParameterTypeIntSlice:
		_, ok := value.(int)
		if !ok {
			return fmt.Errorf("%d is not an int", value)
		}

	case AppParameterTypeFloatSlice:
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
