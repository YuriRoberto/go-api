package routes

import (
	"net/http"

	"github.com/YuriRoberto/go-api/controllers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func HandleRequest() {
	log.Info("Loading routes...")
	r := mux.NewRouter()
	r.HandleFunc("/api/pokemons", controllers.AllPokemons).Methods("GET")
	r.HandleFunc("/api/pokemons/{id}/", controllers.ChosenPokemon).Methods("GET")
	r.HandleFunc("/api/pokemon", controllers.AddPokemon).Methods("POST")
	r.HandleFunc("/api/pokemon/{id}/", controllers.EditPokemon).Methods("PUT")
	//TODO - Fazer endpoint de edição
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatalf("Problema ao subir o ListenAndServer: %s", err.Error())
	}
	log.Info("Rotas carregadas!")
}
