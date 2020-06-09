package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AccessReportActionValue represents the possible values for attribute "action".
type AccessReportActionValue string

const (
	// AccessReportActionAccept represents the value Accept.
	AccessReportActionAccept AccessReportActionValue = "Accept"

	// AccessReportActionReject represents the value Reject.
	AccessReportActionReject AccessReportActionValue = "Reject"
)

// AccessReportTypeValue represents the possible values for attribute "type".
type AccessReportTypeValue string

const (
	// AccessReportTypeSSHLogin represents the value SSHLogin.
	AccessReportTypeSSHLogin AccessReportTypeValue = "SSHLogin"

	// AccessReportTypeSSHLogout represents the value SSHLogout.
	AccessReportTypeSSHLogout AccessReportTypeValue = "SSHLogout"

	// AccessReportTypeSudoEnter represents the value SudoEnter.
	AccessReportTypeSudoEnter AccessReportTypeValue = "SudoEnter"

	// AccessReportTypeSudoExit represents the value SudoExit.
	AccessReportTypeSudoExit AccessReportTypeValue = "SudoExit"
)

// AccessReportIdentity represents the Identity of the object.
var AccessReportIdentity = elemental.Identity{
	Name:     "accessreport",
	Category: "accessreports",
	Package:  "zack",
	Private:  false,
}

// AccessReportsList represents a list of AccessReports
type AccessReportsList []*AccessReport

// Identity returns the identity of the objects in the list.
func (o AccessReportsList) Identity() elemental.Identity {

	return AccessReportIdentity
}

// Copy returns a pointer to a copy the AccessReportsList.
func (o AccessReportsList) Copy() elemental.Identifiables {

	copy := append(AccessReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AccessReportsList.
func (o AccessReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AccessReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AccessReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AccessReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AccessReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AccessReportsList converted to SparseAccessReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AccessReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAccessReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAccessReport)
	}

	return out
}

// Version returns the version of the content.
func (o AccessReportsList) Version() int {

	return 1
}

