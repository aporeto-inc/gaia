package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/golang/constants"

// IntegrationTypeValue represents the possible values for attribute "type".
type IntegrationTypeValue string

const (
	// IntegrationTypeRegistry represents the value Registry.
	IntegrationTypeRegistry IntegrationTypeValue = "Registry"

	// IntegrationTypeVulnerabilityscanner represents the value VulnerabilityScanner.
	IntegrationTypeVulnerabilityscanner IntegrationTypeValue = "VulnerabilityScanner"
)

// IntegrationAuthorizationValue represents the possible values for attribute "authorization".
type IntegrationAuthorizationValue string

const (
	// IntegrationAuthorizationBasic represents the value Basic.
	IntegrationAuthorizationBasic IntegrationAuthorizationValue = "Basic"

	// IntegrationAuthorizationOauth represents the value OAuth.
	IntegrationAuthorizationOauth IntegrationAuthorizationValue = "OAuth"
)

// IntegrationIdentity represents the Identity of the object
var IntegrationIdentity = elemental.Identity{
	Name:     "integration",
	Category: "integrations",
}

// IntegrationsList represents a list of Integrations
type IntegrationsList []*Integration

// Integration represents the model of a integration
type Integration struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"id,primarykey,omitempty" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" cql:"annotation,omitempty" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" cql:"associatedtags,omitempty" bson:"associatedtags"`

	// Authorization refers to type of the HTTP authorization header
	Authorization IntegrationAuthorizationValue `json:"authorization" cql:"authorization,omitempty" bson:"authorization"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" cql:"createdat,omitempty" bson:"createdat"`

	// Deleted marks if the entity has been deleted.
	Deleted bool `json:"-" cql:"deleted,omitempty" bson:"deleted"`

	// Endpoint is the API end point of the service
	Endpoint string `json:"endpoint" cql:"endpoint,omitempty" bson:"endpoint"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty" bson:"_namespace"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" cql:"parentid,omitempty" bson:"parentid"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" cql:"parenttype,omitempty" bson:"parenttype"`

	// Password is the password of the user to be used in the HTTP Authorization header
	Password string `json:"password" cql:"password,omitempty" bson:"password"`

	// Port is the port number of the service
	Port int `json:"port" cql:"port,omitempty" bson:"port"`

	// Server is either the DNS name or IP of the server that provides the service
	Server string `json:"server" cql:"server,omitempty" bson:"server"`

	// SSLEnabled defines if the service is either secured or unsecured
	SslEnabled bool `json:"sslEnabled" cql:"sslenabled,omitempty" bson:"sslenabled"`

	// Status of an entity
	Status constants.EntityStatus `json:"status" cql:"status,omitempty" bson:"status"`

	// Type refers to type of the server
	Type IntegrationTypeValue `json:"type" cql:"type,omitempty" bson:"type"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" cql:"updatedat,omitempty" bson:"updatedat"`

	// Username refers to the username to be used in the HTTP Authorization header
	UserName string `json:"userName" cql:"username,omitempty" bson:"username"`
}

// NewIntegration returns a new *Integration
func NewIntegration() *Integration {

	return &Integration{
		SslEnabled: false,
		Status:     constants.Active,
		Type:       "Registry",
	}
}

// Identity returns the Identity of the object.
func (o *Integration) Identity() elemental.Identity {

	return IntegrationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Integration) Identifier() string {

	return o.ID
}

func (o *Integration) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Integration) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Integration) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Integration) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Integration) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *Integration) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *Integration) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetNamespace returns the namespace of the receiver
func (o *Integration) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Integration) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetParentID returns the parentID of the receiver
func (o *Integration) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *Integration) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *Integration) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *Integration) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetServer returns the server of the receiver
func (o *Integration) GetServer() string {
	return o.Server
}

// GetStatus returns the status of the receiver
func (o *Integration) GetStatus() constants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Integration) SetStatus(status constants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Integration) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Integration) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateStringInList("authorization", string(o.Authorization), []string{"Basic", "OAuth"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("port", o.Port, 65535, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("port", o.Port, 1, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("server", o.Server); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Registry", "VulnerabilityScanner"}, false); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o Integration) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return IntegrationAttributesMap[name]
}

// IntegrationAttributesMap represents the map of attribute for Integration.
var IntegrationAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"Authorization": elemental.AttributeSpecification{
		AllowedChoices: []string{"Basic", "OAuth"},
		Exposed:        true,
		Filterable:     true,
		Name:           "authorization",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
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
	"Password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "password",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Port": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		MaxValue:       65535,
		MinValue:       1,
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
		Getter:         true,
		Name:           "server",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"SslEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "sslEnabled",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "boolean",
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
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Registry", "VulnerabilityScanner"},
		Exposed:        true,
		Name:           "type",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "enum",
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
	"UserName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "userName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}
