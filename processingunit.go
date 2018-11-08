package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ProcessingUnitEnforcementStatusValue represents the possible values for attribute "enforcementStatus".
type ProcessingUnitEnforcementStatusValue string

const (
	// ProcessingUnitEnforcementStatusActive represents the value Active.
	ProcessingUnitEnforcementStatusActive ProcessingUnitEnforcementStatusValue = "Active"

	// ProcessingUnitEnforcementStatusFailed represents the value Failed.
	ProcessingUnitEnforcementStatusFailed ProcessingUnitEnforcementStatusValue = "Failed"

	// ProcessingUnitEnforcementStatusInactive represents the value Inactive.
	ProcessingUnitEnforcementStatusInactive ProcessingUnitEnforcementStatusValue = "Inactive"
)

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
	// ProcessingUnitTypeAPIGateway represents the value APIGateway.
	ProcessingUnitTypeAPIGateway ProcessingUnitTypeValue = "APIGateway"

	// ProcessingUnitTypeDocker represents the value Docker.
	ProcessingUnitTypeDocker ProcessingUnitTypeValue = "Docker"

	// ProcessingUnitTypeLinuxService represents the value LinuxService.
	ProcessingUnitTypeLinuxService ProcessingUnitTypeValue = "LinuxService"

	// ProcessingUnitTypeRKT represents the value RKT.
	ProcessingUnitTypeRKT ProcessingUnitTypeValue = "RKT"

	// ProcessingUnitTypeUser represents the value User.
	ProcessingUnitTypeUser ProcessingUnitTypeValue = "User"
)

// ProcessingUnitIdentity represents the Identity of the object.
var ProcessingUnitIdentity = elemental.Identity{
	Name:     "processingunit",
	Category: "processingunits",
	Package:  "squall",
	Private:  false,
}

// ProcessingUnitsList represents a list of ProcessingUnits
type ProcessingUnitsList []*ProcessingUnit

// Identity returns the identity of the objects in the list.
func (o ProcessingUnitsList) Identity() elemental.Identity {

	return ProcessingUnitIdentity
}

// Copy returns a pointer to a copy the ProcessingUnitsList.
func (o ProcessingUnitsList) Copy() elemental.Identifiables {

	copy := append(ProcessingUnitsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ProcessingUnitsList.
func (o ProcessingUnitsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ProcessingUnitsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ProcessingUnit))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ProcessingUnitsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ProcessingUnitsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the ProcessingUnitsList converted to SparseProcessingUnitsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ProcessingUnitsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o ProcessingUnitsList) Version() int {

	return 1
}

// ProcessingUnit represents the model of a processingunit
type ProcessingUnit struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Archived defines if the object is archived.
	Archived bool `json:"-" bson:"archived" mapstructure:"-,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// EnforcementStatus communicates the state of the enforcer for that PU.
	EnforcementStatus ProcessingUnitEnforcementStatusValue `json:"enforcementStatus" bson:"enforcementstatus" mapstructure:"enforcementStatus,omitempty"`

	// EnforcerID is the ID of the enforcer associated with the processing unit.
	EnforcerID string `json:"enforcerID" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Docker image, or path to executable.
	Image string `json:"image" bson:"image" mapstructure:"image,omitempty"`

	// Last poke is the time when the pu got last poked.
	LastPokeTime time.Time `json:"-" bson:"lastpoketime" mapstructure:"-,omitempty"`

	// LastSyncTime is the time when the policy was last resolved.
	LastSyncTime time.Time `json:"lastSyncTime" bson:"lastsynctime" mapstructure:"lastSyncTime,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NativeContextID is the Docker UUID or service PID.
	NativeContextID string `json:"nativeContextID" bson:"nativecontextid" mapstructure:"nativeContextID,omitempty"`

	// NetworkServices is the list of services that this processing unit has declared
	// that it will be listening to. This can happen either with an activation command
	// or by exposing the ports in a container manifest.
	NetworkServices []*ProcessingUnitService `json:"networkServices" bson:"networkservices" mapstructure:"networkServices,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// OperationalStatus of the processing unit.
	OperationalStatus ProcessingUnitOperationalStatusValue `json:"operationalStatus" bson:"operationalstatus" mapstructure:"operationalStatus,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Type of the container ecosystem.
	Type ProcessingUnitTypeValue `json:"type" bson:"type" mapstructure:"type,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewProcessingUnit returns a new *ProcessingUnit
