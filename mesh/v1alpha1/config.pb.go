// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/config.proto

/*
Package v1alpha1 is a generated protocol buffer package.

It is generated from these files:
	mesh/v1alpha1/config.proto

It has these top-level messages:
	ProxyConfig
	MeshConfig
*/
package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AuthenticationPolicy defines authentication policy. It can be set for
// different scopes (mesh, service …), and the most narrow scope with
// non-INHERIT value will be used.
// Mesh policy cannot be INHERIT.
type AuthenticationPolicy int32

const (
	// Do not encrypt Envoy to Envoy traffic.
	AuthenticationPolicy_NONE AuthenticationPolicy = 0
	// Envoy to Envoy traffic is wrapped into mutual TLS connections.
	AuthenticationPolicy_MUTUAL_TLS AuthenticationPolicy = 1
	// Use the policy defined by the parent scope. Should not be used for mesh
	// policy.
	AuthenticationPolicy_INHERIT AuthenticationPolicy = 1000
)

var AuthenticationPolicy_name = map[int32]string{
	0:    "NONE",
	1:    "MUTUAL_TLS",
	1000: "INHERIT",
}
var AuthenticationPolicy_value = map[string]int32{
	"NONE":       0,
	"MUTUAL_TLS": 1,
	"INHERIT":    1000,
}

func (x AuthenticationPolicy) String() string {
	return proto.EnumName(AuthenticationPolicy_name, int32(x))
}
func (AuthenticationPolicy) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MeshConfig_IngressControllerMode int32

const (
	// Disables Istio ingress controller.
	MeshConfig_OFF MeshConfig_IngressControllerMode = 0
	// Istio ingress controller will act on ingress resources that do not
	// contain any annotation or whose annotations match the value
	// specified in the ingress_class parameter described earlier. Use this
	// mode if Istio ingress controller will be the default ingress
	// controller for the entire kubernetes cluster.
	MeshConfig_DEFAULT MeshConfig_IngressControllerMode = 1
	// Istio ingress controller will only act on ingress resources whose
	// annotations match the value specified in the ingress_class parameter
	// described earlier. Use this mode if Istio ingress controller will be
	// a secondary ingress controller (e.g., in addition to a
	// cloud-provided ingress controller).
	MeshConfig_STRICT MeshConfig_IngressControllerMode = 2
)

var MeshConfig_IngressControllerMode_name = map[int32]string{
	0: "OFF",
	1: "DEFAULT",
	2: "STRICT",
}
var MeshConfig_IngressControllerMode_value = map[string]int32{
	"OFF":     0,
	"DEFAULT": 1,
	"STRICT":  2,
}

