package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// OIDCProviderIdentity represents the Identity of the object.
var OIDCProviderIdentity = elemental.Identity{
	Name:     "oidcprovider",
	Category: "oidcproviders",
	Package:  "vince",
	Private:  false,
}

// OIDCProvidersList represents a list of OIDCProviders
type OIDCProvidersList []*OIDCProvider

// Identity returns the identity of the objects in the list.
func (o OIDCProvidersList) Identity() elemental.Identity {

	return OIDCProviderIdentity
}

// Copy returns a pointer to a copy the OIDCProvidersList.
func (o OIDCProvidersList) Copy() elemental.Identifiables {

	copy := append(OIDCProvidersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the OIDCProvidersList.
func (o OIDCProvidersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(OIDCProvidersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*OIDCProvider))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o OIDCProvidersList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o OIDCProvidersList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o OIDCProvidersList) Version() int {

	return 1
}

// OIDCProvider represents the model of a oidcprovider
type OIDCProvider struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Unique client ID.
	ClientID string `json:"clientID" bson:"clientid" mapstructure:"clientID,omitempty"`

	// Client secret associated with the client ID.
	ClientSecret string `json:"clientSecret" bson:"clientsecret" mapstructure:"clientSecret,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// OIDC information endpoint.
	Endpoint string `json:"endpoint" bson:"endpoint" mapstructure:"endpoint,omitempty"`

	// Name of the provider.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// ParentID contains the parent Vince account ID.
	ParentID string `json:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// ParentName contains the name of the Vince parent Account.
	ParentName string `json:"parentName" bson:"parentname" mapstructure:"parentName,omitempty"`

	// List of scopes to allow.
	Scopes []string `json:"scopes" bson:"scopes" mapstructure:"scopes,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewOIDCProvider returns a new *OIDCProvider
func NewOIDCProvider() *OIDCProvider {

	return &OIDCProvider{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *OIDCProvider) Identity() elemental.Identity {

	return OIDCProviderIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *OIDCProvider) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *OIDCProvider) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *OIDCProvider) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *OIDCProvider) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *OIDCProvider) Doc() string {
	return `Allows to declare a generic OpenID Connect provider that can be used in exchange
for a Midgard token.`
}

func (o *OIDCProvider) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *OIDCProvider) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("clientID", o.ClientID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("clientSecret", o.ClientSecret); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("endpoint", o.Endpoint); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
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
func (*OIDCProvider) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := OIDCProviderAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return OIDCProviderLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*OIDCProvider) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return OIDCProviderAttributesMap
}

// OIDCProviderAttributesMap represents the map of attribute for OIDCProvider.
var OIDCProviderAttributesMap = map[string]elemental.AttributeSpecification{
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
	"ClientID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientID",
		Description:    `Unique client ID.`,
		Exposed:        true,
		Name:           "clientID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ClientSecret": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientSecret",
		Description:    `Client secret associated with the client ID.`,
		Exposed:        true,
		Name:           "clientSecret",
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
	"Endpoint": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Endpoint",
		Description:    `OIDC information endpoint.`,
		Exposed:        true,
		Name:           "endpoint",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		Description:    `Name of the provider.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Stored:         true,
		Type:           "string",
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
	"Scopes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Scopes",
		Description:    `List of scopes to allow.`,
		Exposed:        true,
		Name:           "scopes",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
}

// OIDCProviderLowerCaseAttributesMap represents the map of attribute for OIDCProvider.
var OIDCProviderLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"clientid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientID",
		Description:    `Unique client ID.`,
		Exposed:        true,
		Name:           "clientID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"clientsecret": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientSecret",
		Description:    `Client secret associated with the client ID.`,
		Exposed:        true,
		Name:           "clientSecret",
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
	"endpoint": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Endpoint",
		Description:    `OIDC information endpoint.`,
		Exposed:        true,
		Name:           "endpoint",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		Description:    `Name of the provider.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Stored:         true,
		Type:           "string",
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
	"scopes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Scopes",
		Description:    `List of scopes to allow.`,
		Exposed:        true,
		Name:           "scopes",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
}
