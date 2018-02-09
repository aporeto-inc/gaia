package zackmodels

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// RootIdentity represents the Identity of the object.
var RootIdentity = elemental.Identity{
	Name:     "root",
	Category: "root",
}

// Root represents the model of a root
type Root struct {
	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewRoot returns a new *Root
func NewRoot() *Root {

	return &Root{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Root) Identity() elemental.Identity {

	return RootIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Root) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Root) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Root) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Root) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Root) Doc() string {
	return nodocString
}

func (o *Root) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Root) Validate() error {

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
func (*Root) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RootAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RootLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Root) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RootAttributesMap
}

// RootAttributesMap represents the map of attribute for Root.
var RootAttributesMap = map[string]elemental.AttributeSpecification{}

// RootLowerCaseAttributesMap represents the map of attribute for Root.
var RootLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{}
