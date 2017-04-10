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

// List convert the object to and elemental.IdentifiablesList.
func (o EnforcersList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// Enforcer represents the model of a enforcer
type Enforcer struct {
	// FQDN contains the fqdn of the server where the enforcer is running.
	FQDN string `json:"FQDN" bson:"fqdn"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" bson:"annotation"`

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

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// LastSyncTime holds the last heart beat time.
	LastSyncTime time.Time `json:"lastSyncTime" bson:"lastsynctime"`

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

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewEnforcer returns a new *Enforcer
func NewEnforcer() *Enforcer {

	return &Enforcer{
		ModelVersion:      1.0,
		AssociatedTags:    []string{},
		CertificateStatus: "VALID",
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
func (o *Enforcer) Version() float64 {

	return 1.0
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

// SetCreateTime set the given createTime of the receiver
func (o *Enforcer) SetCreateTime(createTime time.Time) {
	o.CreateTime = createTime
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

	return EnforcerAttributesMap[name]
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
	"Annotation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
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
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
