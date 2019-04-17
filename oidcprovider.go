package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
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

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o OIDCProvidersList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the OIDCProvidersList converted to SparseOIDCProvidersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o OIDCProvidersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseOIDCProvidersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseOIDCProvider)
	}

	return out
}

// Version returns the version of the content.
func (o OIDCProvidersList) Version() int {

	return 1
}

// OIDCProvider represents the model of a oidcprovider
type OIDCProvider struct {
	// ID is the identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Unique client ID.
	ClientID string `json:"clientID" msgpack:"clientID" bson:"clientid" mapstructure:"clientID,omitempty"`

	// Client secret associated with the client ID.
	ClientSecret string `json:"clientSecret" msgpack:"clientSecret" bson:"clientsecret" mapstructure:"clientSecret,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// If set, this will be the default OIDCProvider. There can be only one default
	// provider in your account. When logging in with OIDC, if not provider name is
	// given, the default will be used.
	Default bool `json:"default" msgpack:"default" bson:"default" mapstructure:"default,omitempty"`

	// OIDC information endpoint.
	Endpoint string `json:"endpoint" msgpack:"endpoint" bson:"endpoint" mapstructure:"endpoint,omitempty"`

	// Name of the provider.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// ParentID contains the parent Vince account ID.
	ParentID string `json:"parentID" msgpack:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// ParentName contains the name of the Vince parent Account.
	ParentName string `json:"parentName" msgpack:"parentName" bson:"parentname" mapstructure:"parentName,omitempty"`

	// List of scopes to allow.
	Scopes []string `json:"scopes" msgpack:"scopes" bson:"scopes" mapstructure:"scopes,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewOIDCProvider returns a new *OIDCProvider
func NewOIDCProvider() *OIDCProvider {

	return &OIDCProvider{
		ModelVersion: 1,
		Scopes:       []string{},
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

// GetCreateTime returns the CreateTime of the receiver.
func (o *OIDCProvider) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *OIDCProvider) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *OIDCProvider) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *OIDCProvider) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *OIDCProvider) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseOIDCProvider{
			ID:           &o.ID,
			ClientID:     &o.ClientID,
			ClientSecret: &o.ClientSecret,
			CreateTime:   &o.CreateTime,
			Default:      &o.Default,
			Endpoint:     &o.Endpoint,
			Name:         &o.Name,
			ParentID:     &o.ParentID,
			ParentName:   &o.ParentName,
			Scopes:       &o.Scopes,
			UpdateTime:   &o.UpdateTime,
		}
	}

	sp := &SparseOIDCProvider{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "clientID":
			sp.ClientID = &(o.ClientID)
		case "clientSecret":
			sp.ClientSecret = &(o.ClientSecret)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "default":
			sp.Default = &(o.Default)
		case "endpoint":
			sp.Endpoint = &(o.Endpoint)
		case "name":
			sp.Name = &(o.Name)
		case "parentID":
			sp.ParentID = &(o.ParentID)
		case "parentName":
			sp.ParentName = &(o.ParentName)
		case "scopes":
			sp.Scopes = &(o.Scopes)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseOIDCProvider to the object.
func (o *OIDCProvider) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseOIDCProvider)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.ClientID != nil {
		o.ClientID = *so.ClientID
	}
	if so.ClientSecret != nil {
		o.ClientSecret = *so.ClientSecret
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Default != nil {
		o.Default = *so.Default
	}
	if so.Endpoint != nil {
		o.Endpoint = *so.Endpoint
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.ParentID != nil {
		o.ParentID = *so.ParentID
	}
	if so.ParentName != nil {
		o.ParentName = *so.ParentName
	}
	if so.Scopes != nil {
		o.Scopes = *so.Scopes
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the OIDCProvider.
func (o *OIDCProvider) DeepCopy() *OIDCProvider {

	if o == nil {
		return nil
	}

	out := &OIDCProvider{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *OIDCProvider.
func (o *OIDCProvider) DeepCopyInto(out *OIDCProvider) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy OIDCProvider: %s", err))
	}

	*out = *target.(*OIDCProvider)
}

// Validate valides the current information stored into the structure.
func (o *OIDCProvider) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("clientID", o.ClientID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("clientSecret", o.ClientSecret); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("endpoint", o.Endpoint); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *OIDCProvider) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "clientID":
		return o.ClientID
	case "clientSecret":
		return o.ClientSecret
	case "createTime":
		return o.CreateTime
	case "default":
		return o.Default
	case "endpoint":
		return o.Endpoint
	case "name":
		return o.Name
	case "parentID":
		return o.ParentID
	case "parentName":
		return o.ParentName
	case "scopes":
		return o.Scopes
	case "updateTime":
		return o.UpdateTime
	}

	return nil
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
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Default": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Default",
		Description: `If set, this will be the default OIDCProvider. There can be only one default
provider in your account. When logging in with OIDC, if not provider name is
given, the default will be used.`,
		Exposed: true,
		Name:    "default",
		Stored:  true,
		Type:    "boolean",
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
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
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
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"default": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Default",
		Description: `If set, this will be the default OIDCProvider. There can be only one default
provider in your account. When logging in with OIDC, if not provider name is
given, the default will be used.`,
		Exposed: true,
		Name:    "default",
		Stored:  true,
		Type:    "boolean",
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
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}

// SparseOIDCProvidersList represents a list of SparseOIDCProviders
type SparseOIDCProvidersList []*SparseOIDCProvider

// Identity returns the identity of the objects in the list.
func (o SparseOIDCProvidersList) Identity() elemental.Identity {

	return OIDCProviderIdentity
}

// Copy returns a pointer to a copy the SparseOIDCProvidersList.
func (o SparseOIDCProvidersList) Copy() elemental.Identifiables {

	copy := append(SparseOIDCProvidersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseOIDCProvidersList.
func (o SparseOIDCProvidersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseOIDCProvidersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseOIDCProvider))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseOIDCProvidersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseOIDCProvidersList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseOIDCProvidersList converted to OIDCProvidersList.
func (o SparseOIDCProvidersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseOIDCProvidersList) Version() int {

	return 1
}

// SparseOIDCProvider represents the sparse version of a oidcprovider.
type SparseOIDCProvider struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Unique client ID.
	ClientID *string `json:"clientID,omitempty" msgpack:"clientID,omitempty" bson:"clientid,omitempty" mapstructure:"clientID,omitempty"`

	// Client secret associated with the client ID.
	ClientSecret *string `json:"clientSecret,omitempty" msgpack:"clientSecret,omitempty" bson:"clientsecret,omitempty" mapstructure:"clientSecret,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// If set, this will be the default OIDCProvider. There can be only one default
	// provider in your account. When logging in with OIDC, if not provider name is
	// given, the default will be used.
	Default *bool `json:"default,omitempty" msgpack:"default,omitempty" bson:"default,omitempty" mapstructure:"default,omitempty"`

	// OIDC information endpoint.
	Endpoint *string `json:"endpoint,omitempty" msgpack:"endpoint,omitempty" bson:"endpoint,omitempty" mapstructure:"endpoint,omitempty"`

	// Name of the provider.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// ParentID contains the parent Vince account ID.
	ParentID *string `json:"parentID,omitempty" msgpack:"parentID,omitempty" bson:"parentid,omitempty" mapstructure:"parentID,omitempty"`

	// ParentName contains the name of the Vince parent Account.
	ParentName *string `json:"parentName,omitempty" msgpack:"parentName,omitempty" bson:"parentname,omitempty" mapstructure:"parentName,omitempty"`

	// List of scopes to allow.
	Scopes *[]string `json:"scopes,omitempty" msgpack:"scopes,omitempty" bson:"scopes,omitempty" mapstructure:"scopes,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseOIDCProvider returns a new  SparseOIDCProvider.
func NewSparseOIDCProvider() *SparseOIDCProvider {
	return &SparseOIDCProvider{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseOIDCProvider) Identity() elemental.Identity {

	return OIDCProviderIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseOIDCProvider) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseOIDCProvider) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseOIDCProvider) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseOIDCProvider) ToPlain() elemental.PlainIdentifiable {

	out := NewOIDCProvider()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.ClientID != nil {
		out.ClientID = *o.ClientID
	}
	if o.ClientSecret != nil {
		out.ClientSecret = *o.ClientSecret
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Default != nil {
		out.Default = *o.Default
	}
	if o.Endpoint != nil {
		out.Endpoint = *o.Endpoint
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.ParentID != nil {
		out.ParentID = *o.ParentID
	}
	if o.ParentName != nil {
		out.ParentName = *o.ParentName
	}
	if o.Scopes != nil {
		out.Scopes = *o.Scopes
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseOIDCProvider) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseOIDCProvider) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseOIDCProvider) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseOIDCProvider) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseOIDCProvider.
func (o *SparseOIDCProvider) DeepCopy() *SparseOIDCProvider {

	if o == nil {
		return nil
	}

	out := &SparseOIDCProvider{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseOIDCProvider.
func (o *SparseOIDCProvider) DeepCopyInto(out *SparseOIDCProvider) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseOIDCProvider: %s", err))
	}

	*out = *target.(*SparseOIDCProvider)
}
