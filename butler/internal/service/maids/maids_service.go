package maids

import (
	maiddto "butler/internal/model/dto/maid"
	"butler/internal/model/entity/maid"
	"butler/internal/repository/maids"
	"common/public/pkg"
	"math"
)

// RegisterMaid 注册女仆节点
func RegisterMaid(req *maiddto.RegisterMaidReq) (uint64, error) {
	po := entity.RegisterReqToPO(req)
	err := maids.NewMaid(po)
	if err != nil {
		return 0, err
	}
	return po.MaidId, nil
}

// DismissMaid 解除注册（软删除）
func DismissMaid(maidId uint64) error {
	return maids.DeleteMaid(maidId)
}

// UpdateMaid 更新女仆节点信息
func UpdateMaid(maidId uint64, req *maiddto.UpdateMaidReq) error {
	po := entity.UpdateReqToPO(maidId, req)
	return maids.EditMaid(po)
}

// ListMaids 查看女仆节点列表
func ListMaids(req *maiddto.ListMaidsReq) (*pkg.PageableResp[maiddto.ListMaidsResp], error) {
	maidsData, total, err := maids.ListMaids(req.PageNum, req.PageSize, req.MaidName)
	if err != nil {
		return nil, pkg.InternalServerError(err.Error())
	}

	contents := make([]maiddto.ListMaidsResp, 0, len(maidsData))
	for i := range maidsData {
		resp := entity.ToListMaidsResp(&maidsData[i])
		if resp != nil {
			contents = append(contents, *resp)
		}
	}

	pages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &pkg.PageableResp[maiddto.ListMaidsResp]{
		Total:    total,
		Pages:    pages,
		Contents: contents,
	}, nil
}