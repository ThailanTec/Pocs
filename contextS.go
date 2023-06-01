package main

type CreateDBStrategy struct {
	CreateDataBase DBProvider
}

func NewDataBase(data DBProvider) *CreateDBStrategy {
	return &CreateDBStrategy{
		CreateDataBase: data,
	}
}

func (c *CreateDBStrategy) SetCreated(dbCreate DBProvider) {
	c.CreateDataBase = dbCreate
}

// Refact: Vamos utilizar maps
func (d *CreateDBStrategy) CreateDataB(con CreateDB) {
	d.CreateDataBase.Create(con)
}
