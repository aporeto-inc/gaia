package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// OAUTHInfoIdentity represents the Identity of the object.
var OAUTHInfoIdentity = elemental.Identity{
	Name:     "oauthinfo",
	Category: "oauthinfo",
	Package:  "squall",
	Private:  false,
}

// OAUTHInfosList represents a list of OAUTHInfos
type OAUTHInfosList []*OAUTHInfo

// Identity returns the identity of the objects in the list.
func (o OAUTHInfosList) Identity() elemental.Identity {

	return OAUTHInfoIdentity
}

// Copy returns a pointer to a copy the OAUTHInfosList.
func (o OAUTHInfosList) Copy() elemental.Identifiables {

	copy := append(OAUTHInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the OAUTHInfosList.
func (o OAUTHInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(OAUTHInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*OAUTHInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o OAUTHInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o OAUTHInfosList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the OAUTHInfosList converted to SparseOAUTHInfosList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o OAUTHInfosList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseOAUTHInfosList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseOAUTHInfo)
	}

	return out
}

// Version returns the version of the content.
func (o OAUTHInfosList) Version() int {

	return 1
}

// OAUTHInfo represents the model of a oauthinfo
type OAUTHInfo struct {
	// JWKS_URI is the URI that can be used to retrieve the public keys that will
	// verify a JWT.
	Jwks_uri string `json:"jwks_uri" msgpack:"jwks_uri" bson:"-" mapstructure:"jwks_uri,omitempty"`

	// Keys is a list of keys that can be used to verify a JWT.
	Keys map[string]*JWKS `json:"keys" msgpack:"keys" bson:"-" mapstructure:"keys,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewOAUTHInfo returns a new *OAUTHInfo
func NewOAUTHInfo() *OAUTHInfo {

	return &OAUTHInfo{
		ModelVersion: 1,
		Keys:         map[string]*JWKS{},
	}
}

// Identity returns the Identity of the object.
func (o *OAUTHInfo) Identity() elemental.Identity {

	return OAUTHInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *OAUTHInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *OAUTHInfo) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *OAUTHInfo) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *OAUTHInfo) BleveType() string {

	return "oauthinfo"
}

// DefaultOrder returns the list of default ordering fields.
func (o *OAUTHInfo) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *OAUTHInfo) Doc() string {

	return `OAUTHInfo provides the information for an OAUTH server to retrieve the secrets
that can validate a JWT token issued by us.`
}

func (o *OAUTHInfo) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *OAUTHInfo) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseOAUTHInfo{
			Jwks_uri: &o.Jwks_uri,
			Keys:     &o.Keys,
		}
	}

	sp := &SparseOAUTHInfo{}
	for _, f := range fields {
		switch f {
		case "jwks_uri":
			sp.Jwks_uri = &(o.Jwks_uri)
		case "keys":
			sp.Keys = &(o.Keys)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseOAUTHInfo to the object.
func (o *OAUTHInfo) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseOAUTHInfo)
	if so.Jwks_uri != nil {
		o.Jwks_uri = *so.Jwks_uri
	}
	if so.Keys != nil {
		o.Keys = *so.Keys
	}
}

// DeepCopy returns a deep copy if the OAUTHInfo.
func (o *OAUTHInfo) DeepCopy() *OAUTHInfo {

	if o == nil {
		return nil
	}

	out := &OAUTHInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *OAUTHInfo.
func (o *OAUTHInfo) DeepCopyInto(out *OAUTHInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy OAUTHInfo: %s", err))
	}

	*out = *target.(*OAUTHInfo)
}

// Validate valides the current information stored into the structure.
func (o *OAUTHInfo) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Keys {
		if sub == nil {
			continue
		}
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
func (*OAUTHInfo) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := OAUTHInfoAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return OAUTHInfoLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*OAUTHInfo) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return OAUTHInfoAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *OAUTHInfo) ValueForAttribute(name string) interface{} {

	switch name {
	case "jwks_uri":
		return o.Jwks_uri
	case "keys":
		return o.Keys
	}

	return nil
}

// OAUTHInfoAttributesMap represents the map of attribute for OAUTHInfo.
var OAUTHInfoAttributesMap = map[string]elemental.AttributeSpecification{
	"Jwks_uri": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Jwks_uri",
		Description: `JWKS_URI is the URI that can be used to retrieve the public keys that will
verify a JWT.`,
		Exposed:  true,
		Name:     "jwks_uri",
		ReadOnly: true,
		Type:     "string",
	},
	"Keys": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Keys",
		Description:    `Keys is a list of keys that can be used to verify a JWT.`,
		Exposed:        true,
		Name:           "keys",
		ReadOnly:       true,
		SubType:        "jwks",
		Type:           "refMap",
	},
}

// OAUTHInfoLowerCaseAttributesMap represents the map of attribute for OAUTHInfo.
var OAUTHInfoLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"jwks_uri": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Jwks_uri",
		Description: `JWKS_URI is the URI that can be used to retrieve the public keys that will
verify a JWT.`,
		Exposed:  true,
		Name:     "jwks_uri",
		ReadOnly: true,
		Type:     "string",
	},
	"keys": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Keys",
		Description:    `Keys is a list of keys that can be used to verify a JWT.`,
		Exposed:        true,
		Name:           "keys",
		ReadOnly:       true,
		SubType:        "jwks",
		Type:           "refMap",
	},
}

// SparseOAUTHInfosList represents a list of SparseOAUTHInfos
type SparseOAUTHInfosList []*SparseOAUTHInfo

// Identity returns the identity of the objects in the list.
func (o SparseOAUTHInfosList) Identity() elemental.Identity {

	return OAUTHInfoIdentity
}

// Copy returns a pointer to a copy the SparseOAUTHInfosList.
func (o SparseOAUTHInfosList) Copy() elemental.Identifiables {

	copy := append(SparseOAUTHInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseOAUTHInfosList.
func (o SparseOAUTHInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseOAUTHInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseOAUTHInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseOAUTHInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseOAUTHInfosList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseOAUTHInfosList converted to OAUTHInfosList.
func (o SparseOAUTHInfosList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseOAUTHInfosList) Version() int {

	return 1
}

// SparseOAUTHInfo represents the sparse version of a oauthinfo.
type SparseOAUTHInfo struct {
	// JWKS_URI is the URI that can be used to retrieve the public keys that will
	// verify a JWT.
	Jwks_uri *string `json:"jwks_uri,omitempty" msgpack:"jwks_uri,omitempty" bson:"-" mapstructure:"jwks_uri,omitempty"`

	// Keys is a list of keys that can be used to verify a JWT.
	Keys *map[string]*JWKS `json:"keys,omitempty" msgpack:"keys,omitempty" bson:"-" mapstructure:"keys,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseOAUTHInfo returns a new  SparseOAUTHInfo.
func NewSparseOAUTHInfo() *SparseOAUTHInfo {
	return &SparseOAUTHInfo{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseOAUTHInfo) Identity() elemental.Identity {

	return OAUTHInfoIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseOAUTHInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseOAUTHInfo) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseOAUTHInfo) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseOAUTHInfo) ToPlain() elemental.PlainIdentifiable {

	out := NewOAUTHInfo()
	if o.Jwks_uri != nil {
		out.Jwks_uri = *o.Jwks_uri
	}
	if o.Keys != nil {
		out.Keys = *o.Keys
	}

	return out
}

// DeepCopy returns a deep copy if the SparseOAUTHInfo.
func (o *SparseOAUTHInfo) DeepCopy() *SparseOAUTHInfo {

	if o == nil {
		return nil
	}

	out := &SparseOAUTHInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseOAUTHInfo.
func (o *SparseOAUTHInfo) DeepCopyInto(out *SparseOAUTHInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseOAUTHInfo: %s", err))
	}

	*out = *target.(*SparseOAUTHInfo)
}
