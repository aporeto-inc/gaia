package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EnforcerEnforcementStatusValue represents the possible values for attribute "enforcementStatus".
type EnforcerEnforcementStatusValue string

const (
	// EnforcerEnforcementStatusActive represents the value Active.
	EnforcerEnforcementStatusActive EnforcerEnforcementStatusValue = "Active"

	// EnforcerEnforcementStatusFailed represents the value Failed.
	EnforcerEnforcementStatusFailed EnforcerEnforcementStatusValue = "Failed"

	// EnforcerEnforcementStatusInactive represents the value Inactive.
	EnforcerEnforcementStatusInactive EnforcerEnforcementStatusValue = "Inactive"
)

// EnforcerLogLevelValue represents the possible values for attribute "logLevel".
type EnforcerLogLevelValue string

const (
	// EnforcerLogLevelDebug represents the value Debug.
	EnforcerLogLevelDebug EnforcerLogLevelValue = "Debug"

	// EnforcerLogLevelError represents the value Error.
	EnforcerLogLevelError EnforcerLogLevelValue = "Error"

	// EnforcerLogLevelInfo represents the value Info.
	EnforcerLogLevelInfo EnforcerLogLevelValue = "Info"

	// EnforcerLogLevelTrace represents the value Trace.
	EnforcerLogLevelTrace EnforcerLogLevelValue = "Trace"

	// EnforcerLogLevelWarn represents the value Warn.
	EnforcerLogLevelWarn EnforcerLogLevelValue = "Warn"
)

// EnforcerOperationalStatusValue represents the possible values for attribute "operationalStatus".
type EnforcerOperationalStatusValue string

const (
	// EnforcerOperationalStatusConnected represents the value Connected.
	EnforcerOperationalStatusConnected EnforcerOperationalStatusValue = "Connected"

	// EnforcerOperationalStatusDisconnected represents the value Disconnected.
	EnforcerOperationalStatusDisconnected EnforcerOperationalStatusValue = "Disconnected"

	// EnforcerOperationalStatusInitialized represents the value Initialized.
	EnforcerOperationalStatusInitialized EnforcerOperationalStatusValue = "Initialized"

	// EnforcerOperationalStatusRegistered represents the value Registered.
	EnforcerOperationalStatusRegistered EnforcerOperationalStatusValue = "Registered"
)

// EnforcerIdentity represents the Identity of the object.
var EnforcerIdentity = elemental.Identity{
	Name:     "enforcer",
	Category: "enforcers",
	Package:  "squall",
	Private:  false,
}

// EnforcersList represents a list of Enforcers
type EnforcersList []*Enforcer

// Identity returns the identity of the objects in the list.
func (o EnforcersList) Identity() elemental.Identity {

	return EnforcerIdentity
}

// Copy returns a pointer to a copy the EnforcersList.
func (o EnforcersList) Copy() elemental.Identifiables {

	copy := append(EnforcersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcersList.
func (o EnforcersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EnforcersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Enforcer))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcersList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the EnforcersList converted to SparseEnforcersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EnforcersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseEnforcersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseEnforcer)
	}

	return out
}

// Version returns the version of the content.
func (o EnforcersList) Version() int {

	return 1
}

