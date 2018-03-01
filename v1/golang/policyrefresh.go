package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// PolicyRefreshIdentity represents the Identity of the object.
var PolicyRefreshIdentity = elemental.Identity{
	Name:     "policyrefresh",
	Category: "policyrefreshs",
	Private:  false,
}

// PolicyRefreshsList represents a list of PolicyRefreshs
type PolicyRefreshsList []*PolicyRefresh

// ContentIdentity returns the identity of the objects in the list.
func (o PolicyRefreshsList) ContentIdentity() elemental.Identity {

	return PolicyRefreshIdentity
}

// Copy returns a pointer to a copy the PolicyRefreshsList.
func (o PolicyRefreshsList) Copy() elemental.ContentIdentifiable {

	copy := append(PolicyRefreshsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PolicyRefreshsList.
func (o PolicyRefreshsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(PolicyRefreshsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PolicyRefresh))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PolicyRefreshsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PolicyRefreshsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o PolicyRefreshsList) Version() int {

	return 1
}

// PolicyRefresh represents the model of a policyrefresh
type PolicyRefresh struct {
	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewPolicyRefresh returns a new *PolicyRefresh
func NewPolicyRefresh() *PolicyRefresh {

	return &PolicyRefresh{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PolicyRefresh) Identity() elemental.Identity {

	return PolicyRefreshIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyRefresh) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyRefresh) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *PolicyRefresh) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *PolicyRefresh) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PolicyRefresh) Doc() string {
	return `PolicyRefresh is sent to client when as a push event when a policy refresh is needed on their side.`
}

func (o *PolicyRefresh) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *PolicyRefresh) Validate() error {

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
func (*PolicyRefresh) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PolicyRefreshAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PolicyRefreshLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PolicyRefresh) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PolicyRefreshAttributesMap
}

// PolicyRefreshAttributesMap represents the map of attribute for PolicyRefresh.
var PolicyRefreshAttributesMap = map[string]elemental.AttributeSpecification{}

// PolicyRefreshLowerCaseAttributesMap represents the map of attribute for PolicyRefresh.
var PolicyRefreshLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{}
