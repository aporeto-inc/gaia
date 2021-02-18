package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudGraphNodePolicyTypeValue represents the possible values for attribute "policyType".
type CloudGraphNodePolicyTypeValue string

const (
	// CloudGraphNodePolicyTypeACL represents the value ACL.
	CloudGraphNodePolicyTypeACL CloudGraphNodePolicyTypeValue = "ACL"

	// CloudGraphNodePolicyTypeIdentity represents the value Identity.
	CloudGraphNodePolicyTypeIdentity CloudGraphNodePolicyTypeValue = "Identity"

	// CloudGraphNodePolicyTypeRoute represents the value Route.
	CloudGraphNodePolicyTypeRoute CloudGraphNodePolicyTypeValue = "Route"

	// CloudGraphNodePolicyTypeSG represents the value SG.
	CloudGraphNodePolicyTypeSG CloudGraphNodePolicyTypeValue = "SG"
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
	// Refers to the nodes of the map.
	Node *CloudNode `json:"node" msgpack:"node" bson:"-" mapstructure:"node,omitempty"`

	// The ID of the policies that were used in the path.
	PolicyIDs map[string]string `json:"policyIDs" msgpack:"policyIDs" bson:"-" mapstructure:"policyIDs,omitempty"`

	// The namespace of the policy.
	PolicyNamespace string `json:"policyNamespace" msgpack:"policyNamespace" bson:"-" mapstructure:"policyNamespace,omitempty"`

	// The type of policy that determined forwarding from this node.
	PolicyType CloudGraphNodePolicyTypeValue `json:"policyType" msgpack:"policyType" bson:"-" mapstructure:"policyType,omitempty"`

	// The list of route tables IDs that forwarding was based on for the internal path,
	// if routing was
	// performed.
	RouteTableIDs map[string]string `json:"routeTableIDs" msgpack:"routeTableIDs" bson:"-" mapstructure:"routeTableIDs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudGraphNode returns a new *CloudGraphNode
func NewCloudGraphNode() *CloudGraphNode {

	return &CloudGraphNode{
		ModelVersion:  1,
		Node:          NewCloudNode(),
		PolicyIDs:     map[string]string{},
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
			Node:            o.Node,
			PolicyIDs:       &o.PolicyIDs,
			PolicyNamespace: &o.PolicyNamespace,
			PolicyType:      &o.PolicyType,
			RouteTableIDs:   &o.RouteTableIDs,
		}
	}

	sp := &SparseCloudGraphNode{}
	for _, f := range fields {
		switch f {
		case "node":
			sp.Node = o.Node
		case "policyIDs":
			sp.PolicyIDs = &(o.PolicyIDs)
		case "policyNamespace":
			sp.PolicyNamespace = &(o.PolicyNamespace)
		case "policyType":
			sp.PolicyType = &(o.PolicyType)
		case "routeTableIDs":
			sp.RouteTableIDs = &(o.RouteTableIDs)
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
	if so.Node != nil {
		o.Node = so.Node
	}
	if so.PolicyIDs != nil {
		o.PolicyIDs = *so.PolicyIDs
	}
	if so.PolicyNamespace != nil {
		o.PolicyNamespace = *so.PolicyNamespace
	}
	if so.PolicyType != nil {
		o.PolicyType = *so.PolicyType
	}
	if so.RouteTableIDs != nil {
		o.RouteTableIDs = *so.RouteTableIDs
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

	if o.Node != nil {
		elemental.ResetDefaultForZeroValues(o.Node)
		if err := o.Node.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateStringInList("policyType", string(o.PolicyType), []string{"Route", "ACL", "SG", "Identity"}, true); err != nil {
		errors = errors.Append(err)
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
	case "node":
		return o.Node
	case "policyIDs":
		return o.PolicyIDs
	case "policyNamespace":
		return o.PolicyNamespace
	case "policyType":
		return o.PolicyType
	case "routeTableIDs":
		return o.RouteTableIDs
	}

	return nil
}

// CloudGraphNodeAttributesMap represents the map of attribute for CloudGraphNode.
var CloudGraphNodeAttributesMap = map[string]elemental.AttributeSpecification{
	"Node": {
		AllowedChoices: []string{},
		ConvertedName:  "Node",
		Description:    `Refers to the nodes of the map.`,
		Exposed:        true,
		Name:           "node",
		ReadOnly:       true,
		SubType:        "cloudnode",
		Type:           "ref",
	},
	"PolicyIDs": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PolicyIDs",
		Description:    `The ID of the policies that were used in the path.`,
		Exposed:        true,
		Name:           "policyIDs",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"PolicyNamespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PolicyNamespace",
		Description:    `The namespace of the policy.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Type:           "string",
	},
	"PolicyType": {
		AllowedChoices: []string{"Route", "ACL", "SG", "Identity"},
		Autogenerated:  true,
		ConvertedName:  "PolicyType",
		Description:    `The type of policy that determined forwarding from this node.`,
		Exposed:        true,
		Name:           "policyType",
		Type:           "enum",
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
}

// CloudGraphNodeLowerCaseAttributesMap represents the map of attribute for CloudGraphNode.
var CloudGraphNodeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"node": {
		AllowedChoices: []string{},
		ConvertedName:  "Node",
		Description:    `Refers to the nodes of the map.`,
		Exposed:        true,
		Name:           "node",
		ReadOnly:       true,
		SubType:        "cloudnode",
		Type:           "ref",
	},
	"policyids": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PolicyIDs",
		Description:    `The ID of the policies that were used in the path.`,
		Exposed:        true,
		Name:           "policyIDs",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"policynamespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PolicyNamespace",
		Description:    `The namespace of the policy.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Type:           "string",
	},
	"policytype": {
		AllowedChoices: []string{"Route", "ACL", "SG", "Identity"},
		Autogenerated:  true,
		ConvertedName:  "PolicyType",
		Description:    `The type of policy that determined forwarding from this node.`,
		Exposed:        true,
		Name:           "policyType",
		Type:           "enum",
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
	// Refers to the nodes of the map.
	Node *CloudNode `json:"node,omitempty" msgpack:"node,omitempty" bson:"-" mapstructure:"node,omitempty"`

	// The ID of the policies that were used in the path.
	PolicyIDs *map[string]string `json:"policyIDs,omitempty" msgpack:"policyIDs,omitempty" bson:"-" mapstructure:"policyIDs,omitempty"`

	// The namespace of the policy.
	PolicyNamespace *string `json:"policyNamespace,omitempty" msgpack:"policyNamespace,omitempty" bson:"-" mapstructure:"policyNamespace,omitempty"`

	// The type of policy that determined forwarding from this node.
	PolicyType *CloudGraphNodePolicyTypeValue `json:"policyType,omitempty" msgpack:"policyType,omitempty" bson:"-" mapstructure:"policyType,omitempty"`

	// The list of route tables IDs that forwarding was based on for the internal path,
	// if routing was
	// performed.
	RouteTableIDs *map[string]string `json:"routeTableIDs,omitempty" msgpack:"routeTableIDs,omitempty" bson:"-" mapstructure:"routeTableIDs,omitempty"`

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
	if o.Node != nil {
		out.Node = o.Node
	}
	if o.PolicyIDs != nil {
		out.PolicyIDs = *o.PolicyIDs
	}
	if o.PolicyNamespace != nil {
		out.PolicyNamespace = *o.PolicyNamespace
	}
	if o.PolicyType != nil {
		out.PolicyType = *o.PolicyType
	}
	if o.RouteTableIDs != nil {
		out.RouteTableIDs = *o.RouteTableIDs
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
