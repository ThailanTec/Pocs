package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ThailanTec/factoriesStrategy/model"
	"github.com/ThailanTec/factoriesStrategy/service"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.POST("/pocs/connection", Connection)
	router.GET("/pocs/getdb", GetDB)
	http.ListenAndServe(":8080", router)

}

func Connection(c *gin.Context) {
	c.Set("Content-Type", "application/json")
	data := model.DbDtoRequest{}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind Json" + err.Error(),
		})
		return
	}

	criacao := service.NewDataBase(service.NewDataBaseProviderF(data.Format))
	criacao.CreateDataB(data)

	c.JSON(http.StatusCreated, data)
}

func GetDB(c *gin.Context) {

	data := model.DBStruct{Content: model.Content{
		Body: []model.Body{{UID: "123456", Component: "dataDB", Headline: "DB", Title: "MYSQL"}},
	}}

	db, err := json.Marshal(data)

	if err != nil {
		return
	}

	fmt.Println(db)

	c.SecureJSON(200, data)
}
