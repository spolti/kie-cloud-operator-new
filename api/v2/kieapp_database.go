package v2

// DatabaseType to define what kind of database will be used for the Kie Servers
type DatabaseType string

const (
	// DatabaseH2 H2 Embedded Database deployment
	DatabaseH2 DatabaseType = "h2"
	// DatabaseMySQL MySQL Database deployment
	DatabaseMySQL DatabaseType = "mysql"
	// DatabasePostgreSQL PostgreSQL Database deployment
	DatabasePostgreSQL DatabaseType = "postgresql"
	// DatabaseExternal External Database
	DatabaseExternal DatabaseType = "external"
)

// DatabaseObject Defines how a KieServer will manage and create a new Database
// or connect to an existing one
type DatabaseObject struct {
	InternalDatabaseObject `json:",inline"`
	ExternalConfig         *ExternalDatabaseObject `json:"externalConfig,omitempty"`
}

// InternalDatabaseObject Defines how a deployment will manage and create a new Database
type InternalDatabaseObject struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum:=mysql;postgresql;external;h2
	// Database type to use
	Type DatabaseType `json:"type"`
	// Size of the PersistentVolumeClaim to create. For example, 100Gi
	Size string `json:"size,omitempty"`
	// The storageClassName to use for database pvc's.
	StorageClassName string `json:"storageClassName,omitempty"`
}

// CommonExtDBObjectRequiredURL common configuration definition of an external database
type CommonExtDBObjectRequiredURL struct {
	// +kubebuilder:validation:Required
	// Database JDBC URL. For example, jdbc:mysql:mydb.example.com:3306/rhpam
	JdbcURL                      string `json:"jdbcURL"`
	CommonExternalDatabaseObject `json:",inline"`
}

// CommonExtDBObjectURL common configuration definition of an external database
type CommonExtDBObjectURL struct {
	// Database JDBC URL. For example, jdbc:mysql:mydb.example.com:3306/rhpam
	JdbcURL                      string `json:"jdbcURL,omitempty"`
	CommonExternalDatabaseObject `json:",inline"`
}

// CommonExternalDatabaseObject common configuration definition of an external database
type CommonExternalDatabaseObject struct {
	// +kubebuilder:validation:Required
	// Driver name to use. For example, mysql
	Driver string `json:"driver"`
	// +kubebuilder:validation:Required
	// External database username
	Username string `json:"username"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Format:=password
	// External database password
	Password string `json:"password"`
	// Sets xa-pool/min-pool-size for the configured datasource.
	MinPoolSize string `json:"minPoolSize,omitempty"`
	// Sets xa-pool/max-pool-size for the configured datasource.
	MaxPoolSize string `json:"maxPoolSize,omitempty"`
	// An org.jboss.jca.adapters.jdbc.ValidConnectionChecker that provides a SQLException isValidConnection(Connection e) method to validate if a connection is valid.
	ConnectionChecker string `json:"connectionChecker,omitempty"`
	// An org.jboss.jca.adapters.jdbc.ExceptionSorter that provides a boolean isExceptionFatal(SQLException e) method to validate if an exception should be broadcast to all javax.resource.spi.ConnectionEventListener as a connectionErrorOccurred.
	ExceptionSorter string `json:"exceptionSorter,omitempty"`
	// Sets the sql validation method to background-validation, if set to false the validate-on-match method will be used.
	BackgroundValidation string `json:"backgroundValidation,omitempty"`
	// Defines the interval for the background-validation check for the jdbc connections.
	BackgroundValidationMillis string `json:"backgroundValidationMillis,omitempty"`
}

// ExternalDatabaseObject configuration definition of an external database
type ExternalDatabaseObject struct {
	// +kubebuilder:validation:Required
	// Hibernate dialect class to use. For example, org.hibernate.dialect.MySQL8Dialect
	Dialect string `json:"dialect"`
	// Database Name. For example, rhpam
	Name string `json:"name,omitempty"`
	// Database Host. For example, mydb.example.com. Host is intended to be used with databases running on OCP
	// where the host will correspond to the kubernetes added env *_SERVICE_HOST, it is mostly likely used
	// with PostgreSQL and MySQL variants running on OCP. For Databases Running outside OCP use jdbcUrl instead.
	Host string `json:"host,omitempty"`
	// Database Port. For example, 3306. Port is intended to be used with databases running on OCP
	// where the post will correspond to the kubernetes added env *_SERVICE_PORT, these are mostly likely used
	// with PostgreSQL and MySQL variants running on OCP. For Databases Running outside OCP use jdbcUrl instead.
	Port string `json:"port,omitempty"`
	// Sets the datasources type. It can be XA or NONXA. For non XA set it to true. Default value is false.
	NonXA                string `json:"nonXA,omitempty"`
	CommonExtDBObjectURL `json:",inline"`
}

// DatabaseTemplate contains all the variables used in the yaml templates
type DatabaseTemplate struct {
	InternalDatabaseObject `json:",inline"`
	ServerName             string `json:"serverName,omitempty"`
	Username               string `json:"username,omitempty"`
	DatabaseName           string `json:"databaseName,omitempty"`
}