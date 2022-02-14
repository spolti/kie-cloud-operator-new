package v2

import (
	oappsv1 "github.com/openshift/api/apps/v1"
	buildv1 "github.com/openshift/api/build/v1"
	oimagev1 "github.com/openshift/api/image/v1"
	routev1 "github.com/openshift/api/route/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	//	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KieApp is the Schema for the kieapps API
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=kieapps,scope=Namespaced
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Version",type=string,JSONPath=`.status.version`,description="The version of the application deployment"
// +kubebuilder:printcolumn:name="Environment",type=string,JSONPath=`.spec.environment`,description="The name of the environment used as a baseline"
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.phase`,description="The status of the KieApp deployment"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type KieApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   KieAppSpec   `json:"spec"`
	Status KieAppStatus `json:"status,omitempty"`
}

// KieAppSpec defines the desired state of KieApp
type KieAppSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum:=rhdm-authoring-ha;rhdm-authoring;rhdm-production-immutable;rhdm-trial;rhpam-authoring-ha;rhpam-authoring;rhpam-production-immutable;rhpam-production;rhpam-standalone-dashbuilder;rhpam-trial
	// The name of the environment used as a baseline
	Environment EnvironmentType `json:"environment"`
	// If required imagestreams are missing in both the 'openshift' and local namespaces, the operator will create said imagestreams locally using the registry specified here.
	ImageRegistry *KieAppRegistry `json:"imageRegistry,omitempty"`
	// Configuration of the RHPAM components
	Objects KieAppObjects `json:"objects,omitempty"`
	// Specify the level of product upgrade that should be allowed when an older product version is detected
	Upgrades KieAppUpgrades `json:"upgrades,omitempty"`
	// Set true to enable image tags, disabled by default. This will leverage image tags instead of the image digests.
	UseImageTags bool `json:"useImageTags,omitempty"`
	// Defines which truststore is used by the console, kieservers, smartrouter, and dashbuilder
	Truststore *KieAppTruststore `json:"truststore,omitempty"`
	// The version of the application deployment.
	Version      string            `json:"version,omitempty"`
	CommonConfig CommonConfig      `json:"commonConfig,omitempty"`
	Auth         *KieAppAuthObject `json:"auth,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KieAppList contains a list of KieApp
type KieAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KieApp `json:"items"`
}

// KieAppObjects KIE App deployment objects
type KieAppObjects struct {
	Console *ConsoleObject `json:"console,omitempty"`
	// Configuration of the each individual KIE server
	Servers          []KieServerSet          `json:"servers,omitempty"`
	SmartRouter      *SmartRouterObject      `json:"smartRouter,omitempty"`
	ProcessMigration *ProcessMigrationObject `json:"processMigration,omitempty"`
	Dashbuilder      *DashbuilderObject      `json:"dashbuilder,omitempty"`
}

// KieAppUpgrades KIE App product upgrade flags
type KieAppUpgrades struct {
	// Set true to enable automatic micro version product upgrades, it is disabled by default.
	Enabled bool `json:"enabled,omitempty"`
	// Set true to enable automatic minor product version upgrades, it is disabled by default. Requires spec.upgrades.enabled to be true.
	Minor bool `json:"minor,omitempty"`
}

// KieAppObject generic object definition
type KieAppObject struct {
	Env []corev1.EnvVar `json:"env,omitempty"`
	// Replicas to set for the DeploymentConfig
	Replicas  *int32                       `json:"replicas,omitempty"`
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
	// KeystoreSecret secret name
	KeystoreSecret string `json:"keystoreSecret,omitempty"`
	// ImageContext The image context to use  e.g. rhpam-7, this param is optional for custom image.
	ImageContext string `json:"imageContext,omitempty"`
	// Image The image to use e.g. rhpam-<app>-rhel8, this param is optional for custom image.
	Image string `json:"image,omitempty"`
	// ImageTag The image tag to use e.g. 7.13.0, this param is optional for custom image.
	ImageTag string `json:"imageTag,omitempty"`
	// StorageClassName The storageClassName to use for kie pvc's.
	StorageClassName string `json:"storageClassName,omitempty"`
	// RouteHostname will define the route.spec.host value
	RouteHostname string `json:"routeHostname,omitempty"`
}

type Environment struct {
	Console          CustomObject   `json:"console,omitempty"`
	SmartRouter      CustomObject   `json:"smartRouter,omitempty"`
	Servers          []CustomObject `json:"servers,omitempty"`
	ProcessMigration CustomObject   `json:"processMigration,omitempty"`
	Dashbuilder      CustomObject   `json:"dashbuilder,omitempty"`
	Databases        []CustomObject `json:"databases,omitempty"`
	Others           []CustomObject `json:"others,omitempty"`
}

type CustomObject struct {
	Omit                   bool                           `json:"omit,omitempty"`
	PersistentVolumeClaims []corev1.PersistentVolumeClaim `json:"persistentVolumeClaims,omitempty"`
	ServiceAccounts        []corev1.ServiceAccount        `json:"serviceAccounts,omitempty"`
	Secrets                []corev1.Secret                `json:"secrets,omitempty"`
	Roles                  []rbacv1.Role                  `json:"roles,omitempty"`
	RoleBindings           []rbacv1.RoleBinding           `json:"roleBindings,omitempty"`
	DeploymentConfigs      []oappsv1.DeploymentConfig     `json:"deploymentConfigs,omitempty"`
	StatefulSets           []appsv1.StatefulSet           `json:"statefulSets,omitempty"`
	BuildConfigs           []buildv1.BuildConfig          `json:"buildConfigs,omitempty"`
	ImageStreams           []oimagev1.ImageStream         `json:"imageStreams,omitempty"`
	Services               []corev1.Service               `json:"services,omitempty"`
	Routes                 []routev1.Route                `json:"routes,omitempty"`
	ConfigMaps             []corev1.ConfigMap             `json:"configMaps,omitempty"`
}

