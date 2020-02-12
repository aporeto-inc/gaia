package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ContainerImageIdentity represents the Identity of the object.
var ContainerImageIdentity = elemental.Identity{
	Name:     "containerimage",
	Category: "containerimages",
	Package:  "aki",
	Private:  false,
}

// ContainerImagesList represents a list of ContainerImages
type ContainerImagesList []*ContainerImage

// Identity returns the identity of the objects in the list.
func (o ContainerImagesList) Identity() elemental.Identity {

	return ContainerImageIdentity
}

// Copy returns a pointer to a copy the ContainerImagesList.
func (o ContainerImagesList) Copy() elemental.Identifiables {

	copy := append(ContainerImagesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ContainerImagesList.
func (o ContainerImagesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ContainerImagesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ContainerImage))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ContainerImagesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ContainerImagesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the ContainerImagesList converted to SparseContainerImagesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ContainerImagesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseContainerImagesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseContainerImage)
	}

	return out
}

// Version returns the version of the content.
func (o ContainerImagesList) Version() int {

	return 1
}

// ContainerImage represents the model of a containerimage
type ContainerImage struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Score of the compliance.
	ComplianceRiskScore int `json:"complianceRiskScore" msgpack:"complianceRiskScore" bson:"complianceriskscore" mapstructure:"complianceRiskScore,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Number of critical compliance issue.
	CriticalComplianceIssueCount int `json:"criticalComplianceIssueCount" msgpack:"criticalComplianceIssueCount" bson:"criticalcomplianceissuecount" mapstructure:"criticalComplianceIssueCount,omitempty"`

	// Number of critical vulnerabilities.
	CriticalVulnerabilityCount int `json:"criticalVulnerabilityCount" msgpack:"criticalVulnerabilityCount" bson:"criticalvulnerabilitycount" mapstructure:"criticalVulnerabilityCount,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Full name of the distribution.
	Distro string `json:"distro" msgpack:"distro" bson:"distro" mapstructure:"distro,omitempty"`

	// External Identifier of the container image scan.
	ExternalID string `json:"externalID" msgpack:"externalID" bson:"externalid" mapstructure:"externalID,omitempty"`

	// Number of high compliance issue.
	HighComplianceIssueCount int `json:"highComplianceIssueCount" msgpack:"highComplianceIssueCount" bson:"highcomplianceissuecount" mapstructure:"highComplianceIssueCount,omitempty"`

	// Number of high vulnerabilities.
	HighVulnerabilityCount int `json:"highVulnerabilityCount" msgpack:"highVulnerabilityCount" bson:"highvulnerabilitycount" mapstructure:"highVulnerabilityCount,omitempty"`

	// Number of low compliance issue.
	LowComplianceIssueCount int `json:"lowComplianceIssueCount" msgpack:"lowComplianceIssueCount" bson:"lowcomplianceissuecount" mapstructure:"lowComplianceIssueCount,omitempty"`

	// Number of low vulnerabilities.
	LowVulnerabilityCount int `json:"lowVulnerabilityCount" msgpack:"lowVulnerabilityCount" bson:"lowvulnerabilitycount" mapstructure:"lowVulnerabilityCount,omitempty"`

	// Number of medium compliance issue.
	MediumComplianceIssueCount int `json:"mediumComplianceIssueCount" msgpack:"mediumComplianceIssueCount" bson:"mediumcomplianceissuecount" mapstructure:"mediumComplianceIssueCount,omitempty"`

	// Number of medium vulnerabilities.
	MediumVulnerabilityCount int `json:"mediumVulnerabilityCount" msgpack:"mediumVulnerabilityCount" bson:"mediumvulnerabilitycount" mapstructure:"mediumVulnerabilityCount,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Name of the os distribution.
	OsDistro string `json:"osDistro" msgpack:"osDistro" bson:"osdistro" mapstructure:"osDistro,omitempty"`

	// Name of the release.
	OsDistroRelease string `json:"osDistroRelease" msgpack:"osDistroRelease" bson:"osdistrorelease" mapstructure:"osDistroRelease,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Number of total compliance issue.
	TotalComplianceIssueCount int `json:"totalComplianceIssueCount" msgpack:"totalComplianceIssueCount" bson:"totalcomplianceissuecount" mapstructure:"totalComplianceIssueCount,omitempty"`

	// Number of total vulnerabilities.
	TotalVulnerabilityCount int `json:"totalVulnerabilityCount" msgpack:"totalVulnerabilityCount" bson:"totalvulnerabilitycount" mapstructure:"totalVulnerabilityCount,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// Number of vulnerabilities affecting this image.
	VulnerabilitiesCount int `json:"vulnerabilitiesCount" msgpack:"vulnerabilitiesCount" bson:"vulnerabilitiescount" mapstructure:"vulnerabilitiesCount,omitempty"`

	// Score of the vulnerability.
	VulnerabilityRiskScore int `json:"vulnerabilityRiskScore" msgpack:"vulnerabilityRiskScore" bson:"vulnerabilityriskscore" mapstructure:"vulnerabilityRiskScore,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewContainerImage returns a new *ContainerImage
