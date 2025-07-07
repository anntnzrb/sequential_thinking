# Foxy-Contexts Common Patterns and Best Practices

## Common Use Cases
1. **File System Operations**: Reading directories, file manipulation
2. **Database Queries**: SQL execution and result formatting
3. **API Integrations**: External service calls and data retrieval
4. **System Information**: Gathering system metrics and status
5. **Data Processing**: Transform and format data for AI consumption

## Database Integration Pattern
```go
type DatabaseTool struct {
    db *sql.DB
}

func NewDatabaseTool(db *sql.DB) fxctx.Tool {
    return fxctx.NewTool(
        &mcp.Tool{
            Name: "query-database",
            Description: utils.Ptr("Execute database queries"),
            InputSchema: mcp.ToolInputSchema{
                Type: "object",
                Properties: map[string]map[string]interface{}{
                    "query": {
                        "type": "string",
                        "description": "SQL query to execute",
                    },
                },
                Required: []string{"query"},
            },
        },
        func(ctx context.Context, args map[string]interface{}) *mcp.CallToolResult {
            // Implementation
        },
    )
}
```

## Error Handling Best Practices
- Always wrap operations in error handling
- Use `mcp.CallToolResult` with `IsError: utils.Ptr(true)` for errors
- Log debug info to stderr, not stdout (to avoid polluting JSON-RPC)
- Implement panic recovery in tool callbacks

## Testing with foxytest
- Use `foxytest.Read("testdata")` to load test cases
- Support YAML-driven test scenarios
- Include regex matching for flexible response validation