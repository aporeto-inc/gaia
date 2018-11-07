package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/types"
)

// IsolationProfileIdentity represents the Identity of the object.
var IsolationProfileIdentity = elemental.Identity{
	Name:     "isolationprofile",
	Category: "isolationprofiles",
	Package:  "squall",
	Private:  false,
}

// IsolationProfilesList represents a list of IsolationProfiles
type IsolationProfilesList []*IsolationProfile

// Identity returns the identity of the objects in the list.
func (o IsolationProfilesList) Identity() elemental.Identity {

	return IsolationProfileIdentity
}

// Copy returns a pointer to a copy the IsolationProfilesList.
func (o IsolationProfilesList) Copy() elemental.Identifiables {

	copy := append(IsolationProfilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the IsolationProfilesList.
func (o IsolationProfilesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(IsolationProfilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*IsolationProfile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o IsolationProfilesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o IsolationProfilesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the IsolationProfilesList converted to SparseIsolationProfilesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o IsolationProfilesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o IsolationProfilesList) Version() int {

	return 1
}

// IsolationProfile represents the model of a isolationprofile
type IsolationProfile struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CapabilitiesActions identifies the capabilities that should be added or removed
	// from the processing unit.
	CapabilitiesActions types.CapabilitiesTypeMap `json:"capabilitiesActions" bson:"capabilitiesactions" mapstructure:"capabilitiesActions,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// DefaultAction is the default action applied to all syscalls of this profile.
	// Default is "Allow".
	DefaultSyscallAction types.SyscallEnforcementAction `json:"defaultSyscallAction" bson:"defaultsyscallaction" mapstructure:"defaultSyscallAction,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// SyscallRules is a list of syscall rules that identify actions for particular
	// syscalls.
	SyscallRules types.SyscallEnforcementRulesMap `json:"syscallRules" bson:"syscallrules" mapstructure:"syscallRules,omitempty"`

	// TargetArchitectures is the target processor architectures where this profile can
	// be applied. Default all.
	TargetArchitectures types.ArchitecturesTypeList `json:"targetArchitectures" bson:"targetarchitectures" mapstructure:"targetArchitectures,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewIsolationProfile returns a new *IsolationProfile
func NewIsolationProfile() *IsolationProfile {

	return &IsolationProfile{
		ModelVersion:        1,
		Annotations:         map[string][]string{},
		AssociatedTags:      []string{},
		CapabilitiesActions: types.CapabilitiesTypeMap{},
		Metadata:            []string{},
		NormalizedTags:      []string{},
		SyscallRules:        types.SyscallEnforcementRulesMap{},
		TargetArchitectures: types.ArchitecturesTypeList{},
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

// Doc returns the documentation for the object
func (o *IsolationProfile) Doc() string {
	return `An IsolationProfile needs documentation.`
}

func (o *IsolationProfile) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *IsolationProfile) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *IsolationProfile) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *IsolationProfile) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *IsolationProfile) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *IsolationProfile) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *IsolationProfile) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *IsolationProfile) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *IsolationProfile) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *IsolationProfile) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *IsolationProfile) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *IsolationProfile) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *IsolationProfile) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *IsolationProfile) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
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

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *IsolationProfile) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *IsolationProfile) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseIsolationProfile{
			ID:                   &o.ID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CapabilitiesActions:  &o.CapabilitiesActions,
			CreateTime:           &o.CreateTime,
			DefaultSyscallAction: &o.DefaultSyscallAction,
			Description:          &o.Description,
			Metadata:             &o.Metadata,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Protected:            &o.Protected,
			SyscallRules:         &o.SyscallRules,
			TargetArchitectures:  &o.TargetArchitectures,
			UpdateTime:           &o.UpdateTime,
		}
	}

	sp := &SparseIsolationProfile{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "capabilitiesActions":
			sp.CapabilitiesActions = &(o.CapabilitiesActions)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "defaultSyscallAction":
			sp.DefaultSyscallAction = &(o.DefaultSyscallAction)
		case "description":
			sp.Description = &(o.Description)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "syscallRules":
			sp.SyscallRules = &(o.SyscallRules)
		case "targetArchitectures":
			sp.TargetArchitectures = &(o.TargetArchitectures)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseIsolationProfile to the object.
func (o *IsolationProfile) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseIsolationProfile)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CapabilitiesActions != nil {
		o.CapabilitiesActions = *so.CapabilitiesActions
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.DefaultSyscallAction != nil {
		o.DefaultSyscallAction = *so.DefaultSyscallAction
	}
	if so.Description != nil {
		o.Description = *so.Description
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
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.SyscallRules != nil {
		o.SyscallRules = *so.SyscallRules
	}
	if so.TargetArchitectures != nil {
		o.TargetArchitectures = *so.TargetArchitectures
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the IsolationProfile.
func (o *IsolationProfile) DeepCopy() *IsolationProfile {

	if o == nil {
		return nil
	}

	out := &IsolationProfile{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *IsolationProfile.
func (o *IsolationProfile) DeepCopyInto(out *IsolationProfile) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy IsolationProfile: %s", err))
	}

	*out = *target.(*IsolationProfile)
}

// Validate valides the current information stored into the structure.
func (o *IsolationProfile) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *IsolationProfile) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "capabilitiesActions":
		return o.CapabilitiesActions
	case "createTime":
		return o.CreateTime
	case "defaultSyscallAction":
		return o.DefaultSyscallAction
	case "description":
		return o.Description
	case "metadata":
		return o.Metadata
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "syscallRules":
		return o.SyscallRules
	case "targetArchitectures":
		return o.TargetArchitectures
	case "updateTime":
		return o.UpdateTime
	}

	return nil
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
	"CapabilitiesActions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CapabilitiesActions",
		Description: `CapabilitiesActions identifies the capabilities that should be added or removed
from the processing unit.`,
		Exposed:   true,
		Name:      "capabilitiesActions",
		Orderable: true,
		Stored:    true,
		SubType:   "cap_map",
		Type:      "external",
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
	"DefaultSyscallAction": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DefaultSyscallAction",
		Description: `DefaultAction is the default action applied to all syscalls of this profile.
Default is "Allow".`,
		Exposed: true,
		Name:    "defaultSyscallAction",
		Stored:  true,
		SubType: "syscall_action",
		Type:    "external",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
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
	"SyscallRules": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SyscallRules",
		Description: `SyscallRules is a list of syscall rules that identify actions for particular
syscalls.`,
		Exposed:   true,
		Name:      "syscallRules",
		Orderable: true,
		Stored:    true,
		SubType:   "syscall_rules",
		Type:      "external",
	},
	"TargetArchitectures": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetArchitectures",
		Description: `TargetArchitectures is the target processor architectures where this profile can
be applied. Default all.`,
		Exposed:   true,
		Name:      "targetArchitectures",
		Orderable: true,
		Stored:    true,
		SubType:   "arch_list",
		Type:      "external",
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
	"capabilitiesactions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CapabilitiesActions",
		Description: `CapabilitiesActions identifies the capabilities that should be added or removed
from the processing unit.`,
		Exposed:   true,
		Name:      "capabilitiesActions",
		Orderable: true,
		Stored:    true,
		SubType:   "cap_map",
		Type:      "external",
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
	"defaultsyscallaction": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DefaultSyscallAction",
		Description: `DefaultAction is the default action applied to all syscalls of this profile.
Default is "Allow".`,
		Exposed: true,
		Name:    "defaultSyscallAction",
		Stored:  true,
		SubType: "syscall_action",
		Type:    "external",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
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
	"syscallrules": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SyscallRules",
		Description: `SyscallRules is a list of syscall rules that identify actions for particular
syscalls.`,
		Exposed:   true,
		Name:      "syscallRules",
		Orderable: true,
		Stored:    true,
		SubType:   "syscall_rules",
		Type:      "external",
	},
	"targetarchitectures": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetArchitectures",
		Description: `TargetArchitectures is the target processor architectures where this profile can
be applied. Default all.`,
		Exposed:   true,
		Name:      "targetArchitectures",
		Orderable: true,
		Stored:    true,
		SubType:   "arch_list",
		Type:      "external",
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

// SparseIsolationProfilesList represents a list of SparseIsolationProfiles
type SparseIsolationProfilesList []*SparseIsolationProfile

// Identity returns the identity of the objects in the list.
func (o SparseIsolationProfilesList) Identity() elemental.Identity {

	return IsolationProfileIdentity
}

// Copy returns a pointer to a copy the SparseIsolationProfilesList.
func (o SparseIsolationProfilesList) Copy() elemental.Identifiables {

	copy := append(SparseIsolationProfilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseIsolationProfilesList.
func (o SparseIsolationProfilesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseIsolationProfilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseIsolationProfile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseIsolationProfilesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseIsolationProfilesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseIsolationProfilesList converted to IsolationProfilesList.
func (o SparseIsolationProfilesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseIsolationProfilesList) Version() int {

	return 1
}

// SparseIsolationProfile represents the sparse version of a isolationprofile.
type SparseIsolationProfile struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CapabilitiesActions identifies the capabilities that should be added or removed
	// from the processing unit.
	CapabilitiesActions *types.CapabilitiesTypeMap `json:"capabilitiesActions,omitempty" bson:"capabilitiesactions" mapstructure:"capabilitiesActions,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// DefaultAction is the default action applied to all syscalls of this profile.
	// Default is "Allow".
	DefaultSyscallAction *types.SyscallEnforcementAction `json:"defaultSyscallAction,omitempty" bson:"defaultsyscallaction" mapstructure:"defaultSyscallAction,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected" mapstructure:"protected,omitempty"`

	// SyscallRules is a list of syscall rules that identify actions for particular
	// syscalls.
	SyscallRules *types.SyscallEnforcementRulesMap `json:"syscallRules,omitempty" bson:"syscallrules" mapstructure:"syscallRules,omitempty"`

	// TargetArchitectures is the target processor architectures where this profile can
	// be applied. Default all.
	TargetArchitectures *types.ArchitecturesTypeList `json:"targetArchitectures,omitempty" bson:"targetarchitectures" mapstructure:"targetArchitectures,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseIsolationProfile returns a new  SparseIsolationProfile.
func NewSparseIsolationProfile() *SparseIsolationProfile {
	return &SparseIsolationProfile{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseIsolationProfile) Identity() elemental.Identity {

	return IsolationProfileIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseIsolationProfile) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseIsolationProfile) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseIsolationProfile) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseIsolationProfile) ToPlain() elemental.PlainIdentifiable {

	out := NewIsolationProfile()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CapabilitiesActions != nil {
		out.CapabilitiesActions = *o.CapabilitiesActions
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.DefaultSyscallAction != nil {
		out.DefaultSyscallAction = *o.DefaultSyscallAction
	}
	if o.Description != nil {
		out.Description = *o.Description
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
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.SyscallRules != nil {
		out.SyscallRules = *o.SyscallRules
	}
	if o.TargetArchitectures != nil {
		out.TargetArchitectures = *o.TargetArchitectures
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseIsolationProfile) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseIsolationProfile) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseIsolationProfile) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseIsolationProfile) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseIsolationProfile) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseIsolationProfile) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseIsolationProfile) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseIsolationProfile) GetProtected() bool {

	return *o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseIsolationProfile) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseIsolationProfile) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseIsolationProfile.
func (o *SparseIsolationProfile) DeepCopy() *SparseIsolationProfile {

	if o == nil {
		return nil
	}

	out := &SparseIsolationProfile{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseIsolationProfile.
func (o *SparseIsolationProfile) DeepCopyInto(out *SparseIsolationProfile) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseIsolationProfile: %s", err))
	}

	*out = *target.(*SparseIsolationProfile)
}
