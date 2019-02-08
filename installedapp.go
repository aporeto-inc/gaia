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

	return []string{}
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
	// ID of the installed app.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// AccountName represents the vince account name.
	AccountName string `json:"accountName" bson:"accountname" mapstructure:"accountName,omitempty"`

	// AppIdentifier retains the identifier for the app.
	AppIdentifier string `json:"-" bson:"appidentifier" mapstructure:"-,omitempty"`

	// CategoryID of the app.
	CategoryID string `json:"categoryID" bson:"categoryid" mapstructure:"categoryID,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Version of the installed app.
	CurrentVersion string `json:"currentVersion" bson:"currentversion" mapstructure:"currentVersion,omitempty"`

	// DeploymentCount represents the number of expected deployment for this app.
	DeploymentCount int `json:"-" bson:"deploymentcount" mapstructure:"-,omitempty"`

	// Name of the installed app.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace in which the app is running.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Parameters is a list of parameters to start the app.
	Parameters []*AppParameter `json:"parameters" bson:"parameters" mapstructure:"parameters,omitempty"`

	// Status of the app.
	Status InstalledAppStatusValue `json:"status" bson:"status" mapstructure:"status,omitempty"`

	// Reason for the status of the app.
	StatusMessage string `json:"statusMessage" bson:"statusmessage" mapstructure:"statusMessage,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewInstalledApp returns a new *InstalledApp
func NewInstalledApp() *InstalledApp {

	return &InstalledApp{
		ModelVersion: 1,
		Parameters:   []*AppParameter{},
		Status:       InstalledAppStatusUnknown,
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

	return []string{}
}

// Doc returns the documentation for the object
func (o *InstalledApp) Doc() string {
	return `InstalledApps represents an installed application.`
}

func (o *InstalledApp) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *InstalledApp) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *InstalledApp) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *InstalledApp) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseInstalledApp{
			ID:              &o.ID,
			AccountName:     &o.AccountName,
			AppIdentifier:   &o.AppIdentifier,
			CategoryID:      &o.CategoryID,
			CreateTime:      &o.CreateTime,
			CurrentVersion:  &o.CurrentVersion,
			DeploymentCount: &o.DeploymentCount,
			Name:            &o.Name,
			Namespace:       &o.Namespace,
			Parameters:      &o.Parameters,
			Status:          &o.Status,
			StatusMessage:   &o.StatusMessage,
		}
	}

	sp := &SparseInstalledApp{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "accountName":
			sp.AccountName = &(o.AccountName)
		case "appIdentifier":
			sp.AppIdentifier = &(o.AppIdentifier)
		case "categoryID":
			sp.CategoryID = &(o.CategoryID)
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
		case "parameters":
			sp.Parameters = &(o.Parameters)
		case "status":
			sp.Status = &(o.Status)
		case "statusMessage":
			sp.StatusMessage = &(o.StatusMessage)
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
	if so.AccountName != nil {
		o.AccountName = *so.AccountName
	}
	if so.AppIdentifier != nil {
		o.AppIdentifier = *so.AppIdentifier
	}
	if so.CategoryID != nil {
		o.CategoryID = *so.CategoryID
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
	if so.Parameters != nil {
		o.Parameters = *so.Parameters
	}
	if so.Status != nil {
		o.Status = *so.Status
	}
	if so.StatusMessage != nil {
		o.StatusMessage = *so.StatusMessage
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
	case "accountName":
		return o.AccountName
	case "appIdentifier":
		return o.AppIdentifier
	case "categoryID":
		return o.CategoryID
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
	case "parameters":
		return o.Parameters
	case "status":
		return o.Status
	case "statusMessage":
		return o.StatusMessage
	}

	return nil
}

// InstalledAppAttributesMap represents the map of attribute for InstalledApp.
var InstalledAppAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID of the installed app.`,
		Exposed:        true,
		Identifier:     true,
		Name:           "ID",
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AccountName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccountName",
		CreationOnly:   true,
		Description:    `AccountName represents the vince account name.`,
		Exposed:        true,
		Name:           "accountName",
		Stored:         true,
		Type:           "string",
	},
	"AppIdentifier": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AppIdentifier",
		Description:    `AppIdentifier retains the identifier for the app.`,
		Name:           "appIdentifier",
		Stored:         true,
		Type:           "string",
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
		CreationOnly:   true,
		Description:    `Name of the installed app.`,
		Exposed:        true,
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace in which the app is running.`,
		Exposed:        true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
}

// InstalledAppLowerCaseAttributesMap represents the map of attribute for InstalledApp.
var InstalledAppLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID of the installed app.`,
		Exposed:        true,
		Identifier:     true,
		Name:           "ID",
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"accountname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccountName",
		CreationOnly:   true,
		Description:    `AccountName represents the vince account name.`,
		Exposed:        true,
		Name:           "accountName",
		Stored:         true,
		Type:           "string",
	},
	"appidentifier": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AppIdentifier",
		Description:    `AppIdentifier retains the identifier for the app.`,
		Name:           "appIdentifier",
		Stored:         true,
		Type:           "string",
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
		CreationOnly:   true,
		Description:    `Name of the installed app.`,
		Exposed:        true,
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace in which the app is running.`,
		Exposed:        true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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

	return []string{}
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
	// ID of the installed app.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// AccountName represents the vince account name.
	AccountName *string `json:"accountName,omitempty" bson:"accountname" mapstructure:"accountName,omitempty"`

	// AppIdentifier retains the identifier for the app.
	AppIdentifier *string `json:"-,omitempty" bson:"appidentifier" mapstructure:"-,omitempty"`

	// CategoryID of the app.
	CategoryID *string `json:"categoryID,omitempty" bson:"categoryid" mapstructure:"categoryID,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Version of the installed app.
	CurrentVersion *string `json:"currentVersion,omitempty" bson:"currentversion" mapstructure:"currentVersion,omitempty"`

	// DeploymentCount represents the number of expected deployment for this app.
	DeploymentCount *int `json:"-,omitempty" bson:"deploymentcount" mapstructure:"-,omitempty"`

	// Name of the installed app.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace in which the app is running.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Parameters is a list of parameters to start the app.
	Parameters *[]*AppParameter `json:"parameters,omitempty" bson:"parameters" mapstructure:"parameters,omitempty"`

	// Status of the app.
	Status *InstalledAppStatusValue `json:"status,omitempty" bson:"status" mapstructure:"status,omitempty"`

	// Reason for the status of the app.
	StatusMessage *string `json:"statusMessage,omitempty" bson:"statusmessage" mapstructure:"statusMessage,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
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
	if o.AccountName != nil {
		out.AccountName = *o.AccountName
	}
	if o.AppIdentifier != nil {
		out.AppIdentifier = *o.AppIdentifier
	}
	if o.CategoryID != nil {
		out.CategoryID = *o.CategoryID
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
	if o.Parameters != nil {
		out.Parameters = *o.Parameters
	}
	if o.Status != nil {
		out.Status = *o.Status
	}
	if o.StatusMessage != nil {
		out.StatusMessage = *o.StatusMessage
	}

	return out
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseInstalledApp) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
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
