// Package mcp provides an Adapter that parses MCP tools/list responses.
package mcp

// ListToolsResponse is the top-level MCP tools/list wire format.
type ListToolsResponse struct {
	Tools []Tool `json:"tools"`
}

// Tool is a single tool entry in the MCP tools/list response.
type Tool struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	InputSchema InputSchema `json:"inputSchema"`
}

// InputSchema is the JSON Schema fragment embedded in an MCP Tool.
type InputSchema struct {
	Type        string                    `json:"type,omitempty"`
	Properties  map[string]SchemaProperty `json:"properties,omitempty"`
	Required    []string                  `json:"required,omitempty"`
	Description string                    `json:"description,omitempty"`
}

// SchemaProperty describes a single property within an InputSchema.
type SchemaProperty struct {
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}
