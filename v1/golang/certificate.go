package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
)

// CertificateStatusValue represents the possible values for attribute "status".
type CertificateStatusValue string

const (
	// CertificateStatusRevoked represents the value Revoked.
	CertificateStatusRevoked CertificateStatusValue = "Revoked"

	// CertificateStatusValid represents the value Valid.
	CertificateStatusValid CertificateStatusValue = "Valid"
)

// CertificateIdentity represents the Identity of the object.
var CertificateIdentity = elemental.Identity{
	Name:     "certificate",
	Category: "certificates",
	Private:  false,
}

// CertificatesList represents a list of Certificates
type CertificatesList []*Certificate

// ContentIdentity returns the identity of the objects in the list.
func (o CertificatesList) ContentIdentity() elemental.Identity {

	return CertificateIdentity
}

// Copy returns a pointer to a copy the CertificatesList.
func (o CertificatesList) Copy() elemental.ContentIdentifiable {

	copy := append(CertificatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CertificatesList.
func (o CertificatesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(CertificatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Certificate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CertificatesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CertificatesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o CertificatesList) Version() int {

	return 1
}

// Certificate represents the model of a certificate
type Certificate struct {
	// Admin determines if the certificate must be added to the admin list.
	Admin bool `json:"admin" bson:"admin" mapstructure:"admin,omitempty"`

	// CommonName (CN) for the user certificate
	CommonName string `json:"commonName" bson:"commonname" mapstructure:"commonName,omitempty"`

	// Certificate provides a certificate for the user
	Data string `json:"data" bson:"data" mapstructure:"data,omitempty"`

	// e-mail address of the user
	Email string `json:"email" bson:"email" mapstructure:"email,omitempty"`

	// CertificateExpirationDate indicates the expiration day for the certificate.
	ExpirationDate time.Time `json:"expirationDate" bson:"expirationdate" mapstructure:"expirationDate,omitempty"`

	// CertificateKey provides the key for the user. Only available at create or update time.
	Key string `json:"key" bson:"-" mapstructure:"key,omitempty"`

	// Name of the certificate.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// OrganizationalUnits attribute for the generated certificates
	OrganizationalUnits []string `json:"organizationalUnits" bson:"organizationalunits" mapstructure:"organizationalUnits,omitempty"`

	// ParentID holds the parent account ID.
	ParentID string `json:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// Passphrase to use for the generated p12.
	Passphrase string `json:"passphrase" bson:"-" mapstructure:"passphrase,omitempty"`

	// SerialNumber of the certificate.
	SerialNumber string `json:"serialNumber" bson:"serialnumber" mapstructure:"serialNumber,omitempty"`

	// CertificateStatus provides the status of the certificate. Update with RENEW to get a new certificate.
	Status CertificateStatusValue `json:"status" bson:"status" mapstructure:"status,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Last update date of the object
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewCertificate returns a new *Certificate
func NewCertificate() *Certificate {

	return &Certificate{
		ModelVersion: 1,
		Status:       "Valid",
	}
}

// Identity returns the Identity of the object.
func (o *Certificate) Identity() elemental.Identity {

	return CertificateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Certificate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Certificate) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Certificate) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Certificate) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Certificate) Doc() string {
	return `A User represents the owner of some certificates.`
}

func (o *Certificate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Certificate) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("commonName", o.CommonName); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("commonName", o.CommonName, 64, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("commonName", o.CommonName); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("email", o.Email); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("email", o.Email); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Revoked", "Valid"}, false); err != nil {
		errors = append(errors, err)
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
func (*Certificate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CertificateAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CertificateLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Certificate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CertificateAttributesMap
}

// CertificateAttributesMap represents the map of attribute for Certificate.
var CertificateAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Admin": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Admin",
		Description:    `Admin determines if the certificate must be added to the admin list.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "admin",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"CommonName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CommonName",
		CreationOnly:   true,
		Description:    `CommonName (CN) for the user certificate`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		MaxLength:      64,
		Name:           "commonName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `Certificate provides a certificate for the user`,
		Exposed:        true,
		Format:         "free",
		Name:           "data",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		CreationOnly:   true,
		Description:    `e-mail address of the user`,
		Exposed:        true,
		Filterable:     true,
		Format:         "email",
		Name:           "email",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ExpirationDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description:    `CertificateExpirationDate indicates the expiration day for the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "expirationDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `CertificateKey provides the key for the user. Only available at create or update time.`,
		Exposed:        true,
		Format:         "free",
		Name:           "key",
		ReadOnly:       true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"OrganizationalUnits": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OrganizationalUnits",
		CreationOnly:   true,
		Description:    `OrganizationalUnits attribute for the generated certificates`,
		Exposed:        true,
		Name:           "organizationalUnits",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `ParentID holds the parent account ID.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Passphrase": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Passphrase",
		CreationOnly:   true,
		Description:    `Passphrase to use for the generated p12.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "passphrase",
		Orderable:      true,
		Type:           "string",
	},
	"SerialNumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SerialNumber",
		Description:    `SerialNumber of the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "serialNumber",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Revoked", "Valid"},
		ConvertedName:  "Status",
		DefaultValue:   CertificateStatusValid,
		Description:    `CertificateStatus provides the status of the certificate. Update with RENEW to get a new certificate.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}

// CertificateLowerCaseAttributesMap represents the map of attribute for Certificate.
var CertificateLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"admin": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Admin",
		Description:    `Admin determines if the certificate must be added to the admin list.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "admin",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"commonname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CommonName",
		CreationOnly:   true,
		Description:    `CommonName (CN) for the user certificate`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		MaxLength:      64,
		Name:           "commonName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `Certificate provides a certificate for the user`,
		Exposed:        true,
		Format:         "free",
		Name:           "data",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		CreationOnly:   true,
		Description:    `e-mail address of the user`,
		Exposed:        true,
		Filterable:     true,
		Format:         "email",
		Name:           "email",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"expirationdate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description:    `CertificateExpirationDate indicates the expiration day for the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "expirationDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `CertificateKey provides the key for the user. Only available at create or update time.`,
		Exposed:        true,
		Format:         "free",
		Name:           "key",
		ReadOnly:       true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"organizationalunits": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OrganizationalUnits",
		CreationOnly:   true,
		Description:    `OrganizationalUnits attribute for the generated certificates`,
		Exposed:        true,
		Name:           "organizationalUnits",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"parentid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `ParentID holds the parent account ID.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"passphrase": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Passphrase",
		CreationOnly:   true,
		Description:    `Passphrase to use for the generated p12.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "passphrase",
		Orderable:      true,
		Type:           "string",
	},
	"serialnumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SerialNumber",
		Description:    `SerialNumber of the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "serialNumber",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Revoked", "Valid"},
		ConvertedName:  "Status",
		DefaultValue:   CertificateStatusValid,
		Description:    `CertificateStatus provides the status of the certificate. Update with RENEW to get a new certificate.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}
