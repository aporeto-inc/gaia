package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ClaimsIdentity represents the Identity of the object.
var ClaimsIdentity = elemental.Identity{
	Name:     "claims",
	Category: "claims",
	Package:  "guy",
	Private:  false,
}

// ClaimsList represents a list of Claims
type ClaimsList []*Claims

// Identity returns the identity of the objects in the list.
func (o ClaimsList) Identity() elemental.Identity {

	return ClaimsIdentity
}

// Copy returns a pointer to a copy the ClaimsList.
func (o ClaimsList) Copy() elemental.Identifiables {

	copy := append(ClaimsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ClaimsList.
func (o ClaimsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ClaimsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Claims))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ClaimsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ClaimsList) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// ToSparse returns the ClaimsList converted to SparseClaimsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ClaimsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o ClaimsList) Version() int {

	return 1
}

// Claims represents the model of a claims
type Claims struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Content contains the raw JWT claims.
	Content map[string]string `json:"content" bson:"content" mapstructure:"content,omitempty"`

	// firstSeen contains the date of the first appearance of the claims.
	FirstSeen time.Time `json:"-" bson:"firstseen" mapstructure:"-,omitempty"`

	// XXH64 of the claims content. It will be used as ID. To compute a correct hash,
	// you must first clob Content as an string array in the form `+"`"+`key=value`+"`"+`, sort it
	// then apply the xxhash function.
	Hash string `json:"hash" bson:"-" mapstructure:"hash,omitempty"`

	// lastSeen contains the date of the last appearance of the claims.
	LastSeen time.Time `json:"-" bson:"lastseen" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewClaims returns a new *Claims
func NewClaims() *Claims {

	return &Claims{
		ModelVersion:   1,
		Mutex:          &sync.Mutex{},
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Content:        map[string]string{},
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *Claims) Identity() elemental.Identity {

	return ClaimsIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Claims) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Claims) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Claims) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Claims) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// Doc returns the documentation for the object
func (o *Claims) Doc() string {

	return `This API represents the claims that accessed a service.`
}

func (o *Claims) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Claims) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *Claims) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Claims) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *Claims) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetNamespace returns the Namespace of the receiver.
func (o *Claims) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Claims) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Claims) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *Claims) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Claims) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *Claims) SetProtected(protected bool) {

	o.Protected = protected
}

// GetZHash returns the ZHash of the receiver.
func (o *Claims) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Claims) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Claims) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Claims) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Claims) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseClaims{
			ID:             &o.ID,
			Annotations:    &o.Annotations,
			AssociatedTags: &o.AssociatedTags,
			Content:        &o.Content,
			FirstSeen:      &o.FirstSeen,
			Hash:           &o.Hash,
			LastSeen:       &o.LastSeen,
			Namespace:      &o.Namespace,
			NormalizedTags: &o.NormalizedTags,
			Protected:      &o.Protected,
			ZHash:          &o.ZHash,
			Zone:           &o.Zone,
		}
	}

	sp := &SparseClaims{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "content":
			sp.Content = &(o.Content)
		case "firstSeen":
			sp.FirstSeen = &(o.FirstSeen)
		case "hash":
			sp.Hash = &(o.Hash)
		case "lastSeen":
			sp.LastSeen = &(o.LastSeen)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseClaims to the object.
func (o *Claims) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseClaims)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.Content != nil {
		o.Content = *so.Content
	}
	if so.FirstSeen != nil {
		o.FirstSeen = *so.FirstSeen
	}
	if so.Hash != nil {
		o.Hash = *so.Hash
	}
	if so.LastSeen != nil {
		o.LastSeen = *so.LastSeen
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
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the Claims.
func (o *Claims) DeepCopy() *Claims {

	if o == nil {
		return nil
	}

	out := &Claims{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Claims.
func (o *Claims) DeepCopyInto(out *Claims) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Claims: %s", err))
	}

	*out = *target.(*Claims)
}

// Validate valides the current information stored into the structure.
func (o *Claims) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("hash", o.Hash); err != nil {
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
func (*Claims) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ClaimsAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ClaimsLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Claims) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ClaimsAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Claims) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "content":
		return o.Content
	case "firstSeen":
		return o.FirstSeen
	case "hash":
		return o.Hash
	case "lastSeen":
		return o.LastSeen
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// ClaimsAttributesMap represents the map of attribute for Claims.
var ClaimsAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		CreationOnly:   true,
		Description:    `Content contains the raw JWT claims.`,
		Exposed:        true,
		Name:           "content",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"FirstSeen": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "FirstSeen",
		Description:    `firstSeen contains the date of the first appearance of the claims.`,
		Name:           "firstSeen",
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"Hash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Hash",
		Description: `XXH64 of the claims content. It will be used as ID. To compute a correct hash,