func NewProcessingUnit() *ProcessingUnit {

	return &ProcessingUnit{
		ModelVersion:      1,
		Annotations:       map[string][]string{},
		EnforcementStatus: ProcessingUnitEnforcementStatusInactive,
		AssociatedTags:    []string{},
		Metadata:          []string{},
		NetworkServices:   []*ProcessingUnitService{},
		NormalizedTags:    []string{},
		OperationalStatus: ProcessingUnitOperationalStatusInitialized,
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
func (o *ProcessingUnit) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *ProcessingUnit) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *ProcessingUnit) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *ProcessingUnit) Doc() string {
	return `A Processing Unit reprents anything that can compute. It can be a Docker
container, or a simple Unix process. They are created, updated and deleted by
the system as they come and go. You can only modify its tags.  Processing Units
use Network Access Policies to define which other Processing Units or External
Services they can communicate with and File Access Policies to define what File
Paths they can use.`
}

func (o *ProcessingUnit) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *ProcessingUnit) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *ProcessingUnit) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetArchived returns the Archived of the receiver.
func (o *ProcessingUnit) GetArchived() bool {

	return o.Archived
}

// SetArchived sets the property Archived of the receiver using the given value.
func (o *ProcessingUnit) SetArchived(archived bool) {

	o.Archived = archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *ProcessingUnit) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *ProcessingUnit) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *ProcessingUnit) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *ProcessingUnit) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *ProcessingUnit) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *ProcessingUnit) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *ProcessingUnit) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *ProcessingUnit) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *ProcessingUnit) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *ProcessingUnit) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *ProcessingUnit) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *ProcessingUnit) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *ProcessingUnit) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *ProcessingUnit) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *ProcessingUnit) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *ProcessingUnit) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *ProcessingUnit) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ProcessingUnit) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseProcessingUnit{
			ID:                &o.ID,
			Annotations:       &o.Annotations,
			Archived:          &o.Archived,
			AssociatedTags:    &o.AssociatedTags,
			CreateTime:        &o.CreateTime,
			Description:       &o.Description,
			EnforcementStatus: &o.EnforcementStatus,
			EnforcerID:        &o.EnforcerID,
			Image:             &o.Image,
			LastPokeTime:      &o.LastPokeTime,
			LastSyncTime:      &o.LastSyncTime,
			Metadata:          &o.Metadata,
			Name:              &o.Name,
			Namespace:         &o.Namespace,
			NativeContextID:   &o.NativeContextID,
			NetworkServices:   &o.NetworkServices,
			NormalizedTags:    &o.NormalizedTags,
			OperationalStatus: &o.OperationalStatus,
			Protected:         &o.Protected,
			Type:              &o.Type,
			UpdateTime:        &o.UpdateTime,
		}
	}

	sp := &SparseProcessingUnit{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "archived":
			sp.Archived = &(o.Archived)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "enforcementStatus":
			sp.EnforcementStatus = &(o.EnforcementStatus)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "image":
			sp.Image = &(o.Image)
		case "lastPokeTime":
			sp.LastPokeTime = &(o.LastPokeTime)
		case "lastSyncTime":
			sp.LastSyncTime = &(o.LastSyncTime)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "nativeContextID":
			sp.NativeContextID = &(o.NativeContextID)
		case "networkServices":
			sp.NetworkServices = &(o.NetworkServices)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "operationalStatus":
			sp.OperationalStatus = &(o.OperationalStatus)
		case "protected":
			sp.Protected = &(o.Protected)
		case "type":
			sp.Type = &(o.Type)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseProcessingUnit to the object.
func (o *ProcessingUnit) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseProcessingUnit)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.Archived != nil {
		o.Archived = *so.Archived
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.EnforcementStatus != nil {
		o.EnforcementStatus = *so.EnforcementStatus
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.Image != nil {
		o.Image = *so.Image
	}
	if so.LastPokeTime != nil {
		o.LastPokeTime = *so.LastPokeTime
	}
	if so.LastSyncTime != nil {
		o.LastSyncTime = *so.LastSyncTime
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NativeContextID != nil {
		o.NativeContextID = *so.NativeContextID
	}
	if so.NetworkServices != nil {
		o.NetworkServices = *so.NetworkServices
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.OperationalStatus != nil {
		o.OperationalStatus = *so.OperationalStatus
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the ProcessingUnit.
func (o *ProcessingUnit) DeepCopy() *ProcessingUnit {

	if o == nil {
		return nil
	}

	out := &ProcessingUnit{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ProcessingUnit.
func (o *ProcessingUnit) DeepCopyInto(out *ProcessingUnit) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ProcessingUnit: %s", err))
	}

	*out = *target.(*ProcessingUnit)
}

// Validate valides the current information stored into the structure.
func (o *ProcessingUnit) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("enforcementStatus", string(o.EnforcementStatus), []string{"Active", "Failed", "Inactive"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = append(errors, err)
	}

	for _, sub := range o.NetworkServices {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	if err := ValidateProcessingUnitServicesList("networkServices", o.NetworkServices); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("operationalStatus", string(o.OperationalStatus), []string{"Initialized", "Paused", "Running", "Stopped", "Terminated"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Docker", "LinuxService", "RKT", "User", "APIGateway"}, false); err != nil {
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
func (*ProcessingUnit) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ProcessingUnitAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ProcessingUnitLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ProcessingUnit) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ProcessingUnitAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ProcessingUnit) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "archived":
		return o.Archived
	case "associatedTags":
		return o.AssociatedTags
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "enforcementStatus":
		return o.EnforcementStatus
	case "enforcerID":
		return o.EnforcerID
	case "image":
		return o.Image
	case "lastPokeTime":
		return o.LastPokeTime
	case "lastSyncTime":
		return o.LastSyncTime
	case "metadata":
		return o.Metadata
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "nativeContextID":
		return o.NativeContextID
	case "networkServices":
		return o.NetworkServices
	case "normalizedTags":
		return o.NormalizedTags
	case "operationalStatus":
		return o.OperationalStatus
	case "protected":
		return o.Protected
	case "type":
		return o.Type
	case "updateTime":
		return o.UpdateTime
	}

	return nil
}

// ProcessingUnitAttributesMap represents the map of attribute for ProcessingUnit.
var ProcessingUnitAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"Archived": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Archived",
		Description:    `Archived defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcementStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Failed", "Inactive"},
		ConvertedName:  "EnforcementStatus",
		DefaultValue:   ProcessingUnitEnforcementStatusInactive,
		Description:    `EnforcementStatus communicates the state of the enforcer for that PU.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcementStatus",
		Stored:         true,
		Type:           "enum",
	},
	"EnforcerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `EnforcerID is the ID of the enforcer associated with the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcerID",
		Stored:         true,
		Type:           "string",
	},
	"Image": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Image",
		Description:    `Docker image, or path to executable.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "image",
		Stored:         true,
		Type:           "string",
	},
	"LastPokeTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastPokeTime",
		Description:    `Last poke is the time when the pu got last poked.`,
		Name:           "lastPokeTime",
		Stored:         true,
		Type:           "time",
	},
	"LastSyncTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastSyncTime",
		Description:    `LastSyncTime is the time when the policy was last resolved.`,
		Exposed:        true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "metadata_list",
		Type:       "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NativeContextID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NativeContextID",
		Description:    `NativeContextID is the Docker UUID or service PID.`,
		Exposed:        true,
		Name:           "nativeContextID",
		Stored:         true,
		Type:           "string",
	},
	"NetworkServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkServices",
		Description: `NetworkServices is the list of services that this processing unit has declared
that it will be listening to. This can happen either with an activation command
or by exposing the ports in a container manifest.`,
		Exposed:   true,
		Name:      "networkServices",
		Orderable: true,
		Stored:    true,
		SubType:   "processingunitservice",
		Type:      "refList",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
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
		AllowedChoices: []string{"Initialized", "Paused", "Running", "Stopped", "Terminated"},
		ConvertedName:  "OperationalStatus",
		DefaultValue:   ProcessingUnitOperationalStatusInitialized,
		Description:    `OperationalStatus of the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		Stored:         true,
		Type:           "enum",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Docker", "LinuxService", "RKT", "User", "APIGateway"},
		ConvertedName:  "Type",
		CreationOnly:   true,
		Description:    `Type of the container ecosystem.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}

// ProcessingUnitLowerCaseAttributesMap represents the map of attribute for ProcessingUnit.
var ProcessingUnitLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"archived": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Archived",
		Description:    `Archived defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"enforcementstatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Failed", "Inactive"},
		ConvertedName:  "EnforcementStatus",
		DefaultValue:   ProcessingUnitEnforcementStatusInactive,
		Description:    `EnforcementStatus communicates the state of the enforcer for that PU.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcementStatus",
		Stored:         true,
		Type:           "enum",
	},
	"enforcerid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `EnforcerID is the ID of the enforcer associated with the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "enforcerID",
		Stored:         true,
		Type:           "string",
	},
	"image": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Image",
		Description:    `Docker image, or path to executable.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "image",
		Stored:         true,
		Type:           "string",
	},
	"lastpoketime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastPokeTime",
		Description:    `Last poke is the time when the pu got last poked.`,
		Name:           "lastPokeTime",
		Stored:         true,
		Type:           "time",
	},
	"lastsynctime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastSyncTime",
		Description:    `LastSyncTime is the time when the policy was last resolved.`,
		Exposed:        true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "metadata_list",
		Type:       "external",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"nativecontextid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NativeContextID",
		Description:    `NativeContextID is the Docker UUID or service PID.`,
		Exposed:        true,
		Name:           "nativeContextID",
		Stored:         true,
		Type:           "string",
	},
	"networkservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkServices",
		Description: `NetworkServices is the list of services that this processing unit has declared
that it will be listening to. This can happen either with an activation command
or by exposing the ports in a container manifest.`,
		Exposed:   true,
		Name:      "networkServices",
		Orderable: true,
		Stored:    true,
		SubType:   "processingunitservice",
		Type:      "refList",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
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
	"operationalstatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Initialized", "Paused", "Running", "Stopped", "Terminated"},
		ConvertedName:  "OperationalStatus",
		DefaultValue:   ProcessingUnitOperationalStatusInitialized,
		Description:    `OperationalStatus of the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		Stored:         true,
		Type:           "enum",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Docker", "LinuxService", "RKT", "User", "APIGateway"},
		ConvertedName:  "Type",
		CreationOnly:   true,
		Description:    `Type of the container ecosystem.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}

// SparseProcessingUnitsList represents a list of SparseProcessingUnits
type SparseProcessingUnitsList []*SparseProcessingUnit

// Identity returns the identity of the objects in the list.
func (o SparseProcessingUnitsList) Identity() elemental.Identity {

	return ProcessingUnitIdentity
}

// Copy returns a pointer to a copy the SparseProcessingUnitsList.
func (o SparseProcessingUnitsList) Copy() elemental.Identifiables {

	copy := append(SparseProcessingUnitsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseProcessingUnitsList.
func (o SparseProcessingUnitsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseProcessingUnitsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseProcessingUnit))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseProcessingUnitsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseProcessingUnitsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseProcessingUnitsList converted to ProcessingUnitsList.
func (o SparseProcessingUnitsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseProcessingUnitsList) Version() int {

	return 1
}

// SparseProcessingUnit represents the sparse version of a processingunit.
type SparseProcessingUnit struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Archived defines if the object is archived.
	Archived *bool `json:"-,omitempty" bson:"archived" mapstructure:"-,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// EnforcementStatus communicates the state of the enforcer for that PU.
	EnforcementStatus *ProcessingUnitEnforcementStatusValue `json:"enforcementStatus,omitempty" bson:"enforcementstatus" mapstructure:"enforcementStatus,omitempty"`

	// EnforcerID is the ID of the enforcer associated with the processing unit.
	EnforcerID *string `json:"enforcerID,omitempty" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Docker image, or path to executable.
	Image *string `json:"image,omitempty" bson:"image" mapstructure:"image,omitempty"`

	// Last poke is the time when the pu got last poked.
	LastPokeTime *time.Time `json:"-,omitempty" bson:"lastpoketime" mapstructure:"-,omitempty"`

	// LastSyncTime is the time when the policy was last resolved.
	LastSyncTime *time.Time `json:"lastSyncTime,omitempty" bson:"lastsynctime" mapstructure:"lastSyncTime,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NativeContextID is the Docker UUID or service PID.
	NativeContextID *string `json:"nativeContextID,omitempty" bson:"nativecontextid" mapstructure:"nativeContextID,omitempty"`

	// NetworkServices is the list of services that this processing unit has declared
	// that it will be listening to. This can happen either with an activation command
	// or by exposing the ports in a container manifest.
	NetworkServices *[]*ProcessingUnitService `json:"networkServices,omitempty" bson:"networkservices" mapstructure:"networkServices,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// OperationalStatus of the processing unit.
	OperationalStatus *ProcessingUnitOperationalStatusValue `json:"operationalStatus,omitempty" bson:"operationalstatus" mapstructure:"operationalStatus,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected" mapstructure:"protected,omitempty"`

	// Type of the container ecosystem.
	Type *ProcessingUnitTypeValue `json:"type,omitempty" bson:"type" mapstructure:"type,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseProcessingUnit returns a new  SparseProcessingUnit.
func NewSparseProcessingUnit() *SparseProcessingUnit {
	return &SparseProcessingUnit{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseProcessingUnit) Identity() elemental.Identity {

	return ProcessingUnitIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseProcessingUnit) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseProcessingUnit) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseProcessingUnit) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseProcessingUnit) ToPlain() elemental.PlainIdentifiable {

	out := NewProcessingUnit()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.Archived != nil {
		out.Archived = *o.Archived
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.EnforcementStatus != nil {
		out.EnforcementStatus = *o.EnforcementStatus
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.Image != nil {
		out.Image = *o.Image
	}
	if o.LastPokeTime != nil {
		out.LastPokeTime = *o.LastPokeTime
	}
	if o.LastSyncTime != nil {
		out.LastSyncTime = *o.LastSyncTime
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NativeContextID != nil {
		out.NativeContextID = *o.NativeContextID
	}
	if o.NetworkServices != nil {
		out.NetworkServices = *o.NetworkServices
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.OperationalStatus != nil {
		out.OperationalStatus = *o.OperationalStatus
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Type != nil {
		out.Type = *o.Type
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseProcessingUnit) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetArchived returns the Archived of the receiver.
func (o *SparseProcessingUnit) GetArchived() bool {

	return *o.Archived
}

// SetArchived sets the property Archived of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetArchived(archived bool) {

	o.Archived = &archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseProcessingUnit) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseProcessingUnit) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseProcessingUnit) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseProcessingUnit) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseProcessingUnit) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseProcessingUnit) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseProcessingUnit) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseProcessingUnit) GetProtected() bool {

	return *o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseProcessingUnit) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseProcessingUnit) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseProcessingUnit.
func (o *SparseProcessingUnit) DeepCopy() *SparseProcessingUnit {

	if o == nil {
		return nil
	}

	out := &SparseProcessingUnit{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseProcessingUnit.
func (o *SparseProcessingUnit) DeepCopyInto(out *SparseProcessingUnit) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseProcessingUnit: %s", err))
	}

	*out = *target.(*SparseProcessingUnit)
}
