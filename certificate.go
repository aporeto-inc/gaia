package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
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
	Package:  "vince",
	Private:  false,
}

// CertificatesList represents a list of Certificates
type CertificatesList []*Certificate

// Identity returns the identity of the objects in the list.
func (o CertificatesList) Identity() elemental.Identity {

	return CertificateIdentity
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
		out = append(out, obj.(*Certificate))
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
func (o CertificatesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o CertificatesList) Version() int {

	return 1
}

// Certificate represents the model of a certificate
type Certificate struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Admin determines if the certificate must be added to the admin list.
	Admin bool `json:"admin" bson:"admin" mapstructure:"admin,omitempty"`

	// CommonName (CN) for the user certificate.
	CommonName string `json:"commonName" bson:"commonname" mapstructure:"commonName,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Certificate provides a certificate for the user.
	Data string `json:"data" bson:"data" mapstructure:"data,omitempty"`

	// e-mail address of the user.
	Email string `json:"email" bson:"email" mapstructure:"email,omitempty"`

	// CertificateExpirationDate indicates the expiration day for the certificate.
	ExpirationDate time.Time `json:"expirationDate" bson:"expirationdate" mapstructure:"expirationDate,omitempty"`

	// CertificateKey provides the key for the user. Only available at create or update
	// time.
	Key string `json:"key" bson:"-" mapstructure:"key,omitempty"`

	// Name of the certificate.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// OrganizationalUnits attribute for the generated certificates.
	OrganizationalUnits []string `json:"organizationalUnits" bson:"organizationalunits" mapstructure:"organizationalUnits,omitempty"`

	// ParentID holds the parent account ID.
	ParentID string `json:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// Passphrase to use for the generated p12.
	Passphrase string `json:"passphrase" bson:"-" mapstructure:"passphrase,omitempty"`

	// SerialNumber of the certificate.
	SerialNumber string `json:"serialNumber" bson:"serialnumber" mapstructure:"serialNumber,omitempty"`

	// CertificateStatus provides the status of the certificate. Update with RENEW to
	// get a new certificate.
	Status CertificateStatusValue `json:"status" bson:"status" mapstructure:"status,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewCertificate returns a new *Certificate
func NewCertificate() *Certificate {

	return &Certificate{
		ModelVersion: 1,
		Status:       CertificateStatusValid,
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

// GetZHash returns the ZHash of the receiver.
func (o *Certificate) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Certificate) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Certificate) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Certificate) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Certificate) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCertificate{
			ID:                  &o.ID,
			Admin:               &o.Admin,
			CommonName:          &o.CommonName,
			CreateTime:          &o.CreateTime,
			Data:                &o.Data,
			Email:               &o.Email,
			ExpirationDate:      &o.ExpirationDate,
			Key:                 &o.Key,
			Name:                &o.Name,
			OrganizationalUnits: &o.OrganizationalUnits,
			ParentID:            &o.ParentID,
			Passphrase:          &o.Passphrase,
			SerialNumber:        &o.SerialNumber,
			Status:              &o.Status,
			UpdateTime:          &o.UpdateTime,
			ZHash:               &o.ZHash,
			Zone:                &o.Zone,
		}
	}

	sp := &SparseCertificate{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "admin":
			sp.Admin = &(o.Admin)
		case "commonName":
			sp.CommonName = &(o.CommonName)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "data":
			sp.Data = &(o.Data)
		case "email":
			sp.Email = &(o.Email)
		case "expirationDate":
			sp.ExpirationDate = &(o.ExpirationDate)
		case "key":
			sp.Key = &(o.Key)
		case "name":
			sp.Name = &(o.Name)
		case "organizationalUnits":
			sp.OrganizationalUnits = &(o.OrganizationalUnits)
		case "parentID":
			sp.ParentID = &(o.ParentID)
		case "passphrase":
			sp.Passphrase = &(o.Passphrase)
		case "serialNumber":
			sp.SerialNumber = &(o.SerialNumber)
		case "status":
			sp.Status = &(o.Status)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCertificate to the object.
func (o *Certificate) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCertificate)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Admin != nil {
		o.Admin = *so.Admin
	}
	if so.CommonName != nil {
		o.CommonName = *so.CommonName
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Data != nil {
		o.Data = *so.Data
	}
	if so.Email != nil {
		o.Email = *so.Email
	}
	if so.ExpirationDate != nil {
		o.ExpirationDate = *so.ExpirationDate
	}
	if so.Key != nil {
		o.Key = *so.Key
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.OrganizationalUnits != nil {
		o.OrganizationalUnits = *so.OrganizationalUnits
	}
	if so.ParentID != nil {
		o.ParentID = *so.ParentID
	}
	if so.Passphrase != nil {
		o.Passphrase = *so.Passphrase
	}
	if so.SerialNumber != nil {
		o.SerialNumber = *so.SerialNumber
	}
	if so.Status != nil {
		o.Status = *so.Status
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the Certificate.
func (o *Certificate) DeepCopy() *Certificate {

	if o == nil {
		return nil
	}

	out := &Certificate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Certificate.
func (o *Certificate) DeepCopyInto(out *Certificate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Certificate: %s", err))
	}

	*out = *target.(*Certificate)
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

	if err := elemental.ValidateRequiredString("email", o.Email); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Certificate) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "admin":
		return o.Admin
	case "commonName":
		return o.CommonName
	case "createTime":
		return o.CreateTime
	case "data":
		return o.Data
	case "email":
		return o.Email
	case "expirationDate":
		return o.ExpirationDate
	case "key":
		return o.Key
	case "name":
		return o.Name
	case "organizationalUnits":
		return o.OrganizationalUnits
	case "parentID":
		return o.ParentID
	case "passphrase":
		return o.Passphrase
	case "serialNumber":
		return o.SerialNumber
	case "status":
		return o.Status
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
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
		Description:    `CommonName (CN) for the user certificate.`,
		Exposed:        true,
		Filterable:     true,
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
		Description:    `Certificate provides a certificate for the user.`,
		Exposed:        true,
		Name:           "data",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		CreationOnly:   true,
		Description:    `e-mail address of the user.`,
		Exposed:        true,
		Filterable:     true,
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
		Name:           "expirationDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description: `CertificateKey provides the key for the user. Only available at create or update
time.`,
		Exposed:  true,
		Name:     "key",
		ReadOnly: true,
		Type:     "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the certificate.`,
		Exposed:        true,
		Filterable:     true,
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
		Description:    `OrganizationalUnits attribute for the generated certificates.`,
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
		Description: `CertificateStatus provides the status of the certificate. Update with RENEW to
get a new certificate.`,
		Exposed:    true,
		Filterable: true,
		Name:       "status",
		Orderable:  true,
		Stored:     true,
		Type:       "enum",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"ZHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
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
		Description:    `CommonName (CN) for the user certificate.`,
		Exposed:        true,
		Filterable:     true,
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
		Description:    `Certificate provides a certificate for the user.`,
		Exposed:        true,
		Name:           "data",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		CreationOnly:   true,
		Description:    `e-mail address of the user.`,
		Exposed:        true,
		Filterable:     true,
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
		Name:           "expirationDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description: `CertificateKey provides the key for the user. Only available at create or update
time.`,
		Exposed:  true,
		Name:     "key",
		ReadOnly: true,
		Type:     "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the certificate.`,
		Exposed:        true,
		Filterable:     true,
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
		Description:    `OrganizationalUnits attribute for the generated certificates.`,
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
		Description: `CertificateStatus provides the status of the certificate. Update with RENEW to
get a new certificate.`,
		Exposed:    true,
		Filterable: true,
		Name:       "status",
		Orderable:  true,
		Stored:     true,
		Type:       "enum",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"zhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// SparseCertificatesList represents a list of SparseCertificates
type SparseCertificatesList []*SparseCertificate

// Identity returns the identity of the objects in the list.
func (o SparseCertificatesList) Identity() elemental.Identity {

	return CertificateIdentity
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
		out = append(out, obj.(*SparseCertificate))
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

// SparseCertificate represents the sparse version of a certificate.
type SparseCertificate struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Admin determines if the certificate must be added to the admin list.
	Admin *bool `json:"admin,omitempty" bson:"admin" mapstructure:"admin,omitempty"`

	// CommonName (CN) for the user certificate.
	CommonName *string `json:"commonName,omitempty" bson:"commonname" mapstructure:"commonName,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Certificate provides a certificate for the user.
	Data *string `json:"data,omitempty" bson:"data" mapstructure:"data,omitempty"`

	// e-mail address of the user.
	Email *string `json:"email,omitempty" bson:"email" mapstructure:"email,omitempty"`

	// CertificateExpirationDate indicates the expiration day for the certificate.
	ExpirationDate *time.Time `json:"expirationDate,omitempty" bson:"expirationdate" mapstructure:"expirationDate,omitempty"`

	// CertificateKey provides the key for the user. Only available at create or update
	// time.
	Key *string `json:"key,omitempty" bson:"-" mapstructure:"key,omitempty"`

	// Name of the certificate.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// OrganizationalUnits attribute for the generated certificates.
	OrganizationalUnits *[]string `json:"organizationalUnits,omitempty" bson:"organizationalunits" mapstructure:"organizationalUnits,omitempty"`

	// ParentID holds the parent account ID.
	ParentID *string `json:"parentID,omitempty" bson:"parentid" mapstructure:"parentID,omitempty"`

	// Passphrase to use for the generated p12.
	Passphrase *string `json:"passphrase,omitempty" bson:"-" mapstructure:"passphrase,omitempty"`

	// SerialNumber of the certificate.
	SerialNumber *string `json:"serialNumber,omitempty" bson:"serialnumber" mapstructure:"serialNumber,omitempty"`

	// CertificateStatus provides the status of the certificate. Update with RENEW to
	// get a new certificate.
	Status *CertificateStatusValue `json:"status,omitempty" bson:"status" mapstructure:"status,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-,omitempty" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-,omitempty" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseCertificate returns a new  SparseCertificate.
func NewSparseCertificate() *SparseCertificate {
	return &SparseCertificate{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCertificate) Identity() elemental.Identity {

	return CertificateIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCertificate) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCertificate) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseCertificate) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCertificate) ToPlain() elemental.PlainIdentifiable {

	out := NewCertificate()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Admin != nil {
		out.Admin = *o.Admin
	}
	if o.CommonName != nil {
		out.CommonName = *o.CommonName
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Data != nil {
		out.Data = *o.Data
	}
	if o.Email != nil {
		out.Email = *o.Email
	}
	if o.ExpirationDate != nil {
		out.ExpirationDate = *o.ExpirationDate
	}
	if o.Key != nil {
		out.Key = *o.Key
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.OrganizationalUnits != nil {
		out.OrganizationalUnits = *o.OrganizationalUnits
	}
	if o.ParentID != nil {
		out.ParentID = *o.ParentID
	}
	if o.Passphrase != nil {
		out.Passphrase = *o.Passphrase
	}
	if o.SerialNumber != nil {
		out.SerialNumber = *o.SerialNumber
	}
	if o.Status != nil {
		out.Status = *o.Status
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCertificate) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCertificate) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCertificate) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCertificate) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCertificate.
func (o *SparseCertificate) DeepCopy() *SparseCertificate {

	if o == nil {
		return nil
	}

	out := &SparseCertificate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCertificate.
func (o *SparseCertificate) DeepCopyInto(out *SparseCertificate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCertificate: %s", err))
	}

	*out = *target.(*SparseCertificate)
}
