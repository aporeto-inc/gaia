package types

import (
	"fmt"

	"github.com/aporeto-inc/elemental"
)

// AuditFilterType are the valid IDs of the audit filters.
type AuditFilterType string

// Values of AuditFilterType.
const (
	AuditFilterTypeA0         AuditFilterType = "a0"
	AuditFilterTypeA1         AuditFilterType = "a1"
	AuditFilterTypeA2         AuditFilterType = "a2"
	AuditFilterTypeA3         AuditFilterType = "a3"
	AuditFilterTypeArch       AuditFilterType = "arch"
	AuditFilterTypeDevMajor   AuditFilterType = "devmajor"
	AuditFilterTypeDevMinor   AuditFilterType = "devminor"
	AuditFilterTypeDir        AuditFilterType = "dir"
	AuditFilterTypeEgid       AuditFilterType = "egid"
	AuditFilterTypeEuid       AuditFilterType = "euid"
	AuditFilterTypeExit       AuditFilterType = "exit"
	AuditFilterTypeFiletype   AuditFilterType = "filetye"
	AuditFilterTypeFsgid      AuditFilterType = "fsgid"
	AuditFilterTypeFsuid      AuditFilterType = "fsuid"
	AuditFilterTypeGid        AuditFilterType = "gid"
	AuditFilterTypeInode      AuditFilterType = "inode"
	AuditFilterTypeMsgtype    AuditFilterType = "msgtype"
	AuditFilterTypeObjgid     AuditFilterType = "obj_gid"
	AuditFilterTypeObjlevhigh AuditFilterType = "obj_lev_highj"
	AuditFilterTypeObjlevlow  AuditFilterType = "obj_lev_low"
	AuditFilterTypeObjrole    AuditFilterType = "obj_role"
	AuditFilterTypeObjtype    AuditFilterType = "obj_type"
	AuditFilterTypeObjuid     AuditFilterType = "obj_uid"
	AuditFilterTypeObjuser    AuditFilterType = "obj_user"
	AuditFilterTypePath       AuditFilterType = "path"
	AuditFilterTypePerm       AuditFilterType = "perm"
	AuditFilterTypePers       AuditFilterType = "pers"
	AuditFilterTypePid        AuditFilterType = "pid"
	AuditFilterTypePpid       AuditFilterType = "ppid"
	AuditFilterTypeSgid       AuditFilterType = "sgid"
	AuditFilterTypeSubclr     AuditFilterType = "subj_clr"
	AuditFilterTypeSubjrole   AuditFilterType = "subj_role"
	AuditFilterTypeSubjtype   AuditFilterType = "subj_type"
	AuditFilterTypeSubsen     AuditFilterType = "subj_sen"
	AuditFilterTypeSubuser    AuditFilterType = "subj_user"
	AuditFilterTypeSuccess    AuditFilterType = "success"
	AuditFilterTypeSuid       AuditFilterType = "suid"
	AuditFilterTypeUserid     AuditFilterType = "uid"
	AuditFilterTypeAuid       AuditFilterType = "auid"
	AuditFilterTypeKey        AuditFilterType = "key"
	AuditFilterTypeExe        AuditFilterType = "exe"
)

var auditFilterTypeReverse = map[string]interface{}{
	"a0":            AuditFilterTypeA0,
	"a1":            AuditFilterTypeA1,
	"a2":            AuditFilterTypeA2,
	"a3":            AuditFilterTypeA3,
	"arch":          AuditFilterTypeArch,
	"devmajor":      AuditFilterTypeDevMajor,
	"devminor":      AuditFilterTypeDevMinor,
	"dir":           AuditFilterTypeDir,
	"egid":          AuditFilterTypeEgid,
	"euid":          AuditFilterTypeEuid,
	"exit":          AuditFilterTypeExit,
	"filetye":       AuditFilterTypeFiletype,
	"fsgid":         AuditFilterTypeFsgid,
	"fsuid":         AuditFilterTypeFsuid,
	"gid":           AuditFilterTypeGid,
	"inode":         AuditFilterTypeInode,
	"msg_type":      AuditFilterTypeMsgtype,
	"obj_gid":       AuditFilterTypeObjgid,
	"obj_lev_highj": AuditFilterTypeObjlevhigh,
	"obj_lev_low":   AuditFilterTypeObjlevlow,
	"obj_role":      AuditFilterTypeObjrole,
	"obj_type":      AuditFilterTypeObjtype,
	"obj_uid":       AuditFilterTypeObjuid,
	"obj_user":      AuditFilterTypeObjuser,
	"path":          AuditFilterTypePath,
	"perm":          AuditFilterTypePerm,
	"pers":          AuditFilterTypePers,
	"pid":           AuditFilterTypePid,
	"ppid":          AuditFilterTypePpid,
	"sgid":          AuditFilterTypeSgid,
	"subj_clr":      AuditFilterTypeSubclr,
	"subj_role":     AuditFilterTypeSubjrole,
	"subj_sen":      AuditFilterTypeSubsen,
	"subj_type":     AuditFilterTypeSubjtype,
	"subj_user":     AuditFilterTypeSubuser,
	"success":       AuditFilterTypeSuccess,
	"suid":          AuditFilterTypeSuid,
	"uid":           AuditFilterTypeUserid,
	"auid":          AuditFilterTypeAuid,
	"key":           AuditFilterTypeKey,
	"exe":           AuditFilterTypeExe,
}

// Validate validates the AuditFilterType
func (a AuditFilterType) Validate(attribute string) error {

	return elemental.ValidateStringInMap(attribute, string(a), auditFilterTypeReverse, false)
}

// AuditFilterTypeFromString returns the AuditFilterType from a given string value.
func AuditFilterTypeFromString(value string) (AuditFilterType, error) {

	t, ok := auditFilterTypeReverse[value]
	if !ok {
		return "", fmt.Errorf("unknown AuditFilterType %s", value)
	}

	return t.(AuditFilterType), nil
}

// AuditFilterListType are the types allowed in the list argument of audit filters
type AuditFilterListType string

// Values of AuditFilterListType
const (
	AuditFilterListTypeTask    AuditFilterListType = "task"
	AuditFilterListTypeExit    AuditFilterListType = "exit"
	AuditFilterListTypeUser    AuditFilterListType = "user"
	AuditFilterListTypeExclude AuditFilterListType = "exclude"
)

var auditFilterListTypeReverse = map[string]interface{}{
	"task":    AuditFilterListTypeTask,
	"exit":    AuditFilterListTypeExit,
	"user":    AuditFilterListTypeUser,
	"exclude": AuditFilterListTypeExclude,
}

// Validate validates the AuditFilterListType
func (a AuditFilterListType) Validate(attribute string) error {

	return elemental.ValidateStringInMap(attribute, string(a), auditFilterListTypeReverse, false)
}

// AuditFilterListTypeFromString returns the AuditFilterListType from a given string value.
func AuditFilterListTypeFromString(value string) (AuditFilterListType, error) {

	t, ok := auditFilterListTypeReverse[value]
	if !ok {
		return "", fmt.Errorf("unknown AuditFilterListType %s", value)
	}

	return t.(AuditFilterListType), nil
}

// AuditFilterActionType are the types allowed in the audit filter action
type AuditFilterActionType string

// Values of AuditFilterActionType
const (
	AuditFilterActionTypeNever  AuditFilterActionType = "never"
	AuditFilterActionTypeAlways AuditFilterActionType = "always"
)

var auditFilterActionTypeReverse = map[string]interface{}{
	"never":  AuditFilterActionTypeNever,
	"always": AuditFilterActionTypeAlways,
}

// Validate validates the AuditFilterActionType
func (a AuditFilterActionType) Validate(attribute string) error {

	return elemental.ValidateStringInMap(attribute, string(a), auditFilterActionTypeReverse, false)
}

// AuditFilterActionTypeFromString returns the AuditFilterActionType from a given string value.
func AuditFilterActionTypeFromString(value string) (AuditFilterActionType, error) {

	t, ok := auditFilterActionTypeReverse[value]
	if !ok {
		return "", fmt.Errorf("unknown AuditFilterActionType %s", value)
	}

	return t.(AuditFilterActionType), nil
}

