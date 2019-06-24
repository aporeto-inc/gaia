package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// X509CertificateSignerValue represents the possible values for attribute "signer".
type X509CertificateSignerValue string

const (
	// X509CertificateSignerPublic represents the value Public.
	X509CertificateSignerPublic X509CertificateSignerValue = "Public"

	// X509CertificateSignerSystem represents the value System.
	X509CertificateSignerSystem X509CertificateSignerValue = "System"
)

// X509CertificateUsageValue represents the possible values for attribute "usage".
type X509CertificateUsageValue string

const (
	// X509CertificateUsageClient represents the value Client.
	X509CertificateUsageClient X509CertificateUsageValue = "Client"

	// X509CertificateUsageServer represents the value Server.
	X509CertificateUsageServer X509CertificateUsageValue = "Server"

	// X509CertificateUsageServerClient represents the value ServerClient.
	X509CertificateUsageServerClient X509CertificateUsageValue = "ServerClient"
)

// X509CertificateIdentity represents the Identity of the object.
var X509CertificateIdentity = elemental.Identity{
	Name:     "x509certificate",
	Category: "x509certificates",
	Package:  "barret",
	Private:  true,
}

// X509CertificatesList represents a list of X509Certificates
type X509CertificatesList []*X509Certificate

// Identity returns the identity of the objects in the list.
func (o X509CertificatesList) Identity() elemental.Identity {

	return X509CertificateIdentity
}

