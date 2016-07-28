package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	UserAttributeNameID                  elemental.AttributeSpecificationNameKey = "user/ID"
	UserAttributeNameAnnotation          elemental.AttributeSpecificationNameKey = "user/annotation"
	UserAttributeNameAssociatedTags      elemental.AttributeSpecificationNameKey = "user/associatedTags"
	UserAttributeNameCreatedAt           elemental.AttributeSpecificationNameKey = "user/createdAt"
	UserAttributeNameDeleted             elemental.AttributeSpecificationNameKey = "user/deleted"
	UserAttributeNameEmail               elemental.AttributeSpecificationNameKey = "user/email"
	UserAttributeNameName                elemental.AttributeSpecificationNameKey = "user/name"
	UserAttributeNameNamespace           elemental.AttributeSpecificationNameKey = "user/namespace"
	UserAttributeNameParentAuthenticator elemental.AttributeSpecificationNameKey = "user/parentAuthenticator"
	UserAttributeNameParentID            elemental.AttributeSpecificationNameKey = "user/parentID"
	UserAttributeNameStatus              elemental.AttributeSpecificationNameKey = "user/status"
	UserAttributeNameSubOrganizations    elemental.AttributeSpecificationNameKey = "user/subOrganizations"
	UserAttributeNameUpdatedAt           elemental.AttributeSpecificationNameKey = "user/updatedAt"
	UserAttributeNameUserName            elemental.AttributeSpecificationNameKey = "user/userName"
)

// UserIdentity represents the Identity of the object
var UserIdentity = elemental.Identity{
	Name:     "user",
	Category: "users",
}

// UsersList represents a list of Users
type UsersList []*User

// User represents the model of a user
type User struct {
	ID                  string            `json:"ID,omitempty" cql:"id,primarykey,omitempty"`
	Annotation          map[string]string `json:"annotation,omitempty" cql:"annotation,omitempty"`
	AssociatedTags      []string          `json:"associatedTags,omitempty" cql:"associatedtags,omitempty"`
	CreatedAt           time.Time         `json:"createdAt,omitempty" cql:"createdat,omitempty"`
	Deleted             bool              `json:"-" cql:"deleted,omitempty"`
	Email               string            `json:"email,omitempty" cql:"email,primarykey,omitempty"`
	Name                string            `json:"name,omitempty" cql:"name,omitempty"`
	Namespace           string            `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`
	ParentAuthenticator string            `json:"-" cql:"parentauthenticator,primarykey,omitempty"`
	ParentID            string            `json:"parentID,omitempty" cql:"parentid,omitempty"`
	Status              enum.EntityStatus `json:"status,omitempty" cql:"status,omitempty"`
	SubOrganizations    []string          `json:"subOrganizations,omitempty" cql:"suborganizations,omitempty"`
	UpdatedAt           time.Time         `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
	UserName            string            `json:"userName,omitempty" cql:"username,omitempty"`
}

// NewUser returns a new *User
func NewUser() *User {

	return &User{}
}

// Identity returns the Identity of the object.
func (o *User) Identity() elemental.Identity {

	return UserIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *User) Identifier() string {

	return o.ID
}

func (o *User) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *User) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *User) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *User) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *User) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *User) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *User) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *User) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *User) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *User) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *User) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetStatus returns the status of the receiver
func (o *User) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *User) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *User) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *User) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumLength("userName", o.UserName, 64, false); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o User) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return UserAttributesMap[name]
}

var UserAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	UserAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	UserAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	UserAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	UserAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	UserAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Filterable:     true,
		ForeignKey:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	UserAttributeNameEmail: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "email",
		Name:           "email",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
	UserAttributeNameName: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	UserAttributeNameNamespace: elemental.AttributeSpecification{
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
	UserAttributeNameParentAuthenticator: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Format:         "free",
		Name:           "parentAuthenticator",
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
	UserAttributeNameParentID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentID",
		Stored:         true,
		Type:           "string",
	},
	UserAttributeNameStatus: elemental.AttributeSpecification{
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
	UserAttributeNameSubOrganizations: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "subOrganizations",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	UserAttributeNameUpdatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	UserAttributeNameUserName: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		MaxLength:      64,
		Name:           "userName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}
