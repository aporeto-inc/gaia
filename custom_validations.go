package gaia

import (
	"fmt"
	"net"
	"net/http"

	"go.aporeto.io/elemental"
)

func makeValidationError(attribute string, message string) error {

	err := elemental.NewError(
		"Validation Error",
		fmt.Sprintf(message, attribute),
		"gaia",
		http.StatusUnprocessableEntity,
	)
	err.Data = map[string]interface{}{"attribute": attribute}

	return err
}

// ValidateCIDR validates a CIDR.
func ValidateCIDR(attribute string, cidr string) error {

	if _, _, err := net.ParseCIDR(cidr); err != nil {
		return makeValidationError(attribute, "Attribute '%s' must be a CIDR")
	}

	return nil
}

// ValidateCIDRList validates a list of CIDR.
func ValidateCIDRList(attribute string, cidrs []string) error {

	for _, cidr := range cidrs {
		if err := ValidateCIDR(attribute, cidr); err != nil {
			return err
		}
	}

	return nil
}
