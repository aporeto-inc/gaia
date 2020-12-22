package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudRouteNextHopTypeValue represents the possible values for attribute "nextHopType".
type CloudRouteNextHopTypeValue string

const (
	// CloudRouteNextHopTypeEgressOnlyGateway represents the value EgressOnlyGateway.
	CloudRouteNextHopTypeEgressOnlyGateway CloudRouteNextHopTypeValue = "EgressOnlyGateway"

	// CloudRouteNextHopTypeGateway represents the value Gateway.
	CloudRouteNextHopTypeGateway CloudRouteNextHopTypeValue = "Gateway"

	// CloudRouteNextHopTypeInstance represents the value Instance.
	CloudRouteNextHopTypeInstance CloudRouteNextHopTypeValue = "Instance"

	// CloudRouteNextHopTypeLocalGateway represents the value LocalGateway.
	CloudRouteNextHopTypeLocalGateway CloudRouteNextHopTypeValue = "LocalGateway"

	// CloudRouteNextHopTypeNATGateway represents the value NATGateway.
	CloudRouteNextHopTypeNATGateway CloudRouteNextHopTypeValue = "NATGateway"

	// CloudRouteNextHopTypeNetworkInterface represents the value NetworkInterface.
	CloudRouteNextHopTypeNetworkInterface CloudRouteNextHopTypeValue = "NetworkInterface"

	// CloudRouteNextHopTypeTransitGateway represents the value TransitGateway.
	CloudRouteNextHopTypeTransitGateway CloudRouteNextHopTypeValue = "TransitGateway"

	// CloudRouteNextHopTypeTransitGatewayAttachment represents the value TransitGatewayAttachment.
	CloudRouteNextHopTypeTransitGatewayAttachment CloudRouteNextHopTypeValue = "TransitGatewayAttachment"

	// CloudRouteNextHopTypeVPCPeeringConnection represents the value VPCPeeringConnection.
	CloudRouteNextHopTypeVPCPeeringConnection CloudRouteNextHopTypeValue = "VPCPeeringConnection"
)

// CloudRouteIdentity represents the Identity of the object.
var CloudRouteIdentity = elemental.Identity{
	Name:     "cloudroute",
	Category: "cloudroutes",
	Package:  "pcn",
	Private:  false,
}

// CloudRoutesList represents a list of CloudRoutes
type CloudRoutesList []*CloudRoute

// Identity returns the identity of the objects in the list.
func (o CloudRoutesList) Identity() elemental.Identity {

	return CloudRouteIdentity
}

// Copy returns a pointer to a copy the CloudRoutesList.
func (o CloudRoutesList) Copy() elemental.Identifiables {

	copy := append(CloudRoutesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudRoutesList.
func (o CloudRoutesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudRoutesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudRoute))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudRoutesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudRoutesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudRoutesList converted to SparseCloudRoutesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudRoutesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudRoutesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudRoute)
	}

	return out
}

// Version returns the version of the content.
func (o CloudRoutesList) Version() int {

	return 1
}