// Enforcer represents the model of a enforcer
type Enforcer struct {
	// Contains the fully qualified domain name (FQDN) of the server where the
	// enforcer is running.
	FQDN string `json:"FQDN" msgpack:"FQDN" bson:"fqdn" mapstructure:"FQDN,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// The certificate of the enforcer.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"certificate" mapstructure:"certificate,omitempty"`

	// The expiration date of the certificate. This is an
	// internal attribute, not exposed in the API.
	CertificateExpirationDate time.Time `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	// The certificate key of the enforcer. This is an internal
	// attribute, not exposed in the API.
	CertificateKey string `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	// If not empty during a create or update operation, the provided certificate
	// signing request (CSR) will be validated and signed by the control plane,
	// providing a renewed certificate.
	CertificateRequest string `json:"certificateRequest" msgpack:"certificateRequest" bson:"-" mapstructure:"certificateRequest,omitempty"`

	// Indicates to the enforcer whether or not it needs to collect information.
	CollectInfo bool `json:"collectInfo" msgpack:"collectInfo" bson:"collectinfo" mapstructure:"collectInfo,omitempty"`

	// Represents the latest information collected by the enforcer.
	CollectedInfo map[string]string `json:"collectedInfo" msgpack:"collectedInfo" bson:"collectedinfo" mapstructure:"collectedInfo,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// The version number of the installed enforcer binary.
	CurrentVersion string `json:"currentVersion" msgpack:"currentVersion" bson:"currentversion" mapstructure:"currentVersion,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Status of the enforcement for host services.
	EnforcementStatus EnforcerEnforcementStatusValue `json:"enforcementStatus" msgpack:"enforcementStatus" bson:"enforcementstatus" mapstructure:"enforcementStatus,omitempty"`

	// Identifies when the information was collected.
	LastCollectionTime time.Time `json:"lastCollectionTime" msgpack:"lastCollectionTime" bson:"lastcollectiontime" mapstructure:"lastCollectionTime,omitempty"`

	// The time and date of the last poke.
	LastPokeTime time.Time `json:"-" msgpack:"-" bson:"lastpoketime" mapstructure:"-,omitempty"`

	// The time and date of the last heartbeat.
	LastSyncTime time.Time `json:"lastSyncTime" msgpack:"lastSyncTime" bson:"lastsynctime" mapstructure:"lastSyncTime,omitempty"`

	// LastValidHostServices is a read only attribute that stores the list valid host
	// services that have been applied to this enforcer. This list might be different
	// from the list retrieved through policy, if the dynamically calculated list leads
	// into conflicts.
	LastValidHostServices HostServicesList `json:"-" msgpack:"-" bson:"lastvalidhostservices" mapstructure:"-,omitempty"`

	// Contains the initial chain of trust for the enforcer. This value is only
	// given when you retrieve a single enforcer.
	LocalCA string `json:"localCA" msgpack:"localCA" bson:"-" mapstructure:"localCA,omitempty"`

	// Log level of the enforcer.
	LogLevel EnforcerLogLevelValue `json:"logLevel" msgpack:"logLevel" bson:"loglevel" mapstructure:"logLevel,omitempty"`

	// Determines the duration of which the log level will be active, using [Golang
	// duration syntax](https://golang.org/pkg/time/#example_Duration).
	LogLevelDuration string `json:"logLevelDuration" msgpack:"logLevelDuration" bson:"loglevelduration" mapstructure:"logLevelDuration,omitempty"`

	// A unique identifier for every machine as detected by the enforcer. It is
	// based on hardware information such as the SMBIOS UUID, MAC addresses of
	// interfaces, or cloud provider IDs.
	MachineID string `json:"machineID" msgpack:"machineID" bson:"machineid" mapstructure:"machineID,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// The status of the enforcer.
	OperationalStatus EnforcerOperationalStatusValue `json:"operationalStatus" msgpack:"operationalStatus" bson:"operationalstatus" mapstructure:"operationalStatus,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// The public token of the server that will be included in the datapath and
	// is signed by the private certificate authority.
	PublicToken string `json:"publicToken" msgpack:"publicToken" bson:"publictoken" mapstructure:"publicToken,omitempty"`

	// The time and date on which this enforcer was started. The enforcer reports
	// this and the value is preserved across disconnects.
	StartTime time.Time `json:"startTime" msgpack:"startTime" bson:"starttime" mapstructure:"startTime,omitempty"`

	// Local subnets of this enforcer.
	Subnets []string `json:"subnets" msgpack:"subnets" bson:"subnets" mapstructure:"subnets,omitempty"`

	// The Aporeto control plane sets this value to `true` if it hasn't heard from
	// the enforcer in the last five minutes.
	Unreachable bool `json:"unreachable" msgpack:"unreachable" bson:"unreachable" mapstructure:"unreachable,omitempty"`

	// If `true`, the enforcer version is outdated and should be updated.
	UpdateAvailable bool `json:"updateAvailable" msgpack:"updateAvailable" bson:"updateavailable" mapstructure:"updateAvailable,omitempty"`

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

// NewEnforcer returns a new *Enforcer
func NewEnforcer() *Enforcer {

	return &Enforcer{
		ModelVersion:          1,
		Annotations:           map[string][]string{},
		AssociatedTags:        []string{},
		CollectedInfo:         map[string]string{},
		EnforcementStatus:     EnforcerEnforcementStatusInactive,
		NormalizedTags:        []string{},
		OperationalStatus:     EnforcerOperationalStatusRegistered,
		LogLevelDuration:      "10s",
		Metadata:              []string{},
		LogLevel:              EnforcerLogLevelInfo,
		Subnets:               []string{},
		LastValidHostServices: HostServicesList{},
		MigrationsLog:         map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *Enforcer) Identity() elemental.Identity {

	return EnforcerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Enforcer) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Enforcer) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Enforcer) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesEnforcer{}

	s.FQDN = o.FQDN
	s.ID = bson.ObjectIdHex(o.ID)
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.Certificate = o.Certificate
	s.CollectInfo = o.CollectInfo
	s.CollectedInfo = o.CollectedInfo
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.CurrentVersion = o.CurrentVersion
	s.Description = o.Description
	s.EnforcementStatus = o.EnforcementStatus
	s.LastCollectionTime = o.LastCollectionTime
	s.LastPokeTime = o.LastPokeTime
	s.LastSyncTime = o.LastSyncTime
	s.LastValidHostServices = o.LastValidHostServices
	s.LogLevel = o.LogLevel
	s.LogLevelDuration = o.LogLevelDuration
	s.MachineID = o.MachineID
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.OperationalStatus = o.OperationalStatus
	s.Protected = o.Protected
	s.PublicToken = o.PublicToken
	s.StartTime = o.StartTime
	s.Subnets = o.Subnets
	s.Unreachable = o.Unreachable
	s.UpdateAvailable = o.UpdateAvailable
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Enforcer) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesEnforcer{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.FQDN = s.FQDN
	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.Certificate = s.Certificate
	o.CollectInfo = s.CollectInfo
	o.CollectedInfo = s.CollectedInfo
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.CurrentVersion = s.CurrentVersion
	o.Description = s.Description
	o.EnforcementStatus = s.EnforcementStatus
	o.LastCollectionTime = s.LastCollectionTime
	o.LastPokeTime = s.LastPokeTime
	o.LastSyncTime = s.LastSyncTime
	o.LastValidHostServices = s.LastValidHostServices
	o.LogLevel = s.LogLevel
	o.LogLevelDuration = s.LogLevelDuration
	o.MachineID = s.MachineID
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.OperationalStatus = s.OperationalStatus
	o.Protected = s.Protected
	o.PublicToken = s.PublicToken
	o.StartTime = s.StartTime
	o.Subnets = s.Subnets
	o.Unreachable = s.Unreachable
	o.UpdateAvailable = s.UpdateAvailable
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Enforcer) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Enforcer) BleveType() string {

	return "enforcer"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Enforcer) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Enforcer) Doc() string {

	return `Contains all parameters associated with a registered enforcer. The
object is mainly maintained by the enforcers themselves. Users can read the
object in order to understand the current status of the enforcers.`
}

