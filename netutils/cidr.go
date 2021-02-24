package netutils

import (
	"fmt"
	"net"
	"strings"

	"github.com/yl2chen/cidranger"
)

var (
	opInclude         = "include" // default operation to include cidr in networks
	opExclude         = "exclude" // not include cidr when cidr string has the ! prefix
	multicastv4Subnet = "224.0.0.0/4"
	multicastv6Subnet = "ff00::/8"
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

		// we only are interested if the excluded pfx is present in the included.
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
		if err := ranger.Insert(newCustomRangerEntry(cidr)); err != nil {
			return fmt.Errorf("Error adding CIDR %s", cidr.str)
		}
	}

	// Parse and validate all not CIDRs are included in regular CIDRs
	for _, c := range cidrmap {

		if c.op == opExclude {
			// get all the networks which are containing the network,
			// basically getting all the parent prefixes.
			entries, err := ranger.ContainingNetworks(c.ipNet.IP)
			if err != nil {
				return fmt.Errorf("Cannot find the CIDR: %s", err)
			}

			mask := c.ipNet.Mask
			// make sure the NOT(!) CIDR is part/contained of/by a included CIDR
			present := checkExcPfxContainedInc(entries, mask, c.ipNet)

			// if the excluded CIDR is not contained in the included CIDR then return error
			if !present {
				return fmt.Errorf("%s is not contained in any CIDR", c.str)
			}

			// also make sure there are no included CIDRs in the excluded CIDRs
			// "CoveredNetworks" basically gets all the networks under the particular IP,
			// basically all the children networks.
			entries, err = ranger.CoveredNetworks(c.ipNet)
			if err != nil {
				return fmt.Errorf("Cannot find the CIDR: %s", err)
			}
			// now check if there are any children networks that don't have NOT(!),
			// basically, check if we have any included networks, if yes return error.
			ip, present := checkIncPfxContainedInExc(entries, c.ipNet)
			if present {
				return fmt.Errorf("%s is contained in excluded CIDR %s", ip, c.Network())
			}
		}
	}

	// now if all the CIDR make sense, check for multicast subnet,
	//  check if we have a 224/4 subnet in the pfx tree,
	// if yes then return error

	_, v4network, err := net.ParseCIDR(multicastv4Subnet)
	if err != nil {
		return fmt.Errorf("%s is not a valid CIDR", v4network)
	}
	_, v6network, err := net.ParseCIDR(multicastv6Subnet)
	if err != nil {
		return fmt.Errorf("%s is not a valid CIDR", v6network)
	}
	// get all the networks that have 224/4 and if anyone has a NOT(!) then we are good.
	// else we report ERROR.
	multicastv4Entries, _ := ranger.ContainingNetworks(v4network.IP)
	multicastv6Entries, _ := ranger.ContainingNetworks(v6network.IP)

	multicastv4SubnetExc := false
	wrongEntryv4CIDR := make([]string, 0, len(multicastv4Entries))

	for _, entry := range multicastv4Entries {
		cidr := entry.(*cidr)
		if cidr.op == opExclude {
			multicastv4SubnetExc = true
			break
		}
		wrongEntryv4CIDR = append(wrongEntryv4CIDR, cidr.str)
	}
	if len(multicastv4Entries) > 0 && !multicastv4SubnetExc {
		return fmt.Errorf("The CIDR %s contains the multicast subnet", wrongEntryv4CIDR)
	}

	multicastv6SubnetExc := false
	wrongEntryv6CIDR := make([]string, 0, len(multicastv6Entries))

	for _, entry := range multicastv6Entries {
		cidr := entry.(*cidr)
		if cidr.op == opExclude {
			multicastv6SubnetExc = true
			break
		}
		wrongEntryv6CIDR = append(wrongEntryv6CIDR, cidr.str)
	}
	if len(multicastv6Entries) > 0 && !multicastv6SubnetExc {
		return fmt.Errorf("The CIDR %s contains the multicast subnet", wrongEntryv6CIDR)
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
		if err := ranger.Insert(newCustomRangerEntry(cidr)); err != nil {
			return fmt.Errorf("Error adding CIDR %s", cidr.str)
		}
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
