package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ClaimsAccessIdentity represents the Identity of the object.
var ClaimsAccessIdentity = elemental.Identity{
	Name:     "claimsaccess",
	Category: "claimsaccesses",
	Package:  "guy",
	Private:  false,
}

// ClaimsAccessList represents a list of ClaimsAccess
type ClaimsAccessList []*ClaimsAccess

// Identity returns the identity of the objects in the list.
func (o ClaimsAccessList) Identity() elemental.Identity {

	return ClaimsAccessIdentity
}

// Copy returns a pointer to a copy the ClaimsAccessList.
func (o ClaimsAccessList) Copy() elemental.Identifiables {

	copy := append(ClaimsAccessList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ClaimsAccessList.
func (o ClaimsAccessList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ClaimsAccessList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ClaimsAccess))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ClaimsAccessList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ClaimsAccessList) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// ToSparse returns the ClaimsAccessList converted to SparseClaimsAccessList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ClaimsAccessList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o ClaimsAccessList) Version() int {

	return 1
}

// ClaimsAccess represents the model of a claimsaccess
type ClaimsAccess struct {
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

	// Unique hash of the claim access.
	Hash string `json:"hash" bson:"hash" mapstructure:"hash,omitempty"`

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

// NewClaimsAccess returns a new *ClaimsAccess
func NewClaimsAccess() *ClaimsAccess {

	return &ClaimsAccess{
		ModelVersion:   1,
		Mutex:          &sync.Mutex{},
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Content:        map[string]string{},
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *ClaimsAccess) Identity() elemental.Identity {

	return ClaimsAccessIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ClaimsAccess) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ClaimsAccess) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *ClaimsAccess) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *ClaimsAccess) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// Doc returns the documentation for the object
func (o *ClaimsAccess) Doc() string {
	return `This API represents the claims that accessed a service.`
}

func (o *ClaimsAccess) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *ClaimsAccess) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *ClaimsAccess) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *ClaimsAccess) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *ClaimsAccess) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetNamespace returns the Namespace of the receiver.
func (o *ClaimsAccess) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *ClaimsAccess) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *ClaimsAccess) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *ClaimsAccess) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *ClaimsAccess) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *ClaimsAccess) SetProtected(protected bool) {

	o.Protected = protected
}

// GetZHash returns the ZHash of the receiver.
func (o *ClaimsAccess) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *ClaimsAccess) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *ClaimsAccess) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *ClaimsAccess) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ClaimsAccess) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseClaimsAccess{
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

	sp := &SparseClaimsAccess{}
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

// Patch apply the non nil value of a *SparseClaimsAccess to the object.
func (o *ClaimsAccess) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseClaimsAccess)
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

// DeepCopy returns a deep copy if the ClaimsAccess.
func (o *ClaimsAccess) DeepCopy() *ClaimsAccess {

	if o == nil {
		return nil
	}

	out := &ClaimsAccess{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ClaimsAccess.
func (o *ClaimsAccess) DeepCopyInto(out *ClaimsAccess) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ClaimsAccess: %s", err))
	}

	*out = *target.(*ClaimsAccess)
}

// Validate valides the current information stored into the structure.
func (o *ClaimsAccess) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("content", o.Content); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

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
func (*ClaimsAccess) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ClaimsAccessAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ClaimsAccessLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ClaimsAccess) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ClaimsAccessAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ClaimsAccess) ValueForAttribute(name string) interface{} {

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

// ClaimsAccessAttributesMap represents the map of attribute for ClaimsAccess.
var ClaimsAccessAttributesMap = map[string]elemental.AttributeSpecification{
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
		Required:       true,
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
		Description:    `Unique hash of the claim access.`,
		Exposed:        true,
		Name:           "hash",
		Required:       true,
		Stored:         true,
		Type:           "string",
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

// ClaimsAccessLowerCaseAttributesMap represents the map of attribute for ClaimsAccess.
var ClaimsAccessLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		Required:       true,
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
		Description:    `Unique hash of the claim access.`,
		Exposed:        true,
		Name:           "hash",
		Required:       true,
		Stored:         true,
		Type:           "string",
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

// SparseClaimsAccessList represents a list of SparseClaimsAccess
type SparseClaimsAccessList []*SparseClaimsAccess

// Identity returns the identity of the objects in the list.
func (o SparseClaimsAccessList) Identity() elemental.Identity {

	return ClaimsAccessIdentity
}

// Copy returns a pointer to a copy the SparseClaimsAccessList.
func (o SparseClaimsAccessList) Copy() elemental.Identifiables {

	copy := append(SparseClaimsAccessList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseClaimsAccessList.
func (o SparseClaimsAccessList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseClaimsAccessList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseClaimsAccess))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseClaimsAccessList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseClaimsAccessList) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// ToPlain returns the SparseClaimsAccessList converted to ClaimsAccessList.
func (o SparseClaimsAccessList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseClaimsAccessList) Version() int {

	return 1
}

// SparseClaimsAccess represents the sparse version of a claimsaccess.
type SparseClaimsAccess struct {
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

	// Unique hash of the claim access.
	Hash *string `json:"hash,omitempty" bson:"hash,omitempty" mapstructure:"hash,omitempty"`

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

// NewSparseClaimsAccess returns a new  SparseClaimsAccess.
func NewSparseClaimsAccess() *SparseClaimsAccess {
	return &SparseClaimsAccess{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseClaimsAccess) Identity() elemental.Identity {

	return ClaimsAccessIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseClaimsAccess) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseClaimsAccess) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseClaimsAccess) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseClaimsAccess) ToPlain() elemental.PlainIdentifiable {

	out := NewClaimsAccess()
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
func (o *SparseClaimsAccess) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseClaimsAccess) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseClaimsAccess) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseClaimsAccess) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseClaimsAccess) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseClaimsAccess) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseClaimsAccess) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseClaimsAccess) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseClaimsAccess) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseClaimsAccess) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseClaimsAccess) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseClaimsAccess) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseClaimsAccess) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseClaimsAccess) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseClaimsAccess.
func (o *SparseClaimsAccess) DeepCopy() *SparseClaimsAccess {

	if o == nil {
		return nil
	}

	out := &SparseClaimsAccess{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseClaimsAccess.
func (o *SparseClaimsAccess) DeepCopyInto(out *SparseClaimsAccess) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseClaimsAccess: %s", err))
	}

	*out = *target.(*SparseClaimsAccess)
}
