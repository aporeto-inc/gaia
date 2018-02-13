package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
)

// HookPolicyModeValue represents the possible values for attribute "mode".
type HookPolicyModeValue string

const (
	// HookPolicyModeBoth represents the value Both.
	HookPolicyModeBoth HookPolicyModeValue = "Both"

	// HookPolicyModePost represents the value Post.
	HookPolicyModePost HookPolicyModeValue = "Post"

	// HookPolicyModePre represents the value Pre.
	HookPolicyModePre HookPolicyModeValue = "Pre"
)

// HookPolicyIdentity represents the Identity of the object.
var HookPolicyIdentity = elemental.Identity{
	Name:     "hookpolicy",
	Category: "hookpolicies",
	Private:  false,
}

// HookPoliciesList represents a list of HookPolicies
type HookPoliciesList []*HookPolicy

// ContentIdentity returns the identity of the objects in the list.
func (o HookPoliciesList) ContentIdentity() elemental.Identity {

	return HookPolicyIdentity
}

// Copy returns a pointer to a copy the HookPoliciesList.
func (o HookPoliciesList) Copy() elemental.ContentIdentifiable {

	copy := append(HookPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the HookPoliciesList.
func (o HookPoliciesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(HookPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*HookPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o HookPoliciesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o HookPoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o HookPoliciesList) Version() int {

	return 1
}

// HookPolicy represents the model of a hookpolicy
type HookPolicy struct {
	// CertificateAuthority contains the pem block of the certificate authority used by the remote endpoint.
	CertificateAuthority string `json:"certificateAuthority" bson:"certificateauthority" mapstructure:"certificateAuthority,omitempty"`

	// ClientCertificate contains the client certificate that will be used to connect to the remote endoint.
	ClientCertificate string `json:"clientCertificate" bson:"clientcertificate" mapstructure:"clientCertificate,omitempty"`

	// ClientCertificateKey contains the key associated to the clientCertificate.
	ClientCertificateKey string `json:"clientCertificateKey" bson:"clientcertificatekey" mapstructure:"clientCertificateKey,omitempty"`

	// Endpoint contains the full address of the remote processor endoint.
	Endpoint string `json:"endpoint" bson:"endpoint" mapstructure:"endpoint,omitempty"`

	// Mode define the type of the hook.
	Mode HookPolicyModeValue `json:"mode" bson:"mode" mapstructure:"mode,omitempty"`

	// Subject contains the tag expression that an object must match in order to trigger the hook.
	Subject [][]string `json:"subject" bson:"subject" mapstructure:"subject,omitempty"`

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

	// Disabled defines if the propert is disabled.
	Disabled bool `json:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate bool `json:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// If set to true while the policy is propagating, it won't be visible to children namespace, but still used for policy resolution.
	PropagationHidden bool `json:"propagationHidden" bson:"propagationhidden" mapstructure:"propagationHidden,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewHookPolicy returns a new *HookPolicy
func NewHookPolicy() *HookPolicy {

	return &HookPolicy{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Metadata:       []string{},
		Mode:           "Pre",
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *HookPolicy) Identity() elemental.Identity {

	return HookPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *HookPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *HookPolicy) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *HookPolicy) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *HookPolicy) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *HookPolicy) Doc() string {
	return `Hook allows to to define hooks to the write operations in squall. Hooks are sent to an external Rufus server that will do the processing and eventually return a modified version of the object before we save it.`
}

func (o *HookPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *HookPolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *HookPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *HookPolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *HookPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *HookPolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *HookPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *HookPolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *HookPolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *HookPolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *HookPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *HookPolicy) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *HookPolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *HookPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetDisabled returns the Disabled of the receiver.
func (o *HookPolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the given Disabled of the receiver.
func (o *HookPolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetMetadata returns the Metadata of the receiver.
func (o *HookPolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *HookPolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *HookPolicy) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *HookPolicy) SetName(name string) {

	o.Name = name
}

// GetPropagate returns the Propagate of the receiver.
func (o *HookPolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the given Propagate of the receiver.
func (o *HookPolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetPropagationHidden returns the PropagationHidden of the receiver.
func (o *HookPolicy) GetPropagationHidden() bool {

	return o.PropagationHidden
}

// SetPropagationHidden sets the given PropagationHidden of the receiver.
func (o *HookPolicy) SetPropagationHidden(propagationHidden bool) {

	o.PropagationHidden = propagationHidden
}

// Validate valides the current information stored into the structure.
func (o *HookPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("certificateAuthority", o.CertificateAuthority); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("certificateAuthority", o.CertificateAuthority); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("clientCertificate", o.ClientCertificate); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("clientCertificate", o.ClientCertificate); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("clientCertificateKey", o.ClientCertificateKey); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("clientCertificateKey", o.ClientCertificateKey); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("endpoint", o.Endpoint); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("endpoint", o.Endpoint); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("mode", string(o.Mode), []string{"Both", "Post", "Pre"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("subject", o.Subject); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("subject", o.Subject); err != nil {
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
func (*HookPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := HookPolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return HookPolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*HookPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return HookPolicyAttributesMap
}

// HookPolicyAttributesMap represents the map of attribute for HookPolicy.
var HookPolicyAttributesMap = map[string]elemental.AttributeSpecification{
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
		ReadOnly:       true,
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
	"CertificateAuthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateAuthority",
		Description:    `CertificateAuthority contains the pem block of the certificate authority used by the remote endpoint.`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificateAuthority",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ClientCertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientCertificate",
		Description:    `ClientCertificate contains the client certificate that will be used to connect to the remote endoint.`,
		Exposed:        true,
		Format:         "free",
		Name:           "clientCertificate",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ClientCertificateKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientCertificateKey",
		Description:    `ClientCertificateKey contains the key associated to the clientCertificate.`,
		Exposed:        true,
		Format:         "free",
		Name:           "clientCertificateKey",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
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
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Endpoint": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Endpoint",
		Description:    `Endpoint contains the full address of the remote processor endoint.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "endpoint",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
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
	"Mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Both", "Post", "Pre"},
		ConvertedName:  "Mode",
		DefaultValue:   HookPolicyModePre,
		Description:    `Mode define the type of the hook.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "mode",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
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
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"PropagationHidden": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PropagationHidden",
		Description:    `If set to true while the policy is propagating, it won't be visible to children namespace, but still used for policy resolution.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagationHidden",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `Subject contains the tag expression that an object must match in order to trigger the hook.`,
		Exposed:        true,
		Name:           "subject",
		Required:       true,
		Stored:         true,
		SubType:        "policies_list",
		Type:           "external",
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

// HookPolicyLowerCaseAttributesMap represents the map of attribute for HookPolicy.
var HookPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		ReadOnly:       true,
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
	"certificateauthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateAuthority",
		Description:    `CertificateAuthority contains the pem block of the certificate authority used by the remote endpoint.`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificateAuthority",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"clientcertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientCertificate",
		Description:    `ClientCertificate contains the client certificate that will be used to connect to the remote endoint.`,
		Exposed:        true,
		Format:         "free",
		Name:           "clientCertificate",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"clientcertificatekey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientCertificateKey",
		Description:    `ClientCertificateKey contains the key associated to the clientCertificate.`,
		Exposed:        true,
		Format:         "free",
		Name:           "clientCertificateKey",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
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
	"disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"endpoint": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Endpoint",
		Description:    `Endpoint contains the full address of the remote processor endoint.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "endpoint",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
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
	"mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Both", "Post", "Pre"},
		ConvertedName:  "Mode",
		DefaultValue:   HookPolicyModePre,
		Description:    `Mode define the type of the hook.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "mode",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
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
	"propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"propagationhidden": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PropagationHidden",
		Description:    `If set to true while the policy is propagating, it won't be visible to children namespace, but still used for policy resolution.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagationHidden",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description:    `Subject contains the tag expression that an object must match in order to trigger the hook.`,
		Exposed:        true,
		Name:           "subject",
		Required:       true,
		Stored:         true,
		SubType:        "policies_list",
		Type:           "external",
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