// AuditFilterOperator is the operator for filters.
type AuditFilterOperator string

// Values of AuditFilterOperator
const (
	AuditFilterOperatorEqual           AuditFilterOperator = "="
	AuditFilterOperatorNotEqual        AuditFilterOperator = "!="
	AuditFilterOperatorGreater         AuditFilterOperator = ">"
	AuditFilterOperatorLessThan        AuditFilterOperator = "<"
	AuditFilterOperatorGreaterOrEqual  AuditFilterOperator = ">="
	AuditFilterOperatorLessThanOrEqual AuditFilterOperator = "<="
	AuditFilterOperatorBitMask         AuditFilterOperator = "&"
	AuditFilterOperatorBitTest         AuditFilterOperator = "&="
)

var auditFilterOperatorReverse = map[string]interface{}{
	"=":  AuditFilterOperatorEqual,
	"!=": AuditFilterOperatorNotEqual,
	">":  AuditFilterOperatorGreater,
	"<":  AuditFilterOperatorLessThan,
	">=": AuditFilterOperatorGreaterOrEqual,
	"<=": AuditFilterOperatorLessThanOrEqual,
	"&":  AuditFilterOperatorBitMask,
	"&=": AuditFilterOperatorBitTest,
}

// Validate validates the audit filter operator
func (a AuditFilterOperator) Validate(attribute string) error {

	return elemental.ValidateStringInMap(attribute, string(a), auditFilterOperatorReverse, false)
}

// AuditFilterOperatorFromString returns the AuditFilterOperator from a given string value.
func AuditFilterOperatorFromString(value string) (AuditFilterOperator, error) {

	t, ok := auditFilterOperatorReverse[value]
	if !ok {
		return "", fmt.Errorf("unknown AuditFilterOperator %s", value)
	}

	return t.(AuditFilterOperator), nil
}

// AuditFilePermissions is the type of file permissions
type AuditFilePermissions string

// Values of AuditFilePermissions
const (
	AuditFilePermissionsWrite     AuditFilePermissions = "w"
	AuditFilePermissionsRead      AuditFilePermissions = "r"
	AuditFilePermissionsExecute   AuditFilePermissions = "x"
	AuditFilePermissionsAttribute AuditFilePermissions = "a"
)

// Validate validates the audit file permissions
func (a AuditFilePermissions) Validate(attribute string) error {

	enums := []string{
		string(AuditFilePermissionsWrite),
		string(AuditFilePermissionsRead),
		string(AuditFilePermissionsExecute),
		string(AuditFilePermissionsAttribute),
	}

	return elemental.ValidateStringInList(attribute, string(a), enums, false)
}

// AuditSystemCallType is the type for the system calls.
type AuditSystemCallType string

// Validate validates the AuditSystemCallType
func (a AuditSystemCallType) Validate(attribute string) error {

	return elemental.ValidateStringInMap(attribute, string(a), auditSystemCallTypeReverse, false)
}

// AuditSystemCallTypeFromString returns the AuditSystemCallType from a given string value.
func AuditSystemCallTypeFromString(value string) (AuditSystemCallType, error) {

	t, ok := auditSystemCallTypeReverse[value]
	if !ok {
		return "", fmt.Errorf("unknown AuditSystemCallType %s", value)
	}

	return t.(AuditSystemCallType), nil
}