func NewContainerImage() *ContainerImage {

	return &ContainerImage{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Metadata:       []string{},
		NormalizedTags: []string{},
		MigrationsLog:  map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *ContainerImage) Identity() elemental.Identity {

	return ContainerImageIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ContainerImage) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ContainerImage) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ContainerImage) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesContainerImage{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.ComplianceRiskScore = o.ComplianceRiskScore
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.CriticalComplianceIssueCount = o.CriticalComplianceIssueCount
	s.CriticalVulnerabilityCount = o.CriticalVulnerabilityCount
	s.Description = o.Description
	s.Distro = o.Distro
	s.ExternalID = o.ExternalID
	s.HighComplianceIssueCount = o.HighComplianceIssueCount
	s.HighVulnerabilityCount = o.HighVulnerabilityCount
	s.LowComplianceIssueCount = o.LowComplianceIssueCount
	s.LowVulnerabilityCount = o.LowVulnerabilityCount
	s.MediumComplianceIssueCount = o.MediumComplianceIssueCount
	s.MediumVulnerabilityCount = o.MediumVulnerabilityCount
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.OsDistro = o.OsDistro
	s.OsDistroRelease = o.OsDistroRelease
	s.Protected = o.Protected
	s.TotalComplianceIssueCount = o.TotalComplianceIssueCount
	s.TotalVulnerabilityCount = o.TotalVulnerabilityCount
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.VulnerabilitiesCount = o.VulnerabilitiesCount
	s.VulnerabilityRiskScore = o.VulnerabilityRiskScore
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ContainerImage) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesContainerImage{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.ComplianceRiskScore = s.ComplianceRiskScore
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.CriticalComplianceIssueCount = s.CriticalComplianceIssueCount
	o.CriticalVulnerabilityCount = s.CriticalVulnerabilityCount
	o.Description = s.Description
	o.Distro = s.Distro
	o.ExternalID = s.ExternalID
	o.HighComplianceIssueCount = s.HighComplianceIssueCount
	o.HighVulnerabilityCount = s.HighVulnerabilityCount
	o.LowComplianceIssueCount = s.LowComplianceIssueCount
	o.LowVulnerabilityCount = s.LowVulnerabilityCount
	o.MediumComplianceIssueCount = s.MediumComplianceIssueCount
	o.MediumVulnerabilityCount = s.MediumVulnerabilityCount
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.OsDistro = s.OsDistro
	o.OsDistroRelease = s.OsDistroRelease
	o.Protected = s.Protected
	o.TotalComplianceIssueCount = s.TotalComplianceIssueCount
	o.TotalVulnerabilityCount = s.TotalVulnerabilityCount
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.VulnerabilitiesCount = s.VulnerabilitiesCount
	o.VulnerabilityRiskScore = s.VulnerabilityRiskScore
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ContainerImage) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ContainerImage) BleveType() string {

	return "containerimage"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ContainerImage) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *ContainerImage) Doc() string {

	return `A container image can be affected by vulnerabilities and compliance issues.`
}

