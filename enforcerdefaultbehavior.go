package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EnforcerDefaultBehaviorBehaviorValue represents the possible values for attribute "behavior".
type EnforcerDefaultBehaviorBehaviorValue string

const (
	// EnforcerDefaultBehaviorBehaviorAllow represents the value Allow.
	EnforcerDefaultBehaviorBehaviorAllow EnforcerDefaultBehaviorBehaviorValue = "Allow"

	// EnforcerDefaultBehaviorBehaviorInherit represents the value Inherit.
	EnforcerDefaultBehaviorBehaviorInherit EnforcerDefaultBehaviorBehaviorValue = "Inherit"

	// EnforcerDefaultBehaviorBehaviorReject represents the value Reject.
	EnforcerDefaultBehaviorBehaviorReject EnforcerDefaultBehaviorBehaviorValue = "Reject"
)

// EnforcerDefaultBehaviorIdentity represents the Identity of the object.
var EnforcerDefaultBehaviorIdentity = elemental.Identity{
	Name:     "enforcerdefaultbehavior",
	Category: "enforcerdefaultbehavior",
	Package:  "squall",
	Private:  false,
}

// EnforcerDefaultBehaviorsList represents a list of EnforcerDefaultBehaviors
type EnforcerDefaultBehaviorsList []*EnforcerDefaultBehavior

// Identity returns the identity of the objects in the list.
func (o EnforcerDefaultBehaviorsList) Identity() elemental.Identity {

	return EnforcerDefaultBehaviorIdentity
}

// Copy returns a pointer to a copy the EnforcerDefaultBehaviorsList.
func (o EnforcerDefaultBehaviorsList) Copy() elemental.Identifiables {

	copy := append(EnforcerDefaultBehaviorsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcerDefaultBehaviorsList.
func (o EnforcerDefaultBehaviorsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EnforcerDefaultBehaviorsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EnforcerDefaultBehavior))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerDefaultBehaviorsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerDefaultBehaviorsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the EnforcerDefaultBehaviorsList converted to SparseEnforcerDefaultBehaviorsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EnforcerDefaultBehaviorsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseEnforcerDefaultBehaviorsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseEnforcerDefaultBehavior)
	}

	return out
}

// Version returns the version of the content.
func (o EnforcerDefaultBehaviorsList) Version() int {

	return 1
}

