package user

import userEntity "common/public/model/entity/user"
import userDto "common/public/model/dto/user"
import userRepo "common/internal/repository/user"
import grants "common/public/model/entity/grants"
import grantsRepo "common/internal/repository/grants"


// GetUser 根据用户ID查询获取用户信息,包含用户所在组和拥有的权限
func GetUser(userId uint64) (*userEntity.User, error) {
	po, err := userRepo.SelectById(userId)
	if (err != nil) {
		return nil, err
	}
	//加载权限信息
	grantsPO, err := grantsRepo.SelectByTargetId(string(grants.User), userId)
	return userEntity.LoadFromPO(po, nil, grantsPO), nil
}

//Login 用户登录
func Login(login *userDto.LoginRequest) (*userDto.LoginSuccessResponse, error) {
	return nil, nil
}