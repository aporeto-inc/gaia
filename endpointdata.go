package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EndpointDataTypeValue represents the possible values for attribute "type".
type EndpointDataTypeValue string

const (
	// EndpointDataTypeGateway represents the value Gateway.
	EndpointDataTypeGateway EndpointDataTypeValue = "Gateway"

	// EndpointDataTypeInstance represents the value Instance.
	EndpointDataTypeInstance EndpointDataTypeValue = "Instance"

	// EndpointDataTypeLoadBalancer represents the value LoadBalancer.
	EndpointDataTypeLoadBalancer EndpointDataTypeValue = "LoadBalancer"

	// EndpointDataTypePeeringConnection represents the value PeeringConnection.
	EndpointDataTypePeeringConnection EndpointDataTypeValue = "PeeringConnection"

	// EndpointDataTypeService represents the value Service.
	EndpointDataTypeService EndpointDataTypeValue = "Service"

	// EndpointDataTypeTransitGateway represents the value TransitGateway.
	EndpointDataTypeTransitGateway EndpointDataTypeValue = "TransitGateway"
)

// EndpointData represents the model of a endpointdata
type EndpointData struct {
	// Indicates that the endpoint is directly attached to the VPC. In this case the
	// attachedInterfaces is empty. In general this is only valid for endpoint type
	// Gateway and Peering Connection.
	VPCAttached bool `json:"VPCAttached" msgpack:"VPCAttached" bson:"vpcattached" mapstructure:"VPCAttached,omitempty"`

	// The list of VPCs that this endpoint is directly attached to.
	VPCAttachments []string `json:"VPCAttachments" msgpack:"VPCAttachments" bson:"vpcattachments" mapstructure:"VPCAttachments,omitempty"`

	// List of route tables associated with this endpoint. Depending on cloud provider
	// it can apply in some gateways.
	AssociatedRouteTables []string `json:"associatedRouteTables" msgpack:"associatedRouteTables" bson:"associatedroutetables" mapstructure:"associatedRouteTables,omitempty"`

	// A list of interfaces attached with the endpoint. In some cases endpoints can
	// have more than one interface.
	AttachedInterfaces []string `json:"attachedInterfaces" msgpack:"attachedInterfaces" bson:"attachedinterfaces" mapstructure:"attachedInterfaces,omitempty"`

	// If the endpoint has multiple connections and forwarding can be enabled between
	// them.
	ForwardingEnabled bool `json:"forwardingEnabled" msgpack:"forwardingEnabled" bson:"forwardingenabled" mapstructure:"forwardingEnabled,omitempty"`

	// Type of the endpoint.
	Type EndpointDataTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEndpointData returns a new *EndpointData
func NewEndpointData() *EndpointData {

	return &EndpointData{
		ModelVersion:          1,
		AssociatedRouteTables: []string{},
		VPCAttachments:        []string{},
		AttachedInterfaces:    []string{},
		ForwardingEnabled:     true,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EndpointData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesEndpointData{}

	s.VPCAttached = o.VPCAttached
	s.VPCAttachments = o.VPCAttachments
	s.AssociatedRouteTables = o.AssociatedRouteTables
	s.AttachedInterfaces = o.AttachedInterfaces
	s.ForwardingEnabled = o.ForwardingEnabled
	s.Type = o.Type

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EndpointData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesEndpointData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.VPCAttached = s.VPCAttached
	o.VPCAttachments = s.VPCAttachments
	o.AssociatedRouteTables = s.AssociatedRouteTables
	o.AttachedInterfaces = s.AttachedInterfaces
	o.ForwardingEnabled = s.ForwardingEnabled
	o.Type = s.Type

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *EndpointData) BleveType() string {

	return "endpointdata"
}

// DeepCopy returns a deep copy if the EndpointData.
func (o *EndpointData) DeepCopy() *EndpointData {

	if o == nil {
		return nil
	}

	out := &EndpointData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *EndpointData.
func (o *EndpointData) DeepCopyInto(out *EndpointData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy EndpointData: %s", err))
	}

	*out = *target.(*EndpointData)
}

// Validate valides the current information stored into the structure.
func (o *EndpointData) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Instance", "LoadBalancer", "PeeringConnection", "Service", "Gateway", "TransitGateway"}, false); err != nil {
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

type mongoAttributesEndpointData struct {
	VPCAttached           bool                  `bson:"vpcattached"`
	VPCAttachments        []string              `bson:"vpcattachments"`
	AssociatedRouteTables []string              `bson:"associatedroutetables"`
	AttachedInterfaces    []string              `bson:"attachedinterfaces"`
	ForwardingEnabled     bool                  `bson:"forwardingenabled"`
	Type                  EndpointDataTypeValue `bson:"type"`
}
