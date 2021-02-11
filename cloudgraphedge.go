package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudGraphEdge represents the model of a cloudgraphedge
type CloudGraphEdge struct {
	// ID of the destination `cloud node` of the edge.
	DestinationID string `json:"destinationID" msgpack:"destinationID" bson:"destinationid" mapstructure:"destinationID,omitempty"`

	// ID of the source `cloud node` of the edge.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"sourceid" mapstructure:"sourceID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudGraphEdge returns a new *CloudGraphEdge
func NewCloudGraphEdge() *CloudGraphEdge {

	return &CloudGraphEdge{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphEdge) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudGraphEdge{}

	s.DestinationID = o.DestinationID
	s.SourceID = o.SourceID

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphEdge) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudGraphEdge{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.DestinationID = s.DestinationID
	o.SourceID = s.SourceID

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudGraphEdge) BleveType() string {

	return "cloudgraphedge"
}

// DeepCopy returns a deep copy if the CloudGraphEdge.
func (o *CloudGraphEdge) DeepCopy() *CloudGraphEdge {

	if o == nil {
		return nil
	}

	out := &CloudGraphEdge{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudGraphEdge.
func (o *CloudGraphEdge) DeepCopyInto(out *CloudGraphEdge) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudGraphEdge: %s", err))
	}

	*out = *target.(*CloudGraphEdge)
}

// Validate valides the current information stored into the structure.
func (o *CloudGraphEdge) Validate() error {

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

type mongoAttributesCloudGraphEdge struct {
	DestinationID string `bson:"destinationid"`
	SourceID      string `bson:"sourceid"`
}
