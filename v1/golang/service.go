package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/v1/golang/types"
	"time"
)

// ServiceTypeValue represents the possible values for attribute "type".
type ServiceTypeValue string

const (
	// ServiceTypeHTTP represents the value HTTP.
	ServiceTypeHTTP ServiceTypeValue = "HTTP"

	// ServiceTypeTCP represents the value TCP.
	ServiceTypeTCP ServiceTypeValue = "TCP"
)

// ServiceIdentity represents the Identity of the object.
var ServiceIdentity = elemental.Identity{
	Name:     "service",
	Category: "services",
	Private:  false,
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

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o ServicesList) Version() int {

	return 1
}

// Service represents the model of a service
type Service struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// IPs is the list of IP addresses where the service can be accessed.
	// This is an optional attribute and is only required if no host names are
	// provided.
	// The system will automatically resolve IP addresses from  host names otherwise.
	IPs types.IPList `json:"IPs" bson:"ips" mapstructure:"IPs,omitempty"`

	// JWTSigningCertificate is a certificate that can be used to validate user JWT in
	// HTTP requests. This is an optional field, needed only if user JWT validation is
	// required for this service. The certificate must be in PEM format.
	JWTSigningCertificate string `json:"JWTSigningCertificate" bson:"jwtsigningcertificate" mapstructure:"JWTSigningCertificate,omitempty"`

	// This is a set of all API tags for matching in the DB.
	AllAPITags []string `json:"-" bson:"allapitags" mapstructure:"-,omitempty"`

	// This is a set of all selector tags for matching in the DB.
	AllServiceTags []string `json:"-" bson:"allservicetags" mapstructure:"-,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Archived defines if the object is archived.
	Archived bool `json:"-" bson:"archived" mapstructure:"-,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Endpoints is a read only attribute that actually resolves the API
	// endpoints that the service is exposing. Only valid during policy rendering.
	Endpoints types.ExposedAPIList `json:"endpoints" bson:"-" mapstructure:"endpoints,omitempty"`

	// ExposedAPIs contains a tag expression that will determine which
	// APIs a servie is exposing. The APIs can be defined as the RESTAPISpec or
	// similar specifications for other L7 protocols.
	ExposedAPIs [][]string `json:"exposedAPIs" bson:"exposedapis" mapstructure:"exposedAPIs,omitempty"`

	// ExposedPort is the port that the service can be accessed. Note that
	// this is different from the Port attribute that describes the port that the
	// service is actually listening. For example if a load balancer is used, the
	// ExposedPort is the port that the load balancer is listening for the service,
	// whereas the port that the implementation is listening can be different.
	ExposedPort int `json:"exposedPort" bson:"exposedport" mapstructure:"exposedPort,omitempty"`

	// External is a boolean that indicates if this is an external service.
	External bool `json:"external" bson:"external" mapstructure:"external,omitempty"`

	// Hosts are the names that the service can be accessed with.
	Hosts []string `json:"hosts" bson:"hosts" mapstructure:"hosts,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Port is the port that the implementation of the service is listening to and
	// it can be different than the exposedPorts describing the service. This is needed
	// for port mapping use cases where there is private and public ports.
	Port int `json:"port" bson:"port" mapstructure:"port,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Selectors contains the tag expression that an a processing unit
	// must match in order to implement this particular service.
	Selectors [][]string `json:"selectors" bson:"selectors" mapstructure:"selectors,omitempty"`

	// ServiceCA  is the certificate authority that the service is using. This
	// is needed for external services with private certificate authorities. The
	// field is optional. If provided, this must be a valid PEM CA file.
	ServiceCA string `json:"serviceCA" bson:"serviceca" mapstructure:"serviceCA,omitempty"`

	// Type is the type of the service.
	Type ServiceTypeValue `json:"type" bson:"type" mapstructure:"type,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewService returns a new *Service
func NewService() *Service {

	return &Service{
		ModelVersion:   1,
		AllAPITags:     []string{},
		AllServiceTags: []string{},
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Endpoints:      types.ExposedAPIList{},
		External:       false,
		IPs:            types.IPList{},
		NormalizedTags: []string{},
		Type:           "HTTP",
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

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Service) Doc() string {
	return `A Service defines a generic service object at L4 or L7 that encapsulates
