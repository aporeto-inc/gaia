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

// Copy returns a pointer to a copy the ChecksList.
func (o ChecksList) Copy() elemental.ContentIdentifiable {

	copy := append(ChecksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ChecksList.
func (o ChecksList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(ChecksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Check))
	}

	return out
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

// Version returns the version of the content.
func (o ChecksList) Version() int {

	return 1
}

// Check represents the model of a check
type Check struct {
	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewCheck returns a new *Check
func NewCheck() *Check {

	return &Check{
		ModelVersion: 1,
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
func (o *Check) Version() int {

	return 1
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

	if v, ok := CheckAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CheckLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Check) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CheckAttributesMap
}

// CheckAttributesMap represents the map of attribute for Check.
var CheckAttributesMap = map[string]elemental.AttributeSpecification{}

// CheckLowerCaseAttributesMap represents the map of attribute for Check.
var CheckLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{}
