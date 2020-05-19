package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PingResultIdentity represents the Identity of the object.
var PingResultIdentity = elemental.Identity{
	Name:     "pingresult",
	Category: "pingresult",
	Package:  "guy",
	Private:  false,
}

// PingResultsList represents a list of PingResults
type PingResultsList []*PingResult

// Identity returns the identity of the objects in the list.
func (o PingResultsList) Identity() elemental.Identity {

	return PingResultIdentity
}

// Copy returns a pointer to a copy the PingResultsList.
func (o PingResultsList) Copy() elemental.Identifiables {

	copy := append(PingResultsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PingResultsList.
func (o PingResultsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PingResultsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PingResult))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PingResultsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PingResultsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PingResultsList converted to SparsePingResultsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PingResultsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePingResultsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePingResult)
	}

	return out
}

// Version returns the version of the content.
func (o PingResultsList) Version() int {

	return 1
}

// PingResult represents the model of a pingresult
type PingResult struct {
	// Contains the result of aggregated ping pairs.
	PingPairs []*PingPair `json:"pingPairs" msgpack:"pingPairs" bson:"-" mapstructure:"pingPairs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPingResult returns a new *PingResult
func NewPingResult() *PingResult {

	return &PingResult{
		ModelVersion: 1,
		PingPairs:    []*PingPair{},
	}
}

// Identity returns the Identity of the object.
func (o *PingResult) Identity() elemental.Identity {

	return PingResultIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PingResult) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PingResult) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingResult) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPingResult{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingResult) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPingResult{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PingResult) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PingResult) BleveType() string {

	return "pingresult"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PingResult) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PingResult) Doc() string {

	return `Post a new pu pingresult.`
}

func (o *PingResult) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PingResult) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePingResult{
			PingPairs: &o.PingPairs,
		}
	}

	sp := &SparsePingResult{}
	for _, f := range fields {
		switch f {
		case "pingPairs":
			sp.PingPairs = &(o.PingPairs)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePingResult to the object.
func (o *PingResult) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePingResult)
	if so.PingPairs != nil {
		o.PingPairs = *so.PingPairs
	}
}

// DeepCopy returns a deep copy if the PingResult.
func (o *PingResult) DeepCopy() *PingResult {

	if o == nil {
		return nil
	}

	out := &PingResult{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PingResult.
func (o *PingResult) DeepCopyInto(out *PingResult) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PingResult: %s", err))
	}

	*out = *target.(*PingResult)
}

// Validate valides the current information stored into the structure.
func (o *PingResult) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.PingPairs {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
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
func (*PingResult) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PingResultAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PingResultLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PingResult) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PingResultAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PingResult) ValueForAttribute(name string) interface{} {

	switch name {
	case "pingPairs":
		return o.PingPairs
	}

	return nil
}

// PingResultAttributesMap represents the map of attribute for PingResult.
var PingResultAttributesMap = map[string]elemental.AttributeSpecification{
	"PingPairs": {
		AllowedChoices: []string{},
		ConvertedName:  "PingPairs",
		Description:    `Contains the result of aggregated ping pairs.`,
		Exposed:        true,
		Name:           "pingPairs",
		SubType:        "pingpair",
		Type:           "refList",
	},
}

// PingResultLowerCaseAttributesMap represents the map of attribute for PingResult.
var PingResultLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"pingpairs": {
		AllowedChoices: []string{},
		ConvertedName:  "PingPairs",
		Description:    `Contains the result of aggregated ping pairs.`,
		Exposed:        true,
		Name:           "pingPairs",
		SubType:        "pingpair",
		Type:           "refList",
	},
}

// SparsePingResultsList represents a list of SparsePingResults
type SparsePingResultsList []*SparsePingResult

// Identity returns the identity of the objects in the list.
func (o SparsePingResultsList) Identity() elemental.Identity {

	return PingResultIdentity
}

// Copy returns a pointer to a copy the SparsePingResultsList.
func (o SparsePingResultsList) Copy() elemental.Identifiables {

	copy := append(SparsePingResultsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePingResultsList.
func (o SparsePingResultsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePingResultsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePingResult))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePingResultsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePingResultsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePingResultsList converted to PingResultsList.
func (o SparsePingResultsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePingResultsList) Version() int {

	return 1
}

// SparsePingResult represents the sparse version of a pingresult.
type SparsePingResult struct {
	// Contains the result of aggregated ping pairs.
	PingPairs *[]*PingPair `json:"pingPairs,omitempty" msgpack:"pingPairs,omitempty" bson:"-" mapstructure:"pingPairs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePingResult returns a new  SparsePingResult.
func NewSparsePingResult() *SparsePingResult {
	return &SparsePingResult{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePingResult) Identity() elemental.Identity {

	return PingResultIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePingResult) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePingResult) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePingResult) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePingResult{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePingResult) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePingResult{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePingResult) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePingResult) ToPlain() elemental.PlainIdentifiable {

	out := NewPingResult()
	if o.PingPairs != nil {
		out.PingPairs = *o.PingPairs
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePingResult.
func (o *SparsePingResult) DeepCopy() *SparsePingResult {

	if o == nil {
		return nil
	}

	out := &SparsePingResult{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePingResult.
func (o *SparsePingResult) DeepCopyInto(out *SparsePingResult) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePingResult: %s", err))
	}

	*out = *target.(*SparsePingResult)
}

type mongoAttributesPingResult struct {
}
type mongoAttributesSparsePingResult struct {
}
