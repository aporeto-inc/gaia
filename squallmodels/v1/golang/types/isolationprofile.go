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
	CapabilitiesTypeAuditControl   CapabilitiesType = "CAP_AUDIT_CONTROL"
	CapabilitiesTypeAuditRead      CapabilitiesType = "CAP_AUDIT_READ"
	CapabilitiesTypeAuditWrite     CapabilitiesType = "CAP_AUDIT_WRITE"
	CapabilitiesTypeBlockSuspend   CapabilitiesType = "CAP_BLOCK_SUSPEND"
	CapabilitiesTypeChown          CapabilitiesType = "CAP_CHOWN"
	CapabilitiesTypeDacOverride    CapabilitiesType = "CAP_DAC_OVERRIDE"
	CapabilitiesTypeReadSearch     CapabilitiesType = "CAP_DAC_READ_SEARCH"
	CapabilitiesTypeFowner         CapabilitiesType = "CAP_FOWNER"
	CapabilitiesTypeFsetid         CapabilitiesType = "CAP_FSETID"
	CapabilitiesTypeIPCLock        CapabilitiesType = "CAP_IPC_LOCK"
	CapabilitiesTypeIPCOwner       CapabilitiesType = "CAP_IPC_OWNER"
	CapabilitiesTypeKill           CapabilitiesType = "CAP_KILL"
	CapabilitiesTypeLease          CapabilitiesType = "CAP_LEASE"
	CapabilitiesTypeLinuxImmutable CapabilitiesType = "CAP_LINUX_IMMUTABLE"
	CapabilitiesTypeMacAdmin       CapabilitiesType = "CAP_MAC_ADMIN"
	CapabilitiesTypeMacOverride    CapabilitiesType = "CAP_MAC_OVERRIDE"
	CapabilitiesTypeMknod          CapabilitiesType = "CAP_MKNOD"
	CapabilitiesTypeNetAdmin       CapabilitiesType = "CAP_NET_ADMIN"
	CapabilitiesTypeNetBindService CapabilitiesType = "CAP_NET_BIND_SERVICE"
	CapabilitiesTypeNetBroadcast   CapabilitiesType = "CAP_NET_BROADCAST"
	CapabilitiesTypeNetRaw         CapabilitiesType = "CAP_NET_RAW"
	CapabilitiesTypeSetGid         CapabilitiesType = "CAP_SETGID"
	CapabilitiesTypeSetFcap        CapabilitiesType = "CAP_SETFCAP"
	CapabilitiesTypeSetPcap        CapabilitiesType = "CAP_SETPCAP"
	CapabilitiesTypeSetUID         CapabilitiesType = "CAP_SETUID"
	CapabilitiesTypeSysAdmin       CapabilitiesType = "CAP_SYS_ADMIN"
	CapabilitiesTypeSysBoot        CapabilitiesType = "CAP_SYS_BOOT"
	CapabilitiesTypeSysChroot      CapabilitiesType = "CAP_SYS_CHROOT"
	CapabilitiesTypeSysModule      CapabilitiesType = "CAP_SYS_MODULE"
	CapabilitiesTypeSysNice        CapabilitiesType = "CAP_SYS_NICE"
	CapabilitiesTypeSysPacct       CapabilitiesType = "CAP_SYS_PACCT"
	CapabilitiesTypeSysPtrace      CapabilitiesType = "CAP_SYS_PTRACE"
	CapabilitiesTypeSysRawIO       CapabilitiesType = "CAP_SYS_RAWIO"
	CapabilitiesTypeSysResource    CapabilitiesType = "CAP_SYS_RESOURCE"
	CapabilitiesTypeSysTime        CapabilitiesType = "CAP_SYS_TIME"
	CapabilitiesTypeSysTTYConfig   CapabilitiesType = "CAP_SYS_TTY_CONFIG"
	CapabilitiesTypeCapSyslog      CapabilitiesType = "CAP_SYSLOG"
	CapabilitiesTypeWakeAlarm      CapabilitiesType = "CAP_WAKE_ALARM"
)

var reverseCapabilitiesMap = map[string]interface{}{
	"CAP_AUDIT_CONTROL":    nil,
	"CAP_AUDIT_READ":       nil,
	"CAP_AUDIT_WRITE":      nil,
	"CAP_BLOCK_SUSPEND":    nil,
	"CAP_CHOWN":            nil,
	"CAP_DAC_OVERRIDE":     nil,
	"CAP_DAC_READ_SEARCH":  nil,
	"CAP_FOWNER":           nil,
	"CAP_FSETID":           nil,
	"CAP_IPC_LOCK":         nil,
	"CAP_IPC_OWNER":        nil,
	"CAP_KILL":             nil,
	"CAP_LEASE":            nil,
	"CAP_LINUX_IMMUTABLE":  nil,
	"CAP_MAC_ADMIN":        nil,
	"CAP_MAC_OVERRIDE":     nil,
	"CAP_MKNOD":            nil,
	"CAP_NET_ADMIN":        nil,
	"CAP_NET_BIND_SERVICE": nil,
	"CAP_NET_BROADCAST":    nil,
	"CAP_NET_RAW":          nil,
	"CAP_SETGID":           nil,
	"CAP_SETFCAP":          nil,
	"CAP_SETPCAP":          nil,
	"CAP_SETUID":           nil,
	"CAP_SYS_ADMIN":        nil,
	"CAP_SYS_BOOT":         nil,
	"CAP_SYS_CHROOT":       nil,
	"CAP_SYS_MODULE":       nil,
	"CAP_SYS_NICE":         nil,
	"CAP_SYS_PACCT":        nil,
	"CAP_SYS_PTRACE":       nil,
	"CAP_SYS_RAWIO":        nil,
	"CAP_SYS_RESOURCE":     nil,
	"CAP_SYS_TIME":         nil,
	"CAP_SYS_TTY_CONFIG":   nil,
	"CAP_SYSLOG":           nil,
	"CAP_WAKE_ALARM":       nil,
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
	ArchitectureTYpeX86         ArchitecturesType = "x86"
	ArchitectureTYpeX86_64      ArchitecturesType = "amd64"
	ArchitectureTYpeX32         ArchitecturesType = "x32"
	ArchitectureTYpeARM         ArchitecturesType = "arm"
	ArchitectureTYpeAARCH64     ArchitecturesType = "arm64"
	ArchitectureTYpeMIPS        ArchitecturesType = "mips"
	ArchitectureTYpeMIPS64      ArchitecturesType = "mips64"
	ArchitectureTYpeMIPS64N32   ArchitecturesType = "mips64n32"
	ArchitectureTYpeMIPSEL      ArchitecturesType = "mipsel"
	ArchitectureTYpeMIPSEL64    ArchitecturesType = "mipsel64"
	ArchitectureTYpeMIPSEL64N32 ArchitecturesType = "mipsel64n32"
	ArchitectureTYpePPC         ArchitecturesType = "ppc"
	ArchitectureTYpePPC64       ArchitecturesType = "ppc64"
	ArchitectureTYpePPC64LE     ArchitecturesType = "ppc64le"
	ArchitectureTYpeS390        ArchitecturesType = "s390"
	ArchitectureTYpeS390X       ArchitecturesType = "s390x"
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
