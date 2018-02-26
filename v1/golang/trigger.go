package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// TriggerIdentity represents the Identity of the object.
var TriggerIdentity = elemental.Identity{
	Name:     "trigger",
	Category: "triggers",
	Private:  false,
}

// TriggersList represents a list of Triggers
type TriggersList []*Trigger

// ContentIdentity returns the identity of the objects in the list.
func (o TriggersList) ContentIdentity() elemental.Identity {

	return TriggerIdentity
}

// Copy returns a pointer to a copy the TriggersList.
func (o TriggersList) Copy() elemental.ContentIdentifiable {

	copy := append(TriggersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TriggersList.
func (o TriggersList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(TriggersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Trigger))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TriggersList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TriggersList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o TriggersList) Version() int {

	return 1
}

// Trigger represents the model of a trigger
type Trigger struct {
	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewTrigger returns a new *Trigger
func NewTrigger() *Trigger {

	return &Trigger{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Trigger) Identity() elemental.Identity {

	return TriggerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Trigger) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Trigger) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Trigger) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Trigger) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Trigger) Doc() string {
	return `Trigger can be used to remotely trigger an automation.`
}

func (o *Trigger) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Trigger) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Trigger) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TriggerAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TriggerLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Trigger) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TriggerAttributesMap
}

// TriggerAttributesMap represents the map of attribute for Trigger.
var TriggerAttributesMap = map[string]elemental.AttributeSpecification{}

// TriggerLowerCaseAttributesMap represents the map of attribute for Trigger.
var TriggerLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{}
