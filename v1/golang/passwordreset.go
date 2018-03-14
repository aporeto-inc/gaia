package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// PasswordResetIdentity represents the Identity of the object.
var PasswordResetIdentity = elemental.Identity{
	Name:     "passwordreset",
	Category: "passwordreset",
	Private:  false,
}

// PasswordResetsList represents a list of PasswordResets
type PasswordResetsList []*PasswordReset

// ContentIdentity returns the identity of the objects in the list.
func (o PasswordResetsList) ContentIdentity() elemental.Identity {

	return PasswordResetIdentity
}

// Copy returns a pointer to a copy the PasswordResetsList.
func (o PasswordResetsList) Copy() elemental.ContentIdentifiable {

	copy := append(PasswordResetsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PasswordResetsList.
func (o PasswordResetsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(PasswordResetsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PasswordReset))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PasswordResetsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PasswordResetsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o PasswordResetsList) Version() int {

	return 1
}

// PasswordReset represents the model of a passwordreset
type PasswordReset struct {
	// Password contains the new password.
	Password string `json:"password" bson:"-" mapstructure:"password,omitempty"`

	// Token contains the reset password token
	Token string `json:"token" bson:"-" mapstructure:"token,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewPasswordReset returns a new *PasswordReset
func NewPasswordReset() *PasswordReset {

	return &PasswordReset{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PasswordReset) Identity() elemental.Identity {

	return PasswordResetIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PasswordReset) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PasswordReset) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *PasswordReset) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *PasswordReset) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PasswordReset) Doc() string {
	return `Used to reset an account password.`
}

func (o *PasswordReset) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *PasswordReset) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("password", o.Password); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("password", o.Password); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("token", o.Token); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("token", o.Token); err != nil {
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
func (*PasswordReset) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PasswordResetAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PasswordResetLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PasswordReset) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PasswordResetAttributesMap
}

// PasswordResetAttributesMap represents the map of attribute for PasswordReset.
var PasswordResetAttributesMap = map[string]elemental.AttributeSpecification{
	"Password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Password",
		Description:    `Password contains the new password.`,
		Exposed:        true,
		Format:         "free",
		Name:           "password",
		Required:       true,
		Type:           "string",
	},
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `Token contains the reset password token`,
		Exposed:        true,
		Format:         "free",
		Name:           "token",
		Required:       true,
		Type:           "string",
	},
}

// PasswordResetLowerCaseAttributesMap represents the map of attribute for PasswordReset.
var PasswordResetLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Password",
		Description:    `Password contains the new password.`,
		Exposed:        true,
		Format:         "free",
		Name:           "password",
		Required:       true,
		Type:           "string",
	},
	"token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `Token contains the reset password token`,
		Exposed:        true,
		Format:         "free",
		Name:           "token",
		Required:       true,
		Type:           "string",
	},
}
