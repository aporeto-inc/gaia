package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// SSHAuthorityAlgValue represents the possible values for attribute "alg".
type SSHAuthorityAlgValue string

const (
	// SSHAuthorityAlgECDSA represents the value ECDSA.
	SSHAuthorityAlgECDSA SSHAuthorityAlgValue = "ECDSA"

	// SSHAuthorityAlgRSA represents the value RSA.
	SSHAuthorityAlgRSA SSHAuthorityAlgValue = "RSA"
)

// SSHAuthorityIdentity represents the Identity of the object.
var SSHAuthorityIdentity = elemental.Identity{
	Name:     "sshauthority",
	Category: "sshauthorities",
	Package:  "barret",
	Private:  true,
}

// SSHAuthoritiesList represents a list of SSHAuthorities
type SSHAuthoritiesList []*SSHAuthority

// Identity returns the identity of the objects in the list.
func (o SSHAuthoritiesList) Identity() elemental.Identity {

	return SSHAuthorityIdentity
}

// Copy returns a pointer to a copy the SSHAuthoritiesList.
func (o SSHAuthoritiesList) Copy() elemental.Identifiables {

	copy := append(SSHAuthoritiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SSHAuthoritiesList.
func (o SSHAuthoritiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SSHAuthoritiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SSHAuthority))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SSHAuthoritiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SSHAuthoritiesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the SSHAuthoritiesList converted to SparseSSHAuthoritiesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o SSHAuthoritiesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseSSHAuthoritiesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseSSHAuthority)
	}

	return out
}

// Version returns the version of the content.
func (o SSHAuthoritiesList) Version() int {

	return 1
}

// SSHAuthority represents the model of a sshauthority
type SSHAuthority struct {
	// ID is the identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Algorithm to use for the CA.
	Alg SSHAuthorityAlgValue `json:"alg" msgpack:"alg" bson:"alg" mapstructure:"alg,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Contains the private key of the CA.
	PrivateKey string `json:"-" msgpack:"-" bson:"privatekey" mapstructure:"-,omitempty"`

	// Contains the public key of the CA.
	PublicKey string `json:"publicKey" msgpack:"publicKey" bson:"publickey" mapstructure:"publicKey,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSSHAuthority returns a new *SSHAuthority
func NewSSHAuthority() *SSHAuthority {

	return &SSHAuthority{
		ModelVersion: 1,
		Alg:          SSHAuthorityAlgECDSA,
	}
}

// Identity returns the Identity of the object.
func (o *SSHAuthority) Identity() elemental.Identity {

	return SSHAuthorityIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SSHAuthority) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SSHAuthority) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *SSHAuthority) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *SSHAuthority) BleveType() string {

	return "sshauthority"
}

// DefaultOrder returns the list of default ordering fields.
func (o *SSHAuthority) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *SSHAuthority) Doc() string {

	return `Internal api to deliver SSH CA.`
}

func (o *SSHAuthority) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SSHAuthority) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *SSHAuthority) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetName returns the Name of the receiver.
func (o *SSHAuthority) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *SSHAuthority) SetName(name string) {

	o.Name = name
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SSHAuthority) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *SSHAuthority) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SSHAuthority) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *SSHAuthority) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *SSHAuthority) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *SSHAuthority) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *SSHAuthority) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseSSHAuthority{
			ID:         &o.ID,
			Alg:        &o.Alg,
			CreateTime: &o.CreateTime,
			Name:       &o.Name,
			PrivateKey: &o.PrivateKey,
			PublicKey:  &o.PublicKey,
			UpdateTime: &o.UpdateTime,
			ZHash:      &o.ZHash,
			Zone:       &o.Zone,
		}
	}

	sp := &SparseSSHAuthority{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "alg":
			sp.Alg = &(o.Alg)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "name":
			sp.Name = &(o.Name)
		case "privateKey":
			sp.PrivateKey = &(o.PrivateKey)
		case "publicKey":
			sp.PublicKey = &(o.PublicKey)
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

// Patch apply the non nil value of a *SparseSSHAuthority to the object.
func (o *SSHAuthority) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseSSHAuthority)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Alg != nil {
		o.Alg = *so.Alg
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.PrivateKey != nil {
		o.PrivateKey = *so.PrivateKey
	}
	if so.PublicKey != nil {
		o.PublicKey = *so.PublicKey
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

// DeepCopy returns a deep copy if the SSHAuthority.
func (o *SSHAuthority) DeepCopy() *SSHAuthority {

	if o == nil {
		return nil
	}

	out := &SSHAuthority{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SSHAuthority.
func (o *SSHAuthority) DeepCopyInto(out *SSHAuthority) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SSHAuthority: %s", err))
	}

	*out = *target.(*SSHAuthority)
}

// Validate valides the current information stored into the structure.
func (o *SSHAuthority) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("alg", string(o.Alg), []string{"RSA", "ECDSA"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
func (*SSHAuthority) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := SSHAuthorityAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return SSHAuthorityLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*SSHAuthority) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SSHAuthorityAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *SSHAuthority) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "alg":
		return o.Alg
	case "createTime":
		return o.CreateTime
	case "name":
		return o.Name
	case "privateKey":
		return o.PrivateKey
	case "publicKey":
		return o.PublicKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// SSHAuthorityAttributesMap represents the map of attribute for SSHAuthority.
var SSHAuthorityAttributesMap = map[string]elemental.AttributeSpecification{
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
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Alg": elemental.AttributeSpecification{
		AllowedChoices: []string{"RSA", "ECDSA"},
		ConvertedName:  "Alg",
		DefaultValue:   SSHAuthorityAlgECDSA,
		Description:    `Algorithm to use for the CA.`,
		Exposed:        true,
		Name:           "alg",
		Stored:         true,
		Type:           "enum",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"PrivateKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PrivateKey",
		Description:    `Contains the private key of the CA.`,
		Name:           "privateKey",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"PublicKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PublicKey",
		Description:    `Contains the public key of the CA.`,
		Exposed:        true,
		Name:           "publicKey",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
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
		Exposed:   true,
		Getter:    true,
		Name:      "zone",
		ReadOnly:  true,
		Setter:    true,
		Stored:    true,
		Transient: true,
		Type:      "integer",
	},
}

// SSHAuthorityLowerCaseAttributesMap represents the map of attribute for SSHAuthority.
var SSHAuthorityLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"alg": elemental.AttributeSpecification{
		AllowedChoices: []string{"RSA", "ECDSA"},
		ConvertedName:  "Alg",
		DefaultValue:   SSHAuthorityAlgECDSA,
		Description:    `Algorithm to use for the CA.`,
		Exposed:        true,
		Name:           "alg",
		Stored:         true,
		Type:           "enum",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"privatekey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PrivateKey",
		Description:    `Contains the private key of the CA.`,
		Name:           "privateKey",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"publickey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PublicKey",
		Description:    `Contains the public key of the CA.`,
		Exposed:        true,
		Name:           "publicKey",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
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
		Exposed:   true,
		Getter:    true,
		Name:      "zone",
		ReadOnly:  true,
		Setter:    true,
		Stored:    true,
		Transient: true,
		Type:      "integer",
	},
}

