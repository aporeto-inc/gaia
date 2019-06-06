package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AuthorityIdentity represents the Identity of the object.
var AuthorityIdentity = elemental.Identity{
	Name:     "authority",
	Category: "authorities",
	Package:  "barret",
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

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AuthoritiesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AuthoritiesList converted to SparseAuthoritiesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AuthoritiesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAuthoritiesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAuthority)
	}

	return out
}

// Version returns the version of the content.
func (o AuthoritiesList) Version() int {

	return 1
}

// Authority represents the model of a authority
type Authority struct {
	// ID is the identitfier of the Authority.
	ID string `json:"ID" msgpack:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// PEM encoded certificate data.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"certificate" mapstructure:"certificate,omitempty"`

	// CommonName contains the common name of the CA.
	CommonName string `json:"commonName" msgpack:"commonName" bson:"commonname" mapstructure:"commonName,omitempty"`

	// Date of expiration of the authority.
	ExpirationDate time.Time `json:"expirationDate" msgpack:"expirationDate" bson:"expirationdate" mapstructure:"expirationDate,omitempty"`

	// Encrypted private key of the Authority.
	Key string `json:"-" msgpack:"-" bson:"key" mapstructure:"-,omitempty"`

	// serialNumber of the certificate.
	SerialNumber string `json:"serialNumber" msgpack:"serialNumber" bson:"serialnumber" mapstructure:"serialNumber,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	// Defines what zone the CA should live in.
	Zoning int `json:"zoning" msgpack:"zoning" bson:"zoning" mapstructure:"zoning,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
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

// GetZHash returns the ZHash of the receiver.
func (o *Authority) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Authority) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Authority) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Authority) SetZone(zone int) {

	o.Zone = zone
}

// GetZoning returns the Zoning of the receiver.
func (o *Authority) GetZoning() int {

	return o.Zoning
}

// SetZoning sets the property Zoning of the receiver using the given value.
func (o *Authority) SetZoning(zoning int) {

	o.Zoning = zoning
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Authority) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAuthority{
			ID:             &o.ID,
			Certificate:    &o.Certificate,
			CommonName:     &o.CommonName,
			ExpirationDate: &o.ExpirationDate,
			Key:            &o.Key,
			SerialNumber:   &o.SerialNumber,
			ZHash:          &o.ZHash,
			Zone:           &o.Zone,
			Zoning:         &o.Zoning,
		}
	}

	sp := &SparseAuthority{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "commonName":
			sp.CommonName = &(o.CommonName)
		case "expirationDate":
			sp.ExpirationDate = &(o.ExpirationDate)
		case "key":
			sp.Key = &(o.Key)
		case "serialNumber":
			sp.SerialNumber = &(o.SerialNumber)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		case "zoning":
			sp.Zoning = &(o.Zoning)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAuthority to the object.
func (o *Authority) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAuthority)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.CommonName != nil {
		o.CommonName = *so.CommonName
	}
	if so.ExpirationDate != nil {
		o.ExpirationDate = *so.ExpirationDate
	}
	if so.Key != nil {
		o.Key = *so.Key
	}
	if so.SerialNumber != nil {
		o.SerialNumber = *so.SerialNumber
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
	if so.Zoning != nil {
		o.Zoning = *so.Zoning
	}
}

// DeepCopy returns a deep copy if the Authority.
func (o *Authority) DeepCopy() *Authority {

	if o == nil {
		return nil
	}

	out := &Authority{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Authority.
func (o *Authority) DeepCopyInto(out *Authority) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Authority: %s", err))
	}

	*out = *target.(*Authority)
}

// Validate valides the current information stored into the structure.
func (o *Authority) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("commonName", o.CommonName); err != nil {
		requiredErrors = requiredErrors.Append(err)
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Authority) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "certificate":
		return o.Certificate
	case "commonName":
		return o.CommonName
	case "expirationDate":
		return o.ExpirationDate
	case "key":
		return o.Key
	case "serialNumber":
		return o.SerialNumber
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	case "zoning":
		return o.Zoning
	}

	return nil
}

