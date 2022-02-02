package v2

import corev1 "k8s.io/api/core/v1"

// KieAppBuildObject Data to define how to build an application from source
type KieAppBuildObject struct {
	// Env set environment variables for BuildConfigs
	Env []corev1.EnvVar `json:"env,omitempty"`
	// The Maven GAV to deploy, e.g., rhpam-kieserver-library=org.openshift.quickstarts:rhpam-kieserver-library:1.5.0-SNAPSHOT
	KieServerContainerDeployment string `json:"kieServerContainerDeployment,omitempty"`
	// Disable Maven pull dependencies for immutable KIE Server configurations for S2I and pre built kjars. Useful for pre-compiled kjar.
	DisablePullDeps bool `json:"disablePullDeps,omitempty"`
	// Disable Maven KIE Jar verification. It is recommended to test the kjar manually before disabling this verification.
	DisableKCVerification bool      `json:"disableKCVerification,omitempty"`
	GitSource             GitSource `json:"gitSource,omitempty"`
	// Maven mirror to use for S2I builds
	MavenMirrorURL string `json:"mavenMirrorURL,omitempty"`
	// List of directories from which archives will be copied into the deployment folder. If unspecified, all archives in /target will be copied.
	ArtifactDir string `json:"artifactDir,omitempty"`
	// +kubebuilder:validation:MinItems:=1
	Webhooks []WebhookSecret `json:"webhooks,omitempty"`
	From     *ImageObjRef    `json:"from,omitempty"`
	// ImageStreamTag definition for the image containing the drivers and configuration. For example, custom-driver-image:7.7.0.
	ExtensionImageStreamTag string `json:"extensionImageStreamTag,omitempty"`
	// Namespace within which the ImageStream definition for the image containing the drivers and configuration is located. Defaults to openshift namespace.
	ExtensionImageStreamTagNamespace string `json:"extensionImageStreamTagNamespace,omitempty"`
	// Full path to the directory within the extensions image where the extensions are located (e.g. install.sh, modules/, etc.).  Defaults to '/extension'. Do not change it unless it is necessary.
	ExtensionImageInstallDir string `json:"extensionImageInstallDir,omitempty"`
}

// GitSource Git coordinates to locate the source code to build
type GitSource struct {
	// +kubebuilder:validation:Required
	// Git URI for the s2i source
	URI string `json:"uri"`
	// +kubebuilder:validation:Required
	// Branch to use in the git repository
	Reference string `json:"reference"`
	// Context/subdirectory where the code is located, relatively to repo root
	ContextDir string `json:"contextDir,omitempty"`
}

// WebhookType literal type to distinguish between different types of Webhooks
type WebhookType string

const (
	// GitHubWebhook GitHub webhook
	GitHubWebhook WebhookType = "GitHub"
	// GenericWebhook Generic webhook
	GenericWebhook WebhookType = "Generic"
)

// WebhookSecret Secret to use for a given webhook
type WebhookSecret struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum:=GitHub;Generic
	Type WebhookType `json:"type"`
	// +kubebuilder:validation:Required
	Secret string `json:"secret"`
}

// GitHooksVolume GitHooks volume configuration
type GitHooksVolume struct {
	// Absolute path where the gitHooks folder will be mounted.
	MountPath string  `json:"mountPath,omitempty"`
	From      *ObjRef `json:"from,omitempty"`
	// Secret to use for ssh key and known hosts file. The secret must contain two files: id_rsa and known_hosts.
	SSHSecret string `json:"sshSecret,omitempty"`
}

// BuildTemplate build variables used in the templates
type BuildTemplate struct {
	From                         ImageObjRef `json:"from,omitempty"`
	GitSource                    GitSource   `json:"gitSource,omitempty"`
	GitHubWebhookSecret          string      `json:"githubWebhookSecret,omitempty"`
	GenericWebhookSecret         string      `json:"genericWebhookSecret,omitempty"`
	KieServerContainerDeployment string      `json:"kieServerContainerDeployment,omitempty"`
	DisablePullDeps              bool        `json:"disablePullDeps,omitempty"`
	DisableKCVerification        bool        `json:"disableKCVerification,omitempty"`
	MavenMirrorURL               string      `json:"mavenMirrorURL,omitempty"`
	ArtifactDir                  string      `json:"artifactDir,omitempty"`
	// Extension image configuration which provides custom jdbc drivers to be used
	// by KieServer.
	ExtensionImageStreamTag          string `json:"extensionImageStreamTag,omitempty"`
	ExtensionImageStreamTagNamespace string `json:"extensionImageStreamTagNamespace,omitempty"`
	ExtensionImageInstallDir         string `json:"extensionImageInstallDir,omitempty"`
}

