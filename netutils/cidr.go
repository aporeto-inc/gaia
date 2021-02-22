package netutils

import (
	"fmt"
	"net"
	"strings"

	"github.com/yl2chen/cidranger"
)

var (
	opInclude       = "include" // default operation to include cidr in networks
	opExclude       = "exclude" // not include cidr when cidr string has the ! prefix
	multicastSubnet = "224.0.0.0/4"
)

// cidr represents the cidr network to be included or excluded
type cidr struct {
	// op is the opreation to perform with the cidr: include or exclude
	op string

	// ipNet is the IPNet object that contains ip and mask of the cidr
	ipNet net.IPNet

	// str is the original cidr string
	str string
}

func checkExcPfxContainedInc(entries []cidranger.RangerEntry, mask net.IPMask, ip net.IPNet) bool {

	for _, e := range entries {
		cidr := e.(*cidr)

		// we only are interested if the exculded pfx is present in the included.
		if cidr.op == opExclude {
			continue
		}

		includedOnes, size1 := e.Network().Mask.Size() //  included CIDR subnet
		excludedOnes, size2 := mask.Size()             // exclude CIDR subnet, basically the ip under check.

		if size1 != size2 {
			continue
		}
		// basically, check here is that the,
		// excluded subnet should always be greater than included, if its less, then return ERROR/FALSE.
		if includedOnes <= excludedOnes {
			return true
		}

	}
	return false
}

// checkIncPfxContainedInEx checks if there are any included pfxs in the excluded pfx
func checkIncPfxContainedInExc(entries []cidranger.RangerEntry, ip net.IPNet) (net.IPNet, bool) {
	for _, e := range entries {
		cidr := e.(*cidr)
		if cidr.op == opInclude {
			return cidr.ipNet, true
		}
	}
	return net.IPNet{}, false
}

// parseCIDR converts the given string to cidr. Returns an error if it wasnt able to parse a CIDR
func parseCIDR(s string) (*cidr, error) {

	c := &cidr{op: opInclude, str: s}
	_, network, err := net.ParseCIDR(strings.TrimPrefix(s, "!"))
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid CIDR", s)
	}
	c.ipNet = *network
	if strings.HasPrefix(s, "!") {
		c.op = opExclude
	}

	return c, nil
}

// get function for network
func (b *cidr) Network() net.IPNet {
	return b.ipNet
}

// create customRangerEntry object using net and cidr
func newCustomRangerEntry(c *cidr) cidranger.RangerEntry {
	return c
}

// ValidateUDPCIDRs validates that the list of string provided as a set is a valid CIDR set
func ValidateUDPCIDRs(ss []string) error {

	cidrmap := make(map[string]*cidr)
	// instantiate NewPCTrieRanger
	ranger := cidranger.NewPCTrieRanger()

	for _, s := range ss {
		cidr, err := parseCIDR(s)
		if err != nil {
			return err
		}
		if _, ok := cidrmap[cidr.ipNet.String()]; !ok {
			cidrmap[cidr.ipNet.String()] = cidr
		} else {
			return fmt.Errorf("CIDR subnet parsed from %s is duplicated", cidr.str)
		}
		ranger.Insert(newCustomRangerEntry(cidr))
	}

	// Parse and validate all not CIDRs are included in regular CIDRs
	for _, c := range cidrmap {

		if c.op == opExclude {

			entries, err := ranger.ContainingNetworks(c.ipNet.IP)
			if err != nil {
				return fmt.Errorf("Cannot find the CIDR: %s", err)
			}

			mask := c.ipNet.Mask
			present := checkExcPfxContainedInc(entries, mask, c.ipNet)

			// if the excluded CIDR is not contained in the included CIDR then return error
			if !present {
				return fmt.Errorf("%s is not contained in any CIDR", c.str)
			}

			// also make sure there are no included CIDRs in the excluded CIDRs
			entries, err = ranger.CoveredNetworks(c.ipNet)
			if err != nil {
				return fmt.Errorf("Cannot find the CIDR: %s", err)
			}

			ip, present := checkIncPfxContainedInExc(entries, c.ipNet)
			if present {
				return fmt.Errorf("%s is contained in excluded CIDR %s", ip, c.Network())
			}
		}
	}

	// now if all the CIDR make sense, check for multicast subnet,
	//  check if we have a 224/4 subnet in the pfx tree,
	// if yes then return error

	_, network, err := net.ParseCIDR(multicastSubnet)
	if err != nil {
		return fmt.Errorf("%s is not a valid CIDR", network)
	}
	multicastEntries, err := ranger.ContainingNetworks(network.IP)
	for _, entry := range multicastEntries {
		cidr := entry.(*cidr)
		if cidr.op == opExclude {
			continue
		}
		return fmt.Errorf("The CIDR %s contains the multicast subnet", cidr.str)
	}

	return nil
}

// ValidateCIDRs validates that the list of string provided as a set is a valid CIDR set
func ValidateCIDRs(ss []string) error {

	cidrmap := make(map[string]*cidr)
	// instantiate NewPCTrieRanger
	ranger := cidranger.NewPCTrieRanger()

	for _, s := range ss {
		cidr, err := parseCIDR(s)
		if err != nil {
			return err
		}
		if _, ok := cidrmap[cidr.ipNet.String()]; !ok {
			cidrmap[cidr.ipNet.String()] = cidr
		} else {
			return fmt.Errorf("CIDR subnet parsed from %s is duplicated", cidr.str)
		}
		ranger.Insert(newCustomRangerEntry(cidr))
	}

	// Parse and validate all not CIDRs are included in regular CIDRs
	for _, c := range cidrmap {

		if c.op == opExclude {

			entries, err := ranger.ContainingNetworks(c.ipNet.IP)
			if err != nil {
				return fmt.Errorf("Cannot find the CIDR: %s", err)
			}

			mask := c.ipNet.Mask
			present := checkExcPfxContainedInc(entries, mask, c.ipNet)

			// if the excluded CIDR is not contained in the included CIDR then return error
			if !present {
				return fmt.Errorf("%s is not contained in any CIDR", c.str)
			}

			// also make sure there are no included CIDRs in the excluded CIDRs
			entries, err = ranger.CoveredNetworks(c.ipNet)
			if err != nil {
				return fmt.Errorf("Cannot find the CIDR: %s", err)
			}

			ip, present := checkIncPfxContainedInExc(entries, c.ipNet)
			if present {
				return fmt.Errorf("%s is contained in excluded CIDR %s", ip, c.Network())
			}
		}
	}

	return nil
}