func (x MeshConfig_IngressControllerMode) String() string {
	return proto.EnumName(MeshConfig_IngressControllerMode_name, int32(x))
}
func (MeshConfig_IngressControllerMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

// TODO AuthPolicy needs to be removed and merged with AuthPolicy defined above
type MeshConfig_AuthPolicy int32

const (
	// Do not encrypt Envoy to Envoy traffic.
	MeshConfig_NONE MeshConfig_AuthPolicy = 0
	// Envoy to Envoy traffic is wrapped into mutual TLS connections.
	MeshConfig_MUTUAL_TLS MeshConfig_AuthPolicy = 1
)

var MeshConfig_AuthPolicy_name = map[int32]string{
	0: "NONE",
	1: "MUTUAL_TLS",
}
var MeshConfig_AuthPolicy_value = map[string]int32{
	"NONE":       0,
	"MUTUAL_TLS": 1,
}

func (x MeshConfig_AuthPolicy) String() string {
	return proto.EnumName(MeshConfig_AuthPolicy_name, int32(x))
}
func (MeshConfig_AuthPolicy) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

// ProxyConfig defines variables for individual Envoy instances.
type ProxyConfig struct {
	// Path to the generated configuration file directory.
	// Proxy agent generates the actual configuration and stores it in this directory.
	ConfigPath string `protobuf:"bytes,1,opt,name=config_path,json=configPath" json:"config_path,omitempty"`
	// Path to the proxy binary
	BinaryPath string `protobuf:"bytes,2,opt,name=binary_path,json=binaryPath" json:"binary_path,omitempty"`
	// Service cluster defines the name for the service_cluster that is
	// shared by all Envoy instances. This setting corresponds to
	// _--service-cluster_ flag in Envoy.  In a typical Envoy deployment, the
	// _service-cluster_ flag is used to identify the caller, for
	// source-based routing scenarios.
	//
	// Since Istio does not assign a local service/service version to each
	// Envoy instance, the name is same for all of them.  However, the
	// source/caller's identity (e.g., IP address) is encoded in the
	// _--service-node_ flag when launching Envoy.  When the RDS service
	// receives API calls from Envoy, it uses the value of the _service-node_
	// flag to compute routes that are relative to the service instances
	// located at that IP address.
	ServiceCluster string `protobuf:"bytes,3,opt,name=service_cluster,json=serviceCluster" json:"service_cluster,omitempty"`
	// The time in seconds that Envoy will drain connections during a hot
	// restart. MUST be >=1s (e.g., _1s/1m/1h_)
	DrainDuration *google_protobuf.Duration `protobuf:"bytes,4,opt,name=drain_duration,json=drainDuration" json:"drain_duration,omitempty"`
	// The time in seconds that Envoy will wait before shutting down the
	// parent process during a hot restart. MUST be >=1s (e.g., _1s/1m/1h_).
	// MUST BE greater than _drain_duration_ parameter.
	ParentShutdownDuration *google_protobuf.Duration `protobuf:"bytes,5,opt,name=parent_shutdown_duration,json=parentShutdownDuration" json:"parent_shutdown_duration,omitempty"`
	// Address of the discovery service exposing xDS (e.g. _istio-pilot:8080_).
	DiscoveryAddress string `protobuf:"bytes,6,opt,name=discovery_address,json=discoveryAddress" json:"discovery_address,omitempty"`
	// Polling interval for service discovery (used by EDS, CDS, LDS, but not RDS). (MUST BE >=1ms)
	DiscoveryRefreshDelay *google_protobuf.Duration `protobuf:"bytes,7,opt,name=discovery_refresh_delay,json=discoveryRefreshDelay" json:"discovery_refresh_delay,omitempty"`
	// Address of the Zipkin service (e.g. _zipkin:9411_).
	ZipkinAddress string `protobuf:"bytes,8,opt,name=zipkin_address,json=zipkinAddress" json:"zipkin_address,omitempty"`
	// Connection timeout used by Envoy for supporting services. (MUST BE >=1ms)
	ConnectTimeout *google_protobuf.Duration `protobuf:"bytes,9,opt,name=connect_timeout,json=connectTimeout" json:"connect_timeout,omitempty"`
	// IP Address and Port of a statsd UDP listener (e.g. _10.75.241.127:9125_).
	StatsdUdpAddress string `protobuf:"bytes,10,opt,name=statsd_udp_address,json=statsdUdpAddress" json:"statsd_udp_address,omitempty"`
	// Port on which Envoy should listen for administrative commands.
	ProxyAdminPort int32 `protobuf:"varint,11,opt,name=proxy_admin_port,json=proxyAdminPort" json:"proxy_admin_port,omitempty"`
	// The availability zone where this Envoy instance is running. When running
	// Envoy as a sidecar in Kubernetes, this flag must be one of the availability
	// zones assigned to a node using failure-domain.beta.kubernetes.io/zone annotation.
	AvailabilityZone string `protobuf:"bytes,12,opt,name=availability_zone,json=availabilityZone" json:"availability_zone,omitempty"`
	// Authentication policy defines the global switch to control authentication
	// for Envoy-to-Envoy communication for istio components Mixer and Pilot.
	ControlPlaneAuthPolicy AuthenticationPolicy `protobuf:"varint,13,opt,name=control_plane_auth_policy,json=controlPlaneAuthPolicy,enum=istio.mesh.v1alpha1.AuthenticationPolicy" json:"control_plane_auth_policy,omitempty"`
	// File path of custom proxy configuration, currently used by proxies
	// in front of Mixer and Pilot.
	CustomConfigFile string `protobuf:"bytes,14,opt,name=custom_config_file,json=customConfigFile" json:"custom_config_file,omitempty"`
	// Maximum length of name field in Envoy's metrics. The length of the name field
	// is determined by the length of a name field in a service and the set of labels that
	// comprise a particular version of the service. The default value is set to 189 characters.
	// Envoy's internal metrics take up 67 characters, for a total of 256 character name per metric.
	// Increase the value of this field if you find that the metrics from Envoys are truncated.
	StatNameLength int32 `protobuf:"varint,15,opt,name=stat_name_length,json=statNameLength" json:"stat_name_length,omitempty"`
	// The number of worker threads to run. Default value is number of cores on the machine.
	Concurrency int32 `protobuf:"varint,16,opt,name=concurrency" json:"concurrency,omitempty"`
}

func (m *ProxyConfig) Reset()                    { *m = ProxyConfig{} }
func (m *ProxyConfig) String() string            { return proto.CompactTextString(m) }
func (*ProxyConfig) ProtoMessage()               {}
func (*ProxyConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProxyConfig) GetConfigPath() string {
	if m != nil {
		return m.ConfigPath
	}
	return ""
}

func (m *ProxyConfig) GetBinaryPath() string {
	if m != nil {
		return m.BinaryPath
	}
	return ""
}

func (m *ProxyConfig) GetServiceCluster() string {
	if m != nil {
		return m.ServiceCluster
	}
	return ""
}

func (m *ProxyConfig) GetDrainDuration() *google_protobuf.Duration {
	if m != nil {
		return m.DrainDuration
	}
	return nil
}

func (m *ProxyConfig) GetParentShutdownDuration() *google_protobuf.Duration {
	if m != nil {
		return m.ParentShutdownDuration
	}
	return nil
}

func (m *ProxyConfig) GetDiscoveryAddress() string {
	if m != nil {
		return m.DiscoveryAddress
	}
	return ""
}

func (m *ProxyConfig) GetDiscoveryRefreshDelay() *google_protobuf.Duration {
	if m != nil {
		return m.DiscoveryRefreshDelay
	}
	return nil
}

func (m *ProxyConfig) GetZipkinAddress() string {
	if m != nil {
		return m.ZipkinAddress
	}
	return ""
}

func (m *ProxyConfig) GetConnectTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.ConnectTimeout
	}
	return nil
}

