package squallmodels

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/squallmodels/v1/golang/types"
	"time"
)

// IsolationProfileDefaultSyscallActionValue represents the possible values for attribute "defaultSyscallAction".
type IsolationProfileDefaultSyscallActionValue string

const (
	// IsolationProfileDefaultSyscallActionAllow represents the value Allow.
	IsolationProfileDefaultSyscallActionAllow IsolationProfileDefaultSyscallActionValue = "Allow"

	// IsolationProfileDefaultSyscallActionErrno represents the value Errno.
	IsolationProfileDefaultSyscallActionErrno IsolationProfileDefaultSyscallActionValue = "Errno"

	// IsolationProfileDefaultSyscallActionKill represents the value Kill.
	IsolationProfileDefaultSyscallActionKill IsolationProfileDefaultSyscallActionValue = "Kill"

	// IsolationProfileDefaultSyscallActionTrace represents the value Trace.
	IsolationProfileDefaultSyscallActionTrace IsolationProfileDefaultSyscallActionValue = "Trace"

	// IsolationProfileDefaultSyscallActionTrap represents the value Trap.
	IsolationProfileDefaultSyscallActionTrap IsolationProfileDefaultSyscallActionValue = "Trap"
)

// IsolationProfileIdentity represents the Identity of the object.
var IsolationProfileIdentity = elemental.Identity{
	Name:     "isolationprofile",
	Category: "isolationprofiles",
}

// IsolationProfilesList represents a list of IsolationProfiles
type IsolationProfilesList []*IsolationProfile

// ContentIdentity returns the identity of the objects in the list.
func (o IsolationProfilesList) ContentIdentity() elemental.Identity {

	return IsolationProfileIdentity
}

// Copy returns a pointer to a copy the IsolationProfilesList.
func (o IsolationProfilesList) Copy() elemental.ContentIdentifiable {

	copy := append(IsolationProfilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the IsolationProfilesList.
func (o IsolationProfilesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(IsolationProfilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*IsolationProfile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o IsolationProfilesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o IsolationProfilesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o IsolationProfilesList) Version() int {

	return 1
}

// IsolationProfile represents the model of a isolationprofile
type IsolationProfile struct {
	// Priviledged allows the processing unit to be activated as privileged. Default false.
	AllowPrivileged bool `json:"allowPrivileged" bson:"allowprivileged"`

	// CapabilitiesAdded identifies additional capabilities that must be provided to the processing unit.
	CapabilitiesAdded []types.CapabilitiesType `json:"capabilitiesAdded" bson:"capabilitiesadded"`

	// CapabilitiesDropped identifies capabilities that must be dropped from the processing unit.
	CapabilitiesDropped []types.CapabilitiesType `json:"capabilitiesDropped" bson:"capabilitiesdropped"`

	// DefaultAction is the default action applied to all syscalls of this profile. Default is "Allow".
	DefaultSyscallAction IsolationProfileDefaultSyscallActionValue `json:"defaultSyscallAction" bson:"defaultsyscallaction"`

	// SyscallRules is a list of syscall rules that identify actions for particular syscalls.
	SyscallRules []*types.SyscallEnforcementRulesList `json:"syscallRules" bson:"syscallrules"`

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

// NewIsolationProfile returns a new *IsolationProfile
func NewIsolationProfile() *IsolationProfile {

	return &IsolationProfile{
		ModelVersion:         1,
		AllowPrivileged:      false,
		Annotations:          map[string][]string{},
		AssociatedTags:       []string{},
		CapabilitiesAdded:    []types.CapabilitiesType{},
		CapabilitiesDropped:  []types.CapabilitiesType{},
		DefaultSyscallAction: "Allow",
		Metadata:             []string{},
		NormalizedTags:       []string{},
		SyscallRules:         []*types.SyscallEnforcementRulesList{},
		TargetArchitectures:  []types.ArchitecturesType{},
	}
}

// Identity returns the Identity of the object.
func (o *IsolationProfile) Identity() elemental.Identity {

	return IsolationProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IsolationProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IsolationProfile) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *IsolationProfile) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *IsolationProfile) DefaultOrder() []string {

	return []string{
		"name",
	}
}

func (o *IsolationProfile) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *IsolationProfile) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *IsolationProfile) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *IsolationProfile) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *IsolationProfile) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *IsolationProfile) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *IsolationProfile) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *IsolationProfile) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *IsolationProfile) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *IsolationProfile) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *IsolationProfile) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *IsolationProfile) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *IsolationProfile) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *IsolationProfile) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *IsolationProfile) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *IsolationProfile) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *IsolationProfile) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *IsolationProfile) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *IsolationProfile) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("defaultSyscallAction", string(o.DefaultSyscallAction), []string{"Allow", "Errno", "Kill", "Trace", "Trap"}, false); err != nil {
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
func (*IsolationProfile) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := IsolationProfileAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return IsolationProfileLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*IsolationProfile) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return IsolationProfileAttributesMap
}

// IsolationProfileAttributesMap represents the map of attribute for IsolationProfile.
var IsolationProfileAttributesMap = map[string]elemental.AttributeSpecification{
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
	"CapabilitiesAdded": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CapabilitiesAdded",
		Description:    `CapabilitiesAdded identifies additional capabilities that must be provided to the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "capabilitiesAdded",
		Orderable:      true,
		Stored:         true,
		SubType:        "cap_list",
		Type:           "external",
	},
	"CapabilitiesDropped": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CapabilitiesDropped",
		Description:    `CapabilitiesDropped identifies capabilities that must be dropped from the processing unit. `,
		Exposed:        true,
		Filterable:     true,
		Name:           "capabilitiesDropped",
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
		DefaultValue:   IsolationProfileDefaultSyscallActionAllow,
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
	"SyscallRules": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SyscallRules",
		Description:    `SyscallRules is a list of syscall rules that identify actions for particular syscalls.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "syscallRules",
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

// IsolationProfileLowerCaseAttributesMap represents the map of attribute for IsolationProfile.
var IsolationProfileLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"capabilitiesadded": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CapabilitiesAdded",
		Description:    `CapabilitiesAdded identifies additional capabilities that must be provided to the processing unit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "capabilitiesAdded",
		Orderable:      true,
		Stored:         true,
		SubType:        "cap_list",
		Type:           "external",
	},
	"capabilitiesdropped": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CapabilitiesDropped",
		Description:    `CapabilitiesDropped identifies capabilities that must be dropped from the processing unit. `,
		Exposed:        true,
		Filterable:     true,
		Name:           "capabilitiesDropped",
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
		DefaultValue:   IsolationProfileDefaultSyscallActionAllow,
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
	"syscallrules": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SyscallRules",
		Description:    `SyscallRules is a list of syscall rules that identify actions for particular syscalls.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "syscallRules",
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
