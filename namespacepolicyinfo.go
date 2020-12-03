package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NamespacePolicyInfoBehaviorValue represents the possible values for attribute "behavior".
type NamespacePolicyInfoBehaviorValue string

const (
	// NamespacePolicyInfoBehaviorAllow represents the value Allow.
	NamespacePolicyInfoBehaviorAllow NamespacePolicyInfoBehaviorValue = "Allow"

	// NamespacePolicyInfoBehaviorInherit represents the value Inherit.
	NamespacePolicyInfoBehaviorInherit NamespacePolicyInfoBehaviorValue = "Inherit"

	// NamespacePolicyInfoBehaviorReject represents the value Reject.
	NamespacePolicyInfoBehaviorReject NamespacePolicyInfoBehaviorValue = "Reject"
)

// NamespacePolicyInfoIdentity represents the Identity of the object.
var NamespacePolicyInfoIdentity = elemental.Identity{
	Name:     "namespacepolicyinfo",
	Category: "namespacepolicyinfo",
	Package:  "squall",
	Private:  false,
}

// NamespacePolicyInfosList represents a list of NamespacePolicyInfos
type NamespacePolicyInfosList []*NamespacePolicyInfo

// Identity returns the identity of the objects in the list.
func (o NamespacePolicyInfosList) Identity() elemental.Identity {

	return NamespacePolicyInfoIdentity
}

