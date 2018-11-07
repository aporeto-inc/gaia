package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TabulationIdentity represents the Identity of the object.
var TabulationIdentity = elemental.Identity{
	Name:     "tabulation",
	Category: "tabulations",
	Package:  "elena",
	Private:  false,
}

// TabulationsList represents a list of Tabulations
type TabulationsList []*Tabulation

// Identity returns the identity of the objects in the list.
func (o TabulationsList) Identity() elemental.Identity {

	return TabulationIdentity
}

// Copy returns a pointer to a copy the TabulationsList.
func (o TabulationsList) Copy() elemental.Identifiables {

	copy := append(TabulationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TabulationsList.
func (o TabulationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TabulationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Tabulation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TabulationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TabulationsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TabulationsList converted to SparseTabulationsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TabulationsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o TabulationsList) Version() int {

	return 1
}

// Tabulation represents the model of a tabulation
type Tabulation struct {
	// Headers contains the requests headers that matched.
	Headers []string `json:"headers" bson:"-" mapstructure:"headers,omitempty"`

	// Rows contains the tabulated data.
	Rows [][]interface{} `json:"rows" bson:"-" mapstructure:"rows,omitempty"`

	// TargetIdentity contains the requested target identity.
	TargetIdentity string `json:"targetIdentity" bson:"-" mapstructure:"targetIdentity,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewTabulation returns a new *Tabulation
func NewTabulation() *Tabulation {

	return &Tabulation{
		ModelVersion: 1,
		Rows:         [][]interface{}{},
	}
}

// Identity returns the Identity of the object.
func (o *Tabulation) Identity() elemental.Identity {

	return TabulationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Tabulation) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Tabulation) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Tabulation) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Tabulation) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Tabulation) Doc() string {
	return `Tabulate API allows you to retrieve a custom table view for any identity using
any tags you like as columns.`
}

func (o *Tabulation) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Tabulation) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTabulation{
			Headers:        &o.Headers,
			Rows:           &o.Rows,
			TargetIdentity: &o.TargetIdentity,
		}
	}

	sp := &SparseTabulation{}
	for _, f := range fields {
		switch f {
		case "headers":
			sp.Headers = &(o.Headers)
		case "rows":
			sp.Rows = &(o.Rows)
		case "targetIdentity":
			sp.TargetIdentity = &(o.TargetIdentity)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTabulation to the object.
func (o *Tabulation) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTabulation)
	if so.Headers != nil {
		o.Headers = *so.Headers
	}
	if so.Rows != nil {
		o.Rows = *so.Rows
	}
	if so.TargetIdentity != nil {
		o.TargetIdentity = *so.TargetIdentity
	}
}

// DeepCopy returns a deep copy if the Tabulation.
func (o *Tabulation) DeepCopy() *Tabulation {

	if o == nil {
		return nil
	}

	out := &Tabulation{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Tabulation.
func (o *Tabulation) DeepCopyInto(out *Tabulation) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Tabulation: %s", err))
	}

	*out = *target.(*Tabulation)
}

// Validate valides the current information stored into the structure.
func (o *Tabulation) Validate() error {

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
func (*Tabulation) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TabulationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TabulationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Tabulation) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TabulationAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Tabulation) ValueForAttribute(name string) interface{} {

	switch name {
	case "headers":
		return o.Headers
	case "rows":
		return o.Rows
	case "targetIdentity":
		return o.TargetIdentity
	}

	return nil
}

// TabulationAttributesMap represents the map of attribute for Tabulation.
var TabulationAttributesMap = map[string]elemental.AttributeSpecification{
	"Headers": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Headers",
		Description:    `Headers contains the requests headers that matched.`,
		Exposed:        true,
		Name:           "headers",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"Rows": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Rows",
		Description:    `Rows contains the tabulated data.`,
		Exposed:        true,
		Name:           "rows",
		ReadOnly:       true,
		SubType:        "tabulated_data",
		Type:           "external",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "TargetIdentity",
		Description:    `TargetIdentity contains the requested target identity.`,
		Exposed:        true,
		Name:           "targetIdentity",
		ReadOnly:       true,
		Type:           "string",
	},
}

// TabulationLowerCaseAttributesMap represents the map of attribute for Tabulation.
var TabulationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"headers": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Headers",
		Description:    `Headers contains the requests headers that matched.`,
		Exposed:        true,
		Name:           "headers",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"rows": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Rows",
		Description:    `Rows contains the tabulated data.`,
		Exposed:        true,
		Name:           "rows",
		ReadOnly:       true,
		SubType:        "tabulated_data",
		Type:           "external",
	},
	"targetidentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "TargetIdentity",
		Description:    `TargetIdentity contains the requested target identity.`,
		Exposed:        true,
		Name:           "targetIdentity",
		ReadOnly:       true,
		Type:           "string",
	},
}

// SparseTabulationsList represents a list of SparseTabulations
type SparseTabulationsList []*SparseTabulation

// Identity returns the identity of the objects in the list.
func (o SparseTabulationsList) Identity() elemental.Identity {

	return TabulationIdentity
}

// Copy returns a pointer to a copy the SparseTabulationsList.
func (o SparseTabulationsList) Copy() elemental.Identifiables {

	copy := append(SparseTabulationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTabulationsList.
func (o SparseTabulationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTabulationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTabulation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTabulationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTabulationsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTabulationsList converted to TabulationsList.
func (o SparseTabulationsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTabulationsList) Version() int {

	return 1
}

// SparseTabulation represents the sparse version of a tabulation.
type SparseTabulation struct {
	// Headers contains the requests headers that matched.
	Headers *[]string `json:"headers,omitempty" bson:"-" mapstructure:"headers,omitempty"`

	// Rows contains the tabulated data.
	Rows *[][]interface{} `json:"rows,omitempty" bson:"-" mapstructure:"rows,omitempty"`

	// TargetIdentity contains the requested target identity.
	TargetIdentity *string `json:"targetIdentity,omitempty" bson:"-" mapstructure:"targetIdentity,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseTabulation returns a new  SparseTabulation.
func NewSparseTabulation() *SparseTabulation {
	return &SparseTabulation{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTabulation) Identity() elemental.Identity {

	return TabulationIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTabulation) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTabulation) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseTabulation) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTabulation) ToPlain() elemental.PlainIdentifiable {

	out := NewTabulation()
	if o.Headers != nil {
		out.Headers = *o.Headers
	}
	if o.Rows != nil {
		out.Rows = *o.Rows
	}
	if o.TargetIdentity != nil {
		out.TargetIdentity = *o.TargetIdentity
	}

	return out
}

// DeepCopy returns a deep copy if the SparseTabulation.
func (o *SparseTabulation) DeepCopy() *SparseTabulation {

	if o == nil {
		return nil
	}

	out := &SparseTabulation{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTabulation.
func (o *SparseTabulation) DeepCopyInto(out *SparseTabulation) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTabulation: %s", err))
	}

	*out = *target.(*SparseTabulation)
}
