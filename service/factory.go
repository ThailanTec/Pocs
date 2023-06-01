package service

import (
	"github.com/ThailanTec/factoriesStrategy/model"
	"github.com/ThailanTec/factoriesStrategy/strategy"
)

func NewDataBaseProviderF(format string) model.DBProvider {

	switch format {
	case "mssql":
		return &strategy.CreateSQL{}
	case "postgres":
		return &strategy.CreatePostgres{}
	/* case "oracle":
	return &strategy.CreateOracle{}
	*/
	default:
		return nil
	}

}
