package strategy

import (
	"encoding/json"
	"fmt"

	"github.com/ThailanTec/factoriesStrategy/model"
)

type CreateBigQuery struct {
}

func (sql *CreateBigQuery) Create(dto model.DbDtoRequest) {

	const SourceDefinitionId = "bfd1ddf8-ae8a-4620-b1d7-55597d2ba08c"

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
	valida := valideDataB(dataInput, dto)
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

func valideDataB(input map[string]interface{}, con model.DbDtoRequest) bool {

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

	project_id, ok := dataInput["project_id"].(string)
	if !ok {
		fmt.Println("project_id não encontrado ou não é uma string")
		return false
	} else {
		fmt.Println("project_id:", project_id)
	}

	dataset_id, ok := dataInput["dataset_id"].(string)
	if !ok {
		fmt.Println("dataset_id não encontrado ou não é uma string")
		return false
	} else {
		fmt.Println("dataset_id:", dataset_id)
	}

	credentials_json, ok := dataInput["credentials_json"].(string)
	if !ok {
		fmt.Println("credentials_json não encontrado ou não é uma string")
		return false
	} else {
		fmt.Println("credentials_json:", credentials_json)
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
