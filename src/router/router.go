package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

//gerar as rotas da funcao
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
