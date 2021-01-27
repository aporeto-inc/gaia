package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudEdge represents the model of a cloudedge
type CloudEdge struct {
	// ID of the destination `cloud node` of the edge.
	DestinationID string `json:"destinationID" msgpack:"destinationID" bson:"destinationid" mapstructure:"destinationID,omitempty"`

	// ID of the source `cloud node` of the edge.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"sourceid" mapstructure:"sourceID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudEdge returns a new *CloudEdge
func NewCloudEdge() *CloudEdge {

	return &CloudEdge{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEdge) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudEdge{}

	s.DestinationID = o.DestinationID
	s.SourceID = o.SourceID

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEdge) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudEdge{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.DestinationID = s.DestinationID
	o.SourceID = s.SourceID

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudEdge) BleveType() string {

	return "cloudedge"
}

// DeepCopy returns a deep copy if the CloudEdge.
func (o *CloudEdge) DeepCopy() *CloudEdge {

	if o == nil {
		return nil
	}

	out := &CloudEdge{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudEdge.
func (o *CloudEdge) DeepCopyInto(out *CloudEdge) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudEdge: %s", err))
	}

	*out = *target.(*CloudEdge)
}

// Validate valides the current information stored into the structure.
func (o *CloudEdge) Validate() error {

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

type mongoAttributesCloudEdge struct {
	DestinationID string `bson:"destinationid"`
	SourceID      string `bson:"sourceid"`
}
