package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AWSAPIGatewayIdentity represents the Identity of the object.
var AWSAPIGatewayIdentity = elemental.Identity{
	Name:     "awsapigateway",
	Category: "awsapigateways",
	Package:  "goldrush",
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

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AWSAPIGatewaysList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the AWSAPIGatewaysList converted to SparseAWSAPIGatewaysList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AWSAPIGatewaysList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
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
		"namespace",
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

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *AWSAPIGateway) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *AWSAPIGateway) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *AWSAPIGateway) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *AWSAPIGateway) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *AWSAPIGateway) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *AWSAPIGateway) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *AWSAPIGateway) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *AWSAPIGateway) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *AWSAPIGateway) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *AWSAPIGateway) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *AWSAPIGateway) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *AWSAPIGateway) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *AWSAPIGateway) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *AWSAPIGateway) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
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

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *AWSAPIGateway) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AWSAPIGateway) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAWSAPIGateway{
			APIID:          &o.APIID,
			ID:             &o.ID,
			AccountID:      &o.AccountID,
			Annotations:    &o.Annotations,
			AssociatedTags: &o.AssociatedTags,
			Authorized:     &o.Authorized,
			CreateTime:     &o.CreateTime,
			Description:    &o.Description,
			Metadata:       &o.Metadata,
			Method:         &o.Method,
			Name:           &o.Name,
			Namespace:      &o.Namespace,
			NamespaceID:    &o.NamespaceID,
			NormalizedTags: &o.NormalizedTags,
			Protected:      &o.Protected,
			Resource:       &o.Resource,
			SourceIP:       &o.SourceIP,
			Stage:          &o.Stage,
			Token:          &o.Token,
			UpdateTime:     &o.UpdateTime,
		}
	}

	sp := &SparseAWSAPIGateway{}
	for _, f := range fields {
		switch f {
		case "APIID":
			sp.APIID = &(o.APIID)
		case "ID":
			sp.ID = &(o.ID)
		case "accountID":
			sp.AccountID = &(o.AccountID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "authorized":
			sp.Authorized = &(o.Authorized)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "method":
			sp.Method = &(o.Method)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "namespaceID":
			sp.NamespaceID = &(o.NamespaceID)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "resource":
			sp.Resource = &(o.Resource)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
		case "stage":
			sp.Stage = &(o.Stage)
		case "token":
			sp.Token = &(o.Token)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAWSAPIGateway to the object.
func (o *AWSAPIGateway) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAWSAPIGateway)
	if so.APIID != nil {
		o.APIID = *so.APIID
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.AccountID != nil {
		o.AccountID = *so.AccountID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.Authorized != nil {
		o.Authorized = *so.Authorized
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.Method != nil {
		o.Method = *so.Method
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NamespaceID != nil {
		o.NamespaceID = *so.NamespaceID
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Resource != nil {
		o.Resource = *so.Resource
	}
	if so.SourceIP != nil {
		o.SourceIP = *so.SourceIP
	}
	if so.Stage != nil {
		o.Stage = *so.Stage
	}
	if so.Token != nil {
		o.Token = *so.Token
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the AWSAPIGateway.
func (o *AWSAPIGateway) DeepCopy() *AWSAPIGateway {

	if o == nil {
		return nil
	}

	out := &AWSAPIGateway{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AWSAPIGateway.
func (o *AWSAPIGateway) DeepCopyInto(out *AWSAPIGateway) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AWSAPIGateway: %s", err))
	}

	*out = *target.(*AWSAPIGateway)
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AWSAPIGateway) ValueForAttribute(name string) interface{} {

	switch name {
	case "APIID":
		return o.APIID
	case "ID":
		return o.ID
	case "accountID":
		return o.AccountID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "authorized":
		return o.Authorized
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "metadata":
		return o.Metadata
	case "method":
		return o.Method
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "namespaceID":
		return o.NamespaceID
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "resource":
		return o.Resource
	case "sourceIP":
		return o.SourceIP
	case "stage":
		return o.Stage
	case "token":
		return o.Token
	case "updateTime":
		return o.UpdateTime
	}

	return nil
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
		SubType:        "map[string][]string",
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
		SubType:        "string",
		Type:           "list",
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
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
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
		SubType:    "string",
		Type:       "list",
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
		SubType:        "string",
		Transient:      true,
		Type:           "list",
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
		SubType:        "map[string][]string",
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
		SubType:        "string",
		Type:           "list",
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
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
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
		SubType:    "string",
		Type:       "list",
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
		SubType:        "string",
		Transient:      true,
		Type:           "list",
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

// SparseAWSAPIGatewaysList represents a list of SparseAWSAPIGateways
type SparseAWSAPIGatewaysList []*SparseAWSAPIGateway

// Identity returns the identity of the objects in the list.
func (o SparseAWSAPIGatewaysList) Identity() elemental.Identity {

	return AWSAPIGatewayIdentity
}

// Copy returns a pointer to a copy the SparseAWSAPIGatewaysList.
func (o SparseAWSAPIGatewaysList) Copy() elemental.Identifiables {

	copy := append(SparseAWSAPIGatewaysList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAWSAPIGatewaysList.
func (o SparseAWSAPIGatewaysList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAWSAPIGatewaysList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAWSAPIGateway))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAWSAPIGatewaysList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAWSAPIGatewaysList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparseAWSAPIGatewaysList converted to AWSAPIGatewaysList.
func (o SparseAWSAPIGatewaysList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAWSAPIGatewaysList) Version() int {

	return 1
}

// SparseAWSAPIGateway represents the sparse version of a awsapigateway.
type SparseAWSAPIGateway struct {
	// API ID as defined on AWS for the API that handled this request.
	APIID *string `json:"APIID,omitempty" bson:"-" mapstructure:"APIID,omitempty"`

	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// the account ID for the gateway managing this request.
	AccountID *string `json:"accountID,omitempty" bson:"-" mapstructure:"accountID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// The policy decision for this API flow.
	Authorized *bool `json:"authorized,omitempty" bson:"-" mapstructure:"authorized,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" bson:"metadata" mapstructure:"metadata,omitempty"`

	// API method that handled this request.
	Method *string `json:"method,omitempty" bson:"-" mapstructure:"method,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Link to the cluster namespace where the AWS API gateway is defined.
	NamespaceID *string `json:"namespaceID,omitempty" bson:"-" mapstructure:"namespaceID,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected" mapstructure:"protected,omitempty"`

	// API resource that handled this request.
	Resource *string `json:"resource,omitempty" bson:"-" mapstructure:"resource,omitempty"`

	// the client ip for this request.
	SourceIP *string `json:"sourceIP,omitempty" bson:"-" mapstructure:"sourceIP,omitempty"`

	// the stage name as defined on AWS for the API that handled this request.
	Stage *string `json:"stage,omitempty" bson:"-" mapstructure:"stage,omitempty"`

	// the JWT token that was optionally attached to this request.
	Token *string `json:"token,omitempty" bson:"-" mapstructure:"token,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseAWSAPIGateway returns a new  SparseAWSAPIGateway.
func NewSparseAWSAPIGateway() *SparseAWSAPIGateway {
	return &SparseAWSAPIGateway{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAWSAPIGateway) Identity() elemental.Identity {

	return AWSAPIGatewayIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAWSAPIGateway) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAWSAPIGateway) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseAWSAPIGateway) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAWSAPIGateway) ToPlain() elemental.PlainIdentifiable {

	out := NewAWSAPIGateway()
	if o.APIID != nil {
		out.APIID = *o.APIID
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.AccountID != nil {
		out.AccountID = *o.AccountID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.Authorized != nil {
		out.Authorized = *o.Authorized
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Method != nil {
		out.Method = *o.Method
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NamespaceID != nil {
		out.NamespaceID = *o.NamespaceID
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Resource != nil {
		out.Resource = *o.Resource
	}
	if o.SourceIP != nil {
		out.SourceIP = *o.SourceIP
	}
	if o.Stage != nil {
		out.Stage = *o.Stage
	}
	if o.Token != nil {
		out.Token = *o.Token
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseAWSAPIGateway) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseAWSAPIGateway) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseAWSAPIGateway) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseAWSAPIGateway) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseAWSAPIGateway) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseAWSAPIGateway) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseAWSAPIGateway) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseAWSAPIGateway) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseAWSAPIGateway) GetProtected() bool {

	return *o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseAWSAPIGateway) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseAWSAPIGateway.
func (o *SparseAWSAPIGateway) DeepCopy() *SparseAWSAPIGateway {

	if o == nil {
		return nil
	}

	out := &SparseAWSAPIGateway{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAWSAPIGateway.
func (o *SparseAWSAPIGateway) DeepCopyInto(out *SparseAWSAPIGateway) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAWSAPIGateway: %s", err))
	}

	*out = *target.(*SparseAWSAPIGateway)
}
