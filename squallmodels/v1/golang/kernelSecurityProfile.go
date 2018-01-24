package squallmodels

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/squallmodels/v1/golang/types"
	"time"
)

// KernelSecurityProfileDefaultSyscallActionValue represents the possible values for attribute "defaultSyscallAction".
type KernelSecurityProfileDefaultSyscallActionValue string

const (
	// KernelSecurityProfileDefaultSyscallActionAllow represents the value Allow.
	KernelSecurityProfileDefaultSyscallActionAllow KernelSecurityProfileDefaultSyscallActionValue = "Allow"

	// KernelSecurityProfileDefaultSyscallActionErrno represents the value Errno.
	KernelSecurityProfileDefaultSyscallActionErrno KernelSecurityProfileDefaultSyscallActionValue = "Errno"

	// KernelSecurityProfileDefaultSyscallActionKill represents the value Kill.
	KernelSecurityProfileDefaultSyscallActionKill KernelSecurityProfileDefaultSyscallActionValue = "Kill"

	// KernelSecurityProfileDefaultSyscallActionTrace represents the value Trace.
	KernelSecurityProfileDefaultSyscallActionTrace KernelSecurityProfileDefaultSyscallActionValue = "Trace"

	// KernelSecurityProfileDefaultSyscallActionTrap represents the value Trap.
	KernelSecurityProfileDefaultSyscallActionTrap KernelSecurityProfileDefaultSyscallActionValue = "Trap"
)

// KernelSecurityProfileIdentity represents the Identity of the object.
var KernelSecurityProfileIdentity = elemental.Identity{
	Name:     "kernelSecurityProfile",
	Category: "kernelSecurityProfiles",
}

// KernelSecurityProfilesList represents a list of KernelSecurityProfiles
type KernelSecurityProfilesList []*KernelSecurityProfile

// ContentIdentity returns the identity of the objects in the list.
func (o KernelSecurityProfilesList) ContentIdentity() elemental.Identity {

	return KernelSecurityProfileIdentity
}

