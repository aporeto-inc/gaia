package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
)

// FileAccessPolicyIdentity represents the Identity of the object.
var FileAccessPolicyIdentity = elemental.Identity{
	Name:     "fileaccesspolicy",
	Category: "fileaccesspolicies",
	Private:  false,
}

// FileAccessPoliciesList represents a list of FileAccessPolicies
type FileAccessPoliciesList []*FileAccessPolicy

// ContentIdentity returns the identity of the objects in the list.
func (o FileAccessPoliciesList) ContentIdentity() elemental.Identity {

	return FileAccessPolicyIdentity
}

// Copy returns a pointer to a copy the FileAccessPoliciesList.
func (o FileAccessPoliciesList) Copy() elemental.ContentIdentifiable {

	copy := append(FileAccessPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the FileAccessPoliciesList.
func (o FileAccessPoliciesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(FileAccessPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FileAccessPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FileAccessPoliciesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FileAccessPoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o FileAccessPoliciesList) Version() int {

	return 1
}

// FileAccessPolicy represents the model of a fileaccesspolicy
type FileAccessPolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// ActiveDuration defines for how long the policy will be active according to the
	// activeSchedule.
	ActiveDuration string `json:"activeDuration" bson:"activeduration" mapstructure:"activeDuration,omitempty"`

	// ActiveSchedule defines when the policy should be active using the cron notation.
	// The policy will be active for the given activeDuration.
	ActiveSchedule string `json:"activeSchedule" bson:"activeschedule" mapstructure:"activeSchedule,omitempty"`

	// AllowsExecute allows to execute the files.
	AllowsExecute bool `json:"allowsExecute" bson:"-" mapstructure:"allowsExecute,omitempty"`

	// AllowsRead allows to read the files.
	AllowsRead bool `json:"allowsRead" bson:"-" mapstructure:"allowsRead,omitempty"`

	// AllowsWrite allows to write the files.
	AllowsWrite bool `json:"allowsWrite" bson:"-" mapstructure:"allowsWrite,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Disabled defines if the propert is disabled.
	Disabled bool `json:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// EncryptionEnabled will enable the automatic encryption.
	EncryptionEnabled bool `json:"encryptionEnabled" bson:"-" mapstructure:"encryptionEnabled,omitempty"`

	// Fallback indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback bool `json:"fallback" bson:"fallback" mapstructure:"fallback,omitempty"`

	// LogsEnabled will enable logging when this policy is used.
	LogsEnabled bool `json:"logsEnabled" bson:"-" mapstructure:"logsEnabled,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Object is the object of the policy.
	Object [][]string `json:"object" bson:"-" mapstructure:"object,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate bool `json:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// If set to true while the policy is propagating, it won't be visible to children
	// namespace, but still used for policy resolution.
	PropagationHidden bool `json:"propagationHidden" bson:"propagationhidden" mapstructure:"propagationHidden,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Subject is the subject of the policy.
	Subject [][]string `json:"subject" bson:"-" mapstructure:"subject,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewFileAccessPolicy returns a new *FileAccessPolicy
func NewFileAccessPolicy() *FileAccessPolicy {

	return &FileAccessPolicy{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Metadata:       []string{},
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *FileAccessPolicy) Identity() elemental.Identity {

	return FileAccessPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FileAccessPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FileAccessPolicy) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *FileAccessPolicy) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *FileAccessPolicy) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *FileAccessPolicy) Doc() string {
	return `A File Access Policy allows Processing Units to access various folder and files.
It will use the tags of a File Path to know what is the path of the file or
folder to allow access to. You can allow the Processing Unit to have any
combination of read, write or execute.

When a Processing Unit is Docker container, then it will police the volumes
mount. executewon''t have any effect.

File path are not supported yet for standard Linux processes.`
}

func (o *FileAccessPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *FileAccessPolicy) GetActiveDuration() string {

	return o.ActiveDuration
}

// SetActiveDuration sets the given ActiveDuration of the receiver.
func (o *FileAccessPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *FileAccessPolicy) GetActiveSchedule() string {

	return o.ActiveSchedule
}

// SetActiveSchedule sets the given ActiveSchedule of the receiver.
func (o *FileAccessPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *FileAccessPolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *FileAccessPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *FileAccessPolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *FileAccessPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *FileAccessPolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *FileAccessPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDisabled returns the Disabled of the receiver.
func (o *FileAccessPolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the given Disabled of the receiver.
func (o *FileAccessPolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetFallback returns the Fallback of the receiver.
func (o *FileAccessPolicy) GetFallback() bool {

	return o.Fallback
}

// SetFallback sets the given Fallback of the receiver.
func (o *FileAccessPolicy) SetFallback(fallback bool) {

	o.Fallback = fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *FileAccessPolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *FileAccessPolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *FileAccessPolicy) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *FileAccessPolicy) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *FileAccessPolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *FileAccessPolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *FileAccessPolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *FileAccessPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *FileAccessPolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the given Propagate of the receiver.
func (o *FileAccessPolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetPropagationHidden returns the PropagationHidden of the receiver.
func (o *FileAccessPolicy) GetPropagationHidden() bool {

	return o.PropagationHidden
}

// SetPropagationHidden sets the given PropagationHidden of the receiver.
func (o *FileAccessPolicy) SetPropagationHidden(propagationHidden bool) {

	o.PropagationHidden = propagationHidden
}

// GetProtected returns the Protected of the receiver.
func (o *FileAccessPolicy) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *FileAccessPolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *FileAccessPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *FileAccessPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidatePattern("activeDuration", o.ActiveDuration, `^[0-9]+[smh]$`, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
func (*FileAccessPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FileAccessPolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FileAccessPolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FileAccessPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FileAccessPolicyAttributesMap
}

// FileAccessPolicyAttributesMap represents the map of attribute for FileAccessPolicy.
var FileAccessPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"ActiveDuration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "ActiveDuration",
		Description: `ActiveDuration defines for how long the policy will be active according to the
activeSchedule.`,
		Exposed: true,
		Format:  "free",
		Getter:  true,
		Name:    "activeDuration",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"ActiveSchedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ActiveSchedule",
		Description: `ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		SubType: "cron_expression",
		Type:    "external",
	},
	"AllowsExecute": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsExecute",
		Description:    `AllowsExecute allows to execute the files.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsExecute",
		Orderable:      true,
		Type:           "boolean",
	},
	"AllowsRead": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsRead",
		Description:    `AllowsRead allows to read the files.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsRead",
		Orderable:      true,
		Type:           "boolean",
	},
	"AllowsWrite": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsWrite",
		Description:    `AllowsWrite allows to write the files.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsWrite",
		Orderable:      true,
		Type:           "boolean",
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
		Format:         "free",
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"EncryptionEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EncryptionEnabled",
		Description:    `EncryptionEnabled will enable the automatic encryption.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "encryptionEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"Fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "fallback",
		Orderable:  true,
		Setter:     true,
		Stored:     true,
		Type:       "boolean",
	},
	"LogsEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LogsEnabled",
		Description:    `LogsEnabled will enable logging when this policy is used.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "logsEnabled",
		Orderable:      true,
		Type:           "boolean",
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
		Format:         "free",
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
	"Object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description:    `Object is the object of the policy.`,
		Exposed:        true,
		Name:           "object",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"PropagationHidden": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PropagationHidden",
		Description: `If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "propagationHidden",
		Orderable:  true,
		Setter:     true,
		Stored:     true,
		Type:       "boolean",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `Subject is the subject of the policy.`,
		Exposed:        true,
		Name:           "subject",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
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

// FileAccessPolicyLowerCaseAttributesMap represents the map of attribute for FileAccessPolicy.
var FileAccessPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"activeduration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "ActiveDuration",
		Description: `ActiveDuration defines for how long the policy will be active according to the
activeSchedule.`,
		Exposed: true,
		Format:  "free",
		Getter:  true,
		Name:    "activeDuration",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"activeschedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ActiveSchedule",
		Description: `ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		SubType: "cron_expression",
		Type:    "external",
	},
	"allowsexecute": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsExecute",
		Description:    `AllowsExecute allows to execute the files.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsExecute",
		Orderable:      true,
		Type:           "boolean",
	},
	"allowsread": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsRead",
		Description:    `AllowsRead allows to read the files.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsRead",
		Orderable:      true,
		Type:           "boolean",
	},
	"allowswrite": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsWrite",
		Description:    `AllowsWrite allows to write the files.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsWrite",
		Orderable:      true,
		Type:           "boolean",
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
		Format:         "free",
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"encryptionenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EncryptionEnabled",
		Description:    `EncryptionEnabled will enable the automatic encryption.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "encryptionEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "fallback",
		Orderable:  true,
		Setter:     true,
		Stored:     true,
		Type:       "boolean",
	},
	"logsenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LogsEnabled",
		Description:    `LogsEnabled will enable logging when this policy is used.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "logsEnabled",
		Orderable:      true,
		Type:           "boolean",
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
		Format:         "free",
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
	"object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description:    `Object is the object of the policy.`,
		Exposed:        true,
		Name:           "object",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"propagationhidden": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PropagationHidden",
		Description: `If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "propagationHidden",
		Orderable:  true,
		Setter:     true,
		Stored:     true,
		Type:       "boolean",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `Subject is the subject of the policy.`,
		Exposed:        true,
		Name:           "subject",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
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
