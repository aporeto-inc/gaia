package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
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

	// EnforcerOperationalStatusUnknown represents the value Unknown.
	EnforcerOperationalStatusUnknown EnforcerOperationalStatusValue = "Unknown"
)

// EnforcerIdentity represents the Identity of the object.
var EnforcerIdentity = elemental.Identity{
	Name:     "enforcer",
	Category: "enforcers",
	Private:  false,
}

// EnforcersList represents a list of Enforcers
type EnforcersList []*Enforcer

// ContentIdentity returns the identity of the objects in the list.
func (o EnforcersList) ContentIdentity() elemental.Identity {

	return EnforcerIdentity
}

// Copy returns a pointer to a copy the EnforcersList.
func (o EnforcersList) Copy() elemental.ContentIdentifiable {

	copy := append(EnforcersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcersList.
func (o EnforcersList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(EnforcersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Enforcer))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcersList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcersList) DefaultOrder() []string {

	return []string{
		"name",
	}
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

	// createTime holds the time this enforcerd was created. This is preserved across
	// disconnects.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// CurrentVersion holds the enforcerd binary version that is currently associated
	// to this object.
	CurrentVersion string `json:"currentVersion" bson:"currentversion" mapstructure:"currentVersion,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Contains the ID of the profile used by the instance of enforcerd.
	EnforcerProfileID string `json:"enforcerProfileID" bson:"enforcerprofileid" mapstructure:"enforcerProfileID,omitempty"`

	// initializeTime holds the time this enforcerd completed initialization. This is
	// preserved across disconnects.
	InitializeTime time.Time `json:"initializeTime" bson:"initializetime" mapstructure:"initializeTime,omitempty"`

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
	PublicToken string `json:"publicToken" bson:"-" mapstructure:"publicToken,omitempty"`

	// Tells if the the version of enforcerd is outdated and should be updated.
	UpdateAvailable bool `json:"updateAvailable" bson:"updateavailable" mapstructure:"updateAvailable,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewEnforcer returns a new *Enforcer
func NewEnforcer() *Enforcer {

	return &Enforcer{
		ModelVersion:      1,
		Annotations:       map[string][]string{},
		AssociatedTags:    []string{},
		Metadata:          []string{},
		NormalizedTags:    []string{},
		OperationalStatus: "Initialized",
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
		Filterable:     true,
		Format:         "free",
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
		Format:         "free",
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
		Format:         "free",
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
		Filterable:     true,
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
		Format:    "free",
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
		Format:    "free",
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
		ConvertedName:  "CreateTime",
		Description: `createTime holds the time this enforcerd was created. This is preserved across
disconnects.`,
		Exposed:    true,
		Filterable: true,
		Name:       "createTime",
		Orderable:  true,
		Stored:     true,
		Type:       "time",
	},
	"CurrentVersion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CurrentVersion",
		Description: `CurrentVersion holds the enforcerd binary version that is currently associated
to this object.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
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
		Format:         "free",
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
		Filterable:     true,
		Format:         "free",
		Name:           "enforcerProfileID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"InitializeTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "InitializeTime",
		Description: `initializeTime holds the time this enforcerd completed initialization. This is
preserved across disconnects.`,
		Exposed:    true,
		Filterable: true,
		Name:       "initializeTime",
		Orderable:  true,
		Stored:     true,
		Type:       "time",
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
		Filterable:     true,
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
		Format:    "free",
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
		Format:         "free",
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
		Format:         "free",
		Getter:         true,
		Index:          true,
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
		AllowedChoices: []string{"Connected", "Disconnected", "Initialized", "Unknown"},
		Autogenerated:  true,
		ConvertedName:  "OperationalStatus",
		DefaultValue:   EnforcerOperationalStatusInitialized,
		Description:    `OperationalStatus tells the status of the enforcer.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		ReadOnly:       true,
		Transient:      true,
		Type:           "enum",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Filterable:     true,
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
		Format:    "free",
		Name:      "publicToken",
		ReadOnly:  true,
		Transient: true,
		Type:      "string",
	},
	"UpdateAvailable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "UpdateAvailable",
		Description:    `Tells if the the version of enforcerd is outdated and should be updated.`,
		Exposed:        true,
		Filterable:     true,
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
		Filterable:     true,
		Format:         "free",
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
		Format:         "free",
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
		Format:         "free",
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
		Filterable:     true,
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
		Format:    "free",
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
		Format:    "free",
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
		ConvertedName:  "CreateTime",
		Description: `createTime holds the time this enforcerd was created. This is preserved across
disconnects.`,
		Exposed:    true,
		Filterable: true,
		Name:       "createTime",
		Orderable:  true,
		Stored:     true,
		Type:       "time",
	},
	"currentversion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CurrentVersion",
		Description: `CurrentVersion holds the enforcerd binary version that is currently associated
to this object.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
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
		Format:         "free",
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
		Filterable:     true,
		Format:         "free",
		Name:           "enforcerProfileID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"initializetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "InitializeTime",
		Description: `initializeTime holds the time this enforcerd completed initialization. This is
preserved across disconnects.`,
		Exposed:    true,
		Filterable: true,
		Name:       "initializeTime",
		Orderable:  true,
		Stored:     true,
		Type:       "time",
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
		Filterable:     true,
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
		Format:    "free",
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
		Format:         "free",
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
		Format:         "free",
		Getter:         true,
		Index:          true,
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
		AllowedChoices: []string{"Connected", "Disconnected", "Initialized", "Unknown"},
		Autogenerated:  true,
		ConvertedName:  "OperationalStatus",
		DefaultValue:   EnforcerOperationalStatusInitialized,
		Description:    `OperationalStatus tells the status of the enforcer.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		ReadOnly:       true,
		Transient:      true,
		Type:           "enum",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Filterable:     true,
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
		Format:    "free",
		Name:      "publicToken",
		ReadOnly:  true,
		Transient: true,
		Type:      "string",
	},
	"updateavailable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "UpdateAvailable",
		Description:    `Tells if the the version of enforcerd is outdated and should be updated.`,
		Exposed:        true,
		Filterable:     true,
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
