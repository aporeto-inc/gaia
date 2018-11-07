package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
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

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AuditReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AuditReportsList converted to SparseAuditReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AuditReportsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
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

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AuditReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAuditReport{
			AUID:                    &o.AUID,
			CWD:                     &o.CWD,
			EGID:                    &o.EGID,
			EUID:                    &o.EUID,
			EXE:                     &o.EXE,
			FSGID:                   &o.FSGID,
			FSUID:                   &o.FSUID,
			GID:                     &o.GID,
			PER:                     &o.PER,
			PID:                     &o.PID,
			PPID:                    &o.PPID,
			SGID:                    &o.SGID,
			SUID:                    &o.SUID,
			UID:                     &o.UID,
			A0:                      &o.A0,
			A1:                      &o.A1,
			A2:                      &o.A2,
			A3:                      &o.A3,
			Arch:                    &o.Arch,
			AuditProfileID:          &o.AuditProfileID,
			AuditProfileNamespace:   &o.AuditProfileNamespace,
			Command:                 &o.Command,
			EnforcerID:              &o.EnforcerID,
			EnforcerNamespace:       &o.EnforcerNamespace,
			Exit:                    &o.Exit,
			ProcessingUnitID:        &o.ProcessingUnitID,
			ProcessingUnitNamespace: &o.ProcessingUnitNamespace,
			RecordType:              &o.RecordType,
			Sequence:                &o.Sequence,
			Success:                 &o.Success,
			Syscall:                 &o.Syscall,
			Timestamp:               &o.Timestamp,
		}
	}

	sp := &SparseAuditReport{}
	for _, f := range fields {
		switch f {
		case "AUID":
			sp.AUID = &(o.AUID)
		case "CWD":
			sp.CWD = &(o.CWD)
		case "EGID":
			sp.EGID = &(o.EGID)
		case "EUID":
			sp.EUID = &(o.EUID)
		case "EXE":
			sp.EXE = &(o.EXE)
		case "FSGID":
			sp.FSGID = &(o.FSGID)
		case "FSUID":
			sp.FSUID = &(o.FSUID)
		case "GID":
			sp.GID = &(o.GID)
		case "PER":
			sp.PER = &(o.PER)
		case "PID":
			sp.PID = &(o.PID)
		case "PPID":
			sp.PPID = &(o.PPID)
		case "SGID":
			sp.SGID = &(o.SGID)
		case "SUID":
			sp.SUID = &(o.SUID)
		case "UID":
			sp.UID = &(o.UID)
		case "a0":
			sp.A0 = &(o.A0)
		case "a1":
			sp.A1 = &(o.A1)
		case "a2":
			sp.A2 = &(o.A2)
		case "a3":
			sp.A3 = &(o.A3)
		case "arch":
			sp.Arch = &(o.Arch)
		case "auditProfileID":
			sp.AuditProfileID = &(o.AuditProfileID)
		case "auditProfileNamespace":
			sp.AuditProfileNamespace = &(o.AuditProfileNamespace)
		case "command":
			sp.Command = &(o.Command)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "exit":
			sp.Exit = &(o.Exit)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "recordType":
			sp.RecordType = &(o.RecordType)
		case "sequence":
			sp.Sequence = &(o.Sequence)
		case "success":
			sp.Success = &(o.Success)
		case "syscall":
			sp.Syscall = &(o.Syscall)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAuditReport to the object.
func (o *AuditReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAuditReport)
	if so.AUID != nil {
		o.AUID = *so.AUID
	}
	if so.CWD != nil {
		o.CWD = *so.CWD
	}
	if so.EGID != nil {
		o.EGID = *so.EGID
	}
	if so.EUID != nil {
		o.EUID = *so.EUID
	}
	if so.EXE != nil {
		o.EXE = *so.EXE
	}
	if so.FSGID != nil {
		o.FSGID = *so.FSGID
	}
	if so.FSUID != nil {
		o.FSUID = *so.FSUID
	}
	if so.GID != nil {
		o.GID = *so.GID
	}
	if so.PER != nil {
		o.PER = *so.PER
	}
	if so.PID != nil {
		o.PID = *so.PID
	}
	if so.PPID != nil {
		o.PPID = *so.PPID
	}
	if so.SGID != nil {
		o.SGID = *so.SGID
	}
	if so.SUID != nil {
		o.SUID = *so.SUID
	}
	if so.UID != nil {
		o.UID = *so.UID
	}
	if so.A0 != nil {
		o.A0 = *so.A0
	}
	if so.A1 != nil {
		o.A1 = *so.A1
	}
	if so.A2 != nil {
		o.A2 = *so.A2
	}
	if so.A3 != nil {
		o.A3 = *so.A3
	}
	if so.Arch != nil {
		o.Arch = *so.Arch
	}
	if so.AuditProfileID != nil {
		o.AuditProfileID = *so.AuditProfileID
	}
	if so.AuditProfileNamespace != nil {
		o.AuditProfileNamespace = *so.AuditProfileNamespace
	}
	if so.Command != nil {
		o.Command = *so.Command
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.Exit != nil {
		o.Exit = *so.Exit
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = *so.ProcessingUnitNamespace
	}
	if so.RecordType != nil {
		o.RecordType = *so.RecordType
	}
	if so.Sequence != nil {
		o.Sequence = *so.Sequence
	}
	if so.Success != nil {
		o.Success = *so.Success
	}
	if so.Syscall != nil {
		o.Syscall = *so.Syscall
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
}

// DeepCopy returns a deep copy if the AuditReport.
func (o *AuditReport) DeepCopy() *AuditReport {

	if o == nil {
		return nil
	}

	out := &AuditReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AuditReport.
func (o *AuditReport) DeepCopyInto(out *AuditReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AuditReport: %s", err))
	}

	*out = *target.(*AuditReport)
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AuditReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "AUID":
		return o.AUID
	case "CWD":
		return o.CWD
	case "EGID":
		return o.EGID
	case "EUID":
		return o.EUID
	case "EXE":
		return o.EXE
	case "FSGID":
		return o.FSGID
	case "FSUID":
		return o.FSUID
	case "GID":
		return o.GID
	case "PER":
		return o.PER
	case "PID":
		return o.PID
	case "PPID":
		return o.PPID
	case "SGID":
		return o.SGID
	case "SUID":
		return o.SUID
	case "UID":
		return o.UID
	case "a0":
		return o.A0
	case "a1":
		return o.A1
	case "a2":
		return o.A2
	case "a3":
		return o.A3
	case "arch":
		return o.Arch
	case "auditProfileID":
		return o.AuditProfileID
	case "auditProfileNamespace":
		return o.AuditProfileNamespace
	case "command":
		return o.Command
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "exit":
		return o.Exit
	case "processingUnitID":
		return o.ProcessingUnitID
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "recordType":
		return o.RecordType
	case "sequence":
		return o.Sequence
	case "success":
		return o.Success
	case "syscall":
		return o.Syscall
	case "timestamp":
		return o.Timestamp
	}

	return nil
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

// SparseAuditReportsList represents a list of SparseAuditReports
type SparseAuditReportsList []*SparseAuditReport

// Identity returns the identity of the objects in the list.
func (o SparseAuditReportsList) Identity() elemental.Identity {

	return AuditReportIdentity
}

// Copy returns a pointer to a copy the SparseAuditReportsList.
func (o SparseAuditReportsList) Copy() elemental.Identifiables {

	copy := append(SparseAuditReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAuditReportsList.
func (o SparseAuditReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAuditReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAuditReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAuditReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAuditReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAuditReportsList converted to AuditReportsList.
func (o SparseAuditReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAuditReportsList) Version() int {

	return 1
}

// SparseAuditReport represents the sparse version of a auditreport.
type SparseAuditReport struct {
	// Needs documentation.
	AUID *string `json:"AUID,omitempty" bson:"-" mapstructure:"AUID,omitempty"`

	// Command working directory.
	CWD *string `json:"CWD,omitempty" bson:"-" mapstructure:"CWD,omitempty"`

	// Needs documentation.
	EGID *int `json:"EGID,omitempty" bson:"-" mapstructure:"EGID,omitempty"`

	// Needs documentation.
	EUID *int `json:"EUID,omitempty" bson:"-" mapstructure:"EUID,omitempty"`

	// Path to the executable.
	EXE *string `json:"EXE,omitempty" bson:"-" mapstructure:"EXE,omitempty"`

	// Needs documentation.
	FSGID *int `json:"FSGID,omitempty" bson:"-" mapstructure:"FSGID,omitempty"`

	// Needs documentation.
	FSUID *int `json:"FSUID,omitempty" bson:"-" mapstructure:"FSUID,omitempty"`

	// Needs documentation.
	GID *int `json:"GID,omitempty" bson:"-" mapstructure:"GID,omitempty"`

	// Needs documentation.
	PER *int `json:"PER,omitempty" bson:"-" mapstructure:"PER,omitempty"`

	// PID of the executable.
	PID *int `json:"PID,omitempty" bson:"-" mapstructure:"PID,omitempty"`

	// PID of the parent executable.
	PPID *int `json:"PPID,omitempty" bson:"-" mapstructure:"PPID,omitempty"`

	// Needs documentation.
	SGID *int `json:"SGID,omitempty" bson:"-" mapstructure:"SGID,omitempty"`

	// Needs documentation.
	SUID *int `json:"SUID,omitempty" bson:"-" mapstructure:"SUID,omitempty"`

	// Needs documentation.
	UID *int `json:"UID,omitempty" bson:"-" mapstructure:"UID,omitempty"`

	// Needs documentation.
	A0 *string `json:"a0,omitempty" bson:"-" mapstructure:"a0,omitempty"`

	// Needs documentation.
	A1 *string `json:"a1,omitempty" bson:"-" mapstructure:"a1,omitempty"`

	// Needs documentation.
	A2 *string `json:"a2,omitempty" bson:"-" mapstructure:"a2,omitempty"`

	// Needs documentation.
	A3 *string `json:"a3,omitempty" bson:"-" mapstructure:"a3,omitempty"`

	// Architecture of the system where the syscall happened.
	Arch *string `json:"arch,omitempty" bson:"-" mapstructure:"arch,omitempty"`

	// ID the audit profile that triggered the report.
	AuditProfileID *string `json:"auditProfileID,omitempty" bson:"-" mapstructure:"auditProfileID,omitempty"`

	// Namespace the audit profile that triggered the report.
	AuditProfileNamespace *string `json:"auditProfileNamespace,omitempty" bson:"-" mapstructure:"auditProfileNamespace,omitempty"`

	// Command issued.
	Command *string `json:"command,omitempty" bson:"-" mapstructure:"command,omitempty"`

	// ID of the enforcer reporting.
	EnforcerID *string `json:"enforcerID,omitempty" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer reporting.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" bson:"-" mapstructure:"enforcerNamespace,omitempty"`

	// Exit code of the executable.
	Exit *int `json:"exit,omitempty" bson:"-" mapstructure:"exit,omitempty"`

	// ID of the processing unit originating the report.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit originating the report.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Type of record.
	RecordType *string `json:"recordType,omitempty" bson:"-" mapstructure:"recordType,omitempty"`

	// Needs documentation.
	Sequence *int `json:"sequence,omitempty" bson:"-" mapstructure:"sequence,omitempty"`

	// Tells if the operation has been a success of a failure.
	Success *bool `json:"success,omitempty" bson:"-" mapstructure:"success,omitempty"`

	// Syscall name.
	Syscall *string `json:"syscall,omitempty" bson:"-" mapstructure:"syscall,omitempty"`

	// Date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseAuditReport returns a new  SparseAuditReport.
func NewSparseAuditReport() *SparseAuditReport {
	return &SparseAuditReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAuditReport) Identity() elemental.Identity {

	return AuditReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAuditReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAuditReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseAuditReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAuditReport) ToPlain() elemental.PlainIdentifiable {

	out := NewAuditReport()
	if o.AUID != nil {
		out.AUID = *o.AUID
	}
	if o.CWD != nil {
		out.CWD = *o.CWD
	}
	if o.EGID != nil {
		out.EGID = *o.EGID
	}
	if o.EUID != nil {
		out.EUID = *o.EUID
	}
	if o.EXE != nil {
		out.EXE = *o.EXE
	}
	if o.FSGID != nil {
		out.FSGID = *o.FSGID
	}
	if o.FSUID != nil {
		out.FSUID = *o.FSUID
	}
	if o.GID != nil {
		out.GID = *o.GID
	}
	if o.PER != nil {
		out.PER = *o.PER
	}
	if o.PID != nil {
		out.PID = *o.PID
	}
	if o.PPID != nil {
		out.PPID = *o.PPID
	}
	if o.SGID != nil {
		out.SGID = *o.SGID
	}
	if o.SUID != nil {
		out.SUID = *o.SUID
	}
	if o.UID != nil {
		out.UID = *o.UID
	}
	if o.A0 != nil {
		out.A0 = *o.A0
	}
	if o.A1 != nil {
		out.A1 = *o.A1
	}
	if o.A2 != nil {
		out.A2 = *o.A2
	}
	if o.A3 != nil {
		out.A3 = *o.A3
	}
	if o.Arch != nil {
		out.Arch = *o.Arch
	}
	if o.AuditProfileID != nil {
		out.AuditProfileID = *o.AuditProfileID
	}
	if o.AuditProfileNamespace != nil {
		out.AuditProfileNamespace = *o.AuditProfileNamespace
	}
	if o.Command != nil {
		out.Command = *o.Command
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.Exit != nil {
		out.Exit = *o.Exit
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.ProcessingUnitNamespace != nil {
		out.ProcessingUnitNamespace = *o.ProcessingUnitNamespace
	}
	if o.RecordType != nil {
		out.RecordType = *o.RecordType
	}
	if o.Sequence != nil {
		out.Sequence = *o.Sequence
	}
	if o.Success != nil {
		out.Success = *o.Success
	}
	if o.Syscall != nil {
		out.Syscall = *o.Syscall
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}

	return out
}

// DeepCopy returns a deep copy if the SparseAuditReport.
func (o *SparseAuditReport) DeepCopy() *SparseAuditReport {

	if o == nil {
		return nil
	}

	out := &SparseAuditReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAuditReport.
func (o *SparseAuditReport) DeepCopyInto(out *SparseAuditReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAuditReport: %s", err))
	}

	*out = *target.(*SparseAuditReport)
}
