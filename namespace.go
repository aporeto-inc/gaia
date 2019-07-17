package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NamespaceJWTCertificateTypeValue represents the possible values for attribute "JWTCertificateType".
type NamespaceJWTCertificateTypeValue string

const (
	// NamespaceJWTCertificateTypeEC represents the value EC.
	NamespaceJWTCertificateTypeEC NamespaceJWTCertificateTypeValue = "EC"

	// NamespaceJWTCertificateTypeNone represents the value None.
	NamespaceJWTCertificateTypeNone NamespaceJWTCertificateTypeValue = "None"

	// NamespaceJWTCertificateTypeRSA represents the value RSA.
	NamespaceJWTCertificateTypeRSA NamespaceJWTCertificateTypeValue = "RSA"
)

// NamespaceIdentity represents the Identity of the object.
var NamespaceIdentity = elemental.Identity{
	Name:     "namespace",
	Category: "namespaces",
	Package:  "squall",
	Private:  false,
}

// NamespacesList represents a list of Namespaces
type NamespacesList []*Namespace

// Identity returns the identity of the objects in the list.
func (o NamespacesList) Identity() elemental.Identity {

	return NamespaceIdentity
}

// Copy returns a pointer to a copy the NamespacesList.
func (o NamespacesList) Copy() elemental.Identifiables {

	copy := append(NamespacesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the NamespacesList.
func (o NamespacesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(NamespacesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Namespace))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o NamespacesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o NamespacesList) DefaultOrder() []string {

	return []string{
		"name",
		"namespace",
	}
}

// ToSparse returns the NamespacesList converted to SparseNamespacesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o NamespacesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseNamespacesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseNamespace)
	}

	return out
}

// Version returns the version of the content.
func (o NamespacesList) Version() int {

	return 1
}

// Namespace represents the model of a namespace
type Namespace struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// JWTCertificateType defines the JWT signing certificate that must be created
	// for this namespace. If the type is none no certificate will be created.
	JWTCertificateType NamespaceJWTCertificateTypeValue `json:"JWTCertificateType" msgpack:"JWTCertificateType" bson:"jwtcertificatetype" mapstructure:"JWTCertificateType,omitempty"`

	// JWTCertificates hold the certificates used to sign tokens for this namespace.
	// This is map indexed by the ID of the certificate.
	JWTCertificates map[string]string `json:"JWTCertificates" msgpack:"JWTCertificates" bson:"jwtcertificates" mapstructure:"JWTCertificates,omitempty"`

	// The SSH certificate authority used by the namespace.
	SSHCA string `json:"SSHCA" msgpack:"SSHCA" bson:"sshca" mapstructure:"SSHCA,omitempty"`

	// If `true`, an SSH certificate authority (CA) will be generated for the
	// namespace. This CA
	// can be deployed in SSH server to validate SSH certificates issued by the
	// platform.
	SSHCAEnabled bool `json:"SSHCAEnabled" msgpack:"SSHCAEnabled" bson:"sshcaenabled" mapstructure:"SSHCAEnabled,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedLocalCAID holds the remote ID of the certificate authority to use.
	AssociatedLocalCAID string `json:"-" msgpack:"-" bson:"associatedlocalcaid" mapstructure:"-,omitempty"`

	// The remote ID of the SSH certificate authority to use.
	AssociatedSSHCAID string `json:"associatedSSHCAID" msgpack:"associatedSSHCAID" bson:"associatedsshcaid" mapstructure:"associatedSSHCAID,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Defines if the namespace should inherit its parent zone. If this property is set
	// to `false`,
	// the `zoning` property will be ignored and the namespace will have the same zone
	// as its parent.
	CustomZoning bool `json:"customZoning" msgpack:"customZoning" bson:"customzoning" mapstructure:"customZoning,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// The certificate authority used by this namespace.
	LocalCA string `json:"localCA" msgpack:"localCA" bson:"localca" mapstructure:"localCA,omitempty"`

	// Defines if the namespace should use a local certificate
	// authority (CA). Switching it off and on again will regenerate a new CA.
	LocalCAEnabled bool `json:"localCAEnabled" msgpack:"localCAEnabled" bson:"localcaenabled" mapstructure:"localCAEnabled,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// The name of the namespace.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// List of tags that will be added to every `or` clause of all network access
	// policies in the namespace and its children.
	NetworkAccessPolicyTags []string `json:"networkAccessPolicyTags" msgpack:"networkAccessPolicyTags" bson:"networkaccesspolicytags" mapstructure:"networkAccessPolicyTags,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Determines the length of validity of certificates issued in this namespace using
	// [Golang duration syntax](https://golang.org/pkg/time/#example_Duration). Default
	// value is `1h`.
	ServiceCertificateValidity string `json:"serviceCertificateValidity" msgpack:"serviceCertificateValidity" bson:"servicecertificatevalidity" mapstructure:"serviceCertificateValidity,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	// Defines what zone the namespace should live in.
	Zoning int `json:"zoning" msgpack:"zoning" bson:"zoning" mapstructure:"zoning,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewNamespace returns a new *Namespace
func NewNamespace() *Namespace {

	return &Namespace{
		ModelVersion:               1,
		Annotations:                map[string][]string{},
		AssociatedTags:             []string{},
		NetworkAccessPolicyTags:    []string{},
		NormalizedTags:             []string{},
		JWTCertificates:            map[string]string{},
		Metadata:                   []string{},
		JWTCertificateType:         NamespaceJWTCertificateTypeNone,
		ServiceCertificateValidity: "1h",
	}
}

// Identity returns the Identity of the object.
func (o *Namespace) Identity() elemental.Identity {

	return NamespaceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Namespace) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Namespace) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Namespace) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Namespace) BleveType() string {

	return "namespace"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Namespace) DefaultOrder() []string {

	return []string{
		"name",
		"namespace",
	}
}

// Doc returns the documentation for the object
func (o *Namespace) Doc() string {

	return `A namespace represents the core organizational unit of the system. All objects
