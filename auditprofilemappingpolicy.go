package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AuditProfileMappingPolicyIdentity represents the Identity of the object.
var AuditProfileMappingPolicyIdentity = elemental.Identity{
	Name:     "auditprofilemappingpolicy",
	Category: "auditprofilemappingpolicies",
	Package:  "squall",
	Private:  false,
}

// AuditProfileMappingPoliciesList represents a list of AuditProfileMappingPolicies
type AuditProfileMappingPoliciesList []*AuditProfileMappingPolicy

// Identity returns the identity of the objects in the list.
func (o AuditProfileMappingPoliciesList) Identity() elemental.Identity {

	return AuditProfileMappingPolicyIdentity
}

// Copy returns a pointer to a copy the AuditProfileMappingPoliciesList.
func (o AuditProfileMappingPoliciesList) Copy() elemental.Identifiables {

	copy := append(AuditProfileMappingPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AuditProfileMappingPoliciesList.
func (o AuditProfileMappingPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AuditProfileMappingPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AuditProfileMappingPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AuditProfileMappingPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AuditProfileMappingPoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the AuditProfileMappingPoliciesList converted to SparseAuditProfileMappingPoliciesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AuditProfileMappingPoliciesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o AuditProfileMappingPoliciesList) Version() int {

	return 1
}

// AuditProfileMappingPolicy represents the model of a auditprofilemappingpolicy
type AuditProfileMappingPolicy struct {
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

	// Object of the policy is the selector of the audit profiles that must be applied
	// based on this policy.
	Object [][]string `json:"object" bson:"-" mapstructure:"object,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate bool `json:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Subject of the policy is a selector of the enforcers that must implement the
	// policy.
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

// NewAuditProfileMappingPolicy returns a new *AuditProfileMappingPolicy
func NewAuditProfileMappingPolicy() *AuditProfileMappingPolicy {

	return &AuditProfileMappingPolicy{
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
func (o *AuditProfileMappingPolicy) Identity() elemental.Identity {

	return AuditProfileMappingPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AuditProfileMappingPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AuditProfileMappingPolicy) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *AuditProfileMappingPolicy) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *AuditProfileMappingPolicy) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *AuditProfileMappingPolicy) Doc() string {
	return `Defines an audit policy that determine the sets of enforcers that must implement
a specific audit profile.`
}

func (o *AuditProfileMappingPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *AuditProfileMappingPolicy) GetActiveDuration() string {

	return o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *AuditProfileMappingPolicy) GetActiveSchedule() string {

	return o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *AuditProfileMappingPolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *AuditProfileMappingPolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *AuditProfileMappingPolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *AuditProfileMappingPolicy) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *AuditProfileMappingPolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetFallback returns the Fallback of the receiver.
func (o *AuditProfileMappingPolicy) GetFallback() bool {

	return o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetFallback(fallback bool) {

	o.Fallback = fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *AuditProfileMappingPolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *AuditProfileMappingPolicy) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *AuditProfileMappingPolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *AuditProfileMappingPolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *AuditProfileMappingPolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *AuditProfileMappingPolicy) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *AuditProfileMappingPolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *AuditProfileMappingPolicy) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *AuditProfileMappingPolicy) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *AuditProfileMappingPolicy) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AuditProfileMappingPolicy) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAuditProfileMappingPolicy{
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

	sp := &SparseAuditProfileMappingPolicy{}
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

// Patch apply the non nil value of a *SparseAuditProfileMappingPolicy to the object.
func (o *AuditProfileMappingPolicy) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAuditProfileMappingPolicy)
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

// DeepCopy returns a deep copy if the AuditProfileMappingPolicy.
func (o *AuditProfileMappingPolicy) DeepCopy() *AuditProfileMappingPolicy {

	if o == nil {
		return nil
	}

	out := &AuditProfileMappingPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AuditProfileMappingPolicy.
func (o *AuditProfileMappingPolicy) DeepCopyInto(out *AuditProfileMappingPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AuditProfileMappingPolicy: %s", err))
	}

	*out = *target.(*AuditProfileMappingPolicy)
}

