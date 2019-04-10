package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PUNodeIdentity represents the Identity of the object.
var PUNodeIdentity = elemental.Identity{
	Name:     "punode",
	Category: "punodes",
	Package:  "squall",
	Private:  true,
}

// PUNodesList represents a list of PUNodes
type PUNodesList []*PUNode

// Identity returns the identity of the objects in the list.
func (o PUNodesList) Identity() elemental.Identity {

	return PUNodeIdentity
}

// Copy returns a pointer to a copy the PUNodesList.
func (o PUNodesList) Copy() elemental.Identifiables {

	copy := append(PUNodesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PUNodesList.
func (o PUNodesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PUNodesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PUNode))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PUNodesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PUNodesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PUNodesList converted to SparsePUNodesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PUNodesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o PUNodesList) Version() int {

	return 1
}

// PUNode represents the model of a punode
type PUNode struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Enforcement status of the pu.
	EnforcementStatus string `json:"enforcementStatus" bson:"-" mapstructure:"enforcementStatus,omitempty"`

	// Images of the pu.
	Images []string `json:"images" bson:"-" mapstructure:"images,omitempty"`

	// Last poke time of the pu.
	LastPokeTime time.Time `json:"lastPokeTime" bson:"-" mapstructure:"lastPokeTime,omitempty"`

	// Name of the pu.
	Name string `json:"name" bson:"-" mapstructure:"name,omitempty"`

	// Namespace of the pu.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// Status of the pu.
	Status string `json:"status" bson:"-" mapstructure:"status,omitempty"`

	// Tags of the pu.
	Tags []string `json:"tags" bson:"-" mapstructure:"tags,omitempty"`

	// Type of the pu.
	Type string `json:"type" bson:"-" mapstructure:"type,omitempty"`

	// If true, the pu is unreachable.
	Unreachable bool `json:"unreachable" bson:"-" mapstructure:"unreachable,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewPUNode returns a new *PUNode
func NewPUNode() *PUNode {

	return &PUNode{
		ModelVersion: 1,
		Mutex:        &sync.Mutex{},
		Images:       []string{},
		Tags:         []string{},
	}
}

// Identity returns the Identity of the object.
func (o *PUNode) Identity() elemental.Identity {

	return PUNodeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PUNode) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PUNode) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *PUNode) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *PUNode) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PUNode) Doc() string {

	return `Internal scaling down of a pu for dep map representation.`
}

func (o *PUNode) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PUNode) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePUNode{
			ID:                &o.ID,
			EnforcementStatus: &o.EnforcementStatus,
			Images:            &o.Images,
			LastPokeTime:      &o.LastPokeTime,
			Name:              &o.Name,
			Namespace:         &o.Namespace,
			Status:            &o.Status,
			Tags:              &o.Tags,
			Type:              &o.Type,
			Unreachable:       &o.Unreachable,
		}
	}

	sp := &SparsePUNode{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "enforcementStatus":
			sp.EnforcementStatus = &(o.EnforcementStatus)
		case "images":
			sp.Images = &(o.Images)
		case "lastPokeTime":
			sp.LastPokeTime = &(o.LastPokeTime)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "status":
			sp.Status = &(o.Status)
		case "tags":
			sp.Tags = &(o.Tags)
		case "type":
			sp.Type = &(o.Type)
		case "unreachable":
			sp.Unreachable = &(o.Unreachable)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePUNode to the object.
func (o *PUNode) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePUNode)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.EnforcementStatus != nil {
		o.EnforcementStatus = *so.EnforcementStatus
	}
	if so.Images != nil {
		o.Images = *so.Images
	}
	if so.LastPokeTime != nil {
		o.LastPokeTime = *so.LastPokeTime
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Status != nil {
		o.Status = *so.Status
	}
	if so.Tags != nil {
		o.Tags = *so.Tags
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
	if so.Unreachable != nil {
		o.Unreachable = *so.Unreachable
	}
}

// DeepCopy returns a deep copy if the PUNode.
func (o *PUNode) DeepCopy() *PUNode {

	if o == nil {
		return nil
	}

	out := &PUNode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PUNode.
func (o *PUNode) DeepCopyInto(out *PUNode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PUNode: %s", err))
	}

	*out = *target.(*PUNode)
}

// Validate valides the current information stored into the structure.
func (o *PUNode) Validate() error {

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
func (*PUNode) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PUNodeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PUNodeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PUNode) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PUNodeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PUNode) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "enforcementStatus":
		return o.EnforcementStatus
	case "images":
		return o.Images
	case "lastPokeTime":
		return o.LastPokeTime
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "status":
		return o.Status
	case "tags":
		return o.Tags
	case "type":
		return o.Type
	case "unreachable":
		return o.Unreachable
	}

	return nil
}

// PUNodeAttributesMap represents the map of attribute for PUNode.
var PUNodeAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"EnforcementStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcementStatus",
		Description:    `Enforcement status of the pu.`,
		Exposed:        true,
		Name:           "enforcementStatus",
		Type:           "string",
	},
	"Images": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Images",
		Description:    `Images of the pu.`,
		Exposed:        true,
		Name:           "images",
		SubType:        "string",
		Type:           "list",
	},
	"LastPokeTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastPokeTime",
		Description:    `Last poke time of the pu.`,
		Exposed:        true,
		Name:           "lastPokeTime",
		Type:           "time",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the pu.`,
		Exposed:        true,
		Name:           "name",
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the pu.`,
		Exposed:        true,
		Name:           "namespace",
		Type:           "string",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Status",
		Description:    `Status of the pu.`,
		Exposed:        true,
		Name:           "status",
		Type:           "string",
	},
	"Tags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `Tags of the pu.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "string",
		Type:           "list",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Type",
		Description:    `Type of the pu.`,
		Exposed:        true,
		Name:           "type",
		Type:           "string",
	},
	"Unreachable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Unreachable",
		Description:    `If true, the pu is unreachable.`,
		Exposed:        true,
		Name:           "unreachable",
		Type:           "boolean",
	},
}

