package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FileAccessPolicyIdentity represents the Identity of the object.
var FileAccessPolicyIdentity = elemental.Identity{
	Name:     "fileaccesspolicy",
	Category: "fileaccesspolicies",
	Package:  "squall",
	Private:  false,
}

// FileAccessPoliciesList represents a list of FileAccessPolicies
type FileAccessPoliciesList []*FileAccessPolicy

// Identity returns the identity of the objects in the list.
func (o FileAccessPoliciesList) Identity() elemental.Identity {

	return FileAccessPolicyIdentity
}

// Copy returns a pointer to a copy the FileAccessPoliciesList.
func (o FileAccessPoliciesList) Copy() elemental.Identifiables {

	copy := append(FileAccessPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the FileAccessPoliciesList.
func (o FileAccessPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(FileAccessPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FileAccessPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FileAccessPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FileAccessPoliciesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the FileAccessPoliciesList converted to SparseFileAccessPoliciesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o FileAccessPoliciesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseFileAccessPoliciesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseFileAccessPolicy)
	}

	return out
}

// Version returns the version of the content.
func (o FileAccessPoliciesList) Version() int {

	return 1
}

// FileAccessPolicy represents the model of a fileaccesspolicy
type FileAccessPolicy struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Defines for how long the policy will be active according to the
	// `activeSchedule`.
	ActiveDuration string `json:"activeDuration" msgpack:"activeDuration" bson:"activeduration" mapstructure:"activeDuration,omitempty"`

	// Defines when the policy should be active using the cron notation.
	// The policy will be active for the given `activeDuration`.
	ActiveSchedule string `json:"activeSchedule" msgpack:"activeSchedule" bson:"activeschedule" mapstructure:"activeSchedule,omitempty"`

	// Allows files to be executed.
	AllowsExecute bool `json:"allowsExecute" msgpack:"allowsExecute" bson:"-" mapstructure:"allowsExecute,omitempty"`

	// Allows files to be read.
	AllowsRead bool `json:"allowsRead" msgpack:"allowsRead" bson:"-" mapstructure:"allowsRead,omitempty"`

	// Allows files to be written.
	AllowsWrite bool `json:"allowsWrite" msgpack:"allowsWrite" bson:"-" mapstructure:"allowsWrite,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled bool `json:"disabled" msgpack:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// Set to `true` to enable automatic encryption.
	EncryptionEnabled bool `json:"encryptionEnabled" msgpack:"encryptionEnabled" bson:"-" mapstructure:"encryptionEnabled,omitempty"`

	// If set the policy will be automatically deleted after the given time.
	ExpirationTime time.Time `json:"expirationTime" msgpack:"expirationTime" bson:"expirationtime" mapstructure:"expirationTime,omitempty"`

	// Indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback bool `json:"fallback" msgpack:"fallback" bson:"fallback" mapstructure:"fallback,omitempty"`

	// A value of `true` enables logging.
	LogsEnabled bool `json:"logsEnabled" msgpack:"logsEnabled" bson:"-" mapstructure:"logsEnabled,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// The object of the policy.
	Object [][]string `json:"object" msgpack:"object" bson:"-" mapstructure:"object,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// The subject of the policy.
	Subject [][]string `json:"subject" msgpack:"subject" bson:"-" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFileAccessPolicy returns a new *FileAccessPolicy
func NewFileAccessPolicy() *FileAccessPolicy {

	return &FileAccessPolicy{
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

// BleveType implements the bleve.Classifier Interface.
func (o *FileAccessPolicy) BleveType() string {

	return "fileaccesspolicy"
}

// DefaultOrder returns the list of default ordering fields.
func (o *FileAccessPolicy) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// Doc returns the documentation for the object
func (o *FileAccessPolicy) Doc() string {

	return `A file access policy allows processing units to access various folder and files.
It will use the tags of a file path to know what is the path of the file or
folder to allow access to. You can allow the processing unit to have any
combination of read, write, or execute.

When a processing unit is a Docker container, then it will police the volumes.
Mount and execute won't have any effect.

File paths are not supported yet for standard Linux processes.`
}

func (o *FileAccessPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *FileAccessPolicy) GetActiveDuration() string {

	return o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the given value.
func (o *FileAccessPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *FileAccessPolicy) GetActiveSchedule() string {

	return o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the given value.
func (o *FileAccessPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *FileAccessPolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *FileAccessPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *FileAccessPolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *FileAccessPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *FileAccessPolicy) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *FileAccessPolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *FileAccessPolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *FileAccessPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *FileAccessPolicy) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *FileAccessPolicy) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *FileAccessPolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *FileAccessPolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *FileAccessPolicy) GetExpirationTime() time.Time {

	return o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the given value.
func (o *FileAccessPolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *FileAccessPolicy) GetFallback() bool {

	return o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the given value.
func (o *FileAccessPolicy) SetFallback(fallback bool) {

	o.Fallback = fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *FileAccessPolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *FileAccessPolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *FileAccessPolicy) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *FileAccessPolicy) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *FileAccessPolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *FileAccessPolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *FileAccessPolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *FileAccessPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *FileAccessPolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *FileAccessPolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *FileAccessPolicy) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *FileAccessPolicy) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *FileAccessPolicy) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *FileAccessPolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *FileAccessPolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *FileAccessPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *FileAccessPolicy) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseFileAccessPolicy{
			ID:                   &o.ID,
			ActiveDuration:       &o.ActiveDuration,
			ActiveSchedule:       &o.ActiveSchedule,
			AllowsExecute:        &o.AllowsExecute,
			AllowsRead:           &o.AllowsRead,
			AllowsWrite:          &o.AllowsWrite,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Description:          &o.Description,
			Disabled:             &o.Disabled,
			EncryptionEnabled:    &o.EncryptionEnabled,
			ExpirationTime:       &o.ExpirationTime,
			Fallback:             &o.Fallback,
			LogsEnabled:          &o.LogsEnabled,
			Metadata:             &o.Metadata,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Object:               &o.Object,
			Propagate:            &o.Propagate,
			Protected:            &o.Protected,
			Subject:              &o.Subject,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
		}
	}

	sp := &SparseFileAccessPolicy{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "activeDuration":
			sp.ActiveDuration = &(o.ActiveDuration)
		case "activeSchedule":
			sp.ActiveSchedule = &(o.ActiveSchedule)
		case "allowsExecute":
			sp.AllowsExecute = &(o.AllowsExecute)
		case "allowsRead":
			sp.AllowsRead = &(o.AllowsRead)
		case "allowsWrite":
			sp.AllowsWrite = &(o.AllowsWrite)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "encryptionEnabled":
			sp.EncryptionEnabled = &(o.EncryptionEnabled)
		case "expirationTime":
			sp.ExpirationTime = &(o.ExpirationTime)
		case "fallback":
			sp.Fallback = &(o.Fallback)
		case "logsEnabled":
			sp.LogsEnabled = &(o.LogsEnabled)
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
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseFileAccessPolicy to the object.
func (o *FileAccessPolicy) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseFileAccessPolicy)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.ActiveDuration != nil {
		o.ActiveDuration = *so.ActiveDuration
	}
	if so.ActiveSchedule != nil {
		o.ActiveSchedule = *so.ActiveSchedule
	}
	if so.AllowsExecute != nil {
		o.AllowsExecute = *so.AllowsExecute
	}
	if so.AllowsRead != nil {
		o.AllowsRead = *so.AllowsRead
	}
	if so.AllowsWrite != nil {
		o.AllowsWrite = *so.AllowsWrite
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
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
	if so.EncryptionEnabled != nil {
		o.EncryptionEnabled = *so.EncryptionEnabled
	}
	if so.ExpirationTime != nil {
		o.ExpirationTime = *so.ExpirationTime
	}
	if so.Fallback != nil {
		o.Fallback = *so.Fallback
	}
	if so.LogsEnabled != nil {
		o.LogsEnabled = *so.LogsEnabled
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
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the FileAccessPolicy.
func (o *FileAccessPolicy) DeepCopy() *FileAccessPolicy {

	if o == nil {
		return nil
	}

	out := &FileAccessPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FileAccessPolicy.
func (o *FileAccessPolicy) DeepCopyInto(out *FileAccessPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FileAccessPolicy: %s", err))
	}

	*out = *target.(*FileAccessPolicy)
}

// Validate valides the current information stored into the structure.
func (o *FileAccessPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidatePattern("activeDuration", o.ActiveDuration, `^[0-9]+[smh]$`, `must be a valid duration like <n>s or <n>s or <n>h`, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("object", o.Object); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("subject", o.Subject); err != nil {
		errors = errors.Append(err)
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FileAccessPolicy) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "activeDuration":
		return o.ActiveDuration
	case "activeSchedule":
		return o.ActiveSchedule
	case "allowsExecute":
		return o.AllowsExecute
	case "allowsRead":
		return o.AllowsRead
	case "allowsWrite":
		return o.AllowsWrite
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "disabled":
		return o.Disabled
	case "encryptionEnabled":
		return o.EncryptionEnabled
	case "expirationTime":
		return o.ExpirationTime
	case "fallback":
		return o.Fallback
	case "logsEnabled":
		return o.LogsEnabled
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
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	}

	return nil
}

// FileAccessPolicyAttributesMap represents the map of attribute for FileAccessPolicy.
var FileAccessPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
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
		Description: `Defines for how long the policy will be active according to the
` + "`" + `activeSchedule` + "`" + `.`,
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
		Description: `Defines when the policy should be active using the cron notation.
The policy will be active for the given ` + "`" + `activeDuration` + "`" + `.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"AllowsExecute": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsExecute",
		Description:    `Allows files to be executed.`,
		Exposed:        true,
		Name:           "allowsExecute",
		Orderable:      true,
		Type:           "boolean",
	},
	"AllowsRead": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsRead",
		Description:    `Allows files to be read.`,
		Exposed:        true,
		Name:           "allowsRead",
		Orderable:      true,
		Type:           "boolean",
	},
	"AllowsWrite": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsWrite",
		Description:    `Allows files to be written.`,
		Exposed:        true,
		Name:           "allowsWrite",
		Orderable:      true,
		Type:           "boolean",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
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
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CreateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
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
		Description:    `Description of the object.`,
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
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
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
		Description:    `Set to ` + "`" + `true` + "`" + ` to enable automatic encryption.`,
		Exposed:        true,
		Name:           "encryptionEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"ExpirationTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `If set the policy will be automatically deleted after the given time.`,
		Exposed:        true,
		Getter:         true,
		Name:           "expirationTime",
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Indicates that this is fallback policy. It will only be
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
	"LogsEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LogsEnabled",
		Description:    `A value of ` + "`" + `true` + "`" + ` enables logging.`,
		Exposed:        true,
		Name:           "logsEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
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
		Description:    `Name of the entity.`,
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
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
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
		Description:    `The object of the policy.`,
		Exposed:        true,
		Name:           "object",
		Orderable:      true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
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
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `The subject of the policy.`,
		Exposed:        true,
		Name:           "subject",
		Orderable:      true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"UpdateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
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
		Description:    `Identifier of the object.`,
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
		Description: `Defines for how long the policy will be active according to the
` + "`" + `activeSchedule` + "`" + `.`,
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
		Description: `Defines when the policy should be active using the cron notation.
The policy will be active for the given ` + "`" + `activeDuration` + "`" + `.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"allowsexecute": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsExecute",
		Description:    `Allows files to be executed.`,
		Exposed:        true,
		Name:           "allowsExecute",
		Orderable:      true,
		Type:           "boolean",
	},
	"allowsread": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsRead",
		Description:    `Allows files to be read.`,
		Exposed:        true,
		Name:           "allowsRead",
		Orderable:      true,
		Type:           "boolean",
	},
	"allowswrite": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsWrite",
		Description:    `Allows files to be written.`,
		Exposed:        true,
		Name:           "allowsWrite",
		Orderable:      true,
		Type:           "boolean",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
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
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"createidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
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
		Description:    `Description of the object.`,
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
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
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
		Description:    `Set to ` + "`" + `true` + "`" + ` to enable automatic encryption.`,
		Exposed:        true,
		Name:           "encryptionEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"expirationtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `If set the policy will be automatically deleted after the given time.`,
		Exposed:        true,
		Getter:         true,
		Name:           "expirationTime",
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Indicates that this is fallback policy. It will only be
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
	"logsenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LogsEnabled",
		Description:    `A value of ` + "`" + `true` + "`" + ` enables logging.`,
		Exposed:        true,
		Name:           "logsEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
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
		Description:    `Name of the entity.`,
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
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
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
		Description:    `The object of the policy.`,
		Exposed:        true,
		Name:           "object",
		Orderable:      true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
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
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `The subject of the policy.`,
		Exposed:        true,
		Name:           "subject",
		Orderable:      true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"updateidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
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

// SparseFileAccessPoliciesList represents a list of SparseFileAccessPolicies
type SparseFileAccessPoliciesList []*SparseFileAccessPolicy

// Identity returns the identity of the objects in the list.
func (o SparseFileAccessPoliciesList) Identity() elemental.Identity {

	return FileAccessPolicyIdentity
}

// Copy returns a pointer to a copy the SparseFileAccessPoliciesList.
func (o SparseFileAccessPoliciesList) Copy() elemental.Identifiables {

	copy := append(SparseFileAccessPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseFileAccessPoliciesList.
func (o SparseFileAccessPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseFileAccessPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseFileAccessPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseFileAccessPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseFileAccessPoliciesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparseFileAccessPoliciesList converted to FileAccessPoliciesList.
func (o SparseFileAccessPoliciesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseFileAccessPoliciesList) Version() int {

	return 1
}

// SparseFileAccessPolicy represents the sparse version of a fileaccesspolicy.
type SparseFileAccessPolicy struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Defines for how long the policy will be active according to the
	// `activeSchedule`.
	ActiveDuration *string `json:"activeDuration,omitempty" msgpack:"activeDuration,omitempty" bson:"activeduration,omitempty" mapstructure:"activeDuration,omitempty"`

	// Defines when the policy should be active using the cron notation.
	// The policy will be active for the given `activeDuration`.
	ActiveSchedule *string `json:"activeSchedule,omitempty" msgpack:"activeSchedule,omitempty" bson:"activeschedule,omitempty" mapstructure:"activeSchedule,omitempty"`

	// Allows files to be executed.
	AllowsExecute *bool `json:"allowsExecute,omitempty" msgpack:"allowsExecute,omitempty" bson:"-" mapstructure:"allowsExecute,omitempty"`

	// Allows files to be read.
	AllowsRead *bool `json:"allowsRead,omitempty" msgpack:"allowsRead,omitempty" bson:"-" mapstructure:"allowsRead,omitempty"`

	// Allows files to be written.
	AllowsWrite *bool `json:"allowsWrite,omitempty" msgpack:"allowsWrite,omitempty" bson:"-" mapstructure:"allowsWrite,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled *bool `json:"disabled,omitempty" msgpack:"disabled,omitempty" bson:"disabled,omitempty" mapstructure:"disabled,omitempty"`

	// Set to `true` to enable automatic encryption.
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty" msgpack:"encryptionEnabled,omitempty" bson:"-" mapstructure:"encryptionEnabled,omitempty"`

	// If set the policy will be automatically deleted after the given time.
	ExpirationTime *time.Time `json:"expirationTime,omitempty" msgpack:"expirationTime,omitempty" bson:"expirationtime,omitempty" mapstructure:"expirationTime,omitempty"`

	// Indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback *bool `json:"fallback,omitempty" msgpack:"fallback,omitempty" bson:"fallback,omitempty" mapstructure:"fallback,omitempty"`

	// A value of `true` enables logging.
	LogsEnabled *bool `json:"logsEnabled,omitempty" msgpack:"logsEnabled,omitempty" bson:"-" mapstructure:"logsEnabled,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// The object of the policy.
	Object *[][]string `json:"object,omitempty" msgpack:"object,omitempty" bson:"-" mapstructure:"object,omitempty"`

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// The subject of the policy.
	Subject *[][]string `json:"subject,omitempty" msgpack:"subject,omitempty" bson:"-" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseFileAccessPolicy returns a new  SparseFileAccessPolicy.
