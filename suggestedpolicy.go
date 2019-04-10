package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// SuggestedPolicyIdentity represents the Identity of the object.
var SuggestedPolicyIdentity = elemental.Identity{
	Name:     "suggestedpolicy",
	Category: "suggestedpolicies",
	Package:  "jenova",
	Private:  false,
}

// SuggestedPoliciesList represents a list of SuggestedPolicies
type SuggestedPoliciesList []*SuggestedPolicy

// Identity returns the identity of the objects in the list.
func (o SuggestedPoliciesList) Identity() elemental.Identity {

	return SuggestedPolicyIdentity
}

// Copy returns a pointer to a copy the SuggestedPoliciesList.
func (o SuggestedPoliciesList) Copy() elemental.Identifiables {

	copy := append(SuggestedPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SuggestedPoliciesList.
func (o SuggestedPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SuggestedPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SuggestedPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SuggestedPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SuggestedPoliciesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the SuggestedPoliciesList converted to SparseSuggestedPoliciesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o SuggestedPoliciesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o SuggestedPoliciesList) Version() int {

	return 1
}

// SuggestedPolicy represents the model of a suggestedpolicy
type SuggestedPolicy struct {
	// List of suggested network access policies.
	NetworkAccessPolicies NetworkAccessPoliciesList `json:"networkAccessPolicies" bson:"networkaccesspolicies" mapstructure:"networkAccessPolicies,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSuggestedPolicy returns a new *SuggestedPolicy
func NewSuggestedPolicy() *SuggestedPolicy {

	return &SuggestedPolicy{
		ModelVersion:          1,
		Mutex:                 &sync.Mutex{},
		NetworkAccessPolicies: NetworkAccessPoliciesList{},
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

	return `Allows to get policy suggestions.`
}

func (o *SuggestedPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *SuggestedPolicy) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseSuggestedPolicy{
			NetworkAccessPolicies: &o.NetworkAccessPolicies,
		}
	}

	sp := &SparseSuggestedPolicy{}
	for _, f := range fields {
		switch f {
		case "networkAccessPolicies":
			sp.NetworkAccessPolicies = &(o.NetworkAccessPolicies)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseSuggestedPolicy to the object.
func (o *SuggestedPolicy) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseSuggestedPolicy)
	if so.NetworkAccessPolicies != nil {
		o.NetworkAccessPolicies = *so.NetworkAccessPolicies
	}
}

// DeepCopy returns a deep copy if the SuggestedPolicy.
func (o *SuggestedPolicy) DeepCopy() *SuggestedPolicy {

	if o == nil {
		return nil
	}

	out := &SuggestedPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SuggestedPolicy.
func (o *SuggestedPolicy) DeepCopyInto(out *SuggestedPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SuggestedPolicy: %s", err))
	}

	*out = *target.(*SuggestedPolicy)
}

// Validate valides the current information stored into the structure.
func (o *SuggestedPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.NetworkAccessPolicies {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *SuggestedPolicy) ValueForAttribute(name string) interface{} {

	switch name {
	case "networkAccessPolicies":
		return o.NetworkAccessPolicies
	}

	return nil
}

// SuggestedPolicyAttributesMap represents the map of attribute for SuggestedPolicy.
var SuggestedPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"NetworkAccessPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkAccessPolicies",
		Description:    `List of suggested network access policies.`,
		Exposed:        true,
		Name:           "networkAccessPolicies",
		Orderable:      true,
		Stored:         true,
		SubType:        "networkaccesspolicy",
		Type:           "refList",
	},
}

// SuggestedPolicyLowerCaseAttributesMap represents the map of attribute for SuggestedPolicy.
var SuggestedPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"networkaccesspolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkAccessPolicies",
		Description:    `List of suggested network access policies.`,
		Exposed:        true,
		Name:           "networkAccessPolicies",
		Orderable:      true,
		Stored:         true,
		SubType:        "networkaccesspolicy",
		Type:           "refList",
	},
}

// SparseSuggestedPoliciesList represents a list of SparseSuggestedPolicies
type SparseSuggestedPoliciesList []*SparseSuggestedPolicy

// Identity returns the identity of the objects in the list.
func (o SparseSuggestedPoliciesList) Identity() elemental.Identity {

	return SuggestedPolicyIdentity
}

// Copy returns a pointer to a copy the SparseSuggestedPoliciesList.
func (o SparseSuggestedPoliciesList) Copy() elemental.Identifiables {

	copy := append(SparseSuggestedPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseSuggestedPoliciesList.
func (o SparseSuggestedPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseSuggestedPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseSuggestedPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseSuggestedPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseSuggestedPoliciesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseSuggestedPoliciesList converted to SuggestedPoliciesList.
func (o SparseSuggestedPoliciesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseSuggestedPoliciesList) Version() int {

	return 1
}

// SparseSuggestedPolicy represents the sparse version of a suggestedpolicy.
type SparseSuggestedPolicy struct {
	// List of suggested network access policies.
	NetworkAccessPolicies *NetworkAccessPoliciesList `json:"networkAccessPolicies,omitempty" bson:"networkaccesspolicies,omitempty" mapstructure:"networkAccessPolicies,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSparseSuggestedPolicy returns a new  SparseSuggestedPolicy.
func NewSparseSuggestedPolicy() *SparseSuggestedPolicy {
	return &SparseSuggestedPolicy{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseSuggestedPolicy) Identity() elemental.Identity {

	return SuggestedPolicyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseSuggestedPolicy) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseSuggestedPolicy) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseSuggestedPolicy) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseSuggestedPolicy) ToPlain() elemental.PlainIdentifiable {

	out := NewSuggestedPolicy()
	if o.NetworkAccessPolicies != nil {
		out.NetworkAccessPolicies = *o.NetworkAccessPolicies
	}

	return out
}

// DeepCopy returns a deep copy if the SparseSuggestedPolicy.
func (o *SparseSuggestedPolicy) DeepCopy() *SparseSuggestedPolicy {

	if o == nil {
		return nil
	}

	out := &SparseSuggestedPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseSuggestedPolicy.
func (o *SparseSuggestedPolicy) DeepCopyInto(out *SparseSuggestedPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseSuggestedPolicy: %s", err))
	}

	*out = *target.(*SparseSuggestedPolicy)
}
