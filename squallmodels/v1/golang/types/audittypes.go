package types

import (
	"fmt"
	"net/http"

	"github.com/aporeto-inc/elemental"
)

// AuditPolicyRuleList is a list of AuditPolicyRules
type AuditPolicyRuleList []*AuditPolicyRule

// AuditPolicyRule is a generic audit rule
type AuditPolicyRule struct {
	Type     AuditPolicyRuleType `json:"type"`
	Files    *FileWatchRule      `json:"files,omitempty"`
	Syscalls *SyscallRule        `json:"syscalls,omitempty"`
}

// AuditPolicyRuleType specifies the audit rule type.
type AuditPolicyRuleType int

// The rule types supported by this package.
const (
	DeleteAllRuleType      AuditPolicyRuleType = iota + 1 // DeleteAllRule
	FileWatchRuleType                                     // FileWatchRule
	AppendSyscallRuleType                                 // SyscallRule
	PrependSyscallRuleType                                // SyscallRule
)

// Validate validates an audit rule
func (a *AuditPolicyRule) Validate() error {

	switch a.Type {
	case AppendSyscallRuleType, PrependSyscallRuleType:
		if a.Syscalls == nil {
			return elemental.NewError("Validation Error", "Nil syscall rule not allowed", "elemental", http.StatusUnprocessableEntity)
		}
		return a.Syscalls.Validate()
	case FileWatchRuleType:
		if a.Files == nil || a.Syscalls != nil {
			return elemental.NewError("Validation Error", "Nil file watch rule not allowed", "elemental", http.StatusUnprocessableEntity)
		}
		return a.Files.Validate()
	case DeleteAllRuleType:
		if a.Files != nil || a.Syscalls != nil {
			return elemental.NewError("Validation Error", "Delete rule must not have syscall or file", "elemental", http.StatusUnprocessableEntity)
		}
	default:
		return elemental.NewError("Validation Error", "Invalid rule type", "elemental", http.StatusUnprocessableEntity)
	}

	return nil
}

// Rule is the generic interface that all rule types implement.
type Rule interface {
	Validate() error // Validate validates the rule
}

// DeleteAllRule deletes all existing rules.
type DeleteAllRule struct {
}

// Validate validates the rule
func (r *DeleteAllRule) Validate() error {
	return nil
}

// FileWatchRule is used to audit access to particular files or directories
// that you may be interested in.
type FileWatchRule struct {
	Path        string                 `json:"path" bson:"path"`
	Permissions []AuditFilePermissions `json:"permissions" bson:"permissions"`
}

// Validate validates the filewathc rule.
func (r *FileWatchRule) Validate() error {

	if err := elemental.ValidateMaximumLength("path", r.Path, 128, false); err != nil {
		return err
	}

	if len(r.Permissions) > 4 {
		err := elemental.NewError("Validation Error", "Invalid Permissions", "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": "path"}
		return err
	}

	if len(r.Permissions) == 0 {
		r.Permissions = []AuditFilePermissions{AuditFilePermissionsRead, AuditFilePermissionsWrite, AuditFilePermissionsExecute, AuditFilePermissionsAttribute}
	}

	return nil
}

// SyscallRule is used to audit invocations of specific syscalls.
type SyscallRule struct {
	List     AuditFilterListType   `json:"list"`
	Action   AuditFilterActionType `json:"action"`
	Filters  []AuditFilterSpec     `json:"filters"`
	Syscalls []AuditSystemCallType `json:"syscalls"`
}

// Validate validates the filewathc rule.
func (r *SyscallRule) Validate() error {
	return nil
}

// AccessType specifies the type of file access to audit.
type AccessType uint8

// The access types that can be audited for file watches.
const (
	ReadAccessType AccessType = iota + 1
	WriteAccessType
	ExecuteAccessType
	AttributeChangeAccessType
)

var accessTypeName = map[AccessType]string{
	ReadAccessType:            "read",
	WriteAccessType:           "write",
	ExecuteAccessType:         "execute",
	AttributeChangeAccessType: "attribute",
}

func (t AccessType) String() string {
	name, found := accessTypeName[t]
	if found {
		return name
	}
	return "unknown"
}

// AuditFilterKind specifies a type of filter to apply to a syscall rule.
type AuditFilterKind uint8

// The type of filters that can be applied.
const (
	AuditFilterKindInterFieldFilter AuditFilterKind = iota + 1 // Inter-field comparison filtering (-C).
	AuditFilteRKindValueFilter                                 // Filtering based on values (-F).
)

// AuditFilterSpec defines a filter to apply to a syscall rule.
type AuditFilterSpec struct {
	Kind       AuditFilterKind     `json:"kind"`
	LHS        AuditFilterType     `json:"lhs"`
	Comparator AuditFilterOperator `json:"comparator"`
	RHS        string              `json:"rhs"`
}

func (f *AuditFilterSpec) String() string {
	return fmt.Sprintf("%v %v %v", f.LHS, f.Comparator, f.RHS)
}
