package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// ExternalNetworkIndexes lists the attribute compound indexes.
var ExternalNetworkIndexes = [][]string{}

// ExternalNetworkIdentity represents the Identity of the object.
var ExternalNetworkIdentity = elemental.Identity{
	Name:     "externalnetwork",
	Category: "externalnetworks",
	Private:  false,
}

// ExternalNetworksList represents a list of ExternalNetworks
type ExternalNetworksList []*ExternalNetwork

// Identity returns the identity of the objects in the list.
func (o ExternalNetworksList) Identity() elemental.Identity {

	return ExternalNetworkIdentity
}

// Copy returns a pointer to a copy the ExternalNetworksList.
func (o ExternalNetworksList) Copy() elemental.Identifiables {

	copy := append(ExternalNetworksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ExternalNetworksList.
func (o ExternalNetworksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ExternalNetworksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ExternalNetwork))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ExternalNetworksList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ExternalNetworksList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o ExternalNetworksList) Version() int {

	return 1
}

// ExternalNetwork represents the model of a externalnetwork
type ExternalNetwork struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

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

	// List of CIDRs or domain name.
	Entries []string `json:"entries" bson:"entries" mapstructure:"entries,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// List of single ports or range (xx:yy).
	Ports []string `json:"ports" bson:"ports" mapstructure:"ports,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// List of protocols (tcp, udp, or protocol number).
	Protocols []string `json:"protocols" bson:"protocols" mapstructure:"protocols,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewExternalNetwork returns a new *ExternalNetwork
func NewExternalNetwork() *ExternalNetwork {

	return &ExternalNetwork{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Metadata:       []string{},
		NormalizedTags: []string{},
		Ports: []string{
			"1:65535",
		},
		Protocols: []string{
			"tcp",
		},
	}
}

// Identity returns the Identity of the object.
func (o *ExternalNetwork) Identity() elemental.Identity {

	return ExternalNetworkIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ExternalNetwork) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ExternalNetwork) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *ExternalNetwork) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *ExternalNetwork) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *ExternalNetwork) Doc() string {
	return `An External Network represents a random network or ip that is not managed by the
system. They can be used in Network Access Policies in order to allow traffic
from or to the declared network or IP, using the provided protocol and port or
ports range. If you want to describe the Internet (ie. anywhere), use 0.0.0.0/0
as address, and 1-65000 for the ports. You will need to use the External
Services tags to set some policies.`
}

func (o *ExternalNetwork) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *ExternalNetwork) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *ExternalNetwork) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetArchived returns the Archived of the receiver.
func (o *ExternalNetwork) GetArchived() bool {

	return o.Archived
}

// SetArchived sets the given Archived of the receiver.
func (o *ExternalNetwork) SetArchived(archived bool) {

	o.Archived = archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *ExternalNetwork) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *ExternalNetwork) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *ExternalNetwork) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *ExternalNetwork) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *ExternalNetwork) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *ExternalNetwork) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *ExternalNetwork) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *ExternalNetwork) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *ExternalNetwork) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *ExternalNetwork) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *ExternalNetwork) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *ExternalNetwork) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *ExternalNetwork) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *ExternalNetwork) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *ExternalNetwork) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *ExternalNetwork) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := ValidateNetworkList("entries", o.Entries); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = append(errors, err)
	}

	if err := ValidatePortStringList("ports", o.Ports); err != nil {
		errors = append(errors, err)
	}

	if err := ValidateProtocolList("protocols", o.Protocols); err != nil {
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
func (*ExternalNetwork) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ExternalNetworkAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ExternalNetworkLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ExternalNetwork) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ExternalNetworkAttributesMap
}

// ExternalNetworkAttributesMap represents the map of attribute for ExternalNetwork.
var ExternalNetworkAttributesMap = map[string]elemental.AttributeSpecification{
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
		PrimaryKey:     true,
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
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Entries": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Entries",
		Description:    `List of CIDRs or domain name.`,
		Exposed:        true,
		Name:           "entries",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
		SubType:    "metadata_list",
		Type:       "external",
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
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"Ports": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Ports",
		DefaultValue: []string{
			"1:65535",
		},
		Description: `List of single ports or range (xx:yy).`,
		Exposed:     true,
		Name:        "ports",
		Required:    true,
		Stored:      true,
		SubType:     "string",
		Type:        "list",
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
	"Protocols": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protocols",
		DefaultValue: []string{
			"tcp",
		},
		Description: `List of protocols (tcp, udp, or protocol number).`,
		Exposed:     true,
		Name:        "protocols",
		Required:    true,
		Stored:      true,
		SubType:     "string",
		Type:        "list",
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

// ExternalNetworkLowerCaseAttributesMap represents the map of attribute for ExternalNetwork.
var ExternalNetworkLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		PrimaryKey:     true,
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
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"entries": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Entries",
		Description:    `List of CIDRs or domain name.`,
		Exposed:        true,
		Name:           "entries",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
		SubType:    "metadata_list",
		Type:       "external",
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
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"ports": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Ports",
		DefaultValue: []string{
			"1:65535",
		},
		Description: `List of single ports or range (xx:yy).`,
		Exposed:     true,
		Name:        "ports",
		Required:    true,
		Stored:      true,
		SubType:     "string",
		Type:        "list",
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
	"protocols": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protocols",
		DefaultValue: []string{
			"tcp",
		},
		Description: `List of protocols (tcp, udp, or protocol number).`,
		Exposed:     true,
		Name:        "protocols",
		Required:    true,
		Stored:      true,
		SubType:     "string",
		Type:        "list",
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
