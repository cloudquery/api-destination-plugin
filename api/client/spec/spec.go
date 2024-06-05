package spec

import (
	"fmt"
)

type Spec struct {
	BaseURL string            `json:"base_url" jsonschema:"required,example=https://api.example.com"`
	Headers map[string]string `json:"headers,omitempty" jsonschema:"example={\"Authorization\":\"Bearer token\"}"`
}

func (s *Spec) Validate() error {
	if len(s.BaseURL) == 0 {
		return fmt.Errorf("`base_url` must be set")
	}
	return nil
}
