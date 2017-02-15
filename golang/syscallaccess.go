package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// SyscallAccessIdentity represents the Identity of the object
var SyscallAccessIdentity = elemental.Identity{
	Name:     "syscallaccess",
	Category: "syscallaccesses",
}

// SyscallAccessList represents a list of SyscallAccess
type SyscallAccessList []*SyscallAccess

// SyscallAccess represents the model of a syscallaccess
type SyscallAccess struct {
	// PID is the PID of the process that used the system call.
	PID int `json:"PID" cql:"-" bson:"-"`

	// Propagate indicates that the policy must be propagated to the children namespaces.
	Propagate bool `json:"Propagate" cql:"-" bson:"-"`

	// Actions tells if the system call has been allowed.
	Action string `json:"action" cql:"-" bson:"-"`

	// Count tells how many times the syscall has been sent.
	Count int `json:"count" cql:"-" bson:"-"`

	// Name represents the name of the system call.
	Name string `json:"name" cql:"-" bson:"-"`

	// ProcessName is the name of the process that used the system call.
	ProcessName string `json:"processName" cql:"-" bson:"-"`
}

// NewSyscallAccess returns a new *SyscallAccess
func NewSyscallAccess() *SyscallAccess {

	return &SyscallAccess{
		Propagate: false,
	}
}

// Identity returns the Identity of the object.
func (o *SyscallAccess) Identity() elemental.Identity {

	return SyscallAccessIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SyscallAccess) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SyscallAccess) SetIdentifier(ID string) {

}

func (o *SyscallAccess) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *SyscallAccess) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (SyscallAccess) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return SyscallAccessAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (SyscallAccess) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SyscallAccessAttributesMap
}

// SyscallAccessAttributesMap represents the map of attribute for SyscallAccess.
var SyscallAccessAttributesMap = map[string]elemental.AttributeSpecification{
	"PID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `PID is the PID of the process that used the system call.`,
		Exposed:        true,
		Name:           "PID",
		ReadOnly:       true,
		Type:           "integer",
	},
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Propagate indicates that the policy must be propagated to the children namespaces.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "Propagate",
		Orderable:      true,
		Type:           "boolean",
	},
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Actions tells if the system call has been allowed.`,
		Exposed:        true,
		Format:         "free",
		Name:           "action",
		ReadOnly:       true,
		Type:           "string",
	},
	"Count": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Count tells how many times the syscall has been sent.`,
		Exposed:        true,
		Name:           "count",
		ReadOnly:       true,
		Type:           "integer",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Name represents the name of the system call.`,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		ReadOnly:       true,
		Type:           "string",
	},
	"ProcessName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ProcessName is the name of the process that used the system call.`,
		Exposed:        true,
		Format:         "free",
		Name:           "processName",
		ReadOnly:       true,
		Type:           "string",
	},
}
