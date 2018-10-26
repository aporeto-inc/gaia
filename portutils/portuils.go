package portutils

import (
	"fmt"
	"strconv"
	"strings"
)

// ConvertToPortsRange converts a :-separated string to a min and max port.
func ConvertToPortsRange(ports string) (*PortsRange, error) {

	sp := strings.Split(ports, ":")

	if len(sp) != 2 {
		return nil, fmt.Errorf("%s is not a valid range format. It should be of form fromPort:toPort", ports)
	}

	fromPort, err := ConvertToSinglePort(sp[0])
	if err != nil {
		return nil, err
	}

	toPort, err := ConvertToSinglePort(sp[1])
	if err != nil {
		return nil, err
	}

	if fromPort >= toPort {
		return nil, fmt.Errorf("left bound is greater or equal to right bound")
	}

	return &PortsRange{fromPort, toPort}, nil
}

// ConvertToPortsList convert a , separated string to a list of port.
func ConvertToPortsList(ports string) (*PortsList, error) {

	results := PortsList{}

	p, err := ConvertToSinglePort(ports)
	if err != nil {
		return nil, err
	}

	results = append(results, p)
	return &results, nil

}

// ConvertToSinglePort converts a string to port.
func ConvertToSinglePort(port string) (int64, error) {

	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("%s is not a valid port", port)
	}

	if p < 1 || p > 65535 {
		return -1, fmt.Errorf("%s is not in between 1 and 65535", port)
	}

	return p, nil
}
