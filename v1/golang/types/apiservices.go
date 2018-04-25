package types

import (
	"fmt"
	"strings"
)

// IPList is a list of IP addresses
type IPList []IPAddress

// IPAddress is an IP or subnet expressed as a string
type IPAddress string

// ExposedAPIList is a list of API endpoints and associated tags.
type ExposedAPIList []ExposedAPI

var (
	allowedMethods = map[string]bool{
		"GET":    true,
		"POST":   true,
		"PATCH":  true,
		"DELETE": true,
		"HEAD":   true,
		"PUT":    true,
	}
)

// ExposedAPI is an exposed API defined by the URI, verb, and associated tags.
// The URIs must be valid Golang regular expressions.
type ExposedAPI struct {
	URI     string   `json:"URI" bson:"URI" mapstructure:"URI,omitempty"`
	Methods []string `json:"methods" bson:"methods" mapstructure:"methods,omitempty"`
	Scopes  []string `json:"scopes" bson:"scopes" mapstructure:"scopes,omitempty"`
	Public  bool     `json:"public" bson:"public" mapstructure:"public,omitempty"`
}

// Validate validates the ExposedAPI structure
func (e *ExposedAPI) Validate() error {
	for i, m := range e.Methods {
		method := strings.ToUpper(m)
		if _, ok := allowedMethods[method]; !ok {
			return fmt.Errorf("Invalid method %s", m)
		}
		e.Methods[i] = method
	}
	return nil
}
