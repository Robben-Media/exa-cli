package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/builtbyrobben/exa-cli/internal/exa"
	"github.com/builtbyrobben/exa-cli/internal/outfmt"
)

type SearchCmd struct {
	Query          string   `arg:"" required:"" help:"Search query"`
	Num            int      `help:"Number of results" default:"5"`
	Type           string   `help:"Search type: auto, neural, keyword" default:"auto"`
	Domains        []string `help:"Include only these domains (comma-separated)" sep:","`
	ExcludeDomains []string `help:"Exclude these domains (comma-separated)" name:"exclude-domains" sep:","`
	StartDate      string   `help:"Start published date (YYYY-MM-DD)" name:"start-date"`
	EndDate        string   `help:"End published date (YYYY-MM-DD)" name:"end-date"`
}

func (cmd *SearchCmd) Run(ctx context.Context) error {
	client, err := getExaClient()
	if err != nil {
		return err
	}

	req := exa.SearchRequest{
		Query:          cmd.Query,
		NumResults:     cmd.Num,
		Type:           cmd.Type,
		IncludeDomains: cmd.Domains,
		ExcludeDomains: cmd.ExcludeDomains,
		StartPublished: cmd.StartDate,
		EndPublished:   cmd.EndDate,
	}

	result, err := client.Search(ctx, req)
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

	if result.AutopromptText != "" {
		fmt.Fprintf(os.Stderr, "Autoprompt: %s\n\n", result.AutopromptText)
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

		if r.Author != "" {
			fmt.Fprintf(os.Stdout, "Author: %s\n", r.Author)
		}

		if r.Text != "" {
			text := r.Text
			if len(text) > 200 {
				text = text[:200] + "..."
			}

			fmt.Fprintf(os.Stdout, "\nText:  %s\n", text)
		}

		if len(r.Highlights) > 0 {
			fmt.Fprintf(os.Stdout, "\nHighlights:\n  %s\n", strings.Join(r.Highlights, "\n  "))
		}

		if r.Summary != "" {
			fmt.Fprintf(os.Stdout, "\nSummary: %s\n", r.Summary)
		}

		fmt.Fprintln(os.Stdout, "---")
	}

	return nil
}
