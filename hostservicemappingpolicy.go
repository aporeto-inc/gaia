package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// HostServiceMappingPolicyIdentity represents the Identity of the object.
var HostServiceMappingPolicyIdentity = elemental.Identity{
	Name:     "hostservicemappingpolicy",
	Category: "hostservicemappingpolicies",
	Package:  "squall",
	Private:  false,
}

// HostServiceMappingPoliciesList represents a list of HostServiceMappingPolicies
type HostServiceMappingPoliciesList []*HostServiceMappingPolicy

// Identity returns the identity of the objects in the list.
func (o HostServiceMappingPoliciesList) Identity() elemental.Identity {

	return HostServiceMappingPolicyIdentity
}

// Copy returns a pointer to a copy the HostServiceMappingPoliciesList.
func (o HostServiceMappingPoliciesList) Copy() elemental.Identifiables {

	copy := append(HostServiceMappingPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the HostServiceMappingPoliciesList.
func (o HostServiceMappingPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(HostServiceMappingPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*HostServiceMappingPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o HostServiceMappingPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o HostServiceMappingPoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the HostServiceMappingPoliciesList converted to SparseHostServiceMappingPoliciesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o HostServiceMappingPoliciesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o HostServiceMappingPoliciesList) Version() int {

	return 1
}

// HostServiceMappingPolicy represents the model of a hostservicemappingpolicy
type HostServiceMappingPolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// ActiveDuration defines for how long the policy will be active according to the
	// activeSchedule.
	ActiveDuration string `json:"activeDuration" bson:"activeduration" mapstructure:"activeDuration,omitempty"`

	// ActiveSchedule defines when the policy should be active using the cron notation.
	// The policy will be active for the given activeDuration.
	ActiveSchedule string `json:"activeSchedule" bson:"activeschedule" mapstructure:"activeSchedule,omitempty"`

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

	// Fallback indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback bool `json:"fallback" bson:"fallback" mapstructure:"fallback,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Object of the policy is the selector for the host services that must be applied
	// to this enforcer.
	Object [][]string `json:"object" bson:"-" mapstructure:"object,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate bool `json:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Subject of the policy is the selector of the enforcers that the list of host
	// services must apply to.
	Subject [][]string `json:"subject" bson:"-" mapstructure:"subject,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewHostServiceMappingPolicy returns a new *HostServiceMappingPolicy
func NewHostServiceMappingPolicy() *HostServiceMappingPolicy {

	return &HostServiceMappingPolicy{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Metadata:       []string{},
		NormalizedTags: []string{},
		Object:         [][]string{},
		Subject:        [][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *HostServiceMappingPolicy) Identity() elemental.Identity {

	return HostServiceMappingPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *HostServiceMappingPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *HostServiceMappingPolicy) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *HostServiceMappingPolicy) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *HostServiceMappingPolicy) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *HostServiceMappingPolicy) Doc() string {
	return `Defines a host service mapping policy that provides the relation between
enforcers and host services that they must implement.`
}

func (o *HostServiceMappingPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *HostServiceMappingPolicy) GetActiveDuration() string {

	return o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *HostServiceMappingPolicy) GetActiveSchedule() string {

	return o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *HostServiceMappingPolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *HostServiceMappingPolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *HostServiceMappingPolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *HostServiceMappingPolicy) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *HostServiceMappingPolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetFallback returns the Fallback of the receiver.
func (o *HostServiceMappingPolicy) GetFallback() bool {

	return o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetFallback(fallback bool) {

	o.Fallback = fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *HostServiceMappingPolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *HostServiceMappingPolicy) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *HostServiceMappingPolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *HostServiceMappingPolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *HostServiceMappingPolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *HostServiceMappingPolicy) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *HostServiceMappingPolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *HostServiceMappingPolicy) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *HostServiceMappingPolicy) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *HostServiceMappingPolicy) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *HostServiceMappingPolicy) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseHostServiceMappingPolicy{
			ID:             &o.ID,
			ActiveDuration: &o.ActiveDuration,
			ActiveSchedule: &o.ActiveSchedule,
			Annotations:    &o.Annotations,
			AssociatedTags: &o.AssociatedTags,
			CreateTime:     &o.CreateTime,
			Description:    &o.Description,
			Disabled:       &o.Disabled,
			Fallback:       &o.Fallback,
			Metadata:       &o.Metadata,
			Name:           &o.Name,
			Namespace:      &o.Namespace,
			NormalizedTags: &o.NormalizedTags,
			Object:         &o.Object,
			Propagate:      &o.Propagate,
			Protected:      &o.Protected,
			Subject:        &o.Subject,
			UpdateTime:     &o.UpdateTime,
			ZHash:          &o.ZHash,
			Zone:           &o.Zone,
		}
	}

	sp := &SparseHostServiceMappingPolicy{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "activeDuration":
			sp.ActiveDuration = &(o.ActiveDuration)
		case "activeSchedule":
			sp.ActiveSchedule = &(o.ActiveSchedule)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "fallback":
			sp.Fallback = &(o.Fallback)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "object":
			sp.Object = &(o.Object)
		case "propagate":
			sp.Propagate = &(o.Propagate)
		case "protected":
			sp.Protected = &(o.Protected)
		case "subject":
			sp.Subject = &(o.Subject)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseHostServiceMappingPolicy to the object.
func (o *HostServiceMappingPolicy) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseHostServiceMappingPolicy)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.ActiveDuration != nil {
		o.ActiveDuration = *so.ActiveDuration
	}
	if so.ActiveSchedule != nil {
		o.ActiveSchedule = *so.ActiveSchedule
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
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
	if so.Disabled != nil {
		o.Disabled = *so.Disabled
	}
	if so.Fallback != nil {
		o.Fallback = *so.Fallback
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
	if so.Object != nil {
		o.Object = *so.Object
	}
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Subject != nil {
		o.Subject = *so.Subject
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the HostServiceMappingPolicy.
func (o *HostServiceMappingPolicy) DeepCopy() *HostServiceMappingPolicy {

	if o == nil {
		return nil
	}

	out := &HostServiceMappingPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *HostServiceMappingPolicy.
func (o *HostServiceMappingPolicy) DeepCopyInto(out *HostServiceMappingPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy HostServiceMappingPolicy: %s", err))
	}

	*out = *target.(*HostServiceMappingPolicy)
}

// Validate valides the current information stored into the structure.
func (o *HostServiceMappingPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidatePattern("activeDuration", o.ActiveDuration, `^[0-9]+[smh]$`, `must be a valid duration like <n>s or <n>s or <n>h`, false); err != nil {
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
func (*HostServiceMappingPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := HostServiceMappingPolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return HostServiceMappingPolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*HostServiceMappingPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return HostServiceMappingPolicyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *HostServiceMappingPolicy) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "activeDuration":
		return o.ActiveDuration
	case "activeSchedule":
		return o.ActiveSchedule
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "disabled":
		return o.Disabled
	case "fallback":
		return o.Fallback
	case "metadata":
		return o.Metadata
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "object":
		return o.Object
	case "propagate":
		return o.Propagate
	case "protected":
		return o.Protected
	case "subject":
		return o.Subject
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// HostServiceMappingPolicyAttributesMap represents the map of attribute for HostServiceMappingPolicy.
var HostServiceMappingPolicyAttributesMap = map[string]elemental.AttributeSpecification{
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
		Type:    "string",
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
		SubType:        "map[string][]string",
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
		SubType:        "string",
		Type:           "list",
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
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:   true,
		Getter:    true,
		Name:      "fallback",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
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
		SubType:    "string",
		Type:       "list",
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
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description: `Object of the policy is the selector for the host services that must be applied
to this enforcer.`,
		Exposed:   true,
		Name:      "object",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
	},
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description: `Subject of the policy is the selector of the enforcers that the list of host
services must apply to.`,
		Exposed:   true,
		Name:      "subject",
		Orderable: true,
		SubType:   "[][]string",
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
	"ZHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// HostServiceMappingPolicyLowerCaseAttributesMap represents the map of attribute for HostServiceMappingPolicy.
var HostServiceMappingPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		Type:    "string",
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
		SubType:        "map[string][]string",
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
		SubType:        "string",
		Type:           "list",
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
	"disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:   true,
		Getter:    true,
		Name:      "fallback",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
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
		SubType:    "string",
		Type:       "list",
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
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description: `Object of the policy is the selector for the host services that must be applied
to this enforcer.`,
		Exposed:   true,
		Name:      "object",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
	},
	"propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description: `Subject of the policy is the selector of the enforcers that the list of host
services must apply to.`,
		Exposed:   true,
		Name:      "subject",
		Orderable: true,
		SubType:   "[][]string",
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
	"zhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// SparseHostServiceMappingPoliciesList represents a list of SparseHostServiceMappingPolicies
type SparseHostServiceMappingPoliciesList []*SparseHostServiceMappingPolicy

// Identity returns the identity of the objects in the list.
func (o SparseHostServiceMappingPoliciesList) Identity() elemental.Identity {

	return HostServiceMappingPolicyIdentity
}

// Copy returns a pointer to a copy the SparseHostServiceMappingPoliciesList.
func (o SparseHostServiceMappingPoliciesList) Copy() elemental.Identifiables {

	copy := append(SparseHostServiceMappingPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseHostServiceMappingPoliciesList.
func (o SparseHostServiceMappingPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseHostServiceMappingPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseHostServiceMappingPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseHostServiceMappingPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseHostServiceMappingPoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseHostServiceMappingPoliciesList converted to HostServiceMappingPoliciesList.
func (o SparseHostServiceMappingPoliciesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseHostServiceMappingPoliciesList) Version() int {

	return 1
}

// SparseHostServiceMappingPolicy represents the sparse version of a hostservicemappingpolicy.
type SparseHostServiceMappingPolicy struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// ActiveDuration defines for how long the policy will be active according to the
	// activeSchedule.
	ActiveDuration *string `json:"activeDuration,omitempty" bson:"activeduration" mapstructure:"activeDuration,omitempty"`

	// ActiveSchedule defines when the policy should be active using the cron notation.
	// The policy will be active for the given activeDuration.
	ActiveSchedule *string `json:"activeSchedule,omitempty" bson:"activeschedule" mapstructure:"activeSchedule,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// Disabled defines if the propert is disabled.
	Disabled *bool `json:"disabled,omitempty" bson:"disabled" mapstructure:"disabled,omitempty"`

	// Fallback indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback *bool `json:"fallback,omitempty" bson:"fallback" mapstructure:"fallback,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Object of the policy is the selector for the host services that must be applied
	// to this enforcer.
	Object *[][]string `json:"object,omitempty" bson:"-" mapstructure:"object,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected" mapstructure:"protected,omitempty"`

	// Subject of the policy is the selector of the enforcers that the list of host
	// services must apply to.
	Subject *[][]string `json:"subject,omitempty" bson:"-" mapstructure:"subject,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-,omitempty" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-,omitempty" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseHostServiceMappingPolicy returns a new  SparseHostServiceMappingPolicy.
func NewSparseHostServiceMappingPolicy() *SparseHostServiceMappingPolicy {
	return &SparseHostServiceMappingPolicy{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseHostServiceMappingPolicy) Identity() elemental.Identity {

	return HostServiceMappingPolicyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseHostServiceMappingPolicy) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseHostServiceMappingPolicy) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseHostServiceMappingPolicy) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseHostServiceMappingPolicy) ToPlain() elemental.PlainIdentifiable {

	out := NewHostServiceMappingPolicy()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.ActiveDuration != nil {
		out.ActiveDuration = *o.ActiveDuration
	}
	if o.ActiveSchedule != nil {
		out.ActiveSchedule = *o.ActiveSchedule
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
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
	if o.Disabled != nil {
		out.Disabled = *o.Disabled
	}
	if o.Fallback != nil {
		out.Fallback = *o.Fallback
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
	if o.Object != nil {
		out.Object = *o.Object
	}
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Subject != nil {
		out.Subject = *o.Subject
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *SparseHostServiceMappingPolicy) GetActiveDuration() string {

	return *o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = &activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *SparseHostServiceMappingPolicy) GetActiveSchedule() string {

	return *o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = &activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseHostServiceMappingPolicy) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseHostServiceMappingPolicy) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseHostServiceMappingPolicy) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseHostServiceMappingPolicy) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseHostServiceMappingPolicy) GetDisabled() bool {

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetFallback returns the Fallback of the receiver.
func (o *SparseHostServiceMappingPolicy) GetFallback() bool {

	return *o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetFallback(fallback bool) {

	o.Fallback = &fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseHostServiceMappingPolicy) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseHostServiceMappingPolicy) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseHostServiceMappingPolicy) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseHostServiceMappingPolicy) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseHostServiceMappingPolicy) GetPropagate() bool {

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparseHostServiceMappingPolicy) GetProtected() bool {

	return *o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseHostServiceMappingPolicy) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseHostServiceMappingPolicy) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseHostServiceMappingPolicy) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseHostServiceMappingPolicy) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseHostServiceMappingPolicy.
func (o *SparseHostServiceMappingPolicy) DeepCopy() *SparseHostServiceMappingPolicy {

	if o == nil {
		return nil
	}

	out := &SparseHostServiceMappingPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseHostServiceMappingPolicy.
func (o *SparseHostServiceMappingPolicy) DeepCopyInto(out *SparseHostServiceMappingPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseHostServiceMappingPolicy: %s", err))
	}

	*out = *target.(*SparseHostServiceMappingPolicy)
}
