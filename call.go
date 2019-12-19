package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CallIdentity represents the Identity of the object.
var CallIdentity = elemental.Identity{
	Name:     "call",
	Category: "calls",
	Package:  "relm",
	Private:  false,
}

// CallsList represents a list of Calls
type CallsList []*Call

// Identity returns the identity of the objects in the list.
func (o CallsList) Identity() elemental.Identity {

	return CallIdentity
}

// Copy returns a pointer to a copy the CallsList.
func (o CallsList) Copy() elemental.Identifiables {

	copy := append(CallsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CallsList.
func (o CallsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CallsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Call))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CallsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CallsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CallsList converted to SparseCallsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CallsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCallsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCall)
	}

	return out
}

// Version returns the version of the content.
func (o CallsList) Version() int {

	return 1
}

// Call represents the model of a call
type Call struct {
	// Contains the remote `POST` payload.
	Payload string `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCall returns a new *Call
func NewCall() *Call {

	return &Call{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Call) Identity() elemental.Identity {

	return CallIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Call) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Call) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Call) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCall{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Call) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCall{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Call) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Call) BleveType() string {

	return "call"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Call) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Call) Doc() string {

	return `Can be used to send a remote request to an API proxy.`
}

func (o *Call) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Call) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCall{
			Payload: &o.Payload,
		}
	}

	sp := &SparseCall{}
	for _, f := range fields {
		switch f {
		case "payload":
			sp.Payload = &(o.Payload)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCall to the object.
func (o *Call) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCall)
	if so.Payload != nil {
		o.Payload = *so.Payload
	}
}

// DeepCopy returns a deep copy if the Call.
func (o *Call) DeepCopy() *Call {

	if o == nil {
		return nil
	}

	out := &Call{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Call.
func (o *Call) DeepCopyInto(out *Call) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Call: %s", err))
	}

	*out = *target.(*Call)
}

// Validate valides the current information stored into the structure.
func (o *Call) Validate() error {

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
func (*Call) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CallAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CallLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Call) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CallAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Call) ValueForAttribute(name string) interface{} {

	switch name {
	case "payload":
		return o.Payload
	}

	return nil
}

// CallAttributesMap represents the map of attribute for Call.
var CallAttributesMap = map[string]elemental.AttributeSpecification{
	"Payload": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Payload",
		Description:    `Contains the remote ` + "`" + `POST` + "`" + ` payload.`,
		Name:           "payload",
		Type:           "string",
	},
}

// CallLowerCaseAttributesMap represents the map of attribute for Call.
var CallLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"payload": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Payload",
		Description:    `Contains the remote ` + "`" + `POST` + "`" + ` payload.`,
		Name:           "payload",
		Type:           "string",
	},
}

// SparseCallsList represents a list of SparseCalls
type SparseCallsList []*SparseCall

// Identity returns the identity of the objects in the list.
func (o SparseCallsList) Identity() elemental.Identity {

	return CallIdentity
}

// Copy returns a pointer to a copy the SparseCallsList.
func (o SparseCallsList) Copy() elemental.Identifiables {

	copy := append(SparseCallsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCallsList.
func (o SparseCallsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCallsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCall))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCallsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCallsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCallsList converted to CallsList.
func (o SparseCallsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCallsList) Version() int {

	return 1
}

// SparseCall represents the sparse version of a call.
type SparseCall struct {
	// Contains the remote `POST` payload.
	Payload *string `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCall returns a new  SparseCall.
func NewSparseCall() *SparseCall {
	return &SparseCall{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCall) Identity() elemental.Identity {

	return CallIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCall) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCall) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCall) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCall{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCall) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCall{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCall) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCall) ToPlain() elemental.PlainIdentifiable {

	out := NewCall()
	if o.Payload != nil {
		out.Payload = *o.Payload
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCall.
func (o *SparseCall) DeepCopy() *SparseCall {

	if o == nil {
		return nil
	}

	out := &SparseCall{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCall.
func (o *SparseCall) DeepCopyInto(out *SparseCall) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCall: %s", err))
	}

	*out = *target.(*SparseCall)
}

type mongoAttributesCall struct {
}
type mongoAttributesSparseCall struct {
}
