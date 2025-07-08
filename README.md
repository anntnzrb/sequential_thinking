# Sequential Thinking MCP Server

An MCP server implementation that provides a tool for dynamic and reflective problem-solving through a structured thinking process.

> **Note:** This is a Go rewrite of the original implementation in Javascript, maintaining the same functionality and purpose while offering improved performance.

## Features

- Break down complex problems into manageable steps
- Revise and refine thoughts as understanding deepens
- Branch into alternative paths of reasoning
- Adjust the total number of thoughts dynamically
- Generate and verify solution hypotheses

## Tool

### sequential_thinking

Facilitates a detailed, step-by-step thinking process for problem-solving and analysis.

**Inputs:**
- `thought` (string, **required**): Your current thinking step, which can include regular analytical steps, revisions of previous thoughts, questions about previous decisions, realizations about needing more analysis, changes in approach, hypothesis generation, or hypothesis verification
- `thoughtNumber` (integer, **required**): Current number in sequence (can go beyond initial total if needed)
- `totalThoughts` (integer, **required**): Current estimate of thoughts needed (can be adjusted up/down)
- `nextThoughtNeeded` (boolean, **required**): True if you need more thinking, even if at what seemed like the end
- `isRevision` (boolean, optional): A boolean indicating if this thought revises previous thinking
- `revisesThought` (integer, optional): If isRevision is true, which thought number is being reconsidered
- `branchFromThought` (integer, optional): If branching, which thought number is the branching point
- `branchId` (string, optional): Identifier for the current branch (if any)
- `needsMoreThoughts` (boolean, optional): If reaching end but realizing more thoughts needed

## Usage

The Sequential Thinking tool is designed for:
- Breaking down complex problems into steps
- Planning and design with room for revision
- Analysis that might need course correction
- Problems where the full scope might not be clear initially
- Tasks that need to maintain context over multiple steps
- Situations where irrelevant information needs to be filtered out

## Configuration

#### Docker (Recommended)

```json
{
  "mcpServers": {
    "sequential-thinking": {
      "command": "docker",
      "args": [
        "run",
        "--rm",
        "-i",
        "--init",
        "ghcr.io/anntnzrb/sequential_thinking:latest"
      ]
    }
  }
}
```

#### Binary (if you have Go installed)

```json
{
  "mcpServers": {
    "sequential-thinking": {
      "command": "go",
      "args": [
        "run",
        "github.com/anntnzrb/sequential_thinking@latest"
      ]
    }
  }
}
```

To disable logging of thought information set env var: `DISABLE_THOUGHT_LOGGING` to `true`.

## License

This MCP server is licensed under the MIT License. This means you are free to use, modify, and distribute the software, subject to the terms and conditions of the MIT License. For more details, please see the LICENSE file in the project repository.
