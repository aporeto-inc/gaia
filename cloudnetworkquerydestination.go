package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudNetworkQueryDestination represents the model of a cloudnetworkquerydestination
type CloudNetworkQueryDestination struct {
	// The IP address of the destination if it is a public address.
	IP string `json:"IP" msgpack:"IP" bson:"-" mapstructure:"IP,omitempty"`

	// Returns true of this is just a public IP destination.
	IsPublicIP bool `json:"isPublicIP" msgpack:"isPublicIP" bson:"-" mapstructure:"isPublicIP,omitempty"`

	// The nativeID of the destination if it is a cloud node.
	NativeID string `json:"nativeID" msgpack:"nativeID" bson:"-" mapstructure:"nativeID,omitempty"`

	// Returns true if the destination is reachable through routing.
	Reachable bool `json:"reachable" msgpack:"reachable" bson:"-" mapstructure:"reachable,omitempty"`

	// Returns the network security verdict for the destination.
	Verdict string `json:"verdict" msgpack:"verdict" bson:"-" mapstructure:"verdict,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudNetworkQueryDestination returns a new *CloudNetworkQueryDestination
func NewCloudNetworkQueryDestination() *CloudNetworkQueryDestination {

	return &CloudNetworkQueryDestination{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkQueryDestination) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudNetworkQueryDestination{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkQueryDestination) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudNetworkQueryDestination{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudNetworkQueryDestination) BleveType() string {

	return "cloudnetworkquerydestination"
}

// DeepCopy returns a deep copy if the CloudNetworkQueryDestination.
func (o *CloudNetworkQueryDestination) DeepCopy() *CloudNetworkQueryDestination {

	if o == nil {
		return nil
	}

	out := &CloudNetworkQueryDestination{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudNetworkQueryDestination.
func (o *CloudNetworkQueryDestination) DeepCopyInto(out *CloudNetworkQueryDestination) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudNetworkQueryDestination: %s", err))
	}

	*out = *target.(*CloudNetworkQueryDestination)
}

// Validate valides the current information stored into the structure.
func (o *CloudNetworkQueryDestination) Validate() error {

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

type mongoAttributesCloudNetworkQueryDestination struct {
}
