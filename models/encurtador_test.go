package models
import (
	"testing"
	"github.com/DanielDanteDosSantosViana/hire.me/config"
	"github.com/DanielDanteDosSantosViana/hire.me/db"

	. "github.com/smartystreets/goconvey/convey"
)
func init(){
		config.Conf.Base.Alfabeto = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		config.Conf.Db.MysqlWrite ="root:1234@/encurtar_url?charset=utf8"
		config.Conf.Db.MysqlRead ="root:1234@/encurtar_url?charset=utf8"
}

func TestCriacaoUrlEncurtada(t *testing.T) {
	Convey("Dado uma url válida", t, func() {
		encurtador := NewEncurtador() 
		dbUtil:= db.NewDBUtil()
		var url string = "http://www.bemobi.com.br"

		Convey("Quando faz o encurtamento", func() {
			urlEncurtada, _ := encurtador.Encurtar(url,"")
			Convey("Deve retornar uma url encurtada", func() {
				So(urlEncurtada.Alias, ShouldNotBeBlank)
			})
			Reset(func() {
        	dbUtil.DeletePorAlias(urlEncurtada.Alias)
    		})
		})
		Reset(func() {
        	encurtador.close()
        	dbUtil.Close()
    	})
	})
}


func TestBuscarAliasExistente(t *testing.T) {
	Convey("Ao encurtar uma url", t, func() {
		encurtador := NewEncurtador() 
		dbUtil:= db.NewDBUtil()
		var url string = "http://www.bemobi.com.br"
		urlEncurtada, _ := encurtador.Encurtar(url,"")

		Convey("Quando busco pelo alias criado", func() {
			urlEncurtada, _ := encurtador.BuscarPorAlias(urlEncurtada.Alias)
			Convey("Deve retornar a url encurtada", func() {
				So(urlEncurtada.Alias, ShouldNotBeBlank)
			})
			Reset(func() {
        		dbUtil.DeletePorAlias(urlEncurtada.Alias)
    		})
		})
		Reset(func() {
        	encurtador.close()
        	dbUtil.Close()
    	})
	})
}

func TestAtualizarNumeroDeAcessos(t *testing.T) {
	Convey("Dado uma nova atualização de acesso na url", t, func() {
		encurtador := NewEncurtador() 
		dbUtil:= db.NewDBUtil()
		var url string = "http://www.bemobi.com.br"
		urlEncurtada, _ := encurtador.Encurtar(url,"")

		Convey("Quando atualiza o número de acessos", func() {
			encurtador.atualizarNumeroDeAcessos(urlEncurtada.Alias,urlEncurtada.Acessos)
			urlEncurtada_incrementada, _ := encurtador.BuscarPorAlias(urlEncurtada.Alias)

			Convey("O número de acessos deve ser incrementado ", func() {
				So(urlEncurtada.Acessos, ShouldBeLessThan,urlEncurtada_incrementada.Acessos)
			})
			Reset(func() {
        		dbUtil.DeletePorAlias(urlEncurtada.Alias)
    		})
		})
		Reset(func() {
        	encurtador.close()
        	dbUtil.Close()
    	})
	})
}


