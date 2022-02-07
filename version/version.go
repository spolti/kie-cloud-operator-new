package version

import (
	"github.com/spolti/kie-cloud-operator-new/controllers/kieapp/constants"
)

var (
	// Version - current version
	Version = constants.CurrentVersion
	// CsvVersion - csv release
	CsvVersion = Version + "-1"
	// PriorVersion - prior version
	PriorVersion = constants.PriorVersion
	// CsvPriorVersion - prior csv release
	CsvPriorVersion = PriorVersion + "-1"
)