// Validate valides the current information stored into the structure.
func (o *AuditProfileMappingPolicy) Validate() error {

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
func (*AuditProfileMappingPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AuditProfileMappingPolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AuditProfileMappingPolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AuditProfileMappingPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AuditProfileMappingPolicyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AuditProfileMappingPolicy) ValueForAttribute(name string) interface{} {

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

// AuditProfileMappingPolicyAttributesMap represents the map of attribute for AuditProfileMappingPolicy.
var AuditProfileMappingPolicyAttributesMap = map[string]elemental.AttributeSpecification{
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
		Description: `Object of the policy is the selector of the audit profiles that must be applied
based on this policy.`,
		Exposed: true,
		Name:    "object",
		SubType: "[][]string",
		Type:    "external",
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
		Description: `Subject of the policy is a selector of the enforcers that must implement the
policy.`,
		Exposed: true,
		Name:    "subject",
		SubType: "[][]string",
		Type:    "external",
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

// AuditProfileMappingPolicyLowerCaseAttributesMap represents the map of attribute for AuditProfileMappingPolicy.
var AuditProfileMappingPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		Description: `Object of the policy is the selector of the audit profiles that must be applied
based on this policy.`,
		Exposed: true,
		Name:    "object",
		SubType: "[][]string",
		Type:    "external",
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
		Description: `Subject of the policy is a selector of the enforcers that must implement the
policy.`,
		Exposed: true,
		Name:    "subject",
		SubType: "[][]string",
		Type:    "external",
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

// SparseAuditProfileMappingPoliciesList represents a list of SparseAuditProfileMappingPolicies
type SparseAuditProfileMappingPoliciesList []*SparseAuditProfileMappingPolicy

// Identity returns the identity of the objects in the list.
func (o SparseAuditProfileMappingPoliciesList) Identity() elemental.Identity {

	return AuditProfileMappingPolicyIdentity
}

// Copy returns a pointer to a copy the SparseAuditProfileMappingPoliciesList.
func (o SparseAuditProfileMappingPoliciesList) Copy() elemental.Identifiables {

	copy := append(SparseAuditProfileMappingPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAuditProfileMappingPoliciesList.
func (o SparseAuditProfileMappingPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAuditProfileMappingPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAuditProfileMappingPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAuditProfileMappingPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAuditProfileMappingPoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseAuditProfileMappingPoliciesList converted to AuditProfileMappingPoliciesList.
func (o SparseAuditProfileMappingPoliciesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAuditProfileMappingPoliciesList) Version() int {

	return 1
}

// SparseAuditProfileMappingPolicy represents the sparse version of a auditprofilemappingpolicy.
type SparseAuditProfileMappingPolicy struct {
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

	// Object of the policy is the selector of the audit profiles that must be applied
	// based on this policy.
	Object *[][]string `json:"object,omitempty" bson:"-" mapstructure:"object,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected" mapstructure:"protected,omitempty"`

	// Subject of the policy is a selector of the enforcers that must implement the
	// policy.
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

// NewSparseAuditProfileMappingPolicy returns a new  SparseAuditProfileMappingPolicy.
func NewSparseAuditProfileMappingPolicy() *SparseAuditProfileMappingPolicy {
	return &SparseAuditProfileMappingPolicy{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAuditProfileMappingPolicy) Identity() elemental.Identity {

	return AuditProfileMappingPolicyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAuditProfileMappingPolicy) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAuditProfileMappingPolicy) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseAuditProfileMappingPolicy) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAuditProfileMappingPolicy) ToPlain() elemental.PlainIdentifiable {

	out := NewAuditProfileMappingPolicy()
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
func (o *SparseAuditProfileMappingPolicy) GetActiveDuration() string {

	return *o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = &activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetActiveSchedule() string {

	return *o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = &activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetDisabled() bool {

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetFallback returns the Fallback of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetFallback() bool {

	return *o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetFallback(fallback bool) {

	o.Fallback = &fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetPropagate() bool {

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetProtected() bool {

	return *o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAuditProfileMappingPolicy) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAuditProfileMappingPolicy) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAuditProfileMappingPolicy.
func (o *SparseAuditProfileMappingPolicy) DeepCopy() *SparseAuditProfileMappingPolicy {

	if o == nil {
		return nil
	}

	out := &SparseAuditProfileMappingPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAuditProfileMappingPolicy.
func (o *SparseAuditProfileMappingPolicy) DeepCopyInto(out *SparseAuditProfileMappingPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAuditProfileMappingPolicy: %s", err))
	}

	*out = *target.(*SparseAuditProfileMappingPolicy)
}
