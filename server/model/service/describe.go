package service

import (
	"../../utils"
	"../builder"
	repo "../repository"
	"gorm.io/gorm"
)

type DescribeService struct {
	functionRepo *repo.FunctionRepository
}

func NewDescribeService(functionRepo *repo.FunctionRepository) *DescribeService {
	return &DescribeService{functionRepo}
}

func (describe *DescribeService) Create(lang uint, code string) map[string]interface{} {
	functionBuilder := builder.NewFunctionBuilder()
	function := functionBuilder.
		LanguageID(lang).
		Code(code).
		Build()

	if err := describe.functionRepo.Create(&function); err != nil {
		return utils.ErrorOccured()
	}

	return utils.Created()
}

func (describe *DescribeService) Find(id uint) map[string]interface{} {
	function, err := describe.functionRepo.Find(id)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.NotFound()
	default:
		return utils.ErrorOccured()
	}

	response := utils.Found()
	response["function"] = function
	return response
}

func (describe *DescribeService) UpdateCode(id uint, code string) map[string]interface{} {
	if err := describe.functionRepo.UpdateCode(id, code); err != nil {
		return utils.ErrorOccured()
	}

	return utils.Updated()
}

func (describe *DescribeService) Rating(id uint) map[string]interface{} {
	rating := describe.functionRepo.Rating(id)
	response := utils.Found()
	response["rating"] = rating
	return response
}

func (describe *DescribeService) Like(login string, id uint) map[string]interface{} {
	describe.functionRepo.Like(login, id)
	return utils.Updated()
}

func (describe *DescribeService) Dislike(login string, id uint) map[string]interface{} {
	describe.functionRepo.Dislike(login, id)
	return utils.Updated()
}

func (describe *DescribeService) Indifferent(login string, id uint) map[string]interface{} {
	describe.functionRepo.Indifferent(login, id)
	return utils.Deleted()
}
