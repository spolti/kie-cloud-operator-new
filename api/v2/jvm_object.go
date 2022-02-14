package v2

// JvmObject JVM specification to be used by the KieApp
// +k8s:openapi-gen=true
// +operator-sdk:csv:customresourcedefinitions:displayName="JVM specification"
type JvmObject struct {
	// User specified Java options to be appended to generated options in JAVA_OPTS. e.g. '-Dsome.property=foo'
	JavaOptsAppend string `json:"javaOptsAppend,omitempty"`
	// Is used when no '-Xmx' option is given in JAVA_OPTS. This is used to calculate a default maximal heap memory based on a containers restriction. If used in a container without any memory constraints for the container then this option has no effect. If there is a memory constraint then '-Xmx' is set to a ratio of the container available memory as set here. The default is '50' which means 50% of the available memory is used as an upper boundary. You can skip this mechanism by setting this value to '0' in which case no '-Xmx' option is added.
	JavaMaxMemRatio *int32 `json:"javaMaxMemRatio,omitempty"`
	// Is used when no '-Xms' option is given in JAVA_OPTS. This is used to calculate a default initial heap memory based on the maximum heap memory. If used in a container without any memory constraints for the container then this option has no effect. If there is a memory constraint then '-Xms' is set to a ratio of the '-Xmx' memory as set here. The default is '25' which means 25% of the '-Xmx' is used as the initial heap size. You can skip this mechanism by setting this value to '0' in which case no '-Xms' option is added. e.g. '25'
	JavaInitialMemRatio *int32 `json:"javaInitialMemRatio,omitempty"`
	// Is used when no '-Xms' option is given in JAVA_OPTS. This is used to calculate the maximum value of the initial heap memory. If used in a container without any memory constraints for the container then this option has no effect. If there is a memory constraint then '-Xms' is limited to the value set here. The default is 4096Mb which means the calculated value of '-Xms' never will be greater than 4096Mb. The value of this variable is expressed in MB. e.g. '4096'
	JavaMaxInitialMem *int32 `json:"javaMaxInitialMem,omitempty"`
	// Set this to get some diagnostics information to standard output when things are happening. Disabled by default. e.g. 'true'
	JavaDiagnostics *bool `json:"javaDiagnostics,omitempty"`
	// If set remote debugging will be switched on. Disabled by default. e.g. 'true'
	JavaDebug *bool `json:"javaDebug,omitempty"`
	// Port used for remote debugging. Defaults to 5005. e.g. '8787'
	JavaDebugPort *int32 `json:"javaDebugPort,omitempty"`
	// Minimum percentage of heap free after GC to avoid expansion. e.g. '20'
	GcMinHeapFreeRatio *int32 `json:"gcMinHeapFreeRatio,omitempty"`
	// Maximum percentage of heap free after GC to avoid shrinking. e.g. '40'
	GcMaxHeapFreeRatio *int32 `json:"gcMaxHeapFreeRatio,omitempty"`
	// Specifies the ratio of the time spent outside the garbage collection (for example, the time spent for application execution) to the time spent in the garbage collection, it's desirable that not more than 1 / (1 + n) e.g. 99 and means 1% spent on gc, 4 means spent 20% on gc.
	GcTimeRatio *int32 `json:"gcTimeRatio,omitempty"`
	// The weighting given to the current GC time versus previous GC times  when determining the new heap size. e.g. '90'
	GcAdaptiveSizePolicyWeight *int32 `json:"gcAdaptiveSizePolicyWeight,omitempty"`
	// The maximum metaspace size in Mega bytes unit e.g. 400
	GcMaxMetaspaceSize *int32 `json:"gcMaxMetaspaceSize,omitempty"`
	// Specify Java GC to use. The value of this variable should contain the necessary JRE command-line options to specify the required GC, which will override the default of '-XX:+UseParallelOldGC'. e.g. '-XX:+UseG1GC'
	GcContainerOptions string `json:"gcContainerOptions,omitempty"`
}
