package model

import (
	"encoding/json"

	"gitlab.com/gregsolutions/dobby-cd/server/enum"
)

type PermissionRequest struct {
	Permits []*Permit `json:"permits" mapstructure:"permits"`
}

type PermissionUpdateResponse struct {
	Updated float64 `json:"updated"`
	New     float64 `json:"new"`
}
type PermissionResponse struct {
	Permits []Permit `json:"permits"`
}

type Permit struct {
	GroupDN    string          `mapstructure:"group_dn" json:"group_dn" binding:"required"`
	PermitType enum.PermitType `mapstructure:"permit_type" json:"permit_type" binding:"required"`
}

func (a *PermissionRequest) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}
