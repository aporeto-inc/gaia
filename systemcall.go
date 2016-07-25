// Copyright 2016 Aporeto Inc.

package gaia

import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	SystemCallAttributeNameID             elemental.AttributeSpecificationNameKey = "systemcall/ID"
	SystemCallAttributeNameAnnotation     elemental.AttributeSpecificationNameKey = "systemcall/annotation"
	SystemCallAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "systemcall/associatedTags"
	SystemCallAttributeNameCreatedAt      elemental.AttributeSpecificationNameKey = "systemcall/createdAt"
	SystemCallAttributeNameDeleted        elemental.AttributeSpecificationNameKey = "systemcall/deleted"
	SystemCallAttributeNameDescription    elemental.AttributeSpecificationNameKey = "systemcall/description"
	SystemCallAttributeNameName           elemental.AttributeSpecificationNameKey = "systemcall/name"
	SystemCallAttributeNameNamespace      elemental.AttributeSpecificationNameKey = "systemcall/namespace"
	SystemCallAttributeNameOwner          elemental.AttributeSpecificationNameKey = "systemcall/owner"
	SystemCallAttributeNameStatus         elemental.AttributeSpecificationNameKey = "systemcall/status"
	SystemCallAttributeNameUpdatedAt      elemental.AttributeSpecificationNameKey = "systemcall/updatedAt"
)

// SystemCallIdentity represents the Identity of the object
var SystemCallIdentity = elemental.Identity{
	Name:     "systemcall",
	Category: "systemcalls",
}

// SystemCallsList represents a list of SystemCalls
type SystemCallsList []*SystemCall

// SystemCall represents the model of a systemcall
type SystemCall struct {
	ID             string            `json:"ID,omitempty" cql:"id,primarykey,omitempty"`
	Annotation     map[string]string `json:"annotation,omitempty" cql:"annotation,omitempty"`
	AssociatedTags []string          `json:"associatedTags,omitempty" cql:"associatedtags,omitempty"`
	CreatedAt      time.Time         `json:"createdAt,omitempty" cql:"createdat,omitempty"`
	Deleted        bool              `json:"-" cql:"deleted,omitempty"`
	Description    string            `json:"description,omitempty" cql:"description,omitempty"`
	Name           string            `json:"name,omitempty" cql:"name,omitempty"`
	Namespace      string            `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`
	Owner          []string          `json:"owner,omitempty" cql:"owner,omitempty"`
	Status         enum.EntityStatus `json:"status,omitempty" cql:"status,omitempty"`
	UpdatedAt      time.Time         `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
}

// NewSystemCall returns a new *SystemCall
func NewSystemCall() *SystemCall {

	return &SystemCall{}
}

// Identity returns the Identity of the object.
func (o *SystemCall) Identity() elemental.Identity {

	return SystemCallIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SystemCall) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SystemCall) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *SystemCall) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *SystemCall) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *SystemCall) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *SystemCall) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *SystemCall) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *SystemCall) GetName() string {
	return o.Name
}

// GetNamespace returns the namespace of the receiver
func (o *SystemCall) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *SystemCall) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetStatus returns the status of the receiver
func (o *SystemCall) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *SystemCall) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *SystemCall) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *SystemCall) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o SystemCall) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return SystemCallAttributesMap[name]
}

var SystemCallAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	SystemCallAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
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
	SystemCallAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	SystemCallAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tag_list",
		Type:           "external",
	},
	SystemCallAttributeNameCreatedAt: elemental.AttributeSpecification{
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
	SystemCallAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Filterable:     true,
		ForeignKey:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	SystemCallAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		Stored:         true,
		Type:           "string",
	},
	SystemCallAttributeNameName: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	SystemCallAttributeNameNamespace: elemental.AttributeSpecification{
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
	SystemCallAttributeNameOwner: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "owner",
		Stored:         true,
		SubType:        "tag_list",
		Type:           "external",
	},
	SystemCallAttributeNameStatus: elemental.AttributeSpecification{
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
	SystemCallAttributeNameUpdatedAt: elemental.AttributeSpecification{
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
