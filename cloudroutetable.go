package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudRouteTableCloudTypeValue represents the possible values for attribute "cloudType".
type CloudRouteTableCloudTypeValue string

const (
	// CloudRouteTableCloudTypeALIBABA represents the value ALIBABA.
	CloudRouteTableCloudTypeALIBABA CloudRouteTableCloudTypeValue = "ALIBABA"

	// CloudRouteTableCloudTypeAWS represents the value AWS.
	CloudRouteTableCloudTypeAWS CloudRouteTableCloudTypeValue = "AWS"

	// CloudRouteTableCloudTypeAZURE represents the value AZURE.
	CloudRouteTableCloudTypeAZURE CloudRouteTableCloudTypeValue = "AZURE"

	// CloudRouteTableCloudTypeGCP represents the value GCP.
	CloudRouteTableCloudTypeGCP CloudRouteTableCloudTypeValue = "GCP"
)

// CloudRouteTableIdentity represents the Identity of the object.
var CloudRouteTableIdentity = elemental.Identity{
	Name:     "cloudroutetable",
	Category: "cloudroutetables",
	Package:  "pcn",
	Private:  false,
}

// CloudRouteTablesList represents a list of CloudRouteTables
type CloudRouteTablesList []*CloudRouteTable

// Identity returns the identity of the objects in the list.
func (o CloudRouteTablesList) Identity() elemental.Identity {

	return CloudRouteTableIdentity
}

// Copy returns a pointer to a copy the CloudRouteTablesList.
func (o CloudRouteTablesList) Copy() elemental.Identifiables {

	copy := append(CloudRouteTablesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudRouteTablesList.
func (o CloudRouteTablesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudRouteTablesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudRouteTable))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudRouteTablesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudRouteTablesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the CloudRouteTablesList converted to SparseCloudRouteTablesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudRouteTablesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudRouteTablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudRouteTable)
	}

	return out
}

// Version returns the version of the content.
func (o CloudRouteTablesList) Version() int {

	return 1
}

// CloudRouteTable represents the model of a cloudroutetable
type CloudRouteTable struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Restricted Resource Name.
	RRN string `json:"rrn" msgpack:"rrn" bson:"rrn" mapstructure:"rrn,omitempty"`

	// Cloud account ID associated with the entity (matches Prisma Cloud accountID).
	AccountID string `json:"accountId" msgpack:"accountId" bson:"accountid" mapstructure:"accountId,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Prisma Cloud API ID (matches Prisma Cloud API ID).
	ApiID int `json:"apiID" msgpack:"apiID" bson:"apiid" mapstructure:"apiID,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Cloud type of the entity.
	CloudType CloudRouteTableCloudTypeValue `json:"cloudType" msgpack:"cloudType" bson:"cloudtype" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Customer ID as identified by Prisma Cloud.
	CustomerID int `json:"customerID" msgpack:"customerID" bson:"customerid" mapstructure:"customerID,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// The gateway that this route table is associated with.
	GatewayID string `json:"gatewayID" msgpack:"gatewayID" bson:"gatewayid" mapstructure:"gatewayID,omitempty"`

	// Timestamp of when object was created.
	InsertTS int `json:"insertTs,omitempty" msgpack:"insertTs,omitempty" bson:"insertts,omitempty" mapstructure:"insertTs,omitempty"`

	// Indicates that this is the default route table for the VPC.
	MainTable bool `json:"mainTable" msgpack:"mainTable" bson:"maintable" mapstructure:"mainTable,omitempty"`

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

	// Prisma Cloud Resource ID.
	ResourceID int `json:"resourceID" msgpack:"resourceID" bson:"resourceid" mapstructure:"resourceID,omitempty"`

	// Routes associated with this route table.
	Routelist CloudRoutesList `json:"routelist" msgpack:"routelist" bson:"routelist" mapstructure:"routelist,omitempty"`

	// The list of subnets that this route table is associated with.
	SubnetAssociations []string `json:"subnetAssociations" msgpack:"subnetAssociations" bson:"subnetassociations" mapstructure:"subnetAssociations,omitempty"`

	// Internal representation of object tags.
	Tags map[string]string `json:"tags" msgpack:"tags" bson:"tags" mapstructure:"tags,omitempty"`

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

