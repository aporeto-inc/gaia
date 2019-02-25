package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EnforcerProfileIdentity represents the Identity of the object.
var EnforcerProfileIdentity = elemental.Identity{
	Name:     "enforcerprofile",
	Category: "enforcerprofiles",
	Package:  "squall",
	Private:  false,
}

// EnforcerProfilesList represents a list of EnforcerProfiles
type EnforcerProfilesList []*EnforcerProfile

// Identity returns the identity of the objects in the list.
func (o EnforcerProfilesList) Identity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Copy returns a pointer to a copy the EnforcerProfilesList.
func (o EnforcerProfilesList) Copy() elemental.Identifiables {

	copy := append(EnforcerProfilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcerProfilesList.
func (o EnforcerProfilesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EnforcerProfilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EnforcerProfile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerProfilesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerProfilesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the EnforcerProfilesList converted to SparseEnforcerProfilesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EnforcerProfilesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o EnforcerProfilesList) Version() int {

	return 1
}

// EnforcerProfile represents the model of a enforcerprofile
type EnforcerProfile struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// ExcludedInterfaces is a list of interfaces that must be excluded.
	ExcludedInterfaces []string `json:"excludedInterfaces" bson:"excludedinterfaces" mapstructure:"excludedInterfaces,omitempty"`

	// ExcludedNetworks is the list of networks that must be excluded for this
	// enforcer.
	ExcludedNetworks []string `json:"excludedNetworks" bson:"excludednetworks" mapstructure:"excludedNetworks,omitempty"`

	// IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
	// docker container started with labels matching the rule.
	IgnoreExpression [][]string `json:"ignoreExpression" bson:"ignoreexpression" mapstructure:"ignoreExpression,omitempty"`

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

	// TargetNetworks is the list of networks that authorization should be applied.
	TargetNetworks []string `json:"targetNetworks" bson:"targetnetworks" mapstructure:"targetNetworks,omitempty"`

	// TargetUDPNetworks is the list of UDP networks that authorization should be
	// applied.
	TargetUDPNetworks []string `json:"targetUDPNetworks" bson:"targetudpnetworks" mapstructure:"targetUDPNetworks,omitempty"`

	// List of trusted CA. If empty the main chain of trust will be used.
	TrustedCAs []string `json:"trustedCAs" bson:"trustedcas" mapstructure:"trustedCAs,omitempty"`

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

// NewEnforcerProfile returns a new *EnforcerProfile
func NewEnforcerProfile() *EnforcerProfile {

	return &EnforcerProfile{
		ModelVersion:       1,
		Annotations:        map[string][]string{},
		ExcludedNetworks:   []string{},
		AssociatedTags:     []string{},
		ExcludedInterfaces: []string{},
		IgnoreExpression:   [][]string{},
		Metadata:           []string{},
		NormalizedTags:     []string{},
		TargetNetworks:     []string{},
		TargetUDPNetworks:  []string{},
		TrustedCAs:         []string{},
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerProfile) Identity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerProfile) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *EnforcerProfile) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerProfile) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *EnforcerProfile) Doc() string {
	return `Allows to create reusable configuration profile for your enforcers. Enforcer
Profiles contains various startup information that can (for some) be updated
live. Enforcer Profiles are assigned to some Enforcer using a Enforcer Profile
Mapping Policy.`
}

