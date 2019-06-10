package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FileAccessReportActionValue represents the possible values for attribute "action".
type FileAccessReportActionValue string

const (
	// FileAccessReportActionAccept represents the value Accept.
	FileAccessReportActionAccept FileAccessReportActionValue = "Accept"

	// FileAccessReportActionReject represents the value Reject.
	FileAccessReportActionReject FileAccessReportActionValue = "Reject"
)

// FileAccessReportIdentity represents the Identity of the object.
var FileAccessReportIdentity = elemental.Identity{
	Name:     "fileaccessreport",
	Category: "fileaccessreports",
	Package:  "zack",
	Private:  false,
}

// FileAccessReportsList represents a list of FileAccessReports
type FileAccessReportsList []*FileAccessReport

// Identity returns the identity of the objects in the list.
func (o FileAccessReportsList) Identity() elemental.Identity {

	return FileAccessReportIdentity
}

// Copy returns a pointer to a copy the FileAccessReportsList.
func (o FileAccessReportsList) Copy() elemental.Identifiables {

	copy := append(FileAccessReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the FileAccessReportsList.
func (o FileAccessReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(FileAccessReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FileAccessReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FileAccessReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FileAccessReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the FileAccessReportsList converted to SparseFileAccessReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o FileAccessReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseFileAccessReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseFileAccessReport)
	}

	return out
}

// Version returns the version of the content.
func (o FileAccessReportsList) Version() int {

	return 1
}

// FileAccessReport represents the model of a fileaccessreport
type FileAccessReport struct {
	// Action taken.
	Action FileAccessReportActionValue `json:"action" msgpack:"action" bson:"-" mapstructure:"action,omitempty"`

	// Host of the file.
	Host string `json:"host" msgpack:"host" bson:"-" mapstructure:"host,omitempty"`

	// Mode of the file access.
	Mode string `json:"mode" msgpack:"mode" bson:"-" mapstructure:"mode,omitempty"`

	// Path of the file.
	Path string `json:"path" msgpack:"path" bson:"-" mapstructure:"path,omitempty"`

	// ID of the processing unit.
	ProcessingUnitID string `json:"processingUnitID" msgpack:"processingUnitID" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit.
	ProcessingUnitNamespace string `json:"processingUnitNamespace" msgpack:"processingUnitNamespace" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFileAccessReport returns a new *FileAccessReport
func NewFileAccessReport() *FileAccessReport {

	return &FileAccessReport{
		ModelVersion: 1,
		Host:         "localhost",
		Mode:         "rxw",
		Path:         "/etc/passwd",
	}
}

// Identity returns the Identity of the object.
func (o *FileAccessReport) Identity() elemental.Identity {

	return FileAccessReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FileAccessReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FileAccessReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *FileAccessReport) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *FileAccessReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *FileAccessReport) Doc() string {

	return `Post a new file access statistics report.`
}

func (o *FileAccessReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *FileAccessReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseFileAccessReport{
			Action:                  &o.Action,
			Host:                    &o.Host,
			Mode:                    &o.Mode,
			Path:                    &o.Path,
			ProcessingUnitID:        &o.ProcessingUnitID,
			ProcessingUnitNamespace: &o.ProcessingUnitNamespace,
			Timestamp:               &o.Timestamp,
		}
	}

	sp := &SparseFileAccessReport{}
	for _, f := range fields {
		switch f {
		case "action":
			sp.Action = &(o.Action)
		case "host":
			sp.Host = &(o.Host)
		case "mode":
			sp.Mode = &(o.Mode)
		case "path":
			sp.Path = &(o.Path)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseFileAccessReport to the object.
func (o *FileAccessReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseFileAccessReport)
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.Host != nil {
		o.Host = *so.Host
	}
	if so.Mode != nil {
		o.Mode = *so.Mode
	}
	if so.Path != nil {
		o.Path = *so.Path
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = *so.ProcessingUnitNamespace
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
}

// DeepCopy returns a deep copy if the FileAccessReport.
func (o *FileAccessReport) DeepCopy() *FileAccessReport {

	if o == nil {
		return nil
	}

	out := &FileAccessReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FileAccessReport.
func (o *FileAccessReport) DeepCopyInto(out *FileAccessReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FileAccessReport: %s", err))
	}

	*out = *target.(*FileAccessReport)
}

// Validate valides the current information stored into the structure.
func (o *FileAccessReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("action", string(o.Action)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Accept", "Reject"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("host", o.Host); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("mode", o.Mode); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("path", o.Path); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("processingUnitID", o.ProcessingUnitID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("processingUnitNamespace", o.ProcessingUnitNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredTime("timestamp", o.Timestamp); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*FileAccessReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FileAccessReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FileAccessReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FileAccessReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FileAccessReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FileAccessReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "action":
		return o.Action
	case "host":
		return o.Host
	case "mode":
		return o.Mode
	case "path":
		return o.Path
	case "processingUnitID":
		return o.ProcessingUnitID
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "timestamp":
		return o.Timestamp
	}

	return nil
}

// FileAccessReportAttributesMap represents the map of attribute for FileAccessReport.
var FileAccessReportAttributesMap = map[string]elemental.AttributeSpecification{
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action taken.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Type:           "enum",
	},
	"Host": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Host",
		DefaultValue:   "localhost",
		Description:    `Host of the file.`,
		Exposed:        true,
		Name:           "host",
		Required:       true,
		Type:           "string",
	},
	"Mode": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Mode",
		DefaultValue:   "rxw",
		Description:    `Mode of the file access.`,
		Exposed:        true,
		Name:           "mode",
		Required:       true,
		Type:           "string",
	},
	"Path": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Path",
		DefaultValue:   "/etc/passwd",
		Description:    `Path of the file.`,
		Exposed:        true,
		Name:           "path",
		Required:       true,
		Type:           "string",
	},
	"ProcessingUnitID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Type:           "string",
	},
	"ProcessingUnitNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Required:       true,
		Type:           "string",
	},
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Type:           "time",
	},
}