// Values of AuditSystemCallType.
const (
	AuditSystemCallTypeREAD                   AuditSystemCallType = "read"
	AuditSystemCallTypeWRITE                  AuditSystemCallType = "write"
	AuditSystemCallTypeOPEN                   AuditSystemCallType = "open"
	AuditSystemCallTypeCLOSE                  AuditSystemCallType = "close"
	AuditSystemCallTypeSTAT                   AuditSystemCallType = "stat"
	AuditSystemCallTypeFSTAT                  AuditSystemCallType = "fstat"
	AuditSystemCallTypeLSTAT                  AuditSystemCallType = "lstat"
	AuditSystemCallTypePOLL                   AuditSystemCallType = "poll"
	AuditSystemCallTypeLSEEK                  AuditSystemCallType = "lseek"
	AuditSystemCallTypeMMAP                   AuditSystemCallType = "mmap"
	AuditSystemCallTypeMPROTECT               AuditSystemCallType = "mprotect"
	AuditSystemCallTypeMUNMAP                 AuditSystemCallType = "munmap"
	AuditSystemCallTypeBRK                    AuditSystemCallType = "brk"
	AuditSystemCallTypeRT_SIGACTION           AuditSystemCallType = "rt_sigaction"
	AuditSystemCallTypeRT_SIGPROCMASK         AuditSystemCallType = "rt_sigprocmask"
	AuditSystemCallTypeRT_SIGRETURN           AuditSystemCallType = "rt_sigreturn"
	AuditSystemCallTypeIOCTL                  AuditSystemCallType = "ioctl"
	AuditSystemCallTypePREAD64                AuditSystemCallType = "pread64"
	AuditSystemCallTypePWRITE64               AuditSystemCallType = "pwrite64"
	AuditSystemCallTypeREADV                  AuditSystemCallType = "readv"
	AuditSystemCallTypeWRITEV                 AuditSystemCallType = "writev"
	AuditSystemCallTypeACCESS                 AuditSystemCallType = "access"
	AuditSystemCallTypePIPE                   AuditSystemCallType = "pipe"
	AuditSystemCallTypeSELECT                 AuditSystemCallType = "select"
	AuditSystemCallTypeSCHED_YIELD            AuditSystemCallType = "sched_yield"
	AuditSystemCallTypeMREMAP                 AuditSystemCallType = "mremap"
	AuditSystemCallTypeMSYNC                  AuditSystemCallType = "msync"
	AuditSystemCallTypeMINCORE                AuditSystemCallType = "mincore"
	AuditSystemCallTypeMADVISE                AuditSystemCallType = "madvise"
	AuditSystemCallTypeSHMGET                 AuditSystemCallType = "shmget"
	AuditSystemCallTypeSHMAT                  AuditSystemCallType = "shmat"
	AuditSystemCallTypeSHMCTL                 AuditSystemCallType = "shmctl"
	AuditSystemCallTypeDUP                    AuditSystemCallType = "dup"
	AuditSystemCallTypeDUP2                   AuditSystemCallType = "dup2"
	AuditSystemCallTypePAUSE                  AuditSystemCallType = "pause"
	AuditSystemCallTypeNANOSLEEP              AuditSystemCallType = "nanosleep"
	AuditSystemCallTypeGETITIMER              AuditSystemCallType = "getitimer"
	AuditSystemCallTypeALARM                  AuditSystemCallType = "alarm"
	AuditSystemCallTypeSETITIMER              AuditSystemCallType = "setitimer"
	AuditSystemCallTypeGETPID                 AuditSystemCallType = "getpid"
	AuditSystemCallTypeSENDFILE               AuditSystemCallType = "sendfile"
	AuditSystemCallTypeSOCKET                 AuditSystemCallType = "socket"
	AuditSystemCallTypeCONNECT                AuditSystemCallType = "connect"
	AuditSystemCallTypeACCEPT                 AuditSystemCallType = "accept"
	AuditSystemCallTypeSENDTO                 AuditSystemCallType = "sendto"
	AuditSystemCallTypeRECVFROM               AuditSystemCallType = "recvfrom"
	AuditSystemCallTypeSENDMSG                AuditSystemCallType = "sendmsg"
	AuditSystemCallTypeRECVMSG                AuditSystemCallType = "recvmsg"
	AuditSystemCallTypeSHUTDOWN               AuditSystemCallType = "shutdown"
	AuditSystemCallTypeBIND                   AuditSystemCallType = "bind"
	AuditSystemCallTypeLISTEN                 AuditSystemCallType = "listen"
	AuditSystemCallTypeGETSOCKNAME            AuditSystemCallType = "getsockname"
	AuditSystemCallTypeGETPEERNAME            AuditSystemCallType = "getpeername"
	AuditSystemCallTypeSOCKETPAIR             AuditSystemCallType = "socketpair"
	AuditSystemCallTypeSETSOCKOPT             AuditSystemCallType = "setsockopt"
	AuditSystemCallTypeGETSOCKOPT             AuditSystemCallType = "getsockopt"
	AuditSystemCallTypeCLONE                  AuditSystemCallType = "clone"
	AuditSystemCallTypeFORK                   AuditSystemCallType = "fork"
	AuditSystemCallTypeVFORK                  AuditSystemCallType = "vfork"
	AuditSystemCallTypeEXECVE                 AuditSystemCallType = "execve"
	AuditSystemCallTypeEXIT                   AuditSystemCallType = "exit"
	AuditSystemCallTypeWAIT4                  AuditSystemCallType = "wait4"
	AuditSystemCallTypeKILL                   AuditSystemCallType = "kill"
	AuditSystemCallTypeUNAME                  AuditSystemCallType = "uname"
	AuditSystemCallTypeSEMGET                 AuditSystemCallType = "semget"
	AuditSystemCallTypeSEMOP                  AuditSystemCallType = "semop"
	AuditSystemCallTypeSEMCTL                 AuditSystemCallType = "semctl"
	AuditSystemCallTypeSHMDT                  AuditSystemCallType = "shmdt"
	AuditSystemCallTypeMSGGET                 AuditSystemCallType = "msgget"
	AuditSystemCallTypeMSGSND                 AuditSystemCallType = "msgsnd"
	AuditSystemCallTypeMSGRCV                 AuditSystemCallType = "msgrcv"
	AuditSystemCallTypeMSGCTL                 AuditSystemCallType = "msgctl"
	AuditSystemCallTypeFCNTL                  AuditSystemCallType = "fcntl"
	AuditSystemCallTypeFLOCK                  AuditSystemCallType = "flock"
	AuditSystemCallTypeFSYNC                  AuditSystemCallType = "fsync"
	AuditSystemCallTypeFDATASYNC              AuditSystemCallType = "fdatasync"
	AuditSystemCallTypeTRUNCATE               AuditSystemCallType = "truncate"
	AuditSystemCallTypeFTRUNCATE              AuditSystemCallType = "ftruncate"
	AuditSystemCallTypeGETDENTS               AuditSystemCallType = "getdents"
	AuditSystemCallTypeGETCWD                 AuditSystemCallType = "getcwd"
	AuditSystemCallTypeCHDIR                  AuditSystemCallType = "chdir"
	AuditSystemCallTypeFCHDIR                 AuditSystemCallType = "fchdir"
	AuditSystemCallTypeRENAME                 AuditSystemCallType = "rename"
	AuditSystemCallTypeMKDIR                  AuditSystemCallType = "mkdir"
	AuditSystemCallTypeRMDIR                  AuditSystemCallType = "rmdir"
	AuditSystemCallTypeCREAT                  AuditSystemCallType = "creat"
	AuditSystemCallTypeLINK                   AuditSystemCallType = "link"
	AuditSystemCallTypeUNLINK                 AuditSystemCallType = "unlink"
	AuditSystemCallTypeSYMLINK                AuditSystemCallType = "symlink"
	AuditSystemCallTypeREADLINK               AuditSystemCallType = "readlink"
	AuditSystemCallTypeCHMOD                  AuditSystemCallType = "chmod"
	AuditSystemCallTypeFCHMOD                 AuditSystemCallType = "fchmod"
	AuditSystemCallTypeCHOWN                  AuditSystemCallType = "chown"
	AuditSystemCallTypeFCHOWN                 AuditSystemCallType = "fchown"
	AuditSystemCallTypeLCHOWN                 AuditSystemCallType = "lchown"
	AuditSystemCallTypeUMASK                  AuditSystemCallType = "umask"
	AuditSystemCallTypeGETTIMEOFDAY           AuditSystemCallType = "gettimeofday"
	AuditSystemCallTypeGETRLIMIT              AuditSystemCallType = "getrlimit"
	AuditSystemCallTypeGETRUSAGE              AuditSystemCallType = "getrusage"
	AuditSystemCallTypeSYSINFO                AuditSystemCallType = "sysinfo"
	AuditSystemCallTypeTIMES                  AuditSystemCallType = "times"
	AuditSystemCallTypePTRACE                 AuditSystemCallType = "ptrace"
	AuditSystemCallTypeGETUID                 AuditSystemCallType = "getuid"
	AuditSystemCallTypeSYSLOG                 AuditSystemCallType = "syslog"
	AuditSystemCallTypeGETGID                 AuditSystemCallType = "getgid"
	AuditSystemCallTypeSETUID                 AuditSystemCallType = "setuid"
	AuditSystemCallTypeSETGID                 AuditSystemCallType = "setgid"
	AuditSystemCallTypeGETEUID                AuditSystemCallType = "geteuid"
	AuditSystemCallTypeGETEGID                AuditSystemCallType = "getegid"
	AuditSystemCallTypeSETPGID                AuditSystemCallType = "setpgid"
	AuditSystemCallTypeGETPPID                AuditSystemCallType = "getppid"
	AuditSystemCallTypeGETPGRP                AuditSystemCallType = "getpgrp"
	AuditSystemCallTypeSETSID                 AuditSystemCallType = "setsid"
	AuditSystemCallTypeSETREUID               AuditSystemCallType = "setreuid"
	AuditSystemCallTypeSETREGID               AuditSystemCallType = "setregid"
	AuditSystemCallTypeGETGROUPS              AuditSystemCallType = "getgroups"
	AuditSystemCallTypeSETGROUPS              AuditSystemCallType = "setgroups"
	AuditSystemCallTypeSETRESUID              AuditSystemCallType = "setresuid"
	AuditSystemCallTypeGETRESUID              AuditSystemCallType = "getresuid"
	AuditSystemCallTypeSETRESGID              AuditSystemCallType = "setresgid"
	AuditSystemCallTypeGETRESGID              AuditSystemCallType = "getresgid"
	AuditSystemCallTypeGETPGID                AuditSystemCallType = "getpgid"
	AuditSystemCallTypeSETFSUID               AuditSystemCallType = "setfsuid"
	AuditSystemCallTypeSETFSGID               AuditSystemCallType = "setfsgid"
	AuditSystemCallTypeGETSID                 AuditSystemCallType = "getsid"
	AuditSystemCallTypeCAPGET                 AuditSystemCallType = "capget"
	AuditSystemCallTypeCAPSET                 AuditSystemCallType = "capset"
	AuditSystemCallTypeRT_SIGPENDING          AuditSystemCallType = "rt_sigpending"
	AuditSystemCallTypeRT_SIGTIMEDWAIT        AuditSystemCallType = "rt_sigtimedwait"
	AuditSystemCallTypeRT_SIGQUEUEINFO        AuditSystemCallType = "rt_sigqueueinfo"
	AuditSystemCallTypeRT_SIGSUSPEND          AuditSystemCallType = "rt_sigsuspend"
	AuditSystemCallTypeSIGALTSTACK            AuditSystemCallType = "sigaltstack"
	AuditSystemCallTypeUTIME                  AuditSystemCallType = "utime"
	AuditSystemCallTypeMKNOD                  AuditSystemCallType = "mknod"
	AuditSystemCallTypeUSELIB                 AuditSystemCallType = "uselib"
	AuditSystemCallTypePERSONALITY            AuditSystemCallType = "personality"
	AuditSystemCallTypeUSTAT                  AuditSystemCallType = "ustat"
	AuditSystemCallTypeSTATFS                 AuditSystemCallType = "statfs"
	AuditSystemCallTypeFSTATFS                AuditSystemCallType = "fstatfs"
	AuditSystemCallTypeSYSFS                  AuditSystemCallType = "sysfs"
	AuditSystemCallTypeGETPRIORITY            AuditSystemCallType = "getpriority"
	AuditSystemCallTypeSETPRIORITY            AuditSystemCallType = "setpriority"
	AuditSystemCallTypeSCHED_SETPARAM         AuditSystemCallType = "sched_setparam"
	AuditSystemCallTypeSCHED_GETPARAM         AuditSystemCallType = "sched_getparam"
	AuditSystemCallTypeSCHED_SETSCHEDULER     AuditSystemCallType = "sched_setscheduler"
	AuditSystemCallTypeSCHED_GETSCHEDULER     AuditSystemCallType = "sched_getscheduler"
	AuditSystemCallTypeSCHED_GET_PRIORITY_MAX AuditSystemCallType = "sched_get_priority_max"
	AuditSystemCallTypeSCHED_GET_PRIORITY_MIN AuditSystemCallType = "sched_get_priority_min"
	AuditSystemCallTypeSCHED_RR_GET_INTERVAL  AuditSystemCallType = "sched_rr_get_interval"
	AuditSystemCallTypeMLOCK                  AuditSystemCallType = "mlock"
	AuditSystemCallTypeMUNLOCK                AuditSystemCallType = "munlock"
	AuditSystemCallTypeMLOCKALL               AuditSystemCallType = "mlockall"
	AuditSystemCallTypeMUNLOCKALL             AuditSystemCallType = "munlockall"
	AuditSystemCallTypeVHANGUP                AuditSystemCallType = "vhangup"
	AuditSystemCallTypeMODIFY_LDT             AuditSystemCallType = "modify_ldt"
	AuditSystemCallTypePIVOT_ROOT             AuditSystemCallType = "pivot_root"
	AuditSystemCallType_SYSCTL                AuditSystemCallType = "_sysctl"
	AuditSystemCallTypePRCTL                  AuditSystemCallType = "prctl"
	AuditSystemCallTypeARCH_PRCTL             AuditSystemCallType = "arch_prctl"
	AuditSystemCallTypeADJTIMEX               AuditSystemCallType = "adjtimex"
	AuditSystemCallTypeSETRLIMIT              AuditSystemCallType = "setrlimit"
	AuditSystemCallTypeCHROOT                 AuditSystemCallType = "chroot"
	AuditSystemCallTypeSYNC                   AuditSystemCallType = "sync"
	AuditSystemCallTypeACCT                   AuditSystemCallType = "acct"
	AuditSystemCallTypeSETTIMEOFDAY           AuditSystemCallType = "settimeofday"
	AuditSystemCallTypeMOUNT                  AuditSystemCallType = "mount"
	AuditSystemCallTypeUMOUNT2                AuditSystemCallType = "umount2"
	AuditSystemCallTypeSWAPON                 AuditSystemCallType = "swapon"
	AuditSystemCallTypeSWAPOFF                AuditSystemCallType = "swapoff"
	AuditSystemCallTypeREBOOT                 AuditSystemCallType = "reboot"
	AuditSystemCallTypeSETHOSTNAME            AuditSystemCallType = "sethostname"
	AuditSystemCallTypeSETDOMAINNAME          AuditSystemCallType = "setdomainname"
	AuditSystemCallTypeIOPL                   AuditSystemCallType = "iopl"
	AuditSystemCallTypeIOPERM                 AuditSystemCallType = "ioperm"
	AuditSystemCallTypeCREATE_MODULE          AuditSystemCallType = "create_module"
	AuditSystemCallTypeINIT_MODULE            AuditSystemCallType = "init_module"
	AuditSystemCallTypeDELETE_MODULE          AuditSystemCallType = "delete_module"
	AuditSystemCallTypeGET_KERNEL_SYMS        AuditSystemCallType = "get_kernel_syms"
	AuditSystemCallTypeQUERY_MODULE           AuditSystemCallType = "query_module"
	AuditSystemCallTypeQUOTACTL               AuditSystemCallType = "quotactl"
	AuditSystemCallTypeNFSSERVCTL             AuditSystemCallType = "nfsservctl"
	AuditSystemCallTypeGETPMSG                AuditSystemCallType = "getpmsg"
	AuditSystemCallTypePUTPMSG                AuditSystemCallType = "putpmsg"
	AuditSystemCallTypeAFS_SYSCALL            AuditSystemCallType = "afs_syscall"
	AuditSystemCallTypeTUXCALL                AuditSystemCallType = "tuxcall"
	AuditSystemCallTypeSECURITY               AuditSystemCallType = "security"
	AuditSystemCallTypeGETTID                 AuditSystemCallType = "gettid"
	AuditSystemCallTypeREADAHEAD              AuditSystemCallType = "readahead"
	AuditSystemCallTypeSETXATTR               AuditSystemCallType = "setxattr"
	AuditSystemCallTypeLSETXATTR              AuditSystemCallType = "lsetxattr"
	AuditSystemCallTypeFSETXATTR              AuditSystemCallType = "fsetxattr"
	AuditSystemCallTypeGETXATTR               AuditSystemCallType = "getxattr"
	AuditSystemCallTypeLGETXATTR              AuditSystemCallType = "lgetxattr"
	AuditSystemCallTypeFGETXATTR              AuditSystemCallType = "fgetxattr"
	AuditSystemCallTypeLISTXATTR              AuditSystemCallType = "listxattr"
	AuditSystemCallTypeLLISTXATTR             AuditSystemCallType = "llistxattr"
	AuditSystemCallTypeFLISTXATTR             AuditSystemCallType = "flistxattr"
	AuditSystemCallTypeREMOVEXATTR            AuditSystemCallType = "removexattr"
	AuditSystemCallTypeLREMOVEXATTR           AuditSystemCallType = "lremovexattr"
	AuditSystemCallTypeFREMOVEXATTR           AuditSystemCallType = "fremovexattr"
	AuditSystemCallTypeTKILL                  AuditSystemCallType = "tkill"
	AuditSystemCallTypeTIME                   AuditSystemCallType = "time"
	AuditSystemCallTypeFUTEX                  AuditSystemCallType = "futex"
	AuditSystemCallTypeSCHED_SETAFFINITY      AuditSystemCallType = "sched_setaffinity"
	AuditSystemCallTypeSCHED_GETAFFINITY      AuditSystemCallType = "sched_getaffinity"
	AuditSystemCallTypeSET_THREAD_AREA        AuditSystemCallType = "set_thread_area"
	AuditSystemCallTypeIO_SETUP               AuditSystemCallType = "io_setup"
	AuditSystemCallTypeIO_DESTROY             AuditSystemCallType = "io_destroy"
	AuditSystemCallTypeIO_GETEVENTS           AuditSystemCallType = "io_getevents"
	AuditSystemCallTypeIO_SUBMIT              AuditSystemCallType = "io_submit"
	AuditSystemCallTypeIO_CANCEL              AuditSystemCallType = "io_cancel"
	AuditSystemCallTypeGET_THREAD_AREA        AuditSystemCallType = "get_thread_area"
	AuditSystemCallTypeLOOKUP_DCOOKIE         AuditSystemCallType = "lookup_dcookie"
	AuditSystemCallTypeEPOLL_CREATE           AuditSystemCallType = "epoll_create"
	AuditSystemCallTypeEPOLL_CTL_OLD          AuditSystemCallType = "epoll_ctl_old"
	AuditSystemCallTypeEPOLL_WAIT_OLD         AuditSystemCallType = "epoll_wait_old"
	AuditSystemCallTypeREMAP_FILE_PAGES       AuditSystemCallType = "remap_file_pages"
	AuditSystemCallTypeGETDENTS64             AuditSystemCallType = "getdents64"
	AuditSystemCallTypeSET_TID_ADDRESS        AuditSystemCallType = "set_tid_address"
	AuditSystemCallTypeRESTART_SYSCALL        AuditSystemCallType = "restart_syscall"
	AuditSystemCallTypeSEMTIMEDOP             AuditSystemCallType = "semtimedop"
	AuditSystemCallTypeFADVISE64              AuditSystemCallType = "fadvise64"
	AuditSystemCallTypeTIMER_CREATE           AuditSystemCallType = "timer_create"
	AuditSystemCallTypeTIMER_SETTIME          AuditSystemCallType = "timer_settime"
	AuditSystemCallTypeTIMER_GETTIME          AuditSystemCallType = "timer_gettime"
	AuditSystemCallTypeTIMER_GETOVERRUN       AuditSystemCallType = "timer_getoverrun"
	AuditSystemCallTypeTIMER_DELETE           AuditSystemCallType = "timer_delete"
	AuditSystemCallTypeCLOCK_SETTIME          AuditSystemCallType = "clock_settime"
	AuditSystemCallTypeCLOCK_GETTIME          AuditSystemCallType = "clock_gettime"
	AuditSystemCallTypeCLOCK_GETRES           AuditSystemCallType = "clock_getres"
	AuditSystemCallTypeCLOCK_NANOSLEEP        AuditSystemCallType = "clock_nanosleep"
	AuditSystemCallTypeEXIT_GROUP             AuditSystemCallType = "exit_group"
	AuditSystemCallTypeEPOLL_WAIT             AuditSystemCallType = "epoll_wait"
	AuditSystemCallTypeEPOLL_CTL              AuditSystemCallType = "epoll_ctl"
	AuditSystemCallTypeTGKILL                 AuditSystemCallType = "tgkill"
	AuditSystemCallTypeUTIMES                 AuditSystemCallType = "utimes"
	AuditSystemCallTypeVSERVER                AuditSystemCallType = "vserver"
	AuditSystemCallTypeMBIND                  AuditSystemCallType = "mbind"
	AuditSystemCallTypeSET_MEMPOLICY          AuditSystemCallType = "set_mempolicy"
	AuditSystemCallTypeGET_MEMPOLICY          AuditSystemCallType = "get_mempolicy"
	AuditSystemCallTypeMQ_OPEN                AuditSystemCallType = "mq_open"
	AuditSystemCallTypeMQ_UNLINK              AuditSystemCallType = "mq_unlink"
	AuditSystemCallTypeMQ_TIMEDSEND           AuditSystemCallType = "mq_timedsend"
	AuditSystemCallTypeMQ_TIMEDRECEIVE        AuditSystemCallType = "mq_timedreceive"
	AuditSystemCallTypeMQ_NOTIFY              AuditSystemCallType = "mq_notify"
	AuditSystemCallTypeMQ_GETSETATTR          AuditSystemCallType = "mq_getsetattr"
	AuditSystemCallTypeKEXEC_LOAD             AuditSystemCallType = "kexec_load"
	AuditSystemCallTypeWAITID                 AuditSystemCallType = "waitid"
	AuditSystemCallTypeADD_KEY                AuditSystemCallType = "add_key"
	AuditSystemCallTypeREQUEST_KEY            AuditSystemCallType = "request_key"
	AuditSystemCallTypeKEYCTL                 AuditSystemCallType = "keyctl"
	AuditSystemCallTypeIOPRIO_SET             AuditSystemCallType = "ioprio_set"
	AuditSystemCallTypeIOPRIO_GET             AuditSystemCallType = "ioprio_get"
	AuditSystemCallTypeINOTIFY_INIT           AuditSystemCallType = "inotify_init"
	AuditSystemCallTypeINOTIFY_ADD_WATCH      AuditSystemCallType = "inotify_add_watch"
	AuditSystemCallTypeINOTIFY_RM_WATCH       AuditSystemCallType = "inotify_rm_watch"
	AuditSystemCallTypeMIGRATE_PAGES          AuditSystemCallType = "migrate_pages"
	AuditSystemCallTypeOPENAT                 AuditSystemCallType = "openat"
	AuditSystemCallTypeMKDIRAT                AuditSystemCallType = "mkdirat"
	AuditSystemCallTypeMKNODAT                AuditSystemCallType = "mknodat"
	AuditSystemCallTypeFCHOWNAT               AuditSystemCallType = "fchownat"
	AuditSystemCallTypeFUTIMESAT              AuditSystemCallType = "futimesat"
	AuditSystemCallTypeNEWFSTATAT             AuditSystemCallType = "newfstatat"
	AuditSystemCallTypeUNLINKAT               AuditSystemCallType = "unlinkat"
	AuditSystemCallTypeRENAMEAT               AuditSystemCallType = "renameat"
	AuditSystemCallTypeLINKAT                 AuditSystemCallType = "linkat"
	AuditSystemCallTypeSYMLINKAT              AuditSystemCallType = "symlinkat"
	AuditSystemCallTypeREADLINKAT             AuditSystemCallType = "readlinkat"
	AuditSystemCallTypeFCHMODAT               AuditSystemCallType = "fchmodat"
	AuditSystemCallTypeFACCESSAT              AuditSystemCallType = "faccessat"
	AuditSystemCallTypePSELECT6               AuditSystemCallType = "pselect6"
	AuditSystemCallTypePPOLL                  AuditSystemCallType = "ppoll"
	AuditSystemCallTypeUNSHARE                AuditSystemCallType = "unshare"
	AuditSystemCallTypeSET_ROBUST_LIST        AuditSystemCallType = "set_robust_list"
	AuditSystemCallTypeGET_ROBUST_LIST        AuditSystemCallType = "get_robust_list"
	AuditSystemCallTypeSPLICE                 AuditSystemCallType = "splice"
	AuditSystemCallTypeTEE                    AuditSystemCallType = "tee"
	AuditSystemCallTypeSYNC_FILE_RANGE        AuditSystemCallType = "sync_file_range"
	AuditSystemCallTypeVMSPLICE               AuditSystemCallType = "vmsplice"
	AuditSystemCallTypeMOVE_PAGES             AuditSystemCallType = "move_pages"
	AuditSystemCallTypeUTIMENSAT              AuditSystemCallType = "utimensat"
	AuditSystemCallTypeEPOLL_PWAIT            AuditSystemCallType = "epoll_pwait"
	AuditSystemCallTypeSIGNALFD               AuditSystemCallType = "signalfd"
	AuditSystemCallTypeTIMERFD_CREATE         AuditSystemCallType = "timerfd_create"
	AuditSystemCallTypeEVENTFD                AuditSystemCallType = "eventfd"
	AuditSystemCallTypeFALLOCATE              AuditSystemCallType = "fallocate"
	AuditSystemCallTypeTIMERFD_SETTIME        AuditSystemCallType = "timerfd_settime"
	AuditSystemCallTypeTIMERFD_GETTIME        AuditSystemCallType = "timerfd_gettime"
	AuditSystemCallTypeACCEPT4                AuditSystemCallType = "accept4"
	AuditSystemCallTypeSIGNALFD4              AuditSystemCallType = "signalfd4"
	AuditSystemCallTypeEVENTFD2               AuditSystemCallType = "eventfd2"
	AuditSystemCallTypeEPOLL_CREATE1          AuditSystemCallType = "epoll_create1"
	AuditSystemCallTypeDUP3                   AuditSystemCallType = "dup3"
	AuditSystemCallTypePIPE2                  AuditSystemCallType = "pipe2"
	AuditSystemCallTypeINOTIFY_INIT1          AuditSystemCallType = "inotify_init1"
	AuditSystemCallTypePREADV                 AuditSystemCallType = "preadv"
	AuditSystemCallTypePWRITEV                AuditSystemCallType = "pwritev"
	AuditSystemCallTypeRT_TGSIGQUEUEINFO      AuditSystemCallType = "rt_tgsigqueueinfo"
	AuditSystemCallTypePERF_EVENT_OPEN        AuditSystemCallType = "perf_event_open"
	AuditSystemCallTypeRECVMMSG               AuditSystemCallType = "recvmmsg"
	AuditSystemCallTypeFANOTIFY_INIT          AuditSystemCallType = "fanotify_init"
	AuditSystemCallTypeFANOTIFY_MARK          AuditSystemCallType = "fanotify_mark"
	AuditSystemCallTypePRLIMIT64              AuditSystemCallType = "prlimit64"
	AuditSystemCallTypeNAME_TO_HANDLE_AT      AuditSystemCallType = "name_to_handle_at"
	AuditSystemCallTypeOPEN_BY_HANDLE_AT      AuditSystemCallType = "open_by_handle_at"
	AuditSystemCallTypeCLOCK_ADJTIME          AuditSystemCallType = "clock_adjtime"
	AuditSystemCallTypeSYNCFS                 AuditSystemCallType = "syncfs"
	AuditSystemCallTypeSENDMMSG               AuditSystemCallType = "sendmmsg"
	AuditSystemCallTypeSETNS                  AuditSystemCallType = "setns"
	AuditSystemCallTypeGETCPU                 AuditSystemCallType = "getcpu"
	AuditSystemCallTypePROCESS_VM_READV       AuditSystemCallType = "process_vm_readv"
	AuditSystemCallTypePROCESS_VM_WRITEV      AuditSystemCallType = "process_vm_writev"
	AuditSystemCallTypeKCMP                   AuditSystemCallType = "kcmp"
	AuditSystemCallTypeFINIT_MODULE           AuditSystemCallType = "finit_module"
	AuditSystemCallTypeSTIME                  AuditSystemCallType = "stime"
)

