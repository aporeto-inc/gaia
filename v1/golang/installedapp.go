package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/v1/golang/types"
)

// InstalledAppStatusValue represents the possible values for attribute "status".
type InstalledAppStatusValue string

const (
	// InstalledAppStatusError represents the value Error.
	InstalledAppStatusError InstalledAppStatusValue = "Error"

	// InstalledAppStatusPending represents the value Pending.
	InstalledAppStatusPending InstalledAppStatusValue = "Pending"

	// InstalledAppStatusRunning represents the value Running.
	InstalledAppStatusRunning InstalledAppStatusValue = "Running"
)

// InstalledAppIdentity represents the Identity of the object.
var InstalledAppIdentity = elemental.Identity{
	Name:     "installedapp",
	Category: "installedapps",
	Private:  false,
}

// InstalledAppsList represents a list of InstalledApps
type InstalledAppsList []*InstalledApp

// ContentIdentity returns the identity of the objects in the list.
func (o InstalledAppsList) ContentIdentity() elemental.Identity {

	return InstalledAppIdentity
}

// Copy returns a pointer to a copy the InstalledAppsList.
func (o InstalledAppsList) Copy() elemental.ContentIdentifiable {

	copy := append(InstalledAppsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the InstalledAppsList.
func (o InstalledAppsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(InstalledAppsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*InstalledApp))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o InstalledAppsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o InstalledAppsList) DefaultOrder() []string {

	return []string{}
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

	// CategoryID of the app.
	CategoryID string `json:"categoryID" bson:"categoryid" mapstructure:"categoryID,omitempty"`

	// Data retains all data created to use this service.
	Data interface{} `json:"-" bson:"data" mapstructure:"-,omitempty"`

	// K8SIdentifier retains the identifier for kubernetes.
	K8sIdentifier string `json:"-" bson:"k8sidentifier" mapstructure:"-,omitempty"`

	// Name of the installed app.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace in which the app is running.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Parameters is a list of parameters to start the app.
	Parameters []*types.AppParameter `json:"parameters" bson:"parameters" mapstructure:"parameters,omitempty"`

	// RelatedObjects retains all objects created to use this app.
	RelatedObjects []*types.AppRelatedObject `json:"-" bson:"relatedobjects" mapstructure:"-,omitempty"`

	// Replicas represents the number of replicas for the app.
	Replicas int `json:"replicas" bson:"replicas" mapstructure:"replicas,omitempty"`

	// Status of the app.
	Status InstalledAppStatusValue `json:"status" bson:"status" mapstructure:"status,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewInstalledApp returns a new *InstalledApp
func NewInstalledApp() *InstalledApp {

	return &InstalledApp{
		ModelVersion:   1,
		Data:           nil,
		Parameters:     []*types.AppParameter{},
		RelatedObjects: []*types.AppRelatedObject{},
		Status:         "Pending",
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

// Validate valides the current information stored into the structure.
func (o *InstalledApp) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Error", "Pending", "Running"}, false); err != nil {
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

// InstalledAppAttributesMap represents the map of attribute for InstalledApp.
var InstalledAppAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID of the installed app.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
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
		Format:         "free",
		Name:           "accountName",
		Stored:         true,
		Type:           "string",
	},
	"CategoryID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CategoryID",
		Description:    `CategoryID of the app.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "categoryID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Data retains all data created to use this service.`,
		Name:           "data",
		Stored:         true,
		SubType:        "service_data",
		Type:           "external",
	},
	"K8sIdentifier": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "K8sIdentifier",
		Description:    `K8SIdentifier retains the identifier for kubernetes.`,
		Format:         "free",
		Name:           "k8sIdentifier",
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		Description:    `Name of the installed app.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
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
		Filterable:     true,
		Format:         "free",
		Name:           "namespace",
		Orderable:      true,
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
		SubType:        "app_parameters",
		Type:           "external",
	},
	"RelatedObjects": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RelatedObjects",
		Description:    `RelatedObjects retains all objects created to use this app.`,
		Name:           "relatedObjects",
		Stored:         true,
		SubType:        "app_relatedobjects",
		Type:           "external",
	},
	"Replicas": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Replicas",
		Description:    `Replicas represents the number of replicas for the app.`,
		Exposed:        true,
		Name:           "replicas",
		Stored:         true,
		Type:           "integer",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Error", "Pending", "Running"},
		ConvertedName:  "Status",
		DefaultValue:   InstalledAppStatusPending,
		Description:    `Status of the app.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
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
		Filterable:     true,
		Format:         "free",
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
		Format:         "free",
		Name:           "accountName",
		Stored:         true,
		Type:           "string",
	},
	"categoryid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CategoryID",
		Description:    `CategoryID of the app.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "categoryID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Data retains all data created to use this service.`,
		Name:           "data",
		Stored:         true,
		SubType:        "service_data",
		Type:           "external",
	},
	"k8sidentifier": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "K8sIdentifier",
		Description:    `K8SIdentifier retains the identifier for kubernetes.`,
		Format:         "free",
		Name:           "k8sIdentifier",
		Stored:         true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		Description:    `Name of the installed app.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
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
		Filterable:     true,
		Format:         "free",
		Name:           "namespace",
		Orderable:      true,
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
		SubType:        "app_parameters",
		Type:           "external",
	},
	"relatedobjects": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RelatedObjects",
		Description:    `RelatedObjects retains all objects created to use this app.`,
		Name:           "relatedObjects",
		Stored:         true,
		SubType:        "app_relatedobjects",
		Type:           "external",
	},
	"replicas": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Replicas",
		Description:    `Replicas represents the number of replicas for the app.`,
		Exposed:        true,
		Name:           "replicas",
		Stored:         true,
		Type:           "integer",
	},
	"status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Error", "Pending", "Running"},
		ConvertedName:  "Status",
		DefaultValue:   InstalledAppStatusPending,
		Description:    `Status of the app.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
}
