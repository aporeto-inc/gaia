package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// LocalCAsIdentity represents the Identity of the object.
var LocalCAsIdentity = elemental.Identity{
	Name:     "localca",
	Category: "localcas",
	Package:  "squall",
	Private:  false,
}

// LocalCAsList represents a list of LocalCAs
type LocalCAsList []*LocalCAs

// Identity returns the identity of the objects in the list.
func (o LocalCAsList) Identity() elemental.Identity {

	return LocalCAsIdentity
}

// Copy returns a pointer to a copy the LocalCAsList.
func (o LocalCAsList) Copy() elemental.Identifiables {

	copy := append(LocalCAsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the LocalCAsList.
func (o LocalCAsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(LocalCAsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*LocalCAs))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o LocalCAsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o LocalCAsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the LocalCAsList converted to SparseLocalCAsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o LocalCAsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseLocalCAsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseLocalCAs)
	}

	return out
}

// Version returns the version of the content.
func (o LocalCAsList) Version() int {

	return 1
}

// LocalCAs represents the model of a localca
type LocalCAs struct {
	// The certificate authority used by the namespace.
	Certificate string `json:"Certificate" msgpack:"Certificate" bson:"-" mapstructure:"Certificate,omitempty"`

	// Set to `true` to renew the certificate authority of the namespace.
	CertificateRenew bool `json:"CertificateRenew" msgpack:"CertificateRenew" bson:"-" mapstructure:"CertificateRenew,omitempty"`

	// The SSH certificate authority used by the namespace.
	SSHCertificate string `json:"SSHCertificate" msgpack:"SSHCertificate" bson:"-" mapstructure:"SSHCertificate,omitempty"`

	// Set to `true` to renew the SSH certificate authority of the namespace.
	SSHCertificateRenew bool `json:"SSHCertificateRenew" msgpack:"SSHCertificateRenew" bson:"-" mapstructure:"SSHCertificateRenew,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewLocalCAs returns a new *LocalCAs
func NewLocalCAs() *LocalCAs {

	return &LocalCAs{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *LocalCAs) Identity() elemental.Identity {

	return LocalCAsIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *LocalCAs) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *LocalCAs) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *LocalCAs) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesLocalCAs{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *LocalCAs) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesLocalCAs{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *LocalCAs) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *LocalCAs) BleveType() string {

	return "localca"
}

// DefaultOrder returns the list of default ordering fields.
func (o *LocalCAs) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *LocalCAs) Doc() string {

	return `Can be used to retrieve or renew the local and SSH certificate authorities of
the namespace.`
}

func (o *LocalCAs) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *LocalCAs) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseLocalCAs{
			Certificate:         &o.Certificate,
			CertificateRenew:    &o.CertificateRenew,
			SSHCertificate:      &o.SSHCertificate,
			SSHCertificateRenew: &o.SSHCertificateRenew,
		}
	}

	sp := &SparseLocalCAs{}
	for _, f := range fields {
		switch f {
		case "Certificate":
			sp.Certificate = &(o.Certificate)
		case "CertificateRenew":
			sp.CertificateRenew = &(o.CertificateRenew)
		case "SSHCertificate":
			sp.SSHCertificate = &(o.SSHCertificate)
		case "SSHCertificateRenew":
			sp.SSHCertificateRenew = &(o.SSHCertificateRenew)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseLocalCAs to the object.
func (o *LocalCAs) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseLocalCAs)
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.CertificateRenew != nil {
		o.CertificateRenew = *so.CertificateRenew
	}
	if so.SSHCertificate != nil {
		o.SSHCertificate = *so.SSHCertificate
	}
	if so.SSHCertificateRenew != nil {
		o.SSHCertificateRenew = *so.SSHCertificateRenew
	}
}

// DeepCopy returns a deep copy if the LocalCAs.
func (o *LocalCAs) DeepCopy() *LocalCAs {

	if o == nil {
		return nil
	}

	out := &LocalCAs{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *LocalCAs.
func (o *LocalCAs) DeepCopyInto(out *LocalCAs) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy LocalCAs: %s", err))
	}

	*out = *target.(*LocalCAs)
}

// Validate valides the current information stored into the structure.
func (o *LocalCAs) Validate() error {

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
func (*LocalCAs) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := LocalCAsAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return LocalCAsLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*LocalCAs) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return LocalCAsAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *LocalCAs) ValueForAttribute(name string) interface{} {

	switch name {
	case "Certificate":
		return o.Certificate
	case "CertificateRenew":
		return o.CertificateRenew
	case "SSHCertificate":
		return o.SSHCertificate
	case "SSHCertificateRenew":
		return o.SSHCertificateRenew
	}

	return nil
}

// LocalCAsAttributesMap represents the map of attribute for LocalCAs.
var LocalCAsAttributesMap = map[string]elemental.AttributeSpecification{
	"Certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `The certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "Certificate",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"CertificateRenew": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CertificateRenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "CertificateRenew",
		ReadOnly:       true,
		Transient:      true,
		Type:           "boolean",
	},
	"SSHCertificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCertificate",
		Description:    `The SSH certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "SSHCertificate",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"SSHCertificateRenew": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCertificateRenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the SSH certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "SSHCertificateRenew",
		ReadOnly:       true,
		Transient:      true,
		Type:           "boolean",
	},
}

// LocalCAsLowerCaseAttributesMap represents the map of attribute for LocalCAs.
var LocalCAsLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `The certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "Certificate",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"certificaterenew": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CertificateRenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "CertificateRenew",
		ReadOnly:       true,
		Transient:      true,
		Type:           "boolean",
	},
	"sshcertificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCertificate",
		Description:    `The SSH certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "SSHCertificate",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"sshcertificaterenew": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCertificateRenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the SSH certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "SSHCertificateRenew",
		ReadOnly:       true,
		Transient:      true,
		Type:           "boolean",
	},
}

// SparseLocalCAsList represents a list of SparseLocalCAs
type SparseLocalCAsList []*SparseLocalCAs

// Identity returns the identity of the objects in the list.
func (o SparseLocalCAsList) Identity() elemental.Identity {

	return LocalCAsIdentity
}

// Copy returns a pointer to a copy the SparseLocalCAsList.
func (o SparseLocalCAsList) Copy() elemental.Identifiables {

	copy := append(SparseLocalCAsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseLocalCAsList.
func (o SparseLocalCAsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseLocalCAsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseLocalCAs))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseLocalCAsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseLocalCAsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseLocalCAsList converted to LocalCAsList.
func (o SparseLocalCAsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseLocalCAsList) Version() int {

	return 1
}

// SparseLocalCAs represents the sparse version of a localca.
type SparseLocalCAs struct {
	// The certificate authority used by the namespace.
	Certificate *string `json:"Certificate,omitempty" msgpack:"Certificate,omitempty" bson:"-" mapstructure:"Certificate,omitempty"`

	// Set to `true` to renew the certificate authority of the namespace.
	CertificateRenew *bool `json:"CertificateRenew,omitempty" msgpack:"CertificateRenew,omitempty" bson:"-" mapstructure:"CertificateRenew,omitempty"`

	// The SSH certificate authority used by the namespace.
	SSHCertificate *string `json:"SSHCertificate,omitempty" msgpack:"SSHCertificate,omitempty" bson:"-" mapstructure:"SSHCertificate,omitempty"`

	// Set to `true` to renew the SSH certificate authority of the namespace.
	SSHCertificateRenew *bool `json:"SSHCertificateRenew,omitempty" msgpack:"SSHCertificateRenew,omitempty" bson:"-" mapstructure:"SSHCertificateRenew,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseLocalCAs returns a new  SparseLocalCAs.
