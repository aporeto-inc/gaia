package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ProcessingUnitEnforcementStatusValue represents the possible values for attribute "enforcementStatus".
type ProcessingUnitEnforcementStatusValue string

const (
	// ProcessingUnitEnforcementStatusActive represents the value Active.
	ProcessingUnitEnforcementStatusActive ProcessingUnitEnforcementStatusValue = "Active"

	// ProcessingUnitEnforcementStatusFailed represents the value Failed.
	ProcessingUnitEnforcementStatusFailed ProcessingUnitEnforcementStatusValue = "Failed"

	// ProcessingUnitEnforcementStatusInactive represents the value Inactive.
	ProcessingUnitEnforcementStatusInactive ProcessingUnitEnforcementStatusValue = "Inactive"
)

// ProcessingUnitOperationalStatusValue represents the possible values for attribute "operationalStatus".
type ProcessingUnitOperationalStatusValue string

const (
	// ProcessingUnitOperationalStatusInitialized represents the value Initialized.
	ProcessingUnitOperationalStatusInitialized ProcessingUnitOperationalStatusValue = "Initialized"

	// ProcessingUnitOperationalStatusPaused represents the value Paused.
	ProcessingUnitOperationalStatusPaused ProcessingUnitOperationalStatusValue = "Paused"

	// ProcessingUnitOperationalStatusRunning represents the value Running.
	ProcessingUnitOperationalStatusRunning ProcessingUnitOperationalStatusValue = "Running"

	// ProcessingUnitOperationalStatusStopped represents the value Stopped.
	ProcessingUnitOperationalStatusStopped ProcessingUnitOperationalStatusValue = "Stopped"

	// ProcessingUnitOperationalStatusTerminated represents the value Terminated.
	ProcessingUnitOperationalStatusTerminated ProcessingUnitOperationalStatusValue = "Terminated"
)

// ProcessingUnitTypeValue represents the possible values for attribute "type".
type ProcessingUnitTypeValue string

const (
	// ProcessingUnitTypeAPIGateway represents the value APIGateway.
	ProcessingUnitTypeAPIGateway ProcessingUnitTypeValue = "APIGateway"

	// ProcessingUnitTypeDocker represents the value Docker.
	ProcessingUnitTypeDocker ProcessingUnitTypeValue = "Docker"

	// ProcessingUnitTypeHost represents the value Host.
	ProcessingUnitTypeHost ProcessingUnitTypeValue = "Host"

	// ProcessingUnitTypeHostService represents the value HostService.
	ProcessingUnitTypeHostService ProcessingUnitTypeValue = "HostService"

	// ProcessingUnitTypeLinuxService represents the value LinuxService.
	ProcessingUnitTypeLinuxService ProcessingUnitTypeValue = "LinuxService"

	// ProcessingUnitTypeRKT represents the value RKT.
	ProcessingUnitTypeRKT ProcessingUnitTypeValue = "RKT"

	// ProcessingUnitTypeSSHSession represents the value SSHSession.
	ProcessingUnitTypeSSHSession ProcessingUnitTypeValue = "SSHSession"

	// ProcessingUnitTypeUser represents the value User.
	ProcessingUnitTypeUser ProcessingUnitTypeValue = "User"
)

// ProcessingUnitIdentity represents the Identity of the object.
var ProcessingUnitIdentity = elemental.Identity{
	Name:     "processingunit",
	Category: "processingunits",
	Package:  "squall",
	Private:  false,
}

// ProcessingUnitsList represents a list of ProcessingUnits
type ProcessingUnitsList []*ProcessingUnit

// Identity returns the identity of the objects in the list.
func (o ProcessingUnitsList) Identity() elemental.Identity {

	return ProcessingUnitIdentity
}

// Copy returns a pointer to a copy the ProcessingUnitsList.
func (o ProcessingUnitsList) Copy() elemental.Identifiables {

	copy := append(ProcessingUnitsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ProcessingUnitsList.
func (o ProcessingUnitsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ProcessingUnitsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ProcessingUnit))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ProcessingUnitsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ProcessingUnitsList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the ProcessingUnitsList converted to SparseProcessingUnitsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ProcessingUnitsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseProcessingUnitsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseProcessingUnit)
	}

	return out
}

// Version returns the version of the content.
func (o ProcessingUnitsList) Version() int {

	return 1
}

