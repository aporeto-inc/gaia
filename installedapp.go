package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// InstalledAppStatusValue represents the possible values for attribute "status".
type InstalledAppStatusValue string

const (
	// InstalledAppStatusDeploying represents the value Deploying.
	InstalledAppStatusDeploying InstalledAppStatusValue = "Deploying"

	// InstalledAppStatusError represents the value Error.
	InstalledAppStatusError InstalledAppStatusValue = "Error"

	// InstalledAppStatusInitializing represents the value Initializing.
	InstalledAppStatusInitializing InstalledAppStatusValue = "Initializing"

	// InstalledAppStatusRunning represents the value Running.
	InstalledAppStatusRunning InstalledAppStatusValue = "Running"

	// InstalledAppStatusUndeploying represents the value Undeploying.
	InstalledAppStatusUndeploying InstalledAppStatusValue = "Undeploying"

	// InstalledAppStatusUnknown represents the value Unknown.
	InstalledAppStatusUnknown InstalledAppStatusValue = "Unknown"
)

// InstalledAppIdentity represents the Identity of the object.
var InstalledAppIdentity = elemental.Identity{
	Name:     "installedapp",
	Category: "installedapps",
	Package:  "highwind",
	Private:  false,
}

// InstalledAppsList represents a list of InstalledApps
type InstalledAppsList []*InstalledApp

// Identity returns the identity of the objects in the list.
func (o InstalledAppsList) Identity() elemental.Identity {

	return InstalledAppIdentity
}

