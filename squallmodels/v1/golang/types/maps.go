package types

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
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
	AcceptedFlows   int                    `json:"acceptedFlows"`
	DestinationID   string                 `json:"destinationID"`
	DestinationType GraphEdgeExtremityType `json:"destinationType"`
	ID              string                 `json:"ID"`
	Name            string                 `json:"name"`
	RejectedFlows   int                    `json:"rejectedFlows"`
	SourceID        string                 `json:"sourceID"`
	SourceType      GraphEdgeExtremityType `json:"sourceType"`
	PolicyID        string                 `json:"policyID"`
	Encrypted       int                    `json:"encrypted"`
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

// ResolveHostnames resolves hostnames for the current IP.
func (r *IPRecord) ResolveHostnames() error {

	var err error

	r.Lock()
	r.Hostnames, err = net.LookupAddr(r.IP)
	r.Unlock()

	return err
}

// ResolveLocation tries to geolocate the current IP.
func (r *IPRecord) ResolveLocation(geoipurl string) error {

	r.Lock()
	ip := r.IP
	r.Unlock()

	resp, err := http.Get(geoipurl + "/json/" + ip)
	if err != nil {
		return fmt.Errorf("Unable to retrieve info from geoip service: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unable to retrieve info from geoip service: %s", resp.Status)
	}

	s := struct {
		Latitude  float32
		Longitude float32
		Country   string
		City      string
	}{}

	defer resp.Body.Close() // nolint: errcheck
	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return fmt.Errorf("Unable to decode info from geoip service: %s", err.Error())
	}

	r.Lock()
	r.Latitude = s.Latitude
	r.Longitude = s.Longitude
	r.Country = s.Country
	r.City = s.City
	r.Unlock()

	return nil
}
