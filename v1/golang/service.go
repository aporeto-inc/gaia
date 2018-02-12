package models

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/highwindmodels/v1/golang/types"
)

// ServiceStatusValue represents the possible values for attribute "status".
type ServiceStatusValue string

const (
	// ServiceStatusError represents the value Error.
	ServiceStatusError ServiceStatusValue = "Error"

	// ServiceStatusPending represents the value Pending.
	ServiceStatusPending ServiceStatusValue = "Pending"

	// ServiceStatusRunning represents the value Running.
	ServiceStatusRunning ServiceStatusValue = "Running"
)

// ServiceIdentity represents the Identity of the object.
var ServiceIdentity = elemental.Identity{
	Name:     "service",
	Category: "services",
}

// ServicesList represents a list of Services
type ServicesList []*Service

// ContentIdentity returns the identity of the objects in the list.
func (o ServicesList) ContentIdentity() elemental.Identity {

	return ServiceIdentity
}

// Copy returns a pointer to a copy the ServicesList.
func (o ServicesList) Copy() elemental.ContentIdentifiable {

	copy := append(ServicesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ServicesList.
func (o ServicesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(ServicesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Service))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ServicesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ServicesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o ServicesList) Version() int {

	return 1
}

// Service represents the model of a service
type Service struct {
	// ID of the service
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// AccountName represents the vince account name
	AccountName string `json:"accountName" bson:"accountname" mapstructure:"accountName,omitempty"`

	// CategoryID of the service.
	CategoryID string `json:"categoryID" bson:"categoryid" mapstructure:"categoryID,omitempty"`

	// K8SIdentifier retains the identifier for kubernetes.
	K8sIdentifier string `json:"-" bson:"k8sidentifier" mapstructure:"-,omitempty"`

	// Name of the service
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace in which the service in running.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Parameters is a list of parameters to start the service
	Parameters []*types.ServiceParameter `json:"parameters" bson:"parameters" mapstructure:"parameters,omitempty"`

	// RelatedObjects retains all objects created to use this service.
	RelatedObjects []*types.ServiceRelatedObject `json:"-" bson:"relatedobjects" mapstructure:"-,omitempty"`

	// Replicas represents the number of replicas for the service.
	Replicas int `json:"replicas" bson:"replicas" mapstructure:"replicas,omitempty"`

	// Status of the service.
	Status ServiceStatusValue `json:"status" bson:"status" mapstructure:"status,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewService returns a new *Service
func NewService() *Service {

	return &Service{
		ModelVersion:   1,
		Parameters:     []*types.ServiceParameter{},
		RelatedObjects: []*types.ServiceRelatedObject{},
		Status:         "Pending",
	}
}

// Identity returns the Identity of the object.
func (o *Service) Identity() elemental.Identity {

	return ServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Service) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Service) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Service) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Service) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Service) Doc() string {
	return `Service represents a service that can be launched`
}

func (o *Service) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Service) Validate() error {

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
func (*Service) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ServiceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ServiceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Service) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ServiceAttributesMap
}

// ServiceAttributesMap represents the map of attribute for Service.
var ServiceAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID of the service`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"AccountName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccountName",
		CreationOnly:   true,
		Description:    `AccountName represents the vince account name`,
		Exposed:        true,
		Format:         "free",
		Name:           "accountName",
		Stored:         true,
		Type:           "string",
	},
	"CategoryID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CategoryID",
		Description:    `CategoryID of the service.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "categoryID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
		Description:    `Name of the service`,
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
		Description:    `Namespace in which the service in running.`,
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
		Description:    `Parameters is a list of parameters to start the service`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "service_parameters",
		Type:           "external",
	},
	"RelatedObjects": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RelatedObjects",
		Description:    `RelatedObjects retains all objects created to use this service.`,
		Name:           "relatedObjects",
		Stored:         true,
		SubType:        "service_relatedobjects",
		Type:           "external",
	},
	"Replicas": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Replicas",
		Description:    `Replicas represents the number of replicas for the service.`,
		Exposed:        true,
		Name:           "replicas",
		Stored:         true,
		Type:           "integer",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Error", "Pending", "Running"},
		ConvertedName:  "Status",
		DefaultValue:   ServiceStatusPending,
		Description:    `Status of the service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
}

// ServiceLowerCaseAttributesMap represents the map of attribute for Service.
var ServiceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID of the service`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"accountname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccountName",
		CreationOnly:   true,
		Description:    `AccountName represents the vince account name`,
		Exposed:        true,
		Format:         "free",
		Name:           "accountName",
		Stored:         true,
		Type:           "string",
	},
	"categoryid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CategoryID",
		Description:    `CategoryID of the service.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "categoryID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
		Description:    `Name of the service`,
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
		Description:    `Namespace in which the service in running.`,
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
		Description:    `Parameters is a list of parameters to start the service`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "service_parameters",
		Type:           "external",
	},
	"relatedobjects": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RelatedObjects",
		Description:    `RelatedObjects retains all objects created to use this service.`,
		Name:           "relatedObjects",
		Stored:         true,
		SubType:        "service_relatedobjects",
		Type:           "external",
	},
	"replicas": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Replicas",
		Description:    `Replicas represents the number of replicas for the service.`,
		Exposed:        true,
		Name:           "replicas",
		Stored:         true,
		Type:           "integer",
	},
	"status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Error", "Pending", "Running"},
		ConvertedName:  "Status",
		DefaultValue:   ServiceStatusPending,
		Description:    `Status of the service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
}
