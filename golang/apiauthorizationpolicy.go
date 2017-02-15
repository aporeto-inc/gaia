package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/golang/constants"

// APIAuthorizationPolicyIdentity represents the Identity of the object
var APIAuthorizationPolicyIdentity = elemental.Identity{
	Name:     "apiauthorizationpolicy",
	Category: "apiauthorizationpolicies",
}

// APIAuthorizationPoliciesList represents a list of APIAuthorizationPolicies
type APIAuthorizationPoliciesList []*APIAuthorizationPolicy

// APIAuthorizationPolicy represents the model of a apiauthorizationpolicy
type APIAuthorizationPolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"-" bson:"-"`

	// AllowsDelete defines if DELETE request is authorized.
	AllowsDelete bool `json:"allowsDelete" cql:"-" bson:"-"`

	// AllowsGet defines if GET request is authorized.
	AllowsGet bool `json:"allowsGet" cql:"-" bson:"-"`

	// AllowsHead defines if HEAD request is authorized.
	AllowsHead bool `json:"allowsHead" cql:"-" bson:"-"`

	// AllowsPatch defines if PATCH request is authorized.
	AllowsPatch bool `json:"allowsPatch" cql:"-" bson:"-"`

	// AllowsPost defines if POST request is authorized.
	AllowsPost bool `json:"allowsPost" cql:"-" bson:"-"`

	// AllowsPut defines if PUT request is authorized.
	AllowsPut bool `json:"allowsPut" cql:"-" bson:"-"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" cql:"annotation,omitempty" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" cql:"associatedtags,omitempty" bson:"associatedtags"`

	// AuthorizedIdentities defines the list of api identities the policy applies to.
	AuthorizedIdentities []string `json:"authorizedIdentities" cql:"-" bson:"-"`

	// AuthorizedNamespace defines on what namespace the policy applies.
	AuthorizedNamespace string `json:"authorizedNamespace" cql:"-" bson:"-"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" cql:"createdat,omitempty" bson:"createdat"`

	// Description is the description of the object.
	Description string `json:"description" cql:"description,omitempty" bson:"description"`

	// Name is the name of the entity
	Name string `json:"name" cql:"name,omitempty" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" cql:"normalizedtags,omitempty" bson:"normalizedtags"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" cql:"parentid,omitempty" bson:"parentid"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" cql:"parenttype,omitempty" bson:"parenttype"`

	// Propagate defines if the policy should propagate.
	Propagate bool `json:"propagate" cql:"-" bson:"-"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" cql:"protected,omitempty" bson:"protected"`

	// Status of an entity
	Status constants.EntityStatus `json:"status" cql:"status,omitempty" bson:"status"`

	// Subject is the subject.
	Subject [][]string `json:"subject" cql:"-" bson:"-"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" cql:"updatedat,omitempty" bson:"updatedat"`
}

// NewAPIAuthorizationPolicy returns a new *APIAuthorizationPolicy
func NewAPIAuthorizationPolicy() *APIAuthorizationPolicy {

	return &APIAuthorizationPolicy{
		AssociatedTags:       []string{},
		AuthorizedIdentities: []string{},
		NormalizedTags:       []string{},
		Status:               constants.Active,
	}
}

// Identity returns the Identity of the object.
func (o *APIAuthorizationPolicy) Identity() elemental.Identity {

	return APIAuthorizationPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *APIAuthorizationPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *APIAuthorizationPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

func (o *APIAuthorizationPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *APIAuthorizationPolicy) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *APIAuthorizationPolicy) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *APIAuthorizationPolicy) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetName returns the name of the receiver
func (o *APIAuthorizationPolicy) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *APIAuthorizationPolicy) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *APIAuthorizationPolicy) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *APIAuthorizationPolicy) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *APIAuthorizationPolicy) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *APIAuthorizationPolicy) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetParentID returns the parentID of the receiver
func (o *APIAuthorizationPolicy) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *APIAuthorizationPolicy) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *APIAuthorizationPolicy) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *APIAuthorizationPolicy) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetProtected returns the protected of the receiver
func (o *APIAuthorizationPolicy) GetProtected() bool {
	return o.Protected
}

// GetStatus returns the status of the receiver
func (o *APIAuthorizationPolicy) GetStatus() constants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *APIAuthorizationPolicy) SetStatus(status constants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *APIAuthorizationPolicy) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *APIAuthorizationPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("authorizedIdentities", o.AuthorizedIdentities); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("authorizedIdentities", o.AuthorizedIdentities); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("authorizedNamespace", o.AuthorizedNamespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("authorizedNamespace", o.AuthorizedNamespace); err != nil {
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
func (APIAuthorizationPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return APIAuthorizationPolicyAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (APIAuthorizationPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return APIAuthorizationPolicyAttributesMap
}

// APIAuthorizationPolicyAttributesMap represents the map of attribute for APIAuthorizationPolicy.
var APIAuthorizationPolicyAttributesMap = map[string]elemental.AttributeSpecification{
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
	"AllowsDelete": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AllowsDelete defines if DELETE request is authorized.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsDelete",
		Orderable:      true,
		Type:           "boolean",
	},
	"AllowsGet": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AllowsGet defines if GET request is authorized.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsGet",
		Orderable:      true,
		Type:           "boolean",
	},
	"AllowsHead": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AllowsHead defines if HEAD request is authorized.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsHead",
		Orderable:      true,
		Type:           "boolean",
	},
	"AllowsPatch": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AllowsPatch defines if PATCH request is authorized.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsPatch",
		Orderable:      true,
		Type:           "boolean",
	},
	"AllowsPost": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AllowsPost defines if POST request is authorized.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsPost",
		Orderable:      true,
		Type:           "boolean",
	},
	"AllowsPut": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AllowsPut defines if PUT request is authorized.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsPut",
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
	"AuthorizedIdentities": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AuthorizedIdentities defines the list of api identities the policy applies to. `,
		Exposed:        true,
		Name:           "authorizedIdentities",
		Required:       true,
		SubType:        "identity_list",
		Type:           "external",
	},
	"AuthorizedNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AuthorizedNamespace defines on what namespace the policy applies.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "authorizedNamespace",
		Required:       true,
		Type:           "string",
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
		Description:    `Propagate defines if the policy should propagate.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "propagate",
		Orderable:      true,
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
		Description:    `Subject is the subject.`,
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