// NewCloudRouteTable returns a new *CloudRouteTable
func NewCloudRouteTable() *CloudRouteTable {

	return &CloudRouteTable{
		ModelVersion:       1,
		Annotations:        map[string][]string{},
		AssociatedTags:     []string{},
		Routelist:          CloudRoutesList{},
		MigrationsLog:      map[string]string{},
		SubnetAssociations: []string{},
		Tags:               map[string]string{},
		NormalizedTags:     []string{},
		Metadata:           []string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudRouteTable) Identity() elemental.Identity {

	return CloudRouteTableIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudRouteTable) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudRouteTable) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudRouteTable) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudRouteTable{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.RRN = o.RRN
	s.AccountID = o.AccountID
	s.Annotations = o.Annotations
	s.ApiID = o.ApiID
	s.AssociatedTags = o.AssociatedTags
	s.CloudType = o.CloudType
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CustomerID = o.CustomerID
	s.Description = o.Description
	s.GatewayID = o.GatewayID
	s.InsertTS = o.InsertTS
	s.MainTable = o.MainTable
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NativeID = o.NativeID
	s.NormalizedTags = o.NormalizedTags
	s.Protected = o.Protected
	s.RegionID = o.RegionID
	s.RegionName = o.RegionName
	s.ResourceID = o.ResourceID
	s.Routelist = o.Routelist
	s.SubnetAssociations = o.SubnetAssociations
	s.Tags = o.Tags
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
func (o *CloudRouteTable) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudRouteTable{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.RRN = s.RRN
	o.AccountID = s.AccountID
	o.Annotations = s.Annotations
	o.ApiID = s.ApiID
	o.AssociatedTags = s.AssociatedTags
	o.CloudType = s.CloudType
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CustomerID = s.CustomerID
	o.Description = s.Description
	o.GatewayID = s.GatewayID
	o.InsertTS = s.InsertTS
	o.MainTable = s.MainTable
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NativeID = s.NativeID
	o.NormalizedTags = s.NormalizedTags
	o.Protected = s.Protected
	o.RegionID = s.RegionID
	o.RegionName = s.RegionName
	o.ResourceID = s.ResourceID
	o.Routelist = s.Routelist
	o.SubnetAssociations = s.SubnetAssociations
	o.Tags = s.Tags
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.Url = s.Url
	o.VpcID = s.VpcID
	o.VpcName = s.VpcName
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudRouteTable) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudRouteTable) BleveType() string {

	return "cloudroutetable"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudRouteTable) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *CloudRouteTable) Doc() string {

	return `Manages the list of route tables available in a cloud deployment.`
}

func (o *CloudRouteTable) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetRRN returns the RRN of the receiver.
func (o *CloudRouteTable) GetRRN() string {

	return o.RRN
}

// SetRRN sets the property RRN of the receiver using the given value.
func (o *CloudRouteTable) SetRRN(RRN string) {

	o.RRN = RRN
}

// GetAccountID returns the AccountID of the receiver.
func (o *CloudRouteTable) GetAccountID() string {

	return o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the given value.
func (o *CloudRouteTable) SetAccountID(accountID string) {

	o.AccountID = accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CloudRouteTable) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CloudRouteTable) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetApiID returns the ApiID of the receiver.
func (o *CloudRouteTable) GetApiID() int {

	return o.ApiID
}

// SetApiID sets the property ApiID of the receiver using the given value.
func (o *CloudRouteTable) SetApiID(apiID int) {

	o.ApiID = apiID
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CloudRouteTable) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CloudRouteTable) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *CloudRouteTable) GetCloudType() CloudRouteTableCloudTypeValue {

	return o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the given value.
func (o *CloudRouteTable) SetCloudType(cloudType CloudRouteTableCloudTypeValue) {

	o.CloudType = cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CloudRouteTable) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CloudRouteTable) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *CloudRouteTable) GetCustomerID() int {

	return o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the given value.
func (o *CloudRouteTable) SetCustomerID(customerID int) {

	o.CustomerID = customerID
}

// GetDescription returns the Description of the receiver.
func (o *CloudRouteTable) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *CloudRouteTable) SetDescription(description string) {

	o.Description = description
}

// GetInsertTS returns the InsertTS of the receiver.
func (o *CloudRouteTable) GetInsertTS() int {

	return o.InsertTS
}

// SetInsertTS sets the property InsertTS of the receiver using the given value.
func (o *CloudRouteTable) SetInsertTS(insertTS int) {

	o.InsertTS = insertTS
}

// GetMetadata returns the Metadata of the receiver.
func (o *CloudRouteTable) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *CloudRouteTable) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CloudRouteTable) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CloudRouteTable) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *CloudRouteTable) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *CloudRouteTable) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudRouteTable) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudRouteTable) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *CloudRouteTable) GetNativeID() string {

	return o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the given value.
func (o *CloudRouteTable) SetNativeID(nativeID string) {

	o.NativeID = nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CloudRouteTable) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CloudRouteTable) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *CloudRouteTable) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CloudRouteTable) SetProtected(protected bool) {

	o.Protected = protected
}

