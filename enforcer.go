package gaia

import (
	"fmt"
	"sync"
	"time"

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

	// EnforcerOperationalStatusUnknown represents the value Unknown.
	EnforcerOperationalStatusUnknown EnforcerOperationalStatusValue = "Unknown"
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
		"namespace",
		"name",
	}
}

// ToSparse returns the EnforcersList converted to SparseEnforcersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EnforcersList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o EnforcersList) Version() int {

	return 1
}

// Enforcer represents the model of a enforcer
type Enforcer struct {
	// FQDN contains the fqdn of the server where the enforcer is running.
	FQDN string `json:"FQDN" bson:"fqdn" mapstructure:"FQDN,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Certificate is the certificate of the enforcer.
	Certificate string `json:"certificate" bson:"certificate" mapstructure:"certificate,omitempty"`

	// CertificateExpirationDate is the expiration date of the certificate. This is an
	// internal attribute, not exposed in the API.
	CertificateExpirationDate time.Time `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// CertificateKey is the certificate key of the enforcer. This is an internal
	// attribute, not exposed in the API.
	CertificateKey string `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// If not empty during a create or update generation, the provided CSR will be
	// validated and signed by the backend providing a renewed certificate.
	CertificateRequest string `json:"certificateRequest" bson:"-" mapstructure:"certificateRequest,omitempty"`

	// CollectInfo indicates to the enforcer it needs to collect information.
	CollectInfo bool `json:"collectInfo" bson:"collectinfo" mapstructure:"collectInfo,omitempty"`

	// CollectedInfo represents the latest info collected by the enforcer.
	CollectedInfo map[string]string `json:"collectedInfo" bson:"collectedinfo" mapstructure:"collectedInfo,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// CurrentVersion holds the enforcerd binary version that is currently associated
	// to this object.
	CurrentVersion string `json:"currentVersion" bson:"currentversion" mapstructure:"currentVersion,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Status of enforcement for PU managed directly by enforcerd, like Host PUs.
	EnforcementStatus EnforcerEnforcementStatusValue `json:"enforcementStatus" bson:"enforcementstatus" mapstructure:"enforcementStatus,omitempty"`

	// LastCollectionTime represents the date and time when the info have been
	// collected.
	LastCollectionTime time.Time `json:"lastCollectionTime" bson:"lastcollectiontime" mapstructure:"lastCollectionTime,omitempty"`

	// Last poke is the time when the enforcer got last poked.
	LastPokeTime time.Time `json:"-" bson:"lastpoketime" mapstructure:"-,omitempty"`

	// LastSyncTime holds the last heart beat time.
	LastSyncTime time.Time `json:"lastSyncTime" bson:"lastsynctime" mapstructure:"lastSyncTime,omitempty"`

	// LastValidHostServices is a read only attribute that stores the list valid host
	// services that have been applied to this enforcer. This list might be different
	// from the list retrieved through policy, if the dynamically calculated list leads
	// into conflicts.
	LastValidHostServices HostServicesList `json:"-" bson:"lastvalidhostservices" mapstructure:"-,omitempty"`

	// LocalCA contains the initial chain of trust for the enforcer. This valud is only
	// given when you retrieve a single enforcer.
	LocalCA string `json:"localCA" bson:"-" mapstructure:"localCA,omitempty"`

	// MachineID holds a unique identifier for every machine as detected by the
	// enforcer. It is based on hardware information such as the SMBIOS UUID, MAC
	// addresses of interfaces or cloud provider IDs.
	MachineID string `json:"machineID" bson:"machineid" mapstructure:"machineID,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// OperationalStatus tells the status of the enforcer.
	OperationalStatus EnforcerOperationalStatusValue `json:"operationalStatus" bson:"operationalstatus" mapstructure:"operationalStatus,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// PublicToken is the public token of the server that will be included in the
	// datapath and its signed by the private CA.
	PublicToken string `json:"publicToken" bson:"publictoken" mapstructure:"publicToken,omitempty"`

	// startTime holds the time this enforcerd was started. This time-stamp is reported
	// by the enforcer and is is preserved across disconnects.
	StartTime time.Time `json:"startTime" bson:"starttime" mapstructure:"startTime,omitempty"`

	// Tells if the the version of enforcerd is outdated and should be updated.
	UpdateAvailable bool `json:"updateAvailable" bson:"updateavailable" mapstructure:"updateAvailable,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewEnforcer returns a new *Enforcer
func NewEnforcer() *Enforcer {

	return &Enforcer{
		ModelVersion:          1,
		Mutex:                 &sync.Mutex{},
		Annotations:           map[string][]string{},
		AssociatedTags:        []string{},
		EnforcementStatus:     EnforcerEnforcementStatusInactive,
		CollectedInfo:         map[string]string{},
		NormalizedTags:        []string{},
		OperationalStatus:     EnforcerOperationalStatusRegistered,
		Metadata:              []string{},
		LastValidHostServices: HostServicesList{},
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

// Version returns the hardcoded version of the model.
func (o *Enforcer) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Enforcer) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Enforcer) Doc() string {
	return `An Enforcer contains all parameters associated with a registered enforcer. The
object is mainly by maintained by the enforcers themselves. Users can read the
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
			CreateTime:                &o.CreateTime,
			CurrentVersion:            &o.CurrentVersion,
			Description:               &o.Description,
			EnforcementStatus:         &o.EnforcementStatus,
			LastCollectionTime:        &o.LastCollectionTime,
			LastPokeTime:              &o.LastPokeTime,
			LastSyncTime:              &o.LastSyncTime,
			LastValidHostServices:     &o.LastValidHostServices,
			LocalCA:                   &o.LocalCA,
			MachineID:                 &o.MachineID,
			Metadata:                  &o.Metadata,
			Name:                      &o.Name,
			Namespace:                 &o.Namespace,
			NormalizedTags:            &o.NormalizedTags,
			OperationalStatus:         &o.OperationalStatus,
			Protected:                 &o.Protected,
			PublicToken:               &o.PublicToken,
			StartTime:                 &o.StartTime,
			UpdateAvailable:           &o.UpdateAvailable,
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
		case "machineID":
			sp.MachineID = &(o.MachineID)
		case "metadata":
			sp.Metadata = &(o.Metadata)
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
		case "updateAvailable":
			sp.UpdateAvailable = &(o.UpdateAvailable)
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
	if so.MachineID != nil {
		o.MachineID = *so.MachineID
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
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
	if so.UpdateAvailable != nil {
		o.UpdateAvailable = *so.UpdateAvailable
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
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("enforcementStatus", string(o.EnforcementStatus), []string{"Inactive", "Active", "Failed"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("operationalStatus", string(o.OperationalStatus), []string{"Registered", "Connected", "Disconnected", "Initialized", "Unknown"}, false); err != nil {
		errors = append(errors, err)
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
	case "machineID":
		return o.MachineID
	case "metadata":
		return o.Metadata
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
	case "updateAvailable":
		return o.UpdateAvailable
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
		Description:    `FQDN contains the fqdn of the server where the enforcer is running.`,
		Exposed:        true,
		Name:           "FQDN",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
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
		Description:    `Annotation stores additional information about an entity.`,
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
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
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
		Description:    `Certificate is the certificate of the enforcer.`,
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
		Description: `CertificateExpirationDate is the expiration date of the certificate. This is an
internal attribute, not exposed in the API.`,
		Name:     "certificateExpirationDate",
		ReadOnly: true,
		Type:     "time",
	},
	"CertificateKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CertificateKey",
		Description: `CertificateKey is the certificate key of the enforcer. This is an internal
attribute, not exposed in the API.`,
		Name:     "certificateKey",
		ReadOnly: true,
		Type:     "string",
	},
	"CertificateRequest": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateRequest",
		Description: `If not empty during a create or update generation, the provided CSR will be
validated and signed by the backend providing a renewed certificate.`,
		Exposed:   true,
		Name:      "certificateRequest",
		Transient: true,
		Type:      "string",
	},
	"CollectInfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectInfo",
		Description:    `CollectInfo indicates to the enforcer it needs to collect information.`,
		Exposed:        true,
		Name:           "collectInfo",
		Stored:         true,
		Type:           "boolean",
	},
	"CollectedInfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectedInfo",
		Description:    `CollectedInfo represents the latest info collected by the enforcer.`,
		Exposed:        true,
		Name:           "collectedInfo",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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
		Description: `CurrentVersion holds the enforcerd binary version that is currently associated
to this object.`,
		Exposed:    true,
		Filterable: true,
		Name:       "currentVersion",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
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
		Description:    `Status of enforcement for PU managed directly by enforcerd, like Host PUs.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcementStatus",
		Stored:         true,
		Type:           "enum",
	},
	"LastCollectionTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastCollectionTime",
		Description: `LastCollectionTime represents the date and time when the info have been
collected.`,
		Exposed: true,
		Name:    "lastCollectionTime",
		Stored:  true,
		Type:    "time",
	},
	"LastPokeTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastPokeTime",
		Description:    `Last poke is the time when the enforcer got last poked.`,
		Name:           "lastPokeTime",
		Stored:         true,
		Type:           "time",
	},
	"LastSyncTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastSyncTime",
		Description:    `LastSyncTime holds the last heart beat time.`,
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
		Description: `LocalCA contains the initial chain of trust for the enforcer. This valud is only
given when you retrieve a single enforcer.`,
		Exposed:   true,
		Name:      "localCA",
		Transient: true,
		Type:      "string",
	},
	"MachineID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MachineID",
		Description: `MachineID holds a unique identifier for every machine as detected by the
enforcer. It is based on hardware information such as the SMBIOS UUID, MAC
addresses of interfaces or cloud provider IDs.`,
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
		Description: `Metadata contains tags that can only be set during creation. They must all start
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
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
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
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
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
		AllowedChoices: []string{"Registered", "Connected", "Disconnected", "Initialized", "Unknown"},
		ConvertedName:  "OperationalStatus",
		DefaultValue:   EnforcerOperationalStatusRegistered,
		Description:    `OperationalStatus tells the status of the enforcer.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		Stored:         true,
		Type:           "enum",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
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
		Description: `PublicToken is the public token of the server that will be included in the
datapath and its signed by the private CA.`,
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
		Description: `startTime holds the time this enforcerd was started. This time-stamp is reported
by the enforcer and is is preserved across disconnects.`,
		Exposed:   true,
		Name:      "startTime",
		Orderable: true,
		Stored:    true,
		Type:      "time",
	},
	"UpdateAvailable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "UpdateAvailable",
		Description:    `Tells if the the version of enforcerd is outdated and should be updated.`,
		Exposed:        true,
		Name:           "updateAvailable",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
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
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// EnforcerLowerCaseAttributesMap represents the map of attribute for Enforcer.
var EnforcerLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"fqdn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FQDN",
		CreationOnly:   true,
		Description:    `FQDN contains the fqdn of the server where the enforcer is running.`,
		Exposed:        true,
		Name:           "FQDN",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
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
		Description:    `Annotation stores additional information about an entity.`,
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
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
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
		Description:    `Certificate is the certificate of the enforcer.`,
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
		Description: `CertificateExpirationDate is the expiration date of the certificate. This is an
internal attribute, not exposed in the API.`,
		Name:     "certificateExpirationDate",
		ReadOnly: true,
		Type:     "time",
	},
	"certificatekey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CertificateKey",
		Description: `CertificateKey is the certificate key of the enforcer. This is an internal
attribute, not exposed in the API.`,
		Name:     "certificateKey",
		ReadOnly: true,
		Type:     "string",
	},
	"certificaterequest": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateRequest",
		Description: `If not empty during a create or update generation, the provided CSR will be
validated and signed by the backend providing a renewed certificate.`,
		Exposed:   true,
		Name:      "certificateRequest",
		Transient: true,
		Type:      "string",
	},
	"collectinfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectInfo",
		Description:    `CollectInfo indicates to the enforcer it needs to collect information.`,
		Exposed:        true,
		Name:           "collectInfo",
		Stored:         true,
		Type:           "boolean",
	},
	"collectedinfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectedInfo",
		Description:    `CollectedInfo represents the latest info collected by the enforcer.`,
		Exposed:        true,
		Name:           "collectedInfo",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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
		Description: `CurrentVersion holds the enforcerd binary version that is currently associated
to this object.`,
		Exposed:    true,
		Filterable: true,
		Name:       "currentVersion",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
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
		Description:    `Status of enforcement for PU managed directly by enforcerd, like Host PUs.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcementStatus",
		Stored:         true,
		Type:           "enum",
	},
	"lastcollectiontime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastCollectionTime",
		Description: `LastCollectionTime represents the date and time when the info have been
