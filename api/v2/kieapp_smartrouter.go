package v2

// SmartRouterTemplate contains all the variables used in the yaml templates
type SmartRouterTemplate struct {
	OmitImageStream  bool      `json:"omitImageStream"`
	Replicas         int32     `json:"replicas,omitempty"`
	KeystoreSecret   string    `json:"keystoreSecret,omitempty"`
	Protocol         string    `json:"protocol,omitempty"`
	UseExternalRoute bool      `json:"useExternalRoute,omitempty"`
	ImageContext     string    `json:"imageContext,omitempty"`
	Image            string    `json:"image,omitempty"`
	ImageTag         string    `json:"imageTag,omitempty"`
	ImageURL         string    `json:"imageURL,omitempty"`
	StorageClassName string    `json:"storageClassName,omitempty"`
	RouteHostname    string    `json:"routeHostname,omitempty"`
	Jvm              JvmObject `json:"jvm,omitempty"`
}

// SmartRouterObject configuration of the RHPAM smart router
type SmartRouterObject struct {
	KieAppObject `json:",inline"`
	// +kubebuilder:validation:Enum:=http;https
	// Smart Router protocol, if no value is provided, http is the default protocol.
	Protocol string `json:"protocol,omitempty"`
	// If enabled, Business Central will use the external smartrouter route to communicate with it. Note that, valid SSL certificates should be used.
	UseExternalRoute bool       `json:"useExternalRoute,omitempty"`
	Jvm              *JvmObject `json:"jvm,omitempty"`
}
