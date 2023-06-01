package service

import (
	"github.com/ThailanTec/factoriesStrategy/model"
)

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

func (d *CreateDBStrategy) CreateDataB(dto model.DbDtoRequest) {
	d.CreateDataBase.Create(dto)
}
