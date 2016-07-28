package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	TagAttributeNameDescription    elemental.AttributeSpecificationNameKey = "tag/Description"
	TagAttributeNameID             elemental.AttributeSpecificationNameKey = "tag/ID"
	TagAttributeNameAnnotation     elemental.AttributeSpecificationNameKey = "tag/annotation"
	TagAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "tag/associatedTags"
	TagAttributeNameCreatedAt      elemental.AttributeSpecificationNameKey = "tag/createdAt"
	TagAttributeNameDeleted        elemental.AttributeSpecificationNameKey = "tag/deleted"
	TagAttributeNameNamespace      elemental.AttributeSpecificationNameKey = "tag/namespace"
	TagAttributeNameStatus         elemental.AttributeSpecificationNameKey = "tag/status"
	TagAttributeNameUpdatedAt      elemental.AttributeSpecificationNameKey = "tag/updatedAt"
	TagAttributeNameValue          elemental.AttributeSpecificationNameKey = "tag/value"
)

// TagIdentity represents the Identity of the object
var TagIdentity = elemental.Identity{
	Name:     "tag",
	Category: "tags",
}

// TagsList represents a list of Tags
type TagsList []*Tag

// Tag represents the model of a tag
type Tag struct {
	Description    string            `json:"Description,omitempty" cql:"description,omitempty"`
	ID             string            `json:"ID,omitempty" cql:"id,omitempty"`
	Annotation     map[string]string `json:"annotation,omitempty" cql:"annotation,omitempty"`
	AssociatedTags []string          `json:"associatedTags,omitempty" cql:"associatedtags,omitempty"`
	CreatedAt      time.Time         `json:"createdAt,omitempty" cql:"createdat,omitempty"`
	Deleted        bool              `json:"-" cql:"deleted,omitempty"`
	Namespace      string            `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`
	Status         enum.EntityStatus `json:"status,omitempty" cql:"status,omitempty"`
	UpdatedAt      time.Time         `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
	Value          string            `json:"value,omitempty" cql:"value,primarykey,omitempty"`
}

// NewTag returns a new *Tag
func NewTag() *Tag {

	return &Tag{}
}

// Identity returns the Identity of the object.
func (o *Tag) Identity() elemental.Identity {

	return TagIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Tag) Identifier() string {

	return o.ID
}

func (o *Tag) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Tag) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Tag) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Tag) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Tag) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *Tag) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *Tag) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetNamespace returns the namespace of the receiver
func (o *Tag) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Tag) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetStatus returns the status of the receiver
func (o *Tag) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Tag) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Tag) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Tag) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("value", o.Value); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o Tag) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return TagAttributesMap[name]
}

var TagAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	TagAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "Description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	TagAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	TagAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	TagAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	TagAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	TagAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	TagAttributeNameNamespace: elemental.AttributeSpecification{
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
	TagAttributeNameStatus: elemental.AttributeSpecification{
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
	TagAttributeNameUpdatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	TagAttributeNameValue: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "value",
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}
