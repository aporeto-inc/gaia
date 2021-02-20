package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudGraphNodeIdentity represents the Identity of the object.
var CloudGraphNodeIdentity = elemental.Identity{
	Name:     "cloudgraphnode",
	Category: "cloudgraphnodes",
	Package:  "pcn/infrastructure",
	Private:  false,
}

// CloudGraphNodesList represents a list of CloudGraphNodes
type CloudGraphNodesList []*CloudGraphNode

// Identity returns the identity of the objects in the list.
func (o CloudGraphNodesList) Identity() elemental.Identity {

	return CloudGraphNodeIdentity
}

// Copy returns a pointer to a copy the CloudGraphNodesList.
func (o CloudGraphNodesList) Copy() elemental.Identifiables {

	copy := append(CloudGraphNodesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudGraphNodesList.
func (o CloudGraphNodesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudGraphNodesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudGraphNode))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudGraphNodesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudGraphNodesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudGraphNodesList converted to SparseCloudGraphNodesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudGraphNodesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudGraphNodesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudGraphNode)
	}

	return out
}

// Version returns the version of the content.
func (o CloudGraphNodesList) Version() int {

	return 1
}

// CloudGraphNode represents the model of a cloudgraphnode
type CloudGraphNode struct {
	// The native ID of the node.
	NativeID string `json:"nativeID" msgpack:"nativeID" bson:"-" mapstructure:"nativeID,omitempty"`

	// The policies that were applied to this node for each destination.
	Policies map[string]*CloudGraphNodeAction `json:"policies" msgpack:"policies" bson:"-" mapstructure:"policies,omitempty"`

	// The list of route tables IDs that forwarding was based on for the internal path,
	// if routing was
	// performed.
	RouteTableIDs map[string]string `json:"routeTableIDs" msgpack:"routeTableIDs" bson:"-" mapstructure:"routeTableIDs,omitempty"`

	// The type of the node as a string.
	Type string `json:"type" msgpack:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudGraphNode returns a new *CloudGraphNode
func NewCloudGraphNode() *CloudGraphNode {

	return &CloudGraphNode{
		ModelVersion:  1,
		Policies:      map[string]*CloudGraphNodeAction{},
		RouteTableIDs: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudGraphNode) Identity() elemental.Identity {

	return CloudGraphNodeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudGraphNode) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudGraphNode) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphNode) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudGraphNode{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphNode) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudGraphNode{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudGraphNode) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudGraphNode) BleveType() string {

	return "cloudgraphnode"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudGraphNode) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudGraphNode) Doc() string {

	return `Returns a data structure representing the graph of all cloud nodes 
and their connections in a particular namespace.`
}

func (o *CloudGraphNode) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudGraphNode) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudGraphNode{
			NativeID:      &o.NativeID,
			Policies:      &o.Policies,
			RouteTableIDs: &o.RouteTableIDs,
			Type:          &o.Type,
		}
	}

	sp := &SparseCloudGraphNode{}
	for _, f := range fields {
		switch f {
		case "nativeID":
			sp.NativeID = &(o.NativeID)
		case "policies":
			sp.Policies = &(o.Policies)
		case "routeTableIDs":
			sp.RouteTableIDs = &(o.RouteTableIDs)
		case "type":
			sp.Type = &(o.Type)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudGraphNode to the object.
func (o *CloudGraphNode) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudGraphNode)
	if so.NativeID != nil {
		o.NativeID = *so.NativeID
	}
	if so.Policies != nil {
		o.Policies = *so.Policies
	}
	if so.RouteTableIDs != nil {
		o.RouteTableIDs = *so.RouteTableIDs
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
}

// DeepCopy returns a deep copy if the CloudGraphNode.
func (o *CloudGraphNode) DeepCopy() *CloudGraphNode {

	if o == nil {
		return nil
	}

	out := &CloudGraphNode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudGraphNode.
func (o *CloudGraphNode) DeepCopyInto(out *CloudGraphNode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudGraphNode: %s", err))
	}

	*out = *target.(*CloudGraphNode)
}

// Validate valides the current information stored into the structure.
func (o *CloudGraphNode) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Policies {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
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
func (*CloudGraphNode) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudGraphNodeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudGraphNodeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudGraphNode) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudGraphNodeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudGraphNode) ValueForAttribute(name string) interface{} {

	switch name {
	case "nativeID":
		return o.NativeID
	case "policies":
		return o.Policies
	case "routeTableIDs":
		return o.RouteTableIDs
	case "type":
		return o.Type
	}

	return nil
}

// CloudGraphNodeAttributesMap represents the map of attribute for CloudGraphNode.
var CloudGraphNodeAttributesMap = map[string]elemental.AttributeSpecification{
	"NativeID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NativeID",
		Description:    `The native ID of the node.`,
		Exposed:        true,
		Name:           "nativeID",
		Type:           "string",
	},
	"Policies": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Policies",
		Description:    `The policies that were applied to this node for each destination.`,
		Exposed:        true,
		Name:           "policies",
		SubType:        "cloudgraphnodeaction",
		Type:           "refMap",
	},
	"RouteTableIDs": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RouteTableIDs",
		Description: `The list of route tables IDs that forwarding was based on for the internal path,
if routing was
performed.`,
		Exposed: true,
		Name:    "routeTableIDs",
		SubType: "map[string]string",
		Type:    "external",
	},
	"Type": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Type",
		Description:    `The type of the node as a string.`,
		Exposed:        true,
		Name:           "type",
		Type:           "string",
	},
}

// CloudGraphNodeLowerCaseAttributesMap represents the map of attribute for CloudGraphNode.
var CloudGraphNodeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"nativeid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NativeID",
		Description:    `The native ID of the node.`,
		Exposed:        true,
		Name:           "nativeID",
		Type:           "string",
	},
	"policies": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Policies",
		Description:    `The policies that were applied to this node for each destination.`,
		Exposed:        true,
		Name:           "policies",
		SubType:        "cloudgraphnodeaction",
		Type:           "refMap",
	},
	"routetableids": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RouteTableIDs",
		Description: `The list of route tables IDs that forwarding was based on for the internal path,
if routing was
performed.`,
		Exposed: true,
		Name:    "routeTableIDs",
		SubType: "map[string]string",
		Type:    "external",
	},
	"type": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Type",
		Description:    `The type of the node as a string.`,
		Exposed:        true,
		Name:           "type",
		Type:           "string",
	},
}

// SparseCloudGraphNodesList represents a list of SparseCloudGraphNodes
type SparseCloudGraphNodesList []*SparseCloudGraphNode

// Identity returns the identity of the objects in the list.
func (o SparseCloudGraphNodesList) Identity() elemental.Identity {

	return CloudGraphNodeIdentity
}

// Copy returns a pointer to a copy the SparseCloudGraphNodesList.
func (o SparseCloudGraphNodesList) Copy() elemental.Identifiables {

	copy := append(SparseCloudGraphNodesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudGraphNodesList.
func (o SparseCloudGraphNodesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudGraphNodesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudGraphNode))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudGraphNodesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudGraphNodesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudGraphNodesList converted to CloudGraphNodesList.
func (o SparseCloudGraphNodesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudGraphNodesList) Version() int {

	return 1
}

// SparseCloudGraphNode represents the sparse version of a cloudgraphnode.
type SparseCloudGraphNode struct {
	// The native ID of the node.
	NativeID *string `json:"nativeID,omitempty" msgpack:"nativeID,omitempty" bson:"-" mapstructure:"nativeID,omitempty"`

	// The policies that were applied to this node for each destination.
	Policies *map[string]*CloudGraphNodeAction `json:"policies,omitempty" msgpack:"policies,omitempty" bson:"-" mapstructure:"policies,omitempty"`

	// The list of route tables IDs that forwarding was based on for the internal path,
	// if routing was
	// performed.
	RouteTableIDs *map[string]string `json:"routeTableIDs,omitempty" msgpack:"routeTableIDs,omitempty" bson:"-" mapstructure:"routeTableIDs,omitempty"`

	// The type of the node as a string.
	Type *string `json:"type,omitempty" msgpack:"type,omitempty" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudGraphNode returns a new  SparseCloudGraphNode.
func NewSparseCloudGraphNode() *SparseCloudGraphNode {
	return &SparseCloudGraphNode{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudGraphNode) Identity() elemental.Identity {

	return CloudGraphNodeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudGraphNode) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudGraphNode) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudGraphNode) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudGraphNode{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudGraphNode) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudGraphNode{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCloudGraphNode) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudGraphNode) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudGraphNode()
	if o.NativeID != nil {
		out.NativeID = *o.NativeID
	}
	if o.Policies != nil {
		out.Policies = *o.Policies
	}
	if o.RouteTableIDs != nil {
		out.RouteTableIDs = *o.RouteTableIDs
	}
	if o.Type != nil {
		out.Type = *o.Type
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCloudGraphNode.
func (o *SparseCloudGraphNode) DeepCopy() *SparseCloudGraphNode {

	if o == nil {
		return nil
	}

	out := &SparseCloudGraphNode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudGraphNode.
func (o *SparseCloudGraphNode) DeepCopyInto(out *SparseCloudGraphNode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudGraphNode: %s", err))
	}

	*out = *target.(*SparseCloudGraphNode)
}

type mongoAttributesCloudGraphNode struct {
}
type mongoAttributesSparseCloudGraphNode struct {
}
