package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AWSAccountIdentity represents the Identity of the object.
var AWSAccountIdentity = elemental.Identity{
	Name:     "awsaccount",
	Category: "awsaccounts",
	Package:  "vince",
	Private:  false,
}

// AWSAccountsList represents a list of AWSAccounts
type AWSAccountsList []*AWSAccount

// Identity returns the identity of the objects in the list.
func (o AWSAccountsList) Identity() elemental.Identity {

	return AWSAccountIdentity
}

// Copy returns a pointer to a copy the AWSAccountsList.
func (o AWSAccountsList) Copy() elemental.Identifiables {

	copy := append(AWSAccountsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AWSAccountsList.
func (o AWSAccountsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AWSAccountsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AWSAccount))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AWSAccountsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AWSAccountsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AWSAccountsList converted to SparseAWSAccountsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AWSAccountsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAWSAccountsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAWSAccount)
	}

	return out
}

// Version returns the version of the content.
func (o AWSAccountsList) Version() int {

	return 1
}

// AWSAccount represents the model of a awsaccount
type AWSAccount struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Contains the AWS access key ID. Aporeto uses this just to retrieve your
	// account ID and does not store the value.
	AccessKeyID string `json:"accessKeyID" msgpack:"accessKeyID" bson:"-" mapstructure:"accessKeyID,omitempty"`

	// Contains your AWS access token. Aporeto uses this just to retrieve your
	// account ID and does not store the value.
	AccessToken string `json:"accessToken" msgpack:"accessToken" bson:"-" mapstructure:"accessToken,omitempty"`

	// Contains your verified account ID.
	AccountID string `json:"accountID" msgpack:"accountID" bson:"accountid" mapstructure:"accountID,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Contains the parent Vince account ID.
	ParentID string `json:"parentID" msgpack:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// Contains the name of the Vince parent account.
	ParentName string `json:"parentName" msgpack:"parentName" bson:"parentname" mapstructure:"parentName,omitempty"`

	// Contains your the region where your AWS account is located.
	Region string `json:"region" msgpack:"region" bson:"region" mapstructure:"region,omitempty"`

	// Contains the AWS secret access key. Aporeto uses this just to retrieve your
	// account ID and does not store the value.
	SecretAccessKey string `json:"secretAccessKey" msgpack:"secretAccessKey" bson:"-" mapstructure:"secretAccessKey,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAWSAccount returns a new *AWSAccount
func NewAWSAccount() *AWSAccount {

	return &AWSAccount{
		ModelVersion:  1,
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *AWSAccount) Identity() elemental.Identity {

	return AWSAccountIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AWSAccount) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AWSAccount) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *AWSAccount) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *AWSAccount) BleveType() string {

	return "awsaccount"
}

// DefaultOrder returns the list of default ordering fields.
func (o *AWSAccount) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *AWSAccount) Doc() string {

	return `Allows you to bind an AWS account to your Aporeto account to allow
