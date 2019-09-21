package gaia

import (
	"fmt"

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

	// UIParameterTypeJSON represents the value JSON.
	UIParameterTypeJSON UIParameterTypeValue = "JSON"

	// UIParameterTypePassword represents the value Password.
	UIParameterTypePassword UIParameterTypeValue = "Password"

	// UIParameterTypeString represents the value String.
	UIParameterTypeString UIParameterTypeValue = "String"

	// UIParameterTypeStringSlice represents the value StringSlice.
	UIParameterTypeStringSlice UIParameterTypeValue = "StringSlice"

	// UIParameterTypeTagsExpression represents the value TagsExpression.
	UIParameterTypeTagsExpression UIParameterTypeValue = "TagsExpression"
)

// UIParameter represents the model of a uiparameter
type UIParameter struct {
	// A value of `true` designates the parameter as advanced.
	Advanced bool `json:"advanced" msgpack:"advanced" bson:"advanced" mapstructure:"advanced,omitempty"`

	// Lists all the choices in case of an enum.
	AllowedChoices map[string]string `json:"allowedChoices" msgpack:"allowedChoices" bson:"allowedchoices" mapstructure:"allowedChoices,omitempty"`

	// List of values that can be used.
	AllowedValues []interface{} `json:"allowedValues" msgpack:"allowedValues" bson:"allowedvalues" mapstructure:"allowedValues,omitempty"`

	// Default value of the parameter.
	DefaultValue interface{} `json:"defaultValue" msgpack:"defaultValue" bson:"defaultvalue" mapstructure:"defaultValue,omitempty"`

	// Description of the parameter.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Key identifying the parameter.
	Key string `json:"key" msgpack:"key" bson:"key" mapstructure:"key,omitempty"`

	// Long explanation of the parameter.
	LongDescription string `json:"longDescription" msgpack:"longDescription" bson:"longdescription" mapstructure:"longDescription,omitempty"`

	// Name of the parameter.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// A value of `true` designates the parameter as optional.
	Optional bool `json:"optional" msgpack:"optional" bson:"optional" mapstructure:"optional,omitempty"`

	// The datatype of the parameter.
	Type UIParameterTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	// A function that validates the parameter.
	ValidationFunction string `json:"validationFunction" msgpack:"validationFunction" bson:"validationfunction" mapstructure:"validationFunction,omitempty"`

	// Value of the parameter.
	Value interface{} `json:"value" msgpack:"value" bson:"value" mapstructure:"value,omitempty"`

	// A logical expression consisting of one or more [UIParameterVisibility](#uiparametervisibility)
	// conditions linked together using AND or OR operators. If the expression evaluates to true
	// the parameter is displayed to the user.
	VisibilityCondition [][]*UIParameterVisibility `json:"visibilityCondition" msgpack:"visibilityCondition" bson:"visibilitycondition" mapstructure:"visibilityCondition,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewUIParameter returns a new *UIParameter
func NewUIParameter() *UIParameter {

	return &UIParameter{
		ModelVersion:        1,
		AllowedChoices:      map[string]string{},
		AllowedValues:       []interface{}{},
		VisibilityCondition: [][]*UIParameterVisibility{},
	}
}

// BleveType implements the bleve.Classifier Interface.
func (o *UIParameter) BleveType() string {

	return "uiparameter"
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

	if err := elemental.ValidateRequiredString("key", o.Key); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("type", string(o.Type)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Boolean", "Duration", "Enum", "IntegerSlice", "Integer", "Float", "FloatSlice", "Password", "String", "StringSlice", "CVSSThreshold", "JSON", "TagsExpression"}, false); err != nil {
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

type mongoAttributesUIParameter struct {
	Advanced            bool                       `bson:"advanced"`
	AllowedChoices      map[string]string          `bson:"allowedchoices"`
	AllowedValues       []interface{}              `bson:"allowedvalues"`
	DefaultValue        interface{}                `bson:"defaultvalue"`
	Description         string                     `bson:"description"`
	Key                 string                     `bson:"key"`
	LongDescription     string                     `bson:"longdescription"`
	Name                string                     `bson:"name"`
	Optional            bool                       `bson:"optional"`
	Type                UIParameterTypeValue       `bson:"type"`
	ValidationFunction  string                     `bson:"validationfunction"`
	Value               interface{}                `bson:"value"`
	VisibilityCondition [][]*UIParameterVisibility `bson:"visibilitycondition"`
}

type mongoAttributesSparseUIParameter struct {
	Advanced            *bool                       `bson:"advanced,omitempty"`
	AllowedChoices      *map[string]string          `bson:"allowedchoices,omitempty"`
	AllowedValues       *[]interface{}              `bson:"allowedvalues,omitempty"`
	DefaultValue        *interface{}                `bson:"defaultvalue,omitempty"`
	Description         *string                     `bson:"description,omitempty"`
	Key                 *string                     `bson:"key,omitempty"`
	LongDescription     *string                     `bson:"longdescription,omitempty"`
	Name                *string                     `bson:"name,omitempty"`
	Optional            *bool                       `bson:"optional,omitempty"`
	Type                *UIParameterTypeValue       `bson:"type,omitempty"`
	ValidationFunction  *string                     `bson:"validationfunction,omitempty"`
	Value               *interface{}                `bson:"value,omitempty"`
	VisibilityCondition *[][]*UIParameterVisibility `bson:"visibilitycondition,omitempty"`
}
