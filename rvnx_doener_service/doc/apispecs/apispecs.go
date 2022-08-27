package apispecs

import _ "embed"

var (
	//go:embed v1-api.yaml
	API_V1_SpecsFile []byte
)