// Copy returns a pointer to a copy the X509CertificatesList.
func (o X509CertificatesList) Copy() elemental.Identifiables {

	copy := append(X509CertificatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the X509CertificatesList.
func (o X509CertificatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(X509CertificatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*X509Certificate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o X509CertificatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o X509CertificatesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the X509CertificatesList converted to SparseX509CertificatesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o X509CertificatesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseX509CertificatesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseX509Certificate)
	}

	return out
}

// Version returns the version of the content.
func (o X509CertificatesList) Version() int {

	return 1
}

// X509Certificate represents the model of a x509certificate
type X509Certificate struct {
	// CSR contains the Certificate Signing Request as a PEM encoded string.
	CSR string `json:"CSR" msgpack:"CSR" bson:"-" mapstructure:"CSR,omitempty"`

	// ID contains the identifier of the certificate.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Certificate contains the certificate data in PEM format.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// ExpirationDate contains the requested expiration date.
	ExpirationDate time.Time `json:"expirationDate" msgpack:"expirationDate" bson:"-" mapstructure:"expirationDate,omitempty"`

	// Extensions is a list of extensions that can be added as SAN extensions to the
	// certificate.
	Extensions []string `json:"extensions" msgpack:"extensions" bson:"-" mapstructure:"extensions,omitempty"`

	// Selects what CA should sign the certificate.
	Signer X509CertificateSignerValue `json:"signer" msgpack:"signer" bson:"-" mapstructure:"signer,omitempty"`

	// Additional subject information to use to override the ones in the CSR.
	SubjectOverride *PKIXName `json:"subjectOverride,omitempty" msgpack:"subjectOverride,omitempty" bson:"-" mapstructure:"subjectOverride,omitempty"`

	// If set to true, the certificate is considered short lived and it will not be
	// possible to revoke it.
	Unrevocable bool `json:"unrevocable" msgpack:"unrevocable" bson:"-" mapstructure:"unrevocable,omitempty"`

	// Usage defines the requested key usage.
	Usage X509CertificateUsageValue `json:"usage" msgpack:"usage" bson:"-" mapstructure:"usage,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewX509Certificate returns a new *X509Certificate
func NewX509Certificate() *X509Certificate {

	return &X509Certificate{
		ModelVersion:    1,
		Extensions:      []string{},
		Signer:          X509CertificateSignerPublic,
		SubjectOverride: NewPKIXName(),
		Usage:           X509CertificateUsageClient,
	}
}

// Identity returns the Identity of the object.
func (o *X509Certificate) Identity() elemental.Identity {

	return X509CertificateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *X509Certificate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *X509Certificate) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *X509Certificate) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *X509Certificate) BleveType() string {

	return "x509certificate"
}

// DefaultOrder returns the list of default ordering fields.
func (o *X509Certificate) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *X509Certificate) Doc() string {

	return `This API allows to retrieve an client certifcate for api authentication.`
}

func (o *X509Certificate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *X509Certificate) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseX509Certificate{
			CSR:             &o.CSR,
			ID:              &o.ID,
			Certificate:     &o.Certificate,
			ExpirationDate:  &o.ExpirationDate,
			Extensions:      &o.Extensions,
			Signer:          &o.Signer,
			SubjectOverride: o.SubjectOverride,
			Unrevocable:     &o.Unrevocable,
			Usage:           &o.Usage,
		}
	}

	sp := &SparseX509Certificate{}
	for _, f := range fields {
		switch f {
		case "CSR":
			sp.CSR = &(o.CSR)
		case "ID":
			sp.ID = &(o.ID)
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "expirationDate":
			sp.ExpirationDate = &(o.ExpirationDate)
		case "extensions":
			sp.Extensions = &(o.Extensions)
		case "signer":
			sp.Signer = &(o.Signer)
		case "subjectOverride":
			sp.SubjectOverride = o.SubjectOverride
		case "unrevocable":
			sp.Unrevocable = &(o.Unrevocable)
		case "usage":
			sp.Usage = &(o.Usage)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseX509Certificate to the object.
func (o *X509Certificate) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseX509Certificate)
	if so.CSR != nil {
		o.CSR = *so.CSR
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.ExpirationDate != nil {
		o.ExpirationDate = *so.ExpirationDate
	}
	if so.Extensions != nil {
		o.Extensions = *so.Extensions
	}
	if so.Signer != nil {
		o.Signer = *so.Signer
	}
	if so.SubjectOverride != nil {
		o.SubjectOverride = so.SubjectOverride
	}
	if so.Unrevocable != nil {
		o.Unrevocable = *so.Unrevocable
	}
	if so.Usage != nil {
		o.Usage = *so.Usage
	}
}

// DeepCopy returns a deep copy if the X509Certificate.
func (o *X509Certificate) DeepCopy() *X509Certificate {

	if o == nil {
		return nil
	}

	out := &X509Certificate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *X509Certificate.
func (o *X509Certificate) DeepCopyInto(out *X509Certificate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy X509Certificate: %s", err))
	}

	*out = *target.(*X509Certificate)
}

// Validate valides the current information stored into the structure.
func (o *X509Certificate) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("CSR", o.CSR); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("signer", string(o.Signer), []string{"Public", "System"}, false); err != nil {
		errors = errors.Append(err)
	}

	if o.SubjectOverride != nil {
		if err := o.SubjectOverride.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateStringInList("usage", string(o.Usage), []string{"Client", "Server", "ServerClient"}, false); err != nil {
		errors = errors.Append(err)
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
func (*X509Certificate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := X509CertificateAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return X509CertificateLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*X509Certificate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return X509CertificateAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *X509Certificate) ValueForAttribute(name string) interface{} {

	switch name {
	case "CSR":
		return o.CSR
	case "ID":
		return o.ID
	case "certificate":
		return o.Certificate
	case "expirationDate":
		return o.ExpirationDate
	case "extensions":
		return o.Extensions
	case "signer":
		return o.Signer
	case "subjectOverride":
		return o.SubjectOverride
	case "unrevocable":
		return o.Unrevocable
	case "usage":
		return o.Usage
	}

	return nil
}

// X509CertificateAttributesMap represents the map of attribute for X509Certificate.
var X509CertificateAttributesMap = map[string]elemental.AttributeSpecification{
	"CSR": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CSR",
		CreationOnly:   true,
		Description:    `CSR contains the Certificate Signing Request as a PEM encoded string.`,
		Exposed:        true,
		Name:           "CSR",
		Required:       true,
		Type:           "string",
	},
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID contains the identifier of the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `Certificate contains the certificate data in PEM format.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Type:           "string",
	},
	"ExpirationDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description:    `ExpirationDate contains the requested expiration date.`,
		Exposed:        true,
		Name:           "expirationDate",
		SubType:        "string",
		Type:           "time",
	},
	"Extensions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Extensions",
		CreationOnly:   true,
		Description: `Extensions is a list of extensions that can be added as SAN extensions to the
certificate.`,
		Exposed: true,
		Name:    "extensions",
		SubType: "string",
		Type:    "list",
	},
	"Signer": elemental.AttributeSpecification{
		AllowedChoices: []string{"Public", "System"},
		ConvertedName:  "Signer",
		DefaultValue:   X509CertificateSignerPublic,
		Description:    `Selects what CA should sign the certificate.`,
		Exposed:        true,
		Name:           "signer",
		Type:           "enum",
	},
	"SubjectOverride": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SubjectOverride",
		CreationOnly:   true,
		Description:    `Additional subject information to use to override the ones in the CSR.`,
		Exposed:        true,
		Name:           "subjectOverride",
		SubType:        "pkixname",
		Type:           "ref",
	},
	"Unrevocable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Unrevocable",
		Description: `If set to true, the certificate is considered short lived and it will not be
possible to revoke it.`,
		Exposed: true,
		Name:    "unrevocable",
		Type:    "boolean",
	},
	"Usage": elemental.AttributeSpecification{
		AllowedChoices: []string{"Client", "Server", "ServerClient"},
		ConvertedName:  "Usage",
		DefaultValue:   X509CertificateUsageClient,
		Description:    `Usage defines the requested key usage.`,
		Exposed:        true,
		Name:           "usage",
		Type:           "enum",
	},
}

