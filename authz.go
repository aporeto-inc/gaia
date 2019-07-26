package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AuthzTargetOperationValue represents the possible values for attribute "targetOperation".
type AuthzTargetOperationValue string

const (
	// AuthzTargetOperationAny represents the value Any.
	AuthzTargetOperationAny AuthzTargetOperationValue = "Any"

	// AuthzTargetOperationCreate represents the value Create.
	AuthzTargetOperationCreate AuthzTargetOperationValue = "Create"

	// AuthzTargetOperationDelete represents the value Delete.
	AuthzTargetOperationDelete AuthzTargetOperationValue = "Delete"

	// AuthzTargetOperationInfo represents the value Info.
	AuthzTargetOperationInfo AuthzTargetOperationValue = "Info"

	// AuthzTargetOperationPatch represents the value Patch.
	AuthzTargetOperationPatch AuthzTargetOperationValue = "Patch"

	// AuthzTargetOperationRetrieve represents the value Retrieve.
	AuthzTargetOperationRetrieve AuthzTargetOperationValue = "Retrieve"

	// AuthzTargetOperationRetrieveMany represents the value RetrieveMany.
	AuthzTargetOperationRetrieveMany AuthzTargetOperationValue = "RetrieveMany"

	// AuthzTargetOperationUpdate represents the value Update.
	AuthzTargetOperationUpdate AuthzTargetOperationValue = "Update"
)

// AuthzIdentity represents the Identity of the object.
var AuthzIdentity = elemental.Identity{
	Name:     "authz",
	Category: "authz",
	Package:  "cid",
	Private:  true,
}

// AuthzsList represents a list of Authzs
type AuthzsList []*Authz

// Identity returns the identity of the objects in the list.
func (o AuthzsList) Identity() elemental.Identity {

	return AuthzIdentity
}

// Copy returns a pointer to a copy the AuthzsList.
func (o AuthzsList) Copy() elemental.Identifiables {

	copy := append(AuthzsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AuthzsList.
func (o AuthzsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AuthzsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Authz))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AuthzsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AuthzsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AuthzsList converted to SparseAuthzsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AuthzsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAuthzsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAuthz)
	}

	return out
}

// Version returns the version of the content.
func (o AuthzsList) Version() int {

	return 1
}