type EnvTemplate struct {
	*CommonConfig     `json:",inline"`
	Console           ConsoleTemplate          `json:"console,omitempty"`
	Servers           []ServerTemplate         `json:"servers,omitempty"`
	SmartRouter       SmartRouterTemplate      `json:"smartRouter,omitempty"`
	Auth              AuthTemplate             `json:"auth,omitempty"`
	ProcessMigration  ProcessMigrationTemplate `json:"processMigration,omitempty"`
	Dashbuilder       DashbuilderTemplate      `json:"dashbuilder,omitempty"`
	Databases         []DatabaseTemplate       `json:"databases,omitempty"`
	Constants         TemplateConstants        `json:"constants,omitempty"`
	OpenshiftCaBundle bool                     `json:"openshiftCaBundle,omitempty"`
	RouteProtocol     string                   `json:"routeProtocol,omitempty"`
}

// ImageObjRef contains enough information to let you inspect or modify the referred object.
type ImageObjRef struct {
	// Kind of the referent.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum:=ImageStreamTag;DockerImage
	Kind            string `json:"kind" protobuf:"bytes,1,opt,name=kind"`
	ObjectReference `json:",inline"`
}

// ObjectReference contains enough information to let you inspect or modify the referred object.
type ObjectReference struct {
	// Namespace of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`
	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// +kubebuilder:validation:Required
	Name string `json:"name" protobuf:"bytes,3,opt,name=name"`
	// UID of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
	// +optional
	UID types.UID `json:"uid,omitempty" protobuf:"bytes,4,opt,name=uid,casttype=k8s.io/apimachinery/pkg/types.UID"`
	// API version of the referent.
	// +optional
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,5,opt,name=apiVersion"`
	// Specific resourceVersion to which this reference is made, if any.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
	// +optional
	ResourceVersion string `json:"resourceVersion,omitempty" protobuf:"bytes,6,opt,name=resourceVersion"`

	// If referring to a piece of an object instead of an entire object, this string
	// should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
	// For example, if the object reference is to a container within a pod, this would take on a value like:
	// "spec.containers{name}" (where "name" refers to the name of the container that triggered
	// the event) or if no container name is specified "spec.containers[2]" (container with
	// index 2 in this pod). This syntax is chosen only to have some well-defined way of
	// referencing a part of an object.
	// TODO: this design is not final and this field is subject to change in the future.
	// +optional
	FieldPath string `json:"fieldPath,omitempty" protobuf:"bytes,7,opt,name=fieldPath"`
}

// CORSFiltersObject CORS Cross Origin Resource Sharing configuration to be used by the KieApp for the KIE Server, workbench and RHPAM Dashbuilder
type CORSFiltersObject struct {
	//Enable CORS setting with the default values, default is false
	Default bool `json:"default,omitempty"`
	//Access control Response Headers Filters separated by comma, default is: AC_ALLOW_ORIGIN,AC_ALLOW_METHODS,AC_ALLOW_HEADERS,AC_ALLOW_CREDENTIALS,AC_MAX_AGE
	Filters string `json:"filters,omitempty"`

	//Access Control Origin Response Header Filter Header Name, default is Access-Control-Allow-Origin
	AllowOriginName string `json:"allowOriginName,omitempty"`
	//Access Control Origin Response Header  Filter Header Value, default is: *
	AllowOriginValue string `json:"allowOriginValue,omitempty"`

	//Access Control Allow Methods Response Header Filter Header Name, default is: Access-Control-Allow-Methods
	AllowMethodsName string `json:"allowMethodsName,omitempty"`
	//Access Control Allow Methods Response Headers Filter Header Value, default is: GET, POST, OPTIONS, PUT
	AllowMethodsValue string `json:"allowMethodsValue,omitempty"`

	//Access Control Allow Headers Filter Header Name, default is: Access-Control-Allow-Headers
	AllowHeadersName string `json:"allowHeadersName,omitempty"`
	//Access Control Allow Headers Filter Header Value, default is: Accept, Authorization, Content-Type, X-Requested-With
	AllowHeadersValue string `json:"allowHeadersValue,omitempty"`

	//Access Control Allow Credentials Filter Header Name, default is: Access-Control-Allow-Credentials
	AllowCredentialsName string `json:"allowCredentialsName,omitempty"`
	//Access Control Allow Credentials Filter Header Value, default is: true
	AllowCredentialsValue *bool `json:"allowCredentialsValue,omitempty"`

	//Access Control Max Age Filter Header Name, default is: Access-Control-Max-Age
	MaxAgeName string `json:"maxAgeName,omitempty"`
	//Access Control Max Age Filter Header Value, default is: 1
	MaxAgeValue *int32 `json:"maxAgeValue,omitempty"`
}

// StartupStrategy ...
type StartupStrategy struct {
	// StartupStrategy to use. When set to OpenShiftStartupStrategy, allows KIE server to start up independently used shared state from OpenShift API service, option is ControllerBasedStartupStrategy, default is OpenShiftStartupStrategy
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum:=OpenShiftStartupStrategy;ControllerBasedStartupStrategy
	// +kubebuilder:validation:default:OpenShiftStartupStrategy
	StrategyName string `json:"strategyName,omitempty"`
	// Controller Template Cache TTL to use when the OpenShiftStartupStrategy is choosed and Business Central is deployed, default is 5000
	ControllerTemplateCacheTTL *int `json:"controllerTemplateCacheTTL,omitempty"`
}

type OpenShiftObject interface {
	client.Object
}

func init() {
	SchemeBuilder.Register(&KieApp{}, &KieAppList{})
}
