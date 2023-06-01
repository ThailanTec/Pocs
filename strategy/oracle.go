package strategy

import (
	"fmt"

	"github.com/ThailanTec/factoriesStrategy/model"
)

type CreateOracle struct {
}

func (ps *CreateOracle) Create(con model.CreateDB) {

	fmt.Printf("Create database type: ")
	fmt.Print(con.Format)
}
