package models

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// PrivateKeyIdentity represents the Identity of the object.
var PrivateKeyIdentity = elemental.Identity{
	Name:     "privatekey",
	Category: "privatekeys",
}

// PrivateKeysList represents a list of PrivateKeys
type PrivateKeysList []*PrivateKey

// ContentIdentity returns the identity of the objects in the list.
func (o PrivateKeysList) ContentIdentity() elemental.Identity {

	return PrivateKeyIdentity
}

// Copy returns a pointer to a copy the PrivateKeysList.
func (o PrivateKeysList) Copy() elemental.ContentIdentifiable {

	copy := append(PrivateKeysList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PrivateKeysList.
func (o PrivateKeysList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(PrivateKeysList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PrivateKey))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PrivateKeysList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PrivateKeysList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o PrivateKeysList) Version() int {

	return 1
}

// PrivateKey represents the model of a privatekey
type PrivateKey struct {
	// ID is the internal ID of the key.
	ID string `json:"-" bson:"_id" mapstructure:"-,omitempty"`

	// CertificateSerialNumber represents the certificate serial number associated to this key.
	CertificateSerialNumber string `json:"-" bson:"certificateserialnumber" mapstructure:"-,omitempty"`

	// Data contains the privateKey data.
	Data string `json:"-" bson:"data" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
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

func (o *PrivateKey) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
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
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
	"CertificateSerialNumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateSerialNumber",
		Description:    `CertificateSerialNumber represents the certificate serial number associated to this key.`,
		Format:         "free",
		Name:           "certificateSerialNumber",
		Stored:         true,
		Type:           "string",
	},
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		CreationOnly:   true,
		Description:    `Data contains the privateKey data.`,
		Format:         "free",
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
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
	"certificateserialnumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateSerialNumber",
		Description:    `CertificateSerialNumber represents the certificate serial number associated to this key.`,
		Format:         "free",
		Name:           "certificateSerialNumber",
		Stored:         true,
		Type:           "string",
	},
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		CreationOnly:   true,
		Description:    `Data contains the privateKey data.`,
		Format:         "free",
		Name:           "data",
		Stored:         true,
		Type:           "string",
	},
}
