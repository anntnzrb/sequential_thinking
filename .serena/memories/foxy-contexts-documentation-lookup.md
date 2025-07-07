# Foxy-Contexts Documentation Lookup

## When You Need More Information

If something isn't working or you need more context about the foxy-contexts library, use the context7 MCP tools:

### Step 1: Resolve Library ID
```
Use: mcp__MCP_DOCKER__resolve-library-id
Parameter: libraryName = "foxy-contexts" or "strowk/foxy-contexts"
```

### Step 2: Get Library Documentation
```
Use: mcp__MCP_DOCKER__get-library-docs
Parameter: context7CompatibleLibraryID = (from step 1)
Optional: topic = "specific topic like 'tools', 'transport', 'testing'"
Optional: tokens = higher number for more detailed docs
```

## When to Use This
- Encountering errors not covered in existing memories
- Need specific implementation details
- Want to understand advanced features
- Looking for updated API changes
- Debugging complex integration issues
- Need examples for specific use cases

## Example Usage
```
1. resolve-library-id with "foxy-contexts"
2. get-library-docs with the returned library ID
3. Focus on specific topics like "mcp integration" or "testing"
```

This ensures you always have access to the most current and comprehensive documentation for the foxy-contexts library.