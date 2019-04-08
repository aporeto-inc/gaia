package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AutomationTemplateParameterTypeValue represents the possible values for attribute "type".
type AutomationTemplateParameterTypeValue string

const (
	// AutomationTemplateParameterTypeArray represents the value Array.
	AutomationTemplateParameterTypeArray AutomationTemplateParameterTypeValue = "Array"

	// AutomationTemplateParameterTypeBoolean represents the value Boolean.
	AutomationTemplateParameterTypeBoolean AutomationTemplateParameterTypeValue = "Boolean"

	// AutomationTemplateParameterTypeEnum represents the value Enum.
	AutomationTemplateParameterTypeEnum AutomationTemplateParameterTypeValue = "Enum"

	// AutomationTemplateParameterTypeFilter represents the value Filter.
	AutomationTemplateParameterTypeFilter AutomationTemplateParameterTypeValue = "Filter"

	// AutomationTemplateParameterTypeFloat represents the value Float.
	AutomationTemplateParameterTypeFloat AutomationTemplateParameterTypeValue = "Float"

	// AutomationTemplateParameterTypeInteger represents the value Integer.
	AutomationTemplateParameterTypeInteger AutomationTemplateParameterTypeValue = "Integer"

	// AutomationTemplateParameterTypeObject represents the value Object.
	AutomationTemplateParameterTypeObject AutomationTemplateParameterTypeValue = "Object"

	// AutomationTemplateParameterTypeString represents the value String.
	AutomationTemplateParameterTypeString AutomationTemplateParameterTypeValue = "String"
)

// AutomationTemplateParameter represents the model of a automationtemplateparameter
type AutomationTemplateParameter struct {
	// Set the possible values for the parameter.
	AllowedChoices map[string]interface{} `json:"allowedChoices" bson:"-" mapstructure:"allowedChoices,omitempty"`

	// Default value of the parameter.
	DefaultValue interface{} `json:"defaultValue" bson:"-" mapstructure:"defaultValue,omitempty"`

	// Name of the parameter.
	Description string `json:"description" bson:"-" mapstructure:"description,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Preferred position for the parameter.
	Position int `json:"position" bson:"-" mapstructure:"position,omitempty"`

	// Set if the parameter must be set.
	Required bool `json:"required" bson:"-" mapstructure:"required,omitempty"`

	// Type of the parameter.
	Type AutomationTemplateParameterTypeValue `json:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewAutomationTemplateParameter returns a new *AutomationTemplateParameter
func NewAutomationTemplateParameter() *AutomationTemplateParameter {

	return &AutomationTemplateParameter{
		ModelVersion:   1,
		Mutex:          &sync.Mutex{},
		AllowedChoices: map[string]interface{}{},
	}
}

// GetName returns the Name of the receiver.
func (o *AutomationTemplateParameter) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *AutomationTemplateParameter) SetName(name string) {

	o.Name = name
}

// DeepCopy returns a deep copy if the AutomationTemplateParameter.
func (o *AutomationTemplateParameter) DeepCopy() *AutomationTemplateParameter {

	if o == nil {
		return nil
	}

	out := &AutomationTemplateParameter{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AutomationTemplateParameter.
func (o *AutomationTemplateParameter) DeepCopyInto(out *AutomationTemplateParameter) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AutomationTemplateParameter: %s", err))
	}

	*out = *target.(*AutomationTemplateParameter)
}

// Validate valides the current information stored into the structure.
func (o *AutomationTemplateParameter) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("type", string(o.Type)); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Array", "Boolean", "Enum", "Filter", "Float", "Integer", "Object", "String"}, false); err != nil {
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