// Authz represents the model of a authz
type Authz struct {
	// Return if the request should be authorized.
	Authorized bool `json:"authorized,omitempty" msgpack:"authorized,omitempty" bson:"-" mapstructure:"authorized,omitempty"`

	// The list of verified claims.
	Claims []string `json:"claims" msgpack:"claims" bson:"-" mapstructure:"claims,omitempty"`

	// IP of the client.
	ClientIP string `json:"clientIP" msgpack:"clientIP" bson:"-" mapstructure:"clientIP,omitempty"`

	// Return an eventual error.
	Error string `json:"error,omitempty" msgpack:"error,omitempty" bson:"-" mapstructure:"error,omitempty"`

	// If the parameter permissions=1 is set, targetIdentity and targetOperation are
	// ignored and this attribute will contain all the permission for the given claims.
	Permissions map[string]map[string]bool `json:"permissions,omitempty" msgpack:"permissions,omitempty" bson:"-" mapstructure:"permissions,omitempty"`

	// The identity.
	TargetIdentity string `json:"targetIdentity" msgpack:"targetIdentity" bson:"-" mapstructure:"targetIdentity,omitempty"`

	// description.
	TargetNamespace string `json:"targetNamespace" msgpack:"targetNamespace" bson:"-" mapstructure:"targetNamespace,omitempty"`

	// Operation.
	TargetOperation AuthzTargetOperationValue `json:"targetOperation" msgpack:"targetOperation" bson:"-" mapstructure:"targetOperation,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAuthz returns a new *Authz
func NewAuthz() *Authz {

	return &Authz{
		ModelVersion: 1,
		Claims:       []string{},
		Permissions:  map[string]map[string]bool{},
	}
}

// Identity returns the Identity of the object.
func (o *Authz) Identity() elemental.Identity {

	return AuthzIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Authz) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Authz) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Authz) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Authz) BleveType() string {

	return "authz"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Authz) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Authz) Doc() string {

	return `This is an internal api that is used to resolve to api authorization.`
}

func (o *Authz) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Authz) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAuthz{
			Authorized:      &o.Authorized,
			Claims:          &o.Claims,
			ClientIP:        &o.ClientIP,
			Error:           &o.Error,
			Permissions:     &o.Permissions,
			TargetIdentity:  &o.TargetIdentity,
			TargetNamespace: &o.TargetNamespace,
			TargetOperation: &o.TargetOperation,
		}
	}

	sp := &SparseAuthz{}
	for _, f := range fields {
		switch f {
		case "authorized":
			sp.Authorized = &(o.Authorized)
		case "claims":
			sp.Claims = &(o.Claims)
		case "clientIP":
			sp.ClientIP = &(o.ClientIP)
		case "error":
			sp.Error = &(o.Error)
		case "permissions":
			sp.Permissions = &(o.Permissions)
		case "targetIdentity":
			sp.TargetIdentity = &(o.TargetIdentity)
		case "targetNamespace":
			sp.TargetNamespace = &(o.TargetNamespace)
		case "targetOperation":
			sp.TargetOperation = &(o.TargetOperation)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAuthz to the object.
func (o *Authz) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAuthz)
	if so.Authorized != nil {
		o.Authorized = *so.Authorized
	}
	if so.Claims != nil {
		o.Claims = *so.Claims
	}
	if so.ClientIP != nil {
		o.ClientIP = *so.ClientIP
	}
	if so.Error != nil {
		o.Error = *so.Error
	}
	if so.Permissions != nil {
		o.Permissions = *so.Permissions
	}
	if so.TargetIdentity != nil {
		o.TargetIdentity = *so.TargetIdentity
	}
	if so.TargetNamespace != nil {
		o.TargetNamespace = *so.TargetNamespace
	}
	if so.TargetOperation != nil {
		o.TargetOperation = *so.TargetOperation
	}
}

// DeepCopy returns a deep copy if the Authz.
func (o *Authz) DeepCopy() *Authz {

	if o == nil {
		return nil
	}

	out := &Authz{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Authz.
func (o *Authz) DeepCopyInto(out *Authz) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Authz: %s", err))
	}

	*out = *target.(*Authz)
}

// Validate valides the current information stored into the structure.
func (o *Authz) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("claims", o.Claims); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("targetNamespace", o.TargetNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("targetOperation", string(o.TargetOperation), []string{"Any", "Create", "Delete", "Info", "Patch", "Retrieve", "RetrieveMany", "Update"}, false); err != nil {
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
func (*Authz) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AuthzAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AuthzLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Authz) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AuthzAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Authz) ValueForAttribute(name string) interface{} {

	switch name {
	case "authorized":
		return o.Authorized
	case "claims":
		return o.Claims
	case "clientIP":
		return o.ClientIP
	case "error":
		return o.Error
	case "permissions":
		return o.Permissions
	case "targetIdentity":
		return o.TargetIdentity
	case "targetNamespace":
		return o.TargetNamespace
	case "targetOperation":
		return o.TargetOperation
	}

	return nil
}

// AuthzAttributesMap represents the map of attribute for Authz.
var AuthzAttributesMap = map[string]elemental.AttributeSpecification{
	"Authorized": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Authorized",
		Description:    `Return if the request should be authorized.`,
		Exposed:        true,
		Name:           "authorized",
		ReadOnly:       true,
		Type:           "boolean",
	},
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `The list of verified claims.`,
		Exposed:        true,
		Name:           "claims",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"ClientIP": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientIP",
		Description:    `IP of the client.`,
		Exposed:        true,
		Name:           "clientIP",
		Type:           "string",
	},
	"Error": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Error",
		Description:    `Return an eventual error.`,
		Exposed:        true,
		Name:           "error",
		ReadOnly:       true,
		Type:           "string",
	},
	"Permissions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Permissions",
		Description: `If the parameter permissions=1 is set, targetIdentity and targetOperation are
ignored and this attribute will contain all the permission for the given claims.`,
		Exposed:  true,
		Name:     "permissions",
		ReadOnly: true,
		SubType:  "map[string]map[string]bool",
		Type:     "external",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `The identity.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Type:           "string",
	},
	"TargetNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description:    `description.`,
		Exposed:        true,
		Name:           "targetNamespace",
		Required:       true,
		Type:           "string",
	},
	"TargetOperation": elemental.AttributeSpecification{
		AllowedChoices: []string{"Any", "Create", "Delete", "Info", "Patch", "Retrieve", "RetrieveMany", "Update"},
		ConvertedName:  "TargetOperation",
		Description:    `Operation.`,
		Exposed:        true,
		Name:           "targetOperation",
		Type:           "enum",
	},
}

