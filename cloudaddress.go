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

// CloudAddressIdentity represents the Identity of the object.
var CloudAddressIdentity = elemental.Identity{
	Name:     "cloudaddress",
	Category: "cloudaddresses",
	Package:  "yeul",
	Private:  false,
}

// CloudAddressList represents a list of CloudAddress
type CloudAddressList []*CloudAddress

// Identity returns the identity of the objects in the list.
func (o CloudAddressList) Identity() elemental.Identity {

	return CloudAddressIdentity
}

// Copy returns a pointer to a copy the CloudAddressList.
func (o CloudAddressList) Copy() elemental.Identifiables {

	copy := append(CloudAddressList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudAddressList.
func (o CloudAddressList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudAddressList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudAddress))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudAddressList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudAddressList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudAddressList converted to SparseCloudAddressList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudAddressList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudAddressList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudAddress)
	}

	return out
}

// Version returns the version of the content.
func (o CloudAddressList) Version() int {

	return 1
}

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
		Primary:      false,
	}
}

// Identity returns the Identity of the object.
func (o *CloudAddress) Identity() elemental.Identity {

	return CloudAddressIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudAddress) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudAddress) SetIdentifier(id string) {

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

// Version returns the hardcoded version of the model.
func (o *CloudAddress) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudAddress) BleveType() string {

	return "cloudaddress"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudAddress) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudAddress) Doc() string {

	return `Managed the list of IP addresses associated with an interface.`
}

func (o *CloudAddress) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudAddress) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudAddress{
			IPVersion:      &o.IPVersion,
			Primary:        &o.Primary,
			PrivateDNSName: &o.PrivateDNSName,
			PrivateIP:      &o.PrivateIP,
			PublicDNSName:  &o.PublicDNSName,
			PublicIP:       &o.PublicIP,
		}
	}

	sp := &SparseCloudAddress{}
	for _, f := range fields {
		switch f {
		case "IPVersion":
			sp.IPVersion = &(o.IPVersion)
		case "primary":
			sp.Primary = &(o.Primary)
		case "privateDNSName":
			sp.PrivateDNSName = &(o.PrivateDNSName)
		case "privateIP":
			sp.PrivateIP = &(o.PrivateIP)
		case "publicDNSName":
			sp.PublicDNSName = &(o.PublicDNSName)
		case "publicIP":
			sp.PublicIP = &(o.PublicIP)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudAddress to the object.
func (o *CloudAddress) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudAddress)
	if so.IPVersion != nil {
		o.IPVersion = *so.IPVersion
	}
	if so.Primary != nil {
		o.Primary = *so.Primary
	}
	if so.PrivateDNSName != nil {
		o.PrivateDNSName = *so.PrivateDNSName
	}
	if so.PrivateIP != nil {
		o.PrivateIP = *so.PrivateIP
	}
	if so.PublicDNSName != nil {
		o.PublicDNSName = *so.PublicDNSName
	}
	if so.PublicIP != nil {
		o.PublicIP = *so.PublicIP
	}
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

	if err := elemental.ValidateStringInList("IPVersion", string(o.IPVersion), []string{"IPv4", "IPv6"}, false); err != nil {
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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*CloudAddress) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudAddressAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudAddressLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudAddress) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudAddressAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudAddress) ValueForAttribute(name string) interface{} {

	switch name {
	case "IPVersion":
		return o.IPVersion
	case "primary":
		return o.Primary
	case "privateDNSName":
		return o.PrivateDNSName
	case "privateIP":
		return o.PrivateIP
	case "publicDNSName":
		return o.PublicDNSName
	case "publicIP":
		return o.PublicIP
	}

	return nil
}