// PUNodeLowerCaseAttributesMap represents the map of attribute for PUNode.
var PUNodeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"enforcementstatus": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcementStatus",
		Description:    `Enforcement status of the pu.`,
		Exposed:        true,
		Name:           "enforcementStatus",
		Type:           "string",
	},
	"images": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Images",
		Description:    `Images of the pu.`,
		Exposed:        true,
		Name:           "images",
		SubType:        "string",
		Type:           "list",
	},
	"lastpoketime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastPokeTime",
		Description:    `Last poke time of the pu.`,
		Exposed:        true,
		Name:           "lastPokeTime",
		Type:           "time",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the pu.`,
		Exposed:        true,
		Name:           "name",
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the pu.`,
		Exposed:        true,
		Name:           "namespace",
		Type:           "string",
	},
	"status": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Status",
		Description:    `Status of the pu.`,
		Exposed:        true,
		Name:           "status",
		Type:           "string",
	},
	"tags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `Tags of the pu.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "string",
		Type:           "list",
	},
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Type",
		Description:    `Type of the pu.`,
		Exposed:        true,
		Name:           "type",
		Type:           "string",
	},
	"unreachable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Unreachable",
		Description:    `If true, the pu is unreachable.`,
		Exposed:        true,
		Name:           "unreachable",
		Type:           "boolean",
	},
}

// SparsePUNodesList represents a list of SparsePUNodes
type SparsePUNodesList []*SparsePUNode

// Identity returns the identity of the objects in the list.
func (o SparsePUNodesList) Identity() elemental.Identity {

	return PUNodeIdentity
}

// Copy returns a pointer to a copy the SparsePUNodesList.
func (o SparsePUNodesList) Copy() elemental.Identifiables {

	copy := append(SparsePUNodesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePUNodesList.
func (o SparsePUNodesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePUNodesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePUNode))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePUNodesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePUNodesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePUNodesList converted to PUNodesList.
func (o SparsePUNodesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePUNodesList) Version() int {

	return 1
}

// SparsePUNode represents the sparse version of a punode.
type SparsePUNode struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Enforcement status of the pu.
	EnforcementStatus *string `json:"enforcementStatus,omitempty" bson:"-" mapstructure:"enforcementStatus,omitempty"`

	// Images of the pu.
	Images *[]string `json:"images,omitempty" bson:"-" mapstructure:"images,omitempty"`

	// Last poke time of the pu.
	LastPokeTime *time.Time `json:"lastPokeTime,omitempty" bson:"-" mapstructure:"lastPokeTime,omitempty"`

	// Name of the pu.
	Name *string `json:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// Namespace of the pu.
	Namespace *string `json:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// Status of the pu.
	Status *string `json:"status,omitempty" bson:"-" mapstructure:"status,omitempty"`

	// Tags of the pu.
	Tags *[]string `json:"tags,omitempty" bson:"-" mapstructure:"tags,omitempty"`

	// Type of the pu.
	Type *string `json:"type,omitempty" bson:"-" mapstructure:"type,omitempty"`

	// If true, the pu is unreachable.
	Unreachable *bool `json:"unreachable,omitempty" bson:"-" mapstructure:"unreachable,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSparsePUNode returns a new  SparsePUNode.
func NewSparsePUNode() *SparsePUNode {
	return &SparsePUNode{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePUNode) Identity() elemental.Identity {

	return PUNodeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePUNode) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePUNode) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparsePUNode) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePUNode) ToPlain() elemental.PlainIdentifiable {

	out := NewPUNode()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.EnforcementStatus != nil {
		out.EnforcementStatus = *o.EnforcementStatus
	}
	if o.Images != nil {
		out.Images = *o.Images
	}
	if o.LastPokeTime != nil {
		out.LastPokeTime = *o.LastPokeTime
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Status != nil {
		out.Status = *o.Status
	}
	if o.Tags != nil {
		out.Tags = *o.Tags
	}
	if o.Type != nil {
		out.Type = *o.Type
	}
	if o.Unreachable != nil {
		out.Unreachable = *o.Unreachable
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePUNode.
func (o *SparsePUNode) DeepCopy() *SparsePUNode {

	if o == nil {
		return nil
	}

	out := &SparsePUNode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePUNode.
func (o *SparsePUNode) DeepCopyInto(out *SparsePUNode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePUNode: %s", err))
	}

	*out = *target.(*SparsePUNode)
}
