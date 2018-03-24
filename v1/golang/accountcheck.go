package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// AccountCheckIdentity represents the Identity of the object.
var AccountCheckIdentity = elemental.Identity{
	Name:     "accountcheck",
	Category: "accountchecks",
	Private:  true,
}

// AccountChecksList represents a list of AccountChecks
type AccountChecksList []*AccountCheck

// ContentIdentity returns the identity of the objects in the list.
func (o AccountChecksList) ContentIdentity() elemental.Identity {

	return AccountCheckIdentity
}

// Copy returns a pointer to a copy the AccountChecksList.
func (o AccountChecksList) Copy() elemental.ContentIdentifiable {

	copy := append(AccountChecksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AccountChecksList.
func (o AccountChecksList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(AccountChecksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AccountCheck))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AccountChecksList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AccountChecksList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o AccountChecksList) Version() int {

	return 1
}

// AccountCheck represents the model of a accountcheck
type AccountCheck struct {
	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewAccountCheck returns a new *AccountCheck
func NewAccountCheck() *AccountCheck {

	return &AccountCheck{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *AccountCheck) Identity() elemental.Identity {

	return AccountCheckIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AccountCheck) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AccountCheck) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *AccountCheck) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *AccountCheck) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *AccountCheck) Doc() string {
	return nodocString
}

func (o *AccountCheck) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *AccountCheck) Validate() error {

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
func (*AccountCheck) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AccountCheckAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AccountCheckLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AccountCheck) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AccountCheckAttributesMap
}

// AccountCheckAttributesMap represents the map of attribute for AccountCheck.
var AccountCheckAttributesMap = map[string]elemental.AttributeSpecification{}

// AccountCheckLowerCaseAttributesMap represents the map of attribute for AccountCheck.
var AccountCheckLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{}
