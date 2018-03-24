package types

import (
	"net/http"

	"github.com/aporeto-inc/elemental"
)

// SyscallEnforcementRulesMap is a list of SyscallEnforcementRule.
type SyscallEnforcementRulesMap map[AuditSystemCallType]*SyscallEnforcementRule

// Validate validates a SyscallEnforcementRulesMap
func (s SyscallEnforcementRulesMap) Validate() error {
	var errs []error

	for r := range s {
		if err := r.Validate("syscall"); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return elemental.NewErrors(errs...)
	}

	return nil
}

// SyscallEnforcementRule  is a rule to match a syscall in Seccomp.
type SyscallEnforcementRule struct {
	DefaultAction SyscallEnforcementAction         `json:"action"`
	Args          map[uint]*SyscallEnforcermentArg `json:"args"`
}

// Validate validates a syscall enforcement rule.
func (s *SyscallEnforcementRule) Validate() error {
	var errs []error

	if err := s.DefaultAction.Validate(); err != nil {
		errs = append(errs, err)
	}

	for _, s := range s.Args {
		if err := s.Validate(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 1 {
		return elemental.NewErrors(errs...)
	}
	return nil
}

// SyscallEnforcermentArg is a rule to match a specific syscall argument in Seccomp.
type SyscallEnforcermentArg struct {
	Index    uint
	Value    uint64                     `json:"value"`
	ValueTwo uint64                     `json:"valueTwo"`
	Op       SyscallEnforcementOperator `json:"op"`
	Action   SyscallEnforcementAction
}

// Validate validates the syscall enforcement arguments.
func (s *SyscallEnforcermentArg) Validate() error {
	if err := s.Op.Validate(); err != nil {
		return elemental.NewError("Validation Error", "Unknown syscall operator", "elemental", http.StatusUnprocessableEntity)
	}

	if err := s.Action.Validate(); err != nil {
		return elemental.NewError("Validation Error", "Unknown syscall enforcement action", "elemental", http.StatusUnprocessableEntity)
	}

	return nil
}

// SyscallEnforcementOperator is a comparison operator to be used when matching syscall arguments in Seccomp./
type SyscallEnforcementOperator int

// Values of SyscallEnforcementOperator.
const (
	SyscallEnforcementOperatorEqualTo SyscallEnforcementOperator = iota
	SyscallEnforcementOperatorNotEqualTo
	SyscallEnforcementOperatorGreaterThan
	SyscallEnforcementOperatorGreaterThanOrEqualTo
	SyscallEnforcementOperatorLessThan
	SyscallEnforcementOperatorLessThanOrEqualTo
	SyscallEnforcementOperatorMaskEqualTo
)

// Validate validates the syscall enforcement operator
func (s SyscallEnforcementOperator) Validate() error {
	if s < SyscallEnforcementOperatorEqualTo || s > SyscallEnforcementOperatorMaskEqualTo {
		return elemental.NewError("Validation Error", "Unknown operator", "elemental", http.StatusUnprocessableEntity)
	}
	return nil
}

// SyscallEnforcementAction is the action type.
type SyscallEnforcementAction int

// Values of SyscallEnforcementAction.
const (
	SyscallEnforcementActionKill SyscallEnforcementAction = iota
	SyscallEnforcementActionErrno
	SyscallEnforcementActionTrap
	SyscallEnforcementActionAllow
	SyscallEnforcementActionTrace
)

// Validate validates a syslcall enforcement action.
func (s SyscallEnforcementAction) Validate() error {
	if s < SyscallEnforcementActionKill || s > SyscallEnforcementActionTrace {
		return elemental.NewError("Validation Error", "Unknown syscall enforcement action", "elemental", http.StatusUnprocessableEntity)
	}

	return nil
}

// CapabilitiesType is the type of capabilities.
type CapabilitiesType string

// Values of CapabilitiesType.
const (
	CapabilitiesTypeAuditControl   CapabilitiesType = "AUDIT_CONTROL"
	CapabilitiesTypeAuditRead      CapabilitiesType = "AUDIT_READ"
	CapabilitiesTypeAuditWrite     CapabilitiesType = "AUDIT_WRITE"
	CapabilitiesTypeBlockSuspend   CapabilitiesType = "BLOCK_SUSPEND"
	CapabilitiesTypeChown          CapabilitiesType = "CHOWN"
	CapabilitiesTypeDacOverride    CapabilitiesType = "DAC_OVERRIDE"
	CapabilitiesTypeReadSearch     CapabilitiesType = "DAC_READ_SEARCH"
	CapabilitiesTypeFowner         CapabilitiesType = "FOWNER"
	CapabilitiesTypeFsetid         CapabilitiesType = "FSETID"
	CapabilitiesTypeIPCLock        CapabilitiesType = "IPC_LOCK"
	CapabilitiesTypeIPCOwner       CapabilitiesType = "IPC_OWNER"
	CapabilitiesTypeKill           CapabilitiesType = "KILL"
	CapabilitiesTypeLease          CapabilitiesType = "LEASE"
	CapabilitiesTypeLinuxImmutable CapabilitiesType = "LINUX_IMMUTABLE"
	CapabilitiesTypeMacAdmin       CapabilitiesType = "MAC_ADMIN"
	CapabilitiesTypeMacOverride    CapabilitiesType = "MAC_OVERRIDE"
	CapabilitiesTypeMknod          CapabilitiesType = "MKNOD"
	CapabilitiesTypeNetAdmin       CapabilitiesType = "NET_ADMIN"
	CapabilitiesTypeNetBindService CapabilitiesType = "NET_BIND_SERVICE"
	CapabilitiesTypeNetBroadcast   CapabilitiesType = "NET_BROADCAST"
	CapabilitiesTypeNetRaw         CapabilitiesType = "NET_RAW"
	CapabilitiesTypeSetGid         CapabilitiesType = "SETGID"
	CapabilitiesTypeSetFcap        CapabilitiesType = "SETFCAP"
	CapabilitiesTypeSetPcap        CapabilitiesType = "SETPCAP"
	CapabilitiesTypeSetUID         CapabilitiesType = "SETUID"
	CapabilitiesTypeSysAdmin       CapabilitiesType = "SYS_ADMIN"
	CapabilitiesTypeSysBoot        CapabilitiesType = "SYS_BOOT"
	CapabilitiesTypeSysChroot      CapabilitiesType = "SYS_CHROOT"
	CapabilitiesTypeSysModule      CapabilitiesType = "SYS_MODULE"
	CapabilitiesTypeSysNice        CapabilitiesType = "SYS_NICE"
	CapabilitiesTypeSysPacct       CapabilitiesType = "SYS_PACCT"
	CapabilitiesTypeSysPtrace      CapabilitiesType = "SYS_PTRACE"
	CapabilitiesTypeSysRawIO       CapabilitiesType = "SYS_RAWIO"
	CapabilitiesTypeSysResource    CapabilitiesType = "SYS_RESOURCE"
	CapabilitiesTypeSysTime        CapabilitiesType = "SYS_TIME"
	CapabilitiesTypeSysTTYConfig   CapabilitiesType = "SYS_TTY_CONFIG"
	CapabilitiesTypeCapSyslog      CapabilitiesType = "SYSLOG"
	CapabilitiesTypeWakeAlarm      CapabilitiesType = "WAKE_ALARM"
)

var reverseCapabilitiesMap = map[string]interface{}{
	"AUDIT_CONTROL":    nil,
	"AUDIT_READ":       nil,
	"AUDIT_WRITE":      nil,
	"BLOCK_SUSPEND":    nil,
	"CHOWN":            nil,
	"DAC_OVERRIDE":     nil,
	"DAC_READ_SEARCH":  nil,
	"FOWNER":           nil,
	"FSETID":           nil,
	"IPC_LOCK":         nil,
	"IPC_OWNER":        nil,
	"KILL":             nil,
	"LEASE":            nil,
	"LINUX_IMMUTABLE":  nil,
	"MAC_ADMIN":        nil,
	"MAC_OVERRIDE":     nil,
	"MKNOD":            nil,
	"NET_ADMIN":        nil,
	"NET_BIND_SERVICE": nil,
	"NET_BROADCAST":    nil,
	"NET_RAW":          nil,
	"SETGID":           nil,
	"SETFCAP":          nil,
	"SETPCAP":          nil,
	"SETUID":           nil,
	"SYS_ADMIN":        nil,
	"SYS_BOOT":         nil,
	"SYS_CHROOT":       nil,
	"SYS_MODULE":       nil,
	"SYS_NICE":         nil,
	"SYS_PACCT":        nil,
	"SYS_PTRACE":       nil,
	"SYS_RAWIO":        nil,
	"SYS_RESOURCE":     nil,
	"SYS_TIME":         nil,
	"SYS_TTY_CONFIG":   nil,
	"SYSLOG":           nil,
	"WAKE_ALARM":       nil,
}

// Validate validates the capabilities.
func (c CapabilitiesType) Validate() error {
	return elemental.ValidateStringInMap("capabilities", string(c), reverseCapabilitiesMap, false)
}

// CapabilitiesActionType is add or drop
type CapabilitiesActionType int

// Values for CapabilitiesActionType
const (
	CapabilitiesActionTypeAdd CapabilitiesActionType = iota
	CapabilitiesActionTypeDrop
)

// CapabilitiesTypeMap is a list of capabilities.
type CapabilitiesTypeMap map[CapabilitiesType]CapabilitiesActionType

// Validate validates a capabilities type list.
func (c CapabilitiesTypeMap) Validate() error {
	var errs []error

	for k, v := range c {
		if err := k.Validate(); err != nil {
			errs = append(errs, err)
		}

		if v < CapabilitiesActionTypeAdd || v > CapabilitiesActionTypeDrop {
			errs = append(errs, elemental.NewError("Validation Error", "Unknown capabilities action type", "elemental", http.StatusUnprocessableEntity))
		}
	}

	if len(errs) > 1 {
		return elemental.NewErrors(errs...)
	}

	return nil
}

// ArchitecturesType is the type for different architectures supported.
type ArchitecturesType string

// Values of ArchitecturesType.
const (
	ArchitectureTypeX86         ArchitecturesType = "x86"
	ArchitectureTypeX86_64      ArchitecturesType = "amd64"
	ArchitectureTypeX32         ArchitecturesType = "x32"
	ArchitectureTypeARM         ArchitecturesType = "arm"
	ArchitectureTypeAARCH64     ArchitecturesType = "arm64"
	ArchitectureTypeMIPS        ArchitecturesType = "mips"
	ArchitectureTypeMIPS64      ArchitecturesType = "mips64"
	ArchitectureTypeMIPS64N32   ArchitecturesType = "mips64n32"
	ArchitectureTypeMIPSEL      ArchitecturesType = "mipsel"
	ArchitectureTypeMIPSEL64    ArchitecturesType = "mipsel64"
	ArchitectureTypeMIPSEL64N32 ArchitecturesType = "mipsel64n32"
	ArchitectureTypePPC         ArchitecturesType = "ppc"
	ArchitectureTypePPC64       ArchitecturesType = "ppc64"
	ArchitectureTypePPC64LE     ArchitecturesType = "ppc64le"
	ArchitectureTypeS390        ArchitecturesType = "s390"
	ArchitectureTypeS390X       ArchitecturesType = "s390x"
)

var reverseArchitecturesMap = map[string]interface{}{
	"x86":         nil,
	"amd64":       nil,
	"x32":         nil,
	"arm":         nil,
	"arm64":       nil,
	"mips":        nil,
	"mips64":      nil,
	"mips64n32":   nil,
	"mipsel":      nil,
	"mipsel64":    nil,
	"mipsel64n32": nil,
	"ppc":         nil,
	"ppc64":       nil,
	"ppc64le":     nil,
	"s390":        nil,
	"s390x":       nil,
}

// Validate validates the architectures.
func (a ArchitecturesType) Validate() error {
	return elemental.ValidateStringInMap("capabilities", string(a), reverseArchitecturesMap, false)
}

// ArchitecturesTypeList is a list of ArchitectureTypes.
type ArchitecturesTypeList []ArchitecturesType

// Validate validates an architectures type list.
func (a ArchitecturesTypeList) Validate() error {
	var errs []error

	for _, s := range a {
		if err := s.Validate(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 1 {
		return elemental.NewErrors(errs...)
	}

	return nil
}