// SparseSSHAuthoritiesList represents a list of SparseSSHAuthorities
type SparseSSHAuthoritiesList []*SparseSSHAuthority

// Identity returns the identity of the objects in the list.
func (o SparseSSHAuthoritiesList) Identity() elemental.Identity {

	return SSHAuthorityIdentity
}

// Copy returns a pointer to a copy the SparseSSHAuthoritiesList.
func (o SparseSSHAuthoritiesList) Copy() elemental.Identifiables {

	copy := append(SparseSSHAuthoritiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseSSHAuthoritiesList.
func (o SparseSSHAuthoritiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseSSHAuthoritiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseSSHAuthority))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseSSHAuthoritiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseSSHAuthoritiesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseSSHAuthoritiesList converted to SSHAuthoritiesList.
func (o SparseSSHAuthoritiesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseSSHAuthoritiesList) Version() int {

	return 1
}

// SparseSSHAuthority represents the sparse version of a sshauthority.
type SparseSSHAuthority struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Algorithm to use for the CA.
	Alg *SSHAuthorityAlgValue `json:"alg,omitempty" msgpack:"alg,omitempty" bson:"alg,omitempty" mapstructure:"alg,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Contains the private key of the CA.
	PrivateKey *string `json:"-" msgpack:"-" bson:"privatekey,omitempty" mapstructure:"-,omitempty"`

	// Contains the public key of the CA.
	PublicKey *string `json:"publicKey,omitempty" msgpack:"publicKey,omitempty" bson:"publickey,omitempty" mapstructure:"publicKey,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseSSHAuthority returns a new  SparseSSHAuthority.
func NewSparseSSHAuthority() *SparseSSHAuthority {
	return &SparseSSHAuthority{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseSSHAuthority) Identity() elemental.Identity {

	return SSHAuthorityIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseSSHAuthority) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseSSHAuthority) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseSSHAuthority) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseSSHAuthority) ToPlain() elemental.PlainIdentifiable {

	out := NewSSHAuthority()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Alg != nil {
		out.Alg = *o.Alg
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.PrivateKey != nil {
		out.PrivateKey = *o.PrivateKey
	}
	if o.PublicKey != nil {
		out.PublicKey = *o.PublicKey
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

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseSSHAuthority) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseSSHAuthority) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetName returns the Name of the receiver.
func (o *SparseSSHAuthority) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseSSHAuthority) SetName(name string) {

	o.Name = &name
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseSSHAuthority) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseSSHAuthority) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseSSHAuthority) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseSSHAuthority) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseSSHAuthority) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseSSHAuthority) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseSSHAuthority.
func (o *SparseSSHAuthority) DeepCopy() *SparseSSHAuthority {

	if o == nil {
		return nil
	}

	out := &SparseSSHAuthority{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseSSHAuthority.
func (o *SparseSSHAuthority) DeepCopyInto(out *SparseSSHAuthority) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseSSHAuthority: %s", err))
	}

	*out = *target.(*SparseSSHAuthority)
}
