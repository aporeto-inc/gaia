package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
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
	Package:  "squall",
	Private:  false,
}

// HookPoliciesList represents a list of HookPolicies
type HookPoliciesList []*HookPolicy

// Identity returns the identity of the objects in the list.
func (o HookPoliciesList) Identity() elemental.Identity {

	return HookPolicyIdentity
}

// Copy returns a pointer to a copy the HookPoliciesList.
func (o HookPoliciesList) Copy() elemental.Identifiables {

	copy := append(HookPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the HookPoliciesList.
func (o HookPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(HookPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*HookPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o HookPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o HookPoliciesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the HookPoliciesList converted to SparseHookPoliciesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o HookPoliciesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseHookPoliciesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseHookPolicy)
	}

	return out
}

// Version returns the version of the content.
func (o HookPoliciesList) Version() int {

	return 1
}

// HookPolicy represents the model of a hookpolicy
type HookPolicy struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Contains the PEM block of the certificate authority used by the remote endpoint.
	CertificateAuthority string `json:"certificateAuthority" msgpack:"certificateAuthority" bson:"certificateauthority" mapstructure:"certificateAuthority,omitempty"`

	// Contains the client certificate that will be used to connect
	// to the remote endpoint. If provided, the private key associated with this
	// certificate must also be configured.
	ClientCertificate string `json:"clientCertificate" msgpack:"clientCertificate" bson:"clientcertificate" mapstructure:"clientCertificate,omitempty"`

	// Contains the key associated with the `clientCertificate`. It must be provided
	// only
	// when `clientCertificate` has been configured.
	ClientCertificateKey string `json:"clientCertificateKey" msgpack:"clientCertificateKey" bson:"clientcertificatekey" mapstructure:"clientCertificateKey,omitempty"`

	// If set to `true` and `mode` is in `Pre`, the request will be honored even if
	// calling the hook fails.
	ContinueOnError bool `json:"continueOnError" msgpack:"continueOnError" bson:"continueonerror" mapstructure:"continueOnError,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled bool `json:"disabled" msgpack:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// Contains the full address of the remote processor endpoint.
	Endpoint string `json:"endpoint" msgpack:"endpoint" bson:"endpoint" mapstructure:"endpoint,omitempty"`

	// If set the hook will be automatically deleted after the given time.
	ExpirationTime time.Time `json:"expirationTime" msgpack:"expirationTime" bson:"expirationtime" mapstructure:"expirationTime,omitempty"`

	// Indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback bool `json:"fallback" msgpack:"fallback" bson:"fallback" mapstructure:"fallback,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Defines the type of hook.
	Mode HookPolicyModeValue `json:"mode" msgpack:"mode" bson:"mode" mapstructure:"mode,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// If set to `true` while the policy is propagating, it won't be visible to children
	// namespace, but still used for policy resolution.
	PropagationHidden bool `json:"propagationHidden" msgpack:"propagationHidden" bson:"propagationhidden" mapstructure:"propagationHidden,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Contains the tag expression that an object must match in order to trigger the
	// hook.
	Subject [][]string `json:"subject" msgpack:"subject" bson:"subject" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewHookPolicy returns a new *HookPolicy
func NewHookPolicy() *HookPolicy {

	return &HookPolicy{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Metadata:       []string{},
		Mode:           HookPolicyModePre,
		NormalizedTags: []string{},
		Subject:        [][]string{},
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

// BleveType implements the bleve.Classifier Interface.
func (o *HookPolicy) BleveType() string {

	return "hookpolicy"
}

// DefaultOrder returns the list of default ordering fields.
func (o *HookPolicy) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// Doc returns the documentation for the object
func (o *HookPolicy) Doc() string {

	return `Allows you to define hooks to the write operations in squall. Hooks are sent
to an external Rufus server that will do the processing and eventually return a
modified version of the object before we save it.`
}

func (o *HookPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *HookPolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *HookPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *HookPolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *HookPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *HookPolicy) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *HookPolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *HookPolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *HookPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *HookPolicy) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *HookPolicy) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *HookPolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *HookPolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *HookPolicy) GetExpirationTime() time.Time {

	return o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the given value.
func (o *HookPolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *HookPolicy) GetFallback() bool {

	return o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the given value.
func (o *HookPolicy) SetFallback(fallback bool) {

	o.Fallback = fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *HookPolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *HookPolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *HookPolicy) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *HookPolicy) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *HookPolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *HookPolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *HookPolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *HookPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *HookPolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *HookPolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetPropagationHidden returns the PropagationHidden of the receiver.
func (o *HookPolicy) GetPropagationHidden() bool {

	return o.PropagationHidden
}

// SetPropagationHidden sets the property PropagationHidden of the receiver using the given value.
func (o *HookPolicy) SetPropagationHidden(propagationHidden bool) {

	o.PropagationHidden = propagationHidden
}

// GetProtected returns the Protected of the receiver.
func (o *HookPolicy) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *HookPolicy) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *HookPolicy) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *HookPolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *HookPolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *HookPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *HookPolicy) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseHookPolicy{
			ID:                   &o.ID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CertificateAuthority: &o.CertificateAuthority,
			ClientCertificate:    &o.ClientCertificate,
			ClientCertificateKey: &o.ClientCertificateKey,
			ContinueOnError:      &o.ContinueOnError,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Description:          &o.Description,
			Disabled:             &o.Disabled,
			Endpoint:             &o.Endpoint,
			ExpirationTime:       &o.ExpirationTime,
			Fallback:             &o.Fallback,
			Metadata:             &o.Metadata,
			Mode:                 &o.Mode,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Propagate:            &o.Propagate,
			PropagationHidden:    &o.PropagationHidden,
			Protected:            &o.Protected,
			Subject:              &o.Subject,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
		}
	}

	sp := &SparseHookPolicy{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "certificateAuthority":
			sp.CertificateAuthority = &(o.CertificateAuthority)
		case "clientCertificate":
			sp.ClientCertificate = &(o.ClientCertificate)
		case "clientCertificateKey":
			sp.ClientCertificateKey = &(o.ClientCertificateKey)
		case "continueOnError":
			sp.ContinueOnError = &(o.ContinueOnError)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "endpoint":
			sp.Endpoint = &(o.Endpoint)
		case "expirationTime":
			sp.ExpirationTime = &(o.ExpirationTime)
		case "fallback":
			sp.Fallback = &(o.Fallback)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "mode":
			sp.Mode = &(o.Mode)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "propagate":
			sp.Propagate = &(o.Propagate)
		case "propagationHidden":
			sp.PropagationHidden = &(o.PropagationHidden)
		case "protected":
			sp.Protected = &(o.Protected)
		case "subject":
			sp.Subject = &(o.Subject)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseHookPolicy to the object.
func (o *HookPolicy) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseHookPolicy)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CertificateAuthority != nil {
		o.CertificateAuthority = *so.CertificateAuthority
	}
	if so.ClientCertificate != nil {
		o.ClientCertificate = *so.ClientCertificate
	}
	if so.ClientCertificateKey != nil {
		o.ClientCertificateKey = *so.ClientCertificateKey
	}
	if so.ContinueOnError != nil {
		o.ContinueOnError = *so.ContinueOnError
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Disabled != nil {
		o.Disabled = *so.Disabled
	}
	if so.Endpoint != nil {
		o.Endpoint = *so.Endpoint
	}
	if so.ExpirationTime != nil {
		o.ExpirationTime = *so.ExpirationTime
	}
	if so.Fallback != nil {
		o.Fallback = *so.Fallback
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.Mode != nil {
		o.Mode = *so.Mode
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
	if so.PropagationHidden != nil {
		o.PropagationHidden = *so.PropagationHidden
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Subject != nil {
		o.Subject = *so.Subject
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the HookPolicy.
func (o *HookPolicy) DeepCopy() *HookPolicy {

	if o == nil {
		return nil
	}

	out := &HookPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *HookPolicy.
func (o *HookPolicy) DeepCopyInto(out *HookPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy HookPolicy: %s", err))
	}

	*out = *target.(*HookPolicy)
}

// Validate valides the current information stored into the structure.
func (o *HookPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidatePEM("certificateAuthority", o.CertificateAuthority); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidatePEM("clientCertificate", o.ClientCertificate); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("endpoint", o.Endpoint); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("mode", string(o.Mode), []string{"Both", "Post", "Pre"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("subject", o.Subject); err != nil {
		errors = errors.Append(err)
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *HookPolicy) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "certificateAuthority":
		return o.CertificateAuthority
	case "clientCertificate":
		return o.ClientCertificate
	case "clientCertificateKey":
		return o.ClientCertificateKey
	case "continueOnError":
		return o.ContinueOnError
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "disabled":
		return o.Disabled
	case "endpoint":
		return o.Endpoint
	case "expirationTime":
		return o.ExpirationTime
	case "fallback":
		return o.Fallback
	case "metadata":
		return o.Metadata
	case "mode":
		return o.Mode
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "propagate":
		return o.Propagate
	case "propagationHidden":
		return o.PropagationHidden
	case "protected":
		return o.Protected
	case "subject":
		return o.Subject
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	}

	return nil
}

// HookPolicyAttributesMap represents the map of attribute for HookPolicy.
var HookPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
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
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CertificateAuthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateAuthority",
		Description:    `Contains the PEM block of the certificate authority used by the remote endpoint.`,
		Exposed:        true,
		Name:           "certificateAuthority",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ClientCertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientCertificate",
		Description: `Contains the client certificate that will be used to connect
to the remote endpoint. If provided, the private key associated with this
certificate must also be configured.`,
		Exposed:   true,
		Name:      "clientCertificate",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"ClientCertificateKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientCertificateKey",
		Description: `Contains the key associated with the ` + "`" + `clientCertificate` + "`" + `. It must be provided
only
when ` + "`" + `clientCertificate` + "`" + ` has been configured.`,
		Exposed:   true,
		Name:      "clientCertificateKey",
		Orderable: true,
		Secret:    true,
		Stored:    true,
		Transient: true,
		Type:      "string",
	},
	"ContinueOnError": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ContinueOnError",
		Description: `If set to ` + "`" + `true` + "`" + ` and ` + "`" + `mode` + "`" + ` is in ` + "`" + `Pre` + "`" + `, the request will be honored even if
calling the hook fails.`,
		Exposed: true,
		Name:    "continueOnError",
		Stored:  true,
		Type:    "boolean",
	},
	"CreateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
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
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
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
		Description:    `Contains the full address of the remote processor endpoint.`,
		Exposed:        true,
		Name:           "endpoint",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ExpirationTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `If set the hook will be automatically deleted after the given time.`,
		Exposed:        true,
		Getter:         true,
		Name:           "expirationTime",
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:   true,
		Getter:    true,
		Name:      "fallback",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
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
	"Mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Both", "Post", "Pre"},
		ConvertedName:  "Mode",
		DefaultValue:   HookPolicyModePre,
		Description:    `Defines the type of hook.`,
		Exposed:        true,
		Name:           "mode",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name of the entity.`,
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
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
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
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
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
		Description: `If set to ` + "`" + `true` + "`" + ` while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.`,
		Exposed:   true,
		Getter:    true,
		Name:      "propagationHidden",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description: `Contains the tag expression that an object must match in order to trigger the
hook.`,
		Exposed: true,
		Name:    "subject",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"UpdateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
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
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
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
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"certificateauthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateAuthority",
		Description:    `Contains the PEM block of the certificate authority used by the remote endpoint.`,
		Exposed:        true,
		Name:           "certificateAuthority",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"clientcertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientCertificate",
		Description: `Contains the client certificate that will be used to connect
to the remote endpoint. If provided, the private key associated with this
certificate must also be configured.`,
		Exposed:   true,
		Name:      "clientCertificate",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"clientcertificatekey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClientCertificateKey",
		Description: `Contains the key associated with the ` + "`" + `clientCertificate` + "`" + `. It must be provided
only
when ` + "`" + `clientCertificate` + "`" + ` has been configured.`,
		Exposed:   true,
		Name:      "clientCertificateKey",
		Orderable: true,
		Secret:    true,
		Stored:    true,
		Transient: true,
		Type:      "string",
	},
	"continueonerror": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ContinueOnError",
		Description: `If set to ` + "`" + `true` + "`" + ` and ` + "`" + `mode` + "`" + ` is in ` + "`" + `Pre` + "`" + `, the request will be honored even if
calling the hook fails.`,
		Exposed: true,
		Name:    "continueOnError",
		Stored:  true,
		Type:    "boolean",
	},
	"createidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
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
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
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
		Description:    `Contains the full address of the remote processor endpoint.`,
		Exposed:        true,
		Name:           "endpoint",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"expirationtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `If set the hook will be automatically deleted after the given time.`,
		Exposed:        true,
		Getter:         true,
		Name:           "expirationTime",
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:   true,
		Getter:    true,
		Name:      "fallback",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
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
	"mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Both", "Post", "Pre"},
		ConvertedName:  "Mode",
		DefaultValue:   HookPolicyModePre,
		Description:    `Defines the type of hook.`,
		Exposed:        true,
		Name:           "mode",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name of the entity.`,
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
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
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
	"propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
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
		Description: `If set to ` + "`" + `true` + "`" + ` while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.`,
		Exposed:   true,
		Getter:    true,
		Name:      "propagationHidden",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description: `Contains the tag expression that an object must match in order to trigger the
