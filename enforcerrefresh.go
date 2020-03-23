package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EnforcerRefreshDebugValue represents the possible values for attribute "debug".
type EnforcerRefreshDebugValue string

const (
	// EnforcerRefreshDebugCounters represents the value Counters.
	EnforcerRefreshDebugCounters EnforcerRefreshDebugValue = "Counters"

	// EnforcerRefreshDebugLogs represents the value Logs.
	EnforcerRefreshDebugLogs EnforcerRefreshDebugValue = "Logs"
)

// EnforcerRefreshIdentity represents the Identity of the object.
var EnforcerRefreshIdentity = elemental.Identity{
	Name:     "enforcerrefresh",
	Category: "enforcerrefreshes",
	Package:  "gaga",
	Private:  false,
}

// EnforcerRefreshsList represents a list of EnforcerRefreshs
type EnforcerRefreshsList []*EnforcerRefresh

// Identity returns the identity of the objects in the list.
func (o EnforcerRefreshsList) Identity() elemental.Identity {

	return EnforcerRefreshIdentity
}

// Copy returns a pointer to a copy the EnforcerRefreshsList.
func (o EnforcerRefreshsList) Copy() elemental.Identifiables {

	copy := append(EnforcerRefreshsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcerRefreshsList.
func (o EnforcerRefreshsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EnforcerRefreshsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EnforcerRefresh))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerRefreshsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerRefreshsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the EnforcerRefreshsList converted to SparseEnforcerRefreshsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EnforcerRefreshsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseEnforcerRefreshsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseEnforcerRefresh)
	}

	return out
}

// Version returns the version of the content.
func (o EnforcerRefreshsList) Version() int {

	return 1
}

// EnforcerRefresh represents the model of a enforcerrefresh
type EnforcerRefresh struct {
	// Contains the ID of the target enforcer.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Set the debug information collected by the enforcer.
	Debug EnforcerRefreshDebugValue `json:"debug,omitempty" msgpack:"debug,omitempty" bson:"-" mapstructure:"debug,omitempty"`

	// Contains the original namespace of the enforcer.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEnforcerRefresh returns a new *EnforcerRefresh
func NewEnforcerRefresh() *EnforcerRefresh {

	return &EnforcerRefresh{
		ModelVersion: 1,
		Debug:        EnforcerRefreshDebugCounters,
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerRefresh) Identity() elemental.Identity {

	return EnforcerRefreshIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerRefresh) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerRefresh) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EnforcerRefresh) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesEnforcerRefresh{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EnforcerRefresh) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesEnforcerRefresh{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *EnforcerRefresh) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *EnforcerRefresh) BleveType() string {

	return "enforcerrefresh"
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerRefresh) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *EnforcerRefresh) Doc() string {

	return `Sent to client when a poke has been triggered using the
parameter ` + "`" + `?notify=true` + "`" + `. This is used by instances of enforcerd to notify an
external change on the processing unit must be processed.`
}

func (o *EnforcerRefresh) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetID returns the ID of the receiver.
func (o *EnforcerRefresh) GetID() string {

	return o.ID
}

// SetID sets the property ID of the receiver using the given value.
func (o *EnforcerRefresh) SetID(ID string) {

	o.ID = ID
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *EnforcerRefresh) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEnforcerRefresh{
			ID:        &o.ID,
			Debug:     &o.Debug,
			Namespace: &o.Namespace,
		}
	}

	sp := &SparseEnforcerRefresh{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "debug":
			sp.Debug = &(o.Debug)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseEnforcerRefresh to the object.
func (o *EnforcerRefresh) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEnforcerRefresh)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Debug != nil {
		o.Debug = *so.Debug
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
}

// DeepCopy returns a deep copy if the EnforcerRefresh.
func (o *EnforcerRefresh) DeepCopy() *EnforcerRefresh {

	if o == nil {
		return nil
	}

	out := &EnforcerRefresh{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *EnforcerRefresh.
func (o *EnforcerRefresh) DeepCopyInto(out *EnforcerRefresh) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy EnforcerRefresh: %s", err))
	}

	*out = *target.(*EnforcerRefresh)
}

// Validate valides the current information stored into the structure.
func (o *EnforcerRefresh) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("debug", string(o.Debug), []string{"Counters", "Logs"}, false); err != nil {
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
func (*EnforcerRefresh) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerRefreshAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerRefreshLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerRefresh) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerRefreshAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *EnforcerRefresh) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "debug":
		return o.Debug
	case "namespace":
		return o.Namespace
	}

	return nil
}