// CloudRoute represents the model of a cloudroute
type CloudRoute struct {
	// The Destination CIDR for the route.
	DestinationIPv4CIDR string `json:"destinationIPv4CIDR" msgpack:"destinationIPv4CIDR" bson:"destinationipv4cidr" mapstructure:"destinationIPv4CIDR,omitempty"`

	// The destination IPV6 CIDR for the route.
	DestinationIPv6CIDR string `json:"destinationIPv6CIDR" msgpack:"destinationIPv6CIDR" bson:"destinationipv6cidr" mapstructure:"destinationIPv6CIDR,omitempty"`

	// The ID of the next hop object.
	NextHopID string `json:"nextHopID" msgpack:"nextHopID" bson:"nexthopid" mapstructure:"nextHopID,omitempty"`

	// The type of the next hop.
	NextHopType CloudRouteNextHopTypeValue `json:"nextHopType" msgpack:"nextHopType" bson:"nexthoptype" mapstructure:"nextHopType,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudRoute returns a new *CloudRoute
func NewCloudRoute() *CloudRoute {

	return &CloudRoute{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *CloudRoute) Identity() elemental.Identity {

	return CloudRouteIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudRoute) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudRoute) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudRoute) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudRoute{}

	s.DestinationIPv4CIDR = o.DestinationIPv4CIDR
	s.DestinationIPv6CIDR = o.DestinationIPv6CIDR
	s.NextHopID = o.NextHopID
	s.NextHopType = o.NextHopType

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudRoute) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudRoute{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.DestinationIPv4CIDR = s.DestinationIPv4CIDR
	o.DestinationIPv6CIDR = s.DestinationIPv6CIDR
	o.NextHopID = s.NextHopID
	o.NextHopType = s.NextHopType

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudRoute) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudRoute) BleveType() string {

	return "cloudroute"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudRoute) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudRoute) Doc() string {

	return `Describes a route in a route table.`
}

func (o *CloudRoute) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudRoute) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudRoute{
			DestinationIPv4CIDR: &o.DestinationIPv4CIDR,
			DestinationIPv6CIDR: &o.DestinationIPv6CIDR,
			NextHopID:           &o.NextHopID,
			NextHopType:         &o.NextHopType,
		}
	}

	sp := &SparseCloudRoute{}
	for _, f := range fields {
		switch f {
		case "destinationIPv4CIDR":
			sp.DestinationIPv4CIDR = &(o.DestinationIPv4CIDR)
		case "destinationIPv6CIDR":
			sp.DestinationIPv6CIDR = &(o.DestinationIPv6CIDR)
		case "nextHopID":
			sp.NextHopID = &(o.NextHopID)
		case "nextHopType":
			sp.NextHopType = &(o.NextHopType)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudRoute to the object.
func (o *CloudRoute) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudRoute)
	if so.DestinationIPv4CIDR != nil {
		o.DestinationIPv4CIDR = *so.DestinationIPv4CIDR
	}
	if so.DestinationIPv6CIDR != nil {
		o.DestinationIPv6CIDR = *so.DestinationIPv6CIDR
	}
	if so.NextHopID != nil {
		o.NextHopID = *so.NextHopID
	}
	if so.NextHopType != nil {
		o.NextHopType = *so.NextHopType
	}
}

// DeepCopy returns a deep copy if the CloudRoute.
func (o *CloudRoute) DeepCopy() *CloudRoute {

	if o == nil {
		return nil
	}

	out := &CloudRoute{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudRoute.
func (o *CloudRoute) DeepCopyInto(out *CloudRoute) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudRoute: %s", err))
	}

	*out = *target.(*CloudRoute)
}

// Validate valides the current information stored into the structure.
func (o *CloudRoute) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateOptionalCIDR("destinationIPv4CIDR", o.DestinationIPv4CIDR); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDR("destinationIPv6CIDR", o.DestinationIPv6CIDR); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("nextHopType", string(o.NextHopType), []string{"EgressOnlyGateway", "Gateway", "Instance", "LocalGateway", "NATGateway", "NetworkInterface", "TransitGateway", "VPCPeeringConnection", "TransitGatewayAttachment"}, false); err != nil {
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
func (*CloudRoute) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudRouteAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudRouteLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudRoute) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudRouteAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudRoute) ValueForAttribute(name string) interface{} {

	switch name {
	case "destinationIPv4CIDR":
		return o.DestinationIPv4CIDR
	case "destinationIPv6CIDR":
		return o.DestinationIPv6CIDR
	case "nextHopID":
		return o.NextHopID
	case "nextHopType":
		return o.NextHopType
	}

	return nil
}

// CloudRouteAttributesMap represents the map of attribute for CloudRoute.
var CloudRouteAttributesMap = map[string]elemental.AttributeSpecification{
	"DestinationIPv4CIDR": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationipv4cidr",
		ConvertedName:  "DestinationIPv4CIDR",
		Description:    `The Destination CIDR for the route.`,
		Exposed:        true,
		Name:           "destinationIPv4CIDR",
		Stored:         true,
		Type:           "string",
	},
	"DestinationIPv6CIDR": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationipv6cidr",
		ConvertedName:  "DestinationIPv6CIDR",
		Description:    `The destination IPV6 CIDR for the route.`,
		Exposed:        true,
		Name:           "destinationIPv6CIDR",
		Stored:         true,
		Type:           "string",
	},
	"NextHopID": {
		AllowedChoices: []string{},
		BSONFieldName:  "nexthopid",
		ConvertedName:  "NextHopID",
		Description:    `The ID of the next hop object.`,
		Exposed:        true,
		Name:           "nextHopID",
		Stored:         true,
		Type:           "string",
	},
	"NextHopType": {
		AllowedChoices: []string{"EgressOnlyGateway", "Gateway", "Instance", "LocalGateway", "NATGateway", "NetworkInterface", "TransitGateway", "VPCPeeringConnection", "TransitGatewayAttachment"},
		BSONFieldName:  "nexthoptype",
		ConvertedName:  "NextHopType",
		Description:    `The type of the next hop.`,
		Exposed:        true,
		Name:           "nextHopType",
		Stored:         true,
		Type:           "enum",
	},
}

// CloudRouteLowerCaseAttributesMap represents the map of attribute for CloudRoute.
var CloudRouteLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"destinationipv4cidr": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationipv4cidr",
		ConvertedName:  "DestinationIPv4CIDR",
		Description:    `The Destination CIDR for the route.`,
		Exposed:        true,
		Name:           "destinationIPv4CIDR",
		Stored:         true,
		Type:           "string",
	},
	"destinationipv6cidr": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationipv6cidr",
		ConvertedName:  "DestinationIPv6CIDR",
		Description:    `The destination IPV6 CIDR for the route.`,
		Exposed:        true,
		Name:           "destinationIPv6CIDR",
		Stored:         true,
		Type:           "string",
	},
	"nexthopid": {
		AllowedChoices: []string{},
		BSONFieldName:  "nexthopid",
		ConvertedName:  "NextHopID",
		Description:    `The ID of the next hop object.`,
		Exposed:        true,
		Name:           "nextHopID",
		Stored:         true,
		Type:           "string",
	},
	"nexthoptype": {
		AllowedChoices: []string{"EgressOnlyGateway", "Gateway", "Instance", "LocalGateway", "NATGateway", "NetworkInterface", "TransitGateway", "VPCPeeringConnection", "TransitGatewayAttachment"},
		BSONFieldName:  "nexthoptype",
		ConvertedName:  "NextHopType",
		Description:    `The type of the next hop.`,
		Exposed:        true,
		Name:           "nextHopType",
		Stored:         true,
		Type:           "enum",
	},
}

// SparseCloudRoutesList represents a list of SparseCloudRoutes
type SparseCloudRoutesList []*SparseCloudRoute

// Identity returns the identity of the objects in the list.
func (o SparseCloudRoutesList) Identity() elemental.Identity {

	return CloudRouteIdentity
}

// Copy returns a pointer to a copy the SparseCloudRoutesList.
func (o SparseCloudRoutesList) Copy() elemental.Identifiables {

	copy := append(SparseCloudRoutesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudRoutesList.
func (o SparseCloudRoutesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudRoutesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudRoute))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudRoutesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudRoutesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudRoutesList converted to CloudRoutesList.
func (o SparseCloudRoutesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudRoutesList) Version() int {

	return 1
}

// SparseCloudRoute represents the sparse version of a cloudroute.
type SparseCloudRoute struct {
	// The Destination CIDR for the route.
	DestinationIPv4CIDR *string `json:"destinationIPv4CIDR,omitempty" msgpack:"destinationIPv4CIDR,omitempty" bson:"destinationipv4cidr,omitempty" mapstructure:"destinationIPv4CIDR,omitempty"`

	// The destination IPV6 CIDR for the route.
	DestinationIPv6CIDR *string `json:"destinationIPv6CIDR,omitempty" msgpack:"destinationIPv6CIDR,omitempty" bson:"destinationipv6cidr,omitempty" mapstructure:"destinationIPv6CIDR,omitempty"`

	// The ID of the next hop object.
	NextHopID *string `json:"nextHopID,omitempty" msgpack:"nextHopID,omitempty" bson:"nexthopid,omitempty" mapstructure:"nextHopID,omitempty"`

	// The type of the next hop.
	NextHopType *CloudRouteNextHopTypeValue `json:"nextHopType,omitempty" msgpack:"nextHopType,omitempty" bson:"nexthoptype,omitempty" mapstructure:"nextHopType,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudRoute returns a new  SparseCloudRoute.
func NewSparseCloudRoute() *SparseCloudRoute {
	return &SparseCloudRoute{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudRoute) Identity() elemental.Identity {

	return CloudRouteIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudRoute) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudRoute) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudRoute) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudRoute{}

	if o.DestinationIPv4CIDR != nil {
		s.DestinationIPv4CIDR = o.DestinationIPv4CIDR
	}
	if o.DestinationIPv6CIDR != nil {
		s.DestinationIPv6CIDR = o.DestinationIPv6CIDR
	}
	if o.NextHopID != nil {
		s.NextHopID = o.NextHopID
	}
	if o.NextHopType != nil {
		s.NextHopType = o.NextHopType
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudRoute) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudRoute{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.DestinationIPv4CIDR != nil {
		o.DestinationIPv4CIDR = s.DestinationIPv4CIDR
	}
	if s.DestinationIPv6CIDR != nil {
		o.DestinationIPv6CIDR = s.DestinationIPv6CIDR
	}
	if s.NextHopID != nil {
		o.NextHopID = s.NextHopID
	}
	if s.NextHopType != nil {
		o.NextHopType = s.NextHopType
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCloudRoute) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudRoute) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudRoute()
	if o.DestinationIPv4CIDR != nil {
		out.DestinationIPv4CIDR = *o.DestinationIPv4CIDR
	}
	if o.DestinationIPv6CIDR != nil {
		out.DestinationIPv6CIDR = *o.DestinationIPv6CIDR
	}
	if o.NextHopID != nil {
		out.NextHopID = *o.NextHopID
	}
	if o.NextHopType != nil {
		out.NextHopType = *o.NextHopType
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCloudRoute.
func (o *SparseCloudRoute) DeepCopy() *SparseCloudRoute {

	if o == nil {
		return nil
	}

	out := &SparseCloudRoute{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudRoute.
func (o *SparseCloudRoute) DeepCopyInto(out *SparseCloudRoute) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudRoute: %s", err))
	}

	*out = *target.(*SparseCloudRoute)
}

type mongoAttributesCloudRoute struct {
	DestinationIPv4CIDR string                     `bson:"destinationipv4cidr"`
	DestinationIPv6CIDR string                     `bson:"destinationipv6cidr"`
	NextHopID           string                     `bson:"nexthopid"`
	NextHopType         CloudRouteNextHopTypeValue `bson:"nexthoptype"`
}
type mongoAttributesSparseCloudRoute struct {
	DestinationIPv4CIDR *string                     `bson:"destinationipv4cidr,omitempty"`
	DestinationIPv6CIDR *string                     `bson:"destinationipv6cidr,omitempty"`
	NextHopID           *string                     `bson:"nexthopid,omitempty"`
	NextHopType         *CloudRouteNextHopTypeValue `bson:"nexthoptype,omitempty"`
}