collected.`,
		Exposed: true,
		Name:    "lastCollectionTime",
		Stored:  true,
		Type:    "time",
	},
	"lastpoketime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastPokeTime",
		Description:    `Last poke is the time when the enforcer got last poked.`,
		Name:           "lastPokeTime",
		Stored:         true,
		Type:           "time",
	},
	"lastsynctime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastSyncTime",
		Description:    `LastSyncTime holds the last heart beat time.`,
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
		Description: `LocalCA contains the initial chain of trust for the enforcer. This valud is only
given when you retrieve a single enforcer.`,
		Exposed:   true,
		Name:      "localCA",
		Transient: true,
		Type:      "string",
	},
	"machineid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MachineID",
		Description: `MachineID holds a unique identifier for every machine as detected by the
enforcer. It is based on hardware information such as the SMBIOS UUID, MAC
addresses of interfaces or cloud provider IDs.`,
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
		Description: `Metadata contains tags that can only be set during creation. They must all start
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
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
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
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
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
		AllowedChoices: []string{"Registered", "Connected", "Disconnected", "Initialized", "Unknown"},
		ConvertedName:  "OperationalStatus",
		DefaultValue:   EnforcerOperationalStatusRegistered,
		Description:    `OperationalStatus tells the status of the enforcer.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		Stored:         true,
		Type:           "enum",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
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
		Description: `PublicToken is the public token of the server that will be included in the
