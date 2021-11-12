package service

import (
	"../../utils"
	"../builder"
	repo "../repository"
	"gorm.io/gorm"
)

type DescribeService struct {
	descriptionRepo *repo.DescriptionRepository
}

func NewDescribeService(descriptionRepo *repo.DescriptionRepository) *DescribeService {
	return &DescribeService{descriptionRepo}
}

func (describe *DescribeService) Create(lang uint, code string) map[string]interface{} {
	descriptionBuilder := builder.NewDescriptionBuilder()
	description := descriptionBuilder.
		FunctionID(lang).
		Content(code).
		Build()

	if err := describe.descriptionRepo.Create(&description); err != nil {
		return utils.ErrorOccured()
	}

	return utils.Created()
}

func (describe *DescribeService) ListByFunction(id uint) map[string]interface{} {
	descriptions, err := describe.descriptionRepo.ListByFunction(id)
	switch err {
	case nil:
		break
	case gorm.ErrEmptySlice:
		return utils.NotFound()
	default:
		return utils.ErrorOccured()
	}

	response := utils.Found()
	response["descriptions"] = descriptions
	return response
}

func (describe *DescribeService) Delete(id uint) map[string]interface{} {
	if err := describe.descriptionRepo.Delete(id); err != nil {
		return utils.ErrorOccured()
	}

	return utils.Deleted()
}

func (describe *DescribeService) Rating(id uint) map[string]interface{} {
	rating := describe.descriptionRepo.Rating(id)
	response := utils.Found()
	response["rating"] = rating
	return response
}

func (describe *DescribeService) Like(login string, id uint) map[string]interface{} {
	describe.descriptionRepo.Like(login, id)
	return utils.Updated()
}

func (describe *DescribeService) Dislike(login string, id uint) map[string]interface{} {
	describe.descriptionRepo.Dislike(login, id)
	return utils.Updated()
}

func (describe *DescribeService) Indifferent(login string, id uint) map[string]interface{} {
	describe.descriptionRepo.Indifferent(login, id)
	return utils.Deleted()
}
