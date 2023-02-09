package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/YuriRoberto/go-api/models"
)

func AllPokemons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Pokemons)
	log.Println("all pokemons in json")
}

func ChosenPokemon(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[3]
	for _, pokemon := range models.Pokemons {
		if strconv.Itoa(pokemon.Id) == id {
			log.Println("entreou digraça")
			json.NewEncoder(w).Encode(pokemon)
			log.Println("chosen pokemon in json")
		}
	}
}
