package gaia

import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	ProcessingUnitAttributeNameID                elemental.AttributeSpecificationNameKey = "processingunit/ID"
	ProcessingUnitAttributeNameAnnotation        elemental.AttributeSpecificationNameKey = "processingunit/annotation"
	ProcessingUnitAttributeNameAssociatedTags    elemental.AttributeSpecificationNameKey = "processingunit/associatedTags"
	ProcessingUnitAttributeNameCreatedAt         elemental.AttributeSpecificationNameKey = "processingunit/createdAt"
	ProcessingUnitAttributeNameDeleted           elemental.AttributeSpecificationNameKey = "processingunit/deleted"
	ProcessingUnitAttributeNameDescription       elemental.AttributeSpecificationNameKey = "processingunit/description"
	ProcessingUnitAttributeNameLastSyncTime      elemental.AttributeSpecificationNameKey = "processingunit/lastSyncTime"
	ProcessingUnitAttributeNameMetadata          elemental.AttributeSpecificationNameKey = "processingunit/metadata"
	ProcessingUnitAttributeNameName              elemental.AttributeSpecificationNameKey = "processingunit/name"
	ProcessingUnitAttributeNameNamespace         elemental.AttributeSpecificationNameKey = "processingunit/namespace"
	ProcessingUnitAttributeNameNativeContextID   elemental.AttributeSpecificationNameKey = "processingunit/nativeContextID"
	ProcessingUnitAttributeNameOperationalStatus elemental.AttributeSpecificationNameKey = "processingunit/operationalStatus"
	ProcessingUnitAttributeNameOwner             elemental.AttributeSpecificationNameKey = "processingunit/owner"
	ProcessingUnitAttributeNamePolicyState       elemental.AttributeSpecificationNameKey = "processingunit/policyState"
	ProcessingUnitAttributeNameServerID          elemental.AttributeSpecificationNameKey = "processingunit/serverID"
	ProcessingUnitAttributeNameStatus            elemental.AttributeSpecificationNameKey = "processingunit/status"
	ProcessingUnitAttributeNameType              elemental.AttributeSpecificationNameKey = "processingunit/type"
	ProcessingUnitAttributeNameUpdatedAt         elemental.AttributeSpecificationNameKey = "processingunit/updatedAt"
)

// ProcessingUnitOperationalStatusValue represents the possible values for attribute "operationalStatus".
type ProcessingUnitOperationalStatusValue string

const (
	ProcessingUnitOperationalStatusActive ProcessingUnitOperationalStatusValue = "Active"
	ProcessingUnitOperationalStatusDead   ProcessingUnitOperationalStatusValue = "Dead"
)

// ProcessingUnitTypeValue represents the possible values for attribute "type".
type ProcessingUnitTypeValue string

const (
	ProcessingUnitTypeDocker       ProcessingUnitTypeValue = "Docker"
	ProcessingUnitTypeLinuxservice ProcessingUnitTypeValue = "LinuxService"
	ProcessingUnitTypeRkt          ProcessingUnitTypeValue = "RKT"
)

// ProcessingUnitPolicyStateValue represents the possible values for attribute "policyState".
type ProcessingUnitPolicyStateValue string

const (
	ProcessingUnitPolicyStateDirty        ProcessingUnitPolicyStateValue = "Dirty"
	ProcessingUnitPolicyStateSynchronized ProcessingUnitPolicyStateValue = "Synchronized"
)

// ProcessingUnitIdentity represents the Identity of the object
var ProcessingUnitIdentity = elemental.Identity{
	Name:     "processingunit",
	Category: "processingunits",
}

// ProcessingUnitsList represents a list of ProcessingUnits
type ProcessingUnitsList []*ProcessingUnit

