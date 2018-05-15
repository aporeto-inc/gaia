package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// TagInjectIdentity represents the Identity of the object.
var TagInjectIdentity = elemental.Identity{
	Name:     "taginject",
	Category: "taginjects",
	Private:  true,
}

// TagInjectsList represents a list of TagInjects
type TagInjectsList []*TagInject

// ContentIdentity returns the identity of the objects in the list.
func (o TagInjectsList) ContentIdentity() elemental.Identity {

	return TagInjectIdentity
}

// Copy returns a pointer to a copy the TagInjectsList.
func (o TagInjectsList) Copy() elemental.ContentIdentifiable {

	copy := append(TagInjectsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TagInjectsList.
func (o TagInjectsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(TagInjectsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*TagInject))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TagInjectsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TagInjectsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o TagInjectsList) Version() int {

	return 1
}

// TagInject represents the model of a taginject
type TagInject struct {
	// List of tags to be added.
	AddedTags []string `json:"addedTags" bson:"-" mapstructure:"addedTags,omitempty"`

	// List of tags to be removed.
	RemovedTags []string `json:"removedTags" bson:"-" mapstructure:"removedTags,omitempty"`

	// List of tags to inject.
	TargetNamespace string `json:"targetNamespace" bson:"-" mapstructure:"targetNamespace,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewTagInject returns a new *TagInject
func NewTagInject() *TagInject {

	return &TagInject{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *TagInject) Identity() elemental.Identity {

	return TagInjectIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TagInject) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TagInject) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *TagInject) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *TagInject) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *TagInject) Doc() string {
	return `This internal api is used to inject a new tag in the database.`
}

func (o *TagInject) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *TagInject) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("targetNamespace", o.TargetNamespace); err != nil {
		requiredErrors = append(requiredErrors, err)
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
func (*TagInject) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TagInjectAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TagInjectLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TagInject) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TagInjectAttributesMap
}

// TagInjectAttributesMap represents the map of attribute for TagInject.
var TagInjectAttributesMap = map[string]elemental.AttributeSpecification{
	"AddedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AddedTags",
		Description:    `List of tags to be added.`,
		Exposed:        true,
		Name:           "addedTags",
		SubType:        "string",
		Type:           "list",
	},
	"RemovedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RemovedTags",
		Description:    `List of tags to be removed.`,
		Exposed:        true,
		Name:           "removedTags",
		SubType:        "string",
		Type:           "list",
	},
	"TargetNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description:    `List of tags to inject.`,
		Exposed:        true,
		Name:           "targetNamespace",
		Required:       true,
		Type:           "string",
	},
}

// TagInjectLowerCaseAttributesMap represents the map of attribute for TagInject.
var TagInjectLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"addedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AddedTags",
		Description:    `List of tags to be added.`,
		Exposed:        true,
		Name:           "addedTags",
		SubType:        "string",
		Type:           "list",
	},
	"removedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RemovedTags",
		Description:    `List of tags to be removed.`,
		Exposed:        true,
		Name:           "removedTags",
		SubType:        "string",
		Type:           "list",
	},
	"targetnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description:    `List of tags to inject.`,
		Exposed:        true,
		Name:           "targetNamespace",
		Required:       true,
		Type:           "string",
	},
}