the description of a micro-service. A service exposes APIs and can be
implemented through third party entities (such as a cloud provider) or through 
processign units.`
}

func (o *Service) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Service) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *Service) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetArchived returns the Archived of the receiver.
func (o *Service) GetArchived() bool {

	return o.Archived
}

// SetArchived sets the given Archived of the receiver.
func (o *Service) SetArchived(archived bool) {

	o.Archived = archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Service) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *Service) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Service) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *Service) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetName returns the Name of the receiver.
func (o *Service) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *Service) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Service) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *Service) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Service) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *Service) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Service) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Service) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *Service) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *Service) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("exposedAPIs", o.ExposedAPIs); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("exposedPort", o.ExposedPort, int(65535), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("exposedPort", o.ExposedPort, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("port", o.Port, int(65535), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("port", o.Port, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("selectors", o.Selectors); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"HTTP", "TCP"}, false); err != nil {
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
	},
	"IPs": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IPs",
		Description: `IPs is the list of IP addresses where the service can be accessed.
This is an optional attribute and is only required if no host names are
provided.
The system will automatically resolve IP addresses from  host names otherwise.`,
		Exposed: true,
		Name:    "IPs",
		Stored:  true,
		SubType: "ip_list",
		Type:    "external",
	},
	"JWTSigningCertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "JWTSigningCertificate",
		Description: `JWTSigningCertificate is a certificate that can be used to validate user JWT in
HTTP requests. This is an optional field, needed only if user JWT validation is
required for this service. The certificate must be in PEM format.`,
		Exposed: true,
		Format:  "free",
		Name:    "JWTSigningCertificate",
		Stored:  true,
		Type:    "string",
	},
	"AllAPITags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllAPITags",
		Description:    `This is a set of all API tags for matching in the DB.`,
		Name:           "allAPITags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"AllServiceTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllServiceTags",
		Description:    `This is a set of all selector tags for matching in the DB.`,
		Name:           "allServiceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
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
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
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
		Format:         "free",
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Endpoints": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Endpoints",
		Description: `Endpoints is a read only attribute that actually resolves the API
endpoints that the service is exposing. Only valid during policy rendering.`,
		Exposed:  true,
		Name:     "endpoints",
		ReadOnly: true,
		SubType:  "exposed_api_list",
		Type:     "external",
	},
	"ExposedAPIs": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedAPIs",
		Description: `ExposedAPIs contains a tag expression that will determine which
APIs a servie is exposing. The APIs can be defined as the RESTAPISpec or
similar specifications for other L7 protocols.`,
		Exposed:  true,
		Name:     "exposedAPIs",
		Required: true,
		Stored:   true,
		SubType:  "policies_list",
		Type:     "external",
	},
	"ExposedPort": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedPort",
		Description: `ExposedPort is the port that the service can be accessed. Note that
this is different from the Port attribute that describes the port that the
service is actually listening. For example if a load balancer is used, the
ExposedPort is the port that the load balancer is listening for the service,
whereas the port that the implementation is listening can be different.`,
		Exposed:  true,
		MaxValue: 65535,
		MinValue: 1,
		Name:     "exposedPort",
		Required: true,
		Stored:   true,
		Type:     "integer",
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
	"Hosts": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Hosts",
		Description:    `Hosts are the names that the service can be accessed with.`,
		Exposed:        true,
		Format:         "free",
		Name:           "hosts",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
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
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
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
		Description: `Port is the port that the implementation of the service is listening to and
it can be different than the exposedPorts describing the service. This is needed
for port mapping use cases where there is private and public ports.`,
		Exposed:  true,
		MaxValue: 65535,
		MinValue: 1,
		Name:     "port",
		Required: true,
		Stored:   true,
		Type:     "integer",
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
	"Selectors": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Selectors",
		Description: `Selectors contains the tag expression that an a processing unit
must match in order to implement this particular service.`,
		Exposed:  true,
		Name:     "selectors",
		Required: true,
		Stored:   true,
		SubType:  "policies_list",
		Type:     "external",
	},
	"ServiceCA": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceCA",
		Description: `ServiceCA  is the certificate authority that the service is using. This
