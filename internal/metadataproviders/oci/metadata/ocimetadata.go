package metadata

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type MetadataGetter struct {
	BaseURL          string
	MetadataEndpoint string
	Client           *http.Client
}

func (m *MetadataGetter) Get() (*InstanceMetadata, error) {
	req, err := http.NewRequest("GET", m.BaseURL+m.MetadataEndpoint, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return m.executeRequest(req)
}

func (m *MetadataGetter) executeRequest(req *http.Request) (*InstanceMetadata, error) {
	req.Header.Add("Authorization", "Bearer Oracle")
	resp, err := m.Client.Do(req)
	if err != nil {
		zap.S().With(zap.Error(err)).Warn("Failed to get instance metadata with v2 endpoint.")
		return nil, errors.Wrap(err, "Failed to get instance metadata with v2 endpoint")
	}

	zap.S().Infof("Metadata endpoint %s returned response successfully", req.URL.Path)
	if resp != nil {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("metadata endpoint v2 returned status %d; expected 200 OK", resp.StatusCode)
		}
	}
	md := &InstanceMetadata{}
	err = json.NewDecoder(resp.Body).Decode(md)
	if err != nil {
		return nil, errors.Wrap(err, "decoding instance metadata response")
	}

	return md, nil
}
