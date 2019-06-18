package gaia

import (
	"fmt"

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

// GraphEdge represents the model of a graphedge
type GraphEdge struct {
	// Identifier of the edge.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Number of accepted flows in the edge.
	AcceptedFlows int `json:"acceptedFlows" msgpack:"acceptedFlows" bson:"-" mapstructure:"acceptedFlows,omitempty"`

	// ID of the destination GraphNode of the edge.
	DestinationID string `json:"destinationID" msgpack:"destinationID" bson:"-" mapstructure:"destinationID,omitempty"`

	// Type of the destination GraphNode of the edge.
	DestinationType GraphEdgeDestinationTypeValue `json:"destinationType" msgpack:"destinationType" bson:"-" mapstructure:"destinationType,omitempty"`

	// Tells the number of encrypted flows in the edge.
	Encrypted int `json:"encrypted" msgpack:"encrypted" bson:"-" mapstructure:"encrypted,omitempty"`

	// Number of accepted observed flows.
	ObservedAcceptedFlows int `json:"observedAcceptedFlows" msgpack:"observedAcceptedFlows" bson:"-" mapstructure:"observedAcceptedFlows,omitempty"`

	// Number of encrypted observed flows.
	ObservedEncrypted int `json:"observedEncrypted" msgpack:"observedEncrypted" bson:"-" mapstructure:"observedEncrypted,omitempty"`

	// Information about the observation policies that was hit in the flows
	// represented by that edge.
	ObservedPolicyIDs map[string]*GraphPolicyInfo `json:"observedPolicyIDs" msgpack:"observedPolicyIDs" bson:"-" mapstructure:"observedPolicyIDs,omitempty"`

	// Number of rejected observed flows.
	ObservedRejectedFlows int `json:"observedRejectedFlows" msgpack:"observedRejectedFlows" bson:"-" mapstructure:"observedRejectedFlows,omitempty"`

	// Map of ints...
	ObservedServiceIDs map[string]int `json:"observedServiceIDs" msgpack:"observedServiceIDs" bson:"-" mapstructure:"observedServiceIDs,omitempty"`

	// Information about the policies that was hit in the flows represented by that
	// edge.
	PolicyIDs map[string]*GraphPolicyInfo `json:"policyIDs" msgpack:"policyIDs" bson:"-" mapstructure:"policyIDs,omitempty"`

	// Number of rejected flows in the edge.
	RejectedFlows int `json:"rejectedFlows" msgpack:"rejectedFlows" bson:"-" mapstructure:"rejectedFlows,omitempty"`

	// ID of the source GraphNode of the edge.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"-" mapstructure:"sourceID,omitempty"`

	// Type of the source GraphNode of the edge.
	SourceType GraphEdgeSourceTypeValue `json:"sourceType" msgpack:"sourceType" bson:"-" mapstructure:"sourceType,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewGraphEdge returns a new *GraphEdge
func NewGraphEdge() *GraphEdge {

	return &GraphEdge{
		ModelVersion:       1,
		ObservedPolicyIDs:  map[string]*GraphPolicyInfo{},
		ObservedServiceIDs: map[string]int{},
		PolicyIDs:          map[string]*GraphPolicyInfo{},
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

	for _, sub := range o.ObservedPolicyIDs {
		if sub == nil {
			continue
		}
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.PolicyIDs {
		if sub == nil {
			continue
		}
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
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