is needed for external services with private certificate authorities. The
field is optional. If provided, this must be a valid PEM CA file.`,
		Exposed: true,
		Format:  "free",
		Name:    "serviceCA",
		Stored:  true,
		Type:    "string",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"HTTP", "TCP"},
		ConvertedName:  "Type",
		DefaultValue:   ServiceTypeHTTP,
		Description:    `Type is the type of the service.`,
		Exposed:        true,
		Name:           "type",
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

// ServiceLowerCaseAttributesMap represents the map of attribute for Service.
var ServiceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	},
	"ips": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IPs",
		Description: `IPs is the list of IP addresses where the service can be accessed.
This is an optional attribute and is only required if no host names are
provided.
The system will automatically resolve IP addresses from  host names otherwise.`,
		Exposed: true,
		Name:    "IPs",
		Stored:  true,
		SubType: "ip_list",
		Type:    "external",
	},
	"jwtsigningcertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "JWTSigningCertificate",
		Description: `JWTSigningCertificate is a certificate that can be used to validate user JWT in
HTTP requests. This is an optional field, needed only if user JWT validation is
required for this service. The certificate must be in PEM format.`,
		Exposed: true,
		Format:  "free",
		Name:    "JWTSigningCertificate",
		Stored:  true,
		Type:    "string",
	},
	"allapitags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllAPITags",
		Description:    `This is a set of all API tags for matching in the DB.`,
		Name:           "allAPITags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"allservicetags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllServiceTags",
		Description:    `This is a set of all selector tags for matching in the DB.`,
		Name:           "allServiceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
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
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
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
		Format:         "free",
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"endpoints": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Endpoints",
		Description: `Endpoints is a read only attribute that actually resolves the API
endpoints that the service is exposing. Only valid during policy rendering.`,
		Exposed:  true,
		Name:     "endpoints",
		ReadOnly: true,
		SubType:  "exposed_api_list",
		Type:     "external",
	},
	"exposedapis": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedAPIs",
		Description: `ExposedAPIs contains a tag expression that will determine which
APIs a servie is exposing. The APIs can be defined as the RESTAPISpec or
similar specifications for other L7 protocols.`,
		Exposed:  true,
		Name:     "exposedAPIs",
		Required: true,
		Stored:   true,
		SubType:  "policies_list",
		Type:     "external",
	},
	"exposedport": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedPort",
		Description: `ExposedPort is the port that the service can be accessed. Note that
this is different from the Port attribute that describes the port that the
service is actually listening. For example if a load balancer is used, the
ExposedPort is the port that the load balancer is listening for the service,
whereas the port that the implementation is listening can be different.`,
		Exposed:  true,
		MaxValue: 65535,
		MinValue: 1,
		Name:     "exposedPort",
		Required: true,
		Stored:   true,
		Type:     "integer",
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
	"hosts": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Hosts",
		Description:    `Hosts are the names that the service can be accessed with.`,
		Exposed:        true,
		Format:         "free",
		Name:           "hosts",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
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
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
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
		Description: `Port is the port that the implementation of the service is listening to and
it can be different than the exposedPorts describing the service. This is needed
for port mapping use cases where there is private and public ports.`,
		Exposed:  true,
		MaxValue: 65535,
		MinValue: 1,
		Name:     "port",
		Required: true,
		Stored:   true,
		Type:     "integer",
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
	"selectors": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Selectors",
		Description: `Selectors contains the tag expression that an a processing unit
must match in order to implement this particular service.`,
		Exposed:  true,
		Name:     "selectors",
		Required: true,
		Stored:   true,
		SubType:  "policies_list",
		Type:     "external",
	},
	"serviceca": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceCA",
		Description: `ServiceCA  is the certificate authority that the service is using. This
is needed for external services with private certificate authorities. The
field is optional. If provided, this must be a valid PEM CA file.`,
		Exposed: true,
		Format:  "free",
		Name:    "serviceCA",
		Stored:  true,
		Type:    "string",
	},
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{"HTTP", "TCP"},
		ConvertedName:  "Type",
		DefaultValue:   ServiceTypeHTTP,
		Description:    `Type is the type of the service.`,
		Exposed:        true,
		Name:           "type",
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
