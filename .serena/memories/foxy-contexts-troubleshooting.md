# Foxy-Contexts Troubleshooting Guide

## Common Issues and Solutions

### 1. STDIO Transport Issues
**Problem**: Messages not appearing in output
**Solution**: Ensure debug output goes to stderr, not stdout
```go
// Wrong - pollutes stdout
fmt.Println("Debug message")

// Correct - uses stderr
fmt.Fprintln(os.Stderr, "Debug message")
```

### 2. JSON-RPC Compliance
**Problem**: Invalid JSON-RPC messages
**Solution**: Use SDK serialization methods and proper MCP types

### 3. Dependency Injection Issues
**Problem**: Circular dependencies in FX
**Solution**: Use interfaces and careful module organization

### 4. Testing Challenges
**Solution**: Use the foxytest package for integration testing

## Development Tools
- **MCP Inspector**: `npx @modelcontextprotocol/inspector go run main.go`
- **Debugging**: Use structured logging with zap
- **Testing**: foxytest package for comprehensive testing

## Performance Considerations
- Use connection pooling for database operations
- Implement caching for frequently accessed data
- Set appropriate timeouts for external calls
- Manage goroutine lifecycles carefully

## Configuration Best Practices
- Use environment variables for configuration
- Support multiple config formats (JSON, YAML, TOML)
- Never hardcode secrets in source code
- Choose appropriate transport based on use case