package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/golang/constants"

// ServerCertificateStatusValue represents the possible values for attribute "certificateStatus".
type ServerCertificateStatusValue string

const (
	// ServerCertificateStatusRenew represents the value RENEW.
	ServerCertificateStatusRenew ServerCertificateStatusValue = "RENEW"

	// ServerCertificateStatusRevoked represents the value REVOKED.
	ServerCertificateStatusRevoked ServerCertificateStatusValue = "REVOKED"

	// ServerCertificateStatusValid represents the value VALID.
	ServerCertificateStatusValid ServerCertificateStatusValue = "VALID"
)

// ServerOperationalStatusValue represents the possible values for attribute "operationalStatus".
type ServerOperationalStatusValue string

const (
	// ServerOperationalStatusConnected represents the value Connected.
	ServerOperationalStatusConnected ServerOperationalStatusValue = "Connected"

	// ServerOperationalStatusDisconnected represents the value Disconnected.
	ServerOperationalStatusDisconnected ServerOperationalStatusValue = "Disconnected"

	// ServerOperationalStatusInitialized represents the value Initialized.
	ServerOperationalStatusInitialized ServerOperationalStatusValue = "Initialized"

	// ServerOperationalStatusUnknown represents the value Unknown.
	ServerOperationalStatusUnknown ServerOperationalStatusValue = "Unknown"
)

// ServerIdentity represents the Identity of the object
var ServerIdentity = elemental.Identity{
	Name:     "server",
	Category: "servers",
}

// ServersList represents a list of Servers
type ServersList []*Server

// Server represents the model of a server
type Server struct {
	// FQDN contains the fqdn of the server.
	FQDN string `json:"FQDN" cql:"fqdn,omitempty" bson:"fqdn"`

	// ID is the identifier of the object.
	ID string `json:"ID" cql:"id,primarykey,omitempty" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" cql:"annotation,omitempty" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" cql:"associatedtags,omitempty" bson:"associatedtags"`

	// Certificate is the certificate of the server
	Certificate string `json:"certificate" cql:"certificate,omitempty" bson:"certificate"`

	// CertificateExpirationDate is the expiration date of the certificate.
	CertificateExpirationDate time.Time `json:"certificateExpirationDate" cql:"certificateexpirationdate,omitempty" bson:"certificateexpirationdate"`

	// CertificateKey is the secret key of the server. Returned only when a server is created or the certificate is updated.
	CertificateKey string `json:"certificateKey" cql:"-" bson:"-"`

	// CertificateStatus indicates if the certificate is valid.
	CertificateStatus ServerCertificateStatusValue `json:"certificateStatus" cql:"certificatestatus,omitempty" bson:"certificatestatus"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" cql:"createdat,omitempty" bson:"createdat"`

	// Description is the description of the object.
	Description string `json:"description" cql:"description,omitempty" bson:"description"`

	// LastSyncTime holds the last heart beat time.
	LastSyncTime time.Time `json:"lastSyncTime" cql:"lastsynctime,omitempty" bson:"lastsynctime"`

	// Name is the name of the entity
	Name string `json:"name" cql:"name,omitempty" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" cql:"normalizedtags,omitempty" bson:"normalizedtags"`

	// OperationalStatus tells the status of the server
	OperationalStatus ServerOperationalStatusValue `json:"operationalStatus" cql:"-" bson:"-"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" cql:"parentid,omitempty" bson:"parentid"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" cql:"parenttype,omitempty" bson:"parenttype"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" cql:"protected,omitempty" bson:"protected"`

	// Status of an entity
	Status constants.EntityStatus `json:"status" cql:"status,omitempty" bson:"status"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" cql:"updatedat,omitempty" bson:"updatedat"`
}

// NewServer returns a new *Server
func NewServer() *Server {

	return &Server{
		AssociatedTags:    []string{},
		CertificateStatus: "VALID",
		NormalizedTags:    []string{},
		OperationalStatus: "Initialized",
		Status:            constants.Active,
	}
}

// Identity returns the Identity of the object.
func (o *Server) Identity() elemental.Identity {

	return ServerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Server) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Server) SetIdentifier(ID string) {

	o.ID = ID
}

func (o *Server) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Server) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Server) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Server) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetName returns the name of the receiver
func (o *Server) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *Server) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *Server) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Server) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *Server) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *Server) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetParentID returns the parentID of the receiver
func (o *Server) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *Server) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *Server) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *Server) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetProtected returns the protected of the receiver
func (o *Server) GetProtected() bool {
	return o.Protected
}

// GetStatus returns the status of the receiver
func (o *Server) GetStatus() constants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Server) SetStatus(status constants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Server) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Server) Validate() error {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("FQDN", o.FQDN); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("certificateStatus", string(o.CertificateStatus), []string{"RENEW", "REVOKED", "VALID"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredTime("lastSyncTime", o.LastSyncTime); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (Server) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ServerAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (Server) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ServerAttributesMap
}

// ServerAttributesMap represents the map of attribute for Server.
var ServerAttributesMap = map[string]elemental.AttributeSpecification{
	"FQDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `FQDN contains the fqdn of the server.`,
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
		Description:    `Certificate is the certificate of the server `,
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
		Description:    `CertificateKey is the secret key of the server. Returned only when a server is created or the certificate is updated.`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificateKey",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"CertificateStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"RENEW", "REVOKED", "VALID"},
		Description:    `CertificateStatus indicates if the certificate is valid.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "certificateStatus",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"CreatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CreatedAt is the time at which an entity was created`,
		Exposed:        true,
		Name:           "createdAt",
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
		Description:    `OperationalStatus tells the status of the server`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		ReadOnly:       true,
		Transient:      true,
		Type:           "enum",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentID is the ID of the parent, if any,`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentType": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Status of an entity`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "status",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		SubType:        "status_enum",
		Type:           "external",
	},
	"UpdatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdatedAt is the time at which an entity was updated.`,
		Exposed:        true,
		Name:           "updatedAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
