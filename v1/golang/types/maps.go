package types

import (
	"sync"
)

// GraphEdgeMap is a map of id to GraphEdge
type GraphEdgeMap map[string]*GraphEdge

// GraphEdgeExtremityType represents the source or destination type.
type GraphEdgeExtremityType string

const (
	// GraphEdgeExtremityTypePU represents the value of a pu extremity.
	GraphEdgeExtremityTypePU GraphEdgeExtremityType = "pu"

	// GraphEdgeExtremityTypeExternalService represents the value of a pu external service.
	GraphEdgeExtremityTypeExternalService GraphEdgeExtremityType = "ext"
)

// GraphEdge represents the model of a Edge
type GraphEdge struct {
	AcceptedFlows         int                    `json:"acceptedFlows"`
	DestinationID         string                 `json:"destinationID"`
	DestinationType       GraphEdgeExtremityType `json:"destinationType"`
	DestinationPort       string                 `json:"destinationPort"`
	ID                    string                 `json:"ID"`
	Name                  string                 `json:"name"`
	RejectedFlows         int                    `json:"rejectedFlows"`
	SourceID              string                 `json:"sourceID"`
	SourceType            GraphEdgeExtremityType `json:"sourceType"`
	PolicyIDs             map[string]int         `json:"policyIDs"`
	Encrypted             int                    `json:"encrypted"`
	ObservedAcceptedFlows int                    `json:"observedAcceptedFlows"`
	ObservedRejectedFlows int                    `json:"observedRejectedFlows"`
	ObservedPolicyIDs     map[string]int         `json:"observedPolicyIDs"`
	ObservedEncrypted     int                    `json:"observedEncrypted"`
}

// NewGraphEdge returns a new *GraphEdge
func NewGraphEdge() *GraphEdge {
	return &GraphEdge{}
}

// GraphNodeType represents the possible values for attribute "type".
type GraphNodeType string

const (
	// GraphNodeTypeContainer represents the value Container.
	GraphNodeTypeContainer GraphNodeType = "Container"

	// GraphNodeTypeExternalService represents the value EternalService.
	GraphNodeTypeExternalService GraphNodeType = "ExternalService"

	// GraphNodeTypeVolume represents the value Volume.
	GraphNodeTypeVolume GraphNodeType = "Volume"
)

// GraphNodeMap is a map of id to GraphNode
type GraphNodeMap map[string]*GraphNode

// GraphNode represents the model of a Node
type GraphNode struct {
	Description        string        `json:"description"`
	GroupID            string        `json:"groupID"`
	ID                 string        `json:"ID"`
	Name               string        `json:"name"`
	Namespace          string        `json:"namespace"`
	Status             string        `json:"status"`
	Tags               []string      `json:"tags,omitempty"`
	Type               GraphNodeType `json:"type"`
	VulnerabilityLevel string        `json:"vulnerabilityLevel"`
}

// NewGraphNode returns a new *GraphNode
func NewGraphNode() *GraphNode {

	return &GraphNode{
		Tags: []string{},
	}
}

// GraphGroupMap is a map of id to GraphGroup
type GraphGroupMap map[string]*GraphGroup

// GraphGroup represents the model of a Group
type GraphGroup struct {
	Color    string     `json:"color"`
	ID       string     `json:"ID"`
	Match    [][]string `json:"match"`
	Name     string     `json:"name"`
	ParentID string     `json:"parentID"`
}

// NewGraphGroup returns a new *GraphNode
func NewGraphGroup() *GraphGroup {

	return &GraphGroup{
		Match: [][]string{},
	}
}

// l4Protocols is te list of all IANA protocols
// (Ref https://en.wikipedia.org/wiki/List_of_IP_protocol_numbers)
var l4Protocols = make([]string, 150)

