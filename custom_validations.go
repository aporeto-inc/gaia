package gaia

import (
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/protocols"
)

// ValidatePortString validates a string represents a port or a range of port.
// valid: 443, 443:555
func ValidatePortString(attribute string, portExp string) error {

	ports := strings.Split(portExp, ":")
	if len(ports) == 0 || len(ports) > 2 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a port (xxx) or port range (xxx:yyy)", attribute))
	}

	p1, err := strconv.Atoi(ports[0])
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a port (xxx) or port range (xxx:yyy)", attribute))
	}

	if p1 < 1 || p1 > 65535 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a between 1 and 65535", attribute))
	}

	if len(ports) == 1 {
		return nil
	}

	p2, err := strconv.Atoi(ports[1])
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a port (xxx) or port range (xxx:yyy)", attribute))
	}

	if p2 < 1 || p2 > 65535 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a between 1 and 65535", attribute))
	}

	if p1 >= p2 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' left bound is greater or equal to left bound ", attribute))
	}

	return nil
}

// ValidatePortStringList validates a list of ports.
func ValidatePortStringList(attribute string, ports []string) error {

	if len(ports) == 0 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' port list must not be empty", attribute))
	}

	for _, port := range ports {
		if err := ValidatePortString(attribute, port); err != nil {
			return err
		}
	}

	return nil
}

var rxDNSName = regexp.MustCompile(`^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`)

// ValidateNetwork validates a CIDR.
func ValidateNetwork(attribute string, network string) error {

	if _, _, err := net.ParseCIDR(network); err == nil {
		return nil
	}

	if rxDNSName.MatchString(network) {
		return nil
	}

	return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a CIDR or hostname", attribute))
}

// ValidateNetworkList validates a list of networks.
func ValidateNetworkList(attribute string, networks []string) error {

	if len(networks) == 0 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must not be empty", attribute))
	}

	for _, network := range networks {
		if err := ValidateNetwork(attribute, network); err != nil {
			return err
		}
	}

	return nil
}

// ValidateProtocol validates a string represents netwotk a protocol.
func ValidateProtocol(attribute string, proto string) error {

	upperProto := strings.ToUpper(proto)
	if upperProto == protocols.ALL || protocols.L4ProtocolNumberFromName(upperProto) != -1 {
		return nil
	}

	p, err := strconv.Atoi(proto)
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' valid protocol name or number", attribute))
	}

	if p < 0 || p > 255 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' protocol number must be between 0 and 255 included", attribute))
	}

	return nil
}

// ValidateProtocolList validates a list of protocols.
func ValidateProtocolList(attribute string, protocols []string) error {

	if len(protocols) == 0 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must not be empty", attribute))
	}

	for _, proto := range protocols {
		if err := ValidateProtocol(attribute, proto); err != nil {
			return err
		}
	}

	return nil
}

func makeValidationError(attribute string, message string) error {

	err := elemental.NewError(
		"Validation Error",
		message,
		"gaia",
		http.StatusUnprocessableEntity,
	)
	err.Data = map[string]interface{}{"attribute": attribute}

	return err
}
