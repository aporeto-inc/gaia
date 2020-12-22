package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudEndpointCloudTypeValue represents the possible values for attribute "cloudType".
type CloudEndpointCloudTypeValue string

const (
	// CloudEndpointCloudTypeAWS represents the value AWS.
	CloudEndpointCloudTypeAWS CloudEndpointCloudTypeValue = "AWS"

	// CloudEndpointCloudTypeGCP represents the value GCP.
	CloudEndpointCloudTypeGCP CloudEndpointCloudTypeValue = "GCP"
)

// CloudEndpointTypeValue represents the possible values for attribute "type".
type CloudEndpointTypeValue string

const (
	// CloudEndpointTypeGateway represents the value Gateway.
	CloudEndpointTypeGateway CloudEndpointTypeValue = "Gateway"

	// CloudEndpointTypeInstance represents the value Instance.
	CloudEndpointTypeInstance CloudEndpointTypeValue = "Instance"

	// CloudEndpointTypeLoadBalancer represents the value LoadBalancer.
	CloudEndpointTypeLoadBalancer CloudEndpointTypeValue = "LoadBalancer"

	// CloudEndpointTypePeeringConnection represents the value PeeringConnection.
	CloudEndpointTypePeeringConnection CloudEndpointTypeValue = "PeeringConnection"

	// CloudEndpointTypeService represents the value Service.
	CloudEndpointTypeService CloudEndpointTypeValue = "Service"

	// CloudEndpointTypeTransitGateway represents the value TransitGateway.
	CloudEndpointTypeTransitGateway CloudEndpointTypeValue = "TransitGateway"
)

// CloudEndpointIdentity represents the Identity of the object.
var CloudEndpointIdentity = elemental.Identity{
	Name:     "cloudendpoint",
	Category: "cloudendpoints",
	Package:  "pcn",
	Private:  false,
}

// CloudEndpointsList represents a list of CloudEndpoints
type CloudEndpointsList []*CloudEndpoint

// Identity returns the identity of the objects in the list.
func (o CloudEndpointsList) Identity() elemental.Identity {

	return CloudEndpointIdentity
}

// Copy returns a pointer to a copy the CloudEndpointsList.
func (o CloudEndpointsList) Copy() elemental.Identifiables {

	copy := append(CloudEndpointsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudEndpointsList.
func (o CloudEndpointsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudEndpointsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudEndpoint))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudEndpointsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudEndpointsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the CloudEndpointsList converted to SparseCloudEndpointsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudEndpointsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudEndpointsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudEndpoint)
	}

	return out
}

// Version returns the version of the content.
func (o CloudEndpointsList) Version() int {

	return 1
}

// CloudEndpoint represents the model of a cloudendpoint
type CloudEndpoint struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Restricted Resource Name.
	RRN string `json:"rrn" msgpack:"rrn" bson:"rrn" mapstructure:"rrn,omitempty"`

	// Indicates that the endpoint is directly attached to the VPC. In this case the
	// attachedInterfaces is empty. In general this is only valid for endpoint type
	// Gateway and Peering Connection.
	VPCAttached bool `json:"VPCAttached" msgpack:"VPCAttached" bson:"vpcattached" mapstructure:"VPCAttached,omitempty"`

	// The list of VPCs that this endpoint is directly attached to.
	VPCAttachments []string `json:"VPCAttachments" msgpack:"VPCAttachments" bson:"vpcattachments" mapstructure:"VPCAttachments,omitempty"`

	// Cloud account ID associated with the entity.
	AccountID string `json:"accountId" msgpack:"accountId" bson:"accountid" mapstructure:"accountId,omitempty"`

	// Cloud account name associated with the entity.
	AccountName string `json:"accountName" msgpack:"accountName" bson:"accountname" mapstructure:"accountName,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of route tables associated with this endpoint if it is a transit gateway.
	AssociatedRouteTables []string `json:"associatedRouteTables" msgpack:"associatedRouteTables" bson:"associatedroutetables" mapstructure:"associatedRouteTables,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// A list of interfaces attached with the endpoint. In some cases endpoints can
	// have more than one interface.
	AttachedInterfaces []string `json:"attachedInterfaces" msgpack:"attachedInterfaces" bson:"attachedinterfaces" mapstructure:"attachedInterfaces,omitempty"`

	// Cloud type of the entity.
	CloudType CloudEndpointCloudTypeValue `json:"cloudType" msgpack:"cloudType" bson:"cloudtype" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// If the endpoint has multiple connections and forwarding can be enabled between
	// them.
	ForwardingEnabled bool `json:"forwardingEnabled" msgpack:"forwardingEnabled" bson:"forwardingenabled" mapstructure:"forwardingEnabled,omitempty"`

	// Timestamp of when object was created.
	InsertTS int `json:"insertTs,omitempty" msgpack:"insertTs,omitempty" bson:"insertts,omitempty" mapstructure:"insertTs,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// ID of the cloud provider object.
	NativeID string `json:"nativeID" msgpack:"nativeID" bson:"nativeid" mapstructure:"nativeID,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// ID of the region associated with the entity.
	RegionID string `json:"regionId" msgpack:"regionId" bson:"regionid" mapstructure:"regionId,omitempty"`

	// Region name associated with the entity.
	RegionName string `json:"regionName" msgpack:"regionName" bson:"regionname" mapstructure:"regionName,omitempty"`

	// Internal representation of object tags.
	Tags map[string]string `json:"tags" msgpack:"tags" bson:"tags" mapstructure:"tags,omitempty"`

	// Type of the endpoint.
	Type CloudEndpointTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Object resource URL access.
	Url string `json:"url" msgpack:"url" bson:"url" mapstructure:"url,omitempty"`

	// ID of the host VPC.
	VpcID string `json:"vpcId" msgpack:"vpcId" bson:"vpcid" mapstructure:"vpcId,omitempty"`

	// Name of the host VPC.
	VpcName string `json:"vpcName" msgpack:"vpcName" bson:"vpcname" mapstructure:"vpcName,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudEndpoint returns a new *CloudEndpoint
