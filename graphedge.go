package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// GraphEdgeDestinationTypeValue represents the possible values for attribute "destinationType".
type GraphEdgeDestinationTypeValue string

const (
	// GraphEdgeDestinationTypeExternalNetwork represents the value ExternalNetwork.
	GraphEdgeDestinationTypeExternalNetwork GraphEdgeDestinationTypeValue = "ExternalNetwork"

	// GraphEdgeDestinationTypeNode represents the value Node.
	GraphEdgeDestinationTypeNode GraphEdgeDestinationTypeValue = "Node"

	// GraphEdgeDestinationTypeProcessingUnit represents the value ProcessingUnit.
	GraphEdgeDestinationTypeProcessingUnit GraphEdgeDestinationTypeValue = "ProcessingUnit"
)

// GraphEdgeSourceTypeValue represents the possible values for attribute "sourceType".
type GraphEdgeSourceTypeValue string

const (
	// GraphEdgeSourceTypeExternalNetwork represents the value ExternalNetwork.
	GraphEdgeSourceTypeExternalNetwork GraphEdgeSourceTypeValue = "ExternalNetwork"

	// GraphEdgeSourceTypeNode represents the value Node.
	GraphEdgeSourceTypeNode GraphEdgeSourceTypeValue = "Node"

	// GraphEdgeSourceTypeProcessingUnit represents the value ProcessingUnit.
	GraphEdgeSourceTypeProcessingUnit GraphEdgeSourceTypeValue = "ProcessingUnit"
)

// GraphEdgeIdentity represents the Identity of the object.
var GraphEdgeIdentity = elemental.Identity{
	Name:     "graphedge",
	Category: "graphedges",
	Package:  "meteor",
	Private:  true,
}

// GraphEdgesList represents a list of GraphEdges
type GraphEdgesList []*GraphEdge

// Identity returns the identity of the objects in the list.
func (o GraphEdgesList) Identity() elemental.Identity {

	return GraphEdgeIdentity
}

// Copy returns a pointer to a copy the GraphEdgesList.
func (o GraphEdgesList) Copy() elemental.Identifiables {

	copy := append(GraphEdgesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the GraphEdgesList.
func (o GraphEdgesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(GraphEdgesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*GraphEdge))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o GraphEdgesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o GraphEdgesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the GraphEdgesList converted to SparseGraphEdgesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o GraphEdgesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseGraphEdgesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseGraphEdge)
	}

	return out
}

// Version returns the version of the content.
func (o GraphEdgesList) Version() int {

	return 1
}

