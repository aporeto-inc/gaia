package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudVPCCloudTypeValue represents the possible values for attribute "cloudType".
type CloudVPCCloudTypeValue string

const (
	// CloudVPCCloudTypeAWS represents the value AWS.
	CloudVPCCloudTypeAWS CloudVPCCloudTypeValue = "AWS"

	// CloudVPCCloudTypeGCP represents the value GCP.
	CloudVPCCloudTypeGCP CloudVPCCloudTypeValue = "GCP"
)

// CloudVPCIdentity represents the Identity of the object.
var CloudVPCIdentity = elemental.Identity{
	Name:     "cloudvpc",
	Category: "cloudvpcs",
	Package:  "yeul",
	Private:  false,
}

// CloudVPCsList represents a list of CloudVPCs
type CloudVPCsList []*CloudVPC

// Identity returns the identity of the objects in the list.
func (o CloudVPCsList) Identity() elemental.Identity {

	return CloudVPCIdentity
}

// Copy returns a pointer to a copy the CloudVPCsList.
func (o CloudVPCsList) Copy() elemental.Identifiables {

	copy := append(CloudVPCsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudVPCsList.
func (o CloudVPCsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudVPCsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudVPC))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudVPCsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudVPCsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the CloudVPCsList converted to SparseCloudVPCsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudVPCsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudVPCsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudVPC)
	}

	return out
}

// Version returns the version of the content.
func (o CloudVPCsList) Version() int {

	return 1
}

// CloudVPC represents the model of a cloudvpc
type CloudVPC struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Restricted Resource Name.
	RRN string `json:"rrn" msgpack:"rrn" bson:"rrn" mapstructure:"rrn,omitempty"`

	// Cloud account ID associated with the entity.
	AccountID string `json:"accountId" msgpack:"accountId" bson:"accountid" mapstructure:"accountId,omitempty"`

	// Cloud account name associated with the entity.
	AccountName string `json:"accountName" msgpack:"accountName" bson:"accountname" mapstructure:"accountName,omitempty"`

	// The IP CIDR associated with the VPC.
	Address string `json:"address" msgpack:"address" bson:"address" mapstructure:"address,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Cloud type of the entity.
	CloudType CloudVPCCloudTypeValue `json:"cloudType" msgpack:"cloudType" bson:"cloudtype" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

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

// NewCloudVPC returns a new *CloudVPC
func NewCloudVPC() *CloudVPC {

	return &CloudVPC{
		ModelVersion:   1,
		AssociatedTags: []string{},
		Annotations:    map[string][]string{},
		Metadata:       []string{},
		Tags:           map[string]string{},
		NormalizedTags: []string{},
		MigrationsLog:  map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudVPC) Identity() elemental.Identity {

	return CloudVPCIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudVPC) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudVPC) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudVPC) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudVPC{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.RRN = o.RRN
	s.AccountID = o.AccountID
	s.AccountName = o.AccountName
	s.Address = o.Address
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CloudType = o.CloudType
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.Description = o.Description
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
func (o *CloudVPC) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudVPC{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.RRN = s.RRN
	o.AccountID = s.AccountID
	o.AccountName = s.AccountName
	o.Address = s.Address
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CloudType = s.CloudType
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.Description = s.Description
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
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.Url = s.Url
	o.VpcID = s.VpcID
	o.VpcName = s.VpcName
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudVPC) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudVPC) BleveType() string {

	return "cloudvpc"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudVPC) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *CloudVPC) Doc() string {

	return `A CloudVPC represents a VPC as defined in an cloud provider (AWS/Azure/GCP etc).
The VPC is essentially an L3 routing domain with at least one subnet attached
and it defines an isolated network.`
}

func (o *CloudVPC) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetRRN returns the RRN of the receiver.
func (o *CloudVPC) GetRRN() string {

	return o.RRN
}

// SetRRN sets the property RRN of the receiver using the given value.
func (o *CloudVPC) SetRRN(RRN string) {

	o.RRN = RRN
}