auto-registration
of enforcers running on EC2 instances.`
}

func (o *AWSAccount) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *AWSAccount) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *AWSAccount) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *AWSAccount) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *AWSAccount) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *AWSAccount) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *AWSAccount) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *AWSAccount) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *AWSAccount) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *AWSAccount) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *AWSAccount) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AWSAccount) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAWSAccount{
			ID:              &o.ID,
			AccessKeyID:     &o.AccessKeyID,
			AccessToken:     &o.AccessToken,
			AccountID:       &o.AccountID,
			CreateTime:      &o.CreateTime,
			MigrationsLog:   &o.MigrationsLog,
			ParentID:        &o.ParentID,
			ParentName:      &o.ParentName,
			Region:          &o.Region,
			SecretAccessKey: &o.SecretAccessKey,
			UpdateTime:      &o.UpdateTime,
			ZHash:           &o.ZHash,
			Zone:            &o.Zone,
		}
	}

	sp := &SparseAWSAccount{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "accessKeyID":
			sp.AccessKeyID = &(o.AccessKeyID)
		case "accessToken":
			sp.AccessToken = &(o.AccessToken)
		case "accountID":
			sp.AccountID = &(o.AccountID)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "parentID":
			sp.ParentID = &(o.ParentID)
		case "parentName":
			sp.ParentName = &(o.ParentName)
		case "region":
			sp.Region = &(o.Region)
		case "secretAccessKey":
			sp.SecretAccessKey = &(o.SecretAccessKey)
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

// Patch apply the non nil value of a *SparseAWSAccount to the object.
func (o *AWSAccount) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAWSAccount)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.AccessKeyID != nil {
		o.AccessKeyID = *so.AccessKeyID
	}
	if so.AccessToken != nil {
		o.AccessToken = *so.AccessToken
	}
	if so.AccountID != nil {
		o.AccountID = *so.AccountID
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.ParentID != nil {
		o.ParentID = *so.ParentID
	}
	if so.ParentName != nil {
		o.ParentName = *so.ParentName
	}
	if so.Region != nil {
		o.Region = *so.Region
	}
	if so.SecretAccessKey != nil {
		o.SecretAccessKey = *so.SecretAccessKey
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

// DeepCopy returns a deep copy if the AWSAccount.
func (o *AWSAccount) DeepCopy() *AWSAccount {

	if o == nil {
		return nil
	}

	out := &AWSAccount{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AWSAccount.
func (o *AWSAccount) DeepCopyInto(out *AWSAccount) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AWSAccount: %s", err))
	}

	*out = *target.(*AWSAccount)
}

// Validate valides the current information stored into the structure.
func (o *AWSAccount) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("accessKeyID", o.AccessKeyID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("region", o.Region); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("secretAccessKey", o.SecretAccessKey); err != nil {
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
func (*AWSAccount) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AWSAccountAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AWSAccountLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AWSAccount) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AWSAccountAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AWSAccount) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "accessKeyID":
		return o.AccessKeyID
	case "accessToken":
		return o.AccessToken
	case "accountID":
		return o.AccountID
	case "createTime":
		return o.CreateTime
	case "migrationsLog":
		return o.MigrationsLog
	case "parentID":
		return o.ParentID
	case "parentName":
		return o.ParentName
	case "region":
		return o.Region
	case "secretAccessKey":
		return o.SecretAccessKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// AWSAccountAttributesMap represents the map of attribute for AWSAccount.
var AWSAccountAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AccessKeyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessKeyID",
		CreationOnly:   true,
		Description: `Contains the AWS access key ID. Aporeto uses this just to retrieve your
account ID and does not store the value.`,
		Exposed:  true,
		Name:     "accessKeyID",
		Required: true,
		Type:     "string",
	},
	"AccessToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessToken",
		CreationOnly:   true,
		Description: `Contains your AWS access token. Aporeto uses this just to retrieve your
account ID and does not store the value.`,
		Exposed: true,
		Name:    "accessToken",
		Type:    "string",
	},
	"AccountID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AccountID",
		Description:    `Contains your verified account ID.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "accountID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
	"MigrationsLog": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `Contains the parent Vince account ID.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ParentName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentName",
		Description:    `Contains the name of the Vince parent account.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentName",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Region": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Region",
		CreationOnly:   true,
		Description:    `Contains your the region where your AWS account is located.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "region",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"SecretAccessKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SecretAccessKey",
		CreationOnly:   true,
		Description: `Contains the AWS secret access key. Aporeto uses this just to retrieve your
account ID and does not store the value.`,
		Exposed:  true,
		Name:     "secretAccessKey",
		Required: true,
		Type:     "string",
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
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// AWSAccountLowerCaseAttributesMap represents the map of attribute for AWSAccount.
var AWSAccountLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"accesskeyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessKeyID",
		CreationOnly:   true,
		Description: `Contains the AWS access key ID. Aporeto uses this just to retrieve your
account ID and does not store the value.`,
		Exposed:  true,
		Name:     "accessKeyID",
		Required: true,
		Type:     "string",
	},
	"accesstoken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessToken",
		CreationOnly:   true,
		Description: `Contains your AWS access token. Aporeto uses this just to retrieve your
account ID and does not store the value.`,
		Exposed: true,
		Name:    "accessToken",
		Type:    "string",
	},
	"accountid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AccountID",
		Description:    `Contains your verified account ID.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "accountID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
	"migrationslog": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"parentid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `Contains the parent Vince account ID.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"parentname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentName",
		Description:    `Contains the name of the Vince parent account.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentName",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"region": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Region",
		CreationOnly:   true,
		Description:    `Contains your the region where your AWS account is located.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "region",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"secretaccesskey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SecretAccessKey",
		CreationOnly:   true,
		Description: `Contains the AWS secret access key. Aporeto uses this just to retrieve your
account ID and does not store the value.`,
		Exposed:  true,
		Name:     "secretAccessKey",
		Required: true,
		Type:     "string",
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
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseAWSAccountsList represents a list of SparseAWSAccounts
type SparseAWSAccountsList []*SparseAWSAccount

