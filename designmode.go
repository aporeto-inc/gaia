package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// DesignModeIdentity represents the Identity of the object.
var DesignModeIdentity = elemental.Identity{
	Name:     "designmode",
	Category: "designmodes",
	Package:  "yuna",
	Private:  false,
}

// DesignModesList represents a list of DesignModes
type DesignModesList []*DesignMode

// Identity returns the identity of the objects in the list.
func (o DesignModesList) Identity() elemental.Identity {

	return DesignModeIdentity
}

// Copy returns a pointer to a copy the DesignModesList.
func (o DesignModesList) Copy() elemental.Identifiables {

	copy := append(DesignModesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the DesignModesList.
func (o DesignModesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(DesignModesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*DesignMode))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o DesignModesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o DesignModesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the DesignModesList converted to SparseDesignModesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o DesignModesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseDesignModesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseDesignMode)
	}

	return out
}

// Version returns the version of the content.
func (o DesignModesList) Version() int {

	return 1
}

// DesignMode represents the model of a designmode
type DesignMode struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewDesignMode returns a new *DesignMode
func NewDesignMode() *DesignMode {

	return &DesignMode{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *DesignMode) Identity() elemental.Identity {

	return DesignModeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DesignMode) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DesignMode) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DesignMode) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesDesignMode{}

	s.Propagate = o.Propagate

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DesignMode) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesDesignMode{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Propagate = s.Propagate

	return nil
}

// Version returns the hardcoded version of the model.
func (o *DesignMode) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *DesignMode) BleveType() string {

	return "designmode"
}

// DefaultOrder returns the list of default ordering fields.
func (o *DesignMode) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *DesignMode) Doc() string {

	return `When design mode is enabled, all flows are accepted. Flows which do not match an
existing network policy will be represented by a dotted line in your Platform
view.`
}

func (o *DesignMode) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetPropagate returns the Propagate of the receiver.
func (o *DesignMode) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *DesignMode) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *DesignMode) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseDesignMode{
			ID:        &o.ID,
			Propagate: &o.Propagate,
		}
	}

	sp := &SparseDesignMode{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "propagate":
			sp.Propagate = &(o.Propagate)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseDesignMode to the object.
func (o *DesignMode) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseDesignMode)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
}

// DeepCopy returns a deep copy if the DesignMode.
func (o *DesignMode) DeepCopy() *DesignMode {

	if o == nil {
		return nil
	}

	out := &DesignMode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *DesignMode.
func (o *DesignMode) DeepCopyInto(out *DesignMode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy DesignMode: %s", err))
	}

	*out = *target.(*DesignMode)
}

// Validate valides the current information stored into the structure.
func (o *DesignMode) Validate() error {

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
func (*DesignMode) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := DesignModeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return DesignModeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DesignMode) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DesignModeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *DesignMode) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "propagate":
		return o.Propagate
	}

	return nil
}

// DesignModeAttributesMap represents the map of attribute for DesignMode.
var DesignModeAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"Propagate": {
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
}

// DesignModeLowerCaseAttributesMap represents the map of attribute for DesignMode.
var DesignModeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"propagate": {
		AllowedChoices: []string{},
		BSONFieldName:  "propagate",
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
}

// SparseDesignModesList represents a list of SparseDesignModes
type SparseDesignModesList []*SparseDesignMode

// Identity returns the identity of the objects in the list.
func (o SparseDesignModesList) Identity() elemental.Identity {

	return DesignModeIdentity
}

// Copy returns a pointer to a copy the SparseDesignModesList.
func (o SparseDesignModesList) Copy() elemental.Identifiables {

	copy := append(SparseDesignModesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseDesignModesList.
func (o SparseDesignModesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseDesignModesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseDesignMode))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseDesignModesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseDesignModesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseDesignModesList converted to DesignModesList.
func (o SparseDesignModesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseDesignModesList) Version() int {

	return 1
}

// SparseDesignMode represents the sparse version of a designmode.
type SparseDesignMode struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseDesignMode returns a new  SparseDesignMode.
func NewSparseDesignMode() *SparseDesignMode {
	return &SparseDesignMode{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseDesignMode) Identity() elemental.Identity {

	return DesignModeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseDesignMode) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseDesignMode) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDesignMode) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseDesignMode{}

	if o.Propagate != nil {
		s.Propagate = o.Propagate
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDesignMode) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseDesignMode{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Propagate != nil {
		o.Propagate = s.Propagate
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseDesignMode) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseDesignMode) ToPlain() elemental.PlainIdentifiable {

	out := NewDesignMode()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}

	return out
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseDesignMode) GetPropagate() (out bool) {

	if o.Propagate == nil {
		return
	}

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseDesignMode) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// DeepCopy returns a deep copy if the SparseDesignMode.
func (o *SparseDesignMode) DeepCopy() *SparseDesignMode {

	if o == nil {
		return nil
	}

	out := &SparseDesignMode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseDesignMode.
func (o *SparseDesignMode) DeepCopyInto(out *SparseDesignMode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseDesignMode: %s", err))
	}

	*out = *target.(*SparseDesignMode)
}

type mongoAttributesDesignMode struct {
	Propagate bool `bson:"propagate"`
}
type mongoAttributesSparseDesignMode struct {
	Propagate *bool `bson:"propagate,omitempty"`
}
