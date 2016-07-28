package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	PolicyAttributeNameDescription    elemental.AttributeSpecificationNameKey = "policy/Description"
	PolicyAttributeNameID             elemental.AttributeSpecificationNameKey = "policy/ID"
	PolicyAttributeNameAction         elemental.AttributeSpecificationNameKey = "policy/action"
	PolicyAttributeNameAllObjectTags  elemental.AttributeSpecificationNameKey = "policy/allObjectTags"
	PolicyAttributeNameAllSubjectTags elemental.AttributeSpecificationNameKey = "policy/allSubjectTags"
	PolicyAttributeNameAnnotation     elemental.AttributeSpecificationNameKey = "policy/annotation"
	PolicyAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "policy/associatedTags"
	PolicyAttributeNameCreatedAt      elemental.AttributeSpecificationNameKey = "policy/createdAt"
	PolicyAttributeNameDeleted        elemental.AttributeSpecificationNameKey = "policy/deleted"
	PolicyAttributeNameName           elemental.AttributeSpecificationNameKey = "policy/name"
	PolicyAttributeNameNamespace      elemental.AttributeSpecificationNameKey = "policy/namespace"
	PolicyAttributeNameObject         elemental.AttributeSpecificationNameKey = "policy/object"
	PolicyAttributeNameRelation       elemental.AttributeSpecificationNameKey = "policy/relation"
	PolicyAttributeNameStatus         elemental.AttributeSpecificationNameKey = "policy/status"
	PolicyAttributeNameSubject        elemental.AttributeSpecificationNameKey = "policy/subject"
	PolicyAttributeNameType           elemental.AttributeSpecificationNameKey = "policy/type"
	PolicyAttributeNameUpdatedAt      elemental.AttributeSpecificationNameKey = "policy/updatedAt"
)

// PolicyIdentity represents the Identity of the object
var PolicyIdentity = elemental.Identity{
	Name:     "policy",
	Category: "policies",
}

// PoliciesList represents a list of Policies
type PoliciesList []*Policy

// Policy represents the model of a policy
type Policy struct {
	Description    string                       `json:"Description,omitempty" cql:"description,omitempty"`
	ID             string                       `json:"ID,omitempty" cql:"id,primarykey,omitempty"`
	Action         map[string]map[string]string `json:"action,omitempty" cql:"action,omitempty"`
	AllObjectTags  []string                     `json:"-" cql:"allobjecttags,omitempty"`
	AllSubjectTags []string                     `json:"-" cql:"allsubjecttags,omitempty"`
	Annotation     map[string]string            `json:"annotation,omitempty" cql:"annotation,omitempty"`
	AssociatedTags []string                     `json:"associatedTags,omitempty" cql:"associatedtags,omitempty"`
	CreatedAt      time.Time                    `json:"createdAt,omitempty" cql:"createdat,omitempty"`
	Deleted        bool                         `json:"-" cql:"deleted,omitempty"`
	Name           string                       `json:"name,omitempty" cql:"name,omitempty"`
	Namespace      string                       `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`
	Object         [][]string                   `json:"object,omitempty" cql:"object,omitempty"`
	Relation       []string                     `json:"relation,omitempty" cql:"relation,omitempty"`
	Status         enum.EntityStatus            `json:"status,omitempty" cql:"status,omitempty"`
	Subject        [][]string                   `json:"subject,omitempty" cql:"subject,omitempty"`
	Type           enum.PolicyType              `json:"type,omitempty" cql:"type,primarykey,omitempty"`
	UpdatedAt      time.Time                    `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
}

// NewPolicy returns a new *Policy
func NewPolicy() *Policy {

	return &Policy{}
}

// Identity returns the Identity of the object.
func (o *Policy) Identity() elemental.Identity {

	return PolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Policy) Identifier() string {

	return o.ID
}

func (o *Policy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Policy) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Policy) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Policy) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Policy) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *Policy) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *Policy) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *Policy) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *Policy) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *Policy) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Policy) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetStatus returns the status of the receiver
func (o *Policy) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Policy) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Policy) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Policy) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"APIAuthorization", "ExtendTags", "File", "NamespaceMapping", "Network", "Statistics", "Syscall"}); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o Policy) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return PolicyAttributesMap[name]
}

var PolicyAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	PolicyAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "Description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	PolicyAttributeNameID: elemental.AttributeSpecification{
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
	PolicyAttributeNameAction: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Stored:         true,
		SubType:        "actions_list",
		Type:           "external",
	},
	PolicyAttributeNameAllObjectTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Name:           "allObjectTags",
		Required:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	PolicyAttributeNameAllSubjectTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Name:           "allSubjectTags",
		Required:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	PolicyAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	PolicyAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	PolicyAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	PolicyAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	PolicyAttributeNameName: elemental.AttributeSpecification{
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
	PolicyAttributeNameNamespace: elemental.AttributeSpecification{
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
	PolicyAttributeNameObject: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "object",
		Stored:         true,
		SubType:        "policies_list",
		Type:           "external",
	},
	PolicyAttributeNameRelation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "relation",
		Required:       true,
		Stored:         true,
		SubType:        "relations_list",
		Type:           "external",
	},
	PolicyAttributeNameStatus: elemental.AttributeSpecification{
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
	PolicyAttributeNameSubject: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "subject",
		Required:       true,
		Stored:         true,
		SubType:        "policies_list",
		Type:           "external",
	},
	PolicyAttributeNameType: elemental.AttributeSpecification{
		AllowedChoices: []string{"APIAuthorization", "ExtendTags", "File", "NamespaceMapping", "Network", "Statistics", "Syscall"},
		Exposed:        true,
		Format:         "free",
		Name:           "type",
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		SubType:        "policytype_enum",
		Type:           "external",
	},
	PolicyAttributeNameUpdatedAt: elemental.AttributeSpecification{
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
