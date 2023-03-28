package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/YuriRoberto/go-api/database"
	"github.com/YuriRoberto/go-api/models"
	log "github.com/sirupsen/logrus"
)

func AllPokemons(w http.ResponseWriter, r *http.Request) {
	log.Info("Procurando todos pokemons..")
	allPokemons := models.Pokemons
	result := database.DB.Find(&allPokemons)
	if result.Error != nil {
		log.Panic("Error in find all pokemons")
	}

	json.NewEncoder(w).Encode(allPokemons)
	log.Debugf("Pokemons sendo retornados: %+v", models.Pokemons)
	log.Info("All pokemons are returned!")
}

func ChosenPokemon(w http.ResponseWriter, r *http.Request) {
	log.Info("Search pokemon...")
	id := strings.Split(r.URL.Path, "/")[3]

	pokemon := models.Pokemons
	result := database.DB.First(&pokemon, id)
	if result.Error != nil {
		log.Panic("Error in find a pokemon")
	}

	json.NewEncoder(w).Encode(pokemon)
	log.Info("The pokemon is returned!")
	// for _, pokemon := range models.Pokemons {
	// 	if strconv.Itoa(pokemon.Id) == id {
	// 		json.NewEncoder(w).Encode(pokemon)
	// 		log.Debugf("Pokemon sendo retornado: %+v", pokemon)
	// 		log.Info("The pokemon is returned!")
	// 	}
	// }
}

func AddPokemon(w http.ResponseWriter, r *http.Request) {
	log.Info("Adicionando pokemon...")
	decoder := json.NewDecoder(r.Body)
	var pokemon models.Pokemon
	decoder.Decode(&pokemon)
	json.NewEncoder(w).Encode(pokemon)
	log.Debugf("O Pokemon sendo adicionado: %+v", pokemon)
	log.Info("O Pokemon foi adicionado!")
	// fmt.Println(w)
}

func EditPokemon(w http.ResponseWriter, r *http.Request) {
	log.Info("Editando pokemon...")
	decoder := json.NewDecoder(r.Body)
	var pokemon models.Pokemon
	decoder.Decode(&pokemon)
	json.NewEncoder(w).Encode(pokemon)
	log.Debugf("O Pokemon sendo editado: %+v", pokemon)
	log.Info("O Pokemon foi editado!")
}
