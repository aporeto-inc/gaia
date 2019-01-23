package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EnforcerProfileKubernetesMetadataExtractorValue represents the possible values for attribute "kubernetesMetadataExtractor".
type EnforcerProfileKubernetesMetadataExtractorValue string

const (
	// EnforcerProfileKubernetesMetadataExtractorKubeSquall represents the value KubeSquall.
	EnforcerProfileKubernetesMetadataExtractorKubeSquall EnforcerProfileKubernetesMetadataExtractorValue = "KubeSquall"

	// EnforcerProfileKubernetesMetadataExtractorPodAtomic represents the value PodAtomic.
	EnforcerProfileKubernetesMetadataExtractorPodAtomic EnforcerProfileKubernetesMetadataExtractorValue = "PodAtomic"

	// EnforcerProfileKubernetesMetadataExtractorPodContainers represents the value PodContainers.
	EnforcerProfileKubernetesMetadataExtractorPodContainers EnforcerProfileKubernetesMetadataExtractorValue = "PodContainers"
)

// EnforcerProfileMetadataExtractorValue represents the possible values for attribute "metadataExtractor".
type EnforcerProfileMetadataExtractorValue string

const (
	// EnforcerProfileMetadataExtractorDocker represents the value Docker.
	EnforcerProfileMetadataExtractorDocker EnforcerProfileMetadataExtractorValue = "Docker"

	// EnforcerProfileMetadataExtractorECS represents the value ECS.
	EnforcerProfileMetadataExtractorECS EnforcerProfileMetadataExtractorValue = "ECS"

	// EnforcerProfileMetadataExtractorKubernetes represents the value Kubernetes.
	EnforcerProfileMetadataExtractorKubernetes EnforcerProfileMetadataExtractorValue = "Kubernetes"
)

// EnforcerProfileIdentity represents the Identity of the object.
var EnforcerProfileIdentity = elemental.Identity{
	Name:     "enforcerprofile",
	Category: "enforcerprofiles",
	Package:  "squall",
	Private:  false,
}

// EnforcerProfilesList represents a list of EnforcerProfiles
type EnforcerProfilesList []*EnforcerProfile

