/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha2

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// KubeletConfigSpec defines the kubelet configuration
type KubeletConfigSpec struct {
	// APIServers is not used for clusters version 1.6 and later - flag removed
	APIServers string `json:"apiServers,omitempty" flag:"api-servers"`
	// AnonymousAuth permits you to control auth to the kubelet api
	AnonymousAuth *bool `json:"anonymousAuth,omitempty" flag:"anonymous-auth"`
	// ClientCAFile is the path to a CA certificate
	ClientCAFile string `json:"clientCaFile,omitempty" flag:"client-ca-file"`
	// KubeconfigPath is the path of kubeconfig for the kubelet
	KubeconfigPath string `json:"kubeconfigPath,omitempty" flag:"kubeconfig"`
	// RequireKubeconfig indicates a kubeconfig is required
	RequireKubeconfig *bool `json:"requireKubeconfig,omitempty" flag:"require-kubeconfig"`
	// LogLevel is the logging level of the kubelet
	LogLevel *int32 `json:"logLevel,omitempty" flag:"v" flag-empty:"0"`
	// config is the path to the config file or directory of files
	PodManifestPath string `json:"podManifestPath,omitempty" flag:"pod-manifest-path"`
	// HostnameOverride is the hostname used to identify the kubelet instead of the actual hostname.
	HostnameOverride string `json:"hostnameOverride,omitempty" flag:"hostname-override"`
	// PodInfraContainerImage is the image whose network/ipc containers in each pod will use.
	PodInfraContainerImage string `json:"podInfraContainerImage,omitempty" flag:"pod-infra-container-image"`
	// SeccompProfileRoot is the directory path for seccomp profiles.
	SeccompProfileRoot *string `json:"seccompProfileRoot,omitempty" flag:"seccomp-profile-root"`
	// AllowPrivileged enables containers to request privileged mode (defaults to false)
	AllowPrivileged *bool `json:"allowPrivileged,omitempty" flag:"allow-privileged"`
	// EnableDebuggingHandlers enables server endpoints for log collection and local running of containers and commands
	EnableDebuggingHandlers *bool `json:"enableDebuggingHandlers,omitempty" flag:"enable-debugging-handlers"`
	// RegisterNode enables automatic registration with the apiserver.
	RegisterNode *bool `json:"registerNode,omitempty" flag:"register-node"`
	// ClusterDomain is the DNS domain for this cluster
	ClusterDomain string `json:"clusterDomain,omitempty" flag:"cluster-domain"`
	// ClusterDNS is the IP address for a cluster DNS server
	ClusterDNS string `json:"clusterDNS,omitempty" flag:"cluster-dns"`
	// NetworkPluginName is the name of the network plugin to be invoked for various events in kubelet/pod lifecycle
	NetworkPluginName string `json:"networkPluginName,omitempty" flag:"network-plugin"`
	// CloudProvider is the provider for cloud services.
	CloudProvider string `json:"cloudProvider,omitempty" flag:"cloud-provider"`
	// KubeletCgroups is the absolute name of cgroups to isolate the kubelet in.
	KubeletCgroups string `json:"kubeletCgroups,omitempty" flag:"kubelet-cgroups"`
	// Cgroups that container runtime is expected to be isolated in.
	RuntimeCgroups string `json:"runtimeCgroups,omitempty" flag:"runtime-cgroups"`
	// ReadOnlyPort is the port used by the kubelet api for read-only access (default 10255)
	ReadOnlyPort *int32 `json:"readOnlyPort,omitempty" flag:"read-only-port"`
	// SystemCgroups is absolute name of cgroups in which to place
	// all non-kernel processes that are not already in a container. Empty
	// for no container. Rolling back the flag requires a reboot.
	SystemCgroups string `json:"systemCgroups,omitempty" flag:"system-cgroups"`
	// cgroupRoot is the root cgroup to use for pods. This is handled by the container runtime on a best effort basis.
	CgroupRoot string `json:"cgroupRoot,omitempty" flag:"cgroup-root"`
	// configureCBR0 enables the kublet to configure cbr0 based on Node.Spec.PodCIDR.
	ConfigureCBR0 *bool `json:"configureCbr0,omitempty" flag:"configure-cbr0"`
	// How should the kubelet configure the container bridge for hairpin packets.
	// Setting this flag allows endpoints in a Service to loadbalance back to
	// themselves if they should try to access their own Service. Values:
	//   "promiscuous-bridge": make the container bridge promiscuous.
	//   "hairpin-veth":       set the hairpin flag on container veth interfaces.
	//   "none":               do nothing.
	// Setting --configure-cbr0 to false implies that to achieve hairpin NAT
	// one must set --hairpin-mode=veth-flag, because bridge assumes the
	// existence of a container bridge named cbr0.
	HairpinMode string `json:"hairpinMode,omitempty" flag:"hairpin-mode"`
	// The node has babysitter process monitoring docker and kubelet. Removed as of 1.7
	BabysitDaemons *bool `json:"babysitDaemons,omitempty" flag:"babysit-daemons"`
	// MaxPods is the number of pods that can run on this Kubelet.
	MaxPods *int32 `json:"maxPods,omitempty" flag:"max-pods"`
	// NvidiaGPUs is the number of NVIDIA GPU devices on this node.
	NvidiaGPUs int32 `json:"nvidiaGPUs,omitempty" flag:"experimental-nvidia-gpus" flag-empty:"0"`
	// PodCIDR is the CIDR to use for pod IP addresses, only used in standalone mode.
	// In cluster mode, this is obtained from the master.
	PodCIDR string `json:"podCIDR,omitempty" flag:"pod-cidr"`
	// ResolverConfig is the resolver configuration file used as the basis for the container DNS resolution configuration."), []
	ResolverConfig *string `json:"resolvConf,omitempty" flag:"resolv-conf" flag-include-empty:"true"`
	// ReconcileCIDR is Reconcile node CIDR with the CIDR specified by the
	// API server. No-op if register-node or configure-cbr0 is false.
	ReconcileCIDR *bool `json:"reconcileCIDR,omitempty" flag:"reconcile-cidr"`
	// registerSchedulable tells the kubelet to register the node as schedulable. No-op if register-node is false.
	RegisterSchedulable *bool `json:"registerSchedulable,omitempty" flag:"register-schedulable"`
	//// SerializeImagePulls when enabled, tells the Kubelet to pull images one
	//// at a time. We recommend *not* changing the default value on nodes that
	//// run docker daemon with version  < 1.9 or an Aufs storage backend.
	//// Issue #10959 has more details.
	SerializeImagePulls *bool `json:"serializeImagePulls,omitempty" flag:"serialize-image-pulls"`
	// NodeLabels to add when registering the node in the cluster.
	NodeLabels map[string]string `json:"nodeLabels,omitempty" flag:"node-labels"`
	// NonMasqueradeCIDR configures masquerading: traffic to IPs outside this range will use IP masquerade.
	NonMasqueradeCIDR string `json:"nonMasqueradeCIDR,omitempty" flag:"non-masquerade-cidr"`
	// Enable gathering custom metrics.
	EnableCustomMetrics *bool `json:"enableCustomMetrics,omitempty" flag:"enable-custom-metrics"`
	// NetworkPluginMTU is the MTU to be passed to the network plugin,
	// and overrides the default MTU for cases where it cannot be automatically
	// computed (such as IPSEC).
	NetworkPluginMTU *int32 `json:"networkPluginMTU,omitempty" flag:"network-plugin-mtu"`
	// ImageGCHighThresholdPercent is the percent of disk usage after which
	// image garbage collection is always run.
	ImageGCHighThresholdPercent *int32 `json:"imageGCHighThresholdPercent,omitempty" flag:"image-gc-high-threshold"`
	// ImageGCLowThresholdPercent is the percent of disk usage before which
	// image garbage collection is never run. Lowest disk usage to garbage
	// collect to.
	ImageGCLowThresholdPercent *int32 `json:"imageGCLowThresholdPercent,omitempty" flag:"image-gc-low-threshold"`
	// Comma-delimited list of hard eviction expressions.  For example, 'memory.available<300Mi'.
	EvictionHard *string `json:"evictionHard,omitempty" flag:"eviction-hard"`
	// Comma-delimited list of soft eviction expressions.  For example, 'memory.available<300Mi'.
	EvictionSoft string `json:"evictionSoft,omitempty" flag:"eviction-soft"`
	// Comma-delimited list of grace periods for each soft eviction signal.  For example, 'memory.available=30s'.
	EvictionSoftGracePeriod string `json:"evictionSoftGracePeriod,omitempty" flag:"eviction-soft-grace-period"`
	// Duration for which the kubelet has to wait before transitioning out of an eviction pressure condition.
	EvictionPressureTransitionPeriod *metav1.Duration `json:"evictionPressureTransitionPeriod,omitempty" flag:"eviction-pressure-transition-period" flag-empty:"0s"`
	// Maximum allowed grace period (in seconds) to use when terminating pods in response to a soft eviction threshold being met.
	EvictionMaxPodGracePeriod int32 `json:"evictionMaxPodGracePeriod,omitempty" flag:"eviction-max-pod-grace-period" flag-empty:"0"`
	// Comma-delimited list of minimum reclaims (e.g. imagefs.available=2Gi) that describes the minimum amount of resource the kubelet will reclaim when performing a pod eviction if that resource is under pressure.
	EvictionMinimumReclaim string `json:"evictionMinimumReclaim,omitempty" flag:"eviction-minimum-reclaim"`
	// The full path of the directory in which to search for additional third party volume plugins
	VolumePluginDirectory string `json:"volumePluginDirectory,omitempty" flag:"volume-plugin-dir"`
	// Taints to add when registering a node in the cluster
	Taints []string `json:"taints,omitempty" flag:"register-with-taints"`
	// FeatureGates is set of key=value pairs that describe feature gates for alpha/experimental features.
	FeatureGates map[string]string `json:"featureGates,omitempty" flag:"feature-gates"`
	// Resource reservation for kubernetes system daemons like the kubelet, container runtime, node problem detector, etc.
	KubeReserved map[string]string `json:"kubeReserved,omitempty" flag:"kube-reserved"`
	// Control group for kube daemons.
	KubeReservedCgroup string `json:"kubeReservedCgroup,omitempty" flag:"kube-reserved-cgroup"`
	// Capture resource reservation for OS system daemons like sshd, udev, etc.
	SystemReserved map[string]string `json:"systemReserved,omitempty" flag:"system-reserved"`
	// Parent control group for OS system daemons.
	SystemReservedCgroup string `json:"systemReservedCgroup,omitempty" flag:"system-reserved-cgroup"`
	// Enforce Allocatable across pods whenever the overall usage across all pods exceeds Allocatable.
	EnforceNodeAllocatable string `json:"enforceNodeAllocatable,omitempty" flag:"enforce-node-allocatable"`
	// RuntimeRequestTimeout is timeout for runtime requests on - pull, logs, exec and attach
	RuntimeRequestTimeout *metav1.Duration `json:"runtimeRequestTimeout,omitempty" flag:"runtime-request-timeout"`
	// VolumeStatsAggPeriod is the interval for kubelet to calculate and cache the volume disk usage for all pods and volumes
	VolumeStatsAggPeriod *metav1.Duration `json:"volumeStatsAggPeriod,omitempty" flag:"volume-stats-agg-period"`
}

