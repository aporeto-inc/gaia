package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
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

	// CertificateExpirationDate is the expiration date of the certificate.
	CertificateExpirationDate time.Time `json:"certificateExpirationDate" bson:"certificateexpirationdate" mapstructure:"certificateExpirationDate,omitempty"`

	// CertificateKey is the secret key of the enforcer. Returned only when a enforcer
	// is created or the certificate is updated.
	CertificateKey string `json:"certificateKey" bson:"-" mapstructure:"certificateKey,omitempty"`

	// CertificateRequest can be used to generate a certificate from that CSR instead
	// of letting the server generate your private key for you.
	CertificateRequest string `json:"certificateRequest" bson:"-" mapstructure:"certificateRequest,omitempty"`

	// If set during creation,the server will not initially generate a certificate. In
	// that case you have to provide a valid CSR through certificateRequest during an
	// update.
	CertificateRequestEnabled bool `json:"certificateRequestEnabled" bson:"certificaterequestenabled" mapstructure:"certificateRequestEnabled,omitempty"`

	// CollectInfo indicates to the enforcer it needs to collect information.
	CollectInfo bool `json:"collectInfo" bson:"collectinfo" mapstructure:"collectInfo,omitempty"`

	// CollectedInfo represents the latest info collected by the enforcer.
	CollectedInfo map[string]string `json:"collectedInfo" bson:"collectedinfo" mapstructure:"collectedInfo,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// CurrentVersion holds the enforcerd binary version that is currently associated
	// to this object.
	CurrentVersion string `json:"currentVersion" bson:"currentversion" mapstructure:"currentVersion,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Contains the ID of the profile used by the instance of enforcerd.
	EnforcerProfileID string `json:"enforcerProfileID" bson:"enforcerprofileid" mapstructure:"enforcerProfileID,omitempty"`

	// LastCollectionTime represents the date and time when the info have been
	// collected.
	LastCollectionTime time.Time `json:"lastCollectionTime" bson:"lastcollectiontime" mapstructure:"lastCollectionTime,omitempty"`

	// LastSyncTime holds the last heart beat time.
	LastSyncTime time.Time `json:"lastSyncTime" bson:"lastsynctime" mapstructure:"lastSyncTime,omitempty"`

	// LocalCA contains the initial chain of trust for the enforcer. This valud is only
	// given when you retrieve a single enforcer.
	LocalCA string `json:"localCA" bson:"-" mapstructure:"localCA,omitempty"`

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
	OperationalStatus EnforcerOperationalStatusValue `json:"operationalStatus" bson:"-" mapstructure:"operationalStatus,omitempty"`

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

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewEnforcer returns a new *Enforcer
func NewEnforcer() *Enforcer {

	return &Enforcer{
		ModelVersion:      1,
		Annotations:       map[string][]string{},
		AssociatedTags:    []string{},
		NormalizedTags:    []string{},
		Metadata:          []string{},
		OperationalStatus: EnforcerOperationalStatusRegistered,
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
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Enforcer) Doc() string {
	return `An Enforcer Profile contains a configuration for a Enforcer. It contains various
parameters, like what network should not policeds, what processing units should
be ignored based on their tags and so on. It also contains more advanced
parameters to fine tune the Agent. A Enforcer will decide what profile to apply
using aEnforcer Profile Mapping Policy. This policy will decide according the
Enforcer's tags what profile to use. If an Enforcer tags are matching more than
a single policy, it will refuse to start. Some parameters will be applied
directly to a running agent, some will need to restart it.`
}

func (o *Enforcer) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Enforcer) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *Enforcer) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Enforcer) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *Enforcer) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Enforcer) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *Enforcer) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *Enforcer) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *Enforcer) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *Enforcer) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *Enforcer) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Enforcer) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *Enforcer) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Enforcer) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *Enforcer) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Enforcer) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Enforcer) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *Enforcer) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
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
			CertificateRequestEnabled: &o.CertificateRequestEnabled,
			CollectInfo:               &o.CollectInfo,
			CollectedInfo:             &o.CollectedInfo,
			CreateTime:                &o.CreateTime,
			CurrentVersion:            &o.CurrentVersion,
			Description:               &o.Description,
			EnforcerProfileID:         &o.EnforcerProfileID,
			LastCollectionTime:        &o.LastCollectionTime,
			LastSyncTime:              &o.LastSyncTime,
			LocalCA:                   &o.LocalCA,
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
		case "certificateRequestEnabled":
			sp.CertificateRequestEnabled = &(o.CertificateRequestEnabled)
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
		case "enforcerProfileID":
			sp.EnforcerProfileID = &(o.EnforcerProfileID)
		case "lastCollectionTime":
			sp.LastCollectionTime = &(o.LastCollectionTime)
		case "lastSyncTime":
			sp.LastSyncTime = &(o.LastSyncTime)
		case "localCA":
			sp.LocalCA = &(o.LocalCA)
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
	if so.CertificateRequestEnabled != nil {
		o.CertificateRequestEnabled = *so.CertificateRequestEnabled
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
	if so.EnforcerProfileID != nil {
		o.EnforcerProfileID = *so.EnforcerProfileID
	}
	if so.LastCollectionTime != nil {
		o.LastCollectionTime = *so.LastCollectionTime
	}
	if so.LastSyncTime != nil {
		o.LastSyncTime = *so.LastSyncTime
	}
	if so.LocalCA != nil {
		o.LocalCA = *so.LocalCA
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

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
		PrimaryKey:     true,
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
		SubType:        "annotations",
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
		SubType:        "tags_list",
		Type:           "external",
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
		ConvertedName:  "CertificateExpirationDate",
		Description:    `CertificateExpirationDate is the expiration date of the certificate.`,
		Exposed:        true,
		Name:           "certificateExpirationDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"CertificateKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CertificateKey",
		Description: `CertificateKey is the secret key of the enforcer. Returned only when a enforcer
is created or the certificate is updated.`,
		Exposed:   true,
		Name:      "certificateKey",
		Orderable: true,
		ReadOnly:  true,
		Type:      "string",
	},
	"CertificateRequest": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateRequest",
		Description: `CertificateRequest can be used to generate a certificate from that CSR instead
of letting the server generate your private key for you.`,
		Exposed:   true,
		Name:      "certificateRequest",
		Transient: true,
		Type:      "string",
	},
	"CertificateRequestEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateRequestEnabled",
		CreationOnly:   true,
		Description: `If set during creation,the server will not initially generate a certificate. In
that case you have to provide a valid CSR through certificateRequest during an
update.`,
		Exposed: true,
		Name:    "certificateRequestEnabled",
		Stored:  true,
		Type:    "boolean",
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
		SubType:        "collected_info",
		Type:           "external",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created.`,
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
		Exposed:   true,
		Name:      "currentVersion",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerProfileID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerProfileID",
		Description:    `Contains the ID of the profile used by the instance of enforcerd.`,
		Exposed:        true,
		Name:           "enforcerProfileID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
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
		SubType:    "metadata_list",
		Type:       "external",
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
		CreationOnly:   true,
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
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"OperationalStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Registered", "Connected", "Disconnected", "Initialized", "Unknown"},
		Autogenerated:  true,
		ConvertedName:  "OperationalStatus",
		DefaultValue:   EnforcerOperationalStatusRegistered,
		Description:    `OperationalStatus tells the status of the enforcer.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		Transient:      true,
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
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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
		PrimaryKey:     true,
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
		SubType:        "annotations",
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
		SubType:        "tags_list",
		Type:           "external",
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
		ConvertedName:  "CertificateExpirationDate",
		Description:    `CertificateExpirationDate is the expiration date of the certificate.`,
		Exposed:        true,
		Name:           "certificateExpirationDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"certificatekey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CertificateKey",
		Description: `CertificateKey is the secret key of the enforcer. Returned only when a enforcer
is created or the certificate is updated.`,
		Exposed:   true,
		Name:      "certificateKey",
		Orderable: true,
		ReadOnly:  true,
		Type:      "string",
	},
	"certificaterequest": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateRequest",
		Description: `CertificateRequest can be used to generate a certificate from that CSR instead
of letting the server generate your private key for you.`,
		Exposed:   true,
		Name:      "certificateRequest",
		Transient: true,
		Type:      "string",
	},
	"certificaterequestenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateRequestEnabled",
		CreationOnly:   true,
		Description: `If set during creation,the server will not initially generate a certificate. In
that case you have to provide a valid CSR through certificateRequest during an
update.`,
		Exposed: true,
		Name:    "certificateRequestEnabled",
		Stored:  true,
		Type:    "boolean",
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
		SubType:        "collected_info",
		Type:           "external",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created.`,
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
		Exposed:   true,
		Name:      "currentVersion",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"enforcerprofileid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerProfileID",
		Description:    `Contains the ID of the profile used by the instance of enforcerd.`,
		Exposed:        true,
		Name:           "enforcerProfileID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
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
		SubType:    "metadata_list",
		Type:       "external",
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
		CreationOnly:   true,
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
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"operationalstatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Registered", "Connected", "Disconnected", "Initialized", "Unknown"},
		Autogenerated:  true,
		ConvertedName:  "OperationalStatus",
		DefaultValue:   EnforcerOperationalStatusRegistered,
		Description:    `OperationalStatus tells the status of the enforcer.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		Transient:      true,
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
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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

	// CertificateExpirationDate is the expiration date of the certificate.
	CertificateExpirationDate *time.Time `json:"certificateExpirationDate,omitempty" bson:"certificateexpirationdate" mapstructure:"certificateExpirationDate,omitempty"`

	// CertificateKey is the secret key of the enforcer. Returned only when a enforcer
	// is created or the certificate is updated.
	CertificateKey *string `json:"certificateKey,omitempty" bson:"-" mapstructure:"certificateKey,omitempty"`

	// CertificateRequest can be used to generate a certificate from that CSR instead
	// of letting the server generate your private key for you.
	CertificateRequest *string `json:"certificateRequest,omitempty" bson:"-" mapstructure:"certificateRequest,omitempty"`

	// If set during creation,the server will not initially generate a certificate. In
	// that case you have to provide a valid CSR through certificateRequest during an
	// update.
	CertificateRequestEnabled *bool `json:"certificateRequestEnabled,omitempty" bson:"certificaterequestenabled" mapstructure:"certificateRequestEnabled,omitempty"`

	// CollectInfo indicates to the enforcer it needs to collect information.
	CollectInfo *bool `json:"collectInfo,omitempty" bson:"collectinfo" mapstructure:"collectInfo,omitempty"`

	// CollectedInfo represents the latest info collected by the enforcer.
	CollectedInfo *map[string]string `json:"collectedInfo,omitempty" bson:"collectedinfo" mapstructure:"collectedInfo,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// CurrentVersion holds the enforcerd binary version that is currently associated
	// to this object.
	CurrentVersion *string `json:"currentVersion,omitempty" bson:"currentversion" mapstructure:"currentVersion,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// Contains the ID of the profile used by the instance of enforcerd.
	EnforcerProfileID *string `json:"enforcerProfileID,omitempty" bson:"enforcerprofileid" mapstructure:"enforcerProfileID,omitempty"`

	// LastCollectionTime represents the date and time when the info have been
	// collected.
	LastCollectionTime *time.Time `json:"lastCollectionTime,omitempty" bson:"lastcollectiontime" mapstructure:"lastCollectionTime,omitempty"`

	// LastSyncTime holds the last heart beat time.
	LastSyncTime *time.Time `json:"lastSyncTime,omitempty" bson:"lastsynctime" mapstructure:"lastSyncTime,omitempty"`

	// LocalCA contains the initial chain of trust for the enforcer. This valud is only
	// given when you retrieve a single enforcer.
	LocalCA *string `json:"localCA,omitempty" bson:"-" mapstructure:"localCA,omitempty"`

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
	OperationalStatus *EnforcerOperationalStatusValue `json:"operationalStatus,omitempty" bson:"-" mapstructure:"operationalStatus,omitempty"`

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

	// UpdateTime is the time at which an entity was updated.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
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
	if o.CertificateRequestEnabled != nil {
		out.CertificateRequestEnabled = *o.CertificateRequestEnabled
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
	if o.EnforcerProfileID != nil {
		out.EnforcerProfileID = *o.EnforcerProfileID
	}
	if o.LastCollectionTime != nil {
		out.LastCollectionTime = *o.LastCollectionTime
	}
	if o.LastSyncTime != nil {
		out.LastSyncTime = *o.LastSyncTime
	}
	if o.LocalCA != nil {
		out.LocalCA = *o.LocalCA
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

	return out
}
