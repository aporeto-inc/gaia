package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// JaegerbatchIdentity represents the Identity of the object.
var JaegerbatchIdentity = elemental.Identity{
	Name:     "jaegerbatch",
	Category: "jaegerbatchs",
	Package:  "meister",
	Private:  false,
}

// JaegerbatchsList represents a list of Jaegerbatchs
type JaegerbatchsList []*Jaegerbatch

// Identity returns the identity of the objects in the list.
func (o JaegerbatchsList) Identity() elemental.Identity {

	return JaegerbatchIdentity
}

// Copy returns a pointer to a copy the JaegerbatchsList.
func (o JaegerbatchsList) Copy() elemental.Identifiables {

	copy := append(JaegerbatchsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the JaegerbatchsList.
func (o JaegerbatchsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(JaegerbatchsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Jaegerbatch))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o JaegerbatchsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o JaegerbatchsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the JaegerbatchsList converted to SparseJaegerbatchsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o JaegerbatchsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o JaegerbatchsList) Version() int {

	return 1
}

// Jaegerbatch represents the model of a jaegerbatch
type Jaegerbatch struct {
	// Represents a jaeger batch.
	Batch interface{} `json:"batch" bson:"batch" mapstructure:"batch,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewJaegerbatch returns a new *Jaegerbatch
func NewJaegerbatch() *Jaegerbatch {

	return &Jaegerbatch{
		ModelVersion: 1,
		Mutex:        &sync.Mutex{},
	}
}

// Identity returns the Identity of the object.
func (o *Jaegerbatch) Identity() elemental.Identity {

	return JaegerbatchIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Jaegerbatch) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Jaegerbatch) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Jaegerbatch) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Jaegerbatch) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Jaegerbatch) Doc() string {
	return `A jaegerbatch is a batch of jaeger spans. This is used by external service to
post jaeger span in our private jaeger services.`
}

func (o *Jaegerbatch) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Jaegerbatch) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseJaegerbatch{
			Batch: &o.Batch,
		}
	}

	sp := &SparseJaegerbatch{}
	for _, f := range fields {
		switch f {
		case "batch":
			sp.Batch = &(o.Batch)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseJaegerbatch to the object.
func (o *Jaegerbatch) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseJaegerbatch)
	if so.Batch != nil {
		o.Batch = *so.Batch
	}
}

// DeepCopy returns a deep copy if the Jaegerbatch.
func (o *Jaegerbatch) DeepCopy() *Jaegerbatch {

	if o == nil {
		return nil
	}

	out := &Jaegerbatch{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Jaegerbatch.
func (o *Jaegerbatch) DeepCopyInto(out *Jaegerbatch) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Jaegerbatch: %s", err))
	}

	*out = *target.(*Jaegerbatch)
}

// Validate valides the current information stored into the structure.
func (o *Jaegerbatch) Validate() error {

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
func (*Jaegerbatch) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := JaegerbatchAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return JaegerbatchLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Jaegerbatch) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return JaegerbatchAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Jaegerbatch) ValueForAttribute(name string) interface{} {

	switch name {
	case "batch":
		return o.Batch
	}

	return nil
}

// JaegerbatchAttributesMap represents the map of attribute for Jaegerbatch.
var JaegerbatchAttributesMap = map[string]elemental.AttributeSpecification{
	"Batch": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Batch",
		CreationOnly:   true,
		Description:    `Represents a jaeger batch.`,
		Exposed:        true,
		Name:           "batch",
		Stored:         true,
		Type:           "object",
	},
}

// JaegerbatchLowerCaseAttributesMap represents the map of attribute for Jaegerbatch.
var JaegerbatchLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"batch": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Batch",
		CreationOnly:   true,
		Description:    `Represents a jaeger batch.`,
		Exposed:        true,
		Name:           "batch",
		Stored:         true,
		Type:           "object",
	},
}

// SparseJaegerbatchsList represents a list of SparseJaegerbatchs
type SparseJaegerbatchsList []*SparseJaegerbatch

// Identity returns the identity of the objects in the list.
func (o SparseJaegerbatchsList) Identity() elemental.Identity {

	return JaegerbatchIdentity
}

// Copy returns a pointer to a copy the SparseJaegerbatchsList.
func (o SparseJaegerbatchsList) Copy() elemental.Identifiables {

	copy := append(SparseJaegerbatchsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseJaegerbatchsList.
func (o SparseJaegerbatchsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseJaegerbatchsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseJaegerbatch))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseJaegerbatchsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseJaegerbatchsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseJaegerbatchsList converted to JaegerbatchsList.
func (o SparseJaegerbatchsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseJaegerbatchsList) Version() int {

	return 1
}

// SparseJaegerbatch represents the sparse version of a jaegerbatch.
type SparseJaegerbatch struct {
	// Represents a jaeger batch.
	Batch *interface{} `json:"batch,omitempty" bson:"batch" mapstructure:"batch,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSparseJaegerbatch returns a new  SparseJaegerbatch.
func NewSparseJaegerbatch() *SparseJaegerbatch {
	return &SparseJaegerbatch{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseJaegerbatch) Identity() elemental.Identity {

	return JaegerbatchIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseJaegerbatch) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseJaegerbatch) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseJaegerbatch) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseJaegerbatch) ToPlain() elemental.PlainIdentifiable {

	out := NewJaegerbatch()
	if o.Batch != nil {
		out.Batch = *o.Batch
	}

	return out
}

// DeepCopy returns a deep copy if the SparseJaegerbatch.
func (o *SparseJaegerbatch) DeepCopy() *SparseJaegerbatch {

	if o == nil {
		return nil
	}

	out := &SparseJaegerbatch{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseJaegerbatch.
func (o *SparseJaegerbatch) DeepCopyInto(out *SparseJaegerbatch) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseJaegerbatch: %s", err))
	}

	*out = *target.(*SparseJaegerbatch)
}
