package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/builtbyrobben/exa-cli/internal/exa"
	"github.com/builtbyrobben/exa-cli/internal/outfmt"
)

type AnswerCmd struct {
	Query string `arg:"" required:"" help:"Question to answer"`
	Num   int    `help:"Number of source results" default:"5"`
}

func (cmd *AnswerCmd) Run(ctx context.Context) error {
	client, err := getExaClient()
	if err != nil {
		return err
	}

	req := exa.AnswerRequest{
		Query:      cmd.Query,
		NumResults: cmd.Num,
		Text:       true,
	}

	result, err := client.Answer(ctx, req)
	if err != nil {
		return err
	}

	if outfmt.IsJSON(ctx) {
		return outfmt.WriteJSON(os.Stdout, result)
	}

	if outfmt.IsPlain(ctx) {
		headers := []string{"ANSWER"}
		rows := [][]string{{result.Answer}}

		return outfmt.WritePlain(os.Stdout, headers, rows)
	}

	fmt.Fprintf(os.Stdout, "Answer:\n%s\n", result.Answer)

	if len(result.Citations) > 0 {
		fmt.Fprintln(os.Stdout, "\nSources:")

		for i, c := range result.Citations {
			if c.Title != "" {
				fmt.Fprintf(os.Stdout, "%d. [%s] - %s\n", i+1, c.Title, c.URL)
			} else {
				fmt.Fprintf(os.Stdout, "%d. %s\n", i+1, c.URL)
			}
		}
	}

	return nil
}
