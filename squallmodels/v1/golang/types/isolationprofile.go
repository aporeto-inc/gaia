package types

// SyscallEnforcementRulesList is a list of SyscallEnforcementRule
type SyscallEnforcementRulesList []*SyscallEnforcementRule

// SyscallEnforcementRule  is a rule to match a syscall in Seccomp
type SyscallEnforcementRule struct {
	Name   []AuditSystemCallType     `json:"name"`
	Action SyscallEnforcementAction  `json:"action"`
	Args   []*SyscallEnforcermentArg `json:"args"`
}

// SyscallEnforcermentArg is a rule to match a specific syscall argument in Seccomp
type SyscallEnforcermentArg struct {
	Index    uint                       `json:"index"`
	Value    uint64                     `json:"value"`
	ValueTwo uint64                     `json:"valueTwo"`
	Op       SyscallEnforcementOperator `json:"op"`
}

// SyscallEnforcementOperator is a comparison operator to be used when matching syscall arguments in Seccomp
type SyscallEnforcementOperator int

// Values of SyscallEnforcementOperator
const (
	SyscallEnforcementOperatorEqualTo SyscallEnforcementOperator = iota
	SyscallEnforcementOperatorNotEqualTo
	SyscallEnforcementOperatorGreaterThan
	SyscallEnforcementOperatorGreaterThanOrEqualTo
	SyscallEnforcementOperatorLessThan
	SyscallEnforcementOperatorLessThanOrEqualTo
	SyscallEnforcementOperatorMaskEqualTo
)

// SyscallEnforcementAction is the action type
type SyscallEnforcementAction int

// Values of SyscallEnforcementAction
const (
	SyscallEnforcementActionKill SyscallEnforcementAction = iota
	SyscallEnforcementActionErrno
	SyscallEnforcementActionTrap
	SyscallEnforcementActionAllow
	SyscallEnforcementActionTrace
)

// CapabilitiesType is the type of capabilities
type CapabilitiesType string

// Values of CapabilitiesType
const (
	CapabilitiesTypeCapAdmin CapabilitiesType = "CAP_ADMIN"
)

// ArchitecturesType is the type for different architectures supported
type ArchitecturesType string

// Values of ArchitecturesType
const (
	ArchitecturesTypeX86 ArchitecturesType = "X86"
)