// ProcessingUnit represents the model of a processingunit
type ProcessingUnit struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Defines if the object is archived.
	Archived bool `json:"-" msgpack:"-" bson:"archived" mapstructure:"-,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// A value of `true` indicates to the enforcer that it needs to collect information
	// for
	// this processing unit.
	CollectInfo bool `json:"collectInfo" msgpack:"collectInfo" bson:"collectinfo" mapstructure:"collectInfo,omitempty"`

	// Represents the latest information collected by the enforcer for this processing
	// unit.
	CollectedInfo map[string]string `json:"collectedInfo" msgpack:"collectedInfo" bson:"collectedinfo" mapstructure:"collectedInfo,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Contains the state of the enforcer for the processing unit. `Inactive`
	// (default):
	// the enforcer is not enforcing any host service. `Active`: the enforcer is
	// enforcing
	// a host service. `Failed`.
	EnforcementStatus ProcessingUnitEnforcementStatusValue `json:"enforcementStatus" msgpack:"enforcementStatus" bson:"enforcementstatus" mapstructure:"enforcementStatus,omitempty"`

	// The ID of the enforcer associated with the processing unit.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// The namespace of the enforcer associated with the processing unit.
	EnforcerNamespace string `json:"enforcerNamespace" msgpack:"enforcerNamespace" bson:"enforcernamespace" mapstructure:"enforcerNamespace,omitempty"`

	// This field is deprecated and it is there for backward compatibility. Use
	// `images` instead.
	Image string `json:"image" msgpack:"image" bson:"-" mapstructure:"image,omitempty"`

	// List of images or executable paths used by the processing unit.
	Images []string `json:"images" msgpack:"images" bson:"images" mapstructure:"images,omitempty"`

	// The date and time when the information was collected.
	LastCollectionTime time.Time `json:"lastCollectionTime" msgpack:"lastCollectionTime" bson:"lastcollectiontime" mapstructure:"lastCollectionTime,omitempty"`

	// Last poke is the time when the pu got last poked.
	LastPokeTime time.Time `json:"-" msgpack:"-" bson:"lastpoketime" mapstructure:"-,omitempty"`

	// The date and time of the last policy resolution.
	LastSyncTime time.Time `json:"lastSyncTime" msgpack:"lastSyncTime" bson:"lastsynctime" mapstructure:"lastSyncTime,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// The Docker UUID or service PID.
	NativeContextID string `json:"nativeContextID" msgpack:"nativeContextID" bson:"nativecontextid" mapstructure:"nativeContextID,omitempty"`

	// The list of services that this processing unit has declared that it will be
	// listening to,
	// either in its activation command or by exposing the ports in a container
	// manifest.
	NetworkServices []*ProcessingUnitService `json:"networkServices" msgpack:"networkServices" bson:"networkservices" mapstructure:"networkServices,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Operational status of the processing unit: `Initialized` (default), `Paused`,
	// `Running`,
	// `Stopped`, or `Terminated`.
	OperationalStatus ProcessingUnitOperationalStatusValue `json:"operationalStatus" msgpack:"operationalStatus" bson:"operationalstatus" mapstructure:"operationalStatus,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Indicates if this processing unit must be placed in tracing mode.
	Tracing *TraceMode `json:"tracing" msgpack:"tracing" bson:"tracing" mapstructure:"tracing,omitempty"`

	// Type of processing unit: `APIGateway`, `Docker`, `Host`, `HostService`,
	// `LinuxService`,
	// `RKT`, `User`, or `SSHSession`.
	Type ProcessingUnitTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	// The Aporeto control plane sets this value to `true` if it hasn't heard from the
	// processing
	// unit for more than five minutes.
	Unreachable bool `json:"unreachable" msgpack:"unreachable" bson:"unreachable" mapstructure:"unreachable,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewProcessingUnit returns a new *ProcessingUnit
func NewProcessingUnit() *ProcessingUnit {

	return &ProcessingUnit{
		ModelVersion:      1,
		Annotations:       map[string][]string{},
		AssociatedTags:    []string{},
		CollectedInfo:     map[string]string{},
		EnforcementStatus: ProcessingUnitEnforcementStatusInactive,
		NetworkServices:   []*ProcessingUnitService{},
		NormalizedTags:    []string{},
		Images:            []string{},
		OperationalStatus: ProcessingUnitOperationalStatusInitialized,
		Metadata:          []string{},
		Tracing:           NewTraceMode(),
		MigrationsLog:     map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *ProcessingUnit) Identity() elemental.Identity {

	return ProcessingUnitIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ProcessingUnit) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ProcessingUnit) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ProcessingUnit) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesProcessingUnit{}

	s.ID = bson.ObjectIdHex(o.ID)
	s.Annotations = o.Annotations
	s.Archived = o.Archived
	s.AssociatedTags = o.AssociatedTags
	s.CollectInfo = o.CollectInfo
	s.CollectedInfo = o.CollectedInfo
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.EnforcementStatus = o.EnforcementStatus
	s.EnforcerID = o.EnforcerID
	s.EnforcerNamespace = o.EnforcerNamespace
	s.Images = o.Images
	s.LastCollectionTime = o.LastCollectionTime
	s.LastPokeTime = o.LastPokeTime
	s.LastSyncTime = o.LastSyncTime
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NativeContextID = o.NativeContextID
	s.NetworkServices = o.NetworkServices
	s.NormalizedTags = o.NormalizedTags
	s.OperationalStatus = o.OperationalStatus
	s.Protected = o.Protected
	s.Tracing = o.Tracing
	s.Type = o.Type
	s.Unreachable = o.Unreachable
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ProcessingUnit) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesProcessingUnit{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.Archived = s.Archived
	o.AssociatedTags = s.AssociatedTags
	o.CollectInfo = s.CollectInfo
	o.CollectedInfo = s.CollectedInfo
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.EnforcementStatus = s.EnforcementStatus
	o.EnforcerID = s.EnforcerID
	o.EnforcerNamespace = s.EnforcerNamespace
	o.Images = s.Images
	o.LastCollectionTime = s.LastCollectionTime
	o.LastPokeTime = s.LastPokeTime
	o.LastSyncTime = s.LastSyncTime
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NativeContextID = s.NativeContextID
	o.NetworkServices = s.NetworkServices
	o.NormalizedTags = s.NormalizedTags
	o.OperationalStatus = s.OperationalStatus
	o.Protected = s.Protected
	o.Tracing = s.Tracing
	o.Type = s.Type
	o.Unreachable = s.Unreachable
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ProcessingUnit) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ProcessingUnit) BleveType() string {

	return "processingunit"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ProcessingUnit) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// Doc returns the documentation for the object
