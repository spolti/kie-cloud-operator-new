package v2

// DashbuilderObject configuration of the RHPAM Dashbuilder
type DashbuilderObject struct {
	KieAppObject `json:",inline"`
	SSOClient    *SSOAuthClient     `json:"ssoClient,omitempty"`
	Jvm          *JvmObject         `json:"jvm,omitempty"`
	Config       *DashbuilderConfig `json:"config,omitempty"`
	Cors         *CORSFiltersObject `json:"cors,omitempty"`
}

// DashbuilderTemplate contains all the variables used in the yaml templates
type DashbuilderTemplate struct {
	OmitImageStream  bool              `json:"omitImageStream"`
	Name             string            `json:"name,omitempty"`
	Replicas         int32             `json:"replicas,omitempty"`
	SSOAuthClient    SSOAuthClient     `json:"ssoAuthClient,omitempty"`
	ImageContext     string            `json:"imageContext,omitempty"`
	Image            string            `json:"image,omitempty"`
	ImageTag         string            `json:"imageTag,omitempty"`
	ImageURL         string            `json:"imageURL,omitempty"`
	KeystoreSecret   string            `json:"keystoreSecret,omitempty"`
	Database         DatabaseObject    `json:"database,omitempty"`
	Jvm              JvmObject         `json:"jvm,omitempty"`
	StorageClassName string            `json:"storageClassName,omitempty"`
	RouteHostname    string            `json:"routeHostname,omitempty"`
	Config           DashbuilderConfig `json:"config,omitempty"`
	Cors             CORSFiltersObject `json:"cors,omitempty"`
}

// DashbuilderConfig holds all configurations that can be applied to the Dashbuilder env
type DashbuilderConfig struct {
	// Enables integration with Business Central
	// +optional
	EnableBusinessCentral bool `json:"enableBusinessCentral,omitempty"`
	// Enables integration with KIE Server
	// +optional
	EnableKieServer bool `json:"enableKieServer,omitempty"`
	// Allow download of external (remote) files into runtime. Default value is false
	// +optional
	AllowExternalFileRegister *bool `json:"allowExternalFileRegister,omitempty"`
	// Components will be partitioned by the Runtime Model ID. Default value is true
	// +optional
	ComponentPartition *bool `json:"componentPartition,omitempty"`
	// Datasets IDs will partitioned by the Runtime Model ID. Default value is true
	// +optional
	DataSetPartition *bool `json:"dataSetPartition,omitempty"`
	// Set a static dashboard to run with runtime. When this property is set no new imports are allowed.
	// +optional
	ImportFileLocation string `json:"importFileLocation,omitempty"`
	// Make Dashbuilder not ephemeral. If ImportFileLocation is set PersistentConfigs will be ignored.
	// Default value is true.
	// +optional
	PersistentConfigs *bool `json:"persistentConfigs,omitempty"`
	// Base Directory where dashboards ZIPs are stored. If PersistentConfigs is enabled and ImportsBaseDir is not
	// pointing to a already existing PV the /opt/kie/dashbuilder/imports directory will be used. If ImportFileLocation is set
	// ImportsBaseDir will be ignored.
	// +optional
	ImportsBaseDir string `json:"importsBaseDir,omitempty"`
	// Allows Runtime to check model last update in FS to update its content. Default value is true.
	// +optional
	ModelUpdate *bool `json:"modelUpdate,omitempty"`
	// When enabled will also remove actual model file from file system. Default value is false.
	// +optional
	ModelFileRemoval *bool `json:"modelFileRemoval,omitempty"`
	// Runtime will always allow use of new imports (multi tenancy). Default value is false.
	// +optional
	RuntimeMultipleImport *bool `json:"runtimeMultipleImport,omitempty"`
	// Limits the size of uploaded dashboards (in kb). Default value is 10485760 kb.
	// +optional
	UploadSize *int64 `json:"uploadSize,omitempty"`
	// When set to true enables external components.
	// +optional
	ComponentEnable *bool `json:"componentEnable,omitempty"`
	// Base Directory where dashboards ZIPs are stored. If PersistentConfigs is enabled and ExternalCompDir is not
	// pointing to a already existing PV the /opt/kie/dashbuilder/components directory will be used.
	// +optional
	ExternalCompDir string `json:"externalCompDir,omitempty"`
	// Properties file with Dashbuilder configurations, if set, uniq properties will be appended and, if a property
	// is set mode than once, the one from this property file will be used.
	// +optional
	ConfigMapProps string `json:"configMapProps,omitempty"`
	// Defines the KIE Server Datasets access configurations
	// +optional
	KieServerDataSets []KieServerDataSetOrTemplate `json:"kieServerDataSets,omitempty"`
	// Defines the KIE Server Templates access configurations
	// +optional
	KieServerTemplates []KieServerDataSetOrTemplate `json:"kieServerTemplates,omitempty"`
}

type KieServerDataSetOrTemplate struct {
	Name         string `json:"name,omitempty"`
	Location     string `json:"location,omitempty"`
	User         string `json:"user,omitempty"`
	Password     string `json:"password,omitempty"`
	Token        string `json:"token,omitempty"`
	ReplaceQuery string `json:"replaceQuery,omitempty"`
}
