# Model
model:
  rest_name: auditreport
  resource_name: auditreports
  entity_name: AuditReport
  package: zack
  group: policy/audit
  description: Post a new audit report.

# Attributes
attributes:
  v1:
  - name: AUID
    description: The login ID of the user who started the audited process.
    type: string
    exposed: true
    example_value: xxx-xxx

  - name: CWD
    description: Command working directory.
    type: string
    exposed: true
    example_value: /etc

  - name: EGID
    description: Effective group ID of the user who started the audited process.
    type: integer
    exposed: true

  - name: EUID
    description: Effective user ID of the user who started the audited process.
    type: integer
    exposed: true

  - name: EXE
    description: Path to the executable.
    type: string
    exposed: true
    example_value: /bin/ls

  - name: FSGID
    description: File system group ID of the user who started the audited process.
    type: integer
    exposed: true

  - name: FSUID
    description: File system user ID of the user who started the audited process.
    type: integer
    exposed: true

  - name: FilePath
    description: Full path of the file that was passed to the system call.
    type: string
    exposed: true

  - name: GID
    description: Group ID of the user who started the analyzed process.
    type: integer
    exposed: true

  - name: PER
    description: File or directory permissions.
    type: integer
    exposed: true

  - name: PID
    description: Process ID of the executable.
    type: integer
    exposed: true

  - name: PPID
    description: Process ID of the parent executable.
    type: integer
    exposed: true

  - name: SGID
    description: Set group ID of the user who started the audited process.
    type: integer
    exposed: true

  - name: SUID
    description: Set user ID of the user who started the audited process.
    type: integer
    exposed: true

  - name: UID
    description: User ID.
    type: integer
    exposed: true

  - name: a0
    description: First argument of the executed system call.
    type: string
    exposed: true
    example_value: xxx-xxx

  - name: a1
    description: Second argument of the executed system call.
    type: string
    exposed: true
    example_value: xxx-xxx

  - name: a2
    description: Third argument of the executed system call.
    type: string
    exposed: true
    example_value: xxx-xxx

  - name: a3
    description: Fourth argument of the executed system call.
    type: string
    exposed: true
    example_value: xxx-xxx

  - name: arch
    description: Architecture of the system of the monitored process.
    type: string
    exposed: true
    example_value: x86_64

  - name: arguments
    description: Arguments passed to the command.
    type: list
    exposed: true
    subtype: string

  - name: auditProfileID
    description: ID of the audit profile that triggered the report.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: auditProfileNamespace
    description: Namespace of the audit profile that triggered the report.
    type: string
    exposed: true
    required: true
    example_value: /my/ns

  - name: command
    description: Command issued.
    type: string
    exposed: true
    example_value: ls

  - name: enforcerID
    description: ID of the enforcer reporting.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: enforcerNamespace
    description: Namespace of the enforcer reporting.
    type: string
    exposed: true
    required: true
    example_value: /my/ns

  - name: exit
    description: Exit code of the executed system call.
    type: integer
    exposed: true

  - name: processingUnitID
    description: ID of the processing unit originating the report.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: processingUnitNamespace
    description: Namespace of the processing unit originating the report.
    type: string
    exposed: true
    required: true
    example_value: /my/ns

  - name: recordType
    description: Type of audit record.
    type: string
    exposed: true
    required: true
    example_value: Syscall

  - name: sequence
    description: Needs documentation.
    type: integer
    exposed: true

  - name: success
    description: Tells if the operation has been a success or a failure.
    type: boolean
    exposed: true

  - name: syscall
    description: System call executed.
    type: string
    exposed: true
    example_value: execve

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