func (o *ProcessingUnit) Doc() string {

	return `A processing unit represents anything that can compute. It can be a Docker
container or a simple Unix process. Processing units are created, updated, and
deleted by
the system as they come and go. You can only modify their tags. Processing units
use network policies to define which other processing units or external
networks they can communicate with and file access policies to define what file
paths they can use.`
}

func (o *ProcessingUnit) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *ProcessingUnit) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *ProcessingUnit) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetArchived returns the Archived of the receiver.
func (o *ProcessingUnit) GetArchived() bool {

	return o.Archived
}

// SetArchived sets the property Archived of the receiver using the given value.
func (o *ProcessingUnit) SetArchived(archived bool) {

	o.Archived = archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *ProcessingUnit) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *ProcessingUnit) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *ProcessingUnit) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *ProcessingUnit) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *ProcessingUnit) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *ProcessingUnit) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *ProcessingUnit) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *ProcessingUnit) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *ProcessingUnit) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *ProcessingUnit) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *ProcessingUnit) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *ProcessingUnit) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *ProcessingUnit) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *ProcessingUnit) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *ProcessingUnit) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *ProcessingUnit) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *ProcessingUnit) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *ProcessingUnit) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *ProcessingUnit) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *ProcessingUnit) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *ProcessingUnit) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *ProcessingUnit) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *ProcessingUnit) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *ProcessingUnit) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *ProcessingUnit) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *ProcessingUnit) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *ProcessingUnit) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *ProcessingUnit) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ProcessingUnit) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseProcessingUnit{
			ID:                   &o.ID,
			Annotations:          &o.Annotations,
			Archived:             &o.Archived,
			AssociatedTags:       &o.AssociatedTags,
			CollectInfo:          &o.CollectInfo,
			CollectedInfo:        &o.CollectedInfo,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Description:          &o.Description,
			EnforcementStatus:    &o.EnforcementStatus,
			EnforcerID:           &o.EnforcerID,
			EnforcerNamespace:    &o.EnforcerNamespace,
			Image:                &o.Image,
			Images:               &o.Images,
			LastCollectionTime:   &o.LastCollectionTime,
			LastPokeTime:         &o.LastPokeTime,
			LastSyncTime:         &o.LastSyncTime,
			Metadata:             &o.Metadata,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NativeContextID:      &o.NativeContextID,
			NetworkServices:      &o.NetworkServices,
			NormalizedTags:       &o.NormalizedTags,
			OperationalStatus:    &o.OperationalStatus,
			Protected:            &o.Protected,
			Tracing:              o.Tracing,
			Type:                 &o.Type,
			Unreachable:          &o.Unreachable,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseProcessingUnit{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "archived":
			sp.Archived = &(o.Archived)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "collectInfo":
			sp.CollectInfo = &(o.CollectInfo)
		case "collectedInfo":
			sp.CollectedInfo = &(o.CollectedInfo)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "enforcementStatus":
			sp.EnforcementStatus = &(o.EnforcementStatus)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "image":
			sp.Image = &(o.Image)
		case "images":
			sp.Images = &(o.Images)
		case "lastCollectionTime":
			sp.LastCollectionTime = &(o.LastCollectionTime)
		case "lastPokeTime":
			sp.LastPokeTime = &(o.LastPokeTime)
		case "lastSyncTime":
			sp.LastSyncTime = &(o.LastSyncTime)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "nativeContextID":
			sp.NativeContextID = &(o.NativeContextID)
		case "networkServices":
			sp.NetworkServices = &(o.NetworkServices)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "operationalStatus":
			sp.OperationalStatus = &(o.OperationalStatus)
		case "protected":
			sp.Protected = &(o.Protected)
		case "tracing":
			sp.Tracing = o.Tracing
		case "type":
			sp.Type = &(o.Type)
		case "unreachable":
			sp.Unreachable = &(o.Unreachable)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseProcessingUnit to the object.
func (o *ProcessingUnit) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseProcessingUnit)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.Archived != nil {
		o.Archived = *so.Archived
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CollectInfo != nil {
		o.CollectInfo = *so.CollectInfo
	}
	if so.CollectedInfo != nil {
		o.CollectedInfo = *so.CollectedInfo
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.EnforcementStatus != nil {
		o.EnforcementStatus = *so.EnforcementStatus
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.Image != nil {
		o.Image = *so.Image
	}
	if so.Images != nil {
		o.Images = *so.Images
	}
	if so.LastCollectionTime != nil {
		o.LastCollectionTime = *so.LastCollectionTime
	}
	if so.LastPokeTime != nil {
		o.LastPokeTime = *so.LastPokeTime
	}
	if so.LastSyncTime != nil {
		o.LastSyncTime = *so.LastSyncTime
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
	if so.NativeContextID != nil {
		o.NativeContextID = *so.NativeContextID
	}
	if so.NetworkServices != nil {
		o.NetworkServices = *so.NetworkServices
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.OperationalStatus != nil {
		o.OperationalStatus = *so.OperationalStatus
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Tracing != nil {
		o.Tracing = so.Tracing
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
	if so.Unreachable != nil {
		o.Unreachable = *so.Unreachable
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the ProcessingUnit.
func (o *ProcessingUnit) DeepCopy() *ProcessingUnit {

	if o == nil {
		return nil
	}

	out := &ProcessingUnit{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ProcessingUnit.
func (o *ProcessingUnit) DeepCopyInto(out *ProcessingUnit) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ProcessingUnit: %s", err))
	}

	*out = *target.(*ProcessingUnit)
}

// Validate valides the current information stored into the structure.
func (o *ProcessingUnit) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("enforcementStatus", string(o.EnforcementStatus), []string{"Active", "Failed", "Inactive"}, false); err != nil {
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

	for _, sub := range o.NetworkServices {
		if sub == nil {
			continue
		}
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := ValidateProcessingUnitServicesList("networkServices", o.NetworkServices); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("operationalStatus", string(o.OperationalStatus), []string{"Initialized", "Paused", "Running", "Stopped", "Terminated"}, false); err != nil {
		errors = errors.Append(err)
	}

	if o.Tracing != nil {
		if err := o.Tracing.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"APIGateway", "Docker", "Host", "HostService", "LinuxService", "RKT", "User", "SSHSession"}, false); err != nil {
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
func (*ProcessingUnit) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ProcessingUnitAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ProcessingUnitLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ProcessingUnit) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ProcessingUnitAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ProcessingUnit) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "archived":
		return o.Archived
	case "associatedTags":
		return o.AssociatedTags
	case "collectInfo":
		return o.CollectInfo
	case "collectedInfo":
		return o.CollectedInfo
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "enforcementStatus":
		return o.EnforcementStatus
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "image":
		return o.Image
	case "images":
		return o.Images
	case "lastCollectionTime":
		return o.LastCollectionTime
	case "lastPokeTime":
		return o.LastPokeTime
	case "lastSyncTime":
		return o.LastSyncTime
	case "metadata":
		return o.Metadata
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "nativeContextID":
		return o.NativeContextID
	case "networkServices":
		return o.NetworkServices
	case "normalizedTags":
		return o.NormalizedTags
	case "operationalStatus":
		return o.OperationalStatus
	case "protected":
		return o.Protected
	case "tracing":
		return o.Tracing
	case "type":
		return o.Type
	case "unreachable":
		return o.Unreachable
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// ProcessingUnitAttributesMap represents the map of attribute for ProcessingUnit.
var ProcessingUnitAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"Archived": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Archived",
		Description:    `Defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"CollectInfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectInfo",
		Description: `A value of ` + "`" + `true` + "`" + ` indicates to the enforcer that it needs to collect information
for
this processing unit.`,
		Exposed: true,
		Name:    "collectInfo",
		Stored:  true,
		Type:    "boolean",
	},
	"CollectedInfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectedInfo",
		Description: `Represents the latest information collected by the enforcer for this processing
unit.`,
		Exposed: true,
		Name:    "collectedInfo",
		Stored:  true,
		SubType: "map[string]string",
		Type:    "external",
	},
	"CreateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"EnforcementStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Failed", "Inactive"},
		ConvertedName:  "EnforcementStatus",
		DefaultValue:   ProcessingUnitEnforcementStatusInactive,
		Description: `Contains the state of the enforcer for the processing unit. ` + "`" + `Inactive` + "`" + `
(default):
the enforcer is not enforcing any host service. ` + "`" + `Active` + "`" + `: the enforcer is
enforcing
a host service. ` + "`" + `Failed` + "`" + `.`,
		Exposed:    true,
		Filterable: true,
		Name:       "enforcementStatus",
		Stored:     true,
		Type:       "enum",
	},
	"EnforcerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `The ID of the enforcer associated with the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcerID",
		Stored:         true,
		Type:           "string",
	},
	"EnforcerNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `The namespace of the enforcer associated with the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcerNamespace",
		Stored:         true,
		Type:           "string",
	},
	"Image": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Image",
		Deprecated:     true,
		Description: `This field is deprecated and it is there for backward compatibility. Use
` + "`" + `images` + "`" + ` instead.`,
		Exposed: true,
		Name:    "image",
		Type:    "string",
	},
	"Images": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Images",
		CreationOnly:   true,
		Description:    `List of images or executable paths used by the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "images",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"LastCollectionTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastCollectionTime",
		Description:    `The date and time when the information was collected.`,
		Exposed:        true,
		Name:           "lastCollectionTime",
		Stored:         true,
		Type:           "time",
	},
	"LastPokeTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastPokeTime",
		Description:    `Last poke is the time when the pu got last poked.`,
		Name:           "lastPokeTime",
		Stored:         true,
		Type:           "time",
	},
	"LastSyncTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastSyncTime",
		Description:    `The date and time of the last policy resolution.`,
		Exposed:        true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"MigrationsLog": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"NativeContextID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NativeContextID",
		Description:    `The Docker UUID or service PID.`,
		Exposed:        true,
		Name:           "nativeContextID",
		Stored:         true,
		Type:           "string",
	},
	"NetworkServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkServices",
		Description: `The list of services that this processing unit has declared that it will be
listening to,
either in its activation command or by exposing the ports in a container
manifest.`,
		Exposed:   true,
		Name:      "networkServices",
		Orderable: true,
		Stored:    true,
		SubType:   "processingunitservice",
		Type:      "refList",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"OperationalStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Initialized", "Paused", "Running", "Stopped", "Terminated"},
		ConvertedName:  "OperationalStatus",
		DefaultValue:   ProcessingUnitOperationalStatusInitialized,
		Description: `Operational status of the processing unit: ` + "`" + `Initialized` + "`" + ` (default), ` + "`" + `Paused` + "`" + `,
` + "`" + `Running` + "`" + `,
` + "`" + `Stopped` + "`" + `, or ` + "`" + `Terminated` + "`" + `.`,
		Exposed:    true,
		Filterable: true,
		Name:       "operationalStatus",
		Stored:     true,
		Type:       "enum",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"Tracing": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Tracing",
		Description:    `Indicates if this processing unit must be placed in tracing mode.`,
		Exposed:        true,
		Name:           "tracing",
		Stored:         true,
		SubType:        "tracemode",
		Type:           "ref",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"APIGateway", "Docker", "Host", "HostService", "LinuxService", "RKT", "User", "SSHSession"},
		ConvertedName:  "Type",
		CreationOnly:   true,
		Description: `Type of processing unit: ` + "`" + `APIGateway` + "`" + `, ` + "`" + `Docker` + "`" + `, ` + "`" + `Host` + "`" + `, ` + "`" + `HostService` + "`" + `,
` + "`" + `LinuxService` + "`" + `,
` + "`" + `RKT` + "`" + `, ` + "`" + `User` + "`" + `, or ` + "`" + `SSHSession` + "`" + `.`,
		Exposed:    true,
		Filterable: true,
		Name:       "type",
		Stored:     true,
		Type:       "enum",
	},
	"Unreachable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Unreachable",
		Description: `The Aporeto control plane sets this value to ` + "`" + `true` + "`" + ` if it hasn't heard from the
processing
unit for more than five minutes.`,
		Exposed:  true,
		Name:     "unreachable",
		ReadOnly: true,
		Stored:   true,
		Type:     "boolean",
	},
	"UpdateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"ZHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// ProcessingUnitLowerCaseAttributesMap represents the map of attribute for ProcessingUnit.
var ProcessingUnitLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"archived": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Archived",
		Description:    `Defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"collectinfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectInfo",
		Description: `A value of ` + "`" + `true` + "`" + ` indicates to the enforcer that it needs to collect information
for
this processing unit.`,
		Exposed: true,
		Name:    "collectInfo",
		Stored:  true,
		Type:    "boolean",
	},
	"collectedinfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectedInfo",
		Description: `Represents the latest information collected by the enforcer for this processing
unit.`,
		Exposed: true,
		Name:    "collectedInfo",
		Stored:  true,
		SubType: "map[string]string",
		Type:    "external",
	},
	"createidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"enforcementstatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Failed", "Inactive"},
		ConvertedName:  "EnforcementStatus",
		DefaultValue:   ProcessingUnitEnforcementStatusInactive,
		Description: `Contains the state of the enforcer for the processing unit. ` + "`" + `Inactive` + "`" + `
(default):
the enforcer is not enforcing any host service. ` + "`" + `Active` + "`" + `: the enforcer is
enforcing
a host service. ` + "`" + `Failed` + "`" + `.`,
		Exposed:    true,
		Filterable: true,
		Name:       "enforcementStatus",
		Stored:     true,
		Type:       "enum",
	},
	"enforcerid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `The ID of the enforcer associated with the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcerID",
		Stored:         true,
		Type:           "string",
	},
	"enforcernamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `The namespace of the enforcer associated with the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcerNamespace",
		Stored:         true,
		Type:           "string",
	},
	"image": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Image",
		Deprecated:     true,
		Description: `This field is deprecated and it is there for backward compatibility. Use
` + "`" + `images` + "`" + ` instead.`,
		Exposed: true,
		Name:    "image",
		Type:    "string",
	},
	"images": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Images",
		CreationOnly:   true,
		Description:    `List of images or executable paths used by the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "images",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"lastcollectiontime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastCollectionTime",
		Description:    `The date and time when the information was collected.`,
		Exposed:        true,
		Name:           "lastCollectionTime",
		Stored:         true,
		Type:           "time",
	},
	"lastpoketime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastPokeTime",
		Description:    `Last poke is the time when the pu got last poked.`,
		Name:           "lastPokeTime",
		Stored:         true,
		Type:           "time",
	},
	"lastsynctime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastSyncTime",
		Description:    `The date and time of the last policy resolution.`,
		Exposed:        true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"migrationslog": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"nativecontextid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NativeContextID",
		Description:    `The Docker UUID or service PID.`,
		Exposed:        true,
		Name:           "nativeContextID",
		Stored:         true,
		Type:           "string",
	},
	"networkservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkServices",
		Description: `The list of services that this processing unit has declared that it will be
