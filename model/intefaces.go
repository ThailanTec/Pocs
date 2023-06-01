package model

type DBProvider interface {
	Create(DbDtoRequest)
}