// CloudAddressAttributesMap represents the map of attribute for CloudAddress.
var CloudAddressAttributesMap = map[string]elemental.AttributeSpecification{
	"IPVersion": {
		AllowedChoices: []string{"IPv4", "IPv6"},
		BSONFieldName:  "ipversion",
		ConvertedName:  "IPVersion",
		Description:    `Designates IPv4 or IPv6.`,
		Exposed:        true,
		Name:           "IPVersion",
		Stored:         true,
		Type:           "enum",
	},
	"Primary": {
		AllowedChoices: []string{},
		BSONFieldName:  "primary",
		ConvertedName:  "Primary",
		Description:    `Designates the IP address as the primary IP address.`,
		Exposed:        true,
		Name:           "primary",
		Stored:         true,
		Type:           "boolean",
	},
	"PrivateDNSName": {
		AllowedChoices: []string{},
		BSONFieldName:  "privatednsname",
		ConvertedName:  "PrivateDNSName",
		Description:    `The private DNS name associated with the address.`,
		Exposed:        true,
		Name:           "privateDNSName",
		Stored:         true,
		Type:           "string",
	},
	"PrivateIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "privateip",
		ConvertedName:  "PrivateIP",
		Description:    `The private IP address value.`,
		Exposed:        true,
		Name:           "privateIP",
		Stored:         true,
		Type:           "string",
	},
	"PublicDNSName": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicdnsname",
		ConvertedName:  "PublicDNSName",
		Description:    `The private DNS name associated with the address.`,
		Exposed:        true,
		Name:           "publicDNSName",
		Stored:         true,
		Type:           "string",
	},
	"PublicIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicip",
		ConvertedName:  "PublicIP",
		Description:    `The private IP address value.`,
		Exposed:        true,
		Name:           "publicIP",
		Stored:         true,
		Type:           "string",
	},
}

// CloudAddressLowerCaseAttributesMap represents the map of attribute for CloudAddress.
var CloudAddressLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"ipversion": {
		AllowedChoices: []string{"IPv4", "IPv6"},
		BSONFieldName:  "ipversion",
		ConvertedName:  "IPVersion",
		Description:    `Designates IPv4 or IPv6.`,
		Exposed:        true,
		Name:           "IPVersion",
		Stored:         true,
		Type:           "enum",
	},
	"primary": {
		AllowedChoices: []string{},
		BSONFieldName:  "primary",
		ConvertedName:  "Primary",
		Description:    `Designates the IP address as the primary IP address.`,
		Exposed:        true,
		Name:           "primary",
		Stored:         true,
		Type:           "boolean",
	},
	"privatednsname": {
		AllowedChoices: []string{},
		BSONFieldName:  "privatednsname",
		ConvertedName:  "PrivateDNSName",
		Description:    `The private DNS name associated with the address.`,
		Exposed:        true,
		Name:           "privateDNSName",
		Stored:         true,
		Type:           "string",
	},
	"privateip": {
		AllowedChoices: []string{},
		BSONFieldName:  "privateip",
		ConvertedName:  "PrivateIP",
		Description:    `The private IP address value.`,
		Exposed:        true,
		Name:           "privateIP",
		Stored:         true,
		Type:           "string",
	},
	"publicdnsname": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicdnsname",
		ConvertedName:  "PublicDNSName",
		Description:    `The private DNS name associated with the address.`,
		Exposed:        true,
		Name:           "publicDNSName",
		Stored:         true,
		Type:           "string",
	},
	"publicip": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicip",
		ConvertedName:  "PublicIP",
		Description:    `The private IP address value.`,
		Exposed:        true,
		Name:           "publicIP",
		Stored:         true,
		Type:           "string",
	},
}

// SparseCloudAddressList represents a list of SparseCloudAddress
type SparseCloudAddressList []*SparseCloudAddress

// Identity returns the identity of the objects in the list.
func (o SparseCloudAddressList) Identity() elemental.Identity {

	return CloudAddressIdentity
}