listening to,
either in its activation command or by exposing the ports in a container
manifest.`,
		Exposed:   true,
		Name:      "networkServices",
		Orderable: true,
		Stored:    true,
		SubType:   "processingunitservice",
		Type:      "refList",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"operationalstatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Initialized", "Paused", "Running", "Stopped", "Terminated"},
		ConvertedName:  "OperationalStatus",
		DefaultValue:   ProcessingUnitOperationalStatusInitialized,
		Description: `Operational status of the processing unit: ` + "`" + `Initialized` + "`" + ` (default), ` + "`" + `Paused` + "`" + `,
` + "`" + `Running` + "`" + `,
` + "`" + `Stopped` + "`" + `, or ` + "`" + `Terminated` + "`" + `.`,
		Exposed:    true,
		Filterable: true,
		Name:       "operationalStatus",
		Stored:     true,
		Type:       "enum",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"tracing": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Tracing",
		Description:    `Indicates if this processing unit must be placed in tracing mode.`,
		Exposed:        true,
		Name:           "tracing",
		Stored:         true,
		SubType:        "tracemode",
		Type:           "ref",
	},
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{"APIGateway", "Docker", "Host", "HostService", "LinuxService", "RKT", "User", "SSHSession"},
		ConvertedName:  "Type",
		CreationOnly:   true,
		Description: `Type of processing unit: ` + "`" + `APIGateway` + "`" + `, ` + "`" + `Docker` + "`" + `, ` + "`" + `Host` + "`" + `, ` + "`" + `HostService` + "`" + `,
