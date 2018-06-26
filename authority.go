package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// AuthorityIdentity represents the Identity of the object.
var AuthorityIdentity = elemental.Identity{
	Name:     "authority",
	Category: "authorities",
	Private:  true,
}

// AuthoritiesList represents a list of Authorities
type AuthoritiesList []*Authority

// Identity returns the identity of the objects in the list.
func (o AuthoritiesList) Identity() elemental.Identity {

	return AuthorityIdentity
}

// Copy returns a pointer to a copy the AuthoritiesList.
func (o AuthoritiesList) Copy() elemental.Identifiables {

	copy := append(AuthoritiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AuthoritiesList.
func (o AuthoritiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AuthoritiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Authority))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AuthoritiesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AuthoritiesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o AuthoritiesList) Version() int {

	return 1
}

// Authority represents the model of a authority
type Authority struct {
	// ID is the identitfier of the Authority.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// PEM encoded certificate data.
	Certificate string `json:"certificate" bson:"certificate" mapstructure:"certificate,omitempty"`

	// CommonName contains the common name of the CA.
	CommonName string `json:"commonName" bson:"commonname" mapstructure:"commonName,omitempty"`

	// Date of expiration of the authority.
	ExpirationDate time.Time `json:"expirationDate" bson:"expirationdate" mapstructure:"expirationDate,omitempty"`

	// Encrypted private key of the Authority.
	Key string `json:"-" bson:"key" mapstructure:"-,omitempty"`

	// serialNumber of the certificate.
	SerialNumber string `json:"serialNumber" bson:"serialnumber" mapstructure:"serialNumber,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewAuthority returns a new *Authority
func NewAuthority() *Authority {

	return &Authority{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Authority) Identity() elemental.Identity {

	return AuthorityIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Authority) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Authority) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Authority) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Authority) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Authority) Doc() string {
	return `Authority represents a certificate authority.`
}

func (o *Authority) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Authority) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("commonName", o.CommonName); err != nil {
		requiredErrors = append(requiredErrors, err)
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
func (*Authority) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AuthorityAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AuthorityLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Authority) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AuthorityAttributesMap
}

// AuthorityAttributesMap represents the map of attribute for Authority.
var AuthorityAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identitfier of the Authority.`,
		Exposed:        true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `PEM encoded certificate data.`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificate",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CommonName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CommonName",
		CreationOnly:   true,
		Description:    `CommonName contains the common name of the CA.`,
		Exposed:        true,
		Format:         "free",
		Name:           "commonName",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ExpirationDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description:    `Date of expiration of the authority.`,
		Exposed:        true,
		Name:           "expirationDate",
		Stored:         true,
		Type:           "time",
	},
	"Key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `Encrypted private key of the Authority.`,
		Format:         "free",
		Name:           "key",
		Stored:         true,
		Type:           "string",
	},
	"SerialNumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SerialNumber",
		Description:    `serialNumber of the certificate.`,
		Exposed:        true,
		Format:         "free",
		Name:           "serialNumber",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}

// AuthorityLowerCaseAttributesMap represents the map of attribute for Authority.
var AuthorityLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identitfier of the Authority.`,
		Exposed:        true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `PEM encoded certificate data.`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificate",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"commonname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CommonName",
		CreationOnly:   true,
		Description:    `CommonName contains the common name of the CA.`,
		Exposed:        true,
		Format:         "free",
		Name:           "commonName",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"expirationdate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description:    `Date of expiration of the authority.`,
		Exposed:        true,
		Name:           "expirationDate",
		Stored:         true,
		Type:           "time",
	},
	"key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `Encrypted private key of the Authority.`,
		Format:         "free",
		Name:           "key",
		Stored:         true,
		Type:           "string",
	},
	"serialnumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SerialNumber",
		Description:    `serialNumber of the certificate.`,
		Exposed:        true,
		Format:         "free",
		Name:           "serialNumber",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}
