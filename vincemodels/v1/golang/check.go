package vincemodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// CheckIdentity represents the Identity of the object
var CheckIdentity = elemental.Identity{
	Name:     "check",
	Category: "check",
}

// ChecksList represents a list of Checks
type ChecksList []*Check

// ContentIdentity returns the identity of the objects in the list.
func (o ChecksList) ContentIdentity() elemental.Identity {

	return CheckIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o ChecksList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ChecksList) DefaultOrder() []string {

	return []string{}
}

// Check represents the model of a check
type Check struct {
	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewCheck returns a new *Check
func NewCheck() *Check {

	return &Check{
		ModelVersion: 1.0,
	}
}

// Identity returns the Identity of the object.
func (o *Check) Identity() elemental.Identity {

	return CheckIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Check) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Check) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *Check) Version() float64 {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *Check) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Check) Doc() string {
	return nodocString
}

func (o *Check) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Check) Validate() error {

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
func (*Check) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return CheckAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Check) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CheckAttributesMap
}

// CheckAttributesMap represents the map of attribute for Check.
var CheckAttributesMap = map[string]elemental.AttributeSpecification{}
