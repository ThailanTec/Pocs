package strategy

import (
	"encoding/json"
	"fmt"

	"github.com/ThailanTec/factoriesStrategy/model"
)

type CreateMongoDB struct {
}

func (sql *CreateMongoDB) Create(dto model.DbDtoRequest) {

	const SourceDefinitionId = "b39a7370-74c3-45a6-ac3a-380d48520a83"

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

	port, ok := dataInput["port"].(float64)
	if !ok {
		fmt.Println("port não encontrado ou não é uma inteiro")
		return false
	} else {
		fmt.Println("port:", port)
	}

	sslmode, ok := dataInput["sslMode"]
	if !ok {
		fmt.Println("sslMode não encontrado ou não é uma string")
		return false
	} else {
		fmt.Println("ssl:", sslmode)
	}

	schemas, ok := dataInput["schemas"]
	if !ok {
		fmt.Println("schemas não encontrado ou não é uma lista")
		return false
	} else {
		fmt.Println("schemas:", schemas)
	}

	host, ok := dataInput["host"]
	if !ok {
		fmt.Println("host não encontrado ou não é uma string")
		return false
	} else {
		fmt.Println("host:", host)
	}

	username, ok := dataInput["username"]
	if !ok {
		fmt.Println("username não encontrado ou não é uma string")
		return false
	} else {
		fmt.Println("username:", username)
	}

	password, ok := dataInput["password"]
	if !ok {
		fmt.Println("password não encontrado ou não é uma string")
		return false
	} else {
		fmt.Println("password:", password)
	}

	database, ok := dataInput["database"]
	if !ok {
		fmt.Println("database não encontrado ou não é uma string")
		return false
	} else {
		fmt.Println("database:", database)
	}

	sourceType, ok := dataInput["sourceType"]
	if !ok {
		fmt.Println("sourceType não encontrado ou não é uma string")
		return false
	} else {
		fmt.Println("sourceType:", sourceType)
	}

	return true
}
