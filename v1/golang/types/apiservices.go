package types

// IPList is a list of IP addresses
type IPList []IPAddress

// IPAddress is an IP or subnet expressed as a string
type IPAddress string

// ExposedAPIList is a list of API endpoints and associated tags.
type ExposedAPIList []ExposedAPI

// ExposedAPI is an exposed API defined by the URI, verb, and associated tags.
// The URIs must be valid Golang regular expressions.
type ExposedAPI struct {
	URI    string   `json:"URI" bson:"URI" mapstructure:"URI,omitempty"`
	Verbs  []string `json:"verb" bson:"verb" mapstructure:"verb,omitempty"`
	Scopes []string `json:"tags" bson:"tags" mapstructure:"tags,omitempty"`
}

// PortList is a list of ports for the service.
type PortList []string
