package models

import "errors"

var (
	ErrorUrlNaoInformada = errors.New("URL não foi informada")
	ErrorAliasExiste     = errors.New("Alias existente")
	ErrorInterno         = errors.New("Ocorreu um error interno na API")
)

type ErrorReturn struct {
	CodigoError string `json:"codigo"`
	Descricao   string `json:"descricao"`
}

func NewErrorReturn(codigo string, descricao string) *ErrorReturn {
	return &ErrorReturn{codigo, descricao}
}

func NewErrorUrlNaoInformada() *ErrorReturn {
	return &ErrorReturn{"003", ErrorUrlNaoInformada.Error()}
}

func NewErrorAliasExiste() *ErrorReturn {
	return &ErrorReturn{"001", ErrorAliasExiste.Error()}
}
func NewErrorInterno() *ErrorReturn {
	return &ErrorReturn{"000", ErrorInterno.Error()}
}
