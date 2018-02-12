package models

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// SuggestedPolicyIdentity represents the Identity of the object.
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

// Copy returns a pointer to a copy the SuggestedPoliciesList.
func (o SuggestedPoliciesList) Copy() elemental.ContentIdentifiable {

	copy := append(SuggestedPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SuggestedPoliciesList.
func (o SuggestedPoliciesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(SuggestedPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SuggestedPolicy))
	}

	return out
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

// Version returns the version of the content.
func (o SuggestedPoliciesList) Version() int {

	return 1
}

// SuggestedPolicy represents the model of a suggestedpolicy
type SuggestedPolicy struct {
	// List of suggested network access policies
	NetworkAccessPolicies []*NetworkAccessPolicy `json:"networkAccessPolicies" bson:"networkaccesspolicies" mapstructure:"networkAccessPolicies,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewSuggestedPolicy returns a new *SuggestedPolicy
func NewSuggestedPolicy() *SuggestedPolicy {

	return &SuggestedPolicy{
		ModelVersion: 1,
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
func (o *SuggestedPolicy) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SuggestedPolicy) Version() int {

	return 1
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

	if v, ok := SuggestedPolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return SuggestedPolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*SuggestedPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SuggestedPolicyAttributesMap
}

// SuggestedPolicyAttributesMap represents the map of attribute for SuggestedPolicy.
var SuggestedPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"NetworkAccessPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkAccessPolicies",
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

// SuggestedPolicyLowerCaseAttributesMap represents the map of attribute for SuggestedPolicy.
var SuggestedPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"networkaccesspolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkAccessPolicies",
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