// GraphEdge represents the model of a graphedge
type GraphEdge struct {
	// DB Identifier of the edge.
	ID string `json:"-" msgpack:"-" bson:"_id" mapstructure:"-,omitempty"`

	// Number of accepted flows in the edge.
	AcceptedFlows int `json:"acceptedFlows" msgpack:"acceptedFlows" bson:"acceptedflows" mapstructure:"acceptedFlows,omitempty"`

	// Date on which the edge has been inserted.
	CreateTime time.Time `json:"-" msgpack:"-" bson:"createtime" mapstructure:"-,omitempty"`

	// ID of the destination `GraphNode` of the edge.
	DestinationID string `json:"destinationID" msgpack:"destinationID" bson:"destinationid" mapstructure:"destinationID,omitempty"`

	// Type of the destination `GraphNode` of the edge.
	DestinationType GraphEdgeDestinationTypeValue `json:"destinationType" msgpack:"destinationType" bson:"destinationtype" mapstructure:"destinationType,omitempty"`

	// The number of encrypted flows in the edge.
	Encrypted int `json:"encrypted" msgpack:"encrypted" bson:"encrypted" mapstructure:"encrypted,omitempty"`

	// Contains the date when the edge was first seen.
	FirstSeen time.Time `json:"firstSeen" msgpack:"firstSeen" bson:"firstseen" mapstructure:"firstSeen,omitempty"`

	// Identifier of the edge.
	FlowID string `json:"ID" msgpack:"ID" bson:"flowid" mapstructure:"ID,omitempty"`

	// Contains the date when the edge was last seen.
	LastSeen time.Time `json:"lastSeen" msgpack:"lastSeen" bson:"lastseen" mapstructure:"lastSeen,omitempty"`

	// Namespace of object represented by the node.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Number of accepted observed flows.
	ObservedAcceptedFlows int `json:"observedAcceptedFlows" msgpack:"observedAcceptedFlows" bson:"observedacceptedflows" mapstructure:"observedAcceptedFlows,omitempty"`

	// Number of encrypted observed flows.
	ObservedEncrypted int `json:"observedEncrypted" msgpack:"observedEncrypted" bson:"observedencrypted" mapstructure:"observedEncrypted,omitempty"`

	// Number of rejected observed flows.
	ObservedRejectedFlows int `json:"observedRejectedFlows" msgpack:"observedRejectedFlows" bson:"observedrejectedflows" mapstructure:"observedRejectedFlows,omitempty"`

	// Number of rejected flows in the edge.
	RejectedFlows int `json:"rejectedFlows" msgpack:"rejectedFlows" bson:"rejectedflows" mapstructure:"rejectedFlows,omitempty"`

	// ID of the source `GraphNode` of the edge.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"sourceid" mapstructure:"sourceID,omitempty"`

	// Type of the source `GraphNode` of the edge.
	SourceType GraphEdgeSourceTypeValue `json:"sourceType" msgpack:"sourceType" bson:"sourcetype" mapstructure:"sourceType,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewGraphEdge returns a new *GraphEdge
func NewGraphEdge() *GraphEdge {

	return &GraphEdge{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *GraphEdge) Identity() elemental.Identity {

	return GraphEdgeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GraphEdge) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GraphEdge) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *GraphEdge) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *GraphEdge) BleveType() string {

	return "graphedge"
}

// DefaultOrder returns the list of default ordering fields.
func (o *GraphEdge) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *GraphEdge) Doc() string {

	return `Represents an edge from the dependency map.`
}

func (o *GraphEdge) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetZHash returns the ZHash of the receiver.
func (o *GraphEdge) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *GraphEdge) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *GraphEdge) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *GraphEdge) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *GraphEdge) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseGraphEdge{
			ID:                    &o.ID,
			AcceptedFlows:         &o.AcceptedFlows,
			CreateTime:            &o.CreateTime,
			DestinationID:         &o.DestinationID,
			DestinationType:       &o.DestinationType,
			Encrypted:             &o.Encrypted,
			FirstSeen:             &o.FirstSeen,
			FlowID:                &o.FlowID,
			LastSeen:              &o.LastSeen,
			Namespace:             &o.Namespace,
			ObservedAcceptedFlows: &o.ObservedAcceptedFlows,
			ObservedEncrypted:     &o.ObservedEncrypted,
			ObservedRejectedFlows: &o.ObservedRejectedFlows,
			RejectedFlows:         &o.RejectedFlows,
			SourceID:              &o.SourceID,
			SourceType:            &o.SourceType,
			ZHash:                 &o.ZHash,
			Zone:                  &o.Zone,
		}
	}

	sp := &SparseGraphEdge{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "acceptedFlows":
			sp.AcceptedFlows = &(o.AcceptedFlows)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "destinationID":
			sp.DestinationID = &(o.DestinationID)
		case "destinationType":
			sp.DestinationType = &(o.DestinationType)
		case "encrypted":
			sp.Encrypted = &(o.Encrypted)
		case "firstSeen":
			sp.FirstSeen = &(o.FirstSeen)
		case "flowID":
			sp.FlowID = &(o.FlowID)
		case "lastSeen":
			sp.LastSeen = &(o.LastSeen)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "observedAcceptedFlows":
			sp.ObservedAcceptedFlows = &(o.ObservedAcceptedFlows)
		case "observedEncrypted":
			sp.ObservedEncrypted = &(o.ObservedEncrypted)
		case "observedRejectedFlows":
			sp.ObservedRejectedFlows = &(o.ObservedRejectedFlows)
		case "rejectedFlows":
			sp.RejectedFlows = &(o.RejectedFlows)
		case "sourceID":
			sp.SourceID = &(o.SourceID)
		case "sourceType":
			sp.SourceType = &(o.SourceType)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseGraphEdge to the object.
func (o *GraphEdge) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseGraphEdge)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.AcceptedFlows != nil {
		o.AcceptedFlows = *so.AcceptedFlows
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.DestinationID != nil {
		o.DestinationID = *so.DestinationID
	}
	if so.DestinationType != nil {
		o.DestinationType = *so.DestinationType
	}
	if so.Encrypted != nil {
		o.Encrypted = *so.Encrypted
	}
	if so.FirstSeen != nil {
		o.FirstSeen = *so.FirstSeen
	}
	if so.FlowID != nil {
		o.FlowID = *so.FlowID
	}
	if so.LastSeen != nil {
		o.LastSeen = *so.LastSeen
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.ObservedAcceptedFlows != nil {
		o.ObservedAcceptedFlows = *so.ObservedAcceptedFlows
	}
	if so.ObservedEncrypted != nil {
		o.ObservedEncrypted = *so.ObservedEncrypted
	}
	if so.ObservedRejectedFlows != nil {
		o.ObservedRejectedFlows = *so.ObservedRejectedFlows
	}
	if so.RejectedFlows != nil {
		o.RejectedFlows = *so.RejectedFlows
	}
	if so.SourceID != nil {
		o.SourceID = *so.SourceID
	}
	if so.SourceType != nil {
		o.SourceType = *so.SourceType
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the GraphEdge.
func (o *GraphEdge) DeepCopy() *GraphEdge {

	if o == nil {
		return nil
	}

	out := &GraphEdge{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *GraphEdge.
func (o *GraphEdge) DeepCopyInto(out *GraphEdge) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy GraphEdge: %s", err))
	}

	*out = *target.(*GraphEdge)
}

// Validate valides the current information stored into the structure.
func (o *GraphEdge) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("destinationType", string(o.DestinationType), []string{"ProcessingUnit", "ExternalNetwork", "Node"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("sourceType", string(o.SourceType), []string{"ProcessingUnit", "ExternalNetwork", "Node"}, false); err != nil {
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
func (*GraphEdge) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := GraphEdgeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return GraphEdgeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*GraphEdge) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return GraphEdgeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *GraphEdge) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "acceptedFlows":
		return o.AcceptedFlows
	case "createTime":
		return o.CreateTime
	case "destinationID":
		return o.DestinationID
	case "destinationType":
		return o.DestinationType
	case "encrypted":
		return o.Encrypted
	case "firstSeen":
		return o.FirstSeen
	case "flowID":
		return o.FlowID
	case "lastSeen":
		return o.LastSeen
	case "namespace":
		return o.Namespace
	case "observedAcceptedFlows":
		return o.ObservedAcceptedFlows
	case "observedEncrypted":
		return o.ObservedEncrypted
	case "observedRejectedFlows":
		return o.ObservedRejectedFlows
	case "rejectedFlows":
		return o.RejectedFlows
	case "sourceID":
		return o.SourceID
	case "sourceType":
		return o.SourceType
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// GraphEdgeAttributesMap represents the map of attribute for GraphEdge.
var GraphEdgeAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `DB Identifier of the edge.`,
		Identifier:     true,
		Name:           "ID",
		Stored:         true,
		Type:           "string",
	},
	"AcceptedFlows": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AcceptedFlows",
		Description:    `Number of accepted flows in the edge.`,
		Exposed:        true,
		Name:           "acceptedFlows",
		Stored:         true,
		Type:           "integer",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CreateTime",
		Description:    `Date on which the edge has been inserted.`,
		Name:           "createTime",
		Stored:         true,
		Type:           "time",
	},
	"DestinationID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "destinationID",
		Stored:         true,
		Type:           "string",
	},
	"DestinationType": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Node"},
		ConvertedName:  "DestinationType",
		Description:    `Type of the destination ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "destinationType",
		Stored:         true,
		Type:           "enum",
	},
	"Encrypted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Encrypted",
		Description:    `The number of encrypted flows in the edge.`,
		Exposed:        true,
		Name:           "encrypted",
		Stored:         true,
		Type:           "integer",
	},
	"FirstSeen": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FirstSeen",
		Description:    `Contains the date when the edge was first seen.`,
		Exposed:        true,
		Name:           "firstSeen",
		Stored:         true,
		Type:           "time",
	},
	"FlowID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FlowID",
		Description:    `Identifier of the edge.`,
		Exposed:        true,
		Name:           "flowID",
		Stored:         true,
		Type:           "string",
	},
	"LastSeen": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastSeen",
		Description:    `Contains the date when the edge was last seen.`,
		Exposed:        true,
		Name:           "lastSeen",
		Stored:         true,
		Type:           "time",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of object represented by the node.`,
		Exposed:        true,
		Name:           "namespace",
		Stored:         true,
		Type:           "string",
	},
	"ObservedAcceptedFlows": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedAcceptedFlows",
		Description:    `Number of accepted observed flows.`,
		Exposed:        true,
		Name:           "observedAcceptedFlows",
		Stored:         true,
		Type:           "integer",
	},
	"ObservedEncrypted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedEncrypted",
		Description:    `Number of encrypted observed flows.`,
		Exposed:        true,
		Name:           "observedEncrypted",
		Stored:         true,
		Type:           "integer",
	},
	"ObservedRejectedFlows": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedRejectedFlows",
		Description:    `Number of rejected observed flows.`,
		Exposed:        true,
		Name:           "observedRejectedFlows",
		Stored:         true,
		Type:           "integer",
	},
	"RejectedFlows": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RejectedFlows",
		Description:    `Number of rejected flows in the edge.`,
		Exposed:        true,
		Name:           "rejectedFlows",
		Stored:         true,
		Type:           "integer",
	},
	"SourceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "sourceID",
		Stored:         true,
		Type:           "string",
	},
	"SourceType": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Node"},
		ConvertedName:  "SourceType",
		Description:    `Type of the source ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "sourceType",
		Stored:         true,
		Type:           "enum",
	},
	"ZHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// GraphEdgeLowerCaseAttributesMap represents the map of attribute for GraphEdge.
var GraphEdgeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `DB Identifier of the edge.`,
		Identifier:     true,
		Name:           "ID",
		Stored:         true,
		Type:           "string",
	},
	"acceptedflows": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AcceptedFlows",
		Description:    `Number of accepted flows in the edge.`,
		Exposed:        true,
		Name:           "acceptedFlows",
		Stored:         true,
		Type:           "integer",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CreateTime",
		Description:    `Date on which the edge has been inserted.`,
		Name:           "createTime",
		Stored:         true,
		Type:           "time",
	},
	"destinationid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "destinationID",
		Stored:         true,
		Type:           "string",
	},
	"destinationtype": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Node"},
		ConvertedName:  "DestinationType",
		Description:    `Type of the destination ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "destinationType",
		Stored:         true,
		Type:           "enum",
	},
	"encrypted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Encrypted",
		Description:    `The number of encrypted flows in the edge.`,
		Exposed:        true,
		Name:           "encrypted",
		Stored:         true,
		Type:           "integer",
	},
	"firstseen": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FirstSeen",
		Description:    `Contains the date when the edge was first seen.`,
		Exposed:        true,
		Name:           "firstSeen",
		Stored:         true,
		Type:           "time",
	},
	"flowid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FlowID",
		Description:    `Identifier of the edge.`,
		Exposed:        true,
		Name:           "flowID",
		Stored:         true,
		Type:           "string",
	},
	"lastseen": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastSeen",
		Description:    `Contains the date when the edge was last seen.`,
		Exposed:        true,
		Name:           "lastSeen",
		Stored:         true,
		Type:           "time",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of object represented by the node.`,
		Exposed:        true,
		Name:           "namespace",
		Stored:         true,
		Type:           "string",
	},
	"observedacceptedflows": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedAcceptedFlows",
		Description:    `Number of accepted observed flows.`,
		Exposed:        true,
		Name:           "observedAcceptedFlows",
		Stored:         true,
		Type:           "integer",
	},
	"observedencrypted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedEncrypted",
		Description:    `Number of encrypted observed flows.`,
		Exposed:        true,
		Name:           "observedEncrypted",
		Stored:         true,
		Type:           "integer",
	},
	"observedrejectedflows": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedRejectedFlows",
		Description:    `Number of rejected observed flows.`,
		Exposed:        true,
		Name:           "observedRejectedFlows",
		Stored:         true,
		Type:           "integer",
	},
	"rejectedflows": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RejectedFlows",
		Description:    `Number of rejected flows in the edge.`,
		Exposed:        true,
		Name:           "rejectedFlows",
		Stored:         true,
		Type:           "integer",
	},
	"sourceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "sourceID",
		Stored:         true,
		Type:           "string",
	},
	"sourcetype": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Node"},
		ConvertedName:  "SourceType",
		Description:    `Type of the source ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "sourceType",
		Stored:         true,
		Type:           "enum",
	},
	"zhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseGraphEdgesList represents a list of SparseGraphEdges
type SparseGraphEdgesList []*SparseGraphEdge

// Identity returns the identity of the objects in the list.
func (o SparseGraphEdgesList) Identity() elemental.Identity {

	return GraphEdgeIdentity
}

// Copy returns a pointer to a copy the SparseGraphEdgesList.
func (o SparseGraphEdgesList) Copy() elemental.Identifiables {

	copy := append(SparseGraphEdgesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseGraphEdgesList.
func (o SparseGraphEdgesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseGraphEdgesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseGraphEdge))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseGraphEdgesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseGraphEdgesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseGraphEdgesList converted to GraphEdgesList.
func (o SparseGraphEdgesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseGraphEdgesList) Version() int {

	return 1
}

// SparseGraphEdge represents the sparse version of a graphedge.
type SparseGraphEdge struct {
	// DB Identifier of the edge.
	ID *string `json:"-" msgpack:"-" bson:"_id" mapstructure:"-,omitempty"`

	// Number of accepted flows in the edge.
	AcceptedFlows *int `json:"acceptedFlows,omitempty" msgpack:"acceptedFlows,omitempty" bson:"acceptedflows,omitempty" mapstructure:"acceptedFlows,omitempty"`

	// Date on which the edge has been inserted.
	CreateTime *time.Time `json:"-" msgpack:"-" bson:"createtime,omitempty" mapstructure:"-,omitempty"`

	// ID of the destination `GraphNode` of the edge.
	DestinationID *string `json:"destinationID,omitempty" msgpack:"destinationID,omitempty" bson:"destinationid,omitempty" mapstructure:"destinationID,omitempty"`

	// Type of the destination `GraphNode` of the edge.
	DestinationType *GraphEdgeDestinationTypeValue `json:"destinationType,omitempty" msgpack:"destinationType,omitempty" bson:"destinationtype,omitempty" mapstructure:"destinationType,omitempty"`

	// The number of encrypted flows in the edge.
	Encrypted *int `json:"encrypted,omitempty" msgpack:"encrypted,omitempty" bson:"encrypted,omitempty" mapstructure:"encrypted,omitempty"`

	// Contains the date when the edge was first seen.
	FirstSeen *time.Time `json:"firstSeen,omitempty" msgpack:"firstSeen,omitempty" bson:"firstseen,omitempty" mapstructure:"firstSeen,omitempty"`

	// Identifier of the edge.
	FlowID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"flowid,omitempty" mapstructure:"ID,omitempty"`

	// Contains the date when the edge was last seen.
	LastSeen *time.Time `json:"lastSeen,omitempty" msgpack:"lastSeen,omitempty" bson:"lastseen,omitempty" mapstructure:"lastSeen,omitempty"`

	// Namespace of object represented by the node.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Number of accepted observed flows.
	ObservedAcceptedFlows *int `json:"observedAcceptedFlows,omitempty" msgpack:"observedAcceptedFlows,omitempty" bson:"observedacceptedflows,omitempty" mapstructure:"observedAcceptedFlows,omitempty"`

	// Number of encrypted observed flows.
	ObservedEncrypted *int `json:"observedEncrypted,omitempty" msgpack:"observedEncrypted,omitempty" bson:"observedencrypted,omitempty" mapstructure:"observedEncrypted,omitempty"`

	// Number of rejected observed flows.
	ObservedRejectedFlows *int `json:"observedRejectedFlows,omitempty" msgpack:"observedRejectedFlows,omitempty" bson:"observedrejectedflows,omitempty" mapstructure:"observedRejectedFlows,omitempty"`

	// Number of rejected flows in the edge.
	RejectedFlows *int `json:"rejectedFlows,omitempty" msgpack:"rejectedFlows,omitempty" bson:"rejectedflows,omitempty" mapstructure:"rejectedFlows,omitempty"`

	// ID of the source `GraphNode` of the edge.
	SourceID *string `json:"sourceID,omitempty" msgpack:"sourceID,omitempty" bson:"sourceid,omitempty" mapstructure:"sourceID,omitempty"`

	// Type of the source `GraphNode` of the edge.
	SourceType *GraphEdgeSourceTypeValue `json:"sourceType,omitempty" msgpack:"sourceType,omitempty" bson:"sourcetype,omitempty" mapstructure:"sourceType,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseGraphEdge returns a new  SparseGraphEdge.
func NewSparseGraphEdge() *SparseGraphEdge {
	return &SparseGraphEdge{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseGraphEdge) Identity() elemental.Identity {

	return GraphEdgeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseGraphEdge) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseGraphEdge) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseGraphEdge) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseGraphEdge) ToPlain() elemental.PlainIdentifiable {

	out := NewGraphEdge()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.AcceptedFlows != nil {
		out.AcceptedFlows = *o.AcceptedFlows
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.DestinationID != nil {
		out.DestinationID = *o.DestinationID
	}
	if o.DestinationType != nil {
		out.DestinationType = *o.DestinationType
	}
	if o.Encrypted != nil {
		out.Encrypted = *o.Encrypted
	}
	if o.FirstSeen != nil {
		out.FirstSeen = *o.FirstSeen
	}
	if o.FlowID != nil {
		out.FlowID = *o.FlowID
	}
	if o.LastSeen != nil {
		out.LastSeen = *o.LastSeen
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.ObservedAcceptedFlows != nil {
		out.ObservedAcceptedFlows = *o.ObservedAcceptedFlows
	}
	if o.ObservedEncrypted != nil {
		out.ObservedEncrypted = *o.ObservedEncrypted
	}
	if o.ObservedRejectedFlows != nil {
		out.ObservedRejectedFlows = *o.ObservedRejectedFlows
	}
	if o.RejectedFlows != nil {
		out.RejectedFlows = *o.RejectedFlows
	}
	if o.SourceID != nil {
		out.SourceID = *o.SourceID
	}
	if o.SourceType != nil {
		out.SourceType = *o.SourceType
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseGraphEdge) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseGraphEdge) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseGraphEdge) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseGraphEdge) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseGraphEdge.
func (o *SparseGraphEdge) DeepCopy() *SparseGraphEdge {

	if o == nil {
		return nil
	}

	out := &SparseGraphEdge{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseGraphEdge.
func (o *SparseGraphEdge) DeepCopyInto(out *SparseGraphEdge) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseGraphEdge: %s", err))
	}

	*out = *target.(*SparseGraphEdge)
}
