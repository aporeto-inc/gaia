package gaia

import (
	"fmt"
	"sync"

	"time"

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

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FileAccessReportsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o FileAccessReportsList) Version() int {

	return 1
}

// FileAccessReport represents the model of a fileaccessreport
type FileAccessReport struct {
	// Action taken.
	Action FileAccessReportActionValue `json:"action" bson:"-" mapstructure:"action,omitempty"`

	// Host of the file.
	Host string `json:"host" bson:"-" mapstructure:"host,omitempty"`

	// Mode of the file access.
	Mode string `json:"mode" bson:"-" mapstructure:"mode,omitempty"`

	// Path of the file.
	Path string `json:"path" bson:"-" mapstructure:"path,omitempty"`

	// ID of the processing unit.
	ProcessingUnitID string `json:"processingUnitID" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit.
	ProcessingUnitNamespace string `json:"processingUnitNamespace" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
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

// Validate valides the current information stored into the structure.
func (o *FileAccessReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Accept", "Reject"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("host", o.Host); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("mode", o.Mode); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("path", o.Path); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("processingUnitID", o.ProcessingUnitID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("processingUnitNamespace", o.ProcessingUnitNamespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredTime("timestamp", o.Timestamp); err != nil {
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
