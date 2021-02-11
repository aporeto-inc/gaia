package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudNetworkQueryFilterResourceTypeValue represents the possible values for attribute "ResourceType".
type CloudNetworkQueryFilterResourceTypeValue string

const (
	// CloudNetworkQueryFilterResourceTypeInstance represents the value Instance.
	CloudNetworkQueryFilterResourceTypeInstance CloudNetworkQueryFilterResourceTypeValue = "Instance"

	// CloudNetworkQueryFilterResourceTypeInterface represents the value Interface.
	CloudNetworkQueryFilterResourceTypeInterface CloudNetworkQueryFilterResourceTypeValue = "Interface"

	// CloudNetworkQueryFilterResourceTypeProcessingUnit represents the value ProcessingUnit.
	CloudNetworkQueryFilterResourceTypeProcessingUnit CloudNetworkQueryFilterResourceTypeValue = "ProcessingUnit"

	// CloudNetworkQueryFilterResourceTypeService represents the value Service.
	CloudNetworkQueryFilterResourceTypeService CloudNetworkQueryFilterResourceTypeValue = "Service"
)

// CloudNetworkQueryFilter represents the model of a cloudnetworkqueryfilter
type CloudNetworkQueryFilter struct {
	// The accounts that the search must apply to.
	AccountIDs []string `json:"AccountIDs" msgpack:"AccountIDs" bson:"accountids" mapstructure:"AccountIDs,omitempty"`

	// The cloud types that the search must apply to.
	CloudTypes []string `json:"CloudTypes" msgpack:"CloudTypes" bson:"cloudtypes" mapstructure:"CloudTypes,omitempty"`

	// The exact object that the search applies. If ObjectIDs are defined, the rest of
	// the fields are ignored.
	ObjectIDs []string `json:"ObjectIDs" msgpack:"ObjectIDs" bson:"objectids" mapstructure:"ObjectIDs,omitempty"`

	// The region that the search must apply to.
	Regions []string `json:"Regions" msgpack:"Regions" bson:"regions" mapstructure:"Regions,omitempty"`

	// The type of endpoint resource.
	ResourceType CloudNetworkQueryFilterResourceTypeValue `json:"ResourceType" msgpack:"ResourceType" bson:"resourcetype" mapstructure:"ResourceType,omitempty"`

	// The list of security tags associated with the targets of the query.
	SecurityTags []string `json:"SecurityTags" msgpack:"SecurityTags" bson:"securitytags" mapstructure:"SecurityTags,omitempty"`

	// The subnets where the resources must reside.
	Subnets []string `json:"Subnets" msgpack:"Subnets" bson:"subnets" mapstructure:"Subnets,omitempty"`

	// A list of tags that select the same of endpoints for the query.
	Tags []string `json:"Tags" msgpack:"Tags" bson:"tags" mapstructure:"Tags,omitempty"`

	// The VPC ID of the target resources.
	VPCIDs []string `json:"VPCIDs" msgpack:"VPCIDs" bson:"vpcids" mapstructure:"VPCIDs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudNetworkQueryFilter returns a new *CloudNetworkQueryFilter
func NewCloudNetworkQueryFilter() *CloudNetworkQueryFilter {

	return &CloudNetworkQueryFilter{
		ModelVersion: 1,
		AccountIDs:   []string{},
		CloudTypes:   []string{},
		ObjectIDs:    []string{},
		Regions:      []string{},
		ResourceType: CloudNetworkQueryFilterResourceTypeInstance,
		SecurityTags: []string{},
		Subnets:      []string{},
		Tags:         []string{},
		VPCIDs:       []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkQueryFilter) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudNetworkQueryFilter{}

	s.AccountIDs = o.AccountIDs
	s.CloudTypes = o.CloudTypes
	s.ObjectIDs = o.ObjectIDs
	s.Regions = o.Regions
	s.ResourceType = o.ResourceType
	s.SecurityTags = o.SecurityTags
	s.Subnets = o.Subnets
	s.Tags = o.Tags
	s.VPCIDs = o.VPCIDs

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkQueryFilter) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudNetworkQueryFilter{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.AccountIDs = s.AccountIDs
	o.CloudTypes = s.CloudTypes
	o.ObjectIDs = s.ObjectIDs
	o.Regions = s.Regions
	o.ResourceType = s.ResourceType
	o.SecurityTags = s.SecurityTags
	o.Subnets = s.Subnets
	o.Tags = s.Tags
	o.VPCIDs = s.VPCIDs

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudNetworkQueryFilter) BleveType() string {

	return "cloudnetworkqueryfilter"
}

// DeepCopy returns a deep copy if the CloudNetworkQueryFilter.
func (o *CloudNetworkQueryFilter) DeepCopy() *CloudNetworkQueryFilter {

	if o == nil {
		return nil
	}

	out := &CloudNetworkQueryFilter{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudNetworkQueryFilter.
func (o *CloudNetworkQueryFilter) DeepCopyInto(out *CloudNetworkQueryFilter) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudNetworkQueryFilter: %s", err))
	}

	*out = *target.(*CloudNetworkQueryFilter)
}

// Validate valides the current information stored into the structure.
func (o *CloudNetworkQueryFilter) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("ResourceType", string(o.ResourceType), []string{"Instance", "Interface", "Service", "ProcessingUnit"}, false); err != nil {
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

type mongoAttributesCloudNetworkQueryFilter struct {
	AccountIDs   []string                                 `bson:"accountids"`
	CloudTypes   []string                                 `bson:"cloudtypes"`
	ObjectIDs    []string                                 `bson:"objectids"`
	Regions      []string                                 `bson:"regions"`
	ResourceType CloudNetworkQueryFilterResourceTypeValue `bson:"resourcetype"`
	SecurityTags []string                                 `bson:"securitytags"`
	Subnets      []string                                 `bson:"subnets"`
	Tags         []string                                 `bson:"tags"`
	VPCIDs       []string                                 `bson:"vpcids"`
}