// GetRegionID returns the RegionID of the receiver.
func (o *CloudRouteTable) GetRegionID() string {

	return o.RegionID
}

// SetRegionID sets the property RegionID of the receiver using the given value.
func (o *CloudRouteTable) SetRegionID(regionID string) {

	o.RegionID = regionID
}

// GetRegionName returns the RegionName of the receiver.
func (o *CloudRouteTable) GetRegionName() string {

	return o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the given value.
func (o *CloudRouteTable) SetRegionName(regionName string) {

	o.RegionName = regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *CloudRouteTable) GetResourceID() int {

	return o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the given value.
func (o *CloudRouteTable) SetResourceID(resourceID int) {

	o.ResourceID = resourceID
}

// GetTags returns the Tags of the receiver.
func (o *CloudRouteTable) GetTags() map[string]string {

	return o.Tags
}

// SetTags sets the property Tags of the receiver using the given value.
func (o *CloudRouteTable) SetTags(tags map[string]string) {

	o.Tags = tags
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CloudRouteTable) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CloudRouteTable) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUrl returns the Url of the receiver.
func (o *CloudRouteTable) GetUrl() string {

	return o.Url
}

// SetUrl sets the property Url of the receiver using the given value.
func (o *CloudRouteTable) SetUrl(url string) {

	o.Url = url
}

// GetVpcID returns the VpcID of the receiver.
func (o *CloudRouteTable) GetVpcID() string {

	return o.VpcID
}

// SetVpcID sets the property VpcID of the receiver using the given value.
func (o *CloudRouteTable) SetVpcID(vpcID string) {

	o.VpcID = vpcID
}

// GetVpcName returns the VpcName of the receiver.
func (o *CloudRouteTable) GetVpcName() string {

	return o.VpcName
}

// SetVpcName sets the property VpcName of the receiver using the given value.
func (o *CloudRouteTable) SetVpcName(vpcName string) {

	o.VpcName = vpcName
}

// GetZHash returns the ZHash of the receiver.
func (o *CloudRouteTable) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CloudRouteTable) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CloudRouteTable) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CloudRouteTable) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudRouteTable) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudRouteTable{
			ID:                   &o.ID,
			RRN:                  &o.RRN,
			AccountID:            &o.AccountID,
			Annotations:          &o.Annotations,
			ApiID:                &o.ApiID,
			AssociatedTags:       &o.AssociatedTags,
			CloudType:            &o.CloudType,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CustomerID:           &o.CustomerID,
			Description:          &o.Description,
			GatewayID:            &o.GatewayID,
			InsertTS:             &o.InsertTS,
			MainTable:            &o.MainTable,
			Metadata:             &o.Metadata,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NativeID:             &o.NativeID,
			NormalizedTags:       &o.NormalizedTags,
			Protected:            &o.Protected,
			RegionID:             &o.RegionID,
			RegionName:           &o.RegionName,
			ResourceID:           &o.ResourceID,
			Routelist:            &o.Routelist,
			SubnetAssociations:   &o.SubnetAssociations,
			Tags:                 &o.Tags,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			Url:                  &o.Url,
			VpcID:                &o.VpcID,
			VpcName:              &o.VpcName,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseCloudRouteTable{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "RRN":
			sp.RRN = &(o.RRN)
		case "accountID":
			sp.AccountID = &(o.AccountID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "apiID":
			sp.ApiID = &(o.ApiID)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "cloudType":
			sp.CloudType = &(o.CloudType)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "customerID":
			sp.CustomerID = &(o.CustomerID)
		case "description":
			sp.Description = &(o.Description)
		case "gatewayID":
			sp.GatewayID = &(o.GatewayID)
		case "insertTS":
			sp.InsertTS = &(o.InsertTS)
		case "mainTable":
			sp.MainTable = &(o.MainTable)
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
		case "resourceID":
			sp.ResourceID = &(o.ResourceID)
		case "routelist":
			sp.Routelist = &(o.Routelist)
		case "subnetAssociations":
			sp.SubnetAssociations = &(o.SubnetAssociations)
		case "tags":
			sp.Tags = &(o.Tags)
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

// Patch apply the non nil value of a *SparseCloudRouteTable to the object.
func (o *CloudRouteTable) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudRouteTable)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.RRN != nil {
		o.RRN = *so.RRN
	}
	if so.AccountID != nil {
		o.AccountID = *so.AccountID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.ApiID != nil {
		o.ApiID = *so.ApiID
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CloudType != nil {
		o.CloudType = *so.CloudType
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CustomerID != nil {
		o.CustomerID = *so.CustomerID
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.GatewayID != nil {
		o.GatewayID = *so.GatewayID
	}
	if so.InsertTS != nil {
		o.InsertTS = *so.InsertTS
	}
	if so.MainTable != nil {
		o.MainTable = *so.MainTable
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
	if so.ResourceID != nil {
		o.ResourceID = *so.ResourceID
	}
	if so.Routelist != nil {
		o.Routelist = *so.Routelist
	}
	if so.SubnetAssociations != nil {
		o.SubnetAssociations = *so.SubnetAssociations
	}
	if so.Tags != nil {
		o.Tags = *so.Tags
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

// DeepCopy returns a deep copy if the CloudRouteTable.
func (o *CloudRouteTable) DeepCopy() *CloudRouteTable {

	if o == nil {
		return nil
	}

	out := &CloudRouteTable{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudRouteTable.
func (o *CloudRouteTable) DeepCopyInto(out *CloudRouteTable) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudRouteTable: %s", err))
	}

	*out = *target.(*CloudRouteTable)
}

// Validate valides the current information stored into the structure.
func (o *CloudRouteTable) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("cloudType", string(o.CloudType), []string{"AWS", "GCP", "AZURE", "ALIBABA"}, false); err != nil {
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

	for _, sub := range o.Routelist {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
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
func (*CloudRouteTable) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudRouteTableAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudRouteTableLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudRouteTable) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudRouteTableAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudRouteTable) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "RRN":
		return o.RRN
	case "accountID":
		return o.AccountID
	case "annotations":
		return o.Annotations
	case "apiID":
		return o.ApiID
	case "associatedTags":
		return o.AssociatedTags
	case "cloudType":
		return o.CloudType
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "customerID":
		return o.CustomerID
	case "description":
		return o.Description
	case "gatewayID":
		return o.GatewayID
	case "insertTS":
		return o.InsertTS
	case "mainTable":
		return o.MainTable
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
	case "resourceID":
		return o.ResourceID
	case "routelist":
		return o.Routelist
	case "subnetAssociations":
		return o.SubnetAssociations
	case "tags":
		return o.Tags
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

// CloudRouteTableAttributesMap represents the map of attribute for CloudRouteTable.
var CloudRouteTableAttributesMap = map[string]elemental.AttributeSpecification{
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
	"AccountID": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountid",
		ConvertedName:  "AccountID",
		Description:    `Cloud account ID associated with the entity (matches Prisma Cloud accountID).`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "accountID",
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
	"ApiID": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiid",
		ConvertedName:  "ApiID",
		Description:    `Prisma Cloud API ID (matches Prisma Cloud API ID).`,
		Exposed:        true,
		Getter:         true,
		Name:           "apiID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
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
	"CloudType": {
		AllowedChoices: []string{"AWS", "GCP", "AZURE", "ALIBABA"},
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
	"CustomerID": {
		AllowedChoices: []string{},
		BSONFieldName:  "customerid",
		ConvertedName:  "CustomerID",
		Description:    `Customer ID as identified by Prisma Cloud.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "customerID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
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
	"GatewayID": {
		AllowedChoices: []string{},
		BSONFieldName:  "gatewayid",
		ConvertedName:  "GatewayID",
		Description:    `The gateway that this route table is associated with.`,
		Exposed:        true,
		Name:           "gatewayID",
		Stored:         true,
		Type:           "string",
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
	"MainTable": {
		AllowedChoices: []string{},
		BSONFieldName:  "maintable",
		ConvertedName:  "MainTable",
		Description:    `Indicates that this is the default route table for the VPC.`,
		Exposed:        true,
		Name:           "mainTable",
		Stored:         true,
		Type:           "boolean",
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
	"ResourceID": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourceid",
		ConvertedName:  "ResourceID",
		Description:    `Prisma Cloud Resource ID.`,
		Exposed:        true,
		Getter:         true,
		Name:           "resourceID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"Routelist": {
		AllowedChoices: []string{},
		BSONFieldName:  "routelist",
		ConvertedName:  "Routelist",
		Description:    `Routes associated with this route table.`,
		Exposed:        true,
		Name:           "routelist",
		Stored:         true,
		SubType:        "cloudroute",
		Type:           "refList",
	},
	"SubnetAssociations": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnetassociations",
		ConvertedName:  "SubnetAssociations",
		Description:    `The list of subnets that this route table is associated with.`,
		Exposed:        true,
		Name:           "subnetAssociations",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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

// CloudRouteTableLowerCaseAttributesMap represents the map of attribute for CloudRouteTable.
var CloudRouteTableLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"accountid": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountid",
		ConvertedName:  "AccountID",
		Description:    `Cloud account ID associated with the entity (matches Prisma Cloud accountID).`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "accountID",
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
	"apiid": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiid",
		ConvertedName:  "ApiID",
		Description:    `Prisma Cloud API ID (matches Prisma Cloud API ID).`,
		Exposed:        true,
		Getter:         true,
		Name:           "apiID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
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
	"cloudtype": {
		AllowedChoices: []string{"AWS", "GCP", "AZURE", "ALIBABA"},
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
	"customerid": {
		AllowedChoices: []string{},
		BSONFieldName:  "customerid",
		ConvertedName:  "CustomerID",
		Description:    `Customer ID as identified by Prisma Cloud.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "customerID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
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
	"gatewayid": {
		AllowedChoices: []string{},
		BSONFieldName:  "gatewayid",
		ConvertedName:  "GatewayID",
		Description:    `The gateway that this route table is associated with.`,
		Exposed:        true,
		Name:           "gatewayID",
		Stored:         true,
		Type:           "string",
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
	"maintable": {
		AllowedChoices: []string{},
		BSONFieldName:  "maintable",
		ConvertedName:  "MainTable",
		Description:    `Indicates that this is the default route table for the VPC.`,
		Exposed:        true,
		Name:           "mainTable",
		Stored:         true,
		Type:           "boolean",
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
	"resourceid": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourceid",
		ConvertedName:  "ResourceID",
		Description:    `Prisma Cloud Resource ID.`,
		Exposed:        true,
		Getter:         true,
		Name:           "resourceID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"routelist": {
		AllowedChoices: []string{},
		BSONFieldName:  "routelist",
		ConvertedName:  "Routelist",
		Description:    `Routes associated with this route table.`,
		Exposed:        true,
		Name:           "routelist",
		Stored:         true,
		SubType:        "cloudroute",
		Type:           "refList",
	},
	"subnetassociations": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnetassociations",
		ConvertedName:  "SubnetAssociations",
		Description:    `The list of subnets that this route table is associated with.`,
		Exposed:        true,
		Name:           "subnetAssociations",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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

// SparseCloudRouteTablesList represents a list of SparseCloudRouteTables
type SparseCloudRouteTablesList []*SparseCloudRouteTable

// Identity returns the identity of the objects in the list.
func (o SparseCloudRouteTablesList) Identity() elemental.Identity {

	return CloudRouteTableIdentity
}

// Copy returns a pointer to a copy the SparseCloudRouteTablesList.
func (o SparseCloudRouteTablesList) Copy() elemental.Identifiables {

	copy := append(SparseCloudRouteTablesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudRouteTablesList.
func (o SparseCloudRouteTablesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudRouteTablesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudRouteTable))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudRouteTablesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudRouteTablesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseCloudRouteTablesList converted to CloudRouteTablesList.
func (o SparseCloudRouteTablesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudRouteTablesList) Version() int {

	return 1
}

// SparseCloudRouteTable represents the sparse version of a cloudroutetable.
type SparseCloudRouteTable struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Restricted Resource Name.
	RRN *string `json:"rrn,omitempty" msgpack:"rrn,omitempty" bson:"rrn,omitempty" mapstructure:"rrn,omitempty"`

	// Cloud account ID associated with the entity (matches Prisma Cloud accountID).
	AccountID *string `json:"accountId,omitempty" msgpack:"accountId,omitempty" bson:"accountid,omitempty" mapstructure:"accountId,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// Prisma Cloud API ID (matches Prisma Cloud API ID).
	ApiID *int `json:"apiID,omitempty" msgpack:"apiID,omitempty" bson:"apiid,omitempty" mapstructure:"apiID,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// Cloud type of the entity.
	CloudType *CloudRouteTableCloudTypeValue `json:"cloudType,omitempty" msgpack:"cloudType,omitempty" bson:"cloudtype,omitempty" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Customer ID as identified by Prisma Cloud.
	CustomerID *int `json:"customerID,omitempty" msgpack:"customerID,omitempty" bson:"customerid,omitempty" mapstructure:"customerID,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// The gateway that this route table is associated with.
	GatewayID *string `json:"gatewayID,omitempty" msgpack:"gatewayID,omitempty" bson:"gatewayid,omitempty" mapstructure:"gatewayID,omitempty"`

	// Timestamp of when object was created.
	InsertTS *int `json:"insertTs,omitempty" msgpack:"insertTs,omitempty" bson:"insertts,omitempty" mapstructure:"insertTs,omitempty"`

	// Indicates that this is the default route table for the VPC.
	MainTable *bool `json:"mainTable,omitempty" msgpack:"mainTable,omitempty" bson:"maintable,omitempty" mapstructure:"mainTable,omitempty"`

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

	// Prisma Cloud Resource ID.
	ResourceID *int `json:"resourceID,omitempty" msgpack:"resourceID,omitempty" bson:"resourceid,omitempty" mapstructure:"resourceID,omitempty"`

	// Routes associated with this route table.
	Routelist *CloudRoutesList `json:"routelist,omitempty" msgpack:"routelist,omitempty" bson:"routelist,omitempty" mapstructure:"routelist,omitempty"`

	// The list of subnets that this route table is associated with.
	SubnetAssociations *[]string `json:"subnetAssociations,omitempty" msgpack:"subnetAssociations,omitempty" bson:"subnetassociations,omitempty" mapstructure:"subnetAssociations,omitempty"`

	// Internal representation of object tags.
	Tags *map[string]string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"tags,omitempty" mapstructure:"tags,omitempty"`

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

// NewSparseCloudRouteTable returns a new  SparseCloudRouteTable.
func NewSparseCloudRouteTable() *SparseCloudRouteTable {
	return &SparseCloudRouteTable{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudRouteTable) Identity() elemental.Identity {

	return CloudRouteTableIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudRouteTable) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudRouteTable) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudRouteTable) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudRouteTable{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.RRN != nil {
		s.RRN = o.RRN
	}
	if o.AccountID != nil {
		s.AccountID = o.AccountID
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.ApiID != nil {
		s.ApiID = o.ApiID
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CloudType != nil {
		s.CloudType = o.CloudType
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CustomerID != nil {
		s.CustomerID = o.CustomerID
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.GatewayID != nil {
		s.GatewayID = o.GatewayID
	}
	if o.InsertTS != nil {
		s.InsertTS = o.InsertTS
	}
	if o.MainTable != nil {
		s.MainTable = o.MainTable
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
	if o.ResourceID != nil {
		s.ResourceID = o.ResourceID
	}
	if o.Routelist != nil {
		s.Routelist = o.Routelist
	}
	if o.SubnetAssociations != nil {
		s.SubnetAssociations = o.SubnetAssociations
	}
	if o.Tags != nil {
		s.Tags = o.Tags
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
func (o *SparseCloudRouteTable) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudRouteTable{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.RRN != nil {
		o.RRN = s.RRN
	}
	if s.AccountID != nil {
		o.AccountID = s.AccountID
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.ApiID != nil {
		o.ApiID = s.ApiID
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CloudType != nil {
		o.CloudType = s.CloudType
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CustomerID != nil {
		o.CustomerID = s.CustomerID
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.GatewayID != nil {
		o.GatewayID = s.GatewayID
	}
	if s.InsertTS != nil {
		o.InsertTS = s.InsertTS
	}
	if s.MainTable != nil {
		o.MainTable = s.MainTable
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
	if s.ResourceID != nil {
		o.ResourceID = s.ResourceID
	}
	if s.Routelist != nil {
		o.Routelist = s.Routelist
	}
	if s.SubnetAssociations != nil {
		o.SubnetAssociations = s.SubnetAssociations
	}
	if s.Tags != nil {
		o.Tags = s.Tags
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
func (o *SparseCloudRouteTable) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudRouteTable) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudRouteTable()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.RRN != nil {
		out.RRN = *o.RRN
	}
	if o.AccountID != nil {
		out.AccountID = *o.AccountID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.ApiID != nil {
		out.ApiID = *o.ApiID
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CloudType != nil {
		out.CloudType = *o.CloudType
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CustomerID != nil {
		out.CustomerID = *o.CustomerID
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.GatewayID != nil {
		out.GatewayID = *o.GatewayID
	}
	if o.InsertTS != nil {
		out.InsertTS = *o.InsertTS
	}
	if o.MainTable != nil {
		out.MainTable = *o.MainTable
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
	if o.ResourceID != nil {
		out.ResourceID = *o.ResourceID
	}
	if o.Routelist != nil {
		out.Routelist = *o.Routelist
	}
	if o.SubnetAssociations != nil {
		out.SubnetAssociations = *o.SubnetAssociations
	}
	if o.Tags != nil {
		out.Tags = *o.Tags
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
func (o *SparseCloudRouteTable) GetRRN() (out string) {

	if o.RRN == nil {
		return
	}

	return *o.RRN
}

// SetRRN sets the property RRN of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetRRN(RRN string) {

	o.RRN = &RRN
}

// GetAccountID returns the AccountID of the receiver.
func (o *SparseCloudRouteTable) GetAccountID() (out string) {

	if o.AccountID == nil {
		return
	}

	return *o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetAccountID(accountID string) {

	o.AccountID = &accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseCloudRouteTable) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetApiID returns the ApiID of the receiver.
func (o *SparseCloudRouteTable) GetApiID() (out int) {

	if o.ApiID == nil {
		return
	}

	return *o.ApiID
}

// SetApiID sets the property ApiID of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetApiID(apiID int) {

	o.ApiID = &apiID
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCloudRouteTable) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *SparseCloudRouteTable) GetCloudType() (out CloudRouteTableCloudTypeValue) {

	if o.CloudType == nil {
		return
	}

	return *o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetCloudType(cloudType CloudRouteTableCloudTypeValue) {

	o.CloudType = &cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCloudRouteTable) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *SparseCloudRouteTable) GetCustomerID() (out int) {

	if o.CustomerID == nil {
		return
	}

	return *o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetCustomerID(customerID int) {

	o.CustomerID = &customerID
}

// GetDescription returns the Description of the receiver.
func (o *SparseCloudRouteTable) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetDescription(description string) {

	o.Description = &description
}

// GetInsertTS returns the InsertTS of the receiver.
func (o *SparseCloudRouteTable) GetInsertTS() (out int) {

	if o.InsertTS == nil {
		return
	}

	return *o.InsertTS
}

// SetInsertTS sets the property InsertTS of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetInsertTS(insertTS int) {

	o.InsertTS = &insertTS
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseCloudRouteTable) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCloudRouteTable) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseCloudRouteTable) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudRouteTable) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *SparseCloudRouteTable) GetNativeID() (out string) {

	if o.NativeID == nil {
		return
	}

	return *o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetNativeID(nativeID string) {

	o.NativeID = &nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCloudRouteTable) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCloudRouteTable) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetRegionID returns the RegionID of the receiver.
func (o *SparseCloudRouteTable) GetRegionID() (out string) {

	if o.RegionID == nil {
		return
	}

	return *o.RegionID
}

// SetRegionID sets the property RegionID of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetRegionID(regionID string) {

	o.RegionID = &regionID
}

// GetRegionName returns the RegionName of the receiver.
func (o *SparseCloudRouteTable) GetRegionName() (out string) {

	if o.RegionName == nil {
		return
	}

	return *o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetRegionName(regionName string) {

	o.RegionName = &regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *SparseCloudRouteTable) GetResourceID() (out int) {

	if o.ResourceID == nil {
		return
	}

	return *o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetResourceID(resourceID int) {

	o.ResourceID = &resourceID
}

// GetTags returns the Tags of the receiver.
func (o *SparseCloudRouteTable) GetTags() (out map[string]string) {

	if o.Tags == nil {
		return
	}

	return *o.Tags
}

// SetTags sets the property Tags of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetTags(tags map[string]string) {

	o.Tags = &tags
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCloudRouteTable) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUrl returns the Url of the receiver.
func (o *SparseCloudRouteTable) GetUrl() (out string) {

	if o.Url == nil {
		return
	}

	return *o.Url
}

// SetUrl sets the property Url of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetUrl(url string) {

	o.Url = &url
}

// GetVpcID returns the VpcID of the receiver.
func (o *SparseCloudRouteTable) GetVpcID() (out string) {

	if o.VpcID == nil {
		return
	}

	return *o.VpcID
}

// SetVpcID sets the property VpcID of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetVpcID(vpcID string) {

	o.VpcID = &vpcID
}

// GetVpcName returns the VpcName of the receiver.
func (o *SparseCloudRouteTable) GetVpcName() (out string) {

	if o.VpcName == nil {
		return
	}

	return *o.VpcName
}

// SetVpcName sets the property VpcName of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetVpcName(vpcName string) {

	o.VpcName = &vpcName
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCloudRouteTable) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCloudRouteTable) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCloudRouteTable) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCloudRouteTable.
func (o *SparseCloudRouteTable) DeepCopy() *SparseCloudRouteTable {

	if o == nil {
		return nil
	}

	out := &SparseCloudRouteTable{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudRouteTable.
func (o *SparseCloudRouteTable) DeepCopyInto(out *SparseCloudRouteTable) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudRouteTable: %s", err))
	}

	*out = *target.(*SparseCloudRouteTable)
}

type mongoAttributesCloudRouteTable struct {
	ID                   bson.ObjectId                 `bson:"_id,omitempty"`
	RRN                  string                        `bson:"rrn"`
	AccountID            string                        `bson:"accountid"`
	Annotations          map[string][]string           `bson:"annotations"`
	ApiID                int                           `bson:"apiid"`
	AssociatedTags       []string                      `bson:"associatedtags"`
	CloudType            CloudRouteTableCloudTypeValue `bson:"cloudtype"`
	CreateIdempotencyKey string                        `bson:"createidempotencykey"`
	CustomerID           int                           `bson:"customerid"`
	Description          string                        `bson:"description"`
	GatewayID            string                        `bson:"gatewayid"`
	InsertTS             int                           `bson:"insertts,omitempty"`
	MainTable            bool                          `bson:"maintable"`
	Metadata             []string                      `bson:"metadata"`
	MigrationsLog        map[string]string             `bson:"migrationslog,omitempty"`
	Name                 string                        `bson:"name"`
	Namespace            string                        `bson:"namespace"`
	NativeID             string                        `bson:"nativeid"`
	NormalizedTags       []string                      `bson:"normalizedtags"`
	Protected            bool                          `bson:"protected"`
	RegionID             string                        `bson:"regionid"`
	RegionName           string                        `bson:"regionname"`
	ResourceID           int                           `bson:"resourceid"`
	Routelist            CloudRoutesList               `bson:"routelist"`
	SubnetAssociations   []string                      `bson:"subnetassociations"`
	Tags                 map[string]string             `bson:"tags"`
	UpdateIdempotencyKey string                        `bson:"updateidempotencykey"`
	Url                  string                        `bson:"url"`
	VpcID                string                        `bson:"vpcid"`
	VpcName              string                        `bson:"vpcname"`
	ZHash                int                           `bson:"zhash"`
	Zone                 int                           `bson:"zone"`
}
type mongoAttributesSparseCloudRouteTable struct {
	ID                   bson.ObjectId                  `bson:"_id,omitempty"`
	RRN                  *string                        `bson:"rrn,omitempty"`
	AccountID            *string                        `bson:"accountid,omitempty"`
	Annotations          *map[string][]string           `bson:"annotations,omitempty"`
	ApiID                *int                           `bson:"apiid,omitempty"`
	AssociatedTags       *[]string                      `bson:"associatedtags,omitempty"`
	CloudType            *CloudRouteTableCloudTypeValue `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey *string                        `bson:"createidempotencykey,omitempty"`
	CustomerID           *int                           `bson:"customerid,omitempty"`
	Description          *string                        `bson:"description,omitempty"`
	GatewayID            *string                        `bson:"gatewayid,omitempty"`
	InsertTS             *int                           `bson:"insertts,omitempty"`
	MainTable            *bool                          `bson:"maintable,omitempty"`
	Metadata             *[]string                      `bson:"metadata,omitempty"`
	MigrationsLog        *map[string]string             `bson:"migrationslog,omitempty"`
	Name                 *string                        `bson:"name,omitempty"`
	Namespace            *string                        `bson:"namespace,omitempty"`
	NativeID             *string                        `bson:"nativeid,omitempty"`
	NormalizedTags       *[]string                      `bson:"normalizedtags,omitempty"`
	Protected            *bool                          `bson:"protected,omitempty"`
	RegionID             *string                        `bson:"regionid,omitempty"`
	RegionName           *string                        `bson:"regionname,omitempty"`
	ResourceID           *int                           `bson:"resourceid,omitempty"`
	Routelist            *CloudRoutesList               `bson:"routelist,omitempty"`
	SubnetAssociations   *[]string                      `bson:"subnetassociations,omitempty"`
	Tags                 *map[string]string             `bson:"tags,omitempty"`
	UpdateIdempotencyKey *string                        `bson:"updateidempotencykey,omitempty"`
	Url                  *string                        `bson:"url,omitempty"`
	VpcID                *string                        `bson:"vpcid,omitempty"`
	VpcName              *string                        `bson:"vpcname,omitempty"`
	ZHash                *int                           `bson:"zhash,omitempty"`
	Zone                 *int                           `bson:"zone,omitempty"`
}
