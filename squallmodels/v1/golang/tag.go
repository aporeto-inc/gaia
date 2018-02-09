package squallmodels

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// TagIdentity represents the Identity of the object.
var TagIdentity = elemental.Identity{
	Name:     "tag",
	Category: "tags",
}

// TagsList represents a list of Tags
type TagsList []*Tag

// ContentIdentity returns the identity of the objects in the list.
func (o TagsList) ContentIdentity() elemental.Identity {

	return TagIdentity
}

// Copy returns a pointer to a copy the TagsList.
func (o TagsList) Copy() elemental.ContentIdentifiable {

	copy := append(TagsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TagsList.
func (o TagsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(TagsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Tag))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TagsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TagsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o TagsList) Version() int {

	return 1
}

// Tag represents the model of a tag
type Tag struct {
	// Count represents the number of time the tag is used.
	Count int `json:"count" bson:"count" mapstructure:"count,omitempty"`

	// Namespace represents the namespace of the counted tag.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Value represents the value of the tag.
	Value string `json:"value" bson:"value" mapstructure:"value,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewTag returns a new *Tag
func NewTag() *Tag {

	return &Tag{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Tag) Identity() elemental.Identity {

	return TagIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Tag) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Tag) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Tag) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Tag) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Tag) Doc() string {
	return `A tag is a string in the form of "key=value" that can applied to all objects in the system. They are used for policy resolution. Tags starting by a "$" are derived from the property of an object (for instance an object with ID set to xxx and a name set to "the name" will be tagged by default with "$name=the name" and "$id=xxx"). Tags starting with an "@" have been generated by an external system.`
}

func (o *Tag) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Tag) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("value", o.Value); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidatePattern("value", o.Value, `^[\w\d\*\$\+\.:,|@<>/-]+=[= \w\d\*\$\+\.:,|@~<>#/-]+$`, true); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("value", o.Value); err != nil {
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
func (*Tag) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TagAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TagLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Tag) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TagAttributesMap
}

// TagAttributesMap represents the map of attribute for Tag.
var TagAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Count": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Count",
		Description:    `Count represents the number of time the tag is used.`,
		Exposed:        true,
		Name:           "count",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace represents the namespace of the counted tag.`,
		Exposed:        true,
		Format:         "free",
		Name:           "namespace",
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Value": elemental.AttributeSpecification{
		AllowedChars:   `^[\w\d\*\$\+\.:,|@<>/-]+=[= \w\d\*\$\+\.:,|@~<>#/-]+$`,
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		CreationOnly:   true,
		Description:    `Value represents the value of the tag.`,
		Exposed:        true,
		Format:         "free",
		Name:           "value",
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

// TagLowerCaseAttributesMap represents the map of attribute for Tag.
var TagLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"count": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Count",
		Description:    `Count represents the number of time the tag is used.`,
		Exposed:        true,
		Name:           "count",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace represents the namespace of the counted tag.`,
		Exposed:        true,
		Format:         "free",
		Name:           "namespace",
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"value": elemental.AttributeSpecification{
		AllowedChars:   `^[\w\d\*\$\+\.:,|@<>/-]+=[= \w\d\*\$\+\.:,|@~<>#/-]+$`,
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		CreationOnly:   true,
		Description:    `Value represents the value of the tag.`,
		Exposed:        true,
		Format:         "free",
		Name:           "value",
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}
