package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/YuriRoberto/go-api/controllers"
)

func TestListPokemon(t *testing.T) {
	t.Run("Retornar os pokemons", func(t *testing.T) {
		requisicao, _ := http.NewRequest(http.MethodGet, "/api/pokemons", nil)
		resposta := httptest.NewRecorder()

		controllers.AllPokemons(resposta, requisicao)

		recebido := resposta.Body.String()
		log.Println(recebido)
		esperado := "{}"

		if recebido != esperado {
			t.Errorf("recebido '%s', esperado '%s'", recebido, esperado)
		}
	})
}
