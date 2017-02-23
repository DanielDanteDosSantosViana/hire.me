package util

import (
	"testing"
	"github.com/DanielDanteDosSantosViana/hire.me/config"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	config.Conf.Base.Alfabeto = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
}

func TestConverterInteiro18446744073709551615ParaAliasv8QrKbgkrIp(t *testing.T) {
	Convey("Dado um inteiro (int64) 18446744073709551615", t, func() {
		var inteiro  uint64 = 18446744073709551615
		var alias string = "v8QrKbgkrIp"
		Convey("Quando converto inteiro 18446744073709551615 para string ", func() {
			retorno := InteiroParaString(inteiro)

			Convey("O valor retornado deve ser o alias 'v8QrKbgKrIp' ", func() {
				So(retorno, ShouldEqual, alias)
			})
		})
	})
}

func TestConverterPrimeiroElementoDaSequence(t *testing.T) {
	Convey("Dado o primeiro elemento da sequence (int64) 0", t, func() {
		var inteiro  uint64 = 0
		var alias string = "a"
		Convey("Quando converto inteiro 0 para string ", func() {
			retorno := InteiroParaString(inteiro)

			Convey("O valor retornado deve ser a primeira letra contida em  'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789' ", func() {
				So(retorno, ShouldEqual, alias)
			})
		})
	})
}


