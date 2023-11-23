package main

import (
	"demoapp/application"
	"demoapp/infrastructure/storage/mongodb"
	"demoapp/infrastructure/transport/restapi"
	"fmt"
)

func main() {
	// Simula que uma requisição HTTP chegou
	pReq := restapi.PersonReqRes{
		Name: "John Doe",
		Age:  30,
	}

	// Simula que o controller recebeu a requisição e chamou o caso de uso
	pIn := pReq.ToPersonOutput()

	// Simula que o caso de uso chamou o domínio
	p := pIn.ToPerson()
	p.Birthday()

	// Simula que o caso de uso atualizou a pessoa no repositório
	mp := mongodb.NewMongoPersonFromPerson(p)
	pUpdated := mp.ToPerson()

	// Simula que o caso de uso retornou a pessoa atualizada
	pOutUpdated := application.NewPersonOutputFromPerson(pUpdated)

	// Simula que o controller retornou a resposta HTTP
	pRes := restapi.NewPersonReqResFromPersonOutput(pOutUpdated)

	fmt.Printf("Person response: %#v\n", pRes)
}
