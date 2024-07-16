package plugin_entities

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

func parse_yaml_to_json(data []byte) ([]byte, error) {
	var obj interface{}
	err := yaml.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}

	json_data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return json_data, nil
}

const (
	template = `
provider: openai
label:
  en_US: OpenAI
description:
  en_US: Models provided by OpenAI, such as GPT-3.5-Turbo and GPT-4.
  zh_Hans: OpenAI 提供的模型，例如 GPT-3.5-Turbo 和 GPT-4。
icon_small:
  en_US: icon_s_en.svg
icon_large:
  en_US: icon_l_en.svg
background: "#E5E7EB"
help:
  title:
    en_US: Get your API Key from OpenAI
    zh_Hans: 从 OpenAI 获取 API Key
  url:
    en_US: https://platform.openai.com/account/api-keys
supported_model_types:
  - llm
  - text-embedding
  - speech2text
  - moderation
  - tts
configurate_methods:
  - predefined-model
  - customizable-model
model_credential_schema:
  model:
    label:
      en_US: Model Name
      zh_Hans: 模型名称
    placeholder:
      en_US: Enter your model name
      zh_Hans: 输入模型名称
  credential_form_schemas:
    - variable: openai_api_key
      label:
        en_US: API Key
      type: secret-input
      required: true
      placeholder:
        zh_Hans: 在此输入您的 API Key
        en_US: Enter your API Key
    - variable: openai_organization
      label:
        zh_Hans: 组织 ID
        en_US: Organization
      type: text-input
      required: false
      placeholder:
        zh_Hans: 在此输入您的组织 ID
        en_US: Enter your Organization ID
    - variable: openai_api_base
      label:
        zh_Hans: API Base
        en_US: API Base
      type: text-input
      required: false
      placeholder:
        zh_Hans: 在此输入您的 API Base
        en_US: Enter your API Base
provider_credential_schema:
  credential_form_schemas:
    - variable: openai_api_key
      label:
        en_US: API Key
      type: secret-input
      required: true
      placeholder:
        zh_Hans: 在此输入您的 API Key
        en_US: Enter your API Key
    - variable: openai_organization
      label:
        zh_Hans: 组织 ID
        en_US: Organization
      type: text-input
      required: false
      placeholder:
        zh_Hans: 在此输入您的组织 ID
        en_US: Enter your Organization ID
    - variable: openai_api_base
      label:
        zh_Hans: API Base
        en_US: API Base
      type: text-input
      required: false
      placeholder:
        zh_Hans: 在此输入您的 API Base, 如：https://api.openai.com
        en_US: Enter your API Base, e.g. https://api.openai.com
models:
  - model: gpt-3.5-turbo-16k-0613
    label:
      zh_Hans: gpt-3.5-turbo-16k-0613
      en_US: gpt-3.5-turbo-16k-0613
    model_type: llm
    features:
      - multi-tool-call
      - agent-thought
      - stream-tool-call
    model_properties:
      mode: chat
      context_size: 16385
    parameter_rules:
      - name: temperature
        use_template: temperature
      - name: top_p
        use_template: top_p
      - name: presence_penalty
        use_template: presence_penalty
      - name: frequency_penalty
        use_template: frequency_penalty
      - name: max_tokens
        use_template: max_tokens
        default: 512
        min: 1
        max: 16385
      - name: response_format
        use_template: response_format
    pricing:
      input: '0.003'
      output: '0.004'
      unit: '0.001'
      currency: USD
    `
)

func TestFullFunctionModelProvider_Validate(t *testing.T) {
	json_data, err := parse_yaml_to_json([]byte(template))
	if err != nil {
		t.Error(err)
	}

	_, err = UnmarshalModelProviderConfiguration(json_data)
	if err != nil {
		t.Errorf("UnmarshalModelProviderConfiguration() error = %v", err)
	}

}