// Copy returns a pointer to a copy the SparseCloudAddressList.
func (o SparseCloudAddressList) Copy() elemental.Identifiables {

	copy := append(SparseCloudAddressList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudAddressList.
func (o SparseCloudAddressList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudAddressList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudAddress))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudAddressList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudAddressList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudAddressList converted to CloudAddressList.
func (o SparseCloudAddressList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudAddressList) Version() int {

	return 1
}

// SparseCloudAddress represents the sparse version of a cloudaddress.
type SparseCloudAddress struct {
	// Designates IPv4 or IPv6.
	IPVersion *CloudAddressIPVersionValue `json:"IPVersion,omitempty" msgpack:"IPVersion,omitempty" bson:"ipversion,omitempty" mapstructure:"IPVersion,omitempty"`

	// Designates the IP address as the primary IP address.
	Primary *bool `json:"primary,omitempty" msgpack:"primary,omitempty" bson:"primary,omitempty" mapstructure:"primary,omitempty"`

	// The private DNS name associated with the address.
	PrivateDNSName *string `json:"privateDNSName,omitempty" msgpack:"privateDNSName,omitempty" bson:"privatednsname,omitempty" mapstructure:"privateDNSName,omitempty"`

	// The private IP address value.
	PrivateIP *string `json:"privateIP,omitempty" msgpack:"privateIP,omitempty" bson:"privateip,omitempty" mapstructure:"privateIP,omitempty"`

	// The private DNS name associated with the address.
	PublicDNSName *string `json:"publicDNSName,omitempty" msgpack:"publicDNSName,omitempty" bson:"publicdnsname,omitempty" mapstructure:"publicDNSName,omitempty"`

	// The private IP address value.
	PublicIP *string `json:"publicIP,omitempty" msgpack:"publicIP,omitempty" bson:"publicip,omitempty" mapstructure:"publicIP,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudAddress returns a new  SparseCloudAddress.
func NewSparseCloudAddress() *SparseCloudAddress {
	return &SparseCloudAddress{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudAddress) Identity() elemental.Identity {

	return CloudAddressIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudAddress) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudAddress) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudAddress) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudAddress{}

	if o.IPVersion != nil {
		s.IPVersion = o.IPVersion
	}
	if o.Primary != nil {
		s.Primary = o.Primary
	}
	if o.PrivateDNSName != nil {
		s.PrivateDNSName = o.PrivateDNSName
	}
	if o.PrivateIP != nil {
		s.PrivateIP = o.PrivateIP
	}
	if o.PublicDNSName != nil {
		s.PublicDNSName = o.PublicDNSName
	}
	if o.PublicIP != nil {
		s.PublicIP = o.PublicIP
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudAddress) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudAddress{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.IPVersion != nil {
		o.IPVersion = s.IPVersion
	}
	if s.Primary != nil {
		o.Primary = s.Primary
	}
	if s.PrivateDNSName != nil {
		o.PrivateDNSName = s.PrivateDNSName
	}
	if s.PrivateIP != nil {
		o.PrivateIP = s.PrivateIP
	}
	if s.PublicDNSName != nil {
		o.PublicDNSName = s.PublicDNSName
	}
	if s.PublicIP != nil {
		o.PublicIP = s.PublicIP
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCloudAddress) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudAddress) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudAddress()
	if o.IPVersion != nil {
		out.IPVersion = *o.IPVersion
	}
	if o.Primary != nil {
		out.Primary = *o.Primary
	}
	if o.PrivateDNSName != nil {
		out.PrivateDNSName = *o.PrivateDNSName
	}
	if o.PrivateIP != nil {
		out.PrivateIP = *o.PrivateIP
	}
	if o.PublicDNSName != nil {
		out.PublicDNSName = *o.PublicDNSName
	}
	if o.PublicIP != nil {
		out.PublicIP = *o.PublicIP
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCloudAddress.
func (o *SparseCloudAddress) DeepCopy() *SparseCloudAddress {

	if o == nil {
		return nil
	}

	out := &SparseCloudAddress{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudAddress.
func (o *SparseCloudAddress) DeepCopyInto(out *SparseCloudAddress) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudAddress: %s", err))
	}

	*out = *target.(*SparseCloudAddress)
}

type mongoAttributesCloudAddress struct {
	IPVersion      CloudAddressIPVersionValue `bson:"ipversion"`
	Primary        bool                       `bson:"primary"`
	PrivateDNSName string                     `bson:"privatednsname"`
	PrivateIP      string                     `bson:"privateip"`
	PublicDNSName  string                     `bson:"publicdnsname"`
	PublicIP       string                     `bson:"publicip"`
}
type mongoAttributesSparseCloudAddress struct {
	IPVersion      *CloudAddressIPVersionValue `bson:"ipversion,omitempty"`
	Primary        *bool                       `bson:"primary,omitempty"`
	PrivateDNSName *string                     `bson:"privatednsname,omitempty"`
	PrivateIP      *string                     `bson:"privateip,omitempty"`
	PublicDNSName  *string                     `bson:"publicdnsname,omitempty"`
	PublicIP       *string                     `bson:"publicip,omitempty"`
}