hook.`,
		Exposed: true,
		Name:    "subject",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"updateidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
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

// SparseHookPoliciesList represents a list of SparseHookPolicies
type SparseHookPoliciesList []*SparseHookPolicy

// Identity returns the identity of the objects in the list.
func (o SparseHookPoliciesList) Identity() elemental.Identity {

	return HookPolicyIdentity
}

// Copy returns a pointer to a copy the SparseHookPoliciesList.
func (o SparseHookPoliciesList) Copy() elemental.Identifiables {

	copy := append(SparseHookPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseHookPoliciesList.
func (o SparseHookPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseHookPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseHookPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseHookPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseHookPoliciesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparseHookPoliciesList converted to HookPoliciesList.
func (o SparseHookPoliciesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseHookPoliciesList) Version() int {

	return 1
}

// SparseHookPolicy represents the sparse version of a hookpolicy.
type SparseHookPolicy struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// Contains the PEM block of the certificate authority used by the remote endpoint.
	CertificateAuthority *string `json:"certificateAuthority,omitempty" msgpack:"certificateAuthority,omitempty" bson:"certificateauthority,omitempty" mapstructure:"certificateAuthority,omitempty"`

	// Contains the client certificate that will be used to connect
	// to the remote endpoint. If provided, the private key associated with this
	// certificate must also be configured.
	ClientCertificate *string `json:"clientCertificate,omitempty" msgpack:"clientCertificate,omitempty" bson:"clientcertificate,omitempty" mapstructure:"clientCertificate,omitempty"`

	// Contains the key associated with the `clientCertificate`. It must be provided
	// only
	// when `clientCertificate` has been configured.
	ClientCertificateKey *string `json:"clientCertificateKey,omitempty" msgpack:"clientCertificateKey,omitempty" bson:"clientcertificatekey,omitempty" mapstructure:"clientCertificateKey,omitempty"`

	// If set to `true` and `mode` is in `Pre`, the request will be honored even if
	// calling the hook fails.
	ContinueOnError *bool `json:"continueOnError,omitempty" msgpack:"continueOnError,omitempty" bson:"continueonerror,omitempty" mapstructure:"continueOnError,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled *bool `json:"disabled,omitempty" msgpack:"disabled,omitempty" bson:"disabled,omitempty" mapstructure:"disabled,omitempty"`

	// Contains the full address of the remote processor endpoint.
	Endpoint *string `json:"endpoint,omitempty" msgpack:"endpoint,omitempty" bson:"endpoint,omitempty" mapstructure:"endpoint,omitempty"`

	// If set the hook will be automatically deleted after the given time.
	ExpirationTime *time.Time `json:"expirationTime,omitempty" msgpack:"expirationTime,omitempty" bson:"expirationtime,omitempty" mapstructure:"expirationTime,omitempty"`

	// Indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback *bool `json:"fallback,omitempty" msgpack:"fallback,omitempty" bson:"fallback,omitempty" mapstructure:"fallback,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Defines the type of hook.
	Mode *HookPolicyModeValue `json:"mode,omitempty" msgpack:"mode,omitempty" bson:"mode,omitempty" mapstructure:"mode,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// If set to `true` while the policy is propagating, it won't be visible to children
	// namespace, but still used for policy resolution.
	PropagationHidden *bool `json:"propagationHidden,omitempty" msgpack:"propagationHidden,omitempty" bson:"propagationhidden,omitempty" mapstructure:"propagationHidden,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Contains the tag expression that an object must match in order to trigger the
	// hook.
	Subject *[][]string `json:"subject,omitempty" msgpack:"subject,omitempty" bson:"subject,omitempty" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseHookPolicy returns a new  SparseHookPolicy.
func NewSparseHookPolicy() *SparseHookPolicy {
	return &SparseHookPolicy{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseHookPolicy) Identity() elemental.Identity {

	return HookPolicyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseHookPolicy) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseHookPolicy) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseHookPolicy) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseHookPolicy) ToPlain() elemental.PlainIdentifiable {

	out := NewHookPolicy()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CertificateAuthority != nil {
		out.CertificateAuthority = *o.CertificateAuthority
	}
	if o.ClientCertificate != nil {
		out.ClientCertificate = *o.ClientCertificate
	}
	if o.ClientCertificateKey != nil {
		out.ClientCertificateKey = *o.ClientCertificateKey
	}
	if o.ContinueOnError != nil {
		out.ContinueOnError = *o.ContinueOnError
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Disabled != nil {
		out.Disabled = *o.Disabled
	}
	if o.Endpoint != nil {
		out.Endpoint = *o.Endpoint
	}
	if o.ExpirationTime != nil {
		out.ExpirationTime = *o.ExpirationTime
	}
	if o.Fallback != nil {
		out.Fallback = *o.Fallback
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Mode != nil {
		out.Mode = *o.Mode
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}
	if o.PropagationHidden != nil {
		out.PropagationHidden = *o.PropagationHidden
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Subject != nil {
		out.Subject = *o.Subject
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseHookPolicy) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseHookPolicy) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseHookPolicy) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseHookPolicy) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseHookPolicy) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseHookPolicy) GetDisabled() bool {

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *SparseHookPolicy) GetExpirationTime() time.Time {

	return *o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = &expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *SparseHookPolicy) GetFallback() bool {

	return *o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetFallback(fallback bool) {

	o.Fallback = &fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseHookPolicy) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseHookPolicy) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseHookPolicy) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseHookPolicy) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseHookPolicy) GetPropagate() bool {

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetPropagationHidden returns the PropagationHidden of the receiver.
func (o *SparseHookPolicy) GetPropagationHidden() bool {

	return *o.PropagationHidden
}

// SetPropagationHidden sets the property PropagationHidden of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetPropagationHidden(propagationHidden bool) {

	o.PropagationHidden = &propagationHidden
}

// GetProtected returns the Protected of the receiver.
func (o *SparseHookPolicy) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseHookPolicy) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseHookPolicy) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseHookPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseHookPolicy.
func (o *SparseHookPolicy) DeepCopy() *SparseHookPolicy {

	if o == nil {
		return nil
	}

	out := &SparseHookPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseHookPolicy.
func (o *SparseHookPolicy) DeepCopyInto(out *SparseHookPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseHookPolicy: %s", err))
	}

	*out = *target.(*SparseHookPolicy)
}