datapath and its signed by the private CA.`,
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
		Description: `startTime holds the time this enforcerd was started. This time-stamp is reported
by the enforcer and is is preserved across disconnects.`,
		Exposed:   true,
		Name:      "startTime",
		Orderable: true,
		Stored:    true,
		Type:      "time",
	},
	"updateavailable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "UpdateAvailable",
		Description:    `Tells if the the version of enforcerd is outdated and should be updated.`,
		Exposed:        true,
		Name:           "updateAvailable",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
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
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
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
		"namespace",
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
	// FQDN contains the fqdn of the server where the enforcer is running.
	FQDN *string `json:"FQDN,omitempty" bson:"fqdn" mapstructure:"FQDN,omitempty"`

	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Certificate is the certificate of the enforcer.
	Certificate *string `json:"certificate,omitempty" bson:"certificate" mapstructure:"certificate,omitempty"`

	// CertificateExpirationDate is the expiration date of the certificate. This is an
	// internal attribute, not exposed in the API.
	CertificateExpirationDate *time.Time `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// CertificateKey is the certificate key of the enforcer. This is an internal
	// attribute, not exposed in the API.
	CertificateKey *string `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// If not empty during a create or update generation, the provided CSR will be
	// validated and signed by the backend providing a renewed certificate.
	CertificateRequest *string `json:"certificateRequest,omitempty" bson:"-" mapstructure:"certificateRequest,omitempty"`

	// CollectInfo indicates to the enforcer it needs to collect information.
	CollectInfo *bool `json:"collectInfo,omitempty" bson:"collectinfo" mapstructure:"collectInfo,omitempty"`

	// CollectedInfo represents the latest info collected by the enforcer.
	CollectedInfo *map[string]string `json:"collectedInfo,omitempty" bson:"collectedinfo" mapstructure:"collectedInfo,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// CurrentVersion holds the enforcerd binary version that is currently associated
	// to this object.
	CurrentVersion *string `json:"currentVersion,omitempty" bson:"currentversion" mapstructure:"currentVersion,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// Status of enforcement for PU managed directly by enforcerd, like Host PUs.
	EnforcementStatus *EnforcerEnforcementStatusValue `json:"enforcementStatus,omitempty" bson:"enforcementstatus" mapstructure:"enforcementStatus,omitempty"`

	// LastCollectionTime represents the date and time when the info have been
	// collected.
	LastCollectionTime *time.Time `json:"lastCollectionTime,omitempty" bson:"lastcollectiontime" mapstructure:"lastCollectionTime,omitempty"`

	// Last poke is the time when the enforcer got last poked.
	LastPokeTime *time.Time `json:"-" bson:"lastpoketime" mapstructure:"-,omitempty"`

	// LastSyncTime holds the last heart beat time.
	LastSyncTime *time.Time `json:"lastSyncTime,omitempty" bson:"lastsynctime" mapstructure:"lastSyncTime,omitempty"`

	// LastValidHostServices is a read only attribute that stores the list valid host
	// services that have been applied to this enforcer. This list might be different
	// from the list retrieved through policy, if the dynamically calculated list leads
	// into conflicts.
	LastValidHostServices *HostServicesList `json:"-" bson:"lastvalidhostservices" mapstructure:"-,omitempty"`

	// LocalCA contains the initial chain of trust for the enforcer. This valud is only
	// given when you retrieve a single enforcer.
	LocalCA *string `json:"localCA,omitempty" bson:"-" mapstructure:"localCA,omitempty"`

	// MachineID holds a unique identifier for every machine as detected by the
	// enforcer. It is based on hardware information such as the SMBIOS UUID, MAC
	// addresses of interfaces or cloud provider IDs.
	MachineID *string `json:"machineID,omitempty" bson:"machineid" mapstructure:"machineID,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// OperationalStatus tells the status of the enforcer.
	OperationalStatus *EnforcerOperationalStatusValue `json:"operationalStatus,omitempty" bson:"operationalstatus" mapstructure:"operationalStatus,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected" mapstructure:"protected,omitempty"`

	// PublicToken is the public token of the server that will be included in the
	// datapath and its signed by the private CA.
	PublicToken *string `json:"publicToken,omitempty" bson:"publictoken" mapstructure:"publicToken,omitempty"`

	// startTime holds the time this enforcerd was started. This time-stamp is reported
	// by the enforcer and is is preserved across disconnects.
	StartTime *time.Time `json:"startTime,omitempty" bson:"starttime" mapstructure:"startTime,omitempty"`

	// Tells if the the version of enforcerd is outdated and should be updated.
	UpdateAvailable *bool `json:"updateAvailable,omitempty" bson:"updateavailable" mapstructure:"updateAvailable,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
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
	if o.MachineID != nil {
		out.MachineID = *o.MachineID
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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
	if o.UpdateAvailable != nil {
		out.UpdateAvailable = *o.UpdateAvailable
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
