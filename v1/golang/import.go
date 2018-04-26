package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// ImportModeValue represents the possible values for attribute "mode".
type ImportModeValue string

const (
	// ImportModeAppend represents the value Append.
	ImportModeAppend ImportModeValue = "Append"

	// ImportModeReplaceFull represents the value ReplaceFull.
	ImportModeReplaceFull ImportModeValue = "ReplaceFull"

	// ImportModeReplacePartial represents the value ReplacePartial.
	ImportModeReplacePartial ImportModeValue = "ReplacePartial"
)

// ImportIdentity represents the Identity of the object.
var ImportIdentity = elemental.Identity{
	Name:     "import",
	Category: "import",
	Private:  false,
}

// ImportsList represents a list of Imports
type ImportsList []*Import

// ContentIdentity returns the identity of the objects in the list.
func (o ImportsList) ContentIdentity() elemental.Identity {

	return ImportIdentity
}

// Copy returns a pointer to a copy the ImportsList.
func (o ImportsList) Copy() elemental.ContentIdentifiable {

	copy := append(ImportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ImportsList.
func (o ImportsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(ImportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Import))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ImportsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ImportsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o ImportsList) Version() int {

	return 1
}

// Import represents the model of a import
type Import struct {
	// The data to import.
	Data *Export `json:"data" bson:"-" mapstructure:"data,omitempty"`

	// How to import the data.
	Mode ImportModeValue `json:"mode" bson:"-" mapstructure:"mode,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewImport returns a new *Import
func NewImport() *Import {

	return &Import{
		ModelVersion: 1,
		Data:         NewExport(),
	}
}

// Identity returns the Identity of the object.
func (o *Import) Identity() elemental.Identity {

	return ImportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Import) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Import) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Import) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Import) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Import) Doc() string {
	return `Imports an export of policies and related objects into the namespace.`
}

func (o *Import) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Import) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("data", o.Data); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("mode", string(o.Mode), []string{"Append", "ReplacePartial", "ReplaceFull"}, false); err != nil {
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
func (*Import) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ImportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ImportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Import) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ImportAttributesMap
}

// ImportAttributesMap represents the map of attribute for Import.
var ImportAttributesMap = map[string]elemental.AttributeSpecification{
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `The data to import.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		SubType:        "exported_data",
		Type:           "external",
	},
	"Mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Append", "ReplacePartial", "ReplaceFull"},
		ConvertedName:  "Mode",
		Description:    `How to import the data.`,
		Exposed:        true,
		Name:           "mode",
		Required:       true,
		Type:           "enum",
	},
}

// ImportLowerCaseAttributesMap represents the map of attribute for Import.
var ImportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `The data to import.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		SubType:        "exported_data",
		Type:           "external",
	},
	"mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Append", "ReplacePartial", "ReplaceFull"},
		ConvertedName:  "Mode",
		Description:    `How to import the data.`,
		Exposed:        true,
		Name:           "mode",
		Required:       true,
		Type:           "enum",
	},
}
