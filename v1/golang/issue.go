package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// IssueRealmValue represents the possible values for attribute "realm".
type IssueRealmValue string

const (
	// IssueRealmAwsidentitydocument represents the value AWSIdentityDocument.
	IssueRealmAwsidentitydocument IssueRealmValue = "AWSIdentityDocument"

	// IssueRealmCertificate represents the value Certificate.
	IssueRealmCertificate IssueRealmValue = "Certificate"

	// IssueRealmFacebook represents the value Facebook.
	IssueRealmFacebook IssueRealmValue = "Facebook"

	// IssueRealmGithub represents the value Github.
	IssueRealmGithub IssueRealmValue = "Github"

	// IssueRealmGoogle represents the value Google.
	IssueRealmGoogle IssueRealmValue = "Google"

	// IssueRealmLdap represents the value LDAP.
	IssueRealmLdap IssueRealmValue = "LDAP"

	// IssueRealmTwitter represents the value Twitter.
	IssueRealmTwitter IssueRealmValue = "Twitter"

	// IssueRealmVince represents the value Vince.
	IssueRealmVince IssueRealmValue = "Vince"
)

// IssueIdentity represents the Identity of the object.
var IssueIdentity = elemental.Identity{
	Name:     "issue",
	Category: "issue",
	Private:  false,
}

// IssuesList represents a list of Issues
type IssuesList []*Issue

// ContentIdentity returns the identity of the objects in the list.
func (o IssuesList) ContentIdentity() elemental.Identity {

	return IssueIdentity
}

// Copy returns a pointer to a copy the IssuesList.
func (o IssuesList) Copy() elemental.ContentIdentifiable {

	copy := append(IssuesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the IssuesList.
func (o IssuesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(IssuesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Issue))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o IssuesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o IssuesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o IssuesList) Version() int {

	return 1
}

// Issue represents the model of a issue
type Issue struct {
	// Data contains additional data. The value depends on the issuer type.
	Data string `json:"data" bson:"data" mapstructure:"data,omitempty"`

	// Metadata contains various additional information. Meaning depends on the realm.
	Metadata map[string]interface{} `json:"metadata" bson:"-" mapstructure:"metadata,omitempty"`

	// Token is the token to use for the registration.
	Token string `json:"token" bson:"-" mapstructure:"token,omitempty"`

	// Validity configures the max validity time for a token. If it is bigger than the
	// configured max validity, it will be capped.
	Validity string `json:"validity" bson:"validity" mapstructure:"validity,omitempty"`

	// Realm is the realm
	Realm IssueRealmValue `json:"realm" bson:"-" mapstructure:"realm,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewIssue returns a new *Issue
func NewIssue() *Issue {

	return &Issue{
		ModelVersion: 1,
		Metadata:     map[string]interface{}{},
		Validity:     "24h",
	}
}

// Identity returns the Identity of the object.
func (o *Issue) Identity() elemental.Identity {

	return IssueIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Issue) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Issue) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Issue) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Issue) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Issue) Doc() string {
	return `This API issues a new token according to given data.`
}

func (o *Issue) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Issue) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidatePattern("validity", o.Validity, `^([0-9]+h[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+h[0-9]+s|[0-9]+h[0-9]+m|[0-9]+s|[0-9]+h|[0-9]+m)$`, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("realm", string(o.Realm), []string{"AWSIdentityDocument", "Certificate", "Facebook", "Github", "Google", "LDAP", "Twitter", "Vince"}, false); err != nil {
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
func (*Issue) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := IssueAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return IssueLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Issue) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return IssueAttributesMap
}

// IssueAttributesMap represents the map of attribute for Issue.
var IssueAttributesMap = map[string]elemental.AttributeSpecification{
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Data contains additional data. The value depends on the issuer type.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "data",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		Description:    `Metadata contains various additional information. Meaning depends on the realm.`,
		Exposed:        true,
		Name:           "metadata",
		Orderable:      true,
		SubType:        "metadata",
		Type:           "external",
	},
	"Realm": elemental.AttributeSpecification{
		AllowedChoices: []string{"AWSIdentityDocument", "Certificate", "Facebook", "Github", "Google", "LDAP", "Twitter", "Vince"},
		ConvertedName:  "Realm",
		Description:    `Realm is the realm`,
		Exposed:        true,
		Name:           "realm",
		Required:       true,
		Type:           "enum",
	},
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description:    `Token is the token to use for the registration.`,
		Exposed:        true,
		Format:         "free",
		Name:           "token",
		ReadOnly:       true,
		Type:           "string",
	},
	"Validity": elemental.AttributeSpecification{
		AllowedChars:   `^([0-9]+h[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+h[0-9]+s|[0-9]+h[0-9]+m|[0-9]+s|[0-9]+h|[0-9]+m)$`,
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		DefaultValue:   "24h",
		Description: `Validity configures the max validity time for a token. If it is bigger than the
configured max validity, it will be capped.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "validity",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
}

// IssueLowerCaseAttributesMap represents the map of attribute for Issue.
var IssueLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Data contains additional data. The value depends on the issuer type.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "data",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		Description:    `Metadata contains various additional information. Meaning depends on the realm.`,
		Exposed:        true,
		Name:           "metadata",
		Orderable:      true,
		SubType:        "metadata",
		Type:           "external",
	},
	"realm": elemental.AttributeSpecification{
		AllowedChoices: []string{"AWSIdentityDocument", "Certificate", "Facebook", "Github", "Google", "LDAP", "Twitter", "Vince"},
		ConvertedName:  "Realm",
		Description:    `Realm is the realm`,
		Exposed:        true,
		Name:           "realm",
		Required:       true,
		Type:           "enum",
	},
	"token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description:    `Token is the token to use for the registration.`,
		Exposed:        true,
		Format:         "free",
		Name:           "token",
		ReadOnly:       true,
		Type:           "string",
	},
	"validity": elemental.AttributeSpecification{
		AllowedChars:   `^([0-9]+h[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+h[0-9]+s|[0-9]+h[0-9]+m|[0-9]+s|[0-9]+h|[0-9]+m)$`,
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		DefaultValue:   "24h",
		Description: `Validity configures the max validity time for a token. If it is bigger than the
configured max validity, it will be capped.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "validity",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
}
