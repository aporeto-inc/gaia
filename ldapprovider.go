package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// LDAPProviderConnSecurityProtocolValue represents the possible values for attribute "ConnSecurityProtocol".
type LDAPProviderConnSecurityProtocolValue string

const (
	// LDAPProviderConnSecurityProtocolInbandTLS represents the value InbandTLS.
	LDAPProviderConnSecurityProtocolInbandTLS LDAPProviderConnSecurityProtocolValue = "InbandTLS"

	// LDAPProviderConnSecurityProtocolTLS represents the value TLS.
	LDAPProviderConnSecurityProtocolTLS LDAPProviderConnSecurityProtocolValue = "TLS"
)

// LDAPProviderIdentity represents the Identity of the object.
var LDAPProviderIdentity = elemental.Identity{
	Name:     "ldapprovider",
	Category: "ldapproviders",
	Package:  "vince",
	Private:  false,
}

// LDAPProvidersList represents a list of LDAPProviders
type LDAPProvidersList []*LDAPProvider

// Identity returns the identity of the objects in the list.
func (o LDAPProvidersList) Identity() elemental.Identity {

	return LDAPProviderIdentity
}

// Copy returns a pointer to a copy the LDAPProvidersList.
func (o LDAPProvidersList) Copy() elemental.Identifiables {

	copy := append(LDAPProvidersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the LDAPProvidersList.
func (o LDAPProvidersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(LDAPProvidersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*LDAPProvider))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o LDAPProvidersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o LDAPProvidersList) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// ToSparse returns the LDAPProvidersList converted to SparseLDAPProvidersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o LDAPProvidersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseLDAPProvidersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseLDAPProvider)
	}

	return out
}

// Version returns the version of the content.
func (o LDAPProvidersList) Version() int {

	return 1
}

// LDAPProvider represents the model of a ldapprovider
type LDAPProvider struct {
	// Address holds the account authentication account's private LDAP server.
	Address string `json:"Address" msgpack:"Address" bson:"address" mapstructure:"Address,omitempty"`

	// BaseDN holds the base DN to use to LDAP queries.
	BaseDN string `json:"BaseDN" msgpack:"BaseDN" bson:"basedn" mapstructure:"BaseDN,omitempty"`

	// BindDN holds the account's internal LDAP bind DN for querying auth.
	BindDN string `json:"BindDN" msgpack:"BindDN" bson:"binddn" mapstructure:"BindDN,omitempty"`

	// BindPassword holds the password to the LDAP bind DN.
	BindPassword string `json:"BindPassword" msgpack:"BindPassword" bson:"bindpassword" mapstructure:"BindPassword,omitempty"`

	// BindSearchFilter holds filter to be used to uniquely search a user. For
	// Windows based systems, value may be `+"`"+`sAMAccountName={USERNAME}`+"`"+`. For Linux and
	// other systems, value may be `+"`"+`uid={USERNAME}`+"`"+`.
	BindSearchFilter string `json:"BindSearchFilter" msgpack:"BindSearchFilter" bson:"bindsearchfilter" mapstructure:"BindSearchFilter,omitempty"`

	// CertificateAuthority contains the optional certificate authority that will
	// be used to connect to the LDAP server. It is not needed if the TLS certificate
	// of the LDAP is issued from a public truster CA.
	CertificateAuthority string `json:"CertificateAuthority" msgpack:"CertificateAuthority" bson:"certificateauthority" mapstructure:"CertificateAuthority,omitempty"`

	// ConnSecurityProtocol holds the connection type for the LDAP provider.
	ConnSecurityProtocol LDAPProviderConnSecurityProtocolValue `json:"ConnSecurityProtocol" msgpack:"ConnSecurityProtocol" bson:"connsecurityprotocol" mapstructure:"ConnSecurityProtocol,omitempty"`

	// Enabled determines if the account uses it's own LDAP for authentication.
	Enabled bool `json:"Enabled" msgpack:"Enabled" bson:"enabled" mapstructure:"Enabled,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// IgnoredKeys holds a list of keys that must not be imported into Aporeto
	// authorization system.
	IgnoredKeys []string `json:"IgnoredKeys" msgpack:"IgnoredKeys" bson:"ignoredkeys" mapstructure:"IgnoredKeys,omitempty"`

	// SubjectKey holds key to be used to populate the subject. If you want to
	// use the user as a subject, for Windows based systems you may use
	// 'sAMAccountName' and for Linux and other systems, value may be 'uid'. You can
	// also use any alternate key.
	SubjectKey string `json:"SubjectKey" msgpack:"SubjectKey" bson:"subjectkey" mapstructure:"SubjectKey,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewLDAPProvider returns a new *LDAPProvider
func NewLDAPProvider() *LDAPProvider {

	return &LDAPProvider{
		ModelVersion:         1,
		SubjectKey:           "uid",
		Annotations:          map[string][]string{},
		AssociatedTags:       []string{},
		BindSearchFilter:     "uid={USERNAME}",
		ConnSecurityProtocol: LDAPProviderConnSecurityProtocolInbandTLS,
		IgnoredKeys:          []string{},
		NormalizedTags:       []string{},
	}
}

// Identity returns the Identity of the object.
func (o *LDAPProvider) Identity() elemental.Identity {

	return LDAPProviderIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *LDAPProvider) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *LDAPProvider) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *LDAPProvider) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *LDAPProvider) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// Doc returns the documentation for the object
