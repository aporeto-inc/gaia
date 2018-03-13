package types

// PortMapping represents a public/private port mapping.
type PortMapping struct {
	PublicPort  int `json:"publicPort" bson:"publicPort" mapstructure:"publicPort,omitempty"`
	PrivatePort int `json:"privatePort" bson:"privatePort" mapstructure:"privatePort,omitempty"`
}

// NewPortMapping creates a new port mapping.
func NewPortMapping() *PortMapping {
	return &PortMapping{}
}
