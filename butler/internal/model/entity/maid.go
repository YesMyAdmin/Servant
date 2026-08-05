package entity

import (
	"butler/internal/model/dto"
	"butler/internal/model/po"
)

// RegisterReqToPO 将 RegisterMaidReq 转换为 MaidPO
func RegisterReqToPO(r *dto.RegisterMaidReq) *po.MaidPO {
	return &po.MaidPO{
		MaidName:    r.MaidName,
		HostPort:    r.HostPort,
		AccessSecret: r.AccessSecret,
	}
}

// UpdateReqToPO 将 UpdateMaidReq 转换为 MaidPO
func UpdateReqToPO(maidId uint64, r *dto.UpdateMaidReq) *po.MaidPO {
	return &po.MaidPO{
		MaidId:      maidId,
		MaidName:    r.MaidName,
		HostPort:    r.HostPort,
		AccessSecret: r.AccessSecret,
	}
}

// ToListMaidsResp 将 MaidPO 转换为 ListMaidsResp
func ToListMaidsResp(p *po.MaidPO) *dto.ListMaidsResp {
	if p == nil {
		return nil
	}
	return &dto.ListMaidsResp{
		MaidId:      dto.Uint64ToString(p.MaidId),
		MaidName:    p.MaidName,
		HostPort:    p.HostPort,
		AccessSecret: p.AccessSecret,
		Enabled:     p.Enabled,
		DeletedTime: p.DeletedTime,
		CreateTime:  p.CreateTime,
		OwnerId:     dto.Uint64ToString(uint64(p.OwnerId)),
		UpdateTime:  p.UpdateTime,
	}
}