func (o *ContainerImage) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *ContainerImage) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *ContainerImage) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *ContainerImage) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *ContainerImage) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *ContainerImage) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *ContainerImage) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *ContainerImage) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *ContainerImage) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *ContainerImage) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *ContainerImage) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *ContainerImage) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *ContainerImage) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *ContainerImage) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *ContainerImage) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *ContainerImage) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *ContainerImage) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *ContainerImage) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *ContainerImage) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *ContainerImage) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *ContainerImage) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *ContainerImage) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *ContainerImage) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *ContainerImage) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *ContainerImage) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *ContainerImage) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *ContainerImage) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *ContainerImage) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *ContainerImage) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *ContainerImage) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *ContainerImage) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ContainerImage) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseContainerImage{
			ID:                           &o.ID,
			Annotations:                  &o.Annotations,
			AssociatedTags:               &o.AssociatedTags,
			ComplianceRiskScore:          &o.ComplianceRiskScore,
			CreateIdempotencyKey:         &o.CreateIdempotencyKey,
			CreateTime:                   &o.CreateTime,
			CriticalComplianceIssueCount: &o.CriticalComplianceIssueCount,
			CriticalVulnerabilityCount:   &o.CriticalVulnerabilityCount,
			Description:                  &o.Description,
			Distro:                       &o.Distro,
			ExternalID:                   &o.ExternalID,
			HighComplianceIssueCount:     &o.HighComplianceIssueCount,
			HighVulnerabilityCount:       &o.HighVulnerabilityCount,
			LowComplianceIssueCount:      &o.LowComplianceIssueCount,
			LowVulnerabilityCount:        &o.LowVulnerabilityCount,
			MediumComplianceIssueCount:   &o.MediumComplianceIssueCount,
			MediumVulnerabilityCount:     &o.MediumVulnerabilityCount,
			Metadata:                     &o.Metadata,
			MigrationsLog:                &o.MigrationsLog,
			Name:                         &o.Name,
			Namespace:                    &o.Namespace,
			NormalizedTags:               &o.NormalizedTags,
			OsDistro:                     &o.OsDistro,
			OsDistroRelease:              &o.OsDistroRelease,
			Protected:                    &o.Protected,
			TotalComplianceIssueCount:    &o.TotalComplianceIssueCount,
			TotalVulnerabilityCount:      &o.TotalVulnerabilityCount,
			UpdateIdempotencyKey:         &o.UpdateIdempotencyKey,
			UpdateTime:                   &o.UpdateTime,
			VulnerabilitiesCount:         &o.VulnerabilitiesCount,
			VulnerabilityRiskScore:       &o.VulnerabilityRiskScore,
			ZHash:                        &o.ZHash,
			Zone:                         &o.Zone,
		}
	}

	sp := &SparseContainerImage{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "complianceRiskScore":
			sp.ComplianceRiskScore = &(o.ComplianceRiskScore)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "criticalComplianceIssueCount":
			sp.CriticalComplianceIssueCount = &(o.CriticalComplianceIssueCount)
		case "criticalVulnerabilityCount":
			sp.CriticalVulnerabilityCount = &(o.CriticalVulnerabilityCount)
		case "description":
			sp.Description = &(o.Description)
		case "distro":
			sp.Distro = &(o.Distro)
		case "externalID":
			sp.ExternalID = &(o.ExternalID)
		case "highComplianceIssueCount":
			sp.HighComplianceIssueCount = &(o.HighComplianceIssueCount)
		case "highVulnerabilityCount":
			sp.HighVulnerabilityCount = &(o.HighVulnerabilityCount)
		case "lowComplianceIssueCount":
			sp.LowComplianceIssueCount = &(o.LowComplianceIssueCount)
		case "lowVulnerabilityCount":
			sp.LowVulnerabilityCount = &(o.LowVulnerabilityCount)
		case "mediumComplianceIssueCount":
			sp.MediumComplianceIssueCount = &(o.MediumComplianceIssueCount)
		case "mediumVulnerabilityCount":
			sp.MediumVulnerabilityCount = &(o.MediumVulnerabilityCount)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "osDistro":
			sp.OsDistro = &(o.OsDistro)
		case "osDistroRelease":
			sp.OsDistroRelease = &(o.OsDistroRelease)
		case "protected":
			sp.Protected = &(o.Protected)
		case "totalComplianceIssueCount":
			sp.TotalComplianceIssueCount = &(o.TotalComplianceIssueCount)
		case "totalVulnerabilityCount":
			sp.TotalVulnerabilityCount = &(o.TotalVulnerabilityCount)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "vulnerabilitiesCount":
			sp.VulnerabilitiesCount = &(o.VulnerabilitiesCount)
		case "vulnerabilityRiskScore":
			sp.VulnerabilityRiskScore = &(o.VulnerabilityRiskScore)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseContainerImage to the object.
func (o *ContainerImage) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseContainerImage)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.ComplianceRiskScore != nil {
		o.ComplianceRiskScore = *so.ComplianceRiskScore
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.CriticalComplianceIssueCount != nil {
		o.CriticalComplianceIssueCount = *so.CriticalComplianceIssueCount
	}
	if so.CriticalVulnerabilityCount != nil {
		o.CriticalVulnerabilityCount = *so.CriticalVulnerabilityCount
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Distro != nil {
		o.Distro = *so.Distro
	}
	if so.ExternalID != nil {
		o.ExternalID = *so.ExternalID
	}
	if so.HighComplianceIssueCount != nil {
		o.HighComplianceIssueCount = *so.HighComplianceIssueCount
	}
	if so.HighVulnerabilityCount != nil {
		o.HighVulnerabilityCount = *so.HighVulnerabilityCount
	}
	if so.LowComplianceIssueCount != nil {
		o.LowComplianceIssueCount = *so.LowComplianceIssueCount
	}
	if so.LowVulnerabilityCount != nil {
		o.LowVulnerabilityCount = *so.LowVulnerabilityCount
	}
	if so.MediumComplianceIssueCount != nil {
		o.MediumComplianceIssueCount = *so.MediumComplianceIssueCount
	}
	if so.MediumVulnerabilityCount != nil {
		o.MediumVulnerabilityCount = *so.MediumVulnerabilityCount
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.OsDistro != nil {
		o.OsDistro = *so.OsDistro
	}
	if so.OsDistroRelease != nil {
		o.OsDistroRelease = *so.OsDistroRelease
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.TotalComplianceIssueCount != nil {
		o.TotalComplianceIssueCount = *so.TotalComplianceIssueCount
	}
	if so.TotalVulnerabilityCount != nil {
		o.TotalVulnerabilityCount = *so.TotalVulnerabilityCount
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.VulnerabilitiesCount != nil {
		o.VulnerabilitiesCount = *so.VulnerabilitiesCount
	}
	if so.VulnerabilityRiskScore != nil {
		o.VulnerabilityRiskScore = *so.VulnerabilityRiskScore
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the ContainerImage.
func (o *ContainerImage) DeepCopy() *ContainerImage {

	if o == nil {
		return nil
	}

	out := &ContainerImage{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ContainerImage.
func (o *ContainerImage) DeepCopyInto(out *ContainerImage) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ContainerImage: %s", err))
	}

	*out = *target.(*ContainerImage)
}

// Validate valides the current information stored into the structure.
func (o *ContainerImage) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
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
func (*ContainerImage) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ContainerImageAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ContainerImageLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ContainerImage) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ContainerImageAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ContainerImage) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "complianceRiskScore":
		return o.ComplianceRiskScore
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "criticalComplianceIssueCount":
		return o.CriticalComplianceIssueCount
	case "criticalVulnerabilityCount":
		return o.CriticalVulnerabilityCount
	case "description":
		return o.Description
	case "distro":
		return o.Distro
	case "externalID":
		return o.ExternalID
	case "highComplianceIssueCount":
		return o.HighComplianceIssueCount
	case "highVulnerabilityCount":
		return o.HighVulnerabilityCount
	case "lowComplianceIssueCount":
		return o.LowComplianceIssueCount
	case "lowVulnerabilityCount":
		return o.LowVulnerabilityCount
	case "mediumComplianceIssueCount":
		return o.MediumComplianceIssueCount
	case "mediumVulnerabilityCount":
		return o.MediumVulnerabilityCount
	case "metadata":
		return o.Metadata
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "osDistro":
		return o.OsDistro
	case "osDistroRelease":
		return o.OsDistroRelease
	case "protected":
		return o.Protected
	case "totalComplianceIssueCount":
		return o.TotalComplianceIssueCount
	case "totalVulnerabilityCount":
		return o.TotalVulnerabilityCount
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "vulnerabilitiesCount":
		return o.VulnerabilitiesCount
	case "vulnerabilityRiskScore":
		return o.VulnerabilityRiskScore
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// ContainerImageAttributesMap represents the map of attribute for ContainerImage.
var ContainerImageAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ComplianceRiskScore": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ComplianceRiskScore",
		Description:    `Score of the compliance.`,
		Exposed:        true,
		Name:           "complianceRiskScore",
		Stored:         true,
		Type:           "integer",
	},
	"CreateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"CriticalComplianceIssueCount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CriticalComplianceIssueCount",
		Description:    `Number of critical compliance issue.`,
		Exposed:        true,
		Name:           "criticalComplianceIssueCount",
		Stored:         true,
		Type:           "integer",
	},
	"CriticalVulnerabilityCount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CriticalVulnerabilityCount",
		Description:    `Number of critical vulnerabilities.`,
		Exposed:        true,
		Name:           "criticalVulnerabilityCount",
		Stored:         true,
		Type:           "integer",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Distro": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Distro",
		Description:    `Full name of the distribution.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "distro",
		Stored:         true,
		Type:           "string",
	},
	"ExternalID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExternalID",
		Description:    `External Identifier of the container image scan.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "externalID",
		Stored:         true,
		Type:           "string",
	},
	"HighComplianceIssueCount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "HighComplianceIssueCount",
		Description:    `Number of high compliance issue.`,
		Exposed:        true,
		Name:           "highComplianceIssueCount",
		Stored:         true,
		Type:           "integer",
	},
	"HighVulnerabilityCount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "HighVulnerabilityCount",
		Description:    `Number of high vulnerabilities.`,
		Exposed:        true,
		Name:           "highVulnerabilityCount",
		Stored:         true,
		Type:           "integer",
	},
	"LowComplianceIssueCount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LowComplianceIssueCount",
		Description:    `Number of low compliance issue.`,
		Exposed:        true,
		Name:           "lowComplianceIssueCount",
		Stored:         true,
		Type:           "integer",
	},
	"LowVulnerabilityCount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LowVulnerabilityCount",
		Description:    `Number of low vulnerabilities.`,
		Exposed:        true,
		Name:           "lowVulnerabilityCount",
		Stored:         true,
		Type:           "integer",
	},
	"MediumComplianceIssueCount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MediumComplianceIssueCount",
		Description:    `Number of medium compliance issue.`,
		Exposed:        true,
		Name:           "mediumComplianceIssueCount",
		Stored:         true,
		Type:           "integer",
	},
	"MediumVulnerabilityCount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MediumVulnerabilityCount",
		Description:    `Number of medium vulnerabilities.`,
		Exposed:        true,
		Name:           "mediumVulnerabilityCount",
		Stored:         true,
		Type:           "integer",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"MigrationsLog": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"OsDistro": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OsDistro",
		Description:    `Name of the os distribution.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "osDistro",
		Stored:         true,
		Type:           "string",
	},
	"OsDistroRelease": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OsDistroRelease",
		Description:    `Name of the release.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "osDistroRelease",
		Stored:         true,
		Type:           "string",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"TotalComplianceIssueCount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TotalComplianceIssueCount",
		Description:    `Number of total compliance issue.`,
		Exposed:        true,
		Name:           "totalComplianceIssueCount",
		Stored:         true,
		Type:           "integer",
	},
	"TotalVulnerabilityCount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TotalVulnerabilityCount",
		Description:    `Number of total vulnerabilities.`,
		Exposed:        true,
		Name:           "totalVulnerabilityCount",
		Stored:         true,
		Type:           "integer",
	},
	"UpdateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"VulnerabilitiesCount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "VulnerabilitiesCount",
		Description:    `Number of vulnerabilities affecting this image.`,
		Exposed:        true,
		Name:           "vulnerabilitiesCount",
		Stored:         true,
		Type:           "integer",
	},
	"VulnerabilityRiskScore": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "VulnerabilityRiskScore",
		Description:    `Score of the vulnerability.`,
		Exposed:        true,
		Name:           "vulnerabilityRiskScore",
		Stored:         true,
		Type:           "integer",
	},
	"ZHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// ContainerImageLowerCaseAttributesMap represents the map of attribute for ContainerImage.
var ContainerImageLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"complianceriskscore": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ComplianceRiskScore",
		Description:    `Score of the compliance.`,
		Exposed:        true,
		Name:           "complianceRiskScore",
		Stored:         true,
		Type:           "integer",
	},
	"createidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"criticalcomplianceissuecount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CriticalComplianceIssueCount",
		Description:    `Number of critical compliance issue.`,
		Exposed:        true,
		Name:           "criticalComplianceIssueCount",
		Stored:         true,
		Type:           "integer",
	},
	"criticalvulnerabilitycount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CriticalVulnerabilityCount",
		Description:    `Number of critical vulnerabilities.`,
		Exposed:        true,
		Name:           "criticalVulnerabilityCount",
		Stored:         true,
		Type:           "integer",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"distro": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Distro",
		Description:    `Full name of the distribution.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "distro",
		Stored:         true,
		Type:           "string",
	},
	"externalid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExternalID",
		Description:    `External Identifier of the container image scan.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "externalID",
		Stored:         true,
		Type:           "string",
	},
	"highcomplianceissuecount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "HighComplianceIssueCount",
		Description:    `Number of high compliance issue.`,
		Exposed:        true,
		Name:           "highComplianceIssueCount",
		Stored:         true,
		Type:           "integer",
	},
	"highvulnerabilitycount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "HighVulnerabilityCount",
		Description:    `Number of high vulnerabilities.`,
		Exposed:        true,
		Name:           "highVulnerabilityCount",
		Stored:         true,
		Type:           "integer",
	},
	"lowcomplianceissuecount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LowComplianceIssueCount",
		Description:    `Number of low compliance issue.`,
		Exposed:        true,
		Name:           "lowComplianceIssueCount",
		Stored:         true,
		Type:           "integer",
	},
	"lowvulnerabilitycount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LowVulnerabilityCount",
		Description:    `Number of low vulnerabilities.`,
		Exposed:        true,
		Name:           "lowVulnerabilityCount",
		Stored:         true,
		Type:           "integer",
	},
	"mediumcomplianceissuecount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MediumComplianceIssueCount",
		Description:    `Number of medium compliance issue.`,
		Exposed:        true,
		Name:           "mediumComplianceIssueCount",
		Stored:         true,
		Type:           "integer",
	},
	"mediumvulnerabilitycount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MediumVulnerabilityCount",
		Description:    `Number of medium vulnerabilities.`,
		Exposed:        true,
		Name:           "mediumVulnerabilityCount",
		Stored:         true,
		Type:           "integer",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"migrationslog": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"osdistro": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OsDistro",
		Description:    `Name of the os distribution.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "osDistro",
		Stored:         true,
		Type:           "string",
	},
	"osdistrorelease": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OsDistroRelease",
		Description:    `Name of the release.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "osDistroRelease",
		Stored:         true,
		Type:           "string",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"totalcomplianceissuecount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TotalComplianceIssueCount",
		Description:    `Number of total compliance issue.`,
		Exposed:        true,
		Name:           "totalComplianceIssueCount",
		Stored:         true,
		Type:           "integer",
	},
	"totalvulnerabilitycount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TotalVulnerabilityCount",
		Description:    `Number of total vulnerabilities.`,
		Exposed:        true,
		Name:           "totalVulnerabilityCount",
		Stored:         true,
		Type:           "integer",
	},
	"updateidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"vulnerabilitiescount": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "VulnerabilitiesCount",
		Description:    `Number of vulnerabilities affecting this image.`,
		Exposed:        true,
		Name:           "vulnerabilitiesCount",
		Stored:         true,
		Type:           "integer",
	},
	"vulnerabilityriskscore": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "VulnerabilityRiskScore",
		Description:    `Score of the vulnerability.`,
		Exposed:        true,
		Name:           "vulnerabilityRiskScore",
		Stored:         true,
		Type:           "integer",
	},
	"zhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseContainerImagesList represents a list of SparseContainerImages
