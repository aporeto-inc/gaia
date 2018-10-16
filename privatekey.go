package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
)

// PrivateKeyIdentity represents the Identity of the object.
var PrivateKeyIdentity = elemental.Identity{
	Name:     "privatekey",
	Category: "privatekeys",
	Package:  "barret",
	Private:  true,
}

// PrivateKeysList represents a list of PrivateKeys
type PrivateKeysList []*PrivateKey

// Identity returns the identity of the objects in the list.
func (o PrivateKeysList) Identity() elemental.Identity {

	return PrivateKeyIdentity
}

// Copy returns a pointer to a copy the PrivateKeysList.
func (o PrivateKeysList) Copy() elemental.Identifiables {

	copy := append(PrivateKeysList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PrivateKeysList.
func (o PrivateKeysList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PrivateKeysList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PrivateKey))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PrivateKeysList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PrivateKeysList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PrivateKeysList converted to SparsePrivateKeysList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PrivateKeysList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o PrivateKeysList) Version() int {

	return 1
}

// PrivateKey represents the model of a privatekey
type PrivateKey struct {
	// ID is the internal ID of the key.
	ID string `json:"-" bson:"_id" mapstructure:"-,omitempty"`

	// CertificateSerialNumber represents the certificate serial number associated to
	// this key.
	CertificateSerialNumber string `json:"-" bson:"certificateserialnumber" mapstructure:"-,omitempty"`

	// Data contains the privateKey data.
	Data string `json:"-" bson:"data" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewPrivateKey returns a new *PrivateKey
func NewPrivateKey() *PrivateKey {

	return &PrivateKey{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PrivateKey) Identity() elemental.Identity {

	return PrivateKeyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PrivateKey) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PrivateKey) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *PrivateKey) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *PrivateKey) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PrivateKey) Doc() string {
	return `Internal representation of an private key.`
}

func (o *PrivateKey) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PrivateKey) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePrivateKey{
			ID:                      &o.ID,
			CertificateSerialNumber: &o.CertificateSerialNumber,
			Data:                    &o.Data,
		}
	}

	sp := &SparsePrivateKey{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "certificateSerialNumber":
			sp.CertificateSerialNumber = &(o.CertificateSerialNumber)
		case "data":
			sp.Data = &(o.Data)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePrivateKey to the object.
func (o *PrivateKey) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePrivateKey)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.CertificateSerialNumber != nil {
		o.CertificateSerialNumber = *so.CertificateSerialNumber
	}
	if so.Data != nil {
		o.Data = *so.Data
	}
}

// Validate valides the current information stored into the structure.
func (o *PrivateKey) Validate() error {

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
func (*PrivateKey) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PrivateKeyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PrivateKeyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PrivateKey) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PrivateKeyAttributesMap
}

// PrivateKeyAttributesMap represents the map of attribute for PrivateKey.
var PrivateKeyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the internal ID of the key.`,
		Identifier:     true,
		Name:           "ID",
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
	"CertificateSerialNumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateSerialNumber",
		Description: `CertificateSerialNumber represents the certificate serial number associated to
this key.`,
		Name:   "certificateSerialNumber",
		Stored: true,
		Type:   "string",
	},
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		CreationOnly:   true,
		Description:    `Data contains the privateKey data.`,
		Name:           "data",
		Stored:         true,
		Type:           "string",
	},
}

// PrivateKeyLowerCaseAttributesMap represents the map of attribute for PrivateKey.
var PrivateKeyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the internal ID of the key.`,
		Identifier:     true,
		Name:           "ID",
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
	"certificateserialnumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateSerialNumber",
		Description: `CertificateSerialNumber represents the certificate serial number associated to
this key.`,
		Name:   "certificateSerialNumber",
		Stored: true,
		Type:   "string",
	},
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		CreationOnly:   true,
		Description:    `Data contains the privateKey data.`,
		Name:           "data",
		Stored:         true,
		Type:           "string",
	},
}

// SparsePrivateKeysList represents a list of SparsePrivateKeys
type SparsePrivateKeysList []*SparsePrivateKey

// Identity returns the identity of the objects in the list.
func (o SparsePrivateKeysList) Identity() elemental.Identity {

	return PrivateKeyIdentity
}

// Copy returns a pointer to a copy the SparsePrivateKeysList.
func (o SparsePrivateKeysList) Copy() elemental.Identifiables {

	copy := append(SparsePrivateKeysList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePrivateKeysList.
func (o SparsePrivateKeysList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePrivateKeysList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePrivateKey))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePrivateKeysList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePrivateKeysList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePrivateKeysList converted to PrivateKeysList.
func (o SparsePrivateKeysList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePrivateKeysList) Version() int {

	return 1
}

// SparsePrivateKey represents the sparse version of a privatekey.
type SparsePrivateKey struct {
	// ID is the internal ID of the key.
	ID *string `json:"-,omitempty" bson:"_id" mapstructure:"-,omitempty"`

	// CertificateSerialNumber represents the certificate serial number associated to
	// this key.
	CertificateSerialNumber *string `json:"-,omitempty" bson:"certificateserialnumber" mapstructure:"-,omitempty"`

	// Data contains the privateKey data.
	Data *string `json:"-,omitempty" bson:"data" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparsePrivateKey returns a new  SparsePrivateKey.
func NewSparsePrivateKey() *SparsePrivateKey {
	return &SparsePrivateKey{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePrivateKey) Identity() elemental.Identity {

	return PrivateKeyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePrivateKey) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePrivateKey) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparsePrivateKey) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePrivateKey) ToPlain() elemental.PlainIdentifiable {

	out := NewPrivateKey()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.CertificateSerialNumber != nil {
		out.CertificateSerialNumber = *o.CertificateSerialNumber
	}
	if o.Data != nil {
		out.Data = *o.Data
	}

	return out
}
