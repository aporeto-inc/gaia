package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TagPrefixesIdentity represents the Identity of the object.
var TagPrefixesIdentity = elemental.Identity{
	Name:     "tagprefixes",
	Category: "tagprefixes",
	Package:  "squall",
	Private:  false,
}

// TagPrefixesList represents a list of TagPrefixes
type TagPrefixesList []*TagPrefixes

// Identity returns the identity of the objects in the list.
func (o TagPrefixesList) Identity() elemental.Identity {

	return TagPrefixesIdentity
}

// Copy returns a pointer to a copy the TagPrefixesList.
func (o TagPrefixesList) Copy() elemental.Identifiables {

	copy := append(TagPrefixesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TagPrefixesList.
func (o TagPrefixesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TagPrefixesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*TagPrefixes))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TagPrefixesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TagPrefixesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TagPrefixesList converted to SparseTagPrefixesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TagPrefixesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTagPrefixesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTagPrefixes)
	}

	return out
}

// Version returns the version of the content.
func (o TagPrefixesList) Version() int {

	return 1
}

// TagPrefixes represents the model of a tagprefixes
type TagPrefixes struct {
	// List of tag prefixes that will be used to suggest policies.
	Prefixes []string `json:"prefixes" msgpack:"prefixes" bson:"prefixes" mapstructure:"prefixes,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTagPrefixes returns a new *TagPrefixes
func NewTagPrefixes() *TagPrefixes {

	return &TagPrefixes{
		ModelVersion: 1,
		Prefixes:     []string{},
	}
}

// Identity returns the Identity of the object.
func (o *TagPrefixes) Identity() elemental.Identity {

	return TagPrefixesIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TagPrefixes) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TagPrefixes) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TagPrefixes) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTagPrefixes{}

	s.Prefixes = o.Prefixes

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TagPrefixes) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTagPrefixes{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Prefixes = s.Prefixes

	return nil
}

// Version returns the hardcoded version of the model.
func (o *TagPrefixes) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *TagPrefixes) BleveType() string {

	return "tagprefixes"
}

// DefaultOrder returns the list of default ordering fields.
func (o *TagPrefixes) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *TagPrefixes) Doc() string {

	return `Returns the tag prefixes of the specified namespace.`
}

func (o *TagPrefixes) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *TagPrefixes) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTagPrefixes{
			Prefixes: &o.Prefixes,
		}
	}

	sp := &SparseTagPrefixes{}
	for _, f := range fields {
		switch f {
		case "prefixes":
			sp.Prefixes = &(o.Prefixes)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTagPrefixes to the object.
func (o *TagPrefixes) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTagPrefixes)
	if so.Prefixes != nil {
		o.Prefixes = *so.Prefixes
	}
}

// DeepCopy returns a deep copy if the TagPrefixes.
func (o *TagPrefixes) DeepCopy() *TagPrefixes {

	if o == nil {
		return nil
	}

	out := &TagPrefixes{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TagPrefixes.
func (o *TagPrefixes) DeepCopyInto(out *TagPrefixes) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TagPrefixes: %s", err))
	}

	*out = *target.(*TagPrefixes)
}

// Validate valides the current information stored into the structure.
func (o *TagPrefixes) Validate() error {

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
func (*TagPrefixes) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TagPrefixesAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TagPrefixesLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TagPrefixes) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TagPrefixesAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TagPrefixes) ValueForAttribute(name string) interface{} {

	switch name {
	case "prefixes":
		return o.Prefixes
	}

	return nil
}

// TagPrefixesAttributesMap represents the map of attribute for TagPrefixes.
var TagPrefixesAttributesMap = map[string]elemental.AttributeSpecification{
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

// TagPrefixesLowerCaseAttributesMap represents the map of attribute for TagPrefixes.
var TagPrefixesLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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

// SparseTagPrefixesList represents a list of SparseTagPrefixes
type SparseTagPrefixesList []*SparseTagPrefixes

// Identity returns the identity of the objects in the list.
func (o SparseTagPrefixesList) Identity() elemental.Identity {

	return TagPrefixesIdentity
}

// Copy returns a pointer to a copy the SparseTagPrefixesList.
func (o SparseTagPrefixesList) Copy() elemental.Identifiables {

	copy := append(SparseTagPrefixesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTagPrefixesList.
func (o SparseTagPrefixesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTagPrefixesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTagPrefixes))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTagPrefixesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTagPrefixesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTagPrefixesList converted to TagPrefixesList.
func (o SparseTagPrefixesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTagPrefixesList) Version() int {

	return 1
}

// SparseTagPrefixes represents the sparse version of a tagprefixes.
type SparseTagPrefixes struct {
	// List of tag prefixes that will be used to suggest policies.
	Prefixes *[]string `json:"prefixes,omitempty" msgpack:"prefixes,omitempty" bson:"prefixes,omitempty" mapstructure:"prefixes,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTagPrefixes returns a new  SparseTagPrefixes.
func NewSparseTagPrefixes() *SparseTagPrefixes {
	return &SparseTagPrefixes{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTagPrefixes) Identity() elemental.Identity {

	return TagPrefixesIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTagPrefixes) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTagPrefixes) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTagPrefixes) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTagPrefixes{}

	if o.Prefixes != nil {
		s.Prefixes = o.Prefixes
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTagPrefixes) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTagPrefixes{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Prefixes != nil {
		o.Prefixes = s.Prefixes
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTagPrefixes) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTagPrefixes) ToPlain() elemental.PlainIdentifiable {

	out := NewTagPrefixes()
	if o.Prefixes != nil {
		out.Prefixes = *o.Prefixes
	}

	return out
}

// DeepCopy returns a deep copy if the SparseTagPrefixes.
func (o *SparseTagPrefixes) DeepCopy() *SparseTagPrefixes {

	if o == nil {
		return nil
	}

	out := &SparseTagPrefixes{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTagPrefixes.
func (o *SparseTagPrefixes) DeepCopyInto(out *SparseTagPrefixes) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTagPrefixes: %s", err))
	}

	*out = *target.(*SparseTagPrefixes)
}

type mongoAttributesTagPrefixes struct {
	Prefixes []string `bson:"prefixes"`
}
type mongoAttributesSparseTagPrefixes struct {
	Prefixes *[]string `bson:"prefixes,omitempty"`
}