// X509CertificateLowerCaseAttributesMap represents the map of attribute for X509Certificate.
var X509CertificateLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"csr": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CSR",
		CreationOnly:   true,
		Description:    `CSR contains the Certificate Signing Request as a PEM encoded string.`,
		Exposed:        true,
		Name:           "CSR",
		Required:       true,
		Type:           "string",
	},
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID contains the identifier of the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Type:           "string",
	},
	"certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `Certificate contains the certificate data in PEM format.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Type:           "string",
	},
	"expirationdate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description:    `ExpirationDate contains the requested expiration date.`,
		Exposed:        true,
		Name:           "expirationDate",
		SubType:        "string",
		Type:           "time",
	},
	"extensions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Extensions",
		CreationOnly:   true,
		Description: `Extensions is a list of extensions that can be added as SAN extensions to the
certificate.`,
		Exposed: true,
		Name:    "extensions",
		SubType: "string",
		Type:    "list",
	},
	"signer": elemental.AttributeSpecification{
		AllowedChoices: []string{"Public", "System"},
		ConvertedName:  "Signer",
		DefaultValue:   X509CertificateSignerPublic,
		Description:    `Selects what CA should sign the certificate.`,
		Exposed:        true,
		Name:           "signer",
		Type:           "enum",
	},
	"subjectoverride": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SubjectOverride",
		CreationOnly:   true,
		Description:    `Additional subject information to use to override the ones in the CSR.`,
		Exposed:        true,
		Name:           "subjectOverride",
		SubType:        "pkixname",
		Type:           "ref",
	},
	"unrevocable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Unrevocable",
		Description: `If set to true, the certificate is considered short lived and it will not be
possible to revoke it.`,
		Exposed: true,
		Name:    "unrevocable",
		Type:    "boolean",
	},
	"usage": elemental.AttributeSpecification{
		AllowedChoices: []string{"Client", "Server", "ServerClient"},
		ConvertedName:  "Usage",
		DefaultValue:   X509CertificateUsageClient,
		Description:    `Usage defines the requested key usage.`,
		Exposed:        true,
		Name:           "usage",
		Type:           "enum",
	},
}

