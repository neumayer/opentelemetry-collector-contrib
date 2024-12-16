package oci

import (
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/stretchr/testify/assert"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/oci/metadata"
)

func TestNewOciInstanceMetadataGetter(t *testing.T) {
	baseURL := "http://169.254.169.254"
	metadataEndpoint := "/opc/v2/instance/"
	defaultHTTPTimeout := 5 * time.Second

	metadataGetter := metadata.NewOciInstanceMetadataGetter(defaultHTTPTimeout, baseURL, metadataEndpoint)

	assert.NotNil(t, metadataGetter)
	assert.IsType(t, &metadata.MetadataGetter{}, metadataGetter)

	metadataGetterInstance := metadataGetter.(*metadata.MetadataGetter)
	assert.Equal(t, defaultHTTPTimeout, metadataGetterInstance.Client.Timeout)
	assert.Equal(t, baseURL, metadataGetterInstance.BaseURL)
	assert.Equal(t, metadataEndpoint, metadataGetterInstance.MetadataEndpoint)
	instanceMetadata, err := metadataGetter.Get()
	assert.Nil(t, instanceMetadata)
	assert.NotNil(t, err)
}

func TestGet(t *testing.T) {
	expectedInstanceMetadata := &metadata.InstanceMetadata{
		InstanceID:          "ocid1.instance.oc1.phx.exampleuniqueID",
		CompartmentID:       "ocid1.tenancy.oc1..exampleuniqueID",
		DisplayName:         "my-example-instance",
		Region:              "phx",
		CanonicalRegionName: "us-phoenix-1",
		State:               "Running",
		DefinedTags:         map[string]map[string]interface{}{"Operations": {"CostCenter": "42"}},
		FreeformTags:        map[string]string{"Department": "Finance"},
	}

	localInstanceMetadata := metadata.NewLocalInstanceMetadataGetter("./oci-metadata.json")
	instanceMetaData, err := localInstanceMetadata.Get()
	helpers.FatalIfError(err)

	assert.EqualValuesf(t, expectedInstanceMetadata, instanceMetaData, "%v failed", "metadata structs")
}
