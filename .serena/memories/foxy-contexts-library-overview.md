# Foxy-Contexts Library Overview

## What is Foxy-Contexts?
**Foxy-contexts** is a Go framework for building Model Context Protocol (MCP) servers. It provides a high-level, declarative approach to creating AI-integrated applications that communicate with AI assistants like Claude Desktop.

## Key Features
- **Builder Pattern**: Fluent API for server configuration
- **Dependency Injection**: Built on Uber's FX framework
- **Transport Abstraction**: Supports STDIO and HTTP transports
- **MCP Compliance**: Full adherence to MCP protocol specifications
- **Testing Framework**: Includes `foxytest` package for integration testing

## Core Architecture
```
github.com/strowk/foxy-contexts/pkg/
├── app/          # Application builder and configuration
├── fxctx/        # FX context integration and tool definitions
├── mcp/          # MCP protocol types and structures
├── stdio/        # Standard I/O transport implementation
├── foxytest/     # Testing utilities
└── foxyevent/    # Event handling and logging
```

## Repository Location
- **GitHub**: https://github.com/strowk/foxy-contexts
- **Primary Language**: Go
- **Dependencies**: Uber FX for dependency injection