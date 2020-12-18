# Model
model:
  rest_name: auditreport
  resource_name: auditreports
  entity_name: AuditReport
  package: zack
  group: policy/audit
  description: Post a new audit report.
  extends:
  - '@identifiable-stored'
  - '@zoned-monotonic'
  - '@migratable'

# Ordering
default_order:
- :no-inherit
- timestamp

# Indexes
indexes:
- - namespace
  - timestamp

# Attributes
attributes:
  v1:
  - name: AUID
    description: The login ID of the user who started the audited process.
    type: string
    exposed: true
    stored: true
    example_value: xxx-xxx
    omit_empty: true
    extensions:
      bson_name: a

  - name: CWD
    description: Command working directory.
    type: string
    exposed: true
    stored: true
    example_value: /etc
    omit_empty: true
    extensions:
      bson_name: b

  - name: EGID
    description: Effective group ID of the user who started the audited process.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: c

  - name: EUID
    description: Effective user ID of the user who started the audited process.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: d

  - name: EXE
    description: Path to the executable.
    type: string
    exposed: true
    stored: true
    example_value: /bin/ls
    omit_empty: true
    extensions:
      bson_name: e

  - name: FSGID
    description: File system group ID of the user who started the audited process.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: f

  - name: FSUID
    description: File system user ID of the user who started the audited process.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: g

  - name: FilePath
    description: Full path of the file that was passed to the system call.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: h

  - name: GID
    description: Group ID of the user who started the analyzed process.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: i

  - name: PER
    description: File or directory permissions.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: j

  - name: PID
    description: Process ID of the executable.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: k

  - name: PPID
    description: Process ID of the parent executable.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: l

  - name: SGID
    description: Set group ID of the user who started the audited process.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: m

  - name: SUID
    description: Set user ID of the user who started the audited process.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: "n"

  - name: UID
    description: User ID.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: o

  - name: a0
    description: First argument of the executed system call.
    type: string
    exposed: true
    stored: true
    example_value: xxx-xxx
    omit_empty: true
    extensions:
      bson_name: p

  - name: a1
    description: Second argument of the executed system call.
    type: string
    exposed: true
    stored: true
    example_value: xxx-xxx
    omit_empty: true
    extensions:
      bson_name: q

  - name: a2
    description: Third argument of the executed system call.
    type: string
    exposed: true
    stored: true
    example_value: xxx-xxx
    omit_empty: true
    extensions:
      bson_name: r

  - name: a3
    description: Fourth argument of the executed system call.
    type: string
    exposed: true
    stored: true
    example_value: xxx-xxx
    omit_empty: true
    extensions:
      bson_name: s

  - name: arch
    description: Architecture of the system of the monitored process.
    type: string
    exposed: true
    stored: true
    example_value: x86_64
    omit_empty: true
    extensions:
      bson_name: t

  - name: arguments
    description: Arguments passed to the command.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true
    extensions:
      bson_name: u

  - name: auditProfileID
    description: ID of the audit profile that triggered the report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: v

  - name: auditProfileNamespace
    description: Namespace of the audit profile that triggered the report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/ns
    omit_empty: true
    extensions:
      bson_name: w

  - name: command
    description: Command issued.
    type: string
    exposed: true
    stored: true
    example_value: ls
    omit_empty: true
    extensions:
      bson_name: x

  - name: enforcerID
    description: ID of the enforcer reporting.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: "y"

  - name: enforcerNamespace
    description: Namespace of the enforcer reporting.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/ns
    omit_empty: true
    extensions:
      bson_name: z

  - name: exit
    description: Exit code of the executed system call.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: aa

  - name: processingUnitID
    description: ID of the processing unit originating the report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: ab

  - name: processingUnitNamespace
    description: Namespace of the processing unit originating the report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/ns
    omit_empty: true
    extensions:
      bson_name: ac

  - name: recordType
    description: Type of audit record.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: Syscall
    omit_empty: true
    extensions:
      bson_name: ad

  - name: sequence
    description: Needs documentation.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ae

  - name: success
    description: Tells if the operation has been a success or a failure.
    type: boolean
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: af

  - name: syscall
    description: System call executed.
    type: string
    exposed: true
    stored: true
    example_value: execve
    omit_empty: true
    extensions:
      bson_name: ag

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
    stored: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
    orderable: true
    omit_empty: true
    extensions:
      bson_name: ah
