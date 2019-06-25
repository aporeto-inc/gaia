package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// APICheckOperationValue represents the possible values for attribute "operation".
type APICheckOperationValue string

const (
	// APICheckOperationCreate represents the value Create.
	APICheckOperationCreate APICheckOperationValue = "Create"

	// APICheckOperationDelete represents the value Delete.
	APICheckOperationDelete APICheckOperationValue = "Delete"

	// APICheckOperationInfo represents the value Info.
	APICheckOperationInfo APICheckOperationValue = "Info"

	// APICheckOperationPatch represents the value Patch.
	APICheckOperationPatch APICheckOperationValue = "Patch"

	// APICheckOperationRetrieve represents the value Retrieve.
	APICheckOperationRetrieve APICheckOperationValue = "Retrieve"

	// APICheckOperationRetrieveMany represents the value RetrieveMany.
	APICheckOperationRetrieveMany APICheckOperationValue = "RetrieveMany"

	// APICheckOperationUpdate represents the value Update.
	APICheckOperationUpdate APICheckOperationValue = "Update"
)

// APICheckIdentity represents the Identity of the object.
var APICheckIdentity = elemental.Identity{
	Name:     "apicheck",
	Category: "apichecks",
	Package:  "squall",
	Private:  false,
}

// APIChecksList represents a list of APIChecks
type APIChecksList []*APICheck

// Identity returns the identity of the objects in the list.
func (o APIChecksList) Identity() elemental.Identity {

	return APICheckIdentity
}

// Copy returns a pointer to a copy the APIChecksList.
func (o APIChecksList) Copy() elemental.Identifiables {

	copy := append(APIChecksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the APIChecksList.
func (o APIChecksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(APIChecksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*APICheck))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o APIChecksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o APIChecksList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the APIChecksList converted to SparseAPIChecksList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o APIChecksList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAPIChecksList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAPICheck)
	}

	return out
}

// Version returns the version of the content.
func (o APIChecksList) Version() int {

	return 1
}

// APICheck represents the model of a apicheck
type APICheck struct {
	// Authorized contains the results of the check.
	Authorized map[string]bool `json:"authorized" msgpack:"authorized" bson:"-" mapstructure:"authorized,omitempty"`

	// Claims contains the decoded claims used.
	Claims []string `json:"claims" msgpack:"claims" bson:"-" mapstructure:"claims,omitempty"`

	// Namespace is the namespace to use to check the api authentication.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// Operation is the operation you want to check.
	Operation APICheckOperationValue `json:"operation" msgpack:"operation" bson:"operation" mapstructure:"operation,omitempty"`

	// TargetIdentities contains the list of identities you want to check the
	// authorization.
	TargetIdentities []string `json:"targetIdentities" msgpack:"targetIdentities" bson:"-" mapstructure:"targetIdentities,omitempty"`

	// Token is the token to use to check api authentication.
	Token string `json:"token" msgpack:"token" bson:"-" mapstructure:"token,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAPICheck returns a new *APICheck
func NewAPICheck() *APICheck {

	return &APICheck{
		ModelVersion:     1,
		Authorized:       map[string]bool{},
		Claims:           []string{},
		TargetIdentities: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *APICheck) Identity() elemental.Identity {

	return APICheckIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *APICheck) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *APICheck) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *APICheck) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *APICheck) BleveType() string {

	return "apicheck"
}

// DefaultOrder returns the list of default ordering fields.
func (o *APICheck) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *APICheck) Doc() string {

	return `This API allows to verify is a client identitied by his token is allowed to do
some operations on some apis. For example, it allows third party system to
impersonate a user and ensure a proxfied request should be allowed.`
}

func (o *APICheck) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *APICheck) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAPICheck{
			Authorized:       &o.Authorized,
			Claims:           &o.Claims,
			Namespace:        &o.Namespace,
			Operation:        &o.Operation,
			TargetIdentities: &o.TargetIdentities,
			Token:            &o.Token,
		}
	}

	sp := &SparseAPICheck{}
	for _, f := range fields {
		switch f {
		case "authorized":
			sp.Authorized = &(o.Authorized)
		case "claims":
			sp.Claims = &(o.Claims)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "operation":
			sp.Operation = &(o.Operation)
		case "targetIdentities":
			sp.TargetIdentities = &(o.TargetIdentities)
		case "token":
			sp.Token = &(o.Token)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAPICheck to the object.
func (o *APICheck) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAPICheck)
	if so.Authorized != nil {
		o.Authorized = *so.Authorized
	}
	if so.Claims != nil {
		o.Claims = *so.Claims
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Operation != nil {
		o.Operation = *so.Operation
	}
	if so.TargetIdentities != nil {
		o.TargetIdentities = *so.TargetIdentities
	}
	if so.Token != nil {
		o.Token = *so.Token
	}
}

// DeepCopy returns a deep copy if the APICheck.
func (o *APICheck) DeepCopy() *APICheck {

	if o == nil {
		return nil
	}

	out := &APICheck{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *APICheck.
func (o *APICheck) DeepCopyInto(out *APICheck) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy APICheck: %s", err))
	}

	*out = *target.(*APICheck)
}

// Validate valides the current information stored into the structure.
func (o *APICheck) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("operation", string(o.Operation), []string{"Create", "Delete", "Info", "Patch", "Retrieve", "RetrieveMany", "Update"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredExternal("targetIdentities", o.TargetIdentities); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("token", o.Token); err != nil {
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
func (*APICheck) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := APICheckAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return APICheckLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*APICheck) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return APICheckAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *APICheck) ValueForAttribute(name string) interface{} {

	switch name {
	case "authorized":
		return o.Authorized
	case "claims":
		return o.Claims
	case "namespace":
		return o.Namespace
	case "operation":
		return o.Operation
	case "targetIdentities":
		return o.TargetIdentities
	case "token":
		return o.Token
	}

	return nil
}

// APICheckAttributesMap represents the map of attribute for APICheck.
var APICheckAttributesMap = map[string]elemental.AttributeSpecification{
	"Authorized": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Authorized",
		Description:    `Authorized contains the results of the check.`,
		Exposed:        true,
		Name:           "authorized",
		ReadOnly:       true,
		SubType:        "map[string]bool",
		Type:           "external",
	},
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Claims contains the decoded claims used.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace is the namespace to use to check the api authentication.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"Operation": elemental.AttributeSpecification{
		AllowedChoices: []string{"Create", "Delete", "Info", "Patch", "Retrieve", "RetrieveMany", "Update"},
		ConvertedName:  "Operation",
		Description:    `Operation is the operation you want to check.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operation",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"TargetIdentities": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentities",
		Description: `TargetIdentities contains the list of identities you want to check the
authorization.`,
		Exposed:  true,
		Name:     "targetIdentities",
		Required: true,
		SubType:  "string",
		Type:     "list",
	},
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `Token is the token to use to check api authentication.`,
		Exposed:        true,
		Name:           "token",
		Required:       true,
		Type:           "string",
	},
}

