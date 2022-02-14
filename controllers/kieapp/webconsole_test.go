package kieapp

import (
	"github.com/spolti/kie-cloud-operator-new/controllers/kieapp/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/mod/semver"
)

func TestYamlSampleCreation(t *testing.T) {
	reconciler := &KieAppReconciler{Service: test.MockService(), OcpVersion: semver.MajorMinor("v4.2")}
	err := reconciler.createConsoleYAMLSamples()
	assert.NotNil(t, err)
	assert.Equal(t, "console yaml samples not installed, incompatible ocp version", err.Error())
	reconciler.OcpVersion = semver.MajorMinor("v4.3")
	assert.Nil(t, reconciler.createConsoleYAMLSamples())
}

func TestYamlSampleCreationUnknownClusterVersion(t *testing.T) {
	reconciler := &KieAppReconciler{Service: test.MockService(), OcpVersion: semver.MajorMinor("")}
	assert.Nil(t, reconciler.createConsoleYAMLSamples())
}
