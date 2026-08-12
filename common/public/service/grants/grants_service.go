package grants

import grantsEntity "common/public/model/entity/grants"
import grantsRepo "common/internal/repository/grants"
import grantsConst "common/public/consts/grants"

// ListGrants 根据被授权访问的类型查询权限记录
func ListGrants(authorizedContentType grantsConst.AuthorizedContentType) ([]*grantsEntity.Grants, error) {
	poList, err := grantsRepo.SelectByContentType(authorizedContentType)
	if err != nil {
		return nil, err
	}

	var grantsList []*grantsEntity.Grants
	for _, po := range *poList {
		grantsList = append(grantsList, grantsEntity.LoadFromPO(&po))
	}
	return grantsList, nil
}