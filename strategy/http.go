package strategy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ThailanTec/factoriesStrategy/model"
)

func NewHttpRequest(inputB model.CreateDB) error {

	const (
		Url = "http://34.105.89.111/api/v1/sources/create"
	)

	mash := inputB
	dbMarshal, err := json.Marshal(mash)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(dbMarshal))
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", model.Authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Sprintf("failed to create SQL Server source. Status code: %d", resp.StatusCode)
	}
	return nil
}
