package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/golang/constants"

// UserCertificateStatusValue represents the possible values for attribute "certificateStatus".
type UserCertificateStatusValue string

const (
	// UserCertificateStatusRenew represents the value RENEW.
	UserCertificateStatusRenew UserCertificateStatusValue = "RENEW"

	// UserCertificateStatusRevoked represents the value REVOKED.
	UserCertificateStatusRevoked UserCertificateStatusValue = "REVOKED"

	// UserCertificateStatusValid represents the value VALID.
	UserCertificateStatusValid UserCertificateStatusValue = "VALID"
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
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"id,primarykey,omitempty" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" cql:"annotation,omitempty" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" cql:"associatedtags,omitempty" bson:"associatedtags"`

	// Certificate provides a certificate for the user
	Certificate string `json:"certificate" cql:"certificate,omitempty" bson:"certificate"`

	// CertificateExpirationDate indicates the expiration day for the certificate.
	CertificateExpirationDate time.Time `json:"certificateExpirationDate" cql:"certificateexpirationdate,omitempty" bson:"certificateexpirationdate"`

	// CertificateStatus provides the status of the certificate. Update with RENEW to get a new certificate.
	CertificateStatus UserCertificateStatusValue `json:"certificateStatus" cql:"certificatestatus,omitempty" bson:"certificatestatus"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" cql:"createdat,omitempty" bson:"createdat"`

	// Deleted marks if the entity has been deleted.
	Deleted bool `json:"-" cql:"deleted,omitempty" bson:"deleted"`

	// Description is the description of the object.
	Description string `json:"description" cql:"description,omitempty" bson:"description"`

	// e-mail address of the user
	Email string `json:"email" cql:"email,omitempty" bson:"email"`

	// Key provides the key for the user. Only available at create or update time.
	Key string `json:"-" cql:"key,omitempty" bson:"key"`

	// Name is the name of the entity
	Name string `json:"name" cql:"name,omitempty" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty" bson:"_namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" cql:"-" bson:"-"`

	// ParentAuthenticator is an Internal attribute that points to the parent authenticator.
	ParentAuthenticator string `json:"-" cql:"parentauthenticator,primarykey,omitempty" bson:"_parentauthenticator"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" cql:"parentid,omitempty" bson:"parentid"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" cql:"parenttype,omitempty" bson:"parenttype"`

	// Status of an entity
	Status constants.EntityStatus `json:"status" cql:"status,omitempty" bson:"status"`

	// OU attribute for the generated certificates
	SubOrganizations []string `json:"subOrganizations" cql:"suborganizations,omitempty" bson:"suborganizations"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" cql:"updatedat,omitempty" bson:"updatedat"`

	// CommonName (CN) for the user certificate
	UserName string `json:"userName" cql:"username,omitempty" bson:"username"`
}

// NewUser returns a new *User
func NewUser() *User {

	return &User{
		AssociatedTags: []string{},
		NormalizedTags: []string{},
		Status:         constants.Active,
	}
}

// Identity returns the Identity of the object.
func (o *User) Identity() elemental.Identity {

	return UserIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *User) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *User) SetIdentifier(ID string) {

	o.ID = ID
}

func (o *User) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
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

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *User) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *User) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetParentID returns the parentID of the receiver
func (o *User) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *User) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *User) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *User) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetStatus returns the status of the receiver
func (o *User) GetStatus() constants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *User) SetStatus(status constants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *User) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *User) Validate() error {

	errors := elemental.Errors{}

	if err := elemental.ValidateStringInList("certificateStatus", string(o.CertificateStatus), []string{"RENEW", "REVOKED", "VALID"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumLength("userName", o.UserName, 64, false); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (User) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return UserAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (User) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return UserAttributesMap
}

// UserAttributesMap represents the map of attribute for User.
var UserAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Annotation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AssociatedTags are the list of tags attached to an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"Certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Certificate provides a certificate for the user`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificate",
		Stored:         true,
		Type:           "string",
	},
	"CertificateExpirationDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `CertificateExpirationDate indicates the expiration day for the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "certificateExpirationDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"CertificateStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"RENEW", "REVOKED", "VALID"},
		Description:    `CertificateStatus provides the status of the certificate. Update with RENEW to get a new certificate.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "certificateStatus",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"CreatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CreatedAt is the time at which an entity was created`,
		Exposed:        true,
		Name:           "createdAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Deleted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Deleted marks if the entity has been deleted.`,
		Getter:         true,
		Name:           "deleted",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `e-mail address of the user`,
		Exposed:        true,
		Filterable:     true,
		Format:         "email",
		Name:           "email",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Key provides the key for the user. Only available at create or update time.`,
		Format:         "free",
		Name:           "key",
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `NormalizedTags contains the list of normalized tags of the entities`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"ParentAuthenticator": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentAuthenticator is an Internal attribute that points to the parent authenticator.`,
		Filterable:     true,
		Format:         "free",
		Name:           "parentAuthenticator",
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentID is the ID of the parent, if any,`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentType": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Status of an entity`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "status",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		SubType:        "status_enum",
		Type:           "external",
	},
	"SubOrganizations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `OU attribute for the generated certificates`,
		Exposed:        true,
		Filterable:     true,
		Name:           "subOrganizations",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"UpdatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdatedAt is the time at which an entity was updated.`,
		Exposed:        true,
		Name:           "updatedAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"UserName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `CommonName (CN) for the user certificate`,
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
