package main

import "fmt"

type CreateOracle struct {
}

func (ps *CreateOracle) Create(con CreateDB) {

	fmt.Printf("Create database type: ")
	fmt.Print(con.Format)
}
