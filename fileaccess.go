package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
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
	Package:  "jenova",
	Private:  false,
}

// FileAccessList represents a list of FileAccess
type FileAccessList []*FileAccess

// Identity returns the identity of the objects in the list.
func (o FileAccessList) Identity() elemental.Identity {

	return FileAccessIdentity
}

// Copy returns a pointer to a copy the FileAccessList.
func (o FileAccessList) Copy() elemental.Identifiables {

	copy := append(FileAccessList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the FileAccessList.
func (o FileAccessList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(FileAccessList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FileAccess))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FileAccessList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FileAccessList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the FileAccessList converted to SparseFileAccessList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o FileAccessList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
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

	sync.Mutex `json:"-" bson:"-"`
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

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *FileAccess) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseFileAccess{
			Action:   &o.Action,
			Count:    &o.Count,
			Host:     &o.Host,
			Mode:     &o.Mode,
			Path:     &o.Path,
			Protocol: &o.Protocol,
		}
	}

	sp := &SparseFileAccess{}
	for _, f := range fields {
		switch f {
		case "action":
			sp.Action = &(o.Action)
		case "count":
			sp.Count = &(o.Count)
		case "host":
			sp.Host = &(o.Host)
		case "mode":
			sp.Mode = &(o.Mode)
		case "path":
			sp.Path = &(o.Path)
		case "protocol":
			sp.Protocol = &(o.Protocol)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseFileAccess to the object.
func (o *FileAccess) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseFileAccess)
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.Count != nil {
		o.Count = *so.Count
	}
	if so.Host != nil {
		o.Host = *so.Host
	}
	if so.Mode != nil {
		o.Mode = *so.Mode
	}
	if so.Path != nil {
		o.Path = *so.Path
	}
	if so.Protocol != nil {
		o.Protocol = *so.Protocol
	}
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
		Name:           "protocol",
		ReadOnly:       true,
		Type:           "string",
	},
}

// SparseFileAccessList represents a list of SparseFileAccess
type SparseFileAccessList []*SparseFileAccess

// Identity returns the identity of the objects in the list.
func (o SparseFileAccessList) Identity() elemental.Identity {

	return FileAccessIdentity
}

// Copy returns a pointer to a copy the SparseFileAccessList.
func (o SparseFileAccessList) Copy() elemental.Identifiables {

	copy := append(SparseFileAccessList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseFileAccessList.
func (o SparseFileAccessList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseFileAccessList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseFileAccess))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseFileAccessList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseFileAccessList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseFileAccessList converted to FileAccessList.
func (o SparseFileAccessList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseFileAccessList) Version() int {

	return 1
}

// SparseFileAccess represents the sparse version of a fileaccess.
type SparseFileAccess struct {
	// Action tells if the access has been allowed or not.
	Action *string `json:"action,omitempty" bson:"-" mapstructure:"action,omitempty"`

	// Count tells how many times the file has been accessed.
	Count *int `json:"count,omitempty" bson:"-" mapstructure:"count,omitempty"`

	// Host is the host that served the accessed file.
	Host *string `json:"host,omitempty" bson:"-" mapstructure:"host,omitempty"`

	// Mode is the mode of the accessed file.
	Mode *FileAccessModeValue `json:"mode,omitempty" bson:"-" mapstructure:"mode,omitempty"`

	// Path is the path of the accessed file.
	Path *string `json:"path,omitempty" bson:"-" mapstructure:"path,omitempty"`

	// Protocol is the protocol used to access the file.
	Protocol *string `json:"protocol,omitempty" bson:"-" mapstructure:"protocol,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseFileAccess returns a new  SparseFileAccess.
func NewSparseFileAccess() *SparseFileAccess {
	return &SparseFileAccess{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseFileAccess) Identity() elemental.Identity {

	return FileAccessIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseFileAccess) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseFileAccess) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseFileAccess) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseFileAccess) ToPlain() elemental.PlainIdentifiable {

	out := NewFileAccess()
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Host != nil {
		out.Host = *o.Host
	}
	if o.Mode != nil {
		out.Mode = *o.Mode
	}
	if o.Path != nil {
		out.Path = *o.Path
	}
	if o.Protocol != nil {
		out.Protocol = *o.Protocol
	}

	return out
}