// GetAccountID returns the AccountID of the receiver.
func (o *CloudVPC) GetAccountID() string {

	return o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the given value.
func (o *CloudVPC) SetAccountID(accountID string) {

	o.AccountID = accountID
}

// GetAccountName returns the AccountName of the receiver.
func (o *CloudVPC) GetAccountName() string {

	return o.AccountName
}

// SetAccountName sets the property AccountName of the receiver using the given value.
func (o *CloudVPC) SetAccountName(accountName string) {

	o.AccountName = accountName
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CloudVPC) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CloudVPC) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CloudVPC) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CloudVPC) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *CloudVPC) GetCloudType() CloudVPCCloudTypeValue {

	return o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the given value.
func (o *CloudVPC) SetCloudType(cloudType CloudVPCCloudTypeValue) {

	o.CloudType = cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CloudVPC) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CloudVPC) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetDescription returns the Description of the receiver.
func (o *CloudVPC) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *CloudVPC) SetDescription(description string) {

	o.Description = description
}

// GetInsertTS returns the InsertTS of the receiver.
func (o *CloudVPC) GetInsertTS() int {

	return o.InsertTS
}

// SetInsertTS sets the property InsertTS of the receiver using the given value.
func (o *CloudVPC) SetInsertTS(insertTS int) {

	o.InsertTS = insertTS
}

// GetMetadata returns the Metadata of the receiver.
func (o *CloudVPC) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *CloudVPC) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CloudVPC) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CloudVPC) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *CloudVPC) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *CloudVPC) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudVPC) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudVPC) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *CloudVPC) GetNativeID() string {

	return o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the given value.
func (o *CloudVPC) SetNativeID(nativeID string) {

	o.NativeID = nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CloudVPC) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CloudVPC) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *CloudVPC) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CloudVPC) SetProtected(protected bool) {

	o.Protected = protected
}

// GetRegionID returns the RegionID of the receiver.
func (o *CloudVPC) GetRegionID() string {

	return o.RegionID
}

// SetRegionID sets the property RegionID of the receiver using the given value.
func (o *CloudVPC) SetRegionID(regionID string) {

	o.RegionID = regionID
}

// GetRegionName returns the RegionName of the receiver.
func (o *CloudVPC) GetRegionName() string {

	return o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the given value.
func (o *CloudVPC) SetRegionName(regionName string) {

	o.RegionName = regionName
}

// GetTags returns the Tags of the receiver.
func (o *CloudVPC) GetTags() map[string]string {

	return o.Tags
}

// SetTags sets the property Tags of the receiver using the given value.
func (o *CloudVPC) SetTags(tags map[string]string) {

	o.Tags = tags
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CloudVPC) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CloudVPC) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUrl returns the Url of the receiver.
func (o *CloudVPC) GetUrl() string {

	return o.Url
}

// SetUrl sets the property Url of the receiver using the given value.
func (o *CloudVPC) SetUrl(url string) {

	o.Url = url
}

// GetVpcID returns the VpcID of the receiver.
func (o *CloudVPC) GetVpcID() string {

	return o.VpcID
}

// SetVpcID sets the property VpcID of the receiver using the given value.
func (o *CloudVPC) SetVpcID(vpcID string) {

	o.VpcID = vpcID
}

// GetVpcName returns the VpcName of the receiver.
func (o *CloudVPC) GetVpcName() string {

	return o.VpcName
}

// SetVpcName sets the property VpcName of the receiver using the given value.
func (o *CloudVPC) SetVpcName(vpcName string) {

	o.VpcName = vpcName
}

