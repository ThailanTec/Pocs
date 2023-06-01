package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func mmomo() {

	url := "http://localhost:8000/api/v1/sources/create"
	method := "POST"

	payload := strings.NewReader(`{
    "sourceDefinitionId": "decd338e-5647-4c0b-adf4-da0e75f5a750",
    "sourceId": "646eb6e9-860b-45f5-a042-dbefb488817f",
    "workspaceId": "b2cceae1-84f2-4188-b160-ee393190c88d",
    "connectionConfiguration": {
        "dataset_name": "",
        "host": "enge-sql-server.database.windows.net",
        "schemas": [
            "dto"
        ],
        "port": 1433,
        "database": "mySampleDatabase",
        "username": "azureuser",
        "sslMode": "disable",
        "password": "Pwd@1234",
        "replicationMethod": "true"
    },
    "name": "SQLSERVERSSS",
    "sourceName": "SQLSERVERSSS",
    "icon": "icon"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic YWlyYnl0ZTpwYXNzd29yZA==")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
