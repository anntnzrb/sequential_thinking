// Package main implements a sequential thinking MCP server using foxy-contexts.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/strowk/foxy-contexts/pkg/app"
	"github.com/strowk/foxy-contexts/pkg/fxctx"
	"github.com/strowk/foxy-contexts/pkg/mcp"
	"github.com/strowk/foxy-contexts/pkg/stdio"
)

func ptr[T any](v T) *T {
	return &v
}

// ThoughtData represents the input parameters for sequential thinking operations.
type ThoughtData struct {
	Thought           string `json:"thought"`
	ThoughtNumber     int    `json:"thoughtNumber"`
	TotalThoughts     int    `json:"totalThoughts"`
	IsRevision        *bool  `json:"isRevision,omitempty"`
	RevisesThought    *int   `json:"revisesThought,omitempty"`
	BranchFromThought *int   `json:"branchFromThought,omitempty"`
	BranchID          string `json:"branchId,omitempty"`
	NeedsMoreThoughts *bool  `json:"needsMoreThoughts,omitempty"`
	NextThoughtNeeded bool   `json:"nextThoughtNeeded"`
}

func validateThoughtData(args map[string]any) (*ThoughtData, error) {
	thought, ok := args["thought"].(string)
	if !ok || thought == "" {
		return nil, fmt.Errorf("thought is required and must be a non-empty string")
	}

	thoughtNumberFloat, ok := args["thoughtNumber"].(float64)
	if !ok {
		return nil, fmt.Errorf("thoughtNumber is required and must be a number")
	}
	thoughtNumber := int(thoughtNumberFloat)

	totalThoughtsFloat, ok := args["totalThoughts"].(float64)
	if !ok {
		return nil, fmt.Errorf("totalThoughts is required and must be a number")
	}
	totalThoughts := int(totalThoughtsFloat)

	nextThoughtNeeded, ok := args["nextThoughtNeeded"].(bool)
	if !ok {
		return nil, fmt.Errorf("nextThoughtNeeded is required and must be a boolean")
	}

	if thoughtNumber < 1 {
		return nil, fmt.Errorf("thoughtNumber must be >= 1")
	}
	if totalThoughts < 1 {
		return nil, fmt.Errorf("totalThoughts must be >= 1")
	}
	if thoughtNumber > totalThoughts {
		return nil, fmt.Errorf("thoughtNumber cannot be greater than totalThoughts")
	}

	data := &ThoughtData{
		Thought:           thought,
		ThoughtNumber:     thoughtNumber,
		TotalThoughts:     totalThoughts,
		NextThoughtNeeded: nextThoughtNeeded,
	}

	if isRevision, ok := args["isRevision"].(bool); ok {
		data.IsRevision = &isRevision
	}

	if revisesThoughtFloat, ok := args["revisesThought"].(float64); ok {
		revisesThought := int(revisesThoughtFloat)
		data.RevisesThought = &revisesThought
	}

	if branchFromThoughtFloat, ok := args["branchFromThought"].(float64); ok {
		branchFromThought := int(branchFromThoughtFloat)
		data.BranchFromThought = &branchFromThought
	}

	if branchID, ok := args["branchId"].(string); ok {
		data.BranchID = branchID
	}

	if needsMoreThoughts, ok := args["needsMoreThoughts"].(bool); ok {
		data.NeedsMoreThoughts = &needsMoreThoughts
	}

	return data, nil
}