func (o *Enforcer) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Enforcer) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *Enforcer) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Enforcer) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *Enforcer) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *Enforcer) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *Enforcer) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Enforcer) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Enforcer) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *Enforcer) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *Enforcer) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *Enforcer) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *Enforcer) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *Enforcer) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *Enforcer) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *Enforcer) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Enforcer) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Enforcer) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Enforcer) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Enforcer) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *Enforcer) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Enforcer) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *Enforcer) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *Enforcer) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *Enforcer) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Enforcer) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Enforcer) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *Enforcer) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Enforcer) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Enforcer) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Enforcer) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Enforcer) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEnforcer{
			FQDN:                      &o.FQDN,
			ID:                        &o.ID,
			Annotations:               &o.Annotations,
			AssociatedTags:            &o.AssociatedTags,
			Certificate:               &o.Certificate,
			CertificateExpirationDate: &o.CertificateExpirationDate,
			CertificateKey:            &o.CertificateKey,
			CertificateRequest:        &o.CertificateRequest,
			CollectInfo:               &o.CollectInfo,
			CollectedInfo:             &o.CollectedInfo,
			CreateIdempotencyKey:      &o.CreateIdempotencyKey,
			CreateTime:                &o.CreateTime,
			CurrentVersion:            &o.CurrentVersion,
			Description:               &o.Description,
			EnforcementStatus:         &o.EnforcementStatus,
			LastCollectionTime:        &o.LastCollectionTime,
			LastPokeTime:              &o.LastPokeTime,
			LastSyncTime:              &o.LastSyncTime,
			LastValidHostServices:     &o.LastValidHostServices,
			LocalCA:                   &o.LocalCA,
			LogLevel:                  &o.LogLevel,
			LogLevelDuration:          &o.LogLevelDuration,
			MachineID:                 &o.MachineID,
			Metadata:                  &o.Metadata,
			MigrationsLog:             &o.MigrationsLog,
			Name:                      &o.Name,
			Namespace:                 &o.Namespace,
			NormalizedTags:            &o.NormalizedTags,
			OperationalStatus:         &o.OperationalStatus,
			Protected:                 &o.Protected,
			PublicToken:               &o.PublicToken,
			StartTime:                 &o.StartTime,
			Subnets:                   &o.Subnets,
			Unreachable:               &o.Unreachable,
			UpdateAvailable:           &o.UpdateAvailable,
			UpdateIdempotencyKey:      &o.UpdateIdempotencyKey,
			UpdateTime:                &o.UpdateTime,
			ZHash:                     &o.ZHash,
			Zone:                      &o.Zone,
		}
	}

	sp := &SparseEnforcer{}
	for _, f := range fields {
		switch f {
		case "FQDN":
			sp.FQDN = &(o.FQDN)
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "certificateExpirationDate":
			sp.CertificateExpirationDate = &(o.CertificateExpirationDate)
		case "certificateKey":
			sp.CertificateKey = &(o.CertificateKey)
		case "certificateRequest":
			sp.CertificateRequest = &(o.CertificateRequest)
		case "collectInfo":
			sp.CollectInfo = &(o.CollectInfo)
		case "collectedInfo":
			sp.CollectedInfo = &(o.CollectedInfo)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "currentVersion":
			sp.CurrentVersion = &(o.CurrentVersion)
		case "description":
			sp.Description = &(o.Description)
		case "enforcementStatus":
			sp.EnforcementStatus = &(o.EnforcementStatus)
		case "lastCollectionTime":
			sp.LastCollectionTime = &(o.LastCollectionTime)
		case "lastPokeTime":
			sp.LastPokeTime = &(o.LastPokeTime)
		case "lastSyncTime":
			sp.LastSyncTime = &(o.LastSyncTime)
		case "lastValidHostServices":
			sp.LastValidHostServices = &(o.LastValidHostServices)
		case "localCA":
			sp.LocalCA = &(o.LocalCA)
		case "logLevel":
			sp.LogLevel = &(o.LogLevel)
		case "logLevelDuration":
			sp.LogLevelDuration = &(o.LogLevelDuration)
		case "machineID":
			sp.MachineID = &(o.MachineID)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "operationalStatus":
			sp.OperationalStatus = &(o.OperationalStatus)
		case "protected":
			sp.Protected = &(o.Protected)
		case "publicToken":
			sp.PublicToken = &(o.PublicToken)
		case "startTime":
			sp.StartTime = &(o.StartTime)
		case "subnets":
			sp.Subnets = &(o.Subnets)
		case "unreachable":
			sp.Unreachable = &(o.Unreachable)
		case "updateAvailable":
			sp.UpdateAvailable = &(o.UpdateAvailable)
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

// Patch apply the non nil value of a *SparseEnforcer to the object.
func (o *Enforcer) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEnforcer)
	if so.FQDN != nil {
		o.FQDN = *so.FQDN
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.CertificateExpirationDate != nil {
		o.CertificateExpirationDate = *so.CertificateExpirationDate
	}
	if so.CertificateKey != nil {
		o.CertificateKey = *so.CertificateKey
	}
	if so.CertificateRequest != nil {
		o.CertificateRequest = *so.CertificateRequest
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
	if so.CurrentVersion != nil {
		o.CurrentVersion = *so.CurrentVersion
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.EnforcementStatus != nil {
		o.EnforcementStatus = *so.EnforcementStatus
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
	if so.LastValidHostServices != nil {
		o.LastValidHostServices = *so.LastValidHostServices
	}
	if so.LocalCA != nil {
		o.LocalCA = *so.LocalCA
	}
	if so.LogLevel != nil {
		o.LogLevel = *so.LogLevel
	}
	if so.LogLevelDuration != nil {
		o.LogLevelDuration = *so.LogLevelDuration
	}
	if so.MachineID != nil {
		o.MachineID = *so.MachineID
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
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.OperationalStatus != nil {
		o.OperationalStatus = *so.OperationalStatus
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.PublicToken != nil {
		o.PublicToken = *so.PublicToken
	}
	if so.StartTime != nil {
		o.StartTime = *so.StartTime
	}
	if so.Subnets != nil {
		o.Subnets = *so.Subnets
	}
	if so.Unreachable != nil {
		o.Unreachable = *so.Unreachable
	}
	if so.UpdateAvailable != nil {
		o.UpdateAvailable = *so.UpdateAvailable
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

// DeepCopy returns a deep copy if the Enforcer.
func (o *Enforcer) DeepCopy() *Enforcer {

	if o == nil {
		return nil
	}

	out := &Enforcer{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Enforcer.
func (o *Enforcer) DeepCopyInto(out *Enforcer) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Enforcer: %s", err))
	}

	*out = *target.(*Enforcer)
}

// Validate valides the current information stored into the structure.
func (o *Enforcer) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("FQDN", o.FQDN); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("enforcementStatus", string(o.EnforcementStatus), []string{"Inactive", "Active", "Failed"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("logLevel", string(o.LogLevel), []string{"Info", "Debug", "Warn", "Error", "Trace"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTimeDuration("logLevelDuration", o.LogLevelDuration); err != nil {
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

	if err := elemental.ValidateStringInList("operationalStatus", string(o.OperationalStatus), []string{"Registered", "Connected", "Disconnected", "Initialized"}, false); err != nil {
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
func (*Enforcer) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Enforcer) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Enforcer) ValueForAttribute(name string) interface{} {

	switch name {
	case "FQDN":
		return o.FQDN
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "certificate":
		return o.Certificate
	case "certificateExpirationDate":
		return o.CertificateExpirationDate
	case "certificateKey":
		return o.CertificateKey
	case "certificateRequest":
		return o.CertificateRequest
	case "collectInfo":
		return o.CollectInfo
	case "collectedInfo":
		return o.CollectedInfo
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "currentVersion":
		return o.CurrentVersion
	case "description":
		return o.Description
	case "enforcementStatus":
		return o.EnforcementStatus
	case "lastCollectionTime":
		return o.LastCollectionTime
	case "lastPokeTime":
		return o.LastPokeTime
	case "lastSyncTime":
		return o.LastSyncTime
	case "lastValidHostServices":
		return o.LastValidHostServices
	case "localCA":
		return o.LocalCA
	case "logLevel":
		return o.LogLevel
	case "logLevelDuration":
		return o.LogLevelDuration
	case "machineID":
		return o.MachineID
	case "metadata":
		return o.Metadata
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "operationalStatus":
		return o.OperationalStatus
	case "protected":
		return o.Protected
	case "publicToken":
		return o.PublicToken
	case "startTime":
		return o.StartTime
	case "subnets":
		return o.Subnets
	case "unreachable":
		return o.Unreachable
	case "updateAvailable":
		return o.UpdateAvailable
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

// EnforcerAttributesMap represents the map of attribute for Enforcer.
var EnforcerAttributesMap = map[string]elemental.AttributeSpecification{
	"FQDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FQDN",
		CreationOnly:   true,
		Description: `Contains the fully qualified domain name (FQDN) of the server where the
enforcer is running.`,
		Exposed:   true,
		Name:      "FQDN",
		Orderable: true,
		Required:  true,
		Stored:    true,
		Type:      "string",
	},
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
	"Certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `The certificate of the enforcer.`,
		Exposed:        true,
		Name:           "certificate",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CertificateExpirationDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CertificateExpirationDate",
		Description: `The expiration date of the certificate. This is an
internal attribute, not exposed in the API.`,
		Name:     "certificateExpirationDate",
		ReadOnly: true,
		Type:     "time",
	},
	"CertificateKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CertificateKey",
		Description: `The certificate key of the enforcer. This is an internal
attribute, not exposed in the API.`,
		Name:     "certificateKey",
		ReadOnly: true,
		Type:     "string",
	},
	"CertificateRequest": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateRequest",
		Description: `If not empty during a create or update operation, the provided certificate
signing request (CSR) will be validated and signed by the control plane,
providing a renewed certificate.`,
		Exposed:   true,
		Name:      "certificateRequest",
		Transient: true,
		Type:      "string",
	},
	"CollectInfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectInfo",
		Description:    `Indicates to the enforcer whether or not it needs to collect information.`,
		Exposed:        true,
		Name:           "collectInfo",
		Stored:         true,
		Type:           "boolean",
	},
	"CollectedInfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectedInfo",
		Description:    `Represents the latest information collected by the enforcer.`,
		Exposed:        true,
		Name:           "collectedInfo",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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
	"CurrentVersion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CurrentVersion",
		Description:    `The version number of the installed enforcer binary.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "currentVersion",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
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
		AllowedChoices: []string{"Inactive", "Active", "Failed"},
		ConvertedName:  "EnforcementStatus",
		DefaultValue:   EnforcerEnforcementStatusInactive,
		Description:    `Status of the enforcement for host services.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcementStatus",
		Stored:         true,
		Type:           "enum",
	},
	"LastCollectionTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastCollectionTime",
		Description:    `Identifies when the information was collected.`,
		Exposed:        true,
		Name:           "lastCollectionTime",
		Stored:         true,
		Type:           "time",
	},
	"LastPokeTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastPokeTime",
		Description:    `The time and date of the last poke.`,
		Name:           "lastPokeTime",
		Stored:         true,
		Type:           "time",
	},
	"LastSyncTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastSyncTime",
		Description:    `The time and date of the last heartbeat.`,
		Exposed:        true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"LastValidHostServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastValidHostServices",
		Description: `LastValidHostServices is a read only attribute that stores the list valid host
services that have been applied to this enforcer. This list might be different
from the list retrieved through policy, if the dynamically calculated list leads
into conflicts.`,
		Name:    "lastValidHostServices",
		Stored:  true,
		SubType: "hostservice",
		Type:    "refList",
	},
	"LocalCA": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LocalCA",
		Description: `Contains the initial chain of trust for the enforcer. This value is only
given when you retrieve a single enforcer.`,
		Exposed:   true,
		Name:      "localCA",
		Transient: true,
		Type:      "string",
	},
	"LogLevel": elemental.AttributeSpecification{
		AllowedChoices: []string{"Info", "Debug", "Warn", "Error", "Trace"},
		ConvertedName:  "LogLevel",
		DefaultValue:   EnforcerLogLevelInfo,
		Description:    `Log level of the enforcer.`,
		Exposed:        true,
		Name:           "logLevel",
		Stored:         true,
		Type:           "enum",
	},
	"LogLevelDuration": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LogLevelDuration",
		DefaultValue:   "10s",
		Description: `Determines the duration of which the log level will be active, using [Golang
duration syntax](https://golang.org/pkg/time/#example_Duration).`,
		Exposed: true,
		Name:    "logLevelDuration",
		Stored:  true,
		Type:    "string",
	},
	"MachineID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MachineID",
		Description: `A unique identifier for every machine as detected by the enforcer. It is
based on hardware information such as the SMBIOS UUID, MAC addresses of
interfaces, or cloud provider IDs.`,
		Exposed:    true,
		Filterable: true,
		Name:       "machineID",
		Stored:     true,
		Type:       "string",
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
		AllowedChoices: []string{"Registered", "Connected", "Disconnected", "Initialized"},
		ConvertedName:  "OperationalStatus",
		DefaultValue:   EnforcerOperationalStatusRegistered,
		Description:    `The status of the enforcer.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		Stored:         true,
		Type:           "enum",
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
	"PublicToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PublicToken",
		Description: `The public token of the server that will be included in the datapath and
is signed by the private certificate authority.`,
		Exposed:   true,
		Name:      "publicToken",
		ReadOnly:  true,
		Stored:    true,
		Transient: true,
		Type:      "string",
	},
	"StartTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "StartTime",
		Description: `The time and date on which this enforcer was started. The enforcer reports
this and the value is preserved across disconnects.`,
		Exposed:   true,
		Name:      "startTime",
		Orderable: true,
		Stored:    true,
		Type:      "time",
	},
	"Subnets": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subnets",
		Description:    `Local subnets of this enforcer.`,
		Exposed:        true,
		Name:           "subnets",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Unreachable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Unreachable",
		Description: `The Aporeto control plane sets this value to ` + "`" + `true` + "`" + ` if it hasn't heard from
the enforcer in the last five minutes.`,
		Exposed:  true,
		Name:     "unreachable",
		ReadOnly: true,
		Stored:   true,
		Type:     "boolean",
	},
	"UpdateAvailable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "UpdateAvailable",
		Description:    `If ` + "`" + `true` + "`" + `, the enforcer version is outdated and should be updated.`,
		Exposed:        true,
		Name:           "updateAvailable",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
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

// EnforcerLowerCaseAttributesMap represents the map of attribute for Enforcer.
var EnforcerLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"fqdn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FQDN",
		CreationOnly:   true,
		Description: `Contains the fully qualified domain name (FQDN) of the server where the
enforcer is running.`,
		Exposed:   true,
		Name:      "FQDN",
		Orderable: true,
		Required:  true,
		Stored:    true,
		Type:      "string",
	},
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
	"certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `The certificate of the enforcer.`,
		Exposed:        true,
		Name:           "certificate",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"certificateexpirationdate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CertificateExpirationDate",
		Description: `The expiration date of the certificate. This is an
internal attribute, not exposed in the API.`,
		Name:     "certificateExpirationDate",
		ReadOnly: true,
		Type:     "time",
	},
	"certificatekey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CertificateKey",
		Description: `The certificate key of the enforcer. This is an internal
attribute, not exposed in the API.`,
		Name:     "certificateKey",
		ReadOnly: true,
		Type:     "string",
	},
	"certificaterequest": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateRequest",
		Description: `If not empty during a create or update operation, the provided certificate
signing request (CSR) will be validated and signed by the control plane,
providing a renewed certificate.`,
		Exposed:   true,
		Name:      "certificateRequest",
		Transient: true,
		Type:      "string",
	},
	"collectinfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectInfo",
		Description:    `Indicates to the enforcer whether or not it needs to collect information.`,
		Exposed:        true,
		Name:           "collectInfo",
		Stored:         true,
		Type:           "boolean",
	},
	"collectedinfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectedInfo",
		Description:    `Represents the latest information collected by the enforcer.`,
		Exposed:        true,
		Name:           "collectedInfo",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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
	"currentversion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CurrentVersion",
		Description:    `The version number of the installed enforcer binary.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "currentVersion",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
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
		AllowedChoices: []string{"Inactive", "Active", "Failed"},
		ConvertedName:  "EnforcementStatus",
		DefaultValue:   EnforcerEnforcementStatusInactive,
		Description:    `Status of the enforcement for host services.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcementStatus",
		Stored:         true,
		Type:           "enum",
	},
	"lastcollectiontime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastCollectionTime",
		Description:    `Identifies when the information was collected.`,
		Exposed:        true,
		Name:           "lastCollectionTime",
		Stored:         true,
		Type:           "time",
	},
	"lastpoketime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastPokeTime",
		Description:    `The time and date of the last poke.`,
		Name:           "lastPokeTime",
		Stored:         true,
		Type:           "time",
	},
	"lastsynctime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastSyncTime",
		Description:    `The time and date of the last heartbeat.`,
		Exposed:        true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"lastvalidhostservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastValidHostServices",
		Description: `LastValidHostServices is a read only attribute that stores the list valid host
services that have been applied to this enforcer. This list might be different
from the list retrieved through policy, if the dynamically calculated list leads
into conflicts.`,
		Name:    "lastValidHostServices",
		Stored:  true,
		SubType: "hostservice",
		Type:    "refList",
	},
	"localca": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LocalCA",
		Description: `Contains the initial chain of trust for the enforcer. This value is only
given when you retrieve a single enforcer.`,
		Exposed:   true,
		Name:      "localCA",
		Transient: true,
		Type:      "string",
	},
	"loglevel": elemental.AttributeSpecification{
		AllowedChoices: []string{"Info", "Debug", "Warn", "Error", "Trace"},
		ConvertedName:  "LogLevel",
		DefaultValue:   EnforcerLogLevelInfo,
		Description:    `Log level of the enforcer.`,
		Exposed:        true,
		Name:           "logLevel",
		Stored:         true,
		Type:           "enum",
	},
	"loglevelduration": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LogLevelDuration",
		DefaultValue:   "10s",
		Description: `Determines the duration of which the log level will be active, using [Golang
duration syntax](https://golang.org/pkg/time/#example_Duration).`,
		Exposed: true,
		Name:    "logLevelDuration",
		Stored:  true,
		Type:    "string",
	},
	"machineid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MachineID",
		Description: `A unique identifier for every machine as detected by the enforcer. It is
based on hardware information such as the SMBIOS UUID, MAC addresses of
interfaces, or cloud provider IDs.`,
		Exposed:    true,
		Filterable: true,
		Name:       "machineID",
		Stored:     true,
		Type:       "string",
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
		AllowedChoices: []string{"Registered", "Connected", "Disconnected", "Initialized"},
		ConvertedName:  "OperationalStatus",
		DefaultValue:   EnforcerOperationalStatusRegistered,
		Description:    `The status of the enforcer.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		Stored:         true,
		Type:           "enum",
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
	"publictoken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PublicToken",
		Description: `The public token of the server that will be included in the datapath and
is signed by the private certificate authority.`,
		Exposed:   true,
		Name:      "publicToken",
		ReadOnly:  true,
		Stored:    true,
		Transient: true,
		Type:      "string",
	},
	"starttime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "StartTime",
		Description: `The time and date on which this enforcer was started. The enforcer reports
this and the value is preserved across disconnects.`,
		Exposed:   true,
		Name:      "startTime",
		Orderable: true,
		Stored:    true,
		Type:      "time",
	},
	"subnets": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subnets",
		Description:    `Local subnets of this enforcer.`,
		Exposed:        true,
		Name:           "subnets",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"unreachable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Unreachable",
		Description: `The Aporeto control plane sets this value to ` + "`" + `true` + "`" + ` if it hasn't heard from
the enforcer in the last five minutes.`,
		Exposed:  true,
		Name:     "unreachable",
		ReadOnly: true,
		Stored:   true,
		Type:     "boolean",
	},
	"updateavailable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "UpdateAvailable",
		Description:    `If ` + "`" + `true` + "`" + `, the enforcer version is outdated and should be updated.`,
		Exposed:        true,
		Name:           "updateAvailable",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
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

// SparseEnforcersList represents a list of SparseEnforcers
type SparseEnforcersList []*SparseEnforcer

// Identity returns the identity of the objects in the list.
func (o SparseEnforcersList) Identity() elemental.Identity {

	return EnforcerIdentity
}

// Copy returns a pointer to a copy the SparseEnforcersList.
func (o SparseEnforcersList) Copy() elemental.Identifiables {

	copy := append(SparseEnforcersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEnforcersList.
func (o SparseEnforcersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEnforcersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEnforcer))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEnforcersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEnforcersList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseEnforcersList converted to EnforcersList.
func (o SparseEnforcersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEnforcersList) Version() int {

	return 1
}

// SparseEnforcer represents the sparse version of a enforcer.
type SparseEnforcer struct {
	// Contains the fully qualified domain name (FQDN) of the server where the
	// enforcer is running.
	FQDN *string `json:"FQDN,omitempty" msgpack:"FQDN,omitempty" bson:"fqdn,omitempty" mapstructure:"FQDN,omitempty"`

	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// The certificate of the enforcer.
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"certificate,omitempty" mapstructure:"certificate,omitempty"`

	// The expiration date of the certificate. This is an
	// internal attribute, not exposed in the API.
	CertificateExpirationDate *time.Time `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	// The certificate key of the enforcer. This is an internal
	// attribute, not exposed in the API.
	CertificateKey *string `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	// If not empty during a create or update operation, the provided certificate
	// signing request (CSR) will be validated and signed by the control plane,
	// providing a renewed certificate.
	CertificateRequest *string `json:"certificateRequest,omitempty" msgpack:"certificateRequest,omitempty" bson:"-" mapstructure:"certificateRequest,omitempty"`

	// Indicates to the enforcer whether or not it needs to collect information.
	CollectInfo *bool `json:"collectInfo,omitempty" msgpack:"collectInfo,omitempty" bson:"collectinfo,omitempty" mapstructure:"collectInfo,omitempty"`

	// Represents the latest information collected by the enforcer.
	CollectedInfo *map[string]string `json:"collectedInfo,omitempty" msgpack:"collectedInfo,omitempty" bson:"collectedinfo,omitempty" mapstructure:"collectedInfo,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// The version number of the installed enforcer binary.
	CurrentVersion *string `json:"currentVersion,omitempty" msgpack:"currentVersion,omitempty" bson:"currentversion,omitempty" mapstructure:"currentVersion,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Status of the enforcement for host services.
	EnforcementStatus *EnforcerEnforcementStatusValue `json:"enforcementStatus,omitempty" msgpack:"enforcementStatus,omitempty" bson:"enforcementstatus,omitempty" mapstructure:"enforcementStatus,omitempty"`

	// Identifies when the information was collected.
	LastCollectionTime *time.Time `json:"lastCollectionTime,omitempty" msgpack:"lastCollectionTime,omitempty" bson:"lastcollectiontime,omitempty" mapstructure:"lastCollectionTime,omitempty"`

	// The time and date of the last poke.
	LastPokeTime *time.Time `json:"-" msgpack:"-" bson:"lastpoketime,omitempty" mapstructure:"-,omitempty"`

	// The time and date of the last heartbeat.
	LastSyncTime *time.Time `json:"lastSyncTime,omitempty" msgpack:"lastSyncTime,omitempty" bson:"lastsynctime,omitempty" mapstructure:"lastSyncTime,omitempty"`

	// LastValidHostServices is a read only attribute that stores the list valid host
	// services that have been applied to this enforcer. This list might be different
	// from the list retrieved through policy, if the dynamically calculated list leads
	// into conflicts.
	LastValidHostServices *HostServicesList `json:"-" msgpack:"-" bson:"lastvalidhostservices,omitempty" mapstructure:"-,omitempty"`

	// Contains the initial chain of trust for the enforcer. This value is only
	// given when you retrieve a single enforcer.
	LocalCA *string `json:"localCA,omitempty" msgpack:"localCA,omitempty" bson:"-" mapstructure:"localCA,omitempty"`

	// Log level of the enforcer.
	LogLevel *EnforcerLogLevelValue `json:"logLevel,omitempty" msgpack:"logLevel,omitempty" bson:"loglevel,omitempty" mapstructure:"logLevel,omitempty"`

	// Determines the duration of which the log level will be active, using [Golang
	// duration syntax](https://golang.org/pkg/time/#example_Duration).
	LogLevelDuration *string `json:"logLevelDuration,omitempty" msgpack:"logLevelDuration,omitempty" bson:"loglevelduration,omitempty" mapstructure:"logLevelDuration,omitempty"`

	// A unique identifier for every machine as detected by the enforcer. It is
	// based on hardware information such as the SMBIOS UUID, MAC addresses of
	// interfaces, or cloud provider IDs.
	MachineID *string `json:"machineID,omitempty" msgpack:"machineID,omitempty" bson:"machineid,omitempty" mapstructure:"machineID,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// The status of the enforcer.
	OperationalStatus *EnforcerOperationalStatusValue `json:"operationalStatus,omitempty" msgpack:"operationalStatus,omitempty" bson:"operationalstatus,omitempty" mapstructure:"operationalStatus,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// The public token of the server that will be included in the datapath and
	// is signed by the private certificate authority.
	PublicToken *string `json:"publicToken,omitempty" msgpack:"publicToken,omitempty" bson:"publictoken,omitempty" mapstructure:"publicToken,omitempty"`

	// The time and date on which this enforcer was started. The enforcer reports
	// this and the value is preserved across disconnects.
	StartTime *time.Time `json:"startTime,omitempty" msgpack:"startTime,omitempty" bson:"starttime,omitempty" mapstructure:"startTime,omitempty"`

	// Local subnets of this enforcer.
	Subnets *[]string `json:"subnets,omitempty" msgpack:"subnets,omitempty" bson:"subnets,omitempty" mapstructure:"subnets,omitempty"`

	// The Aporeto control plane sets this value to `true` if it hasn't heard from
	// the enforcer in the last five minutes.
	Unreachable *bool `json:"unreachable,omitempty" msgpack:"unreachable,omitempty" bson:"unreachable,omitempty" mapstructure:"unreachable,omitempty"`

	// If `true`, the enforcer version is outdated and should be updated.
	UpdateAvailable *bool `json:"updateAvailable,omitempty" msgpack:"updateAvailable,omitempty" bson:"updateavailable,omitempty" mapstructure:"updateAvailable,omitempty"`

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

// NewSparseEnforcer returns a new  SparseEnforcer.
func NewSparseEnforcer() *SparseEnforcer {
	return &SparseEnforcer{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEnforcer) Identity() elemental.Identity {

	return EnforcerIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEnforcer) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEnforcer) SetIdentifier(id string) {

	o.ID = &id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEnforcer) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseEnforcer{}

	if o.FQDN != nil {
		s.FQDN = o.FQDN
	}
	s.ID = bson.ObjectIdHex(*o.ID)
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.Certificate != nil {
		s.Certificate = o.Certificate
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
	if o.CurrentVersion != nil {
		s.CurrentVersion = o.CurrentVersion
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.EnforcementStatus != nil {
		s.EnforcementStatus = o.EnforcementStatus
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
	if o.LastValidHostServices != nil {
		s.LastValidHostServices = o.LastValidHostServices
	}
	if o.LogLevel != nil {
		s.LogLevel = o.LogLevel
	}
	if o.LogLevelDuration != nil {
		s.LogLevelDuration = o.LogLevelDuration
	}
	if o.MachineID != nil {
		s.MachineID = o.MachineID
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
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.OperationalStatus != nil {
		s.OperationalStatus = o.OperationalStatus
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.PublicToken != nil {
		s.PublicToken = o.PublicToken
	}
	if o.StartTime != nil {
		s.StartTime = o.StartTime
	}
	if o.Subnets != nil {
		s.Subnets = o.Subnets
	}
	if o.Unreachable != nil {
		s.Unreachable = o.Unreachable
	}
	if o.UpdateAvailable != nil {
		s.UpdateAvailable = o.UpdateAvailable
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
func (o *SparseEnforcer) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseEnforcer{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.FQDN != nil {
		o.FQDN = s.FQDN
	}
	id := s.ID.Hex()
	o.ID = &id
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.Certificate != nil {
		o.Certificate = s.Certificate
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
	if s.CurrentVersion != nil {
		o.CurrentVersion = s.CurrentVersion
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.EnforcementStatus != nil {
		o.EnforcementStatus = s.EnforcementStatus
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
	if s.LastValidHostServices != nil {
		o.LastValidHostServices = s.LastValidHostServices
	}
	if s.LogLevel != nil {
		o.LogLevel = s.LogLevel
	}
	if s.LogLevelDuration != nil {
		o.LogLevelDuration = s.LogLevelDuration
	}
	if s.MachineID != nil {
		o.MachineID = s.MachineID
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
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.OperationalStatus != nil {
		o.OperationalStatus = s.OperationalStatus
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.PublicToken != nil {
		o.PublicToken = s.PublicToken
	}
	if s.StartTime != nil {
		o.StartTime = s.StartTime
	}
	if s.Subnets != nil {
		o.Subnets = s.Subnets
	}
	if s.Unreachable != nil {
		o.Unreachable = s.Unreachable
	}
	if s.UpdateAvailable != nil {
		o.UpdateAvailable = s.UpdateAvailable
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
func (o *SparseEnforcer) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEnforcer) ToPlain() elemental.PlainIdentifiable {

	out := NewEnforcer()
	if o.FQDN != nil {
		out.FQDN = *o.FQDN
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.CertificateExpirationDate != nil {
		out.CertificateExpirationDate = *o.CertificateExpirationDate
	}
	if o.CertificateKey != nil {
		out.CertificateKey = *o.CertificateKey
	}
	if o.CertificateRequest != nil {
		out.CertificateRequest = *o.CertificateRequest
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
	if o.CurrentVersion != nil {
		out.CurrentVersion = *o.CurrentVersion
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.EnforcementStatus != nil {
		out.EnforcementStatus = *o.EnforcementStatus
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
	if o.LastValidHostServices != nil {
		out.LastValidHostServices = *o.LastValidHostServices
	}
	if o.LocalCA != nil {
		out.LocalCA = *o.LocalCA
	}
	if o.LogLevel != nil {
		out.LogLevel = *o.LogLevel
	}
	if o.LogLevelDuration != nil {
		out.LogLevelDuration = *o.LogLevelDuration
	}
	if o.MachineID != nil {
		out.MachineID = *o.MachineID
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
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.OperationalStatus != nil {
		out.OperationalStatus = *o.OperationalStatus
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.PublicToken != nil {
		out.PublicToken = *o.PublicToken
	}
	if o.StartTime != nil {
		out.StartTime = *o.StartTime
	}
	if o.Subnets != nil {
		out.Subnets = *o.Subnets
	}
	if o.Unreachable != nil {
		out.Unreachable = *o.Unreachable
	}
	if o.UpdateAvailable != nil {
		out.UpdateAvailable = *o.UpdateAvailable
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
func (o *SparseEnforcer) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseEnforcer) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseEnforcer) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseEnforcer) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseEnforcer) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseEnforcer) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseEnforcer) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseEnforcer) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseEnforcer) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseEnforcer) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseEnforcer) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseEnforcer) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseEnforcer) GetMigrationsLog() map[string]string {

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseEnforcer) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseEnforcer) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseEnforcer) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseEnforcer) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseEnforcer) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseEnforcer) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseEnforcer) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseEnforcer) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseEnforcer) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseEnforcer) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseEnforcer) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseEnforcer) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseEnforcer) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseEnforcer) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseEnforcer) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseEnforcer) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseEnforcer) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseEnforcer.
func (o *SparseEnforcer) DeepCopy() *SparseEnforcer {

	if o == nil {
		return nil
	}

	out := &SparseEnforcer{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEnforcer.
func (o *SparseEnforcer) DeepCopyInto(out *SparseEnforcer) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEnforcer: %s", err))
	}

	*out = *target.(*SparseEnforcer)
}

