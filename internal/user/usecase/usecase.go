package usecase

import (
	"context"

	"InitiaRe-website/constant"
	authModel "InitiaRe-website/internal/auth/models"
	authUc "InitiaRe-website/internal/auth/usecase"
	userInfoModel "InitiaRe-website/internal/user_info/models"
	userInfoUc "InitiaRe-website/internal/user_info/usecase"
	"InitiaRe-website/pkg/utils"
)

type usecase struct {
	authUc     authUc.IUseCase
	userInfoUc userInfoUc.IUseCase
}

func InitUsecase(authUc authUc.IUseCase, userInfoUc userInfoUc.IUseCase) IUseCase {
	return &usecase{
		authUc,
		userInfoUc,
	}
}

func (u *usecase) Enable(ctx context.Context, userId int) error {
	user, err := u.userInfoUc.GetOne(ctx, &userInfoModel.RequestList{
		UserId: userId,
	})
	if err != nil {
		return err
	}
	if user.UserId == 0 {
		return utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_USER_NOT_FOUND)
	}
	if _, err := u.userInfoUc.Update(ctx, userId, &userInfoModel.SaveRequest{
		Id:     user.Id,
		Status: constant.USER_STATUS_ACTIVE,
	}); err != nil {
		return err
	}
	return nil
}

func (u *usecase) Disable(ctx context.Context, userId int) error {
	user, err := u.userInfoUc.GetOne(ctx, &userInfoModel.RequestList{
		UserId: userId,
	})
	if err != nil {
		return err
	}
	if user.UserId == 0 {
		return utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_USER_NOT_FOUND)
	}
	if _, err := u.userInfoUc.Update(ctx, userId, &userInfoModel.SaveRequest{
		Id:     user.Id,
		Status: constant.USER_STATUS_INACTIVE,
	}); err != nil {
		return err
	}
	return nil
}

func (u *usecase) PromoteAdmin(ctx context.Context, userId int, email string) error {
	// Get user by email
	user, err := u.authUc.GetOne(ctx, &authModel.RequestList{
		Email: email,
	})
	if err != nil {
		return err
	}
	if user.Id == 0 {
		return utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_USER_NOT_FOUND)
	}

	// Get user info
	userInfo, err := u.userInfoUc.GetOne(ctx, &userInfoModel.RequestList{
		UserId: user.Id,
	})
	if err != nil {
		return err
	}
	if userInfo.UserId == 0 {
		return utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_USER_NOT_FOUND)
	}
	if userInfo.Role == constant.USER_ROLE_ADMIN {
		return nil
	}
	if userInfo.Role != constant.USER_ROLE_NORMAL {
		return utils.NewError(constant.STATUS_CODE_BAD_REQUEST, "Only allowed promote from normal user to admin")
	}
	if _, err := u.userInfoUc.Update(ctx, userId, &userInfoModel.SaveRequest{
		Id:   user.Id,
		Role: constant.USER_ROLE_ADMIN,
	}); err != nil {
		return err
	}
	return nil
}

func (u *usecase) IsAdmin(ctx context.Context, userId int) (bool, error) {
	userInfo, err := u.userInfoUc.GetOne(ctx, &userInfoModel.RequestList{
		UserId: userId,
	})
	if err != nil {
		return false, err
	}
	if userInfo.UserId == 0 {
		return false, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_USER_NOT_FOUND)
	}
	if userInfo.Role != constant.USER_ROLE_ADMIN {
		return false, nil
	}
	return true, nil
}
