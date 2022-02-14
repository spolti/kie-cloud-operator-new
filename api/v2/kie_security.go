package v2

// KieAppAuthObject Authentication specification to be used by the KieApp
type KieAppAuthObject struct {
	SSO  *SSOAuthConfig  `json:"sso,omitempty"`
	LDAP *LDAPAuthConfig `json:"ldap,omitempty"`
	// When present, the RoleMapping Login Module will be configured.
	RoleMapper *RoleMapperAuthConfig `json:"roleMapper,omitempty"`
}

// SSOAuthConfig Authentication configuration for SSO
type SSOAuthConfig struct {
	// +kubebuilder:validation:Format:=password
	// RH-SSO Realm Admin Password used to create the Client
	AdminPassword string `json:"adminPassword,omitempty"`
	// RH-SSO Realm Admin Username used to create the Client if it doesn't exist
	AdminUser string `json:"adminUser,omitempty"`
	// +kubebuilder:validation:Required
	// RH-SSO URL
	URL string `json:"url"`
	// +kubebuilder:validation:Required
	// RH-SSO Realm name
	Realm string `json:"realm"`
	// RH-SSO Disable SSL Certificate Validation
	DisableSSLCertValidation bool `json:"disableSSLCertValidation,omitempty"`
	// RH-SSO Principal Attribute to use as username
	PrincipalAttribute string `json:"principalAttribute,omitempty"`
}

// SSOAuthClient Auth client to use for the SSO integration
type SSOAuthClient struct {
	// +kubebuilder:validation:Format:=password
	// Client secret
	Secret string `json:"secret,omitempty"`
	// Client name
	Name string `json:"name,omitempty"`
	// Hostname to set as redirect URL
	HostnameHTTP string `json:"hostnameHTTP,omitempty"`
	// Secure hostname to set as redirect URL
	HostnameHTTPS string `json:"hostnameHTTPS,omitempty"`
}

// LDAPAuthConfig Authentication configuration for LDAP
type LDAPAuthConfig struct {
	// +kubebuilder:validation:Format:=password
	// LDAP Credentials used for authentication
	BindCredential string `json:"bindCredential,omitempty"`
	// +kubebuilder:validation:Required
	// LDAP endpoint to connect for authentication. For failover set two or more LDAP endpoints separated by space
	URL string `json:"url"`
	// +kubebuilder:validation:Enum:=optional
	// LDAP login module flag, adds backward compatibility with the legacy security subsystem on elytron.
	// 'optional' is the only supported value, if set will create a distributed realm with ldap and filesystem realm
	// with the user added using the KIE_ADMIN_USER.
	LoginModule LoginModuleType `json:"loginModule,omitempty"`
	// Enable failover, if Ldap Url is unreachable, it will fail over to the KieFsRealm.
	LoginFailover bool `json:"loginFailover,omitempty"`
	// Bind DN used for authentication
	BindDN string `json:"bindDN,omitempty"`
	// LDAP Base DN of the top-level context to begin the user search.
	BaseCtxDN string `json:"baseCtxDN,omitempty"`
	// DAP search filter used to locate the context of the user to authenticate. The input username or userDN obtained from the login module callback is substituted into the filter anywhere a {0} expression is used. A common example for the search filter is (uid={0}).
	BaseFilter string `json:"baseFilter,omitempty"`
	// Indicates if the user queries are recursive.
	RecursiveSearch bool `json:"recursiveSearch,omitempty"`
	// The timeout in milliseconds for user or role searches.
	SearchTimeLimit int32 `json:"searchTimeLimit,omitempty"`
	// Name of the attribute containing the user roles.
	RoleAttributeID string `json:"roleAttributeID,omitempty"`
	// The fixed DN of the context to search for user roles. This is not the DN where the actual roles are, but the DN where the objects containing the user roles are. For example, in a Microsoft Active Directory server, this is the DN where the user account is.
	RolesCtxDN string `json:"rolesCtxDN,omitempty"`
	// A search filter used to locate the roles associated with the authenticated user. The input username or userDN obtained from the login module callback is substituted into the filter anywhere a {0} expression is used. The authenticated userDN is substituted into the filter anywhere a {1} is used. An example search filter that matches on the input username is (member={0}). An alternative that matches on the authenticated userDN is (member={1}).
	RoleFilter string `json:"roleFilter,omitempty"`
	// +kubebuilder:validation:Format:=int16
	// The number of levels of recursion the role search will go below a matching context. Disable recursion by setting this to 0.
	RoleRecursion int16 `json:"roleRecursion,omitempty"`
	// A role included for all authenticated users
	DefaultRole string `json:"defaultRole,omitempty"`
	// Provide new identities for Ldap  identity mapping, the pattern to be used with this env is 'attribute_name=attribute_value;another_attribute_name=value'
	NewIdentityAttributes string `json:"newIdentityAttributes,omitempty"`
	// +kubebuilder:validation:Enum:=FOLLOW;IGNORE;THROW
	// If LDAP referrals should be followed.
	ReferralMode ReferralModeType `json:"referralMode,omitempty"`
}

// AuthTemplate Authentication definition used in the template
type AuthTemplate struct {
	SSO        SSOAuthConfig      `json:"sso,omitempty"`
	LDAP       LDAPAuthConfig     `json:"ldap,omitempty"`
	RoleMapper RoleMapperTemplate `json:"roleMapper,omitempty"`
}

// RoleMapperTemplate RoleMapper definition used in the template
type RoleMapperTemplate struct {
	MountPath            string `json:"mountPath,omitempty"`
	RoleMapperAuthConfig `json:",inline"`
}

// RoleMapperAuthConfig Configuration for RoleMapper Authentication
type RoleMapperAuthConfig struct {
	// +kubebuilder:validation:Required
	// When present, the RoleMapping will be configured to use the provided properties file or roles. This parameter
	// defines the fully-qualified file path and name of a properties file or a set of roles with the following pattern
	// 'role=role1;another-role=role2'. The format of every entry in the file is original_role=role1,role2,role3
	// expects eiter a .properties file or a content with the patter above.
	RolesProperties string `json:"rolesProperties"`
	// When set to 'true' the mapped roles will retain all roles, that have defined mappings. Defaults to false.
	RolesKeepMapped bool `json:"rolesKeepMapped,omitempty"`
	// When set to 'true' the mapped roles will retain all roles, that have no defined mappings. Defaults to false.
	RolesKeepNonMapped bool    `json:"rolesKeepNonMapped,omitempty"`
	From               *ObjRef `json:"from,omitempty"`
}

// ReferralModeType Type used to define how the LDAP will follow referrals
type ReferralModeType string

const (
	Follow ReferralModeType = "FOLLOW"
	Ignore ReferralModeType = "IGNORE"
	Throw  ReferralModeType = "THROW"
)

// LoginModuleType A flag to set login module to optional.
type LoginModuleType string

const (
	//OptionalLoginModule optional login module
	OptionalLoginModule LoginModuleType = "optional"
)

// ObjRef contains enough information to let you inspect or modify the referred object.
type ObjRef struct {
	// Kind of the referent.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum:=ConfigMap;Secret;PersistentVolumeClaim
	Kind            string `json:"kind" protobuf:"bytes,1,opt,name=kind"`
	ObjectReference `json:",inline"`
}

// KieAppTruststore defines which truststore is used by the console, kieservers, smartrouter, and dashbuilder
type KieAppTruststore struct {
	// Set true to use Openshift's CA Bundle as a truststore, instead of java's cacert.
	OpenshiftCaBundle bool `json:"openshiftCaBundle,omitempty"`
}