always exist in a single namespace. A namespace can also have child namespaces.
They can be used to split the system into organizations, business units,
applications, services or any combination you like.`
}

func (o *Namespace) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Namespace) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *Namespace) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Namespace) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *Namespace) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *Namespace) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *Namespace) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Namespace) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Namespace) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *Namespace) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *Namespace) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *Namespace) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *Namespace) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *Namespace) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Namespace) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Namespace) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Namespace) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Namespace) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *Namespace) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Namespace) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *Namespace) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *Namespace) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *Namespace) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Namespace) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Namespace) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *Namespace) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Namespace) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Namespace) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Namespace) SetZone(zone int) {

	o.Zone = zone
}

// GetZoning returns the Zoning of the receiver.
func (o *Namespace) GetZoning() int {

	return o.Zoning
}

// SetZoning sets the property Zoning of the receiver using the given value.
func (o *Namespace) SetZoning(zoning int) {

	o.Zoning = zoning
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Namespace) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseNamespace{
			ID:                         &o.ID,
			JWTCertificateType:         &o.JWTCertificateType,
			JWTCertificates:            &o.JWTCertificates,
			SSHCA:                      &o.SSHCA,
			SSHCAEnabled:               &o.SSHCAEnabled,
			Annotations:                &o.Annotations,
			AssociatedLocalCAID:        &o.AssociatedLocalCAID,
			AssociatedSSHCAID:          &o.AssociatedSSHCAID,
			AssociatedTags:             &o.AssociatedTags,
			CreateIdempotencyKey:       &o.CreateIdempotencyKey,
			CreateTime:                 &o.CreateTime,
			CustomZoning:               &o.CustomZoning,
			Description:                &o.Description,
			LocalCA:                    &o.LocalCA,
			LocalCAEnabled:             &o.LocalCAEnabled,
			Metadata:                   &o.Metadata,
			Name:                       &o.Name,
			Namespace:                  &o.Namespace,
			NetworkAccessPolicyTags:    &o.NetworkAccessPolicyTags,
			NormalizedTags:             &o.NormalizedTags,
			Protected:                  &o.Protected,
			ServiceCertificateValidity: &o.ServiceCertificateValidity,
			UpdateIdempotencyKey:       &o.UpdateIdempotencyKey,
			UpdateTime:                 &o.UpdateTime,
			ZHash:                      &o.ZHash,
			Zone:                       &o.Zone,
			Zoning:                     &o.Zoning,
		}
	}

	sp := &SparseNamespace{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "JWTCertificateType":
			sp.JWTCertificateType = &(o.JWTCertificateType)
		case "JWTCertificates":
			sp.JWTCertificates = &(o.JWTCertificates)
		case "SSHCA":
			sp.SSHCA = &(o.SSHCA)
		case "SSHCAEnabled":
			sp.SSHCAEnabled = &(o.SSHCAEnabled)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedLocalCAID":
			sp.AssociatedLocalCAID = &(o.AssociatedLocalCAID)
		case "associatedSSHCAID":
			sp.AssociatedSSHCAID = &(o.AssociatedSSHCAID)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "customZoning":
			sp.CustomZoning = &(o.CustomZoning)
		case "description":
			sp.Description = &(o.Description)
		case "localCA":
			sp.LocalCA = &(o.LocalCA)
		case "localCAEnabled":
			sp.LocalCAEnabled = &(o.LocalCAEnabled)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "networkAccessPolicyTags":
			sp.NetworkAccessPolicyTags = &(o.NetworkAccessPolicyTags)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "serviceCertificateValidity":
			sp.ServiceCertificateValidity = &(o.ServiceCertificateValidity)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		case "zoning":
			sp.Zoning = &(o.Zoning)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseNamespace to the object.
func (o *Namespace) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseNamespace)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.JWTCertificateType != nil {
		o.JWTCertificateType = *so.JWTCertificateType
	}
	if so.JWTCertificates != nil {
		o.JWTCertificates = *so.JWTCertificates
	}
	if so.SSHCA != nil {
		o.SSHCA = *so.SSHCA
	}
	if so.SSHCAEnabled != nil {
		o.SSHCAEnabled = *so.SSHCAEnabled
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedLocalCAID != nil {
		o.AssociatedLocalCAID = *so.AssociatedLocalCAID
	}
	if so.AssociatedSSHCAID != nil {
		o.AssociatedSSHCAID = *so.AssociatedSSHCAID
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.CustomZoning != nil {
		o.CustomZoning = *so.CustomZoning
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.LocalCA != nil {
		o.LocalCA = *so.LocalCA
	}
	if so.LocalCAEnabled != nil {
		o.LocalCAEnabled = *so.LocalCAEnabled
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NetworkAccessPolicyTags != nil {
		o.NetworkAccessPolicyTags = *so.NetworkAccessPolicyTags
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.ServiceCertificateValidity != nil {
		o.ServiceCertificateValidity = *so.ServiceCertificateValidity
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
	if so.Zoning != nil {
		o.Zoning = *so.Zoning
	}
}

// DeepCopy returns a deep copy if the Namespace.
func (o *Namespace) DeepCopy() *Namespace {

	if o == nil {
		return nil
	}

	out := &Namespace{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Namespace.
func (o *Namespace) DeepCopyInto(out *Namespace) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Namespace: %s", err))
	}

	*out = *target.(*Namespace)
}

// Validate valides the current information stored into the structure.
func (o *Namespace) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("JWTCertificateType", string(o.JWTCertificateType), []string{"RSA", "EC", "None"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidatePattern("name", o.Name, `^[a-zA-Z0-9-_/]+$`, `must only contain alpha numerical characters, '-' or '_'`, true); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTags("networkAccessPolicyTags", o.NetworkAccessPolicyTags); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTimeDuration("serviceCertificateValidity", o.ServiceCertificateValidity); err != nil {
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
func (*Namespace) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := NamespaceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return NamespaceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Namespace) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NamespaceAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Namespace) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "JWTCertificateType":
		return o.JWTCertificateType
	case "JWTCertificates":
		return o.JWTCertificates
	case "SSHCA":
		return o.SSHCA
	case "SSHCAEnabled":
		return o.SSHCAEnabled
	case "annotations":
		return o.Annotations
	case "associatedLocalCAID":
		return o.AssociatedLocalCAID
	case "associatedSSHCAID":
		return o.AssociatedSSHCAID
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "customZoning":
		return o.CustomZoning
	case "description":
		return o.Description
	case "localCA":
		return o.LocalCA
	case "localCAEnabled":
		return o.LocalCAEnabled
	case "metadata":
		return o.Metadata
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "networkAccessPolicyTags":
		return o.NetworkAccessPolicyTags
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "serviceCertificateValidity":
		return o.ServiceCertificateValidity
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	case "zoning":
		return o.Zoning
	}

	return nil
}

// NamespaceAttributesMap represents the map of attribute for Namespace.
var NamespaceAttributesMap = map[string]elemental.AttributeSpecification{
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
		Stored:         true,
		Type:           "string",
	},
	"JWTCertificateType": elemental.AttributeSpecification{
		AllowedChoices: []string{"RSA", "EC", "None"},
		ConvertedName:  "JWTCertificateType",
		DefaultValue:   NamespaceJWTCertificateTypeNone,
		Description: `JWTCertificateType defines the JWT signing certificate that must be created
