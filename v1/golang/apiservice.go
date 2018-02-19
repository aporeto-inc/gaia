package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/v1/golang/types"
	"time"
)

// APIServiceTypeValue represents the possible values for attribute "type".
type APIServiceTypeValue string

const (
	// APIServiceTypeHttp represents the value HTTP.
	APIServiceTypeHttp APIServiceTypeValue = "HTTP"

	// APIServiceTypeL3 represents the value L3.
	APIServiceTypeL3 APIServiceTypeValue = "L3"

	// APIServiceTypeTcp represents the value TCP.
	APIServiceTypeTcp APIServiceTypeValue = "TCP"
)

// APIServiceIdentity represents the Identity of the object.
var APIServiceIdentity = elemental.Identity{
	Name:     "apiservice",
	Category: "apiservices",
	Private:  false,
}

// APIServicesList represents a list of APIServices
type APIServicesList []*APIService

// ContentIdentity returns the identity of the objects in the list.
func (o APIServicesList) ContentIdentity() elemental.Identity {

	return APIServiceIdentity
}

// Copy returns a pointer to a copy the APIServicesList.
func (o APIServicesList) Copy() elemental.ContentIdentifiable {

	copy := append(APIServicesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the APIServicesList.
func (o APIServicesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(APIServicesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*APIService))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o APIServicesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o APIServicesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o APIServicesList) Version() int {

	return 1
}

// APIService represents the model of a apiservice
type APIService struct {
	// FQDN is the fully qualified domain name of the service. It is required for external API services. It can be deduced from a service discovery system in advanced environments.
	FQDN string `json:"FQDN" bson:"fqdn" mapstructure:"FQDN,omitempty"`

	// IPList is the list of ip address or subnets of the service if available.
	IPList []*types.IPList `json:"IPList" bson:"iplist" mapstructure:"IPList,omitempty"`

	// AllServiceTags is an internal object that summarizes all the implementedBy tags to accelerate database searches. It is not exposed.
	AllServiceTags []string `json:"-" bson:"allservicetags" mapstructure:"-,omitempty"`

	// ExposedAPIs is a list of API endpoints that are exposed for the service.
	ExposedAPIs []*types.ExposedAPIList `json:"exposedAPIs" bson:"exposedapis" mapstructure:"exposedAPIs,omitempty"`

	// External is a boolean that indicates if this is an external service.
	External bool `json:"external" bson:"external" mapstructure:"external,omitempty"`

	// ImplementedBy is a list of tag selectors that identifies that Processing Units that will implement this service.
	ImplementedBy [][]string `json:"implementedBy" bson:"implementedby" mapstructure:"implementedBy,omitempty"`

	// Port is the port of the service. Default 443.
	Port int `json:"port" bson:"port" mapstructure:"port,omitempty"`

	// Type is the type of the service (HTTP, TCP, etc). More types will be added to the system.
	Type APIServiceTypeValue `json:"type" bson:"type" mapstructure:"type,omitempty"`

	// Archived defines if the object is archived.
	Archived bool `json:"-" bson:"archived" mapstructure:"-,omitempty"`

	// Annotation stores additional information about an entity
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewAPIService returns a new *APIService
func NewAPIService() *APIService {

	return &APIService{
		ModelVersion:   1,
		AllServiceTags: []string{},
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		ExposedAPIs:    []*types.ExposedAPIList{},
		External:       false,
		IPList:         []*types.IPList{},
		Metadata:       []string{},
		NormalizedTags: []string{},
		Port:           443,
		Type:           "L3",
	}
}

// Identity returns the Identity of the object.
func (o *APIService) Identity() elemental.Identity {

	return APIServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *APIService) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *APIService) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *APIService) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *APIService) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *APIService) Doc() string {
	return `APIService descibes a L4/L7 service and the corresponding implementation. It allows users to define their services, the APIs that they expose, the implementation of the service. These definitions can be used by network policy in order to define advanced controls based on the APIs.`
}

func (o *APIService) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetArchived returns the Archived of the receiver.
func (o *APIService) GetArchived() bool {

	return o.Archived
}

// SetArchived sets the given Archived of the receiver.
func (o *APIService) SetArchived(archived bool) {

	o.Archived = archived
}