// Identity returns the identity of the objects in the list.
func (o SparseAWSAccountsList) Identity() elemental.Identity {

	return AWSAccountIdentity
}

// Copy returns a pointer to a copy the SparseAWSAccountsList.
func (o SparseAWSAccountsList) Copy() elemental.Identifiables {

	copy := append(SparseAWSAccountsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAWSAccountsList.
func (o SparseAWSAccountsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAWSAccountsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAWSAccount))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAWSAccountsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAWSAccountsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAWSAccountsList converted to AWSAccountsList.
func (o SparseAWSAccountsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAWSAccountsList) Version() int {

	return 1
}

// SparseAWSAccount represents the sparse version of a awsaccount.
type SparseAWSAccount struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Contains the AWS access key ID. Aporeto uses this just to retrieve your
	// account ID and does not store the value.
	AccessKeyID *string `json:"accessKeyID,omitempty" msgpack:"accessKeyID,omitempty" bson:"-" mapstructure:"accessKeyID,omitempty"`

	// Contains your AWS access token. Aporeto uses this just to retrieve your
	// account ID and does not store the value.
	AccessToken *string `json:"accessToken,omitempty" msgpack:"accessToken,omitempty" bson:"-" mapstructure:"accessToken,omitempty"`

	// Contains your verified account ID.
	AccountID *string `json:"accountID,omitempty" msgpack:"accountID,omitempty" bson:"accountid,omitempty" mapstructure:"accountID,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Contains the parent Vince account ID.
	ParentID *string `json:"parentID,omitempty" msgpack:"parentID,omitempty" bson:"parentid,omitempty" mapstructure:"parentID,omitempty"`

	// Contains the name of the Vince parent account.
	ParentName *string `json:"parentName,omitempty" msgpack:"parentName,omitempty" bson:"parentname,omitempty" mapstructure:"parentName,omitempty"`

	// Contains your the region where your AWS account is located.
	Region *string `json:"region,omitempty" msgpack:"region,omitempty" bson:"region,omitempty" mapstructure:"region,omitempty"`

	// Contains the AWS secret access key. Aporeto uses this just to retrieve your
	// account ID and does not store the value.
	SecretAccessKey *string `json:"secretAccessKey,omitempty" msgpack:"secretAccessKey,omitempty" bson:"-" mapstructure:"secretAccessKey,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAWSAccount returns a new  SparseAWSAccount.
func NewSparseAWSAccount() *SparseAWSAccount {
	return &SparseAWSAccount{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAWSAccount) Identity() elemental.Identity {

	return AWSAccountIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAWSAccount) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAWSAccount) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseAWSAccount) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAWSAccount) ToPlain() elemental.PlainIdentifiable {

	out := NewAWSAccount()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.AccessKeyID != nil {
		out.AccessKeyID = *o.AccessKeyID
	}
	if o.AccessToken != nil {
		out.AccessToken = *o.AccessToken
	}
	if o.AccountID != nil {
		out.AccountID = *o.AccountID
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.ParentID != nil {
		out.ParentID = *o.ParentID
	}
	if o.ParentName != nil {
		out.ParentName = *o.ParentName
	}
	if o.Region != nil {
		out.Region = *o.Region
	}
	if o.SecretAccessKey != nil {
		out.SecretAccessKey = *o.SecretAccessKey
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
func (o *SparseAWSAccount) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseAWSAccount) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseAWSAccount) GetMigrationsLog() map[string]string {

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseAWSAccount) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseAWSAccount) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseAWSAccount) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAWSAccount) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAWSAccount) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAWSAccount) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAWSAccount) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAWSAccount.
func (o *SparseAWSAccount) DeepCopy() *SparseAWSAccount {

	if o == nil {
		return nil
	}

	out := &SparseAWSAccount{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAWSAccount.
func (o *SparseAWSAccount) DeepCopyInto(out *SparseAWSAccount) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAWSAccount: %s", err))
	}

	*out = *target.(*SparseAWSAccount)
}
