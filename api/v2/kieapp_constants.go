package v2

const (
	// RhpamTrial RHPAM Trial environment
	RhpamTrial EnvironmentType = "rhpam-trial"
	// RhpamProduction RHPAM Production environment
	RhpamProduction EnvironmentType = "rhpam-production"
	// RhpamProductionImmutable RHPAM Production Immutable environment
	RhpamProductionImmutable EnvironmentType = "rhpam-production-immutable"
	// RhpamAuthoring RHPAM Authoring environment
	RhpamAuthoring EnvironmentType = "rhpam-authoring"
	// RhpamAuthoringHA RHPAM Authoring HA environment
	RhpamAuthoringHA EnvironmentType = "rhpam-authoring-ha"
	// RhpamStandaloneDashbuilder RHPAM Standalone Dashbuilder environment
	RhpamStandaloneDashbuilder EnvironmentType = "rhpam-standalone-dashbuilder"
	// RhdmTrial RHDM Trial environment
	RhdmTrial EnvironmentType = "rhdm-trial"
	// RhdmAuthoring RHDM Authoring environment
	RhdmAuthoring EnvironmentType = "rhdm-authoring"
	// RhdmAuthoringHA RHDM Authoring HA environment
	RhdmAuthoringHA EnvironmentType = "rhdm-authoring-ha"
	// RhdmProductionImmutable RHDM Production Immutable environment
	RhdmProductionImmutable EnvironmentType = "rhdm-production-immutable"
)

// EnvironmentType describes a possible application environment
type EnvironmentType string

// EnvironmentConstants stores both the App and Replica Constants for a given environment
type EnvironmentConstants struct {
	App      AppConstants     `json:"app,omitempty"`
	Replica  ReplicaConstants `json:"replica,omitempty"`
	Database *DatabaseObject  `json:"database,omitempty"`
	Jms      *KieAppJmsObject `json:"jms,omitempty"`
}

// StartupStrategies supported values
const (
	OpenshiftStartupStrategy  = "OpenShiftStartupStrategy"
	ControllerStartupStrategy = "ControllerBasedStartupStrategy"
)

// ReplicaConstants contains the default replica amounts for a component in a given environment type
type ReplicaConstants struct {
	Console     Replicas `json:"console,omitempty"`
	Dashbuilder Replicas `json:"dashbuilder,omitempty"`
	Server      Replicas `json:"server,omitempty"`
	SmartRouter Replicas `json:"smartRouter,omitempty"`
}

// AppConstants data type to store application deployment constants
type AppConstants struct {
	Product      string `json:"name,omitempty"`
	Prefix       string `json:"prefix,omitempty"`
	ImageName    string `json:"imageName,omitempty"`
	ImageVar     string `json:"imageVar,omitempty"`
	MavenRepo    string `json:"mavenRepo,omitempty"`
	FriendlyName string `json:"friendlyName,omitempty"`
}

// TemplateConstants constant values that are used within the different configuration templates
type TemplateConstants struct {
	Product              string `json:"product,omitempty"`
	Major                string `json:"major,omitempty"`
	Minor                string `json:"minor,omitempty"`
	Micro                string `json:"micro,omitempty"`
	MavenRepo            string `json:"mavenRepo,omitempty"`
	KeystoreVolumeSuffix string `json:"keystoreVolumeSuffix"`
	DatabaseVolumeSuffix string `json:"databaseVolumeSuffix"`
	OseCliImageURL       string `json:"oseCliImageURL,omitempty"`
	BrokerImageContext   string `json:"brokerImageContext"`
	BrokerImage          string `json:"brokerImage"`
	BrokerImageTag       string `json:"brokerImageTag"`
	DatagridImageContext string `json:"datagridImageContext"`
	DatagridImage        string `json:"datagridImage"`
	DatagridImageTag     string `json:"datagridImageTag"`
	MySQLImageURL        string `json:"mySQLImageURL"`
	PostgreSQLImageURL   string `json:"postgreSQLImageURL"`
	BrokerImageURL       string `json:"brokerImageURL,omitempty"`
	DatagridImageURL     string `json:"datagridImageURL,omitempty"`
	RoleMapperVolume     string `json:"roleMapperVolume"`
	GitHooksVolume       string `json:"gitHooksVolume,omitempty"`
	GitHooksSSHSecret    string `json:"gitHooksSSHSecret,omitempty"`
}
