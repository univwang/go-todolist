package core

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"user/model"
	"user/service"
)

func BuildUser(item model.User) *service.UserModel {
	return &service.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.Username,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
}

func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest, resp *service.UserDetailResponse) error {
	user := model.User{}
	resp.Code = 200
	if err := model.DB.Where("username = ?", req.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return err
	}
	if user.CheckPassword(req.Password) == false {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest, resp *service.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("password not match")
		return err
	}
	count := 0
	if err := model.DB.Model(&model.User{}).Where("username = ?", req.UserName).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		err := errors.New("user already exist")
		return err
	}
	user := model.User{
		Username: req.UserName,
	}
	//加密密码
	err := user.SetPassword(req.Password)
	if err != nil {
		return err
	}
	//创建用户
	err = model.DB.Create(&user).Error
	if err != nil {
		return err
	}
	resp.UserDetail = BuildUser(user)
	return nil
}
