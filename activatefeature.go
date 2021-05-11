package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ActivateFeatureFeatureValue represents the possible values for attribute "feature".
type ActivateFeatureFeatureValue string

const (
	// ActivateFeatureFeatureNetworkEffectivePermission represents the value NetworkEffectivePermission.
	ActivateFeatureFeatureNetworkEffectivePermission ActivateFeatureFeatureValue = "NetworkEffectivePermission"

	// ActivateFeatureFeatureNetworkSecurity represents the value NetworkSecurity.
	ActivateFeatureFeatureNetworkSecurity ActivateFeatureFeatureValue = "NetworkSecurity"
)

// ActivateFeatureIdentity represents the Identity of the object.
var ActivateFeatureIdentity = elemental.Identity{
	Name:     "activatefeature",
	Category: "activatefeatures",
	Package:  "karl",
	Private:  false,
}

// ActivateFeaturesList represents a list of ActivateFeatures
type ActivateFeaturesList []*ActivateFeature

// Identity returns the identity of the objects in the list.
func (o ActivateFeaturesList) Identity() elemental.Identity {

	return ActivateFeatureIdentity
}

// Copy returns a pointer to a copy the ActivateFeaturesList.
func (o ActivateFeaturesList) Copy() elemental.Identifiables {

	copy := append(ActivateFeaturesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ActivateFeaturesList.
func (o ActivateFeaturesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ActivateFeaturesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ActivateFeature))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ActivateFeaturesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ActivateFeaturesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ActivateFeaturesList converted to SparseActivateFeaturesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ActivateFeaturesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseActivateFeaturesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseActivateFeature)
	}

	return out
}

// Version returns the version of the content.
func (o ActivateFeaturesList) Version() int {

	return 1
}

// ActivateFeature represents the model of a activatefeature
type ActivateFeature struct {
	// Name of the feature to activate for the specified tenant.
	Feature ActivateFeatureFeatureValue `json:"feature" msgpack:"feature" bson:"-" mapstructure:"feature,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewActivateFeature returns a new *ActivateFeature
func NewActivateFeature() *ActivateFeature {

	return &ActivateFeature{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *ActivateFeature) Identity() elemental.Identity {

	return ActivateFeatureIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ActivateFeature) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ActivateFeature) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ActivateFeature) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesActivateFeature{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ActivateFeature) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesActivateFeature{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ActivateFeature) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ActivateFeature) BleveType() string {

	return "activatefeature"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ActivateFeature) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ActivateFeature) Doc() string {

	return `Activates one or multiple features per tenant's Prisma ID.`
}

func (o *ActivateFeature) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ActivateFeature) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseActivateFeature{
			Feature: &o.Feature,
		}
	}

	sp := &SparseActivateFeature{}
	for _, f := range fields {
		switch f {
		case "feature":
			sp.Feature = &(o.Feature)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseActivateFeature to the object.
func (o *ActivateFeature) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseActivateFeature)
	if so.Feature != nil {
		o.Feature = *so.Feature
	}
}

// DeepCopy returns a deep copy if the ActivateFeature.
func (o *ActivateFeature) DeepCopy() *ActivateFeature {

	if o == nil {
		return nil
	}

	out := &ActivateFeature{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ActivateFeature.
func (o *ActivateFeature) DeepCopyInto(out *ActivateFeature) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ActivateFeature: %s", err))
	}

	*out = *target.(*ActivateFeature)
}

// Validate valides the current information stored into the structure.
func (o *ActivateFeature) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("feature", string(o.Feature)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("feature", string(o.Feature), []string{"NetworkEffectivePermission", "NetworkSecurity"}, false); err != nil {
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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*ActivateFeature) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ActivateFeatureAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ActivateFeatureLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ActivateFeature) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ActivateFeatureAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ActivateFeature) ValueForAttribute(name string) interface{} {

	switch name {
	case "feature":
		return o.Feature
	}

	return nil
}

// ActivateFeatureAttributesMap represents the map of attribute for ActivateFeature.
var ActivateFeatureAttributesMap = map[string]elemental.AttributeSpecification{
	"Feature": {
		AllowedChoices: []string{"NetworkEffectivePermission", "NetworkSecurity"},
		ConvertedName:  "Feature",
		Description:    `Name of the feature to activate for the specified tenant.`,
		Exposed:        true,
		Name:           "feature",
		Required:       true,
		Type:           "enum",
	},
}

// ActivateFeatureLowerCaseAttributesMap represents the map of attribute for ActivateFeature.
var ActivateFeatureLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"feature": {
		AllowedChoices: []string{"NetworkEffectivePermission", "NetworkSecurity"},
		ConvertedName:  "Feature",
		Description:    `Name of the feature to activate for the specified tenant.`,
		Exposed:        true,
		Name:           "feature",
		Required:       true,
		Type:           "enum",
	},
}

// SparseActivateFeaturesList represents a list of SparseActivateFeatures
type SparseActivateFeaturesList []*SparseActivateFeature

// Identity returns the identity of the objects in the list.
func (o SparseActivateFeaturesList) Identity() elemental.Identity {

	return ActivateFeatureIdentity
}

// Copy returns a pointer to a copy the SparseActivateFeaturesList.
func (o SparseActivateFeaturesList) Copy() elemental.Identifiables {

	copy := append(SparseActivateFeaturesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseActivateFeaturesList.
func (o SparseActivateFeaturesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseActivateFeaturesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseActivateFeature))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseActivateFeaturesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseActivateFeaturesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseActivateFeaturesList converted to ActivateFeaturesList.
func (o SparseActivateFeaturesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseActivateFeaturesList) Version() int {

	return 1
}

// SparseActivateFeature represents the sparse version of a activatefeature.
type SparseActivateFeature struct {
	// Name of the feature to activate for the specified tenant.
	Feature *ActivateFeatureFeatureValue `json:"feature,omitempty" msgpack:"feature,omitempty" bson:"-" mapstructure:"feature,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseActivateFeature returns a new  SparseActivateFeature.
func NewSparseActivateFeature() *SparseActivateFeature {
	return &SparseActivateFeature{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseActivateFeature) Identity() elemental.Identity {

	return ActivateFeatureIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseActivateFeature) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseActivateFeature) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseActivateFeature) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseActivateFeature{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseActivateFeature) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseActivateFeature{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseActivateFeature) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseActivateFeature) ToPlain() elemental.PlainIdentifiable {

	out := NewActivateFeature()
	if o.Feature != nil {
		out.Feature = *o.Feature
	}

	return out
}

// DeepCopy returns a deep copy if the SparseActivateFeature.
func (o *SparseActivateFeature) DeepCopy() *SparseActivateFeature {

	if o == nil {
		return nil
	}

	out := &SparseActivateFeature{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseActivateFeature.
func (o *SparseActivateFeature) DeepCopyInto(out *SparseActivateFeature) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseActivateFeature: %s", err))
	}

	*out = *target.(*SparseActivateFeature)
}

type mongoAttributesActivateFeature struct {
}
type mongoAttributesSparseActivateFeature struct {
}