// KubeProxyConfig defined the configuration for a proxy
type KubeProxyConfig struct {
	Image string `json:"image,omitempty"`
	// TODO: Better type ?
	CPURequest string `json:"cpuRequest,omitempty"` // e.g. "20m"
	// LogLevel is the logging level of the proxy
	LogLevel int32 `json:"logLevel,omitempty" flag:"v"`
	// ClusterCIDR is the CIDR range of the pods in the cluster
	ClusterCIDR string `json:"clusterCIDR,omitempty" flag:"cluster-cidr"`
	// HostnameOverride, if non-empty, will be used as the identity instead of the actual hostname.
	HostnameOverride string `json:"hostnameOverride,omitempty" flag:"hostname-override"`
	// Master is the address of the Kubernetes API server (overrides any value in kubeconfig)
	Master string `json:"master,omitempty" flag:"master"`
	// FeatureGates is a series of key pairs used to switch on features for the proxy
	FeatureGates map[string]string `json:"featureGates" flag:"feature-gates"`
}

// KubeAPIServerConfig defines the configuration for the kube api
type KubeAPIServerConfig struct {
	// Image is the docker container used
	Image string `json:"image,omitempty"`
	// LogLevel is the logging level of the api
	LogLevel int32 `json:"logLevel,omitempty" flag:"v" flag-empty:"0"`
	// CloudProvider is the name of the cloudProvider we are using, aws, gce etcd
	CloudProvider string `json:"cloudProvider,omitempty" flag:"cloud-provider"`
	// SecurePort is the port the kube runs on
	SecurePort int32 `json:"securePort,omitempty" flag:"secure-port"`
	// InsecurePort is the port the insecure api runs
	InsecurePort int32 `json:"insecurePort,omitempty" flag:"insecure-port"`
	// Address is the binding address for the kube api
	Address string `json:"address,omitempty" flag:"address"`
	// AdmissionControl is a list of admission controllers to user
	AdmissionControl []string `json:"admissionControl,omitempty" flag:"admission-control"`
	// ServiceClusterIPRange is the service address range
	ServiceClusterIPRange string `json:"serviceClusterIPRange,omitempty" flag:"service-cluster-ip-range"`
	// Passed as --service-node-port-range to kube-apiserver. Expects 'startPort-endPort' format. Eg. 30000-33000
	ServiceNodePortRange string `json:"serviceNodePortRange,omitempty" flag:"service-node-port-range"`
	// EtcdServers is a list of the etcd service to connect
	EtcdServers []string `json:"etcdServers,omitempty" flag:"etcd-servers"`
	// EtcdServersOverrides is per-resource etcd servers overrides, comma separated. The individual override format: group/resource#servers, where servers are http://ip:port, semicolon separated
	EtcdServersOverrides []string `json:"etcdServersOverrides,omitempty" flag:"etcd-servers-overrides"`
	// EtcdCAFile is the path to a ca certificate
	EtcdCAFile string `json:"etcdCaFile,omitempty" flag:"etcd-cafile"`
	// EtcdCertFile is the path to a certificate
	EtcdCertFile string `json:"etcdCertFile,omitempty" flag:"etcd-certfile"`
	// EtcdKeyFile is the path to a orivate key
	EtcdKeyFile string `json:"etcdKeyFile,omitempty" flag:"etcd-keyfile"`
	// TODO: Remove unused BasicAuthFile
	BasicAuthFile string `json:"basicAuthFile,omitempty" flag:"basic-auth-file"`
	// TODO: Remove unused ClientCAFile
	ClientCAFile string `json:"clientCAFile,omitempty" flag:"client-ca-file"`
	// TODO: Remove unused TLSCertFile
	TLSCertFile string `json:"tlsCertFile,omitempty" flag:"tls-cert-file"`
	// TODO: Remove unused TLSPrivateKeyFile
	TLSPrivateKeyFile string `json:"tlsPrivateKeyFile,omitempty" flag:"tls-private-key-file"`
	// TODO: Remove unused TokenAuthFile
	TokenAuthFile string `json:"tokenAuthFile,omitempty" flag:"token-auth-file"`
	// AllowPrivileged indicates if we can run privileged containers
	AllowPrivileged *bool `json:"allowPrivileged,omitempty" flag:"allow-privileged"`
	// APIServerCount is the number of api servers
	APIServerCount *int32 `json:"apiServerCount,omitempty" flag:"apiserver-count"`
	// RuntimeConfig is a series of keys/values are parsed into the `--runtime-config` parameters
	RuntimeConfig map[string]string `json:"runtimeConfig,omitempty" flag:"runtime-config"`
	// KubeletClientCertificate is the path of a certificate for secure communication between api and kubelet
	KubeletClientCertificate string `json:"kubeletClientCertificate,omitempty" flag:"kubelet-client-certificate"`
	// KubeletClientKey is the path of a private to secure communication between api and kubelet
	KubeletClientKey string `json:"kubeletClientKey,omitempty" flag:"kubelet-client-key"`
	// AnonymousAuth indicates if anonymous authentication is permitted
	AnonymousAuth *bool `json:"anonymousAuth,omitempty" flag:"anonymous-auth"`
	// KubeletPreferredAddressTypes is a list of the preferred NodeAddressTypes to use for kubelet connections
	KubeletPreferredAddressTypes []string `json:"kubeletPreferredAddressTypes,omitempty" flag:"kubelet-preferred-address-types"`
	// StorageBackend is the backend storage
	StorageBackend *string `json:"storageBackend,omitempty" flag:"storage-backend"`
	// The OpenID claim to use as the user name.
	// Note that claims other than the default ('sub') is not guaranteed to be unique and immutable.
	OIDCUsernameClaim *string `json:"oidcUsernameClaim,omitempty" flag:"oidc-username-claim"`
	// If provided, the name of a custom OpenID Connect claim for specifying user groups.
	// The claim value is expected to be a string or array of strings.
	OIDCGroupsClaim *string `json:"oidcGroupsClaim,omitempty" flag:"oidc-groups-claim"`
	// The URL of the OpenID issuer, only HTTPS scheme will be accepted.
	// If set, it will be used to verify the OIDC JSON Web Token (JWT).
	OIDCIssuerURL *string `json:"oidcIssuerURL,omitempty" flag:"oidc-issuer-url"`
	// The client ID for the OpenID Connect client, must be set if oidc-issuer-url is set.
	OIDCClientID *string `json:"oidcClientID,omitempty" flag:"oidc-client-id"`
	// If set, the OpenID server's certificate will be verified by one of the authorities in the oidc-ca-file
	OIDCCAFile *string `json:"oidcCAFile,omitempty" flag:"oidc-ca-file"`
	// The apiserver's client certificate used for outbound requests.
	ProxyClientCertFile *string `json:"proxyClientCertFile,omitempty" flag:"proxy-client-cert-file"`
	// The apiserver's client key used for outbound requests.
	ProxyClientKeyFile *string `json:"proxyClientKeyFile,omitempty" flag:"proxy-client-key-file"`
	// If set, all requests coming to the apiserver will be logged to this file.
	AuditLogPath *string `json:"auditLogPath,omitempty" flag:"audit-log-path"`
	// The maximum number of days to retain old audit log files based on the timestamp encoded in their filename.
	AuditLogMaxAge *int32 `json:"auditLogMaxAge,omitempty" flag:"audit-log-maxage"`
	// The maximum number of old audit log files to retain.
	AuditLogMaxBackups *int32 `json:"auditLogMaxBackups,omitempty" flag:"audit-log-maxbackup"`
	// The maximum size in megabytes of the audit log file before it gets rotated. Defaults to 100MB.
	AuditLogMaxSize *int32 `json:"auditLogMaxSize,omitempty" flag:"audit-log-maxsize"`
	// File with webhook configuration for token authentication in kubeconfig format. The API server will query the remote service to determine authentication for bearer tokens.
	AuthenticationTokenWebhookConfigFile *string `json:"authenticationTokenWebhookConfigFile,omitempty" flag:"authentication-token-webhook-config-file"`
	// The duration to cache responses from the webhook token authenticator. Default is 2m. (default 2m0s)
	AuthenticationTokenWebhookCacheTTL *metav1.Duration `json:"authenticationTokenWebhookCacheTtl,omitempty" flag:"authentication-token-webhook-cache-ttl"`
	// AuthorizationMode is the authorization mode the kubeapi is running in
	AuthorizationMode *string `json:"authorizationMode,omitempty" flag:"authorization-mode"`
	// AuthorizationRBACSuperUser is the name of the superuser for default rbac
	AuthorizationRBACSuperUser *string `json:"authorizationRbacSuperUser,omitempty" flag:"authorization-rbac-super-user"`
	// ExperimentalEncryptionProviderConfig enables encryption at rest for secrets.
	ExperimentalEncryptionProviderConfig *string `json:"experimentalEncryptionProviderConfig,omitempty" flag:"experimental-encryption-provider-config"`
}

