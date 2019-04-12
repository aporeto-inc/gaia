package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AppParameterTypeValue represents the possible values for attribute "type".
type AppParameterTypeValue string

const (
	// AppParameterTypeBoolean represents the value Boolean.
	AppParameterTypeBoolean AppParameterTypeValue = "Boolean"

	// AppParameterTypeCVSSThreshold represents the value CVSSThreshold.
	AppParameterTypeCVSSThreshold AppParameterTypeValue = "CVSSThreshold"

	// AppParameterTypeDuration represents the value Duration.
	AppParameterTypeDuration AppParameterTypeValue = "Duration"

	// AppParameterTypeEnum represents the value Enum.
	AppParameterTypeEnum AppParameterTypeValue = "Enum"

	// AppParameterTypeFloat represents the value Float.
	AppParameterTypeFloat AppParameterTypeValue = "Float"

	// AppParameterTypeFloatSlice represents the value FloatSlice.
	AppParameterTypeFloatSlice AppParameterTypeValue = "FloatSlice"

	// AppParameterTypeInteger represents the value Integer.
	AppParameterTypeInteger AppParameterTypeValue = "Integer"

	// AppParameterTypeIntegerSlice represents the value IntegerSlice.
	AppParameterTypeIntegerSlice AppParameterTypeValue = "IntegerSlice"

	// AppParameterTypePassword represents the value Password.
	AppParameterTypePassword AppParameterTypeValue = "Password"

	// AppParameterTypeString represents the value String.
	AppParameterTypeString AppParameterTypeValue = "String"

	// AppParameterTypeStringSlice represents the value StringSlice.
	AppParameterTypeStringSlice AppParameterTypeValue = "StringSlice"
)

// AppParameter represents the model of a appparameter
type AppParameter struct {
	// Defines if the parameter is an advanced one.
	Advanced bool `json:"advanced" msgpack:"advanced" bson:"advanced" mapstructure:"advanced,omitempty"`

	// List of values that can be used.
	AllowedValues []interface{} `json:"allowedValues" msgpack:"allowedValues" bson:"allowedvalues" mapstructure:"allowedValues,omitempty"`

	// Default value of the parameter.
	DefaultValue interface{} `json:"defaultValue" msgpack:"defaultValue" bson:"defaultvalue" mapstructure:"defaultValue,omitempty"`

	// Description of the paramerter.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Key identifying the parameter.
	Key string `json:"key" msgpack:"key" bson:"key" mapstructure:"key,omitempty"`

	// Long explanation of the parameter.
	LongDescription string `json:"longDescription" msgpack:"longDescription" bson:"longdescription" mapstructure:"longDescription,omitempty"`

	// Name of the paramerter.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Defines if the parameter is optional.
	Optional bool `json:"optional" msgpack:"optional" bson:"optional" mapstructure:"optional,omitempty"`

	// The type of the parameter.
	Type AppParameterTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	// Value of the parameter.
	Value interface{} `json:"value" msgpack:"value" bson:"value" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAppParameter returns a new *AppParameter
func NewAppParameter() *AppParameter {

	return &AppParameter{
		ModelVersion:  1,
		AllowedValues: []interface{}{},
	}
}

// DeepCopy returns a deep copy if the AppParameter.
func (o *AppParameter) DeepCopy() *AppParameter {

	if o == nil {
		return nil
	}

	out := &AppParameter{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AppParameter.
func (o *AppParameter) DeepCopyInto(out *AppParameter) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AppParameter: %s", err))
	}

	*out = *target.(*AppParameter)
}

// Validate valides the current information stored into the structure.
func (o *AppParameter) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Boolean", "Duration", "Enum", "IntegerSlice", "Integer", "Float", "FloatSlice", "Password", "String", "StringSlice", "CVSSThreshold"}, false); err != nil {
		errors = errors.Append(err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
