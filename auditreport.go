package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// AuditReportIdentity represents the Identity of the object.
var AuditReportIdentity = elemental.Identity{
	Name:     "auditreport",
	Category: "auditreports",
	Package:  "zack",
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
	AUID string `json:"AUID" bson:"-" mapstructure:"AUID,omitempty"`

	// Command working directory.
	CWD string `json:"CWD" bson:"-" mapstructure:"CWD,omitempty"`

	// Needs documentation.
	EGID int `json:"EGID" bson:"-" mapstructure:"EGID,omitempty"`

	// Needs documentation.
	EUID int `json:"EUID" bson:"-" mapstructure:"EUID,omitempty"`

	// Path to the executable.
	EXE string `json:"EXE" bson:"-" mapstructure:"EXE,omitempty"`

	// Needs documentation.
	FSGID int `json:"FSGID" bson:"-" mapstructure:"FSGID,omitempty"`

	// Needs documentation.
	FSUID int `json:"FSUID" bson:"-" mapstructure:"FSUID,omitempty"`

	// Needs documentation.
	GID int `json:"GID" bson:"-" mapstructure:"GID,omitempty"`

	// Needs documentation.
	PER int `json:"PER" bson:"-" mapstructure:"PER,omitempty"`

	// PID of the executable.
	PID int `json:"PID" bson:"-" mapstructure:"PID,omitempty"`

	// PID of the parent executable.
	PPID int `json:"PPID" bson:"-" mapstructure:"PPID,omitempty"`

	// Needs documentation.
	SGID int `json:"SGID" bson:"-" mapstructure:"SGID,omitempty"`

	// Needs documentation.
	SUID int `json:"SUID" bson:"-" mapstructure:"SUID,omitempty"`

	// Needs documentation.
	UID int `json:"UID" bson:"-" mapstructure:"UID,omitempty"`

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

	// Command issued.
	Command string `json:"command" bson:"-" mapstructure:"command,omitempty"`

	// ID of the enforcer reporting.
	EnforcerID string `json:"enforcerID" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer reporting.
	EnforcerNamespace string `json:"enforcerNamespace" bson:"-" mapstructure:"enforcerNamespace,omitempty"`

	// Exit code of the executable.
	Exit int `json:"exit" bson:"-" mapstructure:"exit,omitempty"`

	// ID of the processing unit originating the report.
	ProcessingUnitID string `json:"processingUnitID" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit originating the report.
	ProcessingUnitNamespace string `json:"processingUnitNamespace" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Type of record.
	RecordType string `json:"recordType" bson:"-" mapstructure:"recordType,omitempty"`

	// Needs documentation.
	Sequence int `json:"sequence" bson:"-" mapstructure:"sequence,omitempty"`

	// Tells if the operation has been a success of a failure.
	Success bool `json:"success" bson:"-" mapstructure:"success,omitempty"`

	// Syscall name.
	Syscall string `json:"syscall" bson:"-" mapstructure:"syscall,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

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
	"AUID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AUID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "AUID",
		Type:           "string",
	},
	"CWD": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CWD",
		Description:    `Command working directory.`,
		Exposed:        true,
		Name:           "CWD",
		Type:           "string",
	},
	"EGID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EGID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "EGID",
		Type:           "integer",
	},
	"EUID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EUID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "EUID",
		Type:           "integer",
	},
	"EXE": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EXE",
		Description:    `Path to the executable.`,
		Exposed:        true,
		Name:           "EXE",
		Type:           "string",
	},
	"FSGID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FSGID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "FSGID",
		Type:           "integer",
	},
	"FSUID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FSUID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "FSUID",
		Type:           "integer",
	},
	"GID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "GID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "GID",
		Type:           "integer",
	},
	"PER": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PER",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "PER",
		Type:           "integer",
	},
	"PID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PID",
		Description:    `PID of the executable.`,
		Exposed:        true,
		Name:           "PID",
		Type:           "integer",
	},
	"PPID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PPID",
		Description:    `PID of the parent executable.`,
		Exposed:        true,
		Name:           "PPID",
		Type:           "integer",
	},
	"SGID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SGID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "SGID",
		Type:           "integer",
	},
	"SUID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SUID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "SUID",
		Type:           "integer",
	},
	"UID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "UID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "UID",
		Type:           "integer",
	},
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
	"Command": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Command",
		Description:    `Command issued.`,
		Exposed:        true,
		Name:           "command",
		Type:           "string",
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
	"Exit": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Exit",
		Description:    `Exit code of the executable.`,
		Exposed:        true,
		Name:           "exit",
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
	"Success": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Success",
		Description:    `Tells if the operation has been a success of a failure.`,
		Exposed:        true,
		Name:           "success",
		Required:       true,
		Type:           "boolean",
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
}

// AuditReportLowerCaseAttributesMap represents the map of attribute for AuditReport.
var AuditReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"auid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AUID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "AUID",
		Type:           "string",
	},
	"cwd": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CWD",
		Description:    `Command working directory.`,
		Exposed:        true,
		Name:           "CWD",
		Type:           "string",
	},
	"egid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EGID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "EGID",
		Type:           "integer",
	},
	"euid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EUID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "EUID",
		Type:           "integer",
	},
	"exe": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EXE",
		Description:    `Path to the executable.`,
		Exposed:        true,
		Name:           "EXE",
		Type:           "string",
	},
	"fsgid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FSGID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "FSGID",
		Type:           "integer",
	},
	"fsuid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FSUID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "FSUID",
		Type:           "integer",
	},
	"gid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "GID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "GID",
		Type:           "integer",
	},
	"per": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PER",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "PER",
		Type:           "integer",
	},
	"pid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PID",
		Description:    `PID of the executable.`,
		Exposed:        true,
		Name:           "PID",
		Type:           "integer",
	},
	"ppid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PPID",
		Description:    `PID of the parent executable.`,
		Exposed:        true,
		Name:           "PPID",
		Type:           "integer",
	},
	"sgid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SGID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "SGID",
		Type:           "integer",
	},
	"suid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SUID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "SUID",
		Type:           "integer",
	},
	"uid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "UID",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "UID",
		Type:           "integer",
	},
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
	"command": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Command",
		Description:    `Command issued.`,
		Exposed:        true,
		Name:           "command",
		Type:           "string",
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
	"exit": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Exit",
		Description:    `Exit code of the executable.`,
		Exposed:        true,
		Name:           "exit",
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
	"success": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Success",
		Description:    `Tells if the operation has been a success of a failure.`,
		Exposed:        true,
		Name:           "success",
		Required:       true,
		Type:           "boolean",
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
}
