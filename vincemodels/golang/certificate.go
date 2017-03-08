package vincemodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"

// CertificateStatusValue represents the possible values for attribute "status".
type CertificateStatusValue string

const (
	// CertificateStatusRevoked represents the value Revoked.
	CertificateStatusRevoked CertificateStatusValue = "Revoked"

	// CertificateStatusValid represents the value Valid.
	CertificateStatusValid CertificateStatusValue = "Valid"
)

// CertificateIdentity represents the Identity of the object
var CertificateIdentity = elemental.Identity{
	Name:     "certificate",
	Category: "certificates",
}

// CertificatesList represents a list of Certificates
type CertificatesList []*Certificate

// ContentIdentity returns the identity of the objects in the list.
func (o CertificatesList) ContentIdentity() elemental.Identity {
	return CertificateIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o CertificatesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// Certificate represents the model of a certificate
type Certificate struct {
	// ID of the object.
	ID string `json:"ID" bson:"_id"`

	// Admin determines if the certificate must be added to the admin list.
	Admin bool `json:"admin" bson:"admin"`

	// CommonName (CN) for the user certificate
	CommonName string `json:"commonName" bson:"commonname"`

	// createdAt represents the creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime"`

	// Certificate provides a certificate for the user
	Data string `json:"data" bson:"data"`

	// Mark if the certificate is deleted.
	Deleted bool `json:"-" bson:"deleted"`

	// e-mail address of the user
	Email string `json:"email" bson:"email"`

	// CertificateExpirationDate indicates the expiration day for the certificate.
	ExpirationDate time.Time `json:"expirationDate" bson:"expirationdate"`

	// CertificateKey provides the key for the user. Only available at create or update time.
	Key string `json:"key" bson:"-"`

	// Name of the certificate.
	Name string `json:"name" bson:"name"`

	// OrganizationalUnits attribute for the generated certificates
	OrganizationalUnits []string `json:"organizationalUnits" bson:"organizationalunits"`

	// P12 contains the raw certificate and key in pkcs12 format.
	P12 string `json:"p12" bson:"-"`

	// ParentID holds the parent account ID.
	ParentID string `json:"parentID" bson:"parentid"`

	// Passphrase to use for the generated p12.
	Passphrase string `json:"passphrase" bson:"passphrase"`

	// SerialNumber of the certificate.
	SerialNumber string `json:"serialNumber" bson:"serialnumber"`

	// CertificateStatus provides the status of the certificate. Update with RENEW to get a new certificate.
	Status CertificateStatusValue `json:"status" bson:"status"`

	// UpdateTime represents the last update date of the objct.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`
}

// NewCertificate returns a new *Certificate
func NewCertificate() *Certificate {

	return &Certificate{
		ModelVersion: 1.0,
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
func (o *Certificate) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *Certificate) Version() float64 {

	return 1.0
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

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("passphrase", o.Passphrase); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("passphrase", o.Passphrase); err != nil {
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
func (Certificate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return CertificateAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (Certificate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CertificateAttributesMap
}

// CertificateAttributesMap represents the map of attribute for Certificate.
var CertificateAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID of the object.`,
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
		Description:    `createdAt represents the creation date of the object.`,
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
		Description:    `Certificate provides a certificate for the user`,
		Exposed:        true,
		Format:         "free",
		Name:           "data",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Deleted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Mark if the certificate is deleted.`,
		Name:           "deleted",
		Stored:         true,
		Type:           "boolean",
	},
	"Email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `e-mail address of the user`,
		Exposed:        true,
		Filterable:     true,
		Format:         "email",
		Name:           "email",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ExpirationDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		Description:    `CertificateKey provides the key for the user. Only available at create or update time.`,
		Exposed:        true,
		Format:         "free",
		Name:           "key",
		ReadOnly:       true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		CreationOnly:   true,
		Description:    `OrganizationalUnits attribute for the generated certificates`,
		Exposed:        true,
		Name:           "organizationalUnits",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"P12": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `P12 contains the raw certificate and key in pkcs12 format.`,
		Exposed:        true,
		Format:         "free",
		Name:           "p12",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
		CreationOnly:   true,
		Description:    `Passphrase to use for the generated p12.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "passphrase",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"SerialNumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
		DefaultValue:   CertificateStatusValue("Valid"),
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
		Description:    `UpdateTime represents the last update date of the objct.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}
