# Foxy-Contexts Implementation Pattern

## Basic Server Structure
```go
func main() {
    app.NewBuilder().
        WithName("server-name").
        WithVersion("1.0.0").
        WithTool(NewMyTool()).
        WithTransport(stdio.NewStdioTransport()).
        Build().
        Run(context.Background())
}
```

## Tool Creation Pattern
```go
func NewMyTool() fxctx.Tool {
    return fxctx.NewTool(
        &mcp.Tool{
            Name:        "tool-name",
            Description: utils.Ptr("Tool description"),
            InputSchema: mcp.ToolInputSchema{
                Type:       "object",
                Properties: map[string]map[string]interface{}{},
                Required:   []string{},
            },
        },
        func(ctx context.Context, args map[string]interface{}) *mcp.CallToolResult {
            // Tool implementation
            return &mcp.CallToolResult{
                Content: []interface{}{
                    mcp.TextContent{
                        Type: "text",
                        Text: "response",
                    },
                },
                IsError: utils.Ptr(false),
            }
        },
    )
}
```

## Key APIs
- `app.NewBuilder()`: Creates application builder
- `fxctx.NewTool()`: Creates MCP tool definition
- `stdio.NewStdioTransport()`: Creates STDIO transport for Claude Desktop
- `mcp.CallToolResult`: Standard response structure