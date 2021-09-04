package service

import (
	"../../utils"
	"../builder"
	repo "../repository"
)

type LanguageService struct {
	languageRepo *repo.LanguageRepository
}

func NewLanguageService(languageRepo *repo.LanguageRepository) *LanguageService {
	return &LanguageService{languageRepo}
}

func (lang *LanguageService) Create(name string) map[string]interface{} {
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

func (lang *LanguageService) Delete(id uint) map[string]interface{} {
	if err := lang.languageRepo.Delete(id); err != nil {
		return utils.CantDelete()
	}

	return utils.Deleted()
}
