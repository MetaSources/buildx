package ui

import (
	"fmt"

	"github.com/wagoodman/go-partybus"

	buildxEventParsers "github.com/metasources/buildx/buildx/event/parsers"
)

// handleExit is a UI function for processing the Exit bus event,
// and calling the given function to output the contents.
func handleExit(event partybus.Event) error {
	// show the report to stdout
	fn, err := buildxEventParsers.ParseExit(event)
	if err != nil {
		return fmt.Errorf("bad CatalogerFinished event: %w", err)
	}

	if err := fn(); err != nil {
		return fmt.Errorf("unable to show package catalog report: %w", err)
	}
	return nil
}