// EnforcerDefaultBehavior represents the model of a enforcerdefaultbehavior
type EnforcerDefaultBehavior struct {
	// The default enforcer behavior for the namespace.
	Behavior EnforcerDefaultBehaviorBehaviorValue `json:"behavior" msgpack:"behavior" bson:"behavior" mapstructure:"behavior,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEnforcerDefaultBehavior returns a new *EnforcerDefaultBehavior
func NewEnforcerDefaultBehavior() *EnforcerDefaultBehavior {

	return &EnforcerDefaultBehavior{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerDefaultBehavior) Identity() elemental.Identity {

	return EnforcerDefaultBehaviorIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerDefaultBehavior) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerDefaultBehavior) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EnforcerDefaultBehavior) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesEnforcerDefaultBehavior{}

	s.Behavior = o.Behavior

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EnforcerDefaultBehavior) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesEnforcerDefaultBehavior{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Behavior = s.Behavior

	return nil
}

// Version returns the hardcoded version of the model.
func (o *EnforcerDefaultBehavior) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *EnforcerDefaultBehavior) BleveType() string {

	return "enforcerdefaultbehavior"
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerDefaultBehavior) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *EnforcerDefaultBehavior) Doc() string {

	return `Returns the default enforcer behavior of the specified namespace.`
}

func (o *EnforcerDefaultBehavior) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *EnforcerDefaultBehavior) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEnforcerDefaultBehavior{
			Behavior: &o.Behavior,
		}
	}

	sp := &SparseEnforcerDefaultBehavior{}
	for _, f := range fields {
		switch f {
		case "behavior":
			sp.Behavior = &(o.Behavior)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseEnforcerDefaultBehavior to the object.
func (o *EnforcerDefaultBehavior) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEnforcerDefaultBehavior)
	if so.Behavior != nil {
		o.Behavior = *so.Behavior
	}
}

// DeepCopy returns a deep copy if the EnforcerDefaultBehavior.
func (o *EnforcerDefaultBehavior) DeepCopy() *EnforcerDefaultBehavior {

	if o == nil {
		return nil
	}

	out := &EnforcerDefaultBehavior{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *EnforcerDefaultBehavior.
func (o *EnforcerDefaultBehavior) DeepCopyInto(out *EnforcerDefaultBehavior) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy EnforcerDefaultBehavior: %s", err))
	}

	*out = *target.(*EnforcerDefaultBehavior)
}

// Validate valides the current information stored into the structure.
func (o *EnforcerDefaultBehavior) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("behavior", string(o.Behavior), []string{"Allow", "Reject", "Inherit"}, false); err != nil {
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
func (*EnforcerDefaultBehavior) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerDefaultBehaviorAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerDefaultBehaviorLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerDefaultBehavior) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerDefaultBehaviorAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *EnforcerDefaultBehavior) ValueForAttribute(name string) interface{} {

	switch name {
	case "behavior":
		return o.Behavior
	}

	return nil
}

// EnforcerDefaultBehaviorAttributesMap represents the map of attribute for EnforcerDefaultBehavior.
var EnforcerDefaultBehaviorAttributesMap = map[string]elemental.AttributeSpecification{
	"Behavior": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		BSONFieldName:  "behavior",
		ConvertedName:  "Behavior",
		Description:    `The default enforcer behavior for the namespace.`,
		Exposed:        true,
		Name:           "behavior",
		Stored:         true,
		Type:           "enum",
	},
}

// EnforcerDefaultBehaviorLowerCaseAttributesMap represents the map of attribute for EnforcerDefaultBehavior.
var EnforcerDefaultBehaviorLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"behavior": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		BSONFieldName:  "behavior",
		ConvertedName:  "Behavior",
		Description:    `The default enforcer behavior for the namespace.`,
		Exposed:        true,
		Name:           "behavior",
		Stored:         true,
		Type:           "enum",
	},
}

// SparseEnforcerDefaultBehaviorsList represents a list of SparseEnforcerDefaultBehaviors
type SparseEnforcerDefaultBehaviorsList []*SparseEnforcerDefaultBehavior

// Identity returns the identity of the objects in the list.
func (o SparseEnforcerDefaultBehaviorsList) Identity() elemental.Identity {

	return EnforcerDefaultBehaviorIdentity
}

// Copy returns a pointer to a copy the SparseEnforcerDefaultBehaviorsList.
func (o SparseEnforcerDefaultBehaviorsList) Copy() elemental.Identifiables {

	copy := append(SparseEnforcerDefaultBehaviorsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEnforcerDefaultBehaviorsList.
func (o SparseEnforcerDefaultBehaviorsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEnforcerDefaultBehaviorsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEnforcerDefaultBehavior))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEnforcerDefaultBehaviorsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEnforcerDefaultBehaviorsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseEnforcerDefaultBehaviorsList converted to EnforcerDefaultBehaviorsList.
func (o SparseEnforcerDefaultBehaviorsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEnforcerDefaultBehaviorsList) Version() int {

	return 1
}

// SparseEnforcerDefaultBehavior represents the sparse version of a enforcerdefaultbehavior.
type SparseEnforcerDefaultBehavior struct {
	// The default enforcer behavior for the namespace.
	Behavior *EnforcerDefaultBehaviorBehaviorValue `json:"behavior,omitempty" msgpack:"behavior,omitempty" bson:"behavior,omitempty" mapstructure:"behavior,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseEnforcerDefaultBehavior returns a new  SparseEnforcerDefaultBehavior.
func NewSparseEnforcerDefaultBehavior() *SparseEnforcerDefaultBehavior {
	return &SparseEnforcerDefaultBehavior{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEnforcerDefaultBehavior) Identity() elemental.Identity {

	return EnforcerDefaultBehaviorIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEnforcerDefaultBehavior) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEnforcerDefaultBehavior) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEnforcerDefaultBehavior) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseEnforcerDefaultBehavior{}

	if o.Behavior != nil {
		s.Behavior = o.Behavior
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEnforcerDefaultBehavior) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseEnforcerDefaultBehavior{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Behavior != nil {
		o.Behavior = s.Behavior
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseEnforcerDefaultBehavior) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEnforcerDefaultBehavior) ToPlain() elemental.PlainIdentifiable {

	out := NewEnforcerDefaultBehavior()
	if o.Behavior != nil {
		out.Behavior = *o.Behavior
	}

	return out
}

// DeepCopy returns a deep copy if the SparseEnforcerDefaultBehavior.
func (o *SparseEnforcerDefaultBehavior) DeepCopy() *SparseEnforcerDefaultBehavior {

	if o == nil {
		return nil
	}

	out := &SparseEnforcerDefaultBehavior{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEnforcerDefaultBehavior.
func (o *SparseEnforcerDefaultBehavior) DeepCopyInto(out *SparseEnforcerDefaultBehavior) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEnforcerDefaultBehavior: %s", err))
	}

	*out = *target.(*SparseEnforcerDefaultBehavior)
}

type mongoAttributesEnforcerDefaultBehavior struct {
	Behavior EnforcerDefaultBehaviorBehaviorValue `bson:"behavior"`
}
type mongoAttributesSparseEnforcerDefaultBehavior struct {
	Behavior *EnforcerDefaultBehaviorBehaviorValue `bson:"behavior,omitempty"`
}
