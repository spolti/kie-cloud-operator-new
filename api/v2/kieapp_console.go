package v2

// ConsoleObject configuration of the RHPAM workbench
type ConsoleObject struct {
	KieAppObject `json:",inline"`
	SSOClient    *SSOAuthClient     `json:"ssoClient,omitempty"`
	GitHooks     *GitHooksVolume    `json:"gitHooks,omitempty"`
	Jvm          *JvmObject         `json:"jvm,omitempty"`
	PvSize       string             `json:"pvSize,omitempty"`
	Cors         *CORSFiltersObject `json:"cors,omitempty"`
	DataGridAuth *DataGridAuth      `json:"dataGridAuth,omitempty"`
}

// DataGridAuth ...
type DataGridAuth struct {
	// The user to use for datagrid
	Username string `json:"username,omitempty"`
	// +kubebuilder:validation:Format:=password
	// The password to use for datagrid user
	Password string `json:"password,omitempty"`
}

// ConsoleTemplate contains all the variables used in the yaml templates
type ConsoleTemplate struct {
	OmitImageStream     bool              `json:"omitImageStream"`
	SSOAuthClient       SSOAuthClient     `json:"ssoAuthClient,omitempty"`
	Name                string            `json:"name,omitempty"`
	Replicas            int32             `json:"replicas,omitempty"`
	ImageContext        string            `json:"imageContext,omitempty"`
	Image               string            `json:"image,omitempty"`
	ImageTag            string            `json:"imageTag,omitempty"`
	ImageURL            string            `json:"imageURL,omitempty"`
	KeystoreSecret      string            `json:"keystoreSecret,omitempty"`
	GitHooks            GitHooksVolume    `json:"gitHooks,omitempty"`
	Jvm                 JvmObject         `json:"jvm,omitempty"`
	StorageClassName    string            `json:"storageClassName,omitempty"`
	RouteHostname       string            `json:"routeHostname,omitempty"`
	PvSize              string            `json:"pvSize,omitempty"`
	Simplified          bool              `json:"simplifed"`
	DashbuilderLocation string            `json:"dashbuilderLocation,omitempty"`
	Cors                CORSFiltersObject `json:"cors,omitempty"`
	StartupStrategy     StartupStrategy   `json:"startupStrategy,omitempty"`
	DataGridAuth        DataGridAuth      `json:"dataGridAuth,omitempty"`
}
