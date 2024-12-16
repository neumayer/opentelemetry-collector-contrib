package metadata

import (
	"testing"

	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	expectedInstanceMetadata := &InstanceMetadata{
		AvailabilityDomain:  "EMIr:PHX-AD-1",
		InstanceID:          "ocid1.instance.oc1.phx.exampleuniqueID",
		CompartmentID:       "ocid1.tenancy.oc1..exampleuniqueID",
		DisplayName:         "my-example-instance",
		Region:              "phx",
		CanonicalRegionName: "us-phoenix-1",
		State:               "Running",
		DefinedTags:         map[string]map[string]interface{}{"Operations": {"CostCenter": "42"}},
		FreeformTags:        map[string]string{"Department": "Finance"},
	}

	localInstanceMetadata := NewLocalInstanceMetadataGetter("./oci-metadata.json")
	instanceMetaData, err := localInstanceMetadata.Get()
	helpers.FatalIfError(err)

	assert.EqualValuesf(t, expectedInstanceMetadata, instanceMetaData, "%v failed", "metadata structs")
}
