// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oci // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/oci"

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/oci"
	localMetadata "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/oci/internal/metadata"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
)

const (
	TypeStr = "oci"
)

// NewDetector returns a detector which can detect resource attributes on oci
func NewDetector(set processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	cfg := dcfg.(Config)
	ociProvider, err := oci.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed creating detector: %w", err)
	}
	return &Detector{
		logger:   set.Logger,
		provider: ociProvider,
		rb:       localMetadata.NewResourceBuilder(cfg.ResourceAttributes),
	}, nil
}

var _ internal.Detector = (*Detector)(nil)

type Detector struct {
	provider oci.Provider
	logger   *zap.Logger
	rb       *localMetadata.ResourceBuilder
}

func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	d.rb.SetCloudProvider("oci")
	availabilityDomain, err := d.provider.AvailabilityDomain(ctx)
	if err != nil {
		return pcommon.NewResource(), "", fmt.Errorf("failed getting availability domain: %w", err)
	}
	d.rb.SetOciAvailabilityDomain(availabilityDomain)
	compartmentID, err := d.provider.CompartmentID(ctx)
	if err != nil {
		return pcommon.NewResource(), "", fmt.Errorf("failed getting compartment id type: %w", err)
	}
	d.rb.SetOciCompartmentID(compartmentID)
	instanceID, err := d.provider.InstanceID(ctx)
	if err != nil {
		return pcommon.NewResource(), "", fmt.Errorf("failed getting instance id type: %w", err)
	}
	d.rb.SetOciInstanceID(instanceID)
	displayName, err := d.provider.DisplayName(ctx)
	if err != nil {
		return pcommon.NewResource(), "", fmt.Errorf("failed getting display name type: %w", err)
	}
	d.rb.SetOciInstanceDisplayname(displayName)
	region, err := d.provider.Region(ctx)
	if err != nil {
		return pcommon.NewResource(), "", fmt.Errorf("failed getting region name type: %w", err)
	}
	d.rb.SetOciInstanceRegion(region)
	state, err := d.provider.State(ctx)
	if err != nil {
		return pcommon.NewResource(), "", fmt.Errorf("failed getting instance state type: %w", err)
	}
	d.rb.SetOciInstanceState(state)
	definedTags, err := d.provider.DefinedTags(ctx)
	if err != nil {
		return pcommon.NewResource(), "", fmt.Errorf("failed getting defined tags type: %w", err)
	}
	d.rb.SetOciOwner(definedTags["Oracle-Recommended-Tags"]["ResourceOwner"].(string))
	return d.rb.Emit(), conventions.SchemaURL, nil
}
