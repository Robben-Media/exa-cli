package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/builtbyrobben/exa-cli/internal/exa"
	"github.com/builtbyrobben/exa-cli/internal/outfmt"
)

type ContentsCmd struct {
	Query      string   `arg:"" required:"" help:"Search query"`
	Num        int      `help:"Number of results" default:"3"`
	MaxChars   int      `help:"Maximum characters of text per result" name:"max-chars" default:"3000"`
	Highlights bool     `help:"Include text highlights"`
	Summary    bool     `help:"Include AI summary of each result"`
	Domains    []string `help:"Include only these domains (comma-separated)" sep:","`
}

func (cmd *ContentsCmd) Run(ctx context.Context) error {
	client, err := getExaClient()
	if err != nil {
		return err
	}

	contents := &exa.ContentOptions{
		Text: &exa.TextOptions{
			MaxCharacters: cmd.MaxChars,
		},
		Highlights: cmd.Highlights,
		Summary:    cmd.Summary,
	}

	req := exa.SearchRequest{
		Query:          cmd.Query,
		NumResults:     cmd.Num,
		IncludeDomains: cmd.Domains,
		Contents:       contents,
	}

	result, err := client.SearchAndContents(ctx, req)
	if err != nil {
		return err
	}

	if outfmt.IsJSON(ctx) {
		return outfmt.WriteJSON(os.Stdout, result)
	}

	if len(result.Results) == 0 {
		fmt.Fprintln(os.Stderr, "No results found")
		return nil
	}

	for i, r := range result.Results {
		fmt.Fprintf(os.Stdout, "=== Result %d ===\n", i+1)
		fmt.Fprintf(os.Stdout, "Title: %s\n", r.Title)
		fmt.Fprintf(os.Stdout, "URL:   %s\n", r.URL)

		if r.Score > 0 {
			fmt.Fprintf(os.Stdout, "Score: %.2f\n", r.Score)
		}

		if r.PublishedDate != "" {
			fmt.Fprintf(os.Stdout, "Date:  %s\n", r.PublishedDate)
		}

		if r.Summary != "" {
			fmt.Fprintf(os.Stdout, "\nSummary: %s\n", r.Summary)
		}

		if len(r.Highlights) > 0 {
			fmt.Fprintf(os.Stdout, "\nHighlights:\n  %s\n", strings.Join(r.Highlights, "\n  "))
		}

		if r.Text != "" {
			fmt.Fprintf(os.Stdout, "\nText:\n%s\n", r.Text)
		}

		fmt.Fprintln(os.Stdout, "---")
	}

	return nil
}