func (m *ProxyConfig) GetStatsdUdpAddress() string {
	if m != nil {
		return m.StatsdUdpAddress
	}
	return ""
}

func (m *ProxyConfig) GetProxyAdminPort() int32 {
	if m != nil {
		return m.ProxyAdminPort
	}
	return 0
}

func (m *ProxyConfig) GetAvailabilityZone() string {
	if m != nil {
		return m.AvailabilityZone
	}
	return ""
}

func (m *ProxyConfig) GetControlPlaneAuthPolicy() AuthenticationPolicy {
	if m != nil {
		return m.ControlPlaneAuthPolicy
	}
	return AuthenticationPolicy_NONE
}

func (m *ProxyConfig) GetCustomConfigFile() string {
	if m != nil {
		return m.CustomConfigFile
	}
	return ""
}

func (m *ProxyConfig) GetStatNameLength() int32 {
	if m != nil {
		return m.StatNameLength
	}
	return 0
}

func (m *ProxyConfig) GetConcurrency() int32 {
	if m != nil {
		return m.Concurrency
	}
	return 0
}

// MeshConfig defines mesh-wide variables shared by all Envoy instances in the
// Istio service mesh.
type MeshConfig struct {
	// Address of the egress Envoy service (e.g. _istio-egress:80_).
	EgressProxyAddress string `protobuf:"bytes,1,opt,name=egress_proxy_address,json=egressProxyAddress" json:"egress_proxy_address,omitempty"`
	// Address of the mixer service (e.g. _istio-mixer:15004_).
	// Empty value disables Mixer checks and telemetry.
	MixerAddress string `protobuf:"bytes,2,opt,name=mixer_address,json=mixerAddress" json:"mixer_address,omitempty"`
	// Disable policy checks by the mixer service. Metrics will still be
	// reported to the mixer for HTTP requests and TCP connections. Default
	// is false, i.e. mixer policy check is enabled by default.
	DisablePolicyChecks bool `protobuf:"varint,3,opt,name=disable_policy_checks,json=disablePolicyChecks" json:"disable_policy_checks,omitempty"`
	// Port on which Envoy should listen for incoming connections from
	// other services.
	ProxyListenPort int32 `protobuf:"varint,4,opt,name=proxy_listen_port,json=proxyListenPort" json:"proxy_listen_port,omitempty"`
	// Port on which Envoy should listen for HTTP PROXY requests if set.
	ProxyHttpPort int32 `protobuf:"varint,5,opt,name=proxy_http_port,json=proxyHttpPort" json:"proxy_http_port,omitempty"`
	// Connection timeout used by Envoy. (MUST BE >=1ms)
	ConnectTimeout *google_protobuf.Duration `protobuf:"bytes,6,opt,name=connect_timeout,json=connectTimeout" json:"connect_timeout,omitempty"`
	// Class of ingress resources to be processed by Istio ingress
	// controller.  This corresponds to the value of
	// "kubernetes.io/ingress.class" annotation.
	IngressClass string `protobuf:"bytes,7,opt,name=ingress_class,json=ingressClass" json:"ingress_class,omitempty"`
	// Name of the kubernetes service used for the istio ingress controller.
	IngressService string `protobuf:"bytes,8,opt,name=ingress_service,json=ingressService" json:"ingress_service,omitempty"`
	// Defines whether to use Istio ingress controller for annotated or all ingress resources.
	IngressControllerMode MeshConfig_IngressControllerMode `protobuf:"varint,9,opt,name=ingress_controller_mode,json=ingressControllerMode,enum=istio.mesh.v1alpha1.MeshConfig_IngressControllerMode" json:"ingress_controller_mode,omitempty"`
	// Authentication policy defines the global switch to control authentication
	// for Envoy-to-Envoy communication.
	// Use authentication_policy instead.
	AuthPolicy MeshConfig_AuthPolicy `protobuf:"varint,10,opt,name=auth_policy,json=authPolicy,enum=istio.mesh.v1alpha1.MeshConfig_AuthPolicy" json:"auth_policy,omitempty"`
	// Polling interval for RDS (MUST BE >=1ms)
	RdsRefreshDelay *google_protobuf.Duration `protobuf:"bytes,11,opt,name=rds_refresh_delay,json=rdsRefreshDelay" json:"rds_refresh_delay,omitempty"`
	// Flag to control generation of trace spans and request IDs.
	// Requires a trace span collector defined in the proxy configuration.
	EnableTracing bool `protobuf:"varint,12,opt,name=enable_tracing,json=enableTracing" json:"enable_tracing,omitempty"`
	// File address for the proxy access log (e.g. /dev/stdout).
	// Empty value disables access logging.
	AccessLogFile string `protobuf:"bytes,13,opt,name=access_log_file,json=accessLogFile" json:"access_log_file,omitempty"`
	// Default proxy config used by the proxy injection mechanism operating in the mesh
	// (e.g. Kubernetes admission controller)
	// In case of Kubernetes, the proxy config is applied once during the injection process,
	// and remain constant for the duration of the pod. The rest of the mesh config can be changed
	// at runtime and config gets distributed dynamically.
	DefaultConfig *ProxyConfig `protobuf:"bytes,14,opt,name=default_config,json=defaultConfig" json:"default_config,omitempty"`
	// List of remote services for which mTLS authentication should not be expected by Istio .
	// Typically, these are control services (e.g kubernetes API server) that don't have Istio sidecar
	// to transparently terminate mTLS authentication.
	// It has no effect if the authentication policy is already 'NONE'.
	// DO NOT use this setting for services that are managed by Istio (i.e. using Istio sidecar).
	// Instead, use service-level annotations to overwrite the authentication policy.
	MtlsExcludedServices []string `protobuf:"bytes,15,rep,name=mtls_excluded_services,json=mtlsExcludedServices" json:"mtls_excluded_services,omitempty"`
}

