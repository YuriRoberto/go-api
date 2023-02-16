package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/YuriRoberto/go-api/models"
	log "github.com/sirupsen/logrus"
)

func AllPokemons(w http.ResponseWriter, r *http.Request) {
	log.Info("Procurando todos pokemons..")
	json.NewEncoder(w).Encode(models.Pokemons)
	log.Debugf("Pokemons sendo retornados: %+v", models.Pokemons)
	log.Info("Todos os Pokemons foram retornados!")
}

func ChosenPokemon(w http.ResponseWriter, r *http.Request) {
	log.Info("Procurando pokemon específico...")
	id := strings.Split(r.URL.Path, "/")[3]
	for _, pokemon := range models.Pokemons {
		if strconv.Itoa(pokemon.Id) == id {
			json.NewEncoder(w).Encode(pokemon)
			log.Debugf("Pokemon sendo retornado: %+v", pokemon)
			log.Info("O Pokemon foi retornado!")
		}
	}
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

	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	// if err != nil {
	// 	log.Error("Não foi possível adicionar, motivo: %s", err.Error())
	// }

}
