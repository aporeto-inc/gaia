package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/v1/golang/types"
	"time"
)

// APIserviceTypeValue represents the possible values for attribute "type".
type APIserviceTypeValue string

const (
	// APIserviceTypeHttp represents the value HTTP.
	APIserviceTypeHttp APIserviceTypeValue = "HTTP"

	// APIserviceTypeL3 represents the value L3.
	APIserviceTypeL3 APIserviceTypeValue = "L3"

	// APIserviceTypeTcp represents the value TCP.
	APIserviceTypeTcp APIserviceTypeValue = "TCP"
)

// APIserviceIdentity represents the Identity of the object.
var APIserviceIdentity = elemental.Identity{
	Name:     "apiservice",
	Category: "apiservices",
	Private:  false,
}

// APIservicesList represents a list of APIservices
type APIservicesList []*APIservice

// ContentIdentity returns the identity of the objects in the list.
func (o APIservicesList) ContentIdentity() elemental.Identity {

	return APIserviceIdentity
}

// Copy returns a pointer to a copy the APIservicesList.
func (o APIservicesList) Copy() elemental.ContentIdentifiable {

	copy := append(APIservicesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the APIservicesList.
func (o APIservicesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(APIservicesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*APIservice))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o APIservicesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o APIservicesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o APIservicesList) Version() int {

	return 1
}

// APIservice represents the model of a apiservice
type APIservice struct {
	// ExposedAPIs is a list of API endpoints that are exposed for the service.
	ExposedAPIs []*types.ExposedAPIList `json:"exposedAPIs" bson:"exposedapis" mapstructure:"exposedAPIs,omitempty"`

	// External is a boolean that indicates if this is an external service.
	External bool `json:"external" bson:"external" mapstructure:"external,omitempty"`

	// Fqdn is the fully qualified domain name of the service. It is required for external API services. It can be deduced from a service discovery system in advanced environments.
	Fqdn string `json:"fqdn" bson:"fqdn" mapstructure:"fqdn,omitempty"`

	// ImplementedBy is a list of tag selectors that identifies that Processing Units that will implement this service.
	ImplementedBy [][]string `json:"implementedBy" bson:"implementedby" mapstructure:"implementedBy,omitempty"`

	// IPList is the list of ip address or subnets of the service if available.
	IpList []*types.IPList `json:"ipList" bson:"iplist" mapstructure:"ipList,omitempty"`

	// Port is the port of the service. Default 443.
	Port int `json:"port" bson:"port" mapstructure:"port,omitempty"`

	// Type is the type of the service (HTTP, TCP, etc). More types will be added to the system.
	Type APIserviceTypeValue `json:"type" bson:"type" mapstructure:"type,omitempty"`

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

// NewAPIservice returns a new *APIservice
func NewAPIservice() *APIservice {

	return &APIservice{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		ExposedAPIs:    []*types.ExposedAPIList{},
		External:       false,
		IpList:         []*types.IPList{},
		Metadata:       []string{},
		NormalizedTags: []string{},
		Port:           443,
		Type:           "L3",
	}
}

// Identity returns the Identity of the object.
func (o *APIservice) Identity() elemental.Identity {

	return APIserviceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *APIservice) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *APIservice) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *APIservice) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *APIservice) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *APIservice) Doc() string {
	return `APIService descibes a L4/L7 service and the corresponding implementation. It allows users to define their services, the APIs that they expose, the implementation of the service. These definitions can be used by network policy in order to define advanced controls based on the APIs.`
}

func (o *APIservice) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetArchived returns the Archived of the receiver.
func (o *APIservice) GetArchived() bool {

	return o.Archived
}

// SetArchived sets the given Archived of the receiver.
func (o *APIservice) SetArchived(archived bool) {

	o.Archived = archived
}

// GetAnnotations returns the Annotations of the receiver.
func (o *APIservice) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *APIservice) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *APIservice) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *APIservice) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *APIservice) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *APIservice) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *APIservice) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *APIservice) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *APIservice) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *APIservice) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *APIservice) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *APIservice) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *APIservice) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *APIservice) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *APIservice) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *APIservice) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *APIservice) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *APIservice) Validate() error {

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
func (*APIservice) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := APIserviceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return APIserviceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*APIservice) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return APIserviceAttributesMap
}

// APIserviceAttributesMap represents the map of attribute for APIservice.
var APIserviceAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Fqdn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fqdn",
		Description:    `Fqdn is the fully qualified domain name of the service. It is required for external API services. It can be deduced from a service discovery system in advanced environments.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "fqdn",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
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
	"IpList": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IpList",
		Description:    `IPList is the list of ip address or subnets of the service if available.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "ipList",
		Orderable:      true,
		Stored:         true,
		SubType:        "ip_list",
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
		DefaultValue:   APIserviceTypeL3,
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

// APIserviceLowerCaseAttributesMap represents the map of attribute for APIservice.
var APIserviceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"fqdn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fqdn",
		Description:    `Fqdn is the fully qualified domain name of the service. It is required for external API services. It can be deduced from a service discovery system in advanced environments.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "fqdn",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
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
	"iplist": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IpList",
		Description:    `IPList is the list of ip address or subnets of the service if available.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "ipList",
		Orderable:      true,
		Stored:         true,
		SubType:        "ip_list",
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
		DefaultValue:   APIserviceTypeL3,
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
