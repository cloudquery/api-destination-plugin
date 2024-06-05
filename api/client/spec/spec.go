package spec

import (
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

type Spec struct {
	BaseURL      string               `json:"base_url" jsonschema:"required,example=https://api.example.com"`
	Headers      map[string]string    `json:"headers,omitempty" jsonschema:"example={\"Authorization\":\"Bearer token\"}"`
	BatchSize    *int64               `json:"batch_size" jsonschema:"minimum=1,default=1000"`
	BatchTimeout *configtype.Duration `jsonschema:"default=30s"`
}

func (s *Spec) SetDefaults() {
	if s.BatchSize == nil {
		s.BatchSize = ptr(int64(1000))
	}
	if s.BatchTimeout == nil {
		d := configtype.NewDuration(30 * time.Second)
		s.BatchTimeout = &d
	}
}

func (s *Spec) Validate() error {
	if len(s.BaseURL) == 0 {
		return fmt.Errorf("`base_url` must be set")
	}
	return nil
}

func ptr[A any](a A) *A {
	return &a
}
