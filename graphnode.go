package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// GraphNodeTypeValue represents the possible values for attribute "type".
type GraphNodeTypeValue string

const (
	// GraphNodeTypeClaim represents the value Claim.
	GraphNodeTypeClaim GraphNodeTypeValue = "Claim"

	// GraphNodeTypeDocker represents the value Docker.
	GraphNodeTypeDocker GraphNodeTypeValue = "Docker"

	// GraphNodeTypeExternalNetwork represents the value ExternalNetwork.
	GraphNodeTypeExternalNetwork GraphNodeTypeValue = "ExternalNetwork"

	// GraphNodeTypeVolume represents the value Volume.
	GraphNodeTypeVolume GraphNodeTypeValue = "Volume"
)

// GraphNode represents the model of a graphnode
type GraphNode struct {
	// Identifier of object represented by the node.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Description of object represented by the node.
	Description string `json:"description" bson:"-" mapstructure:"description,omitempty"`

	// Enforcement status of processing unit represented by the node.
	EnforcementStatus string `json:"enforcementStatus" bson:"-" mapstructure:"enforcementStatus,omitempty"`

	// ID of the group the node is eventually part of.
	GroupID string `json:"groupID" bson:"-" mapstructure:"groupID,omitempty"`

	// Last update of the node.
	LastUpdate time.Time `json:"lastUpdate" bson:"-" mapstructure:"lastUpdate,omitempty"`

	// Name of object represented by the node.
	Name string `json:"name" bson:"-" mapstructure:"name,omitempty"`

	// Namespace of object represented by the node.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// Status of object represented by the node.
	Status string `json:"status" bson:"-" mapstructure:"status,omitempty"`

	// Tags of object represented by the node.
	Tags []string `json:"tags" bson:"-" mapstructure:"tags,omitempty"`

	// Type of object represented by the node.
	Type GraphNodeTypeValue `json:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewGraphNode returns a new *GraphNode
func NewGraphNode() *GraphNode {

	return &GraphNode{
		ModelVersion: 1,
	}
}

// DeepCopy returns a deep copy if the GraphNode.
func (o *GraphNode) DeepCopy() *GraphNode {

	if o == nil {
		return nil
	}

	out := &GraphNode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *GraphNode.
func (o *GraphNode) DeepCopyInto(out *GraphNode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy GraphNode: %s", err))
	}

	*out = *target.(*GraphNode)
}

// Validate valides the current information stored into the structure.
func (o *GraphNode) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Docker", "ExternalNetwork", "Volume", "Claim"}, false); err != nil {
		errors = append(errors, err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
