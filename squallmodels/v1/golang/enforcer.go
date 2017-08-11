package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "time"

// EnforcerCertificateStatusValue represents the possible values for attribute "certificateStatus".
type EnforcerCertificateStatusValue string

const (
	// EnforcerCertificateStatusRenew represents the value RENEW.
	EnforcerCertificateStatusRenew EnforcerCertificateStatusValue = "RENEW"

	// EnforcerCertificateStatusRevoked represents the value REVOKED.
	EnforcerCertificateStatusRevoked EnforcerCertificateStatusValue = "REVOKED"

	// EnforcerCertificateStatusValid represents the value VALID.
	EnforcerCertificateStatusValid EnforcerCertificateStatusValue = "VALID"
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

// EnforcerIdentity represents the Identity of the object
var EnforcerIdentity = elemental.Identity{
	Name:     "enforcer",
	Category: "enforcers",
}

// EnforcersList represents a list of Enforcers
type EnforcersList []*Enforcer

// ContentIdentity returns the identity of the objects in the list.
func (o EnforcersList) ContentIdentity() elemental.Identity {

	return EnforcerIdentity
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
	FQDN string `json:"FQDN" bson:"fqdn"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotations map[string][]string `json:"annotations" bson:"annotations"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// Certificate is the certificate of the enforcer.
	Certificate string `json:"certificate" bson:"certificate"`

	// CertificateExpirationDate is the expiration date of the certificate.
	CertificateExpirationDate time.Time `json:"certificateExpirationDate" bson:"certificateexpirationdate"`

	// CertificateKey is the secret key of the enforcer. Returned only when a enforcer is created or the certificate is updated.
	CertificateKey string `json:"certificateKey" bson:"-"`

	// CertificateStatus indicates if the certificate is valid.
	CertificateStatus EnforcerCertificateStatusValue `json:"certificateStatus" bson:"certificatestatus"`

	// CreatedTime is the time at which the object was created
	CreateTime time.Time `json:"createTime" bson:"createtime"`

	// CurrentVersion holds the enforcerd binary version that is currently associated to this object.
	CurrentVersion string `json:"currentVersion" bson:"currentversion"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// Contains the ID of the profile used by the instance of enforcerd.
	EnforcerProfileID string `json:"enforcerProfileID" bson:"enforcerprofileid"`

	// LastSyncTime holds the last heart beat time.
	LastSyncTime time.Time `json:"lastSyncTime" bson:"lastsynctime"`

	// Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// OperationalStatus tells the status of the enforcer.
	OperationalStatus EnforcerOperationalStatusValue `json:"operationalStatus" bson:"-"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// PublicToken is the public token of the server that will be included in the datapath and its signed by the private CA.
	PublicToken string `json:"publicToken" bson:"publictoken"`

	// Tells if the the version of enforcerd is outdated and should be updated.
	UpdateAvailable bool `json:"updateAvailable" bson:"updateavailable"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewEnforcer returns a new *Enforcer
func NewEnforcer() *Enforcer {

	return &Enforcer{
		ModelVersion:      1,
		Annotations:       map[string][]string{},
		AssociatedTags:    []string{},
		CertificateStatus: "VALID",
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
func (o *Enforcer) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
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
	return `A Enforcer Profile contains a configuration for a Enforcer. It contains various parameters, like what network should not policeds, what processing units should be ignored based on their tags and so on. It also contains more advanced parameters to fine tune the Agent. A Enforcer will decide what profile to apply using aEnforcer Profile Mapping Policy. This policy will decide according the Enforcer's tags what profile to use. If an Enforcer tags are matching more than a single policy, it will refuse to start. Some parameters will be applied directly to a running agent, some will need to restart it.`
}

func (o *Enforcer) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Enforcer) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Enforcer) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the createTime of the receiver
func (o *Enforcer) GetCreateTime() time.Time {
	return o.CreateTime
}

