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
	"msgtype":       AuditFilterTypeMsgtype,
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
	AuditSystemCallTypeREAD                AuditSystemCallType = "read"
	AuditSystemCallTypeWRITE               AuditSystemCallType = "write"
	AuditSystemCallTypeOPEN                AuditSystemCallType = "open"
	AuditSystemCallTypeCLOSE               AuditSystemCallType = "close"
	AuditSystemCallTypeSTAT                AuditSystemCallType = "stat"
	AuditSystemCallTypeFSTAT               AuditSystemCallType = "fstat"
	AuditSystemCallTypeLSTAT               AuditSystemCallType = "lstat"
	AuditSystemCallTypePOLL                AuditSystemCallType = "poll"
	AuditSystemCallTypeLSEEK               AuditSystemCallType = "lseek"
	AuditSystemCallTypeMMAP                AuditSystemCallType = "mmap"
	AuditSystemCallTypeMPROTECT            AuditSystemCallType = "mprotect"
	AuditSystemCallTypeMUNMAP              AuditSystemCallType = "munmap"
	AuditSystemCallTypeBRK                 AuditSystemCallType = "brk"
	AuditSystemCallTypeRTSIGACTION         AuditSystemCallType = "rt_sigaction"
	AuditSystemCallTypeRTSIGPROCMASK       AuditSystemCallType = "rt_sigprocmask"
	AuditSystemCallTypeRTSIGRETURN         AuditSystemCallType = "rt_sigreturn"
	AuditSystemCallTypeIOCTL               AuditSystemCallType = "ioctl"
	AuditSystemCallTypePREAD64             AuditSystemCallType = "pread64"
	AuditSystemCallTypePWRITE64            AuditSystemCallType = "pwrite64"
	AuditSystemCallTypeREADV               AuditSystemCallType = "readv"
	AuditSystemCallTypeWRITEV              AuditSystemCallType = "writev"
	AuditSystemCallTypeACCESS              AuditSystemCallType = "access"
	AuditSystemCallTypePIPE                AuditSystemCallType = "pipe"
	AuditSystemCallTypeSELECT              AuditSystemCallType = "select"
	AuditSystemCallTypeSCHEDYIELD          AuditSystemCallType = "sched_yield"
	AuditSystemCallTypeMREMAP              AuditSystemCallType = "mremap"
	AuditSystemCallTypeMSYNC               AuditSystemCallType = "msync"
	AuditSystemCallTypeMINCORE             AuditSystemCallType = "mincore"
	AuditSystemCallTypeMADVISE             AuditSystemCallType = "madvise"
	AuditSystemCallTypeSHMGET              AuditSystemCallType = "shmget"
	AuditSystemCallTypeSHMAT               AuditSystemCallType = "shmat"
	AuditSystemCallTypeSHMCTL              AuditSystemCallType = "shmctl"
	AuditSystemCallTypeDUP                 AuditSystemCallType = "dup"
	AuditSystemCallTypeDUP2                AuditSystemCallType = "dup2"
	AuditSystemCallTypePAUSE               AuditSystemCallType = "pause"
	AuditSystemCallTypeNANOSLEEP           AuditSystemCallType = "nanosleep"
	AuditSystemCallTypeGETITIMER           AuditSystemCallType = "getitimer"
	AuditSystemCallTypeALARM               AuditSystemCallType = "alarm"
	AuditSystemCallTypeSETITIMER           AuditSystemCallType = "setitimer"
	AuditSystemCallTypeGETPID              AuditSystemCallType = "getpid"
	AuditSystemCallTypeSENDFILE            AuditSystemCallType = "sendfile"
	AuditSystemCallTypeSOCKET              AuditSystemCallType = "socket"
	AuditSystemCallTypeCONNECT             AuditSystemCallType = "connect"
	AuditSystemCallTypeACCEPT              AuditSystemCallType = "accept"
	AuditSystemCallTypeSENDTO              AuditSystemCallType = "sendto"
	AuditSystemCallTypeRECVFROM            AuditSystemCallType = "recvfrom"
	AuditSystemCallTypeSENDMSG             AuditSystemCallType = "sendmsg"
	AuditSystemCallTypeRECVMSG             AuditSystemCallType = "recvmsg"
	AuditSystemCallTypeSHUTDOWN            AuditSystemCallType = "shutdown"
	AuditSystemCallTypeBIND                AuditSystemCallType = "bind"
	AuditSystemCallTypeLISTEN              AuditSystemCallType = "listen"
	AuditSystemCallTypeGETSOCKNAME         AuditSystemCallType = "getsockname"
	AuditSystemCallTypeGETPEERNAME         AuditSystemCallType = "getpeername"
	AuditSystemCallTypeSOCKETPAIR          AuditSystemCallType = "socketpair"
	AuditSystemCallTypeSETSOCKOPT          AuditSystemCallType = "setsockopt"
	AuditSystemCallTypeGETSOCKOPT          AuditSystemCallType = "getsockopt"
	AuditSystemCallTypeCLONE               AuditSystemCallType = "clone"
	AuditSystemCallTypeFORK                AuditSystemCallType = "fork"
	AuditSystemCallTypeVFORK               AuditSystemCallType = "vfork"
	AuditSystemCallTypeEXECVE              AuditSystemCallType = "execve"
	AuditSystemCallTypeEXIT                AuditSystemCallType = "exit"
	AuditSystemCallTypeWAIT4               AuditSystemCallType = "wait4"
	AuditSystemCallTypeKILL                AuditSystemCallType = "kill"
	AuditSystemCallTypeUNAME               AuditSystemCallType = "uname"
	AuditSystemCallTypeSEMGET              AuditSystemCallType = "semget"
	AuditSystemCallTypeSEMOP               AuditSystemCallType = "semop"
	AuditSystemCallTypeSEMCTL              AuditSystemCallType = "semctl"
	AuditSystemCallTypeSHMDT               AuditSystemCallType = "shmdt"
	AuditSystemCallTypeMSGGET              AuditSystemCallType = "msgget"
	AuditSystemCallTypeMSGSND              AuditSystemCallType = "msgsnd"
	AuditSystemCallTypeMSGRCV              AuditSystemCallType = "msgrcv"
	AuditSystemCallTypeMSGCTL              AuditSystemCallType = "msgctl"
	AuditSystemCallTypeFCNTL               AuditSystemCallType = "fcntl"
	AuditSystemCallTypeFLOCK               AuditSystemCallType = "flock"
	AuditSystemCallTypeFSYNC               AuditSystemCallType = "fsync"
	AuditSystemCallTypeFDATASYNC           AuditSystemCallType = "fdatasync"
	AuditSystemCallTypeTRUNCATE            AuditSystemCallType = "truncate"
	AuditSystemCallTypeFTRUNCATE           AuditSystemCallType = "ftruncate"
	AuditSystemCallTypeGETDENTS            AuditSystemCallType = "getdents"
	AuditSystemCallTypeGETCWD              AuditSystemCallType = "getcwd"
	AuditSystemCallTypeCHDIR               AuditSystemCallType = "chdir"
	AuditSystemCallTypeFCHDIR              AuditSystemCallType = "fchdir"
	AuditSystemCallTypeRENAME              AuditSystemCallType = "rename"
	AuditSystemCallTypeMKDIR               AuditSystemCallType = "mkdir"
	AuditSystemCallTypeRMDIR               AuditSystemCallType = "rmdir"
	AuditSystemCallTypeCREAT               AuditSystemCallType = "creat"
	AuditSystemCallTypeLINK                AuditSystemCallType = "link"
	AuditSystemCallTypeUNLINK              AuditSystemCallType = "unlink"
	AuditSystemCallTypeSYMLINK             AuditSystemCallType = "symlink"
	AuditSystemCallTypeREADLINK            AuditSystemCallType = "readlink"
	AuditSystemCallTypeCHMOD               AuditSystemCallType = "chmod"
	AuditSystemCallTypeFCHMOD              AuditSystemCallType = "fchmod"
	AuditSystemCallTypeCHOWN               AuditSystemCallType = "chown"
	AuditSystemCallTypeFCHOWN              AuditSystemCallType = "fchown"
	AuditSystemCallTypeLCHOWN              AuditSystemCallType = "lchown"
	AuditSystemCallTypeUMASK               AuditSystemCallType = "umask"
	AuditSystemCallTypeGETTIMEOFDAY        AuditSystemCallType = "gettimeofday"
	AuditSystemCallTypeGETRLIMIT           AuditSystemCallType = "getrlimit"
	AuditSystemCallTypeGETRUSAGE           AuditSystemCallType = "getrusage"
	AuditSystemCallTypeSYSINFO             AuditSystemCallType = "sysinfo"
	AuditSystemCallTypeTIMES               AuditSystemCallType = "times"
	AuditSystemCallTypePTRACE              AuditSystemCallType = "ptrace"
	AuditSystemCallTypeGETUID              AuditSystemCallType = "getuid"
	AuditSystemCallTypeSYSLOG              AuditSystemCallType = "syslog"
	AuditSystemCallTypeGETGID              AuditSystemCallType = "getgid"
	AuditSystemCallTypeSETUID              AuditSystemCallType = "setuid"
	AuditSystemCallTypeSETGID              AuditSystemCallType = "setgid"
	AuditSystemCallTypeGETEUID             AuditSystemCallType = "geteuid"
	AuditSystemCallTypeGETEGID             AuditSystemCallType = "getegid"
	AuditSystemCallTypeSETPGID             AuditSystemCallType = "setpgid"
	AuditSystemCallTypeGETPPID             AuditSystemCallType = "getppid"
	AuditSystemCallTypeGETPGRP             AuditSystemCallType = "getpgrp"
	AuditSystemCallTypeSETSID              AuditSystemCallType = "setsid"
	AuditSystemCallTypeSETREUID            AuditSystemCallType = "setreuid"
	AuditSystemCallTypeSETREGID            AuditSystemCallType = "setregid"
	AuditSystemCallTypeGETGROUPS           AuditSystemCallType = "getgroups"
	AuditSystemCallTypeSETGROUPS           AuditSystemCallType = "setgroups"
	AuditSystemCallTypeSETRESUID           AuditSystemCallType = "setresuid"
	AuditSystemCallTypeGETRESUID           AuditSystemCallType = "getresuid"
	AuditSystemCallTypeSETRESGID           AuditSystemCallType = "setresgid"
	AuditSystemCallTypeGETRESGID           AuditSystemCallType = "getresgid"
	AuditSystemCallTypeGETPGID             AuditSystemCallType = "getpgid"
	AuditSystemCallTypeSETFSUID            AuditSystemCallType = "setfsuid"
	AuditSystemCallTypeSETFSGID            AuditSystemCallType = "setfsgid"
	AuditSystemCallTypeGETSID              AuditSystemCallType = "getsid"
	AuditSystemCallTypeCAPGET              AuditSystemCallType = "capget"
	AuditSystemCallTypeCAPSET              AuditSystemCallType = "capset"
	AuditSystemCallTypeRTSIGPENDING        AuditSystemCallType = "rt_sigpending"
	AuditSystemCallTypeRTSIGTIMEDWAIT      AuditSystemCallType = "rt_sigtimedwait"
	AuditSystemCallTypeRTSIGQUEUEINFO      AuditSystemCallType = "rt_sigqueueinfo"
	AuditSystemCallTypeRTSIGSUSPEND        AuditSystemCallType = "rt_sigsuspend"
	AuditSystemCallTypeSIGALTSTACK         AuditSystemCallType = "sigaltstack"
	AuditSystemCallTypeUTIME               AuditSystemCallType = "utime"
	AuditSystemCallTypeMKNOD               AuditSystemCallType = "mknod"
	AuditSystemCallTypeUSELIB              AuditSystemCallType = "uselib"
	AuditSystemCallTypePERSONALITY         AuditSystemCallType = "personality"
	AuditSystemCallTypeUSTAT               AuditSystemCallType = "ustat"
	AuditSystemCallTypeSTATFS              AuditSystemCallType = "statfs"
	AuditSystemCallTypeFSTATFS             AuditSystemCallType = "fstatfs"
	AuditSystemCallTypeSYSFS               AuditSystemCallType = "sysfs"
	AuditSystemCallTypeGETPRIORITY         AuditSystemCallType = "getpriority"
	AuditSystemCallTypeSETPRIORITY         AuditSystemCallType = "setpriority"
	AuditSystemCallTypeSCHEDSETPARAM       AuditSystemCallType = "sched_setparam"
	AuditSystemCallTypeSCHEDGETPARAM       AuditSystemCallType = "sched_getparam"
	AuditSystemCallTypeSCHEDSETSCHEDULER   AuditSystemCallType = "sched_setscheduler"
	AuditSystemCallTypeSCHEDGETSCHEDULER   AuditSystemCallType = "sched_getscheduler"
	AuditSystemCallTypeSCHEDGETPRIORITYMAX AuditSystemCallType = "sched_get_priority_max"
	AuditSystemCallTypeSCHEDGETPRIORITYMIN AuditSystemCallType = "sched_get_priority_min"
	AuditSystemCallTypeSCHEDRRGETINTERVAL  AuditSystemCallType = "sched_rr_get_interval"
	AuditSystemCallTypeMLOCK               AuditSystemCallType = "mlock"
	AuditSystemCallTypeMUNLOCK             AuditSystemCallType = "munlock"
	AuditSystemCallTypeMLOCKALL            AuditSystemCallType = "mlockall"
	AuditSystemCallTypeMUNLOCKALL          AuditSystemCallType = "munlockall"
	AuditSystemCallTypeVHANGUP             AuditSystemCallType = "vhangup"
	AuditSystemCallTypeMODIFYLDT           AuditSystemCallType = "modify_ldt"
	AuditSystemCallTypePIVOTROOT           AuditSystemCallType = "pivot_root"
	AuditSystemCallTypeSYSCTL              AuditSystemCallType = "_sysctl"
	AuditSystemCallTypePRCTL               AuditSystemCallType = "prctl"
	AuditSystemCallTypeARCHPRCTL           AuditSystemCallType = "arch_prctl"
	AuditSystemCallTypeADJTIMEX            AuditSystemCallType = "adjtimex"
	AuditSystemCallTypeSETRLIMIT           AuditSystemCallType = "setrlimit"
	AuditSystemCallTypeCHROOT              AuditSystemCallType = "chroot"
	AuditSystemCallTypeSYNC                AuditSystemCallType = "sync"
	AuditSystemCallTypeACCT                AuditSystemCallType = "acct"
	AuditSystemCallTypeSETTIMEOFDAY        AuditSystemCallType = "settimeofday"
	AuditSystemCallTypeMOUNT               AuditSystemCallType = "mount"
	AuditSystemCallTypeUMOUNT2             AuditSystemCallType = "umount2"
	AuditSystemCallTypeSWAPON              AuditSystemCallType = "swapon"
	AuditSystemCallTypeSWAPOFF             AuditSystemCallType = "swapoff"
	AuditSystemCallTypeREBOOT              AuditSystemCallType = "reboot"
	AuditSystemCallTypeSETHOSTNAME         AuditSystemCallType = "sethostname"
	AuditSystemCallTypeSETDOMAINNAME       AuditSystemCallType = "setdomainname"
	AuditSystemCallTypeIOPL                AuditSystemCallType = "iopl"
	AuditSystemCallTypeIOPERM              AuditSystemCallType = "ioperm"
	AuditSystemCallTypeCREATEMODULE        AuditSystemCallType = "create_module"
	AuditSystemCallTypeINITMODULE          AuditSystemCallType = "init_module"
	AuditSystemCallTypeDELETEMODULE        AuditSystemCallType = "delete_module"
	AuditSystemCallTypeGETKERNELSYMS       AuditSystemCallType = "get_kernel_syms"
	AuditSystemCallTypeQUERYMODULE         AuditSystemCallType = "query_module"
	AuditSystemCallTypeQUOTACTL            AuditSystemCallType = "quotactl"
	AuditSystemCallTypeNFSSERVCTL          AuditSystemCallType = "nfsservctl"
	AuditSystemCallTypeGETPMSG             AuditSystemCallType = "getpmsg"
	AuditSystemCallTypePUTPMSG             AuditSystemCallType = "putpmsg"
	AuditSystemCallTypeAFSSYSCALL          AuditSystemCallType = "afs_syscall"
	AuditSystemCallTypeTUXCALL             AuditSystemCallType = "tuxcall"
	AuditSystemCallTypeSECURITY            AuditSystemCallType = "security"
	AuditSystemCallTypeGETTID              AuditSystemCallType = "gettid"
	AuditSystemCallTypeREADAHEAD           AuditSystemCallType = "readahead"
	AuditSystemCallTypeSETXATTR            AuditSystemCallType = "setxattr"
	AuditSystemCallTypeLSETXATTR           AuditSystemCallType = "lsetxattr"
	AuditSystemCallTypeFSETXATTR           AuditSystemCallType = "fsetxattr"
	AuditSystemCallTypeGETXATTR            AuditSystemCallType = "getxattr"
	AuditSystemCallTypeLGETXATTR           AuditSystemCallType = "lgetxattr"
	AuditSystemCallTypeFGETXATTR           AuditSystemCallType = "fgetxattr"
	AuditSystemCallTypeLISTXATTR           AuditSystemCallType = "listxattr"
	AuditSystemCallTypeLLISTXATTR          AuditSystemCallType = "llistxattr"
	AuditSystemCallTypeFLISTXATTR          AuditSystemCallType = "flistxattr"
	AuditSystemCallTypeREMOVEXATTR         AuditSystemCallType = "removexattr"
	AuditSystemCallTypeLREMOVEXATTR        AuditSystemCallType = "lremovexattr"
	AuditSystemCallTypeFREMOVEXATTR        AuditSystemCallType = "fremovexattr"
	AuditSystemCallTypeTKILL               AuditSystemCallType = "tkill"
	AuditSystemCallTypeTIME                AuditSystemCallType = "time"
	AuditSystemCallTypeFUTEX               AuditSystemCallType = "futex"
	AuditSystemCallTypeSCHEDSETAFFINITY    AuditSystemCallType = "sched_setaffinity"
	AuditSystemCallTypeSCHEDGETAFFINITY    AuditSystemCallType = "sched_getaffinity"
	AuditSystemCallTypeSETTHREADAREA       AuditSystemCallType = "set_thread_area"
	AuditSystemCallTypeIOSETUP             AuditSystemCallType = "io_setup"
	AuditSystemCallTypeIODESTROY           AuditSystemCallType = "io_destroy"
	AuditSystemCallTypeIOGETEVENTS         AuditSystemCallType = "io_getevents"
	AuditSystemCallTypeIOSUBMIT            AuditSystemCallType = "io_submit"
	AuditSystemCallTypeIOCANCEL            AuditSystemCallType = "io_cancel"
	AuditSystemCallTypeGETTHREADAREA       AuditSystemCallType = "get_thread_area"
	AuditSystemCallTypeLOOKUPDCOOKIE       AuditSystemCallType = "lookup_dcookie"
	AuditSystemCallTypeEPOLLCREATE         AuditSystemCallType = "epoll_create"
	AuditSystemCallTypeEPOLLCTLOLD         AuditSystemCallType = "epoll_ctl_old"
	AuditSystemCallTypeEPOLLWAITOLD        AuditSystemCallType = "epoll_wait_old"
	AuditSystemCallTypeREMAPFILEPAGES      AuditSystemCallType = "remap_file_pages"
	AuditSystemCallTypeGETDENTS64          AuditSystemCallType = "getdents64"
	AuditSystemCallTypeSETTIDADDRESS       AuditSystemCallType = "set_tid_address"
	AuditSystemCallTypeRESTARTSYSCALL      AuditSystemCallType = "restart_syscall"
	AuditSystemCallTypeSEMTIMEDOP          AuditSystemCallType = "semtimedop"
	AuditSystemCallTypeFADVISE64           AuditSystemCallType = "fadvise64"
	AuditSystemCallTypeTIMERCREATE         AuditSystemCallType = "timer_create"
	AuditSystemCallTypeTIMERSETTIME        AuditSystemCallType = "timer_settime"
	AuditSystemCallTypeTIMERGETTIME        AuditSystemCallType = "timer_gettime"
	AuditSystemCallTypeTIMERGETOVERRUN     AuditSystemCallType = "timer_getoverrun"
	AuditSystemCallTypeTIMERDELETE         AuditSystemCallType = "timer_delete"
	AuditSystemCallTypeCLOCKSETTIME        AuditSystemCallType = "clock_settime"
	AuditSystemCallTypeCLOCKGETTIME        AuditSystemCallType = "clock_gettime"
	AuditSystemCallTypeCLOCKGETRES         AuditSystemCallType = "clock_getres"
	AuditSystemCallTypeCLOCKNANOSLEEP      AuditSystemCallType = "clock_nanosleep"
	AuditSystemCallTypeEXITGROUP           AuditSystemCallType = "exit_group"
	AuditSystemCallTypeEPOLLWAIT           AuditSystemCallType = "epoll_wait"
	AuditSystemCallTypeEPOLLCTL            AuditSystemCallType = "epoll_ctl"
	AuditSystemCallTypeTGKILL              AuditSystemCallType = "tgkill"
	AuditSystemCallTypeUTIMES              AuditSystemCallType = "utimes"
	AuditSystemCallTypeVSERVER             AuditSystemCallType = "vserver"
	AuditSystemCallTypeMBIND               AuditSystemCallType = "mbind"
	AuditSystemCallTypeSETMEMPOLICY        AuditSystemCallType = "set_mempolicy"
	AuditSystemCallTypeGETMEMPOLICY        AuditSystemCallType = "get_mempolicy"
	AuditSystemCallTypeMQOPEN              AuditSystemCallType = "mq_open"
	AuditSystemCallTypeMQUNLINK            AuditSystemCallType = "mq_unlink"
	AuditSystemCallTypeMQTIMEDSEND         AuditSystemCallType = "mq_timedsend"
	AuditSystemCallTypeMQTIMEDRECEIVE      AuditSystemCallType = "mq_timedreceive"
	AuditSystemCallTypeMQNOTIFY            AuditSystemCallType = "mq_notify"
	AuditSystemCallTypeMQGETSETATTR        AuditSystemCallType = "mq_getsetattr"
	AuditSystemCallTypeKEXECLOAD           AuditSystemCallType = "kexec_load"
	AuditSystemCallTypeWAITID              AuditSystemCallType = "waitid"
	AuditSystemCallTypeADDKEY              AuditSystemCallType = "add_key"
	AuditSystemCallTypeREQUESTKEY          AuditSystemCallType = "request_key"
	AuditSystemCallTypeKEYCTL              AuditSystemCallType = "keyctl"
	AuditSystemCallTypeIOPRIOSET           AuditSystemCallType = "ioprio_set"
	AuditSystemCallTypeIOPRIOGET           AuditSystemCallType = "ioprio_get"
	AuditSystemCallTypeINOTIFYINIT         AuditSystemCallType = "inotify_init"
	AuditSystemCallTypeINOTIFYADDWATCH     AuditSystemCallType = "inotify_add_watch"
	AuditSystemCallTypeINOTIFYRMWATCH      AuditSystemCallType = "inotify_rm_watch"
	AuditSystemCallTypeMIGRATEPAGES        AuditSystemCallType = "migrate_pages"
	AuditSystemCallTypeOPENAT              AuditSystemCallType = "openat"
	AuditSystemCallTypeMKDIRAT             AuditSystemCallType = "mkdirat"
	AuditSystemCallTypeMKNODAT             AuditSystemCallType = "mknodat"
	AuditSystemCallTypeFCHOWNAT            AuditSystemCallType = "fchownat"
	AuditSystemCallTypeFUTIMESAT           AuditSystemCallType = "futimesat"
	AuditSystemCallTypeNEWFSTATAT          AuditSystemCallType = "newfstatat"
	AuditSystemCallTypeUNLINKAT            AuditSystemCallType = "unlinkat"
	AuditSystemCallTypeRENAMEAT            AuditSystemCallType = "renameat"
	AuditSystemCallTypeLINKAT              AuditSystemCallType = "linkat"
	AuditSystemCallTypeSYMLINKAT           AuditSystemCallType = "symlinkat"
	AuditSystemCallTypeREADLINKAT          AuditSystemCallType = "readlinkat"
	AuditSystemCallTypeFCHMODAT            AuditSystemCallType = "fchmodat"
	AuditSystemCallTypeFACCESSAT           AuditSystemCallType = "faccessat"
	AuditSystemCallTypePSELECT6            AuditSystemCallType = "pselect6"
	AuditSystemCallTypePPOLL               AuditSystemCallType = "ppoll"
	AuditSystemCallTypeUNSHARE             AuditSystemCallType = "unshare"
	AuditSystemCallTypeSETROBUSTLIST       AuditSystemCallType = "set_robust_list"
	AuditSystemCallTypeGETROBUSTLIST       AuditSystemCallType = "get_robust_list"
	AuditSystemCallTypeSPLICE              AuditSystemCallType = "splice"
	AuditSystemCallTypeTEE                 AuditSystemCallType = "tee"
	AuditSystemCallTypeSYNCFILERANGE       AuditSystemCallType = "sync_file_range"
	AuditSystemCallTypeVMSPLICE            AuditSystemCallType = "vmsplice"
	AuditSystemCallTypeMOVEPAGES           AuditSystemCallType = "move_pages"
	AuditSystemCallTypeUTIMENSAT           AuditSystemCallType = "utimensat"
	AuditSystemCallTypeEPOLLPWAIT          AuditSystemCallType = "epoll_pwait"
	AuditSystemCallTypeSIGNALFD            AuditSystemCallType = "signalfd"
	AuditSystemCallTypeTIMERFDCREATE       AuditSystemCallType = "timerfd_create"
	AuditSystemCallTypeEVENTFD             AuditSystemCallType = "eventfd"
	AuditSystemCallTypeFALLOCATE           AuditSystemCallType = "fallocate"
	AuditSystemCallTypeTIMERFDSETTIME      AuditSystemCallType = "timerfd_settime"
	AuditSystemCallTypeTIMERFDGETTIME      AuditSystemCallType = "timerfd_gettime"
	AuditSystemCallTypeACCEPT4             AuditSystemCallType = "accept4"
	AuditSystemCallTypeSIGNALFD4           AuditSystemCallType = "signalfd4"
	AuditSystemCallTypeEVENTFD2            AuditSystemCallType = "eventfd2"
	AuditSystemCallTypeEPOLLCREATE1        AuditSystemCallType = "epoll_create1"
	AuditSystemCallTypeDUP3                AuditSystemCallType = "dup3"
	AuditSystemCallTypePIPE2               AuditSystemCallType = "pipe2"
	AuditSystemCallTypeINOTIFYINIT1        AuditSystemCallType = "inotify_init1"
	AuditSystemCallTypePREADV              AuditSystemCallType = "preadv"
	AuditSystemCallTypePWRITEV             AuditSystemCallType = "pwritev"
	AuditSystemCallTypeRTTGSIGQUEUEINFO    AuditSystemCallType = "rt_tgsigqueueinfo"
	AuditSystemCallTypePERFEVENTOPEN       AuditSystemCallType = "perf_event_open"
	AuditSystemCallTypeRECVMMSG            AuditSystemCallType = "recvmmsg"
	AuditSystemCallTypeFANOTIFYINIT        AuditSystemCallType = "fanotify_init"
	AuditSystemCallTypeFANOTIFYMARK        AuditSystemCallType = "fanotify_mark"
	AuditSystemCallTypePRLIMIT64           AuditSystemCallType = "prlimit64"
	AuditSystemCallTypeNAMETOHANDLEAT      AuditSystemCallType = "name_to_handle_at"
	AuditSystemCallTypeOPENBYHANDLEAT      AuditSystemCallType = "open_by_handle_at"
	AuditSystemCallTypeCLOCKADJTIME        AuditSystemCallType = "clock_adjtime"
	AuditSystemCallTypeSYNCFS              AuditSystemCallType = "syncfs"
	AuditSystemCallTypeSENDMMSG            AuditSystemCallType = "sendmmsg"
	AuditSystemCallTypeSETNS               AuditSystemCallType = "setns"
	AuditSystemCallTypeGETCPU              AuditSystemCallType = "getcpu"
	AuditSystemCallTypePROCESSVMREADV      AuditSystemCallType = "process_vm_readv"
	AuditSystemCallTypePROCESSVMWRITEV     AuditSystemCallType = "process_vm_writev"
	AuditSystemCallTypeKCMP                AuditSystemCallType = "kcmp"
	AuditSystemCallTypeFINITMODULE         AuditSystemCallType = "finit_module"
	AuditSystemCallTypeSTIME               AuditSystemCallType = "stime"
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
	"rt_sigaction":           AuditSystemCallTypeRTSIGACTION,
	"rt_sigprocmask":         AuditSystemCallTypeRTSIGPROCMASK,
	"rt_sigreturn":           AuditSystemCallTypeRTSIGRETURN,
	"ioctl":                  AuditSystemCallTypeIOCTL,
	"pread64":                AuditSystemCallTypePREAD64,
	"pwrite64":               AuditSystemCallTypePWRITE64,
	"readv":                  AuditSystemCallTypeREADV,
	"writev":                 AuditSystemCallTypeWRITEV,
	"access":                 AuditSystemCallTypeACCESS,
	"pipe":                   AuditSystemCallTypePIPE,
	"select":                 AuditSystemCallTypeSELECT,
	"sched_yield":            AuditSystemCallTypeSCHEDYIELD,
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
	"rt_sigpending":          AuditSystemCallTypeRTSIGPENDING,
	"rt_sigtimedwait":        AuditSystemCallTypeRTSIGTIMEDWAIT,
	"rt_sigqueueinfo":        AuditSystemCallTypeRTSIGQUEUEINFO,
	"rt_sigsuspend":          AuditSystemCallTypeRTSIGSUSPEND,
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
	"sched_setparam":         AuditSystemCallTypeSCHEDSETPARAM,
	"sched_getparam":         AuditSystemCallTypeSCHEDGETPARAM,
	"sched_setscheduler":     AuditSystemCallTypeSCHEDSETSCHEDULER,
	"sched_getscheduler":     AuditSystemCallTypeSCHEDGETSCHEDULER,
	"sched_get_priority_max": AuditSystemCallTypeSCHEDGETPRIORITYMAX,
	"sched_get_priority_min": AuditSystemCallTypeSCHEDGETPRIORITYMIN,
	"sched_rr_get_interval":  AuditSystemCallTypeSCHEDRRGETINTERVAL,
	"mlock":                  AuditSystemCallTypeMLOCK,
	"munlock":                AuditSystemCallTypeMUNLOCK,
	"mlockall":               AuditSystemCallTypeMLOCKALL,
	"munlockall":             AuditSystemCallTypeMUNLOCKALL,
	"vhangup":                AuditSystemCallTypeVHANGUP,
	"modify_ldt":             AuditSystemCallTypeMODIFYLDT,
	"pivot_root":             AuditSystemCallTypePIVOTROOT,
	"_sysctl":                AuditSystemCallTypeSYSCTL,
	"prctl":                  AuditSystemCallTypePRCTL,
	"arch_prctl":             AuditSystemCallTypeARCHPRCTL,
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
	"create_module":          AuditSystemCallTypeCREATEMODULE,
	"init_module":            AuditSystemCallTypeINITMODULE,
	"delete_module":          AuditSystemCallTypeDELETEMODULE,
	"get_kernel_syms":        AuditSystemCallTypeGETKERNELSYMS,
	"query_module":           AuditSystemCallTypeQUERYMODULE,
	"quotactl":               AuditSystemCallTypeQUOTACTL,
	"nfsservctl":             AuditSystemCallTypeNFSSERVCTL,
	"getpmsg":                AuditSystemCallTypeGETPMSG,
	"putpmsg":                AuditSystemCallTypePUTPMSG,
	"afs_syscall":            AuditSystemCallTypeAFSSYSCALL,
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
	"sched_setaffinity":      AuditSystemCallTypeSCHEDSETAFFINITY,
	"sched_getaffinity":      AuditSystemCallTypeSCHEDGETAFFINITY,
	"set_thread_area":        AuditSystemCallTypeSETTHREADAREA,
	"io_setup":               AuditSystemCallTypeIOSETUP,
	"io_destroy":             AuditSystemCallTypeIODESTROY,
	"io_getevents":           AuditSystemCallTypeIOGETEVENTS,
	"io_submit":              AuditSystemCallTypeIOSUBMIT,
	"io_cancel":              AuditSystemCallTypeIOCANCEL,
	"get_thread_area":        AuditSystemCallTypeGETTHREADAREA,
	"lookup_dcookie":         AuditSystemCallTypeLOOKUPDCOOKIE,
	"epoll_create":           AuditSystemCallTypeEPOLLCREATE,
	"epoll_ctl_old":          AuditSystemCallTypeEPOLLCTLOLD,
	"epoll_wait_old":         AuditSystemCallTypeEPOLLWAITOLD,
	"remap_file_pages":       AuditSystemCallTypeREMAPFILEPAGES,
	"getdents64":             AuditSystemCallTypeGETDENTS64,
	"set_tid_address":        AuditSystemCallTypeSETTIDADDRESS,
	"restart_syscall":        AuditSystemCallTypeRESTARTSYSCALL,
	"semtimedop":             AuditSystemCallTypeSEMTIMEDOP,
	"fadvise64":              AuditSystemCallTypeFADVISE64,
	"timer_create":           AuditSystemCallTypeTIMERCREATE,
	"timer_settime":          AuditSystemCallTypeTIMERSETTIME,
	"timer_gettime":          AuditSystemCallTypeTIMERGETTIME,
	"timer_getoverrun":       AuditSystemCallTypeTIMERGETOVERRUN,
	"timer_delete":           AuditSystemCallTypeTIMERDELETE,
	"clock_settime":          AuditSystemCallTypeCLOCKSETTIME,
	"clock_gettime":          AuditSystemCallTypeCLOCKGETTIME,
	"clock_getres":           AuditSystemCallTypeCLOCKGETRES,
	"clock_nanosleep":        AuditSystemCallTypeCLOCKNANOSLEEP,
	"exit_group":             AuditSystemCallTypeEXITGROUP,
	"epoll_wait":             AuditSystemCallTypeEPOLLWAIT,
	"epoll_ctl":              AuditSystemCallTypeEPOLLCTL,
	"tgkill":                 AuditSystemCallTypeTGKILL,
	"utimes":                 AuditSystemCallTypeUTIMES,
	"vserver":                AuditSystemCallTypeVSERVER,
	"mbind":                  AuditSystemCallTypeMBIND,
	"set_mempolicy":          AuditSystemCallTypeSETMEMPOLICY,
	"get_mempolicy":          AuditSystemCallTypeGETMEMPOLICY,
	"mq_open":                AuditSystemCallTypeMQOPEN,
	"mq_unlink":              AuditSystemCallTypeMQUNLINK,
	"mq_timedsend":           AuditSystemCallTypeMQTIMEDSEND,
	"mq_timedreceive":        AuditSystemCallTypeMQTIMEDRECEIVE,
	"mq_notify":              AuditSystemCallTypeMQNOTIFY,
	"mq_getsetattr":          AuditSystemCallTypeMQGETSETATTR,
	"kexec_load":             AuditSystemCallTypeKEXECLOAD,
	"waitid":                 AuditSystemCallTypeWAITID,
	"add_key":                AuditSystemCallTypeADDKEY,
	"request_key":            AuditSystemCallTypeREQUESTKEY,
	"keyctl":                 AuditSystemCallTypeKEYCTL,
	"ioprio_set":             AuditSystemCallTypeIOPRIOSET,
	"ioprio_get":             AuditSystemCallTypeIOPRIOGET,
	"inotify_init":           AuditSystemCallTypeINOTIFYINIT,
	"inotify_add_watch":      AuditSystemCallTypeINOTIFYADDWATCH,
	"inotify_rm_watch":       AuditSystemCallTypeINOTIFYRMWATCH,
	"migrate_pages":          AuditSystemCallTypeMIGRATEPAGES,
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
	"set_robust_list":        AuditSystemCallTypeSETROBUSTLIST,
	"get_robust_list":        AuditSystemCallTypeGETROBUSTLIST,
	"splice":                 AuditSystemCallTypeSPLICE,
	"tee":                    AuditSystemCallTypeTEE,
	"sync_file_range":        AuditSystemCallTypeSYNCFILERANGE,
	"vmsplice":               AuditSystemCallTypeVMSPLICE,
	"move_pages":             AuditSystemCallTypeMOVEPAGES,
	"utimensat":              AuditSystemCallTypeUTIMENSAT,
	"epoll_pwait":            AuditSystemCallTypeEPOLLPWAIT,
	"signalfd":               AuditSystemCallTypeSIGNALFD,
	"timerfd_create":         AuditSystemCallTypeTIMERFDCREATE,
	"eventfd":                AuditSystemCallTypeEVENTFD,
	"fallocate":              AuditSystemCallTypeFALLOCATE,
	"timerfd_settime":        AuditSystemCallTypeTIMERFDSETTIME,
	"timerfd_gettime":        AuditSystemCallTypeTIMERFDGETTIME,
	"accept4":                AuditSystemCallTypeACCEPT4,
	"signalfd4":              AuditSystemCallTypeSIGNALFD4,
	"eventfd2":               AuditSystemCallTypeEVENTFD2,
	"epoll_create1":          AuditSystemCallTypeEPOLLCREATE1,
	"dup3":                   AuditSystemCallTypeDUP3,
	"pipe2":                  AuditSystemCallTypePIPE2,
	"inotify_init1":          AuditSystemCallTypeINOTIFYINIT1,
	"preadv":                 AuditSystemCallTypePREADV,
	"pwritev":                AuditSystemCallTypePWRITEV,
	"rt_tgsigqueueinfo":      AuditSystemCallTypeRTTGSIGQUEUEINFO,
	"perf_event_open":        AuditSystemCallTypePERFEVENTOPEN,
	"recvmmsg":               AuditSystemCallTypeRECVMMSG,
	"fanotify_init":          AuditSystemCallTypeFANOTIFYINIT,
	"fanotify_mark":          AuditSystemCallTypeFANOTIFYMARK,
	"prlimit64":              AuditSystemCallTypePRLIMIT64,
	"name_to_handle_at":      AuditSystemCallTypeNAMETOHANDLEAT,
	"open_by_handle_at":      AuditSystemCallTypeOPENBYHANDLEAT,
	"clock_adjtime":          AuditSystemCallTypeCLOCKADJTIME,
	"syncfs":                 AuditSystemCallTypeSYNCFS,
	"sendmmsg":               AuditSystemCallTypeSENDMMSG,
	"setns":                  AuditSystemCallTypeSETNS,
	"getcpu":                 AuditSystemCallTypeGETCPU,
	"process_vm_readv":       AuditSystemCallTypePROCESSVMREADV,
	"process_vm_writev":      AuditSystemCallTypePROCESSVMWRITEV,
	"kcmp":                   AuditSystemCallTypeKCMP,
	"finit_module":           AuditSystemCallTypeFINITMODULE,
	"stime":                  AuditSystemCallTypeSTIME,
}
