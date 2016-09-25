package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/golang/constants"

// ServiceSslValue represents the possible values for attribute "ssl".
type ServiceSslValue string

const (
	// ServiceSslDisabled represents the value Disabled.
	ServiceSslDisabled ServiceSslValue = "Disabled"

	// ServiceSslEnabled represents the value Enabled.
	ServiceSslEnabled ServiceSslValue = "Enabled"
)

// ServiceIdentity represents the Identity of the object
var ServiceIdentity = elemental.Identity{
	Name:     "service",
	Category: "services",
}

// ServicesList represents a list of Services
type ServicesList []*Service

// Service represents the model of a service
type Service struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"id,omitempty"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" cql:"annotation,omitempty"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" cql:"associatedtags,omitempty"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" cql:"createdat,omitempty"`

	// Deleted marks if the entity has been deleted.
	Deleted bool `json:"-" cql:"deleted,omitempty"`

	// Endpoint is the API end point of the service
	Endpoint string `json:"endpoint" cql:"endpoint,omitempty"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" cql:"parentid,omitempty"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" cql:"parenttype,omitempty"`

	// Port is the port number of the service
	Port int `json:"port" cql:"port,omitempty"`

	// Server is either the DNS name or IP of the server that provides the service
	Server string `json:"server" cql:"server,primarykey,omitempty"`

	// SSL defines if the service is either secured or unsecured
	Ssl ServiceSslValue `json:"ssl" cql:"ssl,omitempty"`

	// Status of an entity
	Status constants.EntityStatus `json:"status" cql:"status,omitempty"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" cql:"updatedat,omitempty"`
}

// NewService returns a new *Service
func NewService() *Service {

	return &Service{
		Status: constants.Active,
	}
}

// Identity returns the Identity of the object.
func (o *Service) Identity() elemental.Identity {

	return ServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Service) Identifier() string {

	return o.ID
}

func (o *Service) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Service) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Service) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Service) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Service) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *Service) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *Service) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetNamespace returns the namespace of the receiver
func (o *Service) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Service) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetParentID returns the parentID of the receiver
func (o *Service) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *Service) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *Service) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *Service) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetStatus returns the status of the receiver
func (o *Service) GetStatus() constants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Service) SetStatus(status constants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Service) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Service) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateStringInList("ssl", string(o.Ssl), []string{"Disabled", "Enabled"}, false); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o Service) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ServiceAttributesMap[name]
}

// ServiceAttributesMap represents the map of attribute for Service.
var ServiceAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Annotation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"CreatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Deleted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Getter:         true,
		Name:           "deleted",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Endpoint": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "endpoint",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"Port": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "port",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"Server": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "server",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
	"Ssl": elemental.AttributeSpecification{
		AllowedChoices: []string{"Disabled", "Enabled"},
		Exposed:        true,
		Filterable:     true,
		Name:           "ssl",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