// GetAnnotations returns the Annotations of the receiver.
func (o *APIService) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *APIService) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *APIService) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *APIService) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *APIService) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *APIService) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *APIService) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *APIService) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *APIService) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *APIService) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *APIService) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *APIService) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *APIService) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *APIService) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *APIService) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *APIService) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *APIService) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *APIService) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("implementedBy", o.ImplementedBy); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("implementedBy", o.ImplementedBy); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("port", o.Port, int(65636), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("port", o.Port, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"HTTP", "L3", "TCP"}, false); err != nil {
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
func (*APIService) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := APIServiceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return APIServiceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*APIService) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return APIServiceAttributesMap
}

// APIServiceAttributesMap represents the map of attribute for APIService.
var APIServiceAttributesMap = map[string]elemental.AttributeSpecification{
	"FQDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FQDN",
		Description:    `FQDN is the fully qualified domain name of the service. It is required for external API services. It can be deduced from a service discovery system in advanced environments.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "FQDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
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
	"IPList": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IPList",
		Description:    `IPList is the list of ip address or subnets of the service if available.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "IPList",
		Orderable:      true,
		Stored:         true,
		SubType:        "ip_list",
		Type:           "external",
	},
	"AllServiceTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllServiceTags",
		Description:    `AllServiceTags is an internal object that summarizes all the implementedBy tags to accelerate database searches. It is not exposed.`,
		Name:           "allServiceTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
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
	"Archived": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Archived",
		Description:    `Archived defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"ExposedAPIs": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedAPIs",
		Description:    `ExposedAPIs is a list of API endpoints that are exposed for the service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "exposedAPIs",
		Orderable:      true,
		Stored:         true,
		SubType:        "exposed_api_list",
		Type:           "external",
	},
	"External": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "External",
		DefaultValue:   false,
		Description:    `External is a boolean that indicates if this is an external service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "external",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"ImplementedBy": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ImplementedBy",
		Description:    `ImplementedBy is a list of tag selectors that identifies that Processing Units that will implement this service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "implementedBy",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		SubType:        "target_tags",
		Type:           "external",
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
	"Port": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Port",
		DefaultValue:   443,
		Description:    `Port is the port of the service. Default 443.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       65636,
		MinValue:       1,
		Name:           "port",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "integer",
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
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"HTTP", "L3", "TCP"},
		ConvertedName:  "Type",
		DefaultValue:   APIServiceTypeL3,
		Description:    `Type is the type of the service (HTTP, TCP, etc). More types will be added to the system.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "enum",
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

// APIServiceLowerCaseAttributesMap represents the map of attribute for APIService.
var APIServiceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"fqdn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FQDN",
		Description:    `FQDN is the fully qualified domain name of the service. It is required for external API services. It can be deduced from a service discovery system in advanced environments.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "FQDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
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
	"iplist": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IPList",
		Description:    `IPList is the list of ip address or subnets of the service if available.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "IPList",
		Orderable:      true,
		Stored:         true,
		SubType:        "ip_list",
		Type:           "external",
	},
	"allservicetags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllServiceTags",
		Description:    `AllServiceTags is an internal object that summarizes all the implementedBy tags to accelerate database searches. It is not exposed.`,
		Name:           "allServiceTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
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
	"archived": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Archived",
		Description:    `Archived defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"exposedapis": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedAPIs",
		Description:    `ExposedAPIs is a list of API endpoints that are exposed for the service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "exposedAPIs",
		Orderable:      true,
		Stored:         true,
		SubType:        "exposed_api_list",
		Type:           "external",
	},
	"external": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "External",
		DefaultValue:   false,
		Description:    `External is a boolean that indicates if this is an external service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "external",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"implementedby": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ImplementedBy",
		Description:    `ImplementedBy is a list of tag selectors that identifies that Processing Units that will implement this service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "implementedBy",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		SubType:        "target_tags",
		Type:           "external",
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
	"port": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Port",
		DefaultValue:   443,
		Description:    `Port is the port of the service. Default 443.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       65636,
		MinValue:       1,
		Name:           "port",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "integer",
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
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{"HTTP", "L3", "TCP"},
		ConvertedName:  "Type",
		DefaultValue:   APIServiceTypeL3,
		Description:    `Type is the type of the service (HTTP, TCP, etc). More types will be added to the system.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "enum",
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
