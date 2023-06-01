package main

func NewDataBaseProviderF(dbtype string) DBProvider {

	switch dbtype {
	case "mssql":
		return &CreateSQL{}
	case "postgres":
		return &CreatePostgres{}
	case "oracle":
		return &CreateOracle{}

	default:
		return nil
	}

}
