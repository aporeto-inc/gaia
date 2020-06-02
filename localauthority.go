package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// LocalAuthorityIdentity represents the Identity of the object.
var LocalAuthorityIdentity = elemental.Identity{
	Name:     "localauthority",
	Category: "localauthorities",
	Package:  "squall",
	Private:  false,
}

// LocalAuthoritiesList represents a list of LocalAuthorities
type LocalAuthoritiesList []*LocalAuthority

// Identity returns the identity of the objects in the list.
func (o LocalAuthoritiesList) Identity() elemental.Identity {

	return LocalAuthorityIdentity
}

// Copy returns a pointer to a copy the LocalAuthoritiesList.
func (o LocalAuthoritiesList) Copy() elemental.Identifiables {

	copy := append(LocalAuthoritiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the LocalAuthoritiesList.
func (o LocalAuthoritiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(LocalAuthoritiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*LocalAuthority))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o LocalAuthoritiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o LocalAuthoritiesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the LocalAuthoritiesList converted to SparseLocalAuthoritiesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o LocalAuthoritiesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseLocalAuthoritiesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseLocalAuthority)
	}

	return out
}

// Version returns the version of the content.
func (o LocalAuthoritiesList) Version() int {

	return 1
}

// LocalAuthority represents the model of a localauthority
type LocalAuthority struct {
	// The certificate authority used by the namespace.
	CA string `json:"CA" msgpack:"CA" bson:"-" mapstructure:"CA,omitempty"`

	// Set to `true` to renew the certificate authority of the namespace.
	CARenew bool `json:"CARenew" msgpack:"CARenew" bson:"-" mapstructure:"CARenew,omitempty"`

	// The SSH certificate authority used by the namespace.
	SSHCA string `json:"SSHCA" msgpack:"SSHCA" bson:"-" mapstructure:"SSHCA,omitempty"`

	// Set to `true` to renew the SSH certificate authority of the namespace.
	SSHCARenew bool `json:"SSHCARenew" msgpack:"SSHCARenew" bson:"-" mapstructure:"SSHCARenew,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewLocalAuthority returns a new *LocalAuthority
func NewLocalAuthority() *LocalAuthority {

	return &LocalAuthority{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *LocalAuthority) Identity() elemental.Identity {

	return LocalAuthorityIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *LocalAuthority) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *LocalAuthority) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *LocalAuthority) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesLocalAuthority{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *LocalAuthority) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesLocalAuthority{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *LocalAuthority) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *LocalAuthority) BleveType() string {

	return "localauthority"
}

// DefaultOrder returns the list of default ordering fields.
func (o *LocalAuthority) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *LocalAuthority) Doc() string {

	return `Can be used to retrieve or renew the local and SSH certificate authorities of
the namespace.`
}

func (o *LocalAuthority) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *LocalAuthority) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseLocalAuthority{
			CA:         &o.CA,
			CARenew:    &o.CARenew,
			SSHCA:      &o.SSHCA,
			SSHCARenew: &o.SSHCARenew,
		}
	}

	sp := &SparseLocalAuthority{}
	for _, f := range fields {
		switch f {
		case "CA":
			sp.CA = &(o.CA)
		case "CARenew":
			sp.CARenew = &(o.CARenew)
		case "SSHCA":
			sp.SSHCA = &(o.SSHCA)
		case "SSHCARenew":
			sp.SSHCARenew = &(o.SSHCARenew)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseLocalAuthority to the object.
func (o *LocalAuthority) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseLocalAuthority)
	if so.CA != nil {
		o.CA = *so.CA
	}
	if so.CARenew != nil {
		o.CARenew = *so.CARenew
	}
	if so.SSHCA != nil {
		o.SSHCA = *so.SSHCA
	}
	if so.SSHCARenew != nil {
		o.SSHCARenew = *so.SSHCARenew
	}
}

// DeepCopy returns a deep copy if the LocalAuthority.
func (o *LocalAuthority) DeepCopy() *LocalAuthority {

	if o == nil {
		return nil
	}

	out := &LocalAuthority{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *LocalAuthority.
func (o *LocalAuthority) DeepCopyInto(out *LocalAuthority) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy LocalAuthority: %s", err))
	}

	*out = *target.(*LocalAuthority)
}

// Validate valides the current information stored into the structure.
func (o *LocalAuthority) Validate() error {

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
func (*LocalAuthority) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := LocalAuthorityAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return LocalAuthorityLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*LocalAuthority) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return LocalAuthorityAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *LocalAuthority) ValueForAttribute(name string) interface{} {

	switch name {
	case "CA":
		return o.CA
	case "CARenew":
		return o.CARenew
	case "SSHCA":
		return o.SSHCA
	case "SSHCARenew":
		return o.SSHCARenew
	}

	return nil
}

// LocalAuthorityAttributesMap represents the map of attribute for LocalAuthority.
var LocalAuthorityAttributesMap = map[string]elemental.AttributeSpecification{
	"CA": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CA",
		Description:    `The certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "CA",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"CARenew": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CARenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "CARenew",
		ReadOnly:       true,
		Transient:      true,
		Type:           "boolean",
	},
	"SSHCA": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCA",
		Description:    `The SSH certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "SSHCA",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"SSHCARenew": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCARenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the SSH certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "SSHCARenew",
		ReadOnly:       true,
		Transient:      true,
		Type:           "boolean",
	},
}

// LocalAuthorityLowerCaseAttributesMap represents the map of attribute for LocalAuthority.
var LocalAuthorityLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"ca": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CA",
		Description:    `The certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "CA",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"carenew": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CARenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "CARenew",
		ReadOnly:       true,
		Transient:      true,
		Type:           "boolean",
	},
	"sshca": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCA",
		Description:    `The SSH certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "SSHCA",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"sshcarenew": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCARenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the SSH certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "SSHCARenew",
		ReadOnly:       true,
		Transient:      true,
		Type:           "boolean",
	},
}

// SparseLocalAuthoritiesList represents a list of SparseLocalAuthorities
type SparseLocalAuthoritiesList []*SparseLocalAuthority

// Identity returns the identity of the objects in the list.
func (o SparseLocalAuthoritiesList) Identity() elemental.Identity {

	return LocalAuthorityIdentity
}

// Copy returns a pointer to a copy the SparseLocalAuthoritiesList.
func (o SparseLocalAuthoritiesList) Copy() elemental.Identifiables {

	copy := append(SparseLocalAuthoritiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseLocalAuthoritiesList.
func (o SparseLocalAuthoritiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseLocalAuthoritiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseLocalAuthority))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseLocalAuthoritiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseLocalAuthoritiesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseLocalAuthoritiesList converted to LocalAuthoritiesList.
func (o SparseLocalAuthoritiesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseLocalAuthoritiesList) Version() int {

	return 1
}

// SparseLocalAuthority represents the sparse version of a localauthority.
type SparseLocalAuthority struct {
	// The certificate authority used by the namespace.
	CA *string `json:"CA,omitempty" msgpack:"CA,omitempty" bson:"-" mapstructure:"CA,omitempty"`

	// Set to `true` to renew the certificate authority of the namespace.
	CARenew *bool `json:"CARenew,omitempty" msgpack:"CARenew,omitempty" bson:"-" mapstructure:"CARenew,omitempty"`

	// The SSH certificate authority used by the namespace.
	SSHCA *string `json:"SSHCA,omitempty" msgpack:"SSHCA,omitempty" bson:"-" mapstructure:"SSHCA,omitempty"`

	// Set to `true` to renew the SSH certificate authority of the namespace.
	SSHCARenew *bool `json:"SSHCARenew,omitempty" msgpack:"SSHCARenew,omitempty" bson:"-" mapstructure:"SSHCARenew,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseLocalAuthority returns a new  SparseLocalAuthority.
