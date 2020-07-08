package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RenderNamespaceIdentity represents the Identity of the object.
var RenderNamespaceIdentity = elemental.Identity{
	Name:     "rendernamespace",
	Category: "rendernamespaces",
	Package:  "squall",
	Private:  false,
}

// RenderNamespacesList represents a list of RenderNamespaces
type RenderNamespacesList []*RenderNamespace

// Identity returns the identity of the objects in the list.
func (o RenderNamespacesList) Identity() elemental.Identity {

	return RenderNamespaceIdentity
}

// Copy returns a pointer to a copy the RenderNamespacesList.
func (o RenderNamespacesList) Copy() elemental.Identifiables {

	copy := append(RenderNamespacesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the RenderNamespacesList.
func (o RenderNamespacesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(RenderNamespacesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*RenderNamespace))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o RenderNamespacesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RenderNamespacesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the RenderNamespacesList converted to SparseRenderNamespacesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o RenderNamespacesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseRenderNamespacesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseRenderNamespace)
	}

	return out
}

// Version returns the version of the content.
func (o RenderNamespacesList) Version() int {

	return 1
}

// RenderNamespace represents the model of a rendernamespace
type RenderNamespace struct {
	// The namespace where the object should reside in.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// List of tags of the object to render the namespace for.
	Tags []string `json:"tags" msgpack:"tags" bson:"-" mapstructure:"tags,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewRenderNamespace returns a new *RenderNamespace
func NewRenderNamespace() *RenderNamespace {

	return &RenderNamespace{
		ModelVersion: 1,
		Tags:         []string{},
	}
}

// Identity returns the Identity of the object.
func (o *RenderNamespace) Identity() elemental.Identity {

	return RenderNamespaceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *RenderNamespace) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *RenderNamespace) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RenderNamespace) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesRenderNamespace{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RenderNamespace) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesRenderNamespace{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *RenderNamespace) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *RenderNamespace) BleveType() string {

	return "rendernamespace"
}

// DefaultOrder returns the list of default ordering fields.
func (o *RenderNamespace) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *RenderNamespace) Doc() string {

	return `This object allows you to determine which namespace an object should reside in
based on the tags provided.`
}

func (o *RenderNamespace) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *RenderNamespace) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseRenderNamespace{
			Namespace: &o.Namespace,
			Tags:      &o.Tags,
		}
	}

	sp := &SparseRenderNamespace{}
	for _, f := range fields {
		switch f {
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "tags":
			sp.Tags = &(o.Tags)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseRenderNamespace to the object.
func (o *RenderNamespace) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseRenderNamespace)
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Tags != nil {
		o.Tags = *so.Tags
	}
}

// DeepCopy returns a deep copy if the RenderNamespace.
func (o *RenderNamespace) DeepCopy() *RenderNamespace {

	if o == nil {
		return nil
	}

	out := &RenderNamespace{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *RenderNamespace.
func (o *RenderNamespace) DeepCopyInto(out *RenderNamespace) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy RenderNamespace: %s", err))
	}

	*out = *target.(*RenderNamespace)
}

// Validate valides the current information stored into the structure.
func (o *RenderNamespace) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("tags", o.Tags); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*RenderNamespace) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RenderNamespaceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RenderNamespaceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*RenderNamespace) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RenderNamespaceAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *RenderNamespace) ValueForAttribute(name string) interface{} {

	switch name {
	case "namespace":
		return o.Namespace
	case "tags":
		return o.Tags
	}

	return nil
}

// RenderNamespaceAttributesMap represents the map of attribute for RenderNamespace.
var RenderNamespaceAttributesMap = map[string]elemental.AttributeSpecification{
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `The namespace where the object should reside in.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
	"Tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `List of tags of the object to render the namespace for.`,
		Exposed:        true,
		Name:           "tags",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// RenderNamespaceLowerCaseAttributesMap represents the map of attribute for RenderNamespace.
var RenderNamespaceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `The namespace where the object should reside in.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
	"tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `List of tags of the object to render the namespace for.`,
		Exposed:        true,
		Name:           "tags",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// SparseRenderNamespacesList represents a list of SparseRenderNamespaces
type SparseRenderNamespacesList []*SparseRenderNamespace

// Identity returns the identity of the objects in the list.
func (o SparseRenderNamespacesList) Identity() elemental.Identity {

	return RenderNamespaceIdentity
}

// Copy returns a pointer to a copy the SparseRenderNamespacesList.
func (o SparseRenderNamespacesList) Copy() elemental.Identifiables {

	copy := append(SparseRenderNamespacesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseRenderNamespacesList.
func (o SparseRenderNamespacesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseRenderNamespacesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseRenderNamespace))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseRenderNamespacesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseRenderNamespacesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseRenderNamespacesList converted to RenderNamespacesList.
func (o SparseRenderNamespacesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseRenderNamespacesList) Version() int {

	return 1
}

// SparseRenderNamespace represents the sparse version of a rendernamespace.
type SparseRenderNamespace struct {
	// The namespace where the object should reside in.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// List of tags of the object to render the namespace for.
	Tags *[]string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"-" mapstructure:"tags,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseRenderNamespace returns a new  SparseRenderNamespace.
func NewSparseRenderNamespace() *SparseRenderNamespace {
	return &SparseRenderNamespace{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseRenderNamespace) Identity() elemental.Identity {

	return RenderNamespaceIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseRenderNamespace) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseRenderNamespace) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseRenderNamespace) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseRenderNamespace{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseRenderNamespace) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseRenderNamespace{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseRenderNamespace) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseRenderNamespace) ToPlain() elemental.PlainIdentifiable {

	out := NewRenderNamespace()
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Tags != nil {
		out.Tags = *o.Tags
	}

	return out
}

// DeepCopy returns a deep copy if the SparseRenderNamespace.
func (o *SparseRenderNamespace) DeepCopy() *SparseRenderNamespace {

	if o == nil {
		return nil
	}

	out := &SparseRenderNamespace{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseRenderNamespace.
func (o *SparseRenderNamespace) DeepCopyInto(out *SparseRenderNamespace) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseRenderNamespace: %s", err))
	}

	*out = *target.(*SparseRenderNamespace)
}

type mongoAttributesRenderNamespace struct {
}
type mongoAttributesSparseRenderNamespace struct {
}
