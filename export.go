package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ExportIdentity represents the Identity of the object.
var ExportIdentity = elemental.Identity{
	Name:     "export",
	Category: "export",
	Package:  "yuna",
	Private:  false,
}

// ExportsList represents a list of Exports
type ExportsList []*Export

// Identity returns the identity of the objects in the list.
func (o ExportsList) Identity() elemental.Identity {

	return ExportIdentity
}

// Copy returns a pointer to a copy the ExportsList.
func (o ExportsList) Copy() elemental.Identifiables {

	copy := append(ExportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ExportsList.
func (o ExportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ExportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Export))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ExportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ExportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ExportsList converted to SparseExportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ExportsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o ExportsList) Version() int {

	return 1
}

// Export represents the model of a export
type Export struct {
	// APIVersion of the api used for the exported data.
	APIVersion int `json:"APIVersion" bson:"-" mapstructure:"APIVersion,omitempty"`

	// List of all exported data.
	Data map[string][]map[string]interface{} `json:"data" bson:"-" mapstructure:"data,omitempty"`

	// The list of identities to export.
	Identities []string `json:"identities" bson:"identities" mapstructure:"identities,omitempty"`

	// Label allows to define a unique label for this export. When importing the
	// content of the export, this label will be added as a tag that will be used to
	// recognize imported object in a later import.
	Label string `json:"label" bson:"-" mapstructure:"label,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewExport returns a new *Export
func NewExport() *Export {

	return &Export{
		ModelVersion: 1,
		Data:         map[string][]map[string]interface{}{},
	}
}

// Identity returns the Identity of the object.
func (o *Export) Identity() elemental.Identity {

	return ExportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Export) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Export) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Export) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Export) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Export) Doc() string {
	return `Export the policies and related objects in a given namespace.`
}

func (o *Export) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Export) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseExport{
			APIVersion: &o.APIVersion,
			Data:       &o.Data,
			Identities: &o.Identities,
			Label:      &o.Label,
		}
	}

	sp := &SparseExport{}
	for _, f := range fields {
		switch f {
		case "APIVersion":
			sp.APIVersion = &(o.APIVersion)
		case "data":
			sp.Data = &(o.Data)
		case "identities":
			sp.Identities = &(o.Identities)
		case "label":
			sp.Label = &(o.Label)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseExport to the object.
func (o *Export) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseExport)
	if so.APIVersion != nil {
		o.APIVersion = *so.APIVersion
	}
	if so.Data != nil {
		o.Data = *so.Data
	}
	if so.Identities != nil {
		o.Identities = *so.Identities
	}
	if so.Label != nil {
		o.Label = *so.Label
	}
}

// DeepCopy returns a deep copy if the Export.
func (o *Export) DeepCopy() *Export {

	if o == nil {
		return nil
	}

	out := &Export{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Export.
func (o *Export) DeepCopyInto(out *Export) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Export: %s", err))
	}

	*out = *target.(*Export)
}

// Validate valides the current information stored into the structure.
func (o *Export) Validate() error {

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
func (*Export) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ExportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ExportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Export) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ExportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Export) ValueForAttribute(name string) interface{} {

	switch name {
	case "APIVersion":
		return o.APIVersion
	case "data":
		return o.Data
	case "identities":
		return o.Identities
	case "label":
		return o.Label
	}

	return nil
}

// ExportAttributesMap represents the map of attribute for Export.
var ExportAttributesMap = map[string]elemental.AttributeSpecification{
	"APIVersion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "APIVersion",
		Description:    `APIVersion of the api used for the exported data.`,
		Exposed:        true,
		Name:           "APIVersion",
		ReadOnly:       true,
		Type:           "integer",
	},
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `List of all exported data.`,
		Exposed:        true,
		Name:           "data",
		SubType:        "exported_data_content",
		Type:           "external",
	},
	"Identities": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Identities",
		Description:    `The list of identities to export.`,
		Exposed:        true,
		Name:           "identities",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Label": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Label",
		Description: `Label allows to define a unique label for this export. When importing the
content of the export, this label will be added as a tag that will be used to
recognize imported object in a later import.`,
		Exposed:  true,
		Name:     "label",
		ReadOnly: true,
		Type:     "string",
	},
}

// ExportLowerCaseAttributesMap represents the map of attribute for Export.
var ExportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"apiversion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "APIVersion",
		Description:    `APIVersion of the api used for the exported data.`,
		Exposed:        true,
		Name:           "APIVersion",
		ReadOnly:       true,
		Type:           "integer",
	},
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `List of all exported data.`,
		Exposed:        true,
		Name:           "data",
		SubType:        "exported_data_content",
		Type:           "external",
	},
	"identities": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Identities",
		Description:    `The list of identities to export.`,
		Exposed:        true,
		Name:           "identities",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"label": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Label",
		Description: `Label allows to define a unique label for this export. When importing the
content of the export, this label will be added as a tag that will be used to
recognize imported object in a later import.`,
		Exposed:  true,
		Name:     "label",
		ReadOnly: true,
		Type:     "string",
	},
}

// SparseExportsList represents a list of SparseExports
type SparseExportsList []*SparseExport

// Identity returns the identity of the objects in the list.
func (o SparseExportsList) Identity() elemental.Identity {

	return ExportIdentity
}

// Copy returns a pointer to a copy the SparseExportsList.
func (o SparseExportsList) Copy() elemental.Identifiables {

	copy := append(SparseExportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseExportsList.
func (o SparseExportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseExportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseExport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseExportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseExportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseExportsList converted to ExportsList.
func (o SparseExportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseExportsList) Version() int {

	return 1
}

// SparseExport represents the sparse version of a export.
type SparseExport struct {
	// APIVersion of the api used for the exported data.
	APIVersion *int `json:"APIVersion,omitempty" bson:"-" mapstructure:"APIVersion,omitempty"`

	// List of all exported data.
	Data *map[string][]map[string]interface{} `json:"data,omitempty" bson:"-" mapstructure:"data,omitempty"`

	// The list of identities to export.
	Identities *[]string `json:"identities,omitempty" bson:"identities" mapstructure:"identities,omitempty"`

	// Label allows to define a unique label for this export. When importing the
	// content of the export, this label will be added as a tag that will be used to
	// recognize imported object in a later import.
	Label *string `json:"label,omitempty" bson:"-" mapstructure:"label,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseExport returns a new  SparseExport.
func NewSparseExport() *SparseExport {
	return &SparseExport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseExport) Identity() elemental.Identity {

	return ExportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseExport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseExport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseExport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseExport) ToPlain() elemental.PlainIdentifiable {

	out := NewExport()
	if o.APIVersion != nil {
		out.APIVersion = *o.APIVersion
	}
	if o.Data != nil {
		out.Data = *o.Data
	}
	if o.Identities != nil {
		out.Identities = *o.Identities
	}
	if o.Label != nil {
		out.Label = *o.Label
	}

	return out
}

// DeepCopy returns a deep copy if the SparseExport.
func (o *SparseExport) DeepCopy() *SparseExport {

	if o == nil {
		return nil
	}

	out := &SparseExport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseExport.
func (o *SparseExport) DeepCopyInto(out *SparseExport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseExport: %s", err))
	}

	*out = *target.(*SparseExport)
}