func (o *EnforcerProfile) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *EnforcerProfile) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *EnforcerProfile) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *EnforcerProfile) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *EnforcerProfile) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *EnforcerProfile) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *EnforcerProfile) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *EnforcerProfile) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *EnforcerProfile) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *EnforcerProfile) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *EnforcerProfile) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *EnforcerProfile) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *EnforcerProfile) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *EnforcerProfile) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *EnforcerProfile) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *EnforcerProfile) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *EnforcerProfile) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *EnforcerProfile) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *EnforcerProfile) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *EnforcerProfile) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *EnforcerProfile) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *EnforcerProfile) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *EnforcerProfile) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *EnforcerProfile) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *EnforcerProfile) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEnforcerProfile{
			ID:                 &o.ID,
			Annotations:        &o.Annotations,
			AssociatedTags:     &o.AssociatedTags,
			CreateTime:         &o.CreateTime,
			Description:        &o.Description,
			ExcludedInterfaces: &o.ExcludedInterfaces,
			ExcludedNetworks:   &o.ExcludedNetworks,
			IgnoreExpression:   &o.IgnoreExpression,
			Metadata:           &o.Metadata,
			Name:               &o.Name,
			Namespace:          &o.Namespace,
			NormalizedTags:     &o.NormalizedTags,
			Protected:          &o.Protected,
			TargetNetworks:     &o.TargetNetworks,
			TargetUDPNetworks:  &o.TargetUDPNetworks,
			TrustedCAs:         &o.TrustedCAs,
			UpdateTime:         &o.UpdateTime,
			ZHash:              &o.ZHash,
			Zone:               &o.Zone,
		}
	}

	sp := &SparseEnforcerProfile{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "excludedInterfaces":
			sp.ExcludedInterfaces = &(o.ExcludedInterfaces)
		case "excludedNetworks":
			sp.ExcludedNetworks = &(o.ExcludedNetworks)
		case "ignoreExpression":
			sp.IgnoreExpression = &(o.IgnoreExpression)
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
		case "targetNetworks":
			sp.TargetNetworks = &(o.TargetNetworks)
		case "targetUDPNetworks":
			sp.TargetUDPNetworks = &(o.TargetUDPNetworks)
		case "trustedCAs":
			sp.TrustedCAs = &(o.TrustedCAs)
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

// Patch apply the non nil value of a *SparseEnforcerProfile to the object.
func (o *EnforcerProfile) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEnforcerProfile)
	if so.ID != nil {
		o.ID = *so.ID
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
	if so.ExcludedInterfaces != nil {
		o.ExcludedInterfaces = *so.ExcludedInterfaces
	}
	if so.ExcludedNetworks != nil {
		o.ExcludedNetworks = *so.ExcludedNetworks
	}
	if so.IgnoreExpression != nil {
		o.IgnoreExpression = *so.IgnoreExpression
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
	if so.TargetNetworks != nil {
		o.TargetNetworks = *so.TargetNetworks
	}
	if so.TargetUDPNetworks != nil {
		o.TargetUDPNetworks = *so.TargetUDPNetworks
	}
	if so.TrustedCAs != nil {
		o.TrustedCAs = *so.TrustedCAs
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

// DeepCopy returns a deep copy if the EnforcerProfile.
func (o *EnforcerProfile) DeepCopy() *EnforcerProfile {

	if o == nil {
		return nil
	}

	out := &EnforcerProfile{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *EnforcerProfile.
func (o *EnforcerProfile) DeepCopyInto(out *EnforcerProfile) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy EnforcerProfile: %s", err))
	}

	*out = *target.(*EnforcerProfile)
}

// Validate valides the current information stored into the structure.
func (o *EnforcerProfile) Validate() error {

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

	// Custom object validation.
	if err := ValidateEnforcerProfile(o); err != nil {
		switch e := err.(type) {
		case elemental.Errors:
			errors = append(errors, e...)
		case elemental.Error:
			errors = append(errors, e)
		default:
			errors = append(errors, e)
		}
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
func (*EnforcerProfile) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerProfileAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerProfileLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerProfile) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerProfileAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *EnforcerProfile) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "excludedInterfaces":
		return o.ExcludedInterfaces
	case "excludedNetworks":
		return o.ExcludedNetworks
	case "ignoreExpression":
		return o.IgnoreExpression
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
	case "targetNetworks":
		return o.TargetNetworks
	case "targetUDPNetworks":
		return o.TargetUDPNetworks
	case "trustedCAs":
		return o.TrustedCAs
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// EnforcerProfileAttributesMap represents the map of attribute for EnforcerProfile.
var EnforcerProfileAttributesMap = map[string]elemental.AttributeSpecification{
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
	"ExcludedInterfaces": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedInterfaces",
		Description:    `ExcludedInterfaces is a list of interfaces that must be excluded.`,
		Exposed:        true,
		Name:           "excludedInterfaces",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ExcludedNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedNetworks",
		Description: `ExcludedNetworks is the list of networks that must be excluded for this
enforcer.`,
		Exposed:   true,
		Name:      "excludedNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"IgnoreExpression": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IgnoreExpression",
		Description: `IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
docker container started with labels matching the rule.`,
		Exposed: true,
		Name:    "ignoreExpression",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
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
	"TargetNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNetworks",
		Description:    `TargetNetworks is the list of networks that authorization should be applied.`,
		Exposed:        true,
		Name:           "targetNetworks",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"TargetUDPNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetUDPNetworks",
		Description: `TargetUDPNetworks is the list of UDP networks that authorization should be
applied.`,
		Exposed:   true,
		Name:      "targetUDPNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"TrustedCAs": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TrustedCAs",
		Description:    `List of trusted CA. If empty the main chain of trust will be used.`,
		Exposed:        true,
		Name:           "trustedCAs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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

// EnforcerProfileLowerCaseAttributesMap represents the map of attribute for EnforcerProfile.
var EnforcerProfileLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"excludedinterfaces": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedInterfaces",
		Description:    `ExcludedInterfaces is a list of interfaces that must be excluded.`,
		Exposed:        true,
		Name:           "excludedInterfaces",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"excludednetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedNetworks",
		Description: `ExcludedNetworks is the list of networks that must be excluded for this
enforcer.`,
		Exposed:   true,
		Name:      "excludedNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"ignoreexpression": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IgnoreExpression",
		Description: `IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
docker container started with labels matching the rule.`,
		Exposed: true,
		Name:    "ignoreExpression",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
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
	"targetnetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNetworks",
		Description:    `TargetNetworks is the list of networks that authorization should be applied.`,
		Exposed:        true,
		Name:           "targetNetworks",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"targetudpnetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetUDPNetworks",
		Description: `TargetUDPNetworks is the list of UDP networks that authorization should be
applied.`,
		Exposed:   true,
		Name:      "targetUDPNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"trustedcas": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TrustedCAs",
		Description:    `List of trusted CA. If empty the main chain of trust will be used.`,
		Exposed:        true,
		Name:           "trustedCAs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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

// SparseEnforcerProfilesList represents a list of SparseEnforcerProfiles
type SparseEnforcerProfilesList []*SparseEnforcerProfile

// Identity returns the identity of the objects in the list.
func (o SparseEnforcerProfilesList) Identity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Copy returns a pointer to a copy the SparseEnforcerProfilesList.
func (o SparseEnforcerProfilesList) Copy() elemental.Identifiables {

	copy := append(SparseEnforcerProfilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEnforcerProfilesList.
func (o SparseEnforcerProfilesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEnforcerProfilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEnforcerProfile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEnforcerProfilesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEnforcerProfilesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseEnforcerProfilesList converted to EnforcerProfilesList.
func (o SparseEnforcerProfilesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEnforcerProfilesList) Version() int {

	return 1
}

// SparseEnforcerProfile represents the sparse version of a enforcerprofile.
type SparseEnforcerProfile struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// ExcludedInterfaces is a list of interfaces that must be excluded.
	ExcludedInterfaces *[]string `json:"excludedInterfaces,omitempty" bson:"excludedinterfaces" mapstructure:"excludedInterfaces,omitempty"`

	// ExcludedNetworks is the list of networks that must be excluded for this
	// enforcer.
	ExcludedNetworks *[]string `json:"excludedNetworks,omitempty" bson:"excludednetworks" mapstructure:"excludedNetworks,omitempty"`

	// IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
	// docker container started with labels matching the rule.
	IgnoreExpression *[][]string `json:"ignoreExpression,omitempty" bson:"ignoreexpression" mapstructure:"ignoreExpression,omitempty"`

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

	// TargetNetworks is the list of networks that authorization should be applied.
	TargetNetworks *[]string `json:"targetNetworks,omitempty" bson:"targetnetworks" mapstructure:"targetNetworks,omitempty"`

	// TargetUDPNetworks is the list of UDP networks that authorization should be
	// applied.
	TargetUDPNetworks *[]string `json:"targetUDPNetworks,omitempty" bson:"targetudpnetworks" mapstructure:"targetUDPNetworks,omitempty"`

	// List of trusted CA. If empty the main chain of trust will be used.
	TrustedCAs *[]string `json:"trustedCAs,omitempty" bson:"trustedcas" mapstructure:"trustedCAs,omitempty"`

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

// NewSparseEnforcerProfile returns a new  SparseEnforcerProfile.
func NewSparseEnforcerProfile() *SparseEnforcerProfile {
	return &SparseEnforcerProfile{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEnforcerProfile) Identity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEnforcerProfile) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEnforcerProfile) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseEnforcerProfile) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEnforcerProfile) ToPlain() elemental.PlainIdentifiable {

	out := NewEnforcerProfile()
	if o.ID != nil {
		out.ID = *o.ID
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
	if o.ExcludedInterfaces != nil {
		out.ExcludedInterfaces = *o.ExcludedInterfaces
	}
	if o.ExcludedNetworks != nil {
		out.ExcludedNetworks = *o.ExcludedNetworks
	}
	if o.IgnoreExpression != nil {
		out.IgnoreExpression = *o.IgnoreExpression
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
	if o.TargetNetworks != nil {
		out.TargetNetworks = *o.TargetNetworks
	}
	if o.TargetUDPNetworks != nil {
		out.TargetUDPNetworks = *o.TargetUDPNetworks
	}
	if o.TrustedCAs != nil {
		out.TrustedCAs = *o.TrustedCAs
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

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseEnforcerProfile) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseEnforcerProfile) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseEnforcerProfile) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseEnforcerProfile) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseEnforcerProfile) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseEnforcerProfile) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseEnforcerProfile) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseEnforcerProfile) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseEnforcerProfile) GetProtected() bool {

	return *o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseEnforcerProfile) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseEnforcerProfile) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseEnforcerProfile) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseEnforcerProfile.
func (o *SparseEnforcerProfile) DeepCopy() *SparseEnforcerProfile {

	if o == nil {
		return nil
	}

	out := &SparseEnforcerProfile{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEnforcerProfile.
func (o *SparseEnforcerProfile) DeepCopyInto(out *SparseEnforcerProfile) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEnforcerProfile: %s", err))
	}

	*out = *target.(*SparseEnforcerProfile)
}
