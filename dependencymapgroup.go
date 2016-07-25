// Copyright 2016 Aporeto Inc.

package gaia

import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	DependencyMapGroupAttributeNameID             elemental.AttributeSpecificationNameKey = "dependencymapgroup/ID"
	DependencyMapGroupAttributeNameAnnotation     elemental.AttributeSpecificationNameKey = "dependencymapgroup/annotation"
	DependencyMapGroupAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "dependencymapgroup/associatedTags"
	DependencyMapGroupAttributeNameColor          elemental.AttributeSpecificationNameKey = "dependencymapgroup/color"
	DependencyMapGroupAttributeNameCreatedAt      elemental.AttributeSpecificationNameKey = "dependencymapgroup/createdAt"
	DependencyMapGroupAttributeNameDeleted        elemental.AttributeSpecificationNameKey = "dependencymapgroup/deleted"
	DependencyMapGroupAttributeNameDescription    elemental.AttributeSpecificationNameKey = "dependencymapgroup/description"
	DependencyMapGroupAttributeNameName           elemental.AttributeSpecificationNameKey = "dependencymapgroup/name"
	DependencyMapGroupAttributeNameNamespace      elemental.AttributeSpecificationNameKey = "dependencymapgroup/namespace"
	DependencyMapGroupAttributeNameOwner          elemental.AttributeSpecificationNameKey = "dependencymapgroup/owner"
	DependencyMapGroupAttributeNameStatus         elemental.AttributeSpecificationNameKey = "dependencymapgroup/status"
	DependencyMapGroupAttributeNameUpdatedAt      elemental.AttributeSpecificationNameKey = "dependencymapgroup/updatedAt"
	DependencyMapGroupAttributeNameValues         elemental.AttributeSpecificationNameKey = "dependencymapgroup/values"
)

// DependencyMapGroupIdentity represents the Identity of the object
var DependencyMapGroupIdentity = elemental.Identity{
	Name:     "dependencymapgroup",
	Category: "dependencymapgroups",
}

// DependencyMapGroupsList represents a list of DependencyMapGroups
type DependencyMapGroupsList []*DependencyMapGroup

// DependencyMapGroup represents the model of a dependencymapgroup
type DependencyMapGroup struct {
	ID             string            `json:"ID,omitempty" cql:"id,primarykey,omitempty"`
	Annotation     map[string]string `json:"annotation,omitempty" cql:"annotation,omitempty"`
	AssociatedTags []string          `json:"associatedTags,omitempty" cql:"associatedtags,omitempty"`
	Color          string            `json:"color,omitempty" cql:"color,omitempty"`
	CreatedAt      time.Time         `json:"createdAt,omitempty" cql:"createdat,omitempty"`
	Deleted        bool              `json:"-" cql:"deleted,omitempty"`
	Description    string            `json:"description,omitempty" cql:"description,omitempty"`
	Name           string            `json:"name,omitempty" cql:"name,omitempty"`
	Namespace      string            `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`
	Owner          []string          `json:"owner,omitempty" cql:"owner,omitempty"`
	Status         enum.EntityStatus `json:"status,omitempty" cql:"status,omitempty"`
	UpdatedAt      time.Time         `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
	Values         []string          `json:"values,omitempty" cql:"values,omitempty"`
}

// NewDependencyMapGroup returns a new *DependencyMapGroup
func NewDependencyMapGroup() *DependencyMapGroup {

	return &DependencyMapGroup{}
}

// Identity returns the Identity of the object.
func (o *DependencyMapGroup) Identity() elemental.Identity {

	return DependencyMapGroupIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DependencyMapGroup) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DependencyMapGroup) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *DependencyMapGroup) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *DependencyMapGroup) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *DependencyMapGroup) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *DependencyMapGroup) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *DependencyMapGroup) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *DependencyMapGroup) GetName() string {
	return o.Name
}

// GetNamespace returns the namespace of the receiver
func (o *DependencyMapGroup) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *DependencyMapGroup) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetStatus returns the status of the receiver
func (o *DependencyMapGroup) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *DependencyMapGroup) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *DependencyMapGroup) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *DependencyMapGroup) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("color", o.Color); err != nil {
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
func (o DependencyMapGroup) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return DependencyMapGroupAttributesMap[name]
}

var DependencyMapGroupAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	DependencyMapGroupAttributeNameID: elemental.AttributeSpecification{
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
	DependencyMapGroupAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	DependencyMapGroupAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tag_list",
		Type:           "external",
	},
	DependencyMapGroupAttributeNameColor: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "color",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	DependencyMapGroupAttributeNameCreatedAt: elemental.AttributeSpecification{
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
	DependencyMapGroupAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Filterable:     true,
		ForeignKey:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	DependencyMapGroupAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		Stored:         true,
		Type:           "string",
	},
	DependencyMapGroupAttributeNameName: elemental.AttributeSpecification{
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
	DependencyMapGroupAttributeNameNamespace: elemental.AttributeSpecification{
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
	DependencyMapGroupAttributeNameOwner: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "owner",
		Stored:         true,
		SubType:        "tag_list",
		Type:           "external",
	},
	DependencyMapGroupAttributeNameStatus: elemental.AttributeSpecification{
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
	DependencyMapGroupAttributeNameUpdatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	DependencyMapGroupAttributeNameValues: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "values",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}