` + "`" + `LinuxService` + "`" + `,
` + "`" + `RKT` + "`" + `, ` + "`" + `User` + "`" + `, or ` + "`" + `SSHSession` + "`" + `.`,
		Exposed:    true,
		Filterable: true,
		Name:       "type",
		Stored:     true,
		Type:       "enum",
	},
	"unreachable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Unreachable",
		Description: `The Aporeto control plane sets this value to ` + "`" + `true` + "`" + ` if it hasn't heard from the
processing
unit for more than five minutes.`,
		Exposed:  true,
		Name:     "unreachable",
		ReadOnly: true,
		Stored:   true,
		Type:     "boolean",
	},
	"updateidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"zhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseProcessingUnitsList represents a list of SparseProcessingUnits
type SparseProcessingUnitsList []*SparseProcessingUnit

// Identity returns the identity of the objects in the list.
func (o SparseProcessingUnitsList) Identity() elemental.Identity {

	return ProcessingUnitIdentity
}

// Copy returns a pointer to a copy the SparseProcessingUnitsList.
func (o SparseProcessingUnitsList) Copy() elemental.Identifiables {

	copy := append(SparseProcessingUnitsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseProcessingUnitsList.
func (o SparseProcessingUnitsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseProcessingUnitsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseProcessingUnit))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseProcessingUnitsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseProcessingUnitsList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparseProcessingUnitsList converted to ProcessingUnitsList.
func (o SparseProcessingUnitsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseProcessingUnitsList) Version() int {

	return 1
}

// SparseProcessingUnit represents the sparse version of a processingunit.
type SparseProcessingUnit struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// Defines if the object is archived.
	Archived *bool `json:"-" msgpack:"-" bson:"archived,omitempty" mapstructure:"-,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// A value of `true` indicates to the enforcer that it needs to collect information
	// for
	// this processing unit.
	CollectInfo *bool `json:"collectInfo,omitempty" msgpack:"collectInfo,omitempty" bson:"collectinfo,omitempty" mapstructure:"collectInfo,omitempty"`

	// Represents the latest information collected by the enforcer for this processing
	// unit.
	CollectedInfo *map[string]string `json:"collectedInfo,omitempty" msgpack:"collectedInfo,omitempty" bson:"collectedinfo,omitempty" mapstructure:"collectedInfo,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Contains the state of the enforcer for the processing unit. `Inactive`
	// (default):
	// the enforcer is not enforcing any host service. `Active`: the enforcer is
	// enforcing
	// a host service. `Failed`.
	EnforcementStatus *ProcessingUnitEnforcementStatusValue `json:"enforcementStatus,omitempty" msgpack:"enforcementStatus,omitempty" bson:"enforcementstatus,omitempty" mapstructure:"enforcementStatus,omitempty"`

	// The ID of the enforcer associated with the processing unit.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid,omitempty" mapstructure:"enforcerID,omitempty"`

	// The namespace of the enforcer associated with the processing unit.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"enforcernamespace,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// This field is deprecated and it is there for backward compatibility. Use
	// `images` instead.
	Image *string `json:"image,omitempty" msgpack:"image,omitempty" bson:"-" mapstructure:"image,omitempty"`

	// List of images or executable paths used by the processing unit.
	Images *[]string `json:"images,omitempty" msgpack:"images,omitempty" bson:"images,omitempty" mapstructure:"images,omitempty"`

	// The date and time when the information was collected.
	LastCollectionTime *time.Time `json:"lastCollectionTime,omitempty" msgpack:"lastCollectionTime,omitempty" bson:"lastcollectiontime,omitempty" mapstructure:"lastCollectionTime,omitempty"`

	// Last poke is the time when the pu got last poked.
	LastPokeTime *time.Time `json:"-" msgpack:"-" bson:"lastpoketime,omitempty" mapstructure:"-,omitempty"`

	// The date and time of the last policy resolution.
	LastSyncTime *time.Time `json:"lastSyncTime,omitempty" msgpack:"lastSyncTime,omitempty" bson:"lastsynctime,omitempty" mapstructure:"lastSyncTime,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The Docker UUID or service PID.
	NativeContextID *string `json:"nativeContextID,omitempty" msgpack:"nativeContextID,omitempty" bson:"nativecontextid,omitempty" mapstructure:"nativeContextID,omitempty"`

	// The list of services that this processing unit has declared that it will be
	// listening to,
	// either in its activation command or by exposing the ports in a container
	// manifest.
	NetworkServices *[]*ProcessingUnitService `json:"networkServices,omitempty" msgpack:"networkServices,omitempty" bson:"networkservices,omitempty" mapstructure:"networkServices,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Operational status of the processing unit: `Initialized` (default), `Paused`,
	// `Running`,
	// `Stopped`, or `Terminated`.
	OperationalStatus *ProcessingUnitOperationalStatusValue `json:"operationalStatus,omitempty" msgpack:"operationalStatus,omitempty" bson:"operationalstatus,omitempty" mapstructure:"operationalStatus,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Indicates if this processing unit must be placed in tracing mode.
	Tracing *TraceMode `json:"tracing,omitempty" msgpack:"tracing,omitempty" bson:"tracing,omitempty" mapstructure:"tracing,omitempty"`

	// Type of processing unit: `APIGateway`, `Docker`, `Host`, `HostService`,
	// `LinuxService`,
	// `RKT`, `User`, or `SSHSession`.
	Type *ProcessingUnitTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

	// The Aporeto control plane sets this value to `true` if it hasn't heard from the
	// processing
	// unit for more than five minutes.
	Unreachable *bool `json:"unreachable,omitempty" msgpack:"unreachable,omitempty" bson:"unreachable,omitempty" mapstructure:"unreachable,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseProcessingUnit returns a new  SparseProcessingUnit.