// KubeControllerManagerConfig is the configuration for the controller
type KubeControllerManagerConfig struct {
	// Master is the url for the kube api master
	Master string `json:"master,omitempty" flag:"master"`
	// LogLevel is the defined logLevel
	LogLevel int32 `json:"logLevel,omitempty" flag:"v" flag-empty:"0"`
	// ServiceAccountPrivateKeyFile the location for a certificate for service account signing
	ServiceAccountPrivateKeyFile string `json:"serviceAccountPrivateKeyFile,omitempty" flag:"service-account-private-key-file"`
	// Image is the docker image to use
	Image string `json:"image,omitempty"`
	// CloudProvider is the provider for cloud services.
	CloudProvider string `json:"cloudProvider,omitempty" flag:"cloud-provider"`
	// ClusterName is the instance prefix for the cluster.
	ClusterName string `json:"clusterName,omitempty" flag:"cluster-name"`
	// ClusterCIDR is CIDR Range for Pods in cluster.
	ClusterCIDR string `json:"clusterCIDR,omitempty" flag:"cluster-cidr"`
	// AllocateNodeCIDRs enables CIDRs for Pods to be allocated and, if
	// ConfigureCloudRoutes is true, to be set on the cloud provider.
	AllocateNodeCIDRs *bool `json:"allocateNodeCIDRs,omitempty" flag:"allocate-node-cidrs"`
	// ConfigureCloudRoutes enables CIDRs allocated with to be configured on the cloud provider.
	ConfigureCloudRoutes *bool `json:"configureCloudRoutes,omitempty" flag:"configure-cloud-routes"`
	// rootCAFile is the root certificate authority will be included in service account's token secret. This must be a valid PEM-encoded CA bundle.
	RootCAFile string `json:"rootCAFile,omitempty" flag:"root-ca-file"`
	// LeaderElection defines the configuration of leader election client.
	LeaderElection *LeaderElectionConfiguration `json:"leaderElection,omitempty"`
	// ReconcilerSyncLoopPeriod is the amount of time the reconciler sync states loop
	// wait between successive executions. Is set to 1 min by kops by default
	AttachDetachReconcileSyncPeriod *metav1.Duration `json:"attachDetachReconcileSyncPeriod,omitempty" flag:"attach-detach-reconcile-sync-period"`
	// TerminatedPodGCThreshold is the number of terminated pods that can exist
	// before the terminated pod garbage collector starts deleting terminated pods.
	// If <= 0, the terminated pod garbage collector is disabled.
	TerminatedPodGCThreshold *int32 `json:"terminatedPodGCThreshold,omitempty" flag:"terminated-pod-gc-threshold"`

	// UseServiceAccountCredentials controls whether we use individual service account credentials for each controller.
	UseServiceAccountCredentials *bool `json:"useServiceAccountCredentials,omitempty" flag:"use-service-account-credentials"`
}