type mongoAttributesEnforcer struct {
	FQDN                  string                         `bson:"fqdn"`
	ID                    bson.ObjectId                  `bson:"_id"`
	Annotations           map[string][]string            `bson:"annotations"`
	AssociatedTags        []string                       `bson:"associatedtags"`
	Certificate           string                         `bson:"certificate"`
	CollectInfo           bool                           `bson:"collectinfo"`
	CollectedInfo         map[string]string              `bson:"collectedinfo"`
	CreateIdempotencyKey  string                         `bson:"createidempotencykey"`
	CreateTime            time.Time                      `bson:"createtime"`
	CurrentVersion        string                         `bson:"currentversion"`
	Description           string                         `bson:"description"`
	EnforcementStatus     EnforcerEnforcementStatusValue `bson:"enforcementstatus"`
	LastCollectionTime    time.Time                      `bson:"lastcollectiontime"`
	LastPokeTime          time.Time                      `bson:"lastpoketime"`
	LastSyncTime          time.Time                      `bson:"lastsynctime"`
	LastValidHostServices HostServicesList               `bson:"lastvalidhostservices"`
	LogLevel              EnforcerLogLevelValue          `bson:"loglevel"`
	LogLevelDuration      string                         `bson:"loglevelduration"`
	MachineID             string                         `bson:"machineid"`
	Metadata              []string                       `bson:"metadata"`
	MigrationsLog         map[string]string              `bson:"migrationslog"`
	Name                  string                         `bson:"name"`
	Namespace             string                         `bson:"namespace"`
	NormalizedTags        []string                       `bson:"normalizedtags"`
	OperationalStatus     EnforcerOperationalStatusValue `bson:"operationalstatus"`
	Protected             bool                           `bson:"protected"`
	PublicToken           string                         `bson:"publictoken"`
	StartTime             time.Time                      `bson:"starttime"`
	Subnets               []string                       `bson:"subnets"`
	Unreachable           bool                           `bson:"unreachable"`
	UpdateAvailable       bool                           `bson:"updateavailable"`
	UpdateIdempotencyKey  string                         `bson:"updateidempotencykey"`
	UpdateTime            time.Time                      `bson:"updatetime"`
	ZHash                 int                            `bson:"zhash"`
	Zone                  int                            `bson:"zone"`
}
type mongoAttributesSparseEnforcer struct {
	FQDN                  *string                         `bson:"fqdn,omitempty"`
	ID                    bson.ObjectId                   `bson:"_id"`
	Annotations           *map[string][]string            `bson:"annotations,omitempty"`
	AssociatedTags        *[]string                       `bson:"associatedtags,omitempty"`
	Certificate           *string                         `bson:"certificate,omitempty"`
	CollectInfo           *bool                           `bson:"collectinfo,omitempty"`
	CollectedInfo         *map[string]string              `bson:"collectedinfo,omitempty"`
	CreateIdempotencyKey  *string                         `bson:"createidempotencykey,omitempty"`
	CreateTime            *time.Time                      `bson:"createtime,omitempty"`
	CurrentVersion        *string                         `bson:"currentversion,omitempty"`
	Description           *string                         `bson:"description,omitempty"`
	EnforcementStatus     *EnforcerEnforcementStatusValue `bson:"enforcementstatus,omitempty"`
	LastCollectionTime    *time.Time                      `bson:"lastcollectiontime,omitempty"`
	LastPokeTime          *time.Time                      `bson:"lastpoketime,omitempty"`
	LastSyncTime          *time.Time                      `bson:"lastsynctime,omitempty"`
	LastValidHostServices *HostServicesList               `bson:"lastvalidhostservices,omitempty"`
	LogLevel              *EnforcerLogLevelValue          `bson:"loglevel,omitempty"`
	LogLevelDuration      *string                         `bson:"loglevelduration,omitempty"`
	MachineID             *string                         `bson:"machineid,omitempty"`
	Metadata              *[]string                       `bson:"metadata,omitempty"`
	MigrationsLog         *map[string]string              `bson:"migrationslog,omitempty"`
	Name                  *string                         `bson:"name,omitempty"`
	Namespace             *string                         `bson:"namespace,omitempty"`
	NormalizedTags        *[]string                       `bson:"normalizedtags,omitempty"`
	OperationalStatus     *EnforcerOperationalStatusValue `bson:"operationalstatus,omitempty"`
	Protected             *bool                           `bson:"protected,omitempty"`
	PublicToken           *string                         `bson:"publictoken,omitempty"`
	StartTime             *time.Time                      `bson:"starttime,omitempty"`
	Subnets               *[]string                       `bson:"subnets,omitempty"`
	Unreachable           *bool                           `bson:"unreachable,omitempty"`
	UpdateAvailable       *bool                           `bson:"updateavailable,omitempty"`
	UpdateIdempotencyKey  *string                         `bson:"updateidempotencykey,omitempty"`
	UpdateTime            *time.Time                      `bson:"updatetime,omitempty"`
	ZHash                 *int                            `bson:"zhash,omitempty"`
	Zone                  *int                            `bson:"zone,omitempty"`
}
