package v2

// ProcessMigrationObject configuration of the RHPAM PIM
type ProcessMigrationObject struct {
	KieAppObject  `json:",inline"`
	Jvm           *JvmObject                     `json:"jvm,omitempty"`
	Database      ProcessMigrationDatabaseObject `json:"database,omitempty"`
	RouteHostname string                         `json:"routeHostname,omitempty"`
	// If empty the CommonConfig.AdminUser will be used
	Username string `json:"username,omitempty"`
	// If empty the CommonConfig.AdminPassword will be used
	Password string `json:"password,omitempty"`
	// ExtraClassPath Allows to add extra jars to the application classpath separated by colon. Needs to be mounted
	// on the image before.
	ExtraClassPath string `json:"extraClassPath,omitempty"`
}

// ProcessMigrationTemplate ...
type ProcessMigrationTemplate struct {
	KieAppObject     `json:",inline"`
	OmitImageStream  bool                           `json:"omitImageStream"`
	ImageURL         string                         `json:"imageURL,omitempty"`
	KieServerClients []KieServerClient              `json:"kieServerClients,omitempty"`
	Jvm              JvmObject                      `json:"jvm,omitempty"`
	Database         ProcessMigrationDatabaseObject `json:"database,omitempty"`
	RouteHostname    string                         `json:"routeHostname,omitempty"`
	// PIM Admin username. If empty the CommonConfig.AdminUser will be used
	Username string `json:"username,omitempty"`
	// PIM Admin password. If empty the CommonConfig.AdminPassword will be used
	Password string `json:"password,omitempty"`
	// ExtraClassPath Allows to add extra jars to the application classpath separated by colon. Needs to be mounted
	// on the image before.
	ExtraClassPath string `json:"extraClassPath,omitempty"`
}

// KieServerClient ...
type KieServerClient struct {
	Host     string `json:"host,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// ProcessMigrationDatabaseObject Defines how a Process Migration server will manage
// and create a new Database or connect to an existing one
type ProcessMigrationDatabaseObject struct {
	InternalDatabaseObject `json:",inline"`
	ExternalConfig         *CommonExtDBObjectRequiredURL `json:"externalConfig,omitempty"`
}
