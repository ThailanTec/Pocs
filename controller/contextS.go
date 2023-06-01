package controller

import "github.com/ThailanTec/factoriesStrategy/model"

type CreateDBStrategy struct {
	CreateDataBase model.DBProvider
}

func NewDataBase(data model.DBProvider) *CreateDBStrategy {
	return &CreateDBStrategy{
		CreateDataBase: data,
	}
}

func (c *CreateDBStrategy) SetCreated(dbCreate model.DBProvider) {
	c.CreateDataBase = dbCreate
}

// Refact: Vamos utilizar maps
func (d *CreateDBStrategy) CreateDataB(con model.CreateDB) {
	d.CreateDataBase.Create(con)
}