type SparseContainerImagesList []*SparseContainerImage

// Identity returns the identity of the objects in the list.
func (o SparseContainerImagesList) Identity() elemental.Identity {

	return ContainerImageIdentity
}

// Copy returns a pointer to a copy the SparseContainerImagesList.
func (o SparseContainerImagesList) Copy() elemental.Identifiables {

	copy := append(SparseContainerImagesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseContainerImagesList.
func (o SparseContainerImagesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseContainerImagesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseContainerImage))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseContainerImagesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseContainerImagesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseContainerImagesList converted to ContainerImagesList.
func (o SparseContainerImagesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseContainerImagesList) Version() int {

	return 1
}

// SparseContainerImage represents the sparse version of a containerimage.
type SparseContainerImage struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// Score of the compliance.
	ComplianceRiskScore *int `json:"complianceRiskScore,omitempty" msgpack:"complianceRiskScore,omitempty" bson:"complianceriskscore,omitempty" mapstructure:"complianceRiskScore,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Number of critical compliance issue.
	CriticalComplianceIssueCount *int `json:"criticalComplianceIssueCount,omitempty" msgpack:"criticalComplianceIssueCount,omitempty" bson:"criticalcomplianceissuecount,omitempty" mapstructure:"criticalComplianceIssueCount,omitempty"`

	// Number of critical vulnerabilities.
	CriticalVulnerabilityCount *int `json:"criticalVulnerabilityCount,omitempty" msgpack:"criticalVulnerabilityCount,omitempty" bson:"criticalvulnerabilitycount,omitempty" mapstructure:"criticalVulnerabilityCount,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Full name of the distribution.
	Distro *string `json:"distro,omitempty" msgpack:"distro,omitempty" bson:"distro,omitempty" mapstructure:"distro,omitempty"`

	// External Identifier of the container image scan.
	ExternalID *string `json:"externalID,omitempty" msgpack:"externalID,omitempty" bson:"externalid,omitempty" mapstructure:"externalID,omitempty"`

	// Number of high compliance issue.
	HighComplianceIssueCount *int `json:"highComplianceIssueCount,omitempty" msgpack:"highComplianceIssueCount,omitempty" bson:"highcomplianceissuecount,omitempty" mapstructure:"highComplianceIssueCount,omitempty"`

	// Number of high vulnerabilities.
	HighVulnerabilityCount *int `json:"highVulnerabilityCount,omitempty" msgpack:"highVulnerabilityCount,omitempty" bson:"highvulnerabilitycount,omitempty" mapstructure:"highVulnerabilityCount,omitempty"`

	// Number of low compliance issue.
	LowComplianceIssueCount *int `json:"lowComplianceIssueCount,omitempty" msgpack:"lowComplianceIssueCount,omitempty" bson:"lowcomplianceissuecount,omitempty" mapstructure:"lowComplianceIssueCount,omitempty"`

	// Number of low vulnerabilities.
	LowVulnerabilityCount *int `json:"lowVulnerabilityCount,omitempty" msgpack:"lowVulnerabilityCount,omitempty" bson:"lowvulnerabilitycount,omitempty" mapstructure:"lowVulnerabilityCount,omitempty"`

	// Number of medium compliance issue.
	MediumComplianceIssueCount *int `json:"mediumComplianceIssueCount,omitempty" msgpack:"mediumComplianceIssueCount,omitempty" bson:"mediumcomplianceissuecount,omitempty" mapstructure:"mediumComplianceIssueCount,omitempty"`

	// Number of medium vulnerabilities.
	MediumVulnerabilityCount *int `json:"mediumVulnerabilityCount,omitempty" msgpack:"mediumVulnerabilityCount,omitempty" bson:"mediumvulnerabilitycount,omitempty" mapstructure:"mediumVulnerabilityCount,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Name of the os distribution.
	OsDistro *string `json:"osDistro,omitempty" msgpack:"osDistro,omitempty" bson:"osdistro,omitempty" mapstructure:"osDistro,omitempty"`

	// Name of the release.
	OsDistroRelease *string `json:"osDistroRelease,omitempty" msgpack:"osDistroRelease,omitempty" bson:"osdistrorelease,omitempty" mapstructure:"osDistroRelease,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Number of total compliance issue.
	TotalComplianceIssueCount *int `json:"totalComplianceIssueCount,omitempty" msgpack:"totalComplianceIssueCount,omitempty" bson:"totalcomplianceissuecount,omitempty" mapstructure:"totalComplianceIssueCount,omitempty"`

	// Number of total vulnerabilities.
	TotalVulnerabilityCount *int `json:"totalVulnerabilityCount,omitempty" msgpack:"totalVulnerabilityCount,omitempty" bson:"totalvulnerabilitycount,omitempty" mapstructure:"totalVulnerabilityCount,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// Number of vulnerabilities affecting this image.
	VulnerabilitiesCount *int `json:"vulnerabilitiesCount,omitempty" msgpack:"vulnerabilitiesCount,omitempty" bson:"vulnerabilitiescount,omitempty" mapstructure:"vulnerabilitiesCount,omitempty"`

	// Score of the vulnerability.
	VulnerabilityRiskScore *int `json:"vulnerabilityRiskScore,omitempty" msgpack:"vulnerabilityRiskScore,omitempty" bson:"vulnerabilityriskscore,omitempty" mapstructure:"vulnerabilityRiskScore,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseContainerImage returns a new  SparseContainerImage.
func NewSparseContainerImage() *SparseContainerImage {
	return &SparseContainerImage{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseContainerImage) Identity() elemental.Identity {

	return ContainerImageIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseContainerImage) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseContainerImage) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseContainerImage) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseContainerImage{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.ComplianceRiskScore != nil {
		s.ComplianceRiskScore = o.ComplianceRiskScore
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.CriticalComplianceIssueCount != nil {
		s.CriticalComplianceIssueCount = o.CriticalComplianceIssueCount
	}
	if o.CriticalVulnerabilityCount != nil {
		s.CriticalVulnerabilityCount = o.CriticalVulnerabilityCount
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Distro != nil {
		s.Distro = o.Distro
	}
	if o.ExternalID != nil {
		s.ExternalID = o.ExternalID
	}
	if o.HighComplianceIssueCount != nil {
		s.HighComplianceIssueCount = o.HighComplianceIssueCount
	}
	if o.HighVulnerabilityCount != nil {
		s.HighVulnerabilityCount = o.HighVulnerabilityCount
	}
	if o.LowComplianceIssueCount != nil {
		s.LowComplianceIssueCount = o.LowComplianceIssueCount
	}
	if o.LowVulnerabilityCount != nil {
		s.LowVulnerabilityCount = o.LowVulnerabilityCount
	}
	if o.MediumComplianceIssueCount != nil {
		s.MediumComplianceIssueCount = o.MediumComplianceIssueCount
	}
	if o.MediumVulnerabilityCount != nil {
		s.MediumVulnerabilityCount = o.MediumVulnerabilityCount
	}
	if o.Metadata != nil {
		s.Metadata = o.Metadata
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.OsDistro != nil {
		s.OsDistro = o.OsDistro
	}
	if o.OsDistroRelease != nil {
		s.OsDistroRelease = o.OsDistroRelease
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.TotalComplianceIssueCount != nil {
		s.TotalComplianceIssueCount = o.TotalComplianceIssueCount
	}
	if o.TotalVulnerabilityCount != nil {
		s.TotalVulnerabilityCount = o.TotalVulnerabilityCount
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
	}
	if o.VulnerabilitiesCount != nil {
		s.VulnerabilitiesCount = o.VulnerabilitiesCount
	}
	if o.VulnerabilityRiskScore != nil {
		s.VulnerabilityRiskScore = o.VulnerabilityRiskScore
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseContainerImage) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseContainerImage{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.ComplianceRiskScore != nil {
		o.ComplianceRiskScore = s.ComplianceRiskScore
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.CriticalComplianceIssueCount != nil {
		o.CriticalComplianceIssueCount = s.CriticalComplianceIssueCount
	}
	if s.CriticalVulnerabilityCount != nil {
		o.CriticalVulnerabilityCount = s.CriticalVulnerabilityCount
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Distro != nil {
		o.Distro = s.Distro
	}
	if s.ExternalID != nil {
		o.ExternalID = s.ExternalID
	}
	if s.HighComplianceIssueCount != nil {
		o.HighComplianceIssueCount = s.HighComplianceIssueCount
	}
	if s.HighVulnerabilityCount != nil {
		o.HighVulnerabilityCount = s.HighVulnerabilityCount
	}
	if s.LowComplianceIssueCount != nil {
		o.LowComplianceIssueCount = s.LowComplianceIssueCount
	}
	if s.LowVulnerabilityCount != nil {
		o.LowVulnerabilityCount = s.LowVulnerabilityCount
	}
	if s.MediumComplianceIssueCount != nil {
		o.MediumComplianceIssueCount = s.MediumComplianceIssueCount
	}
	if s.MediumVulnerabilityCount != nil {
		o.MediumVulnerabilityCount = s.MediumVulnerabilityCount
	}
	if s.Metadata != nil {
		o.Metadata = s.Metadata
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.OsDistro != nil {
		o.OsDistro = s.OsDistro
	}
	if s.OsDistroRelease != nil {
		o.OsDistroRelease = s.OsDistroRelease
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.TotalComplianceIssueCount != nil {
		o.TotalComplianceIssueCount = s.TotalComplianceIssueCount
	}
	if s.TotalVulnerabilityCount != nil {
		o.TotalVulnerabilityCount = s.TotalVulnerabilityCount
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}
	if s.VulnerabilitiesCount != nil {
		o.VulnerabilitiesCount = s.VulnerabilitiesCount
	}
	if s.VulnerabilityRiskScore != nil {
		o.VulnerabilityRiskScore = s.VulnerabilityRiskScore
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseContainerImage) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseContainerImage) ToPlain() elemental.PlainIdentifiable {

	out := NewContainerImage()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.ComplianceRiskScore != nil {
		out.ComplianceRiskScore = *o.ComplianceRiskScore
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.CriticalComplianceIssueCount != nil {
		out.CriticalComplianceIssueCount = *o.CriticalComplianceIssueCount
	}
	if o.CriticalVulnerabilityCount != nil {
		out.CriticalVulnerabilityCount = *o.CriticalVulnerabilityCount
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Distro != nil {
		out.Distro = *o.Distro
	}
	if o.ExternalID != nil {
		out.ExternalID = *o.ExternalID
	}
	if o.HighComplianceIssueCount != nil {
		out.HighComplianceIssueCount = *o.HighComplianceIssueCount
	}
	if o.HighVulnerabilityCount != nil {
		out.HighVulnerabilityCount = *o.HighVulnerabilityCount
	}
	if o.LowComplianceIssueCount != nil {
		out.LowComplianceIssueCount = *o.LowComplianceIssueCount
	}
	if o.LowVulnerabilityCount != nil {
		out.LowVulnerabilityCount = *o.LowVulnerabilityCount
	}
	if o.MediumComplianceIssueCount != nil {
		out.MediumComplianceIssueCount = *o.MediumComplianceIssueCount
	}
	if o.MediumVulnerabilityCount != nil {
		out.MediumVulnerabilityCount = *o.MediumVulnerabilityCount
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.OsDistro != nil {
		out.OsDistro = *o.OsDistro
	}
	if o.OsDistroRelease != nil {
		out.OsDistroRelease = *o.OsDistroRelease
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.TotalComplianceIssueCount != nil {
		out.TotalComplianceIssueCount = *o.TotalComplianceIssueCount
	}
	if o.TotalVulnerabilityCount != nil {
		out.TotalVulnerabilityCount = *o.TotalVulnerabilityCount
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.VulnerabilitiesCount != nil {
		out.VulnerabilitiesCount = *o.VulnerabilitiesCount
	}
	if o.VulnerabilityRiskScore != nil {
		out.VulnerabilityRiskScore = *o.VulnerabilityRiskScore
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseContainerImage) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseContainerImage) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseContainerImage) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseContainerImage) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseContainerImage) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseContainerImage) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseContainerImage) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseContainerImage) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseContainerImage) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseContainerImage) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseContainerImage) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseContainerImage) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseContainerImage) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseContainerImage) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseContainerImage) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseContainerImage) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseContainerImage) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseContainerImage) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseContainerImage) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseContainerImage) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseContainerImage) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseContainerImage) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseContainerImage) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseContainerImage) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseContainerImage) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseContainerImage) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseContainerImage) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseContainerImage) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseContainerImage) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseContainerImage) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseContainerImage.
func (o *SparseContainerImage) DeepCopy() *SparseContainerImage {

	if o == nil {
		return nil
	}

	out := &SparseContainerImage{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseContainerImage.
func (o *SparseContainerImage) DeepCopyInto(out *SparseContainerImage) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseContainerImage: %s", err))
	}

	*out = *target.(*SparseContainerImage)
}

