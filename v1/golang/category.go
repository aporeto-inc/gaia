package models

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// CategoryIdentity represents the Identity of the object.
var CategoryIdentity = elemental.Identity{
	Name:     "category",
	Category: "categories",
}

// CategoriesList represents a list of Categories
type CategoriesList []*Category

// ContentIdentity returns the identity of the objects in the list.
func (o CategoriesList) ContentIdentity() elemental.Identity {

	return CategoryIdentity
}

// Copy returns a pointer to a copy the CategoriesList.
func (o CategoriesList) Copy() elemental.ContentIdentifiable {

	copy := append(CategoriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CategoriesList.
func (o CategoriesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(CategoriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Category))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CategoriesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CategoriesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o CategoriesList) Version() int {

	return 1
}

// Category represents the model of a category
type Category struct {
	// ID is the identifier of the category.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Description is the desription of the category.
	Description string `json:"description" bson:"-" mapstructure:"description,omitempty"`

	// Name of the category.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewCategory returns a new *Category
func NewCategory() *Category {

	return &Category{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Category) Identity() elemental.Identity {

	return CategoryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Category) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Category) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Category) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Category) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Category) Doc() string {
	return `Category allows to categorized services`
}

func (o *Category) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Category) Validate() error {

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
func (*Category) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CategoryAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CategoryLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Category) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CategoryAttributesMap
}

// CategoryAttributesMap represents the map of attribute for Category.
var CategoryAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the category.`,
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
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the desription of the category.`,
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the category.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// CategoryLowerCaseAttributesMap represents the map of attribute for Category.
var CategoryLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the category.`,
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
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the desription of the category.`,
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the category.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}
