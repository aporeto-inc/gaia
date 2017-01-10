package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/golang/constants"

// ProcessingUnitOperationalStatusValue represents the possible values for attribute "operationalStatus".
type ProcessingUnitOperationalStatusValue string

const (
	// ProcessingUnitOperationalStatusInitialized represents the value Initialized.
	ProcessingUnitOperationalStatusInitialized ProcessingUnitOperationalStatusValue = "Initialized"

	// ProcessingUnitOperationalStatusPaused represents the value Paused.
	ProcessingUnitOperationalStatusPaused ProcessingUnitOperationalStatusValue = "Paused"

	// ProcessingUnitOperationalStatusRunning represents the value Running.
	ProcessingUnitOperationalStatusRunning ProcessingUnitOperationalStatusValue = "Running"

	// ProcessingUnitOperationalStatusStopped represents the value Stopped.
	ProcessingUnitOperationalStatusStopped ProcessingUnitOperationalStatusValue = "Stopped"

	// ProcessingUnitOperationalStatusTerminated represents the value Terminated.
	ProcessingUnitOperationalStatusTerminated ProcessingUnitOperationalStatusValue = "Terminated"
)

// ProcessingUnitTypeValue represents the possible values for attribute "type".
type ProcessingUnitTypeValue string

const (
	// ProcessingUnitTypeDocker represents the value Docker.
	ProcessingUnitTypeDocker ProcessingUnitTypeValue = "Docker"

	// ProcessingUnitTypeLinuxservice represents the value LinuxService.
	ProcessingUnitTypeLinuxservice ProcessingUnitTypeValue = "LinuxService"

	// ProcessingUnitTypeRkt represents the value RKT.
	ProcessingUnitTypeRkt ProcessingUnitTypeValue = "RKT"
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
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"id,primarykey,omitempty" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" cql:"annotation,omitempty" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" cql:"associatedtags,omitempty" bson:"associatedtags"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" cql:"createdat,omitempty" bson:"createdat"`

	// Deleted marks if the entity has been deleted.
	Deleted bool `json:"-" cql:"deleted,omitempty" bson:"deleted"`

	// Description is the description of the object.
	Description string `json:"description" cql:"description,omitempty" bson:"description"`

	// LastSyncTime is the time when the policy was last resolved
	LastSyncTime time.Time `json:"lastSyncTime" cql:"lastsynctime,omitempty" bson:"lastsynctime"`

	// Metadata are list of tags associated to the processing unit
	Metadata []string `json:"metadata" cql:"metadata,omitempty" bson:"metadata"`

	// Name is the name of the entity
	Name string `json:"name" cql:"name,omitempty" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty" bson:"_namespace"`

	// NativeContextID is the Docker UUID or service PID
	NativeContextID string `json:"nativeContextID" cql:"nativecontextid,omitempty" bson:"nativecontextid"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" cql:"-" bson:"-"`

	// OperationalStatus of the processing unit
	OperationalStatus ProcessingUnitOperationalStatusValue `json:"operationalStatus" cql:"operationalstatus,omitempty" bson:"operationalstatus"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" cql:"parentid,omitempty" bson:"parentid"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" cql:"parenttype,omitempty" bson:"parenttype"`

	// serverID is the ID of the server associated with the processing unit
	ServerID string `json:"serverID" cql:"serverid,omitempty" bson:"serverid"`

	// Status of an entity
	Status constants.EntityStatus `json:"status" cql:"status,omitempty" bson:"status"`

	// Type of the container ecosystem
	Type ProcessingUnitTypeValue `json:"type" cql:"type,omitempty" bson:"type"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" cql:"updatedat,omitempty" bson:"updatedat"`

	// Vulnerabilities contains the list of vulnerabilities of the processing unit.
	Vulnerabilities []string `json:"-" cql:"vulnerabilities,omitempty" bson:"vulnerabilities"`
}

// NewProcessingUnit returns a new *ProcessingUnit
func NewProcessingUnit() *ProcessingUnit {

	return &ProcessingUnit{
		AssociatedTags:    []string{},
		NormalizedTags:    []string{},
		OperationalStatus: "Initialized",
		Status:            constants.Active,
	}
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

func (o *ProcessingUnit) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
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

// SetName set the given name of the receiver
func (o *ProcessingUnit) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *ProcessingUnit) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *ProcessingUnit) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *ProcessingUnit) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *ProcessingUnit) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetParentID returns the parentID of the receiver
func (o *ProcessingUnit) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *ProcessingUnit) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *ProcessingUnit) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *ProcessingUnit) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetStatus returns the status of the receiver
func (o *ProcessingUnit) GetStatus() constants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *ProcessingUnit) SetStatus(status constants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *ProcessingUnit) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *ProcessingUnit) Validate() error {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("operationalStatus", string(o.OperationalStatus), []string{"Initialized", "Paused", "Running", "Stopped", "Terminated"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Docker", "LinuxService", "RKT"}, false); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (ProcessingUnit) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ProcessingUnitAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (ProcessingUnit) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ProcessingUnitAttributesMap
}

// ProcessingUnitAttributesMap represents the map of attribute for ProcessingUnit.
var ProcessingUnitAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Deleted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Deleted marks if the entity has been deleted.`,
		Getter:         true,
		Name:           "deleted",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
		Autogenerated:  true,
		Description:    `LastSyncTime is the time when the policy was last resolved`,
		Exposed:        true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Metadata are list of tags associated to the processing unit`,
		Exposed:        true,
		Filterable:     true,
		Name:           "metadata",
		Required:       true,
		Stored:         true,
		SubType:        "metadata_list",
		Type:           "external",
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
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"NativeContextID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `NativeContextID is the Docker UUID or service PID`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "nativeContextID",
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
		SubType:        "tags_list",
		Type:           "external",
	},
	"OperationalStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Initialized", "Paused", "Running", "Stopped", "Terminated"},
		Description:    `OperationalStatus of the processing unit`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		Stored:         true,
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
	"ServerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `serverID is the ID of the server associated with the processing unit`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "serverID",
		Stored:         true,
		Type:           "string",
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
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Docker", "LinuxService", "RKT"},
		CreationOnly:   true,
		Description:    `Type of the container ecosystem`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
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
	"Vulnerabilities": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Vulnerabilities contains the list of vulnerabilities of the processing unit.`,
		Name:           "vulnerabilities",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}
