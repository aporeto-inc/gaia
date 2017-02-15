package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/golang/constants"

// NetworkAccessPolicyIdentity represents the Identity of the object
var NetworkAccessPolicyIdentity = elemental.Identity{
	Name:     "networkaccesspolicy",
	Category: "networkaccesspolicies",
}

// NetworkAccessPoliciesList represents a list of NetworkAccessPolicies
type NetworkAccessPoliciesList []*NetworkAccessPolicy

// NetworkAccessPolicy represents the model of a networkaccesspolicy
type NetworkAccessPolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"-" bson:"-"`

	// AllowsTraffic if true, the flow will be accepted. Otherwise other actions like "logs" can still be done, but the traffic will be rejected.
	AllowsTraffic bool `json:"allowsTraffic" cql:"-" bson:"-"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" cql:"annotation,omitempty" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" cql:"associatedtags,omitempty" bson:"associatedtags"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" cql:"createdat,omitempty" bson:"createdat"`

	// Description is the description of the object.
	Description string `json:"description" cql:"description,omitempty" bson:"description"`

	// DestinationPorts contains the list of allowed ports and ranges.
	DestinationPorts []string `json:"destinationPorts" cql:"-" bson:"-"`

	// EncryptionEnabled defines if the flow has to be encrypted.
	EncryptionEnabled bool `json:"encryptionEnabled" cql:"-" bson:"-"`

	// LogsEnabled defines if the flow has to be logged.
	LogsEnabled bool `json:"logsEnabled" cql:"-" bson:"-"`

	// Name is the name of the entity
	Name string `json:"name" cql:"name,omitempty" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" cql:"normalizedtags,omitempty" bson:"normalizedtags"`

	// Object of the policy.
	Object [][]string `json:"object" cql:"-" bson:"-"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" cql:"parentid,omitempty" bson:"parentid"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" cql:"parenttype,omitempty" bson:"parenttype"`

	// Propagate will propagate the policy to all of its children.
	Propagate bool `json:"propagate" cql:"propagate,omitempty" bson:"propagate"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" cql:"protected,omitempty" bson:"protected"`

	// Status of an entity
	Status constants.EntityStatus `json:"status" cql:"status,omitempty" bson:"status"`

	// Subject of the policy.
	Subject [][]string `json:"subject" cql:"-" bson:"-"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" cql:"updatedat,omitempty" bson:"updatedat"`
}

// NewNetworkAccessPolicy returns a new *NetworkAccessPolicy
func NewNetworkAccessPolicy() *NetworkAccessPolicy {

	return &NetworkAccessPolicy{
		AssociatedTags:   []string{},
		DestinationPorts: []string{},
		NormalizedTags:   []string{},
		Propagate:        false,
		Status:           constants.Active,
	}
}

// Identity returns the Identity of the object.
func (o *NetworkAccessPolicy) Identity() elemental.Identity {

	return NetworkAccessPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NetworkAccessPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NetworkAccessPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

func (o *NetworkAccessPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *NetworkAccessPolicy) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *NetworkAccessPolicy) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *NetworkAccessPolicy) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetName returns the name of the receiver
func (o *NetworkAccessPolicy) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *NetworkAccessPolicy) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *NetworkAccessPolicy) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *NetworkAccessPolicy) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *NetworkAccessPolicy) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *NetworkAccessPolicy) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetParentID returns the parentID of the receiver
func (o *NetworkAccessPolicy) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *NetworkAccessPolicy) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *NetworkAccessPolicy) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *NetworkAccessPolicy) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetProtected returns the protected of the receiver
func (o *NetworkAccessPolicy) GetProtected() bool {
	return o.Protected
}

// GetStatus returns the status of the receiver
func (o *NetworkAccessPolicy) GetStatus() constants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *NetworkAccessPolicy) SetStatus(status constants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *NetworkAccessPolicy) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *NetworkAccessPolicy) Validate() error {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (NetworkAccessPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return NetworkAccessPolicyAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (NetworkAccessPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NetworkAccessPolicyAttributesMap
}

// NetworkAccessPolicyAttributesMap represents the map of attribute for NetworkAccessPolicy.
var NetworkAccessPolicyAttributesMap = map[string]elemental.AttributeSpecification{
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
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"AllowsTraffic": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AllowsTraffic if true, the flow will be accepted. Otherwise other actions like "logs" can still be done, but the traffic will be rejected.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsTraffic",
		Orderable:      true,
		Type:           "boolean",
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
	"DestinationPorts": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `DestinationPorts contains the list of allowed ports and ranges.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "destinationPorts",
		Orderable:      true,
		SubType:        "ports_list",
		Type:           "external",
	},
	"EncryptionEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `EncryptionEnabled defines if the flow has to be encrypted.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "encryptionEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"LogsEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LogsEnabled defines if the flow has to be logged.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "logsEnabled",
		Orderable:      true,
		Type:           "boolean",
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
	"Object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Object of the policy.`,
		Exposed:        true,
		Name:           "object",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
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
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Propagate will propagate the policy to all of its children. `,
		Exposed:        true,
		Filterable:     true,
		Name:           "propagate",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
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
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Subject of the policy.`,
		Exposed:        true,
		Name:           "subject",
		Orderable:      true,
		SubType:        "policies_list",
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