type mongoAttributesContainerImage struct {
	ID                           bson.ObjectId       `bson:"_id,omitempty"`
	Annotations                  map[string][]string `bson:"annotations"`
	AssociatedTags               []string            `bson:"associatedtags"`
	ComplianceRiskScore          int                 `bson:"complianceriskscore"`
	CreateIdempotencyKey         string              `bson:"createidempotencykey"`
	CreateTime                   time.Time           `bson:"createtime"`
	CriticalComplianceIssueCount int                 `bson:"criticalcomplianceissuecount"`
	CriticalVulnerabilityCount   int                 `bson:"criticalvulnerabilitycount"`
	Description                  string              `bson:"description"`
	Distro                       string              `bson:"distro"`
	ExternalID                   string              `bson:"externalid"`
	HighComplianceIssueCount     int                 `bson:"highcomplianceissuecount"`
	HighVulnerabilityCount       int                 `bson:"highvulnerabilitycount"`
	LowComplianceIssueCount      int                 `bson:"lowcomplianceissuecount"`
	LowVulnerabilityCount        int                 `bson:"lowvulnerabilitycount"`
	MediumComplianceIssueCount   int                 `bson:"mediumcomplianceissuecount"`
	MediumVulnerabilityCount     int                 `bson:"mediumvulnerabilitycount"`
	Metadata                     []string            `bson:"metadata"`
	MigrationsLog                map[string]string   `bson:"migrationslog,omitempty"`
	Name                         string              `bson:"name"`
	Namespace                    string              `bson:"namespace"`
	NormalizedTags               []string            `bson:"normalizedtags"`
	OsDistro                     string              `bson:"osdistro"`
	OsDistroRelease              string              `bson:"osdistrorelease"`
	Protected                    bool                `bson:"protected"`
	TotalComplianceIssueCount    int                 `bson:"totalcomplianceissuecount"`
	TotalVulnerabilityCount      int                 `bson:"totalvulnerabilitycount"`
	UpdateIdempotencyKey         string              `bson:"updateidempotencykey"`
	UpdateTime                   time.Time           `bson:"updatetime"`
	VulnerabilitiesCount         int                 `bson:"vulnerabilitiescount"`
	VulnerabilityRiskScore       int                 `bson:"vulnerabilityriskscore"`
	ZHash                        int                 `bson:"zhash"`
	Zone                         int                 `bson:"zone"`
}
type mongoAttributesSparseContainerImage struct {
	ID                           bson.ObjectId        `bson:"_id,omitempty"`
	Annotations                  *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags               *[]string            `bson:"associatedtags,omitempty"`
	ComplianceRiskScore          *int                 `bson:"complianceriskscore,omitempty"`
	CreateIdempotencyKey         *string              `bson:"createidempotencykey,omitempty"`
	CreateTime                   *time.Time           `bson:"createtime,omitempty"`
	CriticalComplianceIssueCount *int                 `bson:"criticalcomplianceissuecount,omitempty"`
	CriticalVulnerabilityCount   *int                 `bson:"criticalvulnerabilitycount,omitempty"`
	Description                  *string              `bson:"description,omitempty"`
	Distro                       *string              `bson:"distro,omitempty"`
	ExternalID                   *string              `bson:"externalid,omitempty"`
	HighComplianceIssueCount     *int                 `bson:"highcomplianceissuecount,omitempty"`
	HighVulnerabilityCount       *int                 `bson:"highvulnerabilitycount,omitempty"`
	LowComplianceIssueCount      *int                 `bson:"lowcomplianceissuecount,omitempty"`
	LowVulnerabilityCount        *int                 `bson:"lowvulnerabilitycount,omitempty"`
	MediumComplianceIssueCount   *int                 `bson:"mediumcomplianceissuecount,omitempty"`
	MediumVulnerabilityCount     *int                 `bson:"mediumvulnerabilitycount,omitempty"`
	Metadata                     *[]string            `bson:"metadata,omitempty"`
	MigrationsLog                *map[string]string   `bson:"migrationslog,omitempty"`
	Name                         *string              `bson:"name,omitempty"`
	Namespace                    *string              `bson:"namespace,omitempty"`
	NormalizedTags               *[]string            `bson:"normalizedtags,omitempty"`
	OsDistro                     *string              `bson:"osdistro,omitempty"`
	OsDistroRelease              *string              `bson:"osdistrorelease,omitempty"`
	Protected                    *bool                `bson:"protected,omitempty"`
	TotalComplianceIssueCount    *int                 `bson:"totalcomplianceissuecount,omitempty"`
	TotalVulnerabilityCount      *int                 `bson:"totalvulnerabilitycount,omitempty"`
	UpdateIdempotencyKey         *string              `bson:"updateidempotencykey,omitempty"`
	UpdateTime                   *time.Time           `bson:"updatetime,omitempty"`
	VulnerabilitiesCount         *int                 `bson:"vulnerabilitiescount,omitempty"`
	VulnerabilityRiskScore       *int                 `bson:"vulnerabilityriskscore,omitempty"`
	ZHash                        *int                 `bson:"zhash,omitempty"`
	Zone                         *int                 `bson:"zone,omitempty"`
}
