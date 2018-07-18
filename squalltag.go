package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
)

// SquallTagIdentity represents the Identity of the object.
var SquallTagIdentity = elemental.Identity{
	Name:     "squalltag",
	Category: "squalltags",
	Private:  true,
}

// SquallTagsList represents a list of SquallTags
type SquallTagsList []*SquallTag

// Identity returns the identity of the objects in the list.
func (o SquallTagsList) Identity() elemental.Identity {

	return SquallTagIdentity
}

// Copy returns a pointer to a copy the SquallTagsList.
func (o SquallTagsList) Copy() elemental.Identifiables {

	copy := append(SquallTagsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SquallTagsList.
func (o SquallTagsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SquallTagsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SquallTag))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SquallTagsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SquallTagsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o SquallTagsList) Version() int {

	return 1
}

// SquallTag represents the model of a squalltag
type SquallTag struct {
	// Number of time this tag is used.
	Count int `json:"count" bson:"count" mapstructure:"count,omitempty"`

	// namespace containing these tags.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Value of the tag.
	Value string `json:"value" bson:"value" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewSquallTag returns a new *SquallTag
func NewSquallTag() *SquallTag {

	return &SquallTag{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *SquallTag) Identity() elemental.Identity {

	return SquallTagIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SquallTag) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SquallTag) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SquallTag) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *SquallTag) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *SquallTag) Doc() string {
	return `Internal api that retrieve all tags from squall and their count for given
` + "`" + `?identity=<identity>` + "`" + ` parameter with their counts.`
}

func (o *SquallTag) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *SquallTag) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*SquallTag) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := SquallTagAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return SquallTagLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*SquallTag) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SquallTagAttributesMap
}

// SquallTagAttributesMap represents the map of attribute for SquallTag.
var SquallTagAttributesMap = map[string]elemental.AttributeSpecification{
	"Count": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Count",
		Description:    `Number of time this tag is used.`,
		Exposed:        true,
		Name:           "count",
		Stored:         true,
		Type:           "integer",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `namespace containing these tags.`,
		Exposed:        true,
		Name:           "namespace",
		Stored:         true,
		Type:           "string",
	},
	"Value": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Value of the tag.`,
		Exposed:        true,
		Name:           "value",
		Stored:         true,
		Type:           "string",
	},
}

// SquallTagLowerCaseAttributesMap represents the map of attribute for SquallTag.
var SquallTagLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"count": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Count",
		Description:    `Number of time this tag is used.`,
		Exposed:        true,
		Name:           "count",
		Stored:         true,
		Type:           "integer",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `namespace containing these tags.`,
		Exposed:        true,
		Name:           "namespace",
		Stored:         true,
		Type:           "string",
	},
	"value": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Value of the tag.`,
		Exposed:        true,
		Name:           "value",
		Stored:         true,
		Type:           "string",
	},
}
