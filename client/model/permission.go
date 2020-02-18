package model

import "gitlab.com/gregsolutions/dobby-cd/server/enum"

type PermissionResponse struct {
	Permits []Permit `json:"permits"`
}

type Permit struct {
	GroupDN    string          `json:"group_dn"  binding:"required"`
	PermitType enum.PermitType `json:"permit_type" binding:"required"`
}
