package model

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type ProjectRequest struct {
	Name string `json:"name"`
}

type ProjectResponse struct {
	ID   string `mapstructure:"id"`
	Name string `mapstructure:"name"`
}

func (a *ProjectRequest) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}

func (a *ProjectResponse) UnmarshalBinary(data interface{}) error {
	return mapstructure.Decode(data, a)
}
