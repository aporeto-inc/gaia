package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// SubnetData represents the model of a subnetdata
type SubnetData struct {
	// Address CIDR of the Subnet.
	Address string `json:"address" msgpack:"address" bson:"address" mapstructure:"address,omitempty"`

	// The availability zone ID of the subnet.
	ZoneID string `json:"zoneID" msgpack:"zoneID" bson:"zoneid" mapstructure:"zoneID,omitempty"`

	// The availability zone of the subnet.
	ZoneName string `json:"zoneName" msgpack:"zoneName" bson:"zonename" mapstructure:"zoneName,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSubnetData returns a new *SubnetData
func NewSubnetData() *SubnetData {

	return &SubnetData{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SubnetData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSubnetData{}

	s.Address = o.Address
	s.ZoneID = o.ZoneID
	s.ZoneName = o.ZoneName

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SubnetData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSubnetData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Address = s.Address
	o.ZoneID = s.ZoneID
	o.ZoneName = s.ZoneName

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *SubnetData) BleveType() string {

	return "subnetdata"
}

// DeepCopy returns a deep copy if the SubnetData.
func (o *SubnetData) DeepCopy() *SubnetData {

	if o == nil {
		return nil
	}

	out := &SubnetData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SubnetData.
func (o *SubnetData) DeepCopyInto(out *SubnetData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SubnetData: %s", err))
	}

	*out = *target.(*SubnetData)
}

// Validate valides the current information stored into the structure.
func (o *SubnetData) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateCIDR("address", o.Address); err != nil {
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

type mongoAttributesSubnetData struct {
	Address  string `bson:"address"`
	ZoneID   string `bson:"zoneid"`
	ZoneName string `bson:"zonename"`
}
