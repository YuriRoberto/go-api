package main

import (
	"fmt"

	"github.com/YuriRoberto/go-api/models"
	"github.com/YuriRoberto/go-api/routes"
)

//TODO - Revisar logs e definir níveis de logs
//TODO - Rodar testes automatizados no Git
//TODO - Terminar os testes de todos os verbos de requisição
//TODO - Criar banco de dados com docker
//TODO - Subir api no docker(rodar testes no docker)
//TODO - Alimentar a api com o banco de dados do docker
func main() {
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

	fmt.Println("Iniciando algoooooooooo")
	routes.HandleRequest()
}
