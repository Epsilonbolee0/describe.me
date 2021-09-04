package service

import (
	"../../keygen"
	"../../utils"

	repo "../repository"
	"gorm.io/gorm"
)

type ProfileService struct {
	userRepo *repo.UserRepository
}

func NewProfileService(userRepo *repo.UserRepository) *ProfileService {
	return &ProfileService{userRepo: userRepo}
}

func (profile *ProfileService) Info(login string) map[string]interface{} {
	user, err := profile.userRepo.Find(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message("User not found")
	default:
		return utils.Message("Error occured while finding key")
	}

	response := utils.Message("User has been found!")
	response["user"] = user
	return response
}

func (profile *ProfileService) CurrentKey(login string) map[string]interface{} {
	user, err := profile.userRepo.Find(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message("User not found")
	default:
		return utils.Message("Error occured while finding key")
	}

	if !user.IsAdmin {
		return utils.Message("User is not admin!")
	}

	response := utils.Message("Key has been found!")
	response["key"] = keygen.GetKey()
	return response
}

func (profile *ProfileService) UpdateName(login, name string) map[string]interface{} {
	if err := profile.userRepo.UpdateName(login, name); err != nil {
		return utils.Message("Error occured while updating name")
	}

	return utils.Message("Name has been updated!")
}

func (profile *ProfileService) UpdateGroup(login, group string) map[string]interface{} {
	if err := profile.userRepo.UpdateGroup(login, group); err != nil {
		return utils.Message("Error occured while updating group")
	}

	return utils.Message("Group has been updated!")
}

func (profile *ProfileService) UpdateSex(login string, sex bool) map[string]interface{} {
	if err := profile.userRepo.UpdateSex(login, sex); err != nil {
		return utils.Message("Error occured while updating sex")
	}

	return utils.Message("Sex has been updated!")
}
