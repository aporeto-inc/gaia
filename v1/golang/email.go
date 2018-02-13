package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// EmailTypeValue represents the possible values for attribute "type".
type EmailTypeValue string

const (
	// EmailTypeHtml represents the value HTML.
	EmailTypeHtml EmailTypeValue = "HTML"

	// EmailTypePlain represents the value Plain.
	EmailTypePlain EmailTypeValue = "Plain"
)

// EmailIdentity represents the Identity of the object.
var EmailIdentity = elemental.Identity{
	Name:     "email",
	Category: "emails",
	Private:  true,
}

// EmailsList represents a list of Emails
type EmailsList []*Email

// ContentIdentity returns the identity of the objects in the list.
func (o EmailsList) ContentIdentity() elemental.Identity {

	return EmailIdentity
}

// Copy returns a pointer to a copy the EmailsList.
func (o EmailsList) Copy() elemental.ContentIdentifiable {

	copy := append(EmailsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EmailsList.
func (o EmailsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(EmailsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Email))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EmailsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EmailsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o EmailsList) Version() int {

	return 1
}

// Email represents the model of a email
type Email struct {
	// Attachments is a list of attachments to send
	Attachments map[string]string `json:"attachments" bson:"-" mapstructure:"attachments,omitempty"`

	// Bcc represents email that should be in copy but hidden
	Bcc []string `json:"bcc" bson:"-" mapstructure:"bcc,omitempty"`

	// Cc represents the addresses that should be in copy
	Cc []string `json:"cc" bson:"-" mapstructure:"cc,omitempty"`

	// Content of the email to send
	Content string `json:"content" bson:"-" mapstructure:"content,omitempty"`

	// From represents the sender of the email
	From string `json:"from" bson:"-" mapstructure:"from,omitempty"`

	// Subject represents the subject of the email
	Subject string `json:"subject" bson:"-" mapstructure:"subject,omitempty"`

	// To represents receivers of the email
	To []string `json:"to" bson:"-" mapstructure:"to,omitempty"`

	// Type represents the type of the content.
	Type EmailTypeValue `json:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewEmail returns a new *Email
func NewEmail() *Email {

	return &Email{
		ModelVersion: 1,
		Type:         "Plain",
	}
}

// Identity returns the Identity of the object.
func (o *Email) Identity() elemental.Identity {

	return EmailIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Email) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Email) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Email) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Email) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Email) Doc() string {
	return `Email is a message that can be send via email`
}

func (o *Email) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Email) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("from", o.From); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("from", o.From); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"HTML", "Plain"}, false); err != nil {
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
func (*Email) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EmailAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EmailLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Email) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EmailAttributesMap
}

// EmailAttributesMap represents the map of attribute for Email.
var EmailAttributesMap = map[string]elemental.AttributeSpecification{
	"Attachments": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Attachments",
		Description:    `Attachments is a list of attachments to send`,
		Exposed:        true,
		Name:           "attachments",
		SubType:        "list_attachments",
		Type:           "external",
	},
	"Bcc": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Bcc",
		Description:    `Bcc represents email that should be in copy but hidden `,
		Exposed:        true,
		Name:           "bcc",
		SubType:        "list_emails",
		Type:           "external",
	},
	"Cc": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Cc",
		Description:    `Cc represents the addresses that should be in copy`,
		Exposed:        true,
		Name:           "cc",
		SubType:        "list_emails",
		Type:           "external",
	},
	"Content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		Description:    `Content of the email to send`,
		Exposed:        true,
		Format:         "free",
		Name:           "content",
		Type:           "string",
	},
	"From": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "From",
		Description:    `From represents the sender of the email`,
		Exposed:        true,
		Format:         "email",
		Name:           "from",
		Required:       true,
		Type:           "string",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `Subject represents the subject of the email`,
		Exposed:        true,
		Format:         "free",
		Name:           "subject",
		Type:           "string",
	},
	"To": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "To",
		Description:    `To represents receivers of the email `,
		Exposed:        true,
		Name:           "to",
		SubType:        "list_emails",
		Type:           "external",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"HTML", "Plain"},
		ConvertedName:  "Type",
		DefaultValue:   EmailTypePlain,
		Description:    `Type represents the type of the content.`,
		Exposed:        true,
		Name:           "type",
		Type:           "enum",
	},
}

// EmailLowerCaseAttributesMap represents the map of attribute for Email.
var EmailLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"attachments": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Attachments",
		Description:    `Attachments is a list of attachments to send`,
		Exposed:        true,
		Name:           "attachments",
		SubType:        "list_attachments",
		Type:           "external",
	},
	"bcc": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Bcc",
		Description:    `Bcc represents email that should be in copy but hidden `,
		Exposed:        true,
		Name:           "bcc",
		SubType:        "list_emails",
		Type:           "external",
	},
	"cc": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Cc",
		Description:    `Cc represents the addresses that should be in copy`,
		Exposed:        true,
		Name:           "cc",
		SubType:        "list_emails",
		Type:           "external",
	},
	"content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		Description:    `Content of the email to send`,
		Exposed:        true,
		Format:         "free",
		Name:           "content",
		Type:           "string",
	},
	"from": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "From",
		Description:    `From represents the sender of the email`,
		Exposed:        true,
		Format:         "email",
		Name:           "from",
		Required:       true,
		Type:           "string",
	},
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `Subject represents the subject of the email`,
		Exposed:        true,
		Format:         "free",
		Name:           "subject",
		Type:           "string",
	},
	"to": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "To",
		Description:    `To represents receivers of the email `,
		Exposed:        true,
		Name:           "to",
		SubType:        "list_emails",
		Type:           "external",
	},
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{"HTML", "Plain"},
		ConvertedName:  "Type",
		DefaultValue:   EmailTypePlain,
		Description:    `Type represents the type of the content.`,
		Exposed:        true,
		Name:           "type",
		Type:           "enum",
	},
}
