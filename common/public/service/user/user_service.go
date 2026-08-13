package user

import userEntity "common/public/model/entity/user"
import userDto "common/public/model/dto/user"

// GetUser 根据用户ID查询获取用户信息
func GetUser(userId uint64) (*userEntity.User, error) {
	return nil, nil
}

//Login 用户登录
func Login(login *userDto.LoginRequest) (*userDto.LoginSuccessResponse, error) {
	return nil, nil
}