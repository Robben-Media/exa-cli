package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/builtbyrobben/exa-cli/internal/exa"
	"github.com/builtbyrobben/exa-cli/internal/outfmt"
)

type FindSimilarCmd struct {
	URL            string   `arg:"" required:"" help:"URL to find similar pages for"`
	Num            int      `help:"Number of results" default:"5"`
	Domains        []string `help:"Include only these domains (comma-separated)" sep:","`
	ExcludeDomains []string `help:"Exclude these domains (comma-separated)" name:"exclude-domains" sep:","`
}

func (cmd *FindSimilarCmd) Run(ctx context.Context) error {
	client, err := getExaClient()
	if err != nil {
		return err
	}

	req := exa.FindSimilarRequest{
		URL:            cmd.URL,
		NumResults:     cmd.Num,
		IncludeDomains: cmd.Domains,
		ExcludeDomains: cmd.ExcludeDomains,
	}

	result, err := client.FindSimilar(ctx, req)
	if err != nil {
		return err
	}

	if outfmt.IsJSON(ctx) {
		return outfmt.WriteJSON(os.Stdout, result)
	}

	if outfmt.IsPlain(ctx) {
		headers := []string{"TITLE", "URL", "SCORE", "DATE", "AUTHOR"}

		var rows [][]string
		for _, r := range result.Results {
			rows = append(rows, []string{r.Title, r.URL, fmt.Sprintf("%.2f", r.Score), r.PublishedDate, r.Author})
		}

		return outfmt.WritePlain(os.Stdout, headers, rows)
	}

	if len(result.Results) == 0 {
		fmt.Fprintln(os.Stderr, "No similar pages found")
		return nil
	}

	fmt.Fprintf(os.Stderr, "Found %d similar pages\n\n", len(result.Results))

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

		fmt.Fprintln(os.Stdout, "---")
	}

	return nil
}
