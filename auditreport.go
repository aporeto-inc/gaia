package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// AuditReportIndexes lists the attribute compound indexes.
var AuditReportIndexes = [][]string{}

// AuditReportIdentity represents the Identity of the object.
var AuditReportIdentity = elemental.Identity{
	Name:     "auditreport",
	Category: "auditreports",
	Private:  false,
}

// AuditReportsList represents a list of AuditReports
type AuditReportsList []*AuditReport

// Identity returns the identity of the objects in the list.
func (o AuditReportsList) Identity() elemental.Identity {

	return AuditReportIdentity
}

// Copy returns a pointer to a copy the AuditReportsList.
func (o AuditReportsList) Copy() elemental.Identifiables {

	copy := append(AuditReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AuditReportsList.
func (o AuditReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AuditReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AuditReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AuditReportsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AuditReportsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o AuditReportsList) Version() int {

	return 1
}

// AuditReport represents the model of a auditreport
type AuditReport struct {
	// Needs documentation.
	A0 string `json:"a0" bson:"-" mapstructure:"a0,omitempty"`

	// Needs documentation.
	A1 string `json:"a1" bson:"-" mapstructure:"a1,omitempty"`

	// Needs documentation.
	A2 string `json:"a2" bson:"-" mapstructure:"a2,omitempty"`

	// Needs documentation.
	A3 string `json:"a3" bson:"-" mapstructure:"a3,omitempty"`

	// Architecture of the system where the syscall happened.
	Arch string `json:"arch" bson:"-" mapstructure:"arch,omitempty"`

	// ID the audit profile that triggered the report.
	AuditProfileID string `json:"auditProfileID" bson:"-" mapstructure:"auditProfileID,omitempty"`

	// Namespace the audit profile that triggered the report.
	AuditProfileNamespace string `json:"auditProfileNamespace" bson:"-" mapstructure:"auditProfileNamespace,omitempty"`

	// Needs documentation.
	Auid string `json:"auid" bson:"-" mapstructure:"auid,omitempty"`

	// Command issued.
	Command string `json:"command" bson:"-" mapstructure:"command,omitempty"`

	// Command working directory.
	Cwd string `json:"cwd" bson:"-" mapstructure:"cwd,omitempty"`

	// Needs documentation.
	Egid int `json:"egid" bson:"-" mapstructure:"egid,omitempty"`

	// ID of the enforcer reporting.
	EnforcerID string `json:"enforcerID" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer reporting.
	EnforcerNamespace string `json:"enforcerNamespace" bson:"-" mapstructure:"enforcerNamespace,omitempty"`

	// Needs documentation.
	Euid int `json:"euid" bson:"-" mapstructure:"euid,omitempty"`

	// Path to the executable.
	Exe string `json:"exe" bson:"-" mapstructure:"exe,omitempty"`

	// Exit code of the executable.
	Exit int `json:"exit" bson:"-" mapstructure:"exit,omitempty"`

	// Needs documentation.
	Fsgid int `json:"fsgid" bson:"-" mapstructure:"fsgid,omitempty"`

	// Needs documentation.
	Fsuid int `json:"fsuid" bson:"-" mapstructure:"fsuid,omitempty"`

	// Needs documentation.
	Gid int `json:"gid" bson:"-" mapstructure:"gid,omitempty"`

	// Needs documentation.
	Per int `json:"per" bson:"-" mapstructure:"per,omitempty"`

	// PID of the executable.
	Pid int `json:"pid" bson:"-" mapstructure:"pid,omitempty"`

	// PID of the parent executable.
	Ppid int `json:"ppid" bson:"-" mapstructure:"ppid,omitempty"`

	// ID of the processing unit originating the report.
	ProcessingUnitID string `json:"processingUnitID" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit originating the report.
	ProcessingUnitNamespace string `json:"processingUnitNamespace" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Type of record.
	RecordType string `json:"recordType" bson:"-" mapstructure:"recordType,omitempty"`

	// Needs documentation.
	Sequence int `json:"sequence" bson:"-" mapstructure:"sequence,omitempty"`

	// Needs documentation.
	Sgid int `json:"sgid" bson:"-" mapstructure:"sgid,omitempty"`

	// Tells if the operation has been a success of a failure.
	Success bool `json:"success" bson:"-" mapstructure:"success,omitempty"`

	// Needs documentation.
	Suid int `json:"suid" bson:"-" mapstructure:"suid,omitempty"`

	// Syscall name.
	Syscall string `json:"syscall" bson:"-" mapstructure:"syscall,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	// Needs documentation.
	Uid int `json:"uid" bson:"-" mapstructure:"uid,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewAuditReport returns a new *AuditReport
func NewAuditReport() *AuditReport {

	return &AuditReport{
		ModelVersion: 1,
		Success:      false,
	}
}

// Identity returns the Identity of the object.
func (o *AuditReport) Identity() elemental.Identity {

	return AuditReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AuditReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AuditReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *AuditReport) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *AuditReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *AuditReport) Doc() string {
	return `Post a new audit statistics report.`
}

func (o *AuditReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *AuditReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("auditProfileID", o.AuditProfileID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("auditProfileNamespace", o.AuditProfileNamespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("processingUnitID", o.ProcessingUnitID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("processingUnitNamespace", o.ProcessingUnitNamespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("recordType", o.RecordType); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("syscall", o.Syscall); err != nil {
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
func (*AuditReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AuditReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AuditReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AuditReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AuditReportAttributesMap
}

// AuditReportAttributesMap represents the map of attribute for AuditReport.
var AuditReportAttributesMap = map[string]elemental.AttributeSpecification{
	"A0": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "A0",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "a0",
		Type:           "string",
	},
	"A1": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "A1",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "a1",
		Type:           "string",
	},
	"A2": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "A2",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "a2",
		Type:           "string",
	},
	"A3": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "A3",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "a3",
		Type:           "string",
	},
	"Arch": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Arch",
		Description:    `Architecture of the system where the syscall happened.`,
		Exposed:        true,
		Name:           "arch",
		Type:           "string",
	},
	"AuditProfileID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditProfileID",
		Description:    `ID the audit profile that triggered the report.`,
		Exposed:        true,
		Name:           "auditProfileID",
		Required:       true,
		Type:           "string",
	},
	"AuditProfileNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditProfileNamespace",
		Description:    `Namespace the audit profile that triggered the report.`,
		Exposed:        true,
		Name:           "auditProfileNamespace",
		Required:       true,
		Type:           "string",
	},
	"Auid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Auid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "auid",
		Type:           "string",
	},
	"Command": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Command",
		Description:    `Command issued.`,
		Exposed:        true,
		Name:           "command",
		Type:           "string",
	},
	"Cwd": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Cwd",
		Description:    `Command working directory.`,
		Exposed:        true,
		Name:           "cwd",
		Type:           "string",
	},
	"Egid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Egid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "egid",
		Type:           "integer",
	},
	"EnforcerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer reporting.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Type:           "string",
	},
	"EnforcerNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer reporting.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Type:           "string",
	},
	"Euid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Euid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "euid",
		Type:           "integer",
	},
	"Exe": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Exe",
		Description:    `Path to the executable.`,
		Exposed:        true,
		Name:           "exe",
		Type:           "string",
	},
	"Exit": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Exit",
		Description:    `Exit code of the executable.`,
		Exposed:        true,
		Name:           "exit",
		Type:           "integer",
	},
	"Fsgid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fsgid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "fsgid",
		Type:           "integer",
	},
	"Fsuid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fsuid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "fsuid",
		Type:           "integer",
	},
	"Gid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Gid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "gid",
		Type:           "integer",
	},
	"Per": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Per",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "per",
		Type:           "integer",
	},
	"Pid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Pid",
		Description:    `PID of the executable.`,
		Exposed:        true,
		Name:           "pid",
		Type:           "integer",
	},
	"Ppid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Ppid",
		Description:    `PID of the parent executable.`,
		Exposed:        true,
		Name:           "ppid",
		Type:           "integer",
	},
	"ProcessingUnitID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit originating the report.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Type:           "string",
	},
	"ProcessingUnitNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit originating the report.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Required:       true,
		Type:           "string",
	},
	"RecordType": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RecordType",
		Description:    `Type of record.`,
		Exposed:        true,
		Name:           "recordType",
		Required:       true,
		Type:           "string",
	},
	"Sequence": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Sequence",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "sequence",
		Type:           "integer",
	},
	"Sgid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Sgid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "sgid",
		Type:           "integer",
	},
	"Success": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Success",
		Description:    `Tells if the operation has been a success of a failure.`,
		Exposed:        true,
		Name:           "success",
		Required:       true,
		Type:           "boolean",
	},
	"Suid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Suid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "suid",
		Type:           "integer",
	},
	"Syscall": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Syscall",
		Description:    `Syscall name.`,
		Exposed:        true,
		Name:           "syscall",
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
	"Uid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Uid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "uid",
		Type:           "integer",
	},
}

// AuditReportLowerCaseAttributesMap represents the map of attribute for AuditReport.
var AuditReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"a0": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "A0",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "a0",
		Type:           "string",
	},
	"a1": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "A1",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "a1",
		Type:           "string",
	},
	"a2": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "A2",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "a2",
		Type:           "string",
	},
	"a3": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "A3",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "a3",
		Type:           "string",
	},
	"arch": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Arch",
		Description:    `Architecture of the system where the syscall happened.`,
		Exposed:        true,
		Name:           "arch",
		Type:           "string",
	},
	"auditprofileid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditProfileID",
		Description:    `ID the audit profile that triggered the report.`,
		Exposed:        true,
		Name:           "auditProfileID",
		Required:       true,
		Type:           "string",
	},
	"auditprofilenamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditProfileNamespace",
		Description:    `Namespace the audit profile that triggered the report.`,
		Exposed:        true,
		Name:           "auditProfileNamespace",
		Required:       true,
		Type:           "string",
	},
	"auid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Auid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "auid",
		Type:           "string",
	},
	"command": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Command",
		Description:    `Command issued.`,
		Exposed:        true,
		Name:           "command",
		Type:           "string",
	},
	"cwd": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Cwd",
		Description:    `Command working directory.`,
		Exposed:        true,
		Name:           "cwd",
		Type:           "string",
	},
	"egid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Egid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "egid",
		Type:           "integer",
	},
	"enforcerid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer reporting.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Type:           "string",
	},
	"enforcernamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer reporting.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Type:           "string",
	},
	"euid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Euid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "euid",
		Type:           "integer",
	},
	"exe": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Exe",
		Description:    `Path to the executable.`,
		Exposed:        true,
		Name:           "exe",
		Type:           "string",
	},
	"exit": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Exit",
		Description:    `Exit code of the executable.`,
		Exposed:        true,
		Name:           "exit",
		Type:           "integer",
	},
	"fsgid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fsgid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "fsgid",
		Type:           "integer",
	},
	"fsuid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fsuid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "fsuid",
		Type:           "integer",
	},
	"gid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Gid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "gid",
		Type:           "integer",
	},
	"per": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Per",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "per",
		Type:           "integer",
	},
	"pid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Pid",
		Description:    `PID of the executable.`,
		Exposed:        true,
		Name:           "pid",
		Type:           "integer",
	},
	"ppid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Ppid",
		Description:    `PID of the parent executable.`,
		Exposed:        true,
		Name:           "ppid",
		Type:           "integer",
	},
	"processingunitid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit originating the report.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Type:           "string",
	},
	"processingunitnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit originating the report.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Required:       true,
		Type:           "string",
	},
	"recordtype": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RecordType",
		Description:    `Type of record.`,
		Exposed:        true,
		Name:           "recordType",
		Required:       true,
		Type:           "string",
	},
	"sequence": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Sequence",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "sequence",
		Type:           "integer",
	},
	"sgid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Sgid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "sgid",
		Type:           "integer",
	},
	"success": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Success",
		Description:    `Tells if the operation has been a success of a failure.`,
		Exposed:        true,
		Name:           "success",
		Required:       true,
		Type:           "boolean",
	},
	"suid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Suid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "suid",
		Type:           "integer",
	},
	"syscall": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Syscall",
		Description:    `Syscall name.`,
		Exposed:        true,
		Name:           "syscall",
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
	"uid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Uid",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "uid",
		Type:           "integer",
	},
}
