package models

type Pokemon struct {
	Id      int    `json:"id"`
	Tag     string `json:"tag"`
	Name    string `json:"name"`
	Img     string `json:"img"`
	Details string `json:"details"`
}

var Pokemons []Pokemon
