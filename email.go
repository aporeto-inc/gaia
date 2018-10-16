package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
)

// EmailTypeValue represents the possible values for attribute "type".
type EmailTypeValue string

const (
	// EmailTypeHTML represents the value HTML.
	EmailTypeHTML EmailTypeValue = "HTML"

	// EmailTypePlain represents the value Plain.
	EmailTypePlain EmailTypeValue = "Plain"
)

// EmailIdentity represents the Identity of the object.
var EmailIdentity = elemental.Identity{
	Name:     "email",
	Category: "emails",
	Package:  "yuffie",
	Private:  true,
}

// EmailsList represents a list of Emails
type EmailsList []*Email

// Identity returns the identity of the objects in the list.
func (o EmailsList) Identity() elemental.Identity {

	return EmailIdentity
}

// Copy returns a pointer to a copy the EmailsList.
func (o EmailsList) Copy() elemental.Identifiables {

	copy := append(EmailsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EmailsList.
func (o EmailsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EmailsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Email))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EmailsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EmailsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the EmailsList converted to SparseEmailsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EmailsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o EmailsList) Version() int {

	return 1
}

// Email represents the model of a email
type Email struct {
	// Attachments is a list of attachments to send.
	Attachments map[string]string `json:"attachments" bson:"-" mapstructure:"attachments,omitempty"`

	// Bcc represents email that should be in copy but hidden.
	Bcc []string `json:"bcc" bson:"-" mapstructure:"bcc,omitempty"`

	// Cc represents the addresses that should be in copy.
	Cc []string `json:"cc" bson:"-" mapstructure:"cc,omitempty"`

	// Content of the email to send.
	Content string `json:"content" bson:"-" mapstructure:"content,omitempty"`

	// From represents the sender of the email.
	From string `json:"from" bson:"-" mapstructure:"from,omitempty"`

	// Subject represents the subject of the email.
	Subject string `json:"subject" bson:"-" mapstructure:"subject,omitempty"`

	// To represents receivers of the email.
	To []string `json:"to" bson:"-" mapstructure:"to,omitempty"`

	// Type represents the type of the content.
	Type EmailTypeValue `json:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewEmail returns a new *Email
func NewEmail() *Email {

	return &Email{
		ModelVersion: 1,
		Type:         EmailTypePlain,
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
	return `Email is a message that can be send via email.`
}

func (o *Email) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Email) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEmail{
			Attachments: &o.Attachments,
			Bcc:         &o.Bcc,
			Cc:          &o.Cc,
			Content:     &o.Content,
			From:        &o.From,
			Subject:     &o.Subject,
			To:          &o.To,
			Type:        &o.Type,
		}
	}

	sp := &SparseEmail{}
	for _, f := range fields {
		switch f {
		case "attachments":
			sp.Attachments = &(o.Attachments)
		case "bcc":
			sp.Bcc = &(o.Bcc)
		case "cc":
			sp.Cc = &(o.Cc)
		case "content":
			sp.Content = &(o.Content)
		case "from":
			sp.From = &(o.From)
		case "subject":
			sp.Subject = &(o.Subject)
		case "to":
			sp.To = &(o.To)
		case "type":
			sp.Type = &(o.Type)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseEmail to the object.
func (o *Email) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEmail)
	if so.Attachments != nil {
		o.Attachments = *so.Attachments
	}
	if so.Bcc != nil {
		o.Bcc = *so.Bcc
	}
	if so.Cc != nil {
		o.Cc = *so.Cc
	}
	if so.Content != nil {
		o.Content = *so.Content
	}
	if so.From != nil {
		o.From = *so.From
	}
	if so.Subject != nil {
		o.Subject = *so.Subject
	}
	if so.To != nil {
		o.To = *so.To
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
}

// Validate valides the current information stored into the structure.
func (o *Email) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("from", o.From); err != nil {
		requiredErrors = append(requiredErrors, err)
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
		Description:    `Attachments is a list of attachments to send.`,
		Exposed:        true,
		Name:           "attachments",
		SubType:        "list_attachments",
		Type:           "external",
	},
	"Bcc": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Bcc",
		Description:    `Bcc represents email that should be in copy but hidden.`,
		Exposed:        true,
		Name:           "bcc",
		SubType:        "list_emails",
		Type:           "external",
	},
	"Cc": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Cc",
		Description:    `Cc represents the addresses that should be in copy.`,
		Exposed:        true,
		Name:           "cc",
		SubType:        "list_emails",
		Type:           "external",
	},
	"Content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		Description:    `Content of the email to send.`,
		Exposed:        true,
		Name:           "content",
		Type:           "string",
	},
	"From": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "From",
		Description:    `From represents the sender of the email.`,
		Exposed:        true,
		Name:           "from",
		Required:       true,
		Type:           "string",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `Subject represents the subject of the email.`,
		Exposed:        true,
		Name:           "subject",
		Type:           "string",
	},
	"To": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "To",
		Description:    `To represents receivers of the email.`,
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
		Description:    `Attachments is a list of attachments to send.`,
		Exposed:        true,
		Name:           "attachments",
		SubType:        "list_attachments",
		Type:           "external",
	},
	"bcc": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Bcc",
		Description:    `Bcc represents email that should be in copy but hidden.`,
		Exposed:        true,
		Name:           "bcc",
		SubType:        "list_emails",
		Type:           "external",
	},
	"cc": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Cc",
		Description:    `Cc represents the addresses that should be in copy.`,
		Exposed:        true,
		Name:           "cc",
		SubType:        "list_emails",
		Type:           "external",
	},
	"content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		Description:    `Content of the email to send.`,
		Exposed:        true,
		Name:           "content",
		Type:           "string",
	},
	"from": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "From",
		Description:    `From represents the sender of the email.`,
		Exposed:        true,
		Name:           "from",
		Required:       true,
		Type:           "string",
	},
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `Subject represents the subject of the email.`,
		Exposed:        true,
		Name:           "subject",
		Type:           "string",
	},
	"to": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "To",
		Description:    `To represents receivers of the email.`,
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

// SparseEmailsList represents a list of SparseEmails
type SparseEmailsList []*SparseEmail

// Identity returns the identity of the objects in the list.
func (o SparseEmailsList) Identity() elemental.Identity {

	return EmailIdentity
}

// Copy returns a pointer to a copy the SparseEmailsList.
func (o SparseEmailsList) Copy() elemental.Identifiables {

	copy := append(SparseEmailsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEmailsList.
func (o SparseEmailsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEmailsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEmail))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEmailsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEmailsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseEmailsList converted to EmailsList.
func (o SparseEmailsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEmailsList) Version() int {

	return 1
}

// SparseEmail represents the sparse version of a email.
type SparseEmail struct {
	// Attachments is a list of attachments to send.
	Attachments *map[string]string `json:"attachments,omitempty" bson:"-" mapstructure:"attachments,omitempty"`

	// Bcc represents email that should be in copy but hidden.
	Bcc *[]string `json:"bcc,omitempty" bson:"-" mapstructure:"bcc,omitempty"`

	// Cc represents the addresses that should be in copy.
	Cc *[]string `json:"cc,omitempty" bson:"-" mapstructure:"cc,omitempty"`

	// Content of the email to send.
	Content *string `json:"content,omitempty" bson:"-" mapstructure:"content,omitempty"`

	// From represents the sender of the email.
	From *string `json:"from,omitempty" bson:"-" mapstructure:"from,omitempty"`

	// Subject represents the subject of the email.
	Subject *string `json:"subject,omitempty" bson:"-" mapstructure:"subject,omitempty"`

	// To represents receivers of the email.
	To *[]string `json:"to,omitempty" bson:"-" mapstructure:"to,omitempty"`

	// Type represents the type of the content.
	Type *EmailTypeValue `json:"type,omitempty" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseEmail returns a new  SparseEmail.
func NewSparseEmail() *SparseEmail {
	return &SparseEmail{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEmail) Identity() elemental.Identity {

	return EmailIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEmail) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEmail) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseEmail) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEmail) ToPlain() elemental.PlainIdentifiable {

	out := NewEmail()
	if o.Attachments != nil {
		out.Attachments = *o.Attachments
	}
	if o.Bcc != nil {
		out.Bcc = *o.Bcc
	}
	if o.Cc != nil {
		out.Cc = *o.Cc
	}
	if o.Content != nil {
		out.Content = *o.Content
	}
	if o.From != nil {
		out.From = *o.From
	}
	if o.Subject != nil {
		out.Subject = *o.Subject
	}
	if o.To != nil {
		out.To = *o.To
	}
	if o.Type != nil {
		out.Type = *o.Type
	}

	return out
}
