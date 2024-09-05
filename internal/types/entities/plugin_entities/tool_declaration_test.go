package plugin_entities

import (
	"strings"
	"testing"
)

func TestFullFunctionToolProvider_Validate(t *testing.T) {
	const data = `
{
	"identity": {
		"author": "author",
		"name": "name",
		"description": {
			"en_US": "description",
			"zh_Hans": "描述",
			"pt_BR": "descrição"
		},
		"icon": "icon",
		"label": {
			"en_US": "label",
			"zh_Hans": "标签",
			"pt_BR": "etiqueta"
		},
		"tags": [
			"image",
			"videos"
		]
	},
	"credentials_schema": {
		"api_key": {
			"name": "API Key",
			"type": "secret-input",
			"required": false,
			"default": "default",
			"label": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			},
			"helper": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			},
			"url": "https://example.com",
			"placeholder": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			}
		}
	},
	"tools": [
		{
			"identity": {
				"author": "author",
				"name": "tool",
				"label": {
					"en_US": "label",
					"zh_Hans": "标签",
					"pt_BR": "etiqueta"
				}
			},
			"description": {
				"human": {
					"en_US": "description",
					"zh_Hans": "描述",
					"pt_BR": "descrição"
				},
				"llm": "description"
			},
			"parameters": [
				{
					"name": "parameter",
					"type": "string",
					"label": {
						"en_US": "label",
						"zh_Hans": "标签",
						"pt_BR": "etiqueta"
					},
					"human_description": {
						"en_US": "description",
						"zh_Hans": "描述",
						"pt_BR": "descrição"
					},
					"form": "llm",
					"required": true,
					"default": "default",
					"options": [
						{
							"value": "value",
							"label": {
								"en_US": "label",
								"zh_Hans": "标签",
								"pt_BR": "etiqueta"
							}
						}
					]
				}
			]
		}
	]
}
	`

	_, err := UnmarshalToolProviderDeclaration([]byte(data))
	if err != nil {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v", err)
		return
	}
}

func TestWithoutAuthorToolProvider_Validate(t *testing.T) {
	const data = `
{
	"identity": {
		"name": "name",
		"description": {
			"en_US": "description",
			"zh_Hans": "描述",
			"pt_BR": "descrição"
		},
		"icon": "icon",
		"label": {
			"en_US": "label",
			"zh_Hans": "标签",
			"pt_BR": "etiqueta"
		},
		"tags": [
			"image",
			"videos"
		]
	},
	"credentials_schema": {
		"api_key": {
			"name": "API Key",
			"type": "secret-input",
			"required": false,
			"default": "default",
			"label": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			},
			"helper": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			},
			"url": "https://example.com",
			"placeholder": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			}
		}
	},
	"tools": [
	
	]
}
	`

	_, err := UnmarshalToolProviderDeclaration([]byte(data))
	if err == nil {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v, wantErr %v", err, true)
		return
	}
}

func TestWithoutNameToolProvider_Validate(t *testing.T) {
	const data = `
{
	"identity": {
		"author": "author",
		"description": {
			"en_US": "description",
			"zh_Hans": "描述",
			"pt_BR": "descrição"
		},
		"icon": "icon",
		"label": {
			"en_US": "label",
			"zh_Hans": "标签",
			"pt_BR": "etiqueta"
		},
		"tags": [
			"image",
			"videos"
		]
	},
	"credentials_schema": {
		"api_key": {
			"name": "API Key",
			"type": "secret-input",
			"required": false,
			"default": "default",
			"label": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			},
			"helper": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			},
			"url": "https://example.com",
			"placeholder": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			}
		}
	},
	"tools": [
	
	]
}
	`

	_, err := UnmarshalToolProviderDeclaration([]byte(data))
	if err == nil {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v, wantErr %v", err, true)
		return
	}
}

