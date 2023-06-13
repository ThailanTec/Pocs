package strategy

import (
	"encoding/json"
	"fmt"

	"github.com/ThailanTec/factoriesStrategy/model"
)

type CreateMongoDB struct {
}

func (sql *CreateMongoDB) Create(dto model.DbDtoRequest) {

	const SourceDefinitionId = "b2e713cd-cc36-4c0a-b5bd-b47cb8a0561e"

	dataF := dto.ConnectionConfiguration
	dMashal, err := json.Marshal(dataF)
	if err != nil {
		fmt.Println("Não foi possivel realizar o Marshal.")
		return
	}

	dataInput := make(map[string]interface{})
	err = json.Unmarshal([]byte(dMashal), &dataInput)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	// Validação
	valida := valideDataM(dataInput, dto)
	if valida {
		fmt.Printf("Create database type: ")
		fmt.Printf(dto.Format)
	} else {
		fmt.Println("Error:", err)
		return
	}

	createDb := model.DtoMapp(dto, SourceDefinitionId)

	// Chamada do Airbyte
	NewHttpRequest(*createDb)

}

func valideDataM(input map[string]interface{}, con model.DbDtoRequest) bool {

	conn := con.ConnectionConfiguration

	dMashal, err := json.Marshal(conn)
	if err != nil {
		fmt.Println("Não foi possivel realizar o Marshal.")
		return false
	}

	dataInput := make(map[string]interface{})
	err = json.Unmarshal([]byte(dMashal), &dataInput)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return false
	}

	database, ok := dataInput["database"].(string)
	if !ok {
		fmt.Println("database não encontrado ou não é uma string")
		return false
	} else {
		fmt.Println("database:", database)
	}

	instance_type, ok := dataInput["instance_type"]
	if !ok {
		fmt.Println("Instance Type não encontrado ou não é uma lista")
		return false
	} else {
		fmt.Println("schemas:", instance_type)
	}

	user, ok := dataInput["user"].(string)
	if !ok {
		fmt.Println("user não encontrado ou não é uma string")
		return false
	} else {
		fmt.Println("user:", user)
	}

	password, ok := dataInput["password"].(string)
	if !ok {
		fmt.Println("password não encontrado ou não é uma string")
		return false
	} else {
		fmt.Println("password:", password)
	}

	return true
}