// AuthzLowerCaseAttributesMap represents the map of attribute for Authz.
var AuthzLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"authorized": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Authorized",
		Description:    `Return if the request should be authorized.`,
		Exposed:        true,
		Name:           "authorized",
		ReadOnly:       true,
		Type:           "boolean",
	},
	"claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `The list of verified claims.`,
		Exposed:        true,
		Name:           "claims",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"clientip": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientIP",
		Description:    `IP of the client.`,
		Exposed:        true,
		Name:           "clientIP",
		Type:           "string",
	},
	"error": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Error",
		Description:    `Return an eventual error.`,
		Exposed:        true,
		Name:           "error",
		ReadOnly:       true,
		Type:           "string",
	},
	"permissions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Permissions",
		Description: `If the parameter permissions=1 is set, targetIdentity and targetOperation are
ignored and this attribute will contain all the permission for the given claims.`,
		Exposed:  true,
		Name:     "permissions",
		ReadOnly: true,
		SubType:  "map[string]map[string]bool",
		Type:     "external",
	},
	"targetidentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `The identity.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Type:           "string",
	},
	"targetnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description:    `description.`,
		Exposed:        true,
		Name:           "targetNamespace",
		Required:       true,
		Type:           "string",
	},
	"targetoperation": elemental.AttributeSpecification{
		AllowedChoices: []string{"Any", "Create", "Delete", "Info", "Patch", "Retrieve", "RetrieveMany", "Update"},
		ConvertedName:  "TargetOperation",
		Description:    `Operation.`,
		Exposed:        true,
		Name:           "targetOperation",
		Type:           "enum",
	},
}

// SparseAuthzsList represents a list of SparseAuthzs
type SparseAuthzsList []*SparseAuthz

// Identity returns the identity of the objects in the list.
func (o SparseAuthzsList) Identity() elemental.Identity {

	return AuthzIdentity
}

// Copy returns a pointer to a copy the SparseAuthzsList.
func (o SparseAuthzsList) Copy() elemental.Identifiables {

	copy := append(SparseAuthzsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAuthzsList.
func (o SparseAuthzsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAuthzsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAuthz))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAuthzsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAuthzsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAuthzsList converted to AuthzsList.
func (o SparseAuthzsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAuthzsList) Version() int {

	return 1
}

// SparseAuthz represents the sparse version of a authz.
type SparseAuthz struct {
	// Return if the request should be authorized.
	Authorized *bool `json:"authorized,omitempty" msgpack:"authorized,omitempty" bson:"-" mapstructure:"authorized,omitempty"`

	// The list of verified claims.
	Claims *[]string `json:"claims,omitempty" msgpack:"claims,omitempty" bson:"-" mapstructure:"claims,omitempty"`

	// IP of the client.
	ClientIP *string `json:"clientIP,omitempty" msgpack:"clientIP,omitempty" bson:"-" mapstructure:"clientIP,omitempty"`

	// Return an eventual error.
	Error *string `json:"error,omitempty" msgpack:"error,omitempty" bson:"-" mapstructure:"error,omitempty"`

	// If the parameter permissions=1 is set, targetIdentity and targetOperation are
	// ignored and this attribute will contain all the permission for the given claims.
	Permissions *map[string]map[string]bool `json:"permissions,omitempty" msgpack:"permissions,omitempty" bson:"-" mapstructure:"permissions,omitempty"`

	// The identity.
	TargetIdentity *string `json:"targetIdentity,omitempty" msgpack:"targetIdentity,omitempty" bson:"-" mapstructure:"targetIdentity,omitempty"`

	// description.
	TargetNamespace *string `json:"targetNamespace,omitempty" msgpack:"targetNamespace,omitempty" bson:"-" mapstructure:"targetNamespace,omitempty"`

	// Operation.
	TargetOperation *AuthzTargetOperationValue `json:"targetOperation,omitempty" msgpack:"targetOperation,omitempty" bson:"-" mapstructure:"targetOperation,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAuthz returns a new  SparseAuthz.
func NewSparseAuthz() *SparseAuthz {
	return &SparseAuthz{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAuthz) Identity() elemental.Identity {

	return AuthzIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAuthz) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAuthz) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseAuthz) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAuthz) ToPlain() elemental.PlainIdentifiable {

	out := NewAuthz()
	if o.Authorized != nil {
		out.Authorized = *o.Authorized
	}
	if o.Claims != nil {
		out.Claims = *o.Claims
	}
	if o.ClientIP != nil {
		out.ClientIP = *o.ClientIP
	}
	if o.Error != nil {
		out.Error = *o.Error
	}
	if o.Permissions != nil {
		out.Permissions = *o.Permissions
	}
	if o.TargetIdentity != nil {
		out.TargetIdentity = *o.TargetIdentity
	}
	if o.TargetNamespace != nil {
		out.TargetNamespace = *o.TargetNamespace
	}
	if o.TargetOperation != nil {
		out.TargetOperation = *o.TargetOperation
	}

	return out
}

// DeepCopy returns a deep copy if the SparseAuthz.
func (o *SparseAuthz) DeepCopy() *SparseAuthz {

	if o == nil {
		return nil
	}

	out := &SparseAuthz{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAuthz.
func (o *SparseAuthz) DeepCopyInto(out *SparseAuthz) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAuthz: %s", err))
	}

	*out = *target.(*SparseAuthz)
}
