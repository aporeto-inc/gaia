package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ProcessingUnitRefreshIdentity represents the Identity of the object.
var ProcessingUnitRefreshIdentity = elemental.Identity{
	Name:     "processingunitrefresh",
	Category: "processingunitrefreshs",
	Package:  "gaga",
	Private:  false,
}

// ProcessingUnitRefreshsList represents a list of ProcessingUnitRefreshs
type ProcessingUnitRefreshsList []*ProcessingUnitRefresh

// Identity returns the identity of the objects in the list.
func (o ProcessingUnitRefreshsList) Identity() elemental.Identity {

	return ProcessingUnitRefreshIdentity
}

// Copy returns a pointer to a copy the ProcessingUnitRefreshsList.
func (o ProcessingUnitRefreshsList) Copy() elemental.Identifiables {

	copy := append(ProcessingUnitRefreshsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ProcessingUnitRefreshsList.
func (o ProcessingUnitRefreshsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ProcessingUnitRefreshsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ProcessingUnitRefresh))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ProcessingUnitRefreshsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ProcessingUnitRefreshsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ProcessingUnitRefreshsList converted to SparseProcessingUnitRefreshsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ProcessingUnitRefreshsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o ProcessingUnitRefreshsList) Version() int {

	return 1
}

// ProcessingUnitRefresh represents the model of a processingunitrefresh
type ProcessingUnitRefresh struct {
	// ID contains the original ID of the Processing Unit.
	ID string `json:"ID" bson:"id" mapstructure:"ID,omitempty"`

	// Namespace contains the original namespace of the Processing Unit.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewProcessingUnitRefresh returns a new *ProcessingUnitRefresh
func NewProcessingUnitRefresh() *ProcessingUnitRefresh {

	return &ProcessingUnitRefresh{
		ModelVersion: 1,
		Mutex:        &sync.Mutex{},
	}
}

// Identity returns the Identity of the object.
func (o *ProcessingUnitRefresh) Identity() elemental.Identity {

	return ProcessingUnitRefreshIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ProcessingUnitRefresh) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ProcessingUnitRefresh) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *ProcessingUnitRefresh) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *ProcessingUnitRefresh) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ProcessingUnitRefresh) Doc() string {

	return `ProcessingUnitRefresh is sent to client when a poke has been triggered using the
parameter ` + "`" + `?notify=true` + "`" + `. This is used by instances of enforcerd to be notify an
external change on the processing unit must be processed.`
}

func (o *ProcessingUnitRefresh) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ProcessingUnitRefresh) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseProcessingUnitRefresh{
			ID:        &o.ID,
			Namespace: &o.Namespace,
		}
	}

	sp := &SparseProcessingUnitRefresh{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseProcessingUnitRefresh to the object.
func (o *ProcessingUnitRefresh) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseProcessingUnitRefresh)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
}

// DeepCopy returns a deep copy if the ProcessingUnitRefresh.
func (o *ProcessingUnitRefresh) DeepCopy() *ProcessingUnitRefresh {

	if o == nil {
		return nil
	}

	out := &ProcessingUnitRefresh{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ProcessingUnitRefresh.
func (o *ProcessingUnitRefresh) DeepCopyInto(out *ProcessingUnitRefresh) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ProcessingUnitRefresh: %s", err))
	}

	*out = *target.(*ProcessingUnitRefresh)
}

// Validate valides the current information stored into the structure.
func (o *ProcessingUnitRefresh) Validate() error {

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
func (*ProcessingUnitRefresh) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ProcessingUnitRefreshAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ProcessingUnitRefreshLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ProcessingUnitRefresh) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ProcessingUnitRefreshAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ProcessingUnitRefresh) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "namespace":
		return o.Namespace
	}

	return nil
}

// ProcessingUnitRefreshAttributesMap represents the map of attribute for ProcessingUnitRefresh.
var ProcessingUnitRefreshAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID contains the original ID of the Processing Unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace contains the original namespace of the Processing Unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "namespace",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// ProcessingUnitRefreshLowerCaseAttributesMap represents the map of attribute for ProcessingUnitRefresh.
var ProcessingUnitRefreshLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID contains the original ID of the Processing Unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace contains the original namespace of the Processing Unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "namespace",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseProcessingUnitRefreshsList represents a list of SparseProcessingUnitRefreshs
type SparseProcessingUnitRefreshsList []*SparseProcessingUnitRefresh

// Identity returns the identity of the objects in the list.
func (o SparseProcessingUnitRefreshsList) Identity() elemental.Identity {

	return ProcessingUnitRefreshIdentity
}

// Copy returns a pointer to a copy the SparseProcessingUnitRefreshsList.
func (o SparseProcessingUnitRefreshsList) Copy() elemental.Identifiables {

	copy := append(SparseProcessingUnitRefreshsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseProcessingUnitRefreshsList.
func (o SparseProcessingUnitRefreshsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseProcessingUnitRefreshsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseProcessingUnitRefresh))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseProcessingUnitRefreshsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseProcessingUnitRefreshsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseProcessingUnitRefreshsList converted to ProcessingUnitRefreshsList.
func (o SparseProcessingUnitRefreshsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseProcessingUnitRefreshsList) Version() int {

	return 1
}

// SparseProcessingUnitRefresh represents the sparse version of a processingunitrefresh.
type SparseProcessingUnitRefresh struct {
	// ID contains the original ID of the Processing Unit.
	ID *string `json:"ID,omitempty" bson:"id,omitempty" mapstructure:"ID,omitempty"`

	// Namespace contains the original namespace of the Processing Unit.
	Namespace *string `json:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSparseProcessingUnitRefresh returns a new  SparseProcessingUnitRefresh.
func NewSparseProcessingUnitRefresh() *SparseProcessingUnitRefresh {
	return &SparseProcessingUnitRefresh{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseProcessingUnitRefresh) Identity() elemental.Identity {

	return ProcessingUnitRefreshIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseProcessingUnitRefresh) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseProcessingUnitRefresh) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseProcessingUnitRefresh) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseProcessingUnitRefresh) ToPlain() elemental.PlainIdentifiable {

	out := NewProcessingUnitRefresh()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}

	return out
}

// DeepCopy returns a deep copy if the SparseProcessingUnitRefresh.
func (o *SparseProcessingUnitRefresh) DeepCopy() *SparseProcessingUnitRefresh {

	if o == nil {
		return nil
	}

	out := &SparseProcessingUnitRefresh{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseProcessingUnitRefresh.
func (o *SparseProcessingUnitRefresh) DeepCopyInto(out *SparseProcessingUnitRefresh) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseProcessingUnitRefresh: %s", err))
	}

	*out = *target.(*SparseProcessingUnitRefresh)
}
