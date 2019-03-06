package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// UIParameterTypeValue represents the possible values for attribute "type".
type UIParameterTypeValue string

const (
	// UIParameterTypeBoolean represents the value Boolean.
	UIParameterTypeBoolean UIParameterTypeValue = "Boolean"

	// UIParameterTypeCVSSThreshold represents the value CVSSThreshold.
	UIParameterTypeCVSSThreshold UIParameterTypeValue = "CVSSThreshold"

	// UIParameterTypeDuration represents the value Duration.
	UIParameterTypeDuration UIParameterTypeValue = "Duration"

	// UIParameterTypeEnum represents the value Enum.
	UIParameterTypeEnum UIParameterTypeValue = "Enum"

	// UIParameterTypeFloat represents the value Float.
	UIParameterTypeFloat UIParameterTypeValue = "Float"

	// UIParameterTypeFloatSlice represents the value FloatSlice.
	UIParameterTypeFloatSlice UIParameterTypeValue = "FloatSlice"

	// UIParameterTypeInteger represents the value Integer.
	UIParameterTypeInteger UIParameterTypeValue = "Integer"

	// UIParameterTypeIntegerSlice represents the value IntegerSlice.
	UIParameterTypeIntegerSlice UIParameterTypeValue = "IntegerSlice"

	// UIParameterTypePassword represents the value Password.
	UIParameterTypePassword UIParameterTypeValue = "Password"

	// UIParameterTypeString represents the value String.
	UIParameterTypeString UIParameterTypeValue = "String"

	// UIParameterTypeStringSlice represents the value StringSlice.
	UIParameterTypeStringSlice UIParameterTypeValue = "StringSlice"
)

// UIParameter represents the model of a uiparameter
type UIParameter struct {
	// Defines if the parameter is an advanced one.
	Advanced bool `json:"advanced" bson:"advanced" mapstructure:"advanced,omitempty"`

	// List of values that can be used.
	AllowedValues []interface{} `json:"allowedValues" bson:"allowedvalues" mapstructure:"allowedValues,omitempty"`

	// Default value of the parameter.
	DefaultValue interface{} `json:"defaultValue" bson:"defaultvalue" mapstructure:"defaultValue,omitempty"`

	// Description of the paramerter.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Key identifying the parameter.
	Key string `json:"key" bson:"key" mapstructure:"key,omitempty"`

	// Long explanation of the parameter.
	LongDescription string `json:"longDescription" bson:"longdescription" mapstructure:"longDescription,omitempty"`

	// Name of the paramerter.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Defines if the parameter is optional.
	Optional bool `json:"optional" bson:"optional" mapstructure:"optional,omitempty"`

	// Set if the parameter must be set.
	Required bool `json:"required" bson:"required" mapstructure:"required,omitempty"`

	// The type of the parameter.
	Type UIParameterTypeValue `json:"type" bson:"type" mapstructure:"type,omitempty"`

	// Value of the parameter.
	Value interface{} `json:"value" bson:"value" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewUIParameter returns a new *UIParameter
func NewUIParameter() *UIParameter {

	return &UIParameter{
		ModelVersion:  1,
		AllowedValues: []interface{}{},
	}
}

// DeepCopy returns a deep copy if the UIParameter.
func (o *UIParameter) DeepCopy() *UIParameter {

	if o == nil {
		return nil
	}

	out := &UIParameter{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *UIParameter.
func (o *UIParameter) DeepCopyInto(out *UIParameter) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy UIParameter: %s", err))
	}

	*out = *target.(*UIParameter)
}

// Validate valides the current information stored into the structure.
func (o *UIParameter) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Boolean", "Duration", "Enum", "IntegerSlice", "Integer", "Float", "FloatSlice", "Password", "String", "StringSlice", "CVSSThreshold"}, false); err != nil {
		errors = append(errors, err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
