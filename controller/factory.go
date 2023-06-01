package controller

import (
	"github.com/ThailanTec/factoriesStrategy/model"
	"github.com/ThailanTec/factoriesStrategy/strategy"
)

func NewDataBaseProviderF(dbtype string) model.DBProvider {

	switch dbtype {
	case "mssql":
		return &strategy.CreateSQL{}
	case "postgres":
		return &strategy.CreatePostgres{}
	case "oracle":
		return &strategy.CreateOracle{}

	default:
		return nil
	}

}