func (m *MeshConfig) Reset()                    { *m = MeshConfig{} }
func (m *MeshConfig) String() string            { return proto.CompactTextString(m) }
func (*MeshConfig) ProtoMessage()               {}
func (*MeshConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MeshConfig) GetEgressProxyAddress() string {
	if m != nil {
		return m.EgressProxyAddress
	}
	return ""
}

func (m *MeshConfig) GetMixerAddress() string {
	if m != nil {
		return m.MixerAddress
	}
	return ""
}

func (m *MeshConfig) GetDisablePolicyChecks() bool {
	if m != nil {
		return m.DisablePolicyChecks
	}
	return false
}

func (m *MeshConfig) GetProxyListenPort() int32 {
	if m != nil {
		return m.ProxyListenPort
	}
	return 0
}

func (m *MeshConfig) GetProxyHttpPort() int32 {
	if m != nil {
		return m.ProxyHttpPort
	}
	return 0
}

func (m *MeshConfig) GetConnectTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.ConnectTimeout
	}
	return nil
}

func (m *MeshConfig) GetIngressClass() string {
	if m != nil {
		return m.IngressClass
	}
	return ""
}

func (m *MeshConfig) GetIngressService() string {
	if m != nil {
		return m.IngressService
	}
	return ""
}

func (m *MeshConfig) GetIngressControllerMode() MeshConfig_IngressControllerMode {
	if m != nil {
		return m.IngressControllerMode
	}
	return MeshConfig_OFF
}