// SparseX509CertificatesList represents a list of SparseX509Certificates
type SparseX509CertificatesList []*SparseX509Certificate

// Identity returns the identity of the objects in the list.
func (o SparseX509CertificatesList) Identity() elemental.Identity {

	return X509CertificateIdentity
}

// Copy returns a pointer to a copy the SparseX509CertificatesList.
func (o SparseX509CertificatesList) Copy() elemental.Identifiables {

	copy := append(SparseX509CertificatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseX509CertificatesList.
func (o SparseX509CertificatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseX509CertificatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseX509Certificate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseX509CertificatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseX509CertificatesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseX509CertificatesList converted to X509CertificatesList.
func (o SparseX509CertificatesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseX509CertificatesList) Version() int {

	return 1
}

// SparseX509Certificate represents the sparse version of a x509certificate.
type SparseX509Certificate struct {
	// CSR contains the Certificate Signing Request as a PEM encoded string.
	CSR *string `json:"CSR,omitempty" msgpack:"CSR,omitempty" bson:"-" mapstructure:"CSR,omitempty"`

	// ID contains the identifier of the certificate.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Certificate contains the certificate data in PEM format.
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"-" mapstructure:"certificate,omitempty"`

	// ExpirationDate contains the requested expiration date.
	ExpirationDate *time.Time `json:"expirationDate,omitempty" msgpack:"expirationDate,omitempty" bson:"-" mapstructure:"expirationDate,omitempty"`

	// Extensions is a list of extensions that can be added as SAN extensions to the
	// certificate.
	Extensions *[]string `json:"extensions,omitempty" msgpack:"extensions,omitempty" bson:"-" mapstructure:"extensions,omitempty"`

	// Selects what CA should sign the certificate.
	Signer *X509CertificateSignerValue `json:"signer,omitempty" msgpack:"signer,omitempty" bson:"-" mapstructure:"signer,omitempty"`

	// Additional subject information to use to override the ones in the CSR.
	SubjectOverride *PKIXName `json:"subjectOverride,omitempty" msgpack:"subjectOverride,omitempty" bson:"-" mapstructure:"subjectOverride,omitempty"`

	// If set to true, the certificate is considered short lived and it will not be
	// possible to revoke it.
	Unrevocable *bool `json:"unrevocable,omitempty" msgpack:"unrevocable,omitempty" bson:"-" mapstructure:"unrevocable,omitempty"`

	// Usage defines the requested key usage.
	Usage *X509CertificateUsageValue `json:"usage,omitempty" msgpack:"usage,omitempty" bson:"-" mapstructure:"usage,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseX509Certificate returns a new  SparseX509Certificate.
func NewSparseX509Certificate() *SparseX509Certificate {
	return &SparseX509Certificate{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseX509Certificate) Identity() elemental.Identity {

	return X509CertificateIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseX509Certificate) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseX509Certificate) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseX509Certificate) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseX509Certificate) ToPlain() elemental.PlainIdentifiable {

	out := NewX509Certificate()
	if o.CSR != nil {
		out.CSR = *o.CSR
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.ExpirationDate != nil {
		out.ExpirationDate = *o.ExpirationDate
	}
	if o.Extensions != nil {
		out.Extensions = *o.Extensions
	}
	if o.Signer != nil {
		out.Signer = *o.Signer
	}
	if o.SubjectOverride != nil {
		out.SubjectOverride = o.SubjectOverride
	}
	if o.Unrevocable != nil {
		out.Unrevocable = *o.Unrevocable
	}
	if o.Usage != nil {
		out.Usage = *o.Usage
	}

	return out
}

// DeepCopy returns a deep copy if the SparseX509Certificate.
func (o *SparseX509Certificate) DeepCopy() *SparseX509Certificate {

	if o == nil {
		return nil
	}

	out := &SparseX509Certificate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseX509Certificate.
func (o *SparseX509Certificate) DeepCopyInto(out *SparseX509Certificate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseX509Certificate: %s", err))
	}

	*out = *target.(*SparseX509Certificate)
}
