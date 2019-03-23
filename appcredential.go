package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AppCredentialIdentity represents the Identity of the object.
var AppCredentialIdentity = elemental.Identity{
	Name:     "appcredential",
	Category: "appcredentials",
	Package:  "cactuar",
	Private:  false,
}

// AppCredentialsList represents a list of AppCredentials
type AppCredentialsList []*AppCredential

// Identity returns the identity of the objects in the list.
func (o AppCredentialsList) Identity() elemental.Identity {

	return AppCredentialIdentity
}

// Copy returns a pointer to a copy the AppCredentialsList.
func (o AppCredentialsList) Copy() elemental.Identifiables {

	copy := append(AppCredentialsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AppCredentialsList.
func (o AppCredentialsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AppCredentialsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AppCredential))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AppCredentialsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AppCredentialsList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the AppCredentialsList converted to SparseAppCredentialsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AppCredentialsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o AppCredentialsList) Version() int {

	return 1
}

// AppCredential represents the model of a appcredential
type AppCredential struct {
	// CSR contains the Certificate Signing Request as a PEM encoded string. It can
	// only be set during a renew.
	//
	// - The CN **MUST** be `+"`"+`app:credential:<appcred-id>:<appcred-name>`+"`"+`
	// - The O **MUST** be the namespace of the appcred
	//
	// If you send anything else, the signing request will be rejected.
	CSR string `json:"CSR" bson:"-" mapstructure:"CSR,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// If set, the it will only be valid if the request comes from one
	// the declared subnets.
	AuthorizedSubnets []string `json:"authorizedSubnets" bson:"authorizedsubnets" mapstructure:"authorizedSubnets,omitempty"`

	// The string representation of the Certificate used by the application.
	Certificate string `json:"certificate" bson:"certificate" mapstructure:"certificate,omitempty"`

	// Link to the certificate created for this application.
	CertificateSN string `json:"-" bson:"certificatesn" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// The credential data.
	Credentials *Credential `json:"credentials" bson:"-" mapstructure:"credentials,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Disabled defines if the propert is disabled.
	Disabled bool `json:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// The email address that will receive a copy of the application credentials.
	Email string `json:"email" bson:"email" mapstructure:"email,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Contains the ID of the parent Appcred if this is a derived appcred.
	ParentIDs []string `json:"parentIDs" bson:"parentids" mapstructure:"parentIDs,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// List of roles to give the credentials.
	Roles []string `json:"roles" bson:"roles" mapstructure:"roles,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewAppCredential returns a new *AppCredential
func NewAppCredential() *AppCredential {

	return &AppCredential{
		ModelVersion:      1,
		Mutex:             &sync.Mutex{},
		AssociatedTags:    []string{},
		Annotations:       map[string][]string{},
		AuthorizedSubnets: []string{},
		NormalizedTags:    []string{},
		Metadata:          []string{},
		ParentIDs:         []string{},
		Roles:             []string{},
	}
}

// Identity returns the Identity of the object.
func (o *AppCredential) Identity() elemental.Identity {

	return AppCredentialIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AppCredential) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AppCredential) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *AppCredential) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *AppCredential) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// Doc returns the documentation for the object
func (o *AppCredential) Doc() string {
	return `Create a credential for an application.`
}

func (o *AppCredential) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *AppCredential) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *AppCredential) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *AppCredential) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *AppCredential) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *AppCredential) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *AppCredential) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *AppCredential) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *AppCredential) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *AppCredential) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *AppCredential) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetMetadata returns the Metadata of the receiver.
func (o *AppCredential) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *AppCredential) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *AppCredential) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *AppCredential) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *AppCredential) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *AppCredential) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *AppCredential) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *AppCredential) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *AppCredential) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *AppCredential) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *AppCredential) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *AppCredential) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *AppCredential) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *AppCredential) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *AppCredential) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *AppCredential) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AppCredential) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAppCredential{
			CSR:               &o.CSR,
			ID:                &o.ID,
			Annotations:       &o.Annotations,
			AssociatedTags:    &o.AssociatedTags,
			AuthorizedSubnets: &o.AuthorizedSubnets,
			Certificate:       &o.Certificate,
			CertificateSN:     &o.CertificateSN,
			CreateTime:        &o.CreateTime,
			Credentials:       &o.Credentials,
			Description:       &o.Description,
			Disabled:          &o.Disabled,
			Email:             &o.Email,
			Metadata:          &o.Metadata,
			Name:              &o.Name,
			Namespace:         &o.Namespace,
			NormalizedTags:    &o.NormalizedTags,
			ParentIDs:         &o.ParentIDs,
			Protected:         &o.Protected,
			Roles:             &o.Roles,
			UpdateTime:        &o.UpdateTime,
			ZHash:             &o.ZHash,
			Zone:              &o.Zone,
		}
	}

	sp := &SparseAppCredential{}
	for _, f := range fields {
		switch f {
		case "CSR":
			sp.CSR = &(o.CSR)
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "authorizedSubnets":
			sp.AuthorizedSubnets = &(o.AuthorizedSubnets)
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "certificateSN":
			sp.CertificateSN = &(o.CertificateSN)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "credentials":
			sp.Credentials = &(o.Credentials)
		case "description":
			sp.Description = &(o.Description)
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "email":
			sp.Email = &(o.Email)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "parentIDs":
			sp.ParentIDs = &(o.ParentIDs)
		case "protected":
			sp.Protected = &(o.Protected)
		case "roles":
			sp.Roles = &(o.Roles)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAppCredential to the object.
func (o *AppCredential) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAppCredential)
	if so.CSR != nil {
		o.CSR = *so.CSR
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.AuthorizedSubnets != nil {
		o.AuthorizedSubnets = *so.AuthorizedSubnets
	}
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.CertificateSN != nil {
		o.CertificateSN = *so.CertificateSN
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Credentials != nil {
		o.Credentials = *so.Credentials
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Disabled != nil {
		o.Disabled = *so.Disabled
	}
	if so.Email != nil {
		o.Email = *so.Email
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
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.ParentIDs != nil {
		o.ParentIDs = *so.ParentIDs
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Roles != nil {
		o.Roles = *so.Roles
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
}

// DeepCopy returns a deep copy if the AppCredential.
func (o *AppCredential) DeepCopy() *AppCredential {

	if o == nil {
		return nil
	}

	out := &AppCredential{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AppCredential.
func (o *AppCredential) DeepCopyInto(out *AppCredential) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AppCredential: %s", err))
	}

	*out = *target.(*AppCredential)
}

// Validate valides the current information stored into the structure.
func (o *AppCredential) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateOptionalNetworkList("authorizedSubnets", o.AuthorizedSubnets); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("roles", o.Roles); err != nil {
		requiredErrors = append(requiredErrors, err)
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
func (*AppCredential) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AppCredentialAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AppCredentialLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AppCredential) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AppCredentialAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AppCredential) ValueForAttribute(name string) interface{} {

	switch name {
	case "CSR":
		return o.CSR
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "authorizedSubnets":
		return o.AuthorizedSubnets
	case "certificate":
		return o.Certificate
	case "certificateSN":
		return o.CertificateSN
	case "createTime":
		return o.CreateTime
	case "credentials":
		return o.Credentials
	case "description":
		return o.Description
	case "disabled":
		return o.Disabled
	case "email":
		return o.Email
	case "metadata":
		return o.Metadata
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "parentIDs":
		return o.ParentIDs
	case "protected":
		return o.Protected
	case "roles":
		return o.Roles
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// AppCredentialAttributesMap represents the map of attribute for AppCredential.
var AppCredentialAttributesMap = map[string]elemental.AttributeSpecification{
	"CSR": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CSR",
		Description: `CSR contains the Certificate Signing Request as a PEM encoded string. It can
only be set during a renew.

- The CN **MUST** be ` + "`" + `app:credential:<appcred-id>:<appcred-name>` + "`" + `
- The O **MUST** be the namespace of the appcred

If you send anything else, the signing request will be rejected.`,
		Exposed: true,
		Name:    "CSR",
		Type:    "string",
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
	"AuthorizedSubnets": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizedSubnets",
		Description: `If set, the it will only be valid if the request comes from one
the declared subnets.`,
		Exposed: true,
		Name:    "authorizedSubnets",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"Certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		Description:    `The string representation of the Certificate used by the application.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CertificateSN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateSN",
		Description:    `Link to the certificate created for this application.`,
		Name:           "certificateSN",
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
	"Credentials": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Credentials",
		Description:    `The credential data.`,
		Exposed:        true,
		Name:           "credentials",
		Orderable:      true,
		ReadOnly:       true,
		SubType:        "credential",
		Transient:      true,
		Type:           "ref",
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
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		Description:    `The email address that will receive a copy of the application credentials.`,
		Exposed:        true,
		Name:           "email",
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
		SubType:    "string",
		Type:       "list",
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
	"ParentIDs": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentIDs",
		Description:    `Contains the ID of the parent Appcred if this is a derived appcred.`,
		Exposed:        true,
		Name:           "parentIDs",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
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
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Roles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Roles",
		Description:    `List of roles to give the credentials.`,
		Exposed:        true,
		Name:           "roles",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// AppCredentialLowerCaseAttributesMap represents the map of attribute for AppCredential.
var AppCredentialLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"csr": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CSR",
		Description: `CSR contains the Certificate Signing Request as a PEM encoded string. It can
only be set during a renew.

- The CN **MUST** be ` + "`" + `app:credential:<appcred-id>:<appcred-name>` + "`" + `
- The O **MUST** be the namespace of the appcred

If you send anything else, the signing request will be rejected.`,
		Exposed: true,
		Name:    "CSR",
		Type:    "string",
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
	"authorizedsubnets": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizedSubnets",
		Description: `If set, the it will only be valid if the request comes from one
the declared subnets.`,
		Exposed: true,
		Name:    "authorizedSubnets",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		Description:    `The string representation of the Certificate used by the application.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"certificatesn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateSN",
		Description:    `Link to the certificate created for this application.`,
		Name:           "certificateSN",
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
	"credentials": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Credentials",
		Description:    `The credential data.`,
		Exposed:        true,
		Name:           "credentials",
		Orderable:      true,
		ReadOnly:       true,
		SubType:        "credential",
		Transient:      true,
		Type:           "ref",
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
	"disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		Description:    `The email address that will receive a copy of the application credentials.`,
		Exposed:        true,
		Name:           "email",
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
		SubType:    "string",
		Type:       "list",
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
	"parentids": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentIDs",
		Description:    `Contains the ID of the parent Appcred if this is a derived appcred.`,
		Exposed:        true,
		Name:           "parentIDs",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
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
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"roles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Roles",
		Description:    `List of roles to give the credentials.`,
		Exposed:        true,
		Name:           "roles",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// SparseAppCredentialsList represents a list of SparseAppCredentials
type SparseAppCredentialsList []*SparseAppCredential

// Identity returns the identity of the objects in the list.
func (o SparseAppCredentialsList) Identity() elemental.Identity {

	return AppCredentialIdentity
}

// Copy returns a pointer to a copy the SparseAppCredentialsList.
func (o SparseAppCredentialsList) Copy() elemental.Identifiables {

	copy := append(SparseAppCredentialsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAppCredentialsList.
func (o SparseAppCredentialsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAppCredentialsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAppCredential))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAppCredentialsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAppCredentialsList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparseAppCredentialsList converted to AppCredentialsList.
func (o SparseAppCredentialsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAppCredentialsList) Version() int {

	return 1
}

// SparseAppCredential represents the sparse version of a appcredential.
type SparseAppCredential struct {
	// CSR contains the Certificate Signing Request as a PEM encoded string. It can
	// only be set during a renew.
	//
	// - The CN **MUST** be `+"`"+`app:credential:<appcred-id>:<appcred-name>`+"`"+`
	// - The O **MUST** be the namespace of the appcred
	//
	// If you send anything else, the signing request will be rejected.
	CSR *string `json:"CSR,omitempty" bson:"-" mapstructure:"CSR,omitempty"`

	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// If set, the it will only be valid if the request comes from one
	// the declared subnets.
	AuthorizedSubnets *[]string `json:"authorizedSubnets,omitempty" bson:"authorizedsubnets" mapstructure:"authorizedSubnets,omitempty"`

	// The string representation of the Certificate used by the application.
	Certificate *string `json:"certificate,omitempty" bson:"certificate" mapstructure:"certificate,omitempty"`

	// Link to the certificate created for this application.
	CertificateSN *string `json:"-" bson:"certificatesn" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// The credential data.
	Credentials **Credential `json:"credentials,omitempty" bson:"-" mapstructure:"credentials,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// Disabled defines if the propert is disabled.
	Disabled *bool `json:"disabled,omitempty" bson:"disabled" mapstructure:"disabled,omitempty"`

	// The email address that will receive a copy of the application credentials.
	Email *string `json:"email,omitempty" bson:"email" mapstructure:"email,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Contains the ID of the parent Appcred if this is a derived appcred.
	ParentIDs *[]string `json:"parentIDs,omitempty" bson:"parentids" mapstructure:"parentIDs,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected" mapstructure:"protected,omitempty"`

	// List of roles to give the credentials.
	Roles *[]string `json:"roles,omitempty" bson:"roles" mapstructure:"roles,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSparseAppCredential returns a new  SparseAppCredential.
func NewSparseAppCredential() *SparseAppCredential {
	return &SparseAppCredential{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAppCredential) Identity() elemental.Identity {

	return AppCredentialIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAppCredential) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAppCredential) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseAppCredential) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAppCredential) ToPlain() elemental.PlainIdentifiable {

	out := NewAppCredential()
	if o.CSR != nil {
		out.CSR = *o.CSR
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.AuthorizedSubnets != nil {
		out.AuthorizedSubnets = *o.AuthorizedSubnets
	}
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.CertificateSN != nil {
		out.CertificateSN = *o.CertificateSN
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Credentials != nil {
		out.Credentials = *o.Credentials
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Disabled != nil {
		out.Disabled = *o.Disabled
	}
	if o.Email != nil {
		out.Email = *o.Email
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
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.ParentIDs != nil {
		out.ParentIDs = *o.ParentIDs
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Roles != nil {
		out.Roles = *o.Roles
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

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseAppCredential) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseAppCredential) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseAppCredential) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseAppCredential) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseAppCredential) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseAppCredential) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseAppCredential) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseAppCredential) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseAppCredential) GetDisabled() bool {

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseAppCredential) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseAppCredential) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseAppCredential) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseAppCredential) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseAppCredential) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseAppCredential) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseAppCredential) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseAppCredential) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseAppCredential) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseAppCredential) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseAppCredential) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseAppCredential) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseAppCredential) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAppCredential) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAppCredential) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAppCredential) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAppCredential) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAppCredential.
func (o *SparseAppCredential) DeepCopy() *SparseAppCredential {

	if o == nil {
		return nil
	}

	out := &SparseAppCredential{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAppCredential.
func (o *SparseAppCredential) DeepCopyInto(out *SparseAppCredential) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAppCredential: %s", err))
	}

	*out = *target.(*SparseAppCredential)
}
