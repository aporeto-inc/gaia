package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RecipeCookIdentity represents the Identity of the object.
var RecipeCookIdentity = elemental.Identity{
	Name:     "recipecook",
	Category: "recipecooks",
	Package:  "ignis",
	Private:  false,
}

// RecipeCooksList represents a list of RecipeCooks
type RecipeCooksList []*RecipeCook

// Identity returns the identity of the objects in the list.
func (o RecipeCooksList) Identity() elemental.Identity {

	return RecipeCookIdentity
}

// Copy returns a pointer to a copy the RecipeCooksList.
func (o RecipeCooksList) Copy() elemental.Identifiables {

	copy := append(RecipeCooksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the RecipeCooksList.
func (o RecipeCooksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(RecipeCooksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*RecipeCook))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o RecipeCooksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RecipeCooksList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the RecipeCooksList converted to SparseRecipeCooksList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o RecipeCooksList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o RecipeCooksList) Version() int {

	return 1
}

// RecipeCook represents the model of a recipecook
type RecipeCook struct {
	// Parameters contains the computed parameters.
	Parameters map[string]interface{} `json:"parameters" bson:"-" mapstructure:"parameters,omitempty"`

	// Template of the recipe.
	Template string `json:"template" bson:"-" mapstructure:"template,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewRecipeCook returns a new *RecipeCook
func NewRecipeCook() *RecipeCook {

	return &RecipeCook{
		ModelVersion: 1,
		Parameters:   map[string]interface{}{},
	}
}

// Identity returns the Identity of the object.
func (o *RecipeCook) Identity() elemental.Identity {

	return RecipeCookIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *RecipeCook) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *RecipeCook) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *RecipeCook) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *RecipeCook) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *RecipeCook) Doc() string {
	return `A RecipeCook cooks a recipe based on parameters.`
}

func (o *RecipeCook) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *RecipeCook) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseRecipeCook{
			Parameters: &o.Parameters,
			Template:   &o.Template,
		}
	}

	sp := &SparseRecipeCook{}
	for _, f := range fields {
		switch f {
		case "parameters":
			sp.Parameters = &(o.Parameters)
		case "template":
			sp.Template = &(o.Template)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseRecipeCook to the object.
func (o *RecipeCook) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseRecipeCook)
	if so.Parameters != nil {
		o.Parameters = *so.Parameters
	}
	if so.Template != nil {
		o.Template = *so.Template
	}
}

// DeepCopy returns a deep copy if the RecipeCook.
func (o *RecipeCook) DeepCopy() *RecipeCook {

	if o == nil {
		return nil
	}

	out := &RecipeCook{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *RecipeCook.
func (o *RecipeCook) DeepCopyInto(out *RecipeCook) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy RecipeCook: %s", err))
	}

	*out = *target.(*RecipeCook)
}

// Validate valides the current information stored into the structure.
func (o *RecipeCook) Validate() error {

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
func (*RecipeCook) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RecipeCookAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RecipeCookLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*RecipeCook) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RecipeCookAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *RecipeCook) ValueForAttribute(name string) interface{} {

	switch name {
	case "parameters":
		return o.Parameters
	case "template":
		return o.Template
	}

	return nil
}

// RecipeCookAttributesMap represents the map of attribute for RecipeCook.
var RecipeCookAttributesMap = map[string]elemental.AttributeSpecification{
	"Parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters contains the computed parameters.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"Template": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Template",
		Description:    `Template of the recipe.`,
		Exposed:        true,
		Name:           "template",
		Type:           "string",
	},
}

// RecipeCookLowerCaseAttributesMap represents the map of attribute for RecipeCook.
var RecipeCookLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters contains the computed parameters.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"template": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Template",
		Description:    `Template of the recipe.`,
		Exposed:        true,
		Name:           "template",
		Type:           "string",
	},
}

// SparseRecipeCooksList represents a list of SparseRecipeCooks
type SparseRecipeCooksList []*SparseRecipeCook

// Identity returns the identity of the objects in the list.
func (o SparseRecipeCooksList) Identity() elemental.Identity {

	return RecipeCookIdentity
}

// Copy returns a pointer to a copy the SparseRecipeCooksList.
func (o SparseRecipeCooksList) Copy() elemental.Identifiables {

	copy := append(SparseRecipeCooksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseRecipeCooksList.
func (o SparseRecipeCooksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseRecipeCooksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseRecipeCook))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseRecipeCooksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseRecipeCooksList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseRecipeCooksList converted to RecipeCooksList.
func (o SparseRecipeCooksList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseRecipeCooksList) Version() int {

	return 1
}

// SparseRecipeCook represents the sparse version of a recipecook.
type SparseRecipeCook struct {
	// Parameters contains the computed parameters.
	Parameters *map[string]interface{} `json:"parameters,omitempty" bson:"-" mapstructure:"parameters,omitempty"`

	// Template of the recipe.
	Template *string `json:"template,omitempty" bson:"-" mapstructure:"template,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseRecipeCook returns a new  SparseRecipeCook.
func NewSparseRecipeCook() *SparseRecipeCook {
	return &SparseRecipeCook{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseRecipeCook) Identity() elemental.Identity {

	return RecipeCookIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseRecipeCook) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseRecipeCook) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseRecipeCook) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseRecipeCook) ToPlain() elemental.PlainIdentifiable {

	out := NewRecipeCook()
	if o.Parameters != nil {
		out.Parameters = *o.Parameters
	}
	if o.Template != nil {
		out.Template = *o.Template
	}

	return out
}

// DeepCopy returns a deep copy if the SparseRecipeCook.
func (o *SparseRecipeCook) DeepCopy() *SparseRecipeCook {

	if o == nil {
		return nil
	}

	out := &SparseRecipeCook{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseRecipeCook.
func (o *SparseRecipeCook) DeepCopyInto(out *SparseRecipeCook) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseRecipeCook: %s", err))
	}

	*out = *target.(*SparseRecipeCook)
}
