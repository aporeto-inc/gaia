package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// VPCData represents the model of a vpcdata
type VPCData struct {
	// Address CIDR of the VPC.
	Address string `json:"address" msgpack:"address" bson:"address" mapstructure:"address,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewVPCData returns a new *VPCData
func NewVPCData() *VPCData {

	return &VPCData{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *VPCData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesVPCData{}

	s.Address = o.Address

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *VPCData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesVPCData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Address = s.Address

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *VPCData) BleveType() string {

	return "vpcdata"
}

// DeepCopy returns a deep copy if the VPCData.
func (o *VPCData) DeepCopy() *VPCData {

	if o == nil {
		return nil
	}

	out := &VPCData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *VPCData.
func (o *VPCData) DeepCopyInto(out *VPCData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy VPCData: %s", err))
	}

	*out = *target.(*VPCData)
}

// Validate valides the current information stored into the structure.
func (o *VPCData) Validate() error {

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

type mongoAttributesVPCData struct {
	Address string `bson:"address"`
}
