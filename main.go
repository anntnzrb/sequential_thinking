// Package main implements a sequential thinking MCP server using foxy-contexts.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	"github.com/go-viper/mapstructure/v2"
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
	Thought           string `json:"thought" mapstructure:"thought" validate:"required"`
	ThoughtNumber     int    `json:"thoughtNumber" mapstructure:"thoughtNumber" validate:"required,min=1"`
	TotalThoughts     int    `json:"totalThoughts" mapstructure:"totalThoughts" validate:"required,min=1"`
	IsRevision        *bool  `json:"isRevision,omitempty" mapstructure:"isRevision"`
	RevisesThought    *int   `json:"revisesThought,omitempty" mapstructure:"revisesThought"`
	BranchFromThought *int   `json:"branchFromThought,omitempty" mapstructure:"branchFromThought"`
	BranchID          string `json:"branchId,omitempty" mapstructure:"branchId"`
	NeedsMoreThoughts *bool  `json:"needsMoreThoughts,omitempty" mapstructure:"needsMoreThoughts"`
	NextThoughtNeeded bool   `json:"nextThoughtNeeded" mapstructure:"nextThoughtNeeded" validate:"required"`
}

func validateThoughtData(args map[string]any) (*ThoughtData, error) {
	var data ThoughtData

	// Decode using mapstructure
	if err := mapstructure.Decode(args, &data); err != nil {
		return nil, fmt.Errorf("failed to decode input: %v", err)
	}

	// Validate using validator
	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}

	// Custom business logic validation
	if data.ThoughtNumber > data.TotalThoughts {
		return nil, fmt.Errorf("thoughtNumber cannot be greater than totalThoughts")
	}

	return &data, nil
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
