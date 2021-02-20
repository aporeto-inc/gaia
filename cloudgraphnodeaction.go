package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudGraphNodeActionIdentity represents the Identity of the object.
var CloudGraphNodeActionIdentity = elemental.Identity{
	Name:     "cloudgraphnodeaction",
	Category: "cloudgraphactions",
	Package:  "pcn/infrastructure",
	Private:  false,
}

// CloudGraphNodeActionsList represents a list of CloudGraphNodeActions
type CloudGraphNodeActionsList []*CloudGraphNodeAction

// Identity returns the identity of the objects in the list.
func (o CloudGraphNodeActionsList) Identity() elemental.Identity {

	return CloudGraphNodeActionIdentity
}

// Copy returns a pointer to a copy the CloudGraphNodeActionsList.
func (o CloudGraphNodeActionsList) Copy() elemental.Identifiables {

	copy := append(CloudGraphNodeActionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudGraphNodeActionsList.
func (o CloudGraphNodeActionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudGraphNodeActionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudGraphNodeAction))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudGraphNodeActionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudGraphNodeActionsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudGraphNodeActionsList converted to SparseCloudGraphNodeActionsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudGraphNodeActionsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudGraphNodeActionsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudGraphNodeAction)
	}

	return out
}

// Version returns the version of the content.
func (o CloudGraphNodeActionsList) Version() int {

	return 1
}

// CloudGraphNodeAction represents the model of a cloudgraphnodeaction
type CloudGraphNodeAction struct {
	// The action that is been applied for the particular destination.
	Action string `json:"action" msgpack:"action" bson:"-" mapstructure:"action,omitempty"`

	// The ID of the policies that were used in the path.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"-" mapstructure:"policyID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudGraphNodeAction returns a new *CloudGraphNodeAction
func NewCloudGraphNodeAction() *CloudGraphNodeAction {

	return &CloudGraphNodeAction{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *CloudGraphNodeAction) Identity() elemental.Identity {

	return CloudGraphNodeActionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudGraphNodeAction) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudGraphNodeAction) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphNodeAction) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudGraphNodeAction{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudGraphNodeAction) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudGraphNodeAction{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudGraphNodeAction) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudGraphNodeAction) BleveType() string {

	return "cloudgraphnodeaction"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudGraphNodeAction) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudGraphNodeAction) Doc() string {

	return `Describes the action and corresponding policy that resulted in this decision.`
}

func (o *CloudGraphNodeAction) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudGraphNodeAction) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudGraphNodeAction{
			Action:   &o.Action,
			PolicyID: &o.PolicyID,
		}
	}

	sp := &SparseCloudGraphNodeAction{}
	for _, f := range fields {
		switch f {
		case "action":
			sp.Action = &(o.Action)
		case "policyID":
			sp.PolicyID = &(o.PolicyID)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudGraphNodeAction to the object.
func (o *CloudGraphNodeAction) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudGraphNodeAction)
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.PolicyID != nil {
		o.PolicyID = *so.PolicyID
	}
}

// DeepCopy returns a deep copy if the CloudGraphNodeAction.
func (o *CloudGraphNodeAction) DeepCopy() *CloudGraphNodeAction {

	if o == nil {
		return nil
	}

	out := &CloudGraphNodeAction{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudGraphNodeAction.
func (o *CloudGraphNodeAction) DeepCopyInto(out *CloudGraphNodeAction) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudGraphNodeAction: %s", err))
	}

	*out = *target.(*CloudGraphNodeAction)
}

// Validate valides the current information stored into the structure.
func (o *CloudGraphNodeAction) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*CloudGraphNodeAction) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudGraphNodeActionAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudGraphNodeActionLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudGraphNodeAction) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudGraphNodeActionAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudGraphNodeAction) ValueForAttribute(name string) interface{} {

	switch name {
	case "action":
		return o.Action
	case "policyID":
		return o.PolicyID
	}

	return nil
}

