package strategy

import (
	"fmt"

	"github.com/ThailanTec/factoriesStrategy/model"
)

type CreateOracle struct {
}

func (ps *CreateOracle) Create(con model.DtoAirbyteCreateDB) {

	fmt.Printf("Create database type: ")
	fmt.Print(con.Format)
}
