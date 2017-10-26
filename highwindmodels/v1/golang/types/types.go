package types

import (
	"strconv"
	"strings"
)

// ServiceParameterType is the type representing the type of a parameter
type ServiceParameterType string

// Various values for ServiceParameterType.
const (
	ServiceParameterTypeBool        ServiceParameterType = "bool"
	ServiceParameterTypeDuration    ServiceParameterType = "duration"
	ServiceParameterTypeEmum        ServiceParameterType = "enum"
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
	Name            string                  `json:"name"`
	Description     string                  `json:"description"`
	LongDescription string                  `json:"longDescription"`
	Key             string                  `json:"key"`
	Value           interface{}             `json:"value"`
	Env             string                  `json:"-"`
	Type            ServiceParameterType    `json:"type"`
	AllowedValues   []interface{}           `json:"allowedValues"`
	DefaultValue    interface{}             `json:"defaultValue"`
	MountPath       string                  `json:"-"`
	Backend         ServiceParameterBackend `json:"-"`
	Optional        bool                    `json:"optional"`
}

// NewServiceParameter creates a new parameter.
func NewServiceParameter() *ServiceParameter {

	return &ServiceParameter{}
}

// Copy returns a copy of the current paramter.
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

	case ServiceParameterTypeFloat:
		if value, ok := p.Value.(float64); ok {
			return strconv.FormatFloat(value, 'f', -1, 32)
		}

	case ServiceParameterTypeDuration:
		if value, ok := p.Value.(string); ok {
			return value
		}

	case ServiceParameterTypeStringSlice, ServiceParameterTypeEmum:
		values := []string{}
		if vs, ok := p.Value.([]interface{}); ok {
			for _, v := range vs {
				values = append(values, v.(string))
			}
		}
		return strings.Join(values, " ")

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
type ServiceRelatedObjectOption struct{}