// Copy returns a pointer to a copy the NamespacePolicyInfosList.
func (o NamespacePolicyInfosList) Copy() elemental.Identifiables {

	copy := append(NamespacePolicyInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the NamespacePolicyInfosList.
func (o NamespacePolicyInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(NamespacePolicyInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*NamespacePolicyInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o NamespacePolicyInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o NamespacePolicyInfosList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the NamespacePolicyInfosList converted to SparseNamespacePolicyInfosList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o NamespacePolicyInfosList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseNamespacePolicyInfosList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseNamespacePolicyInfo)
	}

	return out
}

// Version returns the version of the content.
func (o NamespacePolicyInfosList) Version() int {

	return 1
}

// NamespacePolicyInfo represents the model of a namespacepolicyinfo
type NamespacePolicyInfo struct {
	// The default enforcer behavior for the namespace.
	Behavior NamespacePolicyInfoBehaviorValue `json:"behavior" msgpack:"behavior" bson:"behavior" mapstructure:"behavior,omitempty"`

	// List of tag prefixes that will be used to suggest policies.
	Prefixes []string `json:"prefixes" msgpack:"prefixes" bson:"prefixes" mapstructure:"prefixes,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewNamespacePolicyInfo returns a new *NamespacePolicyInfo
func NewNamespacePolicyInfo() *NamespacePolicyInfo {

	return &NamespacePolicyInfo{
		ModelVersion: 1,
		Prefixes:     []string{},
	}
}

// Identity returns the Identity of the object.
func (o *NamespacePolicyInfo) Identity() elemental.Identity {

	return NamespacePolicyInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NamespacePolicyInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NamespacePolicyInfo) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespacePolicyInfo) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesNamespacePolicyInfo{}

	s.Behavior = o.Behavior
	s.Prefixes = o.Prefixes

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespacePolicyInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesNamespacePolicyInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Behavior = s.Behavior
	o.Prefixes = s.Prefixes

	return nil
}

// Version returns the hardcoded version of the model.
func (o *NamespacePolicyInfo) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *NamespacePolicyInfo) BleveType() string {

	return "namespacepolicyinfo"
}

// DefaultOrder returns the list of default ordering fields.
func (o *NamespacePolicyInfo) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *NamespacePolicyInfo) Doc() string {

	return `Returns the policy info of the specified namespace.`
}

func (o *NamespacePolicyInfo) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *NamespacePolicyInfo) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseNamespacePolicyInfo{
			Behavior: &o.Behavior,
			Prefixes: &o.Prefixes,
		}
	}

	sp := &SparseNamespacePolicyInfo{}
	for _, f := range fields {
		switch f {
		case "behavior":
			sp.Behavior = &(o.Behavior)
		case "prefixes":
			sp.Prefixes = &(o.Prefixes)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseNamespacePolicyInfo to the object.
func (o *NamespacePolicyInfo) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseNamespacePolicyInfo)
	if so.Behavior != nil {
		o.Behavior = *so.Behavior
	}
	if so.Prefixes != nil {
		o.Prefixes = *so.Prefixes
	}
}

// DeepCopy returns a deep copy if the NamespacePolicyInfo.
func (o *NamespacePolicyInfo) DeepCopy() *NamespacePolicyInfo {

	if o == nil {
		return nil
	}

	out := &NamespacePolicyInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *NamespacePolicyInfo.
func (o *NamespacePolicyInfo) DeepCopyInto(out *NamespacePolicyInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy NamespacePolicyInfo: %s", err))
	}

	*out = *target.(*NamespacePolicyInfo)
}

// Validate valides the current information stored into the structure.
func (o *NamespacePolicyInfo) Validate() error {

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
func (*NamespacePolicyInfo) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := NamespacePolicyInfoAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return NamespacePolicyInfoLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*NamespacePolicyInfo) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NamespacePolicyInfoAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *NamespacePolicyInfo) ValueForAttribute(name string) interface{} {

	switch name {
	case "behavior":
		return o.Behavior
	case "prefixes":
		return o.Prefixes
	}

	return nil
}

// NamespacePolicyInfoAttributesMap represents the map of attribute for NamespacePolicyInfo.
var NamespacePolicyInfoAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Prefixes": {
		AllowedChoices: []string{},
		BSONFieldName:  "prefixes",
		ConvertedName:  "Prefixes",
		Description:    `List of tag prefixes that will be used to suggest policies.`,
		Exposed:        true,
		Name:           "prefixes",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

// NamespacePolicyInfoLowerCaseAttributesMap represents the map of attribute for NamespacePolicyInfo.
var NamespacePolicyInfoLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"prefixes": {
		AllowedChoices: []string{},
		BSONFieldName:  "prefixes",
		ConvertedName:  "Prefixes",
		Description:    `List of tag prefixes that will be used to suggest policies.`,
		Exposed:        true,
		Name:           "prefixes",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

// SparseNamespacePolicyInfosList represents a list of SparseNamespacePolicyInfos
type SparseNamespacePolicyInfosList []*SparseNamespacePolicyInfo

// Identity returns the identity of the objects in the list.
func (o SparseNamespacePolicyInfosList) Identity() elemental.Identity {

	return NamespacePolicyInfoIdentity
}

// Copy returns a pointer to a copy the SparseNamespacePolicyInfosList.
func (o SparseNamespacePolicyInfosList) Copy() elemental.Identifiables {

	copy := append(SparseNamespacePolicyInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseNamespacePolicyInfosList.
func (o SparseNamespacePolicyInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseNamespacePolicyInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseNamespacePolicyInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseNamespacePolicyInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseNamespacePolicyInfosList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseNamespacePolicyInfosList converted to NamespacePolicyInfosList.
func (o SparseNamespacePolicyInfosList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseNamespacePolicyInfosList) Version() int {

	return 1
}

// SparseNamespacePolicyInfo represents the sparse version of a namespacepolicyinfo.
type SparseNamespacePolicyInfo struct {
	// The default enforcer behavior for the namespace.
	Behavior *NamespacePolicyInfoBehaviorValue `json:"behavior,omitempty" msgpack:"behavior,omitempty" bson:"behavior,omitempty" mapstructure:"behavior,omitempty"`

	// List of tag prefixes that will be used to suggest policies.
	Prefixes *[]string `json:"prefixes,omitempty" msgpack:"prefixes,omitempty" bson:"prefixes,omitempty" mapstructure:"prefixes,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseNamespacePolicyInfo returns a new  SparseNamespacePolicyInfo.
func NewSparseNamespacePolicyInfo() *SparseNamespacePolicyInfo {
	return &SparseNamespacePolicyInfo{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseNamespacePolicyInfo) Identity() elemental.Identity {

	return NamespacePolicyInfoIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseNamespacePolicyInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseNamespacePolicyInfo) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespacePolicyInfo) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseNamespacePolicyInfo{}

	if o.Behavior != nil {
		s.Behavior = o.Behavior
	}
	if o.Prefixes != nil {
		s.Prefixes = o.Prefixes
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespacePolicyInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseNamespacePolicyInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Behavior != nil {
		o.Behavior = s.Behavior
	}
	if s.Prefixes != nil {
		o.Prefixes = s.Prefixes
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseNamespacePolicyInfo) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseNamespacePolicyInfo) ToPlain() elemental.PlainIdentifiable {

	out := NewNamespacePolicyInfo()
	if o.Behavior != nil {
		out.Behavior = *o.Behavior
	}
	if o.Prefixes != nil {
		out.Prefixes = *o.Prefixes
	}

	return out
}

// DeepCopy returns a deep copy if the SparseNamespacePolicyInfo.
func (o *SparseNamespacePolicyInfo) DeepCopy() *SparseNamespacePolicyInfo {

	if o == nil {
		return nil
	}

	out := &SparseNamespacePolicyInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseNamespacePolicyInfo.
func (o *SparseNamespacePolicyInfo) DeepCopyInto(out *SparseNamespacePolicyInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseNamespacePolicyInfo: %s", err))
	}

	*out = *target.(*SparseNamespacePolicyInfo)
}

type mongoAttributesNamespacePolicyInfo struct {
	Behavior NamespacePolicyInfoBehaviorValue `bson:"behavior"`
	Prefixes []string                         `bson:"prefixes"`
}
type mongoAttributesSparseNamespacePolicyInfo struct {
	Behavior *NamespacePolicyInfoBehaviorValue `bson:"behavior,omitempty"`
	Prefixes *[]string                         `bson:"prefixes,omitempty"`
}
