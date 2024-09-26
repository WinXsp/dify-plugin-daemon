package model_entities

import "github.com/langgenius/dify-plugin-daemon/internal/types/entities/plugin_entities"

type GetModelSchemasResponse struct {
	ModelSchema *plugin_entities.ModelDeclaration `json:"model_schema" validate:"omitempty"`
}
