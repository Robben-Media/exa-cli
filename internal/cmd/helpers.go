package cmd

import (
	"fmt"
	"os"

	"github.com/builtbyrobben/exa-cli/internal/exa"
	"github.com/builtbyrobben/exa-cli/internal/secrets"
)

func getExaClient() (*exa.Client, error) {
	// 1. Check env var
	if key := os.Getenv("EXA_API_KEY"); key != "" {
		return exa.NewClient(key), nil
	}

	// 2. Check keyring
	store, err := secrets.OpenDefault()
	if err != nil {
		return nil, fmt.Errorf("open credential store: %w", err)
	}

	key, err := store.GetAPIKey()
	if err != nil {
		return nil, fmt.Errorf("no credentials found; run: exa-cli auth set-key --stdin")
	}

	return exa.NewClient(key), nil
}
