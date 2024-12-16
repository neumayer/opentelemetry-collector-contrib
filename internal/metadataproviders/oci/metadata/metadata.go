package metadata

import (
	"net/http"
	"time"
)

type InstanceMetadata struct {
	AvailabilityDomain  string                            `json:"availabilityDomain"`
	CompartmentID       string                            `json:"compartmentId"`
	InstanceID          string                            `json:"id"`
	DisplayName         string                            `json:"displayName"`
	Region              string                            `json:"region"`
	State               string                            `json:"state"`
	DefinedTags         map[string]map[string]interface{} `json:"definedTags"`
	FreeformTags        map[string]string                 `json:"freeformTags"`
	CanonicalRegionName string                            `json:"canonicalRegionName"`
}

type Interface interface {
	Get() (*InstanceMetadata, error)
}

func NewOciInstanceMetadataGetter(Timeout time.Duration, BaseURL string, MetadataEndpoint string) Interface {
	client := &http.Client{Timeout: Timeout}
	return &MetadataGetter{Client: client, BaseURL: BaseURL, MetadataEndpoint: MetadataEndpoint}
}

func NewLocalInstanceMetadataGetter(file string) Interface {
	return &NoopMetadataGetter{File: file}
}
