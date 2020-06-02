package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CertificatesIdentity represents the Identity of the object.
var CertificatesIdentity = elemental.Identity{
	Name:     "certificates",
	Category: "certificates",
	Package:  "squall",
	Private:  false,
}

// CertificatesList represents a list of Certificates
type CertificatesList []*Certificates

// Identity returns the identity of the objects in the list.
func (o CertificatesList) Identity() elemental.Identity {

	return CertificatesIdentity
}

// Copy returns a pointer to a copy the CertificatesList.
func (o CertificatesList) Copy() elemental.Identifiables {

	copy := append(CertificatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CertificatesList.
func (o CertificatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CertificatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Certificates))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CertificatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CertificatesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CertificatesList converted to SparseCertificatesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CertificatesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCertificatesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCertificates)
	}

	return out
}

// Version returns the version of the content.
func (o CertificatesList) Version() int {

	return 1
}

// Certificates represents the model of a certificates
type Certificates struct {
	// The certificate authority used by the namespace.
	LocalCA string `json:"LocalCA" msgpack:"LocalCA" bson:"-" mapstructure:"LocalCA,omitempty"`

	// Set to `true` to renew the local certificate authority of the namespace.
	LocalCARenew bool `json:"LocalCARenew" msgpack:"LocalCARenew" bson:"-" mapstructure:"LocalCARenew,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCertificates returns a new *Certificates
func NewCertificates() *Certificates {

	return &Certificates{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Certificates) Identity() elemental.Identity {

	return CertificatesIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Certificates) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Certificates) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Certificates) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCertificates{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Certificates) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCertificates{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Certificates) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Certificates) BleveType() string {

	return "certificates"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Certificates) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Certificates) Doc() string {

	return `Can be used to retrieve or renew the local certificate authority of the
namespace.`
}

func (o *Certificates) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Certificates) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCertificates{
			LocalCA:      &o.LocalCA,
			LocalCARenew: &o.LocalCARenew,
		}
	}

	sp := &SparseCertificates{}
	for _, f := range fields {
		switch f {
		case "LocalCA":
			sp.LocalCA = &(o.LocalCA)
		case "LocalCARenew":
			sp.LocalCARenew = &(o.LocalCARenew)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCertificates to the object.
func (o *Certificates) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCertificates)
	if so.LocalCA != nil {
		o.LocalCA = *so.LocalCA
	}
	if so.LocalCARenew != nil {
		o.LocalCARenew = *so.LocalCARenew
	}
}

// DeepCopy returns a deep copy if the Certificates.
func (o *Certificates) DeepCopy() *Certificates {

	if o == nil {
		return nil
	}

	out := &Certificates{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Certificates.
func (o *Certificates) DeepCopyInto(out *Certificates) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Certificates: %s", err))
	}

	*out = *target.(*Certificates)
}

// Validate valides the current information stored into the structure.
func (o *Certificates) Validate() error {

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
func (*Certificates) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CertificatesAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CertificatesLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Certificates) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CertificatesAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Certificates) ValueForAttribute(name string) interface{} {

	switch name {
	case "LocalCA":
		return o.LocalCA
	case "LocalCARenew":
		return o.LocalCARenew
	}

	return nil
}

// CertificatesAttributesMap represents the map of attribute for Certificates.
var CertificatesAttributesMap = map[string]elemental.AttributeSpecification{
	"LocalCA": {
		AllowedChoices: []string{},
		ConvertedName:  "LocalCA",
		Description:    `The certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "LocalCA",
		Transient:      true,
		Type:           "string",
	},
	"LocalCARenew": {
		AllowedChoices: []string{},
		ConvertedName:  "LocalCARenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the local certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "LocalCARenew",
		Transient:      true,
		Type:           "boolean",
	},
}

// CertificatesLowerCaseAttributesMap represents the map of attribute for Certificates.
var CertificatesLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"localca": {
		AllowedChoices: []string{},
		ConvertedName:  "LocalCA",
		Description:    `The certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "LocalCA",
		Transient:      true,
		Type:           "string",
	},
	"localcarenew": {
		AllowedChoices: []string{},
		ConvertedName:  "LocalCARenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the local certificate authority of the namespace.`,
		Exposed:        true,
		Name:           "LocalCARenew",
		Transient:      true,
		Type:           "boolean",
	},
}

// SparseCertificatesList represents a list of SparseCertificates
type SparseCertificatesList []*SparseCertificates

// Identity returns the identity of the objects in the list.
func (o SparseCertificatesList) Identity() elemental.Identity {

	return CertificatesIdentity
}

// Copy returns a pointer to a copy the SparseCertificatesList.
func (o SparseCertificatesList) Copy() elemental.Identifiables {

	copy := append(SparseCertificatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCertificatesList.
func (o SparseCertificatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCertificatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCertificates))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCertificatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCertificatesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCertificatesList converted to CertificatesList.
func (o SparseCertificatesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCertificatesList) Version() int {

	return 1
}

// SparseCertificates represents the sparse version of a certificates.
type SparseCertificates struct {
	// The certificate authority used by the namespace.
	LocalCA *string `json:"LocalCA,omitempty" msgpack:"LocalCA,omitempty" bson:"-" mapstructure:"LocalCA,omitempty"`

	// Set to `true` to renew the local certificate authority of the namespace.
	LocalCARenew *bool `json:"LocalCARenew,omitempty" msgpack:"LocalCARenew,omitempty" bson:"-" mapstructure:"LocalCARenew,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCertificates returns a new  SparseCertificates.
func NewSparseCertificates() *SparseCertificates {
	return &SparseCertificates{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCertificates) Identity() elemental.Identity {

	return CertificatesIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCertificates) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCertificates) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCertificates) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCertificates{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCertificates) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCertificates{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCertificates) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCertificates) ToPlain() elemental.PlainIdentifiable {

	out := NewCertificates()
	if o.LocalCA != nil {
		out.LocalCA = *o.LocalCA
	}
	if o.LocalCARenew != nil {
		out.LocalCARenew = *o.LocalCARenew
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCertificates.
func (o *SparseCertificates) DeepCopy() *SparseCertificates {

	if o == nil {
		return nil
	}

	out := &SparseCertificates{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCertificates.
func (o *SparseCertificates) DeepCopyInto(out *SparseCertificates) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCertificates: %s", err))
	}

	*out = *target.(*SparseCertificates)
}

type mongoAttributesCertificates struct {
}
type mongoAttributesSparseCertificates struct {
}
