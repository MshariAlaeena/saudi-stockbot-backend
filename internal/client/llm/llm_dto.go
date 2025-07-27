package llm

import (
	"encoding/json"
	"patient-chatbot/internal/client/stock"
	"patient-chatbot/internal/dto"
)

type ChatMessageBlock struct {
	Role      dto.Role         `json:"role"`
	Content   string           `json:"content"`
	ToolCalls []ToolCallsBlock `json:"tool_calls,omitempty"`
}

type JsonSchemaProperty struct {
	Type  string   `json:"type"`
	Enum  []string `json:"enum,omitempty"`
	Items []string `json:"items,omitempty"`
}

type ParametersRequest struct {
	Type                 string                 `json:"type"`
	Required             []string               `json:"required"`
	AdditionalProperties bool                   `json:"additionalProperties"`
	Properties           map[string]interface{} `json:"properties"`
}

type ToolCallFunctionRequest struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Parameters  ParametersRequest `json:"parameters"`
}

type ToolCallRequest struct {
	Type     string                  `json:"type"`
	Function ToolCallFunctionRequest `json:"function"`
}

type ChatRequest struct {
	Model               string             `json:"model"`
	Messages            []ChatMessageBlock `json:"messages"`
	Temperature         float32            `json:"temperature"`
	MaxCompletionTokens int                `json:"max_completion_tokens"`
	TopP                float32            `json:"top_p"`
	Stream              bool               `json:"stream"`
	Stop                interface{}        `json:"stop"`
	Tools               []ToolCallRequest  `json:"tools"`
	ToolChoice          string             `json:"tool_choice"`
}

type ImageBlock struct {
	URL string `json:"url"`
}

type ExtractTextContentBlock struct {
	Type     string      `json:"type"`
	Text     string      `json:"text,omitempty"`
	ImageURL *ImageBlock `json:"image_url,omitempty"`
}

type ExtractTextMessageBlock struct {
	Role    string                    `json:"role"`
	Content []ExtractTextContentBlock `json:"content"`
}

type ExtractTextRequest struct {
	Model               string                    `json:"model"`
	Messages            []ExtractTextMessageBlock `json:"messages"`
	Temperature         float32                   `json:"temperature"`
	MaxCompletionTokens int                       `json:"max_completion_tokens"`
	TopP                float32                   `json:"top_p"`
	Stream              bool                      `json:"stream"`
	Stop                interface{}               `json:"stop"`
}

type ExtractTextResponse struct {
	Title    string   `json:"title"`
	Category string   `json:"category"`
	Chunks   []string `json:"chunks"`
}

type ToolCallFunction struct {
	Name      stock.Function  `json:"name"`
	Arguments json.RawMessage `json:"arguments"`
}

type ToolCallsBlock struct {
	ID       string           `json:"id"`
	Type     string           `json:"type"`
	Function ToolCallFunction `json:"function"`
}

type ChatChoice struct {
	Message ChatMessageBlock `json:"message"`
}

type ChatResponse struct {
	Choices []ChatChoice `json:"choices"`
}

type QuittingCoachResponse struct {
	DaysSmokeFree          int  `json:"daysSmokeFree"`
	MoneySaved             int  `json:"moneySaved"`
	MentionedDaysSmokeFree bool `json:"mentionedDaysSmokeFree"`
	MentionedMoneySaved    bool `json:"mentionedMoneySaved"`
}
