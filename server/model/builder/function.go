package builder

import (
	"../domain"
)

type functionModifier func(account *domain.Function)
type FunctionBuilder struct {
	actions []functionModifier
}

func NewFunctionBuilder() *FunctionBuilder {
	return &FunctionBuilder{}
}

func (b *FunctionBuilder) Build() domain.Function {
	function := domain.Function{}
	for _, action := range b.actions {
		action(&function)
	}

	return function
}

func (b *FunctionBuilder) Code(code string) *FunctionBuilder {
	b.actions = append(b.actions, func(function *domain.Function) {
		function.Code = code
	})
	return b
}

func (b *FunctionBuilder) LanguageID(languageID uint) *FunctionBuilder {
	b.actions = append(b.actions, func(function *domain.Function) {
		function.LanguageID = languageID
	})
	return b
}
