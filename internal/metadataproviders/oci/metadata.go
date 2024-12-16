package oci // // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/oci"

import (
	"context"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/oci/metadata"
)

type Provider interface {
	AvailabilityDomain(ctx context.Context) (string, error)
	CompartmentID(ctx context.Context) (string, error)
	InstanceID(ctx context.Context) (string, error)
	DisplayName(ctx context.Context) (string, error)
	Region(ctx context.Context) (string, error)
	State(ctx context.Context) (string, error)
	DefinedTags(ctx context.Context) (map[string]map[string]interface{}, error)
	FreeformTags(ctx context.Context) (map[string]string, error)
	CanonicalRegionName(ctx context.Context) (string, error)
}

type ociProvider struct {
	ociInstanceMetadataGetter metadata.Interface
}

func NewProvider() (Provider, error) {
	baseURL := "http://169.254.169.254"
	metadataEndpoint := "/opc/v2/instance/"
	defaultHTTPTimeout := 5 * time.Second
	ociInstanceMetadataGetter := metadata.NewOciInstanceMetadataGetter(defaultHTTPTimeout, baseURL, metadataEndpoint)
	return &ociProvider{ociInstanceMetadataGetter}, nil
}

func (o *ociProvider) AvailabilityDomain(ctx context.Context) (string, error) {
	ociInstanceMetadata, err := o.ociInstanceMetadataGetter.Get()
	if err != nil {
		return "", err
	}
	return ociInstanceMetadata.AvailabilityDomain, err
}

func (o *ociProvider) CompartmentID(ctx context.Context) (string, error) {
	ociInstanceMetadata, err := o.ociInstanceMetadataGetter.Get()
	if err != nil {
		return "", err
	}
	return ociInstanceMetadata.CompartmentID, err
}

func (o *ociProvider) InstanceID(ctx context.Context) (string, error) {
	ociInstanceMetadata, err := o.ociInstanceMetadataGetter.Get()
	if err != nil {
		return "", err
	}
	return ociInstanceMetadata.InstanceID, err
}

func (o *ociProvider) DisplayName(ctx context.Context) (string, error) {
	ociInstanceMetadata, err := o.ociInstanceMetadataGetter.Get()
	if err != nil {
		return "", err
	}
	return ociInstanceMetadata.DisplayName, err
}

func (o *ociProvider) Region(ctx context.Context) (string, error) {
	ociInstanceMetadata, err := o.ociInstanceMetadataGetter.Get()
	if err != nil {
		return "", err
	}
	return ociInstanceMetadata.Region, err
}

func (o *ociProvider) State(ctx context.Context) (string, error) {
	ociInstanceMetadata, err := o.ociInstanceMetadataGetter.Get()
	if err != nil {
		return "", err
	}
	return ociInstanceMetadata.State, err
}

func (o *ociProvider) DefinedTags(ctx context.Context) (map[string]map[string]interface{}, error) {
	ociInstanceMetadata, err := o.ociInstanceMetadataGetter.Get()
	if err != nil {
		return nil, err
	}
	return ociInstanceMetadata.DefinedTags, err
}

func (o *ociProvider) FreeformTags(ctx context.Context) (map[string]string, error) {
	ociInstanceMetadata, err := o.ociInstanceMetadataGetter.Get()
	if err != nil {
		return nil, err
	}
	return ociInstanceMetadata.FreeformTags, err
}

func (o *ociProvider) CanonicalRegionName(ctx context.Context) (string, error) {
	ociInstanceMetadata, err := o.ociInstanceMetadataGetter.Get()
	if err != nil {
		return "", err
	}
	return ociInstanceMetadata.CanonicalRegionName, err
}