func NewSparseFileAccessPolicy() *SparseFileAccessPolicy {
	return &SparseFileAccessPolicy{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseFileAccessPolicy) Identity() elemental.Identity {

	return FileAccessPolicyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseFileAccessPolicy) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseFileAccessPolicy) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseFileAccessPolicy) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseFileAccessPolicy) ToPlain() elemental.PlainIdentifiable {

	out := NewFileAccessPolicy()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.ActiveDuration != nil {
		out.ActiveDuration = *o.ActiveDuration
	}
	if o.ActiveSchedule != nil {
		out.ActiveSchedule = *o.ActiveSchedule
	}
	if o.AllowsExecute != nil {
		out.AllowsExecute = *o.AllowsExecute
	}
	if o.AllowsRead != nil {
		out.AllowsRead = *o.AllowsRead
	}
	if o.AllowsWrite != nil {
		out.AllowsWrite = *o.AllowsWrite
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
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
	if o.EncryptionEnabled != nil {
		out.EncryptionEnabled = *o.EncryptionEnabled
	}
	if o.ExpirationTime != nil {
		out.ExpirationTime = *o.ExpirationTime
	}
	if o.Fallback != nil {
		out.Fallback = *o.Fallback
	}
	if o.LogsEnabled != nil {
		out.LogsEnabled = *o.LogsEnabled
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
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *SparseFileAccessPolicy) GetActiveDuration() string {

	return *o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = &activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *SparseFileAccessPolicy) GetActiveSchedule() string {

	return *o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = &activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseFileAccessPolicy) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseFileAccessPolicy) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseFileAccessPolicy) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseFileAccessPolicy) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseFileAccessPolicy) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseFileAccessPolicy) GetDisabled() bool {

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *SparseFileAccessPolicy) GetExpirationTime() time.Time {

	return *o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = &expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *SparseFileAccessPolicy) GetFallback() bool {

	return *o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetFallback(fallback bool) {

	o.Fallback = &fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseFileAccessPolicy) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseFileAccessPolicy) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseFileAccessPolicy) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseFileAccessPolicy) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseFileAccessPolicy) GetPropagate() bool {

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparseFileAccessPolicy) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseFileAccessPolicy) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseFileAccessPolicy) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseFileAccessPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseFileAccessPolicy.
func (o *SparseFileAccessPolicy) DeepCopy() *SparseFileAccessPolicy {

	if o == nil {
		return nil
	}

	out := &SparseFileAccessPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseFileAccessPolicy.
func (o *SparseFileAccessPolicy) DeepCopyInto(out *SparseFileAccessPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseFileAccessPolicy: %s", err))
	}

	*out = *target.(*SparseFileAccessPolicy)
}
