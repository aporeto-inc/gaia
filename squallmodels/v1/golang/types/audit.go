package types

// FilterID are the valid IDs of the audit filters
type FilterID string

const (
	// Path filter
	Path FilterID = "path"
	// Msgtype filter
	Msgtype FilterID = "msgtype"
	// SbjType filter
	SbjType FilterID = "subjType"
	// Auid filter
	Auid FilterID = "auid"
	// Success filter
	Success FilterID = "success"
	// Perm filter
	Perm FilterID = "perm"
	// Dir filter
	Dir FilterID = "dir"
)

// FilterOperator is the operator for filters
type FilterOperator string

const (
	// Equal operator
	Equal FilterOperator = "="
	// NotEqual operator
	NotEqual FilterOperator = "!="
	// Greater operator
	Greater FilterOperator = ">"
	// LessThan operator
	LessThan FilterOperator = "<"
	// GreaterOrEqual operator
	GreaterOrEqual FilterOperator = ">="
	// LessThanOrEqual operator
	LessThanOrEqual FilterOperator = "<="
)

// AuditFilter is a signle audit filter rule
type AuditFilter struct {
	Key      FilterID
	Operator FilterOperator
	Value    string
}

// SystemCallType is the type for the system calls
type SystemCallType uint16

const (
	READ                   SystemCallType = 0
	WRITE                  SystemCallType = 1
	OPEN                   SystemCallType = 2
	CLOSE                  SystemCallType = 3
	STAT                   SystemCallType = 4
	FSTAT                  SystemCallType = 5
	LSTAT                  SystemCallType = 6
	POLL                   SystemCallType = 7
	LSEEK                  SystemCallType = 8
	MMAP                   SystemCallType = 9
	MPROTECT               SystemCallType = 10
	MUNMAP                 SystemCallType = 11
	BRK                    SystemCallType = 12
	RT_SIGACTION           SystemCallType = 13
	RT_SIGPROCMASK         SystemCallType = 14
	RT_SIGRETURN           SystemCallType = 15
	IOCTL                  SystemCallType = 16
	PREAD64                SystemCallType = 17
	PWRITE64               SystemCallType = 18
	READV                  SystemCallType = 19
	WRITEV                 SystemCallType = 20
	ACCESS                 SystemCallType = 21
	PIPE                   SystemCallType = 22
	SELECT                 SystemCallType = 23
	SCHED_YIELD            SystemCallType = 24
	MREMAP                 SystemCallType = 25
	MSYNC                  SystemCallType = 26
	MINCORE                SystemCallType = 27
	MADVISE                SystemCallType = 28
	SHMGET                 SystemCallType = 29
	SHMAT                  SystemCallType = 30
	SHMCTL                 SystemCallType = 31
	DUP                    SystemCallType = 32
	DUP2                   SystemCallType = 33
	PAUSE                  SystemCallType = 34
	NANOSLEEP              SystemCallType = 35
	GETITIMER              SystemCallType = 36
	ALARM                  SystemCallType = 37
	SETITIMER              SystemCallType = 38
	GETPID                 SystemCallType = 39
	SENDFILE               SystemCallType = 40
	SOCKET                 SystemCallType = 41
	CONNECT                SystemCallType = 42
	ACCEPT                 SystemCallType = 43
	SENDTO                 SystemCallType = 44
	RECVFROM               SystemCallType = 45
	SENDMSG                SystemCallType = 46
	RECVMSG                SystemCallType = 47
	SHUTDOWN               SystemCallType = 48
	BIND                   SystemCallType = 49
	LISTEN                 SystemCallType = 50
	GETSOCKNAME            SystemCallType = 51
	GETPEERNAME            SystemCallType = 52
	SOCKETPAIR             SystemCallType = 53
	SETSOCKOPT             SystemCallType = 54
	GETSOCKOPT             SystemCallType = 55
	CLONE                  SystemCallType = 56
	FORK                   SystemCallType = 57
	VFORK                  SystemCallType = 58
	EXECVE                 SystemCallType = 59
	EXIT                   SystemCallType = 60
	WAIT4                  SystemCallType = 61
	KILL                   SystemCallType = 62
	UNAME                  SystemCallType = 63
	SEMGET                 SystemCallType = 64
	SEMOP                  SystemCallType = 65
	SEMCTL                 SystemCallType = 66
	SHMDT                  SystemCallType = 67
	MSGGET                 SystemCallType = 68
	MSGSND                 SystemCallType = 69
	MSGRCV                 SystemCallType = 70
	MSGCTL                 SystemCallType = 71
	FCNTL                  SystemCallType = 72
	FLOCK                  SystemCallType = 73
	FSYNC                  SystemCallType = 74
	FDATASYNC              SystemCallType = 75
	TRUNCATE               SystemCallType = 76
	FTRUNCATE              SystemCallType = 77
	GETDENTS               SystemCallType = 78
	GETCWD                 SystemCallType = 79
	CHDIR                  SystemCallType = 80
	FCHDIR                 SystemCallType = 81
	RENAME                 SystemCallType = 82
	MKDIR                  SystemCallType = 83
	RMDIR                  SystemCallType = 84
	CREAT                  SystemCallType = 85
	LINK                   SystemCallType = 86
	UNLINK                 SystemCallType = 87
	SYMLINK                SystemCallType = 88
	READLINK               SystemCallType = 89
	CHMOD                  SystemCallType = 90
	FCHMOD                 SystemCallType = 91
	CHOWN                  SystemCallType = 92
	FCHOWN                 SystemCallType = 93
	LCHOWN                 SystemCallType = 94
	UMASK                  SystemCallType = 95
	GETTIMEOFDAY           SystemCallType = 96
	GETRLIMIT              SystemCallType = 97
	GETRUSAGE              SystemCallType = 98
	SYSINFO                SystemCallType = 99
	TIMES                  SystemCallType = 100
	PTRACE                 SystemCallType = 101
	GETUID                 SystemCallType = 102
	SYSLOG                 SystemCallType = 103
	GETGID                 SystemCallType = 104
	SETUID                 SystemCallType = 105
	SETGID                 SystemCallType = 106
	GETEUID                SystemCallType = 107
	GETEGID                SystemCallType = 108
	SETPGID                SystemCallType = 109
	GETPPID                SystemCallType = 110
	GETPGRP                SystemCallType = 111
	SETSID                 SystemCallType = 112
	SETREUID               SystemCallType = 113
	SETREGID               SystemCallType = 114
	GETGROUPS              SystemCallType = 115
	SETGROUPS              SystemCallType = 116
	SETRESUID              SystemCallType = 117
	GETRESUID              SystemCallType = 118
	SETRESGID              SystemCallType = 119
	GETRESGID              SystemCallType = 120
	GETPGID                SystemCallType = 121
	SETFSUID               SystemCallType = 122
	SETFSGID               SystemCallType = 123
	GETSID                 SystemCallType = 124
	CAPGET                 SystemCallType = 125
	CAPSET                 SystemCallType = 126
	RT_SIGPENDING          SystemCallType = 127
	RT_SIGTIMEDWAIT        SystemCallType = 128
	RT_SIGQUEUEINFO        SystemCallType = 129
	RT_SIGSUSPEND          SystemCallType = 130
	SIGALTSTACK            SystemCallType = 131
	UTIME                  SystemCallType = 132
	MKNOD                  SystemCallType = 133
	USELIB                 SystemCallType = 134
	PERSONALITY            SystemCallType = 135
	USTAT                  SystemCallType = 136
	STATFS                 SystemCallType = 137
	FSTATFS                SystemCallType = 138
	SYSFS                  SystemCallType = 139
	GETPRIORITY            SystemCallType = 140
	SETPRIORITY            SystemCallType = 141
	SCHED_SETPARAM         SystemCallType = 142
	SCHED_GETPARAM         SystemCallType = 143
	SCHED_SETSCHEDULER     SystemCallType = 144
	SCHED_GETSCHEDULER     SystemCallType = 145
	SCHED_GET_PRIORITY_MAX SystemCallType = 146
	SCHED_GET_PRIORITY_MIN SystemCallType = 147
	SCHED_RR_GET_INTERVAL  SystemCallType = 148
	MLOCK                  SystemCallType = 149
	MUNLOCK                SystemCallType = 150
	MLOCKALL               SystemCallType = 151
	MUNLOCKALL             SystemCallType = 152
	VHANGUP                SystemCallType = 153
	MODIFY_LDT             SystemCallType = 154
	PIVOT_ROOT             SystemCallType = 155
	_SYSCTL                SystemCallType = 156
	PRCTL                  SystemCallType = 157
	ARCH_PRCTL             SystemCallType = 158
	ADJTIMEX               SystemCallType = 159
	SETRLIMIT              SystemCallType = 160
	CHROOT                 SystemCallType = 161
	SYNC                   SystemCallType = 162
	ACCT                   SystemCallType = 163
	SETTIMEOFDAY           SystemCallType = 164
	MOUNT                  SystemCallType = 165
	UMOUNT2                SystemCallType = 166
	SWAPON                 SystemCallType = 167
	SWAPOFF                SystemCallType = 168
	REBOOT                 SystemCallType = 169
	SETHOSTNAME            SystemCallType = 170
	SETDOMAINNAME          SystemCallType = 171
	IOPL                   SystemCallType = 172
	IOPERM                 SystemCallType = 173
	CREATE_MODULE          SystemCallType = 174
	INIT_MODULE            SystemCallType = 175
	DELETE_MODULE          SystemCallType = 176
	GET_KERNEL_SYMS        SystemCallType = 177
	QUERY_MODULE           SystemCallType = 178
	QUOTACTL               SystemCallType = 179
	NFSSERVCTL             SystemCallType = 180
	GETPMSG                SystemCallType = 181
	PUTPMSG                SystemCallType = 182
	AFS_SYSCALL            SystemCallType = 183
	TUXCALL                SystemCallType = 184
	SECURITY               SystemCallType = 185
	GETTID                 SystemCallType = 186
	READAHEAD              SystemCallType = 187
	SETXATTR               SystemCallType = 188
	LSETXATTR              SystemCallType = 189
	FSETXATTR              SystemCallType = 190
	GETXATTR               SystemCallType = 191
	LGETXATTR              SystemCallType = 192
	FGETXATTR              SystemCallType = 193
	LISTXATTR              SystemCallType = 194
	LLISTXATTR             SystemCallType = 195
	FLISTXATTR             SystemCallType = 196
	REMOVEXATTR            SystemCallType = 197
	LREMOVEXATTR           SystemCallType = 198
	FREMOVEXATTR           SystemCallType = 199
	TKILL                  SystemCallType = 200
	TIME                   SystemCallType = 201
	FUTEX                  SystemCallType = 202
	SCHED_SETAFFINITY      SystemCallType = 203
	SCHED_GETAFFINITY      SystemCallType = 204
	SET_THREAD_AREA        SystemCallType = 205
	IO_SETUP               SystemCallType = 206
	IO_DESTROY             SystemCallType = 207
	IO_GETEVENTS           SystemCallType = 208
	IO_SUBMIT              SystemCallType = 209
	IO_CANCEL              SystemCallType = 210
	GET_THREAD_AREA        SystemCallType = 211
	LOOKUP_DCOOKIE         SystemCallType = 212
	EPOLL_CREATE           SystemCallType = 213
	EPOLL_CTL_OLD          SystemCallType = 214
	EPOLL_WAIT_OLD         SystemCallType = 215
	REMAP_FILE_PAGES       SystemCallType = 216
	GETDENTS64             SystemCallType = 217
	SET_TID_ADDRESS        SystemCallType = 218
	RESTART_SYSCALL        SystemCallType = 219
	SEMTIMEDOP             SystemCallType = 220
	FADVISE64              SystemCallType = 221
	TIMER_CREATE           SystemCallType = 222
	TIMER_SETTIME          SystemCallType = 223
	TIMER_GETTIME          SystemCallType = 224
	TIMER_GETOVERRUN       SystemCallType = 225
	TIMER_DELETE           SystemCallType = 226
	CLOCK_SETTIME          SystemCallType = 227
	CLOCK_GETTIME          SystemCallType = 228
	CLOCK_GETRES           SystemCallType = 229
	CLOCK_NANOSLEEP        SystemCallType = 230
	EXIT_GROUP             SystemCallType = 231
	EPOLL_WAIT             SystemCallType = 232
	EPOLL_CTL              SystemCallType = 233
	TGKILL                 SystemCallType = 234
	UTIMES                 SystemCallType = 235
	VSERVER                SystemCallType = 236
	MBIND                  SystemCallType = 237
	SET_MEMPOLICY          SystemCallType = 238
	GET_MEMPOLICY          SystemCallType = 239
	MQ_OPEN                SystemCallType = 240
	MQ_UNLINK              SystemCallType = 241
	MQ_TIMEDSEND           SystemCallType = 242
	MQ_TIMEDRECEIVE        SystemCallType = 243
	MQ_NOTIFY              SystemCallType = 244
	MQ_GETSETATTR          SystemCallType = 245
	KEXEC_LOAD             SystemCallType = 246
	WAITID                 SystemCallType = 247
	ADD_KEY                SystemCallType = 248
	REQUEST_KEY            SystemCallType = 249
	KEYCTL                 SystemCallType = 250
	IOPRIO_SET             SystemCallType = 251
	IOPRIO_GET             SystemCallType = 252
	INOTIFY_INIT           SystemCallType = 253
	INOTIFY_ADD_WATCH      SystemCallType = 254
	INOTIFY_RM_WATCH       SystemCallType = 255
	MIGRATE_PAGES          SystemCallType = 256
	OPENAT                 SystemCallType = 257
	MKDIRAT                SystemCallType = 258
	MKNODAT                SystemCallType = 259
	FCHOWNAT               SystemCallType = 260
	FUTIMESAT              SystemCallType = 261
	NEWFSTATAT             SystemCallType = 262
	UNLINKAT               SystemCallType = 263
	RENAMEAT               SystemCallType = 264
	LINKAT                 SystemCallType = 265
	SYMLINKAT              SystemCallType = 266
	READLINKAT             SystemCallType = 267
	FCHMODAT               SystemCallType = 268
	FACCESSAT              SystemCallType = 269
	PSELECT6               SystemCallType = 270
	PPOLL                  SystemCallType = 271
	UNSHARE                SystemCallType = 272
	SET_ROBUST_LIST        SystemCallType = 273
	GET_ROBUST_LIST        SystemCallType = 274
	SPLICE                 SystemCallType = 275
	TEE                    SystemCallType = 276
	SYNC_FILE_RANGE        SystemCallType = 277
	VMSPLICE               SystemCallType = 278
	MOVE_PAGES             SystemCallType = 279
	UTIMENSAT              SystemCallType = 280
	EPOLL_PWAIT            SystemCallType = 281
	SIGNALFD               SystemCallType = 282
	TIMERFD_CREATE         SystemCallType = 283
	EVENTFD                SystemCallType = 284
	FALLOCATE              SystemCallType = 285
	TIMERFD_SETTIME        SystemCallType = 286
	TIMERFD_GETTIME        SystemCallType = 287
	ACCEPT4                SystemCallType = 288
	SIGNALFD4              SystemCallType = 289
	EVENTFD2               SystemCallType = 290
	EPOLL_CREATE1          SystemCallType = 291
	DUP3                   SystemCallType = 292
	PIPE2                  SystemCallType = 293
	INOTIFY_INIT1          SystemCallType = 294
	PREADV                 SystemCallType = 295
	PWRITEV                SystemCallType = 296
	RT_TGSIGQUEUEINFO      SystemCallType = 297
	PERF_EVENT_OPEN        SystemCallType = 298
	RECVMMSG               SystemCallType = 299
	FANOTIFY_INIT          SystemCallType = 300
	FANOTIFY_MARK          SystemCallType = 301
	PRLIMIT64              SystemCallType = 302
	NAME_TO_HANDLE_AT      SystemCallType = 303
	OPEN_BY_HANDLE_AT      SystemCallType = 304
	CLOCK_ADJTIME          SystemCallType = 305
	SYNCFS                 SystemCallType = 306
	SENDMMSG               SystemCallType = 307
	SETNS                  SystemCallType = 308
	GETCPU                 SystemCallType = 309
	PROCESS_VM_READV       SystemCallType = 310
	PROCESS_VM_WRITEV      SystemCallType = 311
	KCMP                   SystemCallType = 312
	FINIT_MODULE           SystemCallType = 313
)
