package builder

import (
	"../domain"
)

type descriptionModifier func(description *domain.Description)
type DescriptionBuilder struct {
	actions []descriptionModifier
}

func NewDescriptionBuilder() *DescriptionBuilder {
	return &DescriptionBuilder{}
}

func (b *DescriptionBuilder) Build() domain.Description {
	description := domain.Description{}
	for _, action := range b.actions {
		action(&description)
	}

	return description
}

func (b *DescriptionBuilder) Content(content string) *DescriptionBuilder {
	b.actions = append(b.actions, func(function *domain.Description) {
		function.Content = content
	})
	return b
}

func (b *DescriptionBuilder) Author(authorLogin string) *DescriptionBuilder {
	b.actions = append(b.actions, func(function *domain.Description) {
		function.AuthorLogin = authorLogin
	})
	return b
}

func (b *DescriptionBuilder) FunctionID(functionID uint) *DescriptionBuilder {
	b.actions = append(b.actions, func(function *domain.Description) {
		function.FunctionID = functionID
	})
	return b
}
