# Model Context Protocol (MCP) Fundamentals

## What is MCP?
Model Context Protocol is an open standard introduced by Anthropic in November 2024 that enables seamless integration between LLM applications and external data sources/tools. Think of it as "USB-C for AI applications."

## Key Components
- **MCP Hosts**: Programs like Claude Desktop, IDEs, or AI tools
- **MCP Clients**: Protocol clients that maintain 1:1 connections with servers
- **MCP Servers**: Lightweight programs that expose capabilities through standardized protocol

## Core Primitives
- **Tools**: Functions that LLMs can call to perform actions (e.g., API calls, file operations)
- **Resources**: Data sources that LLMs can access (similar to GET endpoints)
- **Prompts**: Pre-defined templates to use tools/resources optimally

## Technical Foundation
- **Protocol**: JSON-RPC 2.0 for message exchange
- **Transports**: STDIO (local) and HTTP+SSE (remote)
- **Negotiation**: Capability-based negotiation system
- **Schema**: JSON Schema for input/output validation

## Integration with Claude Desktop
Add to Claude Desktop configuration:
```json
{
  "mcpServers": {
    "my-server": {
      "command": "/path/to/executable",
      "args": []
    }
  }
}
```