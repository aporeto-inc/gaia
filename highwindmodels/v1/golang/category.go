package highwindmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// CategoryIdentity represents the Identity of the object
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
	// Name of the category.
	Name string `json:"name" bson:"name"`

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

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Category) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *Category) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Category) DefaultOrder() []string {

	return []string{}
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
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