var auditSystemCallTypeReverse = map[string]interface{}{
	"read":                   AuditSystemCallTypeREAD,
	"write":                  AuditSystemCallTypeWRITE,
	"open":                   AuditSystemCallTypeOPEN,
	"close":                  AuditSystemCallTypeCLOSE,
	"stat":                   AuditSystemCallTypeSTAT,
	"fstat":                  AuditSystemCallTypeFSTAT,
	"lstat":                  AuditSystemCallTypeLSTAT,
	"poll":                   AuditSystemCallTypePOLL,
	"lseek":                  AuditSystemCallTypeLSEEK,
	"mmap":                   AuditSystemCallTypeMMAP,
	"mprotect":               AuditSystemCallTypeMPROTECT,
	"munmap":                 AuditSystemCallTypeMUNMAP,
	"brk":                    AuditSystemCallTypeBRK,
	"rt_sigaction":           AuditSystemCallTypeRT_SIGACTION,
	"rt_sigprocmask":         AuditSystemCallTypeRT_SIGPROCMASK,
	"rt_sigreturn":           AuditSystemCallTypeRT_SIGRETURN,
	"ioctl":                  AuditSystemCallTypeIOCTL,
	"pread64":                AuditSystemCallTypePREAD64,
	"pwrite64":               AuditSystemCallTypePWRITE64,
	"readv":                  AuditSystemCallTypeREADV,
	"writev":                 AuditSystemCallTypeWRITEV,
	"access":                 AuditSystemCallTypeACCESS,
	"pipe":                   AuditSystemCallTypePIPE,
	"select":                 AuditSystemCallTypeSELECT,
	"sched_yield":            AuditSystemCallTypeSCHED_YIELD,
	"mremap":                 AuditSystemCallTypeMREMAP,
	"msync":                  AuditSystemCallTypeMSYNC,
	"mincore":                AuditSystemCallTypeMINCORE,
	"madvise":                AuditSystemCallTypeMADVISE,
	"shmget":                 AuditSystemCallTypeSHMGET,
	"shmat":                  AuditSystemCallTypeSHMAT,
	"shmctl":                 AuditSystemCallTypeSHMCTL,
	"dup":                    AuditSystemCallTypeDUP,
	"dup2":                   AuditSystemCallTypeDUP2,
	"pause":                  AuditSystemCallTypePAUSE,
	"nanosleep":              AuditSystemCallTypeNANOSLEEP,
	"getitimer":              AuditSystemCallTypeGETITIMER,
	"alarm":                  AuditSystemCallTypeALARM,
	"setitimer":              AuditSystemCallTypeSETITIMER,
	"getpid":                 AuditSystemCallTypeGETPID,
	"sendfile":               AuditSystemCallTypeSENDFILE,
	"socket":                 AuditSystemCallTypeSOCKET,
	"connect":                AuditSystemCallTypeCONNECT,
	"accept":                 AuditSystemCallTypeACCEPT,
	"sendto":                 AuditSystemCallTypeSENDTO,
	"recvfrom":               AuditSystemCallTypeRECVFROM,
	"sendmsg":                AuditSystemCallTypeSENDMSG,
	"recvmsg":                AuditSystemCallTypeRECVMSG,
	"shutdown":               AuditSystemCallTypeSHUTDOWN,
	"bind":                   AuditSystemCallTypeBIND,
	"listen":                 AuditSystemCallTypeLISTEN,
	"getsockname":            AuditSystemCallTypeGETSOCKNAME,
	"getpeername":            AuditSystemCallTypeGETPEERNAME,
	"socketpair":             AuditSystemCallTypeSOCKETPAIR,
	"setsockopt":             AuditSystemCallTypeSETSOCKOPT,
	"getsockopt":             AuditSystemCallTypeGETSOCKOPT,
	"clone":                  AuditSystemCallTypeCLONE,
	"fork":                   AuditSystemCallTypeFORK,
	"vfork":                  AuditSystemCallTypeVFORK,
	"execve":                 AuditSystemCallTypeEXECVE,
	"exit":                   AuditSystemCallTypeEXIT,
	"wait4":                  AuditSystemCallTypeWAIT4,
	"kill":                   AuditSystemCallTypeKILL,
	"uname":                  AuditSystemCallTypeUNAME,
	"semget":                 AuditSystemCallTypeSEMGET,
	"semop":                  AuditSystemCallTypeSEMOP,
	"semctl":                 AuditSystemCallTypeSEMCTL,
	"shmdt":                  AuditSystemCallTypeSHMDT,
	"msgget":                 AuditSystemCallTypeMSGGET,
	"msgsnd":                 AuditSystemCallTypeMSGSND,
	"msgrcv":                 AuditSystemCallTypeMSGRCV,
	"msgctl":                 AuditSystemCallTypeMSGCTL,
	"fcntl":                  AuditSystemCallTypeFCNTL,
	"flock":                  AuditSystemCallTypeFLOCK,
	"fsync":                  AuditSystemCallTypeFSYNC,
	"fdatasync":              AuditSystemCallTypeFDATASYNC,
	"truncate":               AuditSystemCallTypeTRUNCATE,
	"ftruncate":              AuditSystemCallTypeFTRUNCATE,
	"getdents":               AuditSystemCallTypeGETDENTS,
	"getcwd":                 AuditSystemCallTypeGETCWD,
	"chdir":                  AuditSystemCallTypeCHDIR,
	"fchdir":                 AuditSystemCallTypeFCHDIR,
	"rename":                 AuditSystemCallTypeRENAME,
	"mkdir":                  AuditSystemCallTypeMKDIR,
	"rmdir":                  AuditSystemCallTypeRMDIR,
	"creat":                  AuditSystemCallTypeCREAT,
	"link":                   AuditSystemCallTypeLINK,
	"unlink":                 AuditSystemCallTypeUNLINK,
	"symlink":                AuditSystemCallTypeSYMLINK,
	"readlink":               AuditSystemCallTypeREADLINK,
	"chmod":                  AuditSystemCallTypeCHMOD,
	"fchmod":                 AuditSystemCallTypeFCHMOD,
	"chown":                  AuditSystemCallTypeCHOWN,
	"fchown":                 AuditSystemCallTypeFCHOWN,
	"lchown":                 AuditSystemCallTypeLCHOWN,
	"umask":                  AuditSystemCallTypeUMASK,
	"gettimeofday":           AuditSystemCallTypeGETTIMEOFDAY,
	"getrlimit":              AuditSystemCallTypeGETRLIMIT,
	"getrusage":              AuditSystemCallTypeGETRUSAGE,
	"sysinfo":                AuditSystemCallTypeSYSINFO,
	"times":                  AuditSystemCallTypeTIMES,
	"ptrace":                 AuditSystemCallTypePTRACE,
	"getuid":                 AuditSystemCallTypeGETUID,
	"syslog":                 AuditSystemCallTypeSYSLOG,
	"getgid":                 AuditSystemCallTypeGETGID,
	"setuid":                 AuditSystemCallTypeSETUID,
	"setgid":                 AuditSystemCallTypeSETGID,
	"geteuid":                AuditSystemCallTypeGETEUID,
	"getegid":                AuditSystemCallTypeGETEGID,
	"setpgid":                AuditSystemCallTypeSETPGID,
	"getppid":                AuditSystemCallTypeGETPPID,
	"getpgrp":                AuditSystemCallTypeGETPGRP,
	"setsid":                 AuditSystemCallTypeSETSID,
	"setreuid":               AuditSystemCallTypeSETREUID,
	"setregid":               AuditSystemCallTypeSETREGID,
	"getgroups":              AuditSystemCallTypeGETGROUPS,
	"setgroups":              AuditSystemCallTypeSETGROUPS,
	"setresuid":              AuditSystemCallTypeSETRESUID,
	"getresuid":              AuditSystemCallTypeGETRESUID,
	"setresgid":              AuditSystemCallTypeSETRESGID,
	"getresgid":              AuditSystemCallTypeGETRESGID,
	"getpgid":                AuditSystemCallTypeGETPGID,
	"setfsuid":               AuditSystemCallTypeSETFSUID,
	"setfsgid":               AuditSystemCallTypeSETFSGID,
	"getsid":                 AuditSystemCallTypeGETSID,
	"capget":                 AuditSystemCallTypeCAPGET,
	"capset":                 AuditSystemCallTypeCAPSET,
	"rt_sigpending":          AuditSystemCallTypeRT_SIGPENDING,
	"rt_sigtimedwait":        AuditSystemCallTypeRT_SIGTIMEDWAIT,
	"rt_sigqueueinfo":        AuditSystemCallTypeRT_SIGQUEUEINFO,
	"rt_sigsuspend":          AuditSystemCallTypeRT_SIGSUSPEND,
	"sigaltstack":            AuditSystemCallTypeSIGALTSTACK,
	"utime":                  AuditSystemCallTypeUTIME,
	"mknod":                  AuditSystemCallTypeMKNOD,
	"uselib":                 AuditSystemCallTypeUSELIB,
	"personality":            AuditSystemCallTypePERSONALITY,
	"ustat":                  AuditSystemCallTypeUSTAT,
	"statfs":                 AuditSystemCallTypeSTATFS,
	"fstatfs":                AuditSystemCallTypeFSTATFS,
	"sysfs":                  AuditSystemCallTypeSYSFS,
	"getpriority":            AuditSystemCallTypeGETPRIORITY,
	"setpriority":            AuditSystemCallTypeSETPRIORITY,
	"sched_setparam":         AuditSystemCallTypeSCHED_SETPARAM,
	"sched_getparam":         AuditSystemCallTypeSCHED_GETPARAM,
	"sched_setscheduler":     AuditSystemCallTypeSCHED_SETSCHEDULER,
	"sched_getscheduler":     AuditSystemCallTypeSCHED_GETSCHEDULER,
	"sched_get_priority_max": AuditSystemCallTypeSCHED_GET_PRIORITY_MAX,
	"sched_get_priority_min": AuditSystemCallTypeSCHED_GET_PRIORITY_MIN,
	"sched_rr_get_interval":  AuditSystemCallTypeSCHED_RR_GET_INTERVAL,
	"mlock":                  AuditSystemCallTypeMLOCK,
	"munlock":                AuditSystemCallTypeMUNLOCK,
	"mlockall":               AuditSystemCallTypeMLOCKALL,
	"munlockall":             AuditSystemCallTypeMUNLOCKALL,
	"vhangup":                AuditSystemCallTypeVHANGUP,
	"modify_ldt":             AuditSystemCallTypeMODIFY_LDT,
	"pivot_root":             AuditSystemCallTypePIVOT_ROOT,
	"_sysctl":                AuditSystemCallType_SYSCTL,
	"prctl":                  AuditSystemCallTypePRCTL,
	"arch_prctl":             AuditSystemCallTypeARCH_PRCTL,
	"adjtimex":               AuditSystemCallTypeADJTIMEX,
	"setrlimit":              AuditSystemCallTypeSETRLIMIT,
	"chroot":                 AuditSystemCallTypeCHROOT,
	"sync":                   AuditSystemCallTypeSYNC,
	"acct":                   AuditSystemCallTypeACCT,
	"settimeofday":           AuditSystemCallTypeSETTIMEOFDAY,
	"mount":                  AuditSystemCallTypeMOUNT,
	"umount2":                AuditSystemCallTypeUMOUNT2,
	"swapon":                 AuditSystemCallTypeSWAPON,
	"swapoff":                AuditSystemCallTypeSWAPOFF,
	"reboot":                 AuditSystemCallTypeREBOOT,
	"sethostname":            AuditSystemCallTypeSETHOSTNAME,
	"setdomainname":          AuditSystemCallTypeSETDOMAINNAME,
	"iopl":                   AuditSystemCallTypeIOPL,
	"ioperm":                 AuditSystemCallTypeIOPERM,
	"create_module":          AuditSystemCallTypeCREATE_MODULE,
	"init_module":            AuditSystemCallTypeINIT_MODULE,
	"delete_module":          AuditSystemCallTypeDELETE_MODULE,
	"get_kernel_syms":        AuditSystemCallTypeGET_KERNEL_SYMS,
	"query_module":           AuditSystemCallTypeQUERY_MODULE,
	"quotactl":               AuditSystemCallTypeQUOTACTL,
	"nfsservctl":             AuditSystemCallTypeNFSSERVCTL,
	"getpmsg":                AuditSystemCallTypeGETPMSG,
	"putpmsg":                AuditSystemCallTypePUTPMSG,
	"afs_syscall":            AuditSystemCallTypeAFS_SYSCALL,
	"tuxcall":                AuditSystemCallTypeTUXCALL,
	"security":               AuditSystemCallTypeSECURITY,
	"gettid":                 AuditSystemCallTypeGETTID,
	"readahead":              AuditSystemCallTypeREADAHEAD,
	"setxattr":               AuditSystemCallTypeSETXATTR,
	"lsetxattr":              AuditSystemCallTypeLSETXATTR,
	"fsetxattr":              AuditSystemCallTypeFSETXATTR,
	"getxattr":               AuditSystemCallTypeGETXATTR,
	"lgetxattr":              AuditSystemCallTypeLGETXATTR,
	"fgetxattr":              AuditSystemCallTypeFGETXATTR,
	"listxattr":              AuditSystemCallTypeLISTXATTR,
	"llistxattr":             AuditSystemCallTypeLLISTXATTR,
	"flistxattr":             AuditSystemCallTypeFLISTXATTR,
	"removexattr":            AuditSystemCallTypeREMOVEXATTR,
	"lremovexattr":           AuditSystemCallTypeLREMOVEXATTR,
	"fremovexattr":           AuditSystemCallTypeFREMOVEXATTR,
	"tkill":                  AuditSystemCallTypeTKILL,
	"time":                   AuditSystemCallTypeTIME,
	"futex":                  AuditSystemCallTypeFUTEX,
	"sched_setaffinity":      AuditSystemCallTypeSCHED_SETAFFINITY,
	"sched_getaffinity":      AuditSystemCallTypeSCHED_GETAFFINITY,
	"set_thread_area":        AuditSystemCallTypeSET_THREAD_AREA,
	"io_setup":               AuditSystemCallTypeIO_SETUP,
	"io_destroy":             AuditSystemCallTypeIO_DESTROY,
	"io_getevents":           AuditSystemCallTypeIO_GETEVENTS,
	"io_submit":              AuditSystemCallTypeIO_SUBMIT,
	"io_cancel":              AuditSystemCallTypeIO_CANCEL,
	"get_thread_area":        AuditSystemCallTypeGET_THREAD_AREA,
	"lookup_dcookie":         AuditSystemCallTypeLOOKUP_DCOOKIE,
	"epoll_create":           AuditSystemCallTypeEPOLL_CREATE,
	"epoll_ctl_old":          AuditSystemCallTypeEPOLL_CTL_OLD,
	"epoll_wait_old":         AuditSystemCallTypeEPOLL_WAIT_OLD,
	"remap_file_pages":       AuditSystemCallTypeREMAP_FILE_PAGES,
	"getdents64":             AuditSystemCallTypeGETDENTS64,
	"set_tid_address":        AuditSystemCallTypeSET_TID_ADDRESS,
	"restart_syscall":        AuditSystemCallTypeRESTART_SYSCALL,
	"semtimedop":             AuditSystemCallTypeSEMTIMEDOP,
	"fadvise64":              AuditSystemCallTypeFADVISE64,
	"timer_create":           AuditSystemCallTypeTIMER_CREATE,
	"timer_settime":          AuditSystemCallTypeTIMER_SETTIME,
	"timer_gettime":          AuditSystemCallTypeTIMER_GETTIME,
	"timer_getoverrun":       AuditSystemCallTypeTIMER_GETOVERRUN,
	"timer_delete":           AuditSystemCallTypeTIMER_DELETE,
	"clock_settime":          AuditSystemCallTypeCLOCK_SETTIME,
	"clock_gettime":          AuditSystemCallTypeCLOCK_GETTIME,
	"clock_getres":           AuditSystemCallTypeCLOCK_GETRES,
	"clock_nanosleep":        AuditSystemCallTypeCLOCK_NANOSLEEP,
	"exit_group":             AuditSystemCallTypeEXIT_GROUP,
	"epoll_wait":             AuditSystemCallTypeEPOLL_WAIT,
	"epoll_ctl":              AuditSystemCallTypeEPOLL_CTL,
	"tgkill":                 AuditSystemCallTypeTGKILL,
	"utimes":                 AuditSystemCallTypeUTIMES,
	"vserver":                AuditSystemCallTypeVSERVER,
	"mbind":                  AuditSystemCallTypeMBIND,
	"set_mempolicy":          AuditSystemCallTypeSET_MEMPOLICY,
	"get_mempolicy":          AuditSystemCallTypeGET_MEMPOLICY,
	"mq_open":                AuditSystemCallTypeMQ_OPEN,
	"mq_unlink":              AuditSystemCallTypeMQ_UNLINK,
	"mq_timedsend":           AuditSystemCallTypeMQ_TIMEDSEND,
	"mq_timedreceive":        AuditSystemCallTypeMQ_TIMEDRECEIVE,
	"mq_notify":              AuditSystemCallTypeMQ_NOTIFY,
	"mq_getsetattr":          AuditSystemCallTypeMQ_GETSETATTR,
	"kexec_load":             AuditSystemCallTypeKEXEC_LOAD,
	"waitid":                 AuditSystemCallTypeWAITID,
	"add_key":                AuditSystemCallTypeADD_KEY,
	"request_key":            AuditSystemCallTypeREQUEST_KEY,
	"keyctl":                 AuditSystemCallTypeKEYCTL,
	"ioprio_set":             AuditSystemCallTypeIOPRIO_SET,
	"ioprio_get":             AuditSystemCallTypeIOPRIO_GET,
	"inotify_init":           AuditSystemCallTypeINOTIFY_INIT,
	"inotify_add_watch":      AuditSystemCallTypeINOTIFY_ADD_WATCH,
	"inotify_rm_watch":       AuditSystemCallTypeINOTIFY_RM_WATCH,
	"migrate_pages":          AuditSystemCallTypeMIGRATE_PAGES,
	"openat":                 AuditSystemCallTypeOPENAT,
	"mkdirat":                AuditSystemCallTypeMKDIRAT,
	"mknodat":                AuditSystemCallTypeMKNODAT,
	"fchownat":               AuditSystemCallTypeFCHOWNAT,
	"futimesat":              AuditSystemCallTypeFUTIMESAT,
	"newfstatat":             AuditSystemCallTypeNEWFSTATAT,
	"unlinkat":               AuditSystemCallTypeUNLINKAT,
	"renameat":               AuditSystemCallTypeRENAMEAT,
	"linkat":                 AuditSystemCallTypeLINKAT,
	"symlinkat":              AuditSystemCallTypeSYMLINKAT,
	"readlinkat":             AuditSystemCallTypeREADLINKAT,
	"fchmodat":               AuditSystemCallTypeFCHMODAT,
	"faccessat":              AuditSystemCallTypeFACCESSAT,
	"pselect6":               AuditSystemCallTypePSELECT6,
	"ppoll":                  AuditSystemCallTypePPOLL,
	"unshare":                AuditSystemCallTypeUNSHARE,
	"set_robust_list":        AuditSystemCallTypeSET_ROBUST_LIST,
	"get_robust_list":        AuditSystemCallTypeGET_ROBUST_LIST,
	"splice":                 AuditSystemCallTypeSPLICE,
	"tee":                    AuditSystemCallTypeTEE,
	"sync_file_range":        AuditSystemCallTypeSYNC_FILE_RANGE,
	"vmsplice":               AuditSystemCallTypeVMSPLICE,
	"move_pages":             AuditSystemCallTypeMOVE_PAGES,
	"utimensat":              AuditSystemCallTypeUTIMENSAT,
	"epoll_pwait":            AuditSystemCallTypeEPOLL_PWAIT,
	"signalfd":               AuditSystemCallTypeSIGNALFD,
	"timerfd_create":         AuditSystemCallTypeTIMERFD_CREATE,
	"eventfd":                AuditSystemCallTypeEVENTFD,
	"fallocate":              AuditSystemCallTypeFALLOCATE,
	"timerfd_settime":        AuditSystemCallTypeTIMERFD_SETTIME,
	"timerfd_gettime":        AuditSystemCallTypeTIMERFD_GETTIME,
	"accept4":                AuditSystemCallTypeACCEPT4,
	"signalfd4":              AuditSystemCallTypeSIGNALFD4,
	"eventfd2":               AuditSystemCallTypeEVENTFD2,
	"epoll_create1":          AuditSystemCallTypeEPOLL_CREATE1,
	"dup3":                   AuditSystemCallTypeDUP3,
	"pipe2":                  AuditSystemCallTypePIPE2,
	"inotify_init1":          AuditSystemCallTypeINOTIFY_INIT1,
	"preadv":                 AuditSystemCallTypePREADV,
	"pwritev":                AuditSystemCallTypePWRITEV,
	"rt_tgsigqueueinfo":      AuditSystemCallTypeRT_TGSIGQUEUEINFO,
	"perf_event_open":        AuditSystemCallTypePERF_EVENT_OPEN,
	"recvmmsg":               AuditSystemCallTypeRECVMMSG,
	"fanotify_init":          AuditSystemCallTypeFANOTIFY_INIT,
	"fanotify_mark":          AuditSystemCallTypeFANOTIFY_MARK,
	"prlimit64":              AuditSystemCallTypePRLIMIT64,
	"name_to_handle_at":      AuditSystemCallTypeNAME_TO_HANDLE_AT,
	"open_by_handle_at":      AuditSystemCallTypeOPEN_BY_HANDLE_AT,
	"clock_adjtime":          AuditSystemCallTypeCLOCK_ADJTIME,
	"syncfs":                 AuditSystemCallTypeSYNCFS,
	"sendmmsg":               AuditSystemCallTypeSENDMMSG,
	"setns":                  AuditSystemCallTypeSETNS,
	"getcpu":                 AuditSystemCallTypeGETCPU,
	"process_vm_readv":       AuditSystemCallTypePROCESS_VM_READV,
	"process_vm_writev":      AuditSystemCallTypePROCESS_VM_WRITEV,
	"kcmp":                   AuditSystemCallTypeKCMP,
	"finit_module":           AuditSystemCallTypeFINIT_MODULE,
	"stime":                  AuditSystemCallTypeSTIME,
}