type CloudControllerManagerConfig struct {
	// Master is the url for the kube api master.
	Master string `json:"master,omitempty" flag:"master"`
	// LogLevel is the verbosity of the logs.
	LogLevel int32 `json:"logLevel,omitempty" flag:"v" flag-empty:"0"`
	// Image is the OCI image of the cloud controller manager.
	Image string `json:"image,omitempty"`
	// CloudProvider is the provider for cloud services.
	CloudProvider string `json:"cloudProvider,omitempty" flag:"cloud-provider"`
	// ClusterName is the instance prefix for the cluster.
	ClusterName string `json:"clusterName,omitempty" flag:"cluster-name"`
	// ClusterCIDR is CIDR Range for Pods in cluster.
	ClusterCIDR string `json:"clusterCIDR,omitempty" flag:"cluster-cidr"`
	// AllocateNodeCIDRs enables CIDRs for Pods to be allocated and, if
	// ConfigureCloudRoutes is true, to be set on the cloud provider.
	AllocateNodeCIDRs *bool `json:"allocateNodeCIDRs,omitempty" flag:"allocate-node-cidrs"`
	// ConfigureCloudRoutes enables CIDRs allocated with to be configured on the cloud provider.
	ConfigureCloudRoutes *bool `json:"configureCloudRoutes,omitempty" flag:"configure-cloud-routes"`
	// LeaderElection defines the configuration of leader election client.
	LeaderElection *LeaderElectionConfiguration `json:"leaderElection,omitempty"`
	// UseServiceAccountCredentials controls whether we use individual service account credentials for each controller.
	UseServiceAccountCredentials *bool `json:"useServiceAccountCredentials,omitempty" flag:"use-service-account-credentials"`
}

