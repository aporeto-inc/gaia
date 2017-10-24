package types

import (
	"fmt"
	"reflect"
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
	copy.Key = p.Key
	copy.Value = p.Value
	copy.Env = p.Env
	copy.MountPath = p.MountPath
	copy.Type = p.Type
	copy.Optional = p.Optional
	copy.Backend = p.Backend
	copy.AllowedValues = append(copy.AllowedValues, copy.AllowedValues...)

	return copy
}

// Validate validates the service parameter.
func (p *ServiceParameter) Validate() error {

	switch p.Type {
	case ServiceParameterTypeString:
		return p.validateStringValue()

	case ServiceParameterTypeBool:
		return p.validateBoolValue()

	case ServiceParameterTypeDuration:
		return p.validateDurationValue()

	case ServiceParameterTypeInt:
		return p.validateIntValue()

	case ServiceParameterTypeFloat:
		if _, ok := p.Value.(float32); !ok {
			return fmt.Errorf("%s is not a valid float", p.Name)
		}
		return nil

	case ServiceParameterTypeIntSlice:
		values, ok := p.Value.([]int)
		if !ok {
			return fmt.Errorf("%s is not a valid array of integers", p.Name)
		}

		for _, v := range values {
			if err := isAllowedValue(p.AllowedValues, v); err != nil {
				return fmt.Errorf("%s has incorrect value: %s", p.Name, err.Error())
			}
		}
		return nil

	case ServiceParameterTypeFloatSlice:
		values, ok := p.Value.([]float32)
		if !ok {
			return fmt.Errorf("%s is not a valid array of float", p.Name)
		}

		for _, v := range values {
			if err := isAllowedValue(p.AllowedValues, v); err != nil {
				return fmt.Errorf("%s has incorrect value: %s", p.Name, err.Error())
			}
		}
		return nil

	case ServiceParameterTypeStringSlice:
		values, ok := p.Value.([]string)
		if !ok {
			return fmt.Errorf("%s is not a valid array of string", p.Name)
		}

		for _, v := range values {
			if err := isAllowedValue(p.AllowedValues, v); err != nil {
				return fmt.Errorf("%s has incorrect value: %s", p.Name, err.Error())
			}
		}
		return nil

	}

	return nil
}

// ValueToString returns the value as a string.
func (p *ServiceParameter) ValueToString() string {

	switch p.Type {
	case ServiceParameterTypeBool:
		if value, ok := p.Value.(bool); ok {
			return strconv.FormatBool(value)
		}

	case ServiceParameterTypeDuration:
		return p.Value.(string)

	case ServiceParameterTypeIntSlice:
		values := []string{}
		if ints, ok := p.Value.([]int); ok {
			for _, i := range ints {
				values = append(values, strconv.Itoa(i))
			}
		}
		return strings.Join(values, " ")

	case ServiceParameterTypeInt:
		if value, ok := p.Value.(int); ok {
			return strconv.Itoa(value)
		}

	case ServiceParameterTypeFloat:
		if value, ok := p.Value.(float64); ok {
			return strconv.FormatFloat(value, 'f', -1, 32)
		}

	case ServiceParameterTypeFloatSlice:
		values := []string{}
		if floats, ok := p.Value.([]float64); ok {
			for _, f := range floats {
				values = append(values, strconv.FormatFloat(f, 'f', -1, 32))
			}
		}
		return strings.Join(values, " ")

	case ServiceParameterTypePassword:
		return p.Value.(string)

	case ServiceParameterTypeString:
		return p.Value.(string)

	case ServiceParameterTypeStringSlice:
		fmt.Println(p.Value)
		values, ok := p.Value.([]string)
		if !ok {
			fmt.Println("NON")
		}

		fmt.Println(values)
		return strings.Join(values, " ")

	case ServiceParameterTypeEmum:
		v := reflect.ValueOf(p.Value)
		s := make([]string, v.Len())

		for i := 0; i < v.Len(); i++ {
			s[i] = v.Index(i).String()
		}

		return strings.Join(s, " ")
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
