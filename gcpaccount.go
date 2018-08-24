package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// GCPAccountIndexes lists the attribute compound indexes.
var GCPAccountIndexes = [][]string{}

// GCPAccountIdentity represents the Identity of the object.
var GCPAccountIdentity = elemental.Identity{
	Name:     "gcpaccount",
	Category: "gcpaccounts",
	Private:  false,
}

// GCPAccountsList represents a list of GCPAccounts
type GCPAccountsList []*GCPAccount

// Identity returns the identity of the objects in the list.
func (o GCPAccountsList) Identity() elemental.Identity {

	return GCPAccountIdentity
}

// Copy returns a pointer to a copy the GCPAccountsList.
func (o GCPAccountsList) Copy() elemental.Identifiables {

	copy := append(GCPAccountsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the GCPAccountsList.
func (o GCPAccountsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(GCPAccountsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*GCPAccount))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o GCPAccountsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o GCPAccountsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o GCPAccountsList) Version() int {

	return 1
}

// GCPAccount represents the model of a gcpaccount
type GCPAccount struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Credential file to use to validate your account. You can create a temporary
	// service account with read only permissions. We do not store this file.
	JSONCredentialFile string `json:"JSONCredentialFile" bson:"-" mapstructure:"JSONCredentialFile,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// ParentID contains the parent Vince account ID.
	ParentID string `json:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// ParentName contains the name of the Vince parent Account.
	ParentName string `json:"parentName" bson:"parentname" mapstructure:"parentName,omitempty"`

	// projectID contains your verified accound id.
	ProjectID string `json:"projectID" bson:"projectid" mapstructure:"projectID,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewGCPAccount returns a new *GCPAccount
func NewGCPAccount() *GCPAccount {

	return &GCPAccount{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *GCPAccount) Identity() elemental.Identity {

	return GCPAccountIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GCPAccount) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GCPAccount) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *GCPAccount) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *GCPAccount) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *GCPAccount) Doc() string {
	return `Allows to bind an GCP account to your Aporeto account to allow auto registration
of enforcers running on Google Cloud. We ask for the credential for the sole
purpose of validating you are the owner the project ID. We don't store this
credentials and we do not send any api call to your GCP project.`
}

func (o *GCPAccount) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *GCPAccount) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("JSONCredentialFile", o.JSONCredentialFile); err != nil {
		requiredErrors = append(requiredErrors, err)
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
func (*GCPAccount) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := GCPAccountAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return GCPAccountLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*GCPAccount) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return GCPAccountAttributesMap
}

// GCPAccountAttributesMap represents the map of attribute for GCPAccount.
var GCPAccountAttributesMap = map[string]elemental.AttributeSpecification{
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
	"JSONCredentialFile": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "JSONCredentialFile",
		CreationOnly:   true,
		Description: `Credential file to use to validate your account. You can create a temporary
service account with read only permissions. We do not store this file.`,
		Exposed:  true,
		Name:     "JSONCredentialFile",
		Required: true,
		Type:     "string",
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
		Name:           "parentName",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ProjectID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ProjectID",
		Description:    `projectID contains your verified accound id.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "projectID",
		Orderable:      true,
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
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}

// GCPAccountLowerCaseAttributesMap represents the map of attribute for GCPAccount.
var GCPAccountLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"jsoncredentialfile": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "JSONCredentialFile",
		CreationOnly:   true,
		Description: `Credential file to use to validate your account. You can create a temporary
service account with read only permissions. We do not store this file.`,
		Exposed:  true,
		Name:     "JSONCredentialFile",
		Required: true,
		Type:     "string",
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
		Name:           "parentName",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"projectid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ProjectID",
		Description:    `projectID contains your verified accound id.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "projectID",
		Orderable:      true,
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
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}