// ProcessingUnit represents the model of a processingunit
type ProcessingUnit struct {
	ID                string                               `json:"ID,omitempty" cql:"id,primarykey,omitempty"`
	Annotation        map[string]string                    `json:"annotation,omitempty" cql:"annotation,omitempty"`
	AssociatedTags    []string                             `json:"associatedTags,omitempty" cql:"associatedtags,omitempty"`
	CreatedAt         time.Time                            `json:"createdAt,omitempty" cql:"createdat,omitempty"`
	Deleted           bool                                 `json:"-" cql:"deleted,omitempty"`
	Description       string                               `json:"description,omitempty" cql:"description,omitempty"`
	LastSyncTime      time.Time                            `json:"lastSyncTime,omitempty" cql:"lastsynctime,omitempty"`
	Metadata          []string                             `json:"metadata,omitempty" cql:"metadata,omitempty"`
	Name              string                               `json:"name,omitempty" cql:"name,omitempty"`
	Namespace         string                               `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`
	NativeContextID   string                               `json:"nativeContextID,omitempty" cql:"nativecontextid,primarykey,omitempty"`
	OperationalStatus ProcessingUnitOperationalStatusValue `json:"operationalStatus,omitempty" cql:"operationalstatus,omitempty"`
	Owner             []string                             `json:"owner,omitempty" cql:"owner,omitempty"`
	PolicyState       ProcessingUnitPolicyStateValue       `json:"policyState,omitempty" cql:"policystate,omitempty"`
	ServerID          string                               `json:"serverID,omitempty" cql:"serverid,primarykey,omitempty"`
	Status            enum.EntityStatus                    `json:"status,omitempty" cql:"status,omitempty"`
	Type              ProcessingUnitTypeValue              `json:"type,omitempty" cql:"type,omitempty"`
	UpdatedAt         time.Time                            `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
}

// NewProcessingUnit returns a new *ProcessingUnit
func NewProcessingUnit() *ProcessingUnit {

	return &ProcessingUnit{}
}

// Identity returns the Identity of the object.
func (o *ProcessingUnit) Identity() elemental.Identity {

	return ProcessingUnitIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ProcessingUnit) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ProcessingUnit) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *ProcessingUnit) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *ProcessingUnit) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *ProcessingUnit) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *ProcessingUnit) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *ProcessingUnit) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *ProcessingUnit) GetName() string {
	return o.Name
}

// GetNamespace returns the namespace of the receiver
func (o *ProcessingUnit) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *ProcessingUnit) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetStatus returns the status of the receiver
func (o *ProcessingUnit) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *ProcessingUnit) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *ProcessingUnit) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *ProcessingUnit) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("nativeContextID", o.NativeContextID); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("operationalStatus", string(o.OperationalStatus), []string{"Active", "Dead"}); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("policyState", string(o.PolicyState), []string{"Dirty", "Synchronized"}); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("serverID", o.ServerID); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Docker", "LinuxService", "RKT"}); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o ProcessingUnit) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return ProcessingUnitAttributesMap[name]
}

var ProcessingUnitAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	ProcessingUnitAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
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
	ProcessingUnitAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	ProcessingUnitAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tag_list",
		Type:           "external",
	},
	ProcessingUnitAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	ProcessingUnitAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Filterable:     true,
		ForeignKey:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	ProcessingUnitAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		Stored:         true,
		Type:           "string",
	},
	ProcessingUnitAttributeNameLastSyncTime: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	ProcessingUnitAttributeNameMetadata: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "metadata",
		Required:       true,
		Stored:         true,
		SubType:        "metadata_list",
		Type:           "external",
	},
	ProcessingUnitAttributeNameName: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	ProcessingUnitAttributeNameNamespace: elemental.AttributeSpecification{
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
	ProcessingUnitAttributeNameNativeContextID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Name:           "nativeContextID",
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	ProcessingUnitAttributeNameOperationalStatus: elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Dead"},
		CreationOnly:   true,
		Exposed:        true,
		Name:           "operationalStatus",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	ProcessingUnitAttributeNameOwner: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "owner",
		Stored:         true,
		SubType:        "tag_list",
		Type:           "external",
	},
	ProcessingUnitAttributeNamePolicyState: elemental.AttributeSpecification{
		AllowedChoices: []string{"Dirty", "Synchronized"},
		CreationOnly:   true,
		Exposed:        true,
		Name:           "policyState",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	ProcessingUnitAttributeNameServerID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Name:           "serverID",
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	ProcessingUnitAttributeNameStatus: elemental.AttributeSpecification{
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
	ProcessingUnitAttributeNameType: elemental.AttributeSpecification{
		AllowedChoices: []string{"Docker", "LinuxService", "RKT"},
		CreationOnly:   true,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	ProcessingUnitAttributeNameUpdatedAt: elemental.AttributeSpecification{
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