func (m *MeshConfig) GetAuthPolicy() MeshConfig_AuthPolicy {
	if m != nil {
		return m.AuthPolicy
	}
	return MeshConfig_NONE
}

func (m *MeshConfig) GetRdsRefreshDelay() *google_protobuf.Duration {
	if m != nil {
		return m.RdsRefreshDelay
	}
	return nil
}

func (m *MeshConfig) GetEnableTracing() bool {
	if m != nil {
		return m.EnableTracing
	}
	return false
}

func (m *MeshConfig) GetAccessLogFile() string {
	if m != nil {
		return m.AccessLogFile
	}
	return ""
}

func (m *MeshConfig) GetDefaultConfig() *ProxyConfig {
	if m != nil {
		return m.DefaultConfig
	}
	return nil
}

func (m *MeshConfig) GetMtlsExcludedServices() []string {
	if m != nil {
		return m.MtlsExcludedServices
	}
	return nil
}

func init() {
	proto.RegisterType((*ProxyConfig)(nil), "istio.mesh.v1alpha1.ProxyConfig")
	proto.RegisterType((*MeshConfig)(nil), "istio.mesh.v1alpha1.MeshConfig")
	proto.RegisterEnum("istio.mesh.v1alpha1.AuthenticationPolicy", AuthenticationPolicy_name, AuthenticationPolicy_value)
	proto.RegisterEnum("istio.mesh.v1alpha1.MeshConfig_IngressControllerMode", MeshConfig_IngressControllerMode_name, MeshConfig_IngressControllerMode_value)
	proto.RegisterEnum("istio.mesh.v1alpha1.MeshConfig_AuthPolicy", MeshConfig_AuthPolicy_name, MeshConfig_AuthPolicy_value)
}

