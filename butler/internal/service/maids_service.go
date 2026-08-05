package service

import (
	"butler/internal/model/dto"
	"butler/internal/model/entity"
	"butler/internal/repository"
	"common/public/pkg"
	"math"
)

// RegisterMaid 注册女仆节点
func RegisterMaid(req *dto.RegisterMaidReq) (uint64, error) {
	po := entity.RegisterReqToPO(req)
	err := repository.NewMaid(po)
	if err != nil {
		return 0, err
	}
	return po.MaidId, nil
}

// DismissMaid 解除注册（软删除）
func DismissMaid(maidId uint64) error {
	return repository.DeleteMaid(maidId)
}

// UpdateMaid 更新女仆节点信息
func UpdateMaid(maidId uint64, req *dto.UpdateMaidReq) error {
	po := entity.UpdateReqToPO(maidId, req)
	return repository.EditMaid(po)
}

// ListMaids 查看女仆节点列表
func ListMaids(req *dto.ListMaidsReq) (*pkg.PageableResp[dto.ListMaidsResp], error) {
	maids, total, err := repository.ListMaids(req.PageNum, req.PageSize, req.MaidName)
	if err != nil {
		return nil, pkg.InternalServerError(err.Error())
	}

	contents := make([]dto.ListMaidsResp, 0, len(maids))
	for i := range maids {
		resp := entity.ToListMaidsResp(&maids[i])
		if resp != nil {
			contents = append(contents, *resp)
		}
	}

	pages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &pkg.PageableResp[dto.ListMaidsResp]{
		Total:    total,
		Pages:    pages,
		Contents: contents,
	}, nil
}