// Copy returns a pointer to a copy the InstalledAppsList.
func (o InstalledAppsList) Copy() elemental.Identifiables {

	copy := append(InstalledAppsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the InstalledAppsList.
func (o InstalledAppsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(InstalledAppsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*InstalledApp))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o InstalledAppsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o InstalledAppsList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the InstalledAppsList converted to SparseInstalledAppsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o InstalledAppsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o InstalledAppsList) Version() int {

	return 1
}

// InstalledApp represents the model of a installedapp
type InstalledApp struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AppIdentifier retains the identifier for the app.
	AppIdentifier string `json:"-" bson:"appidentifier" mapstructure:"-,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CategoryID of the app.
	CategoryID string `json:"categoryID" bson:"categoryid" mapstructure:"categoryID,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Version of the installed app.
	CurrentVersion string `json:"currentVersion" bson:"currentversion" mapstructure:"currentVersion,omitempty"`

	// DeploymentCount represents the number of expected deployment for this app.
	DeploymentCount int `json:"-" bson:"deploymentcount" mapstructure:"-,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Parameters is a list of parameters to start the app.
	Parameters []*AppParameter `json:"parameters" bson:"parameters" mapstructure:"parameters,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Status of the app.
	Status InstalledAppStatusValue `json:"status" bson:"status" mapstructure:"status,omitempty"`

	// Reason for the status of the app.
	StatusMessage string `json:"statusMessage" bson:"statusmessage" mapstructure:"statusMessage,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewInstalledApp returns a new *InstalledApp
func NewInstalledApp() *InstalledApp {

	return &InstalledApp{
		ModelVersion:   1,
		Mutex:          &sync.Mutex{},
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		NormalizedTags: []string{},
		Parameters:     []*AppParameter{},
		Status:         InstalledAppStatusUnknown,
	}
}

// Identity returns the Identity of the object.
func (o *InstalledApp) Identity() elemental.Identity {

	return InstalledAppIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *InstalledApp) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *InstalledApp) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *InstalledApp) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *InstalledApp) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// Doc returns the documentation for the object
func (o *InstalledApp) Doc() string {

	return `InstalledApps represents an installed application.`
}

func (o *InstalledApp) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *InstalledApp) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *InstalledApp) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *InstalledApp) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *InstalledApp) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *InstalledApp) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *InstalledApp) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *InstalledApp) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *InstalledApp) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetName returns the Name of the receiver.
func (o *InstalledApp) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *InstalledApp) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *InstalledApp) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *InstalledApp) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *InstalledApp) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *InstalledApp) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *InstalledApp) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *InstalledApp) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *InstalledApp) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *InstalledApp) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *InstalledApp) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *InstalledApp) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *InstalledApp) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseInstalledApp{
			ID:                   &o.ID,
			Annotations:          &o.Annotations,
			AppIdentifier:        &o.AppIdentifier,
			AssociatedTags:       &o.AssociatedTags,
			CategoryID:           &o.CategoryID,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			CurrentVersion:       &o.CurrentVersion,
			DeploymentCount:      &o.DeploymentCount,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Parameters:           &o.Parameters,
			Protected:            &o.Protected,
			Status:               &o.Status,
			StatusMessage:        &o.StatusMessage,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
		}
	}

	sp := &SparseInstalledApp{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "appIdentifier":
			sp.AppIdentifier = &(o.AppIdentifier)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "categoryID":
			sp.CategoryID = &(o.CategoryID)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "currentVersion":
			sp.CurrentVersion = &(o.CurrentVersion)
		case "deploymentCount":
			sp.DeploymentCount = &(o.DeploymentCount)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "parameters":
			sp.Parameters = &(o.Parameters)
		case "protected":
			sp.Protected = &(o.Protected)
		case "status":
			sp.Status = &(o.Status)
		case "statusMessage":
			sp.StatusMessage = &(o.StatusMessage)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseInstalledApp to the object.
func (o *InstalledApp) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseInstalledApp)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AppIdentifier != nil {
		o.AppIdentifier = *so.AppIdentifier
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CategoryID != nil {
		o.CategoryID = *so.CategoryID
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.CurrentVersion != nil {
		o.CurrentVersion = *so.CurrentVersion
	}
	if so.DeploymentCount != nil {
		o.DeploymentCount = *so.DeploymentCount
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
	if so.Parameters != nil {
		o.Parameters = *so.Parameters
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Status != nil {
		o.Status = *so.Status
	}
	if so.StatusMessage != nil {
		o.StatusMessage = *so.StatusMessage
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the InstalledApp.
func (o *InstalledApp) DeepCopy() *InstalledApp {

	if o == nil {
		return nil
	}

	out := &InstalledApp{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *InstalledApp.
func (o *InstalledApp) DeepCopyInto(out *InstalledApp) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy InstalledApp: %s", err))
	}

	*out = *target.(*InstalledApp)
}

// Validate valides the current information stored into the structure.
func (o *InstalledApp) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = append(errors, err)
	}

	for _, sub := range o.Parameters {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Unknown", "Deploying", "Initializing", "Running", "Undeploying", "Error"}, false); err != nil {
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
func (*InstalledApp) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := InstalledAppAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return InstalledAppLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*InstalledApp) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return InstalledAppAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *InstalledApp) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "appIdentifier":
		return o.AppIdentifier
	case "associatedTags":
		return o.AssociatedTags
	case "categoryID":
		return o.CategoryID
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "currentVersion":
		return o.CurrentVersion
	case "deploymentCount":
		return o.DeploymentCount
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "parameters":
		return o.Parameters
	case "protected":
		return o.Protected
	case "status":
		return o.Status
	case "statusMessage":
		return o.StatusMessage
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	}

	return nil
}

// InstalledAppAttributesMap represents the map of attribute for InstalledApp.
var InstalledAppAttributesMap = map[string]elemental.AttributeSpecification{
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
	"AppIdentifier": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AppIdentifier",
		Description:    `AppIdentifier retains the identifier for the app.`,
		Name:           "appIdentifier",
		Stored:         true,
		Type:           "string",
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
	"CategoryID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CategoryID",
		Description:    `CategoryID of the app.`,
		Exposed:        true,
		Name:           "categoryID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
	"CurrentVersion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CurrentVersion",
		Description:    `Version of the installed app.`,
		Exposed:        true,
		Name:           "currentVersion",
		Stored:         true,
		Type:           "string",
	},
	"DeploymentCount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DeploymentCount",
		Description:    `DeploymentCount represents the number of expected deployment for this app.`,
		Name:           "deploymentCount",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
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
		DefaultOrder:   true,
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
	"Parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters is a list of parameters to start the app.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "appparameter",
		Type:           "refList",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Unknown", "Deploying", "Initializing", "Running", "Undeploying", "Error"},
		ConvertedName:  "Status",
		DefaultValue:   InstalledAppStatusUnknown,
		Description:    `Status of the app.`,
		Exposed:        true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"StatusMessage": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "StatusMessage",
		Description:    `Reason for the status of the app.`,
		Exposed:        true,
		Name:           "statusMessage",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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

// InstalledAppLowerCaseAttributesMap represents the map of attribute for InstalledApp.
var InstalledAppLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"appidentifier": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AppIdentifier",
		Description:    `AppIdentifier retains the identifier for the app.`,
		Name:           "appIdentifier",
		Stored:         true,
		Type:           "string",
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
	"categoryid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CategoryID",
		Description:    `CategoryID of the app.`,
		Exposed:        true,
		Name:           "categoryID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
	"currentversion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CurrentVersion",
		Description:    `Version of the installed app.`,
		Exposed:        true,
		Name:           "currentVersion",
		Stored:         true,
		Type:           "string",
	},
	"deploymentcount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DeploymentCount",
		Description:    `DeploymentCount represents the number of expected deployment for this app.`,
		Name:           "deploymentCount",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
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
		DefaultOrder:   true,
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
	"parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters is a list of parameters to start the app.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "appparameter",
		Type:           "refList",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Unknown", "Deploying", "Initializing", "Running", "Undeploying", "Error"},
		ConvertedName:  "Status",
		DefaultValue:   InstalledAppStatusUnknown,
		Description:    `Status of the app.`,
		Exposed:        true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"statusmessage": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "StatusMessage",
		Description:    `Reason for the status of the app.`,
		Exposed:        true,
		Name:           "statusMessage",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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

// SparseInstalledAppsList represents a list of SparseInstalledApps
type SparseInstalledAppsList []*SparseInstalledApp

// Identity returns the identity of the objects in the list.
func (o SparseInstalledAppsList) Identity() elemental.Identity {

	return InstalledAppIdentity
}

// Copy returns a pointer to a copy the SparseInstalledAppsList.
func (o SparseInstalledAppsList) Copy() elemental.Identifiables {

	copy := append(SparseInstalledAppsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseInstalledAppsList.
func (o SparseInstalledAppsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseInstalledAppsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseInstalledApp))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseInstalledAppsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseInstalledAppsList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparseInstalledAppsList converted to InstalledAppsList.
func (o SparseInstalledAppsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseInstalledAppsList) Version() int {

	return 1
}

// SparseInstalledApp represents the sparse version of a installedapp.
type SparseInstalledApp struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// AppIdentifier retains the identifier for the app.
	AppIdentifier *string `json:"-" bson:"appidentifier,omitempty" mapstructure:"-,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// CategoryID of the app.
	CategoryID *string `json:"categoryID,omitempty" bson:"categoryid,omitempty" mapstructure:"categoryID,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Version of the installed app.
	CurrentVersion *string `json:"currentVersion,omitempty" bson:"currentversion,omitempty" mapstructure:"currentVersion,omitempty"`

	// DeploymentCount represents the number of expected deployment for this app.
	DeploymentCount *int `json:"-" bson:"deploymentcount,omitempty" mapstructure:"-,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Parameters is a list of parameters to start the app.
	Parameters *[]*AppParameter `json:"parameters,omitempty" bson:"parameters,omitempty" mapstructure:"parameters,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Status of the app.
	Status *InstalledAppStatusValue `json:"status,omitempty" bson:"status,omitempty" mapstructure:"status,omitempty"`

	// Reason for the status of the app.
	StatusMessage *string `json:"statusMessage,omitempty" bson:"statusmessage,omitempty" mapstructure:"statusMessage,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSparseInstalledApp returns a new  SparseInstalledApp.
func NewSparseInstalledApp() *SparseInstalledApp {
	return &SparseInstalledApp{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseInstalledApp) Identity() elemental.Identity {

	return InstalledAppIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseInstalledApp) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseInstalledApp) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseInstalledApp) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseInstalledApp) ToPlain() elemental.PlainIdentifiable {

	out := NewInstalledApp()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AppIdentifier != nil {
		out.AppIdentifier = *o.AppIdentifier
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CategoryID != nil {
		out.CategoryID = *o.CategoryID
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.CurrentVersion != nil {
		out.CurrentVersion = *o.CurrentVersion
	}
	if o.DeploymentCount != nil {
		out.DeploymentCount = *o.DeploymentCount
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
	if o.Parameters != nil {
		out.Parameters = *o.Parameters
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Status != nil {
		out.Status = *o.Status
	}
	if o.StatusMessage != nil {
		out.StatusMessage = *o.StatusMessage
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseInstalledApp) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseInstalledApp) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseInstalledApp) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseInstalledApp) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetName returns the Name of the receiver.
func (o *SparseInstalledApp) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseInstalledApp) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseInstalledApp) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseInstalledApp) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseInstalledApp) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseInstalledApp) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseInstalledApp.
func (o *SparseInstalledApp) DeepCopy() *SparseInstalledApp {

	if o == nil {
		return nil
	}

	out := &SparseInstalledApp{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseInstalledApp.
func (o *SparseInstalledApp) DeepCopyInto(out *SparseInstalledApp) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseInstalledApp: %s", err))
	}

	*out = *target.(*SparseInstalledApp)
}
