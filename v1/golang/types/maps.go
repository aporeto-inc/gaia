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
var l4Protocols = map[int]string{
	0:   "HOPOPT",
	1:   "ICMP",
	2:   "IGMP",
	3:   "GGP",
	4:   "IP-in-IP",
	5:   "ST",
	6:   "TCP",
	7:   "CBT",
	8:   "EGP",
	9:   "IGP",
	10:  "BBN-RCC-MON",
	11:  "NVP-II",
	12:  "PUP",
	13:  "ARGUS",
	14:  "EMCON",
	15:  "XNET",
	16:  "CHAOS",
	17:  "UDP",
	18:  "MUX",
	19:  "DCN-MEAS",
	20:  "HMP",
	21:  "PRM",
	22:  "XNS-IDP",
	23:  "TRUNK-1",
	24:  "TRUNK-2",
	25:  "LEAF-1",
	26:  "LEAF-2",
	27:  "RDP",
	28:  "IRTP",
	29:  "ISO-TP4",
	30:  "NETBLT",
	31:  "MFE-NSP",
	32:  "MERIT-INP",
	33:  "DCCP",
	34:  "3PC",
	35:  "IDPR",
	36:  "XTP",
	37:  "DDP",
	38:  "IDPR-CMTP",
	39:  "TP++",
	40:  "IL",
	41:  "IPv6",
	42:  "SDRP",
	43:  "IPv6-Route",
	44:  "IPv6-Frag",
	45:  "IDRP",
	46:  "RSVP",
	47:  "GREs",
	48:  "DSR",
	49:  "BNA",
	50:  "ESP",
	51:  "AH",
	52:  "I-NLSP",
	53:  "SWIPE",
	54:  "NARP",
	55:  "MOBILE",
	56:  "TLSP",
	57:  "SKIP",
	58:  "IPv6-ICMP",
	59:  "IPv6-NoNxt",
	60:  "IPv6-Opts",
	62:  "CFTP",
	64:  "SAT-EXPAK",
	65:  "KRYPTOLAN",
	66:  "RVD",
	67:  "IPPC",
	69:  "SAT-MON",
	70:  "VISA",
	71:  "IPCU",
	72:  "CPNX",
	73:  "CPHB",
	74:  "WSN",
	75:  "PVP",
	76:  "BR-SAT-MON",
	77:  "SUN-ND",
	78:  "WB-MON",
	79:  "WB-EXPAK",
	80:  "ISO-IP",
	81:  "VMTP",
	82:  "SECURE-VMTP",
	83:  "VINES",
	84:  "TTP",
	85:  "NSFNET-IGP",
	86:  "DGP",
	87:  "TCF",
	88:  "EIGRP",
	89:  "OSPF",
	90:  "Sprite-RPC",
	91:  "LARP",
	92:  "MTP",
	93:  "AX.25",
	94:  "OS",
	95:  "MICP",
	96:  "SCC-SP",
	97:  "ETHERIP",
	98:  "ENCAP",
	100: "GMTP",
	101: "IFMP",
	102: "PNNI",
	103: "PIM",
	104: "ARIS",
	105: "SCPS",
	106: "QNX",
	107: "A/N",
	108: "IPComp",
	109: "SNP",
	110: "Compaq-Peer",
	111: "IPX-in-IP",
	112: "VRRP",
	113: "PGM",
	115: "L2TP",
	116: "DDX",
	117: "IATP",
	118: "STP",
	119: "SRP",
	120: "UTI",
	121: "SMP",
	122: "SM",
	123: "PTP",
	124: "IS-IS over IPv4",
	125: "FIRE",
	126: "CRTP",
	127: "CRUDP",
	128: "SSCOPMCE",
	129: "IPLT",
	130: "SPS",
	131: "PIPE",
	132: "SCTP",
	133: "FC",
	134: "RSVP-E2E-IGNORE",
	135: "Mobility Header",
	136: "UDPLite",
	137: "MPLS-in-IP",
	138: "manet",
	139: "HIP",
	140: "Shim6",
	141: "WESP",
	142: "ROHC",
}

// ProtocolName returns the IANA for the protocol
func ProtocolName(n int) string {

	if protocol, ok := l4Protocols[n]; ok {
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