// AccessReport represents the model of a accessreport
type AccessReport struct {
	// Action applied to the access.
	Action AccessReportActionValue `json:"action" msgpack:"action" bson:"-" mapstructure:"action,omitempty"`

	// Hash of the claims used to communicate.
	ClaimHash string `json:"claimHash" msgpack:"claimHash" bson:"-" mapstructure:"claimHash,omitempty"`

	// Identifier of the defender.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Namespace of the defender.
	EnforcerNamespace string `json:"enforcerNamespace" msgpack:"enforcerNamespace" bson:"-" mapstructure:"enforcerNamespace,omitempty"`

	// ID of the processing unit of the report.
	ProcessingUnitID string `json:"processingUnitID" msgpack:"processingUnitID" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Name of the processing unit of the report.
	ProcessingUnitName string `json:"processingUnitName" msgpack:"processingUnitName" bson:"-" mapstructure:"processingUnitName,omitempty"`

	// Namespace of the processing unit of the report.
	ProcessingUnitNamespace string `json:"processingUnitNamespace" msgpack:"processingUnitNamespace" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// This field is only set if `action` is set to `Reject`. It specifies the reason
	// for the rejection.
	Reason string `json:"reason" msgpack:"reason" bson:"-" mapstructure:"reason,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	// Type of the report.
	Type AccessReportTypeValue `json:"type" msgpack:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAccessReport returns a new *AccessReport
func NewAccessReport() *AccessReport {

	return &AccessReport{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *AccessReport) Identity() elemental.Identity {

	return AccessReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AccessReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AccessReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AccessReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAccessReport{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AccessReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAccessReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *AccessReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *AccessReport) BleveType() string {

	return "accessreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *AccessReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *AccessReport) Doc() string {

	return `Represents any access made by the user.`
}

func (o *AccessReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AccessReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAccessReport{
			Action:                  &o.Action,
			ClaimHash:               &o.ClaimHash,
			EnforcerID:              &o.EnforcerID,
			EnforcerNamespace:       &o.EnforcerNamespace,
			ProcessingUnitID:        &o.ProcessingUnitID,
			ProcessingUnitName:      &o.ProcessingUnitName,
			ProcessingUnitNamespace: &o.ProcessingUnitNamespace,
			Reason:                  &o.Reason,
			Timestamp:               &o.Timestamp,
			Type:                    &o.Type,
		}
	}

	sp := &SparseAccessReport{}
	for _, f := range fields {
		switch f {
		case "action":
			sp.Action = &(o.Action)
		case "claimHash":
			sp.ClaimHash = &(o.ClaimHash)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "processingUnitName":
			sp.ProcessingUnitName = &(o.ProcessingUnitName)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "reason":
			sp.Reason = &(o.Reason)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "type":
			sp.Type = &(o.Type)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAccessReport to the object.
func (o *AccessReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAccessReport)
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.ClaimHash != nil {
		o.ClaimHash = *so.ClaimHash
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.ProcessingUnitName != nil {
		o.ProcessingUnitName = *so.ProcessingUnitName
	}
	if so.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = *so.ProcessingUnitNamespace
	}
	if so.Reason != nil {
		o.Reason = *so.Reason
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
}

// DeepCopy returns a deep copy if the AccessReport.
func (o *AccessReport) DeepCopy() *AccessReport {

	if o == nil {
		return nil
	}

	out := &AccessReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AccessReport.
func (o *AccessReport) DeepCopyInto(out *AccessReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AccessReport: %s", err))
	}

	*out = *target.(*AccessReport)
}

// Validate valides the current information stored into the structure.
func (o *AccessReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("action", string(o.Action)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Accept", "Reject"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("type", string(o.Type)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"SSHLogin", "SSHLogout", "SudoEnter", "SudoExit"}, false); err != nil {
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
func (*AccessReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AccessReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AccessReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AccessReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AccessReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AccessReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "action":
		return o.Action
	case "claimHash":
		return o.ClaimHash
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "processingUnitID":
		return o.ProcessingUnitID
	case "processingUnitName":
		return o.ProcessingUnitName
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "reason":
		return o.Reason
	case "timestamp":
		return o.Timestamp
	case "type":
		return o.Type
	}

	return nil
}

// AccessReportAttributesMap represents the map of attribute for AccessReport.
var AccessReportAttributesMap = map[string]elemental.AttributeSpecification{
	"Action": {
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action applied to the access.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Type:           "enum",
	},
	"ClaimHash": {
		AllowedChoices: []string{},
		ConvertedName:  "ClaimHash",
		Description:    `Hash of the claims used to communicate.`,
		Exposed:        true,
		Name:           "claimHash",
		Type:           "string",
	},
	"EnforcerID": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `Identifier of the defender.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Type:           "string",
	},
	"EnforcerNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the defender.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Type:           "string",
	},
	"ProcessingUnitID": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit of the report.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Type:           "string",
	},
	"ProcessingUnitName": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitName",
		Description:    `Name of the processing unit of the report.`,
		Exposed:        true,
		Name:           "processingUnitName",
		Type:           "string",
	},
	"ProcessingUnitNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit of the report.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Type:           "string",
	},
	"Reason": {
		AllowedChoices: []string{},
		ConvertedName:  "Reason",
		Description: `This field is only set if ` + "`" + `action` + "`" + ` is set to ` + "`" + `Reject` + "`" + `. It specifies the reason
for the rejection.`,
		Exposed: true,
		Name:    "reason",
		Type:    "string",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"Type": {
		AllowedChoices: []string{"SSHLogin", "SSHLogout", "SudoEnter", "SudoExit"},
		ConvertedName:  "Type",
		Description:    `Type of the report.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Type:           "enum",
	},
}

// AccessReportLowerCaseAttributesMap represents the map of attribute for AccessReport.
var AccessReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"action": {
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action applied to the access.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Type:           "enum",
	},
	"claimhash": {
		AllowedChoices: []string{},
		ConvertedName:  "ClaimHash",
		Description:    `Hash of the claims used to communicate.`,
		Exposed:        true,
		Name:           "claimHash",
		Type:           "string",
	},
	"enforcerid": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `Identifier of the defender.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Type:           "string",
	},
	"enforcernamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the defender.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Type:           "string",
	},
	"processingunitid": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit of the report.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Type:           "string",
	},
	"processingunitname": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitName",
		Description:    `Name of the processing unit of the report.`,
		Exposed:        true,
		Name:           "processingUnitName",
		Type:           "string",
	},
	"processingunitnamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit of the report.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Type:           "string",
	},
	"reason": {
		AllowedChoices: []string{},
		ConvertedName:  "Reason",
		Description: `This field is only set if ` + "`" + `action` + "`" + ` is set to ` + "`" + `Reject` + "`" + `. It specifies the reason
for the rejection.`,
		Exposed: true,
		Name:    "reason",
		Type:    "string",
	},
	"timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"type": {
		AllowedChoices: []string{"SSHLogin", "SSHLogout", "SudoEnter", "SudoExit"},
		ConvertedName:  "Type",
		Description:    `Type of the report.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Type:           "enum",
	},
}

// SparseAccessReportsList represents a list of SparseAccessReports
type SparseAccessReportsList []*SparseAccessReport

// Identity returns the identity of the objects in the list.
func (o SparseAccessReportsList) Identity() elemental.Identity {

	return AccessReportIdentity
}

// Copy returns a pointer to a copy the SparseAccessReportsList.
func (o SparseAccessReportsList) Copy() elemental.Identifiables {

	copy := append(SparseAccessReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAccessReportsList.
func (o SparseAccessReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAccessReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAccessReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAccessReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAccessReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAccessReportsList converted to AccessReportsList.
func (o SparseAccessReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAccessReportsList) Version() int {

	return 1
}

// SparseAccessReport represents the sparse version of a accessreport.
type SparseAccessReport struct {
	// Action applied to the access.
	Action *AccessReportActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"-" mapstructure:"action,omitempty"`

	// Hash of the claims used to communicate.
	ClaimHash *string `json:"claimHash,omitempty" msgpack:"claimHash,omitempty" bson:"-" mapstructure:"claimHash,omitempty"`

	// Identifier of the defender.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Namespace of the defender.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"-" mapstructure:"enforcerNamespace,omitempty"`

	// ID of the processing unit of the report.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Name of the processing unit of the report.
	ProcessingUnitName *string `json:"processingUnitName,omitempty" msgpack:"processingUnitName,omitempty" bson:"-" mapstructure:"processingUnitName,omitempty"`

	// Namespace of the processing unit of the report.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// This field is only set if `action` is set to `Reject`. It specifies the reason
	// for the rejection.
	Reason *string `json:"reason,omitempty" msgpack:"reason,omitempty" bson:"-" mapstructure:"reason,omitempty"`

	// Date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	// Type of the report.
	Type *AccessReportTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAccessReport returns a new  SparseAccessReport.
func NewSparseAccessReport() *SparseAccessReport {
	return &SparseAccessReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAccessReport) Identity() elemental.Identity {

	return AccessReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAccessReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAccessReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAccessReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAccessReport{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAccessReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAccessReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseAccessReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAccessReport) ToPlain() elemental.PlainIdentifiable {

	out := NewAccessReport()
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.ClaimHash != nil {
		out.ClaimHash = *o.ClaimHash
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.ProcessingUnitName != nil {
		out.ProcessingUnitName = *o.ProcessingUnitName
	}
	if o.ProcessingUnitNamespace != nil {
		out.ProcessingUnitNamespace = *o.ProcessingUnitNamespace
	}
	if o.Reason != nil {
		out.Reason = *o.Reason
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}
	if o.Type != nil {
		out.Type = *o.Type
	}

	return out
}

// DeepCopy returns a deep copy if the SparseAccessReport.
func (o *SparseAccessReport) DeepCopy() *SparseAccessReport {

	if o == nil {
		return nil
	}

	out := &SparseAccessReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAccessReport.
func (o *SparseAccessReport) DeepCopyInto(out *SparseAccessReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAccessReport: %s", err))
	}

	*out = *target.(*SparseAccessReport)
}

type mongoAttributesAccessReport struct {
}
type mongoAttributesSparseAccessReport struct {
}
