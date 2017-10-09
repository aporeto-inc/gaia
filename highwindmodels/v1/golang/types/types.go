package types

import "go.uber.org/zap"

// ServiceDescription describe the service.
type ServiceDescription struct {
	Parameters            ServiceParameters
	Image                 string
	Replicas              int32
	Labels                map[string]string
	UsesClientCertificate bool
}

// ServiceParameters is a map of ServiceParameter.
type ServiceParameters map[string]*ServiceParameter

// SetParameter sets a new value for the key.
func (p ServiceParameters) SetParameter(key string, value string) {

	param, ok := p[key]

	if !ok {
		zap.L().Error("Unable to set parameter value. Key not found", zap.String("key", key))
		return
	}

	param.Value = value
}

// GetParameter sets a new value for the key.
func (p ServiceParameters) GetParameter(key string) string {

	param, ok := p[key]

	if !ok {
		zap.L().Error("Unable to set parameter value. Key not found", zap.String("key", key))
		return ""
	}

	return param.Value
}

// ServiceParameter defines a parameter for the service.
type ServiceParameter struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Env       string `json:"-"`
	Secret    bool   `json:"secret"`
	MountPath string `json:"-"`
}

// NewServiceParameter creates a new parameter.
func NewServiceParameter() *ServiceParameter {

	return &ServiceParameter{}
}

// Copy returns a copy of the current paramter.
func (p *ServiceParameter) Copy() *ServiceParameter {

	copy := NewServiceParameter()
	copy.Key = p.Key
	copy.Value = p.Value
	copy.Env = p.Env
	copy.Secret = p.Secret
	copy.MountPath = p.MountPath

	return copy
}
