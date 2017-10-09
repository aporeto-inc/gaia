package barretmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "time"

// CertificateUsageValue represents the possible values for attribute "usage".
type CertificateUsageValue string

const (
	// CertificateUsageClient represents the value Client.
	CertificateUsageClient CertificateUsageValue = "Client"

	// CertificateUsageServer represents the value Server.
	CertificateUsageServer CertificateUsageValue = "Server"

	// CertificateUsageServerclient represents the value ServerClient.
	CertificateUsageServerclient CertificateUsageValue = "ServerClient"
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
	// CSR contains the Certificate Signing Request as a PEM encoded string.
	CSR string `json:"CSR" bson:"-"`

	// ID contains the identifier of the certificate.
	ID string `json:"ID" bson:"-"`

	// Certificate contains the certificate data in PEM format.
	Certificate string `json:"certificate" bson:"-"`

	// ExpirationDate contains the requested expiration date.
	ExpirationDate time.Time `json:"expirationDate" bson:"-"`

	// Usage defines the requested key usage.
	Usage CertificateUsageValue `json:"usage" bson:"usage"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewCertificate returns a new *Certificate
func NewCertificate() *Certificate {

	return &Certificate{
		ModelVersion: 1,
		Usage:        "Client",
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
func (o *Certificate) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Certificate) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Certificate) Doc() string {
	return `This API allows to retrieve an client certifcate for api authentication.`
}

func (o *Certificate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Certificate) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("CSR", o.CSR); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("CSR", o.CSR); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("usage", string(o.Usage), []string{"Client", "Server", "ServerClient"}, false); err != nil {
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
	"CSR": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `CSR contains the Certificate Signing Request as a PEM encoded string.`,
		Exposed:        true,
		Format:         "free",
		Name:           "CSR",
		Required:       true,
		Type:           "string",
	},
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID contains the identifier of the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Certificate contains the certificate data in PEM format.`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificate",
		ReadOnly:       true,
		Type:           "string",
	},
	"ExpirationDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `ExpirationDate contains the requested expiration date.`,
		Exposed:        true,
		Name:           "expirationDate",
		SubType:        "string",
		Type:           "time",
	},
	"Usage": elemental.AttributeSpecification{
		AllowedChoices: []string{"Client", "Server", "ServerClient"},
		DefaultValue:   CertificateUsageValue("Client"),
		Description:    `Usage defines the requested key usage.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "usage",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
}

// CertificateLowerCaseAttributesMap represents the map of attribute for Certificate.
var CertificateLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"csr": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `CSR contains the Certificate Signing Request as a PEM encoded string.`,
		Exposed:        true,
		Format:         "free",
		Name:           "CSR",
		Required:       true,
		Type:           "string",
	},
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID contains the identifier of the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Type:           "string",
	},
	"certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Certificate contains the certificate data in PEM format.`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificate",
		ReadOnly:       true,
		Type:           "string",
	},
	"expirationdate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `ExpirationDate contains the requested expiration date.`,
		Exposed:        true,
		Name:           "expirationDate",
		SubType:        "string",
		Type:           "time",
	},
	"usage": elemental.AttributeSpecification{
		AllowedChoices: []string{"Client", "Server", "ServerClient"},
		DefaultValue:   CertificateUsageValue("Client"),
		Description:    `Usage defines the requested key usage.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "usage",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
}
