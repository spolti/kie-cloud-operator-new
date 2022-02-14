module github.com/spolti/kie-cloud-operator-new

go 1.16

require (
	github.com/RHsyseng/console-cr-form v0.0.0-00010101000000-000000000000
	github.com/RHsyseng/operator-utils v0.0.0-00010101000000-000000000000
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.4.0
	github.com/go-openapi/spec v0.19.9
	github.com/gobuffalo/packr/v2 v2.7.1
	github.com/google/go-cmp v0.5.5
	github.com/google/uuid v1.1.2
	github.com/imdario/mergo v0.3.12
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.15.0
	github.com/openshift/api v0.0.0-20210521075222-e273a339932a
	github.com/openshift/client-go v0.0.0-20210521082421-73d9475a9142
	github.com/operator-framework/api v0.5.3
	//github.com/operator-framework/operator-sdk v1.13.1
	github.com/pavel-v-chernykh/keystore-go/v4 v4.1.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.50.0
	github.com/prometheus/common v0.26.0
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/gjson v1.6.1
	github.com/tidwall/sjson v1.1.1
	golang.org/x/mod v0.4.2
	k8s.io/api v0.22.1
	k8s.io/apiextensions-apiserver v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
	sigs.k8s.io/controller-runtime v0.10.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v14.2.0+incompatible // Required by OLM

	// Pin RHsyseng library versions
	github.com/RHsyseng/console-cr-form => github.com/RHsyseng/console-cr-form v0.0.0-20210323180350-be2aa15abde0
	github.com/RHsyseng/operator-utils => github.com/RHsyseng/operator-utils v1.4.6-0.20220111145438-f824bea2df43

	github.com/gobuffalo/packr/v2 => github.com/gobuffalo/packr/v2 v2.7.1

	// Versions after v0.3.7 change behaviour
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.7

	// OpenShift release-4.8
	github.com/openshift/api => github.com/openshift/api v0.0.0-20210521075222-e273a339932a
	github.com/openshift/client-go => github.com/openshift/client-go v0.0.0-20210521082421-73d9475a9142

	// Operator Framework v0.19.2
	//github.com/operator-framework/api => github.com/operator-framework/api v0.3.12
	// github.com/operator-framework/operator-sdk => github.com/operator-framework/operator-sdk v0.19.2

	// others
	modernc.org/cc => gitlab.com/cznic/cc v1.0.0
	modernc.org/golex => gitlab.com/cznic/golex v1.0.0
	modernc.org/mathutil => gitlab.com/cznic/mathutil v1.0.0
	modernc.org/strutil => gitlab.com/cznic/strutil v1.0.0
	modernc.org/xc => gitlab.com/cznic/xc v1.0.0
	mvdan.cc/interfacer => github.com/mvdan/interfacer v0.0.0-20180901003855-c20040233aed
	mvdan.cc/lint => github.com/mvdan/lint v0.0.0-20170908181259-adc824a0674b
	mvdan.cc/unparam => github.com/mvdan/unparam v0.0.0-20190209190245-fbb59629db34
)

//go 1.16
//
//require (
//	github.com/RHsyseng/console-cr-form v0.0.0-00010101000000-000000000000
//	github.com/RHsyseng/operator-utils v0.0.0-00010101000000-000000000000
//	github.com/ghodss/yaml v1.0.0
//	github.com/go-logr/logr v0.4.0
//	github.com/go-openapi/spec v0.19.9
//	github.com/gobuffalo/packr/v2 v2.7.1
//	github.com/google/go-cmp v0.5.6
//	github.com/google/uuid v1.1.2
//	github.com/imdario/mergo v0.3.12
//	github.com/onsi/ginkgo v1.16.4
//	github.com/onsi/gomega v1.15.0
//	github.com/openshift/api v0.0.0-20210521075222-e273a339932a
//	github.com/openshift/client-go v0.0.0-20210521082421-73d9475a9142
//	github.com/operator-framework/api v0.3.12
//	github.com/pavel-v-chernykh/keystore-go/v4 v4.3.0
//	github.com/pkg/errors v0.9.1
//	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.50.0
//	github.com/prometheus/common v0.26.0
//	github.com/stretchr/testify v1.7.0
//	github.com/tidwall/gjson v1.14.0
//	github.com/tidwall/sjson v1.2.4
//	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
//	golang.org/x/mod v0.4.2
//	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
//	k8s.io/api v0.21.2
//	k8s.io/apiextensions-apiserver v0.21.2
//	k8s.io/apimachinery v0.21.2
//	k8s.io/client-go v0.21.2
//	sigs.k8s.io/controller-runtime v0.9.1 // update on github.com/RHsyseng/operator-utils and bump ocp to 4.9
//)
//
//replace (
//	// Pin RHsyseng library versions
//	github.com/RHsyseng/console-cr-form => github.com/RHsyseng/console-cr-form v0.0.0-20210323180350-be2aa15abde0
//	github.com/RHsyseng/operator-utils => github.com/RHsyseng/operator-utils v1.4.6-0.20220111145438-f824bea2df43
//
//	github.com/gobuffalo/packr/v2 => github.com/gobuffalo/packr/v2 v2.7.1
//
//	// Versions after v0.3.7 change behaviour
//	github.com/imdario/mergo => github.com/imdario/mergo v0.3.7
//
//	// OpenShift release-4.8
//	github.com/openshift/api => github.com/openshift/api v0.0.0-20210521075222-e273a339932a
//	github.com/openshift/client-go => github.com/openshift/client-go v0.0.0-20210521082421-73d9475a9142
//
//)
