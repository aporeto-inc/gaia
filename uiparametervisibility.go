package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// UIParameterVisibilityOperatorValue represents the possible values for attribute "operator".
type UIParameterVisibilityOperatorValue string

const (
	// UIParameterVisibilityOperatorDefined represents the value Defined.
	UIParameterVisibilityOperatorDefined UIParameterVisibilityOperatorValue = "Defined"

	// UIParameterVisibilityOperatorEqual represents the value Equal.
	UIParameterVisibilityOperatorEqual UIParameterVisibilityOperatorValue = "Equal"

	// UIParameterVisibilityOperatorGreaterThan represents the value GreaterThan.
	UIParameterVisibilityOperatorGreaterThan UIParameterVisibilityOperatorValue = "GreaterThan"

	// UIParameterVisibilityOperatorLesserThan represents the value LesserThan.
	UIParameterVisibilityOperatorLesserThan UIParameterVisibilityOperatorValue = "LesserThan"

	// UIParameterVisibilityOperatorNotEqual represents the value NotEqual.
	UIParameterVisibilityOperatorNotEqual UIParameterVisibilityOperatorValue = "NotEqual"

	// UIParameterVisibilityOperatorUndefined represents the value Undefined.
	UIParameterVisibilityOperatorUndefined UIParameterVisibilityOperatorValue = "Undefined"
)

// UIParameterVisibility represents the model of a uiparametervisibility
type UIParameterVisibility struct {
	// Key holding the value to compare.
	Key string `json:"key" msgpack:"key" bson:"key" mapstructure:"key,omitempty"`

	// Operator to apply.
	Operator UIParameterVisibilityOperatorValue `json:"operator" msgpack:"operator" bson:"operator" mapstructure:"operator,omitempty"`

	// Values that must match the key.
	Value interface{} `json:"value" msgpack:"value" bson:"value" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewUIParameterVisibility returns a new *UIParameterVisibility
func NewUIParameterVisibility() *UIParameterVisibility {

	return &UIParameterVisibility{
		ModelVersion: 1,
	}
}

// DeepCopy returns a deep copy if the UIParameterVisibility.
func (o *UIParameterVisibility) DeepCopy() *UIParameterVisibility {

	if o == nil {
		return nil
	}

	out := &UIParameterVisibility{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *UIParameterVisibility.
func (o *UIParameterVisibility) DeepCopyInto(out *UIParameterVisibility) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy UIParameterVisibility: %s", err))
	}

	*out = *target.(*UIParameterVisibility)
}

// Validate valides the current information stored into the structure.
func (o *UIParameterVisibility) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("key", o.Key); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("operator", string(o.Operator), []string{"Equal", "NotEqual", "GreaterThan", "LesserThan", "Defined", "Undefined"}, false); err != nil {
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
