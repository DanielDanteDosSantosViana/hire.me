package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUrlInvalida(t *testing.T) {
	Convey("Dada uma URL inválida", t, func() {
		var urlInvalida = "invalida.com.br"

		Convey("Quando verifico se é uma URL válida", func() {
			retorno := IsUrl(urlInvalida)

			Convey("O valor retornado deve ser Falso", func() {
				So(retorno, ShouldEqual, false)
			})
		})
	})
}

func TestUrlValida(t *testing.T) {
	Convey("Dada uma URL valida", t, func() {
		var urlValida = "http://google.com.br"

		Convey("Quando verifico se é uma URL válida", func() {
			retorno := IsUrl(urlValida)

			Convey("O valor retornado deve ser true", func() {
				So(retorno, ShouldEqual, true)
			})
		})
	})
}
