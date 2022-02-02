package v2

// KafkaExtObject kafka configuration to be used by the KieApp
type KafkaExtObject struct {
	// Contains the mapping message/signal=topicName for every topic that needs to be mapped globally
	Topics []string `json:"topics,omitempty"`
	// A comma separated list of host/port pairs to use for establishing the initial connection to the Kafka cluster
	BootstrapServers string `json:"bootstrapServers,omitempty"`
	// Identifier to pass to the server when making requests
	ClientID string `json:"clientID,omitempty"`
	// Allow automatic topic creation.
	AutocreateTopics *bool `json:"autocreateTopics,omitempty"`
	// Group identifier the group this consumer belongs
	GroupID string `json:"groupID,omitempty"`
	// The number of acknowledgments the producer requires the leader to have received before considering a request complete.
	Acks *int `json:"acks,omitempty"`
	// Number of milliseconds that indicates how long publish method will bloc
	MaxBlockMs *int32 `json:"maxBlockMs,omitempty"`
}

// KafkaJBPMEventEmittersObject kafka configuration to be used by the KieApp for jBPM Emitter
type KafkaJBPMEventEmittersObject struct {
	// Comma separated list of host/port pairs to use for establishing the initial connection to the Kafka cluster.
	BootstrapServers string `json:"bootstrapServers,omitempty"`
	// This configuration allows users to set an ID to provide a logical application name for logging purposes, not set by default.
	ClientID string `json:"clientID,omitempty"`
	// The number of acknowledgments that the emitter requires the leader to have received before considering a request to be complete, not set by default.
	Acks *int `json:"acks,omitempty"`
	// Value in milliseconds that indicates how long the 'publish' method will block the operation. Default 2000 milliseconds (2 seconds).
	MaxBlockMs *int32 `json:"maxBlockMs,omitempty"`
	// Date and time format to be sent to Kafka. Default format is yyyy-MM-dd'T'HH:mm:ss.SSSZ
	DateFormat string `json:"dateFormat,omitempty"`
	// The topic name for processes event messages. Set up to override the default value jbpm-processes-events.
	ProcessesTopicName string `json:"processesTopicName,omitempty"`
	// The topic name for tasks event messages. Set up to override the default value jbpm-tasks-events.
	TasksTopicName string `json:"tasksTopicName,omitempty"`
	// The topic name for cases event messages. Set up to override the default value jbpm-cases-events.
	CasesTopicName string `json:"casesTopicName,omitempty"`
}
