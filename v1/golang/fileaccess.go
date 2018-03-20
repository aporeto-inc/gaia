package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// FileAccessModeValue represents the possible values for attribute "mode".
type FileAccessModeValue string

const (
	// FileAccessModeRead represents the value Read.
	FileAccessModeRead FileAccessModeValue = "Read"

	// FileAccessModeReadWrite represents the value ReadWrite.
	FileAccessModeReadWrite FileAccessModeValue = "ReadWrite"

	// FileAccessModeWrite represents the value Write.
	FileAccessModeWrite FileAccessModeValue = "Write"
)

// FileAccessIdentity represents the Identity of the object.
var FileAccessIdentity = elemental.Identity{
	Name:     "fileaccess",
	Category: "fileaccesses",
	Private:  false,
}

// FileAccessList represents a list of FileAccess
type FileAccessList []*FileAccess

// ContentIdentity returns the identity of the objects in the list.
func (o FileAccessList) ContentIdentity() elemental.Identity {

	return FileAccessIdentity
}

// Copy returns a pointer to a copy the FileAccessList.
func (o FileAccessList) Copy() elemental.ContentIdentifiable {

	copy := append(FileAccessList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the FileAccessList.
func (o FileAccessList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(FileAccessList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FileAccess))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FileAccessList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FileAccessList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o FileAccessList) Version() int {

	return 1
}

// FileAccess represents the model of a fileaccess
type FileAccess struct {
	// Action tells if the access has been allowed or not.
	Action string `json:"action" bson:"-" mapstructure:"action,omitempty"`

	// Count tells how many times the file has been accessed.
	Count int `json:"count" bson:"-" mapstructure:"count,omitempty"`

	// Host is the host that served the accessed file.
	Host string `json:"host" bson:"-" mapstructure:"host,omitempty"`

	// Mode is the mode of the accessed file.
	Mode FileAccessModeValue `json:"mode" bson:"-" mapstructure:"mode,omitempty"`

	// Path is the path of the accessed file.
	Path string `json:"path" bson:"-" mapstructure:"path,omitempty"`

	// Protocol is the protocol used to access the file.
	Protocol string `json:"protocol" bson:"-" mapstructure:"protocol,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewFileAccess returns a new *FileAccess
func NewFileAccess() *FileAccess {

	return &FileAccess{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *FileAccess) Identity() elemental.Identity {

	return FileAccessIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FileAccess) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FileAccess) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *FileAccess) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *FileAccess) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *FileAccess) Doc() string {
	return `Returns file access statistics on a particular processing unit.`
}

func (o *FileAccess) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *FileAccess) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("mode", string(o.Mode), []string{"Read", "ReadWrite", "Write"}, true); err != nil {
		errors = append(errors, err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*FileAccess) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FileAccessAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FileAccessLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FileAccess) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FileAccessAttributesMap
}

// FileAccessAttributesMap represents the map of attribute for FileAccess.
var FileAccessAttributesMap = map[string]elemental.AttributeSpecification{
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Action",
		Description:    `Action tells if the access has been allowed or not.`,
		Exposed:        true,
		Format:         "free",
		Name:           "action",
		ReadOnly:       true,
		Type:           "string",
	},
	"Count": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Count",
		Description:    `Count tells how many times the file has been accessed.`,
		Exposed:        true,
		Name:           "count",
		ReadOnly:       true,
		Type:           "integer",
	},
	"Host": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Host",
		Description:    `Host is the host that served the accessed file.`,
		Exposed:        true,
		Format:         "free",
		Name:           "host",
		ReadOnly:       true,
		Type:           "string",
	},
	"Mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Read", "ReadWrite", "Write"},
		Autogenerated:  true,
		ConvertedName:  "Mode",
		Description:    `Mode is the mode of the accessed file.`,
		Exposed:        true,
		Name:           "mode",
		ReadOnly:       true,
		Type:           "enum",
	},
	"Path": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Path",
		Description:    `Path is the path of the accessed file.`,
		Exposed:        true,
		Format:         "free",
		Name:           "path",
		ReadOnly:       true,
		Type:           "string",
	},
	"Protocol": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Protocol",
		Description:    `Protocol is the protocol used to access the file.`,
		Exposed:        true,
		Format:         "free",
		Name:           "protocol",
		ReadOnly:       true,
		Type:           "string",
	},
}

// FileAccessLowerCaseAttributesMap represents the map of attribute for FileAccess.
var FileAccessLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"action": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Action",
		Description:    `Action tells if the access has been allowed or not.`,
		Exposed:        true,
		Format:         "free",
		Name:           "action",
		ReadOnly:       true,
		Type:           "string",
	},
	"count": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Count",
		Description:    `Count tells how many times the file has been accessed.`,
		Exposed:        true,
		Name:           "count",
		ReadOnly:       true,
		Type:           "integer",
	},
	"host": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Host",
		Description:    `Host is the host that served the accessed file.`,
		Exposed:        true,
		Format:         "free",
		Name:           "host",
		ReadOnly:       true,
		Type:           "string",
	},
	"mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Read", "ReadWrite", "Write"},
		Autogenerated:  true,
		ConvertedName:  "Mode",
		Description:    `Mode is the mode of the accessed file.`,
		Exposed:        true,
		Name:           "mode",
		ReadOnly:       true,
		Type:           "enum",
	},
	"path": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Path",
		Description:    `Path is the path of the accessed file.`,
		Exposed:        true,
		Format:         "free",
		Name:           "path",
		ReadOnly:       true,
		Type:           "string",
	},
	"protocol": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Protocol",
		Description:    `Protocol is the protocol used to access the file.`,
		Exposed:        true,
		Format:         "free",
		Name:           "protocol",
		ReadOnly:       true,
		Type:           "string",
	},
}