func formatThought(data *ThoughtData) string {
	if os.Getenv("DISABLE_THOUGHT_LOGGING") == "true" {
		return "Thought logging is disabled."
	}

	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	var b strings.Builder
	thoughtNum := strconv.Itoa(data.ThoughtNumber)
	totalNum := strconv.Itoa(data.TotalThoughts)

	fmt.Fprintf(&b, "%s %s (%s/%s)\n",
		cyan("💭 Thought"), yellow(thoughtNum), yellow(thoughtNum), yellow(totalNum))

	if data.IsRevision != nil && *data.IsRevision && data.RevisesThought != nil {
		fmt.Fprintf(&b, "%s %s\n",
			red("🔄 Revising thought"), yellow(strconv.Itoa(*data.RevisesThought)))
	}

	if data.BranchFromThought != nil {
		fmt.Fprintf(&b, "%s %s",
			blue("🌿 Branching from thought"), yellow(strconv.Itoa(*data.BranchFromThought)))
		if data.BranchID != "" {
			fmt.Fprintf(&b, " (%s)", blue(data.BranchID))
		}
		b.WriteString("\n")
	}

	fmt.Fprintf(&b, "\n%s\n", data.Thought)

	if data.NextThoughtNeeded {
		fmt.Fprintf(&b, "\n%s\n", green("→ More thinking needed"))
	} else {
		fmt.Fprintf(&b, "\n%s\n", green("✓ Thinking complete"))
	}

	fmt.Fprintf(&b, "\n%s\n", cyan(fmt.Sprintf(
		"Status: Thought %d/%d | Next needed: %v",
		data.ThoughtNumber, data.TotalThoughts, data.NextThoughtNeeded)))

	return b.String()
}

// NewSequentialThinkingTool creates and returns a new sequential thinking MCP tool.
func NewSequentialThinkingTool() fxctx.Tool {
	return fxctx.NewTool(
		&mcp.Tool{
			Name:        "sequential_thinking",
			Description: ptr("A detailed tool for dynamic and reflective problem-solving through thoughts. This tool helps analyze problems through a flexible thinking process that can adapt and evolve. Each thought can build on, question, or revise previous insights as understanding deepens."),
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]map[string]any{
					"thought": {
						"type":        "string",
						"description": "Your current thinking step, which can include regular analytical steps, revisions of previous thoughts, questions about previous decisions, realizations about needing more analysis, changes in approach, hypothesis generation, or hypothesis verification.",
					},
					"nextThoughtNeeded": {
						"type":        "boolean",
						"description": "True if you need more thinking, even if at what seemed like the end",
					},
					"thoughtNumber": {
						"type":        "integer",
						"minimum":     1,
						"description": "Current number in sequence (can go beyond initial total if needed)",
					},
					"totalThoughts": {
						"type":        "integer",
						"minimum":     1,
						"description": "Current estimate of thoughts needed (can be adjusted up/down)",
					},
					"isRevision": {
						"type":        "boolean",
						"description": "A boolean indicating if this thought revises previous thinking",
					},
					"revisesThought": {
						"type":        "integer",
						"minimum":     1,
						"description": "If isRevision is true, which thought number is being reconsidered",
					},
					"branchFromThought": {
						"type":        "integer",
						"minimum":     1,
						"description": "If branching, which thought number is the branching point",
					},
					"branchId": {
						"type":        "string",
						"description": "Identifier for the current branch (if any)",
					},
					"needsMoreThoughts": {
						"type":        "boolean",
						"description": "If reaching end but realizing more thoughts needed",
					},
				},
				Required: []string{"thought", "nextThoughtNeeded", "thoughtNumber", "totalThoughts"},
			},
		},
		func(args map[string]any) *mcp.CallToolResult {
			data, err := validateThoughtData(args)
			if err != nil {
				return &mcp.CallToolResult{
					IsError: ptr(true),
					Content: []any{
						mcp.TextContent{
							Type: "text",
							Text: fmt.Sprintf("Validation error: %v", err),
						},
					},
				}
			}

			return &mcp.CallToolResult{
				Content: []any{
					mcp.TextContent{
						Type: "text",
						Text: formatThought(data),
					},
				},
				IsError: ptr(false),
				Meta: map[string]any{
					"thoughtNumber":        data.ThoughtNumber,
					"totalThoughts":        data.TotalThoughts,
					"nextThoughtNeeded":    data.NextThoughtNeeded,
					"branches":             []any{},
					"thoughtHistoryLength": 1,
				},
			}
		},
	)
}

func main() {
	if err := app.NewBuilder().
		WithName("sequential-thinking").
		WithVersion("1.0.0").
		WithTool(NewSequentialThinkingTool).
		WithTransport(stdio.NewTransport()).
		Run(); err != nil {
		os.Exit(1)
	}
}
