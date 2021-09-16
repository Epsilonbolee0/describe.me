package service

import (
	"../../utils"
	"../builder"
	repo "../repository"
	"gorm.io/gorm"
)

type FunctionService struct {
	functionRepo *repo.FunctionRepository
	userRepo     *repo.UserRepository
}

func NewFunctionService(functionRepo *repo.FunctionRepository, userRepo *repo.UserRepository) *FunctionService {
	return &FunctionService{functionRepo, userRepo}
}

func (function *FunctionService) List(login string) map[string]interface{} {
	languages := function.userRepo.ListPreferredLanguages(login)

	idSlice := make([]uint, len(languages))
	for index, language := range languages {
		idSlice[index] = language.ID
	}

	functions, err := function.functionRepo.List(idSlice)
	if err != nil {
		return utils.ErrorOccured()
	}

	response := utils.Found()
	response["functions"] = functions
	return response
}

func (function *FunctionService) Create(lang uint, code string) map[string]interface{} {
	functionBuilder := builder.NewFunctionBuilder()
	functionEntity := functionBuilder.
		LanguageID(lang).
		Code(code).
		Build()

	if err := function.functionRepo.Create(&functionEntity); err != nil {
		return utils.ErrorOccured()
	}

	return utils.Created()
}

func (function *FunctionService) Find(id uint) map[string]interface{} {
	functionEntity, err := function.functionRepo.Find(id)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.NotFound()
	default:
		return utils.ErrorOccured()
	}

	response := utils.Found()
	response["function"] = functionEntity
	return response
}

func (function *FunctionService) UpdateCode(id uint, code string) map[string]interface{} {
	if err := function.functionRepo.UpdateCode(id, code); err != nil {
		return utils.ErrorOccured()
	}

	return utils.Updated()
}

func (function *FunctionService) Rating(id uint) map[string]interface{} {
	rating := function.functionRepo.Rating(id)
	response := utils.Found()
	response["rating"] = rating
	return response
}

func (function *FunctionService) Like(login string, id uint) map[string]interface{} {
	function.functionRepo.Like(login, id)
	return utils.Updated()
}

func (function *FunctionService) Dislike(login string, id uint) map[string]interface{} {
	function.functionRepo.Dislike(login, id)
	return utils.Updated()
}

func (function *FunctionService) Indifferent(login string, id uint) map[string]interface{} {
	function.functionRepo.Indifferent(login, id)
	return utils.Deleted()
}
