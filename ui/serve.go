package ui

//go:generate go run -mod=vendor ../controller/kieapp/defaults/.packr/packr.go

import (
	"context"
	"encoding/json"
	api "github.com/spolti/kie-cloud-operator-new/api/v2"
	"github.com/spolti/kie-cloud-operator-new/core/logger"
	"io/ioutil"

	"github.com/RHsyseng/console-cr-form/pkg/web"
	"github.com/ghodss/yaml"
	"github.com/go-openapi/spec"
	"github.com/gobuffalo/packr/v2"

	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

var log = logger.GetLogger("wizard-ui")

// Listen ...
func Listen() {
	config, err := web.NewConfiguration("", 8080, getSchema(), getApiVersion(), getObjectKind(), getForm(), apply)
	if err != nil {
		log.Error(err, "Failed to configure web server")
	}
	if err := web.RunWebServer(config); err != nil {
		log.Error(err, "Failed to run web server")
	}
}

func apply(cr string) error {
	log.Debug("Will deploy KieApp based on yaml %v", cr)
	kieApp := &api.KieApp{}
	err := yaml.Unmarshal([]byte(cr), kieApp)
	if err != nil {
		log.Error(err, "Failed to parse CR based on %s.", cr)
		return err
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Error(err, "Failed to get in-cluster config")
		return err
	}
	err = api.SchemeBuilder.AddToScheme(scheme.Scheme)
	if err != nil {
		log.Debug("Failed to add scheme", err)
		return err
	}
	config.ContentConfig.GroupVersion = &api.GroupVersion
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: scheme.Codecs}
	config.UserAgent = rest.DefaultKubernetesUserAgent()
	restClient, err := rest.UnversionedRESTClientFor(config)
	if err != nil {
		log.Debug("Failed to get REST client", err)
		return err
	}
	kieApp.SetGroupVersionKind(api.GroupVersion.WithKind("KieApp"))
	err = restClient.Post().Namespace(getCurrentNamespace()).Body(kieApp).Resource("kieapps").Do(context.TODO()).Into(kieApp)
	if err != nil {
		log.Debug("Failed to create KIE app", err)
		return err
	}
	log.Info("Created KIE application named %s", kieApp.Name)
	return nil
}

func getCurrentNamespace() string {
	bytes, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		log.Error(err, "Failed to read current namespace and cannot proceed ")
	}
	return string(bytes)
}

type CustomResourceDefinition struct {
	Spec CustomResourceDefinitionSpec `json:"spec,omitempty"`
}

type CustomResourceDefinitionSpec struct {
	Versions   []CustomResourceDefinitionVersion  `json:"versions,omitempty"`
	Validation CustomResourceDefinitionValidation `json:"validation,omitempty"`
}

type CustomResourceDefinitionVersion struct {
	Name   string                             `json:"Name,omitempty"`
	Schema CustomResourceDefinitionValidation `json:"schema,omitempty"`
}

type CustomResourceDefinitionValidation struct {
	OpenAPIV3Schema spec.Schema `json:"openAPIV3Schema,omitempty"`
}

func getSchema() spec.Schema {
	schema := spec.Schema{}
	box := packr.New("CRD", "../../deploy/crds/")
	yamlByte, err := box.Find("kieapp.crd.yaml")
	if err != nil {
		log.Error(err, "Failed to validate kieapp.crd.yaml")
		panic("Failed to retrieve crd, there must be an environment problem!")
	}
	crd := &CustomResourceDefinition{}
	err = yaml.Unmarshal(yamlByte, crd)
	if err != nil {
		log.Error(err, "Failed to unmarshal kieapp.crd.yaml")
		panic("Failed to unmarshal static schema, there must be an environment problem!")
	}
	for _, v := range crd.Spec.Versions {
		if v.Name == api.GroupVersion.Version {
			schema = v.Schema.OpenAPIV3Schema
		}
	}
	return schema
}

func getForm() web.Form {
	box := packr.New("form", "../../deploy/ui/")
	jsonBytes, err := box.Find("form.json")
	if err != nil {
		log.Error(err, "Failed to read form.json.")
		panic("Failed to retrieve ui form, there must be an environment problem!")
	}
	form := &web.Form{}
	err = json.Unmarshal(jsonBytes, form)
	if err != nil {
		log.Error(err, "Failed to unmarshall form.json")
		panic("Failed to unmarshal static ui form, there must be an environment problem!")
	}
	return *form
}

func getObjectKind() string {
	box := packr.New("CRD", "../../deploy/crds/")
	yamlByte, err := box.Find("kieapp.crd.yaml")
	if err != nil {
		log.Error(err, "Failed to read kieapp.crd.yaml")
		panic("Failed to retrieve crd, there must be an environment problem!")
	}
	crd := &extv1.CustomResourceDefinition{}
	err = yaml.Unmarshal(yamlByte, crd)
	if err != nil {
		panic("Failed to unmarshal static schema, there must be an environment problem!")
	}
	return crd.Spec.Names.Kind
}

func getApiVersion() string {
	return api.GroupVersion.String()
}
