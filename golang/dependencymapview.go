package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/golang/constants"

// DependencyMapViewTypeValue represents the possible values for attribute "type".
type DependencyMapViewTypeValue string

const (
	// DependencyMapViewTypeAutomatic represents the value Automatic.
	DependencyMapViewTypeAutomatic DependencyMapViewTypeValue = "Automatic"

	// DependencyMapViewTypeManual represents the value Manual.
	DependencyMapViewTypeManual DependencyMapViewTypeValue = "Manual"
)

// DependencyMapViewIdentity represents the Identity of the object
var DependencyMapViewIdentity = elemental.Identity{
	Name:     "dependencymapview",
	Category: "dependencymapviews",
}

// DependencyMapViewsList represents a list of DependencyMapViews
type DependencyMapViewsList []*DependencyMapView

// DependencyMapView represents the model of a dependencymapview
type DependencyMapView struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"id,primarykey,omitempty" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" cql:"annotation,omitempty" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" cql:"associatedtags,omitempty" bson:"associatedtags"`

	// Boolean to know if the dependency map view was computed by the system or not
	Computed bool `json:"computed" cql:"computed,omitempty" bson:"computed"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" cql:"createdat,omitempty" bson:"createdat"`

	// Description is the description of the object.
	Description string `json:"description" cql:"description,omitempty" bson:"description"`

	// Name is the name of the entity
	Name string `json:"name" cql:"name,omitempty" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" cql:"normalizedtags,omitempty" bson:"normalizedtags"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" cql:"parentid,omitempty" bson:"parentid"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" cql:"parenttype,omitempty" bson:"parenttype"`

	// A map of the tags to apply to processing units
	ProcessingUnitTags map[string][]string `json:"processingUnitTags" cql:"processingunittags,omitempty" bson:"processingunittags"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" cql:"protected,omitempty" bson:"protected"`

	// Status of an entity
	Status constants.EntityStatus `json:"status" cql:"status,omitempty" bson:"status"`

	// Values used by the dependency map group
	Subviews DependencyMapSubviewsList `json:"subviews" cql:"subviews,omitempty" bson:"subviews"`

	// Type represents the type of the dependency map. It could be manual or automatic
	Type DependencyMapViewTypeValue `json:"type" cql:"type,omitempty" bson:"type"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" cql:"updatedat,omitempty" bson:"updatedat"`
}

// NewDependencyMapView returns a new *DependencyMapView
func NewDependencyMapView() *DependencyMapView {

	return &DependencyMapView{
		AssociatedTags: []string{},
		NormalizedTags: []string{},
		Status:         constants.Active,
		Subviews:       DependencyMapSubviewsList{},
		Type:           "Manual",
	}
}

// Identity returns the Identity of the object.
func (o *DependencyMapView) Identity() elemental.Identity {

	return DependencyMapViewIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DependencyMapView) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DependencyMapView) SetIdentifier(ID string) {

	o.ID = ID
}

func (o *DependencyMapView) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *DependencyMapView) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *DependencyMapView) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *DependencyMapView) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetName returns the name of the receiver
func (o *DependencyMapView) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *DependencyMapView) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *DependencyMapView) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *DependencyMapView) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *DependencyMapView) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *DependencyMapView) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetParentID returns the parentID of the receiver
func (o *DependencyMapView) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *DependencyMapView) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *DependencyMapView) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *DependencyMapView) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetProtected returns the protected of the receiver
func (o *DependencyMapView) GetProtected() bool {
	return o.Protected
}

// GetStatus returns the status of the receiver
func (o *DependencyMapView) GetStatus() constants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *DependencyMapView) SetStatus(status constants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *DependencyMapView) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *DependencyMapView) Validate() error {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("subviews", o.Subviews); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Automatic", "Manual"}, false); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (DependencyMapView) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return DependencyMapViewAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (DependencyMapView) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DependencyMapViewAttributesMap
}

// DependencyMapViewAttributesMap represents the map of attribute for DependencyMapView.
var DependencyMapViewAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Annotation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AssociatedTags are the list of tags attached to an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"Computed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Boolean to know if the dependency map view was computed by the system or not`,
		Exposed:        true,
		Filterable:     true,
		Name:           "computed",
		Orderable:      true,
		Stored:         true,
		Transient:      true,
		Type:           "boolean",
	},
	"CreatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CreatedAt is the time at which an entity was created`,
		Exposed:        true,
		Name:           "createdAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Index:          true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `NormalizedTags contains the list of normalized tags of the entities`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentID is the ID of the parent, if any,`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentType": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ProcessingUnitTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `A map of the tags to apply to processing units`,
		Exposed:        true,
		Name:           "processingUnitTags",
		Orderable:      true,
		Stored:         true,
		SubType:        "processingunit_transient_tags_map",
		Type:           "external",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Status of an entity`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "status",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		SubType:        "status_enum",
		Type:           "external",
	},
	"Subviews": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Values used by the dependency map group`,
		Exposed:        true,
		Name:           "subviews",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		SubType:        "dependencymapsubviews_entities",
		Type:           "external",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Automatic", "Manual"},
		Description:    `Type represents the type of the dependency map. It could be manual or automatic`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"UpdatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdatedAt is the time at which an entity was updated.`,
		Exposed:        true,
		Name:           "updatedAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
