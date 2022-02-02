package v2

// CommonConfig variables used in the templates
type CommonConfig struct {
	// The name of the application deployment.
	ApplicationName string `json:"applicationName,omitempty"`
	// +kubebuilder:validation:Format:=password
	// The password to use for keystore generation.
	KeyStorePassword string `json:"keyStorePassword,omitempty"`
	// The user to use for the admin.
	AdminUser string `json:"adminUser,omitempty"`
	// +kubebuilder:validation:Format:=password
	// The password to use for the adminUser.
	AdminPassword string `json:"adminPassword,omitempty"`
	// +kubebuilder:validation:Format:=password
	// The password to use for databases.
	DBPassword string `json:"dbPassword,omitempty"`
	// +kubebuilder:validation:Format:=password
	// The password to use for amq user.
	AMQPassword string `json:"amqPassword,omitempty"`
	// +kubebuilder:validation:Format:=password
	// The password to use for amq cluster user.
	AMQClusterPassword string `json:"amqClusterPassword,omitempty"`
	// If set to true, plain text routes will be configured instead using SSL
	DisableSsl bool `json:"disableSsl,omitempty"`
	// Startup strategy for Console and Kieserver
	StartupStrategy *StartupStrategy `json:"startupStrategy,omitempty"`
}

// VersionConfigs ...
type VersionConfigs struct {
	APIVersion           string `json:"apiVersion,omitempty"`
	OseCliImageURL       string `json:"oseCliImageURL,omitempty"`
	OseCliComponent      string `json:"oseCliComponent,omitempty"`
	BrokerImageContext   string `json:"brokerImageContext,omitempty"`
	BrokerImage          string `json:"brokerImage,omitempty"`
	BrokerImageTag       string `json:"brokerImageTag,omitempty"`
	BrokerImageURL       string `json:"brokerImageURL,omitempty"`
	BrokerComponent      string `json:"brokerComponent,omitempty"`
	DatagridImageContext string `json:"datagridImageContext,omitempty"`
	DatagridImage        string `json:"datagridImage,omitempty"`
	DatagridImageTag     string `json:"datagridImageTag,omitempty"`
	DatagridImageURL     string `json:"datagridImageURL,omitempty"`
	DatagridComponent    string `json:"datagridComponent,omitempty"`
	MySQLImageURL        string `json:"mySQLImageURL,omitempty"`
	MySQLComponent       string `json:"mySQLComponent,omitempty"`
	PostgreSQLImageURL   string `json:"postgreSQLImageURL,omitempty"`
	PostgreSQLComponent  string `json:"postgreSQLComponent,omitempty"`
}

// Replicas contains replica settings
type Replicas struct {
	Replicas  int32 `json:"replicas,omitempty"`
	DenyScale bool  `json:"denyScale,omitempty"`
}

// KieAppRegistry defines the registry that should be used for rhpam images
type KieAppRegistry struct {
	// Image registry's base 'url:port'. e.g. registry.example.com:5000. Defaults to 'registry.redhat.io'.
	Registry string `json:"registry,omitempty"`
	// A flag used to indicate the specified registry is insecure. Defaults to 'false'.
	Insecure bool `json:"insecure,omitempty"`
}