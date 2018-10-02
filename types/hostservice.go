package types

// HostService is a service associated with an enforcer profile for a host.
type HostService struct {
	Name        string                     `json:"name" bson:"name" mapstructure:"name,omitempty"`
	NetworkOnly bool                       `json:"networkonly" bson:"networkonly" mapstructure:"networkonly,omitempty"`
	Services    ProcessingUnitServicesList `json:"services" bson:"services" mapstructure:"services,omitempty"`
}

// HostServicesList is a list of ProcessingUnitServices
type HostServicesList []*HostService
