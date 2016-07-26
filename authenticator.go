package gaia

import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	AuthenticatorAttributeNameID               elemental.AttributeSpecificationNameKey = "authenticator/ID"
	AuthenticatorAttributeNameAnnotation       elemental.AttributeSpecificationNameKey = "authenticator/annotation"
	AuthenticatorAttributeNameAssociatedTags   elemental.AttributeSpecificationNameKey = "authenticator/associatedTags"
	AuthenticatorAttributeNameConfiguration    elemental.AttributeSpecificationNameKey = "authenticator/configuration"
	AuthenticatorAttributeNameCreatedAt        elemental.AttributeSpecificationNameKey = "authenticator/createdAt"
	AuthenticatorAttributeNameDefaultNamespace elemental.AttributeSpecificationNameKey = "authenticator/defaultNamespace"
	AuthenticatorAttributeNameDeleted          elemental.AttributeSpecificationNameKey = "authenticator/deleted"
	AuthenticatorAttributeNameDescription      elemental.AttributeSpecificationNameKey = "authenticator/description"
	AuthenticatorAttributeNameMethod           elemental.AttributeSpecificationNameKey = "authenticator/method"
	AuthenticatorAttributeNameName             elemental.AttributeSpecificationNameKey = "authenticator/name"
	AuthenticatorAttributeNameNamespace        elemental.AttributeSpecificationNameKey = "authenticator/namespace"
	AuthenticatorAttributeNameOwner            elemental.AttributeSpecificationNameKey = "authenticator/owner"
	AuthenticatorAttributeNameStatus           elemental.AttributeSpecificationNameKey = "authenticator/status"
	AuthenticatorAttributeNameUpdatedAt        elemental.AttributeSpecificationNameKey = "authenticator/updatedAt"
)

// AuthenticatorMethodValue represents the possible values for attribute "method".
type AuthenticatorMethodValue string

const (
	AuthenticatorMethodCertificate AuthenticatorMethodValue = "Certificate"
	AuthenticatorMethodKey         AuthenticatorMethodValue = "Key"
	AuthenticatorMethodLdap        AuthenticatorMethodValue = "LDAP"
	AuthenticatorMethodOauth       AuthenticatorMethodValue = "OAUTH"
)

// AuthenticatorIdentity represents the Identity of the object
var AuthenticatorIdentity = elemental.Identity{
	Name:     "authenticator",
	Category: "authenticators",
}

// AuthenticatorsList represents a list of Authenticators
type AuthenticatorsList []*Authenticator

// Authenticator represents the model of a authenticator
type Authenticator struct {
	ID               string                   `json:"ID,omitempty" cql:"id,omitempty"`
	Annotation       map[string]string        `json:"annotation,omitempty" cql:"annotation,omitempty"`
	AssociatedTags   []string                 `json:"associatedTags,omitempty" cql:"associatedtags,omitempty"`
	Configuration    map[string]string        `json:"configuration,omitempty" cql:"configuration,omitempty"`
	CreatedAt        time.Time                `json:"createdAt,omitempty" cql:"createdat,omitempty"`
	DefaultNamespace string                   `json:"defaultNamespace,omitempty" cql:"defaultnamespace,primarykey,omitempty"`
	Deleted          bool                     `json:"-" cql:"deleted,omitempty"`
	Description      string                   `json:"description,omitempty" cql:"description,omitempty"`
	Method           AuthenticatorMethodValue `json:"method,omitempty" cql:"method,omitempty"`
	Name             string                   `json:"name,omitempty" cql:"name,primarykey,omitempty"`
	Namespace        string                   `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`
	Owner            []string                 `json:"owner,omitempty" cql:"owner,omitempty"`
	Status           enum.EntityStatus        `json:"status,omitempty" cql:"status,omitempty"`
	UpdatedAt        time.Time                `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
}

// NewAuthenticator returns a new *Authenticator
func NewAuthenticator() *Authenticator {

	return &Authenticator{}
}

// Identity returns the Identity of the object.
func (o *Authenticator) Identity() elemental.Identity {

	return AuthenticatorIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Authenticator) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Authenticator) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Authenticator) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Authenticator) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Authenticator) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *Authenticator) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *Authenticator) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *Authenticator) GetName() string {
	return o.Name
}

// GetNamespace returns the namespace of the receiver
func (o *Authenticator) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Authenticator) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetStatus returns the status of the receiver
func (o *Authenticator) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Authenticator) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Authenticator) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Authenticator) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateStringInList("method", string(o.Method), []string{"Certificate", "Key", "LDAP", "OAUTH"}); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o Authenticator) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return AuthenticatorAttributesMap[name]
}

var AuthenticatorAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	AuthenticatorAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	AuthenticatorAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	AuthenticatorAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tag_list",
		Type:           "external",
	},
	AuthenticatorAttributeNameConfiguration: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "configuration",
		Required:       true,
		Stored:         true,
		SubType:        "auth_config",
		Type:           "external",
	},
	AuthenticatorAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	AuthenticatorAttributeNameDefaultNamespace: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "defaultNamespace",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
	AuthenticatorAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Filterable:     true,
		ForeignKey:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	AuthenticatorAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		Stored:         true,
		Type:           "string",
	},
	AuthenticatorAttributeNameMethod: elemental.AttributeSpecification{
		AllowedChoices: []string{"Certificate", "Key", "LDAP", "OAUTH"},
		CreationOnly:   true,
		Exposed:        true,
		Name:           "method",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	AuthenticatorAttributeNameName: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	AuthenticatorAttributeNameNamespace: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	AuthenticatorAttributeNameOwner: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "owner",
		Stored:         true,
		SubType:        "tag_list",
		Type:           "external",
	},
	AuthenticatorAttributeNameStatus: elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Candidate", "Disabled"},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		Stored:         true,
		SubType:        "status_enum",
		Type:           "external",
	},
	AuthenticatorAttributeNameUpdatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
}
