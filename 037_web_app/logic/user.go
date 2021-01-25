package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

// SignUp 业务注册
func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 1. 生成UID
	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 2. 密码加密
	// 3. 保存进数据库
	return mysql.InsertUser(user)
}

// Login 登录逻辑
func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	user.Token = token
	return
}