// FileAccessReportLowerCaseAttributesMap represents the map of attribute for FileAccessReport.
var FileAccessReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action taken.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Type:           "enum",
	},
	"host": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Host",
		DefaultValue:   "localhost",
		Description:    `Host of the file.`,
		Exposed:        true,
		Name:           "host",
		Required:       true,
		Type:           "string",
	},
	"mode": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Mode",
		DefaultValue:   "rxw",
		Description:    `Mode of the file access.`,
		Exposed:        true,
		Name:           "mode",
		Required:       true,
		Type:           "string",
	},
	"path": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Path",
		DefaultValue:   "/etc/passwd",
		Description:    `Path of the file.`,
		Exposed:        true,
		Name:           "path",
		Required:       true,
		Type:           "string",
	},
	"processingunitid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Type:           "string",
	},
	"processingunitnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Required:       true,
		Type:           "string",
	},
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Type:           "time",
	},
}

// SparseFileAccessReportsList represents a list of SparseFileAccessReports
type SparseFileAccessReportsList []*SparseFileAccessReport

// Identity returns the identity of the objects in the list.
func (o SparseFileAccessReportsList) Identity() elemental.Identity {

	return FileAccessReportIdentity
}

// Copy returns a pointer to a copy the SparseFileAccessReportsList.
func (o SparseFileAccessReportsList) Copy() elemental.Identifiables {

	copy := append(SparseFileAccessReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseFileAccessReportsList.
func (o SparseFileAccessReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseFileAccessReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseFileAccessReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseFileAccessReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseFileAccessReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseFileAccessReportsList converted to FileAccessReportsList.
func (o SparseFileAccessReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseFileAccessReportsList) Version() int {

	return 1
}

// SparseFileAccessReport represents the sparse version of a fileaccessreport.
type SparseFileAccessReport struct {
	// Action taken.
	Action *FileAccessReportActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"-" mapstructure:"action,omitempty"`

	// Host of the file.
	Host *string `json:"host,omitempty" msgpack:"host,omitempty" bson:"-" mapstructure:"host,omitempty"`

	// Mode of the file access.
	Mode *string `json:"mode,omitempty" msgpack:"mode,omitempty" bson:"-" mapstructure:"mode,omitempty"`

	// Path of the file.
	Path *string `json:"path,omitempty" msgpack:"path,omitempty" bson:"-" mapstructure:"path,omitempty"`

	// ID of the processing unit.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseFileAccessReport returns a new  SparseFileAccessReport.
func NewSparseFileAccessReport() *SparseFileAccessReport {
	return &SparseFileAccessReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseFileAccessReport) Identity() elemental.Identity {

	return FileAccessReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseFileAccessReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseFileAccessReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseFileAccessReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseFileAccessReport) ToPlain() elemental.PlainIdentifiable {

	out := NewFileAccessReport()
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.Host != nil {
		out.Host = *o.Host
	}
	if o.Mode != nil {
		out.Mode = *o.Mode
	}
	if o.Path != nil {
		out.Path = *o.Path
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.ProcessingUnitNamespace != nil {
		out.ProcessingUnitNamespace = *o.ProcessingUnitNamespace
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}

	return out
}

// DeepCopy returns a deep copy if the SparseFileAccessReport.
func (o *SparseFileAccessReport) DeepCopy() *SparseFileAccessReport {

	if o == nil {
		return nil
	}

	out := &SparseFileAccessReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseFileAccessReport.
func (o *SparseFileAccessReport) DeepCopyInto(out *SparseFileAccessReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseFileAccessReport: %s", err))
	}

	*out = *target.(*SparseFileAccessReport)
}
