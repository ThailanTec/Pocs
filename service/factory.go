package service

import (
	"github.com/ThailanTec/factoriesStrategy/model"
	"github.com/ThailanTec/factoriesStrategy/strategy"
)

func NewDataBaseProviderF(format string) model.DBProvider {

	switch format {
	case "mssql":
		return &strategy.CreateSQLServer{}
	case "postgres":
		return &strategy.CreatePostgres{}
	case "oracle":
		return &strategy.CreateOracle{}
	case "mysql":
		return &strategy.CreateMysql{}
	case "bigQuery":
		return &strategy.CreateBigQuery{}
	case "redshift":
		return &strategy.CreatRedShift{}
	default:
		return nil
	}

}
