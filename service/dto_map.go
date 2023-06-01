package service

import "github.com/ThailanTec/factoriesStrategy/model"

func DtoMapp(dto model.DbDtoRequest, SourceDefinitionID string) *model.DtoAirbyteCreateDB {

	return &model.DtoAirbyteCreateDB{
		SourceDefinitionID:      SourceDefinitionID,
		WorkspaceID:             dto.WorkspaceID,
		ConnectionConfiguration: dto.ConnectionConfiguration,
		Name:                    dto.Name,
		Format:                  dto.Format,
		SourceName:              dto.SourceName,
	}

}
