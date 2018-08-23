package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// AWSAPIGatewayIndexes lists the attribute compound indexes.
var AWSAPIGatewayIndexes = [][]string{}

// AWSAPIGatewayIdentity represents the Identity of the object.
var AWSAPIGatewayIdentity = elemental.Identity{
	Name:     "awsapigateway",
	Category: "awsapigateways",
	Private:  false,
}

// AWSAPIGatewaysList represents a list of AWSAPIGateways
type AWSAPIGatewaysList []*AWSAPIGateway

// Identity returns the identity of the objects in the list.
func (o AWSAPIGatewaysList) Identity() elemental.Identity {

	return AWSAPIGatewayIdentity
}

// Copy returns a pointer to a copy the AWSAPIGatewaysList.
func (o AWSAPIGatewaysList) Copy() elemental.Identifiables {

	copy := append(AWSAPIGatewaysList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AWSAPIGatewaysList.
func (o AWSAPIGatewaysList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AWSAPIGatewaysList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AWSAPIGateway))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AWSAPIGatewaysList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AWSAPIGatewaysList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o AWSAPIGatewaysList) Version() int {

	return 1
}

// AWSAPIGateway represents the model of a awsapigateway
type AWSAPIGateway struct {
	// API ID as defined on AWS for the API that handled this request.
	APIID string `json:"APIID" bson:"-" mapstructure:"APIID,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// the account ID for the gateway managing this request.
	AccountID string `json:"accountID" bson:"-" mapstructure:"accountID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// The policy decision for this API flow.
	Authorized bool `json:"authorized" bson:"-" mapstructure:"authorized,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// API method that handled this request.
	Method string `json:"method" bson:"-" mapstructure:"method,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Link to the cluster namespace where the AWS API gateway is defined.
	NamespaceID string `json:"namespaceID" bson:"-" mapstructure:"namespaceID,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// API resource that handled this request.
	Resource string `json:"resource" bson:"-" mapstructure:"resource,omitempty"`

	// the client ip for this request.
	SourceIP string `json:"sourceIP" bson:"-" mapstructure:"sourceIP,omitempty"`

	// the stage name as defined on AWS for the API that handled this request.
	Stage string `json:"stage" bson:"-" mapstructure:"stage,omitempty"`

	// the JWT token that was optionally attached to this request.
	Token string `json:"token" bson:"-" mapstructure:"token,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewAWSAPIGateway returns a new *AWSAPIGateway
func NewAWSAPIGateway() *AWSAPIGateway {

	return &AWSAPIGateway{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Metadata:       []string{},
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *AWSAPIGateway) Identity() elemental.Identity {

	return AWSAPIGatewayIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AWSAPIGateway) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AWSAPIGateway) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *AWSAPIGateway) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *AWSAPIGateway) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *AWSAPIGateway) Doc() string {
	return `managed API decisions for the AWS API Gateway.`
}

func (o *AWSAPIGateway) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *AWSAPIGateway) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *AWSAPIGateway) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *AWSAPIGateway) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *AWSAPIGateway) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *AWSAPIGateway) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *AWSAPIGateway) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *AWSAPIGateway) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *AWSAPIGateway) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *AWSAPIGateway) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *AWSAPIGateway) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *AWSAPIGateway) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *AWSAPIGateway) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *AWSAPIGateway) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *AWSAPIGateway) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *AWSAPIGateway) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *AWSAPIGateway) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *AWSAPIGateway) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *AWSAPIGateway) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
func (*AWSAPIGateway) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AWSAPIGatewayAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AWSAPIGatewayLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AWSAPIGateway) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AWSAPIGatewayAttributesMap
}

// AWSAPIGatewayAttributesMap represents the map of attribute for AWSAPIGateway.
var AWSAPIGatewayAttributesMap = map[string]elemental.AttributeSpecification{
	"APIID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "APIID",
		Description:    `API ID as defined on AWS for the API that handled this request.`,
		Exposed:        true,
		Name:           "APIID",
		Type:           "string",
	},
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
	"AccountID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccountID",
		Description:    `the account ID for the gateway managing this request.`,
		Exposed:        true,
		Name:           "accountID",
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
	"Authorized": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Authorized",
		Description:    `The policy decision for this API flow.`,
		Exposed:        true,
		Name:           "authorized",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "boolean",
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
	"Method": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Method",
		Description:    `API method that handled this request.`,
		Exposed:        true,
		Name:           "method",
		Type:           "string",
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
	"NamespaceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NamespaceID",
		Description:    `Link to the cluster namespace where the AWS API gateway is defined.`,
		Exposed:        true,
		Name:           "namespaceID",
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
	"Resource": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Resource",
		Description:    `API resource that handled this request.`,
		Exposed:        true,
		Name:           "resource",
		Type:           "string",
	},
	"SourceIP": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `the client ip for this request.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"Stage": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Stage",
		Description:    `the stage name as defined on AWS for the API that handled this request.`,
		Exposed:        true,
		Name:           "stage",
		Type:           "string",
	},
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `the JWT token that was optionally attached to this request.`,
		Exposed:        true,
		Name:           "token",
		Type:           "string",
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

// AWSAPIGatewayLowerCaseAttributesMap represents the map of attribute for AWSAPIGateway.
var AWSAPIGatewayLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"apiid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "APIID",
		Description:    `API ID as defined on AWS for the API that handled this request.`,
		Exposed:        true,
		Name:           "APIID",
		Type:           "string",
	},
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
	"accountid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccountID",
		Description:    `the account ID for the gateway managing this request.`,
		Exposed:        true,
		Name:           "accountID",
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
	"authorized": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Authorized",
		Description:    `The policy decision for this API flow.`,
		Exposed:        true,
		Name:           "authorized",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "boolean",
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
	"method": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Method",
		Description:    `API method that handled this request.`,
		Exposed:        true,
		Name:           "method",
		Type:           "string",
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
	"namespaceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NamespaceID",
		Description:    `Link to the cluster namespace where the AWS API gateway is defined.`,
		Exposed:        true,
		Name:           "namespaceID",
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
	"resource": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Resource",
		Description:    `API resource that handled this request.`,
		Exposed:        true,
		Name:           "resource",
		Type:           "string",
	},
	"sourceip": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `the client ip for this request.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"stage": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Stage",
		Description:    `the stage name as defined on AWS for the API that handled this request.`,
		Exposed:        true,
		Name:           "stage",
		Type:           "string",
	},
	"token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `the JWT token that was optionally attached to this request.`,
		Exposed:        true,
		Name:           "token",
		Type:           "string",
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