func TestWithoutDescriptionToolProvider_Validate(t *testing.T) {
	const data = `
{
	"identity": {
		"author": "author",
		"name": "name",
		"icon": "icon",
		"label": {
			"en_US": "label",
			"zh_Hans": "标签",
			"pt_BR": "etiqueta"
		},
		"tags": [
			"image",
			"videos"
		]
	},
	"credentials_schema": {
		"api_key": {
			"name": "API Key",
			"type": "secret-input",
			"required": false,
			"default": "default",
			"label": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			},
			"helper": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			},
			"url": "https://example.com",
			"placeholder": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			}
		}
	},
	"tools": [
	
	]
}
	`

	_, err := UnmarshalToolProviderDeclaration([]byte(data))
	if err == nil {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v, wantErr %v", err, true)
		return
	}
}

func TestWrongCredentialTypeToolProvider_Validate(t *testing.T) {
	const data = `
{
	"identity": {
		"author": "author",
		"name": "name",
		"description": {
			"en_US": "description",
			"zh_Hans": "描述",
			"pt_BR": "descrição"
		},
		"icon": "icon",
		"label": {
			"en_US": "label",
			"zh_Hans": "标签",
			"pt_BR": "etiqueta"
		},
		"tags": [
			"image",
			"videos"
		]
	},
	"credentials_schema": {
		"api_key": {
			"name": "API Key",
			"type": "wrong",
			"required": false,
			"default": "default",
			"label": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			},
			"helper": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			},
			"url": "https://example.com",
			"placeholder": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			}
		}
	},
	"tools": [
	
	]
}
	`

	_, err := UnmarshalToolProviderDeclaration([]byte(data))
	if err == nil {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v, wantErr %v", err, true)
		return
	}
}

func TestWrongIdentityTagsToolProvider_Validate(t *testing.T) {
	const data = `
{
	"identity": {
		"author": "author",
		"name": "name",
		"description": {
			"en_US": "description",
			"zh_Hans": "描述",
			"pt_BR": "descrição"
		},
		"icon": "icon",
		"label": {
			"en_US": "label",
			"zh_Hans": "标签",
			"pt_BR": "etiqueta"
		},
		"tags": [
			"wrong",
			"videos"
		]
	},
	"credentials_schema": {
		"api_key": {
			"name": "API Key",
			"type": "secret-input",
			"required": false,
			"default": "default",
			"label": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			},
			"helper": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			},
			"url": "https://example.com",
			"placeholder": {
				"en_US": "API Key",
				"zh_Hans": "API 密钥",
				"pt_BR": "Chave da API"
			}
		}
	},
	"tools": [
	
	]
}
	`

	_, err := UnmarshalToolProviderDeclaration([]byte(data))
	if err == nil {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v, wantErr %v", err, true)
		return
	}
}

func TestWrongToolParameterTypeToolProvider_Validate(t *testing.T) {
	const data = `
{
	"identity": {
		"author": "author",
		"name": "name",
		"description": {
			"en_US": "description",
			"zh_Hans": "描述",
			"pt_BR": "descrição"
		},
		"icon": "icon",
		"label": {
			"en_US": "label",
			"zh_Hans": "标签",
			"pt_BR": "etiqueta"
		},
		"tags": []
	},
	"credentials_schema": {},
	"tools": [
		{
			"identity": {
				"author": "author",
				"name": "tool",
				"label": {
					"en_US": "label",
					"zh_Hans": "标签",
					"pt_BR": "etiqueta"
				}
			},
			"description": {
				"human": {
					"en_US": "description",
					"zh_Hans": "描述",
					"pt_BR": "descrição"
				},
				"llm": "description"
			},
			"parameters": [
				{
					"name": "parameter",
					"type": "wrong",
					"label": {
						"en_US": "label",
						"zh_Hans": "标签",
						"pt_BR": "etiqueta"
					},
					"human_description": {
						"en_US": "description",
						"zh_Hans": "描述",
						"pt_BR": "descrição"
					},
					"form": "llm",
					"required": true,
					"default": "default",
					"options": []
				}
			]
		}
	]
}
	`

	_, err := UnmarshalToolProviderDeclaration([]byte(data))
	if err == nil {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v, wantErr %v", err, true)
		return
	}
}

