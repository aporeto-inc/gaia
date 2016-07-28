package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	ServerAttributeNameID                elemental.AttributeSpecificationNameKey = "server/ID"
	ServerAttributeNameAddress           elemental.AttributeSpecificationNameKey = "server/address"
	ServerAttributeNameAnnotation        elemental.AttributeSpecificationNameKey = "server/annotation"
	ServerAttributeNameAssociatedTags    elemental.AttributeSpecificationNameKey = "server/associatedTags"
	ServerAttributeNameCreatedAt         elemental.AttributeSpecificationNameKey = "server/createdAt"
	ServerAttributeNameDeleted           elemental.AttributeSpecificationNameKey = "server/deleted"
	ServerAttributeNameDescription       elemental.AttributeSpecificationNameKey = "server/description"
	ServerAttributeNameDomain            elemental.AttributeSpecificationNameKey = "server/domain"
	ServerAttributeNameEnvironment       elemental.AttributeSpecificationNameKey = "server/environment"
	ServerAttributeNameName              elemental.AttributeSpecificationNameKey = "server/name"
	ServerAttributeNameNamespace         elemental.AttributeSpecificationNameKey = "server/namespace"
	ServerAttributeNameOperationalStatus elemental.AttributeSpecificationNameKey = "server/operationalStatus"
	ServerAttributeNameStatus            elemental.AttributeSpecificationNameKey = "server/status"
	ServerAttributeNameUpdatedAt         elemental.AttributeSpecificationNameKey = "server/updatedAt"
)

// ServerEnvironmentValue represents the possible values for attribute "environment".
type ServerEnvironmentValue string

const (
	ServerEnvironmentAws     ServerEnvironmentValue = "AWS"
	ServerEnvironmentGcp     ServerEnvironmentValue = "GCP"
	ServerEnvironmentPrivate ServerEnvironmentValue = "Private"
)

// ServerOperationalStatusValue represents the possible values for attribute "operationalStatus".
type ServerOperationalStatusValue string

const (
	ServerOperationalStatusConnected   ServerOperationalStatusValue = "CONNECTED"
	ServerOperationalStatusInitialized ServerOperationalStatusValue = "INITIALIZED"
	ServerOperationalStatusUnknown     ServerOperationalStatusValue = "UNKNOWN"
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
	ID                string                       `json:"ID,omitempty" cql:"id,primarykey,omitempty"`
	Address           string                       `json:"address,omitempty" cql:"address,omitempty"`
	Annotation        map[string]string            `json:"annotation,omitempty" cql:"annotation,omitempty"`
	AssociatedTags    []string                     `json:"associatedTags,omitempty" cql:"associatedtags,omitempty"`
	CreatedAt         time.Time                    `json:"createdAt,omitempty" cql:"createdat,omitempty"`
	Deleted           bool                         `json:"-" cql:"deleted,omitempty"`
	Description       string                       `json:"description,omitempty" cql:"description,omitempty"`
	Domain            string                       `json:"domain,omitempty" cql:"domain,omitempty"`
	Environment       ServerEnvironmentValue       `json:"environment,omitempty" cql:"environment,omitempty"`
	Name              string                       `json:"name,omitempty" cql:"name,omitempty"`
	Namespace         string                       `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`
	OperationalStatus ServerOperationalStatusValue `json:"operationalStatus,omitempty" cql:"operationalstatus,omitempty"`
	Status            enum.EntityStatus            `json:"status,omitempty" cql:"status,omitempty"`
	UpdatedAt         time.Time                    `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
}

// NewServer returns a new *Server
func NewServer() *Server {

	return &Server{

		OperationalStatus: "UNKNOWN",
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

func (o *Server) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Server) SetIdentifier(ID string) {

	o.ID = ID
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

// GetDeleted returns the deleted of the receiver
func (o *Server) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *Server) SetDeleted(deleted bool) {
	o.Deleted = deleted
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

// GetStatus returns the status of the receiver
func (o *Server) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Server) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Server) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Server) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateStringInList("environment", string(o.Environment), []string{"AWS", "GCP", "Private"}); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("operationalStatus", string(o.OperationalStatus), []string{"CONNECTED", "INITIALIZED", "UNKNOWN"}); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o Server) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return ServerAttributesMap[name]
}

var ServerAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	ServerAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	ServerAttributeNameAddress: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "address",
		Stored:         true,
		Type:           "string",
	},
	ServerAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	ServerAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	ServerAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	ServerAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	ServerAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	ServerAttributeNameDomain: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "domain",
		Stored:         true,
		Type:           "string",
	},
	ServerAttributeNameEnvironment: elemental.AttributeSpecification{
		AllowedChoices: []string{"AWS", "GCP", "Private"},
		CreationOnly:   true,
		Exposed:        true,
		Name:           "environment",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	ServerAttributeNameName: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	ServerAttributeNameNamespace: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	ServerAttributeNameOperationalStatus: elemental.AttributeSpecification{
		AllowedChoices: []string{"CONNECTED", "INITIALIZED", "UNKNOWN"},
		Exposed:        true,
		Name:           "operationalStatus",
		Stored:         true,
		Type:           "enum",
	},
	ServerAttributeNameStatus: elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Candidate", "Disabled"},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		Stored:         true,
		SubType:        "status_enum",
		Type:           "external",
	},
	ServerAttributeNameUpdatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
}
