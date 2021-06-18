package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NamespaceInfoPUIncomingTrafficActionValue represents the possible values for attribute "PUIncomingTrafficAction".
type NamespaceInfoPUIncomingTrafficActionValue string

const (
	// NamespaceInfoPUIncomingTrafficActionAllow represents the value Allow.
	NamespaceInfoPUIncomingTrafficActionAllow NamespaceInfoPUIncomingTrafficActionValue = "Allow"

	// NamespaceInfoPUIncomingTrafficActionInherit represents the value Inherit.
	NamespaceInfoPUIncomingTrafficActionInherit NamespaceInfoPUIncomingTrafficActionValue = "Inherit"

	// NamespaceInfoPUIncomingTrafficActionReject represents the value Reject.
	NamespaceInfoPUIncomingTrafficActionReject NamespaceInfoPUIncomingTrafficActionValue = "Reject"
)

// NamespaceInfoPUOutgoingTrafficActionValue represents the possible values for attribute "PUOutgoingTrafficAction".
type NamespaceInfoPUOutgoingTrafficActionValue string

const (
	// NamespaceInfoPUOutgoingTrafficActionAllow represents the value Allow.
	NamespaceInfoPUOutgoingTrafficActionAllow NamespaceInfoPUOutgoingTrafficActionValue = "Allow"

	// NamespaceInfoPUOutgoingTrafficActionInherit represents the value Inherit.
	NamespaceInfoPUOutgoingTrafficActionInherit NamespaceInfoPUOutgoingTrafficActionValue = "Inherit"

	// NamespaceInfoPUOutgoingTrafficActionReject represents the value Reject.
	NamespaceInfoPUOutgoingTrafficActionReject NamespaceInfoPUOutgoingTrafficActionValue = "Reject"
)

// NamespaceInfoIdentity represents the Identity of the object.
var NamespaceInfoIdentity = elemental.Identity{
	Name:     "namespaceinfo",
	Category: "namespaceinfo",
	Package:  "squall",
	Private:  false,
}

// NamespaceInfosList represents a list of NamespaceInfos
type NamespaceInfosList []*NamespaceInfo

// Identity returns the identity of the objects in the list.
func (o NamespaceInfosList) Identity() elemental.Identity {

	return NamespaceInfoIdentity
}

// Copy returns a pointer to a copy the NamespaceInfosList.
func (o NamespaceInfosList) Copy() elemental.Identifiables {

	copy := append(NamespaceInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the NamespaceInfosList.
func (o NamespaceInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(NamespaceInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*NamespaceInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o NamespaceInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o NamespaceInfosList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the NamespaceInfosList converted to SparseNamespaceInfosList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o NamespaceInfosList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseNamespaceInfosList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseNamespaceInfo)
	}

	return out
}

// Version returns the version of the content.
func (o NamespaceInfosList) Version() int {

	return 1
}

// NamespaceInfo represents the model of a namespaceinfo
type NamespaceInfo struct {
	// The processing unit action for incoming traffic for the namespace.
	PUIncomingTrafficAction NamespaceInfoPUIncomingTrafficActionValue `json:"PUIncomingTrafficAction" msgpack:"PUIncomingTrafficAction" bson:"-" mapstructure:"PUIncomingTrafficAction,omitempty"`

	// The processing unit action for outgoing traffic for the namespace.
	PUOutgoingTrafficAction NamespaceInfoPUOutgoingTrafficActionValue `json:"PUOutgoingTrafficAction" msgpack:"PUOutgoingTrafficAction" bson:"-" mapstructure:"PUOutgoingTrafficAction,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// List of tag prefixes that will be used to suggest policies.
	Prefixes []string `json:"prefixes" msgpack:"prefixes" bson:"-" mapstructure:"prefixes,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewNamespaceInfo returns a new *NamespaceInfo
func NewNamespaceInfo() *NamespaceInfo {

	return &NamespaceInfo{
		ModelVersion: 1,
		Prefixes:     []string{},
	}
}

// Identity returns the Identity of the object.
func (o *NamespaceInfo) Identity() elemental.Identity {

	return NamespaceInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NamespaceInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NamespaceInfo) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespaceInfo) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesNamespaceInfo{}

	s.Description = o.Description
	s.Name = o.Name
	s.Protected = o.Protected

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *NamespaceInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesNamespaceInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Description = s.Description
	o.Name = s.Name
	o.Protected = s.Protected

	return nil
}

// Version returns the hardcoded version of the model.
func (o *NamespaceInfo) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *NamespaceInfo) BleveType() string {

	return "namespaceinfo"
}

// DefaultOrder returns the list of default ordering fields.
func (o *NamespaceInfo) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *NamespaceInfo) Doc() string {

	return `Returns the information of the specified namespace.`
}

