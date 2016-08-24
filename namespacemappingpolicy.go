package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

// NamespaceMappingPolicyIdentity represents the Identity of the object
var NamespaceMappingPolicyIdentity = elemental.Identity{
	Name:     "namespacemappingpolicy",
	Category: "namespacemappingpolicies",
}

// NamespaceMappingPoliciesList represents a list of NamespaceMappingPolicies
type NamespaceMappingPoliciesList []*NamespaceMappingPolicy

// NamespaceMappingPolicy represents the model of a namespacemappingpolicy
type NamespaceMappingPolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"-"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" cql:"annotation,omitempty"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" cql:"associatedtags,omitempty"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" cql:"createdat,omitempty"`

	// Deleted marks if the entity has been deleted.
	Deleted bool `json:"-" cql:"deleted,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" cql:"description,omitempty"`

	// mappedNamespace is the mapped namespace
	MappedNamespace string `json:"mappedNamespace" cql:"mappednamespace,omitempty"`

	// Name is the name of the entity
	Name string `json:"name" cql:"name,omitempty"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" cql:"parentid,omitempty"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" cql:"parenttype,omitempty"`

	// Status of an entity
	Status enum.EntityStatus `json:"status" cql:"status,omitempty"`

	// Subject is the subject.
	Subject [][]string `json:"subject" cql:"-"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" cql:"updatedat,omitempty"`
}

// NewNamespaceMappingPolicy returns a new *NamespaceMappingPolicy
func NewNamespaceMappingPolicy() *NamespaceMappingPolicy {

	return &NamespaceMappingPolicy{
		Status: enum.Active,
	}
}

// Identity returns the Identity of the object.
func (o *NamespaceMappingPolicy) Identity() elemental.Identity {

	return NamespaceMappingPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NamespaceMappingPolicy) Identifier() string {

	return o.ID
}

func (o *NamespaceMappingPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NamespaceMappingPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *NamespaceMappingPolicy) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *NamespaceMappingPolicy) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *NamespaceMappingPolicy) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *NamespaceMappingPolicy) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *NamespaceMappingPolicy) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *NamespaceMappingPolicy) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *NamespaceMappingPolicy) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *NamespaceMappingPolicy) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *NamespaceMappingPolicy) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetParentID returns the parentID of the receiver
func (o *NamespaceMappingPolicy) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *NamespaceMappingPolicy) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *NamespaceMappingPolicy) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *NamespaceMappingPolicy) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetStatus returns the status of the receiver
func (o *NamespaceMappingPolicy) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *NamespaceMappingPolicy) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *NamespaceMappingPolicy) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *NamespaceMappingPolicy) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("mappedNamespace", o.MappedNamespace); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o NamespaceMappingPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return NamespaceMappingPolicyAttributesMap[name]
}

// NamespaceMappingPolicyAttributesMap represents the map of attribute for NamespaceMappingPolicy.
var NamespaceMappingPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"Annotation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"CreatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
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
		Filterable:     true,
		Getter:         true,
		Name:           "deleted",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"MappedNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "mappedNamespace",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "subject",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"UpdatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
