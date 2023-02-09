package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/YuriRoberto/go-api/models"
	"github.com/stretchr/testify/require"
)

func TestListPokemons(t *testing.T) {
	t.Run("Retornar os pokemons", func(t *testing.T) {
		models.Pokemons = []models.Pokemon{
			{
				Id:      1,
				Name:    "Bulbasaur",
				Tag:     "#0001",
				Img:     "https://assets.pokemon.com/assets/cms2/img/pokedex/full/001.png",
				Details: "Pokemon do tipo planta.",
			},
			{
				Id:      4,
				Name:    "Charmander",
				Tag:     "#0004",
				Img:     "https://assets.pokemon.com/assets/cms2/img/pokedex/full/004.png",
				Details: "Pokemon do tipo fogo.",
			},
		}
		requisicao, err := http.NewRequest(http.MethodGet, "/api/pokemons", nil)
		if err != nil {
			t.Fatal("ERROR:", err)
		}
		resposta := httptest.NewRecorder()
		defer resposta.Result().Body.Close()
		AllPokemons(resposta, requisicao)
		recebido := resposta.Body.String()
		esperado, err := json.Marshal(&[]models.Pokemon{
			{
				Id:      1,
				Name:    "Bulbasaur",
				Tag:     "#0001",
				Img:     "https://assets.pokemon.com/assets/cms2/img/pokedex/full/001.png",
				Details: "Pokemon do tipo planta.",
			},
			{
				Id:      4,
				Name:    "Charmander",
				Tag:     "#0004",
				Img:     "https://assets.pokemon.com/assets/cms2/img/pokedex/full/004.png",
				Details: "Pokemon do tipo fogo.",
			},
		})
		if err != nil {
			t.Fatal("ERROR:", err)
		}
		require.JSONEq(t, string(esperado), string(recebido))
	})
}
func TestListOnePokemon(t *testing.T) {
	t.Run("Retornar um pokemon", func(t *testing.T) {
		models.Pokemons = []models.Pokemon{
			{
				Id:      1,
				Name:    "Bulbasaur",
				Tag:     "#0001",
				Img:     "https://assets.pokemon.com/assets/cms2/img/pokedex/full/001.png",
				Details: "Pokemon do tipo planta.",
			},
			{
				Id:      4,
				Name:    "Charmander",
				Tag:     "#0004",
				Img:     "https://assets.pokemon.com/assets/cms2/img/pokedex/full/004.png",
				Details: "Pokemon do tipo fogo.",
			},
		}
		requisicao, err := http.NewRequest(http.MethodGet, "/api/pokemons/4", nil)
		if err != nil {
			t.Fatal("ERROR:", err)
		}
		resposta := httptest.NewRecorder()
		defer resposta.Result().Body.Close()
		ChosenPokemon(resposta, requisicao)
		recebido := resposta.Body.String()
		esperado, err := json.Marshal(&models.Pokemon{
			Id:      4,
			Name:    "Charmander",
			Tag:     "#0004",
			Img:     "https://assets.pokemon.com/assets/cms2/img/pokedex/full/004.png",
			Details: "Pokemon do tipo fogo.",
		})
		if err != nil {
			t.Fatal("ERROR:", err)
		}
		require.JSONEq(t, string(esperado), string(recebido))
	})
}
