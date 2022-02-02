package v2

// KieServerSet KIE Server configuration for a single set, or for multiple sets if deployments is set to >1
type KieServerSet struct {
	// +kubebuilder:validation:Format:=int
	// Number of Server sets that will be deployed
	Deployments *int `json:"deployments,omitempty"` // Number of KieServer DeploymentConfigs (defaults to 1)
	// Server name
	Name string `json:"name,omitempty"`
	// Server ID
	ID           string             `json:"id,omitempty"`
	From         *ImageObjRef       `json:"from,omitempty"`
	Build        *KieAppBuildObject `json:"build,omitempty"` // S2I Build configuration
	SSOClient    *SSOAuthClient     `json:"ssoClient,omitempty"`
	KieAppObject `json:",inline"`
	Database     *DatabaseObject  `json:"database,omitempty"`
	Jms          *KieAppJmsObject `json:"jms,omitempty"`
	Jvm          *JvmObject       `json:"jvm,omitempty"`
	// PersistRepos enables persistent volumes for KIE Server's kie and maven repositories
	PersistRepos bool `json:"persistRepos,omitempty"`
	// ServersM2PvSize the desired size of the Maven persistent volume, the size of the files on this directory
	//can grow fast as all dependencies for KIE Containers will be stored there. Defaults to 1Gi
	ServersM2PvSize string `json:"serversM2PvSize,omitempty"`
	// ServersKiePvSize the desired size of the KIE local repository persistent volume. Defaults to 10Mi
	ServersKiePvSize string `json:"serversKiePvSize,omitempty"`
	// JbpmCluster Enable the KIE Server Jbpm clustering for processes fail-over, it could increase the number of kieservers
	JbpmCluster            bool                          `json:"jbpmCluster,omitempty"`
	Kafka                  *KafkaExtObject               `json:"kafka,omitempty"`
	KafkaJbpmEventEmitters *KafkaJBPMEventEmittersObject `json:"kafkaJbpmEventEmitters,omitempty"`
	Cors                   *CORSFiltersObject            `json:"cors,omitempty"`
	// MDBMaxSession number of KIE Executor sessions
	MDBMaxSession *int `json:"MDBMaxSession,omitempty"`
}

// ServerTemplate contains all the variables used in the yaml templates
type ServerTemplate struct {
	OmitImageStream  bool              `json:"omitImageStream"`
	OmitConsole      bool              `json:"omitConsole"`
	KieName          string            `json:"kieName,omitempty"`
	KieServerID      string            `json:"kieServerID,omitempty"`
	Replicas         int32             `json:"replicas,omitempty"`
	SSOAuthClient    SSOAuthClient     `json:"ssoAuthClient,omitempty"`
	From             ImageObjRef       `json:"from,omitempty"`
	ImageURL         string            `json:"imageURL,omitempty"`
	Build            BuildTemplate     `json:"build,omitempty"`
	KeystoreSecret   string            `json:"keystoreSecret,omitempty"`
	Database         DatabaseObject    `json:"database,omitempty"`
	Jms              KieAppJmsObject   `json:"jms,omitempty"`
	SmartRouter      SmartRouterObject `json:"smartRouter,omitempty"`
	Jvm              JvmObject         `json:"jvm,omitempty"`
	StorageClassName string            `json:"storageClassName,omitempty"`
	RouteHostname    string            `json:"routeHostname,omitempty"`
	PersistRepos     bool              `json:"persistRepos,omitempty"`
	ServersM2PvSize  string            `json:"serversM2PvSize,omitempty"`
	ServersKiePvSize string            `json:"serversKiePvSize,omitempty"`
	// JbpmCluster Enable the KIE Server Jbpm clustering for processes fail-over, it could increase the number of kieservers
	JbpmCluster            bool                          `json:"jbpmCluster,omitempty"`
	Kafka                  *KafkaExtObject               `json:"kafka,omitempty"`
	KafkaJbpmEventEmitters *KafkaJBPMEventEmittersObject `json:"kafkaJbpmEventEmitters,omitempty"`
	Cors                   *CORSFiltersObject            `json:"cors,omitempty"`
	StartupStrategy        *StartupStrategy              `json:"startupStrategy,omitempty"`
	// MDBMaxSession number of KIE Executor sessions
	MDBMaxSession *int `json:"MDBMaxSession,omitempty"`
}
