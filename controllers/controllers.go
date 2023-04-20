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
	log.Info("Search all pokemons...")

	var allPokemons []models.Pokemon

	result := database.DB.Find(&allPokemons)
	if result.Error != nil {
		log.Info("entrou no if de error")
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
	log.Info("Adding pokemon...")
	var pokemon models.Pokemon
	json.NewDecoder(r.Body).Decode(&pokemon)

	log.Infof("pokemon to be added: %+v", pokemon)

	result := database.DB.Create(&pokemon)
	if result.Error != nil {
		log.Panic("Error in attempt to add a pokemon")
	}
	json.NewEncoder(w).Encode(pokemon)
	log.Debugf("The Pokemon: %+v", pokemon)
	log.Info("Pokemon has been added!")
}

func EditPokemon(w http.ResponseWriter, r *http.Request) {
	log.Info("Editing pokemon...")
	var pokemon models.Pokemon
	json.NewDecoder(r.Body).Decode(&pokemon)

	log.Infof("pokemon to be edited: %+v", pokemon)

	result := database.DB.Save(&pokemon)
	if result.Error != nil {
		log.Panic("Error in attempt to edit a pokemon")
	}
	json.NewEncoder(w).Encode(pokemon)
	log.Debugf("The Pokemon: %+v", pokemon)
	log.Info("Pokemons been edited!")
}