// Copy returns a pointer to a copy the KernelSecurityProfilesList.
func (o KernelSecurityProfilesList) Copy() elemental.ContentIdentifiable {

	copy := append(KernelSecurityProfilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the KernelSecurityProfilesList.
func (o KernelSecurityProfilesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(KernelSecurityProfilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*KernelSecurityProfile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o KernelSecurityProfilesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o KernelSecurityProfilesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o KernelSecurityProfilesList) Version() int {

	return 1
}

// KernelSecurityProfile represents the model of a kernelSecurityProfile
type KernelSecurityProfile struct {
	// Priviledged allows the processing unit to be activated as privileged. Default false.
	AllowPrivileged bool `json:"allowPrivileged" bson:"allowprivileged"`

	// Capabilities identifies additional capabilities that must be provided to the container.
	Capabilities []types.CapabilitiesType `json:"capabilities" bson:"capabilities"`

	// DefaultAction is the default action applied to all syscalls of this profile. Default is "Allow".
	DefaultSyscallAction KernelSecurityProfileDefaultSyscallActionValue `json:"defaultSyscallAction" bson:"defaultsyscallaction"`

	// Key identifies the syscall profile. It is any string.
	Key string `json:"key" bson:"key"`

	// Syscalls is a list of syscall rules that identify actions for particular syscalls.
	Syscalls []*types.KernelSyscallRulesList `json:"syscalls" bson:"syscalls"`

	// TargetArchitectures is the target processor architectures where this profile can be applied. Default all.
	TargetArchitectures []types.ArchitecturesType `json:"targetArchitectures" bson:"targetarchitectures"`

	// Annotation stores additional information about an entity
	Annotations map[string][]string `json:"annotations" bson:"annotations"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// CreatedTime is the time at which the object was created
	CreateTime time.Time `json:"createTime" bson:"createtime"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id"`

	// Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewKernelSecurityProfile returns a new *KernelSecurityProfile
func NewKernelSecurityProfile() *KernelSecurityProfile {

	return &KernelSecurityProfile{
		ModelVersion:         1,
		AllowPrivileged:      false,
		Annotations:          map[string][]string{},
		AssociatedTags:       []string{},
		Capabilities:         []types.CapabilitiesType{},
		DefaultSyscallAction: "Allow",
		Metadata:             []string{},
		NormalizedTags:       []string{},
		Syscalls:             []*types.KernelSyscalRulesList{},
		TargetArchitectures:  []types.ArchitecturesType{},
	}
}

// Identity returns the Identity of the object.
func (o *KernelSecurityProfile) Identity() elemental.Identity {

	return KernelSecurityProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *KernelSecurityProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *KernelSecurityProfile) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *KernelSecurityProfile) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *KernelSecurityProfile) DefaultOrder() []string {

	return []string{
		"name",
	}
}

func (o *KernelSecurityProfile) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *KernelSecurityProfile) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *KernelSecurityProfile) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *KernelSecurityProfile) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *KernelSecurityProfile) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *KernelSecurityProfile) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *KernelSecurityProfile) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *KernelSecurityProfile) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *KernelSecurityProfile) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *KernelSecurityProfile) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *KernelSecurityProfile) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *KernelSecurityProfile) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *KernelSecurityProfile) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *KernelSecurityProfile) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *KernelSecurityProfile) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *KernelSecurityProfile) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *KernelSecurityProfile) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *KernelSecurityProfile) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *KernelSecurityProfile) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("defaultSyscallAction", string(o.DefaultSyscallAction), []string{"Allow", "Errno", "Kill", "Trace", "Trap"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumLength("key", o.Key, 64, false); err != nil {
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
func (*KernelSecurityProfile) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := KernelSecurityProfileAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return KernelSecurityProfileLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*KernelSecurityProfile) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return KernelSecurityProfileAttributesMap
}

// KernelSecurityProfileAttributesMap represents the map of attribute for KernelSecurityProfile.
var KernelSecurityProfileAttributesMap = map[string]elemental.AttributeSpecification{
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
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"AllowPrivileged": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowPrivileged",
		DefaultValue:   false,
		Description:    `Priviledged allows the processing unit to be activated as privileged. Default false.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowPrivileged",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity`,
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
		Description:    `AssociatedTags are the list of tags attached to an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"Capabilities": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Capabilities",
		Description:    `Capabilities identifies additional capabilities that must be provided to the container.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "capabilities",
		Orderable:      true,
		Stored:         true,
		SubType:        "cap_list",
		Type:           "external",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"DefaultSyscallAction": elemental.AttributeSpecification{
		AllowedChoices: []string{"Allow", "Errno", "Kill", "Trace", "Trap"},
		ConvertedName:  "DefaultSyscallAction",
		DefaultValue:   KernelSecurityProfileDefaultSyscallActionAllow,
		Description:    `DefaultAction is the default action applied to all syscalls of this profile. Default is "Allow".`,
		Exposed:        true,
		Filterable:     true,
		Name:           "defaultSyscallAction",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Key",
		Description:    `Key identifies the syscall profile. It is any string.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		MaxLength:      64,
		Name:           "key",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description:    `Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "metadata",
		Setter:         true,
		Stored:         true,
		SubType:        "metadata_list",
		Type:           "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
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
		ConvertedName:  "Namespace",
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
		ConvertedName:  "NormalizedTags",
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
	"Syscalls": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Syscalls",
		Description:    `Syscalls is a list of syscall rules that identify actions for particular syscalls.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "syscalls",
		Orderable:      true,
		Stored:         true,
		SubType:        "syscall_rules",
		Type:           "external",
	},
	"TargetArchitectures": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetArchitectures",
		Description:    `TargetArchitectures is the target processor architectures where this profile can be applied. Default all.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetArchitectures",
		Orderable:      true,
		Stored:         true,
		SubType:        "arch_list",
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

// KernelSecurityProfileLowerCaseAttributesMap represents the map of attribute for KernelSecurityProfile.
var KernelSecurityProfileLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"allowprivileged": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowPrivileged",
		DefaultValue:   false,
		Description:    `Priviledged allows the processing unit to be activated as privileged. Default false.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowPrivileged",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity`,
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
		Description:    `AssociatedTags are the list of tags attached to an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"capabilities": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Capabilities",
		Description:    `Capabilities identifies additional capabilities that must be provided to the container.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "capabilities",
		Orderable:      true,
		Stored:         true,
		SubType:        "cap_list",
		Type:           "external",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"defaultsyscallaction": elemental.AttributeSpecification{
		AllowedChoices: []string{"Allow", "Errno", "Kill", "Trace", "Trap"},
		ConvertedName:  "DefaultSyscallAction",
		DefaultValue:   KernelSecurityProfileDefaultSyscallActionAllow,
		Description:    `DefaultAction is the default action applied to all syscalls of this profile. Default is "Allow".`,
		Exposed:        true,
		Filterable:     true,
		Name:           "defaultSyscallAction",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Key",
		Description:    `Key identifies the syscall profile. It is any string.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		MaxLength:      64,
		Name:           "key",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description:    `Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "metadata",
		Setter:         true,
		Stored:         true,
		SubType:        "metadata_list",
		Type:           "external",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
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
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
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
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
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
	"syscalls": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Syscalls",
		Description:    `Syscalls is a list of syscall rules that identify actions for particular syscalls.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "syscalls",
		Orderable:      true,
		Stored:         true,
		SubType:        "syscall_rules",
		Type:           "external",
	},
	"targetarchitectures": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetArchitectures",
		Description:    `TargetArchitectures is the target processor architectures where this profile can be applied. Default all.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetArchitectures",
		Orderable:      true,
		Stored:         true,
		SubType:        "arch_list",
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