func init() {
	l4Protocols[0] = "HOPOPT"
	l4Protocols[1] = "ICMP"
	l4Protocols[2] = "IGMP"
	l4Protocols[3] = "GGP"
	l4Protocols[4] = "IP-in-IP"
	l4Protocols[5] = "ST"
	l4Protocols[6] = "TCP"
	l4Protocols[7] = "CBT"
	l4Protocols[8] = "EGP"
	l4Protocols[9] = "IGP"
	l4Protocols[10] = "BBN-RCC-MON"
	l4Protocols[11] = "NVP-II"
	l4Protocols[12] = "PUP"
	l4Protocols[13] = "ARGUS"
	l4Protocols[14] = "EMCON"
	l4Protocols[15] = "XNET"
	l4Protocols[16] = "CHAOS"
	l4Protocols[17] = "UDP"
	l4Protocols[18] = "MUX"
	l4Protocols[19] = "DCN-MEAS"
	l4Protocols[20] = "HMP"
	l4Protocols[21] = "PRM"
	l4Protocols[22] = "XNS-IDP"
	l4Protocols[23] = "TRUNK-1"
	l4Protocols[24] = "TRUNK-2"
	l4Protocols[25] = "LEAF-1"
	l4Protocols[26] = "LEAF-2"
	l4Protocols[27] = "RDP"
	l4Protocols[28] = "IRTP"
	l4Protocols[29] = "ISO-TP4"
	l4Protocols[30] = "NETBLT"
	l4Protocols[31] = "MFE-NSP"
	l4Protocols[32] = "MERIT-INP"
	l4Protocols[33] = "DCCP"
	l4Protocols[34] = "3PC"
	l4Protocols[35] = "IDPR"
	l4Protocols[36] = "XTP"
	l4Protocols[37] = "DDP"
	l4Protocols[38] = "IDPR-CMTP"
	l4Protocols[39] = "TP++"
	l4Protocols[40] = "IL"
	l4Protocols[41] = "IPv6"
	l4Protocols[42] = "SDRP"
	l4Protocols[43] = "IPv6-Route"
	l4Protocols[44] = "IPv6-Frag"
	l4Protocols[45] = "IDRP"
	l4Protocols[46] = "RSVP"
	l4Protocols[47] = "GREs"
	l4Protocols[48] = "DSR"
	l4Protocols[49] = "BNA"
	l4Protocols[50] = "ESP"
	l4Protocols[51] = "AH"
	l4Protocols[52] = "I-NLSP"
	l4Protocols[53] = "SWIPE"
	l4Protocols[54] = "NARP"
	l4Protocols[55] = "MOBILE"
	l4Protocols[56] = "TLSP"
	l4Protocols[57] = "SKIP"
	l4Protocols[58] = "IPv6-ICMP"
	l4Protocols[59] = "IPv6-NoNxt"
	l4Protocols[60] = "IPv6-Opts"
	l4Protocols[62] = "CFTP"
	l4Protocols[64] = "SAT-EXPAK"
	l4Protocols[65] = "KRYPTOLAN"
	l4Protocols[66] = "RVD"
	l4Protocols[67] = "IPPC"
	l4Protocols[69] = "SAT-MON"
	l4Protocols[70] = "VISA"
	l4Protocols[71] = "IPCU"
	l4Protocols[72] = "CPNX"
	l4Protocols[73] = "CPHB"
	l4Protocols[74] = "WSN"
	l4Protocols[75] = "PVP"
	l4Protocols[76] = "BR-SAT-MON"
	l4Protocols[77] = "SUN-ND"
	l4Protocols[78] = "WB-MON"
	l4Protocols[79] = "WB-EXPAK"
	l4Protocols[80] = "ISO-IP"
	l4Protocols[81] = "VMTP"
	l4Protocols[82] = "SECURE-VMTP"
	l4Protocols[83] = "VINES"
	l4Protocols[84] = "TTP"
	l4Protocols[85] = "NSFNET-IGP"
	l4Protocols[86] = "DGP"
	l4Protocols[87] = "TCF"
	l4Protocols[88] = "EIGRP"
	l4Protocols[89] = "OSPF"
	l4Protocols[90] = "Sprite-RPC"
	l4Protocols[91] = "LARP"
	l4Protocols[92] = "MTP"
	l4Protocols[93] = "AX.25"
	l4Protocols[94] = "OS"
	l4Protocols[95] = "MICP"
	l4Protocols[96] = "SCC-SP"
	l4Protocols[97] = "ETHERIP"
	l4Protocols[98] = "ENCAP"
	l4Protocols[100] = "GMTP"
	l4Protocols[101] = "IFMP"
	l4Protocols[102] = "PNNI"
	l4Protocols[103] = "PIM"
	l4Protocols[104] = "ARIS"
	l4Protocols[105] = "SCPS"
	l4Protocols[106] = "QNX"
	l4Protocols[107] = "A/N"
	l4Protocols[108] = "IPComp"
	l4Protocols[109] = "SNP"
	l4Protocols[110] = "Compaq-Peer"
	l4Protocols[111] = "IPX-in-IP"
	l4Protocols[112] = "VRRP"
	l4Protocols[113] = "PGM"
	l4Protocols[115] = "L2TP"
	l4Protocols[116] = "DDX"
	l4Protocols[117] = "IATP"
	l4Protocols[118] = "STP"
	l4Protocols[119] = "SRP"
	l4Protocols[120] = "UTI"
	l4Protocols[121] = "SMP"
	l4Protocols[122] = "SM"
	l4Protocols[123] = "PTP"
	l4Protocols[124] = "IS-IS over IPv4"
	l4Protocols[125] = "FIRE"
	l4Protocols[126] = "CRTP"
	l4Protocols[127] = "CRUDP"
	l4Protocols[128] = "SSCOPMCE"
	l4Protocols[129] = "IPLT"
	l4Protocols[130] = "SPS"
	l4Protocols[131] = "PIPE"
	l4Protocols[132] = "SCTP"
	l4Protocols[133] = "FC"
	l4Protocols[134] = "RSVP-E2E-IGNORE"
	l4Protocols[135] = "Mobility Header"
	l4Protocols[136] = "UDPLite"
	l4Protocols[137] = "MPLS-in-IP"
	l4Protocols[138] = "manet"
	l4Protocols[139] = "HIP"
	l4Protocols[140] = "Shim6"
	l4Protocols[141] = "WESP"
	l4Protocols[142] = "ROHC"
}