func NewCloudEndpoint() *CloudEndpoint {

	return &CloudEndpoint{
		ModelVersion:          1,
		Annotations:           map[string][]string{},
		AssociatedRouteTables: []string{},
		AssociatedTags:        []string{},
		AttachedInterfaces:    []string{},
		ForwardingEnabled:     true,
		MigrationsLog:         map[string]string{},
		Tags:                  map[string]string{},
		NormalizedTags:        []string{},
		VPCAttachments:        []string{},
		Metadata:              []string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudEndpoint) Identity() elemental.Identity {

	return CloudEndpointIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudEndpoint) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudEndpoint) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEndpoint) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudEndpoint{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.RRN = o.RRN
	s.VPCAttached = o.VPCAttached
	s.VPCAttachments = o.VPCAttachments
	s.AccountID = o.AccountID
	s.AccountName = o.AccountName
	s.Annotations = o.Annotations
	s.AssociatedRouteTables = o.AssociatedRouteTables
	s.AssociatedTags = o.AssociatedTags
	s.AttachedInterfaces = o.AttachedInterfaces
	s.CloudType = o.CloudType
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.Description = o.Description
	s.ForwardingEnabled = o.ForwardingEnabled
	s.InsertTS = o.InsertTS
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NativeID = o.NativeID
	s.NormalizedTags = o.NormalizedTags
	s.Protected = o.Protected
	s.RegionID = o.RegionID
	s.RegionName = o.RegionName
	s.Tags = o.Tags
	s.Type = o.Type
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.Url = o.Url
	s.VpcID = o.VpcID
	s.VpcName = o.VpcName
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEndpoint) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudEndpoint{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.RRN = s.RRN
	o.VPCAttached = s.VPCAttached
	o.VPCAttachments = s.VPCAttachments
	o.AccountID = s.AccountID
	o.AccountName = s.AccountName
	o.Annotations = s.Annotations
	o.AssociatedRouteTables = s.AssociatedRouteTables
	o.AssociatedTags = s.AssociatedTags
	o.AttachedInterfaces = s.AttachedInterfaces
	o.CloudType = s.CloudType
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.Description = s.Description
	o.ForwardingEnabled = s.ForwardingEnabled
	o.InsertTS = s.InsertTS
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NativeID = s.NativeID
	o.NormalizedTags = s.NormalizedTags
	o.Protected = s.Protected
	o.RegionID = s.RegionID
	o.RegionName = s.RegionName
	o.Tags = s.Tags
	o.Type = s.Type
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.Url = s.Url
	o.VpcID = s.VpcID
	o.VpcName = s.VpcName
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudEndpoint) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudEndpoint) BleveType() string {

	return "cloudendpoint"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudEndpoint) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *CloudEndpoint) Doc() string {

	return `Manages the list of endpoints available in a cloud deployment.`
}

func (o *CloudEndpoint) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetRRN returns the RRN of the receiver.
func (o *CloudEndpoint) GetRRN() string {

	return o.RRN
}

// SetRRN sets the property RRN of the receiver using the given value.
func (o *CloudEndpoint) SetRRN(RRN string) {

	o.RRN = RRN
}

// GetAccountID returns the AccountID of the receiver.
func (o *CloudEndpoint) GetAccountID() string {

	return o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the given value.
func (o *CloudEndpoint) SetAccountID(accountID string) {

	o.AccountID = accountID
}

// GetAccountName returns the AccountName of the receiver.
func (o *CloudEndpoint) GetAccountName() string {

	return o.AccountName
}

// SetAccountName sets the property AccountName of the receiver using the given value.
func (o *CloudEndpoint) SetAccountName(accountName string) {

	o.AccountName = accountName
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CloudEndpoint) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CloudEndpoint) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CloudEndpoint) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CloudEndpoint) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *CloudEndpoint) GetCloudType() CloudEndpointCloudTypeValue {

	return o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the given value.
func (o *CloudEndpoint) SetCloudType(cloudType CloudEndpointCloudTypeValue) {

	o.CloudType = cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CloudEndpoint) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CloudEndpoint) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetDescription returns the Description of the receiver.
func (o *CloudEndpoint) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *CloudEndpoint) SetDescription(description string) {

	o.Description = description
}

// GetInsertTS returns the InsertTS of the receiver.
func (o *CloudEndpoint) GetInsertTS() int {

	return o.InsertTS
}

// SetInsertTS sets the property InsertTS of the receiver using the given value.
func (o *CloudEndpoint) SetInsertTS(insertTS int) {

	o.InsertTS = insertTS
}

// GetMetadata returns the Metadata of the receiver.
func (o *CloudEndpoint) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *CloudEndpoint) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CloudEndpoint) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CloudEndpoint) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *CloudEndpoint) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *CloudEndpoint) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudEndpoint) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudEndpoint) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *CloudEndpoint) GetNativeID() string {

	return o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the given value.
func (o *CloudEndpoint) SetNativeID(nativeID string) {

	o.NativeID = nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CloudEndpoint) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CloudEndpoint) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *CloudEndpoint) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CloudEndpoint) SetProtected(protected bool) {

	o.Protected = protected
}

// GetRegionID returns the RegionID of the receiver.
func (o *CloudEndpoint) GetRegionID() string {

	return o.RegionID
}

// SetRegionID sets the property RegionID of the receiver using the given value.
func (o *CloudEndpoint) SetRegionID(regionID string) {

	o.RegionID = regionID
}

