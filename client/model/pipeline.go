package model

import "encoding/json"

type PipelineRequest struct {
	Path       string `json:"path"`
	Token      string `json:"token"`
	Branch     string `json:"branch"`
	Repository string `json:"repository"`
}
type PipelineResponse struct {
	PipelineId string `mapstructure:"pipeline_id"`
}

func (a *PipelineRequest) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}
