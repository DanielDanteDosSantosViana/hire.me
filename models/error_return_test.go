package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRetornoErrorUrlNaoInformada(t *testing.T) {
	Convey("Dado uma excecao de URL não informada", t, func() {
		error := NewErrorUrlNaoInformada() 

		Convey("Quando recupero a descrição do Error", func() {
			descricao := error.Descricao

			Convey("A descrição deve ser igual a existente no ErrorUrlNaoInformada", func() {
				So(descricao, ShouldEqual, ErrorUrlNaoInformada.Error())
			})
		})
	})
}

func TestRetornoErrorAliasExiste(t *testing.T) {
	Convey("Dado uma excecao de alias existente", t, func() {
		error := NewErrorAliasExiste() 

		Convey("Quando recupero a descrição do Error", func() {
			descricao := error.Descricao

			Convey("A descrição deve ser igual a existente no ErrorAliasExiste", func() {
				So(descricao, ShouldEqual, ErrorAliasExiste.Error())
			})
		})
	})
}

func TestRetornoErrorInterno(t *testing.T) {
	Convey("Dado uma excecao de error interno", t, func() {
		error := NewErrorInterno() 

		Convey("Quando recupero a descrição do Error", func() {
			descricao := error.Descricao

			Convey("A descrição deve ser igual a existente no ErrorInterno", func() {
				So(descricao, ShouldEqual, ErrorInterno.Error())
			})
		})
	})
}

func TestRetornoErrorAliasNaoEncontrado(t *testing.T) {
	Convey("Dado uma excecao de alias nao encontrado", t, func() {
		error := NewErrorAliasNaoEncontrado() 

		Convey("Quando recupero a descrição do Error", func() {
			descricao := error.Descricao

			Convey("A descrição deve ser igual a existente no ErrorAliasNaoEncontrado", func() {
				So(descricao, ShouldEqual, ErrorAliasNaoEncontrado.Error())
			})
		})
	})
}

func TestRetornoErrorUrlInvalida(t *testing.T) {
	Convey("Dado uma excecao de url inválida", t, func() {
		error := NewErrorUrlInvalida("www.google.com") 

		Convey("Quando recupero a descrição do Error", func() {
			descricao := error.Descricao

			Convey("A descrição deve ser igual a existente no ErrorUrlInvalida", func() {
				So(descricao, ShouldEqual, ErrorUrlInvalida.Error())
			})
		})
	})
}

func TestRetornoErrorRedirectURL(t *testing.T) {
	Convey("Dado uma excecao de redirect url", t, func() {
		error := NewErrorToRedirect("www.google.com") 

		Convey("Quando recupero a descrição do Error", func() {
			descricao := error.Descricao

			Convey("A descrição deve ser igual a existente no ErrorRedirectURL", func() {
				So(descricao, ShouldEqual, ErrorRedirectURL.Error())
			})
		})
	})
}

