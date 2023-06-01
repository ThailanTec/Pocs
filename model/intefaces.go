package model

type DBProvider interface {
	Create(CreateDB)
}

type CreateDB struct {
	SourceDefinitionID      string                 `json:"sourceDefinitionId"`
	SourceID                string                 `json:"sourceId"`
	WorkspaceID             string                 `json:"workspaceId"`
	ConnectionConfiguration map[string]interface{} `json:"connectionConfiguration"`
	Name                    string                 `json:"name"`
	SourceName              string                 `json:"sourceName"`
	Format                  string                 `json:"format"`
}

const (
	Url           = "http://34.105.89.111/api/v1/sources/create"
	Authorization = "Basic YWlyYnl0ZV9hZG1pbjo3aVVld3p1WiFTWExs"
)