func (o *NamespaceInfo) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetDescription returns the Description of the receiver.
func (o *NamespaceInfo) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *NamespaceInfo) SetDescription(description string) {

	o.Description = description
}

// GetName returns the Name of the receiver.
func (o *NamespaceInfo) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *NamespaceInfo) SetName(name string) {

	o.Name = name
}

// GetProtected returns the Protected of the receiver.
func (o *NamespaceInfo) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *NamespaceInfo) SetProtected(protected bool) {

	o.Protected = protected
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *NamespaceInfo) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseNamespaceInfo{
			PUIncomingTrafficAction: &o.PUIncomingTrafficAction,
			PUOutgoingTrafficAction: &o.PUOutgoingTrafficAction,
			Description:             &o.Description,
			Name:                    &o.Name,
			Prefixes:                &o.Prefixes,
			Protected:               &o.Protected,
		}
	}

	sp := &SparseNamespaceInfo{}
	for _, f := range fields {
		switch f {
		case "PUIncomingTrafficAction":
			sp.PUIncomingTrafficAction = &(o.PUIncomingTrafficAction)
		case "PUOutgoingTrafficAction":
			sp.PUOutgoingTrafficAction = &(o.PUOutgoingTrafficAction)
		case "description":
			sp.Description = &(o.Description)
		case "name":
			sp.Name = &(o.Name)
		case "prefixes":
			sp.Prefixes = &(o.Prefixes)
		case "protected":
			sp.Protected = &(o.Protected)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseNamespaceInfo to the object.
func (o *NamespaceInfo) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseNamespaceInfo)
	if so.PUIncomingTrafficAction != nil {
		o.PUIncomingTrafficAction = *so.PUIncomingTrafficAction
	}
	if so.PUOutgoingTrafficAction != nil {
		o.PUOutgoingTrafficAction = *so.PUOutgoingTrafficAction
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Prefixes != nil {
		o.Prefixes = *so.Prefixes
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
}

// DeepCopy returns a deep copy if the NamespaceInfo.
func (o *NamespaceInfo) DeepCopy() *NamespaceInfo {

	if o == nil {
		return nil
	}

	out := &NamespaceInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *NamespaceInfo.
func (o *NamespaceInfo) DeepCopyInto(out *NamespaceInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy NamespaceInfo: %s", err))
	}

	*out = *target.(*NamespaceInfo)
}

// Validate valides the current information stored into the structure.
func (o *NamespaceInfo) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("PUIncomingTrafficAction", string(o.PUIncomingTrafficAction), []string{"Allow", "Reject", "Inherit"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("PUOutgoingTrafficAction", string(o.PUOutgoingTrafficAction), []string{"Allow", "Reject", "Inherit"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
func (*NamespaceInfo) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := NamespaceInfoAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return NamespaceInfoLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*NamespaceInfo) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NamespaceInfoAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *NamespaceInfo) ValueForAttribute(name string) interface{} {

	switch name {
	case "PUIncomingTrafficAction":
		return o.PUIncomingTrafficAction
	case "PUOutgoingTrafficAction":
		return o.PUOutgoingTrafficAction
	case "description":
		return o.Description
	case "name":
		return o.Name
	case "prefixes":
		return o.Prefixes
	case "protected":
		return o.Protected
	}

	return nil
}

// NamespaceInfoAttributesMap represents the map of attribute for NamespaceInfo.
var NamespaceInfoAttributesMap = map[string]elemental.AttributeSpecification{
	"PUIncomingTrafficAction": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		ConvertedName:  "PUIncomingTrafficAction",
		Description:    `The processing unit action for incoming traffic for the namespace.`,
		Exposed:        true,
		Name:           "PUIncomingTrafficAction",
		ReadOnly:       true,
		Type:           "enum",
	},
	"PUOutgoingTrafficAction": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		ConvertedName:  "PUOutgoingTrafficAction",
		Description:    `The processing unit action for outgoing traffic for the namespace.`,
		Exposed:        true,
		Name:           "PUOutgoingTrafficAction",
		ReadOnly:       true,
		Type:           "enum",
	},
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
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
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
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
	"Prefixes": {
		AllowedChoices: []string{},
		ConvertedName:  "Prefixes",
		Description:    `List of tag prefixes that will be used to suggest policies.`,
		Exposed:        true,
		Name:           "prefixes",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"Protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
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
}

// NamespaceInfoLowerCaseAttributesMap represents the map of attribute for NamespaceInfo.
var NamespaceInfoLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"puincomingtrafficaction": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		ConvertedName:  "PUIncomingTrafficAction",
		Description:    `The processing unit action for incoming traffic for the namespace.`,
		Exposed:        true,
		Name:           "PUIncomingTrafficAction",
		ReadOnly:       true,
		Type:           "enum",
	},
	"puoutgoingtrafficaction": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		ConvertedName:  "PUOutgoingTrafficAction",
		Description:    `The processing unit action for outgoing traffic for the namespace.`,
		Exposed:        true,
		Name:           "PUOutgoingTrafficAction",
		ReadOnly:       true,
		Type:           "enum",
	},
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
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
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
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
	"prefixes": {
		AllowedChoices: []string{},
		ConvertedName:  "Prefixes",
		Description:    `List of tag prefixes that will be used to suggest policies.`,
		Exposed:        true,
		Name:           "prefixes",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
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
}

// SparseNamespaceInfosList represents a list of SparseNamespaceInfos
type SparseNamespaceInfosList []*SparseNamespaceInfo

// Identity returns the identity of the objects in the list.
func (o SparseNamespaceInfosList) Identity() elemental.Identity {

	return NamespaceInfoIdentity
}

// Copy returns a pointer to a copy the SparseNamespaceInfosList.
func (o SparseNamespaceInfosList) Copy() elemental.Identifiables {

	copy := append(SparseNamespaceInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseNamespaceInfosList.
func (o SparseNamespaceInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseNamespaceInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseNamespaceInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseNamespaceInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseNamespaceInfosList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseNamespaceInfosList converted to NamespaceInfosList.
func (o SparseNamespaceInfosList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseNamespaceInfosList) Version() int {

	return 1
}

// SparseNamespaceInfo represents the sparse version of a namespaceinfo.
type SparseNamespaceInfo struct {
	// The processing unit action for incoming traffic for the namespace.
	PUIncomingTrafficAction *NamespaceInfoPUIncomingTrafficActionValue `json:"PUIncomingTrafficAction,omitempty" msgpack:"PUIncomingTrafficAction,omitempty" bson:"-" mapstructure:"PUIncomingTrafficAction,omitempty"`

	// The processing unit action for outgoing traffic for the namespace.
	PUOutgoingTrafficAction *NamespaceInfoPUOutgoingTrafficActionValue `json:"PUOutgoingTrafficAction,omitempty" msgpack:"PUOutgoingTrafficAction,omitempty" bson:"-" mapstructure:"PUOutgoingTrafficAction,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// List of tag prefixes that will be used to suggest policies.
	Prefixes *[]string `json:"prefixes,omitempty" msgpack:"prefixes,omitempty" bson:"-" mapstructure:"prefixes,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseNamespaceInfo returns a new  SparseNamespaceInfo.
func NewSparseNamespaceInfo() *SparseNamespaceInfo {
	return &SparseNamespaceInfo{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseNamespaceInfo) Identity() elemental.Identity {

	return NamespaceInfoIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseNamespaceInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseNamespaceInfo) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespaceInfo) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseNamespaceInfo{}

	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespaceInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseNamespaceInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseNamespaceInfo) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseNamespaceInfo) ToPlain() elemental.PlainIdentifiable {

	out := NewNamespaceInfo()
	if o.PUIncomingTrafficAction != nil {
		out.PUIncomingTrafficAction = *o.PUIncomingTrafficAction
	}
	if o.PUOutgoingTrafficAction != nil {
		out.PUOutgoingTrafficAction = *o.PUOutgoingTrafficAction
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Prefixes != nil {
		out.Prefixes = *o.Prefixes
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}

	return out
}

// GetDescription returns the Description of the receiver.
func (o *SparseNamespaceInfo) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseNamespaceInfo) SetDescription(description string) {

	o.Description = &description
}

// GetName returns the Name of the receiver.
func (o *SparseNamespaceInfo) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseNamespaceInfo) SetName(name string) {

	o.Name = &name
}

// GetProtected returns the Protected of the receiver.
func (o *SparseNamespaceInfo) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseNamespaceInfo) SetProtected(protected bool) {

	o.Protected = &protected
}

// DeepCopy returns a deep copy if the SparseNamespaceInfo.
func (o *SparseNamespaceInfo) DeepCopy() *SparseNamespaceInfo {

	if o == nil {
		return nil
	}

	out := &SparseNamespaceInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseNamespaceInfo.
func (o *SparseNamespaceInfo) DeepCopyInto(out *SparseNamespaceInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseNamespaceInfo: %s", err))
	}

	*out = *target.(*SparseNamespaceInfo)
}

type mongoAttributesNamespaceInfo struct {
	Description string `bson:"description"`
	Name        string `bson:"name"`
	Protected   bool   `bson:"protected"`
}
type mongoAttributesSparseNamespaceInfo struct {
	Description *string `bson:"description,omitempty"`
	Name        *string `bson:"name,omitempty"`
	Protected   *bool   `bson:"protected,omitempty"`
}
