package vincemodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// PasswordResetIdentity represents the Identity of the object
var PasswordResetIdentity = elemental.Identity{
	Name:     "passwordreset",
	Category: "passwordreset",
}

// PasswordResetsList represents a list of PasswordResets
type PasswordResetsList []*PasswordReset

// ContentIdentity returns the identity of the objects in the list.
func (o PasswordResetsList) ContentIdentity() elemental.Identity {

	return PasswordResetIdentity
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

// PasswordReset represents the model of a passwordreset
type PasswordReset struct {
	// Password contains the new password.
	Password string `json:"password" bson:"-"`

	// Token contains the reset password token
	Token string `json:"token" bson:"-"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewPasswordReset returns a new *PasswordReset
func NewPasswordReset() *PasswordReset {

	return &PasswordReset{
		ModelVersion: 1.0,
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
func (o *PasswordReset) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *PasswordReset) Version() float64 {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *PasswordReset) DefaultOrder() []string {

	return []string{}
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

	return PasswordResetAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PasswordReset) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PasswordResetAttributesMap
}

// PasswordResetAttributesMap represents the map of attribute for PasswordReset.
var PasswordResetAttributesMap = map[string]elemental.AttributeSpecification{
	"Password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Password contains the new password.`,
		Exposed:        true,
		Format:         "free",
		Name:           "password",
		Required:       true,
		Type:           "string",
	},
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Token contains the reset password token`,
		Exposed:        true,
		Format:         "free",
		Name:           "token",
		Required:       true,
		Type:           "string",
	},
}
