package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// RevocationIdentity represents the Identity of the object.
var RevocationIdentity = elemental.Identity{
	Name:     "revocation",
	Category: "revocations",
	Package:  "barret",
	Private:  true,
}

// RevocationsList represents a list of Revocations
type RevocationsList []*Revocation

// Identity returns the identity of the objects in the list.
func (o RevocationsList) Identity() elemental.Identity {

	return RevocationIdentity
}

// Copy returns a pointer to a copy the RevocationsList.
func (o RevocationsList) Copy() elemental.Identifiables {

	copy := append(RevocationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the RevocationsList.
func (o RevocationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(RevocationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Revocation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o RevocationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RevocationsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the RevocationsList converted to SparseRevocationsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o RevocationsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o RevocationsList) Version() int {

	return 1
}

// Revocation represents the model of a revocation
type Revocation struct {
	// ID contains the ID of the revocation.
	ID string `json:"-" bson:"_id" mapstructure:"-,omitempty"`

	// Contains the certificate expiration date. This will be used to clean up revoked
	// certificates that have expired.
	ExpirationDate time.Time `json:"expirationDate" bson:"expirationdate" mapstructure:"expirationDate,omitempty"`

	// Set time from when the certificate will be revoked.
	RevokeDate time.Time `json:"revokeDate" bson:"revokedate" mapstructure:"revokeDate,omitempty"`

	// SerialNumber of the revoked certificate.
	SerialNumber string `json:"serialNumber" bson:"serialnumber" mapstructure:"serialNumber,omitempty"`

	// Subject of the certificate related to the revocation.
	Subject string `json:"subject" bson:"subject" mapstructure:"subject,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewRevocation returns a new *Revocation
func NewRevocation() *Revocation {

	return &Revocation{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Revocation) Identity() elemental.Identity {

	return RevocationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Revocation) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Revocation) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Revocation) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Revocation) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Revocation) Doc() string {
	return `Used to revoke a certificate.`
}

func (o *Revocation) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Revocation) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseRevocation{
			ID:             &o.ID,
			ExpirationDate: &o.ExpirationDate,
			RevokeDate:     &o.RevokeDate,
			SerialNumber:   &o.SerialNumber,
			Subject:        &o.Subject,
		}
	}

	sp := &SparseRevocation{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "expirationDate":
			sp.ExpirationDate = &(o.ExpirationDate)
		case "revokeDate":
			sp.RevokeDate = &(o.RevokeDate)
		case "serialNumber":
			sp.SerialNumber = &(o.SerialNumber)
		case "subject":
			sp.Subject = &(o.Subject)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseRevocation to the object.
func (o *Revocation) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseRevocation)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.ExpirationDate != nil {
		o.ExpirationDate = *so.ExpirationDate
	}
	if so.RevokeDate != nil {
		o.RevokeDate = *so.RevokeDate
	}
	if so.SerialNumber != nil {
		o.SerialNumber = *so.SerialNumber
	}
	if so.Subject != nil {
		o.Subject = *so.Subject
	}
}

// Validate valides the current information stored into the structure.
func (o *Revocation) Validate() error {

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
func (*Revocation) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RevocationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RevocationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Revocation) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RevocationAttributesMap
}

// RevocationAttributesMap represents the map of attribute for Revocation.
var RevocationAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID contains the ID of the revocation.`,
		Identifier:     true,
		Name:           "ID",
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ExpirationDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description: `Contains the certificate expiration date. This will be used to clean up revoked
certificates that have expired.`,
		Exposed: true,
		Name:    "expirationDate",
		Stored:  true,
		Type:    "time",
	},
	"RevokeDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RevokeDate",
		Description:    `Set time from when the certificate will be revoked.`,
		Exposed:        true,
		Name:           "revokeDate",
		Stored:         true,
		Type:           "time",
	},
	"SerialNumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SerialNumber",
		CreationOnly:   true,
		Description:    `SerialNumber of the revoked certificate.`,
		Exposed:        true,
		Name:           "serialNumber",
		Stored:         true,
		Type:           "string",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		CreationOnly:   true,
		Description:    `Subject of the certificate related to the revocation.`,
		Exposed:        true,
		Name:           "subject",
		Stored:         true,
		Type:           "string",
	},
}

// RevocationLowerCaseAttributesMap represents the map of attribute for Revocation.
var RevocationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID contains the ID of the revocation.`,
		Identifier:     true,
		Name:           "ID",
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"expirationdate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description: `Contains the certificate expiration date. This will be used to clean up revoked
certificates that have expired.`,
		Exposed: true,
		Name:    "expirationDate",
		Stored:  true,
		Type:    "time",
	},
	"revokedate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RevokeDate",
		Description:    `Set time from when the certificate will be revoked.`,
		Exposed:        true,
		Name:           "revokeDate",
		Stored:         true,
		Type:           "time",
	},
	"serialnumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SerialNumber",
		CreationOnly:   true,
		Description:    `SerialNumber of the revoked certificate.`,
		Exposed:        true,
		Name:           "serialNumber",
		Stored:         true,
		Type:           "string",
	},
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		CreationOnly:   true,
		Description:    `Subject of the certificate related to the revocation.`,
		Exposed:        true,
		Name:           "subject",
		Stored:         true,
		Type:           "string",
	},
}

// SparseRevocationsList represents a list of SparseRevocations
type SparseRevocationsList []*SparseRevocation

// Identity returns the identity of the objects in the list.
func (o SparseRevocationsList) Identity() elemental.Identity {

	return RevocationIdentity
}

// Copy returns a pointer to a copy the SparseRevocationsList.
func (o SparseRevocationsList) Copy() elemental.Identifiables {

	copy := append(SparseRevocationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseRevocationsList.
func (o SparseRevocationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseRevocationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseRevocation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseRevocationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseRevocationsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseRevocationsList converted to RevocationsList.
func (o SparseRevocationsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseRevocationsList) Version() int {

	return 1
}

// SparseRevocation represents the sparse version of a revocation.
type SparseRevocation struct {
	// ID contains the ID of the revocation.
	ID *string `json:"-,omitempty" bson:"_id" mapstructure:"-,omitempty"`

	// Contains the certificate expiration date. This will be used to clean up revoked
	// certificates that have expired.
	ExpirationDate *time.Time `json:"expirationDate,omitempty" bson:"expirationdate" mapstructure:"expirationDate,omitempty"`

	// Set time from when the certificate will be revoked.
	RevokeDate *time.Time `json:"revokeDate,omitempty" bson:"revokedate" mapstructure:"revokeDate,omitempty"`

	// SerialNumber of the revoked certificate.
	SerialNumber *string `json:"serialNumber,omitempty" bson:"serialnumber" mapstructure:"serialNumber,omitempty"`

	// Subject of the certificate related to the revocation.
	Subject *string `json:"subject,omitempty" bson:"subject" mapstructure:"subject,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseRevocation returns a new  SparseRevocation.
func NewSparseRevocation() *SparseRevocation {
	return &SparseRevocation{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseRevocation) Identity() elemental.Identity {

	return RevocationIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseRevocation) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseRevocation) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseRevocation) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseRevocation) ToPlain() elemental.PlainIdentifiable {

	out := NewRevocation()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.ExpirationDate != nil {
		out.ExpirationDate = *o.ExpirationDate
	}
	if o.RevokeDate != nil {
		out.RevokeDate = *o.RevokeDate
	}
	if o.SerialNumber != nil {
		out.SerialNumber = *o.SerialNumber
	}
	if o.Subject != nil {
		out.Subject = *o.Subject
	}

	return out
}