// ProtocolName returns the IANA for the protocol
func ProtocolName(n int64) string {

	l := int64(len(l4Protocols))
	if n < 0 || n >= l {
		return "N/A"
	}

	if protocol := l4Protocols[n]; protocol != "" {
		return protocol
	}

	return "N/A"
}

// IPRecord represent an IP record.
type IPRecord struct {
	Actions          []string `json:"actions"`
	IP               string   `json:"IP"`
	Hostnames        []string `json:"hostnames"`
	DestinationPorts []string `json:"destinationPorts"`
	Hits             int      `json:"hits"`
	Latitude         float32  `json:"latitude"`
	Longitude        float32  `json:"longitude"`
	Country          string   `json:"country"`
	City             string   `json:"city"`
	L4Protocol       string   `json:"l4Protocol"`

	sync.Mutex `json:"-"`
}

// NewIPRecord returns a new IPRecord.
func NewIPRecord() *IPRecord {
	return &IPRecord{
		Actions:          []string{},
		Hostnames:        []string{},
		DestinationPorts: []string{},
	}
}

// >>> TODO DELETE WHEN MERGED

// TagGraphStats represents Tag statistics in a Graph
type TagGraphStats struct {
	Key         string `json:"key"`
	ValuesCount int    `json:"valuesCount"`
	Occurrences int    `json:"Occurrences"`
}

// NewTagGraphStats creates a new NewTagGraphStats
func NewTagGraphStats() *TagGraphStats {
	return &TagGraphStats{}
}

// IsEqual returns true if both TagGraphStats are equal
func (a *TagGraphStats) IsEqual(b *TagGraphStats) bool {
	return a.Key == b.Key && a.ValuesCount == b.ValuesCount && a.Occurrences == b.Occurrences
}

// TagGraphStatsList represents a list of TagGraphStats
type TagGraphStatsList []*TagGraphStats

// Len is the implenimplementationtation for sort.Interface
func (a TagGraphStatsList) Len() int {
	return len(a)
}

// Swap is the implementation for sort.Interface
func (a TagGraphStatsList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less is the implementation for sort.Interface
func (a TagGraphStatsList) Less(i, j int) bool {
	if a[i].ValuesCount > a[j].ValuesCount {
		return true
	}

	if a[i].ValuesCount == a[j].ValuesCount && a[i].Occurrences > a[j].Occurrences {
		return true
	}

	if a[i].Occurrences == a[j].Occurrences {
		return a[i].Key < a[j].Key
	}

	return false
}

// <<< TODO DELETE WHEN MERGED
