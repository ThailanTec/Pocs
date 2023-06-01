package strategy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ThailanTec/factoriesStrategy/model"
)

type CreatePostgres struct {
}

func (ps *CreatePostgres) Create(con model.CreateDB) {
	inputB := con

	const NewSourceDefinitionId = "decd338e-5647-4c0b-adf4-da0e75f5a750"

	dataF := inputB.ConnectionConfiguration
	dMashal, err := json.Marshal(dataF)
	if err != nil {
		fmt.Println("Não foi possivel realizar o Marshal.")
		return
	}
	con.SourceDefinitionID = NewSourceDefinitionId

	dataInput := make(map[string]interface{})
	err = json.Unmarshal([]byte(dMashal), &dataInput)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	// Chamada do Airbyte
	mash := inputB
	sqlMarshal, err := json.Marshal(mash)
	if err != nil {
		return
	}

	url := model.Url
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(sqlMarshal))
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", model.Authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Println(body)
		fmt.Sprintf("failed to postgres source. status code: %d", resp.StatusCode)
	}

	// Validação
	valida := valideDataP(dataInput, con)
	if valida {
		fmt.Printf("Create database type: ")
		fmt.Println(con.Format)
	} else {
		fmt.Println("Error:", err)
		return
	}

}

func valideDataP(input map[string]interface{}, con model.CreateDB) bool {

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

	replicationMethod, ok := dataInput["replicationMethod"]
	if !ok {
		fmt.Println("replicationMethod não encontrado ou não é uma string")
		return false
	} else {
		fmt.Println("replicationMethod:", replicationMethod)
	}

	return true
}
