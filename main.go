package main

import (
	"net/http"

	"github.com/ThailanTec/factoriesStrategy/controller"
	"github.com/ThailanTec/factoriesStrategy/model"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.POST("/pocs/connection", Connection)
	http.ListenAndServe(":8080", router)

}

func Connection(c *gin.Context) {

	data := model.CreateDB{}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind Json" + err.Error(),
		})
		return
	}
	criacao := controller.NewDataBase(controller.NewDataBaseProviderF(data.Format))
	criacao.CreateDataB(data)

	c.JSON(http.StatusCreated, data)
}