// GetZHash returns the ZHash of the receiver.
func (o *CloudVPC) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CloudVPC) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CloudVPC) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CloudVPC) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudVPC) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudVPC{
			ID:                   &o.ID,
			RRN:                  &o.RRN,
			AccountID:            &o.AccountID,
			AccountName:          &o.AccountName,
			Address:              &o.Address,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CloudType:            &o.CloudType,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			Description:          &o.Description,
			InsertTS:             &o.InsertTS,
			Metadata:             &o.Metadata,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NativeID:             &o.NativeID,
			NormalizedTags:       &o.NormalizedTags,
			Protected:            &o.Protected,
			RegionID:             &o.RegionID,
			RegionName:           &o.RegionName,
			Tags:                 &o.Tags,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			Url:                  &o.Url,
			VpcID:                &o.VpcID,
			VpcName:              &o.VpcName,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseCloudVPC{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "RRN":
			sp.RRN = &(o.RRN)
		case "accountID":
			sp.AccountID = &(o.AccountID)
		case "accountName":
			sp.AccountName = &(o.AccountName)
		case "address":
			sp.Address = &(o.Address)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "cloudType":
			sp.CloudType = &(o.CloudType)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "description":
			sp.Description = &(o.Description)
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

// Patch apply the non nil value of a *SparseCloudVPC to the object.
func (o *CloudVPC) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudVPC)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.RRN != nil {
		o.RRN = *so.RRN
	}
	if so.AccountID != nil {
		o.AccountID = *so.AccountID
	}
	if so.AccountName != nil {
		o.AccountName = *so.AccountName
	}
	if so.Address != nil {
		o.Address = *so.Address
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
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
	if so.Description != nil {
		o.Description = *so.Description
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

// DeepCopy returns a deep copy if the CloudVPC.
func (o *CloudVPC) DeepCopy() *CloudVPC {

	if o == nil {
		return nil
	}

	out := &CloudVPC{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudVPC.
func (o *CloudVPC) DeepCopyInto(out *CloudVPC) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudVPC: %s", err))
	}

	*out = *target.(*CloudVPC)
}

// Validate valides the current information stored into the structure.
func (o *CloudVPC) Validate() error {

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
func (*CloudVPC) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudVPCAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudVPCLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudVPC) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudVPCAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudVPC) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "RRN":
		return o.RRN
	case "accountID":
		return o.AccountID
	case "accountName":
		return o.AccountName
	case "address":
		return o.Address
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "cloudType":
		return o.CloudType
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "description":
		return o.Description
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

// CloudVPCAttributesMap represents the map of attribute for CloudVPC.
var CloudVPCAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Address": {
		AllowedChoices: []string{},
		BSONFieldName:  "address",
		ConvertedName:  "Address",
		Description:    `The IP CIDR associated with the VPC.`,
		Exposed:        true,
		Name:           "address",
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

// CloudVPCLowerCaseAttributesMap represents the map of attribute for CloudVPC.
var CloudVPCLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"address": {
		AllowedChoices: []string{},
		BSONFieldName:  "address",
		ConvertedName:  "Address",
		Description:    `The IP CIDR associated with the VPC.`,
		Exposed:        true,
		Name:           "address",
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

// SparseCloudVPCsList represents a list of SparseCloudVPCs
type SparseCloudVPCsList []*SparseCloudVPC

// Identity returns the identity of the objects in the list.
func (o SparseCloudVPCsList) Identity() elemental.Identity {

	return CloudVPCIdentity
}

// Copy returns a pointer to a copy the SparseCloudVPCsList.
func (o SparseCloudVPCsList) Copy() elemental.Identifiables {

	copy := append(SparseCloudVPCsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudVPCsList.
func (o SparseCloudVPCsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudVPCsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudVPC))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudVPCsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudVPCsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseCloudVPCsList converted to CloudVPCsList.
func (o SparseCloudVPCsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudVPCsList) Version() int {

	return 1
}

// SparseCloudVPC represents the sparse version of a cloudvpc.
type SparseCloudVPC struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Restricted Resource Name.
	RRN *string `json:"rrn,omitempty" msgpack:"rrn,omitempty" bson:"rrn,omitempty" mapstructure:"rrn,omitempty"`

	// Cloud account ID associated with the entity.
	AccountID *string `json:"accountId,omitempty" msgpack:"accountId,omitempty" bson:"accountid,omitempty" mapstructure:"accountId,omitempty"`

	// Cloud account name associated with the entity.
	AccountName *string `json:"accountName,omitempty" msgpack:"accountName,omitempty" bson:"accountname,omitempty" mapstructure:"accountName,omitempty"`

	// The IP CIDR associated with the VPC.
	Address *string `json:"address,omitempty" msgpack:"address,omitempty" bson:"address,omitempty" mapstructure:"address,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// Cloud type of the entity.
	CloudType *CloudVPCCloudTypeValue `json:"cloudType,omitempty" msgpack:"cloudType,omitempty" bson:"cloudtype,omitempty" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

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

// NewSparseCloudVPC returns a new  SparseCloudVPC.
func NewSparseCloudVPC() *SparseCloudVPC {
	return &SparseCloudVPC{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudVPC) Identity() elemental.Identity {

	return CloudVPCIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudVPC) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudVPC) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudVPC) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudVPC{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.RRN != nil {
		s.RRN = o.RRN
	}
	if o.AccountID != nil {
		s.AccountID = o.AccountID
	}
	if o.AccountName != nil {
		s.AccountName = o.AccountName
	}
	if o.Address != nil {
		s.Address = o.Address
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
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
	if o.Description != nil {
		s.Description = o.Description
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
func (o *SparseCloudVPC) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudVPC{}
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
	if s.AccountName != nil {
		o.AccountName = s.AccountName
	}
	if s.Address != nil {
		o.Address = s.Address
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
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
	if s.Description != nil {
		o.Description = s.Description
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
func (o *SparseCloudVPC) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudVPC) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudVPC()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.RRN != nil {
		out.RRN = *o.RRN
	}
	if o.AccountID != nil {
		out.AccountID = *o.AccountID
	}
	if o.AccountName != nil {
		out.AccountName = *o.AccountName
	}
	if o.Address != nil {
		out.Address = *o.Address
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
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
	if o.Description != nil {
		out.Description = *o.Description
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
func (o *SparseCloudVPC) GetRRN() (out string) {

	if o.RRN == nil {
		return
	}

	return *o.RRN
}

// SetRRN sets the property RRN of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetRRN(RRN string) {

	o.RRN = &RRN
}

// GetAccountID returns the AccountID of the receiver.
func (o *SparseCloudVPC) GetAccountID() (out string) {

	if o.AccountID == nil {
		return
	}

	return *o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetAccountID(accountID string) {

	o.AccountID = &accountID
}

// GetAccountName returns the AccountName of the receiver.
func (o *SparseCloudVPC) GetAccountName() (out string) {

	if o.AccountName == nil {
		return
	}

	return *o.AccountName
}

// SetAccountName sets the property AccountName of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetAccountName(accountName string) {

	o.AccountName = &accountName
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseCloudVPC) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCloudVPC) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *SparseCloudVPC) GetCloudType() (out CloudVPCCloudTypeValue) {

	if o.CloudType == nil {
		return
	}

	return *o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetCloudType(cloudType CloudVPCCloudTypeValue) {

	o.CloudType = &cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCloudVPC) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetDescription returns the Description of the receiver.
func (o *SparseCloudVPC) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetDescription(description string) {

	o.Description = &description
}

// GetInsertTS returns the InsertTS of the receiver.
func (o *SparseCloudVPC) GetInsertTS() (out int) {

	if o.InsertTS == nil {
		return
	}

	return *o.InsertTS
}

// SetInsertTS sets the property InsertTS of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetInsertTS(insertTS int) {

	o.InsertTS = &insertTS
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseCloudVPC) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCloudVPC) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseCloudVPC) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudVPC) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *SparseCloudVPC) GetNativeID() (out string) {

	if o.NativeID == nil {
		return
	}

	return *o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetNativeID(nativeID string) {

	o.NativeID = &nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCloudVPC) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCloudVPC) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetRegionID returns the RegionID of the receiver.
func (o *SparseCloudVPC) GetRegionID() (out string) {

	if o.RegionID == nil {
		return
	}

	return *o.RegionID
}

// SetRegionID sets the property RegionID of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetRegionID(regionID string) {

	o.RegionID = &regionID
}

// GetRegionName returns the RegionName of the receiver.
func (o *SparseCloudVPC) GetRegionName() (out string) {

	if o.RegionName == nil {
		return
	}

	return *o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetRegionName(regionName string) {

	o.RegionName = &regionName
}

// GetTags returns the Tags of the receiver.
func (o *SparseCloudVPC) GetTags() (out map[string]string) {

	if o.Tags == nil {
		return
	}

	return *o.Tags
}

// SetTags sets the property Tags of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetTags(tags map[string]string) {

	o.Tags = &tags
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCloudVPC) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUrl returns the Url of the receiver.
func (o *SparseCloudVPC) GetUrl() (out string) {

	if o.Url == nil {
		return
	}

	return *o.Url
}

// SetUrl sets the property Url of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetUrl(url string) {

	o.Url = &url
}