func (o *LDAPProvider) Doc() string {

	return `Allows to declare a generic LDAP provider that can be used in exchange
for a Midgard token.`
}

func (o *LDAPProvider) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *LDAPProvider) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *LDAPProvider) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *LDAPProvider) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *LDAPProvider) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *LDAPProvider) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *LDAPProvider) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *LDAPProvider) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *LDAPProvider) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *LDAPProvider) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *LDAPProvider) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *LDAPProvider) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *LDAPProvider) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *LDAPProvider) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *LDAPProvider) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *LDAPProvider) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *LDAPProvider) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *LDAPProvider) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *LDAPProvider) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *LDAPProvider) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseLDAPProvider{
			Address:              &o.Address,
			BaseDN:               &o.BaseDN,
			BindDN:               &o.BindDN,
			BindPassword:         &o.BindPassword,
			BindSearchFilter:     &o.BindSearchFilter,
			CertificateAuthority: &o.CertificateAuthority,
			ConnSecurityProtocol: &o.ConnSecurityProtocol,
			Enabled:              &o.Enabled,
			ID:                   &o.ID,
			IgnoredKeys:          &o.IgnoredKeys,
			SubjectKey:           &o.SubjectKey,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Protected:            &o.Protected,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
		}
	}

	sp := &SparseLDAPProvider{}
	for _, f := range fields {
		switch f {
		case "Address":
			sp.Address = &(o.Address)
		case "BaseDN":
			sp.BaseDN = &(o.BaseDN)
		case "BindDN":
			sp.BindDN = &(o.BindDN)
		case "BindPassword":
			sp.BindPassword = &(o.BindPassword)
		case "BindSearchFilter":
			sp.BindSearchFilter = &(o.BindSearchFilter)
		case "CertificateAuthority":
			sp.CertificateAuthority = &(o.CertificateAuthority)
		case "ConnSecurityProtocol":
			sp.ConnSecurityProtocol = &(o.ConnSecurityProtocol)
		case "Enabled":
			sp.Enabled = &(o.Enabled)
		case "ID":
			sp.ID = &(o.ID)
		case "IgnoredKeys":
			sp.IgnoredKeys = &(o.IgnoredKeys)
		case "SubjectKey":
			sp.SubjectKey = &(o.SubjectKey)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseLDAPProvider to the object.
func (o *LDAPProvider) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseLDAPProvider)
	if so.Address != nil {
		o.Address = *so.Address
	}
	if so.BaseDN != nil {
		o.BaseDN = *so.BaseDN
	}
	if so.BindDN != nil {
		o.BindDN = *so.BindDN
	}
	if so.BindPassword != nil {
		o.BindPassword = *so.BindPassword
	}
	if so.BindSearchFilter != nil {
		o.BindSearchFilter = *so.BindSearchFilter
	}
	if so.CertificateAuthority != nil {
		o.CertificateAuthority = *so.CertificateAuthority
	}
	if so.ConnSecurityProtocol != nil {
		o.ConnSecurityProtocol = *so.ConnSecurityProtocol
	}
	if so.Enabled != nil {
		o.Enabled = *so.Enabled
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.IgnoredKeys != nil {
		o.IgnoredKeys = *so.IgnoredKeys
	}
	if so.SubjectKey != nil {
		o.SubjectKey = *so.SubjectKey
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
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
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the LDAPProvider.
func (o *LDAPProvider) DeepCopy() *LDAPProvider {

	if o == nil {
		return nil
	}

	out := &LDAPProvider{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *LDAPProvider.
func (o *LDAPProvider) DeepCopyInto(out *LDAPProvider) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy LDAPProvider: %s", err))
	}

	*out = *target.(*LDAPProvider)
}

// Validate valides the current information stored into the structure.
func (o *LDAPProvider) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("ConnSecurityProtocol", string(o.ConnSecurityProtocol), []string{"TLS", "InbandTLS"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
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
func (*LDAPProvider) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := LDAPProviderAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return LDAPProviderLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*LDAPProvider) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return LDAPProviderAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *LDAPProvider) ValueForAttribute(name string) interface{} {

	switch name {
	case "Address":
		return o.Address
	case "BaseDN":
		return o.BaseDN
	case "BindDN":
		return o.BindDN
	case "BindPassword":
		return o.BindPassword
	case "BindSearchFilter":
		return o.BindSearchFilter
	case "CertificateAuthority":
		return o.CertificateAuthority
	case "ConnSecurityProtocol":
		return o.ConnSecurityProtocol
	case "Enabled":
		return o.Enabled
	case "ID":
		return o.ID
	case "IgnoredKeys":
		return o.IgnoredKeys
	case "SubjectKey":
		return o.SubjectKey
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	}

	return nil
}

// LDAPProviderAttributesMap represents the map of attribute for LDAPProvider.
var LDAPProviderAttributesMap = map[string]elemental.AttributeSpecification{
	"Address": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Address",
		Description:    `Address holds the account authentication account's private LDAP server.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "Address",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"BaseDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BaseDN",
		Description:    `BaseDN holds the base DN to use to LDAP queries.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "BaseDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"BindDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BindDN",
		Description:    `BindDN holds the account's internal LDAP bind DN for querying auth.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "BindDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"BindPassword": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BindPassword",
		Description:    `BindPassword holds the password to the LDAP bind DN.`,
		Exposed:        true,
		Name:           "BindPassword",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"BindSearchFilter": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BindSearchFilter",
		DefaultValue:   "uid={USERNAME}",
		Description: `BindSearchFilter holds filter to be used to uniquely search a user. For
Windows based systems, value may be ` + "`" + `sAMAccountName={USERNAME}` + "`" + `. For Linux and
other systems, value may be ` + "`" + `uid={USERNAME}` + "`" + `.`,
		Exposed:   true,
		Name:      "BindSearchFilter",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"CertificateAuthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateAuthority",
		Description: `CertificateAuthority contains the optional certificate authority that will
be used to connect to the LDAP server. It is not needed if the TLS certificate
of the LDAP is issued from a public truster CA.`,
		Exposed:   true,
		Name:      "CertificateAuthority",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"ConnSecurityProtocol": elemental.AttributeSpecification{
		AllowedChoices: []string{"TLS", "InbandTLS"},
		ConvertedName:  "ConnSecurityProtocol",
		DefaultValue:   LDAPProviderConnSecurityProtocolInbandTLS,
		Description:    `ConnSecurityProtocol holds the connection type for the LDAP provider.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "ConnSecurityProtocol",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"Enabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Enabled",
		Description:    `Enabled determines if the account uses it's own LDAP for authentication.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "Enabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
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
	"IgnoredKeys": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IgnoredKeys",
		Description: `IgnoredKeys holds a list of keys that must not be imported into Aporeto
authorization system.`,
		Exposed:   true,
		Name:      "IgnoredKeys",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"SubjectKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SubjectKey",
		DefaultValue:   "uid",
		Description: `SubjectKey holds key to be used to populate the subject. If you want to
use the user as a subject, for Windows based systems you may use
'sAMAccountName' and for Linux and other systems, value may be 'uid'. You can
also use any alternate key.`,
		Exposed:   true,
		Name:      "SubjectKey",
		Orderable: true,
		Stored:    true,
		Type:      "string",
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

// LDAPProviderLowerCaseAttributesMap represents the map of attribute for LDAPProvider.
var LDAPProviderLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"address": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Address",
		Description:    `Address holds the account authentication account's private LDAP server.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "Address",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"basedn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BaseDN",
		Description:    `BaseDN holds the base DN to use to LDAP queries.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "BaseDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"binddn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BindDN",
		Description:    `BindDN holds the account's internal LDAP bind DN for querying auth.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "BindDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"bindpassword": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BindPassword",
		Description:    `BindPassword holds the password to the LDAP bind DN.`,
		Exposed:        true,
		Name:           "BindPassword",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"bindsearchfilter": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BindSearchFilter",
		DefaultValue:   "uid={USERNAME}",
		Description: `BindSearchFilter holds filter to be used to uniquely search a user. For
Windows based systems, value may be ` + "`" + `sAMAccountName={USERNAME}` + "`" + `. For Linux and
other systems, value may be ` + "`" + `uid={USERNAME}` + "`" + `.`,
		Exposed:   true,
		Name:      "BindSearchFilter",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"certificateauthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateAuthority",
		Description: `CertificateAuthority contains the optional certificate authority that will
be used to connect to the LDAP server. It is not needed if the TLS certificate
of the LDAP is issued from a public truster CA.`,
		Exposed:   true,
		Name:      "CertificateAuthority",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"connsecurityprotocol": elemental.AttributeSpecification{
		AllowedChoices: []string{"TLS", "InbandTLS"},
		ConvertedName:  "ConnSecurityProtocol",
		DefaultValue:   LDAPProviderConnSecurityProtocolInbandTLS,
		Description:    `ConnSecurityProtocol holds the connection type for the LDAP provider.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "ConnSecurityProtocol",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"enabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Enabled",
		Description:    `Enabled determines if the account uses it's own LDAP for authentication.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "Enabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
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
	"ignoredkeys": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IgnoredKeys",
		Description: `IgnoredKeys holds a list of keys that must not be imported into Aporeto
authorization system.`,
		Exposed:   true,
		Name:      "IgnoredKeys",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"subjectkey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SubjectKey",
		DefaultValue:   "uid",
		Description: `SubjectKey holds key to be used to populate the subject. If you want to
use the user as a subject, for Windows based systems you may use
'sAMAccountName' and for Linux and other systems, value may be 'uid'. You can
also use any alternate key.`,
		Exposed:   true,
		Name:      "SubjectKey",
		Orderable: true,
		Stored:    true,
		Type:      "string",
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

// SparseLDAPProvidersList represents a list of SparseLDAPProviders
type SparseLDAPProvidersList []*SparseLDAPProvider

// Identity returns the identity of the objects in the list.
func (o SparseLDAPProvidersList) Identity() elemental.Identity {

	return LDAPProviderIdentity
}

// Copy returns a pointer to a copy the SparseLDAPProvidersList.
func (o SparseLDAPProvidersList) Copy() elemental.Identifiables {

	copy := append(SparseLDAPProvidersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseLDAPProvidersList.
func (o SparseLDAPProvidersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseLDAPProvidersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseLDAPProvider))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseLDAPProvidersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseLDAPProvidersList) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// ToPlain returns the SparseLDAPProvidersList converted to LDAPProvidersList.
func (o SparseLDAPProvidersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseLDAPProvidersList) Version() int {

	return 1
}

// SparseLDAPProvider represents the sparse version of a ldapprovider.
type SparseLDAPProvider struct {
	// Address holds the account authentication account's private LDAP server.
	Address *string `json:"Address,omitempty" msgpack:"Address,omitempty" bson:"address,omitempty" mapstructure:"Address,omitempty"`

	// BaseDN holds the base DN to use to LDAP queries.
	BaseDN *string `json:"BaseDN,omitempty" msgpack:"BaseDN,omitempty" bson:"basedn,omitempty" mapstructure:"BaseDN,omitempty"`

	// BindDN holds the account's internal LDAP bind DN for querying auth.
	BindDN *string `json:"BindDN,omitempty" msgpack:"BindDN,omitempty" bson:"binddn,omitempty" mapstructure:"BindDN,omitempty"`

	// BindPassword holds the password to the LDAP bind DN.
	BindPassword *string `json:"BindPassword,omitempty" msgpack:"BindPassword,omitempty" bson:"bindpassword,omitempty" mapstructure:"BindPassword,omitempty"`

	// BindSearchFilter holds filter to be used to uniquely search a user. For
	// Windows based systems, value may be `+"`"+`sAMAccountName={USERNAME}`+"`"+`. For Linux and
	// other systems, value may be `+"`"+`uid={USERNAME}`+"`"+`.
	BindSearchFilter *string `json:"BindSearchFilter,omitempty" msgpack:"BindSearchFilter,omitempty" bson:"bindsearchfilter,omitempty" mapstructure:"BindSearchFilter,omitempty"`

	// CertificateAuthority contains the optional certificate authority that will
	// be used to connect to the LDAP server. It is not needed if the TLS certificate
	// of the LDAP is issued from a public truster CA.
	CertificateAuthority *string `json:"CertificateAuthority,omitempty" msgpack:"CertificateAuthority,omitempty" bson:"certificateauthority,omitempty" mapstructure:"CertificateAuthority,omitempty"`

	// ConnSecurityProtocol holds the connection type for the LDAP provider.
	ConnSecurityProtocol *LDAPProviderConnSecurityProtocolValue `json:"ConnSecurityProtocol,omitempty" msgpack:"ConnSecurityProtocol,omitempty" bson:"connsecurityprotocol,omitempty" mapstructure:"ConnSecurityProtocol,omitempty"`

	// Enabled determines if the account uses it's own LDAP for authentication.
	Enabled *bool `json:"Enabled,omitempty" msgpack:"Enabled,omitempty" bson:"enabled,omitempty" mapstructure:"Enabled,omitempty"`

	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// IgnoredKeys holds a list of keys that must not be imported into Aporeto
	// authorization system.
	IgnoredKeys *[]string `json:"IgnoredKeys,omitempty" msgpack:"IgnoredKeys,omitempty" bson:"ignoredkeys,omitempty" mapstructure:"IgnoredKeys,omitempty"`

	// SubjectKey holds key to be used to populate the subject. If you want to
	// use the user as a subject, for Windows based systems you may use
	// 'sAMAccountName' and for Linux and other systems, value may be 'uid'. You can
	// also use any alternate key.
	SubjectKey *string `json:"SubjectKey,omitempty" msgpack:"SubjectKey,omitempty" bson:"subjectkey,omitempty" mapstructure:"SubjectKey,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseLDAPProvider returns a new  SparseLDAPProvider.
func NewSparseLDAPProvider() *SparseLDAPProvider {
	return &SparseLDAPProvider{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseLDAPProvider) Identity() elemental.Identity {

	return LDAPProviderIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseLDAPProvider) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseLDAPProvider) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseLDAPProvider) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseLDAPProvider) ToPlain() elemental.PlainIdentifiable {

	out := NewLDAPProvider()
	if o.Address != nil {
		out.Address = *o.Address
	}
	if o.BaseDN != nil {
		out.BaseDN = *o.BaseDN
	}
	if o.BindDN != nil {
		out.BindDN = *o.BindDN
	}
	if o.BindPassword != nil {
		out.BindPassword = *o.BindPassword
	}
	if o.BindSearchFilter != nil {
		out.BindSearchFilter = *o.BindSearchFilter
	}
	if o.CertificateAuthority != nil {
		out.CertificateAuthority = *o.CertificateAuthority
	}
	if o.ConnSecurityProtocol != nil {
		out.ConnSecurityProtocol = *o.ConnSecurityProtocol
	}
	if o.Enabled != nil {
		out.Enabled = *o.Enabled
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.IgnoredKeys != nil {
		out.IgnoredKeys = *o.IgnoredKeys
	}
	if o.SubjectKey != nil {
		out.SubjectKey = *o.SubjectKey
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
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
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
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
func (o *SparseLDAPProvider) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseLDAPProvider) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseLDAPProvider) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseLDAPProvider) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseLDAPProvider) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseLDAPProvider) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseLDAPProvider) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseLDAPProvider) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseLDAPProvider) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseLDAPProvider.
func (o *SparseLDAPProvider) DeepCopy() *SparseLDAPProvider {

	if o == nil {
		return nil
	}

	out := &SparseLDAPProvider{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseLDAPProvider.
func (o *SparseLDAPProvider) DeepCopyInto(out *SparseLDAPProvider) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseLDAPProvider: %s", err))
	}

	*out = *target.(*SparseLDAPProvider)
}
