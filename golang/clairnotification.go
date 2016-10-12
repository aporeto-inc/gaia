package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/golang/constants"

// ClairNotificationIdentity represents the Identity of the object
var ClairNotificationIdentity = elemental.Identity{
	Name:     "clairnotification",
	Category: "clairnotifications",
}

// ClairNotificationsList represents a list of ClairNotifications
type ClairNotificationsList []*ClairNotification

// ClairNotification represents the model of a clairnotification
type ClairNotification struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"id,primarykey,omitempty" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" cql:"annotation,omitempty" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" cql:"associatedtags,omitempty" bson:"associatedtags"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" cql:"createdat,omitempty" bson:"createdat"`

	// Deleted marks if the entity has been deleted.
	Deleted bool `json:"-" cql:"deleted,omitempty" bson:"deleted"`

	// LayerLimit is the number of layers returned in notification
	Layerlimit int `json:"layerlimit" cql:"layerlimit,omitempty" bson:"layerlimit"`

	// LayersIntroducingNewVulnerability defines layers that are effected by new vulnerability
	LayersIntroducingNewVulnerability []string `json:"layersIntroducingNewVulnerability" cql:"layersintroducingnewvulnerability,omitempty" bson:"layersintroducingnewvulnerability"`

	// LayersIntroducingOldVulnerability defines layers that are effected by old vulnerability
	LayersIntroducingOldVulnerability []string `json:"layersIntroducingOldVulnerability" cql:"layersintroducingoldvulnerability,omitempty" bson:"layersintroducingoldvulnerability"`

	// Name is the name of the notification
	Name string `json:"name" cql:"name,omitempty" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty" bson:"_namespace"`

	// NewVulnerabilityLink is the link that point to the new vulnerability
	NewVulnerabilityLink string `json:"newVulnerabilityLink" cql:"newvulnerabilitylink,omitempty" bson:"newvulnerabilitylink"`

	// NewVulnerabilityName is the name of the new vulnerability
	NewVulnerabilityName string `json:"newVulnerabilityName" cql:"newvulnerabilityname,omitempty" bson:"newvulnerabilityname"`

	// NextPage is the next page number
	NextPage string `json:"nextPage" cql:"nextpage,omitempty" bson:"nextpage"`

	// Notification is the name of the notification sent by Clair using the webhook
	Notification *Notification `json:"notification" cql:"-" bson:"-"`

	// notificationCreatedAt is the time when then notification was created
	NotificationCreatedAt string `json:"notificationCreatedAt" cql:"notificationcreatedat,omitempty" bson:"notificationcreatedat"`

	// NotificationDeletedAt is the time when the notification was deleted
	NotificationDeletedAt string `json:"notificationDeletedAt" cql:"notificationdeletedat,omitempty" bson:"notificationdeletedat"`

	// NotificationNotifiedAt is the time when the notification was sent
	NotificationNotifiedAt string `json:"notificationNotifiedAt" cql:"notificationnotifiedat,omitempty" bson:"notificationnotifiedat"`

	// oldVulnerabilityLink is the link that point to the old vulnerability
	OldVulnerabilityLink string `json:"oldVulnerabilityLink" cql:"oldvulnerabilitylink,omitempty" bson:"oldvulnerabilitylink"`

	// oldVulnerabilityName is the name of the old vulnerability
	OldVulnerabilityName string `json:"oldVulnerabilityName" cql:"oldvulnerabilityname,omitempty" bson:"oldvulnerabilityname"`

	// Page is the page number
	Page string `json:"page" cql:"page,omitempty" bson:"page"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" cql:"parentid,omitempty" bson:"parentid"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" cql:"parenttype,omitempty" bson:"parenttype"`

	// Status of an entity
	Status constants.EntityStatus `json:"status" cql:"status,omitempty" bson:"status"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" cql:"updatedat,omitempty" bson:"updatedat"`
}

// NewClairNotification returns a new *ClairNotification
func NewClairNotification() *ClairNotification {

	return &ClairNotification{
		Status: constants.Active,
	}
}

// Identity returns the Identity of the object.
func (o *ClairNotification) Identity() elemental.Identity {

	return ClairNotificationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ClairNotification) Identifier() string {

	return o.ID
}

func (o *ClairNotification) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ClairNotification) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *ClairNotification) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *ClairNotification) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *ClairNotification) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *ClairNotification) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *ClairNotification) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetNamespace returns the namespace of the receiver
func (o *ClairNotification) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *ClairNotification) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetParentID returns the parentID of the receiver
func (o *ClairNotification) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *ClairNotification) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *ClairNotification) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *ClairNotification) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetStatus returns the status of the receiver
func (o *ClairNotification) GetStatus() constants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *ClairNotification) SetStatus(status constants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *ClairNotification) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *ClairNotification) Validate() error {

	errors := elemental.Errors{}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o ClairNotification) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ClairNotificationAttributesMap[name]
}

// ClairNotificationAttributesMap represents the map of attribute for ClairNotification.
var ClairNotificationAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"CreatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Deleted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Getter:         true,
		Name:           "deleted",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Layerlimit": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "layerlimit",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"LayersIntroducingNewVulnerability": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "layersIntroducingNewVulnerability",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"LayersIntroducingOldVulnerability": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "layersIntroducingOldVulnerability",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"NewVulnerabilityLink": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "newVulnerabilityLink",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"NewVulnerabilityName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "newVulnerabilityName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"NextPage": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "nextPage",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Notification": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "notification",
		Orderable:      true,
		SubType:        "notification",
		Type:           "external",
	},
	"NotificationCreatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "notificationCreatedAt",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"NotificationDeletedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "notificationDeletedAt",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"NotificationNotifiedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "notificationNotifiedAt",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"OldVulnerabilityLink": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "oldVulnerabilityLink",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"OldVulnerabilityName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "oldVulnerabilityName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Page": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "page",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"UpdatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