// GetRegionName returns the RegionName of the receiver.
func (o *CloudEndpoint) GetRegionName() string {

	return o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the given value.
func (o *CloudEndpoint) SetRegionName(regionName string) {

	o.RegionName = regionName
}

// GetTags returns the Tags of the receiver.
func (o *CloudEndpoint) GetTags() map[string]string {

	return o.Tags
}

// SetTags sets the property Tags of the receiver using the given value.
func (o *CloudEndpoint) SetTags(tags map[string]string) {

	o.Tags = tags
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CloudEndpoint) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CloudEndpoint) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUrl returns the Url of the receiver.
func (o *CloudEndpoint) GetUrl() string {

	return o.Url
}

// SetUrl sets the property Url of the receiver using the given value.
func (o *CloudEndpoint) SetUrl(url string) {

	o.Url = url
}

// GetVpcID returns the VpcID of the receiver.
func (o *CloudEndpoint) GetVpcID() string {

	return o.VpcID
}

// SetVpcID sets the property VpcID of the receiver using the given value.
func (o *CloudEndpoint) SetVpcID(vpcID string) {

	o.VpcID = vpcID
}

// GetVpcName returns the VpcName of the receiver.
func (o *CloudEndpoint) GetVpcName() string {

	return o.VpcName
}

// SetVpcName sets the property VpcName of the receiver using the given value.
func (o *CloudEndpoint) SetVpcName(vpcName string) {

	o.VpcName = vpcName
}