func NewSparseProcessingUnit() *SparseProcessingUnit {
	return &SparseProcessingUnit{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseProcessingUnit) Identity() elemental.Identity {

	return ProcessingUnitIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseProcessingUnit) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseProcessingUnit) SetIdentifier(id string) {

	o.ID = &id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseProcessingUnit) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseProcessingUnit{}

	s.ID = bson.ObjectIdHex(*o.ID)
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.Archived != nil {
		s.Archived = o.Archived
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CollectInfo != nil {
		s.CollectInfo = o.CollectInfo
	}
	if o.CollectedInfo != nil {
		s.CollectedInfo = o.CollectedInfo
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.EnforcementStatus != nil {
		s.EnforcementStatus = o.EnforcementStatus
	}
	if o.EnforcerID != nil {
		s.EnforcerID = o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		s.EnforcerNamespace = o.EnforcerNamespace
	}
	if o.Images != nil {
		s.Images = o.Images
	}
	if o.LastCollectionTime != nil {
		s.LastCollectionTime = o.LastCollectionTime
	}
	if o.LastPokeTime != nil {
		s.LastPokeTime = o.LastPokeTime
	}
	if o.LastSyncTime != nil {
		s.LastSyncTime = o.LastSyncTime
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
	if o.NativeContextID != nil {
		s.NativeContextID = o.NativeContextID
	}
	if o.NetworkServices != nil {
		s.NetworkServices = o.NetworkServices
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.OperationalStatus != nil {
		s.OperationalStatus = o.OperationalStatus
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.Tracing != nil {
		s.Tracing = o.Tracing
	}
	if o.Type != nil {
		s.Type = o.Type
	}
	if o.Unreachable != nil {
		s.Unreachable = o.Unreachable
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
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
func (o *SparseProcessingUnit) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseProcessingUnit{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.Archived != nil {
		o.Archived = s.Archived
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CollectInfo != nil {
		o.CollectInfo = s.CollectInfo
	}
	if s.CollectedInfo != nil {
		o.CollectedInfo = s.CollectedInfo
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.EnforcementStatus != nil {
		o.EnforcementStatus = s.EnforcementStatus
	}
	if s.EnforcerID != nil {
		o.EnforcerID = s.EnforcerID
	}
	if s.EnforcerNamespace != nil {
		o.EnforcerNamespace = s.EnforcerNamespace
	}
	if s.Images != nil {
		o.Images = s.Images
	}
	if s.LastCollectionTime != nil {
		o.LastCollectionTime = s.LastCollectionTime
	}
	if s.LastPokeTime != nil {
		o.LastPokeTime = s.LastPokeTime
	}
	if s.LastSyncTime != nil {
		o.LastSyncTime = s.LastSyncTime
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
	if s.NativeContextID != nil {
		o.NativeContextID = s.NativeContextID
	}
	if s.NetworkServices != nil {
		o.NetworkServices = s.NetworkServices
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.OperationalStatus != nil {
		o.OperationalStatus = s.OperationalStatus
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.Tracing != nil {
		o.Tracing = s.Tracing
	}
	if s.Type != nil {
		o.Type = s.Type
	}
	if s.Unreachable != nil {
		o.Unreachable = s.Unreachable
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
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
func (o *SparseProcessingUnit) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseProcessingUnit) ToPlain() elemental.PlainIdentifiable {

	out := NewProcessingUnit()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.Archived != nil {
		out.Archived = *o.Archived
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CollectInfo != nil {
		out.CollectInfo = *o.CollectInfo
	}
	if o.CollectedInfo != nil {
		out.CollectedInfo = *o.CollectedInfo
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.EnforcementStatus != nil {
		out.EnforcementStatus = *o.EnforcementStatus
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.Image != nil {
		out.Image = *o.Image
	}
	if o.Images != nil {
		out.Images = *o.Images
	}
	if o.LastCollectionTime != nil {
		out.LastCollectionTime = *o.LastCollectionTime
	}
	if o.LastPokeTime != nil {
		out.LastPokeTime = *o.LastPokeTime
	}
	if o.LastSyncTime != nil {
		out.LastSyncTime = *o.LastSyncTime
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
	if o.NativeContextID != nil {
		out.NativeContextID = *o.NativeContextID
	}
	if o.NetworkServices != nil {
		out.NetworkServices = *o.NetworkServices
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.OperationalStatus != nil {
		out.OperationalStatus = *o.OperationalStatus
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Tracing != nil {
		out.Tracing = o.Tracing
	}
	if o.Type != nil {
		out.Type = *o.Type
	}
	if o.Unreachable != nil {
		out.Unreachable = *o.Unreachable
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseProcessingUnit) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetArchived returns the Archived of the receiver.
func (o *SparseProcessingUnit) GetArchived() bool {

	return *o.Archived
}

// SetArchived sets the property Archived of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetArchived(archived bool) {

	o.Archived = &archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseProcessingUnit) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseProcessingUnit) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseProcessingUnit) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseProcessingUnit) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseProcessingUnit) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseProcessingUnit) GetMigrationsLog() map[string]string {

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseProcessingUnit) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseProcessingUnit) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseProcessingUnit) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseProcessingUnit) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseProcessingUnit) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseProcessingUnit) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseProcessingUnit) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseProcessingUnit) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseProcessingUnit.
func (o *SparseProcessingUnit) DeepCopy() *SparseProcessingUnit {

	if o == nil {
		return nil
	}

	out := &SparseProcessingUnit{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseProcessingUnit.
func (o *SparseProcessingUnit) DeepCopyInto(out *SparseProcessingUnit) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseProcessingUnit: %s", err))
	}

	*out = *target.(*SparseProcessingUnit)
}

type mongoAttributesProcessingUnit struct {
	ID                   bson.ObjectId                        `bson:"_id"`
	Annotations          map[string][]string                  `bson:"annotations"`
	Archived             bool                                 `bson:"archived"`
	AssociatedTags       []string                             `bson:"associatedtags"`
	CollectInfo          bool                                 `bson:"collectinfo"`
	CollectedInfo        map[string]string                    `bson:"collectedinfo"`
	CreateIdempotencyKey string                               `bson:"createidempotencykey"`
	CreateTime           time.Time                            `bson:"createtime"`
	Description          string                               `bson:"description"`
	EnforcementStatus    ProcessingUnitEnforcementStatusValue `bson:"enforcementstatus"`
	EnforcerID           string                               `bson:"enforcerid"`
	EnforcerNamespace    string                               `bson:"enforcernamespace"`
	Images               []string                             `bson:"images"`
	LastCollectionTime   time.Time                            `bson:"lastcollectiontime"`
	LastPokeTime         time.Time                            `bson:"lastpoketime"`
	LastSyncTime         time.Time                            `bson:"lastsynctime"`
	Metadata             []string                             `bson:"metadata"`
	MigrationsLog        map[string]string                    `bson:"migrationslog"`
	Name                 string                               `bson:"name"`
	Namespace            string                               `bson:"namespace"`
	NativeContextID      string                               `bson:"nativecontextid"`
	NetworkServices      []*ProcessingUnitService             `bson:"networkservices"`
	NormalizedTags       []string                             `bson:"normalizedtags"`
	OperationalStatus    ProcessingUnitOperationalStatusValue `bson:"operationalstatus"`
	Protected            bool                                 `bson:"protected"`
	Tracing              *TraceMode                           `bson:"tracing"`
	Type                 ProcessingUnitTypeValue              `bson:"type"`
	Unreachable          bool                                 `bson:"unreachable"`
	UpdateIdempotencyKey string                               `bson:"updateidempotencykey"`
	UpdateTime           time.Time                            `bson:"updatetime"`
	ZHash                int                                  `bson:"zhash"`
	Zone                 int                                  `bson:"zone"`
}
type mongoAttributesSparseProcessingUnit struct {
	ID                   bson.ObjectId                         `bson:"_id"`
	Annotations          *map[string][]string                  `bson:"annotations,omitempty"`
	Archived             *bool                                 `bson:"archived,omitempty"`
	AssociatedTags       *[]string                             `bson:"associatedtags,omitempty"`
	CollectInfo          *bool                                 `bson:"collectinfo,omitempty"`
	CollectedInfo        *map[string]string                    `bson:"collectedinfo,omitempty"`
	CreateIdempotencyKey *string                               `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time                            `bson:"createtime,omitempty"`
	Description          *string                               `bson:"description,omitempty"`
	EnforcementStatus    *ProcessingUnitEnforcementStatusValue `bson:"enforcementstatus,omitempty"`
	EnforcerID           *string                               `bson:"enforcerid,omitempty"`
	EnforcerNamespace    *string                               `bson:"enforcernamespace,omitempty"`
	Images               *[]string                             `bson:"images,omitempty"`
	LastCollectionTime   *time.Time                            `bson:"lastcollectiontime,omitempty"`
	LastPokeTime         *time.Time                            `bson:"lastpoketime,omitempty"`
	LastSyncTime         *time.Time                            `bson:"lastsynctime,omitempty"`
	Metadata             *[]string                             `bson:"metadata,omitempty"`
	MigrationsLog        *map[string]string                    `bson:"migrationslog,omitempty"`
	Name                 *string                               `bson:"name,omitempty"`
	Namespace            *string                               `bson:"namespace,omitempty"`
	NativeContextID      *string                               `bson:"nativecontextid,omitempty"`
	NetworkServices      *[]*ProcessingUnitService             `bson:"networkservices,omitempty"`
	NormalizedTags       *[]string                             `bson:"normalizedtags,omitempty"`
	OperationalStatus    *ProcessingUnitOperationalStatusValue `bson:"operationalstatus,omitempty"`
	Protected            *bool                                 `bson:"protected,omitempty"`
	Tracing              *TraceMode                            `bson:"tracing,omitempty"`
	Type                 *ProcessingUnitTypeValue              `bson:"type,omitempty"`
	Unreachable          *bool                                 `bson:"unreachable,omitempty"`
	UpdateIdempotencyKey *string                               `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time                            `bson:"updatetime,omitempty"`
	ZHash                *int                                  `bson:"zhash,omitempty"`
	Zone                 *int                                  `bson:"zone,omitempty"`
}