// APICheckLowerCaseAttributesMap represents the map of attribute for APICheck.
var APICheckLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"authorized": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Authorized",
		Description:    `Authorized contains the results of the check.`,
		Exposed:        true,
		Name:           "authorized",
		ReadOnly:       true,
		SubType:        "map[string]bool",
		Type:           "external",
	},
	"claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Claims contains the decoded claims used.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace is the namespace to use to check the api authentication.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"operation": elemental.AttributeSpecification{
		AllowedChoices: []string{"Create", "Delete", "Info", "Patch", "Retrieve", "RetrieveMany", "Update"},
		ConvertedName:  "Operation",
		Description:    `Operation is the operation you want to check.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operation",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"targetidentities": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentities",
		Description: `TargetIdentities contains the list of identities you want to check the
authorization.`,
		Exposed:  true,
		Name:     "targetIdentities",
		Required: true,
		SubType:  "string",
		Type:     "list",
	},
	"token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `Token is the token to use to check api authentication.`,
		Exposed:        true,
		Name:           "token",
		Required:       true,
		Type:           "string",
	},
}

// SparseAPIChecksList represents a list of SparseAPIChecks
type SparseAPIChecksList []*SparseAPICheck

// Identity returns the identity of the objects in the list.
func (o SparseAPIChecksList) Identity() elemental.Identity {

	return APICheckIdentity
}

// Copy returns a pointer to a copy the SparseAPIChecksList.
func (o SparseAPIChecksList) Copy() elemental.Identifiables {

	copy := append(SparseAPIChecksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAPIChecksList.
func (o SparseAPIChecksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAPIChecksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAPICheck))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAPIChecksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAPIChecksList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAPIChecksList converted to APIChecksList.
func (o SparseAPIChecksList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAPIChecksList) Version() int {

	return 1
}

// SparseAPICheck represents the sparse version of a apicheck.
type SparseAPICheck struct {
	// Authorized contains the results of the check.
	Authorized *map[string]bool `json:"authorized,omitempty" msgpack:"authorized,omitempty" bson:"-" mapstructure:"authorized,omitempty"`

	// Claims contains the decoded claims used.
	Claims *[]string `json:"claims,omitempty" msgpack:"claims,omitempty" bson:"-" mapstructure:"claims,omitempty"`

	// Namespace is the namespace to use to check the api authentication.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// Operation is the operation you want to check.
	Operation *APICheckOperationValue `json:"operation,omitempty" msgpack:"operation,omitempty" bson:"operation,omitempty" mapstructure:"operation,omitempty"`

	// TargetIdentities contains the list of identities you want to check the
	// authorization.
	TargetIdentities *[]string `json:"targetIdentities,omitempty" msgpack:"targetIdentities,omitempty" bson:"-" mapstructure:"targetIdentities,omitempty"`

	// Token is the token to use to check api authentication.
	Token *string `json:"token,omitempty" msgpack:"token,omitempty" bson:"-" mapstructure:"token,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAPICheck returns a new  SparseAPICheck.
func NewSparseAPICheck() *SparseAPICheck {
	return &SparseAPICheck{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAPICheck) Identity() elemental.Identity {

	return APICheckIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAPICheck) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAPICheck) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseAPICheck) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAPICheck) ToPlain() elemental.PlainIdentifiable {

	out := NewAPICheck()
	if o.Authorized != nil {
		out.Authorized = *o.Authorized
	}
	if o.Claims != nil {
		out.Claims = *o.Claims
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Operation != nil {
		out.Operation = *o.Operation
	}
	if o.TargetIdentities != nil {
		out.TargetIdentities = *o.TargetIdentities
	}
	if o.Token != nil {
		out.Token = *o.Token
	}

	return out
}

// DeepCopy returns a deep copy if the SparseAPICheck.
func (o *SparseAPICheck) DeepCopy() *SparseAPICheck {

	if o == nil {
		return nil
	}

	out := &SparseAPICheck{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAPICheck.
func (o *SparseAPICheck) DeepCopyInto(out *SparseAPICheck) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAPICheck: %s", err))
	}

	*out = *target.(*SparseAPICheck)
}
