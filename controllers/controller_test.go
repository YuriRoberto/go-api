package controllers

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/YuriRoberto/go-api/models"
	"github.com/stretchr/testify/require"
	"github.com/YuriRoberto/go-api/database"
)

func TestAddPokemon(t *testing.T) {
	t.Run("Adicionar um pokemon", func(t *testing.T) {
		database.ConnectDatabase()
		payload := "{\"tag\":\"#001\",\"name\":\"Bulbasaur\",\"img\":\"https://assets.pokemon.com/assets/cms2/img/pokedex/detail/001.png\",\"details\":\"Este animal tem uma planta nas costas e a usa para se defender.\"}"
		body := strings.NewReader(payload)
		requisicao, err := http.NewRequest(http.MethodPost, "/api/pokemon", body)
		if err != nil {
			t.Fatal("ERROR:", err)
		}
		resposta := httptest.NewRecorder()
		defer resposta.Result().Body.Close()
		AddPokemon(resposta, requisicao)
		recebido := resposta.Body.String()
		esperado, err := json.Marshal(&models.Pokemon{Id:1,Tag:"#001",Name:"Bulbasaur",Img:"https://assets.pokemon.com/assets/cms2/img/pokedex/detail/001.png",Details:"Este animal tem uma planta nas costas e a usa para se defender."})
		if err != nil {
			t.Fatal("ERROR:", err)
		}
		require.JSONEq(t, string(esperado), recebido)
	})
}
func TestListPokemons(t *testing.T) {
	t.Run("Retornar os pokemons", func(t *testing.T) {
		database.ConnectDatabase()
		requisicao, err := http.NewRequest(http.MethodGet, "http://localhost:8000/api/pokemons", nil)
		if err != nil {
			t.Fatal("ERROR:", err)
		}
		resposta := httptest.NewRecorder()
		defer resposta.Result().Body.Close()
		AllPokemons(resposta, requisicao)
		recebido := resposta.Body.String()
		esperado, err := json.Marshal(&[]models.Pokemon{
			{Id:1,Tag:"#001",Name:"Bulbasaur",Img:"https://assets.pokemon.com/assets/cms2/img/pokedex/detail/001.png",Details:"Este animal tem uma planta nas costas e a usa para se defender."},
		})
		if err != nil {
			t.Fatal("ERROR:", err)
		}
		require.JSONEq(t, string(esperado), recebido)
	})
}
// func TestListOnePokemon(t *testing.T) {
// 	t.Run("Listar um pokemon", func(t *testing.T) {
// 		requisicao, err := http.NewRequest(http.MethodGet, "/api/pokemons/4", nil)
// 		if err != nil {
// 			t.Fatal("ERROR:", err)
// 		}
// 		resposta := httptest.NewRecorder()
// 		defer resposta.Result().Body.Close()
// 		ChosenPokemon(resposta, requisicao)
// 		recebido := resposta.Body.String()
// 		esperado, err := json.Marshal(&models.Pokemon{
// 			Id:      4,
// 			Name:    "Charmander",
// 			Tag:     "#0004",
// 			Img:     "https://assets.pokemon.com/assets/cms2/img/pokedex/full/004.png",
// 			Details: "Pokemon do tipo fogo.",
// 		})
// 		if err != nil {
// 			t.Fatal("ERROR:", err)
// 		}
// 		require.JSONEq(t, string(esperado), string(recebido))
// 	})
// }
// func TestEditPokemon(t *testing.T) {
// 	t.Run("Editar um pokemon", func(t *testing.T) {
// 		//TODO - Validar que nao tem nada
// 		//TODO - Criar algum pokemon
// 		//TODO - Criar mais um pokemon
// 		//TODO - Editar um deles e valido q so e somente esse escolhido foi alterado
// 		payload := "{\"id\":4,\"tag\":\"#0004\",\"name\":\"CharmanderNovooo\",\"img\":\"https://assets.pokemon.com/assets/cms2/img/pokedex/full/004.png\",\"details\":\"Pokemon do tipo fogo.\"}"
// 		body := strings.NewReader(payload)
// 		requisicao, err := http.NewRequest(http.MethodPost, "/api/pokemon/4", body)
// 		if err != nil {
// 			t.Fatal("ERROR:", err)
// 		}
// 		resposta := httptest.NewRecorder()
// 		defer resposta.Result().Body.Close()
// 		EditPokemon(resposta, requisicao)
// 		recebido := resposta.Body.String()
// 		fmt.Println(recebido)
// 		esperado, err := json.Marshal(&models.Pokemon{
// 			Id:      4,
// 			Name:    "CharmanderNovooo",
// 			Tag:     "#0004",
// 			Img:     "https://assets.pokemon.com/assets/cms2/img/pokedex/full/004.png",
// 			Details: "Pokemon do tipo fogo.",
// 		})
// 		if err != nil {
// 			t.Fatal("ERROR:", err)
// 		}
// 		fmt.Println(string(esperado))
// 		require.JSONEq(t, string(esperado), string(recebido))
// 	})
// }