// GetVpcID returns the VpcID of the receiver.
func (o *SparseCloudVPC) GetVpcID() (out string) {

	if o.VpcID == nil {
		return
	}

	return *o.VpcID
}

// SetVpcID sets the property VpcID of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetVpcID(vpcID string) {

	o.VpcID = &vpcID
}

// GetVpcName returns the VpcName of the receiver.
func (o *SparseCloudVPC) GetVpcName() (out string) {

	if o.VpcName == nil {
		return
	}

	return *o.VpcName
}

// SetVpcName sets the property VpcName of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetVpcName(vpcName string) {

	o.VpcName = &vpcName
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCloudVPC) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCloudVPC) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCloudVPC) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCloudVPC.
func (o *SparseCloudVPC) DeepCopy() *SparseCloudVPC {

	if o == nil {
		return nil
	}

	out := &SparseCloudVPC{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudVPC.
func (o *SparseCloudVPC) DeepCopyInto(out *SparseCloudVPC) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudVPC: %s", err))
	}

	*out = *target.(*SparseCloudVPC)
}

type mongoAttributesCloudVPC struct {
	ID                   bson.ObjectId          `bson:"_id,omitempty"`
	RRN                  string                 `bson:"rrn"`
	AccountID            string                 `bson:"accountid"`
	AccountName          string                 `bson:"accountname"`
	Address              string                 `bson:"address"`
	Annotations          map[string][]string    `bson:"annotations"`
	AssociatedTags       []string               `bson:"associatedtags"`
	CloudType            CloudVPCCloudTypeValue `bson:"cloudtype"`
	CreateIdempotencyKey string                 `bson:"createidempotencykey"`
	Description          string                 `bson:"description"`
	InsertTS             int                    `bson:"insertts,omitempty"`
	Metadata             []string               `bson:"metadata"`
	MigrationsLog        map[string]string      `bson:"migrationslog,omitempty"`
	Name                 string                 `bson:"name"`
	Namespace            string                 `bson:"namespace"`
	NativeID             string                 `bson:"nativeid"`
	NormalizedTags       []string               `bson:"normalizedtags"`
	Protected            bool                   `bson:"protected"`
	RegionID             string                 `bson:"regionid"`
	RegionName           string                 `bson:"regionname"`
	Tags                 map[string]string      `bson:"tags"`
	UpdateIdempotencyKey string                 `bson:"updateidempotencykey"`
	Url                  string                 `bson:"url"`
	VpcID                string                 `bson:"vpcid"`
	VpcName              string                 `bson:"vpcname"`
	ZHash                int                    `bson:"zhash"`
	Zone                 int                    `bson:"zone"`
}
type mongoAttributesSparseCloudVPC struct {
	ID                   bson.ObjectId           `bson:"_id,omitempty"`
	RRN                  *string                 `bson:"rrn,omitempty"`
	AccountID            *string                 `bson:"accountid,omitempty"`
	AccountName          *string                 `bson:"accountname,omitempty"`
	Address              *string                 `bson:"address,omitempty"`
	Annotations          *map[string][]string    `bson:"annotations,omitempty"`
	AssociatedTags       *[]string               `bson:"associatedtags,omitempty"`
	CloudType            *CloudVPCCloudTypeValue `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey *string                 `bson:"createidempotencykey,omitempty"`
	Description          *string                 `bson:"description,omitempty"`
	InsertTS             *int                    `bson:"insertts,omitempty"`
	Metadata             *[]string               `bson:"metadata,omitempty"`
	MigrationsLog        *map[string]string      `bson:"migrationslog,omitempty"`
	Name                 *string                 `bson:"name,omitempty"`
	Namespace            *string                 `bson:"namespace,omitempty"`
	NativeID             *string                 `bson:"nativeid,omitempty"`
	NormalizedTags       *[]string               `bson:"normalizedtags,omitempty"`
	Protected            *bool                   `bson:"protected,omitempty"`
	RegionID             *string                 `bson:"regionid,omitempty"`
	RegionName           *string                 `bson:"regionname,omitempty"`
	Tags                 *map[string]string      `bson:"tags,omitempty"`
	UpdateIdempotencyKey *string                 `bson:"updateidempotencykey,omitempty"`
	Url                  *string                 `bson:"url,omitempty"`
	VpcID                *string                 `bson:"vpcid,omitempty"`
	VpcName              *string                 `bson:"vpcname,omitempty"`
	ZHash                *int                    `bson:"zhash,omitempty"`
	Zone                 *int                    `bson:"zone,omitempty"`
}
