package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudGraphIdentity represents the Identity of the object.
var CloudGraphIdentity = elemental.Identity{
	Name:     "cloudgraph",
	Category: "cloudgraphs",
	Package:  "pcn/infrastructure",
	Private:  false,
}

// CloudGraphsList represents a list of CloudGraphs
type CloudGraphsList []*CloudGraph

// Identity returns the identity of the objects in the list.
func (o CloudGraphsList) Identity() elemental.Identity {

	return CloudGraphIdentity
}

// Copy returns a pointer to a copy the CloudGraphsList.
func (o CloudGraphsList) Copy() elemental.Identifiables {

	copy := append(CloudGraphsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudGraphsList.
func (o CloudGraphsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudGraphsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudGraph))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudGraphsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudGraphsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudGraphsList converted to SparseCloudGraphsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudGraphsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudGraphsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudGraph)
	}

	return out
}

// Version returns the version of the content.
func (o CloudGraphsList) Version() int {

	return 1
}

// CloudGraph represents the model of a cloudgraph
type CloudGraph struct {
	// The edges of the map.
	Edges map[string]*CloudGraphEdge `json:"edges" msgpack:"edges" bson:"-" mapstructure:"edges,omitempty"`

	// Refers to the nodes of the map.
	Nodes map[string]*CloudGraphNode `json:"nodes" msgpack:"nodes" bson:"-" mapstructure:"nodes,omitempty"`

	// The cloud network query that should be used. This requires a POST operation on
	// the object.
	Query *CloudNetworkQuery `json:"query" msgpack:"query" bson:"-" mapstructure:"query,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudGraph returns a new *CloudGraph
func NewCloudGraph() *CloudGraph {

	return &CloudGraph{
		ModelVersion: 1,
		Edges:        map[string]*CloudGraphEdge{},
		Nodes:        map[string]*CloudGraphNode{},
		Query:        NewCloudNetworkQuery(),
	}
}

// Identity returns the Identity of the object.
func (o *CloudGraph) Identity() elemental.Identity {

	return CloudGraphIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudGraph) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudGraph) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraph) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudGraph{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraph) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudGraph{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudGraph) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudGraph) BleveType() string {

	return "cloudgraph"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudGraph) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudGraph) Doc() string {

	return `Returns a data structure representing the graph of all cloud nodes 
and their connections in a particular namespace.`
}

func (o *CloudGraph) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudGraph) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudGraph{
			Edges: &o.Edges,
			Nodes: &o.Nodes,
			Query: o.Query,
		}
	}

	sp := &SparseCloudGraph{}
	for _, f := range fields {
		switch f {
		case "edges":
			sp.Edges = &(o.Edges)
		case "nodes":
			sp.Nodes = &(o.Nodes)
		case "query":
			sp.Query = o.Query
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudGraph to the object.
func (o *CloudGraph) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudGraph)
	if so.Edges != nil {
		o.Edges = *so.Edges
	}
	if so.Nodes != nil {
		o.Nodes = *so.Nodes
	}
	if so.Query != nil {
		o.Query = so.Query
	}
}

// DeepCopy returns a deep copy if the CloudGraph.
func (o *CloudGraph) DeepCopy() *CloudGraph {

	if o == nil {
		return nil
	}

	out := &CloudGraph{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudGraph.
func (o *CloudGraph) DeepCopyInto(out *CloudGraph) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudGraph: %s", err))
	}

	*out = *target.(*CloudGraph)
}

// Validate valides the current information stored into the structure.
func (o *CloudGraph) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Edges {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.Nodes {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if o.Query != nil {
		elemental.ResetDefaultForZeroValues(o.Query)
		if err := o.Query.Validate(); err != nil {
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
func (*CloudGraph) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudGraphAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudGraphLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudGraph) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudGraphAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudGraph) ValueForAttribute(name string) interface{} {

	switch name {
	case "edges":
		return o.Edges
	case "nodes":
		return o.Nodes
	case "query":
		return o.Query
	}

	return nil
}

// CloudGraphAttributesMap represents the map of attribute for CloudGraph.
var CloudGraphAttributesMap = map[string]elemental.AttributeSpecification{
	"Edges": {
		AllowedChoices: []string{},
		ConvertedName:  "Edges",
		Description:    `The edges of the map.`,
		Exposed:        true,
		Name:           "edges",
		ReadOnly:       true,
		SubType:        "cloudgraphedge",
		Type:           "refMap",
	},
	"Nodes": {
		AllowedChoices: []string{},
		ConvertedName:  "Nodes",
		Description:    `Refers to the nodes of the map.`,
		Exposed:        true,
		Name:           "nodes",
		ReadOnly:       true,
		SubType:        "cloudgraphnode",
		Type:           "refMap",
	},
	"Query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description: `The cloud network query that should be used. This requires a POST operation on
the object.`,
		Exposed: true,
		Name:    "query",
		SubType: "cloudnetworkquery",
		Type:    "ref",
	},
}

// CloudGraphLowerCaseAttributesMap represents the map of attribute for CloudGraph.
var CloudGraphLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"edges": {
		AllowedChoices: []string{},
		ConvertedName:  "Edges",
		Description:    `The edges of the map.`,
		Exposed:        true,
		Name:           "edges",
		ReadOnly:       true,
		SubType:        "cloudgraphedge",
		Type:           "refMap",
	},
	"nodes": {
		AllowedChoices: []string{},
		ConvertedName:  "Nodes",
		Description:    `Refers to the nodes of the map.`,
		Exposed:        true,
		Name:           "nodes",
		ReadOnly:       true,
		SubType:        "cloudgraphnode",
		Type:           "refMap",
	},
	"query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description: `The cloud network query that should be used. This requires a POST operation on
the object.`,
		Exposed: true,
		Name:    "query",
		SubType: "cloudnetworkquery",
		Type:    "ref",
	},
}

// SparseCloudGraphsList represents a list of SparseCloudGraphs
type SparseCloudGraphsList []*SparseCloudGraph

// Identity returns the identity of the objects in the list.
func (o SparseCloudGraphsList) Identity() elemental.Identity {

	return CloudGraphIdentity
}

// Copy returns a pointer to a copy the SparseCloudGraphsList.
func (o SparseCloudGraphsList) Copy() elemental.Identifiables {

	copy := append(SparseCloudGraphsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudGraphsList.
func (o SparseCloudGraphsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudGraphsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudGraph))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudGraphsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudGraphsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudGraphsList converted to CloudGraphsList.
func (o SparseCloudGraphsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudGraphsList) Version() int {

	return 1
}

// SparseCloudGraph represents the sparse version of a cloudgraph.
type SparseCloudGraph struct {
	// The edges of the map.
	Edges *map[string]*CloudGraphEdge `json:"edges,omitempty" msgpack:"edges,omitempty" bson:"-" mapstructure:"edges,omitempty"`

	// Refers to the nodes of the map.
	Nodes *map[string]*CloudGraphNode `json:"nodes,omitempty" msgpack:"nodes,omitempty" bson:"-" mapstructure:"nodes,omitempty"`

	// The cloud network query that should be used. This requires a POST operation on
	// the object.
	Query *CloudNetworkQuery `json:"query,omitempty" msgpack:"query,omitempty" bson:"-" mapstructure:"query,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudGraph returns a new  SparseCloudGraph.
func NewSparseCloudGraph() *SparseCloudGraph {
	return &SparseCloudGraph{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudGraph) Identity() elemental.Identity {

	return CloudGraphIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudGraph) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudGraph) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudGraph) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudGraph{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudGraph) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudGraph{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCloudGraph) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudGraph) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudGraph()
	if o.Edges != nil {
		out.Edges = *o.Edges
	}
	if o.Nodes != nil {
		out.Nodes = *o.Nodes
	}
	if o.Query != nil {
		out.Query = o.Query
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCloudGraph.
func (o *SparseCloudGraph) DeepCopy() *SparseCloudGraph {

	if o == nil {
		return nil
	}

	out := &SparseCloudGraph{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudGraph.
func (o *SparseCloudGraph) DeepCopyInto(out *SparseCloudGraph) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudGraph: %s", err))
	}

	*out = *target.(*SparseCloudGraph)
}

type mongoAttributesCloudGraph struct {
}
type mongoAttributesSparseCloudGraph struct {
}