func NewSparseLocalAuthority() *SparseLocalAuthority {
	return &SparseLocalAuthority{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseLocalAuthority) Identity() elemental.Identity {

	return LocalAuthorityIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseLocalAuthority) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseLocalAuthority) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLocalAuthority) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseLocalAuthority{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLocalAuthority) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseLocalAuthority{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseLocalAuthority) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseLocalAuthority) ToPlain() elemental.PlainIdentifiable {

	out := NewLocalAuthority()
	if o.CA != nil {
		out.CA = *o.CA
	}
	if o.CARenew != nil {
		out.CARenew = *o.CARenew
	}
	if o.SSHCA != nil {
		out.SSHCA = *o.SSHCA
	}
	if o.SSHCARenew != nil {
		out.SSHCARenew = *o.SSHCARenew
	}

	return out
}

// DeepCopy returns a deep copy if the SparseLocalAuthority.
func (o *SparseLocalAuthority) DeepCopy() *SparseLocalAuthority {

	if o == nil {
		return nil
	}

	out := &SparseLocalAuthority{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseLocalAuthority.
func (o *SparseLocalAuthority) DeepCopyInto(out *SparseLocalAuthority) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseLocalAuthority: %s", err))
	}

	*out = *target.(*SparseLocalAuthority)
}

type mongoAttributesLocalAuthority struct {
}
type mongoAttributesSparseLocalAuthority struct {
}
