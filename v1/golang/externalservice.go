package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/v1/golang/types"
	"time"
)

// ExternalServiceTypeValue represents the possible values for attribute "type".
type ExternalServiceTypeValue string

const (
	// ExternalServiceTypeLoadbalancerhttp represents the value LoadBalancerHTTP.
	ExternalServiceTypeLoadbalancerhttp ExternalServiceTypeValue = "LoadBalancerHTTP"

	// ExternalServiceTypeLoadbalancertcp represents the value LoadBalancerTCP.
	ExternalServiceTypeLoadbalancertcp ExternalServiceTypeValue = "LoadBalancerTCP"

	// ExternalServiceTypeNetwork represents the value Network.
	ExternalServiceTypeNetwork ExternalServiceTypeValue = "Network"
)

// ExternalServiceIdentity represents the Identity of the object.
var ExternalServiceIdentity = elemental.Identity{
	Name:     "externalservice",
	Category: "externalservices",
	Private:  false,
}

// ExternalServicesList represents a list of ExternalServices
type ExternalServicesList []*ExternalService

// ContentIdentity returns the identity of the objects in the list.
func (o ExternalServicesList) ContentIdentity() elemental.Identity {

	return ExternalServiceIdentity
}

// Copy returns a pointer to a copy the ExternalServicesList.
func (o ExternalServicesList) Copy() elemental.ContentIdentifiable {

	copy := append(ExternalServicesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ExternalServicesList.
func (o ExternalServicesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(ExternalServicesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ExternalService))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ExternalServicesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ExternalServicesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o ExternalServicesList) Version() int {

	return 1
}

// ExternalService represents the model of a externalservice
type ExternalService struct {
	// LoadbalancerAddresses represents the list of adresses of the external services of type LoadBalancer.
	LoadbalancerAddresses []string `json:"loadbalancerAddresses" bson:"loadbalanceraddresses" mapstructure:"loadbalancerAddresses,omitempty"`

	// LoadbalancerPortsMapping is the list of ports mapped by an extenral service of type load balancer.
	LoadbalancerPortsMapping []*types.PortMapping `json:"loadbalancerPortsMapping" bson:"loadbalancerportsmapping" mapstructure:"loadbalancerPortsMapping,omitempty"`

	// Network refers to either CIDR or domain name
	Network string `json:"network" bson:"network" mapstructure:"network,omitempty"`

	// Port refers to network port which could be a single number or 100:2000 to represent a range of ports
	Port string `json:"port" bson:"port" mapstructure:"port,omitempty"`

	// Protocol refers to network protocol like TCP/UDP or the number of the protocol.
	Protocol string `json:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`

	// Type represents the type of external service.
	Type ExternalServiceTypeValue `json:"type" bson:"type" mapstructure:"type,omitempty"`

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

// NewExternalService returns a new *ExternalService
func NewExternalService() *ExternalService {

	return &ExternalService{
		ModelVersion:             1,
		Annotations:              map[string][]string{},
		AssociatedTags:           []string{},
		LoadbalancerAddresses:    []string{},
		LoadbalancerPortsMapping: []*types.PortMapping{},
		Metadata:                 []string{},
		NormalizedTags:           []string{},
		Port:                     "1:65535",
		Type:                     "Network",
	}
}

// Identity returns the Identity of the object.
func (o *ExternalService) Identity() elemental.Identity {

	return ExternalServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ExternalService) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ExternalService) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *ExternalService) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *ExternalService) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *ExternalService) Doc() string {
	return `An External Service represents a random network or ip that is not managed by the system. They can be used in Network Access Policies in order to allow traffic from or to the declared network or IP, using the provided protocol and port or ports range. If you want to describe the Internet (ie. anywhere), use 0.0.0.0/0 as address, and 1-65000 for the ports. You will need to use the External Services tags to set some policies.`
}

func (o *ExternalService) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetArchived returns the Archived of the receiver.
func (o *ExternalService) GetArchived() bool {

	return o.Archived
}

// SetArchived sets the given Archived of the receiver.
func (o *ExternalService) SetArchived(archived bool) {

	o.Archived = archived
}

