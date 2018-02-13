package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
)

// AWSAccountIdentity represents the Identity of the object.
var AWSAccountIdentity = elemental.Identity{
	Name:     "awsaccount",
	Category: "awsaccounts",
	Private:  false,
}

// AWSAccountsList represents a list of AWSAccounts
type AWSAccountsList []*AWSAccount

// ContentIdentity returns the identity of the objects in the list.
func (o AWSAccountsList) ContentIdentity() elemental.Identity {

	return AWSAccountIdentity
}

// Copy returns a pointer to a copy the AWSAccountsList.
func (o AWSAccountsList) Copy() elemental.ContentIdentifiable {

	copy := append(AWSAccountsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AWSAccountsList.
func (o AWSAccountsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(AWSAccountsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AWSAccount))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AWSAccountsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AWSAccountsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o AWSAccountsList) Version() int {

	return 1
}

// AWSAccount represents the model of a awsaccount
type AWSAccount struct {
	// AccessKeyID contains the aws access key ID. This is used to retrieve your account id, and it is not stored.
	AccessKeyID string `json:"accessKeyID" bson:"-" mapstructure:"accessKeyID,omitempty"`

	// AccessToken contains your aws access token. It is used to retrieve your account id, and it not stored.
	AccessToken string `json:"accessToken" bson:"-" mapstructure:"accessToken,omitempty"`

	// accountID contains your verified accound id.
	AccountID string `json:"accountID" bson:"accountid" mapstructure:"accountID,omitempty"`

	// ParentID contains the parent Vince account ID.
	ParentID string `json:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// ParentName contains the name of the Vince parent Account.
	ParentName string `json:"parentName" bson:"parentname" mapstructure:"parentName,omitempty"`

	// Region contains your the region where your AWS account is located
	Region string `json:"region" bson:"region" mapstructure:"region,omitempty"`

	// secretAccessKey contains the secret key. It is used to retrieve your account id, and it is not stored.
	SecretAccessKey string `json:"secretAccessKey" bson:"-" mapstructure:"secretAccessKey,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Last update date of the object
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewAWSAccount returns a new *AWSAccount
func NewAWSAccount() *AWSAccount {

	return &AWSAccount{
		ModelVersion: 1,
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

// DefaultOrder returns the list of default ordering fields.
func (o *AWSAccount) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *AWSAccount) Doc() string {
	return `Allows to bind an AWS account to your Aporeto account to allow auto registration of enforcers running on EC2 `
}

func (o *AWSAccount) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *AWSAccount) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("accessKeyID", o.AccessKeyID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("accessKeyID", o.AccessKeyID); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("region", o.Region); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("region", o.Region); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("secretAccessKey", o.SecretAccessKey); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("secretAccessKey", o.SecretAccessKey); err != nil {
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

// AWSAccountAttributesMap represents the map of attribute for AWSAccount.
var AWSAccountAttributesMap = map[string]elemental.AttributeSpecification{
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
		Unique:         true,
	},
	"AccessKeyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessKeyID",
		CreationOnly:   true,
		Description:    `AccessKeyID contains the aws access key ID. This is used to retrieve your account id, and it is not stored.`,
		Exposed:        true,
		Format:         "free",
		Name:           "accessKeyID",
		Required:       true,
		Type:           "string",
	},
	"AccessToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessToken",
		CreationOnly:   true,
		Description:    `AccessToken contains your aws access token. It is used to retrieve your account id, and it not stored.`,
		Exposed:        true,
		Format:         "free",
		Name:           "accessToken",
		Type:           "string",
	},
	"AccountID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AccountID",
		Description:    `accountID contains your verified accound id.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
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
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `ParentID contains the parent Vince account ID.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ParentName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentName",
		Description:    `ParentName contains the name of the Vince parent Account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentName",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Region": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Region",
		CreationOnly:   true,
		Description:    `Region contains your the region where your AWS account is located`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
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
		Description:    `secretAccessKey contains the secret key. It is used to retrieve your account id, and it is not stored.`,
		Exposed:        true,
		Format:         "free",
		Name:           "secretAccessKey",
		Required:       true,
		Type:           "string",
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

// AWSAccountLowerCaseAttributesMap represents the map of attribute for AWSAccount.
var AWSAccountLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		Unique:         true,
	},
	"accesskeyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessKeyID",
		CreationOnly:   true,
		Description:    `AccessKeyID contains the aws access key ID. This is used to retrieve your account id, and it is not stored.`,
		Exposed:        true,
		Format:         "free",
		Name:           "accessKeyID",
		Required:       true,
		Type:           "string",
	},
	"accesstoken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessToken",
		CreationOnly:   true,
		Description:    `AccessToken contains your aws access token. It is used to retrieve your account id, and it not stored.`,
		Exposed:        true,
		Format:         "free",
		Name:           "accessToken",
		Type:           "string",
	},
	"accountid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AccountID",
		Description:    `accountID contains your verified accound id.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
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
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"parentid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `ParentID contains the parent Vince account ID.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"parentname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentName",
		Description:    `ParentName contains the name of the Vince parent Account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentName",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"region": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Region",
		CreationOnly:   true,
		Description:    `Region contains your the region where your AWS account is located`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
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
		Description:    `secretAccessKey contains the secret key. It is used to retrieve your account id, and it is not stored.`,
		Exposed:        true,
		Format:         "free",
		Name:           "secretAccessKey",
		Required:       true,
		Type:           "string",
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