// GetZHash returns the ZHash of the receiver.
func (o *CloudEndpoint) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CloudEndpoint) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CloudEndpoint) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CloudEndpoint) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudEndpoint) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudEndpoint{
			ID:                    &o.ID,
			RRN:                   &o.RRN,
			VPCAttached:           &o.VPCAttached,
			VPCAttachments:        &o.VPCAttachments,
			AccountID:             &o.AccountID,
			AccountName:           &o.AccountName,
			Annotations:           &o.Annotations,
			AssociatedRouteTables: &o.AssociatedRouteTables,
			AssociatedTags:        &o.AssociatedTags,
			AttachedInterfaces:    &o.AttachedInterfaces,
			CloudType:             &o.CloudType,
			CreateIdempotencyKey:  &o.CreateIdempotencyKey,
			Description:           &o.Description,
			ForwardingEnabled:     &o.ForwardingEnabled,
			InsertTS:              &o.InsertTS,
			Metadata:              &o.Metadata,
			MigrationsLog:         &o.MigrationsLog,
			Name:                  &o.Name,
			Namespace:             &o.Namespace,
			NativeID:              &o.NativeID,
			NormalizedTags:        &o.NormalizedTags,
			Protected:             &o.Protected,
			RegionID:              &o.RegionID,
			RegionName:            &o.RegionName,
			Tags:                  &o.Tags,
			Type:                  &o.Type,
			UpdateIdempotencyKey:  &o.UpdateIdempotencyKey,
			Url:                   &o.Url,
			VpcID:                 &o.VpcID,
			VpcName:               &o.VpcName,
			ZHash:                 &o.ZHash,
			Zone:                  &o.Zone,
		}
	}

	sp := &SparseCloudEndpoint{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "RRN":
			sp.RRN = &(o.RRN)
		case "VPCAttached":
			sp.VPCAttached = &(o.VPCAttached)
		case "VPCAttachments":
			sp.VPCAttachments = &(o.VPCAttachments)
		case "accountID":
			sp.AccountID = &(o.AccountID)
		case "accountName":
			sp.AccountName = &(o.AccountName)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedRouteTables":
			sp.AssociatedRouteTables = &(o.AssociatedRouteTables)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "attachedInterfaces":
			sp.AttachedInterfaces = &(o.AttachedInterfaces)
		case "cloudType":
			sp.CloudType = &(o.CloudType)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "description":
			sp.Description = &(o.Description)
		case "forwardingEnabled":
			sp.ForwardingEnabled = &(o.ForwardingEnabled)
		case "insertTS":
			sp.InsertTS = &(o.InsertTS)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "nativeID":
			sp.NativeID = &(o.NativeID)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "regionID":
			sp.RegionID = &(o.RegionID)
		case "regionName":
			sp.RegionName = &(o.RegionName)
		case "tags":
			sp.Tags = &(o.Tags)
		case "type":
			sp.Type = &(o.Type)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "url":
			sp.Url = &(o.Url)
		case "vpcID":
			sp.VpcID = &(o.VpcID)
		case "vpcName":
			sp.VpcName = &(o.VpcName)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudEndpoint to the object.
func (o *CloudEndpoint) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudEndpoint)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.RRN != nil {
		o.RRN = *so.RRN
	}
	if so.VPCAttached != nil {
		o.VPCAttached = *so.VPCAttached
	}
	if so.VPCAttachments != nil {
		o.VPCAttachments = *so.VPCAttachments
	}
	if so.AccountID != nil {
		o.AccountID = *so.AccountID
	}
	if so.AccountName != nil {
		o.AccountName = *so.AccountName
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedRouteTables != nil {
		o.AssociatedRouteTables = *so.AssociatedRouteTables
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.AttachedInterfaces != nil {
		o.AttachedInterfaces = *so.AttachedInterfaces
	}
	if so.CloudType != nil {
		o.CloudType = *so.CloudType
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.ForwardingEnabled != nil {
		o.ForwardingEnabled = *so.ForwardingEnabled
	}
	if so.InsertTS != nil {
		o.InsertTS = *so.InsertTS
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NativeID != nil {
		o.NativeID = *so.NativeID
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.RegionID != nil {
		o.RegionID = *so.RegionID
	}
	if so.RegionName != nil {
		o.RegionName = *so.RegionName
	}
	if so.Tags != nil {
		o.Tags = *so.Tags
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.Url != nil {
		o.Url = *so.Url
	}
	if so.VpcID != nil {
		o.VpcID = *so.VpcID
	}
	if so.VpcName != nil {
		o.VpcName = *so.VpcName
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the CloudEndpoint.
func (o *CloudEndpoint) DeepCopy() *CloudEndpoint {

	if o == nil {
		return nil
	}

	out := &CloudEndpoint{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudEndpoint.
func (o *CloudEndpoint) DeepCopyInto(out *CloudEndpoint) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudEndpoint: %s", err))
	}

	*out = *target.(*CloudEndpoint)
}

// Validate valides the current information stored into the structure.
func (o *CloudEndpoint) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("cloudType", string(o.CloudType), []string{"AWS", "GCP"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("nativeID", o.NativeID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("nativeID", o.NativeID, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("regionID", o.RegionID, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("regionName", o.RegionName, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Instance", "LoadBalancer", "PeeringConnection", "Service", "Gateway", "TransitGateway"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("url", o.Url, 256, false); err != nil {
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
func (*CloudEndpoint) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudEndpointAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudEndpointLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudEndpoint) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudEndpointAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudEndpoint) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "RRN":
		return o.RRN
	case "VPCAttached":
		return o.VPCAttached
	case "VPCAttachments":
		return o.VPCAttachments
	case "accountID":
		return o.AccountID
	case "accountName":
		return o.AccountName
	case "annotations":
		return o.Annotations
	case "associatedRouteTables":
		return o.AssociatedRouteTables
	case "associatedTags":
		return o.AssociatedTags
	case "attachedInterfaces":
		return o.AttachedInterfaces
	case "cloudType":
		return o.CloudType
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "description":
		return o.Description
	case "forwardingEnabled":
		return o.ForwardingEnabled
	case "insertTS":
		return o.InsertTS
	case "metadata":
		return o.Metadata
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "nativeID":
		return o.NativeID
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "regionID":
		return o.RegionID
	case "regionName":
		return o.RegionName
	case "tags":
		return o.Tags
	case "type":
		return o.Type
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "url":
		return o.Url
	case "vpcID":
		return o.VpcID
	case "vpcName":
		return o.VpcName
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// CloudEndpointAttributesMap represents the map of attribute for CloudEndpoint.
var CloudEndpointAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"RRN": {
		AllowedChoices: []string{},
		BSONFieldName:  "rrn",
		ConvertedName:  "RRN",
		Description:    `Restricted Resource Name.`,
		Exposed:        true,
		Getter:         true,
		Name:           "RRN",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"VPCAttached": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcattached",
		ConvertedName:  "VPCAttached",
		Description: `Indicates that the endpoint is directly attached to the VPC. In this case the
attachedInterfaces is empty. In general this is only valid for endpoint type
Gateway and Peering Connection.`,
		Exposed: true,
		Name:    "VPCAttached",
		Stored:  true,
		Type:    "boolean",
	},
	"VPCAttachments": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcattachments",
		ConvertedName:  "VPCAttachments",
		Description:    `The list of VPCs that this endpoint is directly attached to.`,
		Exposed:        true,
		Name:           "VPCAttachments",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AccountID": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountid",
		ConvertedName:  "AccountID",
		Description:    `Cloud account ID associated with the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "accountID",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"AccountName": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountname",
		ConvertedName:  "AccountName",
		Description:    `Cloud account name associated with the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "accountName",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedRouteTables": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedroutetables",
		ConvertedName:  "AssociatedRouteTables",
		Description:    `List of route tables associated with this endpoint if it is a transit gateway.`,
		Exposed:        true,
		Name:           "associatedRouteTables",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AssociatedTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AttachedInterfaces": {
		AllowedChoices: []string{},
		BSONFieldName:  "attachedinterfaces",
		ConvertedName:  "AttachedInterfaces",
		Description: `A list of interfaces attached with the endpoint. In some cases endpoints can
have more than one interface.`,
		Exposed: true,
		Name:    "attachedInterfaces",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"CloudType": {
		AllowedChoices: []string{"AWS", "GCP"},
		BSONFieldName:  "cloudtype",
		ConvertedName:  "CloudType",
		Description:    `Cloud type of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "cloudType",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "enum",
	},
	"CreateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ForwardingEnabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "forwardingenabled",
		ConvertedName:  "ForwardingEnabled",
		DefaultValue:   true,
		Description: `If the endpoint has multiple connections and forwarding can be enabled between
them.`,
		Exposed: true,
		Name:    "forwardingEnabled",
		Stored:  true,
		Type:    "boolean",
	},
	"InsertTS": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "insertts",
		ConvertedName:  "InsertTS",
		Description:    `Timestamp of when object was created.`,
		Exposed:        true,
		Getter:         true,
		Name:           "insertTS",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"Metadata": {
		AllowedChoices: []string{},
		BSONFieldName:  "metadata",
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"MigrationsLog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NativeID": {
		AllowedChoices: []string{},
		BSONFieldName:  "nativeid",
		ConvertedName:  "NativeID",
		Description:    `ID of the cloud provider object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      256,
		Name:           "nativeID",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"RegionID": {
		AllowedChoices: []string{},
		BSONFieldName:  "regionid",
		ConvertedName:  "RegionID",
		Description:    `ID of the region associated with the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "regionID",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"RegionName": {
		AllowedChoices: []string{},
		BSONFieldName:  "regionname",
		ConvertedName:  "RegionName",
		Description:    `Region name associated with the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "regionName",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Tags": {
		AllowedChoices: []string{},
		BSONFieldName:  "tags",
		ConvertedName:  "Tags",
		Description:    `Internal representation of object tags.`,
		Exposed:        true,
		Getter:         true,
		Name:           "tags",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Type": {
		AllowedChoices: []string{"Instance", "LoadBalancer", "PeeringConnection", "Service", "Gateway", "TransitGateway"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		Description:    `Type of the endpoint.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "enum",
	},
	"UpdateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Url": {
		AllowedChoices: []string{},
		BSONFieldName:  "url",
		ConvertedName:  "Url",
		Description:    `Object resource URL access.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      256,
		Name:           "url",
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"VpcID": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcid",
		ConvertedName:  "VpcID",
		Description:    `ID of the host VPC.`,
		Exposed:        true,
		Getter:         true,
		Name:           "vpcID",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"VpcName": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcname",
		ConvertedName:  "VpcName",
		Description:    `Name of the host VPC.`,
		Exposed:        true,
		Getter:         true,
		Name:           "vpcName",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ZHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// CloudEndpointLowerCaseAttributesMap represents the map of attribute for CloudEndpoint.
var CloudEndpointLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"rrn": {
		AllowedChoices: []string{},
		BSONFieldName:  "rrn",
		ConvertedName:  "RRN",
		Description:    `Restricted Resource Name.`,
		Exposed:        true,
		Getter:         true,
		Name:           "RRN",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"vpcattached": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcattached",
		ConvertedName:  "VPCAttached",
		Description: `Indicates that the endpoint is directly attached to the VPC. In this case the
attachedInterfaces is empty. In general this is only valid for endpoint type
Gateway and Peering Connection.`,
		Exposed: true,
		Name:    "VPCAttached",
		Stored:  true,
		Type:    "boolean",
	},
	"vpcattachments": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcattachments",
		ConvertedName:  "VPCAttachments",
		Description:    `The list of VPCs that this endpoint is directly attached to.`,
		Exposed:        true,
		Name:           "VPCAttachments",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"accountid": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountid",
		ConvertedName:  "AccountID",
		Description:    `Cloud account ID associated with the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "accountID",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"accountname": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountname",
		ConvertedName:  "AccountName",
		Description:    `Cloud account name associated with the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "accountName",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedroutetables": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedroutetables",
		ConvertedName:  "AssociatedRouteTables",
		Description:    `List of route tables associated with this endpoint if it is a transit gateway.`,
		Exposed:        true,
		Name:           "associatedRouteTables",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"associatedtags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"attachedinterfaces": {
		AllowedChoices: []string{},
		BSONFieldName:  "attachedinterfaces",
		ConvertedName:  "AttachedInterfaces",
		Description: `A list of interfaces attached with the endpoint. In some cases endpoints can
have more than one interface.`,
		Exposed: true,
		Name:    "attachedInterfaces",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"cloudtype": {
		AllowedChoices: []string{"AWS", "GCP"},
		BSONFieldName:  "cloudtype",
		ConvertedName:  "CloudType",
		Description:    `Cloud type of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "cloudType",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "enum",
	},
	"createidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"forwardingenabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "forwardingenabled",
		ConvertedName:  "ForwardingEnabled",
		DefaultValue:   true,
		Description: `If the endpoint has multiple connections and forwarding can be enabled between
them.`,
		Exposed: true,
		Name:    "forwardingEnabled",
		Stored:  true,
		Type:    "boolean",
	},
	"insertts": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "insertts",
		ConvertedName:  "InsertTS",
		Description:    `Timestamp of when object was created.`,
		Exposed:        true,
		Getter:         true,
		Name:           "insertTS",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"metadata": {
		AllowedChoices: []string{},
		BSONFieldName:  "metadata",
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"migrationslog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"nativeid": {
		AllowedChoices: []string{},
		BSONFieldName:  "nativeid",
		ConvertedName:  "NativeID",
		Description:    `ID of the cloud provider object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      256,
		Name:           "nativeID",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"regionid": {
		AllowedChoices: []string{},
		BSONFieldName:  "regionid",
		ConvertedName:  "RegionID",
		Description:    `ID of the region associated with the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "regionID",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"regionname": {
		AllowedChoices: []string{},
		BSONFieldName:  "regionname",
		ConvertedName:  "RegionName",
		Description:    `Region name associated with the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "regionName",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"tags": {
		AllowedChoices: []string{},
		BSONFieldName:  "tags",
		ConvertedName:  "Tags",
		Description:    `Internal representation of object tags.`,
		Exposed:        true,
		Getter:         true,
		Name:           "tags",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"type": {
		AllowedChoices: []string{"Instance", "LoadBalancer", "PeeringConnection", "Service", "Gateway", "TransitGateway"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		Description:    `Type of the endpoint.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "enum",
	},
	"updateidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"url": {
		AllowedChoices: []string{},
		BSONFieldName:  "url",
		ConvertedName:  "Url",
		Description:    `Object resource URL access.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      256,
		Name:           "url",
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"vpcid": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcid",
		ConvertedName:  "VpcID",
		Description:    `ID of the host VPC.`,
		Exposed:        true,
		Getter:         true,
		Name:           "vpcID",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"vpcname": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcname",
		ConvertedName:  "VpcName",
		Description:    `Name of the host VPC.`,
		Exposed:        true,
		Getter:         true,
		Name:           "vpcName",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"zhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseCloudEndpointsList represents a list of SparseCloudEndpoints
type SparseCloudEndpointsList []*SparseCloudEndpoint

// Identity returns the identity of the objects in the list.
func (o SparseCloudEndpointsList) Identity() elemental.Identity {

	return CloudEndpointIdentity
}

// Copy returns a pointer to a copy the SparseCloudEndpointsList.
func (o SparseCloudEndpointsList) Copy() elemental.Identifiables {

	copy := append(SparseCloudEndpointsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudEndpointsList.
func (o SparseCloudEndpointsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudEndpointsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudEndpoint))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudEndpointsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudEndpointsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseCloudEndpointsList converted to CloudEndpointsList.
func (o SparseCloudEndpointsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudEndpointsList) Version() int {

	return 1
}

// SparseCloudEndpoint represents the sparse version of a cloudendpoint.
type SparseCloudEndpoint struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Restricted Resource Name.
	RRN *string `json:"rrn,omitempty" msgpack:"rrn,omitempty" bson:"rrn,omitempty" mapstructure:"rrn,omitempty"`

	// Indicates that the endpoint is directly attached to the VPC. In this case the
	// attachedInterfaces is empty. In general this is only valid for endpoint type
	// Gateway and Peering Connection.
	VPCAttached *bool `json:"VPCAttached,omitempty" msgpack:"VPCAttached,omitempty" bson:"vpcattached,omitempty" mapstructure:"VPCAttached,omitempty"`

	// The list of VPCs that this endpoint is directly attached to.
	VPCAttachments *[]string `json:"VPCAttachments,omitempty" msgpack:"VPCAttachments,omitempty" bson:"vpcattachments,omitempty" mapstructure:"VPCAttachments,omitempty"`

	// Cloud account ID associated with the entity.
	AccountID *string `json:"accountId,omitempty" msgpack:"accountId,omitempty" bson:"accountid,omitempty" mapstructure:"accountId,omitempty"`

	// Cloud account name associated with the entity.
	AccountName *string `json:"accountName,omitempty" msgpack:"accountName,omitempty" bson:"accountname,omitempty" mapstructure:"accountName,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of route tables associated with this endpoint if it is a transit gateway.
	AssociatedRouteTables *[]string `json:"associatedRouteTables,omitempty" msgpack:"associatedRouteTables,omitempty" bson:"associatedroutetables,omitempty" mapstructure:"associatedRouteTables,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// A list of interfaces attached with the endpoint. In some cases endpoints can
	// have more than one interface.
	AttachedInterfaces *[]string `json:"attachedInterfaces,omitempty" msgpack:"attachedInterfaces,omitempty" bson:"attachedinterfaces,omitempty" mapstructure:"attachedInterfaces,omitempty"`

	// Cloud type of the entity.
	CloudType *CloudEndpointCloudTypeValue `json:"cloudType,omitempty" msgpack:"cloudType,omitempty" bson:"cloudtype,omitempty" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// If the endpoint has multiple connections and forwarding can be enabled between
	// them.
	ForwardingEnabled *bool `json:"forwardingEnabled,omitempty" msgpack:"forwardingEnabled,omitempty" bson:"forwardingenabled,omitempty" mapstructure:"forwardingEnabled,omitempty"`

	// Timestamp of when object was created.
	InsertTS *int `json:"insertTs,omitempty" msgpack:"insertTs,omitempty" bson:"insertts,omitempty" mapstructure:"insertTs,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// ID of the cloud provider object.
	NativeID *string `json:"nativeID,omitempty" msgpack:"nativeID,omitempty" bson:"nativeid,omitempty" mapstructure:"nativeID,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// ID of the region associated with the entity.
	RegionID *string `json:"regionId,omitempty" msgpack:"regionId,omitempty" bson:"regionid,omitempty" mapstructure:"regionId,omitempty"`

	// Region name associated with the entity.
	RegionName *string `json:"regionName,omitempty" msgpack:"regionName,omitempty" bson:"regionname,omitempty" mapstructure:"regionName,omitempty"`

	// Internal representation of object tags.
	Tags *map[string]string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"tags,omitempty" mapstructure:"tags,omitempty"`

	// Type of the endpoint.
	Type *CloudEndpointTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Object resource URL access.
	Url *string `json:"url,omitempty" msgpack:"url,omitempty" bson:"url,omitempty" mapstructure:"url,omitempty"`

	// ID of the host VPC.
	VpcID *string `json:"vpcId,omitempty" msgpack:"vpcId,omitempty" bson:"vpcid,omitempty" mapstructure:"vpcId,omitempty"`

	// Name of the host VPC.
	VpcName *string `json:"vpcName,omitempty" msgpack:"vpcName,omitempty" bson:"vpcname,omitempty" mapstructure:"vpcName,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudEndpoint returns a new  SparseCloudEndpoint.
func NewSparseCloudEndpoint() *SparseCloudEndpoint {
	return &SparseCloudEndpoint{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudEndpoint) Identity() elemental.Identity {

	return CloudEndpointIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudEndpoint) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudEndpoint) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudEndpoint) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudEndpoint{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.RRN != nil {
		s.RRN = o.RRN
	}
	if o.VPCAttached != nil {
		s.VPCAttached = o.VPCAttached
	}
	if o.VPCAttachments != nil {
		s.VPCAttachments = o.VPCAttachments
	}
	if o.AccountID != nil {
		s.AccountID = o.AccountID
	}
	if o.AccountName != nil {
		s.AccountName = o.AccountName
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedRouteTables != nil {
		s.AssociatedRouteTables = o.AssociatedRouteTables
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.AttachedInterfaces != nil {
		s.AttachedInterfaces = o.AttachedInterfaces
	}
	if o.CloudType != nil {
		s.CloudType = o.CloudType
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.ForwardingEnabled != nil {
		s.ForwardingEnabled = o.ForwardingEnabled
	}
	if o.InsertTS != nil {
		s.InsertTS = o.InsertTS
	}
	if o.Metadata != nil {
		s.Metadata = o.Metadata
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NativeID != nil {
		s.NativeID = o.NativeID
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.RegionID != nil {
		s.RegionID = o.RegionID
	}
	if o.RegionName != nil {
		s.RegionName = o.RegionName
	}
	if o.Tags != nil {
		s.Tags = o.Tags
	}
	if o.Type != nil {
		s.Type = o.Type
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.Url != nil {
		s.Url = o.Url
	}
	if o.VpcID != nil {
		s.VpcID = o.VpcID
	}
	if o.VpcName != nil {
		s.VpcName = o.VpcName
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudEndpoint) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudEndpoint{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.RRN != nil {
		o.RRN = s.RRN
	}
	if s.VPCAttached != nil {
		o.VPCAttached = s.VPCAttached
	}
	if s.VPCAttachments != nil {
		o.VPCAttachments = s.VPCAttachments
	}
	if s.AccountID != nil {
		o.AccountID = s.AccountID
	}
	if s.AccountName != nil {
		o.AccountName = s.AccountName
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedRouteTables != nil {
		o.AssociatedRouteTables = s.AssociatedRouteTables
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.AttachedInterfaces != nil {
		o.AttachedInterfaces = s.AttachedInterfaces
	}
	if s.CloudType != nil {
		o.CloudType = s.CloudType
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.ForwardingEnabled != nil {
		o.ForwardingEnabled = s.ForwardingEnabled
	}
	if s.InsertTS != nil {
		o.InsertTS = s.InsertTS
	}
	if s.Metadata != nil {
		o.Metadata = s.Metadata
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NativeID != nil {
		o.NativeID = s.NativeID
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.RegionID != nil {
		o.RegionID = s.RegionID
	}
	if s.RegionName != nil {
		o.RegionName = s.RegionName
	}
	if s.Tags != nil {
		o.Tags = s.Tags
	}
	if s.Type != nil {
		o.Type = s.Type
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.Url != nil {
		o.Url = s.Url
	}
	if s.VpcID != nil {
		o.VpcID = s.VpcID
	}
	if s.VpcName != nil {
		o.VpcName = s.VpcName
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCloudEndpoint) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudEndpoint) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudEndpoint()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.RRN != nil {
		out.RRN = *o.RRN
	}
	if o.VPCAttached != nil {
		out.VPCAttached = *o.VPCAttached
	}
	if o.VPCAttachments != nil {
		out.VPCAttachments = *o.VPCAttachments
	}
	if o.AccountID != nil {
		out.AccountID = *o.AccountID
	}
	if o.AccountName != nil {
		out.AccountName = *o.AccountName
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedRouteTables != nil {
		out.AssociatedRouteTables = *o.AssociatedRouteTables
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.AttachedInterfaces != nil {
		out.AttachedInterfaces = *o.AttachedInterfaces
	}
	if o.CloudType != nil {
		out.CloudType = *o.CloudType
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.ForwardingEnabled != nil {
		out.ForwardingEnabled = *o.ForwardingEnabled
	}
	if o.InsertTS != nil {
		out.InsertTS = *o.InsertTS
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NativeID != nil {
		out.NativeID = *o.NativeID
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.RegionID != nil {
		out.RegionID = *o.RegionID
	}
	if o.RegionName != nil {
		out.RegionName = *o.RegionName
	}
	if o.Tags != nil {
		out.Tags = *o.Tags
	}
	if o.Type != nil {
		out.Type = *o.Type
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.Url != nil {
		out.Url = *o.Url
	}
	if o.VpcID != nil {
		out.VpcID = *o.VpcID
	}
	if o.VpcName != nil {
		out.VpcName = *o.VpcName
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetRRN returns the RRN of the receiver.
func (o *SparseCloudEndpoint) GetRRN() (out string) {

	if o.RRN == nil {
		return
	}

	return *o.RRN
}

// SetRRN sets the property RRN of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetRRN(RRN string) {

	o.RRN = &RRN
}

// GetAccountID returns the AccountID of the receiver.
func (o *SparseCloudEndpoint) GetAccountID() (out string) {

	if o.AccountID == nil {
		return
	}

	return *o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetAccountID(accountID string) {

	o.AccountID = &accountID
}

// GetAccountName returns the AccountName of the receiver.
func (o *SparseCloudEndpoint) GetAccountName() (out string) {

	if o.AccountName == nil {
		return
	}

	return *o.AccountName
}

// SetAccountName sets the property AccountName of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetAccountName(accountName string) {

	o.AccountName = &accountName
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseCloudEndpoint) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCloudEndpoint) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *SparseCloudEndpoint) GetCloudType() (out CloudEndpointCloudTypeValue) {

	if o.CloudType == nil {
		return
	}

	return *o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetCloudType(cloudType CloudEndpointCloudTypeValue) {

	o.CloudType = &cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCloudEndpoint) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetDescription returns the Description of the receiver.
func (o *SparseCloudEndpoint) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetDescription(description string) {

	o.Description = &description
}

// GetInsertTS returns the InsertTS of the receiver.
func (o *SparseCloudEndpoint) GetInsertTS() (out int) {

	if o.InsertTS == nil {
		return
	}

	return *o.InsertTS
}

// SetInsertTS sets the property InsertTS of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetInsertTS(insertTS int) {

	o.InsertTS = &insertTS
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseCloudEndpoint) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCloudEndpoint) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseCloudEndpoint) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudEndpoint) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *SparseCloudEndpoint) GetNativeID() (out string) {

	if o.NativeID == nil {
		return
	}

	return *o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetNativeID(nativeID string) {

	o.NativeID = &nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCloudEndpoint) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCloudEndpoint) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetRegionID returns the RegionID of the receiver.
func (o *SparseCloudEndpoint) GetRegionID() (out string) {

	if o.RegionID == nil {
		return
	}

	return *o.RegionID
}

// SetRegionID sets the property RegionID of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetRegionID(regionID string) {

	o.RegionID = &regionID
}

// GetRegionName returns the RegionName of the receiver.
func (o *SparseCloudEndpoint) GetRegionName() (out string) {

	if o.RegionName == nil {
		return
	}

	return *o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetRegionName(regionName string) {

	o.RegionName = &regionName
}

// GetTags returns the Tags of the receiver.
func (o *SparseCloudEndpoint) GetTags() (out map[string]string) {

	if o.Tags == nil {
		return
	}

	return *o.Tags
}

// SetTags sets the property Tags of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetTags(tags map[string]string) {

	o.Tags = &tags
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCloudEndpoint) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUrl returns the Url of the receiver.
func (o *SparseCloudEndpoint) GetUrl() (out string) {

	if o.Url == nil {
		return
	}

	return *o.Url
}

// SetUrl sets the property Url of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetUrl(url string) {

	o.Url = &url
}

// GetVpcID returns the VpcID of the receiver.
func (o *SparseCloudEndpoint) GetVpcID() (out string) {

	if o.VpcID == nil {
		return
	}

	return *o.VpcID
}

// SetVpcID sets the property VpcID of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetVpcID(vpcID string) {

	o.VpcID = &vpcID
}

// GetVpcName returns the VpcName of the receiver.
func (o *SparseCloudEndpoint) GetVpcName() (out string) {

	if o.VpcName == nil {
		return
	}

	return *o.VpcName
}

// SetVpcName sets the property VpcName of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetVpcName(vpcName string) {

	o.VpcName = &vpcName
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCloudEndpoint) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCloudEndpoint) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCloudEndpoint) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCloudEndpoint.
func (o *SparseCloudEndpoint) DeepCopy() *SparseCloudEndpoint {

	if o == nil {
		return nil
	}

	out := &SparseCloudEndpoint{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudEndpoint.
func (o *SparseCloudEndpoint) DeepCopyInto(out *SparseCloudEndpoint) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudEndpoint: %s", err))
	}

	*out = *target.(*SparseCloudEndpoint)
}

type mongoAttributesCloudEndpoint struct {
	ID                    bson.ObjectId               `bson:"_id,omitempty"`
	RRN                   string                      `bson:"rrn"`
	VPCAttached           bool                        `bson:"vpcattached"`
	VPCAttachments        []string                    `bson:"vpcattachments"`
	AccountID             string                      `bson:"accountid"`
	AccountName           string                      `bson:"accountname"`
	Annotations           map[string][]string         `bson:"annotations"`
	AssociatedRouteTables []string                    `bson:"associatedroutetables"`
	AssociatedTags        []string                    `bson:"associatedtags"`
	AttachedInterfaces    []string                    `bson:"attachedinterfaces"`
	CloudType             CloudEndpointCloudTypeValue `bson:"cloudtype"`
	CreateIdempotencyKey  string                      `bson:"createidempotencykey"`
	Description           string                      `bson:"description"`
	ForwardingEnabled     bool                        `bson:"forwardingenabled"`
	InsertTS              int                         `bson:"insertts,omitempty"`
	Metadata              []string                    `bson:"metadata"`
	MigrationsLog         map[string]string           `bson:"migrationslog,omitempty"`
	Name                  string                      `bson:"name"`
	Namespace             string                      `bson:"namespace"`
	NativeID              string                      `bson:"nativeid"`
	NormalizedTags        []string                    `bson:"normalizedtags"`
	Protected             bool                        `bson:"protected"`
	RegionID              string                      `bson:"regionid"`
	RegionName            string                      `bson:"regionname"`
	Tags                  map[string]string           `bson:"tags"`
	Type                  CloudEndpointTypeValue      `bson:"type"`
	UpdateIdempotencyKey  string                      `bson:"updateidempotencykey"`
	Url                   string                      `bson:"url"`
	VpcID                 string                      `bson:"vpcid"`
	VpcName               string                      `bson:"vpcname"`
	ZHash                 int                         `bson:"zhash"`
	Zone                  int                         `bson:"zone"`
}
type mongoAttributesSparseCloudEndpoint struct {
	ID                    bson.ObjectId                `bson:"_id,omitempty"`
	RRN                   *string                      `bson:"rrn,omitempty"`
	VPCAttached           *bool                        `bson:"vpcattached,omitempty"`
	VPCAttachments        *[]string                    `bson:"vpcattachments,omitempty"`
	AccountID             *string                      `bson:"accountid,omitempty"`
	AccountName           *string                      `bson:"accountname,omitempty"`
	Annotations           *map[string][]string         `bson:"annotations,omitempty"`
	AssociatedRouteTables *[]string                    `bson:"associatedroutetables,omitempty"`
	AssociatedTags        *[]string                    `bson:"associatedtags,omitempty"`
	AttachedInterfaces    *[]string                    `bson:"attachedinterfaces,omitempty"`
	CloudType             *CloudEndpointCloudTypeValue `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey  *string                      `bson:"createidempotencykey,omitempty"`
	Description           *string                      `bson:"description,omitempty"`
	ForwardingEnabled     *bool                        `bson:"forwardingenabled,omitempty"`
	InsertTS              *int                         `bson:"insertts,omitempty"`
	Metadata              *[]string                    `bson:"metadata,omitempty"`
	MigrationsLog         *map[string]string           `bson:"migrationslog,omitempty"`
	Name                  *string                      `bson:"name,omitempty"`
	Namespace             *string                      `bson:"namespace,omitempty"`
	NativeID              *string                      `bson:"nativeid,omitempty"`
	NormalizedTags        *[]string                    `bson:"normalizedtags,omitempty"`
	Protected             *bool                        `bson:"protected,omitempty"`
	RegionID              *string                      `bson:"regionid,omitempty"`
	RegionName            *string                      `bson:"regionname,omitempty"`
	Tags                  *map[string]string           `bson:"tags,omitempty"`
	Type                  *CloudEndpointTypeValue      `bson:"type,omitempty"`
	UpdateIdempotencyKey  *string                      `bson:"updateidempotencykey,omitempty"`
	Url                   *string                      `bson:"url,omitempty"`
	VpcID                 *string                      `bson:"vpcid,omitempty"`
	VpcName               *string                      `bson:"vpcname,omitempty"`
	ZHash                 *int                         `bson:"zhash,omitempty"`
	Zone                  *int                         `bson:"zone,omitempty"`
}
