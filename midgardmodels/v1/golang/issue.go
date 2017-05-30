package midgardmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

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

// IssueIdentity represents the Identity of the object
var IssueIdentity = elemental.Identity{
	Name:     "issue",
	Category: "issue",
}

// IssuesList represents a list of Issues
type IssuesList []*Issue

// ContentIdentity returns the identity of the objects in the list.
func (o IssuesList) ContentIdentity() elemental.Identity {

	return IssueIdentity
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

// Issue represents the model of a issue
type Issue struct {
	// Data contains additional data. The value depends on the issuer type.
	Data string `json:"data" bson:"data"`

	// Metadata contains various additional information. Meaning depends on the realm.
	Metadata map[string]interface{} `json:"metadata" bson:"-"`

	// Realm is the realm
	Realm IssueRealmValue `json:"realm" bson:"-"`

	// Token is the token to use for the registration.
	Token string `json:"token" bson:"-"`

	// Validity configures the max validity time for a token. If it is bigger than the configured max validity, it will be capped.
	Validity string `json:"validity" bson:"validity"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewIssue returns a new *Issue
func NewIssue() *Issue {

	return &Issue{
		ModelVersion: 1.0,
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
func (o *Issue) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *Issue) Version() float64 {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *Issue) DefaultOrder() []string {

	return []string{}
}

func (o *Issue) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Issue) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("realm", string(o.Realm), []string{"AWSIdentityDocument", "Certificate", "Facebook", "Github", "Google", "LDAP", "Twitter", "Vince"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("validity", o.Validity, `^([0-9]+h[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+h[0-9]+s|[0-9]+h[0-9]+m|[0-9]+s|[0-9]+h|[0-9]+m)$`); err != nil {
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

	return IssueAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Issue) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return IssueAttributesMap
}

// IssueAttributesMap represents the map of attribute for Issue.
var IssueAttributesMap = map[string]elemental.AttributeSpecification{
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		Description:    `Metadata contains various additional information. Meaning depends on the realm.`,
		Exposed:        true,
		Name:           "metadata",
		Orderable:      true,
		SubType:        "metadata",
		Type:           "external",
	},
	"Realm": elemental.AttributeSpecification{
		AllowedChoices: []string{"AWSIdentityDocument", "Certificate", "Facebook", "Github", "Google", "LDAP", "Twitter", "Vince"},
		Description:    `Realm is the realm`,
		Exposed:        true,
		Name:           "realm",
		Required:       true,
		Type:           "enum",
	},
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
		DefaultValue:   "24h",
		Description:    `Validity configures the max validity time for a token. If it is bigger than the configured max validity, it will be capped.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "validity",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}