// EnforcerRefreshAttributesMap represents the map of attribute for EnforcerRefresh.
var EnforcerRefreshAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `Contains the ID of the target enforcer.`,
		Exposed:        true,
		Getter:         true,
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Setter:         true,
		Type:           "string",
	},
	"Debug": elemental.AttributeSpecification{
		AllowedChoices: []string{"Counters", "Logs"},
		ConvertedName:  "Debug",
		DefaultValue:   EnforcerRefreshDebugCounters,
		Description:    `Set the debug information collected by the enforcer.`,
		Exposed:        true,
		Name:           "debug",
		Type:           "enum",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Contains the original namespace of the enforcer.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
}

// EnforcerRefreshLowerCaseAttributesMap represents the map of attribute for EnforcerRefresh.
var EnforcerRefreshLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `Contains the ID of the target enforcer.`,
		Exposed:        true,
		Getter:         true,
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Setter:         true,
		Type:           "string",
	},
	"debug": elemental.AttributeSpecification{
		AllowedChoices: []string{"Counters", "Logs"},
		ConvertedName:  "Debug",
		DefaultValue:   EnforcerRefreshDebugCounters,
		Description:    `Set the debug information collected by the enforcer.`,
		Exposed:        true,
		Name:           "debug",
		Type:           "enum",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Contains the original namespace of the enforcer.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
}

// SparseEnforcerRefreshsList represents a list of SparseEnforcerRefreshs
type SparseEnforcerRefreshsList []*SparseEnforcerRefresh

// Identity returns the identity of the objects in the list.
func (o SparseEnforcerRefreshsList) Identity() elemental.Identity {

	return EnforcerRefreshIdentity
}

// Copy returns a pointer to a copy the SparseEnforcerRefreshsList.
func (o SparseEnforcerRefreshsList) Copy() elemental.Identifiables {

	copy := append(SparseEnforcerRefreshsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEnforcerRefreshsList.
func (o SparseEnforcerRefreshsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEnforcerRefreshsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEnforcerRefresh))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEnforcerRefreshsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEnforcerRefreshsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseEnforcerRefreshsList converted to EnforcerRefreshsList.
func (o SparseEnforcerRefreshsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEnforcerRefreshsList) Version() int {

	return 1
}

// SparseEnforcerRefresh represents the sparse version of a enforcerrefresh.
type SparseEnforcerRefresh struct {
	// Contains the ID of the target enforcer.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Set the debug information collected by the enforcer.
	Debug *EnforcerRefreshDebugValue `json:"debug,omitempty" msgpack:"debug,omitempty" bson:"-" mapstructure:"debug,omitempty"`

	// Contains the original namespace of the enforcer.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseEnforcerRefresh returns a new  SparseEnforcerRefresh.
func NewSparseEnforcerRefresh() *SparseEnforcerRefresh {
	return &SparseEnforcerRefresh{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEnforcerRefresh) Identity() elemental.Identity {

	return EnforcerRefreshIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEnforcerRefresh) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEnforcerRefresh) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEnforcerRefresh) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseEnforcerRefresh{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEnforcerRefresh) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseEnforcerRefresh{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseEnforcerRefresh) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEnforcerRefresh) ToPlain() elemental.PlainIdentifiable {

	out := NewEnforcerRefresh()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Debug != nil {
		out.Debug = *o.Debug
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}

	return out
}

// GetID returns the ID of the receiver.
func (o *SparseEnforcerRefresh) GetID() (out string) {

	if o.ID == nil {
		return
	}

	return *o.ID
}

// SetID sets the property ID of the receiver using the address of the given value.
func (o *SparseEnforcerRefresh) SetID(ID string) {

	o.ID = &ID
}

// DeepCopy returns a deep copy if the SparseEnforcerRefresh.
func (o *SparseEnforcerRefresh) DeepCopy() *SparseEnforcerRefresh {

	if o == nil {
		return nil
	}

	out := &SparseEnforcerRefresh{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEnforcerRefresh.
func (o *SparseEnforcerRefresh) DeepCopyInto(out *SparseEnforcerRefresh) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEnforcerRefresh: %s", err))
	}

	*out = *target.(*SparseEnforcerRefresh)
}

type mongoAttributesEnforcerRefresh struct {
}
type mongoAttributesSparseEnforcerRefresh struct {
}
