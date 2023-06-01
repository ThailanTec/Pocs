package model

type DbDtoRequest struct {
	WorkspaceID             string                 `json:"workspaceId"`
	ConnectionConfiguration map[string]interface{} `json:"connectionConfiguration"`
	Name                    string                 `json:"name"`
	SourceName              string                 `json:"sourceName"`
	Format                  string                 `json:"format"`
}

type DtoAirbyteCreateDB struct {
	SourceDefinitionID      string                 `json:"sourceDefinitionId"` // Editar para que o back já sete automaticamente
	WorkspaceID             string                 `json:"workspaceId"`
	ConnectionConfiguration map[string]interface{} `json:"connectionConfiguration"`
	Name                    string                 `json:"name"`
	SourceName              string                 `json:"sourceName"`
	Format                  string                 `json:"format"`
}

func DtoMapp(dto DbDtoRequest, SourceDefinitionID string) *DtoAirbyteCreateDB {

	return &DtoAirbyteCreateDB{
		SourceDefinitionID:      SourceDefinitionID,
		WorkspaceID:             dto.WorkspaceID,
		ConnectionConfiguration: dto.ConnectionConfiguration,
		Name:                    dto.Name,
		Format:                  dto.Format,
		SourceName:              dto.SourceName,
	}

}
