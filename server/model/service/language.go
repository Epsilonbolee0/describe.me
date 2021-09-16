package service

import (
	"../../utils"
	"../builder"
	repo "../repository"
)

type LanguageService struct {
	languageRepo *repo.LanguageRepository
	userRepo     *repo.UserRepository
}

func NewLanguageService(languageRepo *repo.LanguageRepository, userRepo *repo.UserRepository) *LanguageService {
	return &LanguageService{languageRepo, userRepo}
}

func (lang *LanguageService) Create(login, name string) map[string]interface{} {
	if ok, err := lang.userRepo.IsAdmin(login); err != nil || !ok {
		return utils.NoRight()
	}

	languageBuilder := builder.NewLanguageBuilder()
	language := languageBuilder.Name(name).Build()

	if err := lang.languageRepo.Create(&language); err != nil {
		return utils.ErrorOccured()
	}

	return utils.Created()
}

func (lang *LanguageService) List() map[string]interface{} {
	languages, err := lang.languageRepo.List()
	if err != nil {
		return utils.ErrorOccured()
	}

	resp := utils.Found()
	resp["languages"] = languages
	return resp
}

func (lang *LanguageService) Find(name string) map[string]interface{} {
	language, err := lang.languageRepo.Find(name)
	if err != nil {
		return utils.ErrorOccured()
	}

	resp := utils.Found()
	resp["language"] = language
	return resp
}

func (lang *LanguageService) Delete(login string, id uint) map[string]interface{} {
	if ok, err := lang.userRepo.IsAdmin(login); err != nil || !ok {
		return utils.NoRight()
	}

	if err := lang.languageRepo.Delete(id); err != nil {
		return utils.CantDelete()
	}

	return utils.Deleted()
}
