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

	if !strings.Contains(ports, ",") {
		p, err := ConvertToSinglePort(ports)
		if err != nil {
			return nil, err
		}

		results = append(results, p)
		return &results, nil
	}

	cache := map[int64]struct{}{}

	sp := strings.Split(ports, ",")
	if len(sp) < 2 {
		return nil, fmt.Errorf("%s format is invalid. Should be port1,port2,port3", ports)
	}

	for _, s := range sp {

		p, err := ConvertToSinglePort(s)
		if err != nil {
			return nil, err
		}

		if _, ok := cache[p]; ok {
			return nil, fmt.Errorf("Port %s already defined", s)
		} else {
			cache[p] = struct{}{}
		}
	}

	for port := range cache {
		results = append(results, port)
	}

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

// // HasPortOverlap returns true if the port overlaps the existing ports
// func HasPortOverlap(port int64, existingPorts map[int64]struct{}) (bool, error) {

// 	if _, ok := existingPorts[port]; ok {
// 		return true, fmt.Errorf("%d port overlaps another port", port)
// 	}

// 	return false, nil
// }

// // HasPortsSlicesOverlap returns true if an overlap is found in the existing ports.
// func HasPortsSlicesOverlap(ports []int64, existingPorts map[int64]struct{}) (bool, error) {

// 	if len(existingPorts) == 0 || len(ports) == 0 {
// 		return false, nil
// 	}

// 	for _, p := range ports {
// 		if b, err := HasPortOverlap(p, existingPorts); b {
// 			return true, err
// 		}
// 	}

// 	return false, nil
// }

// // HasPortsSlicesOverlapWithPortsRanges returns true if an overlap is found in the existing ports.
// func HasPortsSlicesOverlapWithPortsRanges(ports []int64, existingRange []*PortsRange) (bool, error) {

// 	if len(existingRange) == 0 || len(ports) == 0 {
// 		return false, nil
// 	}

// 	for _, pr := range existingRange {
// 		if pr.HasOverlapWithPorts(ports) {
// 			return true, fmt.Errorf("Ports are overlapping with %d:%d", pr.FromPort, pr.ToPort)
// 		}
// 	}

// 	return false, nil
// }

// // HasOverlapInPortsRange returns true if an overlap is found in the existing ports.
// func HasOverlapInPortsRange(pr *PortsRange, existingRanges []*PortsRange) (bool, error) {

// 	if len(existingRanges) == 0 {
// 		return false, nil
// 	}

// 	for _, ex := range existingRanges {
// 		if pr.HasOverlapWithPortsRange(ex) {
// 			return true, fmt.Errorf("%d:%d overlaps range %d:%d", pr.FromPort, pr.ToPort, ex.FromPort, ex.ToPort)
// 		}

// 	}

// 	return false, nil
// }