func TestWrongToolParameterFormToolProvider_Validate(t *testing.T) {
	const data = `
{
	"identity": {
		"author": "author",
		"name": "name",
		"description": {
			"en_US": "description",
			"zh_Hans": "描述",
			"pt_BR": "descrição"
		},
		"icon": "icon",
		"label": {
			"en_US": "label",
			"zh_Hans": "标签",
			"pt_BR": "etiqueta"
		},
		"tags": []
	},
	"credentials_schema": {},
	"tools": [
		{
			"identity": {
				"author": "author",
				"name": "tool",
				"label": {
					"en_US": "label",
					"zh_Hans": "标签",
					"pt_BR": "etiqueta"
				}
			},
			"description": {
				"human": {
					"en_US": "description",
					"zh_Hans": "描述",
					"pt_BR": "descrição"
				},
				"llm": "description"
			},
			"parameters": [
				{
					"name": "parameter",
					"type": "string",
					"label": {
						"en_US": "label",
						"zh_Hans": "标签",
						"pt_BR": "etiqueta"
					},
					"human_description": {
						"en_US": "description",
						"zh_Hans": "描述",
						"pt_BR": "descrição"
					},
					"form": "wrong",
					"required": true,
					"default": "default",
					"options": []
				}
			]
		}
	]
}
	`

	_, err := UnmarshalToolProviderDeclaration([]byte(data))
	if err == nil {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v, wantErr %v", err, true)
		return
	}
}

func TestJSONSchemaTypeToolProvider_Validate(t *testing.T) {
	const data = `
{
	"identity": {
		"author": "author",
		"name": "name",
		"description": {
			"en_US": "description",
			"zh_Hans": "描述",
			"pt_BR": "descrição"
		},
		"icon": "icon",
		"label": {
			"en_US": "label",
			"zh_Hans": "标签",
			"pt_BR": "etiqueta"
		},
		"tags": []
	},
	"credentials_schema": {},
	"tools": [
		{
			"identity": {
				"author": "author",
				"name": "tool",
				"label": {
					"en_US": "label",
					"zh_Hans": "标签",
					"pt_BR": "etiqueta"
				}
			},
			"description": {
				"human": {
					"en_US": "description",
					"zh_Hans": "描述",
					"pt_BR": "descrição"
				},
				"llm": "description"
			},
			"output_schema": {
				"type": "object",
				"properties": {
					"name": {
						"type": "string"
					}
				}
			}
		}
	]
}
	`

	_, err := UnmarshalToolProviderDeclaration([]byte(data))
	if err != nil {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v, wantErr %v", err, true)
		return
	}
}

func TestWrongJSONSchemaToolProvider_Validate(t *testing.T) {
	const data = `
{
	"identity": {
		"author": "author",
		"name": "name",
		"description": {
			"en_US": "description",
			"zh_Hans": "描述",
			"pt_BR": "descrição"
		},
		"icon": "icon",
		"label": {
			"en_US": "label",
			"zh_Hans": "标签",
			"pt_BR": "etiqueta"
		},
		"tags": []
	},
	"credentials_schema": {},
	"tools": [
		{
			"identity": {
				"author": "author",
				"name": "tool",
				"label": {
					"en_US": "label",
					"zh_Hans": "标签",
					"pt_BR": "etiqueta"
				}
			},
			"description": {
				"human": {
					"en_US": "description",
					"zh_Hans": "描述",
					"pt_BR": "descrição"
				},
				"llm": "description"
			},
			"output_schema": {
				"type": "object",
				"properties": {
					"name": {
						"type": "aaa"
					}
				}
			}
		}
	]
}
	`

	_, err := UnmarshalToolProviderDeclaration([]byte(data))
	if err == nil {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v, wantErr %v", err, true)
		return
	}
}

