package model

type DBProvider interface {
	Create(CreateDB)
}

type CreateDB struct {
	SourceDefinitionID      string                 `json:"sourceDefinitionId"` // Editar para que o back já sete automaticamente
	WorkspaceID             string                 `json:"workspaceId"`
	ConnectionConfiguration map[string]interface{} `json:"connectionConfiguration"`
	Name                    string                 `json:"name"`
	SourceName              string                 `json:"sourceName"`
	Format                  string                 `json:"format"`
}

const Authorization = "Basic YWlyYnl0ZV9hZG1pbjo3aVVld3p1WiFTWExs"