// GetAnnotations returns the Annotations of the receiver.
func (o *ExternalService) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *ExternalService) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *ExternalService) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *ExternalService) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *ExternalService) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *ExternalService) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *ExternalService) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *ExternalService) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *ExternalService) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *ExternalService) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *ExternalService) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *ExternalService) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *ExternalService) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *ExternalService) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *ExternalService) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *ExternalService) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *ExternalService) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *ExternalService) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("network", o.Network); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("network", o.Network); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("port", o.Port, `^([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535)(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))?$`, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("protocol", o.Protocol); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidatePattern("protocol", o.Protocol, `^(TCP|UDP|tcp|udp|[1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`, true); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("protocol", o.Protocol); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"LoadBalancerHTTP", "LoadBalancerTCP", "Network"}, false); err != nil {
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
func (*ExternalService) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ExternalServiceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ExternalServiceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ExternalService) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ExternalServiceAttributesMap
}

// ExternalServiceAttributesMap represents the map of attribute for ExternalService.
var ExternalServiceAttributesMap = map[string]elemental.AttributeSpecification{
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
	"LoadbalancerAddresses": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LoadbalancerAddresses",
		Description:    `LoadbalancerAddresses represents the list of adresses of the external services of type LoadBalancer.`,
		Exposed:        true,
		Name:           "loadbalancerAddresses",
		Stored:         true,
		SubType:        "addresses_list",
		Type:           "external",
	},
	"LoadbalancerPortsMapping": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LoadbalancerPortsMapping",
		Description:    `LoadbalancerPortsMapping is the list of ports mapped by an extenral service of type load balancer. `,
		Exposed:        true,
		Name:           "loadbalancerPortsMapping",
		Stored:         true,
		SubType:        "portmapping_list",
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
	"Network": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Network",
		Description:    `Network refers to either CIDR or domain name`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "network",
		Required:       true,
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
		AllowedChars:   `^([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535)(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))?$`,
		AllowedChoices: []string{},
		ConvertedName:  "Port",
		DefaultValue:   "1:65535",
		Description:    `Port refers to network port which could be a single number or 100:2000 to represent a range of ports`,
		Exposed:        true,
		Filterable:     true,
		Name:           "port",
		Stored:         true,
		Type:           "string",
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
	"Protocol": elemental.AttributeSpecification{
		AllowedChars:   `^(TCP|UDP|tcp|udp|[1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`,
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol refers to network protocol like TCP/UDP or the number of the protocol.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "protocol",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"LoadBalancerHTTP", "LoadBalancerTCP", "Network"},
		ConvertedName:  "Type",
		DefaultValue:   ExternalServiceTypeNetwork,
		Description:    `Type represents the type of external service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Orderable:      true,
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

// ExternalServiceLowerCaseAttributesMap represents the map of attribute for ExternalService.
var ExternalServiceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"loadbalanceraddresses": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LoadbalancerAddresses",
		Description:    `LoadbalancerAddresses represents the list of adresses of the external services of type LoadBalancer.`,
		Exposed:        true,
		Name:           "loadbalancerAddresses",
		Stored:         true,
		SubType:        "addresses_list",
		Type:           "external",
	},
	"loadbalancerportsmapping": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LoadbalancerPortsMapping",
		Description:    `LoadbalancerPortsMapping is the list of ports mapped by an extenral service of type load balancer. `,
		Exposed:        true,
		Name:           "loadbalancerPortsMapping",
		Stored:         true,
		SubType:        "portmapping_list",
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
	"network": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Network",
		Description:    `Network refers to either CIDR or domain name`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "network",
		Required:       true,
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
		AllowedChars:   `^([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535)(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))?$`,
		AllowedChoices: []string{},
		ConvertedName:  "Port",
		DefaultValue:   "1:65535",
		Description:    `Port refers to network port which could be a single number or 100:2000 to represent a range of ports`,
		Exposed:        true,
		Filterable:     true,
		Name:           "port",
		Stored:         true,
		Type:           "string",
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
	"protocol": elemental.AttributeSpecification{
		AllowedChars:   `^(TCP|UDP|tcp|udp|[1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`,
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol refers to network protocol like TCP/UDP or the number of the protocol.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "protocol",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{"LoadBalancerHTTP", "LoadBalancerTCP", "Network"},
		ConvertedName:  "Type",
		DefaultValue:   ExternalServiceTypeNetwork,
		Description:    `Type represents the type of external service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Orderable:      true,
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