func init() { proto.RegisterFile("mesh/v1alpha1/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 941 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4d, 0x6f, 0xe3, 0x36,
	0x13, 0x5e, 0xe7, 0xcb, 0xc9, 0x38, 0x96, 0x1d, 0x6d, 0x92, 0xd5, 0x06, 0x2f, 0xde, 0x1a, 0x59,
	0x74, 0xeb, 0xa6, 0x85, 0xdd, 0x4d, 0xdb, 0x43, 0x2f, 0x45, 0xf3, 0xd9, 0x0d, 0xe0, 0x24, 0xae,
	0xe2, 0x5c, 0xf6, 0x42, 0x30, 0x12, 0x63, 0x11, 0x4b, 0x93, 0x02, 0x49, 0xa5, 0xf1, 0xfe, 0xaf,
	0xfe, 0x9e, 0x5e, 0xfb, 0x33, 0x0a, 0x71, 0x28, 0xc7, 0x5d, 0xb8, 0x08, 0x7a, 0xd4, 0x33, 0xcf,
	0xcc, 0x70, 0x3e, 0x9e, 0x11, 0xec, 0x4d, 0x98, 0xc9, 0xfa, 0x0f, 0xef, 0xa8, 0xc8, 0x33, 0xfa,
	0xae, 0x9f, 0x28, 0x79, 0xcf, 0xc7, 0xbd, 0x5c, 0x2b, 0xab, 0xc2, 0x97, 0xdc, 0x58, 0xae, 0x7a,
	0x25, 0xa3, 0x57, 0x31, 0xf6, 0xfe, 0x3f, 0x56, 0x6a, 0x2c, 0x58, 0xdf, 0x51, 0xee, 0x8a, 0xfb,
	0x7e, 0x5a, 0x68, 0x6a, 0xb9, 0x92, 0xe8, 0xb4, 0xff, 0xe7, 0x1a, 0x34, 0x86, 0x5a, 0x3d, 0x4e,
	0x4f, 0x5c, 0xa8, 0xf0, 0x0b, 0x68, 0x60, 0x50, 0x92, 0x53, 0x9b, 0x45, 0xb5, 0x4e, 0xad, 0xbb,
	0x11, 0x03, 0x42, 0x43, 0x6a, 0xb3, 0x92, 0x70, 0xc7, 0x25, 0xd5, 0x53, 0x24, 0x2c, 0x21, 0x01,
	0x21, 0x47, 0xf8, 0x0a, 0x5a, 0x86, 0xe9, 0x07, 0x9e, 0x30, 0x92, 0x88, 0xc2, 0x58, 0xa6, 0xa3,
	0x65, 0x47, 0x0a, 0x3c, 0x7c, 0x82, 0x68, 0xf8, 0x0b, 0x04, 0xa9, 0xa6, 0x5c, 0x92, 0xea, 0x49,
	0xd1, 0x4a, 0xa7, 0xd6, 0x6d, 0x1c, 0xbe, 0xee, 0xe1, 0x9b, 0x7b, 0xd5, 0x9b, 0x7b, 0xa7, 0x9e,
	0x10, 0x37, 0x9d, 0x43, 0xf5, 0x19, 0xde, 0x40, 0x94, 0x53, 0xcd, 0xa4, 0x25, 0x26, 0x2b, 0x6c,
	0xaa, 0x7e, 0x9f, 0x8b, 0xb5, 0xfa, 0x5c, 0xac, 0x5d, 0x74, 0xbd, 0xf1, 0x9e, 0xb3, 0xa0, 0xdf,
	0xc0, 0x56, 0xca, 0x4d, 0xa2, 0x1e, 0x98, 0x9e, 0x12, 0x9a, 0xa6, 0x9a, 0x19, 0x13, 0xad, 0xb9,
	0x0a, 0xda, 0x33, 0xc3, 0x11, 0xe2, 0xe1, 0x6f, 0xf0, 0xea, 0x89, 0xac, 0xd9, 0xbd, 0x66, 0x26,
	0x23, 0x29, 0x13, 0x74, 0x1a, 0xd5, 0x9f, 0x7b, 0xc0, 0xce, 0xcc, 0x33, 0x46, 0xc7, 0xd3, 0xd2,
	0x2f, 0xfc, 0x12, 0x82, 0x4f, 0x3c, 0xff, 0xc8, 0xe5, 0x2c, 0xf9, 0xba, 0x4b, 0xde, 0x44, 0xb4,
	0xca, 0x7c, 0x0c, 0xad, 0x44, 0x49, 0xc9, 0x12, 0x4b, 0x2c, 0x9f, 0x30, 0x55, 0xd8, 0x68, 0xe3,
	0xb9, 0x8c, 0x81, 0xf7, 0x18, 0xa1, 0x43, 0xf8, 0x2d, 0x84, 0xc6, 0x52, 0x6b, 0x52, 0x52, 0xa4,
	0xf9, 0x2c, 0x1d, 0x60, 0xad, 0x68, 0xb9, 0x4d, 0xf3, 0x2a, 0x63, 0x17, 0xda, 0x79, 0xb9, 0x29,
	0x84, 0xa6, 0x13, 0x2e, 0x49, 0xae, 0xb4, 0x8d, 0x1a, 0x9d, 0x5a, 0x77, 0x35, 0x0e, 0x1c, 0x7e,
	0x54, 0xc2, 0x43, 0xa5, 0x6d, 0xd9, 0x42, 0xfa, 0x40, 0xb9, 0xa0, 0x77, 0x5c, 0x70, 0x3b, 0x25,
	0x9f, 0x94, 0x64, 0xd1, 0x26, 0x86, 0x9d, 0x37, 0x7c, 0x50, 0x92, 0x85, 0x29, 0xbc, 0x4e, 0x94,
	0xb4, 0x5a, 0x09, 0x92, 0x0b, 0x2a, 0x19, 0xa1, 0x85, 0xcd, 0x48, 0xae, 0x04, 0x4f, 0xa6, 0x51,
	0xb3, 0x53, 0xeb, 0x06, 0x87, 0x5f, 0xf7, 0x16, 0xac, 0x76, 0xef, 0xa8, 0xb0, 0x19, 0x93, 0x96,
	0x27, 0xae, 0xb8, 0xa1, 0x73, 0x88, 0x77, 0x7d, 0xac, 0x61, 0x19, 0xaa, 0x64, 0x20, 0x5e, 0x96,
	0x9a, 0x14, 0xc6, 0xaa, 0x09, 0xf1, 0xeb, 0x7d, 0xcf, 0x05, 0x8b, 0x02, 0x7c, 0x13, 0x5a, 0x50,
	0x01, 0xe7, 0x5c, 0xb0, 0xb2, 0xd4, 0xb2, 0x7c, 0x22, 0xe9, 0x84, 0x11, 0xc1, 0xe4, 0xd8, 0x66,
	0x51, 0x0b, 0x4b, 0x2d, 0xf1, 0x2b, 0x3a, 0x61, 0x03, 0x87, 0x86, 0x1d, 0xa7, 0x97, 0xa4, 0xd0,
	0x9a, 0xc9, 0x64, 0x1a, 0xb5, 0x1d, 0x69, 0x1e, 0xda, 0xff, 0xa3, 0x0e, 0x70, 0xc9, 0x4c, 0xe6,
	0x05, 0xf6, 0x1d, 0x6c, 0xb3, 0x71, 0xd9, 0x4f, 0x52, 0x35, 0x13, 0xbb, 0x8e, 0x4a, 0x0b, 0xd1,
	0x36, 0xc4, 0x7e, 0x62, 0xdf, 0xdf, 0x40, 0x73, 0xc2, 0x1f, 0x99, 0x9e, 0x51, 0x51, 0x73, 0x9b,
	0x0e, 0xac, 0x48, 0x87, 0x50, 0xae, 0x13, 0xbd, 0x13, 0xcc, 0xb7, 0x8e, 0x24, 0x19, 0x4b, 0x3e,
	0x1a, 0xa7, 0xbd, 0xf5, 0xf8, 0xa5, 0x37, 0x62, 0x37, 0x4e, 0x9c, 0x29, 0x3c, 0x80, 0x2d, 0x7c,
	0x83, 0xe0, 0xc6, 0x32, 0x3f, 0xd1, 0x15, 0x57, 0x41, 0xcb, 0x19, 0x06, 0x0e, 0x77, 0x23, 0x7d,
	0x0b, 0x08, 0x91, 0xcc, 0xda, 0x1c, 0x99, 0xab, 0x8e, 0xd9, 0x74, 0xf0, 0x7b, 0x6b, 0x73, 0xc7,
	0x5b, 0xb0, 0x96, 0x6b, 0xff, 0x75, 0x2d, 0xdf, 0x40, 0x93, 0x4b, 0xec, 0x51, 0x22, 0xa8, 0x31,
	0x4e, 0x4a, 0x1b, 0xf1, 0xa6, 0x07, 0x4f, 0x4a, 0xac, 0x3c, 0x33, 0x15, 0xc9, 0xdf, 0x15, 0xaf,
	0x93, 0xc0, 0xc3, 0x37, 0x88, 0x86, 0x13, 0x78, 0x35, 0x8b, 0x86, 0xbb, 0x21, 0x98, 0x26, 0x13,
	0x95, 0x32, 0x27, 0x98, 0xe0, 0xf0, 0xc7, 0x85, 0xdb, 0xf5, 0x34, 0xb2, 0xde, 0x85, 0xcf, 0x3b,
	0xf3, 0xbe, 0x54, 0x29, 0x8b, 0x77, 0xf8, 0x22, 0x38, 0xbc, 0x86, 0xc6, 0xfc, 0x02, 0x83, 0x4b,
	0x71, 0xf0, 0x5c, 0x8a, 0xa7, 0x4d, 0x3d, 0x5e, 0x8a, 0x6a, 0x31, 0xd0, 0xa7, 0xcd, 0x3d, 0x83,
	0x2d, 0x9d, 0x9a, 0xcf, 0x8e, 0x4b, 0xe3, 0xb9, 0x9e, 0xb6, 0x74, 0x6a, 0x3e, 0x3f, 0x2b, 0x4c,
	0xba, 0xfd, 0xb0, 0x9a, 0x26, 0x5c, 0x8e, 0x9d, 0x20, 0xd7, 0xe3, 0x26, 0xa2, 0x23, 0x04, 0xcb,
	0x39, 0xd3, 0x24, 0x29, 0x9b, 0x25, 0x94, 0x17, 0x49, 0x13, 0xcf, 0x0f, 0xc2, 0x03, 0x85, 0x0a,
	0xf9, 0x15, 0x82, 0x94, 0xdd, 0xd3, 0x42, 0x58, 0x2f, 0x28, 0xa7, 0xa5, 0xc6, 0x61, 0x67, 0x61,
	0xa5, 0x73, 0x7f, 0x98, 0xb8, 0xe9, 0xfd, 0xbc, 0x1e, 0x7e, 0x80, 0xdd, 0x89, 0x15, 0x86, 0xb0,
	0xc7, 0x44, 0x14, 0x29, 0x4b, 0xab, 0x69, 0x9a, 0xa8, 0xd5, 0x59, 0xee, 0x6e, 0xc4, 0xdb, 0xa5,
	0xf5, 0xcc, 0x1b, 0xfd, 0x4c, 0xcd, 0xfe, 0x4f, 0xb0, 0xb3, 0x70, 0x2a, 0x61, 0x1d, 0x96, 0xaf,
	0xcf, 0xcf, 0xdb, 0x2f, 0xc2, 0x06, 0xd4, 0x4f, 0xcf, 0xce, 0x8f, 0x6e, 0x07, 0xa3, 0x76, 0x2d,
	0x04, 0x58, 0xbb, 0x19, 0xc5, 0x17, 0x27, 0xa3, 0xf6, 0xd2, 0xfe, 0x5b, 0x80, 0xb9, 0xbb, 0xb0,
	0x0e, 0x2b, 0x57, 0xd7, 0x57, 0x67, 0xed, 0x17, 0x61, 0x00, 0x70, 0x79, 0x3b, 0xba, 0x3d, 0x1a,
	0x90, 0xd1, 0xe0, 0xa6, 0x5d, 0x3b, 0xf8, 0x19, 0xb6, 0x17, 0x5d, 0x98, 0x7f, 0xf7, 0x08, 0x37,
	0xa1, 0x7e, 0x71, 0xf5, 0xfe, 0x2c, 0xbe, 0x18, 0xb5, 0xff, 0xaa, 0x1f, 0xff, 0xef, 0xc3, 0x1e,
	0xb6, 0x82, 0xab, 0x3e, 0xcd, 0x79, 0xff, 0x1f, 0x7f, 0xee, 0xbb, 0x35, 0x37, 0xb2, 0xef, 0xff,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xbd, 0xe4, 0xe0, 0xd1, 0x07, 0x00, 0x00,
}
