package types

import (
	"fmt"
	"regexp"
)

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
	Scopes []string `json:"scopes" bson:"scopes" mapstructure:"scopes,omitempty"`
	Public bool     `json:"public" bson:"public" mapstructure:"public,omitempty"`
}

var (
	// Proper format a port
	portRegEx      = regexp.MustCompile("^([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535)(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))?$")
	exactPortRegEx = regexp.MustCompile("^([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535)$")
)

// PortList is a list of ports for the service.
type PortList []string

// Validate validates an input slice of ports.
func (p PortList) Validate() error {
	for _, port := range p {
		if !portRegEx.Match([]byte(port)) {
			return fmt.Errorf("Invalid port definition: %s", p)
		}
	}
	return nil
}

// ValidateExactPort will validate that the ports are exact and not ranges.
// Required for HTTP services.
func (p PortList) ValidateExactPort() error {
	for _, port := range p {
		if !exactPortRegEx.Match([]byte(port)) {
			return fmt.Errorf("Invalid port definition: %s", p)
		}
	}
	return nil
}