func NewSparseLocalCAs() *SparseLocalCAs {
	return &SparseLocalCAs{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseLocalCAs) Identity() elemental.Identity {

	return LocalCAsIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseLocalCAs) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseLocalCAs) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLocalCAs) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseLocalCAs{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLocalCAs) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseLocalCAs{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseLocalCAs) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseLocalCAs) ToPlain() elemental.PlainIdentifiable {

	out := NewLocalCAs()
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.CertificateRenew != nil {
		out.CertificateRenew = *o.CertificateRenew
	}
	if o.SSHCertificate != nil {
		out.SSHCertificate = *o.SSHCertificate
	}
	if o.SSHCertificateRenew != nil {
		out.SSHCertificateRenew = *o.SSHCertificateRenew
	}

	return out
}

// DeepCopy returns a deep copy if the SparseLocalCAs.
func (o *SparseLocalCAs) DeepCopy() *SparseLocalCAs {

	if o == nil {
		return nil
	}

	out := &SparseLocalCAs{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseLocalCAs.
func (o *SparseLocalCAs) DeepCopyInto(out *SparseLocalCAs) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseLocalCAs: %s", err))
	}

	*out = *target.(*SparseLocalCAs)
}

type mongoAttributesLocalCAs struct {
}
type mongoAttributesSparseLocalCAs struct {
}