// CloudGraphNodeActionAttributesMap represents the map of attribute for CloudGraphNodeAction.
var CloudGraphNodeActionAttributesMap = map[string]elemental.AttributeSpecification{
	"Action": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Action",
		Description:    `The action that is been applied for the particular destination.`,
		Exposed:        true,
		Name:           "action",
		ReadOnly:       true,
		Type:           "string",
	},
	"PolicyID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PolicyID",
		Description:    `The ID of the policies that were used in the path.`,
		Exposed:        true,
		Name:           "policyID",
		ReadOnly:       true,
		Type:           "string",
	},
}

// CloudGraphNodeActionLowerCaseAttributesMap represents the map of attribute for CloudGraphNodeAction.
var CloudGraphNodeActionLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"action": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Action",
		Description:    `The action that is been applied for the particular destination.`,
		Exposed:        true,
		Name:           "action",
		ReadOnly:       true,
		Type:           "string",
	},
	"policyid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PolicyID",
		Description:    `The ID of the policies that were used in the path.`,
		Exposed:        true,
		Name:           "policyID",
		ReadOnly:       true,
		Type:           "string",
	},
}

// SparseCloudGraphNodeActionsList represents a list of SparseCloudGraphNodeActions
type SparseCloudGraphNodeActionsList []*SparseCloudGraphNodeAction

// Identity returns the identity of the objects in the list.
func (o SparseCloudGraphNodeActionsList) Identity() elemental.Identity {

	return CloudGraphNodeActionIdentity
}

// Copy returns a pointer to a copy the SparseCloudGraphNodeActionsList.
func (o SparseCloudGraphNodeActionsList) Copy() elemental.Identifiables {

	copy := append(SparseCloudGraphNodeActionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudGraphNodeActionsList.
func (o SparseCloudGraphNodeActionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudGraphNodeActionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudGraphNodeAction))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudGraphNodeActionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudGraphNodeActionsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudGraphNodeActionsList converted to CloudGraphNodeActionsList.
func (o SparseCloudGraphNodeActionsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudGraphNodeActionsList) Version() int {

	return 1
}

// SparseCloudGraphNodeAction represents the sparse version of a cloudgraphnodeaction.
type SparseCloudGraphNodeAction struct {
	// The action that is been applied for the particular destination.
	Action *string `json:"action,omitempty" msgpack:"action,omitempty" bson:"-" mapstructure:"action,omitempty"`

	// The ID of the policies that were used in the path.
	PolicyID *string `json:"policyID,omitempty" msgpack:"policyID,omitempty" bson:"-" mapstructure:"policyID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudGraphNodeAction returns a new  SparseCloudGraphNodeAction.
func NewSparseCloudGraphNodeAction() *SparseCloudGraphNodeAction {
	return &SparseCloudGraphNodeAction{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudGraphNodeAction) Identity() elemental.Identity {

	return CloudGraphNodeActionIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudGraphNodeAction) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudGraphNodeAction) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudGraphNodeAction) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudGraphNodeAction{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudGraphNodeAction) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudGraphNodeAction{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCloudGraphNodeAction) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudGraphNodeAction) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudGraphNodeAction()
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.PolicyID != nil {
		out.PolicyID = *o.PolicyID
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCloudGraphNodeAction.
func (o *SparseCloudGraphNodeAction) DeepCopy() *SparseCloudGraphNodeAction {

	if o == nil {
		return nil
	}

	out := &SparseCloudGraphNodeAction{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudGraphNodeAction.
func (o *SparseCloudGraphNodeAction) DeepCopyInto(out *SparseCloudGraphNodeAction) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudGraphNodeAction: %s", err))
	}

	*out = *target.(*SparseCloudGraphNodeAction)
}

type mongoAttributesCloudGraphNodeAction struct {
}
type mongoAttributesSparseCloudGraphNodeAction struct {
}
