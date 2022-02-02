package v2

// KieAppJmsObject messaging specification to be used by the KieApp
type KieAppJmsObject struct {
	// +kubebuilder:validation:Required
	// When set to true will configure the KIE Server with JMS integration, if no configuration is added, the default will be used.
	EnableIntegration bool `json:"enableIntegration"`
	// Set false to disable the JMS executor, it is enabled by default.
	Executor *bool `json:"executor,omitempty"`
	// Enable transactions for JMS executor, disabled by default.
	ExecutorTransacted bool `json:"executorTransacted,omitempty"`
	// JNDI name of request queue for JMS, example queue/CUSTOM.KIE.SERVER.REQUEST, default is queue/KIE.SERVER.REQUEST.
	QueueRequest string `json:"queueRequest,omitempty"`
	// JNDI name of response queue for JMS, example queue/CUSTOM.KIE.SERVER.RESPONSE, default is queue/KIE.SERVER.RESPONSE.
	QueueResponse string `json:"queueResponse,omitempty"`
	// JNDI name of executor queue for JMS, example queue/CUSTOM.KIE.SERVER.EXECUTOR, default is queue/KIE.SERVER.EXECUTOR.
	QueueExecutor string `json:"queueExecutor,omitempty"`
	// Enable the Signal configuration through JMS. Default is false.
	EnableSignal bool `json:"enableSignal,omitempty"`
	// JNDI name of signal queue for JMS, example queue/CUSTOM.KIE.SERVER.SIGNAL, default is queue/KIE.SERVER.SIGNAL.
	QueueSignal string `json:"queueSignal,omitempty"`
	// Enable the Audit logging through JMS. Default is false.
	EnableAudit bool `json:"enableAudit,omitempty"`
	// JNDI name of audit logging queue for JMS, example queue/CUSTOM.KIE.SERVER.AUDIT, default is queue/KIE.SERVER.AUDIT.
	QueueAudit string `json:"queueAudit,omitempty"`
	// Determines if JMS session is transacted or not - default true.
	AuditTransacted *bool `json:"auditTransacted,omitempty"`
	// AMQ broker username to connect do the AMQ, generated if empty.
	Username string `json:"username,omitempty"`
	// +kubebuilder:validation:Format:=password
	// AMQ broker password to connect do the AMQ, generated if empty.
	Password string `json:"password,omitempty"`
	// AMQ broker broker comma separated queues, if empty the values from default queues will be used.
	AMQQueues string `json:"amqQueues,omitempty"` // It will receive the default value for the Executor, Request, Response, Signal and Audit queues.
	// The name of a secret containing AMQ SSL related files.
	AMQSecretName string `json:"amqSecretName,omitempty"` // AMQ SSL parameters
	// The name of the AMQ SSL Trust Store file.
	AMQTruststoreName string `json:"amqTruststoreName,omitempty"`
	// +kubebuilder:validation:Format:=password
	// The password for the AMQ Trust Store.
	AMQTruststorePassword string `json:"amqTruststorePassword,omitempty"`
	// The name of the AMQ keystore file.
	AMQKeystoreName string `json:"amqKeystoreName,omitempty"`
	// +kubebuilder:validation:Format:=password
	// The password for the AMQ keystore and certificate.
	AMQKeystorePassword string `json:"amqKeystorePassword,omitempty"`
	// Not intended to be set by the user, if will be set to true if all required SSL parameters are set.
	AMQEnableSSL bool `json:"amqEnableSSL,omitempty"` // flag will be set to true if all AMQ SSL parameters are correctly set.
}