// AuthorityAttributesMap represents the map of attribute for Authority.
var AuthorityAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identitfier of the Authority.`,
		Exposed:        true,
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
		Name:           "serialNumber",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
	"Zoning": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Zoning",
		CreationOnly:   true,
		Description:    `Defines what zone the CA should live in.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zoning",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
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
		Name:           "serialNumber",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
	"zoning": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Zoning",
		CreationOnly:   true,
		Description:    `Defines what zone the CA should live in.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zoning",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
}

// SparseAuthoritiesList represents a list of SparseAuthorities
type SparseAuthoritiesList []*SparseAuthority

// Identity returns the identity of the objects in the list.
func (o SparseAuthoritiesList) Identity() elemental.Identity {

	return AuthorityIdentity
}

// Copy returns a pointer to a copy the SparseAuthoritiesList.
func (o SparseAuthoritiesList) Copy() elemental.Identifiables {

	copy := append(SparseAuthoritiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAuthoritiesList.
func (o SparseAuthoritiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAuthoritiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAuthority))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAuthoritiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAuthoritiesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAuthoritiesList converted to AuthoritiesList.
func (o SparseAuthoritiesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAuthoritiesList) Version() int {

	return 1
}

// SparseAuthority represents the sparse version of a authority.
type SparseAuthority struct {
	// ID is the identitfier of the Authority.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// PEM encoded certificate data.
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"certificate,omitempty" mapstructure:"certificate,omitempty"`

	// CommonName contains the common name of the CA.
	CommonName *string `json:"commonName,omitempty" msgpack:"commonName,omitempty" bson:"commonname,omitempty" mapstructure:"commonName,omitempty"`

	// Date of expiration of the authority.
	ExpirationDate *time.Time `json:"expirationDate,omitempty" msgpack:"expirationDate,omitempty" bson:"expirationdate,omitempty" mapstructure:"expirationDate,omitempty"`

	// Encrypted private key of the Authority.
	Key *string `json:"-" msgpack:"-" bson:"key,omitempty" mapstructure:"-,omitempty"`

	// serialNumber of the certificate.
	SerialNumber *string `json:"serialNumber,omitempty" msgpack:"serialNumber,omitempty" bson:"serialnumber,omitempty" mapstructure:"serialNumber,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	// Defines what zone the CA should live in.
	Zoning *int `json:"zoning,omitempty" msgpack:"zoning,omitempty" bson:"zoning,omitempty" mapstructure:"zoning,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAuthority returns a new  SparseAuthority.
func NewSparseAuthority() *SparseAuthority {
	return &SparseAuthority{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAuthority) Identity() elemental.Identity {

	return AuthorityIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAuthority) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAuthority) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseAuthority) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAuthority) ToPlain() elemental.PlainIdentifiable {

	out := NewAuthority()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.CommonName != nil {
		out.CommonName = *o.CommonName
	}
	if o.ExpirationDate != nil {
		out.ExpirationDate = *o.ExpirationDate
	}
	if o.Key != nil {
		out.Key = *o.Key
	}
	if o.SerialNumber != nil {
		out.SerialNumber = *o.SerialNumber
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}
	if o.Zoning != nil {
		out.Zoning = *o.Zoning
	}

	return out
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAuthority) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAuthority) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAuthority) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAuthority) SetZone(zone int) {

	o.Zone = &zone
}

// GetZoning returns the Zoning of the receiver.
func (o *SparseAuthority) GetZoning() int {

	return *o.Zoning
}

// SetZoning sets the property Zoning of the receiver using the address of the given value.
func (o *SparseAuthority) SetZoning(zoning int) {

	o.Zoning = &zoning
}

// DeepCopy returns a deep copy if the SparseAuthority.
func (o *SparseAuthority) DeepCopy() *SparseAuthority {

	if o == nil {
		return nil
	}

	out := &SparseAuthority{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAuthority.
func (o *SparseAuthority) DeepCopyInto(out *SparseAuthority) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAuthority: %s", err))
	}

	*out = *target.(*SparseAuthority)
}
