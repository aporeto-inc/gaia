package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudAddressIPVersionValue represents the possible values for attribute "IPVersion".
type CloudAddressIPVersionValue string

const (
	// CloudAddressIPVersionIPv4 represents the value IPv4.
	CloudAddressIPVersionIPv4 CloudAddressIPVersionValue = "IPv4"

	// CloudAddressIPVersionIPv6 represents the value IPv6.
	CloudAddressIPVersionIPv6 CloudAddressIPVersionValue = "IPv6"
)

// CloudAddress represents the model of a cloudaddress
type CloudAddress struct {
	// Designates IPv4 or IPv6.
	IPVersion CloudAddressIPVersionValue `json:"IPVersion" msgpack:"IPVersion" bson:"ipversion" mapstructure:"IPVersion,omitempty"`

	// Designates the IP address as the primary IP address.
	Primary bool `json:"primary" msgpack:"primary" bson:"primary" mapstructure:"primary,omitempty"`

	// The private DNS name associated with the address.
	PrivateDNSName string `json:"privateDNSName" msgpack:"privateDNSName" bson:"privatednsname" mapstructure:"privateDNSName,omitempty"`

	// The private IP address value.
	PrivateIP string `json:"privateIP" msgpack:"privateIP" bson:"privateip" mapstructure:"privateIP,omitempty"`

	// The private DNS name associated with the address.
	PublicDNSName string `json:"publicDNSName" msgpack:"publicDNSName" bson:"publicdnsname" mapstructure:"publicDNSName,omitempty"`

	// The private IP address value.
	PublicIP string `json:"publicIP" msgpack:"publicIP" bson:"publicip" mapstructure:"publicIP,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudAddress returns a new *CloudAddress
func NewCloudAddress() *CloudAddress {

	return &CloudAddress{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAddress) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudAddress{}

	s.IPVersion = o.IPVersion
	s.Primary = o.Primary
	s.PrivateDNSName = o.PrivateDNSName
	s.PrivateIP = o.PrivateIP
	s.PublicDNSName = o.PublicDNSName
	s.PublicIP = o.PublicIP

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAddress) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudAddress{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.IPVersion = s.IPVersion
	o.Primary = s.Primary
	o.PrivateDNSName = s.PrivateDNSName
	o.PrivateIP = s.PrivateIP
	o.PublicDNSName = s.PublicDNSName
	o.PublicIP = s.PublicIP

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudAddress) BleveType() string {

	return "cloudaddress"
}

// DeepCopy returns a deep copy if the CloudAddress.
func (o *CloudAddress) DeepCopy() *CloudAddress {

	if o == nil {
		return nil
	}

	out := &CloudAddress{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudAddress.
func (o *CloudAddress) DeepCopyInto(out *CloudAddress) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudAddress: %s", err))
	}

	*out = *target.(*CloudAddress)
}

// Validate valides the current information stored into the structure.
func (o *CloudAddress) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("IPVersion", string(o.IPVersion)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("IPVersion", string(o.IPVersion), []string{"IPv4", "IPv6"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRorIP("privateIP", o.PrivateIP); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRorIP("publicIP", o.PublicIP); err != nil {
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

type mongoAttributesCloudAddress struct {
	IPVersion      CloudAddressIPVersionValue `bson:"ipversion"`
	Primary        bool                       `bson:"primary"`
	PrivateDNSName string                     `bson:"privatednsname"`
	PrivateIP      string                     `bson:"privateip"`
	PublicDNSName  string                     `bson:"publicdnsname"`
	PublicIP       string                     `bson:"publicip"`
}