// KubeSchedulerConfig is the configuration for the kube-scheduler
type KubeSchedulerConfig struct {
	// Master is a url to the kube master
	Master string `json:"master,omitempty" flag:"master"`
	// LogLevel is the logging level
	LogLevel int32 `json:"logLevel,omitempty" flag:"v"`
	// Image is the docker image to use
	Image string `json:"image,omitempty"`
	// LeaderElection defines the configuration of leader election client.
	LeaderElection *LeaderElectionConfiguration `json:"leaderElection,omitempty"`
}

// LeaderElectionConfiguration defines the configuration of leader election
// clients for components that can run with leader election enabled.
type LeaderElectionConfiguration struct {
	// leaderElect enables a leader election client to gain leadership
	// before executing the main loop. Enable this when running replicated
	// components for high availability.
	LeaderElect *bool `json:"leaderElect,omitempty" flag:"leader-elect"`
}

// CloudConfiguration is defines the cloud provider configuration
type CloudConfiguration struct {
	// GCE cloud-config options
	Multizone          *bool   `json:"multizone,omitempty"`
	NodeTags           *string `json:"nodeTags,omitempty"`
	NodeInstancePrefix *string `json:"nodeInstancePrefix,omitempty"`
	// AWS cloud-config options
	DisableSecurityGroupIngress *bool   `json:"disableSecurityGroupIngress,omitempty"`
	ElbSecurityGroup            *string `json:"elbSecurityGroup,omitempty"`
	// vSphere cloud-config specs
	VSphereUsername      *string `json:"vSphereUsername,omitempty"`
	VSpherePassword      *string `json:"vSpherePassword,omitempty"`
	VSphereServer        *string `json:"vSphereServer,omitempty"`
	VSphereDatacenter    *string `json:"vSphereDatacenter,omitempty"`
	VSphereResourcePool  *string `json:"vSphereResourcePool,omitempty"`
	VSphereDatastore     *string `json:"vSphereDatastore,omitempty"`
	VSphereCoreDNSServer *string `json:"vSphereCoreDNSServer,omitempty"`
}
