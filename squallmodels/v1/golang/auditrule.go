package squallmodels

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/squallmodels/v1/golang/types"
	"time"
)

// AuditRuleArchitectureValue represents the possible values for attribute "architecture".
type AuditRuleArchitectureValue string

const (
	// AuditRuleArchitecture32bit represents the value 32bit.
	AuditRuleArchitecture32bit AuditRuleArchitectureValue = "32bit"

	// AuditRuleArchitecture64bit represents the value 64bit.
	AuditRuleArchitecture64bit AuditRuleArchitectureValue = "64bit"
)

// AuditRuleRuleTypeValue represents the possible values for attribute "ruleType".
type AuditRuleRuleTypeValue string

const (
	// AuditRuleRuleTypeFile represents the value File.
	AuditRuleRuleTypeFile AuditRuleRuleTypeValue = "File"

	// AuditRuleRuleTypeSyscall represents the value Syscall.
	AuditRuleRuleTypeSyscall AuditRuleRuleTypeValue = "Syscall"
)

// AuditRuleIdentity represents the Identity of the object.
var AuditRuleIdentity = elemental.Identity{
	Name:     "auditrule",
	Category: "auditrules",
}

// AuditRulesList represents a list of AuditRules
type AuditRulesList []*AuditRule

// ContentIdentity returns the identity of the objects in the list.
func (o AuditRulesList) ContentIdentity() elemental.Identity {

	return AuditRuleIdentity
}

// Copy returns a pointer to a copy the AuditRulesList.
func (o AuditRulesList) Copy() elemental.ContentIdentifiable {

	copy := append(AuditRulesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AuditRulesList.
func (o AuditRulesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(AuditRulesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AuditRule))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AuditRulesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AuditRulesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o AuditRulesList) Version() int {

	return 1
}

// AuditRule represents the model of a auditrule
type AuditRule struct {
	// Architecture is the processor architecture that this rule applies to. It can be 32-bit or 64-bit.
	Architecture AuditRuleArchitectureValue `json:"architecture" bson:"architecture"`

	// FilePath is the file path for a file system rule.
	FilePath string `json:"filePath" bson:"filepath"`

	// FilterRules is the list of filter rules that must be applied to the auditing rule.
	FilterRules []*types.AuditFilter `json:"filterRules" bson:"filterrules"`

	// GroupName is the name of the group that this rule must be associated with.
	GroupName string `json:"groupName" bson:"groupname"`

	// RuleType is the type of the audit rule and it can be SYSCALL or FILE.
	RuleType AuditRuleRuleTypeValue `json:"ruleType" bson:"ruletype"`

	// SysCalls is the list of system calls that the rule applies to. It is only valid if ruleType is SYSCALL.
	Syscalls []types.AuditSystemCallType `json:"syscalls" bson:"syscalls"`

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

// NewAuditRule returns a new *AuditRule
func NewAuditRule() *AuditRule {

	return &AuditRule{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		Architecture:   "64bit",
		AssociatedTags: []string{},
		FilterRules:    []*types.AuditFilter{},
		Metadata:       []string{},
		NormalizedTags: []string{},
		RuleType:       "Syscall",
		Syscalls:       []types.AuditSystemCallType{},
	}
}

// Identity returns the Identity of the object.
func (o *AuditRule) Identity() elemental.Identity {

	return AuditRuleIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AuditRule) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AuditRule) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *AuditRule) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *AuditRule) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *AuditRule) Doc() string {
	return `AuditRule describes an audit rule that must be applied to an enforer to detect anomalous events. `
}

func (o *AuditRule) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *AuditRule) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *AuditRule) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *AuditRule) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *AuditRule) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *AuditRule) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *AuditRule) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *AuditRule) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *AuditRule) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *AuditRule) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *AuditRule) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *AuditRule) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *AuditRule) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *AuditRule) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *AuditRule) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *AuditRule) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *AuditRule) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *AuditRule) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *AuditRule) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("architecture", string(o.Architecture), []string{"32bit", "64bit"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumLength("filePath", o.FilePath, 128, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("groupName", o.GroupName); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("groupName", o.GroupName, 31, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("groupName", o.GroupName); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("ruleType", string(o.RuleType), []string{"File", "Syscall"}, false); err != nil {
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
func (*AuditRule) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AuditRuleAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AuditRuleLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AuditRule) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AuditRuleAttributesMap
}

// AuditRuleAttributesMap represents the map of attribute for AuditRule.
var AuditRuleAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Architecture": elemental.AttributeSpecification{
		AllowedChoices: []string{"32bit", "64bit"},
		ConvertedName:  "Architecture",
		DefaultValue:   AuditRuleArchitecture64bit,
		Description:    `Architecture is the processor architecture that this rule applies to. It can be 32-bit or 64-bit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "architecture",
		Stored:         true,
		Type:           "enum",
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
	"FilePath": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FilePath",
		Description:    `FilePath is the file path for a file system rule.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		MaxLength:      128,
		Name:           "filePath",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"FilterRules": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FilterRules",
		Description:    `FilterRules is the list of filter rules that must be applied to the auditing rule.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "filterRules",
		Stored:         true,
		SubType:        "audit_filter_list",
		Type:           "external",
	},
	"GroupName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "GroupName",
		Description:    `GroupName is the name of the group that this rule must be associated with.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		MaxLength:      31,
		Name:           "groupName",
		Orderable:      true,
		Required:       true,
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
	"RuleType": elemental.AttributeSpecification{
		AllowedChoices: []string{"File", "Syscall"},
		ConvertedName:  "RuleType",
		DefaultValue:   AuditRuleRuleTypeSyscall,
		Description:    `RuleType is the type of the audit rule and it can be SYSCALL or FILE.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "ruleType",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"Syscalls": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Syscalls",
		Description:    `SysCalls is the list of system calls that the rule applies to. It is only valid if ruleType is SYSCALL.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "syscalls",
		Orderable:      true,
		Stored:         true,
		SubType:        "system_call_list",
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

// AuditRuleLowerCaseAttributesMap represents the map of attribute for AuditRule.
var AuditRuleLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"architecture": elemental.AttributeSpecification{
		AllowedChoices: []string{"32bit", "64bit"},
		ConvertedName:  "Architecture",
		DefaultValue:   AuditRuleArchitecture64bit,
		Description:    `Architecture is the processor architecture that this rule applies to. It can be 32-bit or 64-bit.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "architecture",
		Stored:         true,
		Type:           "enum",
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
	"filepath": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FilePath",
		Description:    `FilePath is the file path for a file system rule.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		MaxLength:      128,
		Name:           "filePath",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"filterrules": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FilterRules",
		Description:    `FilterRules is the list of filter rules that must be applied to the auditing rule.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "filterRules",
		Stored:         true,
		SubType:        "audit_filter_list",
		Type:           "external",
	},
	"groupname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "GroupName",
		Description:    `GroupName is the name of the group that this rule must be associated with.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		MaxLength:      31,
		Name:           "groupName",
		Orderable:      true,
		Required:       true,
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
	"ruletype": elemental.AttributeSpecification{
		AllowedChoices: []string{"File", "Syscall"},
		ConvertedName:  "RuleType",
		DefaultValue:   AuditRuleRuleTypeSyscall,
		Description:    `RuleType is the type of the audit rule and it can be SYSCALL or FILE.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "ruleType",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"syscalls": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Syscalls",
		Description:    `SysCalls is the list of system calls that the rule applies to. It is only valid if ruleType is SYSCALL.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "syscalls",
		Orderable:      true,
		Stored:         true,
		SubType:        "system_call_list",
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
