package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudTopologyIdentity represents the Identity of the object.
var CloudTopologyIdentity = elemental.Identity{
	Name:     "cloudtopology",
	Category: "cloudtopologies",
	Package:  "pcn/infrastructure",
	Private:  false,
}

// CloudTopologiesList represents a list of CloudTopologies
type CloudTopologiesList []*CloudTopology

// Identity returns the identity of the objects in the list.
func (o CloudTopologiesList) Identity() elemental.Identity {

	return CloudTopologyIdentity
}

// Copy returns a pointer to a copy the CloudTopologiesList.
func (o CloudTopologiesList) Copy() elemental.Identifiables {

	copy := append(CloudTopologiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudTopologiesList.
func (o CloudTopologiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudTopologiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudTopology))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudTopologiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudTopologiesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudTopologiesList converted to SparseCloudTopologiesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudTopologiesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudTopologiesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudTopology)
	}

	return out
}

// Version returns the version of the content.
func (o CloudTopologiesList) Version() int {

	return 1
}

// CloudTopology represents the model of a cloudtopology
type CloudTopology struct {
	// The edges of the map.
	Edges map[string]*CloudGraphEdge `json:"edges" msgpack:"edges" bson:"-" mapstructure:"edges,omitempty"`

	// Refers to the nodes of the map.
	Nodes map[string]*CloudNode `json:"nodes" msgpack:"nodes" bson:"-" mapstructure:"nodes,omitempty"`

	// The target VPCs for the topology. If empty, it will return the topology for all
	// VPCs.
	TargetVPCs []string `json:"targetVPCs" msgpack:"targetVPCs" bson:"-" mapstructure:"targetVPCs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudTopology returns a new *CloudTopology
func NewCloudTopology() *CloudTopology {

	return &CloudTopology{
		ModelVersion: 1,
		Edges:        map[string]*CloudGraphEdge{},
		Nodes:        map[string]*CloudNode{},
		TargetVPCs:   []string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudTopology) Identity() elemental.Identity {

	return CloudTopologyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudTopology) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudTopology) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudTopology) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudTopology{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudTopology) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudTopology{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudTopology) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudTopology) BleveType() string {

	return "cloudtopology"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudTopology) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudTopology) Doc() string {

	return `Returns the full topology of all nodes and their relationships.`
}

func (o *CloudTopology) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudTopology) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudTopology{
			Edges:      &o.Edges,
			Nodes:      &o.Nodes,
			TargetVPCs: &o.TargetVPCs,
		}
	}

	sp := &SparseCloudTopology{}
	for _, f := range fields {
		switch f {
		case "edges":
			sp.Edges = &(o.Edges)
		case "nodes":
			sp.Nodes = &(o.Nodes)
		case "targetVPCs":
			sp.TargetVPCs = &(o.TargetVPCs)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudTopology to the object.
func (o *CloudTopology) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudTopology)
	if so.Edges != nil {
		o.Edges = *so.Edges
	}
	if so.Nodes != nil {
		o.Nodes = *so.Nodes
	}
	if so.TargetVPCs != nil {
		o.TargetVPCs = *so.TargetVPCs
	}
}

// DeepCopy returns a deep copy if the CloudTopology.
func (o *CloudTopology) DeepCopy() *CloudTopology {

	if o == nil {
		return nil
	}

	out := &CloudTopology{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudTopology.
func (o *CloudTopology) DeepCopyInto(out *CloudTopology) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudTopology: %s", err))
	}

	*out = *target.(*CloudTopology)
}

// Validate valides the current information stored into the structure.
func (o *CloudTopology) Validate() error {

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

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*CloudTopology) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudTopologyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudTopologyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudTopology) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudTopologyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudTopology) ValueForAttribute(name string) interface{} {

	switch name {
	case "edges":
		return o.Edges
	case "nodes":
		return o.Nodes
	case "targetVPCs":
		return o.TargetVPCs
	}

	return nil
}

// CloudTopologyAttributesMap represents the map of attribute for CloudTopology.
var CloudTopologyAttributesMap = map[string]elemental.AttributeSpecification{
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
		SubType:        "cloudnode",
		Type:           "refMap",
	},
	"TargetVPCs": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetVPCs",
		Description: `The target VPCs for the topology. If empty, it will return the topology for all
VPCs.`,
		Exposed: true,
		Name:    "targetVPCs",
		SubType: "string",
		Type:    "list",
	},
}

// CloudTopologyLowerCaseAttributesMap represents the map of attribute for CloudTopology.
var CloudTopologyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		SubType:        "cloudnode",
		Type:           "refMap",
	},
	"targetvpcs": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetVPCs",
		Description: `The target VPCs for the topology. If empty, it will return the topology for all
VPCs.`,
		Exposed: true,
		Name:    "targetVPCs",
		SubType: "string",
		Type:    "list",
	},
}

// SparseCloudTopologiesList represents a list of SparseCloudTopologies
type SparseCloudTopologiesList []*SparseCloudTopology

// Identity returns the identity of the objects in the list.
func (o SparseCloudTopologiesList) Identity() elemental.Identity {

	return CloudTopologyIdentity
}

// Copy returns a pointer to a copy the SparseCloudTopologiesList.
func (o SparseCloudTopologiesList) Copy() elemental.Identifiables {

	copy := append(SparseCloudTopologiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudTopologiesList.
func (o SparseCloudTopologiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudTopologiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudTopology))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudTopologiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudTopologiesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudTopologiesList converted to CloudTopologiesList.
func (o SparseCloudTopologiesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudTopologiesList) Version() int {

	return 1
}

// SparseCloudTopology represents the sparse version of a cloudtopology.
type SparseCloudTopology struct {
	// The edges of the map.
	Edges *map[string]*CloudGraphEdge `json:"edges,omitempty" msgpack:"edges,omitempty" bson:"-" mapstructure:"edges,omitempty"`

	// Refers to the nodes of the map.
	Nodes *map[string]*CloudNode `json:"nodes,omitempty" msgpack:"nodes,omitempty" bson:"-" mapstructure:"nodes,omitempty"`

	// The target VPCs for the topology. If empty, it will return the topology for all
	// VPCs.
	TargetVPCs *[]string `json:"targetVPCs,omitempty" msgpack:"targetVPCs,omitempty" bson:"-" mapstructure:"targetVPCs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudTopology returns a new  SparseCloudTopology.
func NewSparseCloudTopology() *SparseCloudTopology {
	return &SparseCloudTopology{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudTopology) Identity() elemental.Identity {

	return CloudTopologyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudTopology) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudTopology) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudTopology) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudTopology{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudTopology) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudTopology{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCloudTopology) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudTopology) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudTopology()
	if o.Edges != nil {
		out.Edges = *o.Edges
	}
	if o.Nodes != nil {
		out.Nodes = *o.Nodes
	}
	if o.TargetVPCs != nil {
		out.TargetVPCs = *o.TargetVPCs
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCloudTopology.
func (o *SparseCloudTopology) DeepCopy() *SparseCloudTopology {

	if o == nil {
		return nil
	}

	out := &SparseCloudTopology{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudTopology.
func (o *SparseCloudTopology) DeepCopyInto(out *SparseCloudTopology) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudTopology: %s", err))
	}

	*out = *target.(*SparseCloudTopology)
}

type mongoAttributesCloudTopology struct {
}
type mongoAttributesSparseCloudTopology struct {
}