func TestWrongAppSelectorScopeToolProvider_Validate(t *testing.T) {
	const data = `
{
	"identity": {
		"author": "author",
		"name": "name",
		"description": {
			"en_US": "description",
			"zh_Hans": "描述",
			"pt_BR": "descrição"
		},
		"icon": "icon",
		"label": {
			"en_US": "label",
			"zh_Hans": "标签",
			"pt_BR": "etiqueta"
		},
		"tags": []
	},
	"credentials_schema": {
		"api_key": {
			"name": "app-selector",
			"type": "app-selector",
			"scope": "wrong",
			"required": false,
			"default": null,
			"label": {
				"en_US": "app-selector",
				"zh_Hans": "app-selector",
				"pt_BR": "app-selector"
			},
			"helper": {
				"en_US": "app-selector",
				"zh_Hans": "app-selector",
				"pt_BR": "app-selector"
			},
			"url": "https://example.com",
			"placeholder": {
				"en_US": "app-selector",
				"zh_Hans": "app-selector",
				"pt_BR": "app-selector"
			}
		}
	},
	"tools": [
		{
			"identity": {
				"author": "author",
				"name": "tool",
				"label": {
					"en_US": "label",
					"zh_Hans": "标签",
					"pt_BR": "etiqueta"
				}
			},
			"description": {
				"human": {
					"en_US": "description",
					"zh_Hans": "描述",
					"pt_BR": "descrição"
				},
				"llm": "description"
			},
			"parameters": [
				{
					"name": "parameter-app-selector",
					"label": {
						"en_US": "label",
						"zh_Hans": "标签",
						"pt_BR": "etiqueta"
					},
					"human_description": {
						"en_US": "description",
						"zh_Hans": "描述",
						"pt_BR": "descrição"
					},
					"type": "app-selector",
					"form": "llm",
					"scope": "wrong",
					"required": true,
					"default": "default",
					"options": []
				}
			]
		}
	]
}
	`

	_, err := UnmarshalToolProviderDeclaration([]byte(data))
	if err == nil {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v, wantErr %v", err, true)
		return
	}

	str := err.Error()
	if !strings.Contains(str, "api_key") {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v, wantErr %v", err, true)
		return
	}

	if !strings.Contains(str, "ToolProviderDeclaration.Tools[0].Parameters[0].Scope") {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v, wantErr %v", err, true)
		return
	}
}

func TestAppSelectorScopeToolProvider_Validate(t *testing.T) {
	const data = `
{
	"identity": {
		"author": "author",
		"name": "name",
		"description": {
			"en_US": "description",
			"zh_Hans": "描述",
			"pt_BR": "descrição"
		},
		"icon": "icon",
		"label": {
			"en_US": "label",
			"zh_Hans": "标签",
			"pt_BR": "etiqueta"
		},
		"tags": []
	},
	"credentials_schema": {
		"app-selector": {
			"name": "app-selector",
			"type": "app-selector",
			"scope": "all",
			"required": false,
			"default": null,
			"label": {
				"en_US": "app-selector",
				"zh_Hans": "app-selector",
				"pt_BR": "app-selector"
			},
			"helper": {
				"en_US": "app-selector",
				"zh_Hans": "app-selector",
				"pt_BR": "app-selector"
			},
			"url": "https://example.com",
			"placeholder": {
				"en_US": "app-selector",
				"zh_Hans": "app-selector",
				"pt_BR": "app-selector"
			}
		}
	},
	"tools": [
	
	]
}
	`

	_, err := UnmarshalToolProviderDeclaration([]byte(data))
	if err != nil {
		t.Errorf("UnmarshalToolProviderConfiguration() error = %v, wantErr %v", err, true)
		return
	}
}