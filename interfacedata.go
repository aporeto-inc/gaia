package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// InterfaceDataTypeValue represents the possible values for attribute "type".
type InterfaceDataTypeValue string

const (
	// InterfaceDataTypeGateway represents the value Gateway.
	InterfaceDataTypeGateway InterfaceDataTypeValue = "Gateway"

	// InterfaceDataTypeInstance represents the value Instance.
	InterfaceDataTypeInstance InterfaceDataTypeValue = "Instance"

	// InterfaceDataTypeLoadBalancer represents the value LoadBalancer.
	InterfaceDataTypeLoadBalancer InterfaceDataTypeValue = "LoadBalancer"

	// InterfaceDataTypeService represents the value Service.
	InterfaceDataTypeService InterfaceDataTypeValue = "Service"

	// InterfaceDataTypeTransitGateway represents the value TransitGateway.
	InterfaceDataTypeTransitGateway InterfaceDataTypeValue = "TransitGateway"
)

// InterfaceData represents the model of a interfacedata
type InterfaceData struct {
	// List of IP addresses/subnets (IPv4 or IPv6) associated with the
	// interface.
	Addresses CloudAddressList `json:"addresses" msgpack:"addresses" bson:"addresses" mapstructure:"addresses,omitempty"`

	// If the interface is of type or external, the relatedObjectID identifies the
	// related service or gateway.
	RelatedObjectID string `json:"relatedObjectID" msgpack:"relatedObjectID" bson:"relatedobjectid" mapstructure:"relatedObjectID,omitempty"`

	// Security tags associated with the instance.
	SecurityTags []string `json:"securityTags" msgpack:"securityTags" bson:"securitytags" mapstructure:"securityTags,omitempty"`

	// ID of subnet associated with this interface.
	Subnet []string `json:"subnet" msgpack:"subnet" bson:"subnet" mapstructure:"subnet,omitempty"`

	// Interface type (Instance, Load Balancer, Gateway, etc).
	Type InterfaceDataTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewInterfaceData returns a new *InterfaceData
func NewInterfaceData() *InterfaceData {

	return &InterfaceData{
		ModelVersion: 1,
		Addresses:    CloudAddressList{},
		SecurityTags: []string{},
		Subnet:       []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *InterfaceData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesInterfaceData{}

	s.Addresses = o.Addresses
	s.RelatedObjectID = o.RelatedObjectID
	s.SecurityTags = o.SecurityTags
	s.Subnet = o.Subnet
	s.Type = o.Type

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *InterfaceData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesInterfaceData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Addresses = s.Addresses
	o.RelatedObjectID = s.RelatedObjectID
	o.SecurityTags = s.SecurityTags
	o.Subnet = s.Subnet
	o.Type = s.Type

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *InterfaceData) BleveType() string {

	return "interfacedata"
}

// DeepCopy returns a deep copy if the InterfaceData.
func (o *InterfaceData) DeepCopy() *InterfaceData {

	if o == nil {
		return nil
	}

	out := &InterfaceData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *InterfaceData.
func (o *InterfaceData) DeepCopyInto(out *InterfaceData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy InterfaceData: %s", err))
	}

	*out = *target.(*InterfaceData)
}

// Validate valides the current information stored into the structure.
func (o *InterfaceData) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Addresses {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Instance", "LoadBalancer", "Gateway", "Service", "TransitGateway"}, false); err != nil {
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

type mongoAttributesInterfaceData struct {
	Addresses       CloudAddressList       `bson:"addresses"`
	RelatedObjectID string                 `bson:"relatedobjectid"`
	SecurityTags    []string               `bson:"securitytags"`
	Subnet          []string               `bson:"subnet"`
	Type            InterfaceDataTypeValue `bson:"type"`
}
