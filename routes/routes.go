package routes

import (
	"log"
	"net/http"

	"github.com/YuriRoberto/go-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/api/pokemons", controllers.AllPokemons).Methods("GET")
	r.HandleFunc("/api/pokemons/{id}/", controllers.ChosenPokemon).Methods("GET")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