you must first clob Content as an string array in the form ` + "`" + `key=value` + "`" + `, sort it
then apply the xxhash function.`,
		Exposed:  true,
		Name:     "hash",
		Required: true,
		Type:     "string",
	},
	"LastSeen": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastSeen",
		Description:    `lastSeen contains the date of the last appearance of the claims.`,
		Name:           "lastSeen",
		ReadOnly:       true,
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

// ClaimsLowerCaseAttributesMap represents the map of attribute for Claims.
var ClaimsLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		CreationOnly:   true,
		Description:    `Content contains the raw JWT claims.`,
		Exposed:        true,
		Name:           "content",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"firstseen": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "FirstSeen",
		Description:    `firstSeen contains the date of the first appearance of the claims.`,
		Name:           "firstSeen",
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"hash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Hash",
		Description: `XXH64 of the claims content. It will be used as ID. To compute a correct hash,
you must first clob Content as an string array in the form ` + "`" + `key=value` + "`" + `, sort it
then apply the xxhash function.`,
		Exposed:  true,
		Name:     "hash",
		Required: true,
		Type:     "string",
	},
	"lastseen": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastSeen",
		Description:    `lastSeen contains the date of the last appearance of the claims.`,
		Name:           "lastSeen",
		ReadOnly:       true,
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

// SparseClaimsList represents a list of SparseClaims
type SparseClaimsList []*SparseClaims

// Identity returns the identity of the objects in the list.
func (o SparseClaimsList) Identity() elemental.Identity {

	return ClaimsIdentity
}

// Copy returns a pointer to a copy the SparseClaimsList.
func (o SparseClaimsList) Copy() elemental.Identifiables {

	copy := append(SparseClaimsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseClaimsList.
func (o SparseClaimsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseClaimsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseClaims))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseClaimsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseClaimsList) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// ToPlain returns the SparseClaimsList converted to ClaimsList.
func (o SparseClaimsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseClaimsList) Version() int {

	return 1
}

// SparseClaims represents the sparse version of a claims.
type SparseClaims struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// Content contains the raw JWT claims.
	Content *map[string]string `json:"content,omitempty" bson:"content,omitempty" mapstructure:"content,omitempty"`

	// firstSeen contains the date of the first appearance of the claims.
	FirstSeen *time.Time `json:"-" bson:"firstseen,omitempty" mapstructure:"-,omitempty"`

	// XXH64 of the claims content. It will be used as ID. To compute a correct hash,
	// you must first clob Content as an string array in the form `+"`"+`key=value`+"`"+`, sort it
	// then apply the xxhash function.
	Hash *string `json:"hash,omitempty" bson:"-" mapstructure:"hash,omitempty"`

	// lastSeen contains the date of the last appearance of the claims.
	LastSeen *time.Time `json:"-" bson:"lastseen,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSparseClaims returns a new  SparseClaims.
func NewSparseClaims() *SparseClaims {
	return &SparseClaims{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseClaims) Identity() elemental.Identity {

	return ClaimsIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseClaims) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseClaims) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseClaims) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseClaims) ToPlain() elemental.PlainIdentifiable {

	out := NewClaims()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.Content != nil {
		out.Content = *o.Content
	}
	if o.FirstSeen != nil {
		out.FirstSeen = *o.FirstSeen
	}
	if o.Hash != nil {
		out.Hash = *o.Hash
	}
	if o.LastSeen != nil {
		out.LastSeen = *o.LastSeen
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
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseClaims) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseClaims) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseClaims) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseClaims) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseClaims) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseClaims) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseClaims) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseClaims) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseClaims) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseClaims) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseClaims) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseClaims) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseClaims) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseClaims) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseClaims.
func (o *SparseClaims) DeepCopy() *SparseClaims {

	if o == nil {
		return nil
	}

	out := &SparseClaims{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseClaims.
func (o *SparseClaims) DeepCopyInto(out *SparseClaims) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseClaims: %s", err))
	}

	*out = *target.(*SparseClaims)
}