// Identity returns the identity of the objects in the list.
func (o EnforcerProfilesList) Identity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Copy returns a pointer to a copy the EnforcerProfilesList.
func (o EnforcerProfilesList) Copy() elemental.Identifiables {

	copy := append(EnforcerProfilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcerProfilesList.
func (o EnforcerProfilesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EnforcerProfilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EnforcerProfile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerProfilesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerProfilesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the EnforcerProfilesList converted to SparseEnforcerProfilesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EnforcerProfilesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o EnforcerProfilesList) Version() int {

	return 1
}

// EnforcerProfile represents the model of a enforcerprofile
type EnforcerProfile struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// IptablesMarkValue is the mark value to be used in an iptables implementation.
	IPTablesMarkValue int `json:"IPTablesMarkValue" bson:"iptablesmarkvalue" mapstructure:"IPTablesMarkValue,omitempty"`

	// PUBookkeepingInterval configures how often the PU will be synchronized.
	PUBookkeepingInterval string `json:"PUBookkeepingInterval" bson:"pubookkeepinginterval" mapstructure:"PUBookkeepingInterval,omitempty"`

	// PUHeartbeatInterval configures the heart beat interval.
	PUHeartbeatInterval string `json:"PUHeartbeatInterval" bson:"puheartbeatinterval" mapstructure:"PUHeartbeatInterval,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Port used by aporeto application proxy.
	ApplicationProxyPort int `json:"applicationProxyPort" bson:"applicationproxyport" mapstructure:"applicationProxyPort,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// AuditProfileSelectors is the list of tags (key/value pairs) that define the
	// audit policies that must be implemented by this enforcer. The enforcer will
	// implement all policies that match any of these tags.
	AuditProfileSelectors []string `json:"auditProfileSelectors" bson:"auditprofileselectors" mapstructure:"auditProfileSelectors,omitempty"`

	// AuditProfiles returns the audit rules associated with the enforcer profile. This
	// is a read only attribute when an enforcer profile is resolved for an enforcer.
	AuditProfiles AuditProfilesList `json:"auditProfiles" bson:"-" mapstructure:"auditProfiles,omitempty"`

	// AuditSocketBufferSize is the size of the audit socket buffer. Default 16384.
	AuditSocketBufferSize int `json:"auditSocketBufferSize" bson:"auditsocketbuffersize" mapstructure:"auditSocketBufferSize,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// DockerSocketAddress is the address of the docker daemon.
	DockerSocketAddress string `json:"dockerSocketAddress" bson:"dockersocketaddress" mapstructure:"dockerSocketAddress,omitempty"`

	// ExcludedInterfaces is a list of interfaces that must be excluded.
	ExcludedInterfaces []string `json:"excludedInterfaces" bson:"excludedinterfaces" mapstructure:"excludedInterfaces,omitempty"`

	// ExcludedNetworks is the list of networks that must be excluded for this
	// enforcer.
	ExcludedNetworks []string `json:"excludedNetworks" bson:"excludednetworks" mapstructure:"excludedNetworks,omitempty"`

	// hostModeEnabled enables protection of the complete host. When this option is
	// turned on, all incoming and outgoing flows will be monitored. Flows will
	// be allowed if and only if a network policy has been created to allow the flow.
	HostModeEnabled bool `json:"hostModeEnabled" bson:"hostmodeenabled" mapstructure:"hostModeEnabled,omitempty"`

	// HostServices is a list of services that must be activated by default to all
	// enforcers matching this profile.
	HostServices []*HostService `json:"hostServices" bson:"hostservices" mapstructure:"hostServices,omitempty"`

	// IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
	// docker container started with labels matching the rule.
	IgnoreExpression [][]string `json:"ignoreExpression" bson:"ignoreexpression" mapstructure:"ignoreExpression,omitempty"`

	// KillContainersOnFailure will configure the enforcers to kill any containers if
	// there are policy failures.
	KillContainersOnFailure bool `json:"killContainersOnFailure" bson:"killcontainersonfailure" mapstructure:"killContainersOnFailure,omitempty"`

	// Select which metadata extractor to use to process new processing units from
	// Kubernetes.
	KubernetesMetadataExtractor EnforcerProfileKubernetesMetadataExtractorValue `json:"kubernetesMetadataExtractor" bson:"kubernetesmetadataextractor" mapstructure:"kubernetesMetadataExtractor,omitempty"`

	// KubernetesSupportEnabled enables kubernetes mode for the enforcer.
	KubernetesSupportEnabled bool `json:"kubernetesSupportEnabled" bson:"kubernetessupportenabled" mapstructure:"kubernetesSupportEnabled,omitempty"`

	// LinuxProcessesSupportEnabled configures support for Linux processes.
	LinuxProcessesSupportEnabled bool `json:"linuxProcessesSupportEnabled" bson:"linuxprocessessupportenabled" mapstructure:"linuxProcessesSupportEnabled,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Select which metadata extractor to use to process new processing units.
	MetadataExtractor EnforcerProfileMetadataExtractorValue `json:"metadataExtractor" bson:"metadataextractor" mapstructure:"metadataExtractor,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// PolicySynchronizationInterval configures how often the policy will be
	// resynchronized.
	PolicySynchronizationInterval string `json:"policySynchronizationInterval" bson:"policysynchronizationinterval" mapstructure:"policySynchronizationInterval,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// ProxyListenAddress is the address the enforcer should use to listen for API
	// calls. It can be a port (example :9443) or socket path
	// example:
	//   unix:///var/run/aporeto.sock.
	ProxyListenAddress string `json:"proxyListenAddress" bson:"proxylistenaddress" mapstructure:"proxyListenAddress,omitempty"`

	// ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network
	// receiver starting at the ReceiverQueue.
	ReceiverNumberOfQueues int `json:"receiverNumberOfQueues" bson:"receivernumberofqueues" mapstructure:"receiverNumberOfQueues,omitempty"`

	// ReceiverQueue is the base queue number for traffic from the network.
	ReceiverQueue int `json:"receiverQueue" bson:"receiverqueue" mapstructure:"receiverQueue,omitempty"`

	// ReceiverQueueSize is the queue size of the receiver.
	ReceiverQueueSize int `json:"receiverQueueSize" bson:"receiverqueuesize" mapstructure:"receiverQueueSize,omitempty"`

	// RemoteEnforcerEnabled inidicates whether a single enforcer should be used or a
	// distributed enforcer. True means distributed.
	RemoteEnforcerEnabled bool `json:"remoteEnforcerEnabled" bson:"remoteenforcerenabled" mapstructure:"remoteEnforcerEnabled,omitempty"`

	// TargetNetworks is the list of networks that authorization should be applied.
	TargetNetworks []string `json:"targetNetworks" bson:"targetnetworks" mapstructure:"targetNetworks,omitempty"`

	// TargetUDPNetworks is the list of UDP networks that authorization should be
	// applied.
	TargetUDPNetworks []string `json:"targetUDPNetworks" bson:"targetudpnetworks" mapstructure:"targetUDPNetworks,omitempty"`

	// TransmitterNumberOfQueues is the number of queues for application traffic.
	TransmitterNumberOfQueues int `json:"transmitterNumberOfQueues" bson:"transmitternumberofqueues" mapstructure:"transmitterNumberOfQueues,omitempty"`

	// TransmitterQueue is the queue number for traffic from the applications starting
	// at the transmitterQueue.
	TransmitterQueue int `json:"transmitterQueue" bson:"transmitterqueue" mapstructure:"transmitterQueue,omitempty"`

	// TransmitterQueueSize is the size of the queue for application traffic.
	TransmitterQueueSize int `json:"transmitterQueueSize" bson:"transmitterqueuesize" mapstructure:"transmitterQueueSize,omitempty"`

	// List of trusted CA. If empty the main chain of trust will be used.
	TrustedCAs []string `json:"trustedCAs" bson:"trustedcas" mapstructure:"trustedCAs,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewEnforcerProfile returns a new *EnforcerProfile
func NewEnforcerProfile() *EnforcerProfile {

	return &EnforcerProfile{
		ModelVersion:                  1,
		ApplicationProxyPort:          20992,
		AuditProfiles:                 AuditProfilesList{},
		AuditProfileSelectors:         []string{},
		Annotations:                   map[string][]string{},
		AssociatedTags:                []string{},
		DockerSocketAddress:           "unix:///var/run/docker.sock",
		AuditSocketBufferSize:         16384,
		ExcludedNetworks:              []string{},
		HostModeEnabled:               false,
		HostServices:                  []*HostService{},
		ExcludedInterfaces:            []string{},
		LinuxProcessesSupportEnabled:  true,
		ProxyListenAddress:            "unix:///var/run/aporeto.sock",
		PUBookkeepingInterval:         "15m",
		ReceiverNumberOfQueues:        4,
		IgnoreExpression:              [][]string{},
		KubernetesMetadataExtractor:   EnforcerProfileKubernetesMetadataExtractorPodAtomic,
		RemoteEnforcerEnabled:         true,
		MetadataExtractor:             EnforcerProfileMetadataExtractorDocker,
		ReceiverQueueSize:             500,
		TargetUDPNetworks:             []string{},
		IPTablesMarkValue:             1000,
		TransmitterQueue:              4,
		TargetNetworks:                []string{},
		NormalizedTags:                []string{},
		TransmitterNumberOfQueues:     4,
		TransmitterQueueSize:          500,
		Metadata:                      []string{},
		PUHeartbeatInterval:           "5s",
		PolicySynchronizationInterval: "10m",
		KubernetesSupportEnabled:      false,
		TrustedCAs:                    []string{},
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerProfile) Identity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerProfile) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *EnforcerProfile) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerProfile) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *EnforcerProfile) Doc() string {
	return `Allows to create reusable configuration profile for your enforcers. Enforcer
Profiles contains various startup information that can (for some) be updated
live. Enforcer Profiles are assigned to some Enforcer using a Enforcer Profile
Mapping Policy.`
}

func (o *EnforcerProfile) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *EnforcerProfile) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *EnforcerProfile) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *EnforcerProfile) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *EnforcerProfile) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *EnforcerProfile) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *EnforcerProfile) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *EnforcerProfile) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *EnforcerProfile) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *EnforcerProfile) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *EnforcerProfile) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *EnforcerProfile) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *EnforcerProfile) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *EnforcerProfile) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *EnforcerProfile) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *EnforcerProfile) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *EnforcerProfile) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *EnforcerProfile) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *EnforcerProfile) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *EnforcerProfile) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *EnforcerProfile) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *EnforcerProfile) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *EnforcerProfile) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *EnforcerProfile) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *EnforcerProfile) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEnforcerProfile{
			ID:                            &o.ID,
			IPTablesMarkValue:             &o.IPTablesMarkValue,
			PUBookkeepingInterval:         &o.PUBookkeepingInterval,
			PUHeartbeatInterval:           &o.PUHeartbeatInterval,
			Annotations:                   &o.Annotations,
			ApplicationProxyPort:          &o.ApplicationProxyPort,
			AssociatedTags:                &o.AssociatedTags,
			AuditProfileSelectors:         &o.AuditProfileSelectors,
			AuditProfiles:                 &o.AuditProfiles,
			AuditSocketBufferSize:         &o.AuditSocketBufferSize,
			CreateTime:                    &o.CreateTime,
			Description:                   &o.Description,
			DockerSocketAddress:           &o.DockerSocketAddress,
			ExcludedInterfaces:            &o.ExcludedInterfaces,
			ExcludedNetworks:              &o.ExcludedNetworks,
			HostModeEnabled:               &o.HostModeEnabled,
			HostServices:                  &o.HostServices,
			IgnoreExpression:              &o.IgnoreExpression,
			KillContainersOnFailure:       &o.KillContainersOnFailure,
			KubernetesMetadataExtractor:   &o.KubernetesMetadataExtractor,
			KubernetesSupportEnabled:      &o.KubernetesSupportEnabled,
			LinuxProcessesSupportEnabled:  &o.LinuxProcessesSupportEnabled,
			Metadata:                      &o.Metadata,
			MetadataExtractor:             &o.MetadataExtractor,
			Name:                          &o.Name,
			Namespace:                     &o.Namespace,
			NormalizedTags:                &o.NormalizedTags,
			PolicySynchronizationInterval: &o.PolicySynchronizationInterval,
			Protected:                     &o.Protected,
			ProxyListenAddress:            &o.ProxyListenAddress,
			ReceiverNumberOfQueues:        &o.ReceiverNumberOfQueues,
			ReceiverQueue:                 &o.ReceiverQueue,
			ReceiverQueueSize:             &o.ReceiverQueueSize,
			RemoteEnforcerEnabled:         &o.RemoteEnforcerEnabled,
			TargetNetworks:                &o.TargetNetworks,
			TargetUDPNetworks:             &o.TargetUDPNetworks,
			TransmitterNumberOfQueues:     &o.TransmitterNumberOfQueues,
			TransmitterQueue:              &o.TransmitterQueue,
			TransmitterQueueSize:          &o.TransmitterQueueSize,
			TrustedCAs:                    &o.TrustedCAs,
			UpdateTime:                    &o.UpdateTime,
			ZHash:                         &o.ZHash,
			Zone:                          &o.Zone,
		}
	}

	sp := &SparseEnforcerProfile{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "IPTablesMarkValue":
			sp.IPTablesMarkValue = &(o.IPTablesMarkValue)
		case "PUBookkeepingInterval":
			sp.PUBookkeepingInterval = &(o.PUBookkeepingInterval)
		case "PUHeartbeatInterval":
			sp.PUHeartbeatInterval = &(o.PUHeartbeatInterval)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "applicationProxyPort":
			sp.ApplicationProxyPort = &(o.ApplicationProxyPort)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "auditProfileSelectors":
			sp.AuditProfileSelectors = &(o.AuditProfileSelectors)
		case "auditProfiles":
			sp.AuditProfiles = &(o.AuditProfiles)
		case "auditSocketBufferSize":
			sp.AuditSocketBufferSize = &(o.AuditSocketBufferSize)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "dockerSocketAddress":
			sp.DockerSocketAddress = &(o.DockerSocketAddress)
		case "excludedInterfaces":
			sp.ExcludedInterfaces = &(o.ExcludedInterfaces)
		case "excludedNetworks":
			sp.ExcludedNetworks = &(o.ExcludedNetworks)
		case "hostModeEnabled":
			sp.HostModeEnabled = &(o.HostModeEnabled)
		case "hostServices":
			sp.HostServices = &(o.HostServices)
		case "ignoreExpression":
			sp.IgnoreExpression = &(o.IgnoreExpression)
		case "killContainersOnFailure":
			sp.KillContainersOnFailure = &(o.KillContainersOnFailure)
		case "kubernetesMetadataExtractor":
			sp.KubernetesMetadataExtractor = &(o.KubernetesMetadataExtractor)
		case "kubernetesSupportEnabled":
			sp.KubernetesSupportEnabled = &(o.KubernetesSupportEnabled)
		case "linuxProcessesSupportEnabled":
			sp.LinuxProcessesSupportEnabled = &(o.LinuxProcessesSupportEnabled)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "metadataExtractor":
			sp.MetadataExtractor = &(o.MetadataExtractor)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "policySynchronizationInterval":
			sp.PolicySynchronizationInterval = &(o.PolicySynchronizationInterval)
		case "protected":
			sp.Protected = &(o.Protected)
		case "proxyListenAddress":
			sp.ProxyListenAddress = &(o.ProxyListenAddress)
		case "receiverNumberOfQueues":
			sp.ReceiverNumberOfQueues = &(o.ReceiverNumberOfQueues)
		case "receiverQueue":
			sp.ReceiverQueue = &(o.ReceiverQueue)
		case "receiverQueueSize":
			sp.ReceiverQueueSize = &(o.ReceiverQueueSize)
		case "remoteEnforcerEnabled":
			sp.RemoteEnforcerEnabled = &(o.RemoteEnforcerEnabled)
		case "targetNetworks":
			sp.TargetNetworks = &(o.TargetNetworks)
		case "targetUDPNetworks":
			sp.TargetUDPNetworks = &(o.TargetUDPNetworks)
		case "transmitterNumberOfQueues":
			sp.TransmitterNumberOfQueues = &(o.TransmitterNumberOfQueues)
		case "transmitterQueue":
			sp.TransmitterQueue = &(o.TransmitterQueue)
		case "transmitterQueueSize":
			sp.TransmitterQueueSize = &(o.TransmitterQueueSize)
		case "trustedCAs":
			sp.TrustedCAs = &(o.TrustedCAs)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseEnforcerProfile to the object.
func (o *EnforcerProfile) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEnforcerProfile)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.IPTablesMarkValue != nil {
		o.IPTablesMarkValue = *so.IPTablesMarkValue
	}
	if so.PUBookkeepingInterval != nil {
		o.PUBookkeepingInterval = *so.PUBookkeepingInterval
	}
	if so.PUHeartbeatInterval != nil {
		o.PUHeartbeatInterval = *so.PUHeartbeatInterval
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.ApplicationProxyPort != nil {
		o.ApplicationProxyPort = *so.ApplicationProxyPort
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.AuditProfileSelectors != nil {
		o.AuditProfileSelectors = *so.AuditProfileSelectors
	}
	if so.AuditProfiles != nil {
		o.AuditProfiles = *so.AuditProfiles
	}
	if so.AuditSocketBufferSize != nil {
		o.AuditSocketBufferSize = *so.AuditSocketBufferSize
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.DockerSocketAddress != nil {
		o.DockerSocketAddress = *so.DockerSocketAddress
	}
	if so.ExcludedInterfaces != nil {
		o.ExcludedInterfaces = *so.ExcludedInterfaces
	}
	if so.ExcludedNetworks != nil {
		o.ExcludedNetworks = *so.ExcludedNetworks
	}
	if so.HostModeEnabled != nil {
		o.HostModeEnabled = *so.HostModeEnabled
	}
	if so.HostServices != nil {
		o.HostServices = *so.HostServices
	}
	if so.IgnoreExpression != nil {
		o.IgnoreExpression = *so.IgnoreExpression
	}
	if so.KillContainersOnFailure != nil {
		o.KillContainersOnFailure = *so.KillContainersOnFailure
	}
	if so.KubernetesMetadataExtractor != nil {
		o.KubernetesMetadataExtractor = *so.KubernetesMetadataExtractor
	}
	if so.KubernetesSupportEnabled != nil {
		o.KubernetesSupportEnabled = *so.KubernetesSupportEnabled
	}
	if so.LinuxProcessesSupportEnabled != nil {
		o.LinuxProcessesSupportEnabled = *so.LinuxProcessesSupportEnabled
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.MetadataExtractor != nil {
		o.MetadataExtractor = *so.MetadataExtractor
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
	if so.PolicySynchronizationInterval != nil {
		o.PolicySynchronizationInterval = *so.PolicySynchronizationInterval
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.ProxyListenAddress != nil {
		o.ProxyListenAddress = *so.ProxyListenAddress
	}
	if so.ReceiverNumberOfQueues != nil {
		o.ReceiverNumberOfQueues = *so.ReceiverNumberOfQueues
	}
	if so.ReceiverQueue != nil {
		o.ReceiverQueue = *so.ReceiverQueue
	}
	if so.ReceiverQueueSize != nil {
		o.ReceiverQueueSize = *so.ReceiverQueueSize
	}
	if so.RemoteEnforcerEnabled != nil {
		o.RemoteEnforcerEnabled = *so.RemoteEnforcerEnabled
	}
	if so.TargetNetworks != nil {
		o.TargetNetworks = *so.TargetNetworks
	}
	if so.TargetUDPNetworks != nil {
		o.TargetUDPNetworks = *so.TargetUDPNetworks
	}
	if so.TransmitterNumberOfQueues != nil {
		o.TransmitterNumberOfQueues = *so.TransmitterNumberOfQueues
	}
	if so.TransmitterQueue != nil {
		o.TransmitterQueue = *so.TransmitterQueue
	}
	if so.TransmitterQueueSize != nil {
		o.TransmitterQueueSize = *so.TransmitterQueueSize
	}
	if so.TrustedCAs != nil {
		o.TrustedCAs = *so.TrustedCAs
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the EnforcerProfile.
func (o *EnforcerProfile) DeepCopy() *EnforcerProfile {

	if o == nil {
		return nil
	}

	out := &EnforcerProfile{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *EnforcerProfile.
func (o *EnforcerProfile) DeepCopyInto(out *EnforcerProfile) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy EnforcerProfile: %s", err))
	}

	*out = *target.(*EnforcerProfile)
}

// Validate valides the current information stored into the structure.
func (o *EnforcerProfile) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumInt("IPTablesMarkValue", o.IPTablesMarkValue, int(65000), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("PUBookkeepingInterval", o.PUBookkeepingInterval, `^[0-9]+[smh]$`, `must be a valid duration like <n>s or <n>s or <n>h`, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("PUHeartbeatInterval", o.PUHeartbeatInterval, `^[0-9]+[smh]$`, `must be a valid duration like <n>s or <n>s or <n>h`, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("applicationProxyPort", o.ApplicationProxyPort, int(65535), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("applicationProxyPort", o.ApplicationProxyPort, int(1), false); err != nil {
		errors = append(errors, err)
	}

	for _, sub := range o.AuditProfiles {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	if err := elemental.ValidateMaximumInt("auditSocketBufferSize", o.AuditSocketBufferSize, int(262144), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("dockerSocketAddress", o.DockerSocketAddress, `^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix://(/[^/]{1,16}){1,5}/?)$`, `must be a valid url or path starting by unix://`, false); err != nil {
		errors = append(errors, err)
	}

	for _, sub := range o.HostServices {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	if err := ValidateHostServicesList("hostServices", o.HostServices); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("kubernetesMetadataExtractor", string(o.KubernetesMetadataExtractor), []string{"KubeSquall", "PodAtomic", "PodContainers"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("metadataExtractor", string(o.MetadataExtractor), []string{"Docker", "ECS", "Kubernetes"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("policySynchronizationInterval", o.PolicySynchronizationInterval, `^[0-9]+[smh]$`, `must be a valid duration like <n>s or <n>s or <n>h`, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("proxyListenAddress", o.ProxyListenAddress, `^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix://(/[^/]{1,16}){1,5}/?)$`, `must be a valid url or path starting by unix://`, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("receiverNumberOfQueues", o.ReceiverNumberOfQueues, int(16), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("receiverNumberOfQueues", o.ReceiverNumberOfQueues, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("receiverQueue", o.ReceiverQueue, int(1000), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("receiverQueueSize", o.ReceiverQueueSize, int(5000), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("receiverQueueSize", o.ReceiverQueueSize, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("transmitterNumberOfQueues", o.TransmitterNumberOfQueues, int(16), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("transmitterNumberOfQueues", o.TransmitterNumberOfQueues, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("transmitterQueue", o.TransmitterQueue, int(1000), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("transmitterQueue", o.TransmitterQueue, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("transmitterQueueSize", o.TransmitterQueueSize, int(1000), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("transmitterQueueSize", o.TransmitterQueueSize, int(1), false); err != nil {
		errors = append(errors, err)
	}

	// Custom object validation.
	if err := ValidateEnforcerProfile(o); err != nil {
		switch e := err.(type) {
		case elemental.Errors:
			errors = append(errors, e...)
		case elemental.Error:
			errors = append(errors, e)
		default:
			errors = append(errors, e)
		}
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
func (*EnforcerProfile) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerProfileAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerProfileLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerProfile) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerProfileAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *EnforcerProfile) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "IPTablesMarkValue":
		return o.IPTablesMarkValue
	case "PUBookkeepingInterval":
		return o.PUBookkeepingInterval
	case "PUHeartbeatInterval":
		return o.PUHeartbeatInterval
	case "annotations":
		return o.Annotations
	case "applicationProxyPort":
		return o.ApplicationProxyPort
	case "associatedTags":
		return o.AssociatedTags
	case "auditProfileSelectors":
		return o.AuditProfileSelectors
	case "auditProfiles":
		return o.AuditProfiles
	case "auditSocketBufferSize":
		return o.AuditSocketBufferSize
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "dockerSocketAddress":
		return o.DockerSocketAddress
	case "excludedInterfaces":
		return o.ExcludedInterfaces
	case "excludedNetworks":
		return o.ExcludedNetworks
	case "hostModeEnabled":
		return o.HostModeEnabled
	case "hostServices":
		return o.HostServices
	case "ignoreExpression":
		return o.IgnoreExpression
	case "killContainersOnFailure":
		return o.KillContainersOnFailure
	case "kubernetesMetadataExtractor":
		return o.KubernetesMetadataExtractor
	case "kubernetesSupportEnabled":
		return o.KubernetesSupportEnabled
	case "linuxProcessesSupportEnabled":
		return o.LinuxProcessesSupportEnabled
	case "metadata":
		return o.Metadata
	case "metadataExtractor":
		return o.MetadataExtractor
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "policySynchronizationInterval":
		return o.PolicySynchronizationInterval
	case "protected":
		return o.Protected
	case "proxyListenAddress":
		return o.ProxyListenAddress
	case "receiverNumberOfQueues":
		return o.ReceiverNumberOfQueues
	case "receiverQueue":
		return o.ReceiverQueue
	case "receiverQueueSize":
		return o.ReceiverQueueSize
	case "remoteEnforcerEnabled":
		return o.RemoteEnforcerEnabled
	case "targetNetworks":
		return o.TargetNetworks
	case "targetUDPNetworks":
		return o.TargetUDPNetworks
	case "transmitterNumberOfQueues":
		return o.TransmitterNumberOfQueues
	case "transmitterQueue":
		return o.TransmitterQueue
	case "transmitterQueueSize":
		return o.TransmitterQueueSize
	case "trustedCAs":
		return o.TrustedCAs
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// EnforcerProfileAttributesMap represents the map of attribute for EnforcerProfile.
var EnforcerProfileAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"IPTablesMarkValue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IPTablesMarkValue",
		DefaultValue:   1000,
		Description:    `IptablesMarkValue is the mark value to be used in an iptables implementation.`,
		Exposed:        true,
		MaxValue:       65000,
		Name:           "IPTablesMarkValue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"PUBookkeepingInterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "PUBookkeepingInterval",
		DefaultValue:   "15m",
		Description:    `PUBookkeepingInterval configures how often the PU will be synchronized.`,
		Exposed:        true,
		Name:           "PUBookkeepingInterval",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"PUHeartbeatInterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "PUHeartbeatInterval",
		DefaultValue:   "5s",
		Description:    `PUHeartbeatInterval configures the heart beat interval.`,
		Exposed:        true,
		Name:           "PUHeartbeatInterval",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map_of_string_of_list_of_strings",
		Type:           "external",
	},
	"ApplicationProxyPort": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ApplicationProxyPort",
		DefaultValue:   20992,
		Description:    `Port used by aporeto application proxy.`,
		Exposed:        true,
		MaxValue:       65535,
		MinValue:       1,
		Name:           "applicationProxyPort",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AuditProfileSelectors": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditProfileSelectors",
		Description: `AuditProfileSelectors is the list of tags (key/value pairs) that define the
audit policies that must be implemented by this enforcer. The enforcer will
implement all policies that match any of these tags.`,
		Exposed: true,
		Name:    "auditProfileSelectors",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"AuditProfiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AuditProfiles",
		Description: `AuditProfiles returns the audit rules associated with the enforcer profile. This
is a read only attribute when an enforcer profile is resolved for an enforcer.`,
		Exposed:  true,
		Name:     "auditProfiles",
		ReadOnly: true,
		SubType:  "auditprofile",
		Type:     "refList",
	},
	"AuditSocketBufferSize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditSocketBufferSize",
		DefaultValue:   16384,
		Description:    `AuditSocketBufferSize is the size of the audit socket buffer. Default 16384.`,
		Exposed:        true,
		MaxValue:       262144,
		Name:           "auditSocketBufferSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"DockerSocketAddress": elemental.AttributeSpecification{
		AllowedChars:   `^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix://(/[^/]{1,16}){1,5}/?)$`,
		AllowedChoices: []string{},
		ConvertedName:  "DockerSocketAddress",
		DefaultValue:   "unix:///var/run/docker.sock",
		Description:    `DockerSocketAddress is the address of the docker daemon.`,
		Exposed:        true,
		Name:           "dockerSocketAddress",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ExcludedInterfaces": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedInterfaces",
		Description:    `ExcludedInterfaces is a list of interfaces that must be excluded.`,
		Exposed:        true,
		Name:           "excludedInterfaces",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ExcludedNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedNetworks",
		Description: `ExcludedNetworks is the list of networks that must be excluded for this
enforcer.`,
		Exposed:   true,
		Name:      "excludedNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"HostModeEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "HostModeEnabled",
		Description: `hostModeEnabled enables protection of the complete host. When this option is
turned on, all incoming and outgoing flows will be monitored. Flows will
be allowed if and only if a network policy has been created to allow the flow.`,
		Exposed:   true,
		Name:      "hostModeEnabled",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
	},
	"HostServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "HostServices",
		Description: `HostServices is a list of services that must be activated by default to all
enforcers matching this profile.`,
		Exposed: true,
		Name:    "hostServices",
		Stored:  true,
		SubType: "hostservice",
		Type:    "refList",
	},
	"IgnoreExpression": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IgnoreExpression",
		Description: `IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
docker container started with labels matching the rule.`,
		Exposed: true,
		Name:    "ignoreExpression",
		Stored:  true,
		SubType: "list_of_lists_of_strings",
		Type:    "external",
	},
	"KillContainersOnFailure": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KillContainersOnFailure",
		Description: `KillContainersOnFailure will configure the enforcers to kill any containers if
there are policy failures.`,
		Exposed: true,
		Name:    "killContainersOnFailure",
		Stored:  true,
		Type:    "boolean",
	},
	"KubernetesMetadataExtractor": elemental.AttributeSpecification{
		AllowedChoices: []string{"KubeSquall", "PodAtomic", "PodContainers"},
		ConvertedName:  "KubernetesMetadataExtractor",
		DefaultValue:   EnforcerProfileKubernetesMetadataExtractorPodAtomic,
		Description: `Select which metadata extractor to use to process new processing units from
Kubernetes.`,
		Exposed:   true,
		Name:      "kubernetesMetadataExtractor",
		Orderable: true,
		Stored:    true,
		Type:      "enum",
	},
	"KubernetesSupportEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KubernetesSupportEnabled",
		Description:    `KubernetesSupportEnabled enables kubernetes mode for the enforcer.`,
		Exposed:        true,
		Name:           "kubernetesSupportEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"LinuxProcessesSupportEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LinuxProcessesSupportEnabled",
		DefaultValue:   true,
		Description:    `LinuxProcessesSupportEnabled configures support for Linux processes.`,
		Exposed:        true,
		Name:           "linuxProcessesSupportEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
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
	"MetadataExtractor": elemental.AttributeSpecification{
		AllowedChoices: []string{"Docker", "ECS", "Kubernetes"},
		ConvertedName:  "MetadataExtractor",
		DefaultValue:   EnforcerProfileMetadataExtractorDocker,
		Description:    `Select which metadata extractor to use to process new processing units.`,
		Exposed:        true,
		Name:           "metadataExtractor",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
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
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
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
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
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
	"PolicySynchronizationInterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "PolicySynchronizationInterval",
		DefaultValue:   "10m",
		Description: `PolicySynchronizationInterval configures how often the policy will be
resynchronized.`,
		Exposed:   true,
		Name:      "policySynchronizationInterval",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"ProxyListenAddress": elemental.AttributeSpecification{
		AllowedChars:   `^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix://(/[^/]{1,16}){1,5}/?)$`,
		AllowedChoices: []string{},
		ConvertedName:  "ProxyListenAddress",
		DefaultValue:   "unix:///var/run/aporeto.sock",
		Description: `ProxyListenAddress is the address the enforcer should use to listen for API
calls. It can be a port (example :9443) or socket path
example:
  unix:///var/run/aporeto.sock.`,
		Exposed:   true,
		Name:      "proxyListenAddress",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"ReceiverNumberOfQueues": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReceiverNumberOfQueues",
		DefaultValue:   4,
		Description: `ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network
receiver starting at the ReceiverQueue.`,
		Exposed:   true,
		MaxValue:  16,
		MinValue:  1,
		Name:      "receiverNumberOfQueues",
		Orderable: true,
		Stored:    true,
		Type:      "integer",
	},
	"ReceiverQueue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReceiverQueue",
		Description:    `ReceiverQueue is the base queue number for traffic from the network.`,
		Exposed:        true,
		MaxValue:       1000,
		Name:           "receiverQueue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"ReceiverQueueSize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReceiverQueueSize",
		DefaultValue:   500,
		Description:    `ReceiverQueueSize is the queue size of the receiver.`,
		Exposed:        true,
		MaxValue:       5000,
		MinValue:       1,
		Name:           "receiverQueueSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"RemoteEnforcerEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RemoteEnforcerEnabled",
		DefaultValue:   true,
		Description: `RemoteEnforcerEnabled inidicates whether a single enforcer should be used or a
distributed enforcer. True means distributed.`,
		Exposed:   true,
		Name:      "remoteEnforcerEnabled",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
	},
	"TargetNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNetworks",
		Description:    `TargetNetworks is the list of networks that authorization should be applied.`,
		Exposed:        true,
		Name:           "targetNetworks",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"TargetUDPNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetUDPNetworks",
		Description: `TargetUDPNetworks is the list of UDP networks that authorization should be
applied.`,
		Exposed:   true,
		Name:      "targetUDPNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"TransmitterNumberOfQueues": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TransmitterNumberOfQueues",
		DefaultValue:   4,
		Description:    `TransmitterNumberOfQueues is the number of queues for application traffic.`,
		Exposed:        true,
		MaxValue:       16,
		MinValue:       1,
		Name:           "transmitterNumberOfQueues",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"TransmitterQueue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TransmitterQueue",
		DefaultValue:   4,
		Description: `TransmitterQueue is the queue number for traffic from the applications starting
at the transmitterQueue.`,
		Exposed:   true,
		MaxValue:  1000,
		MinValue:  1,
		Name:      "transmitterQueue",
		Orderable: true,
		Stored:    true,
		Type:      "integer",
	},
	"TransmitterQueueSize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TransmitterQueueSize",
		DefaultValue:   500,
		Description:    `TransmitterQueueSize is the size of the queue for application traffic.`,
		Exposed:        true,
		MaxValue:       1000,
		MinValue:       1,
		Name:           "transmitterQueueSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"TrustedCAs": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TrustedCAs",
		Description:    `List of trusted CA. If empty the main chain of trust will be used.`,
		Exposed:        true,
		Name:           "trustedCAs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// EnforcerProfileLowerCaseAttributesMap represents the map of attribute for EnforcerProfile.
var EnforcerProfileLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"iptablesmarkvalue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IPTablesMarkValue",
		DefaultValue:   1000,
		Description:    `IptablesMarkValue is the mark value to be used in an iptables implementation.`,
		Exposed:        true,
		MaxValue:       65000,
		Name:           "IPTablesMarkValue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"pubookkeepinginterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "PUBookkeepingInterval",
		DefaultValue:   "15m",
		Description:    `PUBookkeepingInterval configures how often the PU will be synchronized.`,
		Exposed:        true,
		Name:           "PUBookkeepingInterval",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"puheartbeatinterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "PUHeartbeatInterval",
		DefaultValue:   "5s",
		Description:    `PUHeartbeatInterval configures the heart beat interval.`,
		Exposed:        true,
		Name:           "PUHeartbeatInterval",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map_of_string_of_list_of_strings",
		Type:           "external",
	},
	"applicationproxyport": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ApplicationProxyPort",
		DefaultValue:   20992,
		Description:    `Port used by aporeto application proxy.`,
		Exposed:        true,
		MaxValue:       65535,
		MinValue:       1,
		Name:           "applicationProxyPort",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"auditprofileselectors": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditProfileSelectors",
		Description: `AuditProfileSelectors is the list of tags (key/value pairs) that define the
audit policies that must be implemented by this enforcer. The enforcer will
implement all policies that match any of these tags.`,
		Exposed: true,
		Name:    "auditProfileSelectors",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"auditprofiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AuditProfiles",
		Description: `AuditProfiles returns the audit rules associated with the enforcer profile. This
is a read only attribute when an enforcer profile is resolved for an enforcer.`,
		Exposed:  true,
		Name:     "auditProfiles",
		ReadOnly: true,
		SubType:  "auditprofile",
		Type:     "refList",
	},
	"auditsocketbuffersize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditSocketBufferSize",
		DefaultValue:   16384,
		Description:    `AuditSocketBufferSize is the size of the audit socket buffer. Default 16384.`,
		Exposed:        true,
		MaxValue:       262144,
		Name:           "auditSocketBufferSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"dockersocketaddress": elemental.AttributeSpecification{
		AllowedChars:   `^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix://(/[^/]{1,16}){1,5}/?)$`,
		AllowedChoices: []string{},
		ConvertedName:  "DockerSocketAddress",
		DefaultValue:   "unix:///var/run/docker.sock",
		Description:    `DockerSocketAddress is the address of the docker daemon.`,
		Exposed:        true,
		Name:           "dockerSocketAddress",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"excludedinterfaces": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedInterfaces",
		Description:    `ExcludedInterfaces is a list of interfaces that must be excluded.`,
		Exposed:        true,
		Name:           "excludedInterfaces",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"excludednetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedNetworks",
		Description: `ExcludedNetworks is the list of networks that must be excluded for this
enforcer.`,
		Exposed:   true,
		Name:      "excludedNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"hostmodeenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "HostModeEnabled",
		Description: `hostModeEnabled enables protection of the complete host. When this option is
turned on, all incoming and outgoing flows will be monitored. Flows will
be allowed if and only if a network policy has been created to allow the flow.`,
		Exposed:   true,
		Name:      "hostModeEnabled",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
	},
	"hostservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "HostServices",
		Description: `HostServices is a list of services that must be activated by default to all
enforcers matching this profile.`,
		Exposed: true,
		Name:    "hostServices",
		Stored:  true,
		SubType: "hostservice",
		Type:    "refList",
	},
	"ignoreexpression": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IgnoreExpression",
		Description: `IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
docker container started with labels matching the rule.`,
		Exposed: true,
		Name:    "ignoreExpression",
		Stored:  true,
		SubType: "list_of_lists_of_strings",
		Type:    "external",
	},
	"killcontainersonfailure": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KillContainersOnFailure",
		Description: `KillContainersOnFailure will configure the enforcers to kill any containers if
there are policy failures.`,
		Exposed: true,
		Name:    "killContainersOnFailure",
		Stored:  true,
		Type:    "boolean",
	},
	"kubernetesmetadataextractor": elemental.AttributeSpecification{
		AllowedChoices: []string{"KubeSquall", "PodAtomic", "PodContainers"},
		ConvertedName:  "KubernetesMetadataExtractor",
		DefaultValue:   EnforcerProfileKubernetesMetadataExtractorPodAtomic,
		Description: `Select which metadata extractor to use to process new processing units from
Kubernetes.`,
		Exposed:   true,
		Name:      "kubernetesMetadataExtractor",
		Orderable: true,
		Stored:    true,
		Type:      "enum",
	},
	"kubernetessupportenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KubernetesSupportEnabled",
		Description:    `KubernetesSupportEnabled enables kubernetes mode for the enforcer.`,
		Exposed:        true,
		Name:           "kubernetesSupportEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"linuxprocessessupportenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LinuxProcessesSupportEnabled",
		DefaultValue:   true,
		Description:    `LinuxProcessesSupportEnabled configures support for Linux processes.`,
		Exposed:        true,
		Name:           "linuxProcessesSupportEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
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
	"metadataextractor": elemental.AttributeSpecification{
		AllowedChoices: []string{"Docker", "ECS", "Kubernetes"},
		ConvertedName:  "MetadataExtractor",
		DefaultValue:   EnforcerProfileMetadataExtractorDocker,
		Description:    `Select which metadata extractor to use to process new processing units.`,
		Exposed:        true,
		Name:           "metadataExtractor",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
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
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
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
	"policysynchronizationinterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "PolicySynchronizationInterval",
		DefaultValue:   "10m",
		Description: `PolicySynchronizationInterval configures how often the policy will be
resynchronized.`,
		Exposed:   true,
		Name:      "policySynchronizationInterval",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"proxylistenaddress": elemental.AttributeSpecification{
		AllowedChars:   `^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix://(/[^/]{1,16}){1,5}/?)$`,
		AllowedChoices: []string{},
		ConvertedName:  "ProxyListenAddress",
		DefaultValue:   "unix:///var/run/aporeto.sock",
		Description: `ProxyListenAddress is the address the enforcer should use to listen for API
calls. It can be a port (example :9443) or socket path
example:
  unix:///var/run/aporeto.sock.`,
		Exposed:   true,
		Name:      "proxyListenAddress",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"receivernumberofqueues": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReceiverNumberOfQueues",
		DefaultValue:   4,
		Description: `ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network
receiver starting at the ReceiverQueue.`,
		Exposed:   true,
		MaxValue:  16,
		MinValue:  1,
		Name:      "receiverNumberOfQueues",
		Orderable: true,
		Stored:    true,
		Type:      "integer",
	},
	"receiverqueue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReceiverQueue",
		Description:    `ReceiverQueue is the base queue number for traffic from the network.`,
		Exposed:        true,
		MaxValue:       1000,
		Name:           "receiverQueue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"receiverqueuesize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReceiverQueueSize",
		DefaultValue:   500,
		Description:    `ReceiverQueueSize is the queue size of the receiver.`,
		Exposed:        true,
		MaxValue:       5000,
		MinValue:       1,
		Name:           "receiverQueueSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"remoteenforcerenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RemoteEnforcerEnabled",
		DefaultValue:   true,
		Description: `RemoteEnforcerEnabled inidicates whether a single enforcer should be used or a
distributed enforcer. True means distributed.`,
		Exposed:   true,
		Name:      "remoteEnforcerEnabled",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
	},
	"targetnetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNetworks",
		Description:    `TargetNetworks is the list of networks that authorization should be applied.`,
		Exposed:        true,
		Name:           "targetNetworks",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"targetudpnetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetUDPNetworks",
		Description: `TargetUDPNetworks is the list of UDP networks that authorization should be
applied.`,
		Exposed:   true,
		Name:      "targetUDPNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"transmitternumberofqueues": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TransmitterNumberOfQueues",
		DefaultValue:   4,
		Description:    `TransmitterNumberOfQueues is the number of queues for application traffic.`,
		Exposed:        true,
		MaxValue:       16,
		MinValue:       1,
		Name:           "transmitterNumberOfQueues",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"transmitterqueue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TransmitterQueue",
		DefaultValue:   4,
		Description: `TransmitterQueue is the queue number for traffic from the applications starting
at the transmitterQueue.`,
		Exposed:   true,
		MaxValue:  1000,
		MinValue:  1,
		Name:      "transmitterQueue",
		Orderable: true,
		Stored:    true,
		Type:      "integer",
	},
	"transmitterqueuesize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TransmitterQueueSize",
		DefaultValue:   500,
		Description:    `TransmitterQueueSize is the size of the queue for application traffic.`,
		Exposed:        true,
		MaxValue:       1000,
		MinValue:       1,
		Name:           "transmitterQueueSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"trustedcas": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TrustedCAs",
		Description:    `List of trusted CA. If empty the main chain of trust will be used.`,
		Exposed:        true,
		Name:           "trustedCAs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// SparseEnforcerProfilesList represents a list of SparseEnforcerProfiles
type SparseEnforcerProfilesList []*SparseEnforcerProfile

// Identity returns the identity of the objects in the list.
func (o SparseEnforcerProfilesList) Identity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Copy returns a pointer to a copy the SparseEnforcerProfilesList.
func (o SparseEnforcerProfilesList) Copy() elemental.Identifiables {

	copy := append(SparseEnforcerProfilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEnforcerProfilesList.
func (o SparseEnforcerProfilesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEnforcerProfilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEnforcerProfile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEnforcerProfilesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEnforcerProfilesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseEnforcerProfilesList converted to EnforcerProfilesList.
func (o SparseEnforcerProfilesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEnforcerProfilesList) Version() int {

	return 1
}

// SparseEnforcerProfile represents the sparse version of a enforcerprofile.
type SparseEnforcerProfile struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// IptablesMarkValue is the mark value to be used in an iptables implementation.
	IPTablesMarkValue *int `json:"IPTablesMarkValue,omitempty" bson:"iptablesmarkvalue" mapstructure:"IPTablesMarkValue,omitempty"`

	// PUBookkeepingInterval configures how often the PU will be synchronized.
	PUBookkeepingInterval *string `json:"PUBookkeepingInterval,omitempty" bson:"pubookkeepinginterval" mapstructure:"PUBookkeepingInterval,omitempty"`

	// PUHeartbeatInterval configures the heart beat interval.
	PUHeartbeatInterval *string `json:"PUHeartbeatInterval,omitempty" bson:"puheartbeatinterval" mapstructure:"PUHeartbeatInterval,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Port used by aporeto application proxy.
	ApplicationProxyPort *int `json:"applicationProxyPort,omitempty" bson:"applicationproxyport" mapstructure:"applicationProxyPort,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// AuditProfileSelectors is the list of tags (key/value pairs) that define the
	// audit policies that must be implemented by this enforcer. The enforcer will
	// implement all policies that match any of these tags.
	AuditProfileSelectors *[]string `json:"auditProfileSelectors,omitempty" bson:"auditprofileselectors" mapstructure:"auditProfileSelectors,omitempty"`

	// AuditProfiles returns the audit rules associated with the enforcer profile. This
	// is a read only attribute when an enforcer profile is resolved for an enforcer.
	AuditProfiles *AuditProfilesList `json:"auditProfiles,omitempty" bson:"-" mapstructure:"auditProfiles,omitempty"`

	// AuditSocketBufferSize is the size of the audit socket buffer. Default 16384.
	AuditSocketBufferSize *int `json:"auditSocketBufferSize,omitempty" bson:"auditsocketbuffersize" mapstructure:"auditSocketBufferSize,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// DockerSocketAddress is the address of the docker daemon.
	DockerSocketAddress *string `json:"dockerSocketAddress,omitempty" bson:"dockersocketaddress" mapstructure:"dockerSocketAddress,omitempty"`

	// ExcludedInterfaces is a list of interfaces that must be excluded.
	ExcludedInterfaces *[]string `json:"excludedInterfaces,omitempty" bson:"excludedinterfaces" mapstructure:"excludedInterfaces,omitempty"`

	// ExcludedNetworks is the list of networks that must be excluded for this
	// enforcer.
	ExcludedNetworks *[]string `json:"excludedNetworks,omitempty" bson:"excludednetworks" mapstructure:"excludedNetworks,omitempty"`

	// hostModeEnabled enables protection of the complete host. When this option is
	// turned on, all incoming and outgoing flows will be monitored. Flows will
	// be allowed if and only if a network policy has been created to allow the flow.
	HostModeEnabled *bool `json:"hostModeEnabled,omitempty" bson:"hostmodeenabled" mapstructure:"hostModeEnabled,omitempty"`

	// HostServices is a list of services that must be activated by default to all
	// enforcers matching this profile.
	HostServices *[]*HostService `json:"hostServices,omitempty" bson:"hostservices" mapstructure:"hostServices,omitempty"`

	// IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
	// docker container started with labels matching the rule.
	IgnoreExpression *[][]string `json:"ignoreExpression,omitempty" bson:"ignoreexpression" mapstructure:"ignoreExpression,omitempty"`

	// KillContainersOnFailure will configure the enforcers to kill any containers if
	// there are policy failures.
	KillContainersOnFailure *bool `json:"killContainersOnFailure,omitempty" bson:"killcontainersonfailure" mapstructure:"killContainersOnFailure,omitempty"`

	// Select which metadata extractor to use to process new processing units from
	// Kubernetes.
	KubernetesMetadataExtractor *EnforcerProfileKubernetesMetadataExtractorValue `json:"kubernetesMetadataExtractor,omitempty" bson:"kubernetesmetadataextractor" mapstructure:"kubernetesMetadataExtractor,omitempty"`

	// KubernetesSupportEnabled enables kubernetes mode for the enforcer.
	KubernetesSupportEnabled *bool `json:"kubernetesSupportEnabled,omitempty" bson:"kubernetessupportenabled" mapstructure:"kubernetesSupportEnabled,omitempty"`

	// LinuxProcessesSupportEnabled configures support for Linux processes.
	LinuxProcessesSupportEnabled *bool `json:"linuxProcessesSupportEnabled,omitempty" bson:"linuxprocessessupportenabled" mapstructure:"linuxProcessesSupportEnabled,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Select which metadata extractor to use to process new processing units.
	MetadataExtractor *EnforcerProfileMetadataExtractorValue `json:"metadataExtractor,omitempty" bson:"metadataextractor" mapstructure:"metadataExtractor,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// PolicySynchronizationInterval configures how often the policy will be
	// resynchronized.
	PolicySynchronizationInterval *string `json:"policySynchronizationInterval,omitempty" bson:"policysynchronizationinterval" mapstructure:"policySynchronizationInterval,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected" mapstructure:"protected,omitempty"`

	// ProxyListenAddress is the address the enforcer should use to listen for API
	// calls. It can be a port (example :9443) or socket path
	// example:
	//   unix:///var/run/aporeto.sock.
	ProxyListenAddress *string `json:"proxyListenAddress,omitempty" bson:"proxylistenaddress" mapstructure:"proxyListenAddress,omitempty"`

	// ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network
	// receiver starting at the ReceiverQueue.
	ReceiverNumberOfQueues *int `json:"receiverNumberOfQueues,omitempty" bson:"receivernumberofqueues" mapstructure:"receiverNumberOfQueues,omitempty"`

	// ReceiverQueue is the base queue number for traffic from the network.
	ReceiverQueue *int `json:"receiverQueue,omitempty" bson:"receiverqueue" mapstructure:"receiverQueue,omitempty"`

	// ReceiverQueueSize is the queue size of the receiver.
	ReceiverQueueSize *int `json:"receiverQueueSize,omitempty" bson:"receiverqueuesize" mapstructure:"receiverQueueSize,omitempty"`

	// RemoteEnforcerEnabled inidicates whether a single enforcer should be used or a
	// distributed enforcer. True means distributed.
	RemoteEnforcerEnabled *bool `json:"remoteEnforcerEnabled,omitempty" bson:"remoteenforcerenabled" mapstructure:"remoteEnforcerEnabled,omitempty"`

	// TargetNetworks is the list of networks that authorization should be applied.
	TargetNetworks *[]string `json:"targetNetworks,omitempty" bson:"targetnetworks" mapstructure:"targetNetworks,omitempty"`

	// TargetUDPNetworks is the list of UDP networks that authorization should be
	// applied.
	TargetUDPNetworks *[]string `json:"targetUDPNetworks,omitempty" bson:"targetudpnetworks" mapstructure:"targetUDPNetworks,omitempty"`

	// TransmitterNumberOfQueues is the number of queues for application traffic.
	TransmitterNumberOfQueues *int `json:"transmitterNumberOfQueues,omitempty" bson:"transmitternumberofqueues" mapstructure:"transmitterNumberOfQueues,omitempty"`

	// TransmitterQueue is the queue number for traffic from the applications starting
	// at the transmitterQueue.
	TransmitterQueue *int `json:"transmitterQueue,omitempty" bson:"transmitterqueue" mapstructure:"transmitterQueue,omitempty"`

	// TransmitterQueueSize is the size of the queue for application traffic.
	TransmitterQueueSize *int `json:"transmitterQueueSize,omitempty" bson:"transmitterqueuesize" mapstructure:"transmitterQueueSize,omitempty"`

	// List of trusted CA. If empty the main chain of trust will be used.
	TrustedCAs *[]string `json:"trustedCAs,omitempty" bson:"trustedcas" mapstructure:"trustedCAs,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-,omitempty" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-,omitempty" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseEnforcerProfile returns a new  SparseEnforcerProfile.
func NewSparseEnforcerProfile() *SparseEnforcerProfile {
	return &SparseEnforcerProfile{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEnforcerProfile) Identity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEnforcerProfile) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEnforcerProfile) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseEnforcerProfile) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEnforcerProfile) ToPlain() elemental.PlainIdentifiable {

	out := NewEnforcerProfile()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.IPTablesMarkValue != nil {
		out.IPTablesMarkValue = *o.IPTablesMarkValue
	}
	if o.PUBookkeepingInterval != nil {
		out.PUBookkeepingInterval = *o.PUBookkeepingInterval
	}
	if o.PUHeartbeatInterval != nil {
		out.PUHeartbeatInterval = *o.PUHeartbeatInterval
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.ApplicationProxyPort != nil {
		out.ApplicationProxyPort = *o.ApplicationProxyPort
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.AuditProfileSelectors != nil {
		out.AuditProfileSelectors = *o.AuditProfileSelectors
	}
	if o.AuditProfiles != nil {
		out.AuditProfiles = *o.AuditProfiles
	}
	if o.AuditSocketBufferSize != nil {
		out.AuditSocketBufferSize = *o.AuditSocketBufferSize
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.DockerSocketAddress != nil {
		out.DockerSocketAddress = *o.DockerSocketAddress
	}
	if o.ExcludedInterfaces != nil {
		out.ExcludedInterfaces = *o.ExcludedInterfaces
	}
	if o.ExcludedNetworks != nil {
		out.ExcludedNetworks = *o.ExcludedNetworks
	}
	if o.HostModeEnabled != nil {
		out.HostModeEnabled = *o.HostModeEnabled
	}
	if o.HostServices != nil {
		out.HostServices = *o.HostServices
	}
	if o.IgnoreExpression != nil {
		out.IgnoreExpression = *o.IgnoreExpression
	}
	if o.KillContainersOnFailure != nil {
		out.KillContainersOnFailure = *o.KillContainersOnFailure
	}
	if o.KubernetesMetadataExtractor != nil {
		out.KubernetesMetadataExtractor = *o.KubernetesMetadataExtractor
	}
	if o.KubernetesSupportEnabled != nil {
		out.KubernetesSupportEnabled = *o.KubernetesSupportEnabled
	}
	if o.LinuxProcessesSupportEnabled != nil {
		out.LinuxProcessesSupportEnabled = *o.LinuxProcessesSupportEnabled
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.MetadataExtractor != nil {
		out.MetadataExtractor = *o.MetadataExtractor
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
	if o.PolicySynchronizationInterval != nil {
		out.PolicySynchronizationInterval = *o.PolicySynchronizationInterval
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.ProxyListenAddress != nil {
		out.ProxyListenAddress = *o.ProxyListenAddress
	}
	if o.ReceiverNumberOfQueues != nil {
		out.ReceiverNumberOfQueues = *o.ReceiverNumberOfQueues
	}
	if o.ReceiverQueue != nil {
		out.ReceiverQueue = *o.ReceiverQueue
	}
	if o.ReceiverQueueSize != nil {
		out.ReceiverQueueSize = *o.ReceiverQueueSize
	}
	if o.RemoteEnforcerEnabled != nil {
		out.RemoteEnforcerEnabled = *o.RemoteEnforcerEnabled
	}
	if o.TargetNetworks != nil {
		out.TargetNetworks = *o.TargetNetworks
	}
	if o.TargetUDPNetworks != nil {
		out.TargetUDPNetworks = *o.TargetUDPNetworks
	}
	if o.TransmitterNumberOfQueues != nil {
		out.TransmitterNumberOfQueues = *o.TransmitterNumberOfQueues
	}
	if o.TransmitterQueue != nil {
		out.TransmitterQueue = *o.TransmitterQueue
	}
	if o.TransmitterQueueSize != nil {
		out.TransmitterQueueSize = *o.TransmitterQueueSize
	}
	if o.TrustedCAs != nil {
		out.TrustedCAs = *o.TrustedCAs
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
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
func (o *SparseEnforcerProfile) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseEnforcerProfile) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseEnforcerProfile) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseEnforcerProfile) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseEnforcerProfile) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseEnforcerProfile) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseEnforcerProfile) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseEnforcerProfile) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseEnforcerProfile) GetProtected() bool {

	return *o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseEnforcerProfile) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseEnforcerProfile) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseEnforcerProfile) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseEnforcerProfile.
func (o *SparseEnforcerProfile) DeepCopy() *SparseEnforcerProfile {

	if o == nil {
		return nil
	}

	out := &SparseEnforcerProfile{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEnforcerProfile.
func (o *SparseEnforcerProfile) DeepCopyInto(out *SparseEnforcerProfile) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEnforcerProfile: %s", err))
	}

	*out = *target.(*SparseEnforcerProfile)
}
