package metadata

import (
	"encoding/json"
	"io"
	"os"

	"github.com/oracle/oci-go-sdk/example/helpers"
)

type NoopMetadataGetter struct {
	File string
}

func (m *NoopMetadataGetter) Get() (*InstanceMetadata, error) {
	md := &InstanceMetadata{}

	jsonFile, err := os.Open(m.File)
	helpers.FatalIfError(err)
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &md)
	return md, nil
}
