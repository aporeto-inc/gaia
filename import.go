package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
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
	Package:  "squall",
	Private:  false,
}

// ImportsList represents a list of Imports
type ImportsList []*Import

// Identity returns the identity of the objects in the list.
func (o ImportsList) Identity() elemental.Identity {

	return ImportIdentity
}

// Copy returns a pointer to a copy the ImportsList.
func (o ImportsList) Copy() elemental.Identifiables {

	copy := append(ImportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ImportsList.
func (o ImportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ImportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Import))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ImportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ImportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ImportsList converted to SparseImportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ImportsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
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

	sync.Mutex `json:"-" bson:"-"`
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

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Import) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseImport{
			Data: &o.Data,
			Mode: &o.Mode,
		}
	}

	sp := &SparseImport{}
	for _, f := range fields {
		switch f {
		case "data":
			sp.Data = &(o.Data)
		case "mode":
			sp.Mode = &(o.Mode)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseImport to the object.
func (o *Import) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseImport)
	if so.Data != nil {
		o.Data = *so.Data
	}
	if so.Mode != nil {
		o.Mode = *so.Mode
	}
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
		Type:           "enum",
	},
}

// SparseImportsList represents a list of SparseImports
type SparseImportsList []*SparseImport

// Identity returns the identity of the objects in the list.
func (o SparseImportsList) Identity() elemental.Identity {

	return ImportIdentity
}

// Copy returns a pointer to a copy the SparseImportsList.
func (o SparseImportsList) Copy() elemental.Identifiables {

	copy := append(SparseImportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseImportsList.
func (o SparseImportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseImportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseImport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseImportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseImportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseImportsList converted to ImportsList.
func (o SparseImportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseImportsList) Version() int {

	return 1
}

// SparseImport represents the sparse version of a import.
type SparseImport struct {
	// The data to import.
	Data **Export `json:"data,omitempty" bson:"-" mapstructure:"data,omitempty"`

	// How to import the data.
	Mode *ImportModeValue `json:"mode,omitempty" bson:"-" mapstructure:"mode,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseImport returns a new  SparseImport.
func NewSparseImport() *SparseImport {
	return &SparseImport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseImport) Identity() elemental.Identity {

	return ImportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseImport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseImport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseImport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseImport) ToPlain() elemental.PlainIdentifiable {

	out := NewImport()
	if o.Data != nil {
		out.Data = *o.Data
	}
	if o.Mode != nil {
		out.Mode = *o.Mode
	}

	return out
}
