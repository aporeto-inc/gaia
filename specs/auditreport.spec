# Model
model:
  rest_name: auditreport
  resource_name: auditreports
  entity_name: AuditReport
  package: zack
  group: policy/audit
  description: Post a new audit statistics report.

# Attributes
attributes:
  v1:
  - name: AUID
    description: Needs documentation.
    type: string
    exposed: true
    example_value: xxx-xxx

  - name: CWD
    description: Command working directory.
    type: string
    exposed: true
    example_value: /etc

  - name: EGID
    description: Needs documentation.
    type: integer
    exposed: true

  - name: EUID
    description: Needs documentation.
    type: integer
    exposed: true

  - name: EXE
    description: Path to the executable.
    type: string
    exposed: true
    example_value: /bin/ls

  - name: FSGID
    description: Needs documentation.
    type: integer
    exposed: true

  - name: FSUID
    description: Needs documentation.
    type: integer
    exposed: true

  - name: GID
    description: Needs documentation.
    type: integer
    exposed: true

  - name: PER
    description: Needs documentation.
    type: integer
    exposed: true

  - name: PID
    description: PID of the executable.
    type: integer
    exposed: true

  - name: PPID
    description: PID of the parent executable.
    type: integer
    exposed: true

  - name: SGID
    description: Needs documentation.
    type: integer
    exposed: true

  - name: SUID
    description: Needs documentation.
    type: integer
    exposed: true

  - name: UID
    description: Needs documentation.
    type: integer
    exposed: true

  - name: a0
    description: Needs documentation.
    type: string
    exposed: true
    example_value: xxx-xxx

  - name: a1
    description: Needs documentation.
    type: string
    exposed: true
    example_value: xxx-xxx

  - name: a2
    description: Needs documentation.
    type: string
    exposed: true
    example_value: xxx-xxx

  - name: a3
    description: Needs documentation.
    type: string
    exposed: true
    example_value: xxx-xxx

  - name: arch
    description: Architecture of the system where the syscall happened.
    type: string
    exposed: true
    example_value: x86_64

  - name: auditProfileID
    description: ID the audit profile that triggered the report.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: auditProfileNamespace
    description: Namespace the audit profile that triggered the report.
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
    description: Exit code of the executable.
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
    description: Type of record.
    type: string
    exposed: true
    required: true
    example_value: Syscall

  - name: sequence
    description: Needs documentation.
    type: integer
    exposed: true

  - name: success
    description: Tells if the operation has been a success of a failure.
    type: boolean
    exposed: true

  - name: syscall
    description: Syscall name.
    type: string
    exposed: true
    example_value: execve

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