for this namespace. If the type is none no certificate will be created.`,
		Exposed: true,
		Name:    "JWTCertificateType",
		Stored:  true,
		Type:    "enum",
	},
	"JWTCertificates": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "JWTCertificates",
		Description: `JWTCertificates hold the certificates used to sign tokens for this namespace.
This is map indexed by the ID of the certificate.`,
		Exposed:  true,
		Name:     "JWTCertificates",
		ReadOnly: true,
		Stored:   true,
		SubType:  "map[string]string",
		Type:     "external",
	},
	"SSHCA": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCA",
		Description:    `The SSH certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "SSHCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"SSHCAEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SSHCAEnabled",
		Description: `If ` + "`" + `true` + "`" + `, an SSH certificate authority (CA) will be generated for the
namespace. This CA 
can be deployed in SSH server to validate SSH certificates issued by the
platform.`,
		Exposed:   true,
		Name:      "SSHCAEnabled",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
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
	"AssociatedLocalCAID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedLocalCAID",
		Description:    `AssociatedLocalCAID holds the remote ID of the certificate authority to use.`,
		Name:           "associatedLocalCAID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AssociatedSSHCAID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedSSHCAID",
		Description:    `The remote ID of the SSH certificate authority to use.`,
		Exposed:        true,
		Name:           "associatedSSHCAID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
	"CustomZoning": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CustomZoning",
		CreationOnly:   true,
		Description: `Defines if the namespace should inherit its parent zone. If this property is set
to ` + "`" + `false` + "`" + `, 
the ` + "`" + `zoning` + "`" + ` property will be ignored and the namespace will have the same zone
as its parent.`,
		Exposed: true,
		Name:    "customZoning",
		Stored:  true,
		Type:    "boolean",
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
	"LocalCA": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LocalCA",
		Description:    `The certificate authority used by this namespace.`,
		Exposed:        true,
		Name:           "localCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"LocalCAEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LocalCAEnabled",
		Description: `Defines if the namespace should use a local certificate
authority (CA). Switching it off and on again will regenerate a new CA.`,
		Exposed:   true,
		Name:      "localCAEnabled",
		Orderable: true,
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
	"Name": elemental.AttributeSpecification{
		AllowedChars:   `^[a-zA-Z0-9-_/]+$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		DefaultOrder:   true,
		Description:    `The name of the namespace.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
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
	"NetworkAccessPolicyTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkAccessPolicyTags",
		Deprecated:     true,
		Description: `List of tags that will be added to every ` + "`" + `or` + "`" + ` clause of all network access
policies in the namespace and its children.`,
		Exposed: true,
		Name:    "networkAccessPolicyTags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
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
	"ServiceCertificateValidity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceCertificateValidity",
		DefaultValue:   "1h",
		Description: `Determines the length of validity of certificates issued in this namespace using
[Golang duration syntax](https://golang.org/pkg/time/#example_Duration). Default
value is ` + "`" + `1h` + "`" + `.`,
		Exposed: true,
		Name:    "serviceCertificateValidity",
		Stored:  true,
		Type:    "string",
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
	"ZHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
	"Zoning": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Zoning",
		CreationOnly:   true,
		Description:    `Defines what zone the namespace should live in.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zoning",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
}

// NamespaceLowerCaseAttributesMap represents the map of attribute for Namespace.
var NamespaceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		Stored:         true,
		Type:           "string",
	},
	"jwtcertificatetype": elemental.AttributeSpecification{
		AllowedChoices: []string{"RSA", "EC", "None"},
		ConvertedName:  "JWTCertificateType",
		DefaultValue:   NamespaceJWTCertificateTypeNone,
		Description: `JWTCertificateType defines the JWT signing certificate that must be created
for this namespace. If the type is none no certificate will be created.`,
		Exposed: true,
		Name:    "JWTCertificateType",
		Stored:  true,
		Type:    "enum",
	},
	"jwtcertificates": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "JWTCertificates",
		Description: `JWTCertificates hold the certificates used to sign tokens for this namespace.
This is map indexed by the ID of the certificate.`,
		Exposed:  true,
		Name:     "JWTCertificates",
		ReadOnly: true,
		Stored:   true,
		SubType:  "map[string]string",
		Type:     "external",
	},
	"sshca": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCA",
		Description:    `The SSH certificate authority used by the namespace.`,
		Exposed:        true,
		Name:           "SSHCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"sshcaenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SSHCAEnabled",
		Description: `If ` + "`" + `true` + "`" + `, an SSH certificate authority (CA) will be generated for the
namespace. This CA 
can be deployed in SSH server to validate SSH certificates issued by the
platform.`,
		Exposed:   true,
		Name:      "SSHCAEnabled",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
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
	"associatedlocalcaid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedLocalCAID",
		Description:    `AssociatedLocalCAID holds the remote ID of the certificate authority to use.`,
		Name:           "associatedLocalCAID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"associatedsshcaid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedSSHCAID",
		Description:    `The remote ID of the SSH certificate authority to use.`,
		Exposed:        true,
		Name:           "associatedSSHCAID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
	"customzoning": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CustomZoning",
		CreationOnly:   true,
		Description: `Defines if the namespace should inherit its parent zone. If this property is set
to ` + "`" + `false` + "`" + `, 
the ` + "`" + `zoning` + "`" + ` property will be ignored and the namespace will have the same zone
as its parent.`,
		Exposed: true,
		Name:    "customZoning",
		Stored:  true,
		Type:    "boolean",
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
	"localca": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LocalCA",
		Description:    `The certificate authority used by this namespace.`,
		Exposed:        true,
		Name:           "localCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"localcaenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LocalCAEnabled",
		Description: `Defines if the namespace should use a local certificate
authority (CA). Switching it off and on again will regenerate a new CA.`,
		Exposed:   true,
		Name:      "localCAEnabled",
		Orderable: true,
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
	"name": elemental.AttributeSpecification{
		AllowedChars:   `^[a-zA-Z0-9-_/]+$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		DefaultOrder:   true,
		Description:    `The name of the namespace.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
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
	"networkaccesspolicytags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkAccessPolicyTags",
		Deprecated:     true,
		Description: `List of tags that will be added to every ` + "`" + `or` + "`" + ` clause of all network access
policies in the namespace and its children.`,
		Exposed: true,
		Name:    "networkAccessPolicyTags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
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
	"servicecertificatevalidity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceCertificateValidity",
		DefaultValue:   "1h",
		Description: `Determines the length of validity of certificates issued in this namespace using
[Golang duration syntax](https://golang.org/pkg/time/#example_Duration). Default
value is ` + "`" + `1h` + "`" + `.`,
		Exposed: true,
		Name:    "serviceCertificateValidity",
		Stored:  true,
		Type:    "string",
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
	"zhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
	"zoning": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Zoning",
		CreationOnly:   true,
		Description:    `Defines what zone the namespace should live in.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zoning",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
}

// SparseNamespacesList represents a list of SparseNamespaces
type SparseNamespacesList []*SparseNamespace

// Identity returns the identity of the objects in the list.
func (o SparseNamespacesList) Identity() elemental.Identity {

	return NamespaceIdentity
}

// Copy returns a pointer to a copy the SparseNamespacesList.
func (o SparseNamespacesList) Copy() elemental.Identifiables {

	copy := append(SparseNamespacesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseNamespacesList.
func (o SparseNamespacesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseNamespacesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseNamespace))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseNamespacesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseNamespacesList) DefaultOrder() []string {

	return []string{
		"name",
		"namespace",
	}
}

// ToPlain returns the SparseNamespacesList converted to NamespacesList.
func (o SparseNamespacesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseNamespacesList) Version() int {

	return 1
}

// SparseNamespace represents the sparse version of a namespace.
type SparseNamespace struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// JWTCertificateType defines the JWT signing certificate that must be created
	// for this namespace. If the type is none no certificate will be created.
	JWTCertificateType *NamespaceJWTCertificateTypeValue `json:"JWTCertificateType,omitempty" msgpack:"JWTCertificateType,omitempty" bson:"jwtcertificatetype,omitempty" mapstructure:"JWTCertificateType,omitempty"`

	// JWTCertificates hold the certificates used to sign tokens for this namespace.
	// This is map indexed by the ID of the certificate.
	JWTCertificates *map[string]string `json:"JWTCertificates,omitempty" msgpack:"JWTCertificates,omitempty" bson:"jwtcertificates,omitempty" mapstructure:"JWTCertificates,omitempty"`

	// The SSH certificate authority used by the namespace.
	SSHCA *string `json:"SSHCA,omitempty" msgpack:"SSHCA,omitempty" bson:"sshca,omitempty" mapstructure:"SSHCA,omitempty"`

	// If `true`, an SSH certificate authority (CA) will be generated for the
	// namespace. This CA
	// can be deployed in SSH server to validate SSH certificates issued by the
	// platform.
	SSHCAEnabled *bool `json:"SSHCAEnabled,omitempty" msgpack:"SSHCAEnabled,omitempty" bson:"sshcaenabled,omitempty" mapstructure:"SSHCAEnabled,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// AssociatedLocalCAID holds the remote ID of the certificate authority to use.
	AssociatedLocalCAID *string `json:"-" msgpack:"-" bson:"associatedlocalcaid,omitempty" mapstructure:"-,omitempty"`

	// The remote ID of the SSH certificate authority to use.
	AssociatedSSHCAID *string `json:"associatedSSHCAID,omitempty" msgpack:"associatedSSHCAID,omitempty" bson:"associatedsshcaid,omitempty" mapstructure:"associatedSSHCAID,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Defines if the namespace should inherit its parent zone. If this property is set
	// to `false`,
	// the `zoning` property will be ignored and the namespace will have the same zone
	// as its parent.
	CustomZoning *bool `json:"customZoning,omitempty" msgpack:"customZoning,omitempty" bson:"customzoning,omitempty" mapstructure:"customZoning,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// The certificate authority used by this namespace.
	LocalCA *string `json:"localCA,omitempty" msgpack:"localCA,omitempty" bson:"localca,omitempty" mapstructure:"localCA,omitempty"`

	// Defines if the namespace should use a local certificate
	// authority (CA). Switching it off and on again will regenerate a new CA.
	LocalCAEnabled *bool `json:"localCAEnabled,omitempty" msgpack:"localCAEnabled,omitempty" bson:"localcaenabled,omitempty" mapstructure:"localCAEnabled,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// The name of the namespace.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// List of tags that will be added to every `or` clause of all network access
	// policies in the namespace and its children.
	NetworkAccessPolicyTags *[]string `json:"networkAccessPolicyTags,omitempty" msgpack:"networkAccessPolicyTags,omitempty" bson:"networkaccesspolicytags,omitempty" mapstructure:"networkAccessPolicyTags,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Determines the length of validity of certificates issued in this namespace using
	// [Golang duration syntax](https://golang.org/pkg/time/#example_Duration). Default
	// value is `1h`.
	ServiceCertificateValidity *string `json:"serviceCertificateValidity,omitempty" msgpack:"serviceCertificateValidity,omitempty" bson:"servicecertificatevalidity,omitempty" mapstructure:"serviceCertificateValidity,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	// Defines what zone the namespace should live in.
	Zoning *int `json:"zoning,omitempty" msgpack:"zoning,omitempty" bson:"zoning,omitempty" mapstructure:"zoning,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseNamespace returns a new  SparseNamespace.
func NewSparseNamespace() *SparseNamespace {
	return &SparseNamespace{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseNamespace) Identity() elemental.Identity {

	return NamespaceIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseNamespace) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseNamespace) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseNamespace) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseNamespace) ToPlain() elemental.PlainIdentifiable {

	out := NewNamespace()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.JWTCertificateType != nil {
		out.JWTCertificateType = *o.JWTCertificateType
	}
	if o.JWTCertificates != nil {
		out.JWTCertificates = *o.JWTCertificates
	}
	if o.SSHCA != nil {
		out.SSHCA = *o.SSHCA
	}
	if o.SSHCAEnabled != nil {
		out.SSHCAEnabled = *o.SSHCAEnabled
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedLocalCAID != nil {
		out.AssociatedLocalCAID = *o.AssociatedLocalCAID
	}
	if o.AssociatedSSHCAID != nil {
		out.AssociatedSSHCAID = *o.AssociatedSSHCAID
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.CustomZoning != nil {
		out.CustomZoning = *o.CustomZoning
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.LocalCA != nil {
		out.LocalCA = *o.LocalCA
	}
	if o.LocalCAEnabled != nil {
		out.LocalCAEnabled = *o.LocalCAEnabled
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NetworkAccessPolicyTags != nil {
		out.NetworkAccessPolicyTags = *o.NetworkAccessPolicyTags
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.ServiceCertificateValidity != nil {
		out.ServiceCertificateValidity = *o.ServiceCertificateValidity
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}
	if o.Zoning != nil {
		out.Zoning = *o.Zoning
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseNamespace) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseNamespace) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseNamespace) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseNamespace) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseNamespace) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseNamespace) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseNamespace) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseNamespace) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseNamespace) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseNamespace) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseNamespace) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseNamespace) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseNamespace) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseNamespace) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseNamespace) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseNamespace) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseNamespace) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseNamespace) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseNamespace) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseNamespace) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseNamespace) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseNamespace) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseNamespace) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseNamespace) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseNamespace) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseNamespace) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseNamespace) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseNamespace) SetZone(zone int) {

	o.Zone = &zone
}

// GetZoning returns the Zoning of the receiver.
func (o *SparseNamespace) GetZoning() int {

	return *o.Zoning
}

// SetZoning sets the property Zoning of the receiver using the address of the given value.
func (o *SparseNamespace) SetZoning(zoning int) {

	o.Zoning = &zoning
}

// DeepCopy returns a deep copy if the SparseNamespace.
func (o *SparseNamespace) DeepCopy() *SparseNamespace {

	if o == nil {
		return nil
	}

	out := &SparseNamespace{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseNamespace.
func (o *SparseNamespace) DeepCopyInto(out *SparseNamespace) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseNamespace: %s", err))
	}

	*out = *target.(*SparseNamespace)
}