// SetCreateTime set the given createTime of the receiver
func (o *Enforcer) SetCreateTime(createTime time.Time) {
	o.CreateTime = createTime
}

// GetMetadata returns the metadata of the receiver
func (o *Enforcer) GetMetadata() []string {
	return o.Metadata
}

// SetMetadata set the given metadata of the receiver
func (o *Enforcer) SetMetadata(metadata []string) {
	o.Metadata = metadata
}

// GetName returns the name of the receiver
func (o *Enforcer) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *Enforcer) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *Enforcer) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Enforcer) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *Enforcer) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *Enforcer) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetProtected returns the protected of the receiver
func (o *Enforcer) GetProtected() bool {
	return o.Protected
}

// GetUpdateTime returns the updateTime of the receiver
func (o *Enforcer) GetUpdateTime() time.Time {
	return o.UpdateTime
}

// SetUpdateTime set the given updateTime of the receiver
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

	if err := elemental.ValidateRequiredString("FQDN", o.FQDN); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("certificateStatus", string(o.CertificateStatus), []string{"RENEW", "REVOKED", "VALID"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredTime("lastSyncTime", o.LastSyncTime); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredTime("lastSyncTime", o.LastSyncTime); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
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
		Unique:         true,
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotations",
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AssociatedTags are the list of tags attached to an entity`,
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
		Description:    `Certificate is the certificate of the enforcer. `,
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
		Description:    `CertificateKey is the secret key of the enforcer. Returned only when a enforcer is created or the certificate is updated.`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificateKey",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"CertificateStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"RENEW", "REVOKED", "VALID"},
		DefaultValue:   EnforcerCertificateStatusValue("VALID"),
		Description:    `CertificateStatus indicates if the certificate is valid.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "certificateStatus",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CreatedTime is the time at which the object was created`,
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
		Description:    `CurrentVersion holds the enforcerd binary version that is currently associated to this object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "currentVersion",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerProfileID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Contains the ID of the profile used by the instance of enforcerd.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "enforcerProfileID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LastSyncTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LastSyncTime holds the last heart beat time.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "time",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "metadata",
		Setter:         true,
		Stored:         true,
		SubType:        "metadata_list",
		Type:           "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultOrder:   true,
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity`,
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
		Description:    `NormalizedTags contains the list of normalized tags of the entities`,
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
		DefaultValue:   EnforcerOperationalStatusValue("Initialized"),
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
		Description:    `PublicToken is the public token of the server that will be included in the datapath and its signed by the private CA.`,
		Exposed:        true,
		Format:         "free",
		Name:           "publicToken",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateAvailable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		Unique:         true,
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotations",
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AssociatedTags are the list of tags attached to an entity`,
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
		Description:    `Certificate is the certificate of the enforcer. `,
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
		Description:    `CertificateKey is the secret key of the enforcer. Returned only when a enforcer is created or the certificate is updated.`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificateKey",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"certificatestatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"RENEW", "REVOKED", "VALID"},
		DefaultValue:   EnforcerCertificateStatusValue("VALID"),
		Description:    `CertificateStatus indicates if the certificate is valid.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "certificateStatus",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CreatedTime is the time at which the object was created`,
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
		Description:    `CurrentVersion holds the enforcerd binary version that is currently associated to this object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "currentVersion",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"enforcerprofileid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Contains the ID of the profile used by the instance of enforcerd.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "enforcerProfileID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"lastsynctime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LastSyncTime holds the last heart beat time.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "time",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "metadata",
		Setter:         true,
		Stored:         true,
		SubType:        "metadata_list",
		Type:           "external",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultOrder:   true,
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity`,
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
		Description:    `NormalizedTags contains the list of normalized tags of the entities`,
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
		DefaultValue:   EnforcerOperationalStatusValue("Initialized"),
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
		Description:    `PublicToken is the public token of the server that will be included in the datapath and its signed by the private CA.`,
		Exposed:        true,
		Format:         "free",
		Name:           "publicToken",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"updateavailable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
