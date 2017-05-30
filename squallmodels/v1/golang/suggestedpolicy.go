package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// SuggestedPolicyIdentity represents the Identity of the object
var SuggestedPolicyIdentity = elemental.Identity{
	Name:     "suggestedpolicy",
	Category: "suggestedpolicies",
}

// SuggestedPoliciesList represents a list of SuggestedPolicies
type SuggestedPoliciesList []*SuggestedPolicy

// ContentIdentity returns the identity of the objects in the list.
func (o SuggestedPoliciesList) ContentIdentity() elemental.Identity {

	return SuggestedPolicyIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o SuggestedPoliciesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SuggestedPoliciesList) DefaultOrder() []string {

	return []string{}
}

// SuggestedPolicy represents the model of a suggestedpolicy
type SuggestedPolicy struct {
	// List of suggested network access policies
	NetworkAccessPolicies []*NetworkAccessPolicy `json:"networkAccessPolicies" bson:"networkaccesspolicies"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewSuggestedPolicy returns a new *SuggestedPolicy
func NewSuggestedPolicy() *SuggestedPolicy {

	return &SuggestedPolicy{
		ModelVersion: 1.0,
	}
}

// Identity returns the Identity of the object.
func (o *SuggestedPolicy) Identity() elemental.Identity {

	return SuggestedPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SuggestedPolicy) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SuggestedPolicy) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *SuggestedPolicy) Version() float64 {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *SuggestedPolicy) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *SuggestedPolicy) Doc() string {
	return `Allows to get policy suggestions`
}

func (o *SuggestedPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *SuggestedPolicy) Validate() error {

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
func (*SuggestedPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return SuggestedPolicyAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*SuggestedPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SuggestedPolicyAttributesMap
}

// SuggestedPolicyAttributesMap represents the map of attribute for SuggestedPolicy.
var SuggestedPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"NetworkAccessPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `List of suggested network access policies`,
		Exposed:        true,
		Filterable:     true,
		Name:           "networkAccessPolicies",
		Orderable:      true,
		Stored:         true,
		SubType:        "network_access_policies_list",
		Type:           "external",
	},
}
