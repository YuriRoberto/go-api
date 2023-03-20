package main

import (
	"os"

	"github.com/YuriRoberto/go-api/database"
	"github.com/YuriRoberto/go-api/routes"
	log "github.com/sirupsen/logrus"
)

//TODO - Revisar logs e definir níveis de logs (ver o q é tempo de execução e compilação)
//TODO - Rodar testes automatizados no Git
//TODO - Terminar os testes de todos os verbos de requisição
//TODO - Criar banco de dados com docker
//TODO - Subir api no docker(rodar testes no docker)
//TODO - Alimentar a api com o banco de dados do docker
func main() {
	level := os.Getenv("LOG_LEVEL")
	if level == "fatal" {
		log.SetLevel(log.FatalLevel)
	}
	if level == "debug" {
		log.SetLevel(log.DebugLevel)
	}
	log.Info("Iniciando API...")

	database.ConnectDataBase()

	log.Debug("Pokemons mockados")
	routes.HandleRequest()
	log.Info("Encerrando api...